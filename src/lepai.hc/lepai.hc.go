package lepai_hc

import (
	"encoding/json"
	"fmt"
	"lepai.net"
	"reflect"
	"strconv"
)

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
	//fmt.Println("SuccessMapList, FailInfo", SuccessMapList, FailMapList)
	return SuccessMapList, FailMapList, Port
}

func CompareSuccessListWithRedisList(SuccessMapList interface{}, RedisListGet string) (myError int) {
	//fmt.Printf("SuccessMapList %v, RedisListGet %v", SuccessMapList, RedisListGet)
	if RedisListGet == "" {
		fmt.Println("RedisListGet CAN'T GET ")
		return 7
	}
	var RedisListGetMap map[string]interface{}
	if err := json.Unmarshal([]byte(RedisListGet), &RedisListGetMap); err != nil {
		fmt.Println("bodyContentByte json change err: ", err)
		return 7
	}
	//fmt.Println("==============json str 转map=======================")
	//fmt.Printf("eeeeee %T", RedisListGetMap)
	//fmt.Println(RedisListGetMap["subsets"].([]interface{})[0])
	subsets := RedisListGetMap["subsets"].([]interface{})[0]
	addresses := subsets.(map[string]interface{})["addresses"]
	if reflect.DeepEqual(addresses, SuccessMapList) {
		fmt.Println("Addresses Is SuccessMapList", addresses, "==", SuccessMapList)
		return
	}
	fmt.Println("Addresses Not SuccessMapList", addresses, "!=", SuccessMapList)
	return 1
}
