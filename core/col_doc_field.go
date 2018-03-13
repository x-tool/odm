package core

import (
	"reflect"
	"sync"
)

type docField struct {
	doc             *doc
	name            string
	selfType        reflect.Type
	kind            Kind
	id              int
	isExtend        bool
	parent          *docField // field golang parent real
	extendParent    *docField // field Handle parent
	childLst        docFieldLst
	dependLst       dependLst
	extendDependLst dependLst
	tag             *odmTag
	funcLst         map[string]string
}

func (d *docField) Name() string {
	return d.name
}

func (d *docField) GetID() int {
	return d.id
}

func (d *docField) Kind() Kind {
	return d.kind
}
func (d *docField) IsExtend() bool {
	return d.isExtend
}

func (d *docField) isSingleType() (b bool) {
	return !d.kind.isGroupType()
}

func (d *docField) isGroupType() (b bool) {
	return d.kind.isGroupType()
}

func newDocField(_doc *doc, d *docFieldLst, t *reflect.StructField, parent *docField) {
	fieldType := *t
	reflectType := fieldType.Type
	tag := fieldType.Tag.Get(tagName)
	kind := reflectToKind(&reflectType)
	isExtend := checkdocFieldisExtend(t)
	var _dependLst dependLst
	var _extendDependLst dependLst
	// if field is root field parent == nil, can't use parent method
	if parent != nil {
		_dependLst = append(parent.dependLst, parent)
		if isExtend {
			_extendDependLst = parent.extendDependLst
		} else {
			_extendDependLst = append(parent.extendDependLst, parent)
		}
	}

	field := &docField{
		doc:             _doc,
		name:            t.Name,
		selfType:        reflectType,
		kind:            kind,
		parent:          parent,
		isExtend:        isExtend,
		dependLst:       _dependLst,
		extendDependLst: _extendDependLst,
		tag:             newTag(tag),
	}

	// add parent childs
	// if field is root field parent == nil, can't use parent method
	if parent != nil {
		if parent.isGroupType() {
			parent.childLst = append(parent.childLst, field)
		}
	}
	// set extendparent
	field.extendParent = d.getExtendParent(field)
	// add item to doc fieldlst, and set field id
	field = d.addItem(field)

	switch field.kind {
	case Array:
		fallthrough
	case Map:
		_fieldType := fieldType.Type.Elem()
		count := _fieldType.NumField()
		for i := 0; i < count; i++ {
			_f := _fieldType.Field(i)
			// addFieldsLock.Add(1)
			// go newDocField(_doc, d, &_f, field)
			newDocField(_doc, d, &_f, field)
		}
	case Struct:
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			_f := fieldType.Type.Field(i)
			// addFieldsLock.Add(1)
			// go newDocField(_doc, d, &_f, field)
			newDocField(_doc, d, &_f, field)
		}

	}
	// addFieldsLock.Done()
}

func checkdocFieldisExtend(r *reflect.StructField) (b bool) {
	return r.Anonymous
}

// lst /////////////////////////
type docFieldLst []*docField
type dependLst docFieldLst

var addFieldLock sync.Mutex

func (d *docFieldLst) addItem(f *docField) *docField {
	// add lock
	// addFieldLock.Lock()
	// defer addFieldLock.Unlock()
	f.id = len(*d)
	*d = append(*d, f)
	return f
}

func (d *docFieldLst) getFieldsByName(name string) (o docFieldLst) {
	for _, v := range *d {
		if v.Name() == name {
			o = append(o, v)
		}
	}
	return
}

func (d *docFieldLst) getExtendFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.IsExtend() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *docFieldLst) getSingleTypeFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.isSingleType() {
			rd = append(rd, v)
		}
	}
	return
}
func (d *docFieldLst) getGroupTypeFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.isGroupType() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *docFieldLst) getExtendParent(field *docField) (f *docField) {
	if field.parent == nil {
		f = nil
	} else {
		if field.parent.isExtend {
			f = d.getExtendParent(field.parent)
		} else {
			f = field.parent
		}
	}
	return
}
