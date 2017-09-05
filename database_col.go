package xodm

func (d *Database) syncCol(colI ColInterface) {
	col := NewCol(d, colI)
	d.Dialect.syncCol(col)
}
