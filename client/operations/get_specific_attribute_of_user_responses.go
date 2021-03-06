package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetSpecificAttributeOfUserReader is a Reader for the GetSpecificAttributeOfUser structure.
type GetSpecificAttributeOfUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSpecificAttributeOfUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSpecificAttributeOfUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSpecificAttributeOfUserOK creates a GetSpecificAttributeOfUserOK with default headers values
func NewGetSpecificAttributeOfUserOK() *GetSpecificAttributeOfUserOK {
	return &GetSpecificAttributeOfUserOK{}
}

/*GetSpecificAttributeOfUserOK handles this case with default header values.

GetSpecificAttributeOfUserOK get specific attribute of user o k
*/
type GetSpecificAttributeOfUserOK struct {
}

func (o *GetSpecificAttributeOfUserOK) Error() string {
	return fmt.Sprintf("[GET /{name}/attributes/{attrName}][%d] getSpecificAttributeOfUserOK ", 200)
}

func (o *GetSpecificAttributeOfUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
