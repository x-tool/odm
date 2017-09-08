package xodm

import "reflect"

type docfield struct {
	fieldName string
	fieldType string
	tagString string
	funcLst   map[string]string
}

func formattag(t reflect.StructField) *docfield {
	_docfield := new(docfield)
	_docfield.fieldName = t.Name
	_docfield.fieldType = t.Type.Kind().String()
	_docfield.tagString = t.Tag.Get(tagName)
	return _docfield
}

type doc struct {
	colName    string
	handleType string
}

func NewDoc(c ColInterface) {

}
