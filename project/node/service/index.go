package service

import (
	"github.com/arkgo/ark"
)

func init() {

	ark.Register("user.Login", ark.Service{
		Name: "用户登录服务", Desc: "用户登录服务",
	})

}
