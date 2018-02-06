package odm

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type doc struct {
	col *col
	docItemType
	fields   DocFields
	colModeJ model.DocModer
}

type docLst []*doc

type docItemType reflect.Type

func NewDoc(c *Col, i interface{}) *Doc {

	// append doc.fields
	docSource := reflect.ValueOf(i)
	docSourceV := docSource.Elem()
	docSourceT := docSourceV.Type()
	doc := &Doc{
		col:         c,
		docItemType: docSourceT,
	}
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			newDocField(doc, &field, -1, -1)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("Doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

// func (doc *Doc) isComplexField(d *DocField) bool {
// 	if d.Type == "struct" || d.Type == "map" || d.Type == "slice" {
// 		return true
// 	}
// 	return false
// }

// func (d *Doc) checkFieldsName() {
// 	FieldsLen := len(d.fields)
// 	for i := 0; i < FieldsLen; i++ {
// 		for j := i + 1; j < FieldsLen; j++ {
// 			if d.fields[i].Name == d.fields[j].Name && d.fields[i].extendPid == d.fields[j].extendPid {
// 				tool.Panic("DB", errors.New("FieldsName Should special, Col Name is "+d.col.name))
// 			}
// 		}
// 	}
// }

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

func checkDocFieldisExtend(name, tag string) bool {
	isMode := isDocMode(name)
	isExtend := tagIsExtend(tag)
	return isMode || isExtend
}
