// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspacesConnectionStatusRequest
type DescribeWorkspacesConnectionStatusInput struct {
	_ struct{} `type:"structure"`

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string `min:"1" type:"string"`

	// The identifiers of the WorkSpaces. You can specify up to 25 WorkSpaces.
	WorkspaceIds []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesConnectionStatusInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeWorkspacesConnectionStatusInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeWorkspacesConnectionStatusInput"}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}
	if s.WorkspaceIds != nil && len(s.WorkspaceIds) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WorkspaceIds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspacesConnectionStatusResult
type DescribeWorkspacesConnectionStatusOutput struct {
	_ struct{} `type:"structure"`

	// The token to use to retrieve the next set of results, or null if no more
	// results are available.
	NextToken *string `min:"1" type:"string"`

	// Information about the connection status of the WorkSpace.
	WorkspacesConnectionStatus []WorkspaceConnectionStatus `type:"list"`
}

// String returns the string representation
func (s DescribeWorkspacesConnectionStatusOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeWorkspacesConnectionStatus = "DescribeWorkspacesConnectionStatus"

// DescribeWorkspacesConnectionStatusRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Describes the connection status of the specified WorkSpaces.
//
//    // Example sending a request using DescribeWorkspacesConnectionStatusRequest.
//    req := client.DescribeWorkspacesConnectionStatusRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/DescribeWorkspacesConnectionStatus
func (c *Client) DescribeWorkspacesConnectionStatusRequest(input *DescribeWorkspacesConnectionStatusInput) DescribeWorkspacesConnectionStatusRequest {
	op := &aws.Operation{
		Name:       opDescribeWorkspacesConnectionStatus,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeWorkspacesConnectionStatusInput{}
	}

	req := c.newRequest(op, input, &DescribeWorkspacesConnectionStatusOutput{})
	return DescribeWorkspacesConnectionStatusRequest{Request: req, Input: input, Copy: c.DescribeWorkspacesConnectionStatusRequest}
}

// DescribeWorkspacesConnectionStatusRequest is the request type for the
// DescribeWorkspacesConnectionStatus API operation.
type DescribeWorkspacesConnectionStatusRequest struct {
	*aws.Request
	Input *DescribeWorkspacesConnectionStatusInput
	Copy  func(*DescribeWorkspacesConnectionStatusInput) DescribeWorkspacesConnectionStatusRequest
}

// Send marshals and sends the DescribeWorkspacesConnectionStatus API request.
func (r DescribeWorkspacesConnectionStatusRequest) Send(ctx context.Context) (*DescribeWorkspacesConnectionStatusResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeWorkspacesConnectionStatusResponse{
		DescribeWorkspacesConnectionStatusOutput: r.Request.Data.(*DescribeWorkspacesConnectionStatusOutput),
		response:                                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeWorkspacesConnectionStatusResponse is the response type for the
// DescribeWorkspacesConnectionStatus API operation.
type DescribeWorkspacesConnectionStatusResponse struct {
	*DescribeWorkspacesConnectionStatusOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeWorkspacesConnectionStatus request.
func (r *DescribeWorkspacesConnectionStatusResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
