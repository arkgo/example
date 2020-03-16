package api

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	Api.Register("index", ark.Router{
		Uri:  "/",
		Name: "首页", Desc: "首页",
		Action: func(ctx *ark.Http) {
			ctx.Text("hello arkgo api")
		},
	})

	Api.Register("login", ark.Router{
		Uri: "/login",
		Method: ark.Routing{
			"get": ark.Router{
				Name: "查询登录信息", Desc: "查询登录信息",
				Action: func(ctx *ark.Http) {
					ctx.Text("hello login")
				},
			},
			"post": ark.Router{
				Name: "用户登录", Desc: "用户登录",
				Args: ark.Params{
					"mobile":   ark.Param{Type: "mobile", Require: true, Name: "手机号"},
					"password": ark.Param{Type: "hash", Require: true, Name: "密码"},
				},
				Action: func(ctx *ark.Http) {
					ctx.Signin("user", 123, "哈哈哈")
					ctx.Answer(nil, ctx.Args)
				},
			},
		},
	})

	Api.Register("logout", ark.Router{
		Uri: "/logout",
		Method: ark.Routing{
			"post": ark.Router{
				Name: "用户登出", Desc: "用户登出",
				Action: func(ctx *ark.Http) {
					ctx.Signout("user")
					ctx.Answer(ark.OK)
				},
			},
		},
	})

	Api.Register("view", ark.Router{
		Uri: "/view/{id}",
		Item: ark.Item{
			"user": ark.Entity{Require: true, Base: DB, Table: "user", Name: "用户"},
		},
		Method: ark.Routing{
			"get": ark.Router{
				Name: "查看用户信息", Desc: "查看用户信息",
				Action: func(ctx *ark.Http) {
					var item Map
					if vv, ok := ctx.Item["user"].(Map); ok {
						item = vv
					}
					ctx.Answer(nil, item)
				},
			},
		},
	})

}