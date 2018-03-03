package core

import (
	"reflect"
)

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	Complex64
	Complex128
	Array
	Map
	Slice
	String
	Time
	Struct
)

func mapTypeToValue(b interface{}, v *reflect.Value) {
	// realV := reflect.Indirect(*v)
	switch b.(type) {
	case string:
		v.SetString(b.(string))
	case int:
		v.SetInt(int64(b.(int)))
	case int32:
		v.SetInt(int64(b.(int32)))
	case int64:
		v.SetInt(b.(int64))
	case bool:
		v.SetBool(b.(bool))
	}

}

func reflectToType(r *reflect.Type) (k Kind) {
	_r := *r
	rKind := _r.Kind()
	switch r {
	case reflect.Bool:
		k = Bool
	case reflect.Int:
		k = Int
	case reflect.Int8:
		k = Int8
	case reflect.Int16:
		k = Int16
	case reflect.Uint:
		k = Uint
	case reflect.Float32:
		k = Float32
	case reflect.Float64:
		k = Float64
	case reflect.Complex64:
		k = Complex64
	case reflect.Complex128:
		k = Complex128
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		k = Array
	case reflect.String:
		k = String
	case reflect.Struct:
		pkgPath := _r.PkgPath()
		if pkgPath == "time" {
			k = Time
		} else {
			k = Struct
		}
	}
	return
}

func isGroupType(k Kind) (b bool) {
	switch k {
	case Array:
		fallthrough
	case Slice:
		fallthrough
	case Map:
		fallthrough
	case Struct:
		b = true
	}
	return
}
