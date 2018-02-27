package odm

import (
	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
)

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

type connect = client.ConnectConfig

func newClient(connect) ODMClient {
	return core.NewClient(connect)
}

func (c *client) Database(name string) *core.Database {
	return newDatabase(name, c)
}
