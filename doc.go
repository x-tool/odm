package odm

type Doc struct {
	Col    *Col
	DB     *Database
	Handle *handle
	Query  *query
	Result *result
	Err    error
}

func newDoc(c *Col, i interface{}) *Doc {
	d := &Doc{
		Col:    c,
		DB:     c.DB,
		Handle: newHandle(),
		Query:  newQuery(),
		Result: newResult(c, i),
		Err:    nil,
	}
	return d
}

func (d *Doc) dbName() string {
	return d.DB.name
}
func (d *Doc) colName() string {
	return d.Col.Name
}
func (d *Doc) insert() (r interface{}, err error) {
	r, err = d.DB.Dialect.Insert(d)
	return
}

func (d *Doc) update(i interface{}) {

}
func (d *Doc) delete(i interface{}) {

}

func (d *Doc) query(i interface{}) {

}

func (d *Doc) Where(s string) *Doc {
	d.Handle.where = s
	return d
}

func (d *Doc) Limit(s string) *Doc {
	d.Handle.limit = s
	return d
}
