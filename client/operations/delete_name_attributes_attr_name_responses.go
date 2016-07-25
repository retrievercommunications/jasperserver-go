package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteNameAttributesAttrNameReader is a Reader for the DeleteNameAttributesAttrName structure.
type DeleteNameAttributesAttrNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteNameAttributesAttrNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteNameAttributesAttrNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNameAttributesAttrNameOK creates a DeleteNameAttributesAttrNameOK with default headers values
func NewDeleteNameAttributesAttrNameOK() *DeleteNameAttributesAttrNameOK {
	return &DeleteNameAttributesAttrNameOK{}
}

/*DeleteNameAttributesAttrNameOK handles this case with default header values.

DeleteNameAttributesAttrNameOK delete name attributes attr name o k
*/
type DeleteNameAttributesAttrNameOK struct {
}

func (o *DeleteNameAttributesAttrNameOK) Error() string {
	return fmt.Sprintf("[DELETE /{name}/attributes/{attrName}][%d] deleteNameAttributesAttrNameOK ", 200)
}

func (o *DeleteNameAttributesAttrNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}