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
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LexModelBuildingService struct {
	lexmodelbuildingserviceiface.LexModelBuildingServiceAPI
	cache *cache.Cache
}

func NewLexModelBuildingService(lexmodelbuildingserviceapi lexmodelbuildingserviceiface.LexModelBuildingServiceAPI) *LexModelBuildingService {
	return &LexModelBuildingService{
		LexModelBuildingServiceAPI: lexmodelbuildingserviceapi,
		cache:                      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LexModelBuildingService) GetBot(input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotAlias(input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotAliasOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotAliasWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotAliasOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotAliases(input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotAliasesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotAliasesPages(input *lexmodelbuildingservice.GetBotAliasesInput, fn func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotAliasesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotAliasesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotAliasesPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasesInput, fn func(*lexmodelbuildingservice.GetBotAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotAliasesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotAliasesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotAliasesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasesInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotAliasesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotChannelAssociation(input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotChannelAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotChannelAssociationWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotChannelAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotChannelAssociations(input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotChannelAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotChannelAssociationsPages(input *lexmodelbuildingservice.GetBotChannelAssociationsInput, fn func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotChannelAssociationsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotChannelAssociationsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotChannelAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotChannelAssociationsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput, fn func(*lexmodelbuildingservice.GetBotChannelAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotChannelAssociationsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotChannelAssociationsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotChannelAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotChannelAssociationsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotChannelAssociationsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotChannelAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotVersions(input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotVersionsPages(input *lexmodelbuildingservice.GetBotVersionsInput, fn func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotVersionsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotVersionsInput, fn func(*lexmodelbuildingservice.GetBotVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotVersionsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotVersionsInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBots(input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBotsPages(input *lexmodelbuildingservice.GetBotsInput, fn func(*lexmodelbuildingservice.GetBotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotsInput, fn func(*lexmodelbuildingservice.GetBotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBotsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBotsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBotsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBotsInput, opts ...request.Option) (*lexmodelbuildingservice.GetBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBotsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinIntent(input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinIntent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinIntentWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput, opts ...request.Option) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinIntentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinIntents(input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinIntents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinIntentsPages(input *lexmodelbuildingservice.GetBuiltinIntentsInput, fn func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBuiltinIntentsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBuiltinIntentsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBuiltinIntentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBuiltinIntentsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput, fn func(*lexmodelbuildingservice.GetBuiltinIntentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBuiltinIntentsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBuiltinIntentsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBuiltinIntentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBuiltinIntentsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput, opts ...request.Option) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinIntentsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinIntentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinSlotTypes(input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinSlotTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetBuiltinSlotTypesPages(input *lexmodelbuildingservice.GetBuiltinSlotTypesInput, fn func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBuiltinSlotTypesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBuiltinSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBuiltinSlotTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBuiltinSlotTypesPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput, fn func(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetBuiltinSlotTypesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetBuiltinSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetBuiltinSlotTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetBuiltinSlotTypesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput, opts ...request.Option) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetBuiltinSlotTypesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetBuiltinSlotTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetExport(input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetExportOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetExportWithContext(ctx context.Context, input *lexmodelbuildingservice.GetExportInput, opts ...request.Option) (*lexmodelbuildingservice.GetExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetExportOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetImport(input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetImportOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetImport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetImportWithContext(ctx context.Context, input *lexmodelbuildingservice.GetImportInput, opts ...request.Option) (*lexmodelbuildingservice.GetImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetImportOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetImportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntent(input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntentVersions(input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntentVersionsPages(input *lexmodelbuildingservice.GetIntentVersionsInput, fn func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetIntentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetIntentVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetIntentVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetIntentVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetIntentVersionsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetIntentVersionsInput, fn func(*lexmodelbuildingservice.GetIntentVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetIntentVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetIntentVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetIntentVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetIntentVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetIntentVersionsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetIntentVersionsInput, opts ...request.Option) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntentWithContext(ctx context.Context, input *lexmodelbuildingservice.GetIntentInput, opts ...request.Option) (*lexmodelbuildingservice.GetIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntents(input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetIntentsPages(input *lexmodelbuildingservice.GetIntentsInput, fn func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetIntentsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetIntentsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetIntentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetIntentsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetIntentsInput, fn func(*lexmodelbuildingservice.GetIntentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetIntentsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetIntentsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetIntentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetIntentsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetIntentsInput, opts ...request.Option) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetIntentsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetIntentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetMigration(input *lexmodelbuildingservice.GetMigrationInput) (*lexmodelbuildingservice.GetMigrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetMigrationOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetMigration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetMigrationWithContext(ctx context.Context, input *lexmodelbuildingservice.GetMigrationInput, opts ...request.Option) (*lexmodelbuildingservice.GetMigrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetMigrationOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetMigrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetMigrations(input *lexmodelbuildingservice.GetMigrationsInput) (*lexmodelbuildingservice.GetMigrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetMigrationsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetMigrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetMigrationsPages(input *lexmodelbuildingservice.GetMigrationsInput, fn func(*lexmodelbuildingservice.GetMigrationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetMigrationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetMigrationsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetMigrationsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetMigrationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetMigrationsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetMigrationsInput, fn func(*lexmodelbuildingservice.GetMigrationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetMigrationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetMigrationsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetMigrationsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetMigrationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetMigrationsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetMigrationsInput, opts ...request.Option) (*lexmodelbuildingservice.GetMigrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetMigrationsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetMigrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotType(input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypeOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotTypeVersions(input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypeVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotTypeVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotTypeVersionsPages(input *lexmodelbuildingservice.GetSlotTypeVersionsInput, fn func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetSlotTypeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetSlotTypeVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetSlotTypeVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetSlotTypeVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetSlotTypeVersionsPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput, fn func(*lexmodelbuildingservice.GetSlotTypeVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetSlotTypeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetSlotTypeVersionsOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetSlotTypeVersionsOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetSlotTypeVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetSlotTypeVersionsWithContext(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput, opts ...request.Option) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypeVersionsOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotTypeVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotTypeWithContext(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeInput, opts ...request.Option) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypeOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotTypes(input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetSlotTypesPages(input *lexmodelbuildingservice.GetSlotTypesInput, fn func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetSlotTypesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetSlotTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetSlotTypesPagesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypesInput, fn func(*lexmodelbuildingservice.GetSlotTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelbuildingservice.GetSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelbuildingservice.GetSlotTypesOutput{}
	fnCacher := func(out *lexmodelbuildingservice.GetSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelBuildingServiceAPI.GetSlotTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelBuildingService) GetSlotTypesWithContext(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypesInput, opts ...request.Option) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetSlotTypesOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetSlotTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetUtterancesView(input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetUtterancesViewOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetUtterancesView(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) GetUtterancesViewWithContext(ctx context.Context, input *lexmodelbuildingservice.GetUtterancesViewInput, opts ...request.Option) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.GetUtterancesViewOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.GetUtterancesViewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) ListTagsForResource(input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.ListTagsForResourceOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelBuildingService) ListTagsForResourceWithContext(ctx context.Context, input *lexmodelbuildingservice.ListTagsForResourceInput, opts ...request.Option) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelbuildingservice.ListTagsForResourceOutput), nil
	}
	out, err := c.LexModelBuildingServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
