package core

import (
	"errors"
	"fmt"
	"reflect"
)

// do some thing before exec sql,like modify docmode
func (d *Handle) execBefore() {
	d.callDocMode()
}
func (d *Handle) callDocMode() {
	callDocMode(d)
}

func (h *Handle) Alias(m map[string]string) *Handle {
	for k, v := range m {
		_odmStruct, err := h.db.getStructByName(v)
		if err != nil {
			h.Err = fmt.Errorf("can't get struct: %v from DB", v)
		}
		h.alias[k] = _odmStruct
	}
	return h
}

func (d *Handle) Insert(r runtimCall, i interface{}) (err error) {
	value := reflect.Indirect(reflect.ValueOf(i))
	d.setColbyValue(&value)
	d.execBefore()
	if d.checkHandleErr() != nil {
		return d.Err
	}
	d.handleType = InsertData
	d.writter = *newWritter(d)
	d.writter.setWritterValue(i)
	return d.Exec()
}

func (d *Handle) Update(i interface{}) (err error) {
	return
}

func (d *Handle) Delete(i interface{}) (err error) {
	// if d.Col.doc.getDeleteFieldName() != "" {
	// 	err = d.DB.Dialect.Update(d)
	// } else {
	// 	err = d.DB.Dialect.Delete(d)
	// }
	return
}

func (d *Handle) Get(i interface{}) (err error) {
	_d := d.Query(i)
	_d.handleType = QueryData
	return _d.Exec()
}

func (d *Handle) Exec() (err error) {
	// if len(d.handleCols) == 0 {
	// 	return errors.New("you should set col")
	// }
	switch d.handleType {
	case InsertData:
		err = d.db.dialect.Insert(d)
	case UpdateData:
		err = d.db.dialect.Update(d)
	case DeleteData:
		err = d.db.dialect.Delete(d)
	case QueryData:
		err = d.db.dialect.Query(d)
	}
	return
}

func (d *Handle) Query(i interface{}) *Handle {
	d.setResult(i)
	return d
}
func (d *Handle) Key(s string) (h *Handle) {
	if d.checkHandleErr() != nil {
		return d
	}
	return
}

func (d *Handle) Where(s string, iLst ...interface{}) (h *Handle) {
	if d.checkHandleErr() != nil {
		return d
	}
	// formatStringToQuery(s, iLst)
	return
}

func (d *Handle) Desc(s string, isSmallFirst bool) (h *Handle) {
	if d.checkHandleErr() != nil {
		return d
	}
	return
}

func (d *Handle) Limit(first int, last int) (h *Handle) {
	if d.checkHandleErr() != nil {
		return d
	}
	return
}

func (d *Handle) Col(name string, i interface{}) (h *Handle) {
	if d.checkHandleErr() != nil {
		return d
	}
	var _col *Col
	switch i.(type) {
	case string:
		_col = d.db.GetColByName(i.(string))
	default:
		iName := reflect.TypeOf(i).Name()
		_col = d.db.GetColByName(iName)
	}
	if _col != nil {
		d.Err = errors.New("can't Find Col")
	} else {
		d.handleCols.add(&handleCol{_col.name, _col})
	}
	return
}

func (h *Handle) Cols(m map[string]interface{}) {
	for k, v := range m {
		h.Col(k, v)
	}
}

func (d *Handle) checkHandleErr() *error {
	if d.Err != nil {
		return &d.Err
	}
	return nil
}
