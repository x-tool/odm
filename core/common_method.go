package core

import "reflect"

// struct path and name in string
func allName(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	s := t.PkgPath() + t.Name()
	return s
}
