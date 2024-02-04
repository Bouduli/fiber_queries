package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Index")
	})

	app.Get("/q", func(c *fiber.Ctx) error {
		// Individual queries are accessed with c.Query
		var q1 = c.Query("q1")
		var q2 = c.Query("q2")
		fmt.Printf("q1: %v, q2: %v", q1, q2)

		var queries string = ""
		//Iterating over Queries with range syntax.
		for k, v := range c.Queries() {
			queries += k + "=" + v + "\r\n"
		}
		if len(queries) == 0 {
			queries = "no queries"
		}
		return c.SendString(queries)
	})

	app.Get("/p/:p1", func(c *fiber.Ctx) error {

		//Individual params are accessed through c.Params().
		var p1 = c.Params("p1")
		fmt.Printf("p1: %v", p1)

		params := ""
		//Params can be iterated over through with range.
		for k, v := range c.AllParams() {
			params += k + "=" + v + "\r\n"
		}
		if len(params) == 0 {
			params = "no params"
		}

		return c.SendString(params)
	})

	app.Listen(":3000")
}
