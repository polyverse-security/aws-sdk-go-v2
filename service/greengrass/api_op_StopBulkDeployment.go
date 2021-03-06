// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package greengrass

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StopBulkDeploymentRequest
type StopBulkDeploymentInput struct {
	_ struct{} `type:"structure"`

	// BulkDeploymentId is a required field
	BulkDeploymentId *string `location:"uri" locationName:"BulkDeploymentId" type:"string" required:"true"`
}

// String returns the string representation
func (s StopBulkDeploymentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopBulkDeploymentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopBulkDeploymentInput"}

	if s.BulkDeploymentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BulkDeploymentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopBulkDeploymentInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.BulkDeploymentId != nil {
		v := *s.BulkDeploymentId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "BulkDeploymentId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StopBulkDeploymentResponse
type StopBulkDeploymentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StopBulkDeploymentOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s StopBulkDeploymentOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opStopBulkDeployment = "StopBulkDeployment"

// StopBulkDeploymentRequest returns a request value for making API operation for
// AWS Greengrass.
//
// Stops the execution of a bulk deployment. This action returns a status of
// ''Stopping'' until the deployment is stopped. You cannot start a new bulk
// deployment while a previous deployment is in the ''Stopping'' state. This
// action doesn't rollback completed deployments or cancel pending deployments.
//
//    // Example sending a request using StopBulkDeploymentRequest.
//    req := client.StopBulkDeploymentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/greengrass-2017-06-07/StopBulkDeployment
func (c *Client) StopBulkDeploymentRequest(input *StopBulkDeploymentInput) StopBulkDeploymentRequest {
	op := &aws.Operation{
		Name:       opStopBulkDeployment,
		HTTPMethod: "PUT",
		HTTPPath:   "/greengrass/bulk/deployments/{BulkDeploymentId}/$stop",
	}

	if input == nil {
		input = &StopBulkDeploymentInput{}
	}

	req := c.newRequest(op, input, &StopBulkDeploymentOutput{})
	return StopBulkDeploymentRequest{Request: req, Input: input, Copy: c.StopBulkDeploymentRequest}
}

// StopBulkDeploymentRequest is the request type for the
// StopBulkDeployment API operation.
type StopBulkDeploymentRequest struct {
	*aws.Request
	Input *StopBulkDeploymentInput
	Copy  func(*StopBulkDeploymentInput) StopBulkDeploymentRequest
}

// Send marshals and sends the StopBulkDeployment API request.
func (r StopBulkDeploymentRequest) Send(ctx context.Context) (*StopBulkDeploymentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StopBulkDeploymentResponse{
		StopBulkDeploymentOutput: r.Request.Data.(*StopBulkDeploymentOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StopBulkDeploymentResponse is the response type for the
// StopBulkDeployment API operation.
type StopBulkDeploymentResponse struct {
	*StopBulkDeploymentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StopBulkDeployment request.
func (r *StopBulkDeploymentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
