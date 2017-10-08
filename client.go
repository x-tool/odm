package odm

type Client struct {
	dbName string
	config ConnectionConfig
}

func NewClient(dbName string, conf ConnectionConfig) *Client {
	var c Client
	c.dbName = dbName
	c.config = conf
	return &c
}

func (c *Client) Database(name string) Database {
	var d Database
	var config ConnectionConfig
	config = c.config
	config.DatabaseName = name
	d.Dialect = initDialect(config)
	return d
}
