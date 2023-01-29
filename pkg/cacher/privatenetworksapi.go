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
	"github.com/aws/aws-sdk-go/service/privatenetworks"
	"github.com/aws/aws-sdk-go/service/privatenetworks/privatenetworksiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PrivateNetworks struct {
	privatenetworksiface.PrivateNetworksAPI
	cache *cache.Cache
}

func NewPrivateNetworks(privatenetworksapi privatenetworksiface.PrivateNetworksAPI) *PrivateNetworks {
	return &PrivateNetworks{
		PrivateNetworksAPI: privatenetworksapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PrivateNetworks) GetDeviceIdentifier(input *privatenetworks.GetDeviceIdentifierInput) (*privatenetworks.GetDeviceIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetDeviceIdentifierOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetDeviceIdentifier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetDeviceIdentifierWithContext(ctx context.Context, input *privatenetworks.GetDeviceIdentifierInput, opts ...request.Option) (*privatenetworks.GetDeviceIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetDeviceIdentifierOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetDeviceIdentifierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetwork(input *privatenetworks.GetNetworkInput) (*privatenetworks.GetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetwork(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetworkResource(input *privatenetworks.GetNetworkResourceInput) (*privatenetworks.GetNetworkResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkResourceOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetworkResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetworkResourceWithContext(ctx context.Context, input *privatenetworks.GetNetworkResourceInput, opts ...request.Option) (*privatenetworks.GetNetworkResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkResourceOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetworkResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetworkSite(input *privatenetworks.GetNetworkSiteInput) (*privatenetworks.GetNetworkSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkSiteOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetworkSite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetworkSiteWithContext(ctx context.Context, input *privatenetworks.GetNetworkSiteInput, opts ...request.Option) (*privatenetworks.GetNetworkSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkSiteOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetworkSiteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetNetworkWithContext(ctx context.Context, input *privatenetworks.GetNetworkInput, opts ...request.Option) (*privatenetworks.GetNetworkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetNetworkOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetNetworkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetOrder(input *privatenetworks.GetOrderInput) (*privatenetworks.GetOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetOrderOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetOrder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) GetOrderWithContext(ctx context.Context, input *privatenetworks.GetOrderInput, opts ...request.Option) (*privatenetworks.GetOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.GetOrderOutput), nil
	}
	out, err := c.PrivateNetworksAPI.GetOrderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListDeviceIdentifiers(input *privatenetworks.ListDeviceIdentifiersInput) (*privatenetworks.ListDeviceIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListDeviceIdentifiersOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListDeviceIdentifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListDeviceIdentifiersPages(input *privatenetworks.ListDeviceIdentifiersInput, fn func(*privatenetworks.ListDeviceIdentifiersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListDeviceIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListDeviceIdentifiersOutput{}
	fnCacher := func(out *privatenetworks.ListDeviceIdentifiersOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListDeviceIdentifiersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListDeviceIdentifiersPagesWithContext(ctx context.Context, input *privatenetworks.ListDeviceIdentifiersInput, fn func(*privatenetworks.ListDeviceIdentifiersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListDeviceIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListDeviceIdentifiersOutput{}
	fnCacher := func(out *privatenetworks.ListDeviceIdentifiersOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListDeviceIdentifiersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListDeviceIdentifiersWithContext(ctx context.Context, input *privatenetworks.ListDeviceIdentifiersInput, opts ...request.Option) (*privatenetworks.ListDeviceIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListDeviceIdentifiersOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListDeviceIdentifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworkResources(input *privatenetworks.ListNetworkResourcesInput) (*privatenetworks.ListNetworkResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworkResourcesOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworkResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworkResourcesPages(input *privatenetworks.ListNetworkResourcesInput, fn func(*privatenetworks.ListNetworkResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworkResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworkResourcesOutput{}
	fnCacher := func(out *privatenetworks.ListNetworkResourcesOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworkResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworkResourcesPagesWithContext(ctx context.Context, input *privatenetworks.ListNetworkResourcesInput, fn func(*privatenetworks.ListNetworkResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworkResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworkResourcesOutput{}
	fnCacher := func(out *privatenetworks.ListNetworkResourcesOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworkResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworkResourcesWithContext(ctx context.Context, input *privatenetworks.ListNetworkResourcesInput, opts ...request.Option) (*privatenetworks.ListNetworkResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworkResourcesOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworkResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworkSites(input *privatenetworks.ListNetworkSitesInput) (*privatenetworks.ListNetworkSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworkSitesOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworkSites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworkSitesPages(input *privatenetworks.ListNetworkSitesInput, fn func(*privatenetworks.ListNetworkSitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworkSitesOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworkSitesOutput{}
	fnCacher := func(out *privatenetworks.ListNetworkSitesOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworkSitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworkSitesPagesWithContext(ctx context.Context, input *privatenetworks.ListNetworkSitesInput, fn func(*privatenetworks.ListNetworkSitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworkSitesOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworkSitesOutput{}
	fnCacher := func(out *privatenetworks.ListNetworkSitesOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworkSitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworkSitesWithContext(ctx context.Context, input *privatenetworks.ListNetworkSitesInput, opts ...request.Option) (*privatenetworks.ListNetworkSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworkSitesOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworkSitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworks(input *privatenetworks.ListNetworksInput) (*privatenetworks.ListNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworksOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListNetworksPages(input *privatenetworks.ListNetworksInput, fn func(*privatenetworks.ListNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworksOutput{}
	fnCacher := func(out *privatenetworks.ListNetworksOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworksPagesWithContext(ctx context.Context, input *privatenetworks.ListNetworksInput, fn func(*privatenetworks.ListNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListNetworksOutput{}
	fnCacher := func(out *privatenetworks.ListNetworksOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListNetworksWithContext(ctx context.Context, input *privatenetworks.ListNetworksInput, opts ...request.Option) (*privatenetworks.ListNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListNetworksOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListOrders(input *privatenetworks.ListOrdersInput) (*privatenetworks.ListOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListOrdersOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListOrders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListOrdersPages(input *privatenetworks.ListOrdersInput, fn func(*privatenetworks.ListOrdersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListOrdersOutput{}
	fnCacher := func(out *privatenetworks.ListOrdersOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListOrdersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListOrdersPagesWithContext(ctx context.Context, input *privatenetworks.ListOrdersInput, fn func(*privatenetworks.ListOrdersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*privatenetworks.ListOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &privatenetworks.ListOrdersOutput{}
	fnCacher := func(out *privatenetworks.ListOrdersOutput, more bool) bool {
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
	if err := c.PrivateNetworksAPI.ListOrdersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PrivateNetworks) ListOrdersWithContext(ctx context.Context, input *privatenetworks.ListOrdersInput, opts ...request.Option) (*privatenetworks.ListOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListOrdersOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListOrdersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListTagsForResource(input *privatenetworks.ListTagsForResourceInput) (*privatenetworks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListTagsForResourceOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PrivateNetworks) ListTagsForResourceWithContext(ctx context.Context, input *privatenetworks.ListTagsForResourceInput, opts ...request.Option) (*privatenetworks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*privatenetworks.ListTagsForResourceOutput), nil
	}
	out, err := c.PrivateNetworksAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
