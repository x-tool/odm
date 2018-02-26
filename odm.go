package odm

const (
	tagName = "xodm"
)

type ODM struct {
}

func New() *ODM {
	return new(ODM)
}

func (o *ODM) NewClient(c ConnectConfig) *client {
	return newClient(c)
}
