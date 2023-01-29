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
	"github.com/aws/aws-sdk-go/service/devicefarm"
	"github.com/aws/aws-sdk-go/service/devicefarm/devicefarmiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DeviceFarm struct {
	devicefarmiface.DeviceFarmAPI
	cache *cache.Cache
}

func NewDeviceFarm(devicefarmapi devicefarmiface.DeviceFarmAPI) *DeviceFarm {
	return &DeviceFarm{
		DeviceFarmAPI: devicefarmapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DeviceFarm) GetAccountSettings(input *devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetAccountSettingsOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetAccountSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetAccountSettingsWithContext(ctx context.Context, input *devicefarm.GetAccountSettingsInput, opts ...request.Option) (*devicefarm.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetAccountSettingsOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetAccountSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDevice(input *devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDeviceOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDeviceInstance(input *devicefarm.GetDeviceInstanceInput) (*devicefarm.GetDeviceInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDeviceInstanceOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDeviceInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDeviceInstanceWithContext(ctx context.Context, input *devicefarm.GetDeviceInstanceInput, opts ...request.Option) (*devicefarm.GetDeviceInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDeviceInstanceOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDeviceInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDevicePool(input *devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDevicePoolOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDevicePool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDevicePoolCompatibility(input *devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDevicePoolCompatibilityOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDevicePoolCompatibility(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDevicePoolCompatibilityWithContext(ctx context.Context, input *devicefarm.GetDevicePoolCompatibilityInput, opts ...request.Option) (*devicefarm.GetDevicePoolCompatibilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDevicePoolCompatibilityOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDevicePoolCompatibilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDevicePoolWithContext(ctx context.Context, input *devicefarm.GetDevicePoolInput, opts ...request.Option) (*devicefarm.GetDevicePoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDevicePoolOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDevicePoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetDeviceWithContext(ctx context.Context, input *devicefarm.GetDeviceInput, opts ...request.Option) (*devicefarm.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetDeviceOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetInstanceProfile(input *devicefarm.GetInstanceProfileInput) (*devicefarm.GetInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetInstanceProfileOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetInstanceProfileWithContext(ctx context.Context, input *devicefarm.GetInstanceProfileInput, opts ...request.Option) (*devicefarm.GetInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetInstanceProfileOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetJob(input *devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetJobOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetJobWithContext(ctx context.Context, input *devicefarm.GetJobInput, opts ...request.Option) (*devicefarm.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetJobOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetNetworkProfile(input *devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetNetworkProfileOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetNetworkProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetNetworkProfileWithContext(ctx context.Context, input *devicefarm.GetNetworkProfileInput, opts ...request.Option) (*devicefarm.GetNetworkProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetNetworkProfileOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetNetworkProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetOfferingStatus(input *devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetOfferingStatusOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetOfferingStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetOfferingStatusPages(input *devicefarm.GetOfferingStatusInput, fn func(*devicefarm.GetOfferingStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.GetOfferingStatusOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.GetOfferingStatusOutput{}
	fnCacher := func(out *devicefarm.GetOfferingStatusOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.GetOfferingStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) GetOfferingStatusPagesWithContext(ctx context.Context, input *devicefarm.GetOfferingStatusInput, fn func(*devicefarm.GetOfferingStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.GetOfferingStatusOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.GetOfferingStatusOutput{}
	fnCacher := func(out *devicefarm.GetOfferingStatusOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.GetOfferingStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) GetOfferingStatusWithContext(ctx context.Context, input *devicefarm.GetOfferingStatusInput, opts ...request.Option) (*devicefarm.GetOfferingStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetOfferingStatusOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetOfferingStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetProject(input *devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetProjectOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetProjectWithContext(ctx context.Context, input *devicefarm.GetProjectInput, opts ...request.Option) (*devicefarm.GetProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetProjectOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetRemoteAccessSession(input *devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetRemoteAccessSessionOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetRemoteAccessSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetRemoteAccessSessionWithContext(ctx context.Context, input *devicefarm.GetRemoteAccessSessionInput, opts ...request.Option) (*devicefarm.GetRemoteAccessSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetRemoteAccessSessionOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetRemoteAccessSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetRun(input *devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetRunOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetRunWithContext(ctx context.Context, input *devicefarm.GetRunInput, opts ...request.Option) (*devicefarm.GetRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetRunOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetSuite(input *devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetSuiteOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetSuite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetSuiteWithContext(ctx context.Context, input *devicefarm.GetSuiteInput, opts ...request.Option) (*devicefarm.GetSuiteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetSuiteOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetSuiteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTest(input *devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTestGridProject(input *devicefarm.GetTestGridProjectInput) (*devicefarm.GetTestGridProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestGridProjectOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTestGridProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTestGridProjectWithContext(ctx context.Context, input *devicefarm.GetTestGridProjectInput, opts ...request.Option) (*devicefarm.GetTestGridProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestGridProjectOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTestGridProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTestGridSession(input *devicefarm.GetTestGridSessionInput) (*devicefarm.GetTestGridSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestGridSessionOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTestGridSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTestGridSessionWithContext(ctx context.Context, input *devicefarm.GetTestGridSessionInput, opts ...request.Option) (*devicefarm.GetTestGridSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestGridSessionOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTestGridSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetTestWithContext(ctx context.Context, input *devicefarm.GetTestInput, opts ...request.Option) (*devicefarm.GetTestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetTestOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetTestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetUpload(input *devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetUploadOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetUpload(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetUploadWithContext(ctx context.Context, input *devicefarm.GetUploadInput, opts ...request.Option) (*devicefarm.GetUploadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetUploadOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetUploadWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetVPCEConfiguration(input *devicefarm.GetVPCEConfigurationInput) (*devicefarm.GetVPCEConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetVPCEConfigurationOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetVPCEConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) GetVPCEConfigurationWithContext(ctx context.Context, input *devicefarm.GetVPCEConfigurationInput, opts ...request.Option) (*devicefarm.GetVPCEConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.GetVPCEConfigurationOutput), nil
	}
	out, err := c.DeviceFarmAPI.GetVPCEConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListArtifacts(input *devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListArtifactsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListArtifactsPages(input *devicefarm.ListArtifactsInput, fn func(*devicefarm.ListArtifactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListArtifactsOutput{}
	fnCacher := func(out *devicefarm.ListArtifactsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListArtifactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListArtifactsPagesWithContext(ctx context.Context, input *devicefarm.ListArtifactsInput, fn func(*devicefarm.ListArtifactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListArtifactsOutput{}
	fnCacher := func(out *devicefarm.ListArtifactsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListArtifactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListArtifactsWithContext(ctx context.Context, input *devicefarm.ListArtifactsInput, opts ...request.Option) (*devicefarm.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListArtifactsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDeviceInstances(input *devicefarm.ListDeviceInstancesInput) (*devicefarm.ListDeviceInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDeviceInstancesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDeviceInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDeviceInstancesWithContext(ctx context.Context, input *devicefarm.ListDeviceInstancesInput, opts ...request.Option) (*devicefarm.ListDeviceInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDeviceInstancesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDeviceInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDevicePools(input *devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDevicePoolsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDevicePools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDevicePoolsPages(input *devicefarm.ListDevicePoolsInput, fn func(*devicefarm.ListDevicePoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListDevicePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListDevicePoolsOutput{}
	fnCacher := func(out *devicefarm.ListDevicePoolsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListDevicePoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListDevicePoolsPagesWithContext(ctx context.Context, input *devicefarm.ListDevicePoolsInput, fn func(*devicefarm.ListDevicePoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListDevicePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListDevicePoolsOutput{}
	fnCacher := func(out *devicefarm.ListDevicePoolsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListDevicePoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListDevicePoolsWithContext(ctx context.Context, input *devicefarm.ListDevicePoolsInput, opts ...request.Option) (*devicefarm.ListDevicePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDevicePoolsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDevicePoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDevices(input *devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDevicesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListDevicesPages(input *devicefarm.ListDevicesInput, fn func(*devicefarm.ListDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListDevicesOutput{}
	fnCacher := func(out *devicefarm.ListDevicesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListDevicesPagesWithContext(ctx context.Context, input *devicefarm.ListDevicesInput, fn func(*devicefarm.ListDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListDevicesOutput{}
	fnCacher := func(out *devicefarm.ListDevicesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListDevicesWithContext(ctx context.Context, input *devicefarm.ListDevicesInput, opts ...request.Option) (*devicefarm.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListDevicesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListInstanceProfiles(input *devicefarm.ListInstanceProfilesInput) (*devicefarm.ListInstanceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListInstanceProfilesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListInstanceProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListInstanceProfilesWithContext(ctx context.Context, input *devicefarm.ListInstanceProfilesInput, opts ...request.Option) (*devicefarm.ListInstanceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListInstanceProfilesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListInstanceProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListJobs(input *devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListJobsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListJobsPages(input *devicefarm.ListJobsInput, fn func(*devicefarm.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListJobsOutput{}
	fnCacher := func(out *devicefarm.ListJobsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListJobsPagesWithContext(ctx context.Context, input *devicefarm.ListJobsInput, fn func(*devicefarm.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListJobsOutput{}
	fnCacher := func(out *devicefarm.ListJobsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListJobsWithContext(ctx context.Context, input *devicefarm.ListJobsInput, opts ...request.Option) (*devicefarm.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListJobsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListNetworkProfiles(input *devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListNetworkProfilesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListNetworkProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListNetworkProfilesWithContext(ctx context.Context, input *devicefarm.ListNetworkProfilesInput, opts ...request.Option) (*devicefarm.ListNetworkProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListNetworkProfilesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListNetworkProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferingPromotions(input *devicefarm.ListOfferingPromotionsInput) (*devicefarm.ListOfferingPromotionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingPromotionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferingPromotions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferingPromotionsWithContext(ctx context.Context, input *devicefarm.ListOfferingPromotionsInput, opts ...request.Option) (*devicefarm.ListOfferingPromotionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingPromotionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferingPromotionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferingTransactions(input *devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingTransactionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferingTransactions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferingTransactionsPages(input *devicefarm.ListOfferingTransactionsInput, fn func(*devicefarm.ListOfferingTransactionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListOfferingTransactionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListOfferingTransactionsOutput{}
	fnCacher := func(out *devicefarm.ListOfferingTransactionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListOfferingTransactionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListOfferingTransactionsPagesWithContext(ctx context.Context, input *devicefarm.ListOfferingTransactionsInput, fn func(*devicefarm.ListOfferingTransactionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListOfferingTransactionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListOfferingTransactionsOutput{}
	fnCacher := func(out *devicefarm.ListOfferingTransactionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListOfferingTransactionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListOfferingTransactionsWithContext(ctx context.Context, input *devicefarm.ListOfferingTransactionsInput, opts ...request.Option) (*devicefarm.ListOfferingTransactionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingTransactionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferingTransactionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferings(input *devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListOfferingsPages(input *devicefarm.ListOfferingsInput, fn func(*devicefarm.ListOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListOfferingsOutput{}
	fnCacher := func(out *devicefarm.ListOfferingsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListOfferingsPagesWithContext(ctx context.Context, input *devicefarm.ListOfferingsInput, fn func(*devicefarm.ListOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListOfferingsOutput{}
	fnCacher := func(out *devicefarm.ListOfferingsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListOfferingsWithContext(ctx context.Context, input *devicefarm.ListOfferingsInput, opts ...request.Option) (*devicefarm.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListOfferingsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListProjects(input *devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListProjectsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListProjectsPages(input *devicefarm.ListProjectsInput, fn func(*devicefarm.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListProjectsOutput{}
	fnCacher := func(out *devicefarm.ListProjectsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListProjectsPagesWithContext(ctx context.Context, input *devicefarm.ListProjectsInput, fn func(*devicefarm.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListProjectsOutput{}
	fnCacher := func(out *devicefarm.ListProjectsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListProjectsWithContext(ctx context.Context, input *devicefarm.ListProjectsInput, opts ...request.Option) (*devicefarm.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListProjectsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListRemoteAccessSessions(input *devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListRemoteAccessSessionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListRemoteAccessSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListRemoteAccessSessionsWithContext(ctx context.Context, input *devicefarm.ListRemoteAccessSessionsInput, opts ...request.Option) (*devicefarm.ListRemoteAccessSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListRemoteAccessSessionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListRemoteAccessSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListRuns(input *devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListRunsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListRunsPages(input *devicefarm.ListRunsInput, fn func(*devicefarm.ListRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListRunsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListRunsOutput{}
	fnCacher := func(out *devicefarm.ListRunsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListRunsPagesWithContext(ctx context.Context, input *devicefarm.ListRunsInput, fn func(*devicefarm.ListRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListRunsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListRunsOutput{}
	fnCacher := func(out *devicefarm.ListRunsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListRunsWithContext(ctx context.Context, input *devicefarm.ListRunsInput, opts ...request.Option) (*devicefarm.ListRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListRunsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListSamples(input *devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListSamplesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListSamples(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListSamplesPages(input *devicefarm.ListSamplesInput, fn func(*devicefarm.ListSamplesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListSamplesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListSamplesOutput{}
	fnCacher := func(out *devicefarm.ListSamplesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListSamplesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListSamplesPagesWithContext(ctx context.Context, input *devicefarm.ListSamplesInput, fn func(*devicefarm.ListSamplesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListSamplesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListSamplesOutput{}
	fnCacher := func(out *devicefarm.ListSamplesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListSamplesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListSamplesWithContext(ctx context.Context, input *devicefarm.ListSamplesInput, opts ...request.Option) (*devicefarm.ListSamplesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListSamplesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListSamplesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListSuites(input *devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListSuitesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListSuites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListSuitesPages(input *devicefarm.ListSuitesInput, fn func(*devicefarm.ListSuitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListSuitesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListSuitesOutput{}
	fnCacher := func(out *devicefarm.ListSuitesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListSuitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListSuitesPagesWithContext(ctx context.Context, input *devicefarm.ListSuitesInput, fn func(*devicefarm.ListSuitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListSuitesOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListSuitesOutput{}
	fnCacher := func(out *devicefarm.ListSuitesOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListSuitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListSuitesWithContext(ctx context.Context, input *devicefarm.ListSuitesInput, opts ...request.Option) (*devicefarm.ListSuitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListSuitesOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListSuitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTagsForResource(input *devicefarm.ListTagsForResourceInput) (*devicefarm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTagsForResourceOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTagsForResourceWithContext(ctx context.Context, input *devicefarm.ListTagsForResourceInput, opts ...request.Option) (*devicefarm.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTagsForResourceOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridProjects(input *devicefarm.ListTestGridProjectsInput) (*devicefarm.ListTestGridProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridProjectsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridProjectsPages(input *devicefarm.ListTestGridProjectsInput, fn func(*devicefarm.ListTestGridProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridProjectsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridProjectsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridProjectsPagesWithContext(ctx context.Context, input *devicefarm.ListTestGridProjectsInput, fn func(*devicefarm.ListTestGridProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridProjectsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridProjectsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridProjectsWithContext(ctx context.Context, input *devicefarm.ListTestGridProjectsInput, opts ...request.Option) (*devicefarm.ListTestGridProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridProjectsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessionActions(input *devicefarm.ListTestGridSessionActionsInput) (*devicefarm.ListTestGridSessionActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionActionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessionActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessionActionsPages(input *devicefarm.ListTestGridSessionActionsInput, fn func(*devicefarm.ListTestGridSessionActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionActionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionActionsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionActionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionActionsPagesWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionActionsInput, fn func(*devicefarm.ListTestGridSessionActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionActionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionActionsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionActionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionActionsWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionActionsInput, opts ...request.Option) (*devicefarm.ListTestGridSessionActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionActionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessionActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessionArtifacts(input *devicefarm.ListTestGridSessionArtifactsInput) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionArtifactsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessionArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessionArtifactsPages(input *devicefarm.ListTestGridSessionArtifactsInput, fn func(*devicefarm.ListTestGridSessionArtifactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionArtifactsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionArtifactsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionArtifactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionArtifactsPagesWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionArtifactsInput, fn func(*devicefarm.ListTestGridSessionArtifactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionArtifactsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionArtifactsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionArtifactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionArtifactsWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionArtifactsInput, opts ...request.Option) (*devicefarm.ListTestGridSessionArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionArtifactsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessionArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessions(input *devicefarm.ListTestGridSessionsInput) (*devicefarm.ListTestGridSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestGridSessionsPages(input *devicefarm.ListTestGridSessionsInput, fn func(*devicefarm.ListTestGridSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionsPagesWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionsInput, fn func(*devicefarm.ListTestGridSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestGridSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestGridSessionsOutput{}
	fnCacher := func(out *devicefarm.ListTestGridSessionsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestGridSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestGridSessionsWithContext(ctx context.Context, input *devicefarm.ListTestGridSessionsInput, opts ...request.Option) (*devicefarm.ListTestGridSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestGridSessionsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestGridSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTests(input *devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListTestsPages(input *devicefarm.ListTestsInput, fn func(*devicefarm.ListTestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestsOutput{}
	fnCacher := func(out *devicefarm.ListTestsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestsPagesWithContext(ctx context.Context, input *devicefarm.ListTestsInput, fn func(*devicefarm.ListTestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListTestsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListTestsOutput{}
	fnCacher := func(out *devicefarm.ListTestsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListTestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListTestsWithContext(ctx context.Context, input *devicefarm.ListTestsInput, opts ...request.Option) (*devicefarm.ListTestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListTestsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListTestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListUniqueProblems(input *devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListUniqueProblemsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListUniqueProblems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListUniqueProblemsPages(input *devicefarm.ListUniqueProblemsInput, fn func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListUniqueProblemsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListUniqueProblemsOutput{}
	fnCacher := func(out *devicefarm.ListUniqueProblemsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListUniqueProblemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListUniqueProblemsPagesWithContext(ctx context.Context, input *devicefarm.ListUniqueProblemsInput, fn func(*devicefarm.ListUniqueProblemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListUniqueProblemsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListUniqueProblemsOutput{}
	fnCacher := func(out *devicefarm.ListUniqueProblemsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListUniqueProblemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListUniqueProblemsWithContext(ctx context.Context, input *devicefarm.ListUniqueProblemsInput, opts ...request.Option) (*devicefarm.ListUniqueProblemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListUniqueProblemsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListUniqueProblemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListUploads(input *devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListUploadsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListUploads(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListUploadsPages(input *devicefarm.ListUploadsInput, fn func(*devicefarm.ListUploadsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListUploadsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListUploadsOutput{}
	fnCacher := func(out *devicefarm.ListUploadsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListUploadsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListUploadsPagesWithContext(ctx context.Context, input *devicefarm.ListUploadsInput, fn func(*devicefarm.ListUploadsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devicefarm.ListUploadsOutput), false)
		return nil
	}
	cachable := true
	output := &devicefarm.ListUploadsOutput{}
	fnCacher := func(out *devicefarm.ListUploadsOutput, more bool) bool {
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
	if err := c.DeviceFarmAPI.ListUploadsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DeviceFarm) ListUploadsWithContext(ctx context.Context, input *devicefarm.ListUploadsInput, opts ...request.Option) (*devicefarm.ListUploadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListUploadsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListUploadsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListVPCEConfigurations(input *devicefarm.ListVPCEConfigurationsInput) (*devicefarm.ListVPCEConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListVPCEConfigurationsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListVPCEConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DeviceFarm) ListVPCEConfigurationsWithContext(ctx context.Context, input *devicefarm.ListVPCEConfigurationsInput, opts ...request.Option) (*devicefarm.ListVPCEConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devicefarm.ListVPCEConfigurationsOutput), nil
	}
	out, err := c.DeviceFarmAPI.ListVPCEConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
