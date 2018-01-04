// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkCreateConfig contains the request for the remote API: POST /networks/create
// swagger:model NetworkCreateConfig

type NetworkCreateConfig struct {

	// Name is the name of the network.
	Name string `json:"Name,omitempty"`

	NetworkCreate
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NetworkCreateConfig) UnmarshalJSON(raw []byte) error {

	var data struct {
		Name string `json:"Name,omitempty"`
	}
	if err := swag.ReadJSON(raw, &data); err != nil {
		return err
	}

	m.Name = data.Name

	var aO1 NetworkCreate
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.NetworkCreate = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NetworkCreateConfig) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	var data struct {
		Name string `json:"Name,omitempty"`
	}

	data.Name = m.Name

	jsonData, err := swag.WriteJSON(data)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jsonData)

	aO1, err := swag.WriteJSON(m.NetworkCreate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this network create config
func (m *NetworkCreateConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.NetworkCreate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkCreateConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkCreateConfig) UnmarshalBinary(b []byte) error {
	var res NetworkCreateConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}