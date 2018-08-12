package core

import (
	"github.com/x-tool/odm/client"
)

// database use
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect
	config
	states  // database some state
	zoneMap map[string]*zone
	odmStructLst
	mapStructs  map[string]*odmStruct // use map to get structs by name, I think struct name should be unique where ever package, if not user should write whole pkgPath and name in one string to get one struct
	defaultZone *zone
	history     *history
}

type config struct {
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

func (d *Database) NewZone(s string) *zone {
	z := newZone(s, d)
	d.zoneMap[s] = z
	return z
}
