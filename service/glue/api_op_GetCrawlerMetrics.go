// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawlerMetricsRequest
type GetCrawlerMetricsInput struct {
	_ struct{} `type:"structure"`

	// A list of the names of crawlers about which to retrieve metrics.
	CrawlerNameList []string `type:"list"`

	// The maximum size of a list to return.
	MaxResults *int64 `min:"1" type:"integer"`

	// A continuation token, if this is a continuation call.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCrawlerMetricsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCrawlerMetricsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCrawlerMetricsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawlerMetricsResponse
type GetCrawlerMetricsOutput struct {
	_ struct{} `type:"structure"`

	// A list of metrics for the specified crawler.
	CrawlerMetricsList []CrawlerMetrics `type:"list"`

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s GetCrawlerMetricsOutput) String() string {
	return awsutil.Prettify(s)
}

const opGetCrawlerMetrics = "GetCrawlerMetrics"

// GetCrawlerMetricsRequest returns a request value for making API operation for
// AWS Glue.
//
// Retrieves metrics about specified crawlers.
//
//    // Example sending a request using GetCrawlerMetricsRequest.
//    req := client.GetCrawlerMetricsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/glue-2017-03-31/GetCrawlerMetrics
func (c *Client) GetCrawlerMetricsRequest(input *GetCrawlerMetricsInput) GetCrawlerMetricsRequest {
	op := &aws.Operation{
		Name:       opGetCrawlerMetrics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &GetCrawlerMetricsInput{}
	}

	req := c.newRequest(op, input, &GetCrawlerMetricsOutput{})
	return GetCrawlerMetricsRequest{Request: req, Input: input, Copy: c.GetCrawlerMetricsRequest}
}

// GetCrawlerMetricsRequest is the request type for the
// GetCrawlerMetrics API operation.
type GetCrawlerMetricsRequest struct {
	*aws.Request
	Input *GetCrawlerMetricsInput
	Copy  func(*GetCrawlerMetricsInput) GetCrawlerMetricsRequest
}

// Send marshals and sends the GetCrawlerMetrics API request.
func (r GetCrawlerMetricsRequest) Send(ctx context.Context) (*GetCrawlerMetricsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetCrawlerMetricsResponse{
		GetCrawlerMetricsOutput: r.Request.Data.(*GetCrawlerMetricsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewGetCrawlerMetricsRequestPaginator returns a paginator for GetCrawlerMetrics.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.GetCrawlerMetricsRequest(input)
//   p := glue.NewGetCrawlerMetricsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewGetCrawlerMetricsPaginator(req GetCrawlerMetricsRequest) GetCrawlerMetricsPaginator {
	return GetCrawlerMetricsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *GetCrawlerMetricsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// GetCrawlerMetricsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type GetCrawlerMetricsPaginator struct {
	aws.Pager
}

func (p *GetCrawlerMetricsPaginator) CurrentPage() *GetCrawlerMetricsOutput {
	return p.Pager.CurrentPage().(*GetCrawlerMetricsOutput)
}

// GetCrawlerMetricsResponse is the response type for the
// GetCrawlerMetrics API operation.
type GetCrawlerMetricsResponse struct {
	*GetCrawlerMetricsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetCrawlerMetrics request.
func (r *GetCrawlerMetricsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
