package core

import (
	"reflect"
)

func (c *Col) GetRootFields() docFieldLst {
	lst := c.doc.GetRootFields()
	return lst
}

func (c *Col) GetRootValues(rootValue *reflect.Value) (valueLst []*reflect.Value) {
	for _, v := range c.GetRootFields() {
		valueLst = append(valueLst, v.GetValueFromRootValue(rootValue))
	}
	return
}
