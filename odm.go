package odm

import (
	"github.com/x-tool/odm/module/model"
)

const (
	tagName = "xodm"
)

type ConnectConfig = model.ConnectConfig

type ODM struct {
}

func New() *ODM {
	return new(ODM)
}

func (o *ODM) NewClient(c ConnectConfig) *client {
	return newClient(c)
}
