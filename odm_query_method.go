package odm

func (q *query) key(s string) {
	q.addWhere("key", "=", s)
}

func (q *query) addWhere(wL ...interface{}) {
	var wLL = len(wL)
	if wLL != 3 && wLL != 4 {
		return
	}
	w := wL[0].(string)
	contrast := wL[1].(string)
	i := wL[2]
	var b bool
	if wLL == 4 {
		b = wL[3].(bool)
	}

	qItem := queryItem{
		dependDoc:  q.dependtoDocOneStr(w),
		whereCheck: contrast,
		whereV:     i,
		whereAnd:   b,
	}
	q.queryFormat.queryLst = append(q.queryFormat.queryLst, qItem)
}
