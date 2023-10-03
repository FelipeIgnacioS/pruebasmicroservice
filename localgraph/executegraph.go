package localgraph

import (
	"github.com/graphql-go/graphql"
)

// ExecuteQueryParams define los parámetros para ejecutar una consulta GraphQL.
type ExecuteQueryParams struct {
	Schema        graphql.Schema
	RequestString string
}

func ExecuteQuery(params ExecuteQueryParams) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        params.Schema,
		RequestString: params.RequestString,
	})

	return result
}
