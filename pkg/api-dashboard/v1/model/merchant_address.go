// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MerchantAddress merchant address
//
// swagger:model merchantAddress
type MerchantAddress struct {

	// address
	// Example: 0xdc22Bb64132fB03467910fc49595F08fCf5C241b
	// Max Length: 128
	Address string `json:"address"`

	// blockchain
	// Example: ETH
	// Enum: [ETH TRON MATIC]
	Blockchain string `json:"blockchain"`

	// Blockchain name
	// Example: Ethereum
	BlockchainName string `json:"blockchainName,omitempty"`

	// Address UUID
	// Example: 123e4567-e89b-12d3-a456-426655440000
	ID string `json:"id"`

	// Name
	// Example: My Wallet
	// Max Length: 128
	// Min Length: 2
	Name string `json:"name"`
}

// Validate validates this merchant address
func (m *MerchantAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockchain(formats); err != nil {
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

func (m *MerchantAddress) validateAddress(formats strfmt.Registry) error {
	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if err := validate.MaxLength("address", "body", m.Address, 128); err != nil {
		return err
	}

	return nil
}

var merchantAddressTypeBlockchainPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ETH","TRON","MATIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		merchantAddressTypeBlockchainPropEnum = append(merchantAddressTypeBlockchainPropEnum, v)
	}
}

const (

	// MerchantAddressBlockchainETH captures enum value "ETH"
	MerchantAddressBlockchainETH string = "ETH"

	// MerchantAddressBlockchainTRON captures enum value "TRON"
	MerchantAddressBlockchainTRON string = "TRON"

	// MerchantAddressBlockchainMATIC captures enum value "MATIC"
	MerchantAddressBlockchainMATIC string = "MATIC"
)

// prop value enum
func (m *MerchantAddress) validateBlockchainEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, merchantAddressTypeBlockchainPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MerchantAddress) validateBlockchain(formats strfmt.Registry) error {
	if swag.IsZero(m.Blockchain) { // not required
		return nil
	}

	// value enum
	if err := m.validateBlockchainEnum("blockchain", "body", m.Blockchain); err != nil {
		return err
	}

	return nil
}

func (m *MerchantAddress) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 2); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 128); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this merchant address based on context it is used
func (m *MerchantAddress) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MerchantAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MerchantAddress) UnmarshalBinary(b []byte) error {
	var res MerchantAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
