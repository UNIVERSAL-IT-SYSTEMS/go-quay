package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetRepoBuildParams creates a new GetRepoBuildParams object
// with the default values initialized.
func NewGetRepoBuildParams() *GetRepoBuildParams {
	var ()
	return &GetRepoBuildParams{}
}

/*GetRepoBuildParams contains all the parameters to send to the API endpoint
for the get repo build operation typically these are written to a http.Request
*/
type GetRepoBuildParams struct {

	/*BuildUUID
	  The UUID of the build

	*/
	BuildUUID string
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
}

// WithBuildUUID adds the buildUuid to the get repo build params
func (o *GetRepoBuildParams) WithBuildUUID(buildUuid string) *GetRepoBuildParams {
	o.BuildUUID = buildUuid
	return o
}

// WithRepository adds the repository to the get repo build params
func (o *GetRepoBuildParams) WithRepository(repository string) *GetRepoBuildParams {
	o.Repository = repository
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoBuildParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
