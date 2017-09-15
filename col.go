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
		name = reflect.TypeOf(i).Name()
	}
	return
}

func (c *Col) Insert(r interface{}, i interface{}) {
	colName := GetColName(i)
	c.DB.Insert(r, doc)
}
func (c *Col) Update(r interface{}) {

}
func (c *Col) Delete(r interface{}) {

}
func (c *Col) Query(r interface{}) {

}
