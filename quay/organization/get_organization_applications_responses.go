package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// GetOrganizationApplicationsReader is a Reader for the GetOrganizationApplications structure.
type GetOrganizationApplicationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetOrganizationApplicationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationApplicationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganizationApplicationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrganizationApplicationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationApplicationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationApplicationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationApplicationsOK creates a GetOrganizationApplicationsOK with default headers values
func NewGetOrganizationApplicationsOK() *GetOrganizationApplicationsOK {
	return &GetOrganizationApplicationsOK{}
}

/*GetOrganizationApplicationsOK handles this case with default header values.

Successful invocation
*/
type GetOrganizationApplicationsOK struct {
}

func (o *GetOrganizationApplicationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications][%d] getOrganizationApplicationsOK ", 200)
}

func (o *GetOrganizationApplicationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationApplicationsBadRequest creates a GetOrganizationApplicationsBadRequest with default headers values
func NewGetOrganizationApplicationsBadRequest() *GetOrganizationApplicationsBadRequest {
	return &GetOrganizationApplicationsBadRequest{}
}

/*GetOrganizationApplicationsBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganizationApplicationsBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications][%d] getOrganizationApplicationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationApplicationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationsUnauthorized creates a GetOrganizationApplicationsUnauthorized with default headers values
func NewGetOrganizationApplicationsUnauthorized() *GetOrganizationApplicationsUnauthorized {
	return &GetOrganizationApplicationsUnauthorized{}
}

/*GetOrganizationApplicationsUnauthorized handles this case with default header values.

Session required
*/
type GetOrganizationApplicationsUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications][%d] getOrganizationApplicationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationApplicationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationsForbidden creates a GetOrganizationApplicationsForbidden with default headers values
func NewGetOrganizationApplicationsForbidden() *GetOrganizationApplicationsForbidden {
	return &GetOrganizationApplicationsForbidden{}
}

/*GetOrganizationApplicationsForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrganizationApplicationsForbidden struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications][%d] getOrganizationApplicationsForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationApplicationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationApplicationsNotFound creates a GetOrganizationApplicationsNotFound with default headers values
func NewGetOrganizationApplicationsNotFound() *GetOrganizationApplicationsNotFound {
	return &GetOrganizationApplicationsNotFound{}
}

/*GetOrganizationApplicationsNotFound handles this case with default header values.

Not found
*/
type GetOrganizationApplicationsNotFound struct {
	Payload *models.APIError
}

func (o *GetOrganizationApplicationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/applications][%d] getOrganizationApplicationsNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationApplicationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
