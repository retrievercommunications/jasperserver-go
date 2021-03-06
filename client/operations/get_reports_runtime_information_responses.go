package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetReportsRuntimeInformationReader is a Reader for the GetReportsRuntimeInformation structure.
type GetReportsRuntimeInformationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetReportsRuntimeInformationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReportsRuntimeInformationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReportsRuntimeInformationOK creates a GetReportsRuntimeInformationOK with default headers values
func NewGetReportsRuntimeInformationOK() *GetReportsRuntimeInformationOK {
	return &GetReportsRuntimeInformationOK{}
}

/*GetReportsRuntimeInformationOK handles this case with default header values.

GetReportsRuntimeInformationOK get reports runtime information o k
*/
type GetReportsRuntimeInformationOK struct {
}

func (o *GetReportsRuntimeInformationOK) Error() string {
	return fmt.Sprintf("[GET /reportExecutions][%d] getReportsRuntimeInformationOK ", 200)
}

func (o *GetReportsRuntimeInformationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
