package trigger

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

// DeleteBuildTriggerReader is a Reader for the DeleteBuildTrigger structure.
type DeleteBuildTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteBuildTriggerReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteBuildTriggerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteBuildTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteBuildTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteBuildTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteBuildTriggerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBuildTriggerNoContent creates a DeleteBuildTriggerNoContent with default headers values
func NewDeleteBuildTriggerNoContent() *DeleteBuildTriggerNoContent {
	return &DeleteBuildTriggerNoContent{}
}

/*DeleteBuildTriggerNoContent handles this case with default header values.

Deleted
*/
type DeleteBuildTriggerNoContent struct {
}

func (o *DeleteBuildTriggerNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/trigger/{trigger_uuid}][%d] deleteBuildTriggerNoContent ", 204)
}

func (o *DeleteBuildTriggerNoContent) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBuildTriggerBadRequest creates a DeleteBuildTriggerBadRequest with default headers values
func NewDeleteBuildTriggerBadRequest() *DeleteBuildTriggerBadRequest {
	return &DeleteBuildTriggerBadRequest{}
}

/*DeleteBuildTriggerBadRequest handles this case with default header values.

Bad Request
*/
type DeleteBuildTriggerBadRequest struct {
	Payload *models.APIError
}

func (o *DeleteBuildTriggerBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/trigger/{trigger_uuid}][%d] deleteBuildTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteBuildTriggerBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBuildTriggerUnauthorized creates a DeleteBuildTriggerUnauthorized with default headers values
func NewDeleteBuildTriggerUnauthorized() *DeleteBuildTriggerUnauthorized {
	return &DeleteBuildTriggerUnauthorized{}
}

/*DeleteBuildTriggerUnauthorized handles this case with default header values.

Session required
*/
type DeleteBuildTriggerUnauthorized struct {
	Payload *models.APIError
}

func (o *DeleteBuildTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/trigger/{trigger_uuid}][%d] deleteBuildTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteBuildTriggerUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBuildTriggerForbidden creates a DeleteBuildTriggerForbidden with default headers values
func NewDeleteBuildTriggerForbidden() *DeleteBuildTriggerForbidden {
	return &DeleteBuildTriggerForbidden{}
}

/*DeleteBuildTriggerForbidden handles this case with default header values.

Unauthorized access
*/
type DeleteBuildTriggerForbidden struct {
	Payload *models.APIError
}

func (o *DeleteBuildTriggerForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/trigger/{trigger_uuid}][%d] deleteBuildTriggerForbidden  %+v", 403, o.Payload)
}

func (o *DeleteBuildTriggerForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBuildTriggerNotFound creates a DeleteBuildTriggerNotFound with default headers values
func NewDeleteBuildTriggerNotFound() *DeleteBuildTriggerNotFound {
	return &DeleteBuildTriggerNotFound{}
}

/*DeleteBuildTriggerNotFound handles this case with default header values.

Not found
*/
type DeleteBuildTriggerNotFound struct {
	Payload *models.APIError
}

func (o *DeleteBuildTriggerNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/repository/{repository}/trigger/{trigger_uuid}][%d] deleteBuildTriggerNotFound  %+v", 404, o.Payload)
}

func (o *DeleteBuildTriggerNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
