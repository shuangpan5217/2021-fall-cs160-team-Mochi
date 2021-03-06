// Code generated by go-swagger; DO NOT EDIT.

package friends_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetFriendsV1HandlerFunc turns a function with the right signature into a get friends v1 handler
type GetFriendsV1HandlerFunc func(GetFriendsV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFriendsV1HandlerFunc) Handle(params GetFriendsV1Params) middleware.Responder {
	return fn(params)
}

// GetFriendsV1Handler interface for that can handle valid get friends v1 params
type GetFriendsV1Handler interface {
	Handle(GetFriendsV1Params) middleware.Responder
}

// NewGetFriendsV1 creates a new http.Handler for the get friends v1 operation
func NewGetFriendsV1(ctx *middleware.Context, handler GetFriendsV1Handler) *GetFriendsV1 {
	return &GetFriendsV1{Context: ctx, Handler: handler}
}

/* GetFriendsV1 swagger:route GET /v1/friends friendsV1 getFriendsV1

get friends info

*/
type GetFriendsV1 struct {
	Context *middleware.Context
	Handler GetFriendsV1Handler
}

func (o *GetFriendsV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetFriendsV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
