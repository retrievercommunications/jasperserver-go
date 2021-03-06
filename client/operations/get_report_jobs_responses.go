package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetReportJobsReader is a Reader for the GetReportJobs structure.
type GetReportJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetReportJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportJobsOK creates a GetReportJobsOK with default headers values
func NewGetReportJobsOK() *GetReportJobsOK {
	return &GetReportJobsOK{}
}

/*GetReportJobsOK handles this case with default header values.

GetReportJobsOK get report jobs o k
*/
type GetReportJobsOK struct {
}

func (o *GetReportJobsOK) Error() string {
	return fmt.Sprintf("[GET /jobs][%d] getReportJobsOK ", 200)
}

func (o *GetReportJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
