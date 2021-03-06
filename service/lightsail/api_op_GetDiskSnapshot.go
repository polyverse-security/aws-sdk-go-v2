// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetDiskSnapshotRequest
type GetDiskSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the disk snapshot (e.g., my-disk-snapshot).
	//
	// DiskSnapshotName is a required field
	DiskSnapshotName *string `locationName:"diskSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDiskSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDiskSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDiskSnapshotInput"}

	if s.DiskSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DiskSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetDiskSnapshotResult
type GetDiskSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about the disk snapshot.
	DiskSnapshot *DiskSnapshot `locationName:"diskSnapshot" type:"structure"`
}

// String returns the string representation
func (s GetDiskSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetDiskSnapshot = "GetDiskSnapshot"

// GetDiskSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Returns information about a specific block storage disk snapshot.
//
//    // Example sending a request using GetDiskSnapshotRequest.
//    req := client.GetDiskSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/GetDiskSnapshot
func (c *Client) GetDiskSnapshotRequest(input *GetDiskSnapshotInput) GetDiskSnapshotRequest {
	op := &aws.Operation{
		Name:       opGetDiskSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetDiskSnapshotInput{}
	}

	req := c.newRequest(op, input, &GetDiskSnapshotOutput{})
	return GetDiskSnapshotRequest{Request: req, Input: input, Copy: c.GetDiskSnapshotRequest}
}

// GetDiskSnapshotRequest is the request type for the
// GetDiskSnapshot API operation.
type GetDiskSnapshotRequest struct {
	*aws.Request
	Input *GetDiskSnapshotInput
	Copy  func(*GetDiskSnapshotInput) GetDiskSnapshotRequest
}

// Send marshals and sends the GetDiskSnapshot API request.
func (r GetDiskSnapshotRequest) Send(ctx context.Context) (*GetDiskSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDiskSnapshotResponse{
		GetDiskSnapshotOutput: r.Request.Data.(*GetDiskSnapshotOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDiskSnapshotResponse is the response type for the
// GetDiskSnapshot API operation.
type GetDiskSnapshotResponse struct {
	*GetDiskSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDiskSnapshot request.
func (r *GetDiskSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
