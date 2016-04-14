package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/coreos/go-quay/models"
)

// CreateOrgRobotReader is a Reader for the CreateOrgRobot structure.
type CreateOrgRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateOrgRobotReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateOrgRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateOrgRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateOrgRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateOrgRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateOrgRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrgRobotOK creates a CreateOrgRobotOK with default headers values
func NewCreateOrgRobotOK() *CreateOrgRobotOK {
	return &CreateOrgRobotOK{}
}

/*CreateOrgRobotOK handles this case with default header values.

Successful invocation
*/
type CreateOrgRobotOK struct {
	Payload *models.Robot
}

func (o *CreateOrgRobotOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] createOrgRobotOK  %+v", 200, o.Payload)
}

func (o *CreateOrgRobotOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Robot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRobotBadRequest creates a CreateOrgRobotBadRequest with default headers values
func NewCreateOrgRobotBadRequest() *CreateOrgRobotBadRequest {
	return &CreateOrgRobotBadRequest{}
}

/*CreateOrgRobotBadRequest handles this case with default header values.

Bad Request
*/
type CreateOrgRobotBadRequest struct {
	Payload *models.APIError
}

func (o *CreateOrgRobotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] createOrgRobotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrgRobotBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRobotUnauthorized creates a CreateOrgRobotUnauthorized with default headers values
func NewCreateOrgRobotUnauthorized() *CreateOrgRobotUnauthorized {
	return &CreateOrgRobotUnauthorized{}
}

/*CreateOrgRobotUnauthorized handles this case with default header values.

Session required
*/
type CreateOrgRobotUnauthorized struct {
	Payload *models.APIError
}

func (o *CreateOrgRobotUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] createOrgRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrgRobotUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRobotForbidden creates a CreateOrgRobotForbidden with default headers values
func NewCreateOrgRobotForbidden() *CreateOrgRobotForbidden {
	return &CreateOrgRobotForbidden{}
}

/*CreateOrgRobotForbidden handles this case with default header values.

Unauthorized access
*/
type CreateOrgRobotForbidden struct {
	Payload *models.APIError
}

func (o *CreateOrgRobotForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] createOrgRobotForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrgRobotForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrgRobotNotFound creates a CreateOrgRobotNotFound with default headers values
func NewCreateOrgRobotNotFound() *CreateOrgRobotNotFound {
	return &CreateOrgRobotNotFound{}
}

/*CreateOrgRobotNotFound handles this case with default header values.

Not found
*/
type CreateOrgRobotNotFound struct {
	Payload *models.APIError
}

func (o *CreateOrgRobotNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/organization/{orgname}/robots/{robot_shortname}][%d] createOrgRobotNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrgRobotNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
