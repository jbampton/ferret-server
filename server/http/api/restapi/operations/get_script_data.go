// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// GetScriptDataHandlerFunc turns a function with the right signature into a get script data handler
type GetScriptDataHandlerFunc func(GetScriptDataParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetScriptDataHandlerFunc) Handle(params GetScriptDataParams) middleware.Responder {
	return fn(params)
}

// GetScriptDataHandler interface for that can handle valid get script data params
type GetScriptDataHandler interface {
	Handle(GetScriptDataParams) middleware.Responder
}

// NewGetScriptData creates a new http.Handler for the get script data operation
func NewGetScriptData(ctx *middleware.Context, handler GetScriptDataHandler) *GetScriptData {
	return &GetScriptData{Context: ctx, Handler: handler}
}

/*GetScriptData swagger:route GET /projects/{projectID}/data/{scriptID}/{dataId} getScriptData

Get Data

*/
type GetScriptData struct {
	Context *middleware.Context
	Handler GetScriptDataHandler
}

func (o *GetScriptData) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetScriptDataParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetScriptDataOKBody Data Output Detailed
//
// The properties that are included when fetching a single Data.
// swagger:model GetScriptDataOKBody
type GetScriptDataOKBody struct {
	GetScriptDataOKBodyAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetScriptDataOKBody) UnmarshalJSON(raw []byte) error {
	// GetScriptDataOKBodyAO0
	var getScriptDataOKBodyAO0 GetScriptDataOKBodyAllOf0
	if err := swag.ReadJSON(raw, &getScriptDataOKBodyAO0); err != nil {
		return err
	}
	o.GetScriptDataOKBodyAllOf0 = getScriptDataOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetScriptDataOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getScriptDataOKBodyAO0, err := swag.WriteJSON(o.GetScriptDataOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getScriptDataOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get script data o k body
func (o *GetScriptDataOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetScriptDataOKBodyAllOf0
	if err := o.GetScriptDataOKBodyAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetScriptDataOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetScriptDataOKBody) UnmarshalBinary(b []byte) error {
	var res GetScriptDataOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetScriptDataOKBodyAllOf0 Data Entity
// swagger:model GetScriptDataOKBodyAllOf0
type GetScriptDataOKBodyAllOf0 struct {
	GetScriptDataOKBodyAllOf0AllOf0

	// job id
	// Required: true
	JobID *string `json:"job_id"`

	// script id
	// Required: true
	ScriptID *string `json:"script_id"`

	// script rev
	// Required: true
	ScriptRev *string `json:"script_rev"`

	// Any
	// Required: true
	Value interface{} `json:"value"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetScriptDataOKBodyAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 GetScriptDataOKBodyAllOf0AllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	o.GetScriptDataOKBodyAllOf0AllOf0 = aO0

	// AO1
	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev *string `json:"script_rev"`

		Value interface{} `json:"value"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.JobID = dataAO1.JobID

	o.ScriptID = dataAO1.ScriptID

	o.ScriptRev = dataAO1.ScriptRev

	o.Value = dataAO1.Value

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetScriptDataOKBodyAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(o.GetScriptDataOKBodyAllOf0AllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		JobID *string `json:"job_id"`

		ScriptID *string `json:"script_id"`

		ScriptRev *string `json:"script_rev"`

		Value interface{} `json:"value"`
	}

	dataAO1.JobID = o.JobID

	dataAO1.ScriptID = o.ScriptID

	dataAO1.ScriptRev = o.ScriptRev

	dataAO1.Value = o.Value

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get script data o k body all of0
func (o *GetScriptDataOKBodyAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetScriptDataOKBodyAllOf0AllOf0
	if err := o.GetScriptDataOKBodyAllOf0AllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScriptID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScriptRev(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetScriptDataOKBodyAllOf0) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", o.JobID); err != nil {
		return err
	}

	return nil
}

func (o *GetScriptDataOKBodyAllOf0) validateScriptID(formats strfmt.Registry) error {

	if err := validate.Required("script_id", "body", o.ScriptID); err != nil {
		return err
	}

	return nil
}

func (o *GetScriptDataOKBodyAllOf0) validateScriptRev(formats strfmt.Registry) error {

	if err := validate.Required("script_rev", "body", o.ScriptRev); err != nil {
		return err
	}

	return nil
}

func (o *GetScriptDataOKBodyAllOf0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetScriptDataOKBodyAllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetScriptDataOKBodyAllOf0) UnmarshalBinary(b []byte) error {
	var res GetScriptDataOKBodyAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetScriptDataOKBodyAllOf0AllOf0 Entity
//
// Represents a database entity
// swagger:model GetScriptDataOKBodyAllOf0AllOf0
type GetScriptDataOKBodyAllOf0AllOf0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// rev
	// Required: true
	Rev *string `json:"rev"`

	// created at
	// Required: true
	CreatedAt *string `json:"created_at"`

	// updated at
	UpdatedAt string `json:"updated_at,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetScriptDataOKBodyAllOf0AllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	o.ID = dataAO0.ID

	o.Rev = dataAO0.Rev

	// AO1
	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.CreatedAt = dataAO1.CreatedAt

	o.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetScriptDataOKBodyAllOf0AllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}

	dataAO0.ID = o.ID

	dataAO0.Rev = o.Rev

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		CreatedAt *string `json:"created_at"`

		UpdatedAt string `json:"updated_at,omitempty"`
	}

	dataAO1.CreatedAt = o.CreatedAt

	dataAO1.UpdatedAt = o.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get script data o k body all of0 all of0
func (o *GetScriptDataOKBodyAllOf0AllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRev(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetScriptDataOKBodyAllOf0AllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetScriptDataOKBodyAllOf0AllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", o.Rev); err != nil {
		return err
	}

	return nil
}

func (o *GetScriptDataOKBodyAllOf0AllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetScriptDataOKBodyAllOf0AllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetScriptDataOKBodyAllOf0AllOf0) UnmarshalBinary(b []byte) error {
	var res GetScriptDataOKBodyAllOf0AllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
