package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutNameAttributesReader is a Reader for the PutNameAttributes structure.
type PutNameAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PutNameAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutNameAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutNameAttributesOK creates a PutNameAttributesOK with default headers values
func NewPutNameAttributesOK() *PutNameAttributesOK {
	return &PutNameAttributesOK{}
}

/*PutNameAttributesOK handles this case with default header values.

PutNameAttributesOK put name attributes o k
*/
type PutNameAttributesOK struct {
}

func (o *PutNameAttributesOK) Error() string {
	return fmt.Sprintf("[PUT /{name}/attributes][%d] putNameAttributesOK ", 200)
}

func (o *PutNameAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
