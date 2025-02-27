// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RestoreProgress restore progress
//
// swagger:model RestoreProgress
type RestoreProgress struct {

	// completed at
	// Format: date-time
	CompletedAt *strfmt.DateTime `json:"completed_at,omitempty"`

	// downloaded
	Downloaded int64 `json:"downloaded,omitempty"`

	// failed
	Failed int64 `json:"failed,omitempty"`

	// keyspaces
	Keyspaces []*RestoreKeyspaceProgress `json:"keyspaces"`

	// restored
	Restored int64 `json:"restored,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// snapshot tag
	SnapshotTag string `json:"snapshot_tag,omitempty"`

	// stage
	Stage string `json:"stage,omitempty"`

	// started at
	// Format: date-time
	StartedAt *strfmt.DateTime `json:"started_at,omitempty"`
}

// Validate validates this restore progress
func (m *RestoreProgress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyspaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RestoreProgress) validateCompletedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("completed_at", "body", "date-time", m.CompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RestoreProgress) validateKeyspaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Keyspaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Keyspaces); i++ {
		if swag.IsZero(m.Keyspaces[i]) { // not required
			continue
		}

		if m.Keyspaces[i] != nil {
			if err := m.Keyspaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("keyspaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *RestoreProgress) validateStartedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("started_at", "body", "date-time", m.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RestoreProgress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RestoreProgress) UnmarshalBinary(b []byte) error {
	var res RestoreProgress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
