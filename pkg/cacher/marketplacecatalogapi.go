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
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"github.com/aws/aws-sdk-go/service/marketplacecatalog/marketplacecatalogiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MarketplaceCatalog struct {
	marketplacecatalogiface.MarketplaceCatalogAPI
	cache *cache.Cache
}

func NewMarketplaceCatalog(marketplacecatalogapi marketplacecatalogiface.MarketplaceCatalogAPI) *MarketplaceCatalog {
	return &MarketplaceCatalog{
		MarketplaceCatalogAPI: marketplacecatalogapi,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MarketplaceCatalog) DescribeChangeSet(input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.DescribeChangeSetOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.DescribeChangeSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) DescribeChangeSetWithContext(ctx context.Context, input *marketplacecatalog.DescribeChangeSetInput, opts ...request.Option) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.DescribeChangeSetOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.DescribeChangeSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) DescribeEntity(input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.DescribeEntityOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.DescribeEntity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) DescribeEntityWithContext(ctx context.Context, input *marketplacecatalog.DescribeEntityInput, opts ...request.Option) (*marketplacecatalog.DescribeEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.DescribeEntityOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.DescribeEntityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListChangeSets(input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListChangeSetsOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListChangeSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListChangeSetsPages(input *marketplacecatalog.ListChangeSetsInput, fn func(*marketplacecatalog.ListChangeSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*marketplacecatalog.ListChangeSetsOutput), false)
		return nil
	}
	cachable := true
	output := &marketplacecatalog.ListChangeSetsOutput{}
	fnCacher := func(out *marketplacecatalog.ListChangeSetsOutput, more bool) bool {
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
	if err := c.MarketplaceCatalogAPI.ListChangeSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MarketplaceCatalog) ListChangeSetsPagesWithContext(ctx context.Context, input *marketplacecatalog.ListChangeSetsInput, fn func(*marketplacecatalog.ListChangeSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*marketplacecatalog.ListChangeSetsOutput), false)
		return nil
	}
	cachable := true
	output := &marketplacecatalog.ListChangeSetsOutput{}
	fnCacher := func(out *marketplacecatalog.ListChangeSetsOutput, more bool) bool {
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
	if err := c.MarketplaceCatalogAPI.ListChangeSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MarketplaceCatalog) ListChangeSetsWithContext(ctx context.Context, input *marketplacecatalog.ListChangeSetsInput, opts ...request.Option) (*marketplacecatalog.ListChangeSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListChangeSetsOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListChangeSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListEntities(input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListEntitiesOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListEntitiesPages(input *marketplacecatalog.ListEntitiesInput, fn func(*marketplacecatalog.ListEntitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*marketplacecatalog.ListEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &marketplacecatalog.ListEntitiesOutput{}
	fnCacher := func(out *marketplacecatalog.ListEntitiesOutput, more bool) bool {
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
	if err := c.MarketplaceCatalogAPI.ListEntitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MarketplaceCatalog) ListEntitiesPagesWithContext(ctx context.Context, input *marketplacecatalog.ListEntitiesInput, fn func(*marketplacecatalog.ListEntitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*marketplacecatalog.ListEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &marketplacecatalog.ListEntitiesOutput{}
	fnCacher := func(out *marketplacecatalog.ListEntitiesOutput, more bool) bool {
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
	if err := c.MarketplaceCatalogAPI.ListEntitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MarketplaceCatalog) ListEntitiesWithContext(ctx context.Context, input *marketplacecatalog.ListEntitiesInput, opts ...request.Option) (*marketplacecatalog.ListEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListEntitiesOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListTagsForResource(input *marketplacecatalog.ListTagsForResourceInput) (*marketplacecatalog.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListTagsForResourceOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MarketplaceCatalog) ListTagsForResourceWithContext(ctx context.Context, input *marketplacecatalog.ListTagsForResourceInput, opts ...request.Option) (*marketplacecatalog.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*marketplacecatalog.ListTagsForResourceOutput), nil
	}
	out, err := c.MarketplaceCatalogAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
