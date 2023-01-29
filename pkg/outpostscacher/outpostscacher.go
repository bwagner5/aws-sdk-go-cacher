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
package outpostscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/outposts"
	"github.com/aws/aws-sdk-go/service/outposts/outpostsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Outposts struct {
	outpostsiface.OutpostsAPI
	cache *cache.Cache
}

func New(outpostsapi outpostsiface.OutpostsAPI) *Outposts {
	return &Outposts{
		OutpostsAPI: outpostsapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Outposts) GetCatalogItem(input *outposts.GetCatalogItemInput) (*outposts.GetCatalogItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetCatalogItemOutput), nil
	}
	out, err := c.OutpostsAPI.GetCatalogItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetCatalogItemWithContext(ctx context.Context, input *outposts.GetCatalogItemInput, opts ...request.Option) (*outposts.GetCatalogItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetCatalogItemOutput), nil
	}
	out, err := c.OutpostsAPI.GetCatalogItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetConnection(input *outposts.GetConnectionInput) (*outposts.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetConnectionOutput), nil
	}
	out, err := c.OutpostsAPI.GetConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetConnectionWithContext(ctx context.Context, input *outposts.GetConnectionInput, opts ...request.Option) (*outposts.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetConnectionOutput), nil
	}
	out, err := c.OutpostsAPI.GetConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOrder(input *outposts.GetOrderInput) (*outposts.GetOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOrderOutput), nil
	}
	out, err := c.OutpostsAPI.GetOrder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOrderWithContext(ctx context.Context, input *outposts.GetOrderInput, opts ...request.Option) (*outposts.GetOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOrderOutput), nil
	}
	out, err := c.OutpostsAPI.GetOrderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOutpost(input *outposts.GetOutpostInput) (*outposts.GetOutpostOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOutpostOutput), nil
	}
	out, err := c.OutpostsAPI.GetOutpost(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOutpostInstanceTypes(input *outposts.GetOutpostInstanceTypesInput) (*outposts.GetOutpostInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOutpostInstanceTypesOutput), nil
	}
	out, err := c.OutpostsAPI.GetOutpostInstanceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOutpostInstanceTypesPages(input *outposts.GetOutpostInstanceTypesInput, fn func(*outposts.GetOutpostInstanceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.GetOutpostInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.GetOutpostInstanceTypesOutput{}
	fnCacher := func(out *outposts.GetOutpostInstanceTypesOutput, more bool) bool {
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
	if err := c.OutpostsAPI.GetOutpostInstanceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) GetOutpostInstanceTypesPagesWithContext(ctx context.Context, input *outposts.GetOutpostInstanceTypesInput, fn func(*outposts.GetOutpostInstanceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.GetOutpostInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.GetOutpostInstanceTypesOutput{}
	fnCacher := func(out *outposts.GetOutpostInstanceTypesOutput, more bool) bool {
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
	if err := c.OutpostsAPI.GetOutpostInstanceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) GetOutpostInstanceTypesWithContext(ctx context.Context, input *outposts.GetOutpostInstanceTypesInput, opts ...request.Option) (*outposts.GetOutpostInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOutpostInstanceTypesOutput), nil
	}
	out, err := c.OutpostsAPI.GetOutpostInstanceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetOutpostWithContext(ctx context.Context, input *outposts.GetOutpostInput, opts ...request.Option) (*outposts.GetOutpostOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetOutpostOutput), nil
	}
	out, err := c.OutpostsAPI.GetOutpostWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetSite(input *outposts.GetSiteInput) (*outposts.GetSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetSiteOutput), nil
	}
	out, err := c.OutpostsAPI.GetSite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetSiteAddress(input *outposts.GetSiteAddressInput) (*outposts.GetSiteAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetSiteAddressOutput), nil
	}
	out, err := c.OutpostsAPI.GetSiteAddress(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetSiteAddressWithContext(ctx context.Context, input *outposts.GetSiteAddressInput, opts ...request.Option) (*outposts.GetSiteAddressOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetSiteAddressOutput), nil
	}
	out, err := c.OutpostsAPI.GetSiteAddressWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) GetSiteWithContext(ctx context.Context, input *outposts.GetSiteInput, opts ...request.Option) (*outposts.GetSiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.GetSiteOutput), nil
	}
	out, err := c.OutpostsAPI.GetSiteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListAssets(input *outposts.ListAssetsInput) (*outposts.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListAssetsOutput), nil
	}
	out, err := c.OutpostsAPI.ListAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListAssetsPages(input *outposts.ListAssetsInput, fn func(*outposts.ListAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListAssetsOutput{}
	fnCacher := func(out *outposts.ListAssetsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListAssetsPagesWithContext(ctx context.Context, input *outposts.ListAssetsInput, fn func(*outposts.ListAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListAssetsOutput{}
	fnCacher := func(out *outposts.ListAssetsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListAssetsWithContext(ctx context.Context, input *outposts.ListAssetsInput, opts ...request.Option) (*outposts.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListAssetsOutput), nil
	}
	out, err := c.OutpostsAPI.ListAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListCatalogItems(input *outposts.ListCatalogItemsInput) (*outposts.ListCatalogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListCatalogItemsOutput), nil
	}
	out, err := c.OutpostsAPI.ListCatalogItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListCatalogItemsPages(input *outposts.ListCatalogItemsInput, fn func(*outposts.ListCatalogItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListCatalogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListCatalogItemsOutput{}
	fnCacher := func(out *outposts.ListCatalogItemsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListCatalogItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListCatalogItemsPagesWithContext(ctx context.Context, input *outposts.ListCatalogItemsInput, fn func(*outposts.ListCatalogItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListCatalogItemsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListCatalogItemsOutput{}
	fnCacher := func(out *outposts.ListCatalogItemsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListCatalogItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListCatalogItemsWithContext(ctx context.Context, input *outposts.ListCatalogItemsInput, opts ...request.Option) (*outposts.ListCatalogItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListCatalogItemsOutput), nil
	}
	out, err := c.OutpostsAPI.ListCatalogItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListOrders(input *outposts.ListOrdersInput) (*outposts.ListOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListOrdersOutput), nil
	}
	out, err := c.OutpostsAPI.ListOrders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListOrdersPages(input *outposts.ListOrdersInput, fn func(*outposts.ListOrdersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListOrdersOutput{}
	fnCacher := func(out *outposts.ListOrdersOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListOrdersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListOrdersPagesWithContext(ctx context.Context, input *outposts.ListOrdersInput, fn func(*outposts.ListOrdersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListOrdersOutput{}
	fnCacher := func(out *outposts.ListOrdersOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListOrdersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListOrdersWithContext(ctx context.Context, input *outposts.ListOrdersInput, opts ...request.Option) (*outposts.ListOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListOrdersOutput), nil
	}
	out, err := c.OutpostsAPI.ListOrdersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListOutposts(input *outposts.ListOutpostsInput) (*outposts.ListOutpostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListOutpostsOutput), nil
	}
	out, err := c.OutpostsAPI.ListOutposts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListOutpostsPages(input *outposts.ListOutpostsInput, fn func(*outposts.ListOutpostsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListOutpostsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListOutpostsOutput{}
	fnCacher := func(out *outposts.ListOutpostsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListOutpostsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListOutpostsPagesWithContext(ctx context.Context, input *outposts.ListOutpostsInput, fn func(*outposts.ListOutpostsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListOutpostsOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListOutpostsOutput{}
	fnCacher := func(out *outposts.ListOutpostsOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListOutpostsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListOutpostsWithContext(ctx context.Context, input *outposts.ListOutpostsInput, opts ...request.Option) (*outposts.ListOutpostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListOutpostsOutput), nil
	}
	out, err := c.OutpostsAPI.ListOutpostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListSites(input *outposts.ListSitesInput) (*outposts.ListSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListSitesOutput), nil
	}
	out, err := c.OutpostsAPI.ListSites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListSitesPages(input *outposts.ListSitesInput, fn func(*outposts.ListSitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListSitesOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListSitesOutput{}
	fnCacher := func(out *outposts.ListSitesOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListSitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListSitesPagesWithContext(ctx context.Context, input *outposts.ListSitesInput, fn func(*outposts.ListSitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*outposts.ListSitesOutput), false)
		return nil
	}
	cachable := true
	output := &outposts.ListSitesOutput{}
	fnCacher := func(out *outposts.ListSitesOutput, more bool) bool {
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
	if err := c.OutpostsAPI.ListSitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Outposts) ListSitesWithContext(ctx context.Context, input *outposts.ListSitesInput, opts ...request.Option) (*outposts.ListSitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListSitesOutput), nil
	}
	out, err := c.OutpostsAPI.ListSitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListTagsForResource(input *outposts.ListTagsForResourceInput) (*outposts.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListTagsForResourceOutput), nil
	}
	out, err := c.OutpostsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Outposts) ListTagsForResourceWithContext(ctx context.Context, input *outposts.ListTagsForResourceInput, opts ...request.Option) (*outposts.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*outposts.ListTagsForResourceOutput), nil
	}
	out, err := c.OutpostsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
