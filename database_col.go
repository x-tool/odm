package odm

func (d *Database) syncCol(colI interface{}) {
	col := d.NewCol(colI)
	d.ColLst = append(d.ColLst, col)
	if !d.checkNativeCol(col.name) {
		d.Dialect.syncCol(col)
	}

}

// check col sync to database
func (d *Database) checkNativeCol(s string) bool {
	for _, v := range d.activeColNameLst {
		if v == s {
			return true
		}
	}
	return false
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
	odm := newODM(c.DB, c)
	err = odm.insert(i)
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

func (d *Database) Key(s string, c *Col) *ODM {
	o := newODM(c.DB, c)
	o.Query.Key(s)
	return o
}
