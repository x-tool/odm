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
type Col string

type Dialect interface {
	Init(ConnectionConfig) Dialect
	Conn() (Conn, error)
	GetTables() ([]string, error)
	SwitchType(string) string
}

type Conn interface {
	Open(...Exec) (Result, error)
	// Close()
	// Begin()
}

type Exec interface{}
type Result interface{}
