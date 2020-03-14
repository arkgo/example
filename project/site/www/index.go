package www

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	// Www.Router("index", Map{
	// 	"uri":  "/",
	// 	"name": "首页", "text": "首页",
	// 	"action": func(ctx *ark.Http) {
	// 		ctx.Text("hello arkgo")
	// 	},
	// })

	ark.Register("www.index", ark.Router{
		Uri:  "/user/login",
		Name: "首页", Desc: "首页",
		Action: func(ctx *ark.Http) {
			ctx.Text("hello arkgo")
		},
	})

	ark.Register("www.login", ark.Router{
		Uri:  "/user/login",
		Name: "登录", Desc: "登录",
		Route: ark.Routes{
			"get": ark.Router{
				Name: "登录查询信息", Desc: "登录查询信息",
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
