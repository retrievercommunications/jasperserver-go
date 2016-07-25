package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// CreateGetReportInputControlValuesViaPostReader is a Reader for the CreateGetReportInputControlValuesViaPost structure.
type CreateGetReportInputControlValuesViaPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateGetReportInputControlValuesViaPostReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGetReportInputControlValuesViaPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGetReportInputControlValuesViaPostOK creates a CreateGetReportInputControlValuesViaPostOK with default headers values
func NewCreateGetReportInputControlValuesViaPostOK() *CreateGetReportInputControlValuesViaPostOK {
	return &CreateGetReportInputControlValuesViaPostOK{}
}

/*CreateGetReportInputControlValuesViaPostOK handles this case with default header values.

CreateGetReportInputControlValuesViaPostOK create get report input control values via post o k
*/
type CreateGetReportInputControlValuesViaPostOK struct {
}

func (o *CreateGetReportInputControlValuesViaPostOK) Error() string {
	return fmt.Sprintf("[POST /{inputControlIds: [^;/]+(;[^;/]+)*}/values][%d] createGetReportInputControlValuesViaPostOK ", 200)
}

func (o *CreateGetReportInputControlValuesViaPostOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}