package controllerGraph

import (
	"BACKEND-MICROSERVICE-AUTHENTICATION/localgraph"

	"github.com/gofiber/fiber/v2"
)

func GraphQLHandler(c *fiber.Ctx) error {
	var params localgraph.ExecuteQueryParams
	params.Schema = localgraph.Schema
	params.RequestString = c.FormValue("query")

	result := localgraph.ExecuteQuery(params)

	if len(result.Errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(result)
	}

	return c.JSON(result)
}
