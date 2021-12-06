package main

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations"
	"flag"
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/spf13/cast"
)

const (
	localhost  = "127.0.0.1"
	dockerhost = "0.0.0.0"
)

var (
	portFlag *string
	hostFlag *string
)

func init() {
	if os.Getenv("ENV") == "production" {
		hostFlag = flag.String("host", dockerhost, "Host to run this service")
	} else {
		hostFlag = flag.String("host", localhost, "Host to run this service")
	}
	portFlag = flag.String("port", "3001", "Port to run this service on")
}

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
