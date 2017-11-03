package odm

func (q *query) key(s string) {
	q.addWhere("key", "=", s)
}

func (q *query) addWhere(w string, contrast string, i interface{}) {
	qItem := queryItem{
		dependDoc:  q.dependtoDocOneStr(w),
		whereCheck: contrast,
		whereV:     i,
	}
	q.queryFormat.queryLst = append(q.queryFormat.queryLst, qItem)
}
