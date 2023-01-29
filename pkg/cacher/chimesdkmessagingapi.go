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
	"github.com/aws/aws-sdk-go/service/chimesdkmessaging"
	"github.com/aws/aws-sdk-go/service/chimesdkmessaging/chimesdkmessagingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ChimeSDKMessaging struct {
	chimesdkmessagingiface.ChimeSDKMessagingAPI
	cache *cache.Cache
}

func NewChimeSDKMessaging(chimesdkmessagingapi chimesdkmessagingiface.ChimeSDKMessagingAPI) *ChimeSDKMessaging {
	return &ChimeSDKMessaging{
		ChimeSDKMessagingAPI: chimesdkmessagingapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ChimeSDKMessaging) BatchCreateChannelMembership(input *chimesdkmessaging.BatchCreateChannelMembershipInput) (*chimesdkmessaging.BatchCreateChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.BatchCreateChannelMembershipOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.BatchCreateChannelMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) BatchCreateChannelMembershipWithContext(ctx context.Context, input *chimesdkmessaging.BatchCreateChannelMembershipInput, opts ...request.Option) (*chimesdkmessaging.BatchCreateChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.BatchCreateChannelMembershipOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.BatchCreateChannelMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannel(input *chimesdkmessaging.DescribeChannelInput) (*chimesdkmessaging.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelBan(input *chimesdkmessaging.DescribeChannelBanInput) (*chimesdkmessaging.DescribeChannelBanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelBanOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelBan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelBanWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelBanInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelBanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelBanOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelBanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelFlow(input *chimesdkmessaging.DescribeChannelFlowInput) (*chimesdkmessaging.DescribeChannelFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelFlowOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelFlow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelFlowWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelFlowInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelFlowOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelFlowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelMembership(input *chimesdkmessaging.DescribeChannelMembershipInput) (*chimesdkmessaging.DescribeChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelMembershipOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelMembershipForAppInstanceUser(input *chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput) (*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelMembershipForAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelMembershipForAppInstanceUserWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelMembershipForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelMembershipForAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelMembershipWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelMembershipInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelMembershipOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelModeratedByAppInstanceUser(input *chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput) (*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelModeratedByAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelModeratedByAppInstanceUserWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelModeratedByAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelModerator(input *chimesdkmessaging.DescribeChannelModeratorInput) (*chimesdkmessaging.DescribeChannelModeratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelModeratorOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelModerator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelModeratorWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelModeratorInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelModeratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelModeratorOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelModeratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) DescribeChannelWithContext(ctx context.Context, input *chimesdkmessaging.DescribeChannelInput, opts ...request.Option) (*chimesdkmessaging.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.DescribeChannelOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMembershipPreferences(input *chimesdkmessaging.GetChannelMembershipPreferencesInput) (*chimesdkmessaging.GetChannelMembershipPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMembershipPreferencesOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMembershipPreferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMembershipPreferencesWithContext(ctx context.Context, input *chimesdkmessaging.GetChannelMembershipPreferencesInput, opts ...request.Option) (*chimesdkmessaging.GetChannelMembershipPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMembershipPreferencesOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMembershipPreferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMessage(input *chimesdkmessaging.GetChannelMessageInput) (*chimesdkmessaging.GetChannelMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMessageOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMessage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMessageStatus(input *chimesdkmessaging.GetChannelMessageStatusInput) (*chimesdkmessaging.GetChannelMessageStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMessageStatusOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMessageStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMessageStatusWithContext(ctx context.Context, input *chimesdkmessaging.GetChannelMessageStatusInput, opts ...request.Option) (*chimesdkmessaging.GetChannelMessageStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMessageStatusOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMessageStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetChannelMessageWithContext(ctx context.Context, input *chimesdkmessaging.GetChannelMessageInput, opts ...request.Option) (*chimesdkmessaging.GetChannelMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetChannelMessageOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetChannelMessageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetMessagingSessionEndpoint(input *chimesdkmessaging.GetMessagingSessionEndpointInput) (*chimesdkmessaging.GetMessagingSessionEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetMessagingSessionEndpointOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetMessagingSessionEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) GetMessagingSessionEndpointWithContext(ctx context.Context, input *chimesdkmessaging.GetMessagingSessionEndpointInput, opts ...request.Option) (*chimesdkmessaging.GetMessagingSessionEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.GetMessagingSessionEndpointOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.GetMessagingSessionEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelBans(input *chimesdkmessaging.ListChannelBansInput) (*chimesdkmessaging.ListChannelBansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelBansOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelBans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelBansPages(input *chimesdkmessaging.ListChannelBansInput, fn func(*chimesdkmessaging.ListChannelBansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelBansOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelBansOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelBansOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelBansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelBansPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelBansInput, fn func(*chimesdkmessaging.ListChannelBansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelBansOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelBansOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelBansOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelBansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelBansWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelBansInput, opts ...request.Option) (*chimesdkmessaging.ListChannelBansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelBansOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelBansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelFlows(input *chimesdkmessaging.ListChannelFlowsInput) (*chimesdkmessaging.ListChannelFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelFlowsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelFlows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelFlowsPages(input *chimesdkmessaging.ListChannelFlowsInput, fn func(*chimesdkmessaging.ListChannelFlowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelFlowsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelFlowsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelFlowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelFlowsPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelFlowsInput, fn func(*chimesdkmessaging.ListChannelFlowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelFlowsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelFlowsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelFlowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelFlowsWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelFlowsInput, opts ...request.Option) (*chimesdkmessaging.ListChannelFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelFlowsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelFlowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMemberships(input *chimesdkmessaging.ListChannelMembershipsInput) (*chimesdkmessaging.ListChannelMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMembershipsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMembershipsForAppInstanceUser(input *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput) (*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMembershipsForAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMembershipsForAppInstanceUserPages(input *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, fn func(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMembershipsForAppInstanceUserPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMembershipsForAppInstanceUserPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, fn func(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMembershipsForAppInstanceUserPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMembershipsForAppInstanceUserWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMembershipsForAppInstanceUserInput, opts ...request.Option) (*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMembershipsForAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMembershipsForAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMembershipsPages(input *chimesdkmessaging.ListChannelMembershipsInput, fn func(*chimesdkmessaging.ListChannelMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMembershipsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMembershipsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMembershipsPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMembershipsInput, fn func(*chimesdkmessaging.ListChannelMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMembershipsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMembershipsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMembershipsWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMembershipsInput, opts ...request.Option) (*chimesdkmessaging.ListChannelMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMembershipsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMessages(input *chimesdkmessaging.ListChannelMessagesInput) (*chimesdkmessaging.ListChannelMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMessagesOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelMessagesPages(input *chimesdkmessaging.ListChannelMessagesInput, fn func(*chimesdkmessaging.ListChannelMessagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMessagesOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMessagesOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMessagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMessagesPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMessagesInput, fn func(*chimesdkmessaging.ListChannelMessagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelMessagesOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelMessagesOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelMessagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelMessagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelMessagesInput, opts ...request.Option) (*chimesdkmessaging.ListChannelMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelMessagesOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelModerators(input *chimesdkmessaging.ListChannelModeratorsInput) (*chimesdkmessaging.ListChannelModeratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelModeratorsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelModerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelModeratorsPages(input *chimesdkmessaging.ListChannelModeratorsInput, fn func(*chimesdkmessaging.ListChannelModeratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelModeratorsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelModeratorsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelModeratorsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelModeratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelModeratorsPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelModeratorsInput, fn func(*chimesdkmessaging.ListChannelModeratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelModeratorsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelModeratorsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelModeratorsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelModeratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelModeratorsWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelModeratorsInput, opts ...request.Option) (*chimesdkmessaging.ListChannelModeratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelModeratorsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelModeratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannels(input *chimesdkmessaging.ListChannelsInput) (*chimesdkmessaging.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelsAssociatedWithChannelFlow(input *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput) (*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelsAssociatedWithChannelFlow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelsAssociatedWithChannelFlowPages(input *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, fn func(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsAssociatedWithChannelFlowPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsAssociatedWithChannelFlowPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, fn func(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsAssociatedWithChannelFlowPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsAssociatedWithChannelFlowWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsAssociatedWithChannelFlowInput, opts ...request.Option) (*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsAssociatedWithChannelFlowOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelsAssociatedWithChannelFlowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelsModeratedByAppInstanceUser(input *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput) (*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelsModeratedByAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelsModeratedByAppInstanceUserPages(input *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, fn func(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsModeratedByAppInstanceUserPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsModeratedByAppInstanceUserPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, fn func(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsModeratedByAppInstanceUserPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsModeratedByAppInstanceUserWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsModeratedByAppInstanceUserInput, opts ...request.Option) (*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsModeratedByAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelsModeratedByAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListChannelsPages(input *chimesdkmessaging.ListChannelsInput, fn func(*chimesdkmessaging.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsInput, fn func(*chimesdkmessaging.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListChannelsWithContext(ctx context.Context, input *chimesdkmessaging.ListChannelsInput, opts ...request.Option) (*chimesdkmessaging.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListSubChannels(input *chimesdkmessaging.ListSubChannelsInput) (*chimesdkmessaging.ListSubChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListSubChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListSubChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListSubChannelsPages(input *chimesdkmessaging.ListSubChannelsInput, fn func(*chimesdkmessaging.ListSubChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListSubChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListSubChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListSubChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListSubChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListSubChannelsPagesWithContext(ctx context.Context, input *chimesdkmessaging.ListSubChannelsInput, fn func(*chimesdkmessaging.ListSubChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.ListSubChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.ListSubChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.ListSubChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.ListSubChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) ListSubChannelsWithContext(ctx context.Context, input *chimesdkmessaging.ListSubChannelsInput, opts ...request.Option) (*chimesdkmessaging.ListSubChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListSubChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListSubChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListTagsForResource(input *chimesdkmessaging.ListTagsForResourceInput) (*chimesdkmessaging.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) ListTagsForResourceWithContext(ctx context.Context, input *chimesdkmessaging.ListTagsForResourceInput, opts ...request.Option) (*chimesdkmessaging.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) SearchChannels(input *chimesdkmessaging.SearchChannelsInput) (*chimesdkmessaging.SearchChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.SearchChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.SearchChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMessaging) SearchChannelsPages(input *chimesdkmessaging.SearchChannelsInput, fn func(*chimesdkmessaging.SearchChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.SearchChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.SearchChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.SearchChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.SearchChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) SearchChannelsPagesWithContext(ctx context.Context, input *chimesdkmessaging.SearchChannelsInput, fn func(*chimesdkmessaging.SearchChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmessaging.SearchChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmessaging.SearchChannelsOutput{}
	fnCacher := func(out *chimesdkmessaging.SearchChannelsOutput, more bool) bool {
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
	if err := c.ChimeSDKMessagingAPI.SearchChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMessaging) SearchChannelsWithContext(ctx context.Context, input *chimesdkmessaging.SearchChannelsInput, opts ...request.Option) (*chimesdkmessaging.SearchChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmessaging.SearchChannelsOutput), nil
	}
	out, err := c.ChimeSDKMessagingAPI.SearchChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
