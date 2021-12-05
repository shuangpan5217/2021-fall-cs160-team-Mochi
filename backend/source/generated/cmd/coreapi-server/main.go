package main

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations"
	"flag"
	"log"

	"github.com/go-openapi/loads"
	"github.com/spf13/cast"
)

var portFlag = flag.String("port", "3000", "Port to run this service on")
var hostFlag = flag.String("host", "0.0.0.0", "Host to run this service")

func main() {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server
	api := operations.NewCoreapiAPI(swaggerSpec)
	server = restapi.NewServer(api)

	server.Port = cast.ToInt(*portFlag)
	server.Host = *hostFlag
	server.EnabledListeners = []string{"http"}
	flag.Parse()

	defer server.Shutdown()
	defer restapi.CloseDBConnection()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
