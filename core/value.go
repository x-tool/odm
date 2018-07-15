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
<<<<<<< HEAD
	_value := *value
	valueType := _value.Type()
=======
	v := *value
	valueType := v.Type()
>>>>>>> 314a364ec47897e822bd43b5e555ea8d557c22ef
	switch valueType.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(v.Bool())
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		s = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		s = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		s = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		s = ""
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			s = ""
		} else {
			s = string(b)
		}
	case reflect.String:
		s = v.String()
	case reflect.Struct:
		pkgPath := valueType.PkgPath()
		switch pkgPath {
		case "time":
			s = v.Interface().(time.Time).String()
		default:
			b, err := json.Marshal(v)
			if err != nil {
				s = ""
			} else {
				s = string(b)
			}
		}

	}
	return
}
