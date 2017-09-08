package xodm

type Client struct {
	db     string
	config ConnectionConfig
}

func NewClient(db string, conf ConnectionConfig) *Client {
	var c Client
	c.db = db
	c.config = conf
	return &c
}

func (c *Client) Database(name string) Database {
	var d Database
	var config ConnectionConfig
	config = c.config
	config.DatabaseName = name
	if c.db == "postgresql" {
		_d := new(dialectpostgre)
		d.name = name
		d.Dialect = _d.Init(config)
	}
	return d
}
