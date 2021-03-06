// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package devicefarm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateVPCEConfigurationRequest
type CreateVPCEConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The DNS name of the service running in your VPC that you want Device Farm
	// to test.
	//
	// ServiceDnsName is a required field
	ServiceDnsName *string `locationName:"serviceDnsName" type:"string" required:"true"`

	// An optional description, providing more details about your VPC endpoint configuration.
	VpceConfigurationDescription *string `locationName:"vpceConfigurationDescription" type:"string"`

	// The friendly name you give to your VPC endpoint configuration, to manage
	// your configurations more easily.
	//
	// VpceConfigurationName is a required field
	VpceConfigurationName *string `locationName:"vpceConfigurationName" type:"string" required:"true"`

	// The name of the VPC endpoint service running inside your AWS account that
	// you want Device Farm to test.
	//
	// VpceServiceName is a required field
	VpceServiceName *string `locationName:"vpceServiceName" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVPCEConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVPCEConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateVPCEConfigurationInput"}

	if s.ServiceDnsName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceDnsName"))
	}

	if s.VpceConfigurationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpceConfigurationName"))
	}

	if s.VpceServiceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpceServiceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateVPCEConfigurationResult
type CreateVPCEConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// An object containing information about your VPC endpoint configuration.
	VpceConfiguration *VPCEConfiguration `locationName:"vpceConfiguration" type:"structure"`
}

// String returns the string representation
func (s CreateVPCEConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateVPCEConfiguration = "CreateVPCEConfiguration"

// CreateVPCEConfigurationRequest returns a request value for making API operation for
// AWS Device Farm.
//
// Creates a configuration record in Device Farm for your Amazon Virtual Private
// Cloud (VPC) endpoint.
//
//    // Example sending a request using CreateVPCEConfigurationRequest.
//    req := client.CreateVPCEConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/devicefarm-2015-06-23/CreateVPCEConfiguration
func (c *Client) CreateVPCEConfigurationRequest(input *CreateVPCEConfigurationInput) CreateVPCEConfigurationRequest {
	op := &aws.Operation{
		Name:       opCreateVPCEConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVPCEConfigurationInput{}
	}

	req := c.newRequest(op, input, &CreateVPCEConfigurationOutput{})
	return CreateVPCEConfigurationRequest{Request: req, Input: input, Copy: c.CreateVPCEConfigurationRequest}
}

// CreateVPCEConfigurationRequest is the request type for the
// CreateVPCEConfiguration API operation.
type CreateVPCEConfigurationRequest struct {
	*aws.Request
	Input *CreateVPCEConfigurationInput
	Copy  func(*CreateVPCEConfigurationInput) CreateVPCEConfigurationRequest
}

// Send marshals and sends the CreateVPCEConfiguration API request.
func (r CreateVPCEConfigurationRequest) Send(ctx context.Context) (*CreateVPCEConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateVPCEConfigurationResponse{
		CreateVPCEConfigurationOutput: r.Request.Data.(*CreateVPCEConfigurationOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateVPCEConfigurationResponse is the response type for the
// CreateVPCEConfiguration API operation.
type CreateVPCEConfigurationResponse struct {
	*CreateVPCEConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateVPCEConfiguration request.
func (r *CreateVPCEConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
