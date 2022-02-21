// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostAuthSignUpOKCode is the HTTP code returned for type PostAuthSignUpOK
const PostAuthSignUpOKCode int = 200

/*PostAuthSignUpOK ok

swagger:response postAuthSignUpOK
*/
type PostAuthSignUpOK struct {
}

// NewPostAuthSignUpOK creates PostAuthSignUpOK with default headers values
func NewPostAuthSignUpOK() *PostAuthSignUpOK {

	return &PostAuthSignUpOK{}
}

// WriteResponse to the client
func (o *PostAuthSignUpOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostAuthSignUpBadRequestCode is the HTTP code returned for type PostAuthSignUpBadRequest
const PostAuthSignUpBadRequestCode int = 400

/*PostAuthSignUpBadRequest Bad request or user with such email already exists

swagger:response postAuthSignUpBadRequest
*/
type PostAuthSignUpBadRequest struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAuthSignUpBadRequest creates PostAuthSignUpBadRequest with default headers values
func NewPostAuthSignUpBadRequest() *PostAuthSignUpBadRequest {

	return &PostAuthSignUpBadRequest{}
}

// WithPayload adds the payload to the post auth sign up bad request response
func (o *PostAuthSignUpBadRequest) WithPayload(payload string) *PostAuthSignUpBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth sign up bad request response
func (o *PostAuthSignUpBadRequest) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthSignUpBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PostAuthSignUpInternalServerErrorCode is the HTTP code returned for type PostAuthSignUpInternalServerError
const PostAuthSignUpInternalServerErrorCode int = 500

/*PostAuthSignUpInternalServerError internal error

swagger:response postAuthSignUpInternalServerError
*/
type PostAuthSignUpInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAuthSignUpInternalServerError creates PostAuthSignUpInternalServerError with default headers values
func NewPostAuthSignUpInternalServerError() *PostAuthSignUpInternalServerError {

	return &PostAuthSignUpInternalServerError{}
}

// WithPayload adds the payload to the post auth sign up internal server error response
func (o *PostAuthSignUpInternalServerError) WithPayload(payload string) *PostAuthSignUpInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth sign up internal server error response
func (o *PostAuthSignUpInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthSignUpInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
