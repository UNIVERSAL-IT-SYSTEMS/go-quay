package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteUserRobotParams creates a new DeleteUserRobotParams object
// with the default values initialized.
func NewDeleteUserRobotParams() *DeleteUserRobotParams {
	var ()
	return &DeleteUserRobotParams{}
}

/*DeleteUserRobotParams contains all the parameters to send to the API endpoint
for the delete user robot operation typically these are written to a http.Request
*/
type DeleteUserRobotParams struct {

	/*RobotShortname
	  The short name for the robot, without any user or organization prefix

	*/
	RobotShortname string
}

// WithRobotShortname adds the robotShortname to the delete user robot params
func (o *DeleteUserRobotParams) WithRobotShortname(robotShortname string) *DeleteUserRobotParams {
	o.RobotShortname = robotShortname
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserRobotParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param robot_shortname
	if err := r.SetPathParam("robot_shortname", o.RobotShortname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
