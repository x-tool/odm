package core

import (
	"reflect"
)

type Col_export = col

type col struct {
	database *Database
	name     string
	doc      *doc
	mode     colMode
	hasMode  bool
}

func (c *col) GetName() string {
	return c.name
}

type ColInterface interface {
	ColName() string
}

func newCol(db *Database, i interface{}) *col {
	c := new(col)
	c.name = GetColName(i)
	c.database = db
	c.doc = NewDoc(c, i)
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

// colLst //////////////////
type colLst []*col

func (cL *colLst) GetCol(i interface{}) (c *col) {
	colName := GetColName(i)
	for _, v := range *cL {
		if v.name == colName {
			c = v
			break
		}
	}
	return
}

func (cL *colLst) GetColByName(name string) *col {
	var col *col
	for _, v := range *cL {
		if v.GetName() == name {
			col = v
			break
		}
	}
	return col
}
