package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteRepoNotificationParams creates a new DeleteRepoNotificationParams object
// with the default values initialized.
func NewDeleteRepoNotificationParams() *DeleteRepoNotificationParams {
	var ()
	return &DeleteRepoNotificationParams{}
}

/*DeleteRepoNotificationParams contains all the parameters to send to the API endpoint
for the delete repo notification operation typically these are written to a http.Request
*/
type DeleteRepoNotificationParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*UUID
	  The UUID of the notification

	*/
	UUID string
}

// WithRepository adds the repository to the delete repo notification params
func (o *DeleteRepoNotificationParams) WithRepository(Repository string) *DeleteRepoNotificationParams {
	o.Repository = Repository
	return o
}

// WithUUID adds the uuid to the delete repo notification params
func (o *DeleteRepoNotificationParams) WithUUID(UUID string) *DeleteRepoNotificationParams {
	o.UUID = UUID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRepoNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
