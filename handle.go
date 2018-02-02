package odm

import "reflect"

type Handle struct {
	Col    *Col
	Query  *query
	Result *result
	R      *reflect.Value
	Err    error
}

func newHandle(c *Col) *Handle {
	d := &Handle{
		Col: c,
	}
	return d
}

func (d *Handle) dbName() string {
	return d.Col.DB.name
}

func (d *Handle) colName() string {
	return d.Col.name
}

func (d *Handle) insert(i interface{}) (err error) {
	r := reflect.Indirect(reflect.ValueOf(i))
	d.R = &r

	d.Query = newQuery(&r, d, "insert")
	d.Result = newResult(&r, d.Col)
	modeInsert(d)
	err = d.Col.DB.pluginInterface.handle(d)
	return
}

func (d *Handle) update(i interface{}) {

}

func (d *Handle) delete(err error) {
	if d.Col.Doc.getDeleteFieldName() != "" {
		err = d.DB.Dialect.Update(d)
	} else {
		err = d.DB.Dialect.Delete(d)
	}

	return
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

func (d *Handle) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}
