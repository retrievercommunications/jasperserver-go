package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetThumbnailsReader is a Reader for the GetThumbnails structure.
type GetThumbnailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetThumbnailsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetThumbnailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetThumbnailsOK creates a GetThumbnailsOK with default headers values
func NewGetThumbnailsOK() *GetThumbnailsOK {
	return &GetThumbnailsOK{}
}

/*GetThumbnailsOK handles this case with default header values.

GetThumbnailsOK get thumbnails o k
*/
type GetThumbnailsOK struct {
}

func (o *GetThumbnailsOK) Error() string {
	return fmt.Sprintf("[GET /thumbnails][%d] getThumbnailsOK ", 200)
}

func (o *GetThumbnailsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}