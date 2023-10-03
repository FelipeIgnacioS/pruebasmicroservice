// schema.go
package localgraph

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.NewNonNull(graphql.ID)},
			"name":     &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"lastName": &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			"email":    &graphql.Field{Type: graphql.NewNonNull(graphql.String)},
			// No exponer campode contraseña
		},
	},
)

var userResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "UserResponse",
		Fields: graphql.Fields{
			"success": &graphql.Field{
				Type: graphql.Boolean,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"user": &graphql.Field{
				Type: userType, // Usar userType aquí
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var userInputType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "UserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"body": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	},
)

var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: resolveUser,
			},
		},
	},
)

var RootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"registerUser": &graphql.Field{
				Type: userResponseType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(userInputType)},
				},
				Resolve: resolveRegisterUser,
			},
			"loginUser": &graphql.Field{
				Type: userResponseType,
				Args: graphql.FieldConfigArgument{
					"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: resolveLoginUser,
			},
		},
	},
)

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    RootQuery,
		Mutation: RootMutation,
	},
)
