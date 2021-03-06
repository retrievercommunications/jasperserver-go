package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteJobReader is a Reader for the DeleteJob structure.
type DeleteJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteJobOK creates a DeleteJobOK with default headers values
func NewDeleteJobOK() *DeleteJobOK {
	return &DeleteJobOK{}
}

/*DeleteJobOK handles this case with default header values.

DeleteJobOK delete job o k
*/
type DeleteJobOK struct {
}

func (o *DeleteJobOK) Error() string {
	return fmt.Sprintf("[DELETE /{id: \\d+}][%d] deleteJobOK ", 200)
}

func (o *DeleteJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
