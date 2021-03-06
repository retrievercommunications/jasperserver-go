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

// NewGetAttributesOfUserParams creates a new GetAttributesOfUserParams object
// with the default values initialized.
func NewGetAttributesOfUserParams() *GetAttributesOfUserParams {
	var ()
	return &GetAttributesOfUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAttributesOfUserParamsWithTimeout creates a new GetAttributesOfUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAttributesOfUserParamsWithTimeout(timeout time.Duration) *GetAttributesOfUserParams {
	var ()
	return &GetAttributesOfUserParams{

		timeout: timeout,
	}
}

/*GetAttributesOfUserParams contains all the parameters to send to the API endpoint
for the get attributes of user operation typically these are written to a http.Request
*/
type GetAttributesOfUserParams struct {

	/*Accept*/
	Accept *string
	/*Embedded*/
	Embedded *string
	/*Name*/
	Name *string

	timeout time.Duration
}

// WithAccept adds the accept to the get attributes of user params
func (o *GetAttributesOfUserParams) WithAccept(Accept *string) *GetAttributesOfUserParams {
	o.Accept = Accept
	return o
}

// WithEmbedded adds the embedded to the get attributes of user params
func (o *GetAttributesOfUserParams) WithEmbedded(Embedded *string) *GetAttributesOfUserParams {
	o.Embedded = Embedded
	return o
}

// WithName adds the name to the get attributes of user params
func (o *GetAttributesOfUserParams) WithName(Name *string) *GetAttributesOfUserParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetAttributesOfUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Accept != nil {

		// query param Accept
		var qrAccept string
		if o.Accept != nil {
			qrAccept = *o.Accept
		}
		qAccept := qrAccept
		if qAccept != "" {
			if err := r.SetQueryParam("Accept", qAccept); err != nil {
				return err
			}
		}

	}

	if o.Embedded != nil {

		// query param _embedded
		var qrEmbedded string
		if o.Embedded != nil {
			qrEmbedded = *o.Embedded
		}
		qEmbedded := qrEmbedded
		if qEmbedded != "" {
			if err := r.SetQueryParam("_embedded", qEmbedded); err != nil {
				return err
			}
		}

	}

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
