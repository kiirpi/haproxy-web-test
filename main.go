package main

import (
	"fmt"
	"log"

	"github.com/namsral/flag"

	"github.com/gofiber/fiber/v2"
)

func main() {

	statusPtr := flag.Int("status", 200, "")
	portPtr := flag.Int("port", 3000, "")
	flag.Parse()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		q := c.Query("q")
		c.SendStatus(*statusPtr)
		resp := fmt.Sprintf("status:%d port:%d q:%s", *statusPtr, *portPtr, q)
		log.Println(resp)
		return c.SendString(resp)
	})

	app.Listen(fmt.Sprintf(":%d", *portPtr))
}
