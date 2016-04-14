package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewCancelRepoBuildParams creates a new CancelRepoBuildParams object
// with the default values initialized.
func NewCancelRepoBuildParams() *CancelRepoBuildParams {
	var ()
	return &CancelRepoBuildParams{}
}

/*CancelRepoBuildParams contains all the parameters to send to the API endpoint
for the cancel repo build operation typically these are written to a http.Request
*/
type CancelRepoBuildParams struct {

	/*BuildUUID
	  The UUID of the build

	*/
	BuildUUID string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithBuildUUID adds the buildUuid to the cancel repo build params
func (o *CancelRepoBuildParams) WithBuildUUID(buildUuid string) *CancelRepoBuildParams {
	o.BuildUUID = buildUuid
	return o
}

// WithRepository adds the repository to the cancel repo build params
func (o *CancelRepoBuildParams) WithRepository(repository string) *CancelRepoBuildParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CancelRepoBuildParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param build_uuid
	if err := r.SetPathParam("build_uuid", o.BuildUUID); err != nil {
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
