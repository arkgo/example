package filter

import (
	"time"

	"github.com/arkgo/ark"
)

func init() {

	ark.Register("*.logger", ark.RequestFilter{
		Name: "日志拦截器", Desc: "日志拦截器",
		Action: func(ctx *ark.Http) {
			begin := time.Now()

			ctx.Next()

			logger := true
			if vv, ok := ctx.Config["logger"].(bool); ok {
				logger = vv
			}

			if logger {
				ark.Trace(ctx.Ip(), ctx.Id, ctx.Method, ctx.Site, ctx.Uri, ctx.Code, time.Now().Sub(begin))
			}
		},
	})

}
