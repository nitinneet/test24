// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SubscriberState EPC state for a subscriber
// swagger:model subscriber_state
type SubscriberState struct {

	// mme
	Mme UntypedMmeState `json:"mme,omitempty"`

	// IP addresses which have been allocated for this subscriber
	Mobility []*SubscriberIPAllocation `json:"mobility,omitempty"`

	// s1ap
	S1ap UntypedMmeState `json:"s1ap,omitempty"`

	// spgw
	Spgw UntypedMmeState `json:"spgw,omitempty"`
}

// Validate validates this subscriber state
func (m *SubscriberState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMobility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriberState) validateMobility(formats strfmt.Registry) error {

	if swag.IsZero(m.Mobility) { // not required
		return nil
	}

	for i := 0; i < len(m.Mobility); i++ {
		if swag.IsZero(m.Mobility[i]) { // not required
			continue
		}

		if m.Mobility[i] != nil {
			if err := m.Mobility[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mobility" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriberState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriberState) UnmarshalBinary(b []byte) error {
	var res SubscriberState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
