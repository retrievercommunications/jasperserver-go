package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetCalendarByNameReader is a Reader for the GetCalendarByName structure.
type GetCalendarByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCalendarByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCalendarByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCalendarByNameOK creates a GetCalendarByNameOK with default headers values
func NewGetCalendarByNameOK() *GetCalendarByNameOK {
	return &GetCalendarByNameOK{}
}

/*GetCalendarByNameOK handles this case with default header values.

GetCalendarByNameOK get calendar by name o k
*/
type GetCalendarByNameOK struct {
}

func (o *GetCalendarByNameOK) Error() string {
	return fmt.Sprintf("[GET /calendars/{calendarName}][%d] getCalendarByNameOK ", 200)
}

func (o *GetCalendarByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
