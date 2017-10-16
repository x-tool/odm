package odm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type Doc struct {
	col    *Col
	fields DocFields
}
type DocFields []*DocField
type DocField struct {
	Name      string
	Type      string
	DBType    string
	Id        int
	Pid       int // field golang parent real ID
	isExtend  bool
	extendPid int // field odm parent ID
	dependLst
	Tag     string
	funcLst map[string]string
}

type dependLst []*DocField

func (doc *Doc) getRootExtendFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.Pid == -1 && v.extendPid != -1 && v.isExtend {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootSinpleFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid == -1 && !v.isExtend {
			d = append(d, v)
		}
	}
	return
}

func (doc *Doc) getRootComplexFields() (d DocFields) {
	for _, v := range doc.fields {
		if v.extendPid != -1 && !v.isExtend {
			d = append(d, v)
		}
	}
	return
}

func (d *Doc) getRootDetails() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !v.isExtend {
			doc = append(doc, v)
		}
	}
	return
}
func (d *Doc) getAllRootDetails() (doc dependLst) {
	for _, v := range d.fields {
		if v.extendPid == -1 {
			doc = append(doc, v)
		}
	}
	return
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
func (d *Doc) DocModel() (docModel string, hasDocModel bool) {
	for _, v := range d.fields {
		if isDocMode(v.Name) {
			return v.Name, true
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

func NewDoc(c *Col, i interface{}) *Doc {
	doc := new(Doc)
	doc.col = c
	// append doc.fields
	docSource := reflect.ValueOf(i)
	docSourceV := docSource.Elem()
	docSourceT := docSourceV.Type()
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

func checkDocFieldisExtend(name, tag string) bool {
	isMode := isDocMode(name)
	isExtend := tagIsExtend(tag)
	return isMode || isExtend
}
