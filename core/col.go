package core

import (
	"reflect"
)

type Col struct {
	database *Database
	name     string
	alias    string
	doc
}

func (c *Col) Name() string {
	return c.name
}

func newCol(db *Database, i interface{}) *Col {
	c := new(Col)
	c.name = GetColName(i)
	c.database = db
	c.doc = *newDoc(c, i)
	return c
}

func (c *Col) GetRootValues(instance *reflect.Value) (RootValues ValueLst) {
	for _, v := range c.GetRootFields() {
		value := newValueByReflect(v.GetValueFromRootValue(instance), v)
		RootValues = append(RootValues, value)
	}
	return
}

// GetColName get interface name
func GetColName(i interface{}) (name string) {
	v := reflect.TypeOf(i)
	name = GetColNameByReflectType(v)
	return
}

// sometimes get col name use reflect.Type Ex: when insert data, must use reflect, use this method could less one time to get name
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
