package core

import (
	"errors"
	"fmt"
	"reflect"

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
	states
	history    *history
	mapCols    map[string]*Col       // use map to get col by name
	mapStructs map[string]*odmStruct // use map to get structs by name, I think struct name should be unique where ever package, if not user should write whole pkgPath and name in one string to get one struct
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
	_d.mapCols = make(map[string]*Col)
	_d.setHistory()
	return _d
}

func (d *Database) Name() string {
	return d.name
}

func (d *Database) GetClient() *client.Client {
	return d.client
}

func (d *Database) GetCol(i interface{}) *Col {
	var name string
	if v, ok := i.(string); !ok {
		name = string(v)
	} else {
		name = reflect.TypeOf(i).Name()
	}
	return d.mapCols[name]
}

func (d *Database) getStructByName(name string) (o *odmStruct, err error) {
	o = d.mapStructs[name]
	if o == nil {
		err = errors.New(fmt.Sprintf("Can't find struct name %d in database", name))
	}
	return
}
