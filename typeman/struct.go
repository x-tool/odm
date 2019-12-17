package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type odmStruct struct {
	name string
	path string
	// parent          *StructField
	fields          StructFieldLst
	rootFields      StructFieldLst
	sourceType      *reflect.Type
	interfaceFields map[string]*StructField
	fieldMarkMap    map[string]*StructField
	fieldNameMap    map[string]StructFieldLst
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
	fields := newStructFieldLst(_odmStruct, odmStructSourceT)
	_odmStruct.fields = *fields
	_odmStruct.fieldMarkMap = makeStructFieldLstMarkMap(_odmStruct)
	_odmStruct.fieldNameMap = makeStructFieldLstNameMap(_odmStruct)
	_odmStruct.rootFields = makerootFieldNameMap(_odmStruct)
	_odmStruct.interfaceFields = makeInterfaceFields(_odmStruct)
	return
}

func newStructFieldLst(d *odmStruct, odmStructSourceT reflect.Type) *StructFieldLst {
	var lst StructFieldLst
	if odmStructSourceT.Kind() == reflect.Struct {
		cont := odmStructSourceT.NumField()
		for i := 0; i < cont; i++ {
			field := odmStructSourceT.Field(i)
			// addFieldsLock.Add(1)
			// go newStructFieldLst(d, &lst, &field, nil)
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

func makeStructFieldLstMarkMap(d *odmStruct) (m map[string]*StructField) {
	_d := d.fields
	for _, v := range _d {
		tagPtr := v.odmTag.mark
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func makeStructFieldLstNameMap(d *odmStruct) map[string]StructFieldLst {
	_d := d.fields
	var _map = make(map[string]StructFieldLst)
	for _, v := range _d {
		name := v.Name()
		// new m[name]
		if _, ok := _map[name]; !ok {
			var temp StructFieldLst
			_map[name] = temp
		}
		_map[name] = append(_map[name], v)
	}
	return _map
}

func makerootFieldNameMap(d *odmStruct) (lst []*StructField) {
	_d := d.fields
	for _, v := range _d {
		if v.logicParent == nil && v.isAnonymous() == false {
			lst = append(lst, v)
		}
	}
	return
}

func makeInterfaceFields(d *odmStruct) (lst map[string]*StructField) {
	for _, v := range d.fields {
		if _, ok := lst[v.name]; !ok {
			if v.Kind() == Interface {
				lst[v.name] = v
			}
		}

	}
	return
}
