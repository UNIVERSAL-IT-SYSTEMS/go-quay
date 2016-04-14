package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewRequestRepoBuildParams creates a new RequestRepoBuildParams object
// with the default values initialized.
func NewRequestRepoBuildParams() *RequestRepoBuildParams {
	var ()
	return &RequestRepoBuildParams{}
}

/*RequestRepoBuildParams contains all the parameters to send to the API endpoint
for the request repo build operation typically these are written to a http.Request
*/
type RequestRepoBuildParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.RepositoryBuildRequest
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithBody adds the body to the request repo build params
func (o *RequestRepoBuildParams) WithBody(body *models.RepositoryBuildRequest) *RequestRepoBuildParams {
	o.Body = body
	return o
}

// WithRepository adds the repository to the request repo build params
func (o *RequestRepoBuildParams) WithRepository(repository string) *RequestRepoBuildParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *RequestRepoBuildParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.RepositoryBuildRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
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
