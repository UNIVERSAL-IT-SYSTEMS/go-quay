package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBuildTriggerParams creates a new DeleteBuildTriggerParams object
// with the default values initialized.
func NewDeleteBuildTriggerParams() *DeleteBuildTriggerParams {
	var ()
	return &DeleteBuildTriggerParams{}
}

/*DeleteBuildTriggerParams contains all the parameters to send to the API endpoint
for the delete build trigger operation typically these are written to a http.Request
*/
type DeleteBuildTriggerParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*TriggerUUID
	  The UUID of the build trigger

	*/
	TriggerUUID string
}

// WithRepository adds the repository to the delete build trigger params
func (o *DeleteBuildTriggerParams) WithRepository(Repository string) *DeleteBuildTriggerParams {
	o.Repository = Repository
	return o
}

// WithTriggerUUID adds the triggerUuid to the delete build trigger params
func (o *DeleteBuildTriggerParams) WithTriggerUUID(TriggerUUID string) *DeleteBuildTriggerParams {
	o.TriggerUUID = TriggerUUID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBuildTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param trigger_uuid
	if err := r.SetPathParam("trigger_uuid", o.TriggerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
