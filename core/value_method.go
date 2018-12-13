package core

import (
	"reflect"

	"github.com/x-tool/tool"
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
	return tool.ReflectValueToString(value)
}
