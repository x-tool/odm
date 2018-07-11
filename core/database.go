package core

import (
	"github.com/x-tool/odm/client"
)

// database use
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect
	ColLst
	odmStructLst
	config
	states      // database some state
	zoneMap     map[string]*zone
	defaultZone *zone
	history     *history
}

type config struct {
	colNameAlias func(string) string
}

type states struct {
	isSyncCols bool
}

type history struct {
	colNames []string
}

func NewDatabase(name string, c *client.Client, d Dialect) *Database {
	_d := new(Database)
	_d.name = name
	_d.client = c
	_d.dialect = d
	_d.setHistory()
	return _d
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) GetClient() *client.Client {
	return d.client
}
