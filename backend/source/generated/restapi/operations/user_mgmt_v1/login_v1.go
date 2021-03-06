// Code generated by go-swagger; DO NOT EDIT.

package user_mgmt_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// LoginV1HandlerFunc turns a function with the right signature into a login v1 handler
type LoginV1HandlerFunc func(LoginV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn LoginV1HandlerFunc) Handle(params LoginV1Params) middleware.Responder {
	return fn(params)
}

// LoginV1Handler interface for that can handle valid login v1 params
type LoginV1Handler interface {
	Handle(LoginV1Params) middleware.Responder
}

// NewLoginV1 creates a new http.Handler for the login v1 operation
func NewLoginV1(ctx *middleware.Context, handler LoginV1Handler) *LoginV1 {
	return &LoginV1{Context: ctx, Handler: handler}
}

/* LoginV1 swagger:route POST /v1/login UserMgmtV1 loginV1

Sign up or log in

handle login request, username and password

*/
type LoginV1 struct {
	Context *middleware.Context
	Handler LoginV1Handler
}

func (o *LoginV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewLoginV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
