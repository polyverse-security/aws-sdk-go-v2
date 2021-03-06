// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the DeleteCertificate operation.
type DeleteCertificateInput struct {
	_ struct{} `type:"structure"`

	// The ID of the certificate. (The last part of the certificate ARN contains
	// the certificate ID.)
	//
	// CertificateId is a required field
	CertificateId *string `location:"uri" locationName:"certificateId" min:"64" type:"string" required:"true"`

	// Forces a certificate request to be deleted.
	ForceDelete *bool `location:"querystring" locationName:"forceDelete" type:"boolean"`
}

// String returns the string representation
func (s DeleteCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCertificateInput"}

	if s.CertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateId"))
	}
	if s.CertificateId != nil && len(*s.CertificateId) < 64 {
		invalidParams.Add(aws.NewErrParamMinLen("CertificateId", 64))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCertificateInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.CertificateId != nil {
		v := *s.CertificateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "certificateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ForceDelete != nil {
		v := *s.ForceDelete

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "forceDelete", protocol.BoolValue(v), metadata)
	}
	return nil
}

type DeleteCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCertificateOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteCertificateOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteCertificate = "DeleteCertificate"

// DeleteCertificateRequest returns a request value for making API operation for
// AWS IoT.
//
// Deletes the specified certificate.
//
// A certificate cannot be deleted if it has a policy attached to it or if its
// status is set to ACTIVE. To delete a certificate, first use the DetachPrincipalPolicy
// API to detach all policies. Next, use the UpdateCertificate API to set the
// certificate to the INACTIVE status.
//
//    // Example sending a request using DeleteCertificateRequest.
//    req := client.DeleteCertificateRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteCertificateRequest(input *DeleteCertificateInput) DeleteCertificateRequest {
	op := &aws.Operation{
		Name:       opDeleteCertificate,
		HTTPMethod: "DELETE",
		HTTPPath:   "/certificates/{certificateId}",
	}

	if input == nil {
		input = &DeleteCertificateInput{}
	}

	req := c.newRequest(op, input, &DeleteCertificateOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteCertificateRequest{Request: req, Input: input, Copy: c.DeleteCertificateRequest}
}

// DeleteCertificateRequest is the request type for the
// DeleteCertificate API operation.
type DeleteCertificateRequest struct {
	*aws.Request
	Input *DeleteCertificateInput
	Copy  func(*DeleteCertificateInput) DeleteCertificateRequest
}

// Send marshals and sends the DeleteCertificate API request.
func (r DeleteCertificateRequest) Send(ctx context.Context) (*DeleteCertificateResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteCertificateResponse{
		DeleteCertificateOutput: r.Request.Data.(*DeleteCertificateOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteCertificateResponse is the response type for the
// DeleteCertificate API operation.
type DeleteCertificateResponse struct {
	*DeleteCertificateOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteCertificate request.
func (r *DeleteCertificateResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
