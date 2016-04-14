package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetAggregateOrgLogsParams creates a new GetAggregateOrgLogsParams object
// with the default values initialized.
func NewGetAggregateOrgLogsParams() *GetAggregateOrgLogsParams {
	var ()
	return &GetAggregateOrgLogsParams{}
}

/*GetAggregateOrgLogsParams contains all the parameters to send to the API endpoint
for the get aggregate org logs operation typically these are written to a http.Request
*/
type GetAggregateOrgLogsParams struct {

	/*Endtime
	  Latest time to which to get logs. (%m/%d/%Y %Z)

	*/
	Endtime *string
	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*Performer
	  Username for which to filter logs.

	*/
	Performer *string
	/*Starttime
	  Earliest time from which to get logs. (%m/%d/%Y %Z)

	*/
	Starttime *string
}

// WithEndtime adds the endtime to the get aggregate org logs params
func (o *GetAggregateOrgLogsParams) WithEndtime(endtime *string) *GetAggregateOrgLogsParams {
	o.Endtime = endtime
	return o
}

// WithOrgname adds the orgname to the get aggregate org logs params
func (o *GetAggregateOrgLogsParams) WithOrgname(orgname string) *GetAggregateOrgLogsParams {
	o.Orgname = orgname
	return o
}

// WithPerformer adds the performer to the get aggregate org logs params
func (o *GetAggregateOrgLogsParams) WithPerformer(performer *string) *GetAggregateOrgLogsParams {
	o.Performer = performer
	return o
}

// WithStarttime adds the starttime to the get aggregate org logs params
func (o *GetAggregateOrgLogsParams) WithStarttime(starttime *string) *GetAggregateOrgLogsParams {
	o.Starttime = starttime
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAggregateOrgLogsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Endtime != nil {

		// query param endtime
		var qrEndtime string
		if o.Endtime != nil {
			qrEndtime = *o.Endtime
		}
		qEndtime := qrEndtime
		if qEndtime != "" {
			if err := r.SetQueryParam("endtime", qEndtime); err != nil {
				return err
			}
		}

	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if o.Performer != nil {

		// query param performer
		var qrPerformer string
		if o.Performer != nil {
			qrPerformer = *o.Performer
		}
		qPerformer := qrPerformer
		if qPerformer != "" {
			if err := r.SetQueryParam("performer", qPerformer); err != nil {
				return err
			}
		}

	}

	if o.Starttime != nil {

		// query param starttime
		var qrStarttime string
		if o.Starttime != nil {
			qrStarttime = *o.Starttime
		}
		qStarttime := qrStarttime
		if qStarttime != "" {
			if err := r.SetQueryParam("starttime", qStarttime); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
