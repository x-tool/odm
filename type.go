package odm

import (
	"reflect"

	"github.com/x-tool/tool"
)

var mapType = map[string]string{
	"bool":    "bool",
	"Time":    "time",
	"string":  "string",
	"int":     "int",
	"float64": "float64",
	"struct":  "struct",
	"slice":   "slice",
	"array":   "array",
	"map":     "map",
}

func formatTypeToString(t *reflect.Type) (s string) {
	b := *t
	s = mapType[b.Name()]
	return
}

func byteMap(b []byte, v reflect.Value) {
	realV := reflect.Indirect(v)
	switch realV.Type().Kind() {
	case reflect.String:
		v.SetString(string(b))
	case reflect.Int:
		i, err := tool.STI(string(b))
		if err == nil {
			i = 0
		}
		v.SetInt(int64(i))
	case reflect.Bool:
		var _v bool
		if string(b) == "true" {
			_v = true
		}
		v.SetBool(_v)
	}

}
