// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteTagOptionInput
type DeleteTagOptionInput struct {
	_ struct{} `type:"structure"`

	// The TagOption identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTagOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTagOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTagOptionInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteTagOptionOutput
type DeleteTagOptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteTagOptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteTagOption = "DeleteTagOption"

// DeleteTagOptionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Deletes the specified TagOption.
//
// You cannot delete a TagOption if it is associated with a product or portfolio.
//
//    // Example sending a request using DeleteTagOptionRequest.
//    req := client.DeleteTagOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DeleteTagOption
func (c *Client) DeleteTagOptionRequest(input *DeleteTagOptionInput) DeleteTagOptionRequest {
	op := &aws.Operation{
		Name:       opDeleteTagOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteTagOptionInput{}
	}

	req := c.newRequest(op, input, &DeleteTagOptionOutput{})
	return DeleteTagOptionRequest{Request: req, Input: input, Copy: c.DeleteTagOptionRequest}
}

// DeleteTagOptionRequest is the request type for the
// DeleteTagOption API operation.
type DeleteTagOptionRequest struct {
	*aws.Request
	Input *DeleteTagOptionInput
	Copy  func(*DeleteTagOptionInput) DeleteTagOptionRequest
}

// Send marshals and sends the DeleteTagOption API request.
func (r DeleteTagOptionRequest) Send(ctx context.Context) (*DeleteTagOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTagOptionResponse{
		DeleteTagOptionOutput: r.Request.Data.(*DeleteTagOptionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTagOptionResponse is the response type for the
// DeleteTagOption API operation.
type DeleteTagOptionResponse struct {
	*DeleteTagOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTagOption request.
func (r *DeleteTagOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
