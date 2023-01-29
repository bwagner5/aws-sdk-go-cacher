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
	"github.com/aws/aws-sdk-go/service/chimesdkmediapipelines"
	"github.com/aws/aws-sdk-go/service/chimesdkmediapipelines/chimesdkmediapipelinesiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ChimeSDKMediaPipelines struct {
	chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI
	cache *cache.Cache
}

func NewChimeSDKMediaPipelines(chimesdkmediapipelinesapi chimesdkmediapipelinesiface.ChimeSDKMediaPipelinesAPI) *ChimeSDKMediaPipelines {
	return &ChimeSDKMediaPipelines{
		ChimeSDKMediaPipelinesAPI: chimesdkmediapipelinesapi,
		cache:                     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ChimeSDKMediaPipelines) GetMediaCapturePipeline(input *chimesdkmediapipelines.GetMediaCapturePipelineInput) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.GetMediaCapturePipelineOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.GetMediaCapturePipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) GetMediaCapturePipelineWithContext(ctx context.Context, input *chimesdkmediapipelines.GetMediaCapturePipelineInput, opts ...request.Option) (*chimesdkmediapipelines.GetMediaCapturePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.GetMediaCapturePipelineOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.GetMediaCapturePipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) GetMediaPipeline(input *chimesdkmediapipelines.GetMediaPipelineInput) (*chimesdkmediapipelines.GetMediaPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.GetMediaPipelineOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.GetMediaPipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) GetMediaPipelineWithContext(ctx context.Context, input *chimesdkmediapipelines.GetMediaPipelineInput, opts ...request.Option) (*chimesdkmediapipelines.GetMediaPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.GetMediaPipelineOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.GetMediaPipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListMediaCapturePipelines(input *chimesdkmediapipelines.ListMediaCapturePipelinesInput) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListMediaCapturePipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListMediaCapturePipelinesPages(input *chimesdkmediapipelines.ListMediaCapturePipelinesInput, fn func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmediapipelines.ListMediaCapturePipelinesOutput{}
	fnCacher := func(out *chimesdkmediapipelines.ListMediaCapturePipelinesOutput, more bool) bool {
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
	if err := c.ChimeSDKMediaPipelinesAPI.ListMediaCapturePipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMediaPipelines) ListMediaCapturePipelinesPagesWithContext(ctx context.Context, input *chimesdkmediapipelines.ListMediaCapturePipelinesInput, fn func(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmediapipelines.ListMediaCapturePipelinesOutput{}
	fnCacher := func(out *chimesdkmediapipelines.ListMediaCapturePipelinesOutput, more bool) bool {
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
	if err := c.ChimeSDKMediaPipelinesAPI.ListMediaCapturePipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMediaPipelines) ListMediaCapturePipelinesWithContext(ctx context.Context, input *chimesdkmediapipelines.ListMediaCapturePipelinesInput, opts ...request.Option) (*chimesdkmediapipelines.ListMediaCapturePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListMediaCapturePipelinesOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListMediaCapturePipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListMediaPipelines(input *chimesdkmediapipelines.ListMediaPipelinesInput) (*chimesdkmediapipelines.ListMediaPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListMediaPipelinesOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListMediaPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListMediaPipelinesPages(input *chimesdkmediapipelines.ListMediaPipelinesInput, fn func(*chimesdkmediapipelines.ListMediaPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmediapipelines.ListMediaPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmediapipelines.ListMediaPipelinesOutput{}
	fnCacher := func(out *chimesdkmediapipelines.ListMediaPipelinesOutput, more bool) bool {
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
	if err := c.ChimeSDKMediaPipelinesAPI.ListMediaPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMediaPipelines) ListMediaPipelinesPagesWithContext(ctx context.Context, input *chimesdkmediapipelines.ListMediaPipelinesInput, fn func(*chimesdkmediapipelines.ListMediaPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmediapipelines.ListMediaPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmediapipelines.ListMediaPipelinesOutput{}
	fnCacher := func(out *chimesdkmediapipelines.ListMediaPipelinesOutput, more bool) bool {
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
	if err := c.ChimeSDKMediaPipelinesAPI.ListMediaPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMediaPipelines) ListMediaPipelinesWithContext(ctx context.Context, input *chimesdkmediapipelines.ListMediaPipelinesInput, opts ...request.Option) (*chimesdkmediapipelines.ListMediaPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListMediaPipelinesOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListMediaPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListTagsForResource(input *chimesdkmediapipelines.ListTagsForResourceInput) (*chimesdkmediapipelines.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMediaPipelines) ListTagsForResourceWithContext(ctx context.Context, input *chimesdkmediapipelines.ListTagsForResourceInput, opts ...request.Option) (*chimesdkmediapipelines.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmediapipelines.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMediaPipelinesAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
