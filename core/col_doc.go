package core

import (
	"errors"
	"reflect"
	"sync"

	"github.com/x-tool/tool"
)

type doc struct {
	name         string
	col          *Col
	fields       docFieldLst
	sourceType   *reflect.Type
	mode         colMode
	fieldTagMap  map[string]*docField
	fieldNameMap map[string]docFieldLst
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

func (d *doc) getFieldByName(name string) (o docFieldLst) {
	return d.fieldNameMap[name]
}

func (d *doc) getFieldByTag(tag string) (o *docField) {
	return d.fieldTagMap[tag]
}

func (d *doc) GetRootFields() (lst docFieldLst) {
	for _, v := range d.fields {
		if v.extendParent == nil && v.IsExtend() == false {
			lst = append(lst, v)
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
	fields := newDocFields(_doc, &docSourceT)
	_doc.fields = *fields
	_doc.mode = checkDocMode(docSourceT)
	_doc.fieldTagMap = makeDocFieldTagMap(&_doc.fields)
	_doc.fieldNameMap = makeDocFieldNameMap(&_doc.fields)
	return
}

var addFieldsLock sync.WaitGroup

func newDocFields(d *doc, docSourceTPtr *reflect.Type) *docFieldLst {
	var lst docFieldLst
	docSourceT := *docSourceTPtr
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			addFieldsLock.Add(1)
			go newDocField(d, &lst, &field, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}
	addFieldsLock.Wait()
	return &lst
}

func makeDocFieldTagMap(d *docFieldLst) (m map[string]*docField) {
	_d := *d
	for _, v := range _d {
		tagPtr := v.tag.ptr
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func makeDocFieldNameMap(d *docFieldLst) (m map[string]docFieldLst) {
	_d := *d
	for _, v := range _d {
		name := v.GetName()
		if _, ok := m[name]; ok {
			m[name] = append(m[name], v)
		} else {
			var temp docFieldLst
			m[name] = append(temp, v)
		}
	}
	return m
}

func checkDocMode(docSourceT reflect.Type) (m colMode) {
	return
}

type docLst []*doc
