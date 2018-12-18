package core

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/x-tool/tool"
)

// func (v *Value) Field() *StructField {
// 	return v.field
// }

// func (v *Value) Kind() Kind {
// 	return v.field.Kind()
// }

// func (v *Value) String() string {
// 	return ValueToString(v)
// }

// func (v *Value) Get() (value *reflect.Value) {
// 	value = v.value
// 	return
// }

// func (v *Value) Addr() (value *reflect.Value) {
// 	if v.value.CanAddr() {
// 		_value := v.value.Addr()
// 		value = &_value
// 	} else {
// 		value = nil
// 	}
// 	return
// }

func ValueToString(value *reflect.Value) (s string) {
	return tool.ReflectValueToString(value)
}

func StringToValue(f *StructField, str string) (r reflect.Value, err error) {
	str = strings.TrimSpace(str)
	switch f.Kind() {
	case Bool:
		var b bool
		if str == "true" || str == "1" {
			b = true
		}
		r = reflect.ValueOf(true)
	case Int:
		_r, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			r = reflect.ValueOf(_r)
		}
	case Byte:
		_r = []byte(str)
		r = reflect.ValueOf(_r)
	case Float:
		_r, err := strconv.ParseFloat(v, 64)
		if err == nil {
			r = reflect.ValueOf(_r)
		}
	case Complex:

	}
	return
	Array
	Map
	String
	Time
	Date
	DateTime
	TimeStamp
	money
	Struct
	Interface
}
