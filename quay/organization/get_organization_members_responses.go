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

// GetOrganizationMembersReader is a Reader for the GetOrganizationMembers structure.
type GetOrganizationMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetOrganizationMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetOrganizationMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetOrganizationMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetOrganizationMembersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetOrganizationMembersOK creates a GetOrganizationMembersOK with default headers values
func NewGetOrganizationMembersOK() *GetOrganizationMembersOK {
	return &GetOrganizationMembersOK{}
}

/*GetOrganizationMembersOK handles this case with default header values.

Successful invocation
*/
type GetOrganizationMembersOK struct {
}

func (o *GetOrganizationMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/members][%d] getOrganizationMembersOK ", 200)
}

func (o *GetOrganizationMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationMembersBadRequest creates a GetOrganizationMembersBadRequest with default headers values
func NewGetOrganizationMembersBadRequest() *GetOrganizationMembersBadRequest {
	return &GetOrganizationMembersBadRequest{}
}

/*GetOrganizationMembersBadRequest handles this case with default header values.

Bad Request
*/
type GetOrganizationMembersBadRequest struct {
	Payload *models.APIError
}

func (o *GetOrganizationMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/members][%d] getOrganizationMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrganizationMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersUnauthorized creates a GetOrganizationMembersUnauthorized with default headers values
func NewGetOrganizationMembersUnauthorized() *GetOrganizationMembersUnauthorized {
	return &GetOrganizationMembersUnauthorized{}
}

/*GetOrganizationMembersUnauthorized handles this case with default header values.

Session required
*/
type GetOrganizationMembersUnauthorized struct {
	Payload *models.APIError
}

func (o *GetOrganizationMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/members][%d] getOrganizationMembersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersForbidden creates a GetOrganizationMembersForbidden with default headers values
func NewGetOrganizationMembersForbidden() *GetOrganizationMembersForbidden {
	return &GetOrganizationMembersForbidden{}
}

/*GetOrganizationMembersForbidden handles this case with default header values.

Unauthorized access
*/
type GetOrganizationMembersForbidden struct {
	Payload *models.APIError
}

func (o *GetOrganizationMembersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/members][%d] getOrganizationMembersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrganizationMembersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationMembersNotFound creates a GetOrganizationMembersNotFound with default headers values
func NewGetOrganizationMembersNotFound() *GetOrganizationMembersNotFound {
	return &GetOrganizationMembersNotFound{}
}

/*GetOrganizationMembersNotFound handles this case with default header values.

Not found
*/
type GetOrganizationMembersNotFound struct {
	Payload *models.APIError
}

func (o *GetOrganizationMembersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/organization/{orgname}/members][%d] getOrganizationMembersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
