// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Request to delete the specified API from your collection.
type DeleteRestApiInput struct {
	_ struct{} `type:"structure"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteRestApiInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteRestApiInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteRestApiInput"}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRestApiInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.RestApiId != nil {
		v := *s.RestApiId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "restapi_id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteRestApiOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteRestApiOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteRestApiOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opDeleteRestApi = "DeleteRestApi"

// DeleteRestApiRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Deletes the specified API.
//
//    // Example sending a request using DeleteRestApiRequest.
//    req := client.DeleteRestApiRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DeleteRestApiRequest(input *DeleteRestApiInput) DeleteRestApiRequest {
	op := &aws.Operation{
		Name:       opDeleteRestApi,
		HTTPMethod: "DELETE",
		HTTPPath:   "/restapis/{restapi_id}",
	}

	if input == nil {
		input = &DeleteRestApiInput{}
	}

	req := c.newRequest(op, input, &DeleteRestApiOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRestApiRequest{Request: req, Input: input, Copy: c.DeleteRestApiRequest}
}

// DeleteRestApiRequest is the request type for the
// DeleteRestApi API operation.
type DeleteRestApiRequest struct {
	*aws.Request
	Input *DeleteRestApiInput
	Copy  func(*DeleteRestApiInput) DeleteRestApiRequest
}

// Send marshals and sends the DeleteRestApi API request.
func (r DeleteRestApiRequest) Send(ctx context.Context) (*DeleteRestApiResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRestApiResponse{
		DeleteRestApiOutput: r.Request.Data.(*DeleteRestApiOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRestApiResponse is the response type for the
// DeleteRestApi API operation.
type DeleteRestApiResponse struct {
	*DeleteRestApiOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRestApi request.
func (r *DeleteRestApiResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
