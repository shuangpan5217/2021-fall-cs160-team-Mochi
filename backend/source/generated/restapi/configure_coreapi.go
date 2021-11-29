// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-chi/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/jinzhu/gorm"

	"2021-fall-cs160-team-Mochi/backend/source/apis/comments"
	"2021-fall-cs160-team-Mochi/backend/source/apis/commonutils"
	"2021-fall-cs160-team-Mochi/backend/source/apis/friends"
	"2021-fall-cs160-team-Mochi/backend/source/apis/groups"
	"2021-fall-cs160-team-Mochi/backend/source/apis/notes"
	"2021-fall-cs160-team-Mochi/backend/source/apis/userImages"
	"2021-fall-cs160-team-Mochi/backend/source/apis/usermgmt"
	"2021-fall-cs160-team-Mochi/backend/source/generated/restapi/operations"
)

var db *gorm.DB

//go:generate swagger generate server --target ../../generated --name Coreapi --spec ../../swagger-specs/api.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.CoreapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CoreapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	db = CreateDBConnection()

	// user management
	api.UserMgmtV1LoginV1Handler = usermgmt.LoginV1Handler(db)
	api.UserMgmtV1GetUserV1Handler = usermgmt.GetUserV1Handler(db)
	api.UserMgmtV1UpdatePasswordV1Handler = usermgmt.UpdatePasswordV1Handler(db)
	api.UserMgmtV1UpdateUserInfoV1Handler = usermgmt.UpdateUserInfoV1Handler(db)
	api.UserMgmtV1SearchUserV1Handler = usermgmt.SearchUserV1Handler(db)

	// friends
	api.FriendsV1GetFriendsV1Handler = friends.GetFriendsV1Handler(db)
	api.FriendsV1RemoveFriendsV1Handler = friends.RemoveFriendsV1Handler(db)
	api.FriendsV1AddFriendsV1Handler = friends.AddFriendsV1Handler(db)

	// notes
	api.NotesV1PostFileV1Handler = notes.PostFileV1Handler(db)
	api.NotesV1GetFileV1Handler = notes.GetFileV1Handler(db)
	api.NotesV1GetMultipleFilesV1Handler = notes.GetMultipleFilesV1Handler(db)
	api.NotesV1UploadNoteV1Handler = notes.UploadNoteV1Handler(db)
	api.NotesV1FindNotesByUsernameHandler = notes.GetNotesByUsernameHandler(db)
	api.NotesV1GetNoteCommentsHandler = notes.GetNoteCommentsHandler(db)
	api.NotesV1GetSingleNoteV1Handler = notes.GetSingleNoteV1Handler(db)
	api.NotesV1DeleteNoteV1Handler = notes.DeleteNoteByIdV1Handler(db)
	api.NotesV1GetNoteMembersV1Handler = notes.GetNoteMembersV1Handler(db)
	api.NotesV1AddNoteMembersV1Handler = notes.AddNoteMembersV1Handler(db)
	api.NotesV1RemoveNoteMembersV1Handler = notes.RemoveNoteMembersV1Handler(db)
	api.NotesV1GetGroupNotesV1Handler = notes.GetGroupNotesV1Handler(db)
	api.NotesV1FindNotesByTagsHandler = notes.FindNoteByTagV1Handler(db)

	// comments
	api.CommentsV1PostCommentsV1Handler = comments.PostCommentsV1Handler(db)
	api.CommentsV1RemoveComnentV1Handler = comments.RemoveComnentV1Handler(db)

	// Groups
	api.GroupsV1CreateGroupV1Handler = groups.CreateGroupV1Handler(db)
	api.GroupsV1GetGroupsV1Handler = groups.GetGroupsV1Handler(db)
	api.GroupsV1GetGroupInfoV1Handler = groups.GetGroupInfoV1Handler(db)
	api.GroupsV1AddGroupUsersV1Handler = groups.AddGroupMembersHandler(db)
	api.GroupsV1DeleteGroupV1Handler = groups.DeleteGroupByIdV1Handler(db)
	api.GroupsV1RemoveGroupUsersV1Handler = groups.RemoveGroupMembersHandler(db)
	api.GroupsV1GetGroupUsersV1Handler = groups.GetGroupMembersHandler(db)

	// user image
	api.UserImagesV1PostUserImagesV1Handler = userImages.UploadUserImagesHandlerV1(db)
	api.UserImagesV1GetUserImagesV1Handler = userImages.GetUserImagesHandlerV1(db)
	api.UserImagesV1GetMultipleUserImagesV1Handler = userImages.GetMultipleUserImages(db)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	corsHandler := cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "PATCH", "PUT", "OPTIONS", "DELETE", "HEAD"},
	})

	return corsHandler(handler)
}

func CreateDBConnection() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=shuangpan user=shuangpan sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
	}
	db.LogMode(true)
	err = commonutils.InsertTables(db)
	if err != nil {
		log.Fatal(err.Error())
	}
	commonutils.AddFKConstraints(db)
	return db
}

func CloseDBConnection() {
	err := db.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}
