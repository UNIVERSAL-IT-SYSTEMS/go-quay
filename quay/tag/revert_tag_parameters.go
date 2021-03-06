package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewRevertTagParams creates a new RevertTagParams object
// with the default values initialized.
func NewRevertTagParams() *RevertTagParams {
	var ()
	return &RevertTagParams{}
}

/*RevertTagParams contains all the parameters to send to the API endpoint
for the revert tag operation typically these are written to a http.Request
*/
type RevertTagParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.RevertTag
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Tag
	  The name of the tag

	*/
	Tag string
}

// WithBody adds the body to the revert tag params
func (o *RevertTagParams) WithBody(Body *models.RevertTag) *RevertTagParams {
	o.Body = Body
	return o
}

// WithRepository adds the repository to the revert tag params
func (o *RevertTagParams) WithRepository(Repository string) *RevertTagParams {
	o.Repository = Repository
	return o
}

// WithTag adds the tag to the revert tag params
func (o *RevertTagParams) WithTag(Tag string) *RevertTagParams {
	o.Tag = Tag
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RevertTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.RevertTag)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

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
