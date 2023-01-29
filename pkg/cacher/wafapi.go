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
package cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/waf/wafiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WAF struct {
	wafiface.WAFAPI
	cache *cache.Cache
}

func NewWAF(wafapi wafiface.WAFAPI) *WAF {
	return &WAF{
		WAFAPI: wafapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WAF) GetByteMatchSet(input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetByteMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetByteMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetByteMatchSetWithContext(ctx context.Context, input *waf.GetByteMatchSetInput, opts ...request.Option) (*waf.GetByteMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetByteMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetByteMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetChangeToken(input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenOutput), nil
	}
	out, err := c.WAFAPI.GetChangeToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetChangeTokenStatus(input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenStatusOutput), nil
	}
	out, err := c.WAFAPI.GetChangeTokenStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetChangeTokenStatusWithContext(ctx context.Context, input *waf.GetChangeTokenStatusInput, opts ...request.Option) (*waf.GetChangeTokenStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenStatusOutput), nil
	}
	out, err := c.WAFAPI.GetChangeTokenStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetChangeTokenWithContext(ctx context.Context, input *waf.GetChangeTokenInput, opts ...request.Option) (*waf.GetChangeTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetChangeTokenOutput), nil
	}
	out, err := c.WAFAPI.GetChangeTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetGeoMatchSet(input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetGeoMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetGeoMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetGeoMatchSetWithContext(ctx context.Context, input *waf.GetGeoMatchSetInput, opts ...request.Option) (*waf.GetGeoMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetGeoMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetGeoMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetIPSet(input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetIPSetOutput), nil
	}
	out, err := c.WAFAPI.GetIPSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetIPSetWithContext(ctx context.Context, input *waf.GetIPSetInput, opts ...request.Option) (*waf.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetIPSetOutput), nil
	}
	out, err := c.WAFAPI.GetIPSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetLoggingConfiguration(input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFAPI.GetLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetLoggingConfigurationWithContext(ctx context.Context, input *waf.GetLoggingConfigurationInput, opts ...request.Option) (*waf.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFAPI.GetLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetPermissionPolicy(input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFAPI.GetPermissionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetPermissionPolicyWithContext(ctx context.Context, input *waf.GetPermissionPolicyInput, opts ...request.Option) (*waf.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFAPI.GetPermissionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRateBasedRule(input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleOutput), nil
	}
	out, err := c.WAFAPI.GetRateBasedRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRateBasedRuleManagedKeys(input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleManagedKeysOutput), nil
	}
	out, err := c.WAFAPI.GetRateBasedRuleManagedKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRateBasedRuleManagedKeysWithContext(ctx context.Context, input *waf.GetRateBasedRuleManagedKeysInput, opts ...request.Option) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleManagedKeysOutput), nil
	}
	out, err := c.WAFAPI.GetRateBasedRuleManagedKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRateBasedRuleWithContext(ctx context.Context, input *waf.GetRateBasedRuleInput, opts ...request.Option) (*waf.GetRateBasedRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRateBasedRuleOutput), nil
	}
	out, err := c.WAFAPI.GetRateBasedRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRegexMatchSet(input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetRegexMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRegexMatchSetWithContext(ctx context.Context, input *waf.GetRegexMatchSetInput, opts ...request.Option) (*waf.GetRegexMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetRegexMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRegexPatternSet(input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFAPI.GetRegexPatternSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRegexPatternSetWithContext(ctx context.Context, input *waf.GetRegexPatternSetInput, opts ...request.Option) (*waf.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFAPI.GetRegexPatternSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRule(input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleOutput), nil
	}
	out, err := c.WAFAPI.GetRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRuleGroup(input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleGroupOutput), nil
	}
	out, err := c.WAFAPI.GetRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRuleGroupWithContext(ctx context.Context, input *waf.GetRuleGroupInput, opts ...request.Option) (*waf.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleGroupOutput), nil
	}
	out, err := c.WAFAPI.GetRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetRuleWithContext(ctx context.Context, input *waf.GetRuleInput, opts ...request.Option) (*waf.GetRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetRuleOutput), nil
	}
	out, err := c.WAFAPI.GetRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSampledRequests(input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFAPI.GetSampledRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSampledRequestsWithContext(ctx context.Context, input *waf.GetSampledRequestsInput, opts ...request.Option) (*waf.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFAPI.GetSampledRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSizeConstraintSet(input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSizeConstraintSetOutput), nil
	}
	out, err := c.WAFAPI.GetSizeConstraintSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSizeConstraintSetWithContext(ctx context.Context, input *waf.GetSizeConstraintSetInput, opts ...request.Option) (*waf.GetSizeConstraintSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSizeConstraintSetOutput), nil
	}
	out, err := c.WAFAPI.GetSizeConstraintSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSqlInjectionMatchSet(input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSqlInjectionMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetSqlInjectionMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetSqlInjectionMatchSetWithContext(ctx context.Context, input *waf.GetSqlInjectionMatchSetInput, opts ...request.Option) (*waf.GetSqlInjectionMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetSqlInjectionMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetSqlInjectionMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetWebACL(input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetWebACLOutput), nil
	}
	out, err := c.WAFAPI.GetWebACL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetWebACLWithContext(ctx context.Context, input *waf.GetWebACLInput, opts ...request.Option) (*waf.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetWebACLOutput), nil
	}
	out, err := c.WAFAPI.GetWebACLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetXssMatchSet(input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetXssMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetXssMatchSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) GetXssMatchSetWithContext(ctx context.Context, input *waf.GetXssMatchSetInput, opts ...request.Option) (*waf.GetXssMatchSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.GetXssMatchSetOutput), nil
	}
	out, err := c.WAFAPI.GetXssMatchSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListActivatedRulesInRuleGroup(input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListActivatedRulesInRuleGroupOutput), nil
	}
	out, err := c.WAFAPI.ListActivatedRulesInRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListActivatedRulesInRuleGroupWithContext(ctx context.Context, input *waf.ListActivatedRulesInRuleGroupInput, opts ...request.Option) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListActivatedRulesInRuleGroupOutput), nil
	}
	out, err := c.WAFAPI.ListActivatedRulesInRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListByteMatchSets(input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListByteMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListByteMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListByteMatchSetsWithContext(ctx context.Context, input *waf.ListByteMatchSetsInput, opts ...request.Option) (*waf.ListByteMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListByteMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListByteMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListGeoMatchSets(input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListGeoMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListGeoMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListGeoMatchSetsWithContext(ctx context.Context, input *waf.ListGeoMatchSetsInput, opts ...request.Option) (*waf.ListGeoMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListGeoMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListGeoMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListIPSets(input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListIPSetsOutput), nil
	}
	out, err := c.WAFAPI.ListIPSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListIPSetsWithContext(ctx context.Context, input *waf.ListIPSetsInput, opts ...request.Option) (*waf.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListIPSetsOutput), nil
	}
	out, err := c.WAFAPI.ListIPSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListLoggingConfigurations(input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFAPI.ListLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListLoggingConfigurationsWithContext(ctx context.Context, input *waf.ListLoggingConfigurationsInput, opts ...request.Option) (*waf.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFAPI.ListLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRateBasedRules(input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRateBasedRulesOutput), nil
	}
	out, err := c.WAFAPI.ListRateBasedRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRateBasedRulesWithContext(ctx context.Context, input *waf.ListRateBasedRulesInput, opts ...request.Option) (*waf.ListRateBasedRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRateBasedRulesOutput), nil
	}
	out, err := c.WAFAPI.ListRateBasedRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRegexMatchSets(input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListRegexMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRegexMatchSetsWithContext(ctx context.Context, input *waf.ListRegexMatchSetsInput, opts ...request.Option) (*waf.ListRegexMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListRegexMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRegexPatternSets(input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFAPI.ListRegexPatternSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRegexPatternSetsWithContext(ctx context.Context, input *waf.ListRegexPatternSetsInput, opts ...request.Option) (*waf.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFAPI.ListRegexPatternSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRuleGroups(input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFAPI.ListRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRuleGroupsWithContext(ctx context.Context, input *waf.ListRuleGroupsInput, opts ...request.Option) (*waf.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFAPI.ListRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRules(input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRulesOutput), nil
	}
	out, err := c.WAFAPI.ListRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListRulesWithContext(ctx context.Context, input *waf.ListRulesInput, opts ...request.Option) (*waf.ListRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListRulesOutput), nil
	}
	out, err := c.WAFAPI.ListRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSizeConstraintSets(input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSizeConstraintSetsOutput), nil
	}
	out, err := c.WAFAPI.ListSizeConstraintSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSizeConstraintSetsWithContext(ctx context.Context, input *waf.ListSizeConstraintSetsInput, opts ...request.Option) (*waf.ListSizeConstraintSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSizeConstraintSetsOutput), nil
	}
	out, err := c.WAFAPI.ListSizeConstraintSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSqlInjectionMatchSets(input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSqlInjectionMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListSqlInjectionMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSqlInjectionMatchSetsWithContext(ctx context.Context, input *waf.ListSqlInjectionMatchSetsInput, opts ...request.Option) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSqlInjectionMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListSqlInjectionMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSubscribedRuleGroups(input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSubscribedRuleGroupsOutput), nil
	}
	out, err := c.WAFAPI.ListSubscribedRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListSubscribedRuleGroupsWithContext(ctx context.Context, input *waf.ListSubscribedRuleGroupsInput, opts ...request.Option) (*waf.ListSubscribedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListSubscribedRuleGroupsOutput), nil
	}
	out, err := c.WAFAPI.ListSubscribedRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListTagsForResource(input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListTagsForResourceWithContext(ctx context.Context, input *waf.ListTagsForResourceInput, opts ...request.Option) (*waf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListWebACLs(input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListWebACLsOutput), nil
	}
	out, err := c.WAFAPI.ListWebACLs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListWebACLsWithContext(ctx context.Context, input *waf.ListWebACLsInput, opts ...request.Option) (*waf.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListWebACLsOutput), nil
	}
	out, err := c.WAFAPI.ListWebACLsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListXssMatchSets(input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListXssMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListXssMatchSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAF) ListXssMatchSetsWithContext(ctx context.Context, input *waf.ListXssMatchSetsInput, opts ...request.Option) (*waf.ListXssMatchSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*waf.ListXssMatchSetsOutput), nil
	}
	out, err := c.WAFAPI.ListXssMatchSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
