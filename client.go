package odm

import (
	"github.com/x-tool/odm/module/dialect"
	"github.com/x-tool/odm/module/dialect/model"
)

type client struct {
	config         ConnectConfig
	dialectConnect model.Dialect
}

func (c *client) Database(name string) *database {
	return newDatabase(name)
}

func NewClient(c ConnectConfig) *client {
	_o := new(client)
	_o.config = c
	_o.dialectConnect = dialect.NewDialect(c)
	return _o
}
