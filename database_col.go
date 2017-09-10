package xodm

func (d *Database) syncCol(colI ColInterface) {
	col := d.NewCol(d, colI)
	d.ColLst = append(d.ColLst, col)
	d.Dialect.syncCol(col)
}

func (d *Database) NewCol(i ColInterface) *Col {
	c := new(Col)
	c.Name = i.ColName()
	c.DB = d
	c.Doc = c.NewDoc(i)
	return c
}
