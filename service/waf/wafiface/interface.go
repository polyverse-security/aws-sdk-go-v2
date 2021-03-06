// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package wafiface provides an interface to enable mocking the AWS WAF service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package wafiface

import (
	"github.com/aws/aws-sdk-go-v2/service/waf"
)

// ClientAPI provides an interface to enable mocking the
// waf.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // WAF.
//    func myFunc(svc wafiface.ClientAPI) bool {
//        // Make svc.CreateByteMatchSet request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := waf.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        wafiface.ClientPI
//    }
//    func (m *mockClientClient) CreateByteMatchSet(input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	CreateByteMatchSetRequest(*waf.CreateByteMatchSetInput) waf.CreateByteMatchSetRequest

	CreateGeoMatchSetRequest(*waf.CreateGeoMatchSetInput) waf.CreateGeoMatchSetRequest

	CreateIPSetRequest(*waf.CreateIPSetInput) waf.CreateIPSetRequest

	CreateRateBasedRuleRequest(*waf.CreateRateBasedRuleInput) waf.CreateRateBasedRuleRequest

	CreateRegexMatchSetRequest(*waf.CreateRegexMatchSetInput) waf.CreateRegexMatchSetRequest

	CreateRegexPatternSetRequest(*waf.CreateRegexPatternSetInput) waf.CreateRegexPatternSetRequest

	CreateRuleRequest(*waf.CreateRuleInput) waf.CreateRuleRequest

	CreateRuleGroupRequest(*waf.CreateRuleGroupInput) waf.CreateRuleGroupRequest

	CreateSizeConstraintSetRequest(*waf.CreateSizeConstraintSetInput) waf.CreateSizeConstraintSetRequest

	CreateSqlInjectionMatchSetRequest(*waf.CreateSqlInjectionMatchSetInput) waf.CreateSqlInjectionMatchSetRequest

	CreateWebACLRequest(*waf.CreateWebACLInput) waf.CreateWebACLRequest

	CreateXssMatchSetRequest(*waf.CreateXssMatchSetInput) waf.CreateXssMatchSetRequest

	DeleteByteMatchSetRequest(*waf.DeleteByteMatchSetInput) waf.DeleteByteMatchSetRequest

	DeleteGeoMatchSetRequest(*waf.DeleteGeoMatchSetInput) waf.DeleteGeoMatchSetRequest

	DeleteIPSetRequest(*waf.DeleteIPSetInput) waf.DeleteIPSetRequest

	DeleteLoggingConfigurationRequest(*waf.DeleteLoggingConfigurationInput) waf.DeleteLoggingConfigurationRequest

	DeletePermissionPolicyRequest(*waf.DeletePermissionPolicyInput) waf.DeletePermissionPolicyRequest

	DeleteRateBasedRuleRequest(*waf.DeleteRateBasedRuleInput) waf.DeleteRateBasedRuleRequest

	DeleteRegexMatchSetRequest(*waf.DeleteRegexMatchSetInput) waf.DeleteRegexMatchSetRequest

	DeleteRegexPatternSetRequest(*waf.DeleteRegexPatternSetInput) waf.DeleteRegexPatternSetRequest

	DeleteRuleRequest(*waf.DeleteRuleInput) waf.DeleteRuleRequest

	DeleteRuleGroupRequest(*waf.DeleteRuleGroupInput) waf.DeleteRuleGroupRequest

	DeleteSizeConstraintSetRequest(*waf.DeleteSizeConstraintSetInput) waf.DeleteSizeConstraintSetRequest

	DeleteSqlInjectionMatchSetRequest(*waf.DeleteSqlInjectionMatchSetInput) waf.DeleteSqlInjectionMatchSetRequest

	DeleteWebACLRequest(*waf.DeleteWebACLInput) waf.DeleteWebACLRequest

	DeleteXssMatchSetRequest(*waf.DeleteXssMatchSetInput) waf.DeleteXssMatchSetRequest

	GetByteMatchSetRequest(*waf.GetByteMatchSetInput) waf.GetByteMatchSetRequest

	GetChangeTokenRequest(*waf.GetChangeTokenInput) waf.GetChangeTokenRequest

	GetChangeTokenStatusRequest(*waf.GetChangeTokenStatusInput) waf.GetChangeTokenStatusRequest

	GetGeoMatchSetRequest(*waf.GetGeoMatchSetInput) waf.GetGeoMatchSetRequest

	GetIPSetRequest(*waf.GetIPSetInput) waf.GetIPSetRequest

	GetLoggingConfigurationRequest(*waf.GetLoggingConfigurationInput) waf.GetLoggingConfigurationRequest

	GetPermissionPolicyRequest(*waf.GetPermissionPolicyInput) waf.GetPermissionPolicyRequest

	GetRateBasedRuleRequest(*waf.GetRateBasedRuleInput) waf.GetRateBasedRuleRequest

	GetRateBasedRuleManagedKeysRequest(*waf.GetRateBasedRuleManagedKeysInput) waf.GetRateBasedRuleManagedKeysRequest

	GetRegexMatchSetRequest(*waf.GetRegexMatchSetInput) waf.GetRegexMatchSetRequest

	GetRegexPatternSetRequest(*waf.GetRegexPatternSetInput) waf.GetRegexPatternSetRequest

	GetRuleRequest(*waf.GetRuleInput) waf.GetRuleRequest

	GetRuleGroupRequest(*waf.GetRuleGroupInput) waf.GetRuleGroupRequest

	GetSampledRequestsRequest(*waf.GetSampledRequestsInput) waf.GetSampledRequestsRequest

	GetSizeConstraintSetRequest(*waf.GetSizeConstraintSetInput) waf.GetSizeConstraintSetRequest

	GetSqlInjectionMatchSetRequest(*waf.GetSqlInjectionMatchSetInput) waf.GetSqlInjectionMatchSetRequest

	GetWebACLRequest(*waf.GetWebACLInput) waf.GetWebACLRequest

	GetXssMatchSetRequest(*waf.GetXssMatchSetInput) waf.GetXssMatchSetRequest

	ListActivatedRulesInRuleGroupRequest(*waf.ListActivatedRulesInRuleGroupInput) waf.ListActivatedRulesInRuleGroupRequest

	ListByteMatchSetsRequest(*waf.ListByteMatchSetsInput) waf.ListByteMatchSetsRequest

	ListGeoMatchSetsRequest(*waf.ListGeoMatchSetsInput) waf.ListGeoMatchSetsRequest

	ListIPSetsRequest(*waf.ListIPSetsInput) waf.ListIPSetsRequest

	ListLoggingConfigurationsRequest(*waf.ListLoggingConfigurationsInput) waf.ListLoggingConfigurationsRequest

	ListRateBasedRulesRequest(*waf.ListRateBasedRulesInput) waf.ListRateBasedRulesRequest

	ListRegexMatchSetsRequest(*waf.ListRegexMatchSetsInput) waf.ListRegexMatchSetsRequest

	ListRegexPatternSetsRequest(*waf.ListRegexPatternSetsInput) waf.ListRegexPatternSetsRequest

	ListRuleGroupsRequest(*waf.ListRuleGroupsInput) waf.ListRuleGroupsRequest

	ListRulesRequest(*waf.ListRulesInput) waf.ListRulesRequest

	ListSizeConstraintSetsRequest(*waf.ListSizeConstraintSetsInput) waf.ListSizeConstraintSetsRequest

	ListSqlInjectionMatchSetsRequest(*waf.ListSqlInjectionMatchSetsInput) waf.ListSqlInjectionMatchSetsRequest

	ListSubscribedRuleGroupsRequest(*waf.ListSubscribedRuleGroupsInput) waf.ListSubscribedRuleGroupsRequest

	ListWebACLsRequest(*waf.ListWebACLsInput) waf.ListWebACLsRequest

	ListXssMatchSetsRequest(*waf.ListXssMatchSetsInput) waf.ListXssMatchSetsRequest

	PutLoggingConfigurationRequest(*waf.PutLoggingConfigurationInput) waf.PutLoggingConfigurationRequest

	PutPermissionPolicyRequest(*waf.PutPermissionPolicyInput) waf.PutPermissionPolicyRequest

	UpdateByteMatchSetRequest(*waf.UpdateByteMatchSetInput) waf.UpdateByteMatchSetRequest

	UpdateGeoMatchSetRequest(*waf.UpdateGeoMatchSetInput) waf.UpdateGeoMatchSetRequest

	UpdateIPSetRequest(*waf.UpdateIPSetInput) waf.UpdateIPSetRequest

	UpdateRateBasedRuleRequest(*waf.UpdateRateBasedRuleInput) waf.UpdateRateBasedRuleRequest

	UpdateRegexMatchSetRequest(*waf.UpdateRegexMatchSetInput) waf.UpdateRegexMatchSetRequest

	UpdateRegexPatternSetRequest(*waf.UpdateRegexPatternSetInput) waf.UpdateRegexPatternSetRequest

	UpdateRuleRequest(*waf.UpdateRuleInput) waf.UpdateRuleRequest

	UpdateRuleGroupRequest(*waf.UpdateRuleGroupInput) waf.UpdateRuleGroupRequest

	UpdateSizeConstraintSetRequest(*waf.UpdateSizeConstraintSetInput) waf.UpdateSizeConstraintSetRequest

	UpdateSqlInjectionMatchSetRequest(*waf.UpdateSqlInjectionMatchSetInput) waf.UpdateSqlInjectionMatchSetRequest

	UpdateWebACLRequest(*waf.UpdateWebACLInput) waf.UpdateWebACLRequest

	UpdateXssMatchSetRequest(*waf.UpdateXssMatchSetInput) waf.UpdateXssMatchSetRequest
}

var _ ClientAPI = (*waf.Client)(nil)
