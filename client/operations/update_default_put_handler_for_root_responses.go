package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateDefaultPutHandlerForRootReader is a Reader for the UpdateDefaultPutHandlerForRoot structure.
type UpdateDefaultPutHandlerForRootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *UpdateDefaultPutHandlerForRootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDefaultPutHandlerForRootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateDefaultPutHandlerForRootOK creates a UpdateDefaultPutHandlerForRootOK with default headers values
func NewUpdateDefaultPutHandlerForRootOK() *UpdateDefaultPutHandlerForRootOK {
	return &UpdateDefaultPutHandlerForRootOK{}
}

/*UpdateDefaultPutHandlerForRootOK handles this case with default header values.

UpdateDefaultPutHandlerForRootOK update default put handler for root o k
*/
type UpdateDefaultPutHandlerForRootOK struct {
}

func (o *UpdateDefaultPutHandlerForRootOK) Error() string {
	return fmt.Sprintf("[PUT /resources][%d] updateDefaultPutHandlerForRootOK ", 200)
}

func (o *UpdateDefaultPutHandlerForRootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
