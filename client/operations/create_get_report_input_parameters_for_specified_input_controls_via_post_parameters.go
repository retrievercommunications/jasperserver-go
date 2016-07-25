package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParams creates a new CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams object
// with the default values initialized.
func NewCreateGetReportInputParametersForSpecifiedInputControlsViaPostParams() *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	var ()
	return &CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams{}
}

/*CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams contains all the parameters to send to the API endpoint
for the create get report input parameters for specified input controls via post operation typically these are written to a http.Request
*/
type CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams struct {

	/*InputControlIds*/
	InputControlIds *string
	/*ReportUnitURI*/
	ReportUnitURI *string
}

// WithInputControlIds adds the inputControlIds to the create get report input parameters for specified input controls via post params
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WithInputControlIds(inputControlIds *string) *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	o.InputControlIds = inputControlIds
	return o
}

// WithReportUnitURI adds the reportUnitUri to the create get report input parameters for specified input controls via post params
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WithReportUnitURI(reportUnitUri *string) *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams {
	o.ReportUnitURI = reportUnitUri
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGetReportInputParametersForSpecifiedInputControlsViaPostParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.InputControlIds != nil {

		// query param inputControlIds
		var qrInputControlIds string
		if o.InputControlIds != nil {
			qrInputControlIds = *o.InputControlIds
		}
		qInputControlIds := qrInputControlIds
		if qInputControlIds != "" {
			if err := r.SetQueryParam("inputControlIds", qInputControlIds); err != nil {
				return err
			}
		}

	}

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