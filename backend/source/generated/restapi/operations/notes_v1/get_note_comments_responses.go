// Code generated by go-swagger; DO NOT EDIT.

package notes_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"2021-fall-cs160-team-Mochi/backend/source/generated/models"
)

// GetNoteCommentsOKCode is the HTTP code returned for type GetNoteCommentsOK
const GetNoteCommentsOKCode int = 200

/*GetNoteCommentsOK Success

swagger:response getNoteCommentsOK
*/
type GetNoteCommentsOK struct {

	/*
	  In: Body
	*/
	Payload *models.NoteCommentsResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsOK creates GetNoteCommentsOK with default headers values
func NewGetNoteCommentsOK() *GetNoteCommentsOK {

	return &GetNoteCommentsOK{}
}

// WithPayload adds the payload to the get note comments o k response
func (o *GetNoteCommentsOK) WithPayload(payload *models.NoteCommentsResponse) *GetNoteCommentsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments o k response
func (o *GetNoteCommentsOK) SetPayload(payload *models.NoteCommentsResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsBadRequestCode is the HTTP code returned for type GetNoteCommentsBadRequest
const GetNoteCommentsBadRequestCode int = 400

/*GetNoteCommentsBadRequest Bad Request

swagger:response getNoteCommentsBadRequest
*/
type GetNoteCommentsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsBadRequest creates GetNoteCommentsBadRequest with default headers values
func NewGetNoteCommentsBadRequest() *GetNoteCommentsBadRequest {

	return &GetNoteCommentsBadRequest{}
}

// WithPayload adds the payload to the get note comments bad request response
func (o *GetNoteCommentsBadRequest) WithPayload(payload *models.ErrResponse) *GetNoteCommentsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments bad request response
func (o *GetNoteCommentsBadRequest) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsUnauthorizedCode is the HTTP code returned for type GetNoteCommentsUnauthorized
const GetNoteCommentsUnauthorizedCode int = 401

/*GetNoteCommentsUnauthorized Unauthorized

swagger:response getNoteCommentsUnauthorized
*/
type GetNoteCommentsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsUnauthorized creates GetNoteCommentsUnauthorized with default headers values
func NewGetNoteCommentsUnauthorized() *GetNoteCommentsUnauthorized {

	return &GetNoteCommentsUnauthorized{}
}

// WithPayload adds the payload to the get note comments unauthorized response
func (o *GetNoteCommentsUnauthorized) WithPayload(payload *models.ErrResponse) *GetNoteCommentsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments unauthorized response
func (o *GetNoteCommentsUnauthorized) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsForbiddenCode is the HTTP code returned for type GetNoteCommentsForbidden
const GetNoteCommentsForbiddenCode int = 403

/*GetNoteCommentsForbidden Forbidden

swagger:response getNoteCommentsForbidden
*/
type GetNoteCommentsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsForbidden creates GetNoteCommentsForbidden with default headers values
func NewGetNoteCommentsForbidden() *GetNoteCommentsForbidden {

	return &GetNoteCommentsForbidden{}
}

// WithPayload adds the payload to the get note comments forbidden response
func (o *GetNoteCommentsForbidden) WithPayload(payload *models.ErrResponse) *GetNoteCommentsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments forbidden response
func (o *GetNoteCommentsForbidden) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsNotFoundCode is the HTTP code returned for type GetNoteCommentsNotFound
const GetNoteCommentsNotFoundCode int = 404

/*GetNoteCommentsNotFound Not Found

swagger:response getNoteCommentsNotFound
*/
type GetNoteCommentsNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsNotFound creates GetNoteCommentsNotFound with default headers values
func NewGetNoteCommentsNotFound() *GetNoteCommentsNotFound {

	return &GetNoteCommentsNotFound{}
}

// WithPayload adds the payload to the get note comments not found response
func (o *GetNoteCommentsNotFound) WithPayload(payload *models.ErrResponse) *GetNoteCommentsNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments not found response
func (o *GetNoteCommentsNotFound) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsConflictCode is the HTTP code returned for type GetNoteCommentsConflict
const GetNoteCommentsConflictCode int = 409

/*GetNoteCommentsConflict Conflict

swagger:response getNoteCommentsConflict
*/
type GetNoteCommentsConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsConflict creates GetNoteCommentsConflict with default headers values
func NewGetNoteCommentsConflict() *GetNoteCommentsConflict {

	return &GetNoteCommentsConflict{}
}

// WithPayload adds the payload to the get note comments conflict response
func (o *GetNoteCommentsConflict) WithPayload(payload *models.ErrResponse) *GetNoteCommentsConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments conflict response
func (o *GetNoteCommentsConflict) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetNoteCommentsInternalServerErrorCode is the HTTP code returned for type GetNoteCommentsInternalServerError
const GetNoteCommentsInternalServerErrorCode int = 500

/*GetNoteCommentsInternalServerError Internal Server Error

swagger:response getNoteCommentsInternalServerError
*/
type GetNoteCommentsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrResponse `json:"body,omitempty"`
}

// NewGetNoteCommentsInternalServerError creates GetNoteCommentsInternalServerError with default headers values
func NewGetNoteCommentsInternalServerError() *GetNoteCommentsInternalServerError {

	return &GetNoteCommentsInternalServerError{}
}

// WithPayload adds the payload to the get note comments internal server error response
func (o *GetNoteCommentsInternalServerError) WithPayload(payload *models.ErrResponse) *GetNoteCommentsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get note comments internal server error response
func (o *GetNoteCommentsInternalServerError) SetPayload(payload *models.ErrResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNoteCommentsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
