package core

import (
	"reflect"
	"strings"
)

const (
	splitStructStr = ":"
)

// var splitStructNameToFieldPath = []string{
// 	".",
// 	tag_Tag
// }

func (d *doc) findDocModeField() (field *structField) {
	for _, v := range d.getExtendFields() {
		_value := reflect.New(v.sourceType)
		_, ok := _value.Interface().(DocMode)
		if ok {
			field = v
			break
		}
	}
	return
}

func (d *doc) getDocMode() *structField {
	return d.mode
}

// "@tag"
// "path"
// "path:@tag"
// "path:path"
func (d *doc) getFieldByAllPath(s string, rootValue *reflect.Value) (value *Value, err error) {
	var rawValue = rootValue
	structLst := strings.Split(s, splitStructStr)
	field, err := d.getFieldByString(structLst[0])
	if err != nil {
		return
	}
	rawValue, err = field.GetValueFromRootValue(rawValue)
	if err != nil {
		return
	}
	for _, v := range structLst[1:] {
		// now just should find "@", if add more sign to split struct in the fultrue, should modify
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
		targetStruct, err = d.col.database.getStructByName(structName)
		if err != nil {
			return
		}
		field, err = targetStruct.getFieldByString(fieldRoute)
		if err != nil {
			return
		}
		rawValue, err = field.GetValueFromRootValue(rawValue)
		if err != nil {
			return
		}
	}
	value = rawValue
	return
}
