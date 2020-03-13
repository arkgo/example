package www

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	Www.Router("index", Map{
		"uri":  "/",
		"name": "首页", "text": "首页",
		"action": func(ctx *ark.Http) {
			ctx.Text("hello arkgo")
		},
	})

}
