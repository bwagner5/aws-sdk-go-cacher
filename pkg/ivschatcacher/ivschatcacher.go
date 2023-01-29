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
package ivschatcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ivschat"
	"github.com/aws/aws-sdk-go/service/ivschat/ivschatiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Ivschat struct {
	ivschatiface.IvschatAPI
	cache *cache.Cache
}

func New(ivschatapi ivschatiface.IvschatAPI) *Ivschat {
	return &Ivschat{
		IvschatAPI: ivschatapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Ivschat) GetLoggingConfiguration(input *ivschat.GetLoggingConfigurationInput) (*ivschat.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.GetLoggingConfigurationOutput), nil
	}
	out, err := c.IvschatAPI.GetLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) GetLoggingConfigurationWithContext(ctx context.Context, input *ivschat.GetLoggingConfigurationInput, opts ...request.Option) (*ivschat.GetLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.GetLoggingConfigurationOutput), nil
	}
	out, err := c.IvschatAPI.GetLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) GetRoom(input *ivschat.GetRoomInput) (*ivschat.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.GetRoomOutput), nil
	}
	out, err := c.IvschatAPI.GetRoom(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) GetRoomWithContext(ctx context.Context, input *ivschat.GetRoomInput, opts ...request.Option) (*ivschat.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.GetRoomOutput), nil
	}
	out, err := c.IvschatAPI.GetRoomWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListLoggingConfigurations(input *ivschat.ListLoggingConfigurationsInput) (*ivschat.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.IvschatAPI.ListLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListLoggingConfigurationsPages(input *ivschat.ListLoggingConfigurationsInput, fn func(*ivschat.ListLoggingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ivschat.ListLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ivschat.ListLoggingConfigurationsOutput{}
	fnCacher := func(out *ivschat.ListLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.IvschatAPI.ListLoggingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Ivschat) ListLoggingConfigurationsPagesWithContext(ctx context.Context, input *ivschat.ListLoggingConfigurationsInput, fn func(*ivschat.ListLoggingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ivschat.ListLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ivschat.ListLoggingConfigurationsOutput{}
	fnCacher := func(out *ivschat.ListLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.IvschatAPI.ListLoggingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Ivschat) ListLoggingConfigurationsWithContext(ctx context.Context, input *ivschat.ListLoggingConfigurationsInput, opts ...request.Option) (*ivschat.ListLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListLoggingConfigurationsOutput), nil
	}
	out, err := c.IvschatAPI.ListLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListRooms(input *ivschat.ListRoomsInput) (*ivschat.ListRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListRoomsOutput), nil
	}
	out, err := c.IvschatAPI.ListRooms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListRoomsPages(input *ivschat.ListRoomsInput, fn func(*ivschat.ListRoomsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ivschat.ListRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &ivschat.ListRoomsOutput{}
	fnCacher := func(out *ivschat.ListRoomsOutput, more bool) bool {
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
	if err := c.IvschatAPI.ListRoomsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Ivschat) ListRoomsPagesWithContext(ctx context.Context, input *ivschat.ListRoomsInput, fn func(*ivschat.ListRoomsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ivschat.ListRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &ivschat.ListRoomsOutput{}
	fnCacher := func(out *ivschat.ListRoomsOutput, more bool) bool {
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
	if err := c.IvschatAPI.ListRoomsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Ivschat) ListRoomsWithContext(ctx context.Context, input *ivschat.ListRoomsInput, opts ...request.Option) (*ivschat.ListRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListRoomsOutput), nil
	}
	out, err := c.IvschatAPI.ListRoomsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListTagsForResource(input *ivschat.ListTagsForResourceInput) (*ivschat.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListTagsForResourceOutput), nil
	}
	out, err := c.IvschatAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Ivschat) ListTagsForResourceWithContext(ctx context.Context, input *ivschat.ListTagsForResourceInput, opts ...request.Option) (*ivschat.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ivschat.ListTagsForResourceOutput), nil
	}
	out, err := c.IvschatAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
