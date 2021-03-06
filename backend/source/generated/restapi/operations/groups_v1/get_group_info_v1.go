// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetGroupInfoV1HandlerFunc turns a function with the right signature into a get group info v1 handler
type GetGroupInfoV1HandlerFunc func(GetGroupInfoV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGroupInfoV1HandlerFunc) Handle(params GetGroupInfoV1Params) middleware.Responder {
	return fn(params)
}

// GetGroupInfoV1Handler interface for that can handle valid get group info v1 params
type GetGroupInfoV1Handler interface {
	Handle(GetGroupInfoV1Params) middleware.Responder
}

// NewGetGroupInfoV1 creates a new http.Handler for the get group info v1 operation
func NewGetGroupInfoV1(ctx *middleware.Context, handler GetGroupInfoV1Handler) *GetGroupInfoV1 {
	return &GetGroupInfoV1{Context: ctx, Handler: handler}
}

/* GetGroupInfoV1 swagger:route GET /v1/groups/{group_id} groupsV1 getGroupInfoV1

get a group info

*/
type GetGroupInfoV1 struct {
	Context *middleware.Context
	Handler GetGroupInfoV1Handler
}

func (o *GetGroupInfoV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetGroupInfoV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
