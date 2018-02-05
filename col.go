package odm

import (
	"reflect"
)

type Col_export = col

type col struct {
	database *database
	name     string
	Doc
	colModeJ model.ColModer
}

type ColInterface interface {
	ColName() string
}

type colLst []*col

func (cL *colLst) GetCol(i interface{}) (c *col) {
	colName := GetColName(i)
	for _, v := range d.ColLst {
		if v.name == colName {
			c = v
			break
		}
	}
	return
}

func (cL *colLst) GetColByName(name string) *col {
	var col *Col
	for _, v := range cL {
		if v.GetName() == name {
			col = v
			break
		}
	}
	return col
}

func newCol(db *database, i interface{}) *col {
	c := new(col)
	c.name = GetColName(i)
	c.db = db
	c.Doc = NewDoc(c, i)
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
