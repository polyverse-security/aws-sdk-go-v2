// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appstream

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/BatchDisassociateUserStackRequest
type BatchDisassociateUserStackInput struct {
	_ struct{} `type:"structure"`

	// The list of UserStackAssociation objects.
	//
	// UserStackAssociations is a required field
	UserStackAssociations []UserStackAssociation `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s BatchDisassociateUserStackInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchDisassociateUserStackInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchDisassociateUserStackInput"}

	if s.UserStackAssociations == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserStackAssociations"))
	}
	if s.UserStackAssociations != nil && len(s.UserStackAssociations) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserStackAssociations", 1))
	}
	if s.UserStackAssociations != nil {
		for i, v := range s.UserStackAssociations {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UserStackAssociations", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/BatchDisassociateUserStackResult
type BatchDisassociateUserStackOutput struct {
	_ struct{} `type:"structure"`

	// The list of UserStackAssociationError objects.
	Errors []UserStackAssociationError `locationName:"errors" type:"list"`
}

// String returns the string representation
func (s BatchDisassociateUserStackOutput) String() string {
	return awsutil.Prettify(s)
}

const opBatchDisassociateUserStack = "BatchDisassociateUserStack"

// BatchDisassociateUserStackRequest returns a request value for making API operation for
// Amazon AppStream.
//
// Disassociates the specified users from the specified stacks.
//
//    // Example sending a request using BatchDisassociateUserStackRequest.
//    req := client.BatchDisassociateUserStackRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appstream-2016-12-01/BatchDisassociateUserStack
func (c *Client) BatchDisassociateUserStackRequest(input *BatchDisassociateUserStackInput) BatchDisassociateUserStackRequest {
	op := &aws.Operation{
		Name:       opBatchDisassociateUserStack,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &BatchDisassociateUserStackInput{}
	}

	req := c.newRequest(op, input, &BatchDisassociateUserStackOutput{})
	return BatchDisassociateUserStackRequest{Request: req, Input: input, Copy: c.BatchDisassociateUserStackRequest}
}

// BatchDisassociateUserStackRequest is the request type for the
// BatchDisassociateUserStack API operation.
type BatchDisassociateUserStackRequest struct {
	*aws.Request
	Input *BatchDisassociateUserStackInput
	Copy  func(*BatchDisassociateUserStackInput) BatchDisassociateUserStackRequest
}

// Send marshals and sends the BatchDisassociateUserStack API request.
func (r BatchDisassociateUserStackRequest) Send(ctx context.Context) (*BatchDisassociateUserStackResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &BatchDisassociateUserStackResponse{
		BatchDisassociateUserStackOutput: r.Request.Data.(*BatchDisassociateUserStackOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// BatchDisassociateUserStackResponse is the response type for the
// BatchDisassociateUserStack API operation.
type BatchDisassociateUserStackResponse struct {
	*BatchDisassociateUserStackOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// BatchDisassociateUserStack request.
func (r *BatchDisassociateUserStackResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
