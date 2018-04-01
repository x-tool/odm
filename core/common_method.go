package core

import "reflect"

func allName(t reflect.Type) string {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	s := t.Name() + t.PkgPath()
	return s
}
