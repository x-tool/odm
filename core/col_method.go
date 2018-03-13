package core

import (
	"reflect"
)

func (c *Col) GetRootFields() docFieldLst {
	lst := c.doc.GetRootFields()
	return lst
}

func (c *Col) GetRootValues(rootValue *reflect.Value) (valueLst ValueLst) {
	for _, v := range c.GetRootFields() {
		value := newValueByReflect(v.GetValueFromRootValue(rootValue), v)
		valueLst = append(valueLst, value)
	}
	return
}
