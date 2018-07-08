package core

import (
	"encoding/json"
	"reflect"
)

func (d *structField) newValue() (v reflect.Value) {
	return reflect.New(d.sourceType)
}

/// error ,should be rewrite
func (d *structField) GetValueFromRootValue(rootValue *reflect.Value) (value *reflect.Value, err error) {
	return getValueByDependLst(d.dependLst, rootValue)
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
