package main

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations"
	"flag"
	"log"

	"github.com/go-openapi/loads"
)

var portFlag = flag.Int("port", 3000, "Port to run this service on")

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server
	api := operations.NewCoreapiAPI(swaggerSpec)
	server = restapi.NewServer(api)

	server.Port = *portFlag
	server.EnabledListeners = []string{"http"}
	flag.Parse()

	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
