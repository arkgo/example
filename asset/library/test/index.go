package test

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

var (
	test = ark.Library("test")
)

func init() {

	test.Method("GetMsg", Map{
		"name": "测试", "text": "测试",
		"action": func(ctx *ark.Service) Map {
			return Map{
				"msg": "msg from test library",
				"setting": ctx.Setting,
			}
		},
	})

}