package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteOrgRobotParams creates a new DeleteOrgRobotParams object
// with the default values initialized.
func NewDeleteOrgRobotParams() *DeleteOrgRobotParams {
	var ()
	return &DeleteOrgRobotParams{}
}

/*DeleteOrgRobotParams contains all the parameters to send to the API endpoint
for the delete org robot operation typically these are written to a http.Request
*/
type DeleteOrgRobotParams struct {

	/*Orgname
	  The name of the organization

	*/
	Orgname string
	/*RobotShortname
	  The short name for the robot, without any user or organization prefix

	*/
	RobotShortname string
}

// WithOrgname adds the orgname to the delete org robot params
func (o *DeleteOrgRobotParams) WithOrgname(Orgname string) *DeleteOrgRobotParams {
	o.Orgname = Orgname
	return o
}

// WithRobotShortname adds the robotShortname to the delete org robot params
func (o *DeleteOrgRobotParams) WithRobotShortname(RobotShortname string) *DeleteOrgRobotParams {
	o.RobotShortname = RobotShortname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrgRobotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
