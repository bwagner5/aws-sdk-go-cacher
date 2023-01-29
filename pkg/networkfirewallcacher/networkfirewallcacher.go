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
package networkfirewallcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
	"github.com/aws/aws-sdk-go/service/networkfirewall/networkfirewalliface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type NetworkFirewall struct {
	networkfirewalliface.NetworkFirewallAPI
	cache *cache.Cache
}

func New(networkfirewallapi networkfirewalliface.NetworkFirewallAPI) *NetworkFirewall {
	return &NetworkFirewall{
		NetworkFirewallAPI: networkfirewallapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *NetworkFirewall) DescribeFirewall(input *networkfirewall.DescribeFirewallInput) (*networkfirewall.DescribeFirewallOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeFirewallOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeFirewall(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeFirewallPolicy(input *networkfirewall.DescribeFirewallPolicyInput) (*networkfirewall.DescribeFirewallPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeFirewallPolicyOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeFirewallPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeFirewallPolicyWithContext(ctx context.Context, input *networkfirewall.DescribeFirewallPolicyInput, opts ...request.Option) (*networkfirewall.DescribeFirewallPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeFirewallPolicyOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeFirewallPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeFirewallWithContext(ctx context.Context, input *networkfirewall.DescribeFirewallInput, opts ...request.Option) (*networkfirewall.DescribeFirewallOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeFirewallOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeFirewallWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeLoggingConfiguration(input *networkfirewall.DescribeLoggingConfigurationInput) (*networkfirewall.DescribeLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeLoggingConfigurationOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeLoggingConfigurationWithContext(ctx context.Context, input *networkfirewall.DescribeLoggingConfigurationInput, opts ...request.Option) (*networkfirewall.DescribeLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeLoggingConfigurationOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeResourcePolicy(input *networkfirewall.DescribeResourcePolicyInput) (*networkfirewall.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeResourcePolicyOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeResourcePolicyWithContext(ctx context.Context, input *networkfirewall.DescribeResourcePolicyInput, opts ...request.Option) (*networkfirewall.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeResourcePolicyOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeRuleGroup(input *networkfirewall.DescribeRuleGroupInput) (*networkfirewall.DescribeRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeRuleGroupOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeRuleGroupMetadata(input *networkfirewall.DescribeRuleGroupMetadataInput) (*networkfirewall.DescribeRuleGroupMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeRuleGroupMetadataOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeRuleGroupMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeRuleGroupMetadataWithContext(ctx context.Context, input *networkfirewall.DescribeRuleGroupMetadataInput, opts ...request.Option) (*networkfirewall.DescribeRuleGroupMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeRuleGroupMetadataOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeRuleGroupMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) DescribeRuleGroupWithContext(ctx context.Context, input *networkfirewall.DescribeRuleGroupInput, opts ...request.Option) (*networkfirewall.DescribeRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.DescribeRuleGroupOutput), nil
	}
	out, err := c.NetworkFirewallAPI.DescribeRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListFirewallPolicies(input *networkfirewall.ListFirewallPoliciesInput) (*networkfirewall.ListFirewallPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListFirewallPoliciesOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListFirewallPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListFirewallPoliciesPages(input *networkfirewall.ListFirewallPoliciesInput, fn func(*networkfirewall.ListFirewallPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListFirewallPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListFirewallPoliciesOutput{}
	fnCacher := func(out *networkfirewall.ListFirewallPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListFirewallPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListFirewallPoliciesPagesWithContext(ctx context.Context, input *networkfirewall.ListFirewallPoliciesInput, fn func(*networkfirewall.ListFirewallPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListFirewallPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListFirewallPoliciesOutput{}
	fnCacher := func(out *networkfirewall.ListFirewallPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListFirewallPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListFirewallPoliciesWithContext(ctx context.Context, input *networkfirewall.ListFirewallPoliciesInput, opts ...request.Option) (*networkfirewall.ListFirewallPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListFirewallPoliciesOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListFirewallPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListFirewalls(input *networkfirewall.ListFirewallsInput) (*networkfirewall.ListFirewallsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListFirewallsOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListFirewalls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListFirewallsPages(input *networkfirewall.ListFirewallsInput, fn func(*networkfirewall.ListFirewallsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListFirewallsOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListFirewallsOutput{}
	fnCacher := func(out *networkfirewall.ListFirewallsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListFirewallsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListFirewallsPagesWithContext(ctx context.Context, input *networkfirewall.ListFirewallsInput, fn func(*networkfirewall.ListFirewallsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListFirewallsOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListFirewallsOutput{}
	fnCacher := func(out *networkfirewall.ListFirewallsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListFirewallsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListFirewallsWithContext(ctx context.Context, input *networkfirewall.ListFirewallsInput, opts ...request.Option) (*networkfirewall.ListFirewallsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListFirewallsOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListFirewallsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListRuleGroups(input *networkfirewall.ListRuleGroupsInput) (*networkfirewall.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListRuleGroupsOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListRuleGroupsPages(input *networkfirewall.ListRuleGroupsInput, fn func(*networkfirewall.ListRuleGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListRuleGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListRuleGroupsOutput{}
	fnCacher := func(out *networkfirewall.ListRuleGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListRuleGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListRuleGroupsPagesWithContext(ctx context.Context, input *networkfirewall.ListRuleGroupsInput, fn func(*networkfirewall.ListRuleGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListRuleGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListRuleGroupsOutput{}
	fnCacher := func(out *networkfirewall.ListRuleGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListRuleGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListRuleGroupsWithContext(ctx context.Context, input *networkfirewall.ListRuleGroupsInput, opts ...request.Option) (*networkfirewall.ListRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListRuleGroupsOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListTagsForResource(input *networkfirewall.ListTagsForResourceInput) (*networkfirewall.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListTagsForResourceOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NetworkFirewall) ListTagsForResourcePages(input *networkfirewall.ListTagsForResourceInput, fn func(*networkfirewall.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListTagsForResourceOutput{}
	fnCacher := func(out *networkfirewall.ListTagsForResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListTagsForResourcePagesWithContext(ctx context.Context, input *networkfirewall.ListTagsForResourceInput, fn func(*networkfirewall.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*networkfirewall.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &networkfirewall.ListTagsForResourceOutput{}
	fnCacher := func(out *networkfirewall.ListTagsForResourceOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.NetworkFirewallAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NetworkFirewall) ListTagsForResourceWithContext(ctx context.Context, input *networkfirewall.ListTagsForResourceInput, opts ...request.Option) (*networkfirewall.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*networkfirewall.ListTagsForResourceOutput), nil
	}
	out, err := c.NetworkFirewallAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
