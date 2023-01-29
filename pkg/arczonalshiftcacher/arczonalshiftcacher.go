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
package arczonalshiftcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/arczonalshift"
	"github.com/aws/aws-sdk-go/service/arczonalshift/arczonalshiftiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ARCZonalShift struct {
	arczonalshiftiface.ARCZonalShiftAPI
	cache *cache.Cache
}

func New(arczonalshiftapi arczonalshiftiface.ARCZonalShiftAPI) *ARCZonalShift {
	return &ARCZonalShift{
		ARCZonalShiftAPI: arczonalshiftapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ARCZonalShift) GetManagedResource(input *arczonalshift.GetManagedResourceInput) (*arczonalshift.GetManagedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.GetManagedResourceOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.GetManagedResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ARCZonalShift) GetManagedResourceWithContext(ctx context.Context, input *arczonalshift.GetManagedResourceInput, opts ...request.Option) (*arczonalshift.GetManagedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.GetManagedResourceOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.GetManagedResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ARCZonalShift) ListManagedResources(input *arczonalshift.ListManagedResourcesInput) (*arczonalshift.ListManagedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.ListManagedResourcesOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.ListManagedResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ARCZonalShift) ListManagedResourcesPages(input *arczonalshift.ListManagedResourcesInput, fn func(*arczonalshift.ListManagedResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*arczonalshift.ListManagedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &arczonalshift.ListManagedResourcesOutput{}
	fnCacher := func(out *arczonalshift.ListManagedResourcesOutput, more bool) bool {
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
	if err := c.ARCZonalShiftAPI.ListManagedResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ARCZonalShift) ListManagedResourcesPagesWithContext(ctx context.Context, input *arczonalshift.ListManagedResourcesInput, fn func(*arczonalshift.ListManagedResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*arczonalshift.ListManagedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &arczonalshift.ListManagedResourcesOutput{}
	fnCacher := func(out *arczonalshift.ListManagedResourcesOutput, more bool) bool {
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
	if err := c.ARCZonalShiftAPI.ListManagedResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ARCZonalShift) ListManagedResourcesWithContext(ctx context.Context, input *arczonalshift.ListManagedResourcesInput, opts ...request.Option) (*arczonalshift.ListManagedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.ListManagedResourcesOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.ListManagedResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ARCZonalShift) ListZonalShifts(input *arczonalshift.ListZonalShiftsInput) (*arczonalshift.ListZonalShiftsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.ListZonalShiftsOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.ListZonalShifts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ARCZonalShift) ListZonalShiftsPages(input *arczonalshift.ListZonalShiftsInput, fn func(*arczonalshift.ListZonalShiftsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*arczonalshift.ListZonalShiftsOutput), false)
		return nil
	}
	cachable := true
	output := &arczonalshift.ListZonalShiftsOutput{}
	fnCacher := func(out *arczonalshift.ListZonalShiftsOutput, more bool) bool {
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
	if err := c.ARCZonalShiftAPI.ListZonalShiftsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ARCZonalShift) ListZonalShiftsPagesWithContext(ctx context.Context, input *arczonalshift.ListZonalShiftsInput, fn func(*arczonalshift.ListZonalShiftsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*arczonalshift.ListZonalShiftsOutput), false)
		return nil
	}
	cachable := true
	output := &arczonalshift.ListZonalShiftsOutput{}
	fnCacher := func(out *arczonalshift.ListZonalShiftsOutput, more bool) bool {
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
	if err := c.ARCZonalShiftAPI.ListZonalShiftsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ARCZonalShift) ListZonalShiftsWithContext(ctx context.Context, input *arczonalshift.ListZonalShiftsInput, opts ...request.Option) (*arczonalshift.ListZonalShiftsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*arczonalshift.ListZonalShiftsOutput), nil
	}
	out, err := c.ARCZonalShiftAPI.ListZonalShiftsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
