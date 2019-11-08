package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
)

var Greeting string = os.Getenv("GREETING")

func main() {
	if Greeting == "" {
		Greeting = "Hello"
	}

	app := iris.Default()
	app.Get("/hello/{name}", hello)

	app.Run(iris.Addr(":80"))
}

type Response struct {
	Message string `json:"message"`
}

func hello(ctx iris.Context) {
	name := ctx.Params().Get("name")
	ctx.JSON(iris.Map{
		"message": fmt.Sprintf("%s %s", Greeting, name),
	})
}
