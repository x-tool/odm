package odm

import (
	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
)

type connect = client.ConnectConfig

type ODM struct {
}

func New() *ODM {
	return new(ODM)
}

func (o *ODM) Client(c ConnectConfig) *ODMClient {
	return newClient(c)
}

type ODMClient struct {
	sourceClient *client.Client
}

func newClient(connect) *ODMClient {
	_c := new(ODMClient)
	_c.sourceClient = core.NewClient(connect)
	return _c
}

func (c *ODMClient) Database(name string) *core.Database {
	return core.NewDatabase(name, c)
}
