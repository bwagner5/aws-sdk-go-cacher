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
	"github.com/aws/aws-sdk-go/service/chimesdkmeetings"
	"github.com/aws/aws-sdk-go/service/chimesdkmeetings/chimesdkmeetingsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ChimeSDKMeetings struct {
	chimesdkmeetingsiface.ChimeSDKMeetingsAPI
	cache *cache.Cache
}

func NewChimeSDKMeetings(chimesdkmeetingsapi chimesdkmeetingsiface.ChimeSDKMeetingsAPI) *ChimeSDKMeetings {
	return &ChimeSDKMeetings{
		ChimeSDKMeetingsAPI: chimesdkmeetingsapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ChimeSDKMeetings) BatchCreateAttendee(input *chimesdkmeetings.BatchCreateAttendeeInput) (*chimesdkmeetings.BatchCreateAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.BatchCreateAttendeeOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.BatchCreateAttendee(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) BatchCreateAttendeeWithContext(ctx context.Context, input *chimesdkmeetings.BatchCreateAttendeeInput, opts ...request.Option) (*chimesdkmeetings.BatchCreateAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.BatchCreateAttendeeOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.BatchCreateAttendeeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) BatchUpdateAttendeeCapabilitiesExcept(input *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput) (*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.BatchUpdateAttendeeCapabilitiesExcept(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) BatchUpdateAttendeeCapabilitiesExceptWithContext(ctx context.Context, input *chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptInput, opts ...request.Option) (*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.BatchUpdateAttendeeCapabilitiesExceptOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.BatchUpdateAttendeeCapabilitiesExceptWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) GetAttendee(input *chimesdkmeetings.GetAttendeeInput) (*chimesdkmeetings.GetAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.GetAttendeeOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.GetAttendee(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) GetAttendeeWithContext(ctx context.Context, input *chimesdkmeetings.GetAttendeeInput, opts ...request.Option) (*chimesdkmeetings.GetAttendeeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.GetAttendeeOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.GetAttendeeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) GetMeeting(input *chimesdkmeetings.GetMeetingInput) (*chimesdkmeetings.GetMeetingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.GetMeetingOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.GetMeeting(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) GetMeetingWithContext(ctx context.Context, input *chimesdkmeetings.GetMeetingInput, opts ...request.Option) (*chimesdkmeetings.GetMeetingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.GetMeetingOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.GetMeetingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) ListAttendees(input *chimesdkmeetings.ListAttendeesInput) (*chimesdkmeetings.ListAttendeesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.ListAttendeesOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.ListAttendees(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) ListAttendeesPages(input *chimesdkmeetings.ListAttendeesInput, fn func(*chimesdkmeetings.ListAttendeesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmeetings.ListAttendeesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmeetings.ListAttendeesOutput{}
	fnCacher := func(out *chimesdkmeetings.ListAttendeesOutput, more bool) bool {
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
	if err := c.ChimeSDKMeetingsAPI.ListAttendeesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMeetings) ListAttendeesPagesWithContext(ctx context.Context, input *chimesdkmeetings.ListAttendeesInput, fn func(*chimesdkmeetings.ListAttendeesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkmeetings.ListAttendeesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkmeetings.ListAttendeesOutput{}
	fnCacher := func(out *chimesdkmeetings.ListAttendeesOutput, more bool) bool {
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
	if err := c.ChimeSDKMeetingsAPI.ListAttendeesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKMeetings) ListAttendeesWithContext(ctx context.Context, input *chimesdkmeetings.ListAttendeesInput, opts ...request.Option) (*chimesdkmeetings.ListAttendeesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.ListAttendeesOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.ListAttendeesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) ListTagsForResource(input *chimesdkmeetings.ListTagsForResourceInput) (*chimesdkmeetings.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKMeetings) ListTagsForResourceWithContext(ctx context.Context, input *chimesdkmeetings.ListTagsForResourceInput, opts ...request.Option) (*chimesdkmeetings.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkmeetings.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKMeetingsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
