package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// CreateUserRobotReader is a Reader for the CreateUserRobot structure.
type CreateUserRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateUserRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateUserRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateUserRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateUserRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateUserRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateUserRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserRobotOK creates a CreateUserRobotOK with default headers values
func NewCreateUserRobotOK() *CreateUserRobotOK {
	return &CreateUserRobotOK{}
}

/*CreateUserRobotOK handles this case with default header values.

Successful invocation
*/
type CreateUserRobotOK struct {
}

func (o *CreateUserRobotOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/user/robots/{robot_shortname}][%d] createUserRobotOK ", 200)
}

func (o *CreateUserRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserRobotBadRequest creates a CreateUserRobotBadRequest with default headers values
func NewCreateUserRobotBadRequest() *CreateUserRobotBadRequest {
	return &CreateUserRobotBadRequest{}
}

/*CreateUserRobotBadRequest handles this case with default header values.

Bad Request
*/
type CreateUserRobotBadRequest struct {
	Payload *models.APIError
}

func (o *CreateUserRobotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v1/user/robots/{robot_shortname}][%d] createUserRobotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserRobotUnauthorized creates a CreateUserRobotUnauthorized with default headers values
func NewCreateUserRobotUnauthorized() *CreateUserRobotUnauthorized {
	return &CreateUserRobotUnauthorized{}
}

/*CreateUserRobotUnauthorized handles this case with default header values.

Session required
*/
type CreateUserRobotUnauthorized struct {
	Payload *models.APIError
}

func (o *CreateUserRobotUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v1/user/robots/{robot_shortname}][%d] createUserRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateUserRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserRobotForbidden creates a CreateUserRobotForbidden with default headers values
func NewCreateUserRobotForbidden() *CreateUserRobotForbidden {
	return &CreateUserRobotForbidden{}
}

/*CreateUserRobotForbidden handles this case with default header values.

Unauthorized access
*/
type CreateUserRobotForbidden struct {
	Payload *models.APIError
}

func (o *CreateUserRobotForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/user/robots/{robot_shortname}][%d] createUserRobotForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserRobotNotFound creates a CreateUserRobotNotFound with default headers values
func NewCreateUserRobotNotFound() *CreateUserRobotNotFound {
	return &CreateUserRobotNotFound{}
}

/*CreateUserRobotNotFound handles this case with default header values.

Not found
*/
type CreateUserRobotNotFound struct {
	Payload *models.APIError
}

func (o *CreateUserRobotNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/user/robots/{robot_shortname}][%d] createUserRobotNotFound  %+v", 404, o.Payload)
}

func (o *CreateUserRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
