package client

type Client struct {
	Config ConnectConfig
}

func (c *Client) GetConnectConfig() *ConnectConfig {
	return &c.Config
}

func NewClient(c ConnectConfig) *Client {
	_o := new(Client)
	_o.Config = c
	return _o
}
