package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteJobsParams creates a new DeleteJobsParams object
// with the default values initialized.
func NewDeleteJobsParams() *DeleteJobsParams {
	var ()
	return &DeleteJobsParams{}
}

/*DeleteJobsParams contains all the parameters to send to the API endpoint
for the delete jobs operation typically these are written to a http.Request
*/
type DeleteJobsParams struct {

	/*ID*/
	ID *string
}

// WithID adds the id to the delete jobs params
func (o *DeleteJobsParams) WithID(id *string) *DeleteJobsParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteJobsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}