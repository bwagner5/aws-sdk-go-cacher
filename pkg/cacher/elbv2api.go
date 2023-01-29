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
	"github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/elbv2/elbv2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ELBV2 struct {
	elbv2iface.ELBV2API
	cache *cache.Cache
}

func NewELBV2(elbv2api elbv2iface.ELBV2API) *ELBV2 {
	return &ELBV2{
		ELBV2API: elbv2api,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ELBV2) DescribeAccountLimits(input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeAccountLimitsOutput), nil
	}
	out, err := c.ELBV2API.DescribeAccountLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeAccountLimitsWithContext(ctx context.Context, input *elbv2.DescribeAccountLimitsInput, opts ...request.Option) (*elbv2.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeAccountLimitsOutput), nil
	}
	out, err := c.ELBV2API.DescribeAccountLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeListenerCertificates(input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeListenerCertificatesOutput), nil
	}
	out, err := c.ELBV2API.DescribeListenerCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeListenerCertificatesWithContext(ctx context.Context, input *elbv2.DescribeListenerCertificatesInput, opts ...request.Option) (*elbv2.DescribeListenerCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeListenerCertificatesOutput), nil
	}
	out, err := c.ELBV2API.DescribeListenerCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeListeners(input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeListenersOutput), nil
	}
	out, err := c.ELBV2API.DescribeListeners(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeListenersPages(input *elbv2.DescribeListenersInput, fn func(*elbv2.DescribeListenersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeListenersOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeListenersOutput{}
	fnCacher := func(out *elbv2.DescribeListenersOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeListenersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeListenersPagesWithContext(ctx context.Context, input *elbv2.DescribeListenersInput, fn func(*elbv2.DescribeListenersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeListenersOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeListenersOutput{}
	fnCacher := func(out *elbv2.DescribeListenersOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeListenersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeListenersWithContext(ctx context.Context, input *elbv2.DescribeListenersInput, opts ...request.Option) (*elbv2.DescribeListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeListenersOutput), nil
	}
	out, err := c.ELBV2API.DescribeListenersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeLoadBalancerAttributes(input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeLoadBalancerAttributesOutput), nil
	}
	out, err := c.ELBV2API.DescribeLoadBalancerAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeLoadBalancerAttributesWithContext(ctx context.Context, input *elbv2.DescribeLoadBalancerAttributesInput, opts ...request.Option) (*elbv2.DescribeLoadBalancerAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeLoadBalancerAttributesOutput), nil
	}
	out, err := c.ELBV2API.DescribeLoadBalancerAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeLoadBalancers(input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeLoadBalancersOutput), nil
	}
	out, err := c.ELBV2API.DescribeLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeLoadBalancersPages(input *elbv2.DescribeLoadBalancersInput, fn func(*elbv2.DescribeLoadBalancersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeLoadBalancersOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeLoadBalancersOutput{}
	fnCacher := func(out *elbv2.DescribeLoadBalancersOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeLoadBalancersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeLoadBalancersPagesWithContext(ctx context.Context, input *elbv2.DescribeLoadBalancersInput, fn func(*elbv2.DescribeLoadBalancersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeLoadBalancersOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeLoadBalancersOutput{}
	fnCacher := func(out *elbv2.DescribeLoadBalancersOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeLoadBalancersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeLoadBalancersWithContext(ctx context.Context, input *elbv2.DescribeLoadBalancersInput, opts ...request.Option) (*elbv2.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeLoadBalancersOutput), nil
	}
	out, err := c.ELBV2API.DescribeLoadBalancersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeRules(input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeRulesOutput), nil
	}
	out, err := c.ELBV2API.DescribeRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeRulesWithContext(ctx context.Context, input *elbv2.DescribeRulesInput, opts ...request.Option) (*elbv2.DescribeRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeRulesOutput), nil
	}
	out, err := c.ELBV2API.DescribeRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeSSLPolicies(input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeSSLPoliciesOutput), nil
	}
	out, err := c.ELBV2API.DescribeSSLPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeSSLPoliciesWithContext(ctx context.Context, input *elbv2.DescribeSSLPoliciesInput, opts ...request.Option) (*elbv2.DescribeSSLPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeSSLPoliciesOutput), nil
	}
	out, err := c.ELBV2API.DescribeSSLPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTags(input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTagsOutput), nil
	}
	out, err := c.ELBV2API.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTagsWithContext(ctx context.Context, input *elbv2.DescribeTagsInput, opts ...request.Option) (*elbv2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTagsOutput), nil
	}
	out, err := c.ELBV2API.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetGroupAttributes(input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetGroupAttributesOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetGroupAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetGroupAttributesWithContext(ctx context.Context, input *elbv2.DescribeTargetGroupAttributesInput, opts ...request.Option) (*elbv2.DescribeTargetGroupAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetGroupAttributesOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetGroupAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetGroups(input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetGroupsOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetGroupsPages(input *elbv2.DescribeTargetGroupsInput, fn func(*elbv2.DescribeTargetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeTargetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeTargetGroupsOutput{}
	fnCacher := func(out *elbv2.DescribeTargetGroupsOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeTargetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeTargetGroupsPagesWithContext(ctx context.Context, input *elbv2.DescribeTargetGroupsInput, fn func(*elbv2.DescribeTargetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elbv2.DescribeTargetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elbv2.DescribeTargetGroupsOutput{}
	fnCacher := func(out *elbv2.DescribeTargetGroupsOutput, more bool) bool {
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
	if err := c.ELBV2API.DescribeTargetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELBV2) DescribeTargetGroupsWithContext(ctx context.Context, input *elbv2.DescribeTargetGroupsInput, opts ...request.Option) (*elbv2.DescribeTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetGroupsOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetHealth(input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetHealthOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELBV2) DescribeTargetHealthWithContext(ctx context.Context, input *elbv2.DescribeTargetHealthInput, opts ...request.Option) (*elbv2.DescribeTargetHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elbv2.DescribeTargetHealthOutput), nil
	}
	out, err := c.ELBV2API.DescribeTargetHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
