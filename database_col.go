package odm

func (d *Database) syncCol(colI interface{}) {
	col := d.NewCol(colI)
	d.ColLst = append(d.ColLst, col)
	d.Dialect.syncCol(col)
}

func (d *Database) NewCol(i interface{}) *Col {
	return NewCol(d, i)
}

func (d *Database) GetCol(i interface{}) (c *Col) {
	colName := GetColName(i)
	for _, v := range d.ColLst {
		if v.name == colName {
			c = v
			break
		}
	}
	return
}

func (d *Database) Insert(c *Col, i interface{}) (r interface{}, err error) {
	doc := newDoc(c, i)
	r, err = doc.insert()
	return r, err
}
func (d *Database) Get(c *Col) (r interface{}, err error) {
	return
}
func (d *Database) Delete(c *Col) (r interface{}, err error) {
	return
}
func (d *Database) Query(c *Col) (r interface{}, err error) {
	return
}
func (d *Database) All(c *Col) (r interface{}, err error) {
	return
}
