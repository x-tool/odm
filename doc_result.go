package odm

import (
	"reflect"
)

type result struct {
	OriginDoc *OriginDoc
	routeR    []routeResult
	routeT    reflect.Type
	raw       interface{}
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
		OriginDoc: c.OriginDocs,
	}
	return r
}

func (r *result) NewResult() (v *reflect.Value) {
	return
}

func (r *result) getRootFields() []*docRootField {
	var rootField []*docRootField
	ivalue := reflect.ValueOf(r.raw).Elem()
	rootDetails := r.OriginDoc.getRootDetails()
	for _, v := range rootDetails.getRootSinpleFields() {
		f := &docRootField{
			name:       v.Name,
			typeName:   v.Type,
			DBtypeName: v.DBType,
			value:      ivalue.FieldByName(v.Name),
		}
		rootField = append(rootField, f)
	}
	for _, v := range rootDetails.getRootComplexFields() {
		fields := r.OriginDoc.getChildFields(v)
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
