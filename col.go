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
	c.Name = GetColName(i)
	c.DB = d
	c.Doc = NewDoc(c, i)
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
