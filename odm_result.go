package odm

import (
	"reflect"
)

type result struct {
	Doc    *Doc
	routeR []routeResult
	routeT reflect.Type
	raw    interface{}
}

type docRootField struct {
	name       string
	typeName   string
	DBtypeName string
	value      reflect.Value
}

type routeResult struct {
	v reflect.Type
}

func newResult(i interface{}, c *Col) *result {
	r := &result{
		Doc: c.Doc,
		raw: i,
	}
	return r
}
func newResultWithoutCol(i interface{}) *result {
	r := &result{
		Doc: nil,
		raw: i,
	}
	return r
}
func (r *result) NewResult() (v *reflect.Value) {
	return
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
			name:       v.Name,
			typeName:   v.Type,
			DBtypeName: v.DBType,
			value:      value,
		}
		rootField = append(rootField, f)
	}
	for _, v := range r.Doc.getRootComplexFields() {
		fields := r.Doc.getChildFields(v)
		for _, val := range fields {
			f := &docRootField{
				name:       val.Name,
				typeName:   val.Type,
				DBtypeName: val.DBType,
				value:      ivalue.FieldByName(val.Name),
			}
			rootField = append(rootField, f)
		}
	}
	return rootField
}

func (r *result) selectValidFields(dLst []*docRootField) (vLst []*docRootField) {
	for _, v := range dLst {
		if v.value.IsValid() {
			vLst = append(vLst, v)
		}
	}
	return
}
