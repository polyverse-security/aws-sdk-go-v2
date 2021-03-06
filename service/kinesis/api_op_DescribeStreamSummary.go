// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStreamSummaryInput
type DescribeStreamSummaryInput struct {
	_ struct{} `type:"structure"`

	// The name of the stream to describe.
	//
	// StreamName is a required field
	StreamName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStreamSummaryInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStreamSummaryInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStreamSummaryInput"}

	if s.StreamName == nil {
		invalidParams.Add(aws.NewErrParamRequired("StreamName"))
	}
	if s.StreamName != nil && len(*s.StreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StreamName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStreamSummaryOutput
type DescribeStreamSummaryOutput struct {
	_ struct{} `type:"structure"`

	// A StreamDescriptionSummary containing information about the stream.
	//
	// StreamDescriptionSummary is a required field
	StreamDescriptionSummary *StreamDescriptionSummary `type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeStreamSummaryOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeStreamSummary = "DescribeStreamSummary"

// DescribeStreamSummaryRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Provides a summarized description of the specified Kinesis data stream without
// the shard list.
//
// The information returned includes the stream name, Amazon Resource Name (ARN),
// status, record retention period, approximate creation time, monitoring, encryption
// details, and open shard count.
//
//    // Example sending a request using DescribeStreamSummaryRequest.
//    req := client.DescribeStreamSummaryRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/DescribeStreamSummary
func (c *Client) DescribeStreamSummaryRequest(input *DescribeStreamSummaryInput) DescribeStreamSummaryRequest {
	op := &aws.Operation{
		Name:       opDescribeStreamSummary,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeStreamSummaryInput{}
	}

	req := c.newRequest(op, input, &DescribeStreamSummaryOutput{})
	return DescribeStreamSummaryRequest{Request: req, Input: input, Copy: c.DescribeStreamSummaryRequest}
}

// DescribeStreamSummaryRequest is the request type for the
// DescribeStreamSummary API operation.
type DescribeStreamSummaryRequest struct {
	*aws.Request
	Input *DescribeStreamSummaryInput
	Copy  func(*DescribeStreamSummaryInput) DescribeStreamSummaryRequest
}

// Send marshals and sends the DescribeStreamSummary API request.
func (r DescribeStreamSummaryRequest) Send(ctx context.Context) (*DescribeStreamSummaryResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeStreamSummaryResponse{
		DescribeStreamSummaryOutput: r.Request.Data.(*DescribeStreamSummaryOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeStreamSummaryResponse is the response type for the
// DescribeStreamSummary API operation.
type DescribeStreamSummaryResponse struct {
	*DescribeStreamSummaryOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeStreamSummary request.
func (r *DescribeStreamSummaryResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
