package core

type writter struct {
	setLst []HandleSetItem
}

type HandleSetItem struct {
	field *structField
	value interface{}
}
