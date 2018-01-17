package model

// Database use
type Database struct {
	name   string
	colLst ColLst
}
type ColLst []*Col

func NewDatabase(name string) *Database {
	_d := new(Database)
	_d.Name = name
	return _d
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) RegisterCol(c interface{}) {
	_col := NewCol(d, c)
	d.colLst = append(d.colLst, _col)
}

func (d *Database) RegisterCols(c ...interface{}) {
	for i := range c {
		go d.RegisterCol(i)
	}
}

func GetColByName(name string, cLst ColLst) *Col {
	var col *Col
	for _, v := range cLst {
		if v.name == name {
			col = v
			break
		}
	}
	return col
}
