package xodm

import "reflect"

type docStruct struct {
	fieldName string
	fieldType string
	tagString string
	funcLst   map[string]string
}

func formattag(t reflect.StructField) *docStruct {
	_docStruct := new(docStruct)
	_docStruct.fieldName = t.Name
	_docStruct.fieldType = t.Type.Kind().String()
	_docStruct.tagString = t.Tag.Get(tagName)
	return _docStruct
}
