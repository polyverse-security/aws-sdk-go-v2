// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/AcceptReservedNodeExchangeInputMessage
type AcceptReservedNodeExchangeInput struct {
	_ struct{} `type:"structure"`

	// A string representing the node identifier of the DC1 Reserved Node to be
	// exchanged.
	//
	// ReservedNodeId is a required field
	ReservedNodeId *string `type:"string" required:"true"`

	// The unique identifier of the DC2 Reserved Node offering to be used for the
	// exchange. You can obtain the value for the parameter by calling GetReservedNodeExchangeOfferings
	//
	// TargetReservedNodeOfferingId is a required field
	TargetReservedNodeOfferingId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AcceptReservedNodeExchangeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AcceptReservedNodeExchangeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AcceptReservedNodeExchangeInput"}

	if s.ReservedNodeId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ReservedNodeId"))
	}

	if s.TargetReservedNodeOfferingId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TargetReservedNodeOfferingId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/AcceptReservedNodeExchangeOutputMessage
type AcceptReservedNodeExchangeOutput struct {
	_ struct{} `type:"structure"`

	// Describes a reserved node. You can call the DescribeReservedNodeOfferings
	// API to obtain the available reserved node offerings.
	ExchangedReservedNode *ReservedNode `type:"structure"`
}

// String returns the string representation
func (s AcceptReservedNodeExchangeOutput) String() string {
	return awsutil.Prettify(s)
}

const opAcceptReservedNodeExchange = "AcceptReservedNodeExchange"

// AcceptReservedNodeExchangeRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Exchanges a DC1 Reserved Node for a DC2 Reserved Node with no changes to
// the configuration (term, payment type, or number of nodes) and no additional
// costs.
//
//    // Example sending a request using AcceptReservedNodeExchangeRequest.
//    req := client.AcceptReservedNodeExchangeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/AcceptReservedNodeExchange
func (c *Client) AcceptReservedNodeExchangeRequest(input *AcceptReservedNodeExchangeInput) AcceptReservedNodeExchangeRequest {
	op := &aws.Operation{
		Name:       opAcceptReservedNodeExchange,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AcceptReservedNodeExchangeInput{}
	}

	req := c.newRequest(op, input, &AcceptReservedNodeExchangeOutput{})
	return AcceptReservedNodeExchangeRequest{Request: req, Input: input, Copy: c.AcceptReservedNodeExchangeRequest}
}

// AcceptReservedNodeExchangeRequest is the request type for the
// AcceptReservedNodeExchange API operation.
type AcceptReservedNodeExchangeRequest struct {
	*aws.Request
	Input *AcceptReservedNodeExchangeInput
	Copy  func(*AcceptReservedNodeExchangeInput) AcceptReservedNodeExchangeRequest
}

// Send marshals and sends the AcceptReservedNodeExchange API request.
func (r AcceptReservedNodeExchangeRequest) Send(ctx context.Context) (*AcceptReservedNodeExchangeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AcceptReservedNodeExchangeResponse{
		AcceptReservedNodeExchangeOutput: r.Request.Data.(*AcceptReservedNodeExchangeOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AcceptReservedNodeExchangeResponse is the response type for the
// AcceptReservedNodeExchange API operation.
type AcceptReservedNodeExchangeResponse struct {
	*AcceptReservedNodeExchangeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AcceptReservedNodeExchange request.
func (r *AcceptReservedNodeExchangeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
