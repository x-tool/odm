package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type doc struct {
	col *col
	docItemType
	fields  docFieldLst
	colMode colModeHook
}

type docLst []*doc

type docItemType reflect.Type

func (d *doc) getChildFields(i *docField) (r docFieldLst) {
	id := i.Id
	for _, v := range d.fields {
		if v.Pid == id {
			r = append(r, v)
		}
	}
	return
}

func (d *doc) getFieldById(id int) (o *docField) {
	for _, v := range d.fields {
		if v.Id == id {
			o = v
			return o
		}
	}
	return
}

func Newdoc(c *col, i interface{}) *doc {

	// append doc.fields
	docSource := reflect.ValueOf(i)
	docSourceV := docSource.Elem()
	docSourceT := docSourceV.Type()
	doc := &doc{
		col:         c,
		docItemType: docSourceT,
	}
	if docSourceT.Kind() == reflect.Struct {
		cont := docSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := docSourceT.Field(i)
			newdocField(doc, &field, rootPid, rootPid)
		}
		// check Fields Name, Can't both same name in one Col
		// doc.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("doc type is "+docSourceT.Kind().String()+"!,Type should be Struct"))
	}

	return doc
}

// func (doc *doc) isComplexField(d *docField) bool {
// 	if d.Type == "struct" || d.Type == "map" || d.Type == "slice" {
// 		return true
// 	}
// 	return false
// }

// func (d *doc) checkFieldsName() {
// 	FieldsLen := len(d.fields)
// 	for i := 0; i < FieldsLen; i++ {
// 		for j := i + 1; j < FieldsLen; j++ {
// 			if d.fields[i].Name == d.fields[j].Name && d.fields[i].extendPid == d.fields[j].extendPid {
// 				tool.Panic("DB", errors.New("FieldsName Should special, Col Name is "+d.col.name))
// 			}
// 		}
// 	}
// }

// func (d *doc) docModel() (hasdocModel bool, docModel string) {
// 	for _, v := range d.fields {
// 		if isdocMode(v.Name) {
// 			return true, v.Name
// 		}
// 	}
// 	return
// }

func checkdocFieldisExtend(tag string) bool {
	return tagIsExtend(tag)
}
