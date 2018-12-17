package core

import (
	"reflect"
)

type StructField struct {
	odmStruct  *odmStruct
	name       string
	sourceType reflect.Type
	kind       Kind
	id         int
	isExtend   bool // is Anonymous field
	odmTag
	// fields relation with golang struct
	childLst      StructFieldLst
	complexParent *StructField // the nearest complex parent field, use for check field path quick
	parent        *StructField // field golang stack parent real
	dependLst     dependLst    // depend chain, include self
	// fields relastion with logic struct
	extendChildLst  StructFieldLst
	extendParent    *StructField // field Handle parent
	extendDependLst dependLst    // depend chain, include self
}

func (d *StructField) Name() string {
	return d.name
}

func (d *StructField) ID() int {
	return d.id
}

func (d *StructField) Kind() Kind {
	return d.kind
}
func (d *StructField) IsExtend() bool {
	return d.isExtend
}

func (d *StructField) isSingleType() (b bool) {
	return !d.kind.isGroupType()
}

func (d *StructField) isGroupType() (b bool) {
	return d.kind.isGroupType()
}

func newStructField(_odmStruct *odmStruct, d *StructFieldLst, t *reflect.StructField, parent *StructField) {
	fieldType := *t
	reflectType := fieldType.Type
	tag := fieldType.Tag.Get(tagName)
	kind := reflectToKind(&reflectType)
	isExtend := checkStructFieldisExtend(t)
	var _dependLst dependLst
	var _extendDependLst dependLst

	field := &StructField{
		odmStruct:       _odmStruct,
		name:            t.Name,
		sourceType:      reflectType,
		kind:            kind,
		parent:          parent,
		isExtend:        isExtend,
		dependLst:       _dependLst,
		extendDependLst: _extendDependLst,
	}
	field.odmTag = *newTag(tag, field)
	// set extendparent, parent extendChildLst
	field.extendParent = getExtendParent(d, field)
	if field.extendParent != nil {
		field.extendParent.extendChildLst = append(field.extendParent.extendChildLst, field)
	}

	// set dependLst, extendDependLst, parent childLst
	// root field's parent is nil
	if parent != nil {
		field.dependLst = append(parent.dependLst, parent)
		if isExtend {
			field.extendDependLst = parent.extendDependLst
		} else {
			field.extendDependLst = append(parent.extendDependLst, parent)
		}
		parent.childLst = append(parent.childLst, field)

	} else {
		field.dependLst = append(field.dependLst, field)
		field.extendDependLst = append(field.extendDependLst, field)
	}

	// set complexParent
	// if get field by mark, this field should be nil
	for i := len(field.dependLst); i > 0; i-- {
		if field.dependLst[i-1].isGroupType() {
			field.complexParent = field.dependLst[i-1]
		}
	}

	// add item to fieldlst, and set field id
	field = d.addItem(field)

	// group type range filed's child
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

func checkStructFieldisExtend(r *reflect.StructField) (b bool) {
	return r.Anonymous
}

// lst /////////////////////////
type StructFieldLst []*StructField
type dependLst []*StructField

func (d *StructFieldLst) addItem(f *StructField) *StructField {
	f.id = len(*d)
	*d = append(*d, f)
	return f
}

func (d *StructFieldLst) getFieldsByName(name string) (o StructFieldLst) {
	for _, v := range *d {
		if v.Name() == name {
			o = append(o, v)
		}
	}
	return
}

func (d *StructFieldLst) getExtendFieldLst() (rd StructFieldLst) {
	for _, v := range *d {
		if v.IsExtend() {
			rd = append(rd, v)
		}
	}
	return
}

func (d *StructFieldLst) getSingleTypeFieldLst() (rd StructFieldLst) {
	for _, v := range *d {
		if v.isSingleType() {
			rd = append(rd, v)
		}
	}
	return
}
func (d *StructFieldLst) getGroupTypeFieldLst() (rd StructFieldLst) {
	for _, v := range *d {
		if v.isGroupType() {
			rd = append(rd, v)
		}
	}
	return
}

func getExtendParent(d *StructFieldLst, field *StructField) (f *StructField) {
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
