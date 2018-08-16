package core

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

var pathSplitStrs = map[string]string{
	"path": ".",
	"mark": tagMark,
}

var pathSplitsRegExpStr = makePathSplitsRegExpStr()

func makePathSplitsRegExpStr() string {
	var slice []string
	for _, v := range pathSplitStrs {
		slice = append(slice, v)
	}
	return strings.Join(slice, "|")
}

type dependLst []*structField

type fieldPathCell struct {
	db *Database
}

func (f *fieldPathCell) getDependLstByNameStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return
	}
	return field.dependLst, err
}

func (f *fieldPathCell) getDependLstByPathStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return
	}
	return field.dependLst, err
}

func (f *fieldPathCell) getDependLstByAllPathStr(str string) (dLst dependLst, err error) {
	splitLst := strings.SplitN(str, pathSplitStrs["path"], 2)
	targetStruct, err := f.db.getStructByName(splitLst[0])
	if err != nil {
		return
	}
	return f.getDependLstByPathStr(targetStruct, splitLst[1])
}

func (f *fieldPathCell) getDependLstByMarkStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field := o.getFieldByMark(str)
	if field == nil {
		return nil, errors.New(fmt.Sprint("Can't Find Mark By String (%v) in Struct (%v)", str, o.name))
	}
	return field.dependLst, err
}

func (f *fieldPathCell) getDependLstByStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return nil, err
	}
	return field.dependLst, err
}

func (f *fieldPathCell) getDependLstByAllStr(str string) (dLst dependLst, err error) {
	reg := regexp.MustCompile(pathSplitsRegExpStr)
	_tempStrLst := reg.Split(str, 2)
	o, err := f.db.getStructByName(_tempStrLst[0])
	if err != nil {
		return nil, err
	}
	return f.getDependLstByStr(o, _tempStrLst[1])
}

// "@mark"
// "path"
// "path:structName@mark"
// "path:structName.path"
// func getDependLstByAllPathInStruct(d Database, o *odmStruct, s string) (dLst dependLst, err error) {
// 	structLst := strings.Split(s, splitStructStr)
// 	// firstPath, without structname
// 	field, err := o.getFieldByString(structLst[0])
// 	if err != nil {
// 		return
// 	}
// 	dLst = field.dependLst
// 	// orther struct path
// 	lstLen := len(structLst)
// 	getDependLstBySplitStr(structLst[1:])
// 	return
// }

// func getDependLstByAllPath(d Database, s string) (dLst dependLst, err error) {

// }

// func getDependLstBySplitStr(d *Database, strLst []string) (dLst dependLst, err error) {
// 	for i, v := range strLst {
// 		// split name and path
// 		// now just need find "@", if add more sign to split struct in the fultrue, should modify use splitStructNameToFieldPath
// 		var sign string
// 		var targetStruct *odmStruct
// 		index := strings.Index(v, "@")
// 		if index == -1 {
// 			sign = "."
// 		} else {
// 			sign = "@"
// 		}

// 		slice := strings.SplitN(v, sign, 2)
// 		structName := slice[0]
// 		fieldRoute := slice[1]

// 		dLst = append(dLst, field.dependLst...)
// 	}
// }

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
