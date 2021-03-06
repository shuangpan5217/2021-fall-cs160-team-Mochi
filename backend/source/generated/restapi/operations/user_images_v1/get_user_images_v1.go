// Code generated by go-swagger; DO NOT EDIT.

package user_images_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserImagesV1HandlerFunc turns a function with the right signature into a get user images v1 handler
type GetUserImagesV1HandlerFunc func(GetUserImagesV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserImagesV1HandlerFunc) Handle(params GetUserImagesV1Params) middleware.Responder {
	return fn(params)
}

// GetUserImagesV1Handler interface for that can handle valid get user images v1 params
type GetUserImagesV1Handler interface {
	Handle(GetUserImagesV1Params) middleware.Responder
}

// NewGetUserImagesV1 creates a new http.Handler for the get user images v1 operation
func NewGetUserImagesV1(ctx *middleware.Context, handler GetUserImagesV1Handler) *GetUserImagesV1 {
	return &GetUserImagesV1{Context: ctx, Handler: handler}
}

/* GetUserImagesV1 swagger:route GET /v1/images userImagesV1 getUserImagesV1

get a user image

get a user image

*/
type GetUserImagesV1 struct {
	Context *middleware.Context
	Handler GetUserImagesV1Handler
}

func (o *GetUserImagesV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserImagesV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
