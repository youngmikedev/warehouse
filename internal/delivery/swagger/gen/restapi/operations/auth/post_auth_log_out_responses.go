// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostAuthLogOutOKCode is the HTTP code returned for type PostAuthLogOutOK
const PostAuthLogOutOKCode int = 200

/*PostAuthLogOutOK ok

swagger:response postAuthLogOutOK
*/
type PostAuthLogOutOK struct {
}

// NewPostAuthLogOutOK creates PostAuthLogOutOK with default headers values
func NewPostAuthLogOutOK() *PostAuthLogOutOK {

	return &PostAuthLogOutOK{}
}

// WriteResponse to the client
func (o *PostAuthLogOutOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostAuthLogOutUnauthorizedCode is the HTTP code returned for type PostAuthLogOutUnauthorized
const PostAuthLogOutUnauthorizedCode int = 401

/*PostAuthLogOutUnauthorized unauthorized

swagger:response postAuthLogOutUnauthorized
*/
type PostAuthLogOutUnauthorized struct {
}

// NewPostAuthLogOutUnauthorized creates PostAuthLogOutUnauthorized with default headers values
func NewPostAuthLogOutUnauthorized() *PostAuthLogOutUnauthorized {

	return &PostAuthLogOutUnauthorized{}
}

// WriteResponse to the client
func (o *PostAuthLogOutUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// PostAuthLogOutInternalServerErrorCode is the HTTP code returned for type PostAuthLogOutInternalServerError
const PostAuthLogOutInternalServerErrorCode int = 500

/*PostAuthLogOutInternalServerError internal error

swagger:response postAuthLogOutInternalServerError
*/
type PostAuthLogOutInternalServerError struct {

	/*
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPostAuthLogOutInternalServerError creates PostAuthLogOutInternalServerError with default headers values
func NewPostAuthLogOutInternalServerError() *PostAuthLogOutInternalServerError {

	return &PostAuthLogOutInternalServerError{}
}

// WithPayload adds the payload to the post auth log out internal server error response
func (o *PostAuthLogOutInternalServerError) WithPayload(payload string) *PostAuthLogOutInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post auth log out internal server error response
func (o *PostAuthLogOutInternalServerError) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostAuthLogOutInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
