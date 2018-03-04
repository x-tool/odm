package core

type HandleSet struct {
	setLst []HandleSetItem
}

type HandleSetItem struct {
	dependDoc *docField
	value     interface{}
}

func (HandleSet *HandleSet) addHandleSet(o *Handle, str string, value interface{}) {
	_HandleSet := HandleSetItem{
		dependDoc: o.dependtoDocOneStr(str),
		value:     value,
	}
	HandleSet.setLst = append(HandleSet.setLst, _HandleSet)
}
