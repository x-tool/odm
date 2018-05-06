package core

func (d *Database) Insert(i interface{}) (err error) {
	return newHandle(d, nil).Insert(i)
}

func (d *Database) Get(i interface{}) (err error) {
	return
}

func (d *Database) Exec() (err error) {
	return
}

func (d *Database) Key(s string) (h *Handle) {
	return
}

func (d *Database) Where(s string) (h *Handle) {
	return
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
