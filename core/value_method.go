package core

import "reflect"

func (v *Value) Field() *structField {
	return v.field
}

func (v *Value) Kind() Kind {
	return v.field.Kind()
}

func (v *Value) String() string {
	return ValueToString(v)
}

func (v *Value) Get() (value *reflect.Value) {
	value = v.value
	return
}

func (v *Value) Addr() (value *reflect.Value) {
	if v.value.CanAddr() {
		_value := v.value.Addr()
		value = &_value
	} else {
		value = nil
	}
	return
}
