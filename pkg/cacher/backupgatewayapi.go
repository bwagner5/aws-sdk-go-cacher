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
	"github.com/aws/aws-sdk-go/service/backupgateway"
	"github.com/aws/aws-sdk-go/service/backupgateway/backupgatewayiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type BackupGateway struct {
	backupgatewayiface.BackupGatewayAPI
	cache *cache.Cache
}

func NewBackupGateway(backupgatewayapi backupgatewayiface.BackupGatewayAPI) *BackupGateway {
	return &BackupGateway{
		BackupGatewayAPI: backupgatewayapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *BackupGateway) GetBandwidthRateLimitSchedule(input *backupgateway.GetBandwidthRateLimitScheduleInput) (*backupgateway.GetBandwidthRateLimitScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetBandwidthRateLimitScheduleOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetBandwidthRateLimitSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetBandwidthRateLimitScheduleWithContext(ctx context.Context, input *backupgateway.GetBandwidthRateLimitScheduleInput, opts ...request.Option) (*backupgateway.GetBandwidthRateLimitScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetBandwidthRateLimitScheduleOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetBandwidthRateLimitScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetGateway(input *backupgateway.GetGatewayInput) (*backupgateway.GetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetGatewayOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetGatewayWithContext(ctx context.Context, input *backupgateway.GetGatewayInput, opts ...request.Option) (*backupgateway.GetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetGatewayOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetHypervisor(input *backupgateway.GetHypervisorInput) (*backupgateway.GetHypervisorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetHypervisorOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetHypervisor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetHypervisorPropertyMappings(input *backupgateway.GetHypervisorPropertyMappingsInput) (*backupgateway.GetHypervisorPropertyMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetHypervisorPropertyMappingsOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetHypervisorPropertyMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetHypervisorPropertyMappingsWithContext(ctx context.Context, input *backupgateway.GetHypervisorPropertyMappingsInput, opts ...request.Option) (*backupgateway.GetHypervisorPropertyMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetHypervisorPropertyMappingsOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetHypervisorPropertyMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetHypervisorWithContext(ctx context.Context, input *backupgateway.GetHypervisorInput, opts ...request.Option) (*backupgateway.GetHypervisorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetHypervisorOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetHypervisorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetVirtualMachine(input *backupgateway.GetVirtualMachineInput) (*backupgateway.GetVirtualMachineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetVirtualMachineOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetVirtualMachine(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) GetVirtualMachineWithContext(ctx context.Context, input *backupgateway.GetVirtualMachineInput, opts ...request.Option) (*backupgateway.GetVirtualMachineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.GetVirtualMachineOutput), nil
	}
	out, err := c.BackupGatewayAPI.GetVirtualMachineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListGateways(input *backupgateway.ListGatewaysInput) (*backupgateway.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListGatewaysOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListGatewaysPages(input *backupgateway.ListGatewaysInput, fn func(*backupgateway.ListGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListGatewaysOutput{}
	fnCacher := func(out *backupgateway.ListGatewaysOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListGatewaysPagesWithContext(ctx context.Context, input *backupgateway.ListGatewaysInput, fn func(*backupgateway.ListGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListGatewaysOutput{}
	fnCacher := func(out *backupgateway.ListGatewaysOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListGatewaysWithContext(ctx context.Context, input *backupgateway.ListGatewaysInput, opts ...request.Option) (*backupgateway.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListGatewaysOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListHypervisors(input *backupgateway.ListHypervisorsInput) (*backupgateway.ListHypervisorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListHypervisorsOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListHypervisors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListHypervisorsPages(input *backupgateway.ListHypervisorsInput, fn func(*backupgateway.ListHypervisorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListHypervisorsOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListHypervisorsOutput{}
	fnCacher := func(out *backupgateway.ListHypervisorsOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListHypervisorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListHypervisorsPagesWithContext(ctx context.Context, input *backupgateway.ListHypervisorsInput, fn func(*backupgateway.ListHypervisorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListHypervisorsOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListHypervisorsOutput{}
	fnCacher := func(out *backupgateway.ListHypervisorsOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListHypervisorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListHypervisorsWithContext(ctx context.Context, input *backupgateway.ListHypervisorsInput, opts ...request.Option) (*backupgateway.ListHypervisorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListHypervisorsOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListHypervisorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListTagsForResource(input *backupgateway.ListTagsForResourceInput) (*backupgateway.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListTagsForResourceOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListTagsForResourceWithContext(ctx context.Context, input *backupgateway.ListTagsForResourceInput, opts ...request.Option) (*backupgateway.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListTagsForResourceOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListVirtualMachines(input *backupgateway.ListVirtualMachinesInput) (*backupgateway.ListVirtualMachinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListVirtualMachinesOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListVirtualMachines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BackupGateway) ListVirtualMachinesPages(input *backupgateway.ListVirtualMachinesInput, fn func(*backupgateway.ListVirtualMachinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListVirtualMachinesOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListVirtualMachinesOutput{}
	fnCacher := func(out *backupgateway.ListVirtualMachinesOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListVirtualMachinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListVirtualMachinesPagesWithContext(ctx context.Context, input *backupgateway.ListVirtualMachinesInput, fn func(*backupgateway.ListVirtualMachinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*backupgateway.ListVirtualMachinesOutput), false)
		return nil
	}
	cachable := true
	output := &backupgateway.ListVirtualMachinesOutput{}
	fnCacher := func(out *backupgateway.ListVirtualMachinesOutput, more bool) bool {
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
	if err := c.BackupGatewayAPI.ListVirtualMachinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BackupGateway) ListVirtualMachinesWithContext(ctx context.Context, input *backupgateway.ListVirtualMachinesInput, opts ...request.Option) (*backupgateway.ListVirtualMachinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*backupgateway.ListVirtualMachinesOutput), nil
	}
	out, err := c.BackupGatewayAPI.ListVirtualMachinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
