package core

type HandleValue struct {
	v    interface{}
	zero bool
	doc  *doc
}

func newValue(v interface{}, doc *doc) (o *HandleValue) {
	_v := &HandleValue{
		v:   v,
		doc: doc,
	}
	return _V
}
