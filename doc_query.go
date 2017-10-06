package odm

type query struct {
	queryItems
}
type queryItems []*queryItem
type queryItem struct{}

func newQuery() (q *query) {
	return
}
