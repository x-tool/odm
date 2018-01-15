package odm

import (
	"github.com/x-tool/odm/module/dialect"
	"github.com/x-tool/odm/module/dialect/model"
)

type client struct {
	config  ConnectConfig
	connect model.Dialect
}

func (c *client) Database(name string) *database {
	return newDatabase(name)
}

func newClient(c ConnectConfig) *client {
	_o := new(client)
	_o.config = c
	_o.connect = dialect.NewDialect(c)
	return _o
}
