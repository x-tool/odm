package core

import (
	"reflect"
)

func (c *Col) findDocModeField() (field *structField) {
	for _, v := range c.getExtendFields() {
		_value := reflect.New(v.sourceType)
		_, ok := _value.Interface().(DocMode)
		if ok {
			field = v
			break
		}
	}
	return
}

func (c *Col) getDocMode() *structField {
	return c.mode
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
