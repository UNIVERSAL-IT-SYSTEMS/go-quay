package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetLoggedInUserParams creates a new GetLoggedInUserParams object
// with the default values initialized.
func NewGetLoggedInUserParams() *GetLoggedInUserParams {

	return &GetLoggedInUserParams{}
}

/*GetLoggedInUserParams contains all the parameters to send to the API endpoint
for the get logged in user operation typically these are written to a http.Request
*/
type GetLoggedInUserParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetLoggedInUserParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
