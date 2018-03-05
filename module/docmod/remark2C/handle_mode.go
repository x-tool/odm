package remark2C

import (
	"time"

	"github.com/x-tool/odm/core"
	"github.com/x-tool/tool"
)

type NormalCol struct {
	Key         string
	CreatedTime time.Time
	UpdateTime  time.Time
	DeleteTime  time.Time `xHandle:"delete"`
	State       int
}

func (n *NormalCol) Create() {
	n.Key = tool.NewUniqueId()
	n.CreatedTime = time.Now()
}

func (n *NormalCol) Update() {
	n.UpdateTime = time.Now()
}
func (n *NormalCol) Delete() {
	n.DeleteTime = time.Now()
}
func (n *NormalCol) Name() (s string) {
	return "NormalCol"
}

type Mode interface {
	Create()
	Update()
	Delete()
	Name()
}

func isDocMode(s string) bool {
	var check bool
	if s == "NormalCol" {
		check = true
	}
	return check
}
func isDelete(s string) bool {
	var check bool
	if s == "NormalCol" {
		check = true
	}
	return check
}
func modeInsert(d *core.Handle) {
	// if d.Col.hasDocModel {
	// 	modeVInterface := d.Query.modeV.Addr().Interface()
	// 	v := modeVInterface.(Mode)
	// 	v.Create()
	// }
}
