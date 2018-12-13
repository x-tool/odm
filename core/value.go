package core

import (
	"reflect"

	"github.com/x-tool/tool"
)

type ValueLst []*Value
type Value struct {
	field    *StructField
	value    *reflect.Value
	hasValue bool
	zero     bool
}

func newValue(v interface{}, field *StructField) (o *Value) {
	_v := reflect.ValueOf(v)
	o = &Value{
		field: field,
		value: &_v,
	}
	return o
}

func (v *Value) toString() (s string) {
	return tool.ReflectValueToString(v.value)
}

func newValueByReflect(v *reflect.Value, field *StructField) (o *Value) {
	o = &Value{
		field: field,
		value: v,
	}
	return o
}
