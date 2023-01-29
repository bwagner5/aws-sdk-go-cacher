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
package mediapackagecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"github.com/aws/aws-sdk-go/service/mediapackage/mediapackageiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaPackage struct {
	mediapackageiface.MediaPackageAPI
	cache *cache.Cache
}

func New(mediapackageapi mediapackageiface.MediaPackageAPI) *MediaPackage {
	return &MediaPackage{
		MediaPackageAPI: mediapackageapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaPackage) DescribeChannel(input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeChannelOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) DescribeChannelWithContext(ctx context.Context, input *mediapackage.DescribeChannelInput, opts ...request.Option) (*mediapackage.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeChannelOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) DescribeHarvestJob(input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeHarvestJobOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeHarvestJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) DescribeHarvestJobWithContext(ctx context.Context, input *mediapackage.DescribeHarvestJobInput, opts ...request.Option) (*mediapackage.DescribeHarvestJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeHarvestJobOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeHarvestJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) DescribeOriginEndpoint(input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeOriginEndpointOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeOriginEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) DescribeOriginEndpointWithContext(ctx context.Context, input *mediapackage.DescribeOriginEndpointInput, opts ...request.Option) (*mediapackage.DescribeOriginEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.DescribeOriginEndpointOutput), nil
	}
	out, err := c.MediaPackageAPI.DescribeOriginEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListChannels(input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListChannelsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListChannelsPages(input *mediapackage.ListChannelsInput, fn func(*mediapackage.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListChannelsOutput{}
	fnCacher := func(out *mediapackage.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListChannelsPagesWithContext(ctx context.Context, input *mediapackage.ListChannelsInput, fn func(*mediapackage.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListChannelsOutput{}
	fnCacher := func(out *mediapackage.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListChannelsWithContext(ctx context.Context, input *mediapackage.ListChannelsInput, opts ...request.Option) (*mediapackage.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListChannelsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListHarvestJobs(input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListHarvestJobsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListHarvestJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListHarvestJobsPages(input *mediapackage.ListHarvestJobsInput, fn func(*mediapackage.ListHarvestJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListHarvestJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListHarvestJobsOutput{}
	fnCacher := func(out *mediapackage.ListHarvestJobsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListHarvestJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListHarvestJobsPagesWithContext(ctx context.Context, input *mediapackage.ListHarvestJobsInput, fn func(*mediapackage.ListHarvestJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListHarvestJobsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListHarvestJobsOutput{}
	fnCacher := func(out *mediapackage.ListHarvestJobsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListHarvestJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListHarvestJobsWithContext(ctx context.Context, input *mediapackage.ListHarvestJobsInput, opts ...request.Option) (*mediapackage.ListHarvestJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListHarvestJobsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListHarvestJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListOriginEndpoints(input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListOriginEndpointsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListOriginEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListOriginEndpointsPages(input *mediapackage.ListOriginEndpointsInput, fn func(*mediapackage.ListOriginEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListOriginEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListOriginEndpointsOutput{}
	fnCacher := func(out *mediapackage.ListOriginEndpointsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListOriginEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListOriginEndpointsPagesWithContext(ctx context.Context, input *mediapackage.ListOriginEndpointsInput, fn func(*mediapackage.ListOriginEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackage.ListOriginEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackage.ListOriginEndpointsOutput{}
	fnCacher := func(out *mediapackage.ListOriginEndpointsOutput, more bool) bool {
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
	if err := c.MediaPackageAPI.ListOriginEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackage) ListOriginEndpointsWithContext(ctx context.Context, input *mediapackage.ListOriginEndpointsInput, opts ...request.Option) (*mediapackage.ListOriginEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListOriginEndpointsOutput), nil
	}
	out, err := c.MediaPackageAPI.ListOriginEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListTagsForResource(input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaPackageAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackage) ListTagsForResourceWithContext(ctx context.Context, input *mediapackage.ListTagsForResourceInput, opts ...request.Option) (*mediapackage.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackage.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaPackageAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
