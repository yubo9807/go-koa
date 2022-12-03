package object

import (
	"encoding/json"
	"reflect"
)

// 深度克隆，可以克隆任意数据类型
func DeepClone(src interface{}) (interface{}, error) {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		dst := reflect.New(typ).Elem()
		b, _ := json.Marshal(src)
		if err := json.Unmarshal(b, dst.Addr().Interface()); err != nil {
			return nil, err
		}
		return dst.Addr().Interface(), nil
	} else {
		dst := reflect.New(typ).Elem()
		b, _ := json.Marshal(src)
		if err := json.Unmarshal(b, dst.Addr().Interface()); err != nil {
			return nil, err
		}
		return dst.Interface(), nil
	}
}
