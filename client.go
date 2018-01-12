package odm

import (
	"github.com/x-tool/odm/module"
	"github.com/x-tool/odm/module/dialect"
)

type ConnectConfig = module.ConnectionConfig

type client struct {
	config  ConnectConfig
	connect dialect.Dialect
}

func (c *client) Database(name string) *Database {
	return NewDatabase(name)
}

func NewClient(c ConnectConfig) *client {
	_o := new(client)
	_o.config = c
	_o.connect = dialect.NewDialect(c)
	return _o
}
