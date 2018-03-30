package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type doc struct {
	name         string
	col          *Col
	fields       docFieldLst
	sourceType   *reflect.Type
	mode         *docField
	fieldTagMap  map[string]*docField
	fieldNameMap map[string]docFieldLst
	rootFields   docFieldLst
}
type docLst []*doc

func (d *doc) getChildFields(i *docField) (r docFieldLst) {
	return i.childLst
}

func (d *doc) getChildFieldByName(i *docField, s string) (r *docField) {
	for _, v := range i.childLst {
		if v.Name() == s {
			r = v
			break
		}
	}
	return
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

func (d *doc) GetRootFields() docFieldLst {
	return d.rootFields
}
func (d *doc) getStructRootFields() (lst docFieldLst) {
	for _, v := range d.fields {
		if v.parent == nil {
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
	fields := newDocFields(_doc, docSourceT)
	_doc.fields = *fields
	_doc.fieldTagMap = _doc.makeDocFieldTagMap()
	_doc.fieldNameMap = _doc.makeDocFieldNameMap()
	_doc.rootFields = _doc.makerootFieldNameMap()
	_doc.mode = _doc.findDocMode()
	return
}

// var addFieldsLock sync.WaitGroup

func newDocFields(d *doc, docSourceT reflect.Type) *docFieldLst {
	var lst docFieldLst
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			// addFieldsLock.Add(1)
			// go newDocField(d, &lst, &field, nil)
			newDocField(d, &lst, &field, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}
	// addFieldsLock.Wait()
	return &lst
}

func (d *doc) makeDocFieldTagMap() (m map[string]*docField) {
	_d := d.fields
	for _, v := range _d {
		tagPtr := v.tag.ptr
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func (d *doc) makeDocFieldNameMap() map[string]docFieldLst {
	_d := d.fields
	var _map = make(map[string]docFieldLst)
	for _, v := range _d {
		name := v.Name()
		// new m[name]
		if _, ok := _map[name]; !ok {
			var temp docFieldLst
			_map[name] = temp
		}
		_map[name] = append(_map[name], v)
	}
	return _map
}

func (d *doc) makerootFieldNameMap() (lst []*docField) {
	_d := d.fields
	for _, v := range _d {
		if v.extendParent == nil && v.IsExtend() == false {
			lst = append(lst, v)
		}
	}
	return
}
