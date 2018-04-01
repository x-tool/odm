package core

import (
	"reflect"
)

type Col struct {
	database *Database
	name     string
	doc
}

func (c *Col) Name() string {
	return c.name
}

// ColInterface to get name quick from interface
type ColInterface interface {
	ColName() string
}

func newCol(db *Database, i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.database = db
	c.doc = *newDoc(c, i)
	return c
}

func (c *Col) GetRootValues(rootValue *reflect.Value) (valueLst ValueLst) {
	for _, v := range c.GetRootFields() {
		value := newValueByReflect(v.GetValueFromRootValue(rootValue), v)
		valueLst = append(valueLst, value)
	}
	return
}

// GetColName get interface name
func GetColName(i interface{}) (name string) {
	v := reflect.TypeOf(i)
	name = GetColNameByReflectType(v)
	return
}

// GetColName get interface name
func GetColNameByReflectType(t reflect.Type) (name string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	name = t.Name()
	return
}

// ColLst //////////////////
type ColLst []*Col

func (cL *ColLst) GetCol(i interface{}) (c *Col) {
	ColName := GetColName(i)
	return cL.GetColByName(ColName)
}

func (cL *ColLst) GetColByName(name string) *Col {
	var Col *Col
	for _, v := range *cL {
		if v.Name() == name {
			Col = v
			break
		}
	}
	return Col
}
