package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func (d *structField) newValue() (v reflect.Value) {
	return reflect.New(d.sourceType)
}

/// error ,should be rewrite
func (d *structField) GetValueFromRootValue(rootValue *reflect.Value) (value *Value, err error) {
	rawValue := *rootValue
	for _, v := range d.dependLst {
		if v.kind == Struct {
			rawValue = rawValue.FieldByName(v.Name())
		} else {
			// can't get field in slice or map, but should be
			err = errors.New(fmt.Sprintf("Can't get Values in struct %d, because fieldName %d parent type is %d", d.name, v.name, v.kind.String()))
			break
		}
	}
	rawValue = rawValue.FieldByName(d.Name())
	value = newValueByReflect(&rawValue, d)
	return
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
