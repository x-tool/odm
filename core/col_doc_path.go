package core

import (
	"strings"
)

var (
	splitStruct = ":"
	splitField  = []string{".", "@"}
)

type pathKind  int

const (
	pathKindTag pathKind = iota
	pathKindFieldPath
)

type target struct{
	structName string
	kind pathKind
	route []string	
}

// use three kind to get field from struct
// @tag
// fieldName
// fieldPath

// use ":" to split struct
// field:structName@tag
// field:structName.fieldName
// field:structName.fieldPath
func formatPath(allPath string) {
	structLst := strings.Split(allPath, splitStruct)
	if len(structLst) === 1{
		
	}
}