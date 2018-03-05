package odm

import (
	"github.com/x-tool/odm/module/dialect/postgresql"
	"github.com/x-tool/odm/module/docmod/remark2C"
)

type NormalCol = remark2C.NormalCol

var (
	defaultPostgre = postgresql.New()
)
