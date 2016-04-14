package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewListBuildTriggersParams creates a new ListBuildTriggersParams object
// with the default values initialized.
func NewListBuildTriggersParams() *ListBuildTriggersParams {
	var ()
	return &ListBuildTriggersParams{}
}

/*ListBuildTriggersParams contains all the parameters to send to the API endpoint
for the list build triggers operation typically these are written to a http.Request
*/
type ListBuildTriggersParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithRepository adds the repository to the list build triggers params
func (o *ListBuildTriggersParams) WithRepository(repository string) *ListBuildTriggersParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ListBuildTriggersParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
