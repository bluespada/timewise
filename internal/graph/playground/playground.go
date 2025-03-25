// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains code for GraphQL Playground
package playground

import (
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// HandlerGraphQLPlayground returns a new Fiber handler that serves GraphQL Playground.
// It requires a GraphQL schema as argument.
//
// The handler will return a HTML page with a form to input GraphQL query and
// variables. When the form is submitted, the handler will execute the query
// against the given schema and return the result in JSON format.
//
// The handler will also serve the static files for GraphQL Playground, which are
// embedded in the binary using the //go:embed directive.
func HandlerGraphQLPlayground(schema graphql.Schema) fiber.Handler {
	return func(c *fiber.Ctx) error {

		_ = handler.New(&handler.Config{
			Schema: &schema,
			Pretty: true,
		})

		html, err := GraphFsStatic.ReadFile("static/index.html")
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error: " + err.Error())
		}

		c.Set("Content-Type", "text/html")
		return c.Send(html)
	}
}
