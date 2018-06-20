package core

import (
	"encoding/json"
	"reflect"
	"strconv"
	"time"
)

type ValueLst []*Value
type Value struct {
	field    *structField
	reflect  *reflect.Value
	hasValue bool
	zero     bool
}

func newValue(v interface{}, field *structField) (o *Value) {
	_v := reflect.ValueOf(v)
	o = &Value{
		field:   field,
		reflect: &_v,
	}
	return o
}
func newValueByReflect(v *reflect.Value, field *structField) (o *Value) {
	o = &Value{
		field:   field,
		reflect: v,
	}
	return o
}
func (v *Value) Value() *structField {
	return v.field
}

func (v *Value) Kind() Kind {
	return v.field.Kind()
}

func (v *Value) FieldName() string {
	return v.field.Name()
}

func (v *Value) ReflectValue() *reflect.Value {
	return v.reflect
}

func (v *Value) ReflectType() reflect.Type {
	return v.field.sourceType
}

func (v *Value) Interface() interface{} {
	return v.reflect.Interface()
}

func (v *Value) String() string {
	return ValueToString(v)
}

func ValueToString(value *Value) (s string) {
	_value := *value
	v := *_value.ReflectValue()
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
