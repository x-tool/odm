package client

// config Struct
type ConnectConfig struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	DatabaseName string
	Database     string
	TLs          bool
}

type Client struct {
	config ConnectConfig
}

func NewClient(c ConnectConfig) *Client {
	_o := new(Client)
	_o.config = c
	//_o.dialectConnect = dialect.NewDialect(c)
	return _o
}
