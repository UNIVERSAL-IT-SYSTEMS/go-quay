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

// DeleteUserPermissionsReader is a Reader for the DeleteUserPermissions structure.
type DeleteUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserPermissionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserPermissionsNoContent creates a DeleteUserPermissionsNoContent with default headers values
func NewDeleteUserPermissionsNoContent() *DeleteUserPermissionsNoContent {
	return &DeleteUserPermissionsNoContent{}
}

/*DeleteUserPermissionsNoContent handles this case with default header values.

Deleted
*/
type DeleteUserPermissionsNoContent struct {
}

func (o *DeleteUserPermissionsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsNoContent ", 204)
}

func (o *DeleteUserPermissionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserPermissionsBadRequest creates a DeleteUserPermissionsBadRequest with default headers values
func NewDeleteUserPermissionsBadRequest() *DeleteUserPermissionsBadRequest {
	return &DeleteUserPermissionsBadRequest{}
}

/*DeleteUserPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type DeleteUserPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteUserPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserPermissionsUnauthorized creates a DeleteUserPermissionsUnauthorized with default headers values
func NewDeleteUserPermissionsUnauthorized() *DeleteUserPermissionsUnauthorized {
	return &DeleteUserPermissionsUnauthorized{}
}

/*DeleteUserPermissionsUnauthorized handles this case with default header values.

Session required
*/
type DeleteUserPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserPermissionsForbidden creates a DeleteUserPermissionsForbidden with default headers values
func NewDeleteUserPermissionsForbidden() *DeleteUserPermissionsForbidden {
	return &DeleteUserPermissionsForbidden{}
}

/*DeleteUserPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteUserPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *DeleteUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *DeleteUserPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteUserPermissionsNotFound creates a DeleteUserPermissionsNotFound with default headers values
func NewDeleteUserPermissionsNotFound() *DeleteUserPermissionsNotFound {
	return &DeleteUserPermissionsNotFound{}
}

/*DeleteUserPermissionsNotFound handles this case with default header values.

Not found
*/
type DeleteUserPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *DeleteUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/permissions/user/{username}][%d] deleteUserPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *DeleteUserPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
