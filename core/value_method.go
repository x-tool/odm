package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	_value := *value
	if value.Kind() == reflect.Invalid {
		return ""
	}
	valueType := _value.Type()
	switch valueType.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(value.Bool())
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		s = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		s = strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		s = strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		s = ""
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		b, err := json.Marshal(value.Interface())
		if err != nil {
			s = ""
		} else {
			s = string(b)
		}
	case reflect.String:
		s = value.String()
	case reflect.Struct:
		pkgPath := valueType.PkgPath()
		switch pkgPath {
		case "time":
			s = value.Interface().(time.Time).String()
		default:
			b, err := json.Marshal(value)
			if err != nil {
				s = ""
			} else {
				s = string(b)
			}
		}

	}
	return
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
