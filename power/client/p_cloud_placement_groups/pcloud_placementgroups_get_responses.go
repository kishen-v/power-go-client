// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPlacementgroupsGetReader is a Reader for the PcloudPlacementgroupsGet structure.
type PcloudPlacementgroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudPlacementgroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudPlacementgroupsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudPlacementgroupsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudPlacementgroupsGetOK creates a PcloudPlacementgroupsGetOK with default headers values
func NewPcloudPlacementgroupsGetOK() *PcloudPlacementgroupsGetOK {
	return &PcloudPlacementgroupsGetOK{}
}

/*PcloudPlacementgroupsGetOK handles this case with default header values.

OK
*/
type PcloudPlacementgroupsGetOK struct {
	Payload *models.PlacementGroup
}

func (o *PcloudPlacementgroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudPlacementgroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlacementGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetBadRequest creates a PcloudPlacementgroupsGetBadRequest with default headers values
func NewPcloudPlacementgroupsGetBadRequest() *PcloudPlacementgroupsGetBadRequest {
	return &PcloudPlacementgroupsGetBadRequest{}
}

/*PcloudPlacementgroupsGetBadRequest handles this case with default header values.

Bad Request
*/
type PcloudPlacementgroupsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudPlacementgroupsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsGetInternalServerError creates a PcloudPlacementgroupsGetInternalServerError with default headers values
func NewPcloudPlacementgroupsGetInternalServerError() *PcloudPlacementgroupsGetInternalServerError {
	return &PcloudPlacementgroupsGetInternalServerError{}
}

/*PcloudPlacementgroupsGetInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudPlacementgroupsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudPlacementgroupsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
