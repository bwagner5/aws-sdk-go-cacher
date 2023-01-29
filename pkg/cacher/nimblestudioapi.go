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
	"github.com/aws/aws-sdk-go/service/nimblestudio"
	"github.com/aws/aws-sdk-go/service/nimblestudio/nimblestudioiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type NimbleStudio struct {
	nimblestudioiface.NimbleStudioAPI
	cache *cache.Cache
}

func NewNimbleStudio(nimblestudioapi nimblestudioiface.NimbleStudioAPI) *NimbleStudio {
	return &NimbleStudio{
		NimbleStudioAPI: nimblestudioapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *NimbleStudio) GetEula(input *nimblestudio.GetEulaInput) (*nimblestudio.GetEulaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetEulaOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetEula(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetEulaWithContext(ctx context.Context, input *nimblestudio.GetEulaInput, opts ...request.Option) (*nimblestudio.GetEulaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetEulaOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetEulaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfile(input *nimblestudio.GetLaunchProfileInput) (*nimblestudio.GetLaunchProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileDetails(input *nimblestudio.GetLaunchProfileDetailsInput) (*nimblestudio.GetLaunchProfileDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileDetailsOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileDetailsWithContext(ctx context.Context, input *nimblestudio.GetLaunchProfileDetailsInput, opts ...request.Option) (*nimblestudio.GetLaunchProfileDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileDetailsOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileInitialization(input *nimblestudio.GetLaunchProfileInitializationInput) (*nimblestudio.GetLaunchProfileInitializationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileInitializationOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileInitialization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileInitializationWithContext(ctx context.Context, input *nimblestudio.GetLaunchProfileInitializationInput, opts ...request.Option) (*nimblestudio.GetLaunchProfileInitializationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileInitializationOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileInitializationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileMember(input *nimblestudio.GetLaunchProfileMemberInput) (*nimblestudio.GetLaunchProfileMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileMemberOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileMemberWithContext(ctx context.Context, input *nimblestudio.GetLaunchProfileMemberInput, opts ...request.Option) (*nimblestudio.GetLaunchProfileMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileMemberOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetLaunchProfileWithContext(ctx context.Context, input *nimblestudio.GetLaunchProfileInput, opts ...request.Option) (*nimblestudio.GetLaunchProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetLaunchProfileOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetLaunchProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingImage(input *nimblestudio.GetStreamingImageInput) (*nimblestudio.GetStreamingImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingImageOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingImageWithContext(ctx context.Context, input *nimblestudio.GetStreamingImageInput, opts ...request.Option) (*nimblestudio.GetStreamingImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingImageOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSession(input *nimblestudio.GetStreamingSessionInput) (*nimblestudio.GetStreamingSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSessionBackup(input *nimblestudio.GetStreamingSessionBackupInput) (*nimblestudio.GetStreamingSessionBackupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionBackupOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSessionBackup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSessionBackupWithContext(ctx context.Context, input *nimblestudio.GetStreamingSessionBackupInput, opts ...request.Option) (*nimblestudio.GetStreamingSessionBackupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionBackupOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSessionBackupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSessionStream(input *nimblestudio.GetStreamingSessionStreamInput) (*nimblestudio.GetStreamingSessionStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionStreamOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSessionStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSessionStreamWithContext(ctx context.Context, input *nimblestudio.GetStreamingSessionStreamInput, opts ...request.Option) (*nimblestudio.GetStreamingSessionStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionStreamOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSessionStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStreamingSessionWithContext(ctx context.Context, input *nimblestudio.GetStreamingSessionInput, opts ...request.Option) (*nimblestudio.GetStreamingSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStreamingSessionOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStreamingSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudio(input *nimblestudio.GetStudioInput) (*nimblestudio.GetStudioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudio(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudioComponent(input *nimblestudio.GetStudioComponentInput) (*nimblestudio.GetStudioComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioComponentOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudioComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudioComponentWithContext(ctx context.Context, input *nimblestudio.GetStudioComponentInput, opts ...request.Option) (*nimblestudio.GetStudioComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioComponentOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudioComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudioMember(input *nimblestudio.GetStudioMemberInput) (*nimblestudio.GetStudioMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioMemberOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudioMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudioMemberWithContext(ctx context.Context, input *nimblestudio.GetStudioMemberInput, opts ...request.Option) (*nimblestudio.GetStudioMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioMemberOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudioMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) GetStudioWithContext(ctx context.Context, input *nimblestudio.GetStudioInput, opts ...request.Option) (*nimblestudio.GetStudioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.GetStudioOutput), nil
	}
	out, err := c.NimbleStudioAPI.GetStudioWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListEulaAcceptances(input *nimblestudio.ListEulaAcceptancesInput) (*nimblestudio.ListEulaAcceptancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListEulaAcceptancesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListEulaAcceptances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListEulaAcceptancesPages(input *nimblestudio.ListEulaAcceptancesInput, fn func(*nimblestudio.ListEulaAcceptancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListEulaAcceptancesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListEulaAcceptancesOutput{}
	fnCacher := func(out *nimblestudio.ListEulaAcceptancesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListEulaAcceptancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListEulaAcceptancesPagesWithContext(ctx context.Context, input *nimblestudio.ListEulaAcceptancesInput, fn func(*nimblestudio.ListEulaAcceptancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListEulaAcceptancesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListEulaAcceptancesOutput{}
	fnCacher := func(out *nimblestudio.ListEulaAcceptancesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListEulaAcceptancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListEulaAcceptancesWithContext(ctx context.Context, input *nimblestudio.ListEulaAcceptancesInput, opts ...request.Option) (*nimblestudio.ListEulaAcceptancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListEulaAcceptancesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListEulaAcceptancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListEulas(input *nimblestudio.ListEulasInput) (*nimblestudio.ListEulasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListEulasOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListEulas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListEulasPages(input *nimblestudio.ListEulasInput, fn func(*nimblestudio.ListEulasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListEulasOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListEulasOutput{}
	fnCacher := func(out *nimblestudio.ListEulasOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListEulasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListEulasPagesWithContext(ctx context.Context, input *nimblestudio.ListEulasInput, fn func(*nimblestudio.ListEulasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListEulasOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListEulasOutput{}
	fnCacher := func(out *nimblestudio.ListEulasOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListEulasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListEulasWithContext(ctx context.Context, input *nimblestudio.ListEulasInput, opts ...request.Option) (*nimblestudio.ListEulasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListEulasOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListEulasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListLaunchProfileMembers(input *nimblestudio.ListLaunchProfileMembersInput) (*nimblestudio.ListLaunchProfileMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListLaunchProfileMembersOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListLaunchProfileMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListLaunchProfileMembersPages(input *nimblestudio.ListLaunchProfileMembersInput, fn func(*nimblestudio.ListLaunchProfileMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListLaunchProfileMembersOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListLaunchProfileMembersOutput{}
	fnCacher := func(out *nimblestudio.ListLaunchProfileMembersOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListLaunchProfileMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListLaunchProfileMembersPagesWithContext(ctx context.Context, input *nimblestudio.ListLaunchProfileMembersInput, fn func(*nimblestudio.ListLaunchProfileMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListLaunchProfileMembersOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListLaunchProfileMembersOutput{}
	fnCacher := func(out *nimblestudio.ListLaunchProfileMembersOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListLaunchProfileMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListLaunchProfileMembersWithContext(ctx context.Context, input *nimblestudio.ListLaunchProfileMembersInput, opts ...request.Option) (*nimblestudio.ListLaunchProfileMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListLaunchProfileMembersOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListLaunchProfileMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListLaunchProfiles(input *nimblestudio.ListLaunchProfilesInput) (*nimblestudio.ListLaunchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListLaunchProfilesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListLaunchProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListLaunchProfilesPages(input *nimblestudio.ListLaunchProfilesInput, fn func(*nimblestudio.ListLaunchProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListLaunchProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListLaunchProfilesOutput{}
	fnCacher := func(out *nimblestudio.ListLaunchProfilesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListLaunchProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListLaunchProfilesPagesWithContext(ctx context.Context, input *nimblestudio.ListLaunchProfilesInput, fn func(*nimblestudio.ListLaunchProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListLaunchProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListLaunchProfilesOutput{}
	fnCacher := func(out *nimblestudio.ListLaunchProfilesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListLaunchProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListLaunchProfilesWithContext(ctx context.Context, input *nimblestudio.ListLaunchProfilesInput, opts ...request.Option) (*nimblestudio.ListLaunchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListLaunchProfilesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListLaunchProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingImages(input *nimblestudio.ListStreamingImagesInput) (*nimblestudio.ListStreamingImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingImagesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingImagesPages(input *nimblestudio.ListStreamingImagesInput, fn func(*nimblestudio.ListStreamingImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingImagesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingImagesOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingImagesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingImagesPagesWithContext(ctx context.Context, input *nimblestudio.ListStreamingImagesInput, fn func(*nimblestudio.ListStreamingImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingImagesOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingImagesOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingImagesOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingImagesWithContext(ctx context.Context, input *nimblestudio.ListStreamingImagesInput, opts ...request.Option) (*nimblestudio.ListStreamingImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingImagesOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingSessionBackups(input *nimblestudio.ListStreamingSessionBackupsInput) (*nimblestudio.ListStreamingSessionBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingSessionBackupsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingSessionBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingSessionBackupsPages(input *nimblestudio.ListStreamingSessionBackupsInput, fn func(*nimblestudio.ListStreamingSessionBackupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingSessionBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingSessionBackupsOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingSessionBackupsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingSessionBackupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingSessionBackupsPagesWithContext(ctx context.Context, input *nimblestudio.ListStreamingSessionBackupsInput, fn func(*nimblestudio.ListStreamingSessionBackupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingSessionBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingSessionBackupsOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingSessionBackupsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingSessionBackupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingSessionBackupsWithContext(ctx context.Context, input *nimblestudio.ListStreamingSessionBackupsInput, opts ...request.Option) (*nimblestudio.ListStreamingSessionBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingSessionBackupsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingSessionBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingSessions(input *nimblestudio.ListStreamingSessionsInput) (*nimblestudio.ListStreamingSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingSessionsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStreamingSessionsPages(input *nimblestudio.ListStreamingSessionsInput, fn func(*nimblestudio.ListStreamingSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingSessionsOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingSessionsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingSessionsPagesWithContext(ctx context.Context, input *nimblestudio.ListStreamingSessionsInput, fn func(*nimblestudio.ListStreamingSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStreamingSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStreamingSessionsOutput{}
	fnCacher := func(out *nimblestudio.ListStreamingSessionsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStreamingSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStreamingSessionsWithContext(ctx context.Context, input *nimblestudio.ListStreamingSessionsInput, opts ...request.Option) (*nimblestudio.ListStreamingSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStreamingSessionsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStreamingSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudioComponents(input *nimblestudio.ListStudioComponentsInput) (*nimblestudio.ListStudioComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudioComponentsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudioComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudioComponentsPages(input *nimblestudio.ListStudioComponentsInput, fn func(*nimblestudio.ListStudioComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudioComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudioComponentsOutput{}
	fnCacher := func(out *nimblestudio.ListStudioComponentsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudioComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudioComponentsPagesWithContext(ctx context.Context, input *nimblestudio.ListStudioComponentsInput, fn func(*nimblestudio.ListStudioComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudioComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudioComponentsOutput{}
	fnCacher := func(out *nimblestudio.ListStudioComponentsOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudioComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudioComponentsWithContext(ctx context.Context, input *nimblestudio.ListStudioComponentsInput, opts ...request.Option) (*nimblestudio.ListStudioComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudioComponentsOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudioComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudioMembers(input *nimblestudio.ListStudioMembersInput) (*nimblestudio.ListStudioMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudioMembersOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudioMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudioMembersPages(input *nimblestudio.ListStudioMembersInput, fn func(*nimblestudio.ListStudioMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudioMembersOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudioMembersOutput{}
	fnCacher := func(out *nimblestudio.ListStudioMembersOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudioMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudioMembersPagesWithContext(ctx context.Context, input *nimblestudio.ListStudioMembersInput, fn func(*nimblestudio.ListStudioMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudioMembersOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudioMembersOutput{}
	fnCacher := func(out *nimblestudio.ListStudioMembersOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudioMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudioMembersWithContext(ctx context.Context, input *nimblestudio.ListStudioMembersInput, opts ...request.Option) (*nimblestudio.ListStudioMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudioMembersOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudioMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudios(input *nimblestudio.ListStudiosInput) (*nimblestudio.ListStudiosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudiosOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudios(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListStudiosPages(input *nimblestudio.ListStudiosInput, fn func(*nimblestudio.ListStudiosOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudiosOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudiosOutput{}
	fnCacher := func(out *nimblestudio.ListStudiosOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudiosPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudiosPagesWithContext(ctx context.Context, input *nimblestudio.ListStudiosInput, fn func(*nimblestudio.ListStudiosOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*nimblestudio.ListStudiosOutput), false)
		return nil
	}
	cachable := true
	output := &nimblestudio.ListStudiosOutput{}
	fnCacher := func(out *nimblestudio.ListStudiosOutput, more bool) bool {
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
	if err := c.NimbleStudioAPI.ListStudiosPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *NimbleStudio) ListStudiosWithContext(ctx context.Context, input *nimblestudio.ListStudiosInput, opts ...request.Option) (*nimblestudio.ListStudiosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListStudiosOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListStudiosWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListTagsForResource(input *nimblestudio.ListTagsForResourceInput) (*nimblestudio.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListTagsForResourceOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *NimbleStudio) ListTagsForResourceWithContext(ctx context.Context, input *nimblestudio.ListTagsForResourceInput, opts ...request.Option) (*nimblestudio.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*nimblestudio.ListTagsForResourceOutput), nil
	}
	out, err := c.NimbleStudioAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
