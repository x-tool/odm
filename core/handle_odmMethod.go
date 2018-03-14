package core

import "reflect"

func (d *Handle) insert(i interface{}) (err error) {
	value := reflect.Indirect(reflect.ValueOf(i))
	col := d.db.GetColByName(value.Type().Name())
	d.col = col
	d.setValue = &value
	d.execBefore()
	err = d.col.database.dialect.Insert(d)
	return
}

func (d *Handle) update(i interface{}) {

}

func (d *Handle) delete(err error) {
	// if d.Col.doc.getDeleteFieldName() != "" {
	// 	err = d.DB.Dialect.Update(d)
	// } else {
	// 	err = d.DB.Dialect.Delete(d)
	// }
	// return
}

func (d *Handle) get(i interface{}) {

}

func (d *Handle) Key(s string) (h *Handle) {
	return
}

func (d *Handle) Where(s string) (h *Handle) {
	return
}

func (d *Handle) Desc(s string, isSmallFirst bool) (h *Handle) {
	return
}

func (d *Handle) Limit(first int, last int) (h *Handle) {
	return
}

func (d *Handle) Col(i interface{}) (h *Handle) {
	return
}

func (d *Handle) execBefore() {
	d.callDocMode()
}
func (d *Handle) callDocMode() {
	callDocMode(d)
}
