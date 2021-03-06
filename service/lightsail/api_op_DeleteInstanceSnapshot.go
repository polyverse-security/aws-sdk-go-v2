// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteInstanceSnapshotRequest
type DeleteInstanceSnapshotInput struct {
	_ struct{} `type:"structure"`

	// The name of the snapshot to delete.
	//
	// InstanceSnapshotName is a required field
	InstanceSnapshotName *string `locationName:"instanceSnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteInstanceSnapshotInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteInstanceSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteInstanceSnapshotInput"}

	if s.InstanceSnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceSnapshotName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteInstanceSnapshotResult
type DeleteInstanceSnapshotOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the results of your
	// delete instance snapshot request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteInstanceSnapshotOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteInstanceSnapshot = "DeleteInstanceSnapshot"

// DeleteInstanceSnapshotRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Deletes a specific snapshot of a virtual private server (or instance).
//
// The delete instance snapshot operation supports tag-based access control
// via resource tags applied to the resource identified by instanceSnapshotName.
// For more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using DeleteInstanceSnapshotRequest.
//    req := client.DeleteInstanceSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/DeleteInstanceSnapshot
func (c *Client) DeleteInstanceSnapshotRequest(input *DeleteInstanceSnapshotInput) DeleteInstanceSnapshotRequest {
	op := &aws.Operation{
		Name:       opDeleteInstanceSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteInstanceSnapshotInput{}
	}

	req := c.newRequest(op, input, &DeleteInstanceSnapshotOutput{})
	return DeleteInstanceSnapshotRequest{Request: req, Input: input, Copy: c.DeleteInstanceSnapshotRequest}
}

// DeleteInstanceSnapshotRequest is the request type for the
// DeleteInstanceSnapshot API operation.
type DeleteInstanceSnapshotRequest struct {
	*aws.Request
	Input *DeleteInstanceSnapshotInput
	Copy  func(*DeleteInstanceSnapshotInput) DeleteInstanceSnapshotRequest
}

// Send marshals and sends the DeleteInstanceSnapshot API request.
func (r DeleteInstanceSnapshotRequest) Send(ctx context.Context) (*DeleteInstanceSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteInstanceSnapshotResponse{
		DeleteInstanceSnapshotOutput: r.Request.Data.(*DeleteInstanceSnapshotOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteInstanceSnapshotResponse is the response type for the
// DeleteInstanceSnapshot API operation.
type DeleteInstanceSnapshotResponse struct {
	*DeleteInstanceSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteInstanceSnapshot request.
func (r *DeleteInstanceSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
