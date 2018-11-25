package core

import (
	"encoding/json"
	"reflect"
)

func (d *structField) newValue() (v reflect.Value) {
	return reflect.New(d.sourceType)
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
