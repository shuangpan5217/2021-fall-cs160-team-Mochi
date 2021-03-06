// Code generated by go-swagger; DO NOT EDIT.

package user_mgmt_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SearchUserV1HandlerFunc turns a function with the right signature into a search user v1 handler
type SearchUserV1HandlerFunc func(SearchUserV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn SearchUserV1HandlerFunc) Handle(params SearchUserV1Params) middleware.Responder {
	return fn(params)
}

// SearchUserV1Handler interface for that can handle valid search user v1 params
type SearchUserV1Handler interface {
	Handle(SearchUserV1Params) middleware.Responder
}

// NewSearchUserV1 creates a new http.Handler for the search user v1 operation
func NewSearchUserV1(ctx *middleware.Context, handler SearchUserV1Handler) *SearchUserV1 {
	return &SearchUserV1{Context: ctx, Handler: handler}
}

/* SearchUserV1 swagger:route GET /v1/user/username/{username} UserMgmtV1 searchUserV1

search by username

search by username

*/
type SearchUserV1 struct {
	Context *middleware.Context
	Handler SearchUserV1Handler
}

func (o *SearchUserV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSearchUserV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
