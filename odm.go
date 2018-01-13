package odm

import (
	"github.com/x-tool/odm/module"
	"github.com/x-tool/odm/module/model"
)

const (
	tagName = "xodm"
)

type ConnectConfig = model.ConnectionConfig

type client struct {
	config  ConnectConfig
	connect module.Dialect
}

func (c *client) Database(name string) *Database {
	return NewDatabase(name)
}

func NewClient(c ConnectConfig) *client {
	_o := new(client)
	_o.config = c
	_o.connect = module.NewDialect(c)
	return _o
}
