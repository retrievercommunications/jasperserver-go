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

// NewDeleteAttributesParams creates a new DeleteAttributesParams object
// with the default values initialized.
func NewDeleteAttributesParams() *DeleteAttributesParams {
	var ()
	return &DeleteAttributesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAttributesParamsWithTimeout creates a new DeleteAttributesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAttributesParamsWithTimeout(timeout time.Duration) *DeleteAttributesParams {
	var ()
	return &DeleteAttributesParams{

		timeout: timeout,
	}
}

/*DeleteAttributesParams contains all the parameters to send to the API endpoint
for the delete attributes operation typically these are written to a http.Request
*/
type DeleteAttributesParams struct {

	/*Name*/
	Name *string

	timeout time.Duration
}

// WithName adds the name to the delete attributes params
func (o *DeleteAttributesParams) WithName(Name *string) *DeleteAttributesParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAttributesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
