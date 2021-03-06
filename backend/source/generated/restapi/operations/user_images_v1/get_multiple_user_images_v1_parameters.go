// Code generated by go-swagger; DO NOT EDIT.

package user_images_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// NewGetMultipleUserImagesV1Params creates a new GetMultipleUserImagesV1Params object
//
// There are no default values defined in the spec.
func NewGetMultipleUserImagesV1Params() GetMultipleUserImagesV1Params {

	return GetMultipleUserImagesV1Params{}
}

// GetMultipleUserImagesV1Params contains all the bound params for the get multiple user images v1 operation
// typically these are obtained from a http.Request
//
// swagger:parameters getMultipleUserImagesV1
type GetMultipleUserImagesV1Params struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Bearer token based Authorization
	  Required: true
	  In: header
	*/
	Authorization string
	/*array of users
	  In: body
	*/
	Body *models.GroupMembers
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetMultipleUserImagesV1Params() beforehand.
func (o *GetMultipleUserImagesV1Params) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := o.bindAuthorization(r.Header[http.CanonicalHeaderKey("Authorization")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.GroupMembers
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAuthorization binds and validates parameter Authorization from header.
func (o *GetMultipleUserImagesV1Params) bindAuthorization(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
