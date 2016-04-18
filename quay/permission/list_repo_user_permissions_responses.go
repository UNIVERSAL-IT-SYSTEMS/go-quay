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

// ListRepoUserPermissionsReader is a Reader for the ListRepoUserPermissions structure.
type ListRepoUserPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ListRepoUserPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListRepoUserPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewListRepoUserPermissionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewListRepoUserPermissionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewListRepoUserPermissionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewListRepoUserPermissionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRepoUserPermissionsOK creates a ListRepoUserPermissionsOK with default headers values
func NewListRepoUserPermissionsOK() *ListRepoUserPermissionsOK {
	return &ListRepoUserPermissionsOK{}
}

/*ListRepoUserPermissionsOK handles this case with default header values.

Successful invocation
*/
type ListRepoUserPermissionsOK struct {
}

func (o *ListRepoUserPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/][%d] listRepoUserPermissionsOK ", 200)
}

func (o *ListRepoUserPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListRepoUserPermissionsBadRequest creates a ListRepoUserPermissionsBadRequest with default headers values
func NewListRepoUserPermissionsBadRequest() *ListRepoUserPermissionsBadRequest {
	return &ListRepoUserPermissionsBadRequest{}
}

/*ListRepoUserPermissionsBadRequest handles this case with default header values.

Bad Request
*/
type ListRepoUserPermissionsBadRequest struct {
	Payload *models.APIError
}

func (o *ListRepoUserPermissionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/][%d] listRepoUserPermissionsBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepoUserPermissionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoUserPermissionsUnauthorized creates a ListRepoUserPermissionsUnauthorized with default headers values
func NewListRepoUserPermissionsUnauthorized() *ListRepoUserPermissionsUnauthorized {
	return &ListRepoUserPermissionsUnauthorized{}
}

/*ListRepoUserPermissionsUnauthorized handles this case with default header values.

Session required
*/
type ListRepoUserPermissionsUnauthorized struct {
	Payload *models.APIError
}

func (o *ListRepoUserPermissionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/][%d] listRepoUserPermissionsUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepoUserPermissionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoUserPermissionsForbidden creates a ListRepoUserPermissionsForbidden with default headers values
func NewListRepoUserPermissionsForbidden() *ListRepoUserPermissionsForbidden {
	return &ListRepoUserPermissionsForbidden{}
}

/*ListRepoUserPermissionsForbidden handles this case with default header values.

Unauthorized access
*/
type ListRepoUserPermissionsForbidden struct {
	Payload *models.APIError
}

func (o *ListRepoUserPermissionsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/][%d] listRepoUserPermissionsForbidden  %+v", 403, o.Payload)
}

func (o *ListRepoUserPermissionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRepoUserPermissionsNotFound creates a ListRepoUserPermissionsNotFound with default headers values
func NewListRepoUserPermissionsNotFound() *ListRepoUserPermissionsNotFound {
	return &ListRepoUserPermissionsNotFound{}
}

/*ListRepoUserPermissionsNotFound handles this case with default header values.

Not found
*/
type ListRepoUserPermissionsNotFound struct {
	Payload *models.APIError
}

func (o *ListRepoUserPermissionsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/repository/{repository}/permissions/user/][%d] listRepoUserPermissionsNotFound  %+v", 404, o.Payload)
}

func (o *ListRepoUserPermissionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
