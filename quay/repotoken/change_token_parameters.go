package repotoken

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewChangeTokenParams creates a new ChangeTokenParams object
// with the default values initialized.
func NewChangeTokenParams() *ChangeTokenParams {
	var ()
	return &ChangeTokenParams{}
}

/*ChangeTokenParams contains all the parameters to send to the API endpoint
for the change token operation typically these are written to a http.Request
*/
type ChangeTokenParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.TokenPermission
	/*Code
	  The token code

	*/
	Code string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithBody adds the body to the change token params
func (o *ChangeTokenParams) WithBody(body *models.TokenPermission) *ChangeTokenParams {
	o.Body = body
	return o
}

// WithCode adds the code to the change token params
func (o *ChangeTokenParams) WithCode(code string) *ChangeTokenParams {
	o.Code = code
	return o
}

// WithRepository adds the repository to the change token params
func (o *ChangeTokenParams) WithRepository(repository string) *ChangeTokenParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTokenParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.TokenPermission)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param code
	if err := r.SetPathParam("code", o.Code); err != nil {
		return err
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
