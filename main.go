package main

import "github.com/gofiber/fiber" //cmd: go mod init example/go.mod
//then cmd: go get `this-url`

func main() {
	app := fiber.New()

	app.Get("/", hello) //can no longer code function within get command, has to be own function

	app.Listen(":8000")
}

func hello(c *fiber.Ctx) {
	c.SendString("Hello World")
}
