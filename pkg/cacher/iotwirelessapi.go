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
	"github.com/aws/aws-sdk-go/service/iotwireless"
	"github.com/aws/aws-sdk-go/service/iotwireless/iotwirelessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTWireless struct {
	iotwirelessiface.IoTWirelessAPI
	cache *cache.Cache
}

func NewIoTWireless(iotwirelessapi iotwirelessiface.IoTWirelessAPI) *IoTWireless {
	return &IoTWireless{
		IoTWirelessAPI: iotwirelessapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTWireless) GetDestination(input *iotwireless.GetDestinationInput) (*iotwireless.GetDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetDestinationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetDestinationWithContext(ctx context.Context, input *iotwireless.GetDestinationInput, opts ...request.Option) (*iotwireless.GetDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetDestinationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetDeviceProfile(input *iotwireless.GetDeviceProfileInput) (*iotwireless.GetDeviceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetDeviceProfileOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetDeviceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetDeviceProfileWithContext(ctx context.Context, input *iotwireless.GetDeviceProfileInput, opts ...request.Option) (*iotwireless.GetDeviceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetDeviceProfileOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetDeviceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetEventConfigurationByResourceTypes(input *iotwireless.GetEventConfigurationByResourceTypesInput) (*iotwireless.GetEventConfigurationByResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetEventConfigurationByResourceTypesOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetEventConfigurationByResourceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetEventConfigurationByResourceTypesWithContext(ctx context.Context, input *iotwireless.GetEventConfigurationByResourceTypesInput, opts ...request.Option) (*iotwireless.GetEventConfigurationByResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetEventConfigurationByResourceTypesOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetEventConfigurationByResourceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetFuotaTask(input *iotwireless.GetFuotaTaskInput) (*iotwireless.GetFuotaTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetFuotaTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetFuotaTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetFuotaTaskWithContext(ctx context.Context, input *iotwireless.GetFuotaTaskInput, opts ...request.Option) (*iotwireless.GetFuotaTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetFuotaTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetFuotaTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetLogLevelsByResourceTypes(input *iotwireless.GetLogLevelsByResourceTypesInput) (*iotwireless.GetLogLevelsByResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetLogLevelsByResourceTypesOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetLogLevelsByResourceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetLogLevelsByResourceTypesWithContext(ctx context.Context, input *iotwireless.GetLogLevelsByResourceTypesInput, opts ...request.Option) (*iotwireless.GetLogLevelsByResourceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetLogLevelsByResourceTypesOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetLogLevelsByResourceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetMulticastGroup(input *iotwireless.GetMulticastGroupInput) (*iotwireless.GetMulticastGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetMulticastGroupOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetMulticastGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetMulticastGroupSession(input *iotwireless.GetMulticastGroupSessionInput) (*iotwireless.GetMulticastGroupSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetMulticastGroupSessionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetMulticastGroupSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetMulticastGroupSessionWithContext(ctx context.Context, input *iotwireless.GetMulticastGroupSessionInput, opts ...request.Option) (*iotwireless.GetMulticastGroupSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetMulticastGroupSessionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetMulticastGroupSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetMulticastGroupWithContext(ctx context.Context, input *iotwireless.GetMulticastGroupInput, opts ...request.Option) (*iotwireless.GetMulticastGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetMulticastGroupOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetMulticastGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetNetworkAnalyzerConfiguration(input *iotwireless.GetNetworkAnalyzerConfigurationInput) (*iotwireless.GetNetworkAnalyzerConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetNetworkAnalyzerConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetNetworkAnalyzerConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetNetworkAnalyzerConfigurationWithContext(ctx context.Context, input *iotwireless.GetNetworkAnalyzerConfigurationInput, opts ...request.Option) (*iotwireless.GetNetworkAnalyzerConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetNetworkAnalyzerConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetNetworkAnalyzerConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPartnerAccount(input *iotwireless.GetPartnerAccountInput) (*iotwireless.GetPartnerAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPartnerAccountOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPartnerAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPartnerAccountWithContext(ctx context.Context, input *iotwireless.GetPartnerAccountInput, opts ...request.Option) (*iotwireless.GetPartnerAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPartnerAccountOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPartnerAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPosition(input *iotwireless.GetPositionInput) (*iotwireless.GetPositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPositionConfiguration(input *iotwireless.GetPositionConfigurationInput) (*iotwireless.GetPositionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPositionConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPositionConfigurationWithContext(ctx context.Context, input *iotwireless.GetPositionConfigurationInput, opts ...request.Option) (*iotwireless.GetPositionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPositionConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPositionEstimate(input *iotwireless.GetPositionEstimateInput) (*iotwireless.GetPositionEstimateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionEstimateOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPositionEstimate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPositionEstimateWithContext(ctx context.Context, input *iotwireless.GetPositionEstimateInput, opts ...request.Option) (*iotwireless.GetPositionEstimateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionEstimateOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPositionEstimateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetPositionWithContext(ctx context.Context, input *iotwireless.GetPositionInput, opts ...request.Option) (*iotwireless.GetPositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetPositionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetPositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourceEventConfiguration(input *iotwireless.GetResourceEventConfigurationInput) (*iotwireless.GetResourceEventConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourceEventConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourceEventConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourceEventConfigurationWithContext(ctx context.Context, input *iotwireless.GetResourceEventConfigurationInput, opts ...request.Option) (*iotwireless.GetResourceEventConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourceEventConfigurationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourceEventConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourceLogLevel(input *iotwireless.GetResourceLogLevelInput) (*iotwireless.GetResourceLogLevelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourceLogLevelOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourceLogLevel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourceLogLevelWithContext(ctx context.Context, input *iotwireless.GetResourceLogLevelInput, opts ...request.Option) (*iotwireless.GetResourceLogLevelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourceLogLevelOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourceLogLevelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourcePosition(input *iotwireless.GetResourcePositionInput) (*iotwireless.GetResourcePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourcePositionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourcePosition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetResourcePositionWithContext(ctx context.Context, input *iotwireless.GetResourcePositionInput, opts ...request.Option) (*iotwireless.GetResourcePositionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetResourcePositionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetResourcePositionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetServiceEndpoint(input *iotwireless.GetServiceEndpointInput) (*iotwireless.GetServiceEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetServiceEndpointOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetServiceEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetServiceEndpointWithContext(ctx context.Context, input *iotwireless.GetServiceEndpointInput, opts ...request.Option) (*iotwireless.GetServiceEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetServiceEndpointOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetServiceEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetServiceProfile(input *iotwireless.GetServiceProfileInput) (*iotwireless.GetServiceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetServiceProfileOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetServiceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetServiceProfileWithContext(ctx context.Context, input *iotwireless.GetServiceProfileInput, opts ...request.Option) (*iotwireless.GetServiceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetServiceProfileOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetServiceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessDevice(input *iotwireless.GetWirelessDeviceInput) (*iotwireless.GetWirelessDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessDeviceOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessDeviceStatistics(input *iotwireless.GetWirelessDeviceStatisticsInput) (*iotwireless.GetWirelessDeviceStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessDeviceStatisticsOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessDeviceStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessDeviceStatisticsWithContext(ctx context.Context, input *iotwireless.GetWirelessDeviceStatisticsInput, opts ...request.Option) (*iotwireless.GetWirelessDeviceStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessDeviceStatisticsOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessDeviceStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessDeviceWithContext(ctx context.Context, input *iotwireless.GetWirelessDeviceInput, opts ...request.Option) (*iotwireless.GetWirelessDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessDeviceOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGateway(input *iotwireless.GetWirelessGatewayInput) (*iotwireless.GetWirelessGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayCertificate(input *iotwireless.GetWirelessGatewayCertificateInput) (*iotwireless.GetWirelessGatewayCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayCertificateOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayCertificateWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayCertificateInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayCertificateOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayFirmwareInformation(input *iotwireless.GetWirelessGatewayFirmwareInformationInput) (*iotwireless.GetWirelessGatewayFirmwareInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayFirmwareInformationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayFirmwareInformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayFirmwareInformationWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayFirmwareInformationInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayFirmwareInformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayFirmwareInformationOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayFirmwareInformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayStatistics(input *iotwireless.GetWirelessGatewayStatisticsInput) (*iotwireless.GetWirelessGatewayStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayStatisticsOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayStatisticsWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayStatisticsInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayStatisticsOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayTask(input *iotwireless.GetWirelessGatewayTaskInput) (*iotwireless.GetWirelessGatewayTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayTaskDefinition(input *iotwireless.GetWirelessGatewayTaskDefinitionInput) (*iotwireless.GetWirelessGatewayTaskDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayTaskDefinitionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayTaskDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayTaskDefinitionWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayTaskDefinitionInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayTaskDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayTaskDefinitionOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayTaskDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayTaskWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayTaskInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) GetWirelessGatewayWithContext(ctx context.Context, input *iotwireless.GetWirelessGatewayInput, opts ...request.Option) (*iotwireless.GetWirelessGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.GetWirelessGatewayOutput), nil
	}
	out, err := c.IoTWirelessAPI.GetWirelessGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListDestinations(input *iotwireless.ListDestinationsInput) (*iotwireless.ListDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListDestinationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListDestinationsPages(input *iotwireless.ListDestinationsInput, fn func(*iotwireless.ListDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListDestinationsOutput{}
	fnCacher := func(out *iotwireless.ListDestinationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListDestinationsPagesWithContext(ctx context.Context, input *iotwireless.ListDestinationsInput, fn func(*iotwireless.ListDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListDestinationsOutput{}
	fnCacher := func(out *iotwireless.ListDestinationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListDestinationsWithContext(ctx context.Context, input *iotwireless.ListDestinationsInput, opts ...request.Option) (*iotwireless.ListDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListDestinationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListDeviceProfiles(input *iotwireless.ListDeviceProfilesInput) (*iotwireless.ListDeviceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListDeviceProfilesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListDeviceProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListDeviceProfilesPages(input *iotwireless.ListDeviceProfilesInput, fn func(*iotwireless.ListDeviceProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListDeviceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListDeviceProfilesOutput{}
	fnCacher := func(out *iotwireless.ListDeviceProfilesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListDeviceProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListDeviceProfilesPagesWithContext(ctx context.Context, input *iotwireless.ListDeviceProfilesInput, fn func(*iotwireless.ListDeviceProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListDeviceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListDeviceProfilesOutput{}
	fnCacher := func(out *iotwireless.ListDeviceProfilesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListDeviceProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListDeviceProfilesWithContext(ctx context.Context, input *iotwireless.ListDeviceProfilesInput, opts ...request.Option) (*iotwireless.ListDeviceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListDeviceProfilesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListDeviceProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListEventConfigurations(input *iotwireless.ListEventConfigurationsInput) (*iotwireless.ListEventConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListEventConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListEventConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListEventConfigurationsWithContext(ctx context.Context, input *iotwireless.ListEventConfigurationsInput, opts ...request.Option) (*iotwireless.ListEventConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListEventConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListEventConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListFuotaTasks(input *iotwireless.ListFuotaTasksInput) (*iotwireless.ListFuotaTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListFuotaTasksOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListFuotaTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListFuotaTasksPages(input *iotwireless.ListFuotaTasksInput, fn func(*iotwireless.ListFuotaTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListFuotaTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListFuotaTasksOutput{}
	fnCacher := func(out *iotwireless.ListFuotaTasksOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListFuotaTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListFuotaTasksPagesWithContext(ctx context.Context, input *iotwireless.ListFuotaTasksInput, fn func(*iotwireless.ListFuotaTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListFuotaTasksOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListFuotaTasksOutput{}
	fnCacher := func(out *iotwireless.ListFuotaTasksOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListFuotaTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListFuotaTasksWithContext(ctx context.Context, input *iotwireless.ListFuotaTasksInput, opts ...request.Option) (*iotwireless.ListFuotaTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListFuotaTasksOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListFuotaTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListMulticastGroups(input *iotwireless.ListMulticastGroupsInput) (*iotwireless.ListMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListMulticastGroupsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListMulticastGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListMulticastGroupsByFuotaTask(input *iotwireless.ListMulticastGroupsByFuotaTaskInput) (*iotwireless.ListMulticastGroupsByFuotaTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListMulticastGroupsByFuotaTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListMulticastGroupsByFuotaTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListMulticastGroupsByFuotaTaskPages(input *iotwireless.ListMulticastGroupsByFuotaTaskInput, fn func(*iotwireless.ListMulticastGroupsByFuotaTaskOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListMulticastGroupsByFuotaTaskOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListMulticastGroupsByFuotaTaskOutput{}
	fnCacher := func(out *iotwireless.ListMulticastGroupsByFuotaTaskOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListMulticastGroupsByFuotaTaskPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListMulticastGroupsByFuotaTaskPagesWithContext(ctx context.Context, input *iotwireless.ListMulticastGroupsByFuotaTaskInput, fn func(*iotwireless.ListMulticastGroupsByFuotaTaskOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListMulticastGroupsByFuotaTaskOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListMulticastGroupsByFuotaTaskOutput{}
	fnCacher := func(out *iotwireless.ListMulticastGroupsByFuotaTaskOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListMulticastGroupsByFuotaTaskPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListMulticastGroupsByFuotaTaskWithContext(ctx context.Context, input *iotwireless.ListMulticastGroupsByFuotaTaskInput, opts ...request.Option) (*iotwireless.ListMulticastGroupsByFuotaTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListMulticastGroupsByFuotaTaskOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListMulticastGroupsByFuotaTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListMulticastGroupsPages(input *iotwireless.ListMulticastGroupsInput, fn func(*iotwireless.ListMulticastGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListMulticastGroupsOutput{}
	fnCacher := func(out *iotwireless.ListMulticastGroupsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListMulticastGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListMulticastGroupsPagesWithContext(ctx context.Context, input *iotwireless.ListMulticastGroupsInput, fn func(*iotwireless.ListMulticastGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListMulticastGroupsOutput{}
	fnCacher := func(out *iotwireless.ListMulticastGroupsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListMulticastGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListMulticastGroupsWithContext(ctx context.Context, input *iotwireless.ListMulticastGroupsInput, opts ...request.Option) (*iotwireless.ListMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListMulticastGroupsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListMulticastGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListNetworkAnalyzerConfigurations(input *iotwireless.ListNetworkAnalyzerConfigurationsInput) (*iotwireless.ListNetworkAnalyzerConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListNetworkAnalyzerConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListNetworkAnalyzerConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListNetworkAnalyzerConfigurationsPages(input *iotwireless.ListNetworkAnalyzerConfigurationsInput, fn func(*iotwireless.ListNetworkAnalyzerConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListNetworkAnalyzerConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListNetworkAnalyzerConfigurationsOutput{}
	fnCacher := func(out *iotwireless.ListNetworkAnalyzerConfigurationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListNetworkAnalyzerConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListNetworkAnalyzerConfigurationsPagesWithContext(ctx context.Context, input *iotwireless.ListNetworkAnalyzerConfigurationsInput, fn func(*iotwireless.ListNetworkAnalyzerConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListNetworkAnalyzerConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListNetworkAnalyzerConfigurationsOutput{}
	fnCacher := func(out *iotwireless.ListNetworkAnalyzerConfigurationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListNetworkAnalyzerConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListNetworkAnalyzerConfigurationsWithContext(ctx context.Context, input *iotwireless.ListNetworkAnalyzerConfigurationsInput, opts ...request.Option) (*iotwireless.ListNetworkAnalyzerConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListNetworkAnalyzerConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListNetworkAnalyzerConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListPartnerAccounts(input *iotwireless.ListPartnerAccountsInput) (*iotwireless.ListPartnerAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListPartnerAccountsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListPartnerAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListPartnerAccountsWithContext(ctx context.Context, input *iotwireless.ListPartnerAccountsInput, opts ...request.Option) (*iotwireless.ListPartnerAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListPartnerAccountsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListPartnerAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListPositionConfigurations(input *iotwireless.ListPositionConfigurationsInput) (*iotwireless.ListPositionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListPositionConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListPositionConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListPositionConfigurationsPages(input *iotwireless.ListPositionConfigurationsInput, fn func(*iotwireless.ListPositionConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListPositionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListPositionConfigurationsOutput{}
	fnCacher := func(out *iotwireless.ListPositionConfigurationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListPositionConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListPositionConfigurationsPagesWithContext(ctx context.Context, input *iotwireless.ListPositionConfigurationsInput, fn func(*iotwireless.ListPositionConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListPositionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListPositionConfigurationsOutput{}
	fnCacher := func(out *iotwireless.ListPositionConfigurationsOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListPositionConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListPositionConfigurationsWithContext(ctx context.Context, input *iotwireless.ListPositionConfigurationsInput, opts ...request.Option) (*iotwireless.ListPositionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListPositionConfigurationsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListPositionConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListQueuedMessages(input *iotwireless.ListQueuedMessagesInput) (*iotwireless.ListQueuedMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListQueuedMessagesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListQueuedMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListQueuedMessagesPages(input *iotwireless.ListQueuedMessagesInput, fn func(*iotwireless.ListQueuedMessagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListQueuedMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListQueuedMessagesOutput{}
	fnCacher := func(out *iotwireless.ListQueuedMessagesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListQueuedMessagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListQueuedMessagesPagesWithContext(ctx context.Context, input *iotwireless.ListQueuedMessagesInput, fn func(*iotwireless.ListQueuedMessagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListQueuedMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListQueuedMessagesOutput{}
	fnCacher := func(out *iotwireless.ListQueuedMessagesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListQueuedMessagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListQueuedMessagesWithContext(ctx context.Context, input *iotwireless.ListQueuedMessagesInput, opts ...request.Option) (*iotwireless.ListQueuedMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListQueuedMessagesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListQueuedMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListServiceProfiles(input *iotwireless.ListServiceProfilesInput) (*iotwireless.ListServiceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListServiceProfilesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListServiceProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListServiceProfilesPages(input *iotwireless.ListServiceProfilesInput, fn func(*iotwireless.ListServiceProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListServiceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListServiceProfilesOutput{}
	fnCacher := func(out *iotwireless.ListServiceProfilesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListServiceProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListServiceProfilesPagesWithContext(ctx context.Context, input *iotwireless.ListServiceProfilesInput, fn func(*iotwireless.ListServiceProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListServiceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListServiceProfilesOutput{}
	fnCacher := func(out *iotwireless.ListServiceProfilesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListServiceProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListServiceProfilesWithContext(ctx context.Context, input *iotwireless.ListServiceProfilesInput, opts ...request.Option) (*iotwireless.ListServiceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListServiceProfilesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListServiceProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListTagsForResource(input *iotwireless.ListTagsForResourceInput) (*iotwireless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListTagsForResourceWithContext(ctx context.Context, input *iotwireless.ListTagsForResourceInput, opts ...request.Option) (*iotwireless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessDevices(input *iotwireless.ListWirelessDevicesInput) (*iotwireless.ListWirelessDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessDevicesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessDevicesPages(input *iotwireless.ListWirelessDevicesInput, fn func(*iotwireless.ListWirelessDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListWirelessDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListWirelessDevicesOutput{}
	fnCacher := func(out *iotwireless.ListWirelessDevicesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListWirelessDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListWirelessDevicesPagesWithContext(ctx context.Context, input *iotwireless.ListWirelessDevicesInput, fn func(*iotwireless.ListWirelessDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListWirelessDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListWirelessDevicesOutput{}
	fnCacher := func(out *iotwireless.ListWirelessDevicesOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListWirelessDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListWirelessDevicesWithContext(ctx context.Context, input *iotwireless.ListWirelessDevicesInput, opts ...request.Option) (*iotwireless.ListWirelessDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessDevicesOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessGatewayTaskDefinitions(input *iotwireless.ListWirelessGatewayTaskDefinitionsInput) (*iotwireless.ListWirelessGatewayTaskDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessGatewayTaskDefinitionsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessGatewayTaskDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessGatewayTaskDefinitionsWithContext(ctx context.Context, input *iotwireless.ListWirelessGatewayTaskDefinitionsInput, opts ...request.Option) (*iotwireless.ListWirelessGatewayTaskDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessGatewayTaskDefinitionsOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessGatewayTaskDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessGateways(input *iotwireless.ListWirelessGatewaysInput) (*iotwireless.ListWirelessGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessGatewaysOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTWireless) ListWirelessGatewaysPages(input *iotwireless.ListWirelessGatewaysInput, fn func(*iotwireless.ListWirelessGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListWirelessGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListWirelessGatewaysOutput{}
	fnCacher := func(out *iotwireless.ListWirelessGatewaysOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListWirelessGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListWirelessGatewaysPagesWithContext(ctx context.Context, input *iotwireless.ListWirelessGatewaysInput, fn func(*iotwireless.ListWirelessGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotwireless.ListWirelessGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &iotwireless.ListWirelessGatewaysOutput{}
	fnCacher := func(out *iotwireless.ListWirelessGatewaysOutput, more bool) bool {
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
	if err := c.IoTWirelessAPI.ListWirelessGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTWireless) ListWirelessGatewaysWithContext(ctx context.Context, input *iotwireless.ListWirelessGatewaysInput, opts ...request.Option) (*iotwireless.ListWirelessGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotwireless.ListWirelessGatewaysOutput), nil
	}
	out, err := c.IoTWirelessAPI.ListWirelessGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
