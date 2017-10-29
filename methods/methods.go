package methods

import (
	"github.com/kataras/iris"
)

func Add(ctx iris.Context) {
	ctx.Text("add method")
}
