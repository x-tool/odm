package core

import "reflect"

func (d *Handle) Insert(i interface{}) (err error) {
	value := reflect.Indirect(reflect.ValueOf(i))
	d.setColbyValue(&value)
	d.setValue = &value
	d.execBefore()
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

func (d *Handle) Query(i interface{}) (err error) {
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
	switch i.(type) {
	case string:
		d.col = d.db.getColByName(i.(string))
	default:
		name := reflect.TypeOf(i).Name()
		d.col = d.db.getColByName(name)
	}
	return
}

func (d *Handle) execBefore() {
	d.callDocMode()
}
func (d *Handle) callDocMode() {
	callDocMode(d)
}
