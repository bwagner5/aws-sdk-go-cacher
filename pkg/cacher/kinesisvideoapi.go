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
	"github.com/aws/aws-sdk-go/service/kinesisvideo"
	"github.com/aws/aws-sdk-go/service/kinesisvideo/kinesisvideoiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KinesisVideo struct {
	kinesisvideoiface.KinesisVideoAPI
	cache *cache.Cache
}

func NewKinesisVideo(kinesisvideoapi kinesisvideoiface.KinesisVideoAPI) *KinesisVideo {
	return &KinesisVideo{
		KinesisVideoAPI: kinesisvideoapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KinesisVideo) DescribeEdgeConfiguration(input *kinesisvideo.DescribeEdgeConfigurationInput) (*kinesisvideo.DescribeEdgeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeEdgeConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeEdgeConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeEdgeConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeEdgeConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeEdgeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeEdgeConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeEdgeConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeImageGenerationConfiguration(input *kinesisvideo.DescribeImageGenerationConfigurationInput) (*kinesisvideo.DescribeImageGenerationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeImageGenerationConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeImageGenerationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeImageGenerationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeImageGenerationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeImageGenerationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeImageGenerationConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeImageGenerationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeMappedResourceConfiguration(input *kinesisvideo.DescribeMappedResourceConfigurationInput) (*kinesisvideo.DescribeMappedResourceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeMappedResourceConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeMappedResourceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeMappedResourceConfigurationPages(input *kinesisvideo.DescribeMappedResourceConfigurationInput, fn func(*kinesisvideo.DescribeMappedResourceConfigurationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.DescribeMappedResourceConfigurationOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.DescribeMappedResourceConfigurationOutput{}
	fnCacher := func(out *kinesisvideo.DescribeMappedResourceConfigurationOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.DescribeMappedResourceConfigurationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) DescribeMappedResourceConfigurationPagesWithContext(ctx context.Context, input *kinesisvideo.DescribeMappedResourceConfigurationInput, fn func(*kinesisvideo.DescribeMappedResourceConfigurationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.DescribeMappedResourceConfigurationOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.DescribeMappedResourceConfigurationOutput{}
	fnCacher := func(out *kinesisvideo.DescribeMappedResourceConfigurationOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.DescribeMappedResourceConfigurationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) DescribeMappedResourceConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeMappedResourceConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeMappedResourceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeMappedResourceConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeMappedResourceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeMediaStorageConfiguration(input *kinesisvideo.DescribeMediaStorageConfigurationInput) (*kinesisvideo.DescribeMediaStorageConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeMediaStorageConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeMediaStorageConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeMediaStorageConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeMediaStorageConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeMediaStorageConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeMediaStorageConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeMediaStorageConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeNotificationConfiguration(input *kinesisvideo.DescribeNotificationConfigurationInput) (*kinesisvideo.DescribeNotificationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeNotificationConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeNotificationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeNotificationConfigurationWithContext(ctx context.Context, input *kinesisvideo.DescribeNotificationConfigurationInput, opts ...request.Option) (*kinesisvideo.DescribeNotificationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeNotificationConfigurationOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeNotificationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeSignalingChannel(input *kinesisvideo.DescribeSignalingChannelInput) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeSignalingChannelOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeSignalingChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeSignalingChannelWithContext(ctx context.Context, input *kinesisvideo.DescribeSignalingChannelInput, opts ...request.Option) (*kinesisvideo.DescribeSignalingChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeSignalingChannelOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeSignalingChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeStream(input *kinesisvideo.DescribeStreamInput) (*kinesisvideo.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeStreamOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) DescribeStreamWithContext(ctx context.Context, input *kinesisvideo.DescribeStreamInput, opts ...request.Option) (*kinesisvideo.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.DescribeStreamOutput), nil
	}
	out, err := c.KinesisVideoAPI.DescribeStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) GetDataEndpoint(input *kinesisvideo.GetDataEndpointInput) (*kinesisvideo.GetDataEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.GetDataEndpointOutput), nil
	}
	out, err := c.KinesisVideoAPI.GetDataEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) GetDataEndpointWithContext(ctx context.Context, input *kinesisvideo.GetDataEndpointInput, opts ...request.Option) (*kinesisvideo.GetDataEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.GetDataEndpointOutput), nil
	}
	out, err := c.KinesisVideoAPI.GetDataEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) GetSignalingChannelEndpoint(input *kinesisvideo.GetSignalingChannelEndpointInput) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.GetSignalingChannelEndpointOutput), nil
	}
	out, err := c.KinesisVideoAPI.GetSignalingChannelEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) GetSignalingChannelEndpointWithContext(ctx context.Context, input *kinesisvideo.GetSignalingChannelEndpointInput, opts ...request.Option) (*kinesisvideo.GetSignalingChannelEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.GetSignalingChannelEndpointOutput), nil
	}
	out, err := c.KinesisVideoAPI.GetSignalingChannelEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListSignalingChannels(input *kinesisvideo.ListSignalingChannelsInput) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListSignalingChannelsOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListSignalingChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListSignalingChannelsPages(input *kinesisvideo.ListSignalingChannelsInput, fn func(*kinesisvideo.ListSignalingChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.ListSignalingChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.ListSignalingChannelsOutput{}
	fnCacher := func(out *kinesisvideo.ListSignalingChannelsOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.ListSignalingChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) ListSignalingChannelsPagesWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, fn func(*kinesisvideo.ListSignalingChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.ListSignalingChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.ListSignalingChannelsOutput{}
	fnCacher := func(out *kinesisvideo.ListSignalingChannelsOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.ListSignalingChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) ListSignalingChannelsWithContext(ctx context.Context, input *kinesisvideo.ListSignalingChannelsInput, opts ...request.Option) (*kinesisvideo.ListSignalingChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListSignalingChannelsOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListSignalingChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListStreams(input *kinesisvideo.ListStreamsInput) (*kinesisvideo.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListStreamsOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListStreamsPages(input *kinesisvideo.ListStreamsInput, fn func(*kinesisvideo.ListStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.ListStreamsOutput{}
	fnCacher := func(out *kinesisvideo.ListStreamsOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.ListStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) ListStreamsPagesWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, fn func(*kinesisvideo.ListStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesisvideo.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesisvideo.ListStreamsOutput{}
	fnCacher := func(out *kinesisvideo.ListStreamsOutput, more bool) bool {
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
	if err := c.KinesisVideoAPI.ListStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KinesisVideo) ListStreamsWithContext(ctx context.Context, input *kinesisvideo.ListStreamsInput, opts ...request.Option) (*kinesisvideo.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListStreamsOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListTagsForResource(input *kinesisvideo.ListTagsForResourceInput) (*kinesisvideo.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListTagsForResourceWithContext(ctx context.Context, input *kinesisvideo.ListTagsForResourceInput, opts ...request.Option) (*kinesisvideo.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListTagsForResourceOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListTagsForStream(input *kinesisvideo.ListTagsForStreamInput) (*kinesisvideo.ListTagsForStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListTagsForStreamOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListTagsForStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KinesisVideo) ListTagsForStreamWithContext(ctx context.Context, input *kinesisvideo.ListTagsForStreamInput, opts ...request.Option) (*kinesisvideo.ListTagsForStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesisvideo.ListTagsForStreamOutput), nil
	}
	out, err := c.KinesisVideoAPI.ListTagsForStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
