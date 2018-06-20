package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type odmStruct struct {
	name            string
	path            string
	allName         string // name+path
	fields          structFieldLst
	rootFields      structFieldLst
	sourceType      *reflect.Type
	interfaceFields map[string]*structField
	fieldTagMap     map[string]*structField
	fieldNameMap    map[string]structFieldLst
}
type odmStructLst []*odmStruct

func newOdmStruct(i interface{}) (_odmStruct *odmStruct) {

	// append odmStruct.fields
	_odmStructSourceT := reflect.TypeOf(i)
	odmStructSourceT := _odmStructSourceT.Elem()
	_odmStruct = &odmStruct{
		name:       odmStructSourceT.Name(),
		path:       odmStructSourceT.PkgPath(),
		allName:    allName(odmStructSourceT),
		sourceType: &odmStructSourceT,
	}
	fields := newstructFieldLst(_odmStruct, odmStructSourceT)
	_odmStruct.fields = *fields
	_odmStruct.fieldTagMap = makestructFieldLstTagMap(_odmStruct)
	_odmStruct.fieldNameMap = makestructFieldLstNameMap(_odmStruct)
	_odmStruct.rootFields = makerootFieldNameMap(_odmStruct)
	_odmStruct.interfaceFields = makeInterfaceFields(_odmStruct)
	return
}

func newstructFieldLst(d *odmStruct, odmStructSourceT reflect.Type) *structFieldLst {
	var lst structFieldLst
	if odmStructSourceT.Kind() == reflect.Struct {
		cont := odmStructSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := odmStructSourceT.Field(i)
			// addFieldsLock.Add(1)
			// go newstructFieldLst(d, &lst, &field, nil)
			newStructField(d, &lst, &field, nil)
		}
		// check Fields Name, Can't both same name in one Col
		// odmStruct.checkFieldsName()
	} else {
		tool.Panic("DB", errors.New("odmStruct type is "+odmStructSourceT.Kind().String()+"!,Type should be Struct"))
	}
	// addFieldsLock.Wait()
	return &lst
}

func makestructFieldLstTagMap(d *odmStruct) (m map[string]*structField) {
	_d := d.fields
	for _, v := range _d {
		tagPtr := v.tag.sign
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func makestructFieldLstNameMap(d *odmStruct) map[string]structFieldLst {
	_d := d.fields
	var _map = make(map[string]structFieldLst)
	for _, v := range _d {
		name := v.Name()
		// new m[name]
		if _, ok := _map[name]; !ok {
			var temp structFieldLst
			_map[name] = temp
		}
		_map[name] = append(_map[name], v)
	}
	return _map
}

func makerootFieldNameMap(d *odmStruct) (lst []*structField) {
	_d := d.fields
	for _, v := range _d {
		if v.extendParent == nil && v.IsExtend() == false {
			lst = append(lst, v)
		}
	}
	return
}

func makeInterfaceFields(d *odmStruct) (lst map[string]*structField) {
	for _, v := range d.fields {
		if _, ok := lst[v.name]; !ok {
			if v.Kind() == Interface {
				lst[v.name] = v
			}
		}

	}
	return
}
