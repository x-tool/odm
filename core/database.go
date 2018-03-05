package core

import (
	"sync"

	"github.com/x-tool/odm/client"
)

// database use
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect
	ColLst
}

type history struct {
	colLst []string
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

func (d *Database) GetClient() *client.Client {
	return d.client
}

func (d *Database) RegisterCol(c interface{}) {
	_col := newCol(d, c)
	d.ColLst = append(d.ColLst, _col)
	syncLock.Done()
}

var syncLock sync.WaitGroup

func (d *Database) RegisterCols(c ...interface{}) {
	for _, v := range c {
		syncLock.Add(1)
		go d.RegisterCol(v)
	}
	syncLock.Wait()
}

func (d *Database) SyncCols() {
	d.dialect.SyncCols(d.ColLst)
}
