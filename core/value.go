package odm

type ODMValue struct {
	v    interface{}
	zero bool
	doc  *Doc
}

func newValue(v interface{}, doc *Doc) (o *ODMValue) {
	_v := &ODMValue{
		v:   v,
		doc: doc,
	}
	return _V
}
