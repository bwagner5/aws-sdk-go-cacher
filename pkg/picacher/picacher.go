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
package picacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pi/piiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PI struct {
	piiface.PIAPI
	cache *cache.Cache
}

func New(piapi piiface.PIAPI) *PI {
	return &PI{
		PIAPI: piapi,
		cache: cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PI) DescribeDimensionKeys(input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.DescribeDimensionKeysOutput), nil
	}
	out, err := c.PIAPI.DescribeDimensionKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) DescribeDimensionKeysPages(input *pi.DescribeDimensionKeysInput, fn func(*pi.DescribeDimensionKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.DescribeDimensionKeysOutput), false)
		return nil
	}
	cachable := true
	output := &pi.DescribeDimensionKeysOutput{}
	fnCacher := func(out *pi.DescribeDimensionKeysOutput, more bool) bool {
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
	if err := c.PIAPI.DescribeDimensionKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) DescribeDimensionKeysPagesWithContext(ctx context.Context, input *pi.DescribeDimensionKeysInput, fn func(*pi.DescribeDimensionKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.DescribeDimensionKeysOutput), false)
		return nil
	}
	cachable := true
	output := &pi.DescribeDimensionKeysOutput{}
	fnCacher := func(out *pi.DescribeDimensionKeysOutput, more bool) bool {
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
	if err := c.PIAPI.DescribeDimensionKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) DescribeDimensionKeysWithContext(ctx context.Context, input *pi.DescribeDimensionKeysInput, opts ...request.Option) (*pi.DescribeDimensionKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.DescribeDimensionKeysOutput), nil
	}
	out, err := c.PIAPI.DescribeDimensionKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetDimensionKeyDetails(input *pi.GetDimensionKeyDetailsInput) (*pi.GetDimensionKeyDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetDimensionKeyDetailsOutput), nil
	}
	out, err := c.PIAPI.GetDimensionKeyDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetDimensionKeyDetailsWithContext(ctx context.Context, input *pi.GetDimensionKeyDetailsInput, opts ...request.Option) (*pi.GetDimensionKeyDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetDimensionKeyDetailsOutput), nil
	}
	out, err := c.PIAPI.GetDimensionKeyDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetResourceMetadata(input *pi.GetResourceMetadataInput) (*pi.GetResourceMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetResourceMetadataOutput), nil
	}
	out, err := c.PIAPI.GetResourceMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetResourceMetadataWithContext(ctx context.Context, input *pi.GetResourceMetadataInput, opts ...request.Option) (*pi.GetResourceMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetResourceMetadataOutput), nil
	}
	out, err := c.PIAPI.GetResourceMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetResourceMetrics(input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetResourceMetricsOutput), nil
	}
	out, err := c.PIAPI.GetResourceMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) GetResourceMetricsPages(input *pi.GetResourceMetricsInput, fn func(*pi.GetResourceMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.GetResourceMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.GetResourceMetricsOutput{}
	fnCacher := func(out *pi.GetResourceMetricsOutput, more bool) bool {
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
	if err := c.PIAPI.GetResourceMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) GetResourceMetricsPagesWithContext(ctx context.Context, input *pi.GetResourceMetricsInput, fn func(*pi.GetResourceMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.GetResourceMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.GetResourceMetricsOutput{}
	fnCacher := func(out *pi.GetResourceMetricsOutput, more bool) bool {
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
	if err := c.PIAPI.GetResourceMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) GetResourceMetricsWithContext(ctx context.Context, input *pi.GetResourceMetricsInput, opts ...request.Option) (*pi.GetResourceMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.GetResourceMetricsOutput), nil
	}
	out, err := c.PIAPI.GetResourceMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) ListAvailableResourceDimensions(input *pi.ListAvailableResourceDimensionsInput) (*pi.ListAvailableResourceDimensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.ListAvailableResourceDimensionsOutput), nil
	}
	out, err := c.PIAPI.ListAvailableResourceDimensions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) ListAvailableResourceDimensionsPages(input *pi.ListAvailableResourceDimensionsInput, fn func(*pi.ListAvailableResourceDimensionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.ListAvailableResourceDimensionsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.ListAvailableResourceDimensionsOutput{}
	fnCacher := func(out *pi.ListAvailableResourceDimensionsOutput, more bool) bool {
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
	if err := c.PIAPI.ListAvailableResourceDimensionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) ListAvailableResourceDimensionsPagesWithContext(ctx context.Context, input *pi.ListAvailableResourceDimensionsInput, fn func(*pi.ListAvailableResourceDimensionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.ListAvailableResourceDimensionsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.ListAvailableResourceDimensionsOutput{}
	fnCacher := func(out *pi.ListAvailableResourceDimensionsOutput, more bool) bool {
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
	if err := c.PIAPI.ListAvailableResourceDimensionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) ListAvailableResourceDimensionsWithContext(ctx context.Context, input *pi.ListAvailableResourceDimensionsInput, opts ...request.Option) (*pi.ListAvailableResourceDimensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.ListAvailableResourceDimensionsOutput), nil
	}
	out, err := c.PIAPI.ListAvailableResourceDimensionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) ListAvailableResourceMetrics(input *pi.ListAvailableResourceMetricsInput) (*pi.ListAvailableResourceMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.ListAvailableResourceMetricsOutput), nil
	}
	out, err := c.PIAPI.ListAvailableResourceMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PI) ListAvailableResourceMetricsPages(input *pi.ListAvailableResourceMetricsInput, fn func(*pi.ListAvailableResourceMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.ListAvailableResourceMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.ListAvailableResourceMetricsOutput{}
	fnCacher := func(out *pi.ListAvailableResourceMetricsOutput, more bool) bool {
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
	if err := c.PIAPI.ListAvailableResourceMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) ListAvailableResourceMetricsPagesWithContext(ctx context.Context, input *pi.ListAvailableResourceMetricsInput, fn func(*pi.ListAvailableResourceMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pi.ListAvailableResourceMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &pi.ListAvailableResourceMetricsOutput{}
	fnCacher := func(out *pi.ListAvailableResourceMetricsOutput, more bool) bool {
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
	if err := c.PIAPI.ListAvailableResourceMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PI) ListAvailableResourceMetricsWithContext(ctx context.Context, input *pi.ListAvailableResourceMetricsInput, opts ...request.Option) (*pi.ListAvailableResourceMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pi.ListAvailableResourceMetricsOutput), nil
	}
	out, err := c.PIAPI.ListAvailableResourceMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
