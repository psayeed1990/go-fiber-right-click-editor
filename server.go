package main

import (
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)
func main() {
  app := fiber.New()

  app.Static("/", filepath.Base("public"))

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendFile(filepath.FromSlash("index.html"))
  })

  app.Get("/iframe", func(c *fiber.Ctx) error {
    return c.SendFile(filepath.FromSlash("iframe.html"))
  })

  app.Listen(":3000")
}