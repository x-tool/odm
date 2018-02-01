package model

import (
	"github.com/x-tool/odm/core"
	"github.com/x-tool/odm/module"
)

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

type Dialect interface {
	Init(module.ConnectConfig) Dialect
	// Conn() (Conn, error)
	GetColNames() ([]string, error)
	SwitchType(string) string
	syncCol(*core.Col)
	// base handel
	Insert(*core.Handle) error
	Update(*core.Handle) error
	Delete(*core.Handle) error
	Query(*core.Handle) (interface{}, error)
	LogSql(string)
	Session() *Session
}
