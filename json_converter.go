package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	yaml "gopkg.in/mikefarah/yaml.v2"
)

func jsonToString(context interface{}) (string, error) {
	out, err := json.Marshal(toJSON(context))
	if err != nil {
		return "", fmt.Errorf("error printing yaml as json: %v", err)
	}
	return string(out), nil
}

func toJSON(context interface{}) interface{} {
	switch context.(type) {
	case []interface{}:
		oldArray := context.([]interface{})
		newArray := make([]interface{}, len(oldArray))
		for index, value := range oldArray {
			newArray[index] = toJSON(value)
		}
		return newArray
	case yaml.MapSlice:
		oldMap := context.(yaml.MapSlice)
		newMap := make(map[string]interface{})
		for _, entry := range oldMap {
			if str, ok := entry.Key.(string); ok {
				newMap[str] = toJSON(entry.Value)
			} else if i, ok := entry.Key.(int); ok {
				newMap[strconv.Itoa(i)] = toJSON(entry.Value)
			} else if b, ok := entry.Key.(bool); ok {
				newMap[strconv.FormatBool(b)] = toJSON(entry.Value)
			}
		}
		return newMap
	default:
		return context
	}
}
