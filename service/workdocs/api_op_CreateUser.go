// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateUserRequest
type CreateUserInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Do not set this field when using administrative
	// API actions, as in accessing the API using AWS credentials.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string"`

	// The email address of the user.
	EmailAddress *string `min:"1" type:"string"`

	// The given name of the user.
	//
	// GivenName is a required field
	GivenName *string `min:"1" type:"string" required:"true"`

	// The ID of the organization.
	OrganizationId *string `min:"1" type:"string"`

	// The password of the user.
	//
	// Password is a required field
	Password *string `min:"4" type:"string" required:"true"`

	// The amount of storage for the user.
	StorageRule *StorageRuleType `type:"structure"`

	// The surname of the user.
	//
	// Surname is a required field
	Surname *string `min:"1" type:"string" required:"true"`

	// The time zone ID of the user.
	TimeZoneId *string `min:"1" type:"string"`

	// The login name of the user.
	//
	// Username is a required field
	Username *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateUserInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.EmailAddress != nil && len(*s.EmailAddress) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("EmailAddress", 1))
	}

	if s.GivenName == nil {
		invalidParams.Add(aws.NewErrParamRequired("GivenName"))
	}
	if s.GivenName != nil && len(*s.GivenName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GivenName", 1))
	}
	if s.OrganizationId != nil && len(*s.OrganizationId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationId", 1))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 4 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 4))
	}

	if s.Surname == nil {
		invalidParams.Add(aws.NewErrParamRequired("Surname"))
	}
	if s.Surname != nil && len(*s.Surname) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Surname", 1))
	}
	if s.TimeZoneId != nil && len(*s.TimeZoneId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TimeZoneId", 1))
	}

	if s.Username == nil {
		invalidParams.Add(aws.NewErrParamRequired("Username"))
	}
	if s.Username != nil && len(*s.Username) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Username", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.EmailAddress != nil {
		v := *s.EmailAddress

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "EmailAddress", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.GivenName != nil {
		v := *s.GivenName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "GivenName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.OrganizationId != nil {
		v := *s.OrganizationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "OrganizationId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Password != nil {
		v := *s.Password

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Password", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.StorageRule != nil {
		v := s.StorageRule

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "StorageRule", v, metadata)
	}
	if s.Surname != nil {
		v := *s.Surname

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Surname", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TimeZoneId != nil {
		v := *s.TimeZoneId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TimeZoneId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Username != nil {
		v := *s.Username

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Username", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateUserResponse
type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	// The user information.
	User *User `type:"structure"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateUserOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.User != nil {
		v := s.User

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "User", v, metadata)
	}
	return nil
}

const opCreateUser = "CreateUser"

// CreateUserRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Creates a user in a Simple AD or Microsoft AD directory. The status of a
// newly created user is "ACTIVE". New users can access Amazon WorkDocs.
//
//    // Example sending a request using CreateUserRequest.
//    req := client.CreateUserRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/CreateUser
func (c *Client) CreateUserRequest(input *CreateUserInput) CreateUserRequest {
	op := &aws.Operation{
		Name:       opCreateUser,
		HTTPMethod: "POST",
		HTTPPath:   "/api/v1/users",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	req := c.newRequest(op, input, &CreateUserOutput{})
	return CreateUserRequest{Request: req, Input: input, Copy: c.CreateUserRequest}
}

// CreateUserRequest is the request type for the
// CreateUser API operation.
type CreateUserRequest struct {
	*aws.Request
	Input *CreateUserInput
	Copy  func(*CreateUserInput) CreateUserRequest
}

// Send marshals and sends the CreateUser API request.
func (r CreateUserRequest) Send(ctx context.Context) (*CreateUserResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResponse{
		CreateUserOutput: r.Request.Data.(*CreateUserOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateUserResponse is the response type for the
// CreateUser API operation.
type CreateUserResponse struct {
	*CreateUserOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateUser request.
func (r *CreateUserResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
