package core

type filterTree struct {
	children []filterTree
	isConnectOr bool
	isConnectNot bool
	isGroup bool
	filed *StructField
	operator 
	value interface{}
}

func newFilter(values ...Interface{}) (f *filterTree, err error){
	if len(values) <= 1 {
		return _, fmt.Errorf("can't")
	}
}