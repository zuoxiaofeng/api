package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/zuoxiaofeng/api/methods"
)

func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome</h1>")
	})

	app.Get("/zxf/v1/desc", func(ctx iris.Context) {
		ctx.Text("My web api")
	})

	app.Post("/add", methods.Add)

	app.Run(iris.Addr(":9090"))
	iris.WithoutServerError(iris.ErrServerClosed)
}
