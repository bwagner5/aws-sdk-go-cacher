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
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ELB struct {
	elbiface.ELBAPI
	cache *cache.Cache
}

func NewELB(elbapi elbiface.ELBAPI) *ELB {
	return &ELB{
		ELBAPI: elbapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ELB) DescribeAccountLimits(input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeAccountLimitsOutput), nil
	}
	out, err := c.ELBAPI.DescribeAccountLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeAccountLimitsWithContext(ctx context.Context, input *elb.DescribeAccountLimitsInput, opts ...request.Option) (*elb.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeAccountLimitsOutput), nil
	}
	out, err := c.ELBAPI.DescribeAccountLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeInstanceHealth(input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeInstanceHealthOutput), nil
	}
	out, err := c.ELBAPI.DescribeInstanceHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeInstanceHealthWithContext(ctx context.Context, input *elb.DescribeInstanceHealthInput, opts ...request.Option) (*elb.DescribeInstanceHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeInstanceHealthOutput), nil
	}
	out, err := c.ELBAPI.DescribeInstanceHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerAttributes(input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerAttributesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerAttributesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerAttributesInput, opts ...request.Option) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerAttributesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerPolicies(input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerPoliciesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerPoliciesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPoliciesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerPoliciesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerPolicyTypes(input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerPolicyTypesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerPolicyTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancerPolicyTypesWithContext(ctx context.Context, input *elb.DescribeLoadBalancerPolicyTypesInput, opts ...request.Option) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancerPolicyTypesOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancerPolicyTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancers(input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancersOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeLoadBalancersPages(input *elb.DescribeLoadBalancersInput, fn func(*elb.DescribeLoadBalancersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elb.DescribeLoadBalancersOutput), false)
		return nil
	}
	cachable := true
	output := &elb.DescribeLoadBalancersOutput{}
	fnCacher := func(out *elb.DescribeLoadBalancersOutput, more bool) bool {
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
	if err := c.ELBAPI.DescribeLoadBalancersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELB) DescribeLoadBalancersPagesWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, fn func(*elb.DescribeLoadBalancersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elb.DescribeLoadBalancersOutput), false)
		return nil
	}
	cachable := true
	output := &elb.DescribeLoadBalancersOutput{}
	fnCacher := func(out *elb.DescribeLoadBalancersOutput, more bool) bool {
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
	if err := c.ELBAPI.DescribeLoadBalancersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ELB) DescribeLoadBalancersWithContext(ctx context.Context, input *elb.DescribeLoadBalancersInput, opts ...request.Option) (*elb.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeLoadBalancersOutput), nil
	}
	out, err := c.ELBAPI.DescribeLoadBalancersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeTags(input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeTagsOutput), nil
	}
	out, err := c.ELBAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ELB) DescribeTagsWithContext(ctx context.Context, input *elb.DescribeTagsInput, opts ...request.Option) (*elb.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elb.DescribeTagsOutput), nil
	}
	out, err := c.ELBAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
