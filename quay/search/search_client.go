package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConductSearch Get a list of entities and resources that match the specified query.
*/
func (a *Client) ConductSearch(params *ConductSearchParams, authInfo runtime.ClientAuthInfoWriter) (*ConductSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConductSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "conductSearch",
		Method:             "GET",
		PathPattern:        "/api/v1/find/all",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConductSearchReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConductSearchOK), nil
}

/*
GetMatchingEntities Get a list of entities that match the specified prefix.
*/
func (a *Client) GetMatchingEntities(params *GetMatchingEntitiesParams) (*GetMatchingEntitiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMatchingEntitiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMatchingEntities",
		Method:             "GET",
		PathPattern:        "/api/v1/entities/{prefix}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMatchingEntitiesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMatchingEntitiesOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
