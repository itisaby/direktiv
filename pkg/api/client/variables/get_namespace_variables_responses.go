// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNamespaceVariablesReader is a Reader for the GetNamespaceVariables structure.
type GetNamespaceVariablesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNamespaceVariablesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNamespaceVariablesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNamespaceVariablesOK creates a GetNamespaceVariablesOK with default headers values
func NewGetNamespaceVariablesOK() *GetNamespaceVariablesOK {
	return &GetNamespaceVariablesOK{}
}

/* GetNamespaceVariablesOK describes a response with status code 200, with default header values.

successfully got namespace variables
*/
type GetNamespaceVariablesOK struct {
}

func (o *GetNamespaceVariablesOK) Error() string {
	return fmt.Sprintf("[GET /api/namespaces/{namespace}/vars][%d] getNamespaceVariablesOK ", 200)
}

func (o *GetNamespaceVariablesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}