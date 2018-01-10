package odm

import "github.com/x-tool/odm/core"

func NewClient(c core.ConnectionConfig) core.Client {
	_o := new(core.Client)
	_o.Config = c
	return _o
}
