package base

import (
	"github.com/arkgo/ark"
)

const (
	DB = "db"

	REMOVED = "removed"

	USER_PLATFORM = "platform"
	USER_CUSTOMER = "customer"
)

var (
	Www  = ark.Site("www")
	File = ark.Site("file")
	Api  = ark.Site("api")
)
