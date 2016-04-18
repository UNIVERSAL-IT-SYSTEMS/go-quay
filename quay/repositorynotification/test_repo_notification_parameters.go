package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTestRepoNotificationParams creates a new TestRepoNotificationParams object
// with the default values initialized.
func NewTestRepoNotificationParams() *TestRepoNotificationParams {
	var ()
	return &TestRepoNotificationParams{}
}

/*TestRepoNotificationParams contains all the parameters to send to the API endpoint
for the test repo notification operation typically these are written to a http.Request
*/
type TestRepoNotificationParams struct {

	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*UUID
	  The UUID of the notification

	*/
	UUID string
}

// WithRepository adds the repository to the test repo notification params
func (o *TestRepoNotificationParams) WithRepository(Repository string) *TestRepoNotificationParams {
	o.Repository = Repository
	return o
}

// WithUUID adds the uuid to the test repo notification params
func (o *TestRepoNotificationParams) WithUUID(UUID string) *TestRepoNotificationParams {
	o.UUID = UUID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TestRepoNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
