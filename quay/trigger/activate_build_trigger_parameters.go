package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// NewActivateBuildTriggerParams creates a new ActivateBuildTriggerParams object
// with the default values initialized.
func NewActivateBuildTriggerParams() *ActivateBuildTriggerParams {
	var ()
	return &ActivateBuildTriggerParams{}
}

/*ActivateBuildTriggerParams contains all the parameters to send to the API endpoint
for the activate build trigger operation typically these are written to a http.Request
*/
type ActivateBuildTriggerParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.BuildTriggerActivateRequest
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*TriggerUUID
	  The UUID of the build trigger

	*/
	TriggerUUID string
}

// WithBody adds the body to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithBody(body *models.BuildTriggerActivateRequest) *ActivateBuildTriggerParams {
	o.Body = body
	return o
}

// WithRepository adds the repository to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithRepository(repository string) *ActivateBuildTriggerParams {
	o.Repository = repository
	return o
}

// WithTriggerUUID adds the triggerUuid to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithTriggerUUID(triggerUuid string) *ActivateBuildTriggerParams {
	o.TriggerUUID = triggerUuid
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *ActivateBuildTriggerParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.BuildTriggerActivateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

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
