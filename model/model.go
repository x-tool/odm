package model

import "reflect"

// Database use
type Database struct {
	Name string
	ColLst
}

type ColLst []*Col

type Col struct {
	DB             *Database
	name           string
	hasDocModel    bool
	DocModel       string
	hasDeleteField bool
	deleteField    string
	Doc            *Doc
}

type Doc struct {
	col    *Col
	t      *reflect.Type
	fields DocFields
}
type DocFields []*DocField
type DocField struct {
	Name      string
	Type      string
	DBType    string
	Id        int
	Pid       int // field golang parent real ID; default:-1
	isExtend  bool
	extendPid int // field Handle parent ID; default:-1
	dependLst
	Tag     string
	funcLst map[string]string
}

type dependLst []*DocField
