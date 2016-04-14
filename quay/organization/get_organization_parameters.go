package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetOrganizationParams creates a new GetOrganizationParams object
// with the default values initialized.
func NewGetOrganizationParams() *GetOrganizationParams {
	var ()
	return &GetOrganizationParams{}
}

/*GetOrganizationParams contains all the parameters to send to the API endpoint
for the get organization operation typically these are written to a http.Request
*/
type GetOrganizationParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithOrgname adds the orgname to the get organization params
func (o *GetOrganizationParams) WithOrgname(orgname string) *GetOrganizationParams {
	o.Orgname = orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
