// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewGetFileV1Params creates a new GetFileV1Params object
//
// There are no default values defined in the spec.
func NewGetFileV1Params() GetFileV1Params {

	return GetFileV1Params{}
}

// GetFileV1Params contains all the bound params for the get file v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters getFileV1
type GetFileV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Bearer token based Authorization
	  Required: true
	  In: header
	*/
	Authorization string
	/*file path
	  Required: true
	  In: path
	*/
	Path string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFileV1Params() beforehand.
func (o *GetFileV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	rPath, rhkPath, _ := route.Params.GetOK("path")
	if err := o.bindPath(rPath, rhkPath, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *GetFileV1Params) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("Authorization", "header", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true

	if err := validate.RequiredString("Authorization", "header", raw); err != nil {
		return err
	}
	o.Authorization = raw

	return nil
}

// bindPath binds and validates parameter Path from path.
func (o *GetFileV1Params) bindPath(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route
	o.Path = raw

	return nil
}
