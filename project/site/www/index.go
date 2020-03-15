package www

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	Www.Register("index", ark.Router{
		Uri:  "/",
		Name: "首页", Desc: "首页",
		Action: func(ctx *ark.Http) {
			ctx.Text("hello arkgo")
		},
	})

	Www.Register("login", ark.Router{
		Uri: "/login",
		Routing: ark.Routing{
			"get": ark.Router{
				Name: "查询登录信息", Desc: "查询登录信息",
				Action: func(ctx *ark.Http) {
					ctx.Text("hello arkgo")
				},
			},
			"post": ark.Router{
				Name: "用户登录", Desc: "用户登录",
				Action: func(ctx *ark.Http) {
					ctx.Text("hello arkgo")
				},
			},
		},
	})

}
