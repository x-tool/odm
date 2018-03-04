package client

type Client struct {
	config ConnectConfig
}

func (c *Client) GetConnectConfig() ConnectConfig {
	return c.config
}

func NewClient(c ConnectConfig) *Client {
	_o := new(Client)
	_o.config = c
	//_o.dialectConnect = dialect.NewDialect(c)
	return _o
}
