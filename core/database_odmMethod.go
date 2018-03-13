package core

import "reflect"

func (d *Database) Insert(i interface{}) (err error) {
	value := reflect.Indirect(reflect.ValueOf(i))
	col := d.GetColByName(value.Type().Name())
	handle := newHandle(col, insertData, nil)
	return handle.insert(&value)
}
