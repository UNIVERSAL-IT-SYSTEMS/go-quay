package prototype

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// CreateOrganizationPrototypePermissionReader is a Reader for the CreateOrganizationPrototypePermission structure.
type CreateOrganizationPrototypePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateOrganizationPrototypePermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateOrganizationPrototypePermissionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateOrganizationPrototypePermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateOrganizationPrototypePermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateOrganizationPrototypePermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateOrganizationPrototypePermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateOrganizationPrototypePermissionCreated creates a CreateOrganizationPrototypePermissionCreated with default headers values
func NewCreateOrganizationPrototypePermissionCreated() *CreateOrganizationPrototypePermissionCreated {
	return &CreateOrganizationPrototypePermissionCreated{}
}

/*CreateOrganizationPrototypePermissionCreated handles this case with default header values.

Successful creation
*/
type CreateOrganizationPrototypePermissionCreated struct {
}

func (o *CreateOrganizationPrototypePermissionCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/prototypes][%d] createOrganizationPrototypePermissionCreated ", 201)
}

func (o *CreateOrganizationPrototypePermissionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateOrganizationPrototypePermissionBadRequest creates a CreateOrganizationPrototypePermissionBadRequest with default headers values
func NewCreateOrganizationPrototypePermissionBadRequest() *CreateOrganizationPrototypePermissionBadRequest {
	return &CreateOrganizationPrototypePermissionBadRequest{}
}

/*CreateOrganizationPrototypePermissionBadRequest handles this case with default header values.

Bad Request
*/
type CreateOrganizationPrototypePermissionBadRequest struct {
	Payload *models.APIError
}

func (o *CreateOrganizationPrototypePermissionBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/prototypes][%d] createOrganizationPrototypePermissionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateOrganizationPrototypePermissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationPrototypePermissionUnauthorized creates a CreateOrganizationPrototypePermissionUnauthorized with default headers values
func NewCreateOrganizationPrototypePermissionUnauthorized() *CreateOrganizationPrototypePermissionUnauthorized {
	return &CreateOrganizationPrototypePermissionUnauthorized{}
}

/*CreateOrganizationPrototypePermissionUnauthorized handles this case with default header values.

Session required
*/
type CreateOrganizationPrototypePermissionUnauthorized struct {
	Payload *models.APIError
}

func (o *CreateOrganizationPrototypePermissionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/prototypes][%d] createOrganizationPrototypePermissionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateOrganizationPrototypePermissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationPrototypePermissionForbidden creates a CreateOrganizationPrototypePermissionForbidden with default headers values
func NewCreateOrganizationPrototypePermissionForbidden() *CreateOrganizationPrototypePermissionForbidden {
	return &CreateOrganizationPrototypePermissionForbidden{}
}

/*CreateOrganizationPrototypePermissionForbidden handles this case with default header values.

Unauthorized access
*/
type CreateOrganizationPrototypePermissionForbidden struct {
	Payload *models.APIError
}

func (o *CreateOrganizationPrototypePermissionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/prototypes][%d] createOrganizationPrototypePermissionForbidden  %+v", 403, o.Payload)
}

func (o *CreateOrganizationPrototypePermissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateOrganizationPrototypePermissionNotFound creates a CreateOrganizationPrototypePermissionNotFound with default headers values
func NewCreateOrganizationPrototypePermissionNotFound() *CreateOrganizationPrototypePermissionNotFound {
	return &CreateOrganizationPrototypePermissionNotFound{}
}

/*CreateOrganizationPrototypePermissionNotFound handles this case with default header values.

Not found
*/
type CreateOrganizationPrototypePermissionNotFound struct {
	Payload *models.APIError
}

func (o *CreateOrganizationPrototypePermissionNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/organization/{orgname}/prototypes][%d] createOrganizationPrototypePermissionNotFound  %+v", 404, o.Payload)
}

func (o *CreateOrganizationPrototypePermissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
