package user

import (
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/base"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	ark.Register("user", ark.Table{
		Name: "用户", Desc: "用户",
		Fields: ark.Params{
			"id": ark.Param{
				Type: "id", Require: false, Name: "编号",
			},
			"type": ark.Param{
				Type: "enum", Require: true, Name: "类型", Default: USER_CUSTOMER,
				Option: ark.Option{
					USER_PLATFORM: "平台",
					USER_CUSTOMER: "客户",
				},
			},

			"name": ark.Param{
				Type: "string", Require: true, Name: "昵称",
			},

			"setting": ark.Param{
				Type: "json", Require: true, Name: "类型",
				Children: ark.Params{
					"showad": ark.Param{Type: "bool", Require: true, Name: "是否显示广告", Default: false},
				},
			},

			"status": ark.Param{
				Type: "enum", Require: false, Name: "状态",
				Option: ark.Option{
					REMOVED: "已删除",
				},
			},
			"changed": ark.Param{
				Type: "datetime", Require: true, Name: "修改时间", Default: time.Now,
			},
			"created": ark.Param{
				Type: "datetime", Require: true, Name: "创建时间", Default: time.Now,
			},
		},
	})

	ark.Register("user.search", ark.View{
		Name: "用户搜索视图", Desc: "用户搜索视图",
		Fields: ark.Params{
			"id": ark.Param{
				Type: "id", Require: false, Name: "编号",
			},
			"type": ark.Param{
				Type: "enum", Require: true, Name: "类型", Default: USER_CUSTOMER,
				Option: ark.Option{
					USER_PLATFORM: "平台",
					USER_CUSTOMER: "客户",
				},
			},
			"name": ark.Param{
				Type: "string", Require: true, Name: "昵称",
			},
		},
	})

	ark.Register("user.test", ark.Model{
		Name: "用户测试模型", Desc: "用户测试模型",
		Fields: ark.Params{
			"name": ark.Param{
				Type: "string", Require: true, Name: "名称",
			},
			"remark": ark.Param{
				Type: "string", Require: false, Name: "备注",
			},
		},
	})
}
