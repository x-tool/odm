package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

// type ValueLst []*Value
// type Value struct {
// 	field    *structField
// 	value    *reflect.Value
// 	hasValue bool
// 	zero     bool
// }

// func newValue(v interface{}, field *structField) (o *Value) {
// 	_v := reflect.ValueOf(v)
// 	o = &Value{
// 		field: field,
// 		value: &_v,
// 	}
// 	return o
// }
// func newValueByReflect(v *reflect.Value, field *structField) (o *Value) {
// 	o = &Value{
// 		field: field,
// 		value: v,
// 	}
// 	return o
// }

func ValueToString(value *reflect.Value) (s string) {
	_value := *value
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
