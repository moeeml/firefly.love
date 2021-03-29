package lib

import "github.com/kataras/iris/v12"

func Home(c iris.Context) {
	c.View("home.tpl")
}
