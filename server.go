package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Favicon("./public/favicon.ico")
	app.Use(iris.Gzip)
	app.StaticWeb("/", "./public")
	app.Run(iris.Addr(":9080"))
}

