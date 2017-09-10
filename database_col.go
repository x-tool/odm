package xodm

func (d *Database) syncCol(colI ColInterface) {
	col := d.NewCol(colI)
	d.ColLst = append(d.ColLst, col)
	d.Dialect.syncCol(col)
}

func (d *Database) NewCol(i ColInterface) *Col {
	return NewCol(d, i)
}
