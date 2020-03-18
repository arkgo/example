package user

import (
	"time"

	"github.com/arkgo/ark"
	. "github.com/arkgo/asset"
	. "github.com/arkgo/example/asset/base"
)

func init() {

	ark.Register("user", ark.Table{
		Name: "用户", Desc: "用户",
		Fields: Vars{
			"id": Var{
				Type: "id", Require: false, Name: "编号",
			},
			"type": Var{
				Type: "enum", Require: true, Name: "类型", Default: USER_CUSTOMER,
				Option: Map{
					USER_PLATFORM: "平台",
					USER_CUSTOMER: "客户",
				},
			},

			"name": Var{
				Type: "string", Require: true, Name: "昵称",
			},

			"setting": Var{
				Type: "json", Require: true, Name: "类型",
				Children: Vars{
					"showad": Var{Type: "bool", Require: true, Name: "是否显示广告", Default: false},
				},
			},

			"status": Var{
				Type: "enum", Require: false, Name: "状态",
				Option: Map{
					REMOVED: "已删除",
				},
			},
			"changed": Var{
				Type: "datetime", Require: true, Name: "修改时间", Default: time.Now,
			},
			"created": Var{
				Type: "datetime", Require: true, Name: "创建时间", Default: time.Now,
			},
		},
	})

	ark.Register("user.search", ark.View{
		Name: "用户搜索视图", Desc: "用户搜索视图",
		Fields: Vars{
			"id": Var{
				Type: "id", Require: false, Name: "编号",
			},
			"type": Var{
				Type: "enum", Require: true, Name: "类型", Default: USER_CUSTOMER,
				Option: Map{
					USER_PLATFORM: "平台",
					USER_CUSTOMER: "客户",
				},
			},
			"name": Var{
				Type: "string", Require: true, Name: "昵称",
			},
		},
	})

	ark.Register("user.test", ark.Model{
		Name: "用户测试模型", Desc: "用户测试模型",
		Fields: Vars{
			"name": Var{
				Type: "string", Require: true, Name: "名称",
			},
			"remark": Var{
				Type: "string", Require: false, Name: "备注",
			},
		},
	})
}
