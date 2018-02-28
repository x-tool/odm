package core

import "github.com/x-tool/odm/client"

// database use
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect
	colLst
}

func NewDatabase(name string, c *client.Client, d Dialect) *Database {
	_d := new(Database)
	_d.name = name
	_d.client = c
	_d.dialect = d
	return _d
}

func (d *Database) GetName() string {
	return d.name
}

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.colLst = append(d.colLst, _col)
}

func (d *Database) RegisterCols(c ...interface{}) {
	for i := range c {
		go d.RegisterCol(i)
	}
}

func (d *Database) SyncCol() {
	d.dialect.SyncCol(d.colLst)
}
