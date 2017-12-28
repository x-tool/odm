package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type Doc struct {
	col    *Col
	t      *reflect.Type
	fields DocFields
}
type DocFields []*DocField
type DocField struct {
	Name      string
	Type      string
	DBType    string
	Id        int
	Pid       int // field golang parent real ID; default:-1
	isExtend  bool
	extendPid int // field odm parent ID; default:-1
	dependLst
	Tag     string
	funcLst map[string]string
}

type dependLst []*DocField

func (o *DocField) getRootFieldDB() (r *DocField) {
	switch len(o.dependLst) {
	case 0:
		return 0
	default:
		if o.dependLst[0].isExtend {
			return o.dependLst[1]
		} else {
			return o.dependLst[0]
		}
	}
}
func (o *DocField) getDependLstDB() (r []*DocField) {
	for _, v := range o.dependLst {
		if v.isExtend {
			r = append(r, v)
		}
	}
	return
}
func NewDoc(c *Col, i interface{}) *Doc {

	// append doc.fields
	docSource := reflect.ValueOf(i)
	docSourceV := docSource.Elem()
	docSourceT := docSourceV.Type()
	doc := &Doc{
		col: c,
		t:   &docSourceT,
	}
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			newDocField(doc, &field, -1, -1)
		}
		// check Fields Name, Can't both same name in one Col
		doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

func newDocField(d *Doc, t *reflect.StructField, Pid int, extendPid int) {
	fieldType := *t
	fieldTypeStr := formatTypeToString(&fieldType.Type)
	id := len(d.fields)
	tag := fieldType.Tag.Get(tagName)
	isExtend := checkDocFieldisExtend(fieldType.Name, tag)
	extendField := d.getFieldById(extendPid)
	var dependLst dependLst
	if extendField == nil {
	} else {
		dependLst = append(extendField.dependLst, extendField)
	}

	field := &DocField{
		Name:      fieldType.Name,
		Type:      fieldTypeStr,
		DBType:    d.col.DB.SwitchType(fieldTypeStr),
		Id:        id,
		Pid:       Pid,
		isExtend:  isExtend,
		dependLst: dependLst,
		extendPid: extendPid,
		Tag:       tag,
	}
	d.fields = append(d.fields, field)
	switch t.Type.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		_fieldType := fieldType.Type.Elem()
		count := _fieldType.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = Pid
			} else {
				extendPid = id
			}
			field := _fieldType.Field(i)
			newDocField(d, &field, id, extendPid)
		}
	case reflect.Struct:
		// if time package not range time struct
		if t.Type.PkgPath() == "time" {
			return
		}
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = Pid
			} else {
				extendPid = id
			}
			field := fieldType.Type.Field(i)
			newDocField(d, &field, id, extendPid)
		}

	}
}

func (doc *Doc) checkComplexField(d *DocField) bool {
	if d.Type == "struct" || d.Type == "map" || d.Type == "slice" {
		return true
	}
	return false

}

func (d *Doc) checkFieldsName() {
	FieldsLen := len(d.fields)
	for i := 0; i < FieldsLen; i++ {
		for j := i + 1; j < FieldsLen; j++ {
			if d.fields[i].Name == d.fields[j].Name && d.fields[i].extendPid == d.fields[j].extendPid {
				tool.Panic("DB", errors.New("FieldsName Should special, Col Name is "+d.col.name))
			}
		}
	}
}
func (d *Doc) DocModel() (hasDocModel bool, docModel string) {
	for _, v := range d.fields {
		if isDocMode(v.Name) {
			return true, v.Name
		}
	}
	return
}

func (d *Doc) getChildFields(i *DocField) (r DocFields) {
	id := i.Id
	for _, v := range d.fields {
		if v.Pid == id {
			r = append(r, v)
		}
	}
	return
}

func (d *Doc) getFieldById(id int) (o *DocField) {
	for _, v := range d.fields {
		if v.Id == id {
			o = v
			return o
		}
	}
	return
}
func (d *Doc) getFieldByName(name string) (o DocFields) {
	for _, v := range d.fields {
		if v.Name == name {
			o = append(o, v)
		}
	}
	return
}

func (d *Doc) getDeleteFieldName() (name string) {
	for _, v := range d.fields {
		if tagIsDelete(v.Tag) {
			return v.Name
		}
	}
	return
}

func checkDocFieldisExtend(name, tag string) bool {
	isMode := isDocMode(name)
	isExtend := tagIsExtend(tag)
	return isMode || isExtend
}
