package core

import (
	"reflect"
)

type queryRootField struct {
	DocField *structField
	zero     bool
	value    reflect.Value
}
