package odm

import "github.com/x-tool/odm/core"

type client struct {
	*client.Client
}

type connect = client.ConnectConfig

func NewClient() *core.Client {
	return core.NewClient()
}
