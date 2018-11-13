package core

type Hook struct {
	db            *Database
	getFieldByStr []func(o *odmStruct, str string) (dLst dependLst, err error)
}

func newHook(db *Database) (h Hook) {
	_h := new(Hook)
	_h.db = db
	_h.RegisterGetField(db.getDependLstByStr)
	return *_h
}
func (h *Hook) RegisterGetField(f func(o *odmStruct, str string) (dLst dependLst, err error)) {
	h.getFieldByStr = append(h.getFieldByStr, f)
}
