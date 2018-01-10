package odm

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
	Config ConnectionConfig
}

func (c *Client) Database(name string) *Database {
	return NewDatabase(name)
}
