package repotoken

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/coreos/go-quay/models"
)

// CreateTokenReader is a Reader for the CreateToken structure.
type CreateTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateTokenCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateTokenBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateTokenCreated creates a CreateTokenCreated with default headers values
func NewCreateTokenCreated() *CreateTokenCreated {
	return &CreateTokenCreated{}
}

/*CreateTokenCreated handles this case with default header values.

Successful creation
*/
type CreateTokenCreated struct {
}

func (o *CreateTokenCreated) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenCreated ", 201)
}

func (o *CreateTokenCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateTokenBadRequest creates a CreateTokenBadRequest with default headers values
func NewCreateTokenBadRequest() *CreateTokenBadRequest {
	return &CreateTokenBadRequest{}
}

/*CreateTokenBadRequest handles this case with default header values.

Bad Request
*/
type CreateTokenBadRequest struct {
	Payload *models.APIError
}

func (o *CreateTokenBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenBadRequest  %+v", 400, o.Payload)
}

func (o *CreateTokenBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenUnauthorized creates a CreateTokenUnauthorized with default headers values
func NewCreateTokenUnauthorized() *CreateTokenUnauthorized {
	return &CreateTokenUnauthorized{}
}

/*CreateTokenUnauthorized handles this case with default header values.

Session required
*/
type CreateTokenUnauthorized struct {
	Payload *models.APIError
}

func (o *CreateTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenForbidden creates a CreateTokenForbidden with default headers values
func NewCreateTokenForbidden() *CreateTokenForbidden {
	return &CreateTokenForbidden{}
}

/*CreateTokenForbidden handles this case with default header values.

Unauthorized access
*/
type CreateTokenForbidden struct {
	Payload *models.APIError
}

func (o *CreateTokenForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenForbidden  %+v", 403, o.Payload)
}

func (o *CreateTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTokenNotFound creates a CreateTokenNotFound with default headers values
func NewCreateTokenNotFound() *CreateTokenNotFound {
	return &CreateTokenNotFound{}
}

/*CreateTokenNotFound handles this case with default header values.

Not found
*/
type CreateTokenNotFound struct {
	Payload *models.APIError
}

func (o *CreateTokenNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/repository/{repository}/tokens/][%d] createTokenNotFound  %+v", 404, o.Payload)
}

func (o *CreateTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
