package core

import (
	"reflect"
	"strings"
)

type odmSet struct {
	setLst []odmSetItem
}

type odmSetItem struct {
	dependDoc *Doc
	value interface{}
}

func (odmSet *odmSet)addODMSet(o *ODM, str string, value interface{}){
	_odmSet := odmSetItem{
		dependDoc: o.dependtoDocOneStr(str),
		value: value,
	}
	odmSet.setLst = append(odmSet.setLst, _odmSet)
}
