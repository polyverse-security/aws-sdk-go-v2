// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateServerRequest
type CreateServerInput struct {
	_ struct{} `type:"structure"`

	// The virtual private cloud (VPC) endpoint settings that you want to configure
	// for your SFTP server.
	EndpointDetails *EndpointDetails `type:"structure"`

	// The type of VPC endpoint that you want your SFTP server connect to. If you
	// connect to a VPC endpoint, your SFTP server isn't accessible over the public
	// internet.
	EndpointType EndpointType `type:"string" enum:"true"`

	// The RSA private key as generated by ssh-keygen -N "" -f my-new-server-key
	// command.
	//
	// If you aren't planning to migrate existing users from an existing SFTP server
	// to a new AWS SFTP server, don't update the host key. Accidentally changing
	// a server's host key can be disruptive. For more information, see change-host-key
	// in the AWS SFTP User Guide.
	HostKey *string `type:"string"`

	// An array containing all of the information required to call a customer-supplied
	// authentication API. This parameter is not required when the IdentityProviderType
	// value of server that is created uses the SERVICE_MANAGED authentication method.
	IdentityProviderDetails *IdentityProviderDetails `type:"structure"`

	// The mode of authentication enabled for this service. The default value is
	// SERVICE_MANAGED, which allows you to store and access SFTP user credentials
	// within the service. An IdentityProviderType value of API_GATEWAY indicates
	// that user authentication requires a call to an API Gateway endpoint URL provided
	// by you to integrate an identity provider of your choice.
	IdentityProviderType IdentityProviderType `type:"string" enum:"true"`

	// A value that allows the service to write your SFTP users' activity to your
	// Amazon CloudWatch logs for monitoring and auditing purposes.
	LoggingRole *string `type:"string"`

	// Key-value pairs that can be used to group and search for servers.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateServerInput"}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateServerResponse
type CreateServerOutput struct {
	_ struct{} `type:"structure"`

	// The service-assigned ID of the SFTP server that is created.
	//
	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateServer = "CreateServer"

// CreateServerRequest returns a request value for making API operation for
// AWS Transfer for SFTP.
//
// Instantiates an autoscaling virtual server based on Secure File Transfer
// Protocol (SFTP) in AWS. The call returns the ServerId property assigned by
// the service to the newly created server. Reference this ServerId property
// when you make updates to your server, or work with users.
//
// The response returns the ServerId value for the newly created server.
//
//    // Example sending a request using CreateServerRequest.
//    req := client.CreateServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transfer-2018-11-05/CreateServer
func (c *Client) CreateServerRequest(input *CreateServerInput) CreateServerRequest {
	op := &aws.Operation{
		Name:       opCreateServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateServerInput{}
	}

	req := c.newRequest(op, input, &CreateServerOutput{})
	return CreateServerRequest{Request: req, Input: input, Copy: c.CreateServerRequest}
}

// CreateServerRequest is the request type for the
// CreateServer API operation.
type CreateServerRequest struct {
	*aws.Request
	Input *CreateServerInput
	Copy  func(*CreateServerInput) CreateServerRequest
}

// Send marshals and sends the CreateServer API request.
func (r CreateServerRequest) Send(ctx context.Context) (*CreateServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateServerResponse{
		CreateServerOutput: r.Request.Data.(*CreateServerOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateServerResponse is the response type for the
// CreateServer API operation.
type CreateServerResponse struct {
	*CreateServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateServer request.
func (r *CreateServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
