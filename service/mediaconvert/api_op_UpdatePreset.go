// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Modify a preset by sending a request with the preset name and any of the
// following that you wish to change: description, category, and transcoding
// settings.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/UpdatePresetRequest
type UpdatePresetInput struct {
	_ struct{} `type:"structure"`

	// The new category for the preset, if you are changing it.
	Category *string `locationName:"category" type:"string"`

	// The new description for the preset, if you are changing it.
	Description *string `locationName:"description" type:"string"`

	// The name of the preset you are modifying.
	//
	// Name is a required field
	Name *string `location:"uri" locationName:"name" type:"string" required:"true"`

	// Settings for preset
	Settings *PresetSettings `locationName:"settings" type:"structure"`
}

// String returns the string representation
func (s UpdatePresetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdatePresetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdatePresetInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			invalidParams.AddNested("Settings", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePresetInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Category != nil {
		v := *s.Category

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "category", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Settings != nil {
		v := s.Settings

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "settings", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Successful update preset requests will return the new preset JSON.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/UpdatePresetResponse
type UpdatePresetOutput struct {
	_ struct{} `type:"structure"`

	// A preset is a collection of preconfigured media conversion settings that
	// you want MediaConvert to apply to the output during the conversion process.
	Preset *Preset `locationName:"preset" type:"structure"`
}

// String returns the string representation
func (s UpdatePresetOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdatePresetOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Preset != nil {
		v := s.Preset

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "preset", v, metadata)
	}
	return nil
}

const opUpdatePreset = "UpdatePreset"

// UpdatePresetRequest returns a request value for making API operation for
// AWS Elemental MediaConvert.
//
// Modify one of your existing presets.
//
//    // Example sending a request using UpdatePresetRequest.
//    req := client.UpdatePresetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediaconvert-2017-08-29/UpdatePreset
func (c *Client) UpdatePresetRequest(input *UpdatePresetInput) UpdatePresetRequest {
	op := &aws.Operation{
		Name:       opUpdatePreset,
		HTTPMethod: "PUT",
		HTTPPath:   "/2017-08-29/presets/{name}",
	}

	if input == nil {
		input = &UpdatePresetInput{}
	}

	req := c.newRequest(op, input, &UpdatePresetOutput{})
	return UpdatePresetRequest{Request: req, Input: input, Copy: c.UpdatePresetRequest}
}

// UpdatePresetRequest is the request type for the
// UpdatePreset API operation.
type UpdatePresetRequest struct {
	*aws.Request
	Input *UpdatePresetInput
	Copy  func(*UpdatePresetInput) UpdatePresetRequest
}

// Send marshals and sends the UpdatePreset API request.
func (r UpdatePresetRequest) Send(ctx context.Context) (*UpdatePresetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdatePresetResponse{
		UpdatePresetOutput: r.Request.Data.(*UpdatePresetOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdatePresetResponse is the response type for the
// UpdatePreset API operation.
type UpdatePresetResponse struct {
	*UpdatePresetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdatePreset request.
func (r *UpdatePresetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
