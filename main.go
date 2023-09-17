package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"imooc-iris/web/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./web/views", ".html"))

	// controller register
	mvc.New(app.Party("/hello")).Handle(new(controllers.MovieController))
	app.Run(iris.Addr("localhost:8080"))
}
