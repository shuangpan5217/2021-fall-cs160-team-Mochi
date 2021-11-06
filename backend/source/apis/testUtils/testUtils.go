package testUtils

import (
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi"
	"log"
	"net/http/httptest"
	"os"

	"github.com/jinzhu/gorm"
)

func SetupTestServer() (db *gorm.DB, testServer *httptest.Server) {
	os.Setenv("TESTING", "true")
	apiHandler, err := restapi.GetAPIHandler()
	if err != nil {
		log.Fatal(err.Error())
	}
	testServer = httptest.NewServer(apiHandler)
	db = restapi.CreateDBConnection()
	return
}

func ShutDownTestServer(db *gorm.DB, testServer *httptest.Server) {
	db.Close()
	testServer.Close()
}
