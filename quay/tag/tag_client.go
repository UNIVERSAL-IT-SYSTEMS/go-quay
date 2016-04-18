package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tag API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeTagImage Change which image a tag points to or create a new tag.
*/
func (a *Client) ChangeTagImage(params *ChangeTagImageParams, authInfo runtime.ClientAuthInfoWriter) (*ChangeTagImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTagImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeTagImage",
		Method:             "PUT",
		PathPattern:        "/api/v1/repository/{repository}/tag/{tag}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ChangeTagImageReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTagImageOK), nil
}

/*
DeleteFullTag Delete the specified repository tag.
*/
func (a *Client) DeleteFullTag(params *DeleteFullTagParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFullTagNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFullTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteFullTag",
		Method:             "DELETE",
		PathPattern:        "/api/v1/repository/{repository}/tag/{tag}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteFullTagReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFullTagNoContent), nil
}

/*
ListRepoTags list repo tags API
*/
func (a *Client) ListRepoTags(params *ListRepoTagsParams, authInfo runtime.ClientAuthInfoWriter) (*ListRepoTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRepoTagsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRepoTags",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/tag/",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRepoTagsReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRepoTagsOK), nil
}

/*
ListTagImages List the images for the specified repository tag.
*/
func (a *Client) ListTagImages(params *ListTagImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListTagImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTagImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listTagImages",
		Method:             "GET",
		PathPattern:        "/api/v1/repository/{repository}/tag/{tag}/images",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTagImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListTagImagesOK), nil
}

/*
RevertTag Reverts a repository tag back to a previous image in the repository.
*/
func (a *Client) RevertTag(params *RevertTagParams, authInfo runtime.ClientAuthInfoWriter) (*RevertTagCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevertTagParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "revertTag",
		Method:             "POST",
		PathPattern:        "/api/v1/repository/{repository}/tag/{tag}/revert",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RevertTagReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RevertTagCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
