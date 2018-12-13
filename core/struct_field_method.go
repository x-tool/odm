package core

import (
	"encoding/json"
	"reflect"
)

func (d *StructField) newValue() (v reflect.Value) {
	return reflect.New(d.sourceType)
}

func (d *StructField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
