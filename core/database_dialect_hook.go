package core

import (
	"github.com/x-tool/odm/client"
)

type colModeHook interface {
}

type Dialect interface {
	// Init(client.ConnectConfig) Dialect
	Init(*client.Client)
	// Conn() (Conn, error)
	GetColNames() ([]string, error)
	// SwitchType(string) string
	SyncCols(ColLst)
	// base handel
	Insert(*Handle) error
	Update(*Handle) error
	Delete(*Handle) error
	Query(*Handle) error
	LogSql(string)
	Session() *Session
}

type Session struct{}
type Exec interface{}
type Result interface{}

type DocModer interface {
	config()
}
