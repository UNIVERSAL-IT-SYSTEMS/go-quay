package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*UpdateOrg Description of updates for an existing organization

swagger:model UpdateOrg
*/
type UpdateOrg struct {

	/* Organization contact email
	 */
	Email string `json:"email,omitempty"`

	/* Whether the organization desires to receive emails for invoices
	 */
	InvoiceEmail bool `json:"invoice_email,omitempty"`

	/* The email address at which to receive invoices
	 */
	InvoiceEmailAddress string `json:"invoice_email_address,omitempty"`

	/* tag expiration

	Maximum: 2.592e+06
	Minimum: 0
	*/
	TagExpiration *int64 `json:"tag_expiration,omitempty"`
}

// Validate validates this update org
func (m *UpdateOrg) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTagExpiration(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateOrg) validateTagExpiration(formats strfmt.Registry) error {

	if swag.IsZero(m.TagExpiration) { // not required
		return nil
	}

	if err := validate.MinimumInt("tag_expiration", "body", int64(*m.TagExpiration), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("tag_expiration", "body", int64(*m.TagExpiration), 2.592e+06, false); err != nil {
		return err
	}

	return nil
}
