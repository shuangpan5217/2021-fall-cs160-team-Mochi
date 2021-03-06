// Code generated by go-swagger; DO NOT EDIT.

package friends_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// GetFriendsV1OKCode is the HTTP code returned for type GetFriendsV1OK
const GetFriendsV1OKCode int = 200

/*GetFriendsV1OK Success

swagger:response getFriendsV1OK
*/
type GetFriendsV1OK struct {

	/*
	  In: Body
	*/
	Payload *models.GetFriendObject `json:"body,omitempty"`
}

// NewGetFriendsV1OK creates GetFriendsV1OK with default headers values
func NewGetFriendsV1OK() *GetFriendsV1OK {

	return &GetFriendsV1OK{}
}

// WithPayload adds the payload to the get friends v1 o k response
func (o *GetFriendsV1OK) WithPayload(payload *models.GetFriendObject) *GetFriendsV1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 o k response
func (o *GetFriendsV1OK) SetPayload(payload *models.GetFriendObject) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1BadRequestCode is the HTTP code returned for type GetFriendsV1BadRequest
const GetFriendsV1BadRequestCode int = 400

/*GetFriendsV1BadRequest Bad Request

swagger:response getFriendsV1BadRequest
*/
type GetFriendsV1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1BadRequest creates GetFriendsV1BadRequest with default headers values
func NewGetFriendsV1BadRequest() *GetFriendsV1BadRequest {

	return &GetFriendsV1BadRequest{}
}

// WithPayload adds the payload to the get friends v1 bad request response
func (o *GetFriendsV1BadRequest) WithPayload(payload *models.ErrResponse) *GetFriendsV1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 bad request response
func (o *GetFriendsV1BadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1UnauthorizedCode is the HTTP code returned for type GetFriendsV1Unauthorized
const GetFriendsV1UnauthorizedCode int = 401

/*GetFriendsV1Unauthorized Unauthorized

swagger:response getFriendsV1Unauthorized
*/
type GetFriendsV1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1Unauthorized creates GetFriendsV1Unauthorized with default headers values
func NewGetFriendsV1Unauthorized() *GetFriendsV1Unauthorized {

	return &GetFriendsV1Unauthorized{}
}

// WithPayload adds the payload to the get friends v1 unauthorized response
func (o *GetFriendsV1Unauthorized) WithPayload(payload *models.ErrResponse) *GetFriendsV1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 unauthorized response
func (o *GetFriendsV1Unauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1ForbiddenCode is the HTTP code returned for type GetFriendsV1Forbidden
const GetFriendsV1ForbiddenCode int = 403

/*GetFriendsV1Forbidden Forbidden

swagger:response getFriendsV1Forbidden
*/
type GetFriendsV1Forbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1Forbidden creates GetFriendsV1Forbidden with default headers values
func NewGetFriendsV1Forbidden() *GetFriendsV1Forbidden {

	return &GetFriendsV1Forbidden{}
}

// WithPayload adds the payload to the get friends v1 forbidden response
func (o *GetFriendsV1Forbidden) WithPayload(payload *models.ErrResponse) *GetFriendsV1Forbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 forbidden response
func (o *GetFriendsV1Forbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1NotFoundCode is the HTTP code returned for type GetFriendsV1NotFound
const GetFriendsV1NotFoundCode int = 404

/*GetFriendsV1NotFound Not Found

swagger:response getFriendsV1NotFound
*/
type GetFriendsV1NotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1NotFound creates GetFriendsV1NotFound with default headers values
func NewGetFriendsV1NotFound() *GetFriendsV1NotFound {

	return &GetFriendsV1NotFound{}
}

// WithPayload adds the payload to the get friends v1 not found response
func (o *GetFriendsV1NotFound) WithPayload(payload *models.ErrResponse) *GetFriendsV1NotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 not found response
func (o *GetFriendsV1NotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1ConflictCode is the HTTP code returned for type GetFriendsV1Conflict
const GetFriendsV1ConflictCode int = 409

/*GetFriendsV1Conflict Conflict

swagger:response getFriendsV1Conflict
*/
type GetFriendsV1Conflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1Conflict creates GetFriendsV1Conflict with default headers values
func NewGetFriendsV1Conflict() *GetFriendsV1Conflict {

	return &GetFriendsV1Conflict{}
}

// WithPayload adds the payload to the get friends v1 conflict response
func (o *GetFriendsV1Conflict) WithPayload(payload *models.ErrResponse) *GetFriendsV1Conflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 conflict response
func (o *GetFriendsV1Conflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1Conflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFriendsV1InternalServerErrorCode is the HTTP code returned for type GetFriendsV1InternalServerError
const GetFriendsV1InternalServerErrorCode int = 500

/*GetFriendsV1InternalServerError Internal Server Error

swagger:response getFriendsV1InternalServerError
*/
type GetFriendsV1InternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetFriendsV1InternalServerError creates GetFriendsV1InternalServerError with default headers values
func NewGetFriendsV1InternalServerError() *GetFriendsV1InternalServerError {

	return &GetFriendsV1InternalServerError{}
}

// WithPayload adds the payload to the get friends v1 internal server error response
func (o *GetFriendsV1InternalServerError) WithPayload(payload *models.ErrResponse) *GetFriendsV1InternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get friends v1 internal server error response
func (o *GetFriendsV1InternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFriendsV1InternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
