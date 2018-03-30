package core

import "reflect"

type odmStruct struct {
	name string
	path string
	Type reflect.Type
}

type structLst []odmStruct
