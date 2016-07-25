package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteRoleParams creates a new DeleteRoleParams object
// with the default values initialized.
func NewDeleteRoleParams() *DeleteRoleParams {
	var ()
	return &DeleteRoleParams{}
}

/*DeleteRoleParams contains all the parameters to send to the API endpoint
for the delete role operation typically these are written to a http.Request
*/
type DeleteRoleParams struct {

	/*Name*/
	Name *string
}

// WithName adds the name to the delete role params
func (o *DeleteRoleParams) WithName(name *string) *DeleteRoleParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRoleParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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