package core

type HandleValue struct {
	v    interface{}
	zero bool
	doc  *doc
}

func newValue(v interface{}, doc *doc) (o *HandleValue) {
	o = &HandleValue{
		v:   v,
		doc: doc,
	}
	return o
}
