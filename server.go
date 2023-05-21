package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
    "chat/pkg/common"
    "log"
    "fmt"
)

func main() {
    
    c, err := config.LoadConfig()
    if err != nil {
        log.Fatalln("Failed at config", err)
    }

    app := fiber.New()
	app.Use(cors.New())

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is not an api route, this is the root route. You must have a problem")
    })

    fmt.Print("Port: ", c.Port)

    if err := app.Listen(c.Port); err != nil {
        fmt.Println(err)
    }
}