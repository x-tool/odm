package core

import (
	"reflect"
)

const (
	Invalid Kind = iota
	Bool
	Int
	Uint
	Byte
	Float
	Complex
	Array
	Map
	String
	Time
	Date
	DateTime
	TimeStamp
	Struct
)

var typeStringMap = map[Kind]string{
	Bool:      "bool",
	Int:       "int",
	Uint:      "unit",
	Byte:      "byte",
	Float:     "float",
	Complex:   "complex",
	Array:     "array",
	Map:       "map",
	String:    "string",
	Time:      "time",
	Date:      "date",
	DateTime:  "datetime",
	TimeStamp: "timestamp",
	Struct:    "struct",
}

type Kind uint

func (k Kind) String() (s string) {
	return typeStringMap[k]
}

func (k Kind) isGroupType() (b bool) {
	switch k {
	case Array:
		fallthrough
	case Map:
		fallthrough
	case Struct:
		b = true
	}
	return
}

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

func reflectToKind(r *reflect.Type) (k Kind) {
	_r := *r
	rKind := _r.Kind()
	switch rKind {
	case reflect.Bool:
		k = Bool
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		k = Int
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		k = Uint
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		k = Float
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		k = Complex
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
