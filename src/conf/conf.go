package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func InitLoadConfig() (interface{}, interface{}, interface{}, interface{}, interface{}) { // 初始化数据
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	projectPath := strings.Replace(pwd, "\\", "/", -1)
	confFile := filepath.Join(projectPath+"/src", "config.yaml")
	yamlContent, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Fatal(err)
	}
	ConfigMap := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(yamlContent), &ConfigMap)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Println(ConfigMap)
	serviceHealthCheckList := ConfigMap["service_healthcheck_list"]
	serviceInfo := ConfigMap["service_info"]
	endpointTemplate := ConfigMap["endpoint_template"]
	kubernetesConf := ConfigMap["kubernetes_conf"]
	storageConf := ConfigMap["storage_conf"]
	return serviceHealthCheckList, serviceInfo, endpointTemplate, kubernetesConf, storageConf
}
