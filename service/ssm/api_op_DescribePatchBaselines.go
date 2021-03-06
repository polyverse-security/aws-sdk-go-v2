// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchBaselinesRequest
type DescribePatchBaselinesInput struct {
	_ struct{} `type:"structure"`

	// Each element in the array is a structure containing:
	//
	// Key: (string, "NAME_PREFIX" or "OWNER")
	//
	// Value: (array of strings, exactly 1 entry, between 1 and 255 characters)
	Filters []PatchOrchestratorFilter `type:"list"`

	// The maximum number of patch baselines to return (per page).
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribePatchBaselinesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePatchBaselinesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePatchBaselinesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchBaselinesResult
type DescribePatchBaselinesOutput struct {
	_ struct{} `type:"structure"`

	// An array of PatchBaselineIdentity elements.
	BaselineIdentities []PatchBaselineIdentity `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribePatchBaselinesOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribePatchBaselines = "DescribePatchBaselines"

// DescribePatchBaselinesRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// Lists the patch baselines in your AWS account.
//
//    // Example sending a request using DescribePatchBaselinesRequest.
//    req := client.DescribePatchBaselinesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/DescribePatchBaselines
func (c *Client) DescribePatchBaselinesRequest(input *DescribePatchBaselinesInput) DescribePatchBaselinesRequest {
	op := &aws.Operation{
		Name:       opDescribePatchBaselines,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePatchBaselinesInput{}
	}

	req := c.newRequest(op, input, &DescribePatchBaselinesOutput{})
	return DescribePatchBaselinesRequest{Request: req, Input: input, Copy: c.DescribePatchBaselinesRequest}
}

// DescribePatchBaselinesRequest is the request type for the
// DescribePatchBaselines API operation.
type DescribePatchBaselinesRequest struct {
	*aws.Request
	Input *DescribePatchBaselinesInput
	Copy  func(*DescribePatchBaselinesInput) DescribePatchBaselinesRequest
}

// Send marshals and sends the DescribePatchBaselines API request.
func (r DescribePatchBaselinesRequest) Send(ctx context.Context) (*DescribePatchBaselinesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribePatchBaselinesResponse{
		DescribePatchBaselinesOutput: r.Request.Data.(*DescribePatchBaselinesOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribePatchBaselinesResponse is the response type for the
// DescribePatchBaselines API operation.
type DescribePatchBaselinesResponse struct {
	*DescribePatchBaselinesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribePatchBaselines request.
func (r *DescribePatchBaselinesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
