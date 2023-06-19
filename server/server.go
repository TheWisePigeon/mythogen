package server

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
)

func Serve( port *string ){
    mythogen := fiber.New()
    fmt.Printf("Mythogen server launched on port %s", *port)
    mythogen.Get("/test", func(c *fiber.Ctx) error {
        return c.SendString("Hello from Mythogen server")
    })
    err := mythogen.Listen(fmt.Sprintf(":%s", *port))
    if err != nil {
        panic(err)
    }
}
