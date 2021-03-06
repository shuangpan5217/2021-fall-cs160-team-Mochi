// Code generated by go-swagger; DO NOT EDIT.

package user_mgmt_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// SearchUserV1OKCode is the HTTP code returned for type SearchUserV1OK
const SearchUserV1OKCode int = 200

/*SearchUserV1OK Success

swagger:response searchUserV1OK
*/
type SearchUserV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.UserObj `json:"body,omitempty"`
}

// NewSearchUserV1OK creates SearchUserV1OK with default headers values
func NewSearchUserV1OK() *SearchUserV1OK {

	return &SearchUserV1OK{}
}

// WithPayload adds the payload to the search user v1 o k response
func (o *SearchUserV1OK) WithPayload(payload *models.UserObj) *SearchUserV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 o k response
func (o *SearchUserV1OK) SetPayload(payload *models.UserObj) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1BadRequestCode is the HTTP code returned for type SearchUserV1BadRequest
const SearchUserV1BadRequestCode int = 400

/*SearchUserV1BadRequest Bad Request

swagger:response searchUserV1BadRequest
*/
type SearchUserV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1BadRequest creates SearchUserV1BadRequest with default headers values
func NewSearchUserV1BadRequest() *SearchUserV1BadRequest {

	return &SearchUserV1BadRequest{}
}

// WithPayload adds the payload to the search user v1 bad request response
func (o *SearchUserV1BadRequest) WithPayload(payload *models.ErrResponse) *SearchUserV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 bad request response
func (o *SearchUserV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1UnauthorizedCode is the HTTP code returned for type SearchUserV1Unauthorized
const SearchUserV1UnauthorizedCode int = 401

/*SearchUserV1Unauthorized Unauthorized

swagger:response searchUserV1Unauthorized
*/
type SearchUserV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1Unauthorized creates SearchUserV1Unauthorized with default headers values
func NewSearchUserV1Unauthorized() *SearchUserV1Unauthorized {

	return &SearchUserV1Unauthorized{}
}

// WithPayload adds the payload to the search user v1 unauthorized response
func (o *SearchUserV1Unauthorized) WithPayload(payload *models.ErrResponse) *SearchUserV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 unauthorized response
func (o *SearchUserV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1ForbiddenCode is the HTTP code returned for type SearchUserV1Forbidden
const SearchUserV1ForbiddenCode int = 403

/*SearchUserV1Forbidden Forbidden

swagger:response searchUserV1Forbidden
*/
type SearchUserV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1Forbidden creates SearchUserV1Forbidden with default headers values
func NewSearchUserV1Forbidden() *SearchUserV1Forbidden {

	return &SearchUserV1Forbidden{}
}

// WithPayload adds the payload to the search user v1 forbidden response
func (o *SearchUserV1Forbidden) WithPayload(payload *models.ErrResponse) *SearchUserV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 forbidden response
func (o *SearchUserV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1NotFoundCode is the HTTP code returned for type SearchUserV1NotFound
const SearchUserV1NotFoundCode int = 404

/*SearchUserV1NotFound Not Found

swagger:response searchUserV1NotFound
*/
type SearchUserV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1NotFound creates SearchUserV1NotFound with default headers values
func NewSearchUserV1NotFound() *SearchUserV1NotFound {

	return &SearchUserV1NotFound{}
}

// WithPayload adds the payload to the search user v1 not found response
func (o *SearchUserV1NotFound) WithPayload(payload *models.ErrResponse) *SearchUserV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 not found response
func (o *SearchUserV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1ConflictCode is the HTTP code returned for type SearchUserV1Conflict
const SearchUserV1ConflictCode int = 409

/*SearchUserV1Conflict Conflict

swagger:response searchUserV1Conflict
*/
type SearchUserV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1Conflict creates SearchUserV1Conflict with default headers values
func NewSearchUserV1Conflict() *SearchUserV1Conflict {

	return &SearchUserV1Conflict{}
}

// WithPayload adds the payload to the search user v1 conflict response
func (o *SearchUserV1Conflict) WithPayload(payload *models.ErrResponse) *SearchUserV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 conflict response
func (o *SearchUserV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SearchUserV1InternalServerErrorCode is the HTTP code returned for type SearchUserV1InternalServerError
const SearchUserV1InternalServerErrorCode int = 500

/*SearchUserV1InternalServerError Internal Server Error

swagger:response searchUserV1InternalServerError
*/
type SearchUserV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewSearchUserV1InternalServerError creates SearchUserV1InternalServerError with default headers values
func NewSearchUserV1InternalServerError() *SearchUserV1InternalServerError {

	return &SearchUserV1InternalServerError{}
}

// WithPayload adds the payload to the search user v1 internal server error response
func (o *SearchUserV1InternalServerError) WithPayload(payload *models.ErrResponse) *SearchUserV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the search user v1 internal server error response
func (o *SearchUserV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SearchUserV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
