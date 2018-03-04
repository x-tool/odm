package core

import (
	"reflect"
)

type Col struct {
	database *Database
	name     string
	doc      *doc
}

func (c *Col) GetName() string {
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
	c.doc = NewDoc(c, i)
	return c
}

// GetColName get interface name
func GetColName(i interface{}) (name string) {
	// if i = ColInterface use method to get name
	if ColI, ok := i.(ColInterface); ok {
		name = ColI.ColName()
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

// ColLst //////////////////
type ColLst []*Col

func (cL *ColLst) GetCol(i interface{}) (c *Col) {
	ColName := GetColName(i)
	return cL.GetColByName(ColName)
}

func (cL *ColLst) GetColByName(name string) *Col {
	var Col *Col
	for _, v := range *cL {
		if v.GetName() == name {
			Col = v
			break
		}
	}
	return Col
}
