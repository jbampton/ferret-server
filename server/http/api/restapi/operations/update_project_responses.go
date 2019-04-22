// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// UpdateProjectOKCode is the HTTP code returned for type UpdateProjectOK
const UpdateProjectOKCode int = 200

/*UpdateProjectOK update project o k

swagger:response updateProjectOK
*/
type UpdateProjectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Entity `json:"body,omitempty"`
}

// NewUpdateProjectOK creates UpdateProjectOK with default headers values
func NewUpdateProjectOK() *UpdateProjectOK {

	return &UpdateProjectOK{}
}

// WithPayload adds the payload to the update project o k response
func (o *UpdateProjectOK) WithPayload(payload *models.Entity) *UpdateProjectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update project o k response
func (o *UpdateProjectOK) SetPayload(payload *models.Entity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateProjectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
