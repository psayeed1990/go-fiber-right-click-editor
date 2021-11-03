package main

import (
	"io/ioutil"

	"path/filepath"

	"log"

	"strings"

	"github.com/gofiber/fiber/v2"
)
func main() {

	app := fiber.New()

	app.Static("/", filepath.Base("public"))

	app.Get("/", func(c *fiber.Ctx) error {
    	return c.SendFile("/index.html")
  	})


	app.Post("/change-iframe", func(c *fiber.Ctx) error {
		//payload
		payload := struct {
			Content string `json:"content"`
		}{}
			
		//parse body
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		//read iframe file
		input, err := ioutil.ReadFile(filepath.FromSlash("public/iframe.html"))
		
		//log error
		if err != nil {
			log.Fatalln(err)
		}

		//get all lines in a array
		lines := strings.Split(string(input), "\n")

		//check text
        for i, line := range lines {
                if strings.Contains(line, "<h1>") {
                        lines[i] = "<h1>" + payload.Content + "</h1>"
                }
        }
        output := strings.Join(lines, "\n")
        err = ioutil.WriteFile(filepath.FromSlash("public/iframe.html"), []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
        }

		return c.JSON(payload)

  	})

  	app.Listen(":3000")
}