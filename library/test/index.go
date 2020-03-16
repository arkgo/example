package test

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
)

var (
	test = ark.Library("test")
)

func init() {

	test.Register("GetMsg", ark.Method{
		Name: "测试", Desc: "测试",
		Action: func(ctx *ark.Service) Map {
			return Map{
				"msg":     "msg from test library",
				"setting": ctx.Setting,
			}
		},
	})

}
