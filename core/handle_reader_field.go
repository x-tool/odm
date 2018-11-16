package core

import "reflect"

type ReaderField struct {
	reader *Reader
	name   string
	dependLst
	Addr         reflect.Value
	complexValue map[int]string // slice id or map key
}

func newReaderField(r *Reader, addr reflect.Value, o *odmStruct, s string) *ReaderField {
	f := &ReaderField{
		reader: r,
		Addr:   addr,
	}
	return f
}
