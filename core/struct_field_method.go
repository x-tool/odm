package core

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func (d *StructField) newValue() (v reflect.Value) {
	return reflect.New(d.sourceType)
}

func (d *StructField) json(v *reflect.Value) ([]byte, error) {
	_v := *v
	return json.Marshal(_v.Interface())
}

func (d *StructField) String(value *reflect.Value) (s string) {
	_value := *value
	if value.Kind() == reflect.Invalid {
		return ""
	}
	valueType := _value.Type()
	switch valueType.Kind() {
	case reflect.Bool:
		s = strconv.FormatBool(value.Bool())
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		s = strconv.FormatInt(value.Int(), 10)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		s = strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		s = strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Complex64:
		fallthrough
	case reflect.Complex128:
		s = ""
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		b, err := json.Marshal(value.Interface())
		if err != nil {
			s = ""
		} else {
			s = string(b)
		}
	case reflect.String:
		s = value.String()
	case reflect.Struct:
		pkgPath := valueType.PkgPath()
		switch pkgPath {
		case "time":
			s = value.Interface().(time.Time).String()
		default:
			b, err := json.Marshal(value)
			if err != nil {
				s = ""
			} else {
				s = string(b)
			}
		}

	}
	return
}

func (f *StructField) Parse(b []byte, valueptr reflect.Value) (err error) {
	rType := valueptr.Type()
	Errfunc := func(s string) error {
		return fmt.Errorf("can't parse field: %v in %v, use value type: '%v', %v", f.name, f.odmStruct.name, rType.Kind(), s)
	}
	if rType.Kind() != reflect.Ptr {
		return Errfunc("should use ptr to set value")
	}
	realValue := reflect.Indirect(valueptr)
	if !realValue.CanSet() {
		return Errfunc("input value can't set, please check")
	}

	switch realValue.Kind() {
	case reflect.Bool:
		str := string(b)
		if str == "true" || str == "1" {
			realValue.SetBool(true)
		}
		realValue.SetBool(false)
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		buf := bytes.NewBuffer(b)
		i64, err := binary.ReadVarint(buf)
		if err != nil {
			return err
		}
		realValue.SetInt(i64)
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		buf := bytes.NewBuffer(b)
		ui64, err := binary.ReadUvarint(buf)
		if err != nil {
			return err
		}
		realValue.SetUint(ui64)
	case reflect.Float32:
		str := string(b)
		f64, err := strconv.ParseFloat(str, 32)
		if err != nil {
			return err
		}
		realValue.SetFloat(f64)
	case reflect.Float64:
		str := string(b)
		f64, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		realValue.SetFloat(f64)
	case reflect.String:
		realValue.SetString(string(b))
	case reflect.Array:
		fallthrough
	case reflect.Slice:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Struct:
		err = json.Unmarshal(b, realValue.Interface())
		if err != nil {
			return err
		}
	case reflect.Interface:
		// set raw byte here, because can't know what type user wirte in, user should parse by self
		realValue.SetBytes(b)
	}
	return nil
}
