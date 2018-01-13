package odm

import (
	"reflect"

	"github.com/x-tool/odm/model"
)

type ColInterface interface {
	ColName() string
}
type Col = model.Col

func NewCol(d *Database, i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.DB = d
	c.Doc = NewDoc(c, i)
	c.hasDocModel, c.DocModel = c.Doc.DocModel()
	c.deleteField = c.Doc.getDeleteFieldName()
	if c.deleteField != "" {
		c.hasDeleteField = true
	}
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

func (c *Col) Insert(i interface{}) error {
	Handle := newHandle(c)
	err = Handle.insert(i)
	return err
}
func (c *Col) Update() {

}
func (c *Col) Delete() {

}
func (c *Col) Query() {

}
func (c *Col) Key(s string) (o *Handle) {
	return c.DB.Key(s, c)
}
