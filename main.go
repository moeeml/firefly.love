package main

import (
	"firefly.love/lib"
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	tpl := iris.HTML("./resource/view", ".tpl")
	app.RegisterView(tpl)

	app.Favicon("./public/favicon.ico")
	app.HandleDir("/public", iris.Dir("./public"))

	app.Get("/", lib.Home)
	app.Get("/line", lib.CreateWebsocket())

	app.Listen(":8000")
}
