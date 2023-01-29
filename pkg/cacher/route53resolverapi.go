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
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"github.com/aws/aws-sdk-go/service/route53resolver/route53resolveriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Route53Resolver struct {
	route53resolveriface.Route53ResolverAPI
	cache *cache.Cache
}

func NewRoute53Resolver(route53resolverapi route53resolveriface.Route53ResolverAPI) *Route53Resolver {
	return &Route53Resolver{
		Route53ResolverAPI: route53resolverapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Route53Resolver) GetFirewallConfig(input *route53resolver.GetFirewallConfigInput) (*route53resolver.GetFirewallConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallConfigWithContext(ctx context.Context, input *route53resolver.GetFirewallConfigInput, opts ...request.Option) (*route53resolver.GetFirewallConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallDomainList(input *route53resolver.GetFirewallDomainListInput) (*route53resolver.GetFirewallDomainListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallDomainListOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallDomainList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallDomainListWithContext(ctx context.Context, input *route53resolver.GetFirewallDomainListInput, opts ...request.Option) (*route53resolver.GetFirewallDomainListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallDomainListOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallDomainListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroup(input *route53resolver.GetFirewallRuleGroupInput) (*route53resolver.GetFirewallRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroupAssociation(input *route53resolver.GetFirewallRuleGroupAssociationInput) (*route53resolver.GetFirewallRuleGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroupAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroupAssociationWithContext(ctx context.Context, input *route53resolver.GetFirewallRuleGroupAssociationInput, opts ...request.Option) (*route53resolver.GetFirewallRuleGroupAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroupAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroupPolicy(input *route53resolver.GetFirewallRuleGroupPolicyInput) (*route53resolver.GetFirewallRuleGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupPolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroupPolicyWithContext(ctx context.Context, input *route53resolver.GetFirewallRuleGroupPolicyInput, opts ...request.Option) (*route53resolver.GetFirewallRuleGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupPolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetFirewallRuleGroupWithContext(ctx context.Context, input *route53resolver.GetFirewallRuleGroupInput, opts ...request.Option) (*route53resolver.GetFirewallRuleGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetFirewallRuleGroupOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetFirewallRuleGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverConfig(input *route53resolver.GetResolverConfigInput) (*route53resolver.GetResolverConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverConfigWithContext(ctx context.Context, input *route53resolver.GetResolverConfigInput, opts ...request.Option) (*route53resolver.GetResolverConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverDnssecConfig(input *route53resolver.GetResolverDnssecConfigInput) (*route53resolver.GetResolverDnssecConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverDnssecConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverDnssecConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverDnssecConfigWithContext(ctx context.Context, input *route53resolver.GetResolverDnssecConfigInput, opts ...request.Option) (*route53resolver.GetResolverDnssecConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverDnssecConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverDnssecConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverEndpoint(input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverEndpointOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverEndpointWithContext(ctx context.Context, input *route53resolver.GetResolverEndpointInput, opts ...request.Option) (*route53resolver.GetResolverEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverEndpointOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfig(input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfigAssociation(input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfigAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfigAssociationWithContext(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput, opts ...request.Option) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfigAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfigPolicy(input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigPolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfigPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfigPolicyWithContext(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput, opts ...request.Option) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigPolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfigPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverQueryLogConfigWithContext(ctx context.Context, input *route53resolver.GetResolverQueryLogConfigInput, opts ...request.Option) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverQueryLogConfigOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverQueryLogConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRule(input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRuleOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRuleAssociation(input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRuleAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRuleAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRuleAssociationWithContext(ctx context.Context, input *route53resolver.GetResolverRuleAssociationInput, opts ...request.Option) (*route53resolver.GetResolverRuleAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRuleAssociationOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRuleAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRulePolicy(input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRulePolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRulePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRulePolicyWithContext(ctx context.Context, input *route53resolver.GetResolverRulePolicyInput, opts ...request.Option) (*route53resolver.GetResolverRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRulePolicyOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRulePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) GetResolverRuleWithContext(ctx context.Context, input *route53resolver.GetResolverRuleInput, opts ...request.Option) (*route53resolver.GetResolverRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.GetResolverRuleOutput), nil
	}
	out, err := c.Route53ResolverAPI.GetResolverRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallConfigs(input *route53resolver.ListFirewallConfigsInput) (*route53resolver.ListFirewallConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallConfigsPages(input *route53resolver.ListFirewallConfigsInput, fn func(*route53resolver.ListFirewallConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallConfigsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallConfigsPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallConfigsInput, fn func(*route53resolver.ListFirewallConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallConfigsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallConfigsWithContext(ctx context.Context, input *route53resolver.ListFirewallConfigsInput, opts ...request.Option) (*route53resolver.ListFirewallConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallDomainLists(input *route53resolver.ListFirewallDomainListsInput) (*route53resolver.ListFirewallDomainListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallDomainListsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallDomainLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallDomainListsPages(input *route53resolver.ListFirewallDomainListsInput, fn func(*route53resolver.ListFirewallDomainListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallDomainListsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallDomainListsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallDomainListsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallDomainListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallDomainListsPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallDomainListsInput, fn func(*route53resolver.ListFirewallDomainListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallDomainListsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallDomainListsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallDomainListsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallDomainListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallDomainListsWithContext(ctx context.Context, input *route53resolver.ListFirewallDomainListsInput, opts ...request.Option) (*route53resolver.ListFirewallDomainListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallDomainListsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallDomainListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallDomains(input *route53resolver.ListFirewallDomainsInput) (*route53resolver.ListFirewallDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallDomainsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallDomainsPages(input *route53resolver.ListFirewallDomainsInput, fn func(*route53resolver.ListFirewallDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallDomainsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallDomainsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallDomainsPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallDomainsInput, fn func(*route53resolver.ListFirewallDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallDomainsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallDomainsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallDomainsWithContext(ctx context.Context, input *route53resolver.ListFirewallDomainsInput, opts ...request.Option) (*route53resolver.ListFirewallDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallDomainsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRuleGroupAssociations(input *route53resolver.ListFirewallRuleGroupAssociationsInput) (*route53resolver.ListFirewallRuleGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRuleGroupAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRuleGroupAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRuleGroupAssociationsPages(input *route53resolver.ListFirewallRuleGroupAssociationsInput, fn func(*route53resolver.ListFirewallRuleGroupAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRuleGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRuleGroupAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRuleGroupAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRuleGroupAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRuleGroupAssociationsPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallRuleGroupAssociationsInput, fn func(*route53resolver.ListFirewallRuleGroupAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRuleGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRuleGroupAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRuleGroupAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRuleGroupAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRuleGroupAssociationsWithContext(ctx context.Context, input *route53resolver.ListFirewallRuleGroupAssociationsInput, opts ...request.Option) (*route53resolver.ListFirewallRuleGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRuleGroupAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRuleGroupAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRuleGroups(input *route53resolver.ListFirewallRuleGroupsInput) (*route53resolver.ListFirewallRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRuleGroupsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRuleGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRuleGroupsPages(input *route53resolver.ListFirewallRuleGroupsInput, fn func(*route53resolver.ListFirewallRuleGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRuleGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRuleGroupsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRuleGroupsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRuleGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRuleGroupsPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallRuleGroupsInput, fn func(*route53resolver.ListFirewallRuleGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRuleGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRuleGroupsOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRuleGroupsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRuleGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRuleGroupsWithContext(ctx context.Context, input *route53resolver.ListFirewallRuleGroupsInput, opts ...request.Option) (*route53resolver.ListFirewallRuleGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRuleGroupsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRuleGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRules(input *route53resolver.ListFirewallRulesInput) (*route53resolver.ListFirewallRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRulesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListFirewallRulesPages(input *route53resolver.ListFirewallRulesInput, fn func(*route53resolver.ListFirewallRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRulesOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRulesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRulesPagesWithContext(ctx context.Context, input *route53resolver.ListFirewallRulesInput, fn func(*route53resolver.ListFirewallRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListFirewallRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListFirewallRulesOutput{}
	fnCacher := func(out *route53resolver.ListFirewallRulesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListFirewallRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListFirewallRulesWithContext(ctx context.Context, input *route53resolver.ListFirewallRulesInput, opts ...request.Option) (*route53resolver.ListFirewallRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListFirewallRulesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListFirewallRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverConfigs(input *route53resolver.ListResolverConfigsInput) (*route53resolver.ListResolverConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverConfigsPages(input *route53resolver.ListResolverConfigsInput, fn func(*route53resolver.ListResolverConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverConfigsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverConfigsInput, fn func(*route53resolver.ListResolverConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverConfigsWithContext(ctx context.Context, input *route53resolver.ListResolverConfigsInput, opts ...request.Option) (*route53resolver.ListResolverConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverDnssecConfigs(input *route53resolver.ListResolverDnssecConfigsInput) (*route53resolver.ListResolverDnssecConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverDnssecConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverDnssecConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverDnssecConfigsPages(input *route53resolver.ListResolverDnssecConfigsInput, fn func(*route53resolver.ListResolverDnssecConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverDnssecConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverDnssecConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverDnssecConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverDnssecConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverDnssecConfigsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverDnssecConfigsInput, fn func(*route53resolver.ListResolverDnssecConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverDnssecConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverDnssecConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverDnssecConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverDnssecConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverDnssecConfigsWithContext(ctx context.Context, input *route53resolver.ListResolverDnssecConfigsInput, opts ...request.Option) (*route53resolver.ListResolverDnssecConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverDnssecConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverDnssecConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverEndpointIpAddresses(input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverEndpointIpAddressesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverEndpointIpAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverEndpointIpAddressesPages(input *route53resolver.ListResolverEndpointIpAddressesInput, fn func(*route53resolver.ListResolverEndpointIpAddressesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverEndpointIpAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverEndpointIpAddressesOutput{}
	fnCacher := func(out *route53resolver.ListResolverEndpointIpAddressesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverEndpointIpAddressesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverEndpointIpAddressesPagesWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput, fn func(*route53resolver.ListResolverEndpointIpAddressesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverEndpointIpAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverEndpointIpAddressesOutput{}
	fnCacher := func(out *route53resolver.ListResolverEndpointIpAddressesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverEndpointIpAddressesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverEndpointIpAddressesWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointIpAddressesInput, opts ...request.Option) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverEndpointIpAddressesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverEndpointIpAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverEndpoints(input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverEndpointsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverEndpointsPages(input *route53resolver.ListResolverEndpointsInput, fn func(*route53resolver.ListResolverEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverEndpointsOutput{}
	fnCacher := func(out *route53resolver.ListResolverEndpointsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverEndpointsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointsInput, fn func(*route53resolver.ListResolverEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverEndpointsOutput{}
	fnCacher := func(out *route53resolver.ListResolverEndpointsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverEndpointsWithContext(ctx context.Context, input *route53resolver.ListResolverEndpointsInput, opts ...request.Option) (*route53resolver.ListResolverEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverEndpointsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverQueryLogConfigAssociations(input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverQueryLogConfigAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverQueryLogConfigAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverQueryLogConfigAssociationsPages(input *route53resolver.ListResolverQueryLogConfigAssociationsInput, fn func(*route53resolver.ListResolverQueryLogConfigAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverQueryLogConfigAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverQueryLogConfigAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListResolverQueryLogConfigAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverQueryLogConfigAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverQueryLogConfigAssociationsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput, fn func(*route53resolver.ListResolverQueryLogConfigAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverQueryLogConfigAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverQueryLogConfigAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListResolverQueryLogConfigAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverQueryLogConfigAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverQueryLogConfigAssociationsWithContext(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput, opts ...request.Option) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverQueryLogConfigAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverQueryLogConfigAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverQueryLogConfigs(input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverQueryLogConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverQueryLogConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverQueryLogConfigsPages(input *route53resolver.ListResolverQueryLogConfigsInput, fn func(*route53resolver.ListResolverQueryLogConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverQueryLogConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverQueryLogConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverQueryLogConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverQueryLogConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverQueryLogConfigsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigsInput, fn func(*route53resolver.ListResolverQueryLogConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverQueryLogConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverQueryLogConfigsOutput{}
	fnCacher := func(out *route53resolver.ListResolverQueryLogConfigsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverQueryLogConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverQueryLogConfigsWithContext(ctx context.Context, input *route53resolver.ListResolverQueryLogConfigsInput, opts ...request.Option) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverQueryLogConfigsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverQueryLogConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverRuleAssociations(input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverRuleAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverRuleAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverRuleAssociationsPages(input *route53resolver.ListResolverRuleAssociationsInput, fn func(*route53resolver.ListResolverRuleAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverRuleAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverRuleAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListResolverRuleAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverRuleAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverRuleAssociationsPagesWithContext(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput, fn func(*route53resolver.ListResolverRuleAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverRuleAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverRuleAssociationsOutput{}
	fnCacher := func(out *route53resolver.ListResolverRuleAssociationsOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverRuleAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverRuleAssociationsWithContext(ctx context.Context, input *route53resolver.ListResolverRuleAssociationsInput, opts ...request.Option) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverRuleAssociationsOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverRuleAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverRules(input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverRulesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListResolverRulesPages(input *route53resolver.ListResolverRulesInput, fn func(*route53resolver.ListResolverRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverRulesOutput{}
	fnCacher := func(out *route53resolver.ListResolverRulesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverRulesPagesWithContext(ctx context.Context, input *route53resolver.ListResolverRulesInput, fn func(*route53resolver.ListResolverRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListResolverRulesOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListResolverRulesOutput{}
	fnCacher := func(out *route53resolver.ListResolverRulesOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListResolverRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListResolverRulesWithContext(ctx context.Context, input *route53resolver.ListResolverRulesInput, opts ...request.Option) (*route53resolver.ListResolverRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListResolverRulesOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListResolverRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListTagsForResource(input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Route53Resolver) ListTagsForResourcePages(input *route53resolver.ListTagsForResourceInput, fn func(*route53resolver.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListTagsForResourceOutput{}
	fnCacher := func(out *route53resolver.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListTagsForResourcePagesWithContext(ctx context.Context, input *route53resolver.ListTagsForResourceInput, fn func(*route53resolver.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*route53resolver.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &route53resolver.ListTagsForResourceOutput{}
	fnCacher := func(out *route53resolver.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.Route53ResolverAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Route53Resolver) ListTagsForResourceWithContext(ctx context.Context, input *route53resolver.ListTagsForResourceInput, opts ...request.Option) (*route53resolver.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*route53resolver.ListTagsForResourceOutput), nil
	}
	out, err := c.Route53ResolverAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
