package testing

import "github.com/graphql-go/graphql"

var testHello = &graphql.Field{
	Type: graphql.String,
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return "World", nil
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
