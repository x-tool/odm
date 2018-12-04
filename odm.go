package odm

import (
	"errors"

	"github.com/x-tool/odm/client"
	"github.com/x-tool/odm/core"
	"github.com/x-tool/tool"
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
	config := c.sourceClient.GetConnectConfig()
	lenD := len(d)
	if lenD >= 1 {
		dialect = d[0]
	} else if lenD == 0 {
		// check default database
		if config.Database == "postgresql" && d == nil {
			dialect = defaultPostgre
		} else {
			tool.Panic("DB", errors.New("Cannot find Database dialect"))
		}
	}

	dialect.Init(c.sourceClient)
	return core.NewDatabase(c.sourceClient, dialect)
}
