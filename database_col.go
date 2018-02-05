package odm

func (d *database) syncCol() {
	d.Dialect.syncCol(d.colLst)
}
