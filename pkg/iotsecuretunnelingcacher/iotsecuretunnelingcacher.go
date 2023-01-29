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
package iotsecuretunnelingcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling/iotsecuretunnelingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTSecureTunneling struct {
	iotsecuretunnelingiface.IoTSecureTunnelingAPI
	cache *cache.Cache
}

func New(iotsecuretunnelingapi iotsecuretunnelingiface.IoTSecureTunnelingAPI) *IoTSecureTunneling {
	return &IoTSecureTunneling{
		IoTSecureTunnelingAPI: iotsecuretunnelingapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTSecureTunneling) DescribeTunnel(input *iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.DescribeTunnelOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.DescribeTunnel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSecureTunneling) DescribeTunnelWithContext(ctx context.Context, input *iotsecuretunneling.DescribeTunnelInput, opts ...request.Option) (*iotsecuretunneling.DescribeTunnelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.DescribeTunnelOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.DescribeTunnelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSecureTunneling) ListTagsForResource(input *iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSecureTunneling) ListTagsForResourceWithContext(ctx context.Context, input *iotsecuretunneling.ListTagsForResourceInput, opts ...request.Option) (*iotsecuretunneling.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSecureTunneling) ListTunnels(input *iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.ListTunnelsOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.ListTunnels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSecureTunneling) ListTunnelsPages(input *iotsecuretunneling.ListTunnelsInput, fn func(*iotsecuretunneling.ListTunnelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsecuretunneling.ListTunnelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsecuretunneling.ListTunnelsOutput{}
	fnCacher := func(out *iotsecuretunneling.ListTunnelsOutput, more bool) bool {
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
	if err := c.IoTSecureTunnelingAPI.ListTunnelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSecureTunneling) ListTunnelsPagesWithContext(ctx context.Context, input *iotsecuretunneling.ListTunnelsInput, fn func(*iotsecuretunneling.ListTunnelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsecuretunneling.ListTunnelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsecuretunneling.ListTunnelsOutput{}
	fnCacher := func(out *iotsecuretunneling.ListTunnelsOutput, more bool) bool {
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
	if err := c.IoTSecureTunnelingAPI.ListTunnelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSecureTunneling) ListTunnelsWithContext(ctx context.Context, input *iotsecuretunneling.ListTunnelsInput, opts ...request.Option) (*iotsecuretunneling.ListTunnelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsecuretunneling.ListTunnelsOutput), nil
	}
	out, err := c.IoTSecureTunnelingAPI.ListTunnelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
