package xodm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type docfield struct {
	Name    string
	Type    string
	DBType  string
	Id      int
	Pid     int
	isRoot  bool
	Tag     string
	funcLst map[string]string
}

type Doc struct {
	DB     *Database
	Col    *Col
	fields []*docfield
}

func (c *Doc) getRootDetails() (doc []*docfield) {
	for _, v := range c.fields {
		if v.Pid == -1 {
			doc = append(doc, v)
		}
	}
	return
}

func NewDoc(c *Col, i ColInterface) *Doc {
	doc := new(Doc)
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
			NewDocField(doc, &field, -1, true)
		}
	} else {
		tool.Panic("DB", errors.New("Database Collection type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

func NewDocField(d *Doc, t *reflect.StructField, Pid int, isRoot bool) {
	fieldType := *t
	fieldTypeStr := fieldType.Type.Kind().String()
	id := len(d.fields)
	tag := fieldType.Tag.Get(tagName)
	field := &docfield{
		Name:   fieldType.Name,
		Type:   fieldTypeStr,
		DBType: d.DB.SwitchType(fieldTypeStr),
		Id:     id,
		Pid:    Pid,
		isRoot: isRoot,
		Tag:    tag,
	}
	d.fields = append(d.fields, field)
	switch t.Type.Kind() {
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Struct:
		count := fieldType.Type.NumField()
		for i := 0; i < count; i++ {
			// tag extend mode
			if Pid == -1 || tagExtend(tag) {
				isRoot = true
			}
			field := fieldType.Type.Field(i)
			NewDocField(d, &field, id, isRoot)
		}

	}
}
