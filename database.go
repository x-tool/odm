package odm

import "github.com/x-tool/odm/module/dialect"

// database use
type database struct {
	client         *client
	name           string
	dialectConnect dialect.Dialect
	colLst
}

type databaseRelation interface {
	GetColByName(string) *Col
}

func newDatabase(name string, c *client) *database {
	_d := new(database)
	_d.Name = name
	_d.client = c
	return _d
}

func (d *database) GetName() string {
	return d.name
}

func (d *database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.colLst = append(d.colLst, _col)
}

func (d *database) RegisterCols(c ...interface{}) {
	for i := range c {
		go d.RegisterCol(i)
	}
}

func (d *database) GetName() string {
	return d.name
}

func (d *database) SyncCol() {
	d.Dialect.syncCol(d.colLst)
}
