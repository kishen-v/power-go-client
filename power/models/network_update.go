// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkUpdate network update
//
// swagger:model NetworkUpdate
type NetworkUpdate struct {

	// Replaces the current DNS Servers
	DNSServers []string `json:"dnsServers"`

	// Replaces the current Gateway IP Address
	Gateway *string `json:"gateway,omitempty"`

	// Replaces the current IP Address Ranges
	IPAddressRanges []*IPAddressRange `json:"ipAddressRanges"`

	// Replaces the current Network Name
	Name *string `json:"name,omitempty"`
}

// Validate validates this network update
func (m *NetworkUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddressRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkUpdate) validateIPAddressRanges(formats strfmt.Registry) error {
	if swag.IsZero(m.IPAddressRanges) { // not required
		return nil
	}

	for i := 0; i < len(m.IPAddressRanges); i++ {
		if swag.IsZero(m.IPAddressRanges[i]) { // not required
			continue
		}

		if m.IPAddressRanges[i] != nil {
			if err := m.IPAddressRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this network update based on the context it is used
func (m *NetworkUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPAddressRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkUpdate) contextValidateIPAddressRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IPAddressRanges); i++ {

		if m.IPAddressRanges[i] != nil {

			if swag.IsZero(m.IPAddressRanges[i]) { // not required
				return nil
			}

			if err := m.IPAddressRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipAddressRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkUpdate) UnmarshalBinary(b []byte) error {
	var res NetworkUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
