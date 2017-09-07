package xodm

import "reflect"

type tagStruct struct {
	fieldName string
	fieldType string
	tagString string
	funcLst   map[string]string
}

func formattag(t reflect.StructField) *tagStruct {
	_tagStruct := new(tagStruct)
	_tagStruct.fieldName = t.Name
	_tagStruct.fieldType = t.Type.Kind().String()
	_tagStruct.tagString = t.Tag.Get(tagName)
	return _tagStruct
}
