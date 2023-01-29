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
package medialivecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/medialive/medialiveiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaLive struct {
	medialiveiface.MediaLiveAPI
	cache *cache.Cache
}

func New(medialiveapi medialiveiface.MediaLiveAPI) *MediaLive {
	return &MediaLive{
		MediaLiveAPI: medialiveapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaLive) BatchDelete(input *medialive.BatchDeleteInput) (*medialive.BatchDeleteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchDeleteOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchDelete(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchDeleteWithContext(ctx context.Context, input *medialive.BatchDeleteInput, opts ...request.Option) (*medialive.BatchDeleteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchDeleteOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchDeleteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchStart(input *medialive.BatchStartInput) (*medialive.BatchStartOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchStartOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchStart(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchStartWithContext(ctx context.Context, input *medialive.BatchStartInput, opts ...request.Option) (*medialive.BatchStartOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchStartOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchStartWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchStop(input *medialive.BatchStopInput) (*medialive.BatchStopOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchStopOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchStop(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchStopWithContext(ctx context.Context, input *medialive.BatchStopInput, opts ...request.Option) (*medialive.BatchStopOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchStopOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchStopWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchUpdateSchedule(input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchUpdateScheduleOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchUpdateSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) BatchUpdateScheduleWithContext(ctx context.Context, input *medialive.BatchUpdateScheduleInput, opts ...request.Option) (*medialive.BatchUpdateScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.BatchUpdateScheduleOutput), nil
	}
	out, err := c.MediaLiveAPI.BatchUpdateScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeChannel(input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeChannelOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeChannelWithContext(ctx context.Context, input *medialive.DescribeChannelInput, opts ...request.Option) (*medialive.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeChannelOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInput(input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputDevice(input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputDeviceOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputDeviceThumbnail(input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputDeviceThumbnailOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputDeviceThumbnail(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputDeviceThumbnailWithContext(ctx context.Context, input *medialive.DescribeInputDeviceThumbnailInput, opts ...request.Option) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputDeviceThumbnailOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputDeviceThumbnailWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputDeviceWithContext(ctx context.Context, input *medialive.DescribeInputDeviceInput, opts ...request.Option) (*medialive.DescribeInputDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputDeviceOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputSecurityGroup(input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputSecurityGroupOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputSecurityGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputSecurityGroupWithContext(ctx context.Context, input *medialive.DescribeInputSecurityGroupInput, opts ...request.Option) (*medialive.DescribeInputSecurityGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputSecurityGroupOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputSecurityGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeInputWithContext(ctx context.Context, input *medialive.DescribeInputInput, opts ...request.Option) (*medialive.DescribeInputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeInputOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeInputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeMultiplex(input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeMultiplexOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeMultiplex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeMultiplexProgram(input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeMultiplexProgramOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeMultiplexProgram(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeMultiplexProgramWithContext(ctx context.Context, input *medialive.DescribeMultiplexProgramInput, opts ...request.Option) (*medialive.DescribeMultiplexProgramOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeMultiplexProgramOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeMultiplexProgramWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeMultiplexWithContext(ctx context.Context, input *medialive.DescribeMultiplexInput, opts ...request.Option) (*medialive.DescribeMultiplexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeMultiplexOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeMultiplexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeOffering(input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeOfferingOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeOffering(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeOfferingWithContext(ctx context.Context, input *medialive.DescribeOfferingInput, opts ...request.Option) (*medialive.DescribeOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeOfferingOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeOfferingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeReservation(input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeReservationOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeReservationWithContext(ctx context.Context, input *medialive.DescribeReservationInput, opts ...request.Option) (*medialive.DescribeReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeReservationOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeSchedule(input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeScheduleOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) DescribeSchedulePages(input *medialive.DescribeScheduleInput, fn func(*medialive.DescribeScheduleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.DescribeScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.DescribeScheduleOutput{}
	fnCacher := func(out *medialive.DescribeScheduleOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.DescribeSchedulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) DescribeSchedulePagesWithContext(ctx context.Context, input *medialive.DescribeScheduleInput, fn func(*medialive.DescribeScheduleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.DescribeScheduleOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.DescribeScheduleOutput{}
	fnCacher := func(out *medialive.DescribeScheduleOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.DescribeSchedulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) DescribeScheduleWithContext(ctx context.Context, input *medialive.DescribeScheduleInput, opts ...request.Option) (*medialive.DescribeScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.DescribeScheduleOutput), nil
	}
	out, err := c.MediaLiveAPI.DescribeScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListChannels(input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListChannelsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListChannelsPages(input *medialive.ListChannelsInput, fn func(*medialive.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListChannelsOutput{}
	fnCacher := func(out *medialive.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListChannelsPagesWithContext(ctx context.Context, input *medialive.ListChannelsInput, fn func(*medialive.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListChannelsOutput{}
	fnCacher := func(out *medialive.ListChannelsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListChannelsWithContext(ctx context.Context, input *medialive.ListChannelsInput, opts ...request.Option) (*medialive.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListChannelsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputDeviceTransfers(input *medialive.ListInputDeviceTransfersInput) (*medialive.ListInputDeviceTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputDeviceTransfersOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputDeviceTransfers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputDeviceTransfersPages(input *medialive.ListInputDeviceTransfersInput, fn func(*medialive.ListInputDeviceTransfersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputDeviceTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputDeviceTransfersOutput{}
	fnCacher := func(out *medialive.ListInputDeviceTransfersOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputDeviceTransfersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputDeviceTransfersPagesWithContext(ctx context.Context, input *medialive.ListInputDeviceTransfersInput, fn func(*medialive.ListInputDeviceTransfersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputDeviceTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputDeviceTransfersOutput{}
	fnCacher := func(out *medialive.ListInputDeviceTransfersOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputDeviceTransfersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputDeviceTransfersWithContext(ctx context.Context, input *medialive.ListInputDeviceTransfersInput, opts ...request.Option) (*medialive.ListInputDeviceTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputDeviceTransfersOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputDeviceTransfersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputDevices(input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputDevicesOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputDevicesPages(input *medialive.ListInputDevicesInput, fn func(*medialive.ListInputDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputDevicesOutput{}
	fnCacher := func(out *medialive.ListInputDevicesOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputDevicesPagesWithContext(ctx context.Context, input *medialive.ListInputDevicesInput, fn func(*medialive.ListInputDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputDevicesOutput{}
	fnCacher := func(out *medialive.ListInputDevicesOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputDevicesWithContext(ctx context.Context, input *medialive.ListInputDevicesInput, opts ...request.Option) (*medialive.ListInputDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputDevicesOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputSecurityGroups(input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputSecurityGroupsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputSecurityGroupsPages(input *medialive.ListInputSecurityGroupsInput, fn func(*medialive.ListInputSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputSecurityGroupsOutput{}
	fnCacher := func(out *medialive.ListInputSecurityGroupsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputSecurityGroupsPagesWithContext(ctx context.Context, input *medialive.ListInputSecurityGroupsInput, fn func(*medialive.ListInputSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputSecurityGroupsOutput{}
	fnCacher := func(out *medialive.ListInputSecurityGroupsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputSecurityGroupsWithContext(ctx context.Context, input *medialive.ListInputSecurityGroupsInput, opts ...request.Option) (*medialive.ListInputSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputSecurityGroupsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputs(input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListInputsPages(input *medialive.ListInputsInput, fn func(*medialive.ListInputsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputsOutput{}
	fnCacher := func(out *medialive.ListInputsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputsPagesWithContext(ctx context.Context, input *medialive.ListInputsInput, fn func(*medialive.ListInputsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListInputsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListInputsOutput{}
	fnCacher := func(out *medialive.ListInputsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListInputsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListInputsWithContext(ctx context.Context, input *medialive.ListInputsInput, opts ...request.Option) (*medialive.ListInputsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListInputsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListInputsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListMultiplexPrograms(input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListMultiplexProgramsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListMultiplexPrograms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListMultiplexProgramsPages(input *medialive.ListMultiplexProgramsInput, fn func(*medialive.ListMultiplexProgramsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListMultiplexProgramsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListMultiplexProgramsOutput{}
	fnCacher := func(out *medialive.ListMultiplexProgramsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListMultiplexProgramsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListMultiplexProgramsPagesWithContext(ctx context.Context, input *medialive.ListMultiplexProgramsInput, fn func(*medialive.ListMultiplexProgramsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListMultiplexProgramsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListMultiplexProgramsOutput{}
	fnCacher := func(out *medialive.ListMultiplexProgramsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListMultiplexProgramsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListMultiplexProgramsWithContext(ctx context.Context, input *medialive.ListMultiplexProgramsInput, opts ...request.Option) (*medialive.ListMultiplexProgramsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListMultiplexProgramsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListMultiplexProgramsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListMultiplexes(input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListMultiplexesOutput), nil
	}
	out, err := c.MediaLiveAPI.ListMultiplexes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListMultiplexesPages(input *medialive.ListMultiplexesInput, fn func(*medialive.ListMultiplexesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListMultiplexesOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListMultiplexesOutput{}
	fnCacher := func(out *medialive.ListMultiplexesOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListMultiplexesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListMultiplexesPagesWithContext(ctx context.Context, input *medialive.ListMultiplexesInput, fn func(*medialive.ListMultiplexesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListMultiplexesOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListMultiplexesOutput{}
	fnCacher := func(out *medialive.ListMultiplexesOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListMultiplexesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListMultiplexesWithContext(ctx context.Context, input *medialive.ListMultiplexesInput, opts ...request.Option) (*medialive.ListMultiplexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListMultiplexesOutput), nil
	}
	out, err := c.MediaLiveAPI.ListMultiplexesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListOfferings(input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListOfferingsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListOfferingsPages(input *medialive.ListOfferingsInput, fn func(*medialive.ListOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListOfferingsOutput{}
	fnCacher := func(out *medialive.ListOfferingsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListOfferingsPagesWithContext(ctx context.Context, input *medialive.ListOfferingsInput, fn func(*medialive.ListOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListOfferingsOutput{}
	fnCacher := func(out *medialive.ListOfferingsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListOfferingsWithContext(ctx context.Context, input *medialive.ListOfferingsInput, opts ...request.Option) (*medialive.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListOfferingsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListReservations(input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListReservationsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListReservationsPages(input *medialive.ListReservationsInput, fn func(*medialive.ListReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListReservationsOutput{}
	fnCacher := func(out *medialive.ListReservationsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListReservationsPagesWithContext(ctx context.Context, input *medialive.ListReservationsInput, fn func(*medialive.ListReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*medialive.ListReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &medialive.ListReservationsOutput{}
	fnCacher := func(out *medialive.ListReservationsOutput, more bool) bool {
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
	if err := c.MediaLiveAPI.ListReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaLive) ListReservationsWithContext(ctx context.Context, input *medialive.ListReservationsInput, opts ...request.Option) (*medialive.ListReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListReservationsOutput), nil
	}
	out, err := c.MediaLiveAPI.ListReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListTagsForResource(input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaLiveAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaLive) ListTagsForResourceWithContext(ctx context.Context, input *medialive.ListTagsForResourceInput, opts ...request.Option) (*medialive.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*medialive.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaLiveAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
