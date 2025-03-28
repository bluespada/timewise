// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains code for GraphQL
package graph

import (
	"context"
	"log"

	graphtesting "github.com/bluespada/timewise/internal/graph/schema/testing"
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

// define Schema
var Schema graphql.Schema

// define GraphQLRequest type
type GraphQLRequest struct {
	Query         string                 `query:"query"`
	OperationName string                 `query:"operationName"`
	Variables     map[string]interface{} `query:"variables"`
}

// GraphHandler processes incoming HTTP requests and executes GraphQL queries.
// It supports both GET and POST methods. For GET requests, it parses query parameters
// into a GraphQLRequest struct, while for POST requests, it parses the request body.
// After parsing, it executes the GraphQL query using the defined schema and returns
// the result in JSON format. It sets the content type to "application/json" and
// returns an error message with a 500 status code if parsing fails.
func GraphHandler(c *fiber.Ctx) error {
	var input GraphQLRequest

	// curl 'http://localhost:9090/?query=query%7Bhello%7D'
	if c.Method() == fiber.MethodGet {
		if err := c.QueryParser(&input); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Cannot parser query parameters: " + err.Error())
		}
	}

	// curl 'http://localhost:9090/' --header 'content-type: application/json' --data-raw '{"query":"query{hello}"}'
	if c.Method() == fiber.MethodPost {
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Cannot parser query parameters: " + err.Error())
		}
	}
	ctx := context.Background()
	ctx = context.WithValue(ctx, "context", c)
	// Do GraphQL
	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  input.Query,
		OperationName:  input.OperationName,
		VariableValues: input.Variables,
		Context:        ctx,
	})

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(result)
}

// Initialize graphql and registering the function
func init() {
	var err error

	fields := graphql.Fields{
		"test": &graphql.Field{
			Type: graphtesting.TestQuery,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]interface{}{}, nil
			},
		},
	}

	mutation := graphql.Fields{
		"test": &graphql.Field{
			Type: graphtesting.TestMutation,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]interface{}{}, nil
			},
		},
	}

	// creating mutation for query and mutation
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	rootMutation := graphql.ObjectConfig{Name: "RootMutation", Fields: mutation}

	// creating schema.
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query:    graphql.NewObject(rootQuery),
		Mutation: graphql.NewObject(rootMutation),
	})

	if err != nil {
		log.Fatalln("Failed to create schema, error:", err)
	}

}
