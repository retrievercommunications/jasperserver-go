package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetSpecificAttributeOfUserParams creates a new GetSpecificAttributeOfUserParams object
// with the default values initialized.
func NewGetSpecificAttributeOfUserParams() *GetSpecificAttributeOfUserParams {
	var ()
	return &GetSpecificAttributeOfUserParams{}
}

/*GetSpecificAttributeOfUserParams contains all the parameters to send to the API endpoint
for the get specific attribute of user operation typically these are written to a http.Request
*/
type GetSpecificAttributeOfUserParams struct {

	/*Accept*/
	Accept *string
	/*Embedded*/
	Embedded *string
	/*AttrName*/
	AttrName *string
	/*Name*/
	Name *string
}

// WithAccept adds the accept to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithAccept(accept *string) *GetSpecificAttributeOfUserParams {
	o.Accept = accept
	return o
}

// WithEmbedded adds the embedded to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithEmbedded(embedded *string) *GetSpecificAttributeOfUserParams {
	o.Embedded = embedded
	return o
}

// WithAttrName adds the attrName to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithAttrName(attrName *string) *GetSpecificAttributeOfUserParams {
	o.AttrName = attrName
	return o
}

// WithName adds the name to the get specific attribute of user params
func (o *GetSpecificAttributeOfUserParams) WithName(name *string) *GetSpecificAttributeOfUserParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpecificAttributeOfUserParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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

	if o.AttrName != nil {

		// path param attrName
		if err := r.SetPathParam("attrName", *o.AttrName); err != nil {
			return err
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