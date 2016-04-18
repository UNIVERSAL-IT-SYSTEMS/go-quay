package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// GetTeamPermissionsReader is a Reader for the GetTeamPermissions structure.
type GetTeamPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetTeamPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTeamPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetTeamPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetTeamPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTeamPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetTeamPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTeamPermissionsOK creates a GetTeamPermissionsOK with default headers values
func NewGetTeamPermissionsOK() *GetTeamPermissionsOK {
	return &GetTeamPermissionsOK{}
}

/*GetTeamPermissionsOK handles this case with default header values.

Successful invocation
*/
type GetTeamPermissionsOK struct {
}

func (o *GetTeamPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/{teamname}][%d] getTeamPermissionsOK ", 200)
}

func (o *GetTeamPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTeamPermissionsBadRequest creates a GetTeamPermissionsBadRequest with default headers values
func NewGetTeamPermissionsBadRequest() *GetTeamPermissionsBadRequest {
	return &GetTeamPermissionsBadRequest{}
}

/*GetTeamPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type GetTeamPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *GetTeamPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/{teamname}][%d] getTeamPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTeamPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamPermissionsUnauthorized creates a GetTeamPermissionsUnauthorized with default headers values
func NewGetTeamPermissionsUnauthorized() *GetTeamPermissionsUnauthorized {
	return &GetTeamPermissionsUnauthorized{}
}

/*GetTeamPermissionsUnauthorized handles this case with default header values.

Session required
*/
type GetTeamPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetTeamPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/{teamname}][%d] getTeamPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTeamPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamPermissionsForbidden creates a GetTeamPermissionsForbidden with default headers values
func NewGetTeamPermissionsForbidden() *GetTeamPermissionsForbidden {
	return &GetTeamPermissionsForbidden{}
}

/*GetTeamPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type GetTeamPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *GetTeamPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/{teamname}][%d] getTeamPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *GetTeamPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTeamPermissionsNotFound creates a GetTeamPermissionsNotFound with default headers values
func NewGetTeamPermissionsNotFound() *GetTeamPermissionsNotFound {
	return &GetTeamPermissionsNotFound{}
}

/*GetTeamPermissionsNotFound handles this case with default header values.

Not found
*/
type GetTeamPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *GetTeamPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/team/{teamname}][%d] getTeamPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *GetTeamPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
