package core

type Client struct {
	config ConnectConfig
}

func (c *Client) Database(name string) *database {
	return newDatabase(name, c)
}

func NewClient(c ConnectConfig) *Client {
	_o := new(Client)
	_o.config = c
	//_o.dialectConnect = dialect.NewDialect(c)
	return _o
}
