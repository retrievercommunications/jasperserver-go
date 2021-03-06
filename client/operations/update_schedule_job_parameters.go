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

// NewUpdateScheduleJobParams creates a new UpdateScheduleJobParams object
// with the default values initialized.
func NewUpdateScheduleJobParams() *UpdateScheduleJobParams {

	return &UpdateScheduleJobParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateScheduleJobParamsWithTimeout creates a new UpdateScheduleJobParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateScheduleJobParamsWithTimeout(timeout time.Duration) *UpdateScheduleJobParams {

	return &UpdateScheduleJobParams{

		timeout: timeout,
	}
}

/*UpdateScheduleJobParams contains all the parameters to send to the API endpoint
for the update schedule job operation typically these are written to a http.Request
*/
type UpdateScheduleJobParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateScheduleJobParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
