package odm

import (
	"reflect"
)

var mapType = map[string]string{
	"bool":    "bool",
	"Time":    "time",
	"string":  "string",
	"int":     "int",
	"float64": "float64",
	"struct":  "struct",
	"slice":   "slice",
	"map":     "map",
}

func formatTypeToString(t *reflect.Type) (s string) {
	b := *t
	s = mapType[b.Name()]
	return
}

func mapTypeToValue(b interface{}, v *reflect.Value) {
	// realV := reflect.Indirect(*v)
	switch b.(type) {
	case string:
		v.SetString(b.(string))
	case int:
		v.SetInt(int64(b.(int)))
	case int32:
		v.SetInt(int64(b.(int32)))
	case int64:
		v.SetInt(b.(int64))
	case bool:
		v.SetBool(b.(bool))
	}

}
