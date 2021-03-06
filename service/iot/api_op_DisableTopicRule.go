// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// The input for the DisableTopicRuleRequest operation.
type DisableTopicRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the rule to disable.
	//
	// RuleName is a required field
	RuleName *string `location:"uri" locationName:"ruleName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DisableTopicRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisableTopicRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisableTopicRuleInput"}

	if s.RuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RuleName"))
	}
	if s.RuleName != nil && len(*s.RuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RuleName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisableTopicRuleInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RuleName != nil {
		v := *s.RuleName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ruleName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DisableTopicRuleOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DisableTopicRuleOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DisableTopicRuleOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDisableTopicRule = "DisableTopicRule"

// DisableTopicRuleRequest returns a request value for making API operation for
// AWS IoT.
//
// Disables the rule.
//
//    // Example sending a request using DisableTopicRuleRequest.
//    req := client.DisableTopicRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DisableTopicRuleRequest(input *DisableTopicRuleInput) DisableTopicRuleRequest {
	op := &aws.Operation{
		Name:       opDisableTopicRule,
		HTTPMethod: "POST",
		HTTPPath:   "/rules/{ruleName}/disable",
	}

	if input == nil {
		input = &DisableTopicRuleInput{}
	}

	req := c.newRequest(op, input, &DisableTopicRuleOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DisableTopicRuleRequest{Request: req, Input: input, Copy: c.DisableTopicRuleRequest}
}

// DisableTopicRuleRequest is the request type for the
// DisableTopicRule API operation.
type DisableTopicRuleRequest struct {
	*aws.Request
	Input *DisableTopicRuleInput
	Copy  func(*DisableTopicRuleInput) DisableTopicRuleRequest
}

// Send marshals and sends the DisableTopicRule API request.
func (r DisableTopicRuleRequest) Send(ctx context.Context) (*DisableTopicRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisableTopicRuleResponse{
		DisableTopicRuleOutput: r.Request.Data.(*DisableTopicRuleOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisableTopicRuleResponse is the response type for the
// DisableTopicRule API operation.
type DisableTopicRuleResponse struct {
	*DisableTopicRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisableTopicRule request.
func (r *DisableTopicRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
