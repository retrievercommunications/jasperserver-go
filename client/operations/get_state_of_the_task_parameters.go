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

// NewGetStateOfTheTaskParams creates a new GetStateOfTheTaskParams object
// with the default values initialized.
func NewGetStateOfTheTaskParams() *GetStateOfTheTaskParams {
	var ()
	return &GetStateOfTheTaskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStateOfTheTaskParamsWithTimeout creates a new GetStateOfTheTaskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStateOfTheTaskParamsWithTimeout(timeout time.Duration) *GetStateOfTheTaskParams {
	var ()
	return &GetStateOfTheTaskParams{

		timeout: timeout,
	}
}

/*GetStateOfTheTaskParams contains all the parameters to send to the API endpoint
for the get state of the task operation typically these are written to a http.Request
*/
type GetStateOfTheTaskParams struct {

	/*ID*/
	ID *string

	timeout time.Duration
}

// WithID adds the id to the get state of the task params
func (o *GetStateOfTheTaskParams) WithID(ID *string) *GetStateOfTheTaskParams {
	o.ID = ID
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStateOfTheTaskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ID != nil {

		// path param id
		if err := r.SetPathParam("id", *o.ID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
