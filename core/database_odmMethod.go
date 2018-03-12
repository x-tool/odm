package core

import "reflect"

func (d *Database) Insert(i interface{}) (h *Handle) {
	value := reflect.Indirect(reflect.ValueOf(i))
	col := d.GetColByName(value.Type().Name())
	return newHandle(col, insertData, &value)
}
