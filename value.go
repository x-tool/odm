package odm

type HandleValue struct {
	v    interface{}
	zero bool
	doc  *Doc
}

func newValue(v interface{}, doc *Doc) (o *HandleValue) {
	_v := &HandleValue{
		v:   v,
		doc: doc,
	}
	return _V
}
