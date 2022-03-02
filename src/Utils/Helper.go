package Utils

import (
	"bytes"
	"reflect"
	"strings"
)

func Join(str ...string) string {
	ret := ""
	for _, s := range str {
		ret += s
	}
	return ret
}

func GoJoin(strs ...string) string {
	return strings.Join(strs, "")
}

func BufferJoin(strs ...string) string {
	var bf bytes.Buffer
	for _, str := range strs {
		bf.WriteString(str)
	}
	return bf.String()
}

func Map2Struct(m map[string]interface{}, u interface{}) {
	v := reflect.ValueOf(u)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		if v.Kind() != reflect.Struct {
			panic("must struct")
		} else {
			findFromMap := func(key string) interface{} {
				for k, v := range m {
					if k == key {
						return v
					}
				}
				return nil
			}
			for i := 0; i < v.NumField(); i++ {
				val := findFromMap(v.Type().Field(i).Name)
				if val != nil && reflect.ValueOf(val).Kind() == v.Field(i).Kind() {
					v.Field(i).Set(reflect.ValueOf(val))
				}
			}
		}
	} else {
		panic("mush ptr")
	}
}
