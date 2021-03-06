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

// NewPostReportsReportUnitURIInputControlsParams creates a new PostReportsReportUnitURIInputControlsParams object
// with the default values initialized.
func NewPostReportsReportUnitURIInputControlsParams() *PostReportsReportUnitURIInputControlsParams {
	var ()
	return &PostReportsReportUnitURIInputControlsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostReportsReportUnitURIInputControlsParamsWithTimeout creates a new PostReportsReportUnitURIInputControlsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostReportsReportUnitURIInputControlsParamsWithTimeout(timeout time.Duration) *PostReportsReportUnitURIInputControlsParams {
	var ()
	return &PostReportsReportUnitURIInputControlsParams{

		timeout: timeout,
	}
}

/*PostReportsReportUnitURIInputControlsParams contains all the parameters to send to the API endpoint
for the post reports report unit URI input controls operation typically these are written to a http.Request
*/
type PostReportsReportUnitURIInputControlsParams struct {

	/*ReportUnitURI*/
	ReportUnitURI *string

	timeout time.Duration
}

// WithReportUnitURI adds the reportUnitUri to the post reports report unit URI input controls params
func (o *PostReportsReportUnitURIInputControlsParams) WithReportUnitURI(ReportUnitURI *string) *PostReportsReportUnitURIInputControlsParams {
	o.ReportUnitURI = ReportUnitURI
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostReportsReportUnitURIInputControlsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.ReportUnitURI != nil {

		// query param reportUnitURI
		var qrReportUnitURI string
		if o.ReportUnitURI != nil {
			qrReportUnitURI = *o.ReportUnitURI
		}
		qReportUnitURI := qrReportUnitURI
		if qReportUnitURI != "" {
			if err := r.SetQueryParam("reportUnitURI", qReportUnitURI); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
