// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyContactRolesDeleteReader is a Reader for the TenancyContactRolesDelete structure.
type TenancyContactRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyContactRolesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesDeleteNoContent creates a TenancyContactRolesDeleteNoContent with default headers values
func NewTenancyContactRolesDeleteNoContent() *TenancyContactRolesDeleteNoContent {
	return &TenancyContactRolesDeleteNoContent{}
}

/* TenancyContactRolesDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactRolesDeleteNoContent tenancy contact roles delete no content
*/
type TenancyContactRolesDeleteNoContent struct {
}

func (o *TenancyContactRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/{id}/][%d] tenancyContactRolesDeleteNoContent ", 204)
}

func (o *TenancyContactRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactRolesDeleteDefault creates a TenancyContactRolesDeleteDefault with default headers values
func NewTenancyContactRolesDeleteDefault(code int) *TenancyContactRolesDeleteDefault {
	return &TenancyContactRolesDeleteDefault{
		_statusCode: code,
	}
}

/* TenancyContactRolesDeleteDefault describes a response with status code -1, with default header values.

TenancyContactRolesDeleteDefault tenancy contact roles delete default
*/
type TenancyContactRolesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact roles delete default response
func (o *TenancyContactRolesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TenancyContactRolesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_delete default  %+v", o._statusCode, o.Payload)
}
func (o *TenancyContactRolesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}