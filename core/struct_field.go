package core

import (
	"reflect"
)

type structField struct {
	odmStruct  *odmStruct
	name       string
	sourceType reflect.Type
	kind       Kind
	id         int
	isExtend   bool // is Anonymous field
	tag        *odmTag
	// fields relation with golang struct
	childLst      structFieldLst
	complexParent *structField // the nearest complex parent field
	parent        *structField // field golang stack parent real
	dependLst     dependLst
	// fields relastion with logic struct
	extendChildLst  structFieldLst
	extendParent    *structField // field Handle parent
	extendDependLst dependLst
}

func (d *structField) Name() string {
	return d.name
}

func (d *structField) ID() int {
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

	field := &structField{
		odmStruct:       _odmStruct,
		name:            t.Name,
		sourceType:      reflectType,
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
		field.dependLst = append(parent.dependLst, parent)
		if isExtend {
			field.extendDependLst = parent.extendDependLst
		} else {
			field.extendDependLst = append(parent.extendDependLst, parent)
		}
		if !kind.isGroupType() {
			field.dependLst = append(field.dependLst, field)
			field.extendDependLst = append(field.extendDependLst, field)
		}
		if parent.isGroupType() {
			parent.childLst = append(parent.childLst, field)
		}
	}
	// set extendparent
	field.extendParent = getExtendParent(d, field)
	// add item to doc fieldlst, and set field id
	field = d.addItem(field)

	// set field nearest complex parent field
	// if get field by mark, this field should be nil
	for i := len(field.dependLst); i > 0; i-- {
		if field.dependLst[i-1].isGroupType() {
			field.complexParent = field.dependLst[i-1]
		}
	}
	switch field.kind {
	case Array:
		fallthrough
	case Map:
		_fieldType := fieldType.Type.Elem()
		count := _fieldType.NumField()
		for i := 0; i < count; i++ {
			_f := _fieldType.Field(i)
			newStructField(_odmStruct, d, &_f, field)
		}
	case Struct:
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			_f := fieldType.Type.Field(i)
			newStructField(_odmStruct, d, &_f, field)
		}

	}
}

func checkstructFieldisExtend(r *reflect.StructField) (b bool) {
	return r.Anonymous
}

// lst /////////////////////////
type structFieldLst []*structField
type dependLst []*structField

func (d *structFieldLst) addItem(f *structField) *structField {
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

func getExtendParent(d *structFieldLst, field *structField) (f *structField) {
	if field.parent == nil {
		f = nil
	} else {
		if field.parent.isExtend {
			f = getExtendParent(d, field.parent)
		} else {
			f = field.parent
		}
	}
	return
}
