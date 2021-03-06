// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudwatchevents

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/DescribeEventBusRequest
type DescribeEventBusInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeEventBusInput) String() string {
	return awsutil.Prettify(s)
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/DescribeEventBusResponse
type DescribeEventBusOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the account permitted to write events to
	// the current account.
	Arn *string `type:"string"`

	// The name of the event bus. Currently, this is always default.
	Name *string `type:"string"`

	// The policy that enables the external account to send events to your account.
	Policy *string `type:"string"`
}

// String returns the string representation
func (s DescribeEventBusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeEventBus = "DescribeEventBus"

// DescribeEventBusRequest returns a request value for making API operation for
// Amazon CloudWatch Events.
//
// Displays the external AWS accounts that are permitted to write events to
// your account using your account's event bus, and the associated policy. To
// enable your account to receive events from other accounts, use PutPermission.
//
//    // Example sending a request using DescribeEventBusRequest.
//    req := client.DescribeEventBusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/events-2015-10-07/DescribeEventBus
func (c *Client) DescribeEventBusRequest(input *DescribeEventBusInput) DescribeEventBusRequest {
	op := &aws.Operation{
		Name:       opDescribeEventBus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEventBusInput{}
	}

	req := c.newRequest(op, input, &DescribeEventBusOutput{})
	return DescribeEventBusRequest{Request: req, Input: input, Copy: c.DescribeEventBusRequest}
}

// DescribeEventBusRequest is the request type for the
// DescribeEventBus API operation.
type DescribeEventBusRequest struct {
	*aws.Request
	Input *DescribeEventBusInput
	Copy  func(*DescribeEventBusInput) DescribeEventBusRequest
}

// Send marshals and sends the DescribeEventBus API request.
func (r DescribeEventBusRequest) Send(ctx context.Context) (*DescribeEventBusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeEventBusResponse{
		DescribeEventBusOutput: r.Request.Data.(*DescribeEventBusOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeEventBusResponse is the response type for the
// DescribeEventBus API operation.
type DescribeEventBusResponse struct {
	*DescribeEventBusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeEventBus request.
func (r *DescribeEventBusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
