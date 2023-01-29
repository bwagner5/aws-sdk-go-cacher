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
package pinpointemailcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pinpointemail"
	"github.com/aws/aws-sdk-go/service/pinpointemail/pinpointemailiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PinpointEmail struct {
	pinpointemailiface.PinpointEmailAPI
	cache *cache.Cache
}

func New(pinpointemailapi pinpointemailiface.PinpointEmailAPI) *PinpointEmail {
	return &PinpointEmail{
		PinpointEmailAPI: pinpointemailapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PinpointEmail) GetAccount(input *pinpointemail.GetAccountInput) (*pinpointemail.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetAccountOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetAccountWithContext(ctx context.Context, input *pinpointemail.GetAccountInput, opts ...request.Option) (*pinpointemail.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetAccountOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetBlacklistReports(input *pinpointemail.GetBlacklistReportsInput) (*pinpointemail.GetBlacklistReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetBlacklistReportsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetBlacklistReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetBlacklistReportsWithContext(ctx context.Context, input *pinpointemail.GetBlacklistReportsInput, opts ...request.Option) (*pinpointemail.GetBlacklistReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetBlacklistReportsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetBlacklistReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetConfigurationSet(input *pinpointemail.GetConfigurationSetInput) (*pinpointemail.GetConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetConfigurationSetOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetConfigurationSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetConfigurationSetEventDestinations(input *pinpointemail.GetConfigurationSetEventDestinationsInput) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetConfigurationSetEventDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetConfigurationSetEventDestinationsWithContext(ctx context.Context, input *pinpointemail.GetConfigurationSetEventDestinationsInput, opts ...request.Option) (*pinpointemail.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetConfigurationSetEventDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetConfigurationSetWithContext(ctx context.Context, input *pinpointemail.GetConfigurationSetInput, opts ...request.Option) (*pinpointemail.GetConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetConfigurationSetOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetConfigurationSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDedicatedIp(input *pinpointemail.GetDedicatedIpInput) (*pinpointemail.GetDedicatedIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDedicatedIpOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDedicatedIp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDedicatedIpWithContext(ctx context.Context, input *pinpointemail.GetDedicatedIpInput, opts ...request.Option) (*pinpointemail.GetDedicatedIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDedicatedIpOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDedicatedIpWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDedicatedIps(input *pinpointemail.GetDedicatedIpsInput) (*pinpointemail.GetDedicatedIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDedicatedIpsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDedicatedIps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDedicatedIpsPages(input *pinpointemail.GetDedicatedIpsInput, fn func(*pinpointemail.GetDedicatedIpsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.GetDedicatedIpsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.GetDedicatedIpsOutput{}
	fnCacher := func(out *pinpointemail.GetDedicatedIpsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.GetDedicatedIpsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) GetDedicatedIpsPagesWithContext(ctx context.Context, input *pinpointemail.GetDedicatedIpsInput, fn func(*pinpointemail.GetDedicatedIpsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.GetDedicatedIpsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.GetDedicatedIpsOutput{}
	fnCacher := func(out *pinpointemail.GetDedicatedIpsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.GetDedicatedIpsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) GetDedicatedIpsWithContext(ctx context.Context, input *pinpointemail.GetDedicatedIpsInput, opts ...request.Option) (*pinpointemail.GetDedicatedIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDedicatedIpsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDedicatedIpsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDeliverabilityDashboardOptions(input *pinpointemail.GetDeliverabilityDashboardOptionsInput) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDeliverabilityDashboardOptionsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDeliverabilityDashboardOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDeliverabilityDashboardOptionsWithContext(ctx context.Context, input *pinpointemail.GetDeliverabilityDashboardOptionsInput, opts ...request.Option) (*pinpointemail.GetDeliverabilityDashboardOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDeliverabilityDashboardOptionsOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDeliverabilityDashboardOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDeliverabilityTestReport(input *pinpointemail.GetDeliverabilityTestReportInput) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDeliverabilityTestReportOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDeliverabilityTestReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDeliverabilityTestReportWithContext(ctx context.Context, input *pinpointemail.GetDeliverabilityTestReportInput, opts ...request.Option) (*pinpointemail.GetDeliverabilityTestReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDeliverabilityTestReportOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDeliverabilityTestReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDomainDeliverabilityCampaign(input *pinpointemail.GetDomainDeliverabilityCampaignInput) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDomainDeliverabilityCampaignOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDomainDeliverabilityCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDomainDeliverabilityCampaignWithContext(ctx context.Context, input *pinpointemail.GetDomainDeliverabilityCampaignInput, opts ...request.Option) (*pinpointemail.GetDomainDeliverabilityCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDomainDeliverabilityCampaignOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDomainDeliverabilityCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDomainStatisticsReport(input *pinpointemail.GetDomainStatisticsReportInput) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDomainStatisticsReportOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDomainStatisticsReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetDomainStatisticsReportWithContext(ctx context.Context, input *pinpointemail.GetDomainStatisticsReportInput, opts ...request.Option) (*pinpointemail.GetDomainStatisticsReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetDomainStatisticsReportOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetDomainStatisticsReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetEmailIdentity(input *pinpointemail.GetEmailIdentityInput) (*pinpointemail.GetEmailIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetEmailIdentityOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetEmailIdentity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) GetEmailIdentityWithContext(ctx context.Context, input *pinpointemail.GetEmailIdentityInput, opts ...request.Option) (*pinpointemail.GetEmailIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.GetEmailIdentityOutput), nil
	}
	out, err := c.PinpointEmailAPI.GetEmailIdentityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListConfigurationSets(input *pinpointemail.ListConfigurationSetsInput) (*pinpointemail.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListConfigurationSetsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListConfigurationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListConfigurationSetsPages(input *pinpointemail.ListConfigurationSetsInput, fn func(*pinpointemail.ListConfigurationSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListConfigurationSetsOutput{}
	fnCacher := func(out *pinpointemail.ListConfigurationSetsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListConfigurationSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListConfigurationSetsPagesWithContext(ctx context.Context, input *pinpointemail.ListConfigurationSetsInput, fn func(*pinpointemail.ListConfigurationSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListConfigurationSetsOutput{}
	fnCacher := func(out *pinpointemail.ListConfigurationSetsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListConfigurationSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListConfigurationSetsWithContext(ctx context.Context, input *pinpointemail.ListConfigurationSetsInput, opts ...request.Option) (*pinpointemail.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListConfigurationSetsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListConfigurationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDedicatedIpPools(input *pinpointemail.ListDedicatedIpPoolsInput) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDedicatedIpPoolsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDedicatedIpPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDedicatedIpPoolsPages(input *pinpointemail.ListDedicatedIpPoolsInput, fn func(*pinpointemail.ListDedicatedIpPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDedicatedIpPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDedicatedIpPoolsOutput{}
	fnCacher := func(out *pinpointemail.ListDedicatedIpPoolsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDedicatedIpPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDedicatedIpPoolsPagesWithContext(ctx context.Context, input *pinpointemail.ListDedicatedIpPoolsInput, fn func(*pinpointemail.ListDedicatedIpPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDedicatedIpPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDedicatedIpPoolsOutput{}
	fnCacher := func(out *pinpointemail.ListDedicatedIpPoolsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDedicatedIpPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDedicatedIpPoolsWithContext(ctx context.Context, input *pinpointemail.ListDedicatedIpPoolsInput, opts ...request.Option) (*pinpointemail.ListDedicatedIpPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDedicatedIpPoolsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDedicatedIpPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDeliverabilityTestReports(input *pinpointemail.ListDeliverabilityTestReportsInput) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDeliverabilityTestReportsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDeliverabilityTestReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDeliverabilityTestReportsPages(input *pinpointemail.ListDeliverabilityTestReportsInput, fn func(*pinpointemail.ListDeliverabilityTestReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDeliverabilityTestReportsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDeliverabilityTestReportsOutput{}
	fnCacher := func(out *pinpointemail.ListDeliverabilityTestReportsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDeliverabilityTestReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDeliverabilityTestReportsPagesWithContext(ctx context.Context, input *pinpointemail.ListDeliverabilityTestReportsInput, fn func(*pinpointemail.ListDeliverabilityTestReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDeliverabilityTestReportsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDeliverabilityTestReportsOutput{}
	fnCacher := func(out *pinpointemail.ListDeliverabilityTestReportsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDeliverabilityTestReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDeliverabilityTestReportsWithContext(ctx context.Context, input *pinpointemail.ListDeliverabilityTestReportsInput, opts ...request.Option) (*pinpointemail.ListDeliverabilityTestReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDeliverabilityTestReportsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDeliverabilityTestReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDomainDeliverabilityCampaigns(input *pinpointemail.ListDomainDeliverabilityCampaignsInput) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDomainDeliverabilityCampaignsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDomainDeliverabilityCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListDomainDeliverabilityCampaignsPages(input *pinpointemail.ListDomainDeliverabilityCampaignsInput, fn func(*pinpointemail.ListDomainDeliverabilityCampaignsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDomainDeliverabilityCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDomainDeliverabilityCampaignsOutput{}
	fnCacher := func(out *pinpointemail.ListDomainDeliverabilityCampaignsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDomainDeliverabilityCampaignsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDomainDeliverabilityCampaignsPagesWithContext(ctx context.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput, fn func(*pinpointemail.ListDomainDeliverabilityCampaignsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListDomainDeliverabilityCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListDomainDeliverabilityCampaignsOutput{}
	fnCacher := func(out *pinpointemail.ListDomainDeliverabilityCampaignsOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListDomainDeliverabilityCampaignsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListDomainDeliverabilityCampaignsWithContext(ctx context.Context, input *pinpointemail.ListDomainDeliverabilityCampaignsInput, opts ...request.Option) (*pinpointemail.ListDomainDeliverabilityCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListDomainDeliverabilityCampaignsOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListDomainDeliverabilityCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListEmailIdentities(input *pinpointemail.ListEmailIdentitiesInput) (*pinpointemail.ListEmailIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListEmailIdentitiesOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListEmailIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListEmailIdentitiesPages(input *pinpointemail.ListEmailIdentitiesInput, fn func(*pinpointemail.ListEmailIdentitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListEmailIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListEmailIdentitiesOutput{}
	fnCacher := func(out *pinpointemail.ListEmailIdentitiesOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListEmailIdentitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListEmailIdentitiesPagesWithContext(ctx context.Context, input *pinpointemail.ListEmailIdentitiesInput, fn func(*pinpointemail.ListEmailIdentitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointemail.ListEmailIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointemail.ListEmailIdentitiesOutput{}
	fnCacher := func(out *pinpointemail.ListEmailIdentitiesOutput, more bool) bool {
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
	if err := c.PinpointEmailAPI.ListEmailIdentitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointEmail) ListEmailIdentitiesWithContext(ctx context.Context, input *pinpointemail.ListEmailIdentitiesInput, opts ...request.Option) (*pinpointemail.ListEmailIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListEmailIdentitiesOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListEmailIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListTagsForResource(input *pinpointemail.ListTagsForResourceInput) (*pinpointemail.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointEmail) ListTagsForResourceWithContext(ctx context.Context, input *pinpointemail.ListTagsForResourceInput, opts ...request.Option) (*pinpointemail.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointemail.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointEmailAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
