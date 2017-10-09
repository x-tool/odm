package odm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type OriginDocfield struct {
	Name      string
	Type      string
	DBType    string
	Id        int
	Pid       int
	isExtend  bool
	extendPid int
	dependLst OriginDocfields
	Tag       string
	funcLst   map[string]string
}

type OriginDocfields []*OriginDocfield

func (o OriginDocfields) getRootExtendFields() (returns OriginDocfields) {
	for _, v := range o {
		if v.Pid == -1 && v.extendPid != -1 && v.isExtend {
			returns = append(returns, v)
		}
	}
	return
}

func (o OriginDocfields) getRootSinpleFields() (returns OriginDocfields) {
	for _, v := range o {
		if v.Pid == -1 && v.extendPid == -1 && !v.isExtend {
			returns = append(returns, v)
		}
	}
	return
}

func (o OriginDocfields) getRootComplexFields() (returns []*OriginDocfield) {
	for _, v := range o {
		if v.Pid == -1 && v.extendPid != -1 && !v.isExtend {
			returns = append(returns, v)
		}
	}
	return
}

type OriginDoc struct {
	DB     *Database
	Col    *Col
	fields OriginDocfields
}

func (d *OriginDoc) getRootDetails() (doc OriginDocfields) {
	for _, v := range d.fields {
		if v.extendPid == -1 && !v.isExtend {
			doc = append(doc, v)
		}
	}
	return
}
func (d *OriginDoc) getAllRootDetails() (doc OriginDocfields) {
	for _, v := range d.fields {
		if v.extendPid == -1 {
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
				tool.Panic("DB", errors.New("FieldsName Should special, Col Name is "+d.Col.name))
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

func (d *OriginDoc) getChildFields(i *OriginDocfield) (r OriginDocfields) {
	id := i.Id
	for _, v := range d.fields {
		if v.Pid == id {
			r = append(r, v)
		}
	}
	return
}
func (d *OriginDoc) getFieldsById(id int) (o *OriginDocfield) {
	for _, v := range d.fields {
		if v.Id == id {
			o = v
			return o
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
			NewOriginDocField(doc, &field, -1, -1)
		}
		// check Fields Name, Can't both same name in one Col
		doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

func NewOriginDocField(d *OriginDoc, t *reflect.StructField, Pid int, extendPid int) {
	fieldType := *t
	fieldTypeStr := formatTypeToString(&fieldType.Type)
	id := len(d.fields)
	tag := fieldType.Tag.Get(tagName)
	isExtend := checkOriginDocFieldisExtend(fieldType.Name, tag)
	extendField := d.getFieldsById(extendPid)
	var dependLst OriginDocfields
	if extendField == nil {
	} else {
		dependLst = append(extendField.dependLst, extendField)
	}

	field := &OriginDocfield{
		Name:      fieldType.Name,
		Type:      fieldTypeStr,
		DBType:    d.DB.SwitchType(fieldTypeStr),
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
			NewOriginDocField(d, &field, id, extendPid)
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
			NewOriginDocField(d, &field, id, extendPid)
		}

	}
}

func checkOriginDocFieldisExtend(name, tag string) bool {
	isMode := isDocMode(name)
	isExtend := tagIsExtend(tag)
	return isMode || isExtend
}
