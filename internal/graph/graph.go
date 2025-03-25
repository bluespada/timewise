// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains code for GraphQL
package graph

import (
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

// Gofiber GraphQL Handler
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

	// Do GraphQL
	result := graphql.Do(graphql.Params{
		Schema:         Schema,
		RequestString:  input.Query,
		OperationName:  input.OperationName,
		VariableValues: input.Variables,
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
