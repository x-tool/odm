package odm

type handle struct {
	handleType string
}

func newHandle(handleType string) *handle {
	h := &handle{
		handleType: handleType,
	}
	return
}
