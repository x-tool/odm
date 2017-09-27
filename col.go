package odm

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

func (c *Col) Insert(i interface{}) (interface{}, error) {
	return c.DB.Insert(c, i)
}
func (c *Col) Update() {

}
func (c *Col) Delete() {

}
func (c *Col) Query() {

}
