// Code generated by go-swagger; DO NOT EDIT.

package v1

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
	"github.com/ory/hydra/internal/httpclient/models"
)

// NewAdminUpdateJSONWebKeySetParams creates a new AdminUpdateJSONWebKeySetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminUpdateJSONWebKeySetParams() *AdminUpdateJSONWebKeySetParams {
	return &AdminUpdateJSONWebKeySetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminUpdateJSONWebKeySetParamsWithTimeout creates a new AdminUpdateJSONWebKeySetParams object
// with the ability to set a timeout on a request.
func NewAdminUpdateJSONWebKeySetParamsWithTimeout(timeout time.Duration) *AdminUpdateJSONWebKeySetParams {
	return &AdminUpdateJSONWebKeySetParams{
		timeout: timeout,
	}
}

// NewAdminUpdateJSONWebKeySetParamsWithContext creates a new AdminUpdateJSONWebKeySetParams object
// with the ability to set a context for a request.
func NewAdminUpdateJSONWebKeySetParamsWithContext(ctx context.Context) *AdminUpdateJSONWebKeySetParams {
	return &AdminUpdateJSONWebKeySetParams{
		Context: ctx,
	}
}

// NewAdminUpdateJSONWebKeySetParamsWithHTTPClient creates a new AdminUpdateJSONWebKeySetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminUpdateJSONWebKeySetParamsWithHTTPClient(client *http.Client) *AdminUpdateJSONWebKeySetParams {
	return &AdminUpdateJSONWebKeySetParams{
		HTTPClient: client,
	}
}

/* AdminUpdateJSONWebKeySetParams contains all the parameters to send to the API endpoint
   for the admin update Json web key set operation.

   Typically these are written to a http.Request.
*/
type AdminUpdateJSONWebKeySetParams struct {

	// Body.
	Body *models.JSONWebKeySet

	/* Set.

	   The JSON Web Key Set
	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin update Json web key set params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateJSONWebKeySetParams) WithDefaults() *AdminUpdateJSONWebKeySetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin update Json web key set params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminUpdateJSONWebKeySetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) WithTimeout(timeout time.Duration) *AdminUpdateJSONWebKeySetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) WithContext(ctx context.Context) *AdminUpdateJSONWebKeySetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) WithHTTPClient(client *http.Client) *AdminUpdateJSONWebKeySetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) WithBody(body *models.JSONWebKeySet) *AdminUpdateJSONWebKeySetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) SetBody(body *models.JSONWebKeySet) {
	o.Body = body
}

// WithSet adds the set to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) WithSet(set string) *AdminUpdateJSONWebKeySetParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the admin update Json web key set params
func (o *AdminUpdateJSONWebKeySetParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *AdminUpdateJSONWebKeySetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
