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
package iot1clickdevicesservicecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice"
	"github.com/aws/aws-sdk-go/service/iot1clickdevicesservice/iot1clickdevicesserviceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoT1ClickDevicesService struct {
	iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI
	cache *cache.Cache
}

func New(iot1clickdevicesserviceapi iot1clickdevicesserviceiface.IoT1ClickDevicesServiceAPI) *IoT1ClickDevicesService {
	return &IoT1ClickDevicesService{
		IoT1ClickDevicesServiceAPI: iot1clickdevicesserviceapi,
		cache:                      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoT1ClickDevicesService) DescribeDevice(input *iot1clickdevicesservice.DescribeDeviceInput) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.DescribeDeviceOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.DescribeDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) DescribeDeviceWithContext(ctx context.Context, input *iot1clickdevicesservice.DescribeDeviceInput, opts ...request.Option) (*iot1clickdevicesservice.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.DescribeDeviceOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.DescribeDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) GetDeviceMethods(input *iot1clickdevicesservice.GetDeviceMethodsInput) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.GetDeviceMethodsOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.GetDeviceMethods(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) GetDeviceMethodsWithContext(ctx context.Context, input *iot1clickdevicesservice.GetDeviceMethodsInput, opts ...request.Option) (*iot1clickdevicesservice.GetDeviceMethodsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.GetDeviceMethodsOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.GetDeviceMethodsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListDeviceEvents(input *iot1clickdevicesservice.ListDeviceEventsInput) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListDeviceEventsOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListDeviceEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListDeviceEventsWithContext(ctx context.Context, input *iot1clickdevicesservice.ListDeviceEventsInput, opts ...request.Option) (*iot1clickdevicesservice.ListDeviceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListDeviceEventsOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListDeviceEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListDevices(input *iot1clickdevicesservice.ListDevicesInput) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListDevicesOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListDevicesWithContext(ctx context.Context, input *iot1clickdevicesservice.ListDevicesInput, opts ...request.Option) (*iot1clickdevicesservice.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListDevicesOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListTagsForResource(input *iot1clickdevicesservice.ListTagsForResourceInput) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListTagsForResourceOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoT1ClickDevicesService) ListTagsForResourceWithContext(ctx context.Context, input *iot1clickdevicesservice.ListTagsForResourceInput, opts ...request.Option) (*iot1clickdevicesservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iot1clickdevicesservice.ListTagsForResourceOutput), nil
	}
	out, err := c.IoT1ClickDevicesServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
