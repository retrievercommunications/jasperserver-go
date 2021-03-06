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

// NewGetCustomDataSourceDefinitionParams creates a new GetCustomDataSourceDefinitionParams object
// with the default values initialized.
func NewGetCustomDataSourceDefinitionParams() *GetCustomDataSourceDefinitionParams {
	var ()
	return &GetCustomDataSourceDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCustomDataSourceDefinitionParamsWithTimeout creates a new GetCustomDataSourceDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCustomDataSourceDefinitionParamsWithTimeout(timeout time.Duration) *GetCustomDataSourceDefinitionParams {
	var ()
	return &GetCustomDataSourceDefinitionParams{

		timeout: timeout,
	}
}

/*GetCustomDataSourceDefinitionParams contains all the parameters to send to the API endpoint
for the get custom data source definition operation typically these are written to a http.Request
*/
type GetCustomDataSourceDefinitionParams struct {

	/*Name*/
	Name *string

	timeout time.Duration
}

// WithName adds the name to the get custom data source definition params
func (o *GetCustomDataSourceDefinitionParams) WithName(Name *string) *GetCustomDataSourceDefinitionParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCustomDataSourceDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Name != nil {

		// path param name
		if err := r.SetPathParam("name", *o.Name); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
