package odm

import (
	"github.com/x-tool/odm/module/dialect"
)

type client struct {
	config ConnectConfig
}

func (c *client) Database(name string) *database {
	return newDatabase(name)
}

type ConnectConfig = dialect.ConnectConfig

func NewClient(c ConnectConfig) *client {
	_o := new(client)
	_o.config = c
	_o.dialectConnect = dialect.NewDialect(c)
	return _o
}
