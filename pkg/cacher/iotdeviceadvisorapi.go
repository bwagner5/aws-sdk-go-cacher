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
	"github.com/aws/aws-sdk-go/service/iotdeviceadvisor"
	"github.com/aws/aws-sdk-go/service/iotdeviceadvisor/iotdeviceadvisoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTDeviceAdvisor struct {
	iotdeviceadvisoriface.IoTDeviceAdvisorAPI
	cache *cache.Cache
}

func NewIoTDeviceAdvisor(iotdeviceadvisorapi iotdeviceadvisoriface.IoTDeviceAdvisorAPI) *IoTDeviceAdvisor {
	return &IoTDeviceAdvisor{
		IoTDeviceAdvisorAPI: iotdeviceadvisorapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTDeviceAdvisor) GetEndpoint(input *iotdeviceadvisor.GetEndpointInput) (*iotdeviceadvisor.GetEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetEndpointOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetEndpointWithContext(ctx context.Context, input *iotdeviceadvisor.GetEndpointInput, opts ...request.Option) (*iotdeviceadvisor.GetEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetEndpointOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteDefinition(input *iotdeviceadvisor.GetSuiteDefinitionInput) (*iotdeviceadvisor.GetSuiteDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteDefinitionOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteDefinitionWithContext(ctx context.Context, input *iotdeviceadvisor.GetSuiteDefinitionInput, opts ...request.Option) (*iotdeviceadvisor.GetSuiteDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteDefinitionOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteRun(input *iotdeviceadvisor.GetSuiteRunInput) (*iotdeviceadvisor.GetSuiteRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteRunOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteRunReport(input *iotdeviceadvisor.GetSuiteRunReportInput) (*iotdeviceadvisor.GetSuiteRunReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteRunReportOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteRunReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteRunReportWithContext(ctx context.Context, input *iotdeviceadvisor.GetSuiteRunReportInput, opts ...request.Option) (*iotdeviceadvisor.GetSuiteRunReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteRunReportOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteRunReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) GetSuiteRunWithContext(ctx context.Context, input *iotdeviceadvisor.GetSuiteRunInput, opts ...request.Option) (*iotdeviceadvisor.GetSuiteRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.GetSuiteRunOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.GetSuiteRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListSuiteDefinitions(input *iotdeviceadvisor.ListSuiteDefinitionsInput) (*iotdeviceadvisor.ListSuiteDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListSuiteDefinitionsOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListSuiteDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListSuiteDefinitionsPages(input *iotdeviceadvisor.ListSuiteDefinitionsInput, fn func(*iotdeviceadvisor.ListSuiteDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdeviceadvisor.ListSuiteDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotdeviceadvisor.ListSuiteDefinitionsOutput{}
	fnCacher := func(out *iotdeviceadvisor.ListSuiteDefinitionsOutput, more bool) bool {
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
	if err := c.IoTDeviceAdvisorAPI.ListSuiteDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDeviceAdvisor) ListSuiteDefinitionsPagesWithContext(ctx context.Context, input *iotdeviceadvisor.ListSuiteDefinitionsInput, fn func(*iotdeviceadvisor.ListSuiteDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdeviceadvisor.ListSuiteDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotdeviceadvisor.ListSuiteDefinitionsOutput{}
	fnCacher := func(out *iotdeviceadvisor.ListSuiteDefinitionsOutput, more bool) bool {
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
	if err := c.IoTDeviceAdvisorAPI.ListSuiteDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDeviceAdvisor) ListSuiteDefinitionsWithContext(ctx context.Context, input *iotdeviceadvisor.ListSuiteDefinitionsInput, opts ...request.Option) (*iotdeviceadvisor.ListSuiteDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListSuiteDefinitionsOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListSuiteDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListSuiteRuns(input *iotdeviceadvisor.ListSuiteRunsInput) (*iotdeviceadvisor.ListSuiteRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListSuiteRunsOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListSuiteRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListSuiteRunsPages(input *iotdeviceadvisor.ListSuiteRunsInput, fn func(*iotdeviceadvisor.ListSuiteRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdeviceadvisor.ListSuiteRunsOutput), false)
		return nil
	}
	cachable := true
	output := &iotdeviceadvisor.ListSuiteRunsOutput{}
	fnCacher := func(out *iotdeviceadvisor.ListSuiteRunsOutput, more bool) bool {
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
	if err := c.IoTDeviceAdvisorAPI.ListSuiteRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDeviceAdvisor) ListSuiteRunsPagesWithContext(ctx context.Context, input *iotdeviceadvisor.ListSuiteRunsInput, fn func(*iotdeviceadvisor.ListSuiteRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotdeviceadvisor.ListSuiteRunsOutput), false)
		return nil
	}
	cachable := true
	output := &iotdeviceadvisor.ListSuiteRunsOutput{}
	fnCacher := func(out *iotdeviceadvisor.ListSuiteRunsOutput, more bool) bool {
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
	if err := c.IoTDeviceAdvisorAPI.ListSuiteRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTDeviceAdvisor) ListSuiteRunsWithContext(ctx context.Context, input *iotdeviceadvisor.ListSuiteRunsInput, opts ...request.Option) (*iotdeviceadvisor.ListSuiteRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListSuiteRunsOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListSuiteRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListTagsForResource(input *iotdeviceadvisor.ListTagsForResourceInput) (*iotdeviceadvisor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTDeviceAdvisor) ListTagsForResourceWithContext(ctx context.Context, input *iotdeviceadvisor.ListTagsForResourceInput, opts ...request.Option) (*iotdeviceadvisor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotdeviceadvisor.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTDeviceAdvisorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
