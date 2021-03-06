// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductInput
type DescribeProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProductInput"}

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

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProductOutput
type DescribeProductOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associated budgets.
	Budgets []BudgetDetail `type:"list"`

	// Summary information about the product view.
	ProductViewSummary *ProductViewSummary `type:"structure"`

	// Information about the provisioning artifacts for the specified product.
	ProvisioningArtifacts []ProvisioningArtifact `type:"list"`
}

// String returns the string representation
func (s DescribeProductOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProduct = "DescribeProduct"

// DescribeProductRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Gets information about the specified product.
//
//    // Example sending a request using DescribeProductRequest.
//    req := client.DescribeProductRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/DescribeProduct
func (c *Client) DescribeProductRequest(input *DescribeProductInput) DescribeProductRequest {
	op := &aws.Operation{
		Name:       opDescribeProduct,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeProductInput{}
	}

	req := c.newRequest(op, input, &DescribeProductOutput{})
	return DescribeProductRequest{Request: req, Input: input, Copy: c.DescribeProductRequest}
}

// DescribeProductRequest is the request type for the
// DescribeProduct API operation.
type DescribeProductRequest struct {
	*aws.Request
	Input *DescribeProductInput
	Copy  func(*DescribeProductInput) DescribeProductRequest
}

// Send marshals and sends the DescribeProduct API request.
func (r DescribeProductRequest) Send(ctx context.Context) (*DescribeProductResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProductResponse{
		DescribeProductOutput: r.Request.Data.(*DescribeProductOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProductResponse is the response type for the
// DescribeProduct API operation.
type DescribeProductResponse struct {
	*DescribeProductOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProduct request.
func (r *DescribeProductResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
