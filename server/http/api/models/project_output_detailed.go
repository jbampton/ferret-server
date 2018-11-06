// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectOutputDetailed Project Entity
//
// Represents a full Project entity.
// swagger:model project-output-detailed
type ProjectOutputDetailed struct {
	ProjectOutputDetailedAllOf0

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
func (m *ProjectOutputDetailed) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ProjectOutputDetailedAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ProjectOutputDetailedAllOf0 = aO0

	// AO1
	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Name *string `json:"name"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Description = dataAO1.Description

	m.Name = dataAO1.Name

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProjectOutputDetailed) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ProjectOutputDetailedAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Description string `json:"description,omitempty"`

		Name *string `json:"name"`
	}

	dataAO1.Description = m.Description

	dataAO1.Name = m.Name

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this project output detailed
func (m *ProjectOutputDetailed) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ProjectOutputDetailedAllOf0
	if err := m.ProjectOutputDetailedAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectOutputDetailed) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(m.Description), 10); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(m.Description), 255); err != nil {
		return err
	}

	return nil
}

func (m *ProjectOutputDetailed) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectOutputDetailed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectOutputDetailed) UnmarshalBinary(b []byte) error {
	var res ProjectOutputDetailed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ProjectOutputDetailedAllOf0 Entity
//
// Represents a database entity
// swagger:model ProjectOutputDetailedAllOf0
type ProjectOutputDetailedAllOf0 struct {

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
func (m *ProjectOutputDetailedAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}
	if err := swag.ReadJSON(raw, &dataAO0); err != nil {
		return err
	}

	m.ID = dataAO0.ID

	m.Rev = dataAO0.Rev

	// AO1
	var dataAO1 struct {
		CreatedAt *string `json:"createdAt"`

		UpdatedAt string `json:"updatedAt,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.CreatedAt = dataAO1.CreatedAt

	m.UpdatedAt = dataAO1.UpdatedAt

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProjectOutputDetailedAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	var dataAO0 struct {
		ID *string `json:"id"`

		Rev *string `json:"rev"`
	}

	dataAO0.ID = m.ID

	dataAO0.Rev = m.Rev

	jsonDataAO0, errAO0 := swag.WriteJSON(dataAO0)
	if errAO0 != nil {
		return nil, errAO0
	}
	_parts = append(_parts, jsonDataAO0)

	var dataAO1 struct {
		CreatedAt *string `json:"createdAt"`

		UpdatedAt string `json:"updatedAt,omitempty"`
	}

	dataAO1.CreatedAt = m.CreatedAt

	dataAO1.UpdatedAt = m.UpdatedAt

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this project output detailed all of0
func (m *ProjectOutputDetailedAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRev(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectOutputDetailedAllOf0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectOutputDetailedAllOf0) validateRev(formats strfmt.Registry) error {

	if err := validate.Required("rev", "body", m.Rev); err != nil {
		return err
	}

	return nil
}

func (m *ProjectOutputDetailedAllOf0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("createdAt", "body", m.CreatedAt); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectOutputDetailedAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectOutputDetailedAllOf0) UnmarshalBinary(b []byte) error {
	var res ProjectOutputDetailedAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
