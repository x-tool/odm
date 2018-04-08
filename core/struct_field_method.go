package core

import (
	"encoding/json"
	"reflect"
)

func (d *structField) newValue() (v reflect.Value) {
	return reflect.New(d.selfType)
}

/// error
func (d *structField) GetValueFromRootValue(rootValue *reflect.Value) *reflect.Value {
	_r := *rootValue
	if len(d.dependLst) != 0 {
		for _, v := range d.dependLst {
			// if dependLst parent is extend, should get field from grandparent
			if v.kind == Struct {
				_r = _r.FieldByName(v.Name())
			} else {
				break // !!!!!! mark wait modify !!!!!!!
			}
		}
		_r = _r.FieldByName(d.Name())
	} else {
		_r = _r.FieldByName(d.Name())
	}

	return &_r
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
