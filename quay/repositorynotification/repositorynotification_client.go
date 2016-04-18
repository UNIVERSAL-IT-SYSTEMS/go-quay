package repositorynotification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repositorynotification API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repositorynotification API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRepoNotification Create a new notification for the specified repository.
*/
func (a *Client) CreateRepoNotification(params *CreateRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateRepoNotificationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRepoNotification",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/notification/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRepoNotificationCreated), nil
}

/*
DeleteRepoNotification Deletes the specified notification.
*/
func (a *Client) DeleteRepoNotification(params *DeleteRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRepoNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRepoNotification",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRepoNotificationNoContent), nil
}

/*
GetRepoNotification Get information for the specified notification.
*/
func (a *Client) GetRepoNotification(params *GetRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*GetRepoNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRepoNotification",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoNotificationOK), nil
}

/*
ListRepoNotifications List the notifications for the specified repository.
*/
func (a *Client) ListRepoNotifications(params *ListRepoNotificationsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepoNotifications",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/notification/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepoNotificationsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepoNotificationsOK), nil
}

/*
TestRepoNotification Queues a test notification for this repository.
*/
func (a *Client) TestRepoNotification(params *TestRepoNotificationParams, authInfo runtime.ClientAuthInfoWriter) (*TestRepoNotificationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestRepoNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testRepoNotification",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/notification/{uuid}/test",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TestRepoNotificationReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestRepoNotificationCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
