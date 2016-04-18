package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOrganizationApplicationParams creates a new DeleteOrganizationApplicationParams object
// with the default values initialized.
func NewDeleteOrganizationApplicationParams() *DeleteOrganizationApplicationParams {
	var ()
	return &DeleteOrganizationApplicationParams{}
}

/*DeleteOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the delete organization application operation typically these are written to a http.Request
*/
type DeleteOrganizationApplicationParams struct {

	/*ClientID
	  The OAuth client ID

	*/
	ClientID string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
}

// WithClientID adds the clientId to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithClientID(ClientID string) *DeleteOrganizationApplicationParams {
	o.ClientID = ClientID
	return o
}

// WithOrgname adds the orgname to the delete organization application params
func (o *DeleteOrganizationApplicationParams) WithOrgname(Orgname string) *DeleteOrganizationApplicationParams {
	o.Orgname = Orgname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
