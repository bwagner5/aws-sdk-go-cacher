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
package iotfleetwisecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotfleetwise"
	"github.com/aws/aws-sdk-go/service/iotfleetwise/iotfleetwiseiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTFleetWise struct {
	iotfleetwiseiface.IoTFleetWiseAPI
	cache *cache.Cache
}

func New(iotfleetwiseapi iotfleetwiseiface.IoTFleetWiseAPI) *IoTFleetWise {
	return &IoTFleetWise{
		IoTFleetWiseAPI: iotfleetwiseapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTFleetWise) BatchCreateVehicle(input *iotfleetwise.BatchCreateVehicleInput) (*iotfleetwise.BatchCreateVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.BatchCreateVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.BatchCreateVehicle(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) BatchCreateVehicleWithContext(ctx context.Context, input *iotfleetwise.BatchCreateVehicleInput, opts ...request.Option) (*iotfleetwise.BatchCreateVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.BatchCreateVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.BatchCreateVehicleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) BatchUpdateVehicle(input *iotfleetwise.BatchUpdateVehicleInput) (*iotfleetwise.BatchUpdateVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.BatchUpdateVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.BatchUpdateVehicle(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) BatchUpdateVehicleWithContext(ctx context.Context, input *iotfleetwise.BatchUpdateVehicleInput, opts ...request.Option) (*iotfleetwise.BatchUpdateVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.BatchUpdateVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.BatchUpdateVehicleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetCampaign(input *iotfleetwise.GetCampaignInput) (*iotfleetwise.GetCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetCampaignOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetCampaignWithContext(ctx context.Context, input *iotfleetwise.GetCampaignInput, opts ...request.Option) (*iotfleetwise.GetCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetCampaignOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetDecoderManifest(input *iotfleetwise.GetDecoderManifestInput) (*iotfleetwise.GetDecoderManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetDecoderManifestOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetDecoderManifest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetDecoderManifestWithContext(ctx context.Context, input *iotfleetwise.GetDecoderManifestInput, opts ...request.Option) (*iotfleetwise.GetDecoderManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetDecoderManifestOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetDecoderManifestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetFleet(input *iotfleetwise.GetFleetInput) (*iotfleetwise.GetFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetFleetOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetFleetWithContext(ctx context.Context, input *iotfleetwise.GetFleetInput, opts ...request.Option) (*iotfleetwise.GetFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetFleetOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetLoggingOptions(input *iotfleetwise.GetLoggingOptionsInput) (*iotfleetwise.GetLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetLoggingOptionsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetLoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetLoggingOptionsWithContext(ctx context.Context, input *iotfleetwise.GetLoggingOptionsInput, opts ...request.Option) (*iotfleetwise.GetLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetLoggingOptionsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetLoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetModelManifest(input *iotfleetwise.GetModelManifestInput) (*iotfleetwise.GetModelManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetModelManifestOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetModelManifest(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetModelManifestWithContext(ctx context.Context, input *iotfleetwise.GetModelManifestInput, opts ...request.Option) (*iotfleetwise.GetModelManifestOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetModelManifestOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetModelManifestWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetRegisterAccountStatus(input *iotfleetwise.GetRegisterAccountStatusInput) (*iotfleetwise.GetRegisterAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetRegisterAccountStatusOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetRegisterAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetRegisterAccountStatusWithContext(ctx context.Context, input *iotfleetwise.GetRegisterAccountStatusInput, opts ...request.Option) (*iotfleetwise.GetRegisterAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetRegisterAccountStatusOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetRegisterAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetSignalCatalog(input *iotfleetwise.GetSignalCatalogInput) (*iotfleetwise.GetSignalCatalogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetSignalCatalogOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetSignalCatalog(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetSignalCatalogWithContext(ctx context.Context, input *iotfleetwise.GetSignalCatalogInput, opts ...request.Option) (*iotfleetwise.GetSignalCatalogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetSignalCatalogOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetSignalCatalogWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetVehicle(input *iotfleetwise.GetVehicleInput) (*iotfleetwise.GetVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetVehicle(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetVehicleStatus(input *iotfleetwise.GetVehicleStatusInput) (*iotfleetwise.GetVehicleStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetVehicleStatusOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetVehicleStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetVehicleStatusPages(input *iotfleetwise.GetVehicleStatusInput, fn func(*iotfleetwise.GetVehicleStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.GetVehicleStatusOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.GetVehicleStatusOutput{}
	fnCacher := func(out *iotfleetwise.GetVehicleStatusOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.GetVehicleStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) GetVehicleStatusPagesWithContext(ctx context.Context, input *iotfleetwise.GetVehicleStatusInput, fn func(*iotfleetwise.GetVehicleStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.GetVehicleStatusOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.GetVehicleStatusOutput{}
	fnCacher := func(out *iotfleetwise.GetVehicleStatusOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.GetVehicleStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) GetVehicleStatusWithContext(ctx context.Context, input *iotfleetwise.GetVehicleStatusInput, opts ...request.Option) (*iotfleetwise.GetVehicleStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetVehicleStatusOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetVehicleStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) GetVehicleWithContext(ctx context.Context, input *iotfleetwise.GetVehicleInput, opts ...request.Option) (*iotfleetwise.GetVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.GetVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.GetVehicleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListCampaigns(input *iotfleetwise.ListCampaignsInput) (*iotfleetwise.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListCampaignsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListCampaignsPages(input *iotfleetwise.ListCampaignsInput, fn func(*iotfleetwise.ListCampaignsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListCampaignsOutput{}
	fnCacher := func(out *iotfleetwise.ListCampaignsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListCampaignsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListCampaignsPagesWithContext(ctx context.Context, input *iotfleetwise.ListCampaignsInput, fn func(*iotfleetwise.ListCampaignsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListCampaignsOutput{}
	fnCacher := func(out *iotfleetwise.ListCampaignsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListCampaignsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListCampaignsWithContext(ctx context.Context, input *iotfleetwise.ListCampaignsInput, opts ...request.Option) (*iotfleetwise.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListCampaignsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifestNetworkInterfaces(input *iotfleetwise.ListDecoderManifestNetworkInterfacesInput) (*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifestNetworkInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifestNetworkInterfacesPages(input *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, fn func(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestNetworkInterfacesOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestNetworkInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestNetworkInterfacesPagesWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, fn func(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestNetworkInterfacesOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestNetworkInterfacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestNetworkInterfacesWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestNetworkInterfacesInput, opts ...request.Option) (*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestNetworkInterfacesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifestNetworkInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifestSignals(input *iotfleetwise.ListDecoderManifestSignalsInput) (*iotfleetwise.ListDecoderManifestSignalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestSignalsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifestSignals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifestSignalsPages(input *iotfleetwise.ListDecoderManifestSignalsInput, fn func(*iotfleetwise.ListDecoderManifestSignalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestSignalsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestSignalsOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestSignalsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestSignalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestSignalsPagesWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestSignalsInput, fn func(*iotfleetwise.ListDecoderManifestSignalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestSignalsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestSignalsOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestSignalsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestSignalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestSignalsWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestSignalsInput, opts ...request.Option) (*iotfleetwise.ListDecoderManifestSignalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestSignalsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifestSignalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifests(input *iotfleetwise.ListDecoderManifestsInput) (*iotfleetwise.ListDecoderManifestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListDecoderManifestsPages(input *iotfleetwise.ListDecoderManifestsInput, fn func(*iotfleetwise.ListDecoderManifestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestsOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestsPagesWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestsInput, fn func(*iotfleetwise.ListDecoderManifestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListDecoderManifestsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListDecoderManifestsOutput{}
	fnCacher := func(out *iotfleetwise.ListDecoderManifestsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListDecoderManifestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListDecoderManifestsWithContext(ctx context.Context, input *iotfleetwise.ListDecoderManifestsInput, opts ...request.Option) (*iotfleetwise.ListDecoderManifestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListDecoderManifestsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListDecoderManifestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListFleets(input *iotfleetwise.ListFleetsInput) (*iotfleetwise.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListFleetsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListFleetsForVehicle(input *iotfleetwise.ListFleetsForVehicleInput) (*iotfleetwise.ListFleetsForVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListFleetsForVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListFleetsForVehicle(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListFleetsForVehiclePages(input *iotfleetwise.ListFleetsForVehicleInput, fn func(*iotfleetwise.ListFleetsForVehicleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListFleetsForVehicleOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListFleetsForVehicleOutput{}
	fnCacher := func(out *iotfleetwise.ListFleetsForVehicleOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListFleetsForVehiclePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListFleetsForVehiclePagesWithContext(ctx context.Context, input *iotfleetwise.ListFleetsForVehicleInput, fn func(*iotfleetwise.ListFleetsForVehicleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListFleetsForVehicleOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListFleetsForVehicleOutput{}
	fnCacher := func(out *iotfleetwise.ListFleetsForVehicleOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListFleetsForVehiclePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListFleetsForVehicleWithContext(ctx context.Context, input *iotfleetwise.ListFleetsForVehicleInput, opts ...request.Option) (*iotfleetwise.ListFleetsForVehicleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListFleetsForVehicleOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListFleetsForVehicleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListFleetsPages(input *iotfleetwise.ListFleetsInput, fn func(*iotfleetwise.ListFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListFleetsOutput{}
	fnCacher := func(out *iotfleetwise.ListFleetsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListFleetsPagesWithContext(ctx context.Context, input *iotfleetwise.ListFleetsInput, fn func(*iotfleetwise.ListFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListFleetsOutput{}
	fnCacher := func(out *iotfleetwise.ListFleetsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListFleetsWithContext(ctx context.Context, input *iotfleetwise.ListFleetsInput, opts ...request.Option) (*iotfleetwise.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListFleetsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListModelManifestNodes(input *iotfleetwise.ListModelManifestNodesInput) (*iotfleetwise.ListModelManifestNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListModelManifestNodesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListModelManifestNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListModelManifestNodesPages(input *iotfleetwise.ListModelManifestNodesInput, fn func(*iotfleetwise.ListModelManifestNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListModelManifestNodesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListModelManifestNodesOutput{}
	fnCacher := func(out *iotfleetwise.ListModelManifestNodesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListModelManifestNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListModelManifestNodesPagesWithContext(ctx context.Context, input *iotfleetwise.ListModelManifestNodesInput, fn func(*iotfleetwise.ListModelManifestNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListModelManifestNodesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListModelManifestNodesOutput{}
	fnCacher := func(out *iotfleetwise.ListModelManifestNodesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListModelManifestNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListModelManifestNodesWithContext(ctx context.Context, input *iotfleetwise.ListModelManifestNodesInput, opts ...request.Option) (*iotfleetwise.ListModelManifestNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListModelManifestNodesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListModelManifestNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListModelManifests(input *iotfleetwise.ListModelManifestsInput) (*iotfleetwise.ListModelManifestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListModelManifestsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListModelManifests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListModelManifestsPages(input *iotfleetwise.ListModelManifestsInput, fn func(*iotfleetwise.ListModelManifestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListModelManifestsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListModelManifestsOutput{}
	fnCacher := func(out *iotfleetwise.ListModelManifestsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListModelManifestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListModelManifestsPagesWithContext(ctx context.Context, input *iotfleetwise.ListModelManifestsInput, fn func(*iotfleetwise.ListModelManifestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListModelManifestsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListModelManifestsOutput{}
	fnCacher := func(out *iotfleetwise.ListModelManifestsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListModelManifestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListModelManifestsWithContext(ctx context.Context, input *iotfleetwise.ListModelManifestsInput, opts ...request.Option) (*iotfleetwise.ListModelManifestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListModelManifestsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListModelManifestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListSignalCatalogNodes(input *iotfleetwise.ListSignalCatalogNodesInput) (*iotfleetwise.ListSignalCatalogNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListSignalCatalogNodesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListSignalCatalogNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListSignalCatalogNodesPages(input *iotfleetwise.ListSignalCatalogNodesInput, fn func(*iotfleetwise.ListSignalCatalogNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListSignalCatalogNodesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListSignalCatalogNodesOutput{}
	fnCacher := func(out *iotfleetwise.ListSignalCatalogNodesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListSignalCatalogNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListSignalCatalogNodesPagesWithContext(ctx context.Context, input *iotfleetwise.ListSignalCatalogNodesInput, fn func(*iotfleetwise.ListSignalCatalogNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListSignalCatalogNodesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListSignalCatalogNodesOutput{}
	fnCacher := func(out *iotfleetwise.ListSignalCatalogNodesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListSignalCatalogNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListSignalCatalogNodesWithContext(ctx context.Context, input *iotfleetwise.ListSignalCatalogNodesInput, opts ...request.Option) (*iotfleetwise.ListSignalCatalogNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListSignalCatalogNodesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListSignalCatalogNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListSignalCatalogs(input *iotfleetwise.ListSignalCatalogsInput) (*iotfleetwise.ListSignalCatalogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListSignalCatalogsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListSignalCatalogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListSignalCatalogsPages(input *iotfleetwise.ListSignalCatalogsInput, fn func(*iotfleetwise.ListSignalCatalogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListSignalCatalogsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListSignalCatalogsOutput{}
	fnCacher := func(out *iotfleetwise.ListSignalCatalogsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListSignalCatalogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListSignalCatalogsPagesWithContext(ctx context.Context, input *iotfleetwise.ListSignalCatalogsInput, fn func(*iotfleetwise.ListSignalCatalogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListSignalCatalogsOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListSignalCatalogsOutput{}
	fnCacher := func(out *iotfleetwise.ListSignalCatalogsOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListSignalCatalogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListSignalCatalogsWithContext(ctx context.Context, input *iotfleetwise.ListSignalCatalogsInput, opts ...request.Option) (*iotfleetwise.ListSignalCatalogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListSignalCatalogsOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListSignalCatalogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListTagsForResource(input *iotfleetwise.ListTagsForResourceInput) (*iotfleetwise.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListTagsForResourceWithContext(ctx context.Context, input *iotfleetwise.ListTagsForResourceInput, opts ...request.Option) (*iotfleetwise.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListVehicles(input *iotfleetwise.ListVehiclesInput) (*iotfleetwise.ListVehiclesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListVehiclesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListVehicles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListVehiclesInFleet(input *iotfleetwise.ListVehiclesInFleetInput) (*iotfleetwise.ListVehiclesInFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListVehiclesInFleetOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListVehiclesInFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListVehiclesInFleetPages(input *iotfleetwise.ListVehiclesInFleetInput, fn func(*iotfleetwise.ListVehiclesInFleetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListVehiclesInFleetOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListVehiclesInFleetOutput{}
	fnCacher := func(out *iotfleetwise.ListVehiclesInFleetOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListVehiclesInFleetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListVehiclesInFleetPagesWithContext(ctx context.Context, input *iotfleetwise.ListVehiclesInFleetInput, fn func(*iotfleetwise.ListVehiclesInFleetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListVehiclesInFleetOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListVehiclesInFleetOutput{}
	fnCacher := func(out *iotfleetwise.ListVehiclesInFleetOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListVehiclesInFleetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListVehiclesInFleetWithContext(ctx context.Context, input *iotfleetwise.ListVehiclesInFleetInput, opts ...request.Option) (*iotfleetwise.ListVehiclesInFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListVehiclesInFleetOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListVehiclesInFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTFleetWise) ListVehiclesPages(input *iotfleetwise.ListVehiclesInput, fn func(*iotfleetwise.ListVehiclesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListVehiclesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListVehiclesOutput{}
	fnCacher := func(out *iotfleetwise.ListVehiclesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListVehiclesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListVehiclesPagesWithContext(ctx context.Context, input *iotfleetwise.ListVehiclesInput, fn func(*iotfleetwise.ListVehiclesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotfleetwise.ListVehiclesOutput), false)
		return nil
	}
	cachable := true
	output := &iotfleetwise.ListVehiclesOutput{}
	fnCacher := func(out *iotfleetwise.ListVehiclesOutput, more bool) bool {
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
	if err := c.IoTFleetWiseAPI.ListVehiclesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTFleetWise) ListVehiclesWithContext(ctx context.Context, input *iotfleetwise.ListVehiclesInput, opts ...request.Option) (*iotfleetwise.ListVehiclesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotfleetwise.ListVehiclesOutput), nil
	}
	out, err := c.IoTFleetWiseAPI.ListVehiclesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
