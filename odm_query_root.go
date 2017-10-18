package odm

import (
	"reflect"

	"github.com/x-tool/tool"
)

type docRootField struct {
	DocField *DocField
	zero     bool
	value    reflect.Value
}

func (q *query) getRootFields() []*docRootField {
	var rootField []*docRootField
	ivalue := *q.queryValue
	if ivalue.Kind() == reflect.Ptr || ivalue.Kind() == reflect.Interface {
		ivalue = ivalue.Elem()
	}
	_d := q.Col.Doc.getRootSinpleFields()
	for _, v := range _d {
		var value reflect.Value
		if ivalue.Kind() == reflect.Struct {
			value = ivalue.FieldByName(v.Name)
		} else {
			value = ivalue
		}
		f := &docRootField{
			DocField: v,
			zero:     tool.CheckZero(value),
			value:    value,
		}
		rootField = append(rootField, f)
	}
	for _, val := range q.Col.Doc.getRootComplexFields() {
		fields := q.Col.Doc.getChildFields(val)
		for _, v := range fields {
			value := ivalue.FieldByName(v.Name)
			f := &docRootField{
				DocField: v,
				zero:     tool.CheckZero(value),
				value:    value,
			}
			rootField = append(rootField, f)
		}
	}
	return rootField
}
