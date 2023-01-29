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
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediapackagevod/mediapackagevodiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaPackageVod struct {
	mediapackagevodiface.MediaPackageVodAPI
	cache *cache.Cache
}

func NewMediaPackageVod(mediapackagevodapi mediapackagevodiface.MediaPackageVodAPI) *MediaPackageVod {
	return &MediaPackageVod{
		MediaPackageVodAPI: mediapackagevodapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaPackageVod) DescribeAsset(input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribeAssetOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribeAsset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) DescribeAssetWithContext(ctx context.Context, input *mediapackagevod.DescribeAssetInput, opts ...request.Option) (*mediapackagevod.DescribeAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribeAssetOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribeAssetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) DescribePackagingConfiguration(input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribePackagingConfigurationOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribePackagingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) DescribePackagingConfigurationWithContext(ctx context.Context, input *mediapackagevod.DescribePackagingConfigurationInput, opts ...request.Option) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribePackagingConfigurationOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribePackagingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) DescribePackagingGroup(input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribePackagingGroupOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribePackagingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) DescribePackagingGroupWithContext(ctx context.Context, input *mediapackagevod.DescribePackagingGroupInput, opts ...request.Option) (*mediapackagevod.DescribePackagingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.DescribePackagingGroupOutput), nil
	}
	out, err := c.MediaPackageVodAPI.DescribePackagingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListAssets(input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListAssetsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListAssetsPages(input *mediapackagevod.ListAssetsInput, fn func(*mediapackagevod.ListAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListAssetsOutput{}
	fnCacher := func(out *mediapackagevod.ListAssetsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListAssetsPagesWithContext(ctx context.Context, input *mediapackagevod.ListAssetsInput, fn func(*mediapackagevod.ListAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListAssetsOutput{}
	fnCacher := func(out *mediapackagevod.ListAssetsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListAssetsWithContext(ctx context.Context, input *mediapackagevod.ListAssetsInput, opts ...request.Option) (*mediapackagevod.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListAssetsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListPackagingConfigurations(input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListPackagingConfigurationsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListPackagingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListPackagingConfigurationsPages(input *mediapackagevod.ListPackagingConfigurationsInput, fn func(*mediapackagevod.ListPackagingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListPackagingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListPackagingConfigurationsOutput{}
	fnCacher := func(out *mediapackagevod.ListPackagingConfigurationsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListPackagingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListPackagingConfigurationsPagesWithContext(ctx context.Context, input *mediapackagevod.ListPackagingConfigurationsInput, fn func(*mediapackagevod.ListPackagingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListPackagingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListPackagingConfigurationsOutput{}
	fnCacher := func(out *mediapackagevod.ListPackagingConfigurationsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListPackagingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListPackagingConfigurationsWithContext(ctx context.Context, input *mediapackagevod.ListPackagingConfigurationsInput, opts ...request.Option) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListPackagingConfigurationsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListPackagingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListPackagingGroups(input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListPackagingGroupsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListPackagingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListPackagingGroupsPages(input *mediapackagevod.ListPackagingGroupsInput, fn func(*mediapackagevod.ListPackagingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListPackagingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListPackagingGroupsOutput{}
	fnCacher := func(out *mediapackagevod.ListPackagingGroupsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListPackagingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListPackagingGroupsPagesWithContext(ctx context.Context, input *mediapackagevod.ListPackagingGroupsInput, fn func(*mediapackagevod.ListPackagingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediapackagevod.ListPackagingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &mediapackagevod.ListPackagingGroupsOutput{}
	fnCacher := func(out *mediapackagevod.ListPackagingGroupsOutput, more bool) bool {
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
	if err := c.MediaPackageVodAPI.ListPackagingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaPackageVod) ListPackagingGroupsWithContext(ctx context.Context, input *mediapackagevod.ListPackagingGroupsInput, opts ...request.Option) (*mediapackagevod.ListPackagingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListPackagingGroupsOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListPackagingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListTagsForResource(input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaPackageVod) ListTagsForResourceWithContext(ctx context.Context, input *mediapackagevod.ListTagsForResourceInput, opts ...request.Option) (*mediapackagevod.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediapackagevod.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaPackageVodAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
