package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetOrganizationApplicationParams creates a new GetOrganizationApplicationParams object
// with the default values initialized.
func NewGetOrganizationApplicationParams() *GetOrganizationApplicationParams {
	var ()
	return &GetOrganizationApplicationParams{}
}

/*GetOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the get organization application operation typically these are written to a http.Request
*/
type GetOrganizationApplicationParams struct {

	/*ClientID
	  The OAuth client ID

	*/
	ClientID string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithClientID adds the clientId to the get organization application params
func (o *GetOrganizationApplicationParams) WithClientID(clientId string) *GetOrganizationApplicationParams {
	o.ClientID = clientId
	return o
}

// WithOrgname adds the orgname to the get organization application params
func (o *GetOrganizationApplicationParams) WithOrgname(orgname string) *GetOrganizationApplicationParams {
	o.Orgname = orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationApplicationParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param client_id
	if err := r.SetPathParam("client_id", o.ClientID); err != nil {
		return err
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
