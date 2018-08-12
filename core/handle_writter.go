package core

type writter struct {
	setLst []*writeItem
	raw    interface{} // if insert, use this value
}

type writeItem struct {
	dependLstToRoot dependLst
	value           interface{}
}

func newWritter(insertRawValue interface{}) *writter {
	w := &writter{
		raw: insertRawValue,
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
