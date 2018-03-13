package core

import "reflect"

func (d *Handle) insert(v *reflect.Value) (err error) {
	d.execBefore()
	d.setValue = v
	err = d.Col.database.dialect.Insert(d)
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

func (d *Handle) Where(s string) *Handle {
	// d.Handle.where = s
	return d
}

func (d *Handle) Limit(s string) *Handle {
	// d.Handle.limit = s
	return d
}

func (d *Handle) execBefore() {
	d.callDocMode()
}
func (d *Handle) callDocMode() {
	callDocMode(d)
}
