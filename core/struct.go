package core

import (
	"errors"
	"reflect"

	"github.com/x-tool/tool"
)

type odmStruct struct {
	name         string
	fields       structFieldLst
	sourceType   *reflect.Type
	mode         *structField
	fieldTagMap  map[string]*structField
	fieldNameMap map[string]structFieldLst
	rootFields   structFieldLst
}
type odmStructLst []*odmStruct

func (d *odmStruct) getChildFields(i *structField) (r structFieldLst) {
	return i.childLst
}

func (d *odmStruct) getChildFieldByName(i *structField, s string) (r *structField) {
	for _, v := range i.childLst {
		if v.Name() == s {
			r = v
			break
		}
	}
	return
}

func (d *odmStruct) getFieldById(id int) (o *structField) {
	for _, v := range d.fields {
		if v.GetID() == id {
			o = v
			return o
		}
	}
	return
}

func (d *odmStruct) getFieldByName(name string) (o structFieldLst) {
	return d.fieldNameMap[name]
}

func (d *odmStruct) getFieldByTag(tag string) (o *structField) {
	return d.fieldTagMap[tag]
}

func (d *odmStruct) GetRootFields() structFieldLst {
	return d.rootFields
}
func (d *odmStruct) getStructRootFields() (lst structFieldLst) {
	for _, v := range d.fields {
		if v.parent == nil {
			lst = append(lst, v)
		}
	}
	return
}
func NewOdmStruct(i interface{}) (_odmStruct *odmStruct) {

	// append odmStruct.fields
	_odmStructSourceT := reflect.TypeOf(i)
	odmStructSourceT := _odmStructSourceT.Elem()
	_odmStruct = &odmStruct{
		name:       odmStructSourceT.Name(),
		sourceType: &odmStructSourceT,
	}
	fields := newstructFieldLst(_odmStruct, odmStructSourceT)
	_odmStruct.fields = *fields
	_odmStruct.fieldTagMap = _odmStruct.makestructFieldLstTagMap()
	_odmStruct.fieldNameMap = _odmStruct.makestructFieldLstNameMap()
	_odmStruct.rootFields = _odmStruct.makerootFieldNameMap()
	return
}

// var addFieldsLock sync.WaitGroup

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

func (d *odmStruct) makestructFieldLstTagMap() (m map[string]*structField) {
	_d := d.fields
	for _, v := range _d {
		tagPtr := v.tag.ptr
		if tagPtr != "" {
			m[tagPtr] = v
		}
	}
	return m
}

func (d *odmStruct) makestructFieldLstNameMap() map[string]structFieldLst {
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

func (d *odmStruct) makerootFieldNameMap() (lst []*structField) {
	_d := d.fields
	for _, v := range _d {
		if v.extendParent == nil && v.IsExtend() == false {
			lst = append(lst, v)
		}
	}
	return
}
