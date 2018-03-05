package core

import (
	"errors"
	"reflect"
	"sync"

	"github.com/x-tool/tool"
)

type doc struct {
	name       string
	col        *Col
	fields     docFieldLst
	sourceType *reflect.Type
	mode       colMode
}

func (d *doc) getChildFields(i *docField) (r docFieldLst) {
	return i.childLst
}

func (d *doc) getFieldById(id int) (o *docField) {
	for _, v := range d.fields {
		if v.GetID() == id {
			o = v
			return o
		}
	}
	return
}

func NewDoc(c *Col, i interface{}) (_doc *doc) {

	// append doc.fields
	_docSourceT := reflect.TypeOf(i)
	docSourceT := _docSourceT.Elem()
	_doc = &doc{
		name:       docSourceT.Name(),
		col:        c,
		sourceType: &docSourceT,
	}
	_doc.fields = newDocFields(_doc, &docSourceT)
	_doc.mode = checkDocMode(&docSourceT)

	return
}

var addFieldsLock sync.WaitGroup

func newDocFields(d *doc, docSourceTPtr *reflect.Type) (lst docFieldLst) {
	docSourceT := *docSourceTPtr
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			addFieldsLock.Add(1)
			go newDocField(lst, &field, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}
	addFieldsLock.Wait()
	return
}

func checkDocMode(docSourceT *reflect.Type) (m colMode) {
	return
}

type docLst []*doc
