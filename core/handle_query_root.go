package core

import (
	"reflect"

	"github.com/x-tool/tool"
)

type queryRootField struct {
	DocField *docField
	zero     bool
	value    reflect.Value
}

func (r *query) getRootFields() []*queryRootField {
	var rootField []*queryRootField
	ivalue := *r.queryV
	if ivalue.Kind() == reflect.Ptr || ivalue.Kind() == reflect.Interface {
		ivalue = ivalue.Elem()
	}
	_d := r.Col.Doc.getRootSinpleFields()
	for _, v := range _d {
		var value reflect.Value
		if ivalue.Kind() == reflect.Struct {
			value = *r.Col.Doc.getRootDetailValue(&ivalue, v)
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
	for _, val := range r.Col.Doc.getRootComplexFields() {
		fields := r.Col.Doc.getChildFields(val)
		for _, v := range fields {
			value := *r.Col.Doc.getRootDetailValue(&ivalue, v)
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
