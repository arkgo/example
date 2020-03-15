package main

import (
	"github.com/arkgo/ark"
	_ "github.com/arkgo/builtin"
	_ "github.com/arkgo/driver"
	_ "github.com/arkgo/example/library"
	_ "github.com/arkgo/example/library/test"
	_ "github.com/arkgo/example/project/model"
	_ "github.com/arkgo/example/project/model/user"
	_ "github.com/arkgo/example/project/site"
	_ "github.com/arkgo/example/project/site/www"
)

func main() {
	ark.Go()
}
