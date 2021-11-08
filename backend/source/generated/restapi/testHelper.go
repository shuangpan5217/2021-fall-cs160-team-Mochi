package restapi

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations"
	"net/http"

	"github.com/go-openapi/loads"
)

func getAPI() (*operations.CoreapiAPI, error) {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		return nil, err
	}
	api := operations.NewCoreapiAPI(swaggerSpec)
	return api, nil
}

func GetAPIHandler() (http.Handler, error) {
	api, err := getAPI()
	if err != nil {
		return nil, err
	}
	handler := configureAPI(api)
	err = api.Validate()
	if err != nil {
		return nil, err
	}
	return handler, nil
}
