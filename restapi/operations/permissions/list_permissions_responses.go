package permissions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cyverse-de/permissions/models"
)

// ListPermissionsOKCode is the HTTP code returned for type ListPermissionsOK
const ListPermissionsOKCode int = 200

/*ListPermissionsOK OK

swagger:response listPermissionsOK
*/
type ListPermissionsOK struct {

	/*
	  In: Body
	*/
	Payload *models.PermissionList `json:"body,omitempty"`
}

// NewListPermissionsOK creates ListPermissionsOK with default headers values
func NewListPermissionsOK() *ListPermissionsOK {
	return &ListPermissionsOK{}
}

// WithPayload adds the payload to the list permissions o k response
func (o *ListPermissionsOK) WithPayload(payload *models.PermissionList) *ListPermissionsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list permissions o k response
func (o *ListPermissionsOK) SetPayload(payload *models.PermissionList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPermissionsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListPermissionsInternalServerErrorCode is the HTTP code returned for type ListPermissionsInternalServerError
const ListPermissionsInternalServerErrorCode int = 500

/*ListPermissionsInternalServerError Internal Server Error

swagger:response listPermissionsInternalServerError
*/
type ListPermissionsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorOut `json:"body,omitempty"`
}

// NewListPermissionsInternalServerError creates ListPermissionsInternalServerError with default headers values
func NewListPermissionsInternalServerError() *ListPermissionsInternalServerError {
	return &ListPermissionsInternalServerError{}
}

// WithPayload adds the payload to the list permissions internal server error response
func (o *ListPermissionsInternalServerError) WithPayload(payload *models.ErrorOut) *ListPermissionsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list permissions internal server error response
func (o *ListPermissionsInternalServerError) SetPayload(payload *models.ErrorOut) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListPermissionsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}