package core

type writter struct {
	setLst   []*writeItem
	rawValue interface{} // if insert, use this value
}

type writeItem struct {
	dependLstToRoot dependLst
	value           interface{}
}

func newWritter(insertRawValue interface{}) *writter {
	w := &writter{
		rawValue: insertRawValue,
	}
	return w
}

func (w *writter) add(d dependLst, value interface{}) {
	item := &writeItem{
		dependLstToRoot: d,
		value:           value,
	}
	w.setLst = append(w.setLst, item)
}
