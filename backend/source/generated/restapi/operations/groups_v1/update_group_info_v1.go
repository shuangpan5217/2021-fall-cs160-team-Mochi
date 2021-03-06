// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateGroupInfoV1HandlerFunc turns a function with the right signature into a update group info v1 handler
type UpdateGroupInfoV1HandlerFunc func(UpdateGroupInfoV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateGroupInfoV1HandlerFunc) Handle(params UpdateGroupInfoV1Params) middleware.Responder {
	return fn(params)
}

// UpdateGroupInfoV1Handler interface for that can handle valid update group info v1 params
type UpdateGroupInfoV1Handler interface {
	Handle(UpdateGroupInfoV1Params) middleware.Responder
}

// NewUpdateGroupInfoV1 creates a new http.Handler for the update group info v1 operation
func NewUpdateGroupInfoV1(ctx *middleware.Context, handler UpdateGroupInfoV1Handler) *UpdateGroupInfoV1 {
	return &UpdateGroupInfoV1{Context: ctx, Handler: handler}
}

/* UpdateGroupInfoV1 swagger:route PATCH /v1/groups/{group_id} groupsV1 updateGroupInfoV1

update a group info

*/
type UpdateGroupInfoV1 struct {
	Context *middleware.Context
	Handler UpdateGroupInfoV1Handler
}

func (o *UpdateGroupInfoV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateGroupInfoV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
