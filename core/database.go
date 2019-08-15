package core

import (
	"github.com/x-tool/odm/module/client"
)

// database
type Database struct {
	client  *client.Client
	name    string
	dialect Dialect // third module hook
	config          // database config
	states          // database some state when start xodm
	odmStructLst
	// use map to get structs by name, I think struct name should be unique where ever package,
	// if not user should write whole pkgPath and name in one string to get one struct
	mapStructs map[string]*odmStruct
	ColLst
	mapCols   map[string]*Col // use map to get col by name
	aliasFunc func(string) string
	customType
	Hook // use for plugin
}

type config struct {
	tagSplit string
}

type states struct {
	isSyncCols bool
}

func NewDatabase(c *client.Client, dialect Dialect) *Database {
	d := &Database{
		name:       c.config.DatabaseName,
		client:     c,
		dialect:    dialect,
		mapStructs: make(map[string]*odmStruct),
		mapCols:    make(map[string]*Col),
		Hook:       newHook(d),
	}
	return d
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) GetClient() *client.Client {
	return d.client
}
