package main

import (
	"conf"
	"lepai.hc"
	"lepai.k8sapi"
	"lepai.storage"
	"lepai.transformation"
	"lepai.yaml"
	"log"
	"reflect"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

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

func secondLoop(pingTimeout int, ipPort, passWord, url, tokenFile, contentType, endPointApi string, endpointTemplate, serviceHealthCheckList, serviceInfo interface{}) {
	for name, serviceList := range serviceHealthCheckList.(map[interface{}]interface{}) {
		SuccessMapList, _, Port := lepai_hc.HealthCheck(serviceList, pingTimeout)
		if IsNil(SuccessMapList) == true {
			log.Printf("SuccessMapList Of HealthCheck Name %v is %v !\n", name, SuccessMapList)
			continue
		}
		RedisListGet := lepai_storage.RedisGet(ipPort, passWord, lepai_transfromation.Any2Str(name))
		err := lepai_hc.CompareSuccessListWithRedisList(name, SuccessMapList, RedisListGet)
		if err == 0 || err == 7 {
			continue
		}
		name, nameSpace := lepai_yaml.YamlServiceInfo(serviceInfo, name)
		yamlConverter := lepai_yaml.YamlFactory(name, nameSpace, Port, SuccessMapList, endpointTemplate)
		_, bodyContent := lepai_k8sapi.APIServerPut(url, name, nameSpace, endPointApi, tokenFile, contentType, yamlConverter)
		lepai_storage.RedisSet(ipPort, passWord, name, bodyContent)
	}
}

func main() {
	//conf.InitLoadConfig()
	// read yaml conf from file
	serviceHealthCheckList, serviceInfo, endpointTemplate, kubernetesConf, storageConf, MyConf := conf.InitLoadConfig()
	//_, _, _, _, _, MyConf := conf.InitLoadConfig()
	// get my conf info
	interval, pingTimeout := lepai_yaml.YamlMyConf(MyConf)
	// get k8sapiserver info
	url, contentType, endPointApi, tokenFile := lepai_yaml.YamlKubernetesInfo(kubernetesConf)
	//  change serviceInfo's structure to serviceInfoMap
	// get redis info
	ipPort, passWord := lepai_yaml.YamlstorageConf(storageConf)
	// 第一次for循环，将数据从apiserver拉取，存入到redis
	firstLoop(ipPort, passWord, url, tokenFile, endPointApi, serviceHealthCheckList, serviceInfo)
	// 第二次循环，将存入redis的数据和已有数据进行比对，将合格结果写入到APIServer
	// 定时任务为什么选择sleep而不是tick，是因为我们需要将前一次的结果作为下一次运行的参考依据，因此不建议使用tick
	//tick其实是每次固定时间执行，但是执行命令的时间实际上不关系的，只会在下一时刻再次触发执行。
	//经过测试之后，发现通过go协程来实现，可以快速控制语句执行。
	//t := time.NewTicker(time.Duration(interval) * time.Millisecond)
	//for {
	//	select {
	//	case <-t.C:
	//		log.Println("t1定时器")
	//		time.Sleep(time.Duration(interval) * time.Millisecond)
	//		secondLoop(ipPort, passWord, url, tokenFile, contentType, endPointApi, endpointTemplate, serviceHealthCheckList, serviceInfo)
	//		//t.Stop()
	//	}
	//}
	var count int
	count = 1
	for {
		log.Println("检查次数:[", count, "]")
		count = count + 1
		time.Sleep(time.Duration(interval) * time.Millisecond)
		go secondLoop(pingTimeout, ipPort, passWord, url, tokenFile, contentType, endPointApi, endpointTemplate, serviceHealthCheckList, serviceInfo)
	}
}
