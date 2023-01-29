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
	"github.com/aws/aws-sdk-go/service/snowdevicemanagement"
	"github.com/aws/aws-sdk-go/service/snowdevicemanagement/snowdevicemanagementiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SnowDeviceManagement struct {
	snowdevicemanagementiface.SnowDeviceManagementAPI
	cache *cache.Cache
}

func NewSnowDeviceManagement(snowdevicemanagementapi snowdevicemanagementiface.SnowDeviceManagementAPI) *SnowDeviceManagement {
	return &SnowDeviceManagement{
		SnowDeviceManagementAPI: snowdevicemanagementapi,
		cache:                   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SnowDeviceManagement) DescribeDevice(input *snowdevicemanagement.DescribeDeviceInput) (*snowdevicemanagement.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeDeviceOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeDeviceEc2Instances(input *snowdevicemanagement.DescribeDeviceEc2InstancesInput) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeDeviceEc2InstancesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeDeviceEc2Instances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeDeviceEc2InstancesWithContext(ctx context.Context, input *snowdevicemanagement.DescribeDeviceEc2InstancesInput, opts ...request.Option) (*snowdevicemanagement.DescribeDeviceEc2InstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeDeviceEc2InstancesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeDeviceEc2InstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeDeviceWithContext(ctx context.Context, input *snowdevicemanagement.DescribeDeviceInput, opts ...request.Option) (*snowdevicemanagement.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeDeviceOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeExecution(input *snowdevicemanagement.DescribeExecutionInput) (*snowdevicemanagement.DescribeExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeExecutionOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeExecutionWithContext(ctx context.Context, input *snowdevicemanagement.DescribeExecutionInput, opts ...request.Option) (*snowdevicemanagement.DescribeExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeExecutionOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeTask(input *snowdevicemanagement.DescribeTaskInput) (*snowdevicemanagement.DescribeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeTaskOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) DescribeTaskWithContext(ctx context.Context, input *snowdevicemanagement.DescribeTaskInput, opts ...request.Option) (*snowdevicemanagement.DescribeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.DescribeTaskOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.DescribeTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListDeviceResources(input *snowdevicemanagement.ListDeviceResourcesInput) (*snowdevicemanagement.ListDeviceResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListDeviceResourcesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListDeviceResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListDeviceResourcesPages(input *snowdevicemanagement.ListDeviceResourcesInput, fn func(*snowdevicemanagement.ListDeviceResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListDeviceResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListDeviceResourcesOutput{}
	fnCacher := func(out *snowdevicemanagement.ListDeviceResourcesOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListDeviceResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListDeviceResourcesPagesWithContext(ctx context.Context, input *snowdevicemanagement.ListDeviceResourcesInput, fn func(*snowdevicemanagement.ListDeviceResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListDeviceResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListDeviceResourcesOutput{}
	fnCacher := func(out *snowdevicemanagement.ListDeviceResourcesOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListDeviceResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListDeviceResourcesWithContext(ctx context.Context, input *snowdevicemanagement.ListDeviceResourcesInput, opts ...request.Option) (*snowdevicemanagement.ListDeviceResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListDeviceResourcesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListDeviceResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListDevices(input *snowdevicemanagement.ListDevicesInput) (*snowdevicemanagement.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListDevicesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListDevicesPages(input *snowdevicemanagement.ListDevicesInput, fn func(*snowdevicemanagement.ListDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListDevicesOutput{}
	fnCacher := func(out *snowdevicemanagement.ListDevicesOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListDevicesPagesWithContext(ctx context.Context, input *snowdevicemanagement.ListDevicesInput, fn func(*snowdevicemanagement.ListDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListDevicesOutput{}
	fnCacher := func(out *snowdevicemanagement.ListDevicesOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListDevicesWithContext(ctx context.Context, input *snowdevicemanagement.ListDevicesInput, opts ...request.Option) (*snowdevicemanagement.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListDevicesOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListExecutions(input *snowdevicemanagement.ListExecutionsInput) (*snowdevicemanagement.ListExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListExecutionsOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListExecutionsPages(input *snowdevicemanagement.ListExecutionsInput, fn func(*snowdevicemanagement.ListExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListExecutionsOutput{}
	fnCacher := func(out *snowdevicemanagement.ListExecutionsOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListExecutionsPagesWithContext(ctx context.Context, input *snowdevicemanagement.ListExecutionsInput, fn func(*snowdevicemanagement.ListExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListExecutionsOutput{}
	fnCacher := func(out *snowdevicemanagement.ListExecutionsOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListExecutionsWithContext(ctx context.Context, input *snowdevicemanagement.ListExecutionsInput, opts ...request.Option) (*snowdevicemanagement.ListExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListExecutionsOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListTagsForResource(input *snowdevicemanagement.ListTagsForResourceInput) (*snowdevicemanagement.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListTagsForResourceOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListTagsForResourceWithContext(ctx context.Context, input *snowdevicemanagement.ListTagsForResourceInput, opts ...request.Option) (*snowdevicemanagement.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListTagsForResourceOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListTasks(input *snowdevicemanagement.ListTasksInput) (*snowdevicemanagement.ListTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListTasksOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SnowDeviceManagement) ListTasksPages(input *snowdevicemanagement.ListTasksInput, fn func(*snowdevicemanagement.ListTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListTasksOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListTasksOutput{}
	fnCacher := func(out *snowdevicemanagement.ListTasksOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListTasksPagesWithContext(ctx context.Context, input *snowdevicemanagement.ListTasksInput, fn func(*snowdevicemanagement.ListTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*snowdevicemanagement.ListTasksOutput), false)
		return nil
	}
	cachable := true
	output := &snowdevicemanagement.ListTasksOutput{}
	fnCacher := func(out *snowdevicemanagement.ListTasksOutput, more bool) bool {
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
	if err := c.SnowDeviceManagementAPI.ListTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SnowDeviceManagement) ListTasksWithContext(ctx context.Context, input *snowdevicemanagement.ListTasksInput, opts ...request.Option) (*snowdevicemanagement.ListTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*snowdevicemanagement.ListTasksOutput), nil
	}
	out, err := c.SnowDeviceManagementAPI.ListTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
