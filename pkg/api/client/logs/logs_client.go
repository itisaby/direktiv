// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new logs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for logs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetWorkflowLogs(params *GetWorkflowLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkflowLogsOK, error)

	InstanceLogs(params *InstanceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InstanceLogsOK, error)

	NamespaceLogs(params *NamespaceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NamespaceLogsOK, error)

	ServerLogs(params *ServerLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerLogsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetWorkflowLogs gets workflow level logs

  Get workflow level logs.

*/
func (a *Client) GetWorkflowLogs(params *GetWorkflowLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetWorkflowLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getWorkflowLogs",
		Method:             "GET",
		PathPattern:        "/api/namespaces/{namespace}/tree/{workflow}?op=logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkflowLogsReader{formats: a.formats},
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
	success, ok := result.(*GetWorkflowLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getWorkflowLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InstanceLogs gets instance logs

  Gets the logs of an executed instance.

*/
func (a *Client) InstanceLogs(params *InstanceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InstanceLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInstanceLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "instanceLogs",
		Method:             "GET",
		PathPattern:        "/api/namespaces/{namespace}/instances/{instance}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InstanceLogsReader{formats: a.formats},
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
	success, ok := result.(*InstanceLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*InstanceLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  NamespaceLogs gets namespace level logs

  Gets Namespace Level Logs.

*/
func (a *Client) NamespaceLogs(params *NamespaceLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*NamespaceLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNamespaceLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "namespaceLogs",
		Method:             "GET",
		PathPattern:        "/api/namespaces/{namespace}/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NamespaceLogsReader{formats: a.formats},
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
	success, ok := result.(*NamespaceLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for namespaceLogs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServerLogs gets direktiv server logs

  Gets Direktiv Server Logs.

*/
func (a *Client) ServerLogs(params *ServerLogsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerLogsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serverLogs",
		Method:             "GET",
		PathPattern:        "/api/logs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerLogsReader{formats: a.formats},
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
	success, ok := result.(*ServerLogsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerLogsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}