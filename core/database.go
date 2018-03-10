package core

import (
	"errors"
	"sync"

	"github.com/x-tool/tool"

	"github.com/x-tool/odm/client"
)

// database use
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect
	ColLst
	history *history
	mapCols map[string]*Col
}

type history struct {
	colNames []string
}

func NewDatabase(name string, c *client.Client, d Dialect) *Database {
	_d := new(Database)
	_d.name = name
	_d.client = c
	_d.dialect = d
	_d.mapCols = make(map[string]*Col)
	_d.setHistory()
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
	if _, ok := d.mapCols[_col.GetName()]; !ok {
		d.mapCols[_col.GetName()] = _col
	}
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

func (d *Database) setHistory() {
	var err error
	d.history = new(history)
	d.history.colNames, err = d.dialect.GetColNames()
	if err != nil {
		tool.Panic("DB", errors.New("Get colNames ERROR"))
	}
}

func (d *Database) getColByName(name string) *Col {
	return d.mapCols[name]
}

func (d *Database) Insert(i interface{}) {
	col := d.GetCol(i)
	handle := newHandle(col)
	handle.insert(d, i)
}
