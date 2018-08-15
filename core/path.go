package core

import (
	"reflect"
	"strings"
)

// this vars could be user modify, so use var not const
var (
	dependPathSplit = "."
)

type dependLst []*structField

type fieldPathCell struct {
	db *Database
}

func (f *fieldPathCell) getDependLstByPathStr(o *odmStruct, str string) (dLst dependLst, err error) {
	field, err := o.getFieldByString(str)
	if err != nil {
		return
	}
	return field.dependLst, err
}
func (f *fieldPathCell) getDependLstByAllPathStr(str string) (dLst dependLst, err error) {
	splitLst := strings.SplitN(str, dependPathSplit, 2)
	targetStruct, err := f.db.getStructByName(splitLst[0])
	if err != nil {
		return
	}
	return f.getDependLstByPathStr(targetStruct, splitLst[1])
}

// "@mark"
// "path"
// "path:structName@mark"
// "path:structName.path"
func getDependLstByAllPathInStruct(d Database, o *odmStruct, s string) (dLst dependLst, err error) {
	structLst := strings.Split(s, splitStructStr)
	// firstPath, without structname
	field, err := o.getFieldByString(structLst[0])
	if err != nil {
		return
	}
	dLst = field.dependLst
	// orther struct path
	lstLen := len(structLst)
	getDependLstBySplitStr(structLst[1:])
	return
}

func getDependLstByAllPath(d Database, s string) (dLst dependLst, err error) {

}

func getDependLstBySplitStr(d *Database, strLst []string) (dLst dependLst, err error) {
	for i, v := range strLst {
		// split name and path
		// now just need find "@", if add more sign to split struct in the fultrue, should modify use splitStructNameToFieldPath
		var sign string
		var targetStruct *odmStruct
		index := strings.Index(v, "@")
		if index == -1 {
			sign = "."
		} else {
			sign = "@"
		}

		slice := strings.SplitN(v, sign, 2)
		structName := slice[0]
		fieldRoute := slice[1]

		dLst = append(dLst, field.dependLst...)
	}
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
