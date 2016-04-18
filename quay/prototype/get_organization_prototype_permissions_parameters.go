package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOrganizationPrototypePermissionsParams creates a new GetOrganizationPrototypePermissionsParams object
// with the default values initialized.
func NewGetOrganizationPrototypePermissionsParams() *GetOrganizationPrototypePermissionsParams {
	var ()
	return &GetOrganizationPrototypePermissionsParams{}
}

/*GetOrganizationPrototypePermissionsParams contains all the parameters to send to the API endpoint
for the get organization prototype permissions operation typically these are written to a http.Request
*/
type GetOrganizationPrototypePermissionsParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithOrgname adds the orgname to the get organization prototype permissions params
func (o *GetOrganizationPrototypePermissionsParams) WithOrgname(Orgname string) *GetOrganizationPrototypePermissionsParams {
	o.Orgname = Orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationPrototypePermissionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
