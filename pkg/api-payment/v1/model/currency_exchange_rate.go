// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CurrencyExchangeRate currency exchange rate
//
// swagger:model currencyExchangeRate
type CurrencyExchangeRate struct {

	// crypto amount
	// Example: 0.1231232453453
	CryptoAmount string `json:"cryptoAmount"`

	// crypto currency
	// Example: DAI
	CryptoCurrency string `json:"cryptoCurrency"`

	// display name
	// Example: USD → ETH_DAI
	DisplayName string `json:"displayName"`

	// exchange rate
	// Example: 51.1
	ExchangeRate float64 `json:"exchangeRate"`

	// fiat amount
	// Example: 49.9
	// Minimum: 0.01
	FiatAmount float64 `json:"fiatAmount"`

	// fiat currency
	// Example: USD
	FiatCurrency string `json:"fiatCurrency"`

	// network
	// Example: ETH
	Network string `json:"network"`
}

// Validate validates this currency exchange rate
func (m *CurrencyExchangeRate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiatAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CurrencyExchangeRate) validateFiatAmount(formats strfmt.Registry) error {
	if swag.IsZero(m.FiatAmount) { // not required
		return nil
	}

	if err := validate.Minimum("fiatAmount", "body", m.FiatAmount, 0.01, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this currency exchange rate based on context it is used
func (m *CurrencyExchangeRate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CurrencyExchangeRate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CurrencyExchangeRate) UnmarshalBinary(b []byte) error {
	var res CurrencyExchangeRate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
