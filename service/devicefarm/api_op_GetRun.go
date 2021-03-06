// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents a request to the get run operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetRunRequest
type GetRunInput struct {
	_ struct{} `type:"structure"`

	// The run's ARN.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRunInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRunInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRunInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the result of a get run request.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetRunResult
type GetRunOutput struct {
	_ struct{} `type:"structure"`

	// The run you wish to get results from.
	Run *Run `locationName:"run" type:"structure"`
}

// String returns the string representation
func (s GetRunOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetRun = "GetRun"

// GetRunRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Gets information about a run.
//
//    // Example sending a request using GetRunRequest.
//    req := client.GetRunRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/GetRun
func (c *Client) GetRunRequest(input *GetRunInput) GetRunRequest {
	op := &aws.Operation{
		Name:       opGetRun,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetRunInput{}
	}

	req := c.newRequest(op, input, &GetRunOutput{})
	return GetRunRequest{Request: req, Input: input, Copy: c.GetRunRequest}
}

// GetRunRequest is the request type for the
// GetRun API operation.
type GetRunRequest struct {
	*aws.Request
	Input *GetRunInput
	Copy  func(*GetRunInput) GetRunRequest
}

// Send marshals and sends the GetRun API request.
func (r GetRunRequest) Send(ctx context.Context) (*GetRunResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetRunResponse{
		GetRunOutput: r.Request.Data.(*GetRunOutput),
		response:     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetRunResponse is the response type for the
// GetRun API operation.
type GetRunResponse struct {
	*GetRunOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetRun request.
func (r *GetRunResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
