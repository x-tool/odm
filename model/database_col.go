package model

func (d *database) syncCol(colI interface{}) {
	col := d.NewCol(colI)
	d.ColLst = append(d.ColLst, col)
	if !d.checkNativeCol(col.name) {
		d.Dialect.syncCol(col)
	}

}

// check col sync to database
func (d *database) checkNativeCol(s string) bool {
	for _, v := range d.activeColNameLst {
		if v == s {
			return true
		}
	}
	return false
}

func (d *database) NewCol(i interface{}) *Col {
	return NewCol(d, i)
}

func (d *database) GetCol(i interface{}) (c *Col) {
	colName := GetColName(i)
	for _, v := range d.ColLst {
		if v.name == colName {
			c = v
			break
		}
	}
	return
}
