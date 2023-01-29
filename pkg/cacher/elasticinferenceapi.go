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
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"github.com/aws/aws-sdk-go/service/elasticinference/elasticinferenceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ElasticInference struct {
	elasticinferenceiface.ElasticInferenceAPI
	cache *cache.Cache
}

func NewElasticInference(elasticinferenceapi elasticinferenceiface.ElasticInferenceAPI) *ElasticInference {
	return &ElasticInference{
		ElasticInferenceAPI: elasticinferenceapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ElasticInference) DescribeAcceleratorOfferings(input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorOfferingsOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAcceleratorOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) DescribeAcceleratorOfferingsWithContext(ctx context.Context, input *elasticinference.DescribeAcceleratorOfferingsInput, opts ...request.Option) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorOfferingsOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAcceleratorOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) DescribeAcceleratorTypes(input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorTypesOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAcceleratorTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) DescribeAcceleratorTypesWithContext(ctx context.Context, input *elasticinference.DescribeAcceleratorTypesInput, opts ...request.Option) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorTypesOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAcceleratorTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) DescribeAccelerators(input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorsOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAccelerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) DescribeAcceleratorsPages(input *elasticinference.DescribeAcceleratorsInput, fn func(*elasticinference.DescribeAcceleratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticinference.DescribeAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticinference.DescribeAcceleratorsOutput{}
	fnCacher := func(out *elasticinference.DescribeAcceleratorsOutput, more bool) bool {
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
	if err := c.ElasticInferenceAPI.DescribeAcceleratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticInference) DescribeAcceleratorsPagesWithContext(ctx context.Context, input *elasticinference.DescribeAcceleratorsInput, fn func(*elasticinference.DescribeAcceleratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticinference.DescribeAcceleratorsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticinference.DescribeAcceleratorsOutput{}
	fnCacher := func(out *elasticinference.DescribeAcceleratorsOutput, more bool) bool {
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
	if err := c.ElasticInferenceAPI.DescribeAcceleratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticInference) DescribeAcceleratorsWithContext(ctx context.Context, input *elasticinference.DescribeAcceleratorsInput, opts ...request.Option) (*elasticinference.DescribeAcceleratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.DescribeAcceleratorsOutput), nil
	}
	out, err := c.ElasticInferenceAPI.DescribeAcceleratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) ListTagsForResource(input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.ListTagsForResourceOutput), nil
	}
	out, err := c.ElasticInferenceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticInference) ListTagsForResourceWithContext(ctx context.Context, input *elasticinference.ListTagsForResourceInput, opts ...request.Option) (*elasticinference.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticinference.ListTagsForResourceOutput), nil
	}
	out, err := c.ElasticInferenceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
