package site

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {

	ark.Register("user.Serial", ark.Method{
		Name: "生成序列", Desc: "生成序列",
		Action: func(ctx *ark.Service) (Map, *Res) {
			serial := ark.CacheSerial("hahha", 0, 1)
			return Map{
				"serial": serial,
				"test":   ctx.Invoke("user.Test"),
			}, ark.OK
		},
	})

	ark.Register("user.Test", ark.Method{
		Name: "测试", Desc: "测试",
		Action: func(ctx *ark.Service) Map {

			bll := ctx.Logic("test", Map{
				"msg": "msg from setting",
			})

			return bll.Invoke("GetMsg")
		},
	})

	ark.Register("user.Notify", ark.Method{
		Name: "用户通知", Desc: "用户通知",
		Action: func(ctx *ark.Service) *Res {
			ark.Debug("收到用户通知事件", ctx.Value)

			return ark.OK
		},
	})

	ark.Register("user.Update", ark.Method{
		Name: "用户更新", Desc: "用户更新",
		Action: func(ctx *ark.Service) *Res {
			ark.Debug("收到用户更新队列", ctx.Value)
			return ark.OK
		},
	})

}
