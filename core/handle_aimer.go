package core

type relation = string

const (
	relationLike    = "like"  // like
	relationLikeL   = "?like" // ??like
	relationLikeR   = "like?" // like??
	relationequal   = "=="    // ==
	relationNull    = "isNull"
	relationNotNull = "notNull"
	relationBetween = "between"
	relationIn      = "in"
)

type whereItem struct {
	f *structField
	relation
	values []interface{}
}

type aimer struct {
	handleCols
	whereLst []whereItem
}

func newAimer(lst *handleCols) *aimer {
	return &aimer{
		handleCols: *lst,
	}
}
