package xodm

import (
	"reflect"
)

type ColInterface interface {
	ColName() string
}

type Col struct {
	DB   *Database
	Name string
	Doc  *Doc
}

func NewCol(d *Database, i interface{}) *Col {
	c := new(Col)
	var name string
	if colI, ok := i.(ColInterface); ok {
		name = colI.ColName()
	} else {
		name = reflect.TypeOf(i).Name()
	}
	c.Name = name
	c.DB = d
	c.Doc = NewDoc(c, i)
	return c
}
