package core

type odmError struct {
	str string
}

func (e *odmError) String() string {
	return e.str
}
