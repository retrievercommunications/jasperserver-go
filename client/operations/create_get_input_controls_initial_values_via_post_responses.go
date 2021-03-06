package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CreateGetInputControlsInitialValuesViaPostReader is a Reader for the CreateGetInputControlsInitialValuesViaPost structure.
type CreateGetInputControlsInitialValuesViaPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *CreateGetInputControlsInitialValuesViaPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateGetInputControlsInitialValuesViaPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGetInputControlsInitialValuesViaPostOK creates a CreateGetInputControlsInitialValuesViaPostOK with default headers values
func NewCreateGetInputControlsInitialValuesViaPostOK() *CreateGetInputControlsInitialValuesViaPostOK {
	return &CreateGetInputControlsInitialValuesViaPostOK{}
}

/*CreateGetInputControlsInitialValuesViaPostOK handles this case with default header values.

CreateGetInputControlsInitialValuesViaPostOK create get input controls initial values via post o k
*/
type CreateGetInputControlsInitialValuesViaPostOK struct {
}

func (o *CreateGetInputControlsInitialValuesViaPostOK) Error() string {
	return fmt.Sprintf("[POST /values][%d] createGetInputControlsInitialValuesViaPostOK ", 200)
}

func (o *CreateGetInputControlsInitialValuesViaPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
