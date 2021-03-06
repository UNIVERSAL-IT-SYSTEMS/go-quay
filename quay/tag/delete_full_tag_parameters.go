package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteFullTagParams creates a new DeleteFullTagParams object
// with the default values initialized.
func NewDeleteFullTagParams() *DeleteFullTagParams {
	var ()
	return &DeleteFullTagParams{}
}

/*DeleteFullTagParams contains all the parameters to send to the API endpoint
for the delete full tag operation typically these are written to a http.Request
*/
type DeleteFullTagParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string
}

// WithRepository adds the repository to the delete full tag params
func (o *DeleteFullTagParams) WithRepository(Repository string) *DeleteFullTagParams {
	o.Repository = Repository
	return o
}

// WithTag adds the tag to the delete full tag params
func (o *DeleteFullTagParams) WithTag(Tag string) *DeleteFullTagParams {
	o.Tag = Tag
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFullTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param tag
	if err := r.SetPathParam("tag", o.Tag); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
