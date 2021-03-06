// Code generated by go-swagger; DO NOT EDIT.

package user_images_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// GetUserImagesV1OKCode is the HTTP code returned for type GetUserImagesV1OK
const GetUserImagesV1OKCode int = 200

/*GetUserImagesV1OK Success

swagger:response getUserImagesV1OK
*/
type GetUserImagesV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.UserImagesResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1OK creates GetUserImagesV1OK with default headers values
func NewGetUserImagesV1OK() *GetUserImagesV1OK {

	return &GetUserImagesV1OK{}
}

// WithPayload adds the payload to the get user images v1 o k response
func (o *GetUserImagesV1OK) WithPayload(payload *models.UserImagesResponse) *GetUserImagesV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 o k response
func (o *GetUserImagesV1OK) SetPayload(payload *models.UserImagesResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1BadRequestCode is the HTTP code returned for type GetUserImagesV1BadRequest
const GetUserImagesV1BadRequestCode int = 400

/*GetUserImagesV1BadRequest Bad Request

swagger:response getUserImagesV1BadRequest
*/
type GetUserImagesV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1BadRequest creates GetUserImagesV1BadRequest with default headers values
func NewGetUserImagesV1BadRequest() *GetUserImagesV1BadRequest {

	return &GetUserImagesV1BadRequest{}
}

// WithPayload adds the payload to the get user images v1 bad request response
func (o *GetUserImagesV1BadRequest) WithPayload(payload *models.ErrResponse) *GetUserImagesV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 bad request response
func (o *GetUserImagesV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1UnauthorizedCode is the HTTP code returned for type GetUserImagesV1Unauthorized
const GetUserImagesV1UnauthorizedCode int = 401

/*GetUserImagesV1Unauthorized Unauthorized

swagger:response getUserImagesV1Unauthorized
*/
type GetUserImagesV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1Unauthorized creates GetUserImagesV1Unauthorized with default headers values
func NewGetUserImagesV1Unauthorized() *GetUserImagesV1Unauthorized {

	return &GetUserImagesV1Unauthorized{}
}

// WithPayload adds the payload to the get user images v1 unauthorized response
func (o *GetUserImagesV1Unauthorized) WithPayload(payload *models.ErrResponse) *GetUserImagesV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 unauthorized response
func (o *GetUserImagesV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1ForbiddenCode is the HTTP code returned for type GetUserImagesV1Forbidden
const GetUserImagesV1ForbiddenCode int = 403

/*GetUserImagesV1Forbidden Forbidden

swagger:response getUserImagesV1Forbidden
*/
type GetUserImagesV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1Forbidden creates GetUserImagesV1Forbidden with default headers values
func NewGetUserImagesV1Forbidden() *GetUserImagesV1Forbidden {

	return &GetUserImagesV1Forbidden{}
}

// WithPayload adds the payload to the get user images v1 forbidden response
func (o *GetUserImagesV1Forbidden) WithPayload(payload *models.ErrResponse) *GetUserImagesV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 forbidden response
func (o *GetUserImagesV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1NotFoundCode is the HTTP code returned for type GetUserImagesV1NotFound
const GetUserImagesV1NotFoundCode int = 404

/*GetUserImagesV1NotFound Not Found

swagger:response getUserImagesV1NotFound
*/
type GetUserImagesV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1NotFound creates GetUserImagesV1NotFound with default headers values
func NewGetUserImagesV1NotFound() *GetUserImagesV1NotFound {

	return &GetUserImagesV1NotFound{}
}

// WithPayload adds the payload to the get user images v1 not found response
func (o *GetUserImagesV1NotFound) WithPayload(payload *models.ErrResponse) *GetUserImagesV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 not found response
func (o *GetUserImagesV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1ConflictCode is the HTTP code returned for type GetUserImagesV1Conflict
const GetUserImagesV1ConflictCode int = 409

/*GetUserImagesV1Conflict Conflict

swagger:response getUserImagesV1Conflict
*/
type GetUserImagesV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1Conflict creates GetUserImagesV1Conflict with default headers values
func NewGetUserImagesV1Conflict() *GetUserImagesV1Conflict {

	return &GetUserImagesV1Conflict{}
}

// WithPayload adds the payload to the get user images v1 conflict response
func (o *GetUserImagesV1Conflict) WithPayload(payload *models.ErrResponse) *GetUserImagesV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 conflict response
func (o *GetUserImagesV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserImagesV1InternalServerErrorCode is the HTTP code returned for type GetUserImagesV1InternalServerError
const GetUserImagesV1InternalServerErrorCode int = 500

/*GetUserImagesV1InternalServerError Internal Server Error

swagger:response getUserImagesV1InternalServerError
*/
type GetUserImagesV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetUserImagesV1InternalServerError creates GetUserImagesV1InternalServerError with default headers values
func NewGetUserImagesV1InternalServerError() *GetUserImagesV1InternalServerError {

	return &GetUserImagesV1InternalServerError{}
}

// WithPayload adds the payload to the get user images v1 internal server error response
func (o *GetUserImagesV1InternalServerError) WithPayload(payload *models.ErrResponse) *GetUserImagesV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user images v1 internal server error response
func (o *GetUserImagesV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserImagesV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
