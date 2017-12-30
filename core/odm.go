package core

import "reflect"

type ODM struct {
	Col    *Col
	Query  *query
	Result *result
	R      *reflect.Value
	Err    error
}

func newODM(c *Col) *ODM {
	d := &ODM{
		Col: c,
	}
	return d
}

func (d *ODM) dbName() string {
	return d.Col.DB.name
}

func (d *ODM) colName() string {
	return d.Col.name
}

func (d *ODM) insert(i interface{}) (err error) {
	r := reflect.Indirect(reflect.ValueOf(i))
	d.R = &r

	d.Query = newQuery(&r, d, "insert")
	d.Result = newResult(&r, d.Col)
	modeInsert(d)
	err = d.Col.DB.pluginInterface.handle(d)
	return
}

func (d *ODM) update(i interface{}) {

}

func (d *ODM) delete(err error) {
	if d.Col.Doc.getDeleteFieldName() != "" {
		err = d.DB.Dialect.Update(d)
	} else {
		err = d.DB.Dialect.Delete(d)
	}

	return
}

func (d *ODM) get(i interface{}) {

}

func (d *ODM) Where(s string) *ODM {
	// d.Handle.where = s
	return d
}

func (d *ODM) Limit(s string) *ODM {
	// d.Handle.limit = s
	return d
}

func (d *ODM) selectValidFields(dLst []*queryRootField) (vLst []*queryRootField) {
	for _, v := range dLst {
		if !v.zero {
			vLst = append(vLst, v)
		}
	}
	return
}
