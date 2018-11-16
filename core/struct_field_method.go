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
	return getValueFromDependLst(d.dependLst, rootValue)
}

//****** can't get field in slice or map
func getValueFromDependLst(dLst dependLst, rootValue *reflect.Value) (value *reflect.Value, err error) {
	_value := *rootValue
	for _, v := range dLst {
		if v.kind == Struct {
			_value = _value.FieldByName(v.Name())
		} else {
			// can't get field in slice or map
			// err = errors.New(fmt.Sprintf("Can't get Values in struct %d, because fieldName %d parent type is %d", d.name, v.name, v.kind.String()))
			break
		}
	}
	value = &_value
	return
}

func (d *structField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}
