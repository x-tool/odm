package core

import (
	"reflect"
	"sync"
)

type docField struct {
	name         string
	selfType     *reflect.Type
	kind         Kind
	id           int
	parent       *docField // field golang parent real ID; default:-1
	isExtend     bool
	extendParent *docField // field Handle parent ID; default:-1
	childLst     docFieldLst
	dependLst    dependLst
	Tag          string
	funcLst      map[string]string
}

func (d *docField) GetName() string {
	return d.name
}

func (d *docField) GetId() int {
	return d.id
}

func (d *docField) IsExtend() bool {
	return d.isExtend
}

func (o *docField) getRootFieldDB() (r *docField) {
	switch len(o.dependLst) {
	case 0:
		return nil
	default:
		if o.dependLst[0].isExtend {
			return o.dependLst[1]
		} else {
			return o.dependLst[0]
		}
	}
}

func (o *docField) getDependLstDB() (r docFieldLst) {
	for _, v := range o.dependLst {
		if v.isExtend {
			r = append(r, v)
		}
	}
	return
}

func (o *docField) isSingleType() (b bool) {
	return !isGroupType(o.kind)
}

func (o *docField) isGroupType() (b bool) {
	return isGroupType(o.kind)
}

func newDocField(d *docFieldLst, t *reflect.StructField, parant *docField, extendParent *docField) {
	fieldType := *t
	reflectType := fieldType.Type
	tag := fieldType.Tag.Get(tagName)
	kind := reflectToType(&reflectType)
	isExtend := checkdocFieldisExtend(t)
	var _dependLst dependLst
	if extendField == nil {
	} else {
		_dependLst = append(extendField.dependLst, extendField)
	}

	field := &docField{
		name:     reflectType.Name(),
		selfType: &reflectType,
		kind:     kind,
		// kind:    d.col.DB.SwitchType(fieldTypeStr),
		// Id:        id,
		parent:       parant,
		isExtend:     isExtend,
		dependLst:    _dependLst,
		extendParent: extendParent,
		Tag:          tag,
	}
	// add item to doc fieldlst, and set field id
	d.fields.addItem(field)
	switch field.kind {
	case Array:
		fallthrough
	case Map:
		_fieldType := fieldType.Type.Elem()
		count := _fieldType.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = pid
			} else {
				extendPid = id
			}
			field := _fieldType.Field(i)
			newDocField(d, &field, id, extendPid)
		}
	case Struct:
		// if time package not range time struct
		if t.Type.PkgPath() == "time" {
			return
		}
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = pid
			} else {
				extendPid = id
			}
			field := fieldType.Type.Field(i)
			newDocField(d, &field, id, extendPid)
		}

	}
}

func checkdocFieldisExtend(r *reflect.StructField) (b bool) {
	return r.Anonymous
}

// lst /////////////////////////
type docFieldLst []*docField
type dependLst docFieldLst

var lock *sync.Mutex

func (d docFieldLst) addItem(f *docField) *docField {
	// add lock
	lock.Lock()
	defer lock.Unlock()
	f.id = len(d)
	d = append(d, f)
	return f
}

func (d *docFieldLst) getFieldsByName(name string) (o docFieldLst) {
	for _, v := range *d {
		if v.GetName() == name {
			o = append(o, v)
		}
	}
	return
}

func (d *docFieldLst) getRootFieldLst() (rd docFieldLst) {
	for _, v := range *d {
		if v.pid == rootPid {
			rd = append(rd, v)
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
