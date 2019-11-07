package conf

import (
	"flag"
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

func conf(confFilePath string) string {
	confFilePath = confFilePath + "config.yaml"
	var absoluteconf string
	var f *flag.Flag
	flag.StringVar(&absoluteconf, "f", confFilePath, "default absolute conf is 'relative path+config.yaml'")
	flag.Parse()
	f = flag.Lookup("f")
	if f.DefValue != f.Value.String() {
		return f.Value.String()
	}
	return ""
}

func InitLoadConfig() (serviceHealthCheckList, serviceInfo, endpointTemplate, kubernetesConf, storageConf, MyConf interface{}) {
	var (
		yamlContent            []uint8
		confFilePath, confName string
	)
	ConfigMap := make(map[interface{}]interface{})
	confName = "config.yaml"
	pwd, err := os.Getwd()
	log.Println("Script Execute Path", pwd)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	executePath := changePath(pwd)
	confFilePath = executePath + "/"
	NewConfFilePath := conf(confFilePath)
	if NewConfFilePath != "" {
		yamlContent, err = ioutil.ReadFile(NewConfFilePath)
		goto AbsoluteConf
	}
	//第一次尝试读取配置
	yamlContent, err = ioutil.ReadFile(confFilePath + confName)
	if err != nil {
		log.Println(err)
		//第二次尝试读取配置
		confFilePath = executePath + "/conf/"
	} else {
		goto AbsoluteConf
	}
	yamlContent, err = ioutil.ReadFile(confFilePath + confName)
	if err != nil {
		log.Println(err)
		//第三次尝试读取配置
		confFilePath = executePath + "/src/"
	} else {
		goto AbsoluteConf
	}
	yamlContent, err = ioutil.ReadFile(confFilePath + confName)
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(yamlContent)
AbsoluteConf:
	if yamlContent == nil {
		panicInfo := "\nCan't Not Get The file Named 'config.yaml' From The Path \n1." +
			executePath + "/\n2." + executePath + "/conf/\n3." + confFilePath + "\nPlease Check it !"
		panic(panicInfo)
	}
	err = yaml.Unmarshal([]byte(yamlContent), &ConfigMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Println(ConfigMap)
	serviceHealthCheckList = ConfigMap["service_healthcheck_list"]
	serviceInfo = ConfigMap["service_info"]
	endpointTemplate = ConfigMap["endpoint_template"]
	kubernetesConf = ConfigMap["kubernetes_conf"]
	storageConf = ConfigMap["storage_conf"]
	MyConf = ConfigMap["my_conf"]
	return
}
