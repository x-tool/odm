package odm

import "reflect"

type docRootField struct {
	DocField *DocField
	zero     bool
	value    reflect.Value
}

func (r *result) getRootFields() []*docRootField {
	var rootField []*docRootField
	ivalue := reflect.ValueOf(r.raw)
	if ivalue.Kind() == reflect.Ptr || ivalue.Kind() == reflect.Interface {
		ivalue = ivalue.Elem()
	}
	for _, v := range r.Doc.getRootSinpleFields() {
		var value reflect.Value
		if ivalue.Kind() == reflect.Struct {
			value = ivalue.FieldByName(v.Name)
		} else {
			value = ivalue
		}
		f := &docRootField{
			DocField: v,
			zero:     r.checkZero(value),
			value:    value,
		}
		rootField = append(rootField, f)
	}
	for _, val := range r.Doc.getRootComplexFields() {
		fields := r.Doc.getChildFields(val)
		for _, v := range fields {
			f := &docRootField{
				DocField: v,
				zero:     r.checkZero(value),
				value:    ivalue.FieldByName(v.Name),
			}
			rootField = append(rootField, f)
		}
	}
	return rootField
}
