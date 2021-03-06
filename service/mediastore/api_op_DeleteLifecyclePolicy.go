// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteLifecyclePolicyInput
type DeleteLifecyclePolicyInput struct {
	_ struct{} `type:"structure"`

	// The name of the container that holds the object lifecycle policy.
	//
	// ContainerName is a required field
	ContainerName *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLifecyclePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLifecyclePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLifecyclePolicyInput"}

	if s.ContainerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContainerName"))
	}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteLifecyclePolicyOutput
type DeleteLifecyclePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteLifecyclePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteLifecyclePolicy = "DeleteLifecyclePolicy"

// DeleteLifecyclePolicyRequest returns a request value for making API operation for
// AWS Elemental MediaStore.
//
// Removes an object lifecycle policy from a container. It takes up to 20 minutes
// for the change to take effect.
//
//    // Example sending a request using DeleteLifecyclePolicyRequest.
//    req := client.DeleteLifecyclePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediastore-2017-09-01/DeleteLifecyclePolicy
func (c *Client) DeleteLifecyclePolicyRequest(input *DeleteLifecyclePolicyInput) DeleteLifecyclePolicyRequest {
	op := &aws.Operation{
		Name:       opDeleteLifecyclePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLifecyclePolicyInput{}
	}

	req := c.newRequest(op, input, &DeleteLifecyclePolicyOutput{})
	return DeleteLifecyclePolicyRequest{Request: req, Input: input, Copy: c.DeleteLifecyclePolicyRequest}
}

// DeleteLifecyclePolicyRequest is the request type for the
// DeleteLifecyclePolicy API operation.
type DeleteLifecyclePolicyRequest struct {
	*aws.Request
	Input *DeleteLifecyclePolicyInput
	Copy  func(*DeleteLifecyclePolicyInput) DeleteLifecyclePolicyRequest
}

// Send marshals and sends the DeleteLifecyclePolicy API request.
func (r DeleteLifecyclePolicyRequest) Send(ctx context.Context) (*DeleteLifecyclePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteLifecyclePolicyResponse{
		DeleteLifecyclePolicyOutput: r.Request.Data.(*DeleteLifecyclePolicyOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteLifecyclePolicyResponse is the response type for the
// DeleteLifecyclePolicy API operation.
type DeleteLifecyclePolicyResponse struct {
	*DeleteLifecyclePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteLifecyclePolicy request.
func (r *DeleteLifecyclePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
