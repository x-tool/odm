package xodm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type OriginDocfieldType struct {
	Name      string
	Type      string
	DBType    string
	Id        int
	Pid       int
	isExtend  bool
	extendPid int
	Tag       string
	funcLst   map[string]string
}

type OriginDoc struct {
	DB     *Database
	Col    *Col
	fields []*OriginDocfieldType
}

func (d *OriginDoc) getRootDetails() (doc []*OriginDocfieldType) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !tagExtend(v.Tag) {
			doc = append(doc, v)
		}
	}
	return
}

func (d *OriginDoc) checkFieldsName() {
	FieldsLen := len(d.fields)
	for i := 0; i < FieldsLen; i++ {
		for j := i + 1; j < FieldsLen; j++ {
			if d.fields[i].Name == d.fields[j].Name && d.fields[i].extendPid == d.fields[j].extendPid {
				tool.Panic("DB", errors.New("FieldsName Should special, Col Name is "+d.Col.Name))
			}
		}
	}
}
func (d *OriginDoc) DocModel() (docModel string, hasDocModel bool) {
	for _, v := range d.fields {
		if isDocMode(v.Name) {
			return v.Name, true
		}
	}
	return
}
func NewOriginDoc(c *Col, i interface{}) *OriginDoc {
	doc := new(OriginDoc)
	doc.Col = c
	doc.DB = c.DB
	// append doc.fields
	docSource := reflect.ValueOf(i)
	docSourceV := docSource.Elem()
	docSourceT := docSourceV.Type()
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			NewOriginDocFieldType(doc, &field, -1, -1)
		}
		// check Fields Name, Can't both same name in one Col
		doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

func NewOriginDocFieldType(d *OriginDoc, t *reflect.StructField, Pid int, extendPid int) {
	fieldType := *t
	fieldTypeStr := fieldType.Type.Kind().String()
	id := len(d.fields)
	tag := fieldType.Tag.Get(tagName)
	isExtend := tagExtend(tag)
	field := &OriginDocfieldType{
		Name:      fieldType.Name,
		Type:      fieldTypeStr,
		DBType:    d.DB.SwitchType(fieldTypeStr),
		Id:        id,
		Pid:       Pid,
		isExtend:  isExtend,
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
			NewOriginDocFieldType(d, &field, id, extendPid)
		}
	case reflect.Struct:
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			if isExtend {
				extendPid = Pid
			} else {
				extendPid = id
			}
			field := fieldType.Type.Field(i)
			NewOriginDocFieldType(d, &field, id, extendPid)
		}

	}
}
