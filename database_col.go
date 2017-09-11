package xodm

func (d *Database) syncCol(colI interface{}) {
	col := d.NewCol(colI)
	d.ColLst = append(d.ColLst, col)
	d.Dialect.syncCol(col)
}

func (d *Database) NewCol(i interface{}) *Col {
	return NewCol(d, i)
}
