package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

type HandleValue struct {
	v    interface{}
	zero bool
	doc  *doc
}

func newValue(v interface{}, doc *doc) (o *HandleValue) {
	o = &HandleValue{
		v:   v,
		doc: doc,
	}
	return o
}

type Value interface {
	Name() string
	Value() interface{}
	Kind() Kind
}

func ValueToString(value *reflect.Value) (s string) {
	v := *value
	valueType := v.Type()
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
		s = value.String()
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
