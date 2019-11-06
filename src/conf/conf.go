package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func changePath(pwd string) string {
	operating := runtime.GOOS
	if operating == "windows" {
		pwd = strings.Replace(pwd, "\\", "/", -1)
		return pwd
	}
	return pwd
}

func InitLoadConfig() (interface{}, interface{}, interface{}, interface{}, interface{}) {
	var (
		yamlContent  []uint8
		confFilePath string
	)
	pwd, err := os.Getwd()
	log.Println("Script Execute Path", pwd)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	executePath := changePath(pwd)
	//第一次尝试读取配置
	confFilePath = executePath + "/"
	//fmt.Println(executePath)
	//fmt.Println(confFilePath)
	yamlContent, err = ioutil.ReadFile(confFilePath + "config.yaml")
	if err != nil {
		log.Println(err)
		//第二次尝试读取配置
		confFilePath = executePath + "/conf/"
	}
	yamlContent, err = ioutil.ReadFile(confFilePath + "config.yaml")
	if err != nil {
		log.Println(err)
		//第三次尝试读取配置
		confFilePath = executePath + "/src/"
	}
	yamlContent, err = ioutil.ReadFile(confFilePath + "config.yamls")
	if err != nil {
		log.Println(err)
	}
	if yamlContent == nil {
		panicInfo := "\nCan't Not Get The file Named 'config.yaml' From The Path \n1." +
			executePath + "/\n2." + executePath + "/conf/\n3." + confFilePath + "\nPlease Check it !"
		panic(panicInfo)
	}
	ConfigMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(yamlContent), &ConfigMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(ConfigMap)
	serviceHealthCheckList := ConfigMap["service_healthcheck_list"]
	serviceInfo := ConfigMap["service_info"]
	endpointTemplate := ConfigMap["endpoint_template"]
	kubernetesConf := ConfigMap["kubernetes_conf"]
	storageConf := ConfigMap["storage_conf"]
	return serviceHealthCheckList, serviceInfo, endpointTemplate, kubernetesConf, storageConf
}
