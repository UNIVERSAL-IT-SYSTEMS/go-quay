package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRepoBuildsParams creates a new GetRepoBuildsParams object
// with the default values initialized.
func NewGetRepoBuildsParams() *GetRepoBuildsParams {
	var ()
	return &GetRepoBuildsParams{}
}

/*GetRepoBuildsParams contains all the parameters to send to the API endpoint
for the get repo builds operation typically these are written to a http.Request
*/
type GetRepoBuildsParams struct {

	/*Limit
	  The maximum number of builds to return

	*/
	Limit *int64
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*Since
	  Returns all builds since the given unix timecode

	*/
	Since *int64
}

// WithLimit adds the limit to the get repo builds params
func (o *GetRepoBuildsParams) WithLimit(Limit *int64) *GetRepoBuildsParams {
	o.Limit = Limit
	return o
}

// WithRepository adds the repository to the get repo builds params
func (o *GetRepoBuildsParams) WithRepository(Repository string) *GetRepoBuildsParams {
	o.Repository = Repository
	return o
}

// WithSince adds the since to the get repo builds params
func (o *GetRepoBuildsParams) WithSince(Since *int64) *GetRepoBuildsParams {
	o.Since = Since
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepoBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
