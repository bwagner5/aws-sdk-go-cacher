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
package mediatailorcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mediatailor"
	"github.com/aws/aws-sdk-go/service/mediatailor/mediatailoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaTailor struct {
	mediatailoriface.MediaTailorAPI
	cache *cache.Cache
}

func New(mediatailorapi mediatailoriface.MediaTailorAPI) *MediaTailor {
	return &MediaTailor{
		MediaTailorAPI: mediatailorapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaTailor) DescribeChannel(input *mediatailor.DescribeChannelInput) (*mediatailor.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeChannelOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeChannelWithContext(ctx context.Context, input *mediatailor.DescribeChannelInput, opts ...request.Option) (*mediatailor.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeChannelOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeLiveSource(input *mediatailor.DescribeLiveSourceInput) (*mediatailor.DescribeLiveSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeLiveSourceOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeLiveSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeLiveSourceWithContext(ctx context.Context, input *mediatailor.DescribeLiveSourceInput, opts ...request.Option) (*mediatailor.DescribeLiveSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeLiveSourceOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeLiveSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeProgram(input *mediatailor.DescribeProgramInput) (*mediatailor.DescribeProgramOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeProgramOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeProgram(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeProgramWithContext(ctx context.Context, input *mediatailor.DescribeProgramInput, opts ...request.Option) (*mediatailor.DescribeProgramOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeProgramOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeProgramWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeSourceLocation(input *mediatailor.DescribeSourceLocationInput) (*mediatailor.DescribeSourceLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeSourceLocationOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeSourceLocation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeSourceLocationWithContext(ctx context.Context, input *mediatailor.DescribeSourceLocationInput, opts ...request.Option) (*mediatailor.DescribeSourceLocationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeSourceLocationOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeSourceLocationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeVodSource(input *mediatailor.DescribeVodSourceInput) (*mediatailor.DescribeVodSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeVodSourceOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeVodSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) DescribeVodSourceWithContext(ctx context.Context, input *mediatailor.DescribeVodSourceInput, opts ...request.Option) (*mediatailor.DescribeVodSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.DescribeVodSourceOutput), nil
	}
	out, err := c.MediaTailorAPI.DescribeVodSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetChannelPolicy(input *mediatailor.GetChannelPolicyInput) (*mediatailor.GetChannelPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetChannelPolicyOutput), nil
	}
	out, err := c.MediaTailorAPI.GetChannelPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetChannelPolicyWithContext(ctx context.Context, input *mediatailor.GetChannelPolicyInput, opts ...request.Option) (*mediatailor.GetChannelPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetChannelPolicyOutput), nil
	}
	out, err := c.MediaTailorAPI.GetChannelPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetChannelSchedule(input *mediatailor.GetChannelScheduleInput) (*mediatailor.GetChannelScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetChannelScheduleOutput), nil
	}
	out, err := c.MediaTailorAPI.GetChannelSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetChannelSchedulePages(input *mediatailor.GetChannelScheduleInput, fn func(*mediatailor.GetChannelScheduleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.GetChannelScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.GetChannelScheduleOutput{}
	fnCacher := func(out *mediatailor.GetChannelScheduleOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.GetChannelSchedulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) GetChannelSchedulePagesWithContext(ctx context.Context, input *mediatailor.GetChannelScheduleInput, fn func(*mediatailor.GetChannelScheduleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.GetChannelScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.GetChannelScheduleOutput{}
	fnCacher := func(out *mediatailor.GetChannelScheduleOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.GetChannelSchedulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) GetChannelScheduleWithContext(ctx context.Context, input *mediatailor.GetChannelScheduleInput, opts ...request.Option) (*mediatailor.GetChannelScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetChannelScheduleOutput), nil
	}
	out, err := c.MediaTailorAPI.GetChannelScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetPlaybackConfiguration(input *mediatailor.GetPlaybackConfigurationInput) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetPlaybackConfigurationOutput), nil
	}
	out, err := c.MediaTailorAPI.GetPlaybackConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetPlaybackConfigurationWithContext(ctx context.Context, input *mediatailor.GetPlaybackConfigurationInput, opts ...request.Option) (*mediatailor.GetPlaybackConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetPlaybackConfigurationOutput), nil
	}
	out, err := c.MediaTailorAPI.GetPlaybackConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetPrefetchSchedule(input *mediatailor.GetPrefetchScheduleInput) (*mediatailor.GetPrefetchScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetPrefetchScheduleOutput), nil
	}
	out, err := c.MediaTailorAPI.GetPrefetchSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) GetPrefetchScheduleWithContext(ctx context.Context, input *mediatailor.GetPrefetchScheduleInput, opts ...request.Option) (*mediatailor.GetPrefetchScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.GetPrefetchScheduleOutput), nil
	}
	out, err := c.MediaTailorAPI.GetPrefetchScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListAlerts(input *mediatailor.ListAlertsInput) (*mediatailor.ListAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListAlertsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListAlerts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListAlertsPages(input *mediatailor.ListAlertsInput, fn func(*mediatailor.ListAlertsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListAlertsOutput{}
	fnCacher := func(out *mediatailor.ListAlertsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListAlertsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListAlertsPagesWithContext(ctx context.Context, input *mediatailor.ListAlertsInput, fn func(*mediatailor.ListAlertsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListAlertsOutput{}
	fnCacher := func(out *mediatailor.ListAlertsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListAlertsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListAlertsWithContext(ctx context.Context, input *mediatailor.ListAlertsInput, opts ...request.Option) (*mediatailor.ListAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListAlertsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListAlertsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListChannels(input *mediatailor.ListChannelsInput) (*mediatailor.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListChannelsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListChannelsPages(input *mediatailor.ListChannelsInput, fn func(*mediatailor.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListChannelsOutput{}
	fnCacher := func(out *mediatailor.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListChannelsPagesWithContext(ctx context.Context, input *mediatailor.ListChannelsInput, fn func(*mediatailor.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListChannelsOutput{}
	fnCacher := func(out *mediatailor.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListChannelsWithContext(ctx context.Context, input *mediatailor.ListChannelsInput, opts ...request.Option) (*mediatailor.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListChannelsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListLiveSources(input *mediatailor.ListLiveSourcesInput) (*mediatailor.ListLiveSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListLiveSourcesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListLiveSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListLiveSourcesPages(input *mediatailor.ListLiveSourcesInput, fn func(*mediatailor.ListLiveSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListLiveSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListLiveSourcesOutput{}
	fnCacher := func(out *mediatailor.ListLiveSourcesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListLiveSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListLiveSourcesPagesWithContext(ctx context.Context, input *mediatailor.ListLiveSourcesInput, fn func(*mediatailor.ListLiveSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListLiveSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListLiveSourcesOutput{}
	fnCacher := func(out *mediatailor.ListLiveSourcesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListLiveSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListLiveSourcesWithContext(ctx context.Context, input *mediatailor.ListLiveSourcesInput, opts ...request.Option) (*mediatailor.ListLiveSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListLiveSourcesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListLiveSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListPlaybackConfigurations(input *mediatailor.ListPlaybackConfigurationsInput) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListPlaybackConfigurationsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListPlaybackConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListPlaybackConfigurationsPages(input *mediatailor.ListPlaybackConfigurationsInput, fn func(*mediatailor.ListPlaybackConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListPlaybackConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListPlaybackConfigurationsOutput{}
	fnCacher := func(out *mediatailor.ListPlaybackConfigurationsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListPlaybackConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListPlaybackConfigurationsPagesWithContext(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput, fn func(*mediatailor.ListPlaybackConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListPlaybackConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListPlaybackConfigurationsOutput{}
	fnCacher := func(out *mediatailor.ListPlaybackConfigurationsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListPlaybackConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListPlaybackConfigurationsWithContext(ctx context.Context, input *mediatailor.ListPlaybackConfigurationsInput, opts ...request.Option) (*mediatailor.ListPlaybackConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListPlaybackConfigurationsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListPlaybackConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListPrefetchSchedules(input *mediatailor.ListPrefetchSchedulesInput) (*mediatailor.ListPrefetchSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListPrefetchSchedulesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListPrefetchSchedules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListPrefetchSchedulesPages(input *mediatailor.ListPrefetchSchedulesInput, fn func(*mediatailor.ListPrefetchSchedulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListPrefetchSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListPrefetchSchedulesOutput{}
	fnCacher := func(out *mediatailor.ListPrefetchSchedulesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListPrefetchSchedulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListPrefetchSchedulesPagesWithContext(ctx context.Context, input *mediatailor.ListPrefetchSchedulesInput, fn func(*mediatailor.ListPrefetchSchedulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListPrefetchSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListPrefetchSchedulesOutput{}
	fnCacher := func(out *mediatailor.ListPrefetchSchedulesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListPrefetchSchedulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListPrefetchSchedulesWithContext(ctx context.Context, input *mediatailor.ListPrefetchSchedulesInput, opts ...request.Option) (*mediatailor.ListPrefetchSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListPrefetchSchedulesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListPrefetchSchedulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListSourceLocations(input *mediatailor.ListSourceLocationsInput) (*mediatailor.ListSourceLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListSourceLocationsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListSourceLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListSourceLocationsPages(input *mediatailor.ListSourceLocationsInput, fn func(*mediatailor.ListSourceLocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListSourceLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListSourceLocationsOutput{}
	fnCacher := func(out *mediatailor.ListSourceLocationsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListSourceLocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListSourceLocationsPagesWithContext(ctx context.Context, input *mediatailor.ListSourceLocationsInput, fn func(*mediatailor.ListSourceLocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListSourceLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListSourceLocationsOutput{}
	fnCacher := func(out *mediatailor.ListSourceLocationsOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListSourceLocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListSourceLocationsWithContext(ctx context.Context, input *mediatailor.ListSourceLocationsInput, opts ...request.Option) (*mediatailor.ListSourceLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListSourceLocationsOutput), nil
	}
	out, err := c.MediaTailorAPI.ListSourceLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListTagsForResource(input *mediatailor.ListTagsForResourceInput) (*mediatailor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaTailorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListTagsForResourceWithContext(ctx context.Context, input *mediatailor.ListTagsForResourceInput, opts ...request.Option) (*mediatailor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaTailorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListVodSources(input *mediatailor.ListVodSourcesInput) (*mediatailor.ListVodSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListVodSourcesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListVodSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaTailor) ListVodSourcesPages(input *mediatailor.ListVodSourcesInput, fn func(*mediatailor.ListVodSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListVodSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListVodSourcesOutput{}
	fnCacher := func(out *mediatailor.ListVodSourcesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListVodSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListVodSourcesPagesWithContext(ctx context.Context, input *mediatailor.ListVodSourcesInput, fn func(*mediatailor.ListVodSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediatailor.ListVodSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &mediatailor.ListVodSourcesOutput{}
	fnCacher := func(out *mediatailor.ListVodSourcesOutput, more bool) bool {
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
	if err := c.MediaTailorAPI.ListVodSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaTailor) ListVodSourcesWithContext(ctx context.Context, input *mediatailor.ListVodSourcesInput, opts ...request.Option) (*mediatailor.ListVodSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediatailor.ListVodSourcesOutput), nil
	}
	out, err := c.MediaTailorAPI.ListVodSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
