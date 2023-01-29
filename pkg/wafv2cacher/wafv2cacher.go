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
package wafv2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/wafv2"
	"github.com/aws/aws-sdk-go/service/wafv2/wafv2iface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WAFV2 struct {
	wafv2iface.WAFV2API
	cache *cache.Cache
}

func New(wafv2api wafv2iface.WAFV2API) *WAFV2 {
	return &WAFV2{
		WAFV2API: wafv2api,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WAFV2) DescribeManagedRuleGroup(input *wafv2.DescribeManagedRuleGroupInput) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.DescribeManagedRuleGroupOutput), nil
	}
	out, err := c.WAFV2API.DescribeManagedRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) DescribeManagedRuleGroupWithContext(ctx context.Context, input *wafv2.DescribeManagedRuleGroupInput, opts ...request.Option) (*wafv2.DescribeManagedRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.DescribeManagedRuleGroupOutput), nil
	}
	out, err := c.WAFV2API.DescribeManagedRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetIPSet(input *wafv2.GetIPSetInput) (*wafv2.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetIPSetOutput), nil
	}
	out, err := c.WAFV2API.GetIPSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetIPSetWithContext(ctx context.Context, input *wafv2.GetIPSetInput, opts ...request.Option) (*wafv2.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetIPSetOutput), nil
	}
	out, err := c.WAFV2API.GetIPSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetLoggingConfiguration(input *wafv2.GetLoggingConfigurationInput) (*wafv2.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFV2API.GetLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetLoggingConfigurationWithContext(ctx context.Context, input *wafv2.GetLoggingConfigurationInput, opts ...request.Option) (*wafv2.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetLoggingConfigurationOutput), nil
	}
	out, err := c.WAFV2API.GetLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetManagedRuleSet(input *wafv2.GetManagedRuleSetInput) (*wafv2.GetManagedRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetManagedRuleSetOutput), nil
	}
	out, err := c.WAFV2API.GetManagedRuleSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetManagedRuleSetWithContext(ctx context.Context, input *wafv2.GetManagedRuleSetInput, opts ...request.Option) (*wafv2.GetManagedRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetManagedRuleSetOutput), nil
	}
	out, err := c.WAFV2API.GetManagedRuleSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetMobileSdkRelease(input *wafv2.GetMobileSdkReleaseInput) (*wafv2.GetMobileSdkReleaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetMobileSdkReleaseOutput), nil
	}
	out, err := c.WAFV2API.GetMobileSdkRelease(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetMobileSdkReleaseWithContext(ctx context.Context, input *wafv2.GetMobileSdkReleaseInput, opts ...request.Option) (*wafv2.GetMobileSdkReleaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetMobileSdkReleaseOutput), nil
	}
	out, err := c.WAFV2API.GetMobileSdkReleaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetPermissionPolicy(input *wafv2.GetPermissionPolicyInput) (*wafv2.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFV2API.GetPermissionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetPermissionPolicyWithContext(ctx context.Context, input *wafv2.GetPermissionPolicyInput, opts ...request.Option) (*wafv2.GetPermissionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetPermissionPolicyOutput), nil
	}
	out, err := c.WAFV2API.GetPermissionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRateBasedStatementManagedKeys(input *wafv2.GetRateBasedStatementManagedKeysInput) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRateBasedStatementManagedKeysOutput), nil
	}
	out, err := c.WAFV2API.GetRateBasedStatementManagedKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRateBasedStatementManagedKeysWithContext(ctx context.Context, input *wafv2.GetRateBasedStatementManagedKeysInput, opts ...request.Option) (*wafv2.GetRateBasedStatementManagedKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRateBasedStatementManagedKeysOutput), nil
	}
	out, err := c.WAFV2API.GetRateBasedStatementManagedKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRegexPatternSet(input *wafv2.GetRegexPatternSetInput) (*wafv2.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFV2API.GetRegexPatternSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRegexPatternSetWithContext(ctx context.Context, input *wafv2.GetRegexPatternSetInput, opts ...request.Option) (*wafv2.GetRegexPatternSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRegexPatternSetOutput), nil
	}
	out, err := c.WAFV2API.GetRegexPatternSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRuleGroup(input *wafv2.GetRuleGroupInput) (*wafv2.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRuleGroupOutput), nil
	}
	out, err := c.WAFV2API.GetRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetRuleGroupWithContext(ctx context.Context, input *wafv2.GetRuleGroupInput, opts ...request.Option) (*wafv2.GetRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetRuleGroupOutput), nil
	}
	out, err := c.WAFV2API.GetRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetSampledRequests(input *wafv2.GetSampledRequestsInput) (*wafv2.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFV2API.GetSampledRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetSampledRequestsWithContext(ctx context.Context, input *wafv2.GetSampledRequestsInput, opts ...request.Option) (*wafv2.GetSampledRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetSampledRequestsOutput), nil
	}
	out, err := c.WAFV2API.GetSampledRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetWebACL(input *wafv2.GetWebACLInput) (*wafv2.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetWebACLOutput), nil
	}
	out, err := c.WAFV2API.GetWebACL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetWebACLForResource(input *wafv2.GetWebACLForResourceInput) (*wafv2.GetWebACLForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetWebACLForResourceOutput), nil
	}
	out, err := c.WAFV2API.GetWebACLForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetWebACLForResourceWithContext(ctx context.Context, input *wafv2.GetWebACLForResourceInput, opts ...request.Option) (*wafv2.GetWebACLForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetWebACLForResourceOutput), nil
	}
	out, err := c.WAFV2API.GetWebACLForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) GetWebACLWithContext(ctx context.Context, input *wafv2.GetWebACLInput, opts ...request.Option) (*wafv2.GetWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.GetWebACLOutput), nil
	}
	out, err := c.WAFV2API.GetWebACLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListAvailableManagedRuleGroupVersions(input *wafv2.ListAvailableManagedRuleGroupVersionsInput) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListAvailableManagedRuleGroupVersionsOutput), nil
	}
	out, err := c.WAFV2API.ListAvailableManagedRuleGroupVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListAvailableManagedRuleGroupVersionsWithContext(ctx context.Context, input *wafv2.ListAvailableManagedRuleGroupVersionsInput, opts ...request.Option) (*wafv2.ListAvailableManagedRuleGroupVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListAvailableManagedRuleGroupVersionsOutput), nil
	}
	out, err := c.WAFV2API.ListAvailableManagedRuleGroupVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListAvailableManagedRuleGroups(input *wafv2.ListAvailableManagedRuleGroupsInput) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListAvailableManagedRuleGroupsOutput), nil
	}
	out, err := c.WAFV2API.ListAvailableManagedRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListAvailableManagedRuleGroupsWithContext(ctx context.Context, input *wafv2.ListAvailableManagedRuleGroupsInput, opts ...request.Option) (*wafv2.ListAvailableManagedRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListAvailableManagedRuleGroupsOutput), nil
	}
	out, err := c.WAFV2API.ListAvailableManagedRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListIPSets(input *wafv2.ListIPSetsInput) (*wafv2.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListIPSetsOutput), nil
	}
	out, err := c.WAFV2API.ListIPSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListIPSetsWithContext(ctx context.Context, input *wafv2.ListIPSetsInput, opts ...request.Option) (*wafv2.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListIPSetsOutput), nil
	}
	out, err := c.WAFV2API.ListIPSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListLoggingConfigurations(input *wafv2.ListLoggingConfigurationsInput) (*wafv2.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFV2API.ListLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListLoggingConfigurationsWithContext(ctx context.Context, input *wafv2.ListLoggingConfigurationsInput, opts ...request.Option) (*wafv2.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.WAFV2API.ListLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListManagedRuleSets(input *wafv2.ListManagedRuleSetsInput) (*wafv2.ListManagedRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListManagedRuleSetsOutput), nil
	}
	out, err := c.WAFV2API.ListManagedRuleSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListManagedRuleSetsWithContext(ctx context.Context, input *wafv2.ListManagedRuleSetsInput, opts ...request.Option) (*wafv2.ListManagedRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListManagedRuleSetsOutput), nil
	}
	out, err := c.WAFV2API.ListManagedRuleSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListMobileSdkReleases(input *wafv2.ListMobileSdkReleasesInput) (*wafv2.ListMobileSdkReleasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListMobileSdkReleasesOutput), nil
	}
	out, err := c.WAFV2API.ListMobileSdkReleases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListMobileSdkReleasesWithContext(ctx context.Context, input *wafv2.ListMobileSdkReleasesInput, opts ...request.Option) (*wafv2.ListMobileSdkReleasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListMobileSdkReleasesOutput), nil
	}
	out, err := c.WAFV2API.ListMobileSdkReleasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListRegexPatternSets(input *wafv2.ListRegexPatternSetsInput) (*wafv2.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFV2API.ListRegexPatternSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListRegexPatternSetsWithContext(ctx context.Context, input *wafv2.ListRegexPatternSetsInput, opts ...request.Option) (*wafv2.ListRegexPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListRegexPatternSetsOutput), nil
	}
	out, err := c.WAFV2API.ListRegexPatternSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListResourcesForWebACL(input *wafv2.ListResourcesForWebACLInput) (*wafv2.ListResourcesForWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListResourcesForWebACLOutput), nil
	}
	out, err := c.WAFV2API.ListResourcesForWebACL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListResourcesForWebACLWithContext(ctx context.Context, input *wafv2.ListResourcesForWebACLInput, opts ...request.Option) (*wafv2.ListResourcesForWebACLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListResourcesForWebACLOutput), nil
	}
	out, err := c.WAFV2API.ListResourcesForWebACLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListRuleGroups(input *wafv2.ListRuleGroupsInput) (*wafv2.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFV2API.ListRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListRuleGroupsWithContext(ctx context.Context, input *wafv2.ListRuleGroupsInput, opts ...request.Option) (*wafv2.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListRuleGroupsOutput), nil
	}
	out, err := c.WAFV2API.ListRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListTagsForResource(input *wafv2.ListTagsForResourceInput) (*wafv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListTagsForResourceWithContext(ctx context.Context, input *wafv2.ListTagsForResourceInput, opts ...request.Option) (*wafv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListTagsForResourceOutput), nil
	}
	out, err := c.WAFV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListWebACLs(input *wafv2.ListWebACLsInput) (*wafv2.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListWebACLsOutput), nil
	}
	out, err := c.WAFV2API.ListWebACLs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WAFV2) ListWebACLsWithContext(ctx context.Context, input *wafv2.ListWebACLsInput, opts ...request.Option) (*wafv2.ListWebACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wafv2.ListWebACLsOutput), nil
	}
	out, err := c.WAFV2API.ListWebACLsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
