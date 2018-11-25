package core

type Hook struct {
	db *Database
}

func newHook(db *Database) (h Hook) {
	_h := new(Hook)
	_h.db = db
	return *_h
}
