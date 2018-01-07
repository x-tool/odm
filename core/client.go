package core

import "github.com/x-tool/odm/odm"

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

type Client struct {
	config ConnectionConfig
}

func (c *Client) Database(name string) *odm.Database {
	return odm.NewDatabase(name)
}
