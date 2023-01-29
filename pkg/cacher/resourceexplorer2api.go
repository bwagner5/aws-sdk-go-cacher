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
	"github.com/aws/aws-sdk-go/service/resourceexplorer2"
	"github.com/aws/aws-sdk-go/service/resourceexplorer2/resourceexplorer2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ResourceExplorer2 struct {
	resourceexplorer2iface.ResourceExplorer2API
	cache *cache.Cache
}

func NewResourceExplorer2(resourceexplorer2api resourceexplorer2iface.ResourceExplorer2API) *ResourceExplorer2 {
	return &ResourceExplorer2{
		ResourceExplorer2API: resourceexplorer2api,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ResourceExplorer2) BatchGetView(input *resourceexplorer2.BatchGetViewInput) (*resourceexplorer2.BatchGetViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.BatchGetViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.BatchGetView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) BatchGetViewWithContext(ctx context.Context, input *resourceexplorer2.BatchGetViewInput, opts ...request.Option) (*resourceexplorer2.BatchGetViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.BatchGetViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.BatchGetViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetDefaultView(input *resourceexplorer2.GetDefaultViewInput) (*resourceexplorer2.GetDefaultViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetDefaultViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetDefaultView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetDefaultViewWithContext(ctx context.Context, input *resourceexplorer2.GetDefaultViewInput, opts ...request.Option) (*resourceexplorer2.GetDefaultViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetDefaultViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetDefaultViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetIndex(input *resourceexplorer2.GetIndexInput) (*resourceexplorer2.GetIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetIndexOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetIndexWithContext(ctx context.Context, input *resourceexplorer2.GetIndexInput, opts ...request.Option) (*resourceexplorer2.GetIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetIndexOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetView(input *resourceexplorer2.GetViewInput) (*resourceexplorer2.GetViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) GetViewWithContext(ctx context.Context, input *resourceexplorer2.GetViewInput, opts ...request.Option) (*resourceexplorer2.GetViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.GetViewOutput), nil
	}
	out, err := c.ResourceExplorer2API.GetViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListIndexes(input *resourceexplorer2.ListIndexesInput) (*resourceexplorer2.ListIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListIndexesOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListIndexes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListIndexesPages(input *resourceexplorer2.ListIndexesInput, fn func(*resourceexplorer2.ListIndexesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListIndexesOutput{}
	fnCacher := func(out *resourceexplorer2.ListIndexesOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListIndexesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListIndexesPagesWithContext(ctx context.Context, input *resourceexplorer2.ListIndexesInput, fn func(*resourceexplorer2.ListIndexesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListIndexesOutput{}
	fnCacher := func(out *resourceexplorer2.ListIndexesOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListIndexesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListIndexesWithContext(ctx context.Context, input *resourceexplorer2.ListIndexesInput, opts ...request.Option) (*resourceexplorer2.ListIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListIndexesOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListIndexesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListSupportedResourceTypes(input *resourceexplorer2.ListSupportedResourceTypesInput) (*resourceexplorer2.ListSupportedResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListSupportedResourceTypesOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListSupportedResourceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListSupportedResourceTypesPages(input *resourceexplorer2.ListSupportedResourceTypesInput, fn func(*resourceexplorer2.ListSupportedResourceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListSupportedResourceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListSupportedResourceTypesOutput{}
	fnCacher := func(out *resourceexplorer2.ListSupportedResourceTypesOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListSupportedResourceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListSupportedResourceTypesPagesWithContext(ctx context.Context, input *resourceexplorer2.ListSupportedResourceTypesInput, fn func(*resourceexplorer2.ListSupportedResourceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListSupportedResourceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListSupportedResourceTypesOutput{}
	fnCacher := func(out *resourceexplorer2.ListSupportedResourceTypesOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListSupportedResourceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListSupportedResourceTypesWithContext(ctx context.Context, input *resourceexplorer2.ListSupportedResourceTypesInput, opts ...request.Option) (*resourceexplorer2.ListSupportedResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListSupportedResourceTypesOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListSupportedResourceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListTagsForResource(input *resourceexplorer2.ListTagsForResourceInput) (*resourceexplorer2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListTagsForResourceOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListTagsForResourceWithContext(ctx context.Context, input *resourceexplorer2.ListTagsForResourceInput, opts ...request.Option) (*resourceexplorer2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListTagsForResourceOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListViews(input *resourceexplorer2.ListViewsInput) (*resourceexplorer2.ListViewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListViewsOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListViews(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) ListViewsPages(input *resourceexplorer2.ListViewsInput, fn func(*resourceexplorer2.ListViewsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListViewsOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListViewsOutput{}
	fnCacher := func(out *resourceexplorer2.ListViewsOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListViewsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListViewsPagesWithContext(ctx context.Context, input *resourceexplorer2.ListViewsInput, fn func(*resourceexplorer2.ListViewsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.ListViewsOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.ListViewsOutput{}
	fnCacher := func(out *resourceexplorer2.ListViewsOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.ListViewsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) ListViewsWithContext(ctx context.Context, input *resourceexplorer2.ListViewsInput, opts ...request.Option) (*resourceexplorer2.ListViewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.ListViewsOutput), nil
	}
	out, err := c.ResourceExplorer2API.ListViewsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) Search(input *resourceexplorer2.SearchInput) (*resourceexplorer2.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.SearchOutput), nil
	}
	out, err := c.ResourceExplorer2API.Search(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceExplorer2) SearchPages(input *resourceexplorer2.SearchInput, fn func(*resourceexplorer2.SearchOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.SearchOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.SearchOutput{}
	fnCacher := func(out *resourceexplorer2.SearchOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.SearchPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) SearchPagesWithContext(ctx context.Context, input *resourceexplorer2.SearchInput, fn func(*resourceexplorer2.SearchOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourceexplorer2.SearchOutput), false)
		return nil
	}
	cachable := true
	output := &resourceexplorer2.SearchOutput{}
	fnCacher := func(out *resourceexplorer2.SearchOutput, more bool) bool {
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
	if err := c.ResourceExplorer2API.SearchPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceExplorer2) SearchWithContext(ctx context.Context, input *resourceexplorer2.SearchInput, opts ...request.Option) (*resourceexplorer2.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourceexplorer2.SearchOutput), nil
	}
	out, err := c.ResourceExplorer2API.SearchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
