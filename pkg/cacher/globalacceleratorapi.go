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
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"github.com/aws/aws-sdk-go/service/globalaccelerator/globalacceleratoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GlobalAccelerator struct {
	globalacceleratoriface.GlobalAcceleratorAPI
	cache *cache.Cache
}

func NewGlobalAccelerator(globalacceleratorapi globalacceleratoriface.GlobalAcceleratorAPI) *GlobalAccelerator {
	return &GlobalAccelerator{
		GlobalAcceleratorAPI: globalacceleratorapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GlobalAccelerator) DescribeAccelerator(input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeAcceleratorOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeAccelerator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeAcceleratorAttributes(input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeAcceleratorAttributesOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeAcceleratorAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeAcceleratorAttributesWithContext(ctx context.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput, opts ...request.Option) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeAcceleratorAttributesOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeAcceleratorAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeAcceleratorWithContext(ctx context.Context, input *globalaccelerator.DescribeAcceleratorInput, opts ...request.Option) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeAcceleratorOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeAcceleratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingAccelerator(input *globalaccelerator.DescribeCustomRoutingAcceleratorInput) (*globalaccelerator.DescribeCustomRoutingAcceleratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingAcceleratorOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingAccelerator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingAcceleratorAttributes(input *globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput) (*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingAcceleratorAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingAcceleratorAttributesWithContext(ctx context.Context, input *globalaccelerator.DescribeCustomRoutingAcceleratorAttributesInput, opts ...request.Option) (*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingAcceleratorAttributesOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingAcceleratorAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingAcceleratorWithContext(ctx context.Context, input *globalaccelerator.DescribeCustomRoutingAcceleratorInput, opts ...request.Option) (*globalaccelerator.DescribeCustomRoutingAcceleratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingAcceleratorOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingAcceleratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingEndpointGroup(input *globalaccelerator.DescribeCustomRoutingEndpointGroupInput) (*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingEndpointGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingEndpointGroupWithContext(ctx context.Context, input *globalaccelerator.DescribeCustomRoutingEndpointGroupInput, opts ...request.Option) (*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingEndpointGroupOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingEndpointGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingListener(input *globalaccelerator.DescribeCustomRoutingListenerInput) (*globalaccelerator.DescribeCustomRoutingListenerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingListenerOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingListener(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeCustomRoutingListenerWithContext(ctx context.Context, input *globalaccelerator.DescribeCustomRoutingListenerInput, opts ...request.Option) (*globalaccelerator.DescribeCustomRoutingListenerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeCustomRoutingListenerOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeCustomRoutingListenerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeEndpointGroup(input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeEndpointGroupOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeEndpointGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeEndpointGroupWithContext(ctx context.Context, input *globalaccelerator.DescribeEndpointGroupInput, opts ...request.Option) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeEndpointGroupOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeEndpointGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeListener(input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeListenerOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeListener(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) DescribeListenerWithContext(ctx context.Context, input *globalaccelerator.DescribeListenerInput, opts ...request.Option) (*globalaccelerator.DescribeListenerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.DescribeListenerOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.DescribeListenerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListAccelerators(input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListAcceleratorsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListAccelerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListAcceleratorsPages(input *globalaccelerator.ListAcceleratorsInput, fn func(*globalaccelerator.ListAcceleratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListAcceleratorsOutput{}
	fnCacher := func(out *globalaccelerator.ListAcceleratorsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListAcceleratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListAcceleratorsPagesWithContext(ctx context.Context, input *globalaccelerator.ListAcceleratorsInput, fn func(*globalaccelerator.ListAcceleratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListAcceleratorsOutput{}
	fnCacher := func(out *globalaccelerator.ListAcceleratorsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListAcceleratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListAcceleratorsWithContext(ctx context.Context, input *globalaccelerator.ListAcceleratorsInput, opts ...request.Option) (*globalaccelerator.ListAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListAcceleratorsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListAcceleratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListByoipCidrs(input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListByoipCidrsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListByoipCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListByoipCidrsPages(input *globalaccelerator.ListByoipCidrsInput, fn func(*globalaccelerator.ListByoipCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListByoipCidrsOutput{}
	fnCacher := func(out *globalaccelerator.ListByoipCidrsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListByoipCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListByoipCidrsPagesWithContext(ctx context.Context, input *globalaccelerator.ListByoipCidrsInput, fn func(*globalaccelerator.ListByoipCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListByoipCidrsOutput{}
	fnCacher := func(out *globalaccelerator.ListByoipCidrsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListByoipCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListByoipCidrsWithContext(ctx context.Context, input *globalaccelerator.ListByoipCidrsInput, opts ...request.Option) (*globalaccelerator.ListByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListByoipCidrsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListByoipCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingAccelerators(input *globalaccelerator.ListCustomRoutingAcceleratorsInput) (*globalaccelerator.ListCustomRoutingAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingAcceleratorsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingAccelerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingAcceleratorsPages(input *globalaccelerator.ListCustomRoutingAcceleratorsInput, fn func(*globalaccelerator.ListCustomRoutingAcceleratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingAcceleratorsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingAcceleratorsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingAcceleratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingAcceleratorsPagesWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingAcceleratorsInput, fn func(*globalaccelerator.ListCustomRoutingAcceleratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingAcceleratorsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingAcceleratorsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingAcceleratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingAcceleratorsWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingAcceleratorsInput, opts ...request.Option) (*globalaccelerator.ListCustomRoutingAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingAcceleratorsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingAcceleratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingEndpointGroups(input *globalaccelerator.ListCustomRoutingEndpointGroupsInput) (*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingEndpointGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingEndpointGroupsPages(input *globalaccelerator.ListCustomRoutingEndpointGroupsInput, fn func(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingEndpointGroupsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingEndpointGroupsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingEndpointGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingEndpointGroupsPagesWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingEndpointGroupsInput, fn func(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingEndpointGroupsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingEndpointGroupsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingEndpointGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingEndpointGroupsWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingEndpointGroupsInput, opts ...request.Option) (*globalaccelerator.ListCustomRoutingEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingEndpointGroupsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingEndpointGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingListeners(input *globalaccelerator.ListCustomRoutingListenersInput) (*globalaccelerator.ListCustomRoutingListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingListenersOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingListeners(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingListenersPages(input *globalaccelerator.ListCustomRoutingListenersInput, fn func(*globalaccelerator.ListCustomRoutingListenersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingListenersOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingListenersOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingListenersOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingListenersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingListenersPagesWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingListenersInput, fn func(*globalaccelerator.ListCustomRoutingListenersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingListenersOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingListenersOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingListenersOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingListenersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingListenersWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingListenersInput, opts ...request.Option) (*globalaccelerator.ListCustomRoutingListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingListenersOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingListenersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappings(input *globalaccelerator.ListCustomRoutingPortMappingsInput) (*globalaccelerator.ListCustomRoutingPortMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsByDestination(input *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput) (*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsByDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsByDestinationPages(input *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, fn func(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsByDestinationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsByDestinationPagesWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, fn func(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsByDestinationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsByDestinationWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingPortMappingsByDestinationInput, opts ...request.Option) (*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsByDestinationOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsByDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsPages(input *globalaccelerator.ListCustomRoutingPortMappingsInput, fn func(*globalaccelerator.ListCustomRoutingPortMappingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingPortMappingsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingPortMappingsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsPagesWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingPortMappingsInput, fn func(*globalaccelerator.ListCustomRoutingPortMappingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListCustomRoutingPortMappingsOutput{}
	fnCacher := func(out *globalaccelerator.ListCustomRoutingPortMappingsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListCustomRoutingPortMappingsWithContext(ctx context.Context, input *globalaccelerator.ListCustomRoutingPortMappingsInput, opts ...request.Option) (*globalaccelerator.ListCustomRoutingPortMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListCustomRoutingPortMappingsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListCustomRoutingPortMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListEndpointGroups(input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListEndpointGroupsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListEndpointGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListEndpointGroupsPages(input *globalaccelerator.ListEndpointGroupsInput, fn func(*globalaccelerator.ListEndpointGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListEndpointGroupsOutput{}
	fnCacher := func(out *globalaccelerator.ListEndpointGroupsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListEndpointGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListEndpointGroupsPagesWithContext(ctx context.Context, input *globalaccelerator.ListEndpointGroupsInput, fn func(*globalaccelerator.ListEndpointGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListEndpointGroupsOutput{}
	fnCacher := func(out *globalaccelerator.ListEndpointGroupsOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListEndpointGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListEndpointGroupsWithContext(ctx context.Context, input *globalaccelerator.ListEndpointGroupsInput, opts ...request.Option) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListEndpointGroupsOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListEndpointGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListListeners(input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListListenersOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListListeners(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListListenersPages(input *globalaccelerator.ListListenersInput, fn func(*globalaccelerator.ListListenersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListListenersOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListListenersOutput{}
	fnCacher := func(out *globalaccelerator.ListListenersOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListListenersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListListenersPagesWithContext(ctx context.Context, input *globalaccelerator.ListListenersInput, fn func(*globalaccelerator.ListListenersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*globalaccelerator.ListListenersOutput), false)
		return nil
	}
	cachable := true
	output := &globalaccelerator.ListListenersOutput{}
	fnCacher := func(out *globalaccelerator.ListListenersOutput, more bool) bool {
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
	if err := c.GlobalAcceleratorAPI.ListListenersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlobalAccelerator) ListListenersWithContext(ctx context.Context, input *globalaccelerator.ListListenersInput, opts ...request.Option) (*globalaccelerator.ListListenersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListListenersOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListListenersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListTagsForResource(input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListTagsForResourceOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlobalAccelerator) ListTagsForResourceWithContext(ctx context.Context, input *globalaccelerator.ListTagsForResourceInput, opts ...request.Option) (*globalaccelerator.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*globalaccelerator.ListTagsForResourceOutput), nil
	}
	out, err := c.GlobalAcceleratorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
