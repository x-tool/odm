package odm

import (
	"reflect"

	"github.com/x-tool/tool"
)

type queryRootField struct {
	DocField *DocField
	zero     bool
	value    reflect.Value
}

func (q *query) getRootFields() []*queryRootField {
	var rootField []*queryRootField
	ivalue := *q.queryValue
	if ivalue.Kind() == reflect.Ptr || ivalue.Kind() == reflect.Interface {
		ivalue = ivalue.Elem()
	}
	_d := q.Col.Doc.getRootSinpleFields()
	for _, v := range _d {
		var value reflect.Value
		if ivalue.Kind() == reflect.Struct {
			value = *q.Col.Doc.getRootDetailValue(&ivalue, v)
		} else {
			value = ivalue
		}
		f := &queryRootField{
			DocField: v,
			zero:     tool.IsZero(value),
			value:    value,
		}
		rootField = append(rootField, f)
	}
	for _, val := range q.Col.Doc.getRootComplexFields() {
		fields := q.Col.Doc.getChildFields(val)
		for _, v := range fields {
			value := *q.Col.Doc.getRootDetailValue(&ivalue, v)
			f := &queryRootField{
				DocField: v,
				zero:     tool.IsZero(value),
				value:    value,
			}
			rootField = append(rootField, f)
		}
	}
	return rootField
}
