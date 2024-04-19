// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnapmirrorRelationshipCreateReader is a Reader for the SnapmirrorRelationshipCreate structure.
type SnapmirrorRelationshipCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapmirrorRelationshipCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSnapmirrorRelationshipCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapmirrorRelationshipCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapmirrorRelationshipCreateAccepted creates a SnapmirrorRelationshipCreateAccepted with default headers values
func NewSnapmirrorRelationshipCreateAccepted() *SnapmirrorRelationshipCreateAccepted {
	return &SnapmirrorRelationshipCreateAccepted{}
}

/*
SnapmirrorRelationshipCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SnapmirrorRelationshipCreateAccepted struct {

	/* Useful for tracking the resource location
	 */
	Location string

	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this snapmirror relationship create accepted response has a 2xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapmirror relationship create accepted response has a 3xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapmirror relationship create accepted response has a 4xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapmirror relationship create accepted response has a 5xx status code
func (o *SnapmirrorRelationshipCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this snapmirror relationship create accepted response a status code equal to that given
func (o *SnapmirrorRelationshipCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SnapmirrorRelationshipCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateAccepted  %+v", 202, o.Payload)
}

func (o *SnapmirrorRelationshipCreateAccepted) String() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirrorRelationshipCreateAccepted  %+v", 202, o.Payload)
}

func (o *SnapmirrorRelationshipCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnapmirrorRelationshipCreateDefault creates a SnapmirrorRelationshipCreateDefault with default headers values
func NewSnapmirrorRelationshipCreateDefault(code int) *SnapmirrorRelationshipCreateDefault {
	return &SnapmirrorRelationshipCreateDefault{
		_statusCode: code,
	}
}

/*
	SnapmirrorRelationshipCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1115545 | Access token has an invalid signature. |
| 1115546 | Access token expired at a time. |
| 1115547 | Access token not valid until a time. |
| 1115548 | Access token is malformed. |
| 1115549 | Internal error. Failed to validate access token. |
| 6619637 | Relationship with specified destination volume already exists. |
| 6619699 | Schedule not found. |
| 6620374 | Internal error. Failed to get SVM information. |
| 6620478 | Internal error. Failed to check SnapMirror capability. |
| 6621834 | Object store configuration does not exist for specified vserver. |
| 13303819 | Could not retrieve SnapMirror policy information. |
| 13303821 | Invalid SnapMirror policy UUID. |
| 13303841 | This operation is not supported for SnapMirror relationships between these endpoints. |
| 13303852 | destination.path provided does not contain \\\":\\\". |
| 13303853 | Restore relationships are not supported for SVM-DR endpoints. |
| 13303868 | Create of destination endpoint and SnapMirror relationship failed. |
| 13303869 | Creating a destination endpoint is not supported for restore relationships. |
| 13303870 | A tiering policy cannot be specified if tiering is not being set to supported. |
| 13303871 | Storage service properties cannot be specified if the storage service is not being enabled. |
| 13303872 | Specified property requires a later effective cluster version. |
| 13303873 | Specifying a state when creating a relationship is only supported when creating a destination endpoint. |
| 13303874 | Specified state is not supported when creating this relationship. |
| 13303875 | Destination aggregates do not have sufficient space for hosting copies of source volumes. |
| 13303876 | Destination cluster does not have composite aggregates. |
| 13303877 | Source or destination cluster must be specified. |
| 13303878 | The specified fields do not match. |
| 13303879 | Source cluster name or UUID is needed to provision a destination SVM on the local cluster. |
| 13303880 | Source cluster must be remote for provisioning a destination SVM on the local cluster. |
| 13303881 | Network validation failed. |
| 13303882 | SVM validation failed. |
| 13303883 | Encryption is not enabled on the destination cluster. |
| 13303887 | Synchronous SnapMirror relationships between FlexGroup volumes are not supported. |
| 13303888 | Synchronous SnapMirror relationships require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303889 | Asynchronous SnapMirror relationships between FlexGroup volumes require an effective cluster version of 9.5 or later on both the source and destination clusters. |
| 13303890 | Asynchronous SnapMirror relationships between FlexVol volumes require an effective cluster version of 9.3, 9.5, or later on both the source and destination clusters. |
| 13303891 | Creating a destination endpoint with storage service requires an effective cluster version of 9.7 or later. |
| 13303892 | Fetching remote information from the destination cluster failed. |
| 13303893 | Updating job description failed. |
| 13303894 | Destination volume name is invalid. It must contain the source volume name and have a suffix when creating a destination endpoint on a cluster with an effective cluster version of 9.6 or earlier. |
| 13303895 | Operation on the remote destination cluster is not supported. |
| 13303897 | Specifying transfer_schedule is only supported for an asynchronous SnapMirror relationship with a remote destination cluster that has an effective cluster version of 9.6 or earlier. |
| 13303916 | FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303918 | No suitable destination aggregate type is available. |
| 13303919 | Only FabricPool enabled aggregates are available on the destination. |
| 13303920 | Only SnapLock aggregates are available on the destination. FlexGroup volumes are not supported on SnapLock aggregates. |
| 13303921 | Unable to retrieve the SnapMirror capabilities of the destination cluster. |
| 13303922 | Specified source SVM is not a data SVM. |
| 13303923 | Specified destination SVM is not a data SVM. |
| 13303924 | Source SVM has an invalid Snapshot copy policy. |
| 13303925 | SnapMirror validation has failed. |
| 13303930 | The specified tiering policy is not supported for destination volumes of Synchronous relationships. |
| 13303938 | Fetching information from the local cluster failed. |
| 13303939 | Could not create an SVM peer relationship. |
| 13303944 | An SVM-DR relationship is not supported because the source SVM has CIFS configured and the associated SnapMirror policy has either the "identity_preservation" property not set or set to "exclude_network_and_protocol_config". |
| 13303945 | Schedule specified cannot be associated with the relationship because the policy used for the relationship has a schedule. |
| 13303953 | SnapMirror relationships between FlexGroup volumes and object store endpoints are not supported. |
| 13303957 | Restore to the specified destination endpoint is not supported. |
| 13303958 | The \\\"source.uuid\\\" property must be specified. |
| 13303959 | The \\\"source.uuid\\\" property can only be specified when using the \\\"restore\\\" property to create a SnapMirror relationship from an object store endpoint. |
| 13303960 | The \\\"destination.uuid\\\" property cannot be specified for these endpoints. |
| 13303961 | Creating a destination endpoint is not supported with the \\\"source.uuid\\\" property. |
| 13303962 | Creating a destination endpoint is not supported with the \\\"destination.uuid\\\" property. |
| 13303964 | Specified property is not supported for object store relationships. |
| 13303965 | Specified endpoint not found. |
| 13303966 | Consistency Group relationships require a policy of type \"sync\" with a sync_type of \"automated_failover\". |
| 13303967 | Consistency Group volume is not a FlexVol volume. |
| 13303968 | Unsupported volume type for the Consistency Group. |
| 13303969 | SnapMirror relationships between SVM endpoints and object store endpoints are not supported. |
| 13303970 | Unsupported policy type for the Consistency Group. |
| 13303971 | SnapMirror relationships between Consistency Group endpoints and object store endpoints are not supported. |
| 13303976 | Source or destination SVM is already part of an SVM-DR relation. |
| 13303977 | Destination Consistency Group volume UUIDs are not expected while provisioning the destination volumes. |
| 13303978 | Number of Consistency Group volume names and UUIDs does not match. |
| 13303979 | Number of Consistency Group volumes exceeds the allowed limit. |
| 13303980 | Number of source and destination Consistency Group volumes do not match. |
| 13303981 | ISCSI or FCP protocol is not configured. |
| 13303982 | SAN data interface is not configured on the SVM. |
| 13304021 | No suitable storage can be found meeting the specified requirements. No FabricPool enabled aggregates are available on the destination. |
| 13304022 | No suitable storage can be found meeting the specified requirements. No non-root, non-taken-over, non-SnapLock, non-composite aggregates are available on the destination. |
| 13304026 | API license token is required for this operation. |
| 13304027 | Invalid API license token specified. |
| 13304032 | In an "All SAN Array", an SVM-DR relationship is not supported when the associated SnapMirror policy does not have the "identity_preservation" property set to "exclude_network_and_protocol_config". |
| 13304034 | SnapMirror Cloud license must be installed to use this feature. |
| 13304035 | SnapMirror Cloud license capacity limit has been exceeded. |
| 13304036 | SnapmMirror Cloud license term has expired. |
| 13304038 | SnapMirror Cloud License capacity limit would be exceeded with this SnapMirror operation. |
| 13304042 | The specified destination or source cluster for relationships with an object store destination must be local. |
| 13304043 | The property must be specified for a restore from an object store endpoint. |
| 13304044 | File restore from object store endpoints is not supported for relationship restore. |
| 13304045 | Incremental restore from an object store endpoint is not supported. |
| 13304077 | Specified property is only supported for creating a destination endpoint for FabricLink SnapMirror relationships.</private> |
| 13304080 | Specified UUID and name do not match. |
| 13304082 | Specified properties are mutually exclusive. |
| 13304083 | The specified property is not supported because all nodes in the cluster are not capable of supporting the property.<private> |
| 13304098 | This SnapMirror policy is not supported for SnapMirror relationhips with SnapLock volumes. |
| 13304099 | SnapLock Compliance Clock is not running on all nodes in the destination cluster. |
| 13304108 | Schedule not found in the Administrative SVM or the SVM for the relationship. |
*/
type SnapmirrorRelationshipCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapmirror relationship create default response
func (o *SnapmirrorRelationshipCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snapmirror relationship create default response has a 2xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapmirror relationship create default response has a 3xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapmirror relationship create default response has a 4xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapmirror relationship create default response has a 5xx status code
func (o *SnapmirrorRelationshipCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapmirror relationship create default response a status code equal to that given
func (o *SnapmirrorRelationshipCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnapmirrorRelationshipCreateDefault) Error() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirror_relationship_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnapmirrorRelationshipCreateDefault) String() string {
	return fmt.Sprintf("[POST /snapmirror/relationships][%d] snapmirror_relationship_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnapmirrorRelationshipCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapmirrorRelationshipCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}