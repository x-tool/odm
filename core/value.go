package core

import (
	"encoding/binary"
	"unsafe"
)

// type ValueLst []*Value
// type Value struct {
// 	field    *StructField
// 	value    *reflect.Value
// 	hasValue bool
// 	zero     bool
// }

// func newValue(v interface{}, field *StructField) (o *Value) {
// 	_v := reflect.ValueOf(v)
// 	o = &Value{
// 		field: field,
// 		value: &_v,
// 	}
// 	return o
// }

// func (v *Value) toString() (s string) {
// 	return tool.ReflectValueToString(v.value)
// }

// func newValueByReflect(v *reflect.Value, field *StructField) (o *Value) {
// 	o = &Value{
// 		field: field,
// 		value: v,
// 	}
// 	return o
// }

// check Endian
var Endian binary.ByteOrder

func systemEdian() {
	var i int = 0x1
	bs := (*[int(unsafe.Sizeof(0))]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		Endian = binary.LittleEndian
	} else {
		Endian = binary.BigEndian
	}
}

func init() {
	systemEdian()
}
