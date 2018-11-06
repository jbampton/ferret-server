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

// GetProjectHandlerFunc turns a function with the right signature into a get project handler
type GetProjectHandlerFunc func(GetProjectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetProjectHandlerFunc) Handle(params GetProjectParams) middleware.Responder {
	return fn(params)
}

// GetProjectHandler interface for that can handle valid get project params
type GetProjectHandler interface {
	Handle(GetProjectParams) middleware.Responder
}

// NewGetProject creates a new http.Handler for the get project operation
func NewGetProject(ctx *middleware.Context, handler GetProjectHandler) *GetProject {
	return &GetProject{Context: ctx, Handler: handler}
}

/*GetProject swagger:route GET /projects/{projectId} getProject

Get Project

*/
type GetProject struct {
	Context *middleware.Context
	Handler GetProjectHandler
}

func (o *GetProject) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetProjectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetProjectOKBody Project Entity
//
// Represents a full Project entity.
// swagger:model GetProjectOKBody
type GetProjectOKBody struct {
	GetProjectOKBodyAllOf0

	// description
	// Max Length: 255
	// Min Length: 10
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	// Max Length: 100
	// Min Length: 3
	Name *string `json:"name"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProjectOKBody) UnmarshalJSON(raw []byte) error {
	// GetProjectOKBodyAO0
	var getProjectOKBodyAO0 GetProjectOKBodyAllOf0
	if err := swag.ReadJSON(raw, &getProjectOKBodyAO0); err != nil {
		return err
	}
	o.GetProjectOKBodyAllOf0 = getProjectOKBodyAO0

	// GetProjectOKBodyAO1
	var dataGetProjectOKBodyAO1 struct {
		Description string `json:"description,omitempty"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataGetProjectOKBodyAO1); err != nil {
		return err
	}

	o.Description = dataGetProjectOKBodyAO1.Description

	o.Name = dataGetProjectOKBodyAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProjectOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	getProjectOKBodyAO0, err := swag.WriteJSON(o.GetProjectOKBodyAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getProjectOKBodyAO0)

	var dataGetProjectOKBodyAO1 struct {
		Description string `json:"description,omitempty"`

		Name *string `json:"name"`
	}

	dataGetProjectOKBodyAO1.Description = o.Description

	dataGetProjectOKBodyAO1.Name = o.Name

	jsonDataGetProjectOKBodyAO1, errGetProjectOKBodyAO1 := swag.WriteJSON(dataGetProjectOKBodyAO1)
	if errGetProjectOKBodyAO1 != nil {
		return nil, errGetProjectOKBodyAO1
	}
	_parts = append(_parts, jsonDataGetProjectOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get project o k body
func (o *GetProjectOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with GetProjectOKBodyAllOf0
	if err := o.GetProjectOKBodyAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetProjectOKBody) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(o.Description) { // not required
		return nil
	}

	if err := validate.MinLength("getProjectOK"+"."+"description", "body", string(o.Description), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("getProjectOK"+"."+"description", "body", string(o.Description), 255); err != nil {
		return err
	}

	return nil
}

func (o *GetProjectOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("getProjectOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	if err := validate.MinLength("getProjectOK"+"."+"name", "body", string(*o.Name), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("getProjectOK"+"."+"name", "body", string(*o.Name), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProjectOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProjectOKBody) UnmarshalBinary(b []byte) error {
	var res GetProjectOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// GetProjectOKBodyAllOf0 Entity
//
// Represents a database entity
// swagger:model GetProjectOKBodyAllOf0
type GetProjectOKBodyAllOf0 struct {

	// id
	// Required: true
	ID *string `json:"id"`

	// rev
	// Required: true
	Rev *string `json:"rev"`

	// created at
	// Required: true
	CreatedAt *string `json:"createdAt"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetProjectOKBodyAllOf0) UnmarshalJSON(raw []byte) error {
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
		CreatedAt *string `json:"createdAt"`

		UpdatedAt string `json:"updatedAt,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	o.CreatedAt = dataAO1.CreatedAt

	o.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetProjectOKBodyAllOf0) MarshalJSON() ([]byte, error) {
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
		CreatedAt *string `json:"createdAt"`

		UpdatedAt string `json:"updatedAt,omitempty"`
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

// Validate validates this get project o k body all of0
func (o *GetProjectOKBodyAllOf0) Validate(formats strfmt.Registry) error {
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

func (o *GetProjectOKBodyAllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetProjectOKBodyAllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", o.Rev); err != nil {
		return err
	}

	return nil
}

func (o *GetProjectOKBodyAllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetProjectOKBodyAllOf0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetProjectOKBodyAllOf0) UnmarshalBinary(b []byte) error {
	var res GetProjectOKBodyAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
