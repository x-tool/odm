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
	String() string
}

func valueToString(v interface{}, t reflect.Type) (s string) {
	switch t.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(v.(bool))
	case reflect.Int:
		s = strconv.Itoa(v.(int))
	case reflect.Int8:
		s = strconv.FormatInt(int64(v.(int8)), 10)
	case reflect.Int16:
		s = strconv.FormatInt(int64(v.(int16)), 10)
	case reflect.Int32:
		s = strconv.FormatInt(int64(v.(int32)), 10)
	case reflect.Int64:
		s = strconv.FormatInt(v.(int64), 10)
	case reflect.Uint:
		s = strconv.FormatUint(uint64(v.(uint)), 10)
	case reflect.Uint8:
		s = strconv.FormatUint(uint64(v.(uint8)), 10)
	case reflect.Uint16:
		s = strconv.FormatUint(uint64(v.(uint16)), 10)
	case reflect.Uint32:
		s = strconv.FormatUint(uint64(v.(uint32)), 10)
	case reflect.Uint64:
		s = strconv.FormatUint(v.(uint64), 10)
	case reflect.Float32:
		s = strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case reflect.Float64:
		s = strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		s = ""
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.String:
		b, err := json.Marshal(v)
		if err != nil {
			s = ""
		} else {
			s = string(b)
		}
	case reflect.Struct:
		pkgPath := t.PkgPath()
		switch pkgPath {
		case "time":
			s = v.(time.Time).String()
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
