package core

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// "@tag"
// "path"
// "path:structName@tag"
// "path:structName.path"
func getDependLstByAllPath(d Database, o *odmStruct, s string) (dLst dependLst, err error) {
	var rawValue = rootValue
	structLst := strings.Split(s, splitStructStr)
	// firstPath, without structname
	field, err := o.getFieldByString(structLst[0])
	if err != nil {
		return
	}
	dLst = field.dependLst
	// orther struct path
	lstLen := len(structLst)
	for i, v := range structLst[1:] {
		// split name and path
		// now just need find "@", if add more sign to split struct in the fultrue, should modify use splitStructNameToFieldPath
		var sign string
		var targetStruct *odmStruct
		index := strings.Index(v, "@")
		if index == -1 {
			sign = "."
		} else {
			sign = "@"
		}

		slice := strings.SplitN(v, sign, 2)
		structName := slice[0]
		fieldRoute := slice[1]
		targetStruct, err = d.getStructByName(structName)
		if err != nil {
			return
		}
		field, err = targetStruct.getFieldByString(fieldRoute)
		if err != nil {
			return
		}
		dLst = append(dLst, field.dependLst)
	}
	return
}

//****** can't get field in slice or map
func getValueByDependLst(dLst dependLst, rootValue *reflect.Value) (value *reflect.Value, err error) {
	_value := *rootValue
	for _, v := range dLst {
		if v.kind == Struct {
			_value = _value.FieldByName(v.Name())
		} else {
			// can't get field in slice or map
			err = errors.New(fmt.Sprintf("Can't get Values in struct %d, because fieldName %d parent type is %d", d.name, v.name, v.kind.String()))
			break
		}
	}
	value = &_value
	return
}
