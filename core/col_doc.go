package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type doc struct {
	name       string
	col        *col
	fields     docFieldLst
	sourceType *reflect.Type
	mode       colMode
}

func (d *doc) getChildFields(i *docField) (r docFieldLst) {
	id := i.GetId()
	for _, v := range d.fields {
		if v.pid == id {
			r = append(r, v)
		}
	}
	return
}

func (d *doc) getFieldById(id int) (o *docField) {
	for _, v := range d.fields {
		if v.GetId() == id {
			o = v
			return o
		}
	}
	return
}

func NewDoc(c *col, i interface{}) *doc {

	// append doc.fields
	docSourceT := reflect.TypeOf(i)
	_doc := &doc{
		name:       docSourceT.Name(),
		col:        c,
		sourceType: &docSourceT,
		fields:     newDocFields(&docSourceT),
		mode:       checkDocMode(&docSourceT),
	}

	return _doc
}

func newDocFields(d *doc, docSourceTPtr *reflect.Type) (lst docFieldLst) {
	docSourceT := *docSourceTPtr
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			newDocField(lst, &field, nil, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}
	return
}

func checkDocMode(docSourceT *reflect.Type) (m colMode) {
	return
}

type docLst []*doc
