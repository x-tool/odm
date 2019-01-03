package core

func (d *Database) Insert(i interface{}) (err error) {
	handle := newHandle(d)
	return handle.Insert(i)
}

func (d *Database) Delete(i interface{}) (err error) {
	handle := newHandle(d)
	return handle.Delete(i)
}

func (d *Database) Get(i interface{}) (err error) {
	return
}

func (d *Database) Query(i interface{}) (err error) {
	handle := newHandle(d)
	return handle.Get(i)
}

func (d *Database) Exec() (err error) {
	return
}

func (d *Database) Key(s string) (h *Handle) {
	handle := newHandle(d)
	return handle.Key(s)
}

func (d *Database) Where(s string, iLst ...interface{}) (h *Handle) {
	handle := newHandle(d)
	return handle.Where(s, iLst)
}

func (d *Database) Desc(s string, isSmallFirst bool) (h *Handle) {
	return
}

func (d *Database) Limit(first int, last int) (h *Handle) {
	return
}
func (d *Database) UnLimit() (h *Handle) {
	return
}

func (d *Database) Col(i interface{}) (h *Handle) {
	return
}
