package site

import (
	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
)

func init() {

	ark.Method("user.Serial", Map{
		"name": "生成序列", "text": "生成序列",
		"action": func(ctx *ark.Service) (Map,*Res) {
			serial := ark.CacheSerial("hahha", 0, 1)
			return Map{
				"serial": serial,
				"test": ctx.Invoke("user.Test"),
			}, ark.OK
		},
	})

	ark.Method("user.Test", Map{
		"name": "测试", "text": "测试",
		"action": func(ctx *ark.Service) (Map) {

			bll := ctx.Logic("test", Map{
				"msg": "msg from setting",
			})

			return bll.Invoke("GetMsg")
		},
	})



	ark.Method("user.Notify", Map{
		"name": "用户通知", "text": "用户通知",
		"action": func(ctx *ark.Service) (*Res) {
			ark.Debug("收到用户通知事件", ctx.Value)
			
			return ark.OK
		},
	})



	ark.Method("user.Update", Map{
		"name": "用户更新", "text": "用户更新",
		"action": func(ctx *ark.Service) (*Res) {
			ark.Debug("收到用户更新队列", ctx.Value)
			return ark.OK
		},
	})
}
