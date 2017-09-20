package xodm

import "reflect"

type Doc struct {
}
type rootField struct {
	name       string
	typeName   string
	DBtypeName string
	value      interface{}
}

func (c *Col) getRootfields(i interface{}) (r []*rootField) {
	ivalue := reflect.ValueOf(i).Elem()
	rootDetails := c.OriginDocs.getRootDetails()
	for _, v := range rootDetails.getRootSinpleFields() {
		f := &rootField{
			name:       v.Name,
			typeName:   v.Type,
			DBtypeName: v.DBType,
			value:      ivalue.FieldByName(v.Name),
		}
		r = append(r, f)
	}
	for _, v := range rootDetails.getRootComplexFields() {
		fields := c.OriginDocs.getChildFields(v)
		for _, val := range fields {
			f := &rootField{
				name:       val.Name,
				typeName:   val.Type,
				DBtypeName: val.DBType,
				value:      ivalue.FieldByName(val.Name),
			}
			r = append(r, f)
		}
	}
	return
}
