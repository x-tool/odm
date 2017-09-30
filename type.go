package odm

import (
	"reflect"
)

var mapType = map[string]string{
	"Time":    "time",
	"string":  "string",
	"int":     "int",
	"float64": "float64",
	"struct":  "struct",
}

func formatTypeToString(t *reflect.Type) (s string) {
	b := *t
	s = mapType[b.Name()]
	return
}
