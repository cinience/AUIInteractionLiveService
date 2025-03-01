package utils

func DeepCopy(value interface{}, changeKeyStyleFunc func(k string) string) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newKey := changeKeyStyleFunc(k)
			newMap[newKey] = DeepCopy(v, changeKeyStyleFunc)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v, changeKeyStyleFunc)
		}

		return newSlice
	}

	return value
}
