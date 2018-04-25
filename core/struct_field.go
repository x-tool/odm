package core

import (
	"reflect"
)

type structField struct {
	odmStruct             *odmStruct
	name                  string
	selfType              reflect.Type
	kind                  Kind
	id                    int
	isExtend              bool
	childLst              structFieldLst
	parent                *structField // field golang parent real
	dependLst             dependLst
	AcrossStructdependLst dependLst    // struct can be value in col's interface field,use this field to get all path
	extendParent          *structField // field Handle parent
	extendDependLst       dependLst
	tag                   *odmTag
	funcLst               map[string]string
}

func (d *structField) Name() string {
	return d.name
}

func (d *structField) GetID() int {
	return d.id
}

func (d *structField) Kind() Kind {
	return d.kind
}
func (d *structField) IsExtend() bool {
	return d.isExtend
}

func (d *structField) isSingleType() (b bool) {
	return !d.kind.isGroupType()
}

func (d *structField) isGroupType() (b bool) {
	return d.kind.isGroupType()
}

func newStructField(_odmStruct *odmStruct, d *structFieldLst, t *reflect.StructField, parent *structField) {
	fieldType := *t
	reflectType := fieldType.Type
	tag := fieldType.Tag.Get(tagName)
	kind := reflectToKind(&reflectType)
	isExtend := checkstructFieldisExtend(t)
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

	field := &structField{
		odmStruct:       _odmStruct,
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
			// go newStructField(_doc, d, &_f, field)
			newStructField(_odmStruct, d, &_f, field)
		}
	case Struct:
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			_f := fieldType.Type.Field(i)
			// addFieldsLock.Add(1)
			// go newStructField(_doc, d, &_f, field)
			newStructField(_odmStruct, d, &_f, field)
		}

	}
	// addFieldsLock.Done()
}

func checkstructFieldisExtend(r *reflect.StructField) (b bool) {
	return r.Anonymous
}

// lst /////////////////////////
type structFieldLst []*structField
type dependLst []*structField

// var addFieldLock sync.Mutex

func (d *structFieldLst) addItem(f *structField) *structField {
	// add lock
	// addFieldLock.Lock()
	// defer addFieldLock.Unlock()
	f.id = len(*d)
	*d = append(*d, f)
	return f
}

func (d *structFieldLst) getFieldsByName(name string) (o structFieldLst) {
	for _, v := range *d {
		if v.Name() == name {
			o = append(o, v)
		}
	}
	return
}

func (d *structFieldLst) getExtendFieldLst() (rd structFieldLst) {
	for _, v := range *d {
		if v.IsExtend() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *structFieldLst) getSingleTypeFieldLst() (rd structFieldLst) {
	for _, v := range *d {
		if v.isSingleType() {
			rd = append(rd, v)
		}
	}
	return
}
func (d *structFieldLst) getGroupTypeFieldLst() (rd structFieldLst) {
	for _, v := range *d {
		if v.isGroupType() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *structFieldLst) getExtendParent(field *structField) (f *structField) {
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
