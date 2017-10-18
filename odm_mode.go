package odm

import (
	"time"

	"github.com/x-tool/tool"
)

type NormalCol struct {
	Key         string
	CreatedTime time.Time
	UpdateTime  time.Time
	DeleteTime  time.Time
	State       int
}

func (n *NormalCol) Create() {
	n.Key = tool.NewUniqueId()
	n.CreatedTime = time.Now()
}

type NormalMode interface {
	Create()
}

func isDocMode(s string) bool {
	var check bool
	if s == "NormalCol" {
		check = true
	}
	return check
}

