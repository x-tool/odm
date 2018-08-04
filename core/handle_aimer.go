package core

type relation = string

const (
	relationLike    = "like"  // like
	relationLikeL   = "?like" // ??like
	relationLikeR   = "like?" // like??
	relationequal   = "=="    // ==
	relationNull    = "isNull"
	relationBetween = "between"
	relationIn      = "in"
)

type aimer struct {
	whereLst []whereItem
}

type whereItem struct {
	f *structField
	relation
	values interface{}
}
