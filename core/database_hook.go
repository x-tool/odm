package core

type Hook struct {
	db            *Database
	getFieldByStr []func(o *odmStruct, str string) (dLst dependLst, err error)
}

func newHook(db *Database) (h *Hook) {
	h = new(Hook)
	h.db = db
	h.RegisterGetField(db.getDependLstByStr)
	return h
}
func (h *Hook) RegisterGetField(f func(o *odmStruct, str string) (dLst dependLst, err error)) {
	h.getFieldByStr = append(h.getFieldByStr, f)
}
