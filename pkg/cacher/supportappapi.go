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
	"github.com/aws/aws-sdk-go/service/supportapp"
	"github.com/aws/aws-sdk-go/service/supportapp/supportappiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SupportApp struct {
	supportappiface.SupportAppAPI
	cache *cache.Cache
}

func NewSupportApp(supportappapi supportappiface.SupportAppAPI) *SupportApp {
	return &SupportApp{
		SupportAppAPI: supportappapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SupportApp) GetAccountAlias(input *supportapp.GetAccountAliasInput) (*supportapp.GetAccountAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.GetAccountAliasOutput), nil
	}
	out, err := c.SupportAppAPI.GetAccountAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SupportApp) GetAccountAliasWithContext(ctx context.Context, input *supportapp.GetAccountAliasInput, opts ...request.Option) (*supportapp.GetAccountAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.GetAccountAliasOutput), nil
	}
	out, err := c.SupportAppAPI.GetAccountAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SupportApp) ListSlackChannelConfigurations(input *supportapp.ListSlackChannelConfigurationsInput) (*supportapp.ListSlackChannelConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.ListSlackChannelConfigurationsOutput), nil
	}
	out, err := c.SupportAppAPI.ListSlackChannelConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SupportApp) ListSlackChannelConfigurationsPages(input *supportapp.ListSlackChannelConfigurationsInput, fn func(*supportapp.ListSlackChannelConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*supportapp.ListSlackChannelConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &supportapp.ListSlackChannelConfigurationsOutput{}
	fnCacher := func(out *supportapp.ListSlackChannelConfigurationsOutput, more bool) bool {
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
	if err := c.SupportAppAPI.ListSlackChannelConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SupportApp) ListSlackChannelConfigurationsPagesWithContext(ctx context.Context, input *supportapp.ListSlackChannelConfigurationsInput, fn func(*supportapp.ListSlackChannelConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*supportapp.ListSlackChannelConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &supportapp.ListSlackChannelConfigurationsOutput{}
	fnCacher := func(out *supportapp.ListSlackChannelConfigurationsOutput, more bool) bool {
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
	if err := c.SupportAppAPI.ListSlackChannelConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SupportApp) ListSlackChannelConfigurationsWithContext(ctx context.Context, input *supportapp.ListSlackChannelConfigurationsInput, opts ...request.Option) (*supportapp.ListSlackChannelConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.ListSlackChannelConfigurationsOutput), nil
	}
	out, err := c.SupportAppAPI.ListSlackChannelConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SupportApp) ListSlackWorkspaceConfigurations(input *supportapp.ListSlackWorkspaceConfigurationsInput) (*supportapp.ListSlackWorkspaceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.ListSlackWorkspaceConfigurationsOutput), nil
	}
	out, err := c.SupportAppAPI.ListSlackWorkspaceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SupportApp) ListSlackWorkspaceConfigurationsPages(input *supportapp.ListSlackWorkspaceConfigurationsInput, fn func(*supportapp.ListSlackWorkspaceConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*supportapp.ListSlackWorkspaceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &supportapp.ListSlackWorkspaceConfigurationsOutput{}
	fnCacher := func(out *supportapp.ListSlackWorkspaceConfigurationsOutput, more bool) bool {
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
	if err := c.SupportAppAPI.ListSlackWorkspaceConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SupportApp) ListSlackWorkspaceConfigurationsPagesWithContext(ctx context.Context, input *supportapp.ListSlackWorkspaceConfigurationsInput, fn func(*supportapp.ListSlackWorkspaceConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*supportapp.ListSlackWorkspaceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &supportapp.ListSlackWorkspaceConfigurationsOutput{}
	fnCacher := func(out *supportapp.ListSlackWorkspaceConfigurationsOutput, more bool) bool {
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
	if err := c.SupportAppAPI.ListSlackWorkspaceConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SupportApp) ListSlackWorkspaceConfigurationsWithContext(ctx context.Context, input *supportapp.ListSlackWorkspaceConfigurationsInput, opts ...request.Option) (*supportapp.ListSlackWorkspaceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*supportapp.ListSlackWorkspaceConfigurationsOutput), nil
	}
	out, err := c.SupportAppAPI.ListSlackWorkspaceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
