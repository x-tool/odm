package odm

const (
	HandleQuery handleType = iota
	HandleUpdate
	HandleInsert
	HandleDelete
)

type handleType int
type handle struct {
	handleType
}

func newHandle(hT handleType) *handle {
	h := &handle{
		handleType: hT,
	}
	return h
}

func (h *handle) where() {

}
func (h *handle) limit() {

}
