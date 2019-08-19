package core

type odmError struct {
	str string
	functionCall
}

func (e *odmError) String() string {
	return e.str
}

func newOdmError(err error) (resultErr *odmError) {
	resultErr = &odmError{
		str:          err.String(),
		functionCall: functionCall{},
	}
}
