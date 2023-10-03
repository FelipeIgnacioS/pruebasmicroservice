package localgraph

import (
	"BACKEND-MICROSERVICE-AUTHENTICATION/controllers"

	"github.com/graphql-go/graphql"
)

func resolveUser(params graphql.ResolveParams) (interface{}, error) {
	// Obtener el ID del usuario de los argumentos de la consulta
	userID := params.Args["id"].(string)

	// Invocar la función del controlador para obtener el usuario por ID
	user, err := controllers.GetUser(userID)
	if err != nil {
		return nil, err
	}

	// Devolver el usuario obtenido
	return user, nil
}

func resolveRegisterUser(params graphql.ResolveParams) (interface{}, error) {
	// Obtener los datos del usuario de los argumentos de la mutación
	body := params.Args["input"].(map[string]interface{})["body"].(string)

	// Invocar la función del controlador para registrar un nuevo usuario
	user, err := controllers.SingUp([]byte(body))
	if err != nil {
		return nil, err
	}

	// Devolver una estructura de respuesta de éxito
	return user, nil
}

func resolveLoginUser(params graphql.ResolveParams) (interface{}, error) {
	// Obtener el cuerpo (body) JSON del usuario de los argumentos de la mutación
	body := params.Args["input"].(map[string]interface{})["body"].(string)

	// Invocar la función del controlador para iniciar sesión del usuario
	user, err := controllers.Login([]byte(body))
	if err != nil {
		return nil, err
	}

	// Devolver una estructura de respuesta de éxito con el token
	return user, nil
}
