package core

import (
	"reflect"
)

type Kind uint

// xodm type
const (
	Bool Kind = iota
	Int
	Byte
	Float
	Complex
	Array
	Map
	String
	Time
	Date
	DateTime
	TimeStamp
	money
	Struct
	IP
	Interface
	Custom
)

var typeStringMap = map[Kind]string{
	Bool:      "bool",
	Int:       "int",
	Byte:      "byte",
	Float:     "float",
	Complex:   "complex",
	Array:     "array",
	Map:       "map",
	String:    "string",
	Time:      "time",
	Date:      "date",
	DateTime:  "datetime",
	TimeStamp: "timestamp",
	Struct:    "struct",
	IP:        "ip",
}

func (k Kind) isGroupType() (b bool) {
	switch k {
	case Array:
		fallthrough
	case Map:
		fallthrough
	case Struct:
		b = true
	}
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

// golang type to xodm type sign
func reflectToKind(r *reflect.Type) (k Kind) {
	_r := *r
	rKind := _r.Kind()
	switch rKind {
	case reflect.Bool:
		k = Bool
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		fallthrough
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		k = Int
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		k = Float
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		k = Complex
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		k = Array
	case reflect.String:
		k = String
	case reflect.Struct:
		pkgPath := _r.PkgPath()
		if pkgPath == "time" {
			k = Time
		} else {
			k = Struct
		}
	case reflect.Interface:
		k = Interface
	}

	return
}

///////////// Custom Type

type customType struct {
	defaultFuncMap map[string]func() interface{}
	typeLst        map[string]reflect.Type
}

// user custom type, Ex: uuid
type customTypeItem struct {
	sourceType reflect.Type
	method     customTypeInterface
}

type customTypeInterface interface {
	String() string
	Parse([]byte) (interface{}, error)
}

func newCustomType() customType {
	c := customType{
		defaultFuncMap: make(map[string]func() interface{}),
		typeLst:        make(map[string]reflect.Type),
	}
	return c
}

func (c *customType) RegisterType(name string, r reflect.Type) {
	if _, ok := c.typeLst[name]; !ok {
		c.typeLst[name] = r
	}
}

func (c *customType) RegisterDefaultValueFunc(name string, f func() interface{}) {
	if _, ok := c.defaultFuncMap[name]; !ok {
		c.defaultFuncMap[name] = f
	}
}
