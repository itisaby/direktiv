// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteRegistryParams creates a new DeleteRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteRegistryParams() *DeleteRegistryParams {
	return &DeleteRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRegistryParamsWithTimeout creates a new DeleteRegistryParams object
// with the ability to set a timeout on a request.
func NewDeleteRegistryParamsWithTimeout(timeout time.Duration) *DeleteRegistryParams {
	return &DeleteRegistryParams{
		timeout: timeout,
	}
}

// NewDeleteRegistryParamsWithContext creates a new DeleteRegistryParams object
// with the ability to set a context for a request.
func NewDeleteRegistryParamsWithContext(ctx context.Context) *DeleteRegistryParams {
	return &DeleteRegistryParams{
		Context: ctx,
	}
}

// NewDeleteRegistryParamsWithHTTPClient creates a new DeleteRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteRegistryParamsWithHTTPClient(client *http.Client) *DeleteRegistryParams {
	return &DeleteRegistryParams{
		HTTPClient: client,
	}
}

/* DeleteRegistryParams contains all the parameters to send to the API endpoint
   for the delete registry operation.

   Typically these are written to a http.Request.
*/
type DeleteRegistryParams struct {

	/* RegistryPayload.

	   Payload that contains registry data
	*/
	RegistryPayload DeleteRegistryBody

	/* Namespace.

	   target namespace
	*/
	Namespace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRegistryParams) WithDefaults() *DeleteRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete registry params
func (o *DeleteRegistryParams) WithTimeout(timeout time.Duration) *DeleteRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete registry params
func (o *DeleteRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete registry params
func (o *DeleteRegistryParams) WithContext(ctx context.Context) *DeleteRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete registry params
func (o *DeleteRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete registry params
func (o *DeleteRegistryParams) WithHTTPClient(client *http.Client) *DeleteRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete registry params
func (o *DeleteRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistryPayload adds the registryPayload to the delete registry params
func (o *DeleteRegistryParams) WithRegistryPayload(registryPayload DeleteRegistryBody) *DeleteRegistryParams {
	o.SetRegistryPayload(registryPayload)
	return o
}

// SetRegistryPayload adds the registryPayload to the delete registry params
func (o *DeleteRegistryParams) SetRegistryPayload(registryPayload DeleteRegistryBody) {
	o.RegistryPayload = registryPayload
}

// WithNamespace adds the namespace to the delete registry params
func (o *DeleteRegistryParams) WithNamespace(namespace string) *DeleteRegistryParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the delete registry params
func (o *DeleteRegistryParams) SetNamespace(namespace string) {
	o.Namespace = namespace
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.RegistryPayload); err != nil {
		return err
	}

	// path param namespace
	if err := r.SetPathParam("namespace", o.Namespace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
