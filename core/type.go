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
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
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

func formatTypeToString(k Kind) string {
	// switch k {

	// case Bool:

	// Int
	// Int8
	// Int16
	// Int32
	// Int64
	// Uint
	// Uint8
	// Uint16
	// Uint32
	// Uint64
	// Uintptr
	// Float32
	// Float64
	// Complex64
	// Complex128
	// Array
	// Chan
	// Func
	// Interface
	// Map
	// Ptr
	// Slice
	// String
	// }
	return ""
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
