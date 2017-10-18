package odm

import (
	"reflect"
)

type query struct {
	Col        *Col
	queryValue *reflect.Value
	queryItems
}
type queryItems []*queryItem
type queryItem struct{}

func newQuery(v *reflect.Value, c *Col) (q *query) {
	q = &query{
		Col:        c,
		queryValue: v,
	}
	return
}
