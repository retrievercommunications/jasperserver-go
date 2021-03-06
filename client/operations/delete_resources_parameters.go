package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteResourcesParams creates a new DeleteResourcesParams object
// with the default values initialized.
func NewDeleteResourcesParams() *DeleteResourcesParams {
	var ()
	return &DeleteResourcesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResourcesParamsWithTimeout creates a new DeleteResourcesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteResourcesParamsWithTimeout(timeout time.Duration) *DeleteResourcesParams {
	var ()
	return &DeleteResourcesParams{

		timeout: timeout,
	}
}

/*DeleteResourcesParams contains all the parameters to send to the API endpoint
for the delete resources operation typically these are written to a http.Request
*/
type DeleteResourcesParams struct {

	/*ResourceURI*/
	ResourceURI *string

	timeout time.Duration
}

// WithResourceURI adds the resourceUri to the delete resources params
func (o *DeleteResourcesParams) WithResourceURI(ResourceURI *string) *DeleteResourcesParams {
	o.ResourceURI = ResourceURI
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ResourceURI != nil {

		// query param resourceUri
		var qrResourceURI string
		if o.ResourceURI != nil {
			qrResourceURI = *o.ResourceURI
		}
		qResourceURI := qrResourceURI
		if qResourceURI != "" {
			if err := r.SetQueryParam("resourceUri", qResourceURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
