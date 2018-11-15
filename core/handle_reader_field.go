package core

import "reflect"

type ReaderField struct {
	dependLst
	Addr         reflect.Value
	complexValue map[int]string // slice id or map key
}
