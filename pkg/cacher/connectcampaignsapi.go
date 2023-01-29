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
	"github.com/aws/aws-sdk-go/service/connectcampaigns"
	"github.com/aws/aws-sdk-go/service/connectcampaigns/connectcampaignsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConnectCampaigns struct {
	connectcampaignsiface.ConnectCampaignsAPI
	cache *cache.Cache
}

func NewConnectCampaigns(connectcampaignsapi connectcampaignsiface.ConnectCampaignsAPI) *ConnectCampaigns {
	return &ConnectCampaigns{
		ConnectCampaignsAPI: connectcampaignsapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConnectCampaigns) DescribeCampaign(input *connectcampaigns.DescribeCampaignInput) (*connectcampaigns.DescribeCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.DescribeCampaignOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.DescribeCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) DescribeCampaignWithContext(ctx context.Context, input *connectcampaigns.DescribeCampaignInput, opts ...request.Option) (*connectcampaigns.DescribeCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.DescribeCampaignOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.DescribeCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetCampaignState(input *connectcampaigns.GetCampaignStateInput) (*connectcampaigns.GetCampaignStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetCampaignStateOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetCampaignState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetCampaignStateBatch(input *connectcampaigns.GetCampaignStateBatchInput) (*connectcampaigns.GetCampaignStateBatchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetCampaignStateBatchOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetCampaignStateBatch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetCampaignStateBatchWithContext(ctx context.Context, input *connectcampaigns.GetCampaignStateBatchInput, opts ...request.Option) (*connectcampaigns.GetCampaignStateBatchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetCampaignStateBatchOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetCampaignStateBatchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetCampaignStateWithContext(ctx context.Context, input *connectcampaigns.GetCampaignStateInput, opts ...request.Option) (*connectcampaigns.GetCampaignStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetCampaignStateOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetCampaignStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetConnectInstanceConfig(input *connectcampaigns.GetConnectInstanceConfigInput) (*connectcampaigns.GetConnectInstanceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetConnectInstanceConfigOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetConnectInstanceConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetConnectInstanceConfigWithContext(ctx context.Context, input *connectcampaigns.GetConnectInstanceConfigInput, opts ...request.Option) (*connectcampaigns.GetConnectInstanceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetConnectInstanceConfigOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetConnectInstanceConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetInstanceOnboardingJobStatus(input *connectcampaigns.GetInstanceOnboardingJobStatusInput) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetInstanceOnboardingJobStatusOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetInstanceOnboardingJobStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) GetInstanceOnboardingJobStatusWithContext(ctx context.Context, input *connectcampaigns.GetInstanceOnboardingJobStatusInput, opts ...request.Option) (*connectcampaigns.GetInstanceOnboardingJobStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.GetInstanceOnboardingJobStatusOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.GetInstanceOnboardingJobStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) ListCampaigns(input *connectcampaigns.ListCampaignsInput) (*connectcampaigns.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.ListCampaignsOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.ListCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) ListCampaignsPages(input *connectcampaigns.ListCampaignsInput, fn func(*connectcampaigns.ListCampaignsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcampaigns.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcampaigns.ListCampaignsOutput{}
	fnCacher := func(out *connectcampaigns.ListCampaignsOutput, more bool) bool {
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
	if err := c.ConnectCampaignsAPI.ListCampaignsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCampaigns) ListCampaignsPagesWithContext(ctx context.Context, input *connectcampaigns.ListCampaignsInput, fn func(*connectcampaigns.ListCampaignsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*connectcampaigns.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &connectcampaigns.ListCampaignsOutput{}
	fnCacher := func(out *connectcampaigns.ListCampaignsOutput, more bool) bool {
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
	if err := c.ConnectCampaignsAPI.ListCampaignsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConnectCampaigns) ListCampaignsWithContext(ctx context.Context, input *connectcampaigns.ListCampaignsInput, opts ...request.Option) (*connectcampaigns.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.ListCampaignsOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.ListCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) ListTagsForResource(input *connectcampaigns.ListTagsForResourceInput) (*connectcampaigns.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConnectCampaigns) ListTagsForResourceWithContext(ctx context.Context, input *connectcampaigns.ListTagsForResourceInput, opts ...request.Option) (*connectcampaigns.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*connectcampaigns.ListTagsForResourceOutput), nil
	}
	out, err := c.ConnectCampaignsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
