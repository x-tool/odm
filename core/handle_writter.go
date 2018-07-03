package core

type writter struct {
	setLst   []*writeItem
	rawValue interface{} // if insert, use this value
}

type writeItem struct {
	dependLstToRoot dependLst
	value           interface{}
	expression      string        // ******must change it, It's not used in different database
	expressionValue []interface{} // ******must change it, It's not used in different database
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

// *****must change it, It's not used in different database
func (w *writter) addByExpression(d dependLst, expression string, value []interface{}) {
	item := &writeItem{
		dependLstToRoot: d,
		expression:      expression,
		expressionValue: value,
	}
	w.setLst = append(w.setLst, item)
}
