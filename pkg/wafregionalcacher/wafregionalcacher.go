/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package wafregionalcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WAFRegional struct {
	wafregionaliface.WAFRegionalAPI
	cache *cache.Cache
}

func New(wafregionalapi wafregionaliface.WAFRegionalAPI) *WAFRegional {
	return &WAFRegional{
		WAFRegionalAPI: wafregionalapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WAFRegional) GetByteMatchSet(input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetByteMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetByteMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetByteMatchSetWithContext(ctx context.Context, input *waf.GetByteMatchSetInput, opts ...request.Option) (*waf.GetByteMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetByteMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetByteMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetChangeToken(input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetChangeToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetChangeTokenStatus(input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenStatusOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetChangeTokenStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetChangeTokenStatusWithContext(ctx context.Context, input *waf.GetChangeTokenStatusInput, opts ...request.Option) (*waf.GetChangeTokenStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenStatusOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetChangeTokenStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetChangeTokenWithContext(ctx context.Context, input *waf.GetChangeTokenInput, opts ...request.Option) (*waf.GetChangeTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetChangeTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetGeoMatchSet(input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetGeoMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetGeoMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetGeoMatchSetWithContext(ctx context.Context, input *waf.GetGeoMatchSetInput, opts ...request.Option) (*waf.GetGeoMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetGeoMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetGeoMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetIPSet(input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetIPSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetIPSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetIPSetWithContext(ctx context.Context, input *waf.GetIPSetInput, opts ...request.Option) (*waf.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetIPSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetIPSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetLoggingConfiguration(input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetLoggingConfigurationWithContext(ctx context.Context, input *waf.GetLoggingConfigurationInput, opts ...request.Option) (*waf.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetPermissionPolicy(input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetPermissionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetPermissionPolicyWithContext(ctx context.Context, input *waf.GetPermissionPolicyInput, opts ...request.Option) (*waf.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetPermissionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRateBasedRule(input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRateBasedRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRateBasedRuleManagedKeys(input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleManagedKeysOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRateBasedRuleManagedKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRateBasedRuleManagedKeysWithContext(ctx context.Context, input *waf.GetRateBasedRuleManagedKeysInput, opts ...request.Option) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleManagedKeysOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRateBasedRuleManagedKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRateBasedRuleWithContext(ctx context.Context, input *waf.GetRateBasedRuleInput, opts ...request.Option) (*waf.GetRateBasedRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRateBasedRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRegexMatchSet(input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRegexMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRegexMatchSetWithContext(ctx context.Context, input *waf.GetRegexMatchSetInput, opts ...request.Option) (*waf.GetRegexMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRegexMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRegexPatternSet(input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRegexPatternSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRegexPatternSetWithContext(ctx context.Context, input *waf.GetRegexPatternSetInput, opts ...request.Option) (*waf.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRegexPatternSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRule(input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRuleGroup(input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleGroupOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRuleGroupWithContext(ctx context.Context, input *waf.GetRuleGroupInput, opts ...request.Option) (*waf.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleGroupOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetRuleWithContext(ctx context.Context, input *waf.GetRuleInput, opts ...request.Option) (*waf.GetRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSampledRequests(input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSampledRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSampledRequestsWithContext(ctx context.Context, input *waf.GetSampledRequestsInput, opts ...request.Option) (*waf.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSampledRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSizeConstraintSet(input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSizeConstraintSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSizeConstraintSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSizeConstraintSetWithContext(ctx context.Context, input *waf.GetSizeConstraintSetInput, opts ...request.Option) (*waf.GetSizeConstraintSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSizeConstraintSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSizeConstraintSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSqlInjectionMatchSet(input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSqlInjectionMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSqlInjectionMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetSqlInjectionMatchSetWithContext(ctx context.Context, input *waf.GetSqlInjectionMatchSetInput, opts ...request.Option) (*waf.GetSqlInjectionMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSqlInjectionMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetSqlInjectionMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetWebACL(input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetWebACLOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetWebACL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetWebACLForResource(input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafregional.GetWebACLForResourceOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetWebACLForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetWebACLForResourceWithContext(ctx context.Context, input *wafregional.GetWebACLForResourceInput, opts ...request.Option) (*wafregional.GetWebACLForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafregional.GetWebACLForResourceOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetWebACLForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetWebACLWithContext(ctx context.Context, input *waf.GetWebACLInput, opts ...request.Option) (*waf.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetWebACLOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetWebACLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetXssMatchSet(input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetXssMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetXssMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) GetXssMatchSetWithContext(ctx context.Context, input *waf.GetXssMatchSetInput, opts ...request.Option) (*waf.GetXssMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetXssMatchSetOutput), nil
	}
	out, err := c.WAFRegionalAPI.GetXssMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListActivatedRulesInRuleGroup(input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListActivatedRulesInRuleGroupOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListActivatedRulesInRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListActivatedRulesInRuleGroupWithContext(ctx context.Context, input *waf.ListActivatedRulesInRuleGroupInput, opts ...request.Option) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListActivatedRulesInRuleGroupOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListActivatedRulesInRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListByteMatchSets(input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListByteMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListByteMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListByteMatchSetsWithContext(ctx context.Context, input *waf.ListByteMatchSetsInput, opts ...request.Option) (*waf.ListByteMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListByteMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListByteMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListGeoMatchSets(input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListGeoMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListGeoMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListGeoMatchSetsWithContext(ctx context.Context, input *waf.ListGeoMatchSetsInput, opts ...request.Option) (*waf.ListGeoMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListGeoMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListGeoMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListIPSets(input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListIPSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListIPSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListIPSetsWithContext(ctx context.Context, input *waf.ListIPSetsInput, opts ...request.Option) (*waf.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListIPSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListIPSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListLoggingConfigurations(input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListLoggingConfigurationsWithContext(ctx context.Context, input *waf.ListLoggingConfigurationsInput, opts ...request.Option) (*waf.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRateBasedRules(input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRateBasedRulesOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRateBasedRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRateBasedRulesWithContext(ctx context.Context, input *waf.ListRateBasedRulesInput, opts ...request.Option) (*waf.ListRateBasedRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRateBasedRulesOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRateBasedRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRegexMatchSets(input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRegexMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRegexMatchSetsWithContext(ctx context.Context, input *waf.ListRegexMatchSetsInput, opts ...request.Option) (*waf.ListRegexMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRegexMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRegexPatternSets(input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRegexPatternSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRegexPatternSetsWithContext(ctx context.Context, input *waf.ListRegexPatternSetsInput, opts ...request.Option) (*waf.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRegexPatternSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListResourcesForWebACL(input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafregional.ListResourcesForWebACLOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListResourcesForWebACL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListResourcesForWebACLWithContext(ctx context.Context, input *wafregional.ListResourcesForWebACLInput, opts ...request.Option) (*wafregional.ListResourcesForWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafregional.ListResourcesForWebACLOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListResourcesForWebACLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRuleGroups(input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRuleGroupsWithContext(ctx context.Context, input *waf.ListRuleGroupsInput, opts ...request.Option) (*waf.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRules(input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRulesOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListRulesWithContext(ctx context.Context, input *waf.ListRulesInput, opts ...request.Option) (*waf.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRulesOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSizeConstraintSets(input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSizeConstraintSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSizeConstraintSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSizeConstraintSetsWithContext(ctx context.Context, input *waf.ListSizeConstraintSetsInput, opts ...request.Option) (*waf.ListSizeConstraintSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSizeConstraintSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSizeConstraintSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSqlInjectionMatchSets(input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSqlInjectionMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSqlInjectionMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSqlInjectionMatchSetsWithContext(ctx context.Context, input *waf.ListSqlInjectionMatchSetsInput, opts ...request.Option) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSqlInjectionMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSqlInjectionMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSubscribedRuleGroups(input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSubscribedRuleGroupsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSubscribedRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListSubscribedRuleGroupsWithContext(ctx context.Context, input *waf.ListSubscribedRuleGroupsInput, opts ...request.Option) (*waf.ListSubscribedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSubscribedRuleGroupsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListSubscribedRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListTagsForResource(input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListTagsForResourceWithContext(ctx context.Context, input *waf.ListTagsForResourceInput, opts ...request.Option) (*waf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListWebACLs(input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListWebACLsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListWebACLs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListWebACLsWithContext(ctx context.Context, input *waf.ListWebACLsInput, opts ...request.Option) (*waf.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListWebACLsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListWebACLsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListXssMatchSets(input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListXssMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListXssMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFRegional) ListXssMatchSetsWithContext(ctx context.Context, input *waf.ListXssMatchSetsInput, opts ...request.Option) (*waf.ListXssMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListXssMatchSetsOutput), nil
	}
	out, err := c.WAFRegionalAPI.ListXssMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
