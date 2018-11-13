package core

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type dependLst []*structField

var splitStructFromPath = "|"

func (d *Database) getDependLstByStr(o *odmStruct, str string) (dLst dependLst, err error) {
	for _, v := range d.hook.getFieldByStr {
		l, err := v(o, str)
		if len(l) != 0 {
			return l, nil
		}
	}
}

func (f *Database) getDependLstBy(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return
	}
	return field.dependLst, err
}

func (f *Database) getDependLstByPathStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return
	}
	return field.dependLst, err
}

func (f *Database) getDependLstByAllPathStr(str string) (dLst dependLst, err error) {
	splitLst := strings.SplitN(str, pathSplitStrs["path"], 2)
	targetStruct, err := f.getStructByName(splitLst[0])
	if err != nil {
		return
	}
	return f.getDependLstByPathStr(targetStruct, splitLst[1])
}

func (f *Database) getDependLstByMarkStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field := o.getFieldByMark(str)
	if field == nil {
		return nil, errors.New(fmt.Sprint("Can't Find Mark By String (%v) in Struct (%v)", str, o.name))
	}
	return field.dependLst, err
}

func (f *Database) getDependLstByStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return nil, err
	}
	return field.dependLst, err
}

func (f *Database) getDependLstByAllStr(str string) (dLst dependLst, err error) {
	reg := regexp.MustCompile(pathSplitsRegExpStr)
	_tempStrLst := reg.Split(str, 2)
	o, err := f.db.getStructByName(_tempStrLst[0])
	if err != nil {
		return nil, err
	}
	return f.getDependLstByStr(o, _tempStrLst[1])
}

//****** can't get field in slice or map
func getValueByDependLst(dLst dependLst, rootValue *reflect.Value) (value *reflect.Value, err error) {
	_value := *rootValue
	for _, v := range dLst {
		if v.kind == Struct {
			_value = _value.FieldByName(v.Name())
		} else {
			// can't get field in slice or map
			// err = errors.New(fmt.Sprintf("Can't get Values in struct %d, because fieldName %d parent type is %d", d.name, v.name, v.kind.String()))
			break
		}
	}
	value = &_value
	return
}
