// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/MontFerret/ferret-server/server/http/api/models"
)

// GetScriptDataOKCode is the HTTP code returned for type GetScriptDataOK
const GetScriptDataOKCode int = 200

/*GetScriptDataOK get script data o k

swagger:response getScriptDataOK
*/
type GetScriptDataOK struct {

	/*
	  In: Body
	*/
	Payload *models.DataOutputDetailed `json:"body,omitempty"`
}

// NewGetScriptDataOK creates GetScriptDataOK with default headers values
func NewGetScriptDataOK() *GetScriptDataOK {

	return &GetScriptDataOK{}
}

// WithPayload adds the payload to the get script data o k response
func (o *GetScriptDataOK) WithPayload(payload *models.DataOutputDetailed) *GetScriptDataOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get script data o k response
func (o *GetScriptDataOK) SetPayload(payload *models.DataOutputDetailed) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetScriptDataOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
