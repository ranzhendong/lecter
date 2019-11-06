package lepai_hc

import (
	"encoding/json"
	"lepai.net"
	"log"
	"reflect"
	"strconv"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func HealthCheck(serviceList interface{}) (interface{}, map[string]string, int) {
	var (
		SuccessMapList []interface{}
		Port           int
	)
	FailMapList := make(map[string]string)
	Port = serviceList.([]interface{})[0].(map[interface{}]interface{})["port"].(int)
	ipList := serviceList.([]interface{})[1].(map[interface{}]interface{})["iplist"].([]interface{})
	for _, v := range ipList {
		ip := v.(map[interface{}]interface{})["ip"].(string)
		name := v.(map[interface{}]interface{})["name"].(string)
		IPPort := ip + ":" + strconv.Itoa(Port)
		if err := lepai_net.PingIPPort(IPPort); err != true {
			FailMapList[name] = ip
			continue
		}
		// 下面的方式会导致追加的新数组将原来的数值覆盖，因此需要下面的方式来进行追加
		//SuccessMapList = append(SuccessMapList, IpMap)
		SuccessMapList = append(SuccessMapList, map[string]interface{}{"ip": ip})
	}
	return SuccessMapList, FailMapList, Port
}

func CompareSuccessListWithRedisList(name, SuccessMapList interface{}, RedisListGet string) (myError int) {
	//fmt.Printf("SuccessMapList %v, RedisListGet %v", SuccessMapList, RedisListGet)
	if RedisListGet == "" {
		log.Println("RedisListGet CAN'T GET ")
		return 7
	}
	var RedisListGetMap map[string]interface{}
	if err := json.Unmarshal([]byte(RedisListGet), &RedisListGetMap); err != nil {
		log.Println("bodyContentByte json change err: ", err)
		return 7
	}
	subsets := RedisListGetMap["subsets"].([]interface{})[0]
	addresses := subsets.(map[string]interface{})["addresses"]
	if reflect.DeepEqual(addresses, SuccessMapList) {
		log.Printf("{Name:%v} {Addresses:%v == SuccessMapList: %v}\n", name, addresses, SuccessMapList)
		return
	}
	log.Printf("{Name:%v} {Addresses:%v != SuccessMapList: %v}\n", name, addresses, SuccessMapList)
	return 1
}
