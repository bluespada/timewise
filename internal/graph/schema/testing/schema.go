// Copyright (c) 2025 Bluespada <pentingmain@gmail.com>
//
// This software is licensed under MIT License, please read accompany file copy
// or read online at https://opensource.org/license/mit
//
// This file contains testing code schema for GraphQL
package testing

import (
	"fmt"

	"github.com/bluespada/timewise/internal/graph/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

var testHello = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		context := p.Context.Value("context").(*fiber.Ctx)
		claims, err := middleware.ValidateJWT(context)
		if err != nil {
			return nil, err
		}
		return fmt.Sprintf("%f", claims["user"].(float64)), nil
	},
}

var testGreeting = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"name": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		name := p.Args["name"].(string)
		return "Hi " + name, nil
	},
}

var mutationSend = &graphql.Field{
	Type: graphql.String,
	Args: graphql.FieldConfigArgument{
		"message": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		message := p.Args["message"].(string)
		return "Message: " + message, nil
	},
}

var TestQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "TestQuery",
	Fields: graphql.Fields{
		"hello":    testHello,
		"greeting": testGreeting,
	},
})

var TestMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "TestMutation",
	Fields: graphql.Fields{
		"send": mutationSend,
	},
})
