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
package mediaconvertcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/mediaconvert/mediaconvertiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaConvert struct {
	mediaconvertiface.MediaConvertAPI
	cache *cache.Cache
}

func New(mediaconvertapi mediaconvertiface.MediaConvertAPI) *MediaConvert {
	return &MediaConvert{
		MediaConvertAPI: mediaconvertapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaConvert) DescribeEndpoints(input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.DescribeEndpointsOutput), nil
	}
	out, err := c.MediaConvertAPI.DescribeEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) DescribeEndpointsPages(input *mediaconvert.DescribeEndpointsInput, fn func(*mediaconvert.DescribeEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.DescribeEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.DescribeEndpointsOutput{}
	fnCacher := func(out *mediaconvert.DescribeEndpointsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.DescribeEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) DescribeEndpointsPagesWithContext(ctx context.Context, input *mediaconvert.DescribeEndpointsInput, fn func(*mediaconvert.DescribeEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.DescribeEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.DescribeEndpointsOutput{}
	fnCacher := func(out *mediaconvert.DescribeEndpointsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.DescribeEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) DescribeEndpointsWithContext(ctx context.Context, input *mediaconvert.DescribeEndpointsInput, opts ...request.Option) (*mediaconvert.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.DescribeEndpointsOutput), nil
	}
	out, err := c.MediaConvertAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetJob(input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetJobOutput), nil
	}
	out, err := c.MediaConvertAPI.GetJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetJobTemplate(input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetJobTemplateOutput), nil
	}
	out, err := c.MediaConvertAPI.GetJobTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetJobTemplateWithContext(ctx context.Context, input *mediaconvert.GetJobTemplateInput, opts ...request.Option) (*mediaconvert.GetJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetJobTemplateOutput), nil
	}
	out, err := c.MediaConvertAPI.GetJobTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetJobWithContext(ctx context.Context, input *mediaconvert.GetJobInput, opts ...request.Option) (*mediaconvert.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetJobOutput), nil
	}
	out, err := c.MediaConvertAPI.GetJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetPolicy(input *mediaconvert.GetPolicyInput) (*mediaconvert.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetPolicyOutput), nil
	}
	out, err := c.MediaConvertAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetPolicyWithContext(ctx context.Context, input *mediaconvert.GetPolicyInput, opts ...request.Option) (*mediaconvert.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetPolicyOutput), nil
	}
	out, err := c.MediaConvertAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetPreset(input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetPresetOutput), nil
	}
	out, err := c.MediaConvertAPI.GetPreset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetPresetWithContext(ctx context.Context, input *mediaconvert.GetPresetInput, opts ...request.Option) (*mediaconvert.GetPresetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetPresetOutput), nil
	}
	out, err := c.MediaConvertAPI.GetPresetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetQueue(input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetQueueOutput), nil
	}
	out, err := c.MediaConvertAPI.GetQueue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) GetQueueWithContext(ctx context.Context, input *mediaconvert.GetQueueInput, opts ...request.Option) (*mediaconvert.GetQueueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.GetQueueOutput), nil
	}
	out, err := c.MediaConvertAPI.GetQueueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListJobTemplates(input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListJobTemplatesOutput), nil
	}
	out, err := c.MediaConvertAPI.ListJobTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListJobTemplatesPages(input *mediaconvert.ListJobTemplatesInput, fn func(*mediaconvert.ListJobTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListJobTemplatesOutput{}
	fnCacher := func(out *mediaconvert.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListJobTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListJobTemplatesPagesWithContext(ctx context.Context, input *mediaconvert.ListJobTemplatesInput, fn func(*mediaconvert.ListJobTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListJobTemplatesOutput{}
	fnCacher := func(out *mediaconvert.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListJobTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListJobTemplatesWithContext(ctx context.Context, input *mediaconvert.ListJobTemplatesInput, opts ...request.Option) (*mediaconvert.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListJobTemplatesOutput), nil
	}
	out, err := c.MediaConvertAPI.ListJobTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListJobs(input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListJobsOutput), nil
	}
	out, err := c.MediaConvertAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListJobsPages(input *mediaconvert.ListJobsInput, fn func(*mediaconvert.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListJobsOutput{}
	fnCacher := func(out *mediaconvert.ListJobsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListJobsPagesWithContext(ctx context.Context, input *mediaconvert.ListJobsInput, fn func(*mediaconvert.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListJobsOutput{}
	fnCacher := func(out *mediaconvert.ListJobsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListJobsWithContext(ctx context.Context, input *mediaconvert.ListJobsInput, opts ...request.Option) (*mediaconvert.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListJobsOutput), nil
	}
	out, err := c.MediaConvertAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListPresets(input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListPresetsOutput), nil
	}
	out, err := c.MediaConvertAPI.ListPresets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListPresetsPages(input *mediaconvert.ListPresetsInput, fn func(*mediaconvert.ListPresetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListPresetsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListPresetsOutput{}
	fnCacher := func(out *mediaconvert.ListPresetsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListPresetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListPresetsPagesWithContext(ctx context.Context, input *mediaconvert.ListPresetsInput, fn func(*mediaconvert.ListPresetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListPresetsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListPresetsOutput{}
	fnCacher := func(out *mediaconvert.ListPresetsOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListPresetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListPresetsWithContext(ctx context.Context, input *mediaconvert.ListPresetsInput, opts ...request.Option) (*mediaconvert.ListPresetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListPresetsOutput), nil
	}
	out, err := c.MediaConvertAPI.ListPresetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListQueues(input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListQueuesOutput), nil
	}
	out, err := c.MediaConvertAPI.ListQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListQueuesPages(input *mediaconvert.ListQueuesInput, fn func(*mediaconvert.ListQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListQueuesOutput{}
	fnCacher := func(out *mediaconvert.ListQueuesOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListQueuesPagesWithContext(ctx context.Context, input *mediaconvert.ListQueuesInput, fn func(*mediaconvert.ListQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconvert.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconvert.ListQueuesOutput{}
	fnCacher := func(out *mediaconvert.ListQueuesOutput, more bool) bool {
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
	if err := c.MediaConvertAPI.ListQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConvert) ListQueuesWithContext(ctx context.Context, input *mediaconvert.ListQueuesInput, opts ...request.Option) (*mediaconvert.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListQueuesOutput), nil
	}
	out, err := c.MediaConvertAPI.ListQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListTagsForResource(input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaConvertAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConvert) ListTagsForResourceWithContext(ctx context.Context, input *mediaconvert.ListTagsForResourceInput, opts ...request.Option) (*mediaconvert.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconvert.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaConvertAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
