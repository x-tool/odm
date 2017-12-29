package odm

func NewClient(conf ConnectionConfig) *client {
	c := new(client)
	c.config = conf
	return &c
}
