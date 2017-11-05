package odm

import "reflect"

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
		queryRootField: queryRootField{
			DocField: q.dependtoDocOneStr(w),
			zero:     false,
			value:    reflect.ValueOf(i),
		},
		whereCheck: contrast,
		whereAnd:   b,
	}
	q.querySet = append(q.querySet, qItem)
}
