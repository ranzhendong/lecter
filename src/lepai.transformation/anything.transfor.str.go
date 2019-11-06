package lepai_transfromation

func Any2Str(any interface{}) string {
	switch any.(type) { //多选语句switch
	case string:
		return any.(interface{}).(string)
	case interface{}:
		return any.(interface{}).(string)
	}
	return "nil"
}

//func Any2Map(any interface{}) interface{} {
//	switch any.(type) { //多选语句switch
//	case string:
//		ConfigMap := make(map[interface{}]interface{})
//		err = yaml.Unmarshal([]byte(any), &ConfigMap)
//		return any.(interface{}).(string)
//	case interface{}:
//		return any.(interface{}).(string)
//	}
//	return "nil"
//}
