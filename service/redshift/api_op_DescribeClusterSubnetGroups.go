// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterSubnetGroupsMessage
type DescribeClusterSubnetGroupsInput struct {
	_ struct{} `type:"structure"`

	// The name of the cluster subnet group for which information is requested.
	ClusterSubnetGroupName *string `type:"string"`

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterSubnetGroups request
	// exceed the value specified in MaxRecords, AWS returns a value in the Marker
	// field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying
	// the request.
	Marker *string `type:"string"`

	// The maximum number of response records to return in each call. If the number
	// of remaining response records exceeds the specified MaxRecords value, a value
	// is returned in a marker field of the response. You can retrieve the next
	// set of records by retrying the command with the returned marker value.
	//
	// Default: 100
	//
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// A tag key or keys for which you want to return all matching cluster subnet
	// groups that are associated with the specified key or keys. For example, suppose
	// that you have subnet groups that are tagged with keys called owner and environment.
	// If you specify both of these tag keys in the request, Amazon Redshift returns
	// a response with the subnet groups that have either or both of these tag keys
	// associated with them.
	TagKeys []string `locationNameList:"TagKey" type:"list"`

	// A tag value or values for which you want to return all matching cluster subnet
	// groups that are associated with the specified tag value or values. For example,
	// suppose that you have subnet groups that are tagged with values called admin
	// and test. If you specify both of these tag values in the request, Amazon
	// Redshift returns a response with the subnet groups that have either or both
	// of these tag values associated with them.
	TagValues []string `locationNameList:"TagValue" type:"list"`
}

// String returns the string representation
func (s DescribeClusterSubnetGroupsInput) String() string {
	return awsutil.Prettify(s)
}

// Contains the output from the DescribeClusterSubnetGroups action.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/ClusterSubnetGroupMessage
type DescribeClusterSubnetGroupsOutput struct {
	_ struct{} `type:"structure"`

	// A list of ClusterSubnetGroup instances.
	ClusterSubnetGroups []ClusterSubnetGroup `locationNameList:"ClusterSubnetGroup" type:"list"`

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string `type:"string"`
}

// String returns the string representation
func (s DescribeClusterSubnetGroupsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeClusterSubnetGroups = "DescribeClusterSubnetGroups"

// DescribeClusterSubnetGroupsRequest returns a request value for making API operation for
// Amazon Redshift.
//
// Returns one or more cluster subnet group objects, which contain metadata
// about your cluster subnet groups. By default, this operation returns information
// about all cluster subnet groups that are defined in you AWS account.
//
// If you specify both tag keys and tag values in the same request, Amazon Redshift
// returns all subnet groups that match any combination of the specified keys
// and values. For example, if you have owner and environment for tag keys,
// and admin and test for tag values, all subnet groups that have any combination
// of those values are returned.
//
// If both tag keys and values are omitted from the request, subnet groups are
// returned regardless of whether they have tag keys or values associated with
// them.
//
//    // Example sending a request using DescribeClusterSubnetGroupsRequest.
//    req := client.DescribeClusterSubnetGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/redshift-2012-12-01/DescribeClusterSubnetGroups
func (c *Client) DescribeClusterSubnetGroupsRequest(input *DescribeClusterSubnetGroupsInput) DescribeClusterSubnetGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeClusterSubnetGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeClusterSubnetGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeClusterSubnetGroupsOutput{})
	return DescribeClusterSubnetGroupsRequest{Request: req, Input: input, Copy: c.DescribeClusterSubnetGroupsRequest}
}

// DescribeClusterSubnetGroupsRequest is the request type for the
// DescribeClusterSubnetGroups API operation.
type DescribeClusterSubnetGroupsRequest struct {
	*aws.Request
	Input *DescribeClusterSubnetGroupsInput
	Copy  func(*DescribeClusterSubnetGroupsInput) DescribeClusterSubnetGroupsRequest
}

// Send marshals and sends the DescribeClusterSubnetGroups API request.
func (r DescribeClusterSubnetGroupsRequest) Send(ctx context.Context) (*DescribeClusterSubnetGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeClusterSubnetGroupsResponse{
		DescribeClusterSubnetGroupsOutput: r.Request.Data.(*DescribeClusterSubnetGroupsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeClusterSubnetGroupsRequestPaginator returns a paginator for DescribeClusterSubnetGroups.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeClusterSubnetGroupsRequest(input)
//   p := redshift.NewDescribeClusterSubnetGroupsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeClusterSubnetGroupsPaginator(req DescribeClusterSubnetGroupsRequest) DescribeClusterSubnetGroupsPaginator {
	return DescribeClusterSubnetGroupsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeClusterSubnetGroupsInput
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

// DescribeClusterSubnetGroupsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeClusterSubnetGroupsPaginator struct {
	aws.Pager
}

func (p *DescribeClusterSubnetGroupsPaginator) CurrentPage() *DescribeClusterSubnetGroupsOutput {
	return p.Pager.CurrentPage().(*DescribeClusterSubnetGroupsOutput)
}

// DescribeClusterSubnetGroupsResponse is the response type for the
// DescribeClusterSubnetGroups API operation.
type DescribeClusterSubnetGroupsResponse struct {
	*DescribeClusterSubnetGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeClusterSubnetGroups request.
func (r *DescribeClusterSubnetGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
