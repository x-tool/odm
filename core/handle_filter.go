package core

type filterTree struct {
	children     []filterTree
	isConnectOr  bool
	isConnectNot bool
}

func newFilter(values ...interface{}) (f *filterTree, err error) {

	return
}
