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
	"github.com/aws/aws-sdk-go/service/gamesparks"
	"github.com/aws/aws-sdk-go/service/gamesparks/gamesparksiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GameSparks struct {
	gamesparksiface.GameSparksAPI
	cache *cache.Cache
}

func NewGameSparks(gamesparksapi gamesparksiface.GameSparksAPI) *GameSparks {
	return &GameSparks{
		GameSparksAPI: gamesparksapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GameSparks) GetExtension(input *gamesparks.GetExtensionInput) (*gamesparks.GetExtensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetExtensionOutput), nil
	}
	out, err := c.GameSparksAPI.GetExtension(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetExtensionVersion(input *gamesparks.GetExtensionVersionInput) (*gamesparks.GetExtensionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetExtensionVersionOutput), nil
	}
	out, err := c.GameSparksAPI.GetExtensionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetExtensionVersionWithContext(ctx context.Context, input *gamesparks.GetExtensionVersionInput, opts ...request.Option) (*gamesparks.GetExtensionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetExtensionVersionOutput), nil
	}
	out, err := c.GameSparksAPI.GetExtensionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetExtensionWithContext(ctx context.Context, input *gamesparks.GetExtensionInput, opts ...request.Option) (*gamesparks.GetExtensionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetExtensionOutput), nil
	}
	out, err := c.GameSparksAPI.GetExtensionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGame(input *gamesparks.GetGameInput) (*gamesparks.GetGameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGameOutput), nil
	}
	out, err := c.GameSparksAPI.GetGame(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGameConfiguration(input *gamesparks.GetGameConfigurationInput) (*gamesparks.GetGameConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGameConfigurationOutput), nil
	}
	out, err := c.GameSparksAPI.GetGameConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGameConfigurationWithContext(ctx context.Context, input *gamesparks.GetGameConfigurationInput, opts ...request.Option) (*gamesparks.GetGameConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGameConfigurationOutput), nil
	}
	out, err := c.GameSparksAPI.GetGameConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGameWithContext(ctx context.Context, input *gamesparks.GetGameInput, opts ...request.Option) (*gamesparks.GetGameOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGameOutput), nil
	}
	out, err := c.GameSparksAPI.GetGameWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGeneratedCodeJob(input *gamesparks.GetGeneratedCodeJobInput) (*gamesparks.GetGeneratedCodeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGeneratedCodeJobOutput), nil
	}
	out, err := c.GameSparksAPI.GetGeneratedCodeJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetGeneratedCodeJobWithContext(ctx context.Context, input *gamesparks.GetGeneratedCodeJobInput, opts ...request.Option) (*gamesparks.GetGeneratedCodeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetGeneratedCodeJobOutput), nil
	}
	out, err := c.GameSparksAPI.GetGeneratedCodeJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetPlayerConnectionStatus(input *gamesparks.GetPlayerConnectionStatusInput) (*gamesparks.GetPlayerConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetPlayerConnectionStatusOutput), nil
	}
	out, err := c.GameSparksAPI.GetPlayerConnectionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetPlayerConnectionStatusWithContext(ctx context.Context, input *gamesparks.GetPlayerConnectionStatusInput, opts ...request.Option) (*gamesparks.GetPlayerConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetPlayerConnectionStatusOutput), nil
	}
	out, err := c.GameSparksAPI.GetPlayerConnectionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetSnapshot(input *gamesparks.GetSnapshotInput) (*gamesparks.GetSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetSnapshotOutput), nil
	}
	out, err := c.GameSparksAPI.GetSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetSnapshotWithContext(ctx context.Context, input *gamesparks.GetSnapshotInput, opts ...request.Option) (*gamesparks.GetSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetSnapshotOutput), nil
	}
	out, err := c.GameSparksAPI.GetSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetStage(input *gamesparks.GetStageInput) (*gamesparks.GetStageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetStageOutput), nil
	}
	out, err := c.GameSparksAPI.GetStage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetStageDeployment(input *gamesparks.GetStageDeploymentInput) (*gamesparks.GetStageDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetStageDeploymentOutput), nil
	}
	out, err := c.GameSparksAPI.GetStageDeployment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetStageDeploymentWithContext(ctx context.Context, input *gamesparks.GetStageDeploymentInput, opts ...request.Option) (*gamesparks.GetStageDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetStageDeploymentOutput), nil
	}
	out, err := c.GameSparksAPI.GetStageDeploymentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) GetStageWithContext(ctx context.Context, input *gamesparks.GetStageInput, opts ...request.Option) (*gamesparks.GetStageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.GetStageOutput), nil
	}
	out, err := c.GameSparksAPI.GetStageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListExtensionVersions(input *gamesparks.ListExtensionVersionsInput) (*gamesparks.ListExtensionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListExtensionVersionsOutput), nil
	}
	out, err := c.GameSparksAPI.ListExtensionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListExtensionVersionsPages(input *gamesparks.ListExtensionVersionsInput, fn func(*gamesparks.ListExtensionVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListExtensionVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListExtensionVersionsOutput{}
	fnCacher := func(out *gamesparks.ListExtensionVersionsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListExtensionVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListExtensionVersionsPagesWithContext(ctx context.Context, input *gamesparks.ListExtensionVersionsInput, fn func(*gamesparks.ListExtensionVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListExtensionVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListExtensionVersionsOutput{}
	fnCacher := func(out *gamesparks.ListExtensionVersionsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListExtensionVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListExtensionVersionsWithContext(ctx context.Context, input *gamesparks.ListExtensionVersionsInput, opts ...request.Option) (*gamesparks.ListExtensionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListExtensionVersionsOutput), nil
	}
	out, err := c.GameSparksAPI.ListExtensionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListExtensions(input *gamesparks.ListExtensionsInput) (*gamesparks.ListExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListExtensionsOutput), nil
	}
	out, err := c.GameSparksAPI.ListExtensions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListExtensionsPages(input *gamesparks.ListExtensionsInput, fn func(*gamesparks.ListExtensionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListExtensionsOutput{}
	fnCacher := func(out *gamesparks.ListExtensionsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListExtensionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListExtensionsPagesWithContext(ctx context.Context, input *gamesparks.ListExtensionsInput, fn func(*gamesparks.ListExtensionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListExtensionsOutput{}
	fnCacher := func(out *gamesparks.ListExtensionsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListExtensionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListExtensionsWithContext(ctx context.Context, input *gamesparks.ListExtensionsInput, opts ...request.Option) (*gamesparks.ListExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListExtensionsOutput), nil
	}
	out, err := c.GameSparksAPI.ListExtensionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListGames(input *gamesparks.ListGamesInput) (*gamesparks.ListGamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListGamesOutput), nil
	}
	out, err := c.GameSparksAPI.ListGames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListGamesPages(input *gamesparks.ListGamesInput, fn func(*gamesparks.ListGamesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListGamesOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListGamesOutput{}
	fnCacher := func(out *gamesparks.ListGamesOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListGamesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListGamesPagesWithContext(ctx context.Context, input *gamesparks.ListGamesInput, fn func(*gamesparks.ListGamesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListGamesOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListGamesOutput{}
	fnCacher := func(out *gamesparks.ListGamesOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListGamesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListGamesWithContext(ctx context.Context, input *gamesparks.ListGamesInput, opts ...request.Option) (*gamesparks.ListGamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListGamesOutput), nil
	}
	out, err := c.GameSparksAPI.ListGamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListGeneratedCodeJobs(input *gamesparks.ListGeneratedCodeJobsInput) (*gamesparks.ListGeneratedCodeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListGeneratedCodeJobsOutput), nil
	}
	out, err := c.GameSparksAPI.ListGeneratedCodeJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListGeneratedCodeJobsPages(input *gamesparks.ListGeneratedCodeJobsInput, fn func(*gamesparks.ListGeneratedCodeJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListGeneratedCodeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListGeneratedCodeJobsOutput{}
	fnCacher := func(out *gamesparks.ListGeneratedCodeJobsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListGeneratedCodeJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListGeneratedCodeJobsPagesWithContext(ctx context.Context, input *gamesparks.ListGeneratedCodeJobsInput, fn func(*gamesparks.ListGeneratedCodeJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListGeneratedCodeJobsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListGeneratedCodeJobsOutput{}
	fnCacher := func(out *gamesparks.ListGeneratedCodeJobsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListGeneratedCodeJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListGeneratedCodeJobsWithContext(ctx context.Context, input *gamesparks.ListGeneratedCodeJobsInput, opts ...request.Option) (*gamesparks.ListGeneratedCodeJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListGeneratedCodeJobsOutput), nil
	}
	out, err := c.GameSparksAPI.ListGeneratedCodeJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListSnapshots(input *gamesparks.ListSnapshotsInput) (*gamesparks.ListSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListSnapshotsOutput), nil
	}
	out, err := c.GameSparksAPI.ListSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListSnapshotsPages(input *gamesparks.ListSnapshotsInput, fn func(*gamesparks.ListSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListSnapshotsOutput{}
	fnCacher := func(out *gamesparks.ListSnapshotsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListSnapshotsPagesWithContext(ctx context.Context, input *gamesparks.ListSnapshotsInput, fn func(*gamesparks.ListSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListSnapshotsOutput{}
	fnCacher := func(out *gamesparks.ListSnapshotsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListSnapshotsWithContext(ctx context.Context, input *gamesparks.ListSnapshotsInput, opts ...request.Option) (*gamesparks.ListSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListSnapshotsOutput), nil
	}
	out, err := c.GameSparksAPI.ListSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListStageDeployments(input *gamesparks.ListStageDeploymentsInput) (*gamesparks.ListStageDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListStageDeploymentsOutput), nil
	}
	out, err := c.GameSparksAPI.ListStageDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListStageDeploymentsPages(input *gamesparks.ListStageDeploymentsInput, fn func(*gamesparks.ListStageDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListStageDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListStageDeploymentsOutput{}
	fnCacher := func(out *gamesparks.ListStageDeploymentsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListStageDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListStageDeploymentsPagesWithContext(ctx context.Context, input *gamesparks.ListStageDeploymentsInput, fn func(*gamesparks.ListStageDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListStageDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListStageDeploymentsOutput{}
	fnCacher := func(out *gamesparks.ListStageDeploymentsOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListStageDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListStageDeploymentsWithContext(ctx context.Context, input *gamesparks.ListStageDeploymentsInput, opts ...request.Option) (*gamesparks.ListStageDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListStageDeploymentsOutput), nil
	}
	out, err := c.GameSparksAPI.ListStageDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListStages(input *gamesparks.ListStagesInput) (*gamesparks.ListStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListStagesOutput), nil
	}
	out, err := c.GameSparksAPI.ListStages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListStagesPages(input *gamesparks.ListStagesInput, fn func(*gamesparks.ListStagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListStagesOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListStagesOutput{}
	fnCacher := func(out *gamesparks.ListStagesOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListStagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListStagesPagesWithContext(ctx context.Context, input *gamesparks.ListStagesInput, fn func(*gamesparks.ListStagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamesparks.ListStagesOutput), false)
		return nil
	}
	cachable := true
	output := &gamesparks.ListStagesOutput{}
	fnCacher := func(out *gamesparks.ListStagesOutput, more bool) bool {
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
	if err := c.GameSparksAPI.ListStagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameSparks) ListStagesWithContext(ctx context.Context, input *gamesparks.ListStagesInput, opts ...request.Option) (*gamesparks.ListStagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListStagesOutput), nil
	}
	out, err := c.GameSparksAPI.ListStagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListTagsForResource(input *gamesparks.ListTagsForResourceInput) (*gamesparks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListTagsForResourceOutput), nil
	}
	out, err := c.GameSparksAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameSparks) ListTagsForResourceWithContext(ctx context.Context, input *gamesparks.ListTagsForResourceInput, opts ...request.Option) (*gamesparks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamesparks.ListTagsForResourceOutput), nil
	}
	out, err := c.GameSparksAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
