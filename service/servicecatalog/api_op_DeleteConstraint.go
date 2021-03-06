// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteConstraintInput
type DeleteConstraintInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The identifier of the constraint.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteConstraintInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteConstraintInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteConstraintInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteConstraintOutput
type DeleteConstraintOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteConstraintOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteConstraint = "DeleteConstraint"

// DeleteConstraintRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes the specified constraint.
//
//    // Example sending a request using DeleteConstraintRequest.
//    req := client.DeleteConstraintRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteConstraint
func (c *Client) DeleteConstraintRequest(input *DeleteConstraintInput) DeleteConstraintRequest {
	op := &aws.Operation{
		Name:       opDeleteConstraint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteConstraintInput{}
	}

	req := c.newRequest(op, input, &DeleteConstraintOutput{})
	return DeleteConstraintRequest{Request: req, Input: input, Copy: c.DeleteConstraintRequest}
}

// DeleteConstraintRequest is the request type for the
// DeleteConstraint API operation.
type DeleteConstraintRequest struct {
	*aws.Request
	Input *DeleteConstraintInput
	Copy  func(*DeleteConstraintInput) DeleteConstraintRequest
}

// Send marshals and sends the DeleteConstraint API request.
func (r DeleteConstraintRequest) Send(ctx context.Context) (*DeleteConstraintResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteConstraintResponse{
		DeleteConstraintOutput: r.Request.Data.(*DeleteConstraintOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteConstraintResponse is the response type for the
// DeleteConstraint API operation.
type DeleteConstraintResponse struct {
	*DeleteConstraintOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteConstraint request.
func (r *DeleteConstraintResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
