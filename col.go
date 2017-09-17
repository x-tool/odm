package xodm

import (
	"reflect"
)

type ColInterface interface {
	ColName() string
}

type Col struct {
	DB          *Database
	Name        string
	hasDocModel bool
	DocModel    string
	OriginDocs  *OriginDoc
}

func NewCol(d *Database, i interface{}) *Col {
	c := new(Col)
	c.Name = GetColName(i)
	c.DB = d
	c.OriginDocs = NewOriginDoc(c, i)
	c.DocModel, c.hasDocModel = c.OriginDocs.DocModel()
	return c
}

func GetColName(i interface{}) (name string) {
	if colI, ok := i.(ColInterface); ok {
		name = colI.ColName()
	} else {
		v := reflect.TypeOf(i)
		if v.Kind() == reflect.Ptr {
			name = v.Elem().Name()
		} else {
			name = v.Name()
		}

	}
	return
}

func (c *Col) Insert(r interface{}, i interface{}) {
	// colName := GetColName(i)
	c.DB.Insert(c, r, i)
}
func (c *Col) Update(r interface{}) {

}
func (c *Col) Delete(r interface{}) {

}
func (c *Col) Query(r interface{}) {

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
