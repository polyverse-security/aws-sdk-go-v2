// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUICustomizationRequest
type GetUICustomizationInput struct {
	_ struct{} `type:"structure"`

	// The client ID for the client app.
	ClientId *string `min:"1" type:"string"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetUICustomizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUICustomizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetUICustomizationInput"}
	if s.ClientId != nil && len(*s.ClientId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientId", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUICustomizationResponse
type GetUICustomizationOutput struct {
	_ struct{} `type:"structure"`

	// The UI customization information.
	//
	// UICustomization is a required field
	UICustomization *UICustomizationType `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetUICustomizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetUICustomization = "GetUICustomization"

// GetUICustomizationRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Gets the UI Customization information for a particular app client's app UI,
// if there is something set. If nothing is set for the particular client, but
// there is an existing pool level customization (app clientId will be ALL),
// then that is returned. If nothing is present, then an empty shape is returned.
//
//    // Example sending a request using GetUICustomizationRequest.
//    req := client.GetUICustomizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/GetUICustomization
func (c *Client) GetUICustomizationRequest(input *GetUICustomizationInput) GetUICustomizationRequest {
	op := &aws.Operation{
		Name:       opGetUICustomization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUICustomizationInput{}
	}

	req := c.newRequest(op, input, &GetUICustomizationOutput{})
	return GetUICustomizationRequest{Request: req, Input: input, Copy: c.GetUICustomizationRequest}
}

// GetUICustomizationRequest is the request type for the
// GetUICustomization API operation.
type GetUICustomizationRequest struct {
	*aws.Request
	Input *GetUICustomizationInput
	Copy  func(*GetUICustomizationInput) GetUICustomizationRequest
}

// Send marshals and sends the GetUICustomization API request.
func (r GetUICustomizationRequest) Send(ctx context.Context) (*GetUICustomizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetUICustomizationResponse{
		GetUICustomizationOutput: r.Request.Data.(*GetUICustomizationOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetUICustomizationResponse is the response type for the
// GetUICustomization API operation.
type GetUICustomizationResponse struct {
	*GetUICustomizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetUICustomization request.
func (r *GetUICustomizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
