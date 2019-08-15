package core

import (
	"encoding/binary"
	"reflect"
	"unsafe"
)

// struct path and name in string
func allName(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	s := t.PkgPath() + t.Name()
	return s
}

//////////////// check user system Endian
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
