package odm

import (
	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
)

type Connect = client.ConnectConfig

func New(c Connect) *ODMClient {
	_c := new(ODMClient)
	_c.sourceClient = core.NewClient(connect)
	return _c
}

type ODMClient struct {
	sourceClient *client.Client
}

// if d == nil use default dialect
func (c *ODMClient) Database(d core.Dialect) *core.Database {
	// check default database
	if c.sourceClient.GetConnectConfig().Database == "postgre" && d == nil {
		d = defaultPostgre
	}
	d.SetConnectConfig(c.sourceClient.GetConnectConfig())
	return core.NewDatabase(name, c, d)
}
