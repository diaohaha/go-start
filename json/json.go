package main

import (
	"encoding/json"
	"fmt"
)

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
	}
	return value
}

func ModifyJsonField(jsonBytes []byte, fieldTree []string, data interface{}) (after []byte, err error) {
	// 更改序列化的json某个字符，不丢失总体结构
	var m map[string]interface{}
	err = json.Unmarshal(jsonBytes, &m)
	if err != nil {
		return jsonBytes, err
	}
	n, err := modifyJsonField(&m, fieldTree, data)
	if err != nil {
		return jsonBytes, err
	}
	after, err = json.Marshal(n)
	if err != nil {
		return jsonBytes, err
	}
	return after, nil
}

func modifyJsonField(before *map[string]interface{}, fieldTree []string, data interface{}) (after *map[string]interface{}, err error) {
	var JsonMap map[string]interface{} = DeepCopy(*before).(map[string]interface{})
	if len(fieldTree) == 1 {
		(JsonMap)[fieldTree[0]] = data
		return &JsonMap, nil
	} else {
		SubJsonMap := new(map[string]interface{})
		bytes, ierr := json.Marshal((JsonMap)[fieldTree[0]])
		if err != nil {
			return &JsonMap, ierr
		}
		err = json.Unmarshal(bytes, &SubJsonMap)
		if err != nil {
			return &JsonMap, err
		}
		(JsonMap)[fieldTree[0]], err = modifyJsonField(SubJsonMap, fieldTree[1:], data)
		if err != nil {
			return &JsonMap, err
		}
		return &JsonMap, nil
	}
}

func main() {
	a := "{\"ext\":{\"b\":{\"bd\":\"51\",\"be\":\"父\"},\"head\":\"https\",\"headcolor\":\"#48842D\",\"tv\":{\"tid\":\"577\",\"xcey\":\"hPrRT\",\"xcxm\":\"108000\",\"xcath\":\"/pagvideo\",\"xcxry\":\"h5aldbd\"}},\"sou\":\"knge\",\"type\":\"word\"}"
	after, _ := ModifyJsonField([]byte(a), []string{"sou"}, "ddd")
	// after, _ := ModifyJsonField([]byte(a), []string{"ext", "bk", "bk_id"}, "123")
	fmt.Println(string(after))
}
