package deepcopy

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueMaps, ok := value.(map[interface{}]interface{}); ok {
		newMaps := make(map[interface{}]interface{})
		for k, v := range valueMaps {
			newMaps[k] = DeepCopy(v)
		}
		return newMaps
	}
	return value
}
