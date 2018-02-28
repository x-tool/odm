package core

import "github.com/x-tool/odm/client"

type colModeHook interface {
}

type Dialect interface {
	Init(client.ConnectConfig) Dialect
	// Conn() (Conn, error)
	GetColNames() ([]string, error)
	SwitchType(string) string
	SyncCol(colLst)
	// base handel
	Insert(*Handle) error
	Update(*Handle) error
	Delete(*Handle) error
	Query(*Handle) (interface{}, error)
	LogSql(string)
	Session() *Session
}

// func NewDialect(c *client.ConnectConfig) (d Dialect) {
// 	switch c.Database {
// 	case "postgresql":
// 		fallthrough
// 	default:
// 		_d := new(dialectpostgre)
// 		d = _d.Init(c)
// 		return d
// 	}
// }

// type Conn interface {
// 	Open(sql string) error
// 	// Close()
// 	// Begin()
// }

type Session struct{}
type Exec interface{}
type Result interface{}

type DocModer interface {
	config()
}
