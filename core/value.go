package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

type Value struct {
	field    *docField
	origin   interface{}
	reflect  *reflect.Value
	hasValue bool
	zero     bool
}

func newValue(v interface{}, ptr string) (o *Value) {
	_v := reflect.ValueOf(v)
	o = &Value{
		origin:  v,
		reflect: &_v,
	}
	return o
}

func (v *Value) Value() *docField {
	return v.field
}

func (v *Value) Kind() Kind {
	return v.field.GetKind()
}

func (v *Value) FieldName() string {
	return v.field.GetName()
}

func (v *Value) ReflectValue() *reflect.Value {
	return v.reflect
}

func (v *Value) ReflectType() reflect.Type {
	return v.field.selfType
}

func ValueToString(value Value) (s string) {
	v := *value.ReflectValue()
	valueType := value.ReflectType()
	switch valueType.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(v.Bool())
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		s = strconv.FormatInt(v.Int(), 10)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		s = strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		s = strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		s = ""
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		b, err := json.Marshal(v.Interface())
		if err != nil {
			s = ""
		} else {
			s = string(b)
		}
	case reflect.String:
		s = v.String()
	case reflect.Struct:
		pkgPath := valueType.PkgPath()
		switch pkgPath {
		case "time":
			s = v.Interface().(time.Time).String()
		default:
			b, err := json.Marshal(v)
			if err != nil {
				s = ""
			} else {
				s = string(b)
			}
		}

	}
	return
}
