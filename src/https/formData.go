/*
tester

This file was automatically generated for Stamplay by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package https

import (
    "encoding/json"
    "fmt"
    "log"
    "net/url"
    "reflect"
)


func structToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func formEncodeMap(name string, value interface{}, keys *[]map[string]interface{}) []map[string]interface{} {
	if keys == nil {
		keys = &[]map[string]interface{}{}
	}

	if value == nil {
		return append(*keys, make(map[string]interface{}))
	} else if reflect.TypeOf(value).Kind() == reflect.Struct ||
		reflect.TypeOf(value).Kind() == reflect.Ptr {
		structMap, err := structToMap(value)
		if err != nil {
			log.Fatalln(err)
		}
		for k, v := range structMap {
			var fullName string = k
			if name != "" {
				fullName = name + "[" + k + "]"
			}
			formEncodeMap(fullName, v, keys)
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Map {
		for k, v := range value.(map[string]interface{}) {
			var fullName string = k
			if name != "" {
				fullName = name + "[" + k + "]"
			}
			formEncodeMap(fullName, v, keys)
		}
	} else if reflect.TypeOf(value).Kind() == reflect.Slice {
		if reflect.TypeOf(value).Elem().Kind() == reflect.Interface {
			for num, val := range value.([]interface{}) {
				fullName := name + "[" + fmt.Sprintf("%v", num) + "]"
				formEncodeMap(fullName, val, keys)
			}
		} else {
			for num, val := range value.([]string) {
				fullName := name + "[" + fmt.Sprintf("%v", num) + "]"
				formEncodeMap(fullName, val, keys)
			}
		}
	} else {
		*keys = append(*keys, map[string]interface{}{name: fmt.Sprintf("%v", value)})
	}

	return *keys
}

func PrepareFormFields(key string, value interface{}, form url.Values) url.Values {
	if form == nil {
		form = url.Values{}
	}

	switch x := value.(type) {
	case []string:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []int:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []int16:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []int32:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []int64:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []bool:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []float32:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	case []float64:
		for _, val := range x {
			form.Add(key, fmt.Sprintf("%v", val))
		}
	default:
		//form.Add(key, fmt.Sprintf("%v", x))
		k := formEncodeMap(key, value, nil)

		for num := range k {
			for key, val := range k[num] {
				form.Add(key, fmt.Sprintf("%v", val))
			}
		}
	}

	return form
}