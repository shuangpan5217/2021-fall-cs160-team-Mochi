// Code generated by go-swagger; DO NOT EDIT.

package groups_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// RemoveGroupUsersV1OKCode is the HTTP code returned for type RemoveGroupUsersV1OK
const RemoveGroupUsersV1OKCode int = 200

/*RemoveGroupUsersV1OK Success

swagger:response removeGroupUsersV1OK
*/
type RemoveGroupUsersV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.GroupResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1OK creates RemoveGroupUsersV1OK with default headers values
func NewRemoveGroupUsersV1OK() *RemoveGroupUsersV1OK {

	return &RemoveGroupUsersV1OK{}
}

// WithPayload adds the payload to the remove group users v1 o k response
func (o *RemoveGroupUsersV1OK) WithPayload(payload *models.GroupResponse) *RemoveGroupUsersV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 o k response
func (o *RemoveGroupUsersV1OK) SetPayload(payload *models.GroupResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1BadRequestCode is the HTTP code returned for type RemoveGroupUsersV1BadRequest
const RemoveGroupUsersV1BadRequestCode int = 400

/*RemoveGroupUsersV1BadRequest Bad Request

swagger:response removeGroupUsersV1BadRequest
*/
type RemoveGroupUsersV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1BadRequest creates RemoveGroupUsersV1BadRequest with default headers values
func NewRemoveGroupUsersV1BadRequest() *RemoveGroupUsersV1BadRequest {

	return &RemoveGroupUsersV1BadRequest{}
}

// WithPayload adds the payload to the remove group users v1 bad request response
func (o *RemoveGroupUsersV1BadRequest) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 bad request response
func (o *RemoveGroupUsersV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1UnauthorizedCode is the HTTP code returned for type RemoveGroupUsersV1Unauthorized
const RemoveGroupUsersV1UnauthorizedCode int = 401

/*RemoveGroupUsersV1Unauthorized Unauthorized

swagger:response removeGroupUsersV1Unauthorized
*/
type RemoveGroupUsersV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1Unauthorized creates RemoveGroupUsersV1Unauthorized with default headers values
func NewRemoveGroupUsersV1Unauthorized() *RemoveGroupUsersV1Unauthorized {

	return &RemoveGroupUsersV1Unauthorized{}
}

// WithPayload adds the payload to the remove group users v1 unauthorized response
func (o *RemoveGroupUsersV1Unauthorized) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 unauthorized response
func (o *RemoveGroupUsersV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1ForbiddenCode is the HTTP code returned for type RemoveGroupUsersV1Forbidden
const RemoveGroupUsersV1ForbiddenCode int = 403

/*RemoveGroupUsersV1Forbidden Forbidden

swagger:response removeGroupUsersV1Forbidden
*/
type RemoveGroupUsersV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1Forbidden creates RemoveGroupUsersV1Forbidden with default headers values
func NewRemoveGroupUsersV1Forbidden() *RemoveGroupUsersV1Forbidden {

	return &RemoveGroupUsersV1Forbidden{}
}

// WithPayload adds the payload to the remove group users v1 forbidden response
func (o *RemoveGroupUsersV1Forbidden) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 forbidden response
func (o *RemoveGroupUsersV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1NotFoundCode is the HTTP code returned for type RemoveGroupUsersV1NotFound
const RemoveGroupUsersV1NotFoundCode int = 404

/*RemoveGroupUsersV1NotFound Not Found

swagger:response removeGroupUsersV1NotFound
*/
type RemoveGroupUsersV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1NotFound creates RemoveGroupUsersV1NotFound with default headers values
func NewRemoveGroupUsersV1NotFound() *RemoveGroupUsersV1NotFound {

	return &RemoveGroupUsersV1NotFound{}
}

// WithPayload adds the payload to the remove group users v1 not found response
func (o *RemoveGroupUsersV1NotFound) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 not found response
func (o *RemoveGroupUsersV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1ConflictCode is the HTTP code returned for type RemoveGroupUsersV1Conflict
const RemoveGroupUsersV1ConflictCode int = 409

/*RemoveGroupUsersV1Conflict Conflict

swagger:response removeGroupUsersV1Conflict
*/
type RemoveGroupUsersV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1Conflict creates RemoveGroupUsersV1Conflict with default headers values
func NewRemoveGroupUsersV1Conflict() *RemoveGroupUsersV1Conflict {

	return &RemoveGroupUsersV1Conflict{}
}

// WithPayload adds the payload to the remove group users v1 conflict response
func (o *RemoveGroupUsersV1Conflict) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 conflict response
func (o *RemoveGroupUsersV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// RemoveGroupUsersV1InternalServerErrorCode is the HTTP code returned for type RemoveGroupUsersV1InternalServerError
const RemoveGroupUsersV1InternalServerErrorCode int = 500

/*RemoveGroupUsersV1InternalServerError Internal Server Error

swagger:response removeGroupUsersV1InternalServerError
*/
type RemoveGroupUsersV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewRemoveGroupUsersV1InternalServerError creates RemoveGroupUsersV1InternalServerError with default headers values
func NewRemoveGroupUsersV1InternalServerError() *RemoveGroupUsersV1InternalServerError {

	return &RemoveGroupUsersV1InternalServerError{}
}

// WithPayload adds the payload to the remove group users v1 internal server error response
func (o *RemoveGroupUsersV1InternalServerError) WithPayload(payload *models.ErrResponse) *RemoveGroupUsersV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the remove group users v1 internal server error response
func (o *RemoveGroupUsersV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RemoveGroupUsersV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
