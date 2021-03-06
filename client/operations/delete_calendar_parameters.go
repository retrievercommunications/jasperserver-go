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

// NewDeleteCalendarParams creates a new DeleteCalendarParams object
// with the default values initialized.
func NewDeleteCalendarParams() *DeleteCalendarParams {
	var ()
	return &DeleteCalendarParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCalendarParamsWithTimeout creates a new DeleteCalendarParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCalendarParamsWithTimeout(timeout time.Duration) *DeleteCalendarParams {
	var ()
	return &DeleteCalendarParams{

		timeout: timeout,
	}
}

/*DeleteCalendarParams contains all the parameters to send to the API endpoint
for the delete calendar operation typically these are written to a http.Request
*/
type DeleteCalendarParams struct {

	/*CalendarName*/
	CalendarName *string

	timeout time.Duration
}

// WithCalendarName adds the calendarName to the delete calendar params
func (o *DeleteCalendarParams) WithCalendarName(CalendarName *string) *DeleteCalendarParams {
	o.CalendarName = CalendarName
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCalendarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.CalendarName != nil {

		// path param calendarName
		if err := r.SetPathParam("calendarName", *o.CalendarName); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
