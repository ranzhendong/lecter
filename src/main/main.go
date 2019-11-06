package main

import (
	"conf"
	"fmt"
	"lepai.hc"
	"lepai.k8sapi"
	"lepai.storage"
	"lepai.transformation"
	"lepai.yaml"
	"time"
)

func firstLoop(ipPort, passWord, url, tokenFile, endPointApi string, serviceHealthCheckList, serviceInfo interface{}) {
	//每次重启脚本加入重置命令，将所有已经配置好的service进行删除
	lepai_storage.RedisReset(ipPort, passWord)
	//按照yaml文件配置好的serviceHealthCheckList name进行循环
	for name, _ := range serviceHealthCheckList.(map[interface{}]interface{}) {
		//每次循环获取name对应的name和namespace资源
		name, nameSpace := lepai_yaml.YamlServiceInfo(serviceInfo, name)
		//从APIServer获取资源，如果不存在返回空字符串，但不是nil
		_, bodyContent := lepai_k8sapi.APIServerGet(url, name, nameSpace, endPointApi, tokenFile)
		//只将可以获取到的资源进行redis写入
		if bodyContent != "" {
			lepai_storage.RedisSet(ipPort, passWord, name, bodyContent)
		}
	}
}

func secondLoop(ipPort, passWord, url, tokenFile, contentType, endPointApi string, endpointTemplate, serviceHealthCheckList, serviceInfo interface{}) {
	for name, serviceList := range serviceHealthCheckList.(map[interface{}]interface{}) {
		SuccessMapList, _, Port := lepai_hc.HealthCheck(serviceList)
		RedisListGet := lepai_storage.RedisGet(ipPort, passWord, lepai_transfromation.Any2Str(name))
		err := lepai_hc.CompareSuccessListWithRedisList(SuccessMapList, RedisListGet)
		if err == 0 || err == 7 {
			continue
		}
		fmt.Println("CompareSuccessListWithRedisList: ", err)
		name, nameSpace := lepai_yaml.YamlServiceInfo(serviceInfo, name)
		yamlConverter := lepai_yaml.YamlFactory(name, nameSpace, Port, SuccessMapList, endpointTemplate)
		_, bodyContent := lepai_k8sapi.APIServerPut(url, name, nameSpace, endPointApi, tokenFile, contentType, yamlConverter)
		lepai_storage.RedisSet(ipPort, passWord, name, bodyContent)
	}
}

func main() {
	// read yaml conf from file
	serviceHealthCheckList, serviceInfo, endpointTemplate, kubernetesConf, storageConf := conf.InitLoadConfig()
	// get k8sapiserver info
	url, contentType, endPointApi, tokenFile := lepai_yaml.YamlKubernetesInfo(kubernetesConf)
	//  change serviceInfo's structure to serviceInfoMap
	ipPort, passWord := lepai_yaml.YamlstorageConf(storageConf)
	// 第一次for循环，将数据从apiserver拉取，存入到redis
	firstLoop(ipPort, passWord, url, tokenFile, endPointApi, serviceHealthCheckList, serviceInfo)
	// 第二次循环，将存入redis的数据和已有数据进行比对，将合格结果写入到APIServer
	for {
		time.Sleep(2000 * time.Millisecond)
		secondLoop(ipPort, passWord, url, tokenFile, contentType, endPointApi, endpointTemplate, serviceHealthCheckList, serviceInfo)
	}
}
