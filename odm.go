package odm

import (
	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
)

type Connect = client.ConnectConfig

func New(c Connect) *ODMClient {
	_c := new(ODMClient)
	_c.sourceClient = client.NewClient(c)
	return _c
}

type ODMClient struct {
	sourceClient *client.Client
}

// if d == nil use default dialect
func (c *ODMClient) Database(d ...core.Dialect) *core.Database {
	var dialect core.Dialect
	if len(d) >= 1 {
		dialect = d[0]
	}
	// check default database
	config := c.sourceClient.GetConnectConfig()
	if config.Database == "postgresql" && d == nil {
		dialect = defaultPostgre
	}
	dialect.Init(c.sourceClient)
	return core.NewDatabase(config.DatabaseName, c.sourceClient, dialect)
}
