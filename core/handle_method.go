package core

import (
	"errors"
	"reflect"
)

func (d *Handle) Insert(i interface{}) (err error) {
	value := reflect.Indirect(reflect.ValueOf(i))
	d.setColbyValue(&value)
	d.addSetValue(newSetValue(&value, *new(filter)))
	d.execBefore()
	if d.Err != nil {
		return d.Err
	}
	err = d.col.database.dialect.Insert(d)
	return
}

func (d *Handle) Update(i interface{}) (err error) {
	return
}

func (d *Handle) Delete(err error) {
	// if d.Col.doc.getDeleteFieldName() != "" {
	// 	err = d.DB.Dialect.Update(d)
	// } else {
	// 	err = d.DB.Dialect.Delete(d)
	// }
	// return
}

func (d *Handle) Get(i interface{}) (err error) {
	_d := d.Query(i)
	return _d.Exec()
}

func (d *Handle) Exec() (err error) {
	switch d.handleType {
	case insertData:
		err = d.col.database.dialect.Insert(d)
	case updateData:
		err = d.col.database.dialect.Update(d)
	case deleteData:
		err = d.col.database.dialect.Delete(d)
	case queryData:
		err = d.col.database.dialect.Query(d)
	}
	return
}

func (d *Handle) Query(i interface{}) (h *Handle) {
	return
}
func (d *Handle) Key(s string) (h *Handle) {
	if d.Err != nil {
		return d
	}
	return
}

func (d *Handle) Where(s string, iLst ...interface{}) (h *Handle) {
	if d.Err != nil {
		return d
	}
	// formatStringToQuery(s, iLst)
	return
}

func (d *Handle) Desc(s string, isSmallFirst bool) (h *Handle) {
	if d.Err != nil {
		return d
	}
	return
}

func (d *Handle) Limit(first int, last int) (h *Handle) {
	if d.Err != nil {
		return d
	}
	return
}

func (d *Handle) Col(i interface{}) (h *Handle) {
	if d.Err != nil {
		return d
	}
	var _col *Col
	switch i.(type) {
	case string:
		_col = d.db.getColByName(i.(string))
	default:
		name := reflect.TypeOf(i).Name()
		_col = d.db.getColByName(name)
	}
	if _col != nil {
		d.Err = errors.New("can't Find Col")
	}
	d.col = _col
	return
}

// do some thing before exec sql,like modify docmode
func (d *Handle) execBefore() {
	d.callDocMode()
}
func (d *Handle) callDocMode() {
	callDocMode(d)
}
