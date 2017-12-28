package odm

import "github.com/x-tool/odm/core"

// config Struct
type ConnectionConfig struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	DatabaseName string
	Database     string
	TLs          bool
}

type client struct {
	config ConnectionConfig
}

func (c *client) Database(name string) core.Database {
	var d core.Database
	var config ConnectionConfig
	config = c.config
	config.DatabaseName = name
	d.Dialect = initDialect(config)
	return d
}
