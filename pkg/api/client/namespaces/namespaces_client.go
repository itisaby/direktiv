// Code generated by go-swagger; DO NOT EDIT.

package namespaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new namespaces API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for namespaces API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateNamespace(params *CreateNamespaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNamespaceOK, error)

	DeleteNamespace(params *DeleteNamespaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceOK, error)

	GetNamespaceConfig(params *GetNamespaceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceConfigOK, error)

	GetNamespaces(params *GetNamespacesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespacesOK, error)

	SetNamespaceConfig(params *SetNamespaceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetNamespaceConfigOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateNamespace creates a namespace

  Creates a new namespace.

*/
func (a *Client) CreateNamespace(params *CreateNamespaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNamespaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNamespace",
		Method:             "PUT",
		PathPattern:        "/api/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateNamespaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateNamespaceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteNamespace deletes a namespace

  Delete a namespace.

*/
func (a *Client) DeleteNamespace(params *DeleteNamespaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNamespaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNamespaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNamespace",
		Method:             "DELETE",
		PathPattern:        "/api/namespaces/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNamespaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNamespaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNamespaceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetNamespaceConfig gets a namespace config

  Gets a namespace config.

*/
func (a *Client) GetNamespaceConfig(params *GetNamespaceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespaceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespaceConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaceConfig",
		Method:             "GET",
		PathPattern:        "/api/namespaces/{namespace}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNamespaceConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNamespaceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaceConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNamespaces gets the list of namespaces

  Gets the list of namespaces.

*/
func (a *Client) GetNamespaces(params *GetNamespacesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNamespacesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNamespacesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNamespaces",
		Method:             "GET",
		PathPattern:        "/api/namespaces",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetNamespacesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNamespacesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNamespaces: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetNamespaceConfig sets a namespace config

  Sets a namespace config.

*/
func (a *Client) SetNamespaceConfig(params *SetNamespaceConfigParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SetNamespaceConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetNamespaceConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "setNamespaceConfig",
		Method:             "PATCH",
		PathPattern:        "/api/namespaces/{namespace}/config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SetNamespaceConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetNamespaceConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setNamespaceConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}