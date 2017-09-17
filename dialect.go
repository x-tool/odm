package xodm

import "reflect"

// database table
type Table struct {
	Name    string
	Type    reflect.Type
	colsMap map[string][]*Col
	Indexes map[string]*Index
	Created string
	Updated string
	Deleted string
	Version string
}

// config Struct
type ConnectionConfig struct {
	Host         string
	Port         int64
	User         string
	Passwd       string
	DatabaseName string
	TLs          bool
}

type Index string

type Dialect interface {
	Init(ConnectionConfig) Dialect
	Conn() (Conn, error)
	GetTables() ([]string, error)
	SwitchType(string) string
	syncCol(*Col)
	// base handel
	Insert(col *Col, result interface{}, i interface{})
	Update(col *Col, result interface{}, i interface{})
	Delete(col *Col, result interface{}, i interface{})
	Query(col *Col, result interface{}, i interface{})
	Session() *Session
}

type Conn interface {
	Open(...Exec) (Result, error)
	// Close()
	// Begin()
}

type Session struct{}
type Exec interface{}
type Result interface{}
