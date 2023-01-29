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
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"github.com/aws/aws-sdk-go/service/elastictranscoder/elastictranscoderiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ElasticTranscoder struct {
	elastictranscoderiface.ElasticTranscoderAPI
	cache *cache.Cache
}

func NewElasticTranscoder(elastictranscoderapi elastictranscoderiface.ElasticTranscoderAPI) *ElasticTranscoder {
	return &ElasticTranscoder{
		ElasticTranscoderAPI: elastictranscoderapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ElasticTranscoder) ListJobsByPipeline(input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListJobsByPipelineOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListJobsByPipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListJobsByPipelinePages(input *elastictranscoder.ListJobsByPipelineInput, fn func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListJobsByPipelineOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListJobsByPipelineOutput{}
	fnCacher := func(out *elastictranscoder.ListJobsByPipelineOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListJobsByPipelinePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListJobsByPipelinePagesWithContext(ctx context.Context, input *elastictranscoder.ListJobsByPipelineInput, fn func(*elastictranscoder.ListJobsByPipelineOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListJobsByPipelineOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListJobsByPipelineOutput{}
	fnCacher := func(out *elastictranscoder.ListJobsByPipelineOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListJobsByPipelinePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListJobsByPipelineWithContext(ctx context.Context, input *elastictranscoder.ListJobsByPipelineInput, opts ...request.Option) (*elastictranscoder.ListJobsByPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListJobsByPipelineOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListJobsByPipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListJobsByStatus(input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListJobsByStatusOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListJobsByStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListJobsByStatusPages(input *elastictranscoder.ListJobsByStatusInput, fn func(*elastictranscoder.ListJobsByStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListJobsByStatusOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListJobsByStatusOutput{}
	fnCacher := func(out *elastictranscoder.ListJobsByStatusOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListJobsByStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListJobsByStatusPagesWithContext(ctx context.Context, input *elastictranscoder.ListJobsByStatusInput, fn func(*elastictranscoder.ListJobsByStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListJobsByStatusOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListJobsByStatusOutput{}
	fnCacher := func(out *elastictranscoder.ListJobsByStatusOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListJobsByStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListJobsByStatusWithContext(ctx context.Context, input *elastictranscoder.ListJobsByStatusInput, opts ...request.Option) (*elastictranscoder.ListJobsByStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListJobsByStatusOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListJobsByStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListPipelines(input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListPipelinesOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListPipelinesPages(input *elastictranscoder.ListPipelinesInput, fn func(*elastictranscoder.ListPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListPipelinesOutput{}
	fnCacher := func(out *elastictranscoder.ListPipelinesOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListPipelinesPagesWithContext(ctx context.Context, input *elastictranscoder.ListPipelinesInput, fn func(*elastictranscoder.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListPipelinesOutput{}
	fnCacher := func(out *elastictranscoder.ListPipelinesOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListPipelinesWithContext(ctx context.Context, input *elastictranscoder.ListPipelinesInput, opts ...request.Option) (*elastictranscoder.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListPipelinesOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListPresets(input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListPresetsOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListPresets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticTranscoder) ListPresetsPages(input *elastictranscoder.ListPresetsInput, fn func(*elastictranscoder.ListPresetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListPresetsOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListPresetsOutput{}
	fnCacher := func(out *elastictranscoder.ListPresetsOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListPresetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListPresetsPagesWithContext(ctx context.Context, input *elastictranscoder.ListPresetsInput, fn func(*elastictranscoder.ListPresetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elastictranscoder.ListPresetsOutput), false)
		return nil
	}
	cachable := true
	output := &elastictranscoder.ListPresetsOutput{}
	fnCacher := func(out *elastictranscoder.ListPresetsOutput, more bool) bool {
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
	if err := c.ElasticTranscoderAPI.ListPresetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticTranscoder) ListPresetsWithContext(ctx context.Context, input *elastictranscoder.ListPresetsInput, opts ...request.Option) (*elastictranscoder.ListPresetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elastictranscoder.ListPresetsOutput), nil
	}
	out, err := c.ElasticTranscoderAPI.ListPresetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
