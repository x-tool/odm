package core

import (
	"reflect"
)

type StructField struct {
	odmStruct   *odmStruct
	name        string
	sourceType  reflect.Type
	kind        Kind
	id          int
	isAnonymous bool // is Anonymous field
	odmTag
	// fields relation with golang struct
	childLst      StructFieldLst
	complexParent *StructField // the nearest complex parent field, use for check field path quick
	parent        *StructField // field golang stack parent real
	dependLst     dependLst    // depend chain, include self
	// fields relastion with logic struct
	logicChildLst  StructFieldLst
	logicParent    *StructField // field Handle parent
	logicDependLst dependLst    // depend chain, include self
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
func (d *StructField) isAnonymous() bool {
	return d.isAnonymous
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
	isAnonymous := checkStructFieldisAnonymous(t)
	var _dependLst dependLst
	var _logicDependLst dependLst

	field := &StructField{
		odmStruct:      _odmStruct,
		name:           t.Name,
		sourceType:     reflectType,
		kind:           kind,
		parent:         parent,
		isAnonymous:    isAnonymous,
		dependLst:      _dependLst,
		logicDependLst: _logicDependLst,
	}
	field.odmTag = *newTag(tag, field)
	// set logicParent, parent logicChildLst
	field.logicParent = getlogicParent(d, field)
	if field.logicParent != nil {
		field.logicParent.logicChildLst = append(field.logicParent.logicChildLst, field)
	}

	// set dependLst, logicDependLst, parent childLst
	// root field's parent is nil
	if parent != nil {
		field.dependLst = append(parent.dependLst, parent)
		if isAnonymous {
			field.logicDependLst = parent.logicDependLst
		} else {
			field.logicDependLst = append(parent.logicDependLst, parent)
		}
		parent.childLst = append(parent.childLst, field)

	} else {
		field.dependLst = append(field.dependLst, field)
		field.logicDependLst = append(field.logicDependLst, field)
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

func checkStructFieldisAnonymous(r *reflect.StructField) (b bool) {
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
		if v.isAnonymous() {
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

func getlogicParent(d *StructFieldLst, field *StructField) (f *StructField) {
	if field.parent == nil {
		f = nil
	} else {
		if field.parent.isAnonymous {
			f = getlogicParent(d, field.parent)
		} else {
			f = field.parent
		}
	}
	return
}
