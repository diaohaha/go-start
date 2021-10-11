package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func HttpBuildQuery(params interface{}) (url string, err error) {
	// 请求结构体转url参数
	tmp := []string{}
	paramsTyp := reflect.TypeOf(params)
	if paramsTyp.Kind() != reflect.Struct {
		err = errors.New("TypeError")
		return
	}
	paramsVal := reflect.ValueOf(params)
	for i := 0; i < paramsVal.NumField(); i++ {
		itemKey := paramsTyp.Field(i).Tag.Get("json")
		itemVal := paramsVal.Field(i).Interface()
		typ := paramsVal.Field(i).Type()
		if typ.Kind() == reflect.Map {
			tmp = append(tmp, httpBuildQuery(itemVal, itemKey))
		} else if typ.Kind() == reflect.Slice {
			tmp = append(tmp, httpBuildQuery(itemVal, itemKey))
		} else if typ.Kind() == reflect.Struct {
			tmp = append(tmp, httpBuildQuery(itemVal, itemKey))
		} else if typ.Kind() == reflect.String {
			tmp = append(tmp, fmt.Sprintf("%s=%s", itemKey, itemVal.(string)))
		} else if typ.Kind() == reflect.Int64 {
			tmp = append(tmp, fmt.Sprintf("%s=%d", itemKey, itemVal.(int64)))
		} else if typ.Kind() == reflect.Int {
			tmp = append(tmp, fmt.Sprintf("%s=%d", itemKey, itemVal.(int)))
		} else {
			err = errors.New("TypeError")
			return
		}
	}
	return strings.Join(tmp, "&"), nil
}

func httpBuildQuery(params interface{}, key string) (url string) {
	tmp := []string{}
	typ := reflect.TypeOf(params)
	val := reflect.ValueOf(params)
	if typ.Kind() == reflect.Map || typ.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			itemKey := typ.Field(i).Tag.Get("json")
			itemTyp := val.Field(i).Type()
			itemVal := val.Field(i).Interface()
			if itemTyp.Kind() == reflect.Map {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%s]", key, itemKey)))
			} else if itemTyp.Kind() == reflect.Struct {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%s]", key, itemKey)))
			} else if itemTyp.Kind() == reflect.Array {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%d]", key, i)))
			} else if itemTyp.Kind() == reflect.String {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%s", key, itemKey, itemVal.(string)))
			} else if itemTyp.Kind() == reflect.Int64 {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%d", key, itemKey, itemVal.(int64)))
			} else if itemTyp.Kind() == reflect.Int {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%d", key, itemKey, itemVal.(int)))
			}
		}
	} else if typ.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			itemKey := strconv.Itoa(i)
			itemTyp := val.Index(i).Type()
			itemVal := val.Index(i).Interface()
			if itemTyp.Kind() == reflect.Map {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%s]", key, itemKey)))
			} else if itemTyp.Kind() == reflect.Struct {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%s]", key, itemKey)))
			} else if itemTyp.Kind() == reflect.Array {
				tmp = append(tmp, httpBuildQuery(itemVal, fmt.Sprintf("%s[%d]", key, i)))
			} else if itemTyp.Kind() == reflect.String {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%s", key, itemKey, itemVal.(string)))
			} else if itemTyp.Kind() == reflect.Int64 {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%s", key, itemKey, itemVal.(int64)))
			} else if itemTyp.Kind() == reflect.Int {
				tmp = append(tmp, fmt.Sprintf("%s[%s]=%d", key, itemKey, itemVal.(int)))
			}
		}
	} else {
		return ""
	}
	return strings.Join(tmp, "&")
}

type TestReq struct {
	Name     string   `json:"name"`
	Comments []string `json:"comments"`
	Imgs     []Img    `json:"imgs"`
}

type Img struct {
	Src    string `json:"src"`
	Height int    `json:"height"`
}

func testHttpBulidQuery() {
	imgs := []Img{
		{
			Src:    "apple",
			Height: 100,
		}, {
			Src:    "banana",
			Height: 200,
		},
	}
	test := TestReq{
		Name:     "gaoda",
		Comments: []string{"hello", "nihao"},
		Imgs:     imgs,
	}
	a, e := HttpBuildQuery(test)
	fmt.Println(a)
	fmt.Println(e)
}
