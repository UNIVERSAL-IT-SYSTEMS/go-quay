package secscan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new secscan API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secscan API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
GetRepoImageSecurity Fetches the features and vulnerabilities (if any) for a repository tag.
*/
func (a *Client) GetRepoImageSecurity(params *GetRepoImageSecurityParams, authInfo client.AuthInfoWriter) (*GetRepoImageSecurityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoImageSecurityParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getRepoImageSecurity",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/image/{imageid}/security",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRepoImageSecurityReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoImageSecurityOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}
