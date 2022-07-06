// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	"github.com/ory/hydra/internal/httpclient/client/v1"
)

// NewDeleteTrustedJwtGrantIssuerParams creates a new DeleteTrustedJwtGrantIssuerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTrustedJwtGrantIssuerParams() *DeleteTrustedJwtGrantIssuerParams {
	return &DeleteTrustedJwtGrantIssuerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTrustedJwtGrantIssuerParamsWithTimeout creates a new DeleteTrustedJwtGrantIssuerParams object
// with the ability to set a timeout on a request.
func NewDeleteTrustedJwtGrantIssuerParamsWithTimeout(timeout time.Duration) *DeleteTrustedJwtGrantIssuerParams {
	return &DeleteTrustedJwtGrantIssuerParams{
		timeout: timeout,
	}
}

// NewDeleteTrustedJwtGrantIssuerParamsWithContext creates a new DeleteTrustedJwtGrantIssuerParams object
// with the ability to set a context for a request.
func NewDeleteTrustedJwtGrantIssuerParamsWithContext(ctx context.Context) *DeleteTrustedJwtGrantIssuerParams {
	return &DeleteTrustedJwtGrantIssuerParams{
		Context: ctx,
	}
}

// NewDeleteTrustedJwtGrantIssuerParamsWithHTTPClient creates a new DeleteTrustedJwtGrantIssuerParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTrustedJwtGrantIssuerParamsWithHTTPClient(client *http.Client) *DeleteTrustedJwtGrantIssuerParams {
	return &DeleteTrustedJwtGrantIssuerParams{
		HTTPClient: client,
	}
}

/* DeleteTrustedJwtGrantIssuerParams contains all the parameters to send to the API endpoint
   for the delete trusted jwt grant issuer operation.

   Typically these are written to a http.Request.
*/
type DeleteTrustedJwtGrantIssuerParams struct {

	/* ID.

	   The id of the desired grant
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete trusted jwt grant issuer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTrustedJwtGrantIssuerParams) WithDefaults() *DeleteTrustedJwtGrantIssuerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete trusted jwt grant issuer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTrustedJwtGrantIssuerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) WithTimeout(timeout time.Duration) *DeleteTrustedJwtGrantIssuerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) WithContext(ctx context.Context) *DeleteTrustedJwtGrantIssuerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) WithHTTPClient(client *http.Client) *DeleteTrustedJwtGrantIssuerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) WithID(id string) *DeleteTrustedJwtGrantIssuerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete trusted jwt grant issuer params
func (o *DeleteTrustedJwtGrantIssuerParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTrustedJwtGrantIssuerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
