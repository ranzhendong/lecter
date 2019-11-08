package manageyaml

import (
	"deepcopy"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type error interface {
	Error() string
}

func IsNil(i interface{}) bool {
	defer func() {
		recover()
	}()
	vi := reflect.ValueOf(i)
	return vi.IsNil()
}

func changePath(pwd string) string {
	operating := runtime.GOOS
	if operating == "windows" {
		pwd = strings.Replace(pwd, "\\", "/", -1)
		return pwd
	}
	return pwd
}

func readTokenFile() string {
	var TokenFile, linuxTokenFile, winTokenFile string
	operating := runtime.GOOS
	linuxTokenFile = "var/run/secrets/kubernetes.io/serviceaccount/token"
	winTokenFile = "C:\\Users\\Administrator\\go\\my_script\\src\\token\\token"
	if operating == "windows" {
		TokenFile = strings.Replace(winTokenFile, "\\", "/", -1)
	} else {
		TokenFile = linuxTokenFile
	}
	tokenFileContent, err := ioutil.ReadFile(TokenFile)
	if err != nil {
		log.Fatal(err)
	}
	if tokenFileContent == nil {
		panicInfo := "\nCan't Not Get The Token file Name From The \n1." +
			linuxTokenFile + "\n2." + winTokenFile + "\n" + "\nPlease Check it !"
		panic(panicInfo)
	}
	return string(tokenFileContent)
}

func readTokenFileDefault() string {
	var (
		tokenFileContent []uint8
		tokenFilePath    string
	)
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	executePath := changePath(pwd)
	//第一次尝试读取token
	tokenFilePath = executePath + "/conf/token"
	tokenFileContent, err = ioutil.ReadFile(tokenFilePath)
	if err != nil {
		log.Println(err)

	} else {
		goto tokenFile
	}
	//第二次尝试读取token
	tokenFilePath = executePath + "/token"
	tokenFileContent, err = ioutil.ReadFile(tokenFilePath)
	if err != nil {
		log.Println(err)
	}
tokenFile:
	if tokenFileContent == nil {
		panicInfo := "\nCan't Not Get The Token file Name From The Path \n1." +
			executePath + "/\n2." + executePath + "/conf/\n" + "\nPlease Check it !"
		panic(panicInfo)
	}
	return string(tokenFileContent)
}

func YamlServiceInfoMap(serviceInfo interface{}) map[string][]string {
	serviceInfoMap := make(map[string][]string)
	for c, v := range serviceInfo.(map[interface{}]interface{}) {
		a := make([]string, 2)
		k := c.(string)
		name := v.(map[interface{}]interface{})["name"].(string)
		namespace := v.(map[interface{}]interface{})["namespace"].(string)
		a[0] = name
		a[1] = namespace
		serviceInfoMap[k] = a
	}
	return serviceInfoMap
}

func YamlMyConf(MyConf interface{}) (interval, pingTimeout int) {
	interval = MyConf.(map[interface{}]interface{})["interval"].(int)
	pingTimeout = MyConf.(map[interface{}]interface{})["pingtimeout"].(int)
	return
}

func YamlServiceInfo(serviceInfo, name interface{}) (string, string) {
	myService := serviceInfo.(map[interface{}]interface{})[name]
	Name := myService.(map[interface{}]interface{})["name"].(string)
	nameSpace := myService.(map[interface{}]interface{})["namespace"].(string)
	return Name, nameSpace
}

func YamlstorageConf(storageConf interface{}) (IpPort, PassWord string) {
	defer func() {
		err := recover()
		if PassWord == "" {
			PassWord = ""
		}
		if err != nil {
			fmt.Println(err) // runtime error: index out of range
		}
	}()
	Redis := storageConf.(map[interface{}]interface{})["redis"].(map[interface{}]interface{})
	IpPort = Redis["ipport"].(string)
	PassWord = Redis["password"].(string)
	fmt.Println(IpPort)
	return IpPort, PassWord
}

func YamlKubernetesInfo(kubernetesConf interface{}) (url, contentType, endPointApi, tokenFileContent string) {
	var tokenFile string
	defer func() {
		// recover() 捕获panic异常，获得程序执行权。
		_ = recover()
		tokenFileContent = readTokenFile()
		fmt.Println(tokenFileContent)
	}()
	// 异常处理，通过recover获取执行权
	contentType = kubernetesConf.(map[interface{}]interface{})["content_type"].(string)
	endPointApi = kubernetesConf.(map[interface{}]interface{})["endpointapi"].(string)
	url = kubernetesConf.(map[interface{}]interface{})["url"].(string)
	tokenFile = kubernetesConf.(map[interface{}]interface{})["token_file"].(string)
	if tokenFile == "default" {
		tokenFileContent = readTokenFileDefault()
	} else {
		tokenFileContent = readTokenFile()
	}
	fmt.Println(tokenFileContent)
	return url, contentType, endPointApi, tokenFileContent
}

func YamlFactory(name, nameSpace string, Port int, SuccessMapList, endpointTemplate interface{}) string {
	return YamlConverter(name, nameSpace, Port, SuccessMapList, endpointTemplate)
}

func YamlConverter(name, namespace string, Port int, SuccessMapList, EndpointTemplate interface{}) string {
	//map深拷贝
	NewEndpointTemplate := deepcopy.DeepCopy(EndpointTemplate)
	// make sure type is right
	metadata := make(map[interface{}]interface{})
	subsets := make(map[interface{}]interface{})
	//保证数据结构一致，有些数据是嵌套列表
	subsetsArrary := make([]interface{}, 1)
	portArrary := make([]interface{}, 1)
	ports := make(map[interface{}]interface{})
	//对matadata修改
	metadata["name"] = name
	metadata["namespace"] = namespace
	NewEndpointTemplate.(map[interface{}]interface{})["metadata"] = metadata
	//对subsets修改
	protocol := NewEndpointTemplate.(map[interface{}]interface{})["subsets"].([]interface{})[0].(map[interface{}]interface{})["ports"].([]interface{})[0].(map[interface{}]interface{})["protocol"].(string)
	//如果SuccessMapList为空，但是需要提交EndpointTemplateYaml，
	//不过这部分可以优化，将在前面进行空值判断，无需执行后续操作。
	if IsNil(SuccessMapList) == true {
		goto HEADLE
	}
	ports["port"] = Port
	ports["protocol"] = protocol
	portArrary[0] = ports
	subsets["addresses"] = SuccessMapList
	subsets["ports"] = portArrary
	subsetsArrary[0] = subsets
	NewEndpointTemplate.(map[interface{}]interface{})["subsets"] = subsetsArrary
HEADLE:
	// 直接修改内存地址变量，导致出现不可预计错误，因此需要进行变量赋值
	d, err := yaml.Marshal(NewEndpointTemplate)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)
}
