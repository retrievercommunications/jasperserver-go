package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// GetCustomDataSourceDefinitionReader is a Reader for the GetCustomDataSourceDefinition structure.
type GetCustomDataSourceDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetCustomDataSourceDefinitionReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCustomDataSourceDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCustomDataSourceDefinitionOK creates a GetCustomDataSourceDefinitionOK with default headers values
func NewGetCustomDataSourceDefinitionOK() *GetCustomDataSourceDefinitionOK {
	return &GetCustomDataSourceDefinitionOK{}
}

/*GetCustomDataSourceDefinitionOK handles this case with default header values.

GetCustomDataSourceDefinitionOK get custom data source definition o k
*/
type GetCustomDataSourceDefinitionOK struct {
}

func (o *GetCustomDataSourceDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /{name}][%d] getCustomDataSourceDefinitionOK ", 200)
}

func (o *GetCustomDataSourceDefinitionOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}