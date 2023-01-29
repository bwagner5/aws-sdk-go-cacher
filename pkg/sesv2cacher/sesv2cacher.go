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
package sesv2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sesv2"
	"github.com/aws/aws-sdk-go/service/sesv2/sesv2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SESV2 struct {
	sesv2iface.SESV2API
	cache *cache.Cache
}

func New(sesv2api sesv2iface.SESV2API) *SESV2 {
	return &SESV2{
		SESV2API: sesv2api,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SESV2) BatchGetMetricData(input *sesv2.BatchGetMetricDataInput) (*sesv2.BatchGetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.BatchGetMetricDataOutput), nil
	}
	out, err := c.SESV2API.BatchGetMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) BatchGetMetricDataWithContext(ctx context.Context, input *sesv2.BatchGetMetricDataInput, opts ...request.Option) (*sesv2.BatchGetMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.BatchGetMetricDataOutput), nil
	}
	out, err := c.SESV2API.BatchGetMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetAccount(input *sesv2.GetAccountInput) (*sesv2.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetAccountOutput), nil
	}
	out, err := c.SESV2API.GetAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetAccountWithContext(ctx context.Context, input *sesv2.GetAccountInput, opts ...request.Option) (*sesv2.GetAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetAccountOutput), nil
	}
	out, err := c.SESV2API.GetAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetBlacklistReports(input *sesv2.GetBlacklistReportsInput) (*sesv2.GetBlacklistReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetBlacklistReportsOutput), nil
	}
	out, err := c.SESV2API.GetBlacklistReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetBlacklistReportsWithContext(ctx context.Context, input *sesv2.GetBlacklistReportsInput, opts ...request.Option) (*sesv2.GetBlacklistReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetBlacklistReportsOutput), nil
	}
	out, err := c.SESV2API.GetBlacklistReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetConfigurationSet(input *sesv2.GetConfigurationSetInput) (*sesv2.GetConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetConfigurationSetOutput), nil
	}
	out, err := c.SESV2API.GetConfigurationSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetConfigurationSetEventDestinations(input *sesv2.GetConfigurationSetEventDestinationsInput) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.SESV2API.GetConfigurationSetEventDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetConfigurationSetEventDestinationsWithContext(ctx context.Context, input *sesv2.GetConfigurationSetEventDestinationsInput, opts ...request.Option) (*sesv2.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.SESV2API.GetConfigurationSetEventDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetConfigurationSetWithContext(ctx context.Context, input *sesv2.GetConfigurationSetInput, opts ...request.Option) (*sesv2.GetConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetConfigurationSetOutput), nil
	}
	out, err := c.SESV2API.GetConfigurationSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetContact(input *sesv2.GetContactInput) (*sesv2.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetContactOutput), nil
	}
	out, err := c.SESV2API.GetContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetContactList(input *sesv2.GetContactListInput) (*sesv2.GetContactListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetContactListOutput), nil
	}
	out, err := c.SESV2API.GetContactList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetContactListWithContext(ctx context.Context, input *sesv2.GetContactListInput, opts ...request.Option) (*sesv2.GetContactListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetContactListOutput), nil
	}
	out, err := c.SESV2API.GetContactListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetContactWithContext(ctx context.Context, input *sesv2.GetContactInput, opts ...request.Option) (*sesv2.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetContactOutput), nil
	}
	out, err := c.SESV2API.GetContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetCustomVerificationEmailTemplate(input *sesv2.GetCustomVerificationEmailTemplateInput) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetCustomVerificationEmailTemplateOutput), nil
	}
	out, err := c.SESV2API.GetCustomVerificationEmailTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetCustomVerificationEmailTemplateWithContext(ctx context.Context, input *sesv2.GetCustomVerificationEmailTemplateInput, opts ...request.Option) (*sesv2.GetCustomVerificationEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetCustomVerificationEmailTemplateOutput), nil
	}
	out, err := c.SESV2API.GetCustomVerificationEmailTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIp(input *sesv2.GetDedicatedIpInput) (*sesv2.GetDedicatedIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIpPool(input *sesv2.GetDedicatedIpPoolInput) (*sesv2.GetDedicatedIpPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpPoolOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIpPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIpPoolWithContext(ctx context.Context, input *sesv2.GetDedicatedIpPoolInput, opts ...request.Option) (*sesv2.GetDedicatedIpPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpPoolOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIpPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIpWithContext(ctx context.Context, input *sesv2.GetDedicatedIpInput, opts ...request.Option) (*sesv2.GetDedicatedIpOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIpWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIps(input *sesv2.GetDedicatedIpsInput) (*sesv2.GetDedicatedIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpsOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDedicatedIpsPages(input *sesv2.GetDedicatedIpsInput, fn func(*sesv2.GetDedicatedIpsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.GetDedicatedIpsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.GetDedicatedIpsOutput{}
	fnCacher := func(out *sesv2.GetDedicatedIpsOutput, more bool) bool {
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
	if err := c.SESV2API.GetDedicatedIpsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) GetDedicatedIpsPagesWithContext(ctx context.Context, input *sesv2.GetDedicatedIpsInput, fn func(*sesv2.GetDedicatedIpsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.GetDedicatedIpsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.GetDedicatedIpsOutput{}
	fnCacher := func(out *sesv2.GetDedicatedIpsOutput, more bool) bool {
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
	if err := c.SESV2API.GetDedicatedIpsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) GetDedicatedIpsWithContext(ctx context.Context, input *sesv2.GetDedicatedIpsInput, opts ...request.Option) (*sesv2.GetDedicatedIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDedicatedIpsOutput), nil
	}
	out, err := c.SESV2API.GetDedicatedIpsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDeliverabilityDashboardOptions(input *sesv2.GetDeliverabilityDashboardOptionsInput) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDeliverabilityDashboardOptionsOutput), nil
	}
	out, err := c.SESV2API.GetDeliverabilityDashboardOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDeliverabilityDashboardOptionsWithContext(ctx context.Context, input *sesv2.GetDeliverabilityDashboardOptionsInput, opts ...request.Option) (*sesv2.GetDeliverabilityDashboardOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDeliverabilityDashboardOptionsOutput), nil
	}
	out, err := c.SESV2API.GetDeliverabilityDashboardOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDeliverabilityTestReport(input *sesv2.GetDeliverabilityTestReportInput) (*sesv2.GetDeliverabilityTestReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDeliverabilityTestReportOutput), nil
	}
	out, err := c.SESV2API.GetDeliverabilityTestReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDeliverabilityTestReportWithContext(ctx context.Context, input *sesv2.GetDeliverabilityTestReportInput, opts ...request.Option) (*sesv2.GetDeliverabilityTestReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDeliverabilityTestReportOutput), nil
	}
	out, err := c.SESV2API.GetDeliverabilityTestReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDomainDeliverabilityCampaign(input *sesv2.GetDomainDeliverabilityCampaignInput) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDomainDeliverabilityCampaignOutput), nil
	}
	out, err := c.SESV2API.GetDomainDeliverabilityCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDomainDeliverabilityCampaignWithContext(ctx context.Context, input *sesv2.GetDomainDeliverabilityCampaignInput, opts ...request.Option) (*sesv2.GetDomainDeliverabilityCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDomainDeliverabilityCampaignOutput), nil
	}
	out, err := c.SESV2API.GetDomainDeliverabilityCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDomainStatisticsReport(input *sesv2.GetDomainStatisticsReportInput) (*sesv2.GetDomainStatisticsReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDomainStatisticsReportOutput), nil
	}
	out, err := c.SESV2API.GetDomainStatisticsReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetDomainStatisticsReportWithContext(ctx context.Context, input *sesv2.GetDomainStatisticsReportInput, opts ...request.Option) (*sesv2.GetDomainStatisticsReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetDomainStatisticsReportOutput), nil
	}
	out, err := c.SESV2API.GetDomainStatisticsReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailIdentity(input *sesv2.GetEmailIdentityInput) (*sesv2.GetEmailIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailIdentityOutput), nil
	}
	out, err := c.SESV2API.GetEmailIdentity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailIdentityPolicies(input *sesv2.GetEmailIdentityPoliciesInput) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailIdentityPoliciesOutput), nil
	}
	out, err := c.SESV2API.GetEmailIdentityPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailIdentityPoliciesWithContext(ctx context.Context, input *sesv2.GetEmailIdentityPoliciesInput, opts ...request.Option) (*sesv2.GetEmailIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailIdentityPoliciesOutput), nil
	}
	out, err := c.SESV2API.GetEmailIdentityPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailIdentityWithContext(ctx context.Context, input *sesv2.GetEmailIdentityInput, opts ...request.Option) (*sesv2.GetEmailIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailIdentityOutput), nil
	}
	out, err := c.SESV2API.GetEmailIdentityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailTemplate(input *sesv2.GetEmailTemplateInput) (*sesv2.GetEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailTemplateOutput), nil
	}
	out, err := c.SESV2API.GetEmailTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetEmailTemplateWithContext(ctx context.Context, input *sesv2.GetEmailTemplateInput, opts ...request.Option) (*sesv2.GetEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetEmailTemplateOutput), nil
	}
	out, err := c.SESV2API.GetEmailTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetImportJob(input *sesv2.GetImportJobInput) (*sesv2.GetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetImportJobOutput), nil
	}
	out, err := c.SESV2API.GetImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetImportJobWithContext(ctx context.Context, input *sesv2.GetImportJobInput, opts ...request.Option) (*sesv2.GetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetImportJobOutput), nil
	}
	out, err := c.SESV2API.GetImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetSuppressedDestination(input *sesv2.GetSuppressedDestinationInput) (*sesv2.GetSuppressedDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetSuppressedDestinationOutput), nil
	}
	out, err := c.SESV2API.GetSuppressedDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) GetSuppressedDestinationWithContext(ctx context.Context, input *sesv2.GetSuppressedDestinationInput, opts ...request.Option) (*sesv2.GetSuppressedDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.GetSuppressedDestinationOutput), nil
	}
	out, err := c.SESV2API.GetSuppressedDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListConfigurationSets(input *sesv2.ListConfigurationSetsInput) (*sesv2.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListConfigurationSetsOutput), nil
	}
	out, err := c.SESV2API.ListConfigurationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListConfigurationSetsPages(input *sesv2.ListConfigurationSetsInput, fn func(*sesv2.ListConfigurationSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListConfigurationSetsOutput{}
	fnCacher := func(out *sesv2.ListConfigurationSetsOutput, more bool) bool {
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
	if err := c.SESV2API.ListConfigurationSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListConfigurationSetsPagesWithContext(ctx context.Context, input *sesv2.ListConfigurationSetsInput, fn func(*sesv2.ListConfigurationSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListConfigurationSetsOutput{}
	fnCacher := func(out *sesv2.ListConfigurationSetsOutput, more bool) bool {
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
	if err := c.SESV2API.ListConfigurationSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListConfigurationSetsWithContext(ctx context.Context, input *sesv2.ListConfigurationSetsInput, opts ...request.Option) (*sesv2.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListConfigurationSetsOutput), nil
	}
	out, err := c.SESV2API.ListConfigurationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListContactLists(input *sesv2.ListContactListsInput) (*sesv2.ListContactListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListContactListsOutput), nil
	}
	out, err := c.SESV2API.ListContactLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListContactListsPages(input *sesv2.ListContactListsInput, fn func(*sesv2.ListContactListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListContactListsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListContactListsOutput{}
	fnCacher := func(out *sesv2.ListContactListsOutput, more bool) bool {
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
	if err := c.SESV2API.ListContactListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListContactListsPagesWithContext(ctx context.Context, input *sesv2.ListContactListsInput, fn func(*sesv2.ListContactListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListContactListsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListContactListsOutput{}
	fnCacher := func(out *sesv2.ListContactListsOutput, more bool) bool {
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
	if err := c.SESV2API.ListContactListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListContactListsWithContext(ctx context.Context, input *sesv2.ListContactListsInput, opts ...request.Option) (*sesv2.ListContactListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListContactListsOutput), nil
	}
	out, err := c.SESV2API.ListContactListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListContacts(input *sesv2.ListContactsInput) (*sesv2.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListContactsOutput), nil
	}
	out, err := c.SESV2API.ListContacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListContactsPages(input *sesv2.ListContactsInput, fn func(*sesv2.ListContactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListContactsOutput{}
	fnCacher := func(out *sesv2.ListContactsOutput, more bool) bool {
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
	if err := c.SESV2API.ListContactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListContactsPagesWithContext(ctx context.Context, input *sesv2.ListContactsInput, fn func(*sesv2.ListContactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListContactsOutput{}
	fnCacher := func(out *sesv2.ListContactsOutput, more bool) bool {
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
	if err := c.SESV2API.ListContactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListContactsWithContext(ctx context.Context, input *sesv2.ListContactsInput, opts ...request.Option) (*sesv2.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListContactsOutput), nil
	}
	out, err := c.SESV2API.ListContactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListCustomVerificationEmailTemplates(input *sesv2.ListCustomVerificationEmailTemplatesInput) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListCustomVerificationEmailTemplatesOutput), nil
	}
	out, err := c.SESV2API.ListCustomVerificationEmailTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListCustomVerificationEmailTemplatesPages(input *sesv2.ListCustomVerificationEmailTemplatesInput, fn func(*sesv2.ListCustomVerificationEmailTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListCustomVerificationEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListCustomVerificationEmailTemplatesOutput{}
	fnCacher := func(out *sesv2.ListCustomVerificationEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESV2API.ListCustomVerificationEmailTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListCustomVerificationEmailTemplatesPagesWithContext(ctx context.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput, fn func(*sesv2.ListCustomVerificationEmailTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListCustomVerificationEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListCustomVerificationEmailTemplatesOutput{}
	fnCacher := func(out *sesv2.ListCustomVerificationEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESV2API.ListCustomVerificationEmailTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListCustomVerificationEmailTemplatesWithContext(ctx context.Context, input *sesv2.ListCustomVerificationEmailTemplatesInput, opts ...request.Option) (*sesv2.ListCustomVerificationEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListCustomVerificationEmailTemplatesOutput), nil
	}
	out, err := c.SESV2API.ListCustomVerificationEmailTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDedicatedIpPools(input *sesv2.ListDedicatedIpPoolsInput) (*sesv2.ListDedicatedIpPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDedicatedIpPoolsOutput), nil
	}
	out, err := c.SESV2API.ListDedicatedIpPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDedicatedIpPoolsPages(input *sesv2.ListDedicatedIpPoolsInput, fn func(*sesv2.ListDedicatedIpPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDedicatedIpPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDedicatedIpPoolsOutput{}
	fnCacher := func(out *sesv2.ListDedicatedIpPoolsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDedicatedIpPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDedicatedIpPoolsPagesWithContext(ctx context.Context, input *sesv2.ListDedicatedIpPoolsInput, fn func(*sesv2.ListDedicatedIpPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDedicatedIpPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDedicatedIpPoolsOutput{}
	fnCacher := func(out *sesv2.ListDedicatedIpPoolsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDedicatedIpPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDedicatedIpPoolsWithContext(ctx context.Context, input *sesv2.ListDedicatedIpPoolsInput, opts ...request.Option) (*sesv2.ListDedicatedIpPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDedicatedIpPoolsOutput), nil
	}
	out, err := c.SESV2API.ListDedicatedIpPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDeliverabilityTestReports(input *sesv2.ListDeliverabilityTestReportsInput) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDeliverabilityTestReportsOutput), nil
	}
	out, err := c.SESV2API.ListDeliverabilityTestReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDeliverabilityTestReportsPages(input *sesv2.ListDeliverabilityTestReportsInput, fn func(*sesv2.ListDeliverabilityTestReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDeliverabilityTestReportsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDeliverabilityTestReportsOutput{}
	fnCacher := func(out *sesv2.ListDeliverabilityTestReportsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDeliverabilityTestReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDeliverabilityTestReportsPagesWithContext(ctx context.Context, input *sesv2.ListDeliverabilityTestReportsInput, fn func(*sesv2.ListDeliverabilityTestReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDeliverabilityTestReportsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDeliverabilityTestReportsOutput{}
	fnCacher := func(out *sesv2.ListDeliverabilityTestReportsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDeliverabilityTestReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDeliverabilityTestReportsWithContext(ctx context.Context, input *sesv2.ListDeliverabilityTestReportsInput, opts ...request.Option) (*sesv2.ListDeliverabilityTestReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDeliverabilityTestReportsOutput), nil
	}
	out, err := c.SESV2API.ListDeliverabilityTestReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDomainDeliverabilityCampaigns(input *sesv2.ListDomainDeliverabilityCampaignsInput) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDomainDeliverabilityCampaignsOutput), nil
	}
	out, err := c.SESV2API.ListDomainDeliverabilityCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListDomainDeliverabilityCampaignsPages(input *sesv2.ListDomainDeliverabilityCampaignsInput, fn func(*sesv2.ListDomainDeliverabilityCampaignsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDomainDeliverabilityCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDomainDeliverabilityCampaignsOutput{}
	fnCacher := func(out *sesv2.ListDomainDeliverabilityCampaignsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDomainDeliverabilityCampaignsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDomainDeliverabilityCampaignsPagesWithContext(ctx context.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput, fn func(*sesv2.ListDomainDeliverabilityCampaignsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListDomainDeliverabilityCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListDomainDeliverabilityCampaignsOutput{}
	fnCacher := func(out *sesv2.ListDomainDeliverabilityCampaignsOutput, more bool) bool {
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
	if err := c.SESV2API.ListDomainDeliverabilityCampaignsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListDomainDeliverabilityCampaignsWithContext(ctx context.Context, input *sesv2.ListDomainDeliverabilityCampaignsInput, opts ...request.Option) (*sesv2.ListDomainDeliverabilityCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListDomainDeliverabilityCampaignsOutput), nil
	}
	out, err := c.SESV2API.ListDomainDeliverabilityCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListEmailIdentities(input *sesv2.ListEmailIdentitiesInput) (*sesv2.ListEmailIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListEmailIdentitiesOutput), nil
	}
	out, err := c.SESV2API.ListEmailIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListEmailIdentitiesPages(input *sesv2.ListEmailIdentitiesInput, fn func(*sesv2.ListEmailIdentitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListEmailIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListEmailIdentitiesOutput{}
	fnCacher := func(out *sesv2.ListEmailIdentitiesOutput, more bool) bool {
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
	if err := c.SESV2API.ListEmailIdentitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListEmailIdentitiesPagesWithContext(ctx context.Context, input *sesv2.ListEmailIdentitiesInput, fn func(*sesv2.ListEmailIdentitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListEmailIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListEmailIdentitiesOutput{}
	fnCacher := func(out *sesv2.ListEmailIdentitiesOutput, more bool) bool {
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
	if err := c.SESV2API.ListEmailIdentitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListEmailIdentitiesWithContext(ctx context.Context, input *sesv2.ListEmailIdentitiesInput, opts ...request.Option) (*sesv2.ListEmailIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListEmailIdentitiesOutput), nil
	}
	out, err := c.SESV2API.ListEmailIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListEmailTemplates(input *sesv2.ListEmailTemplatesInput) (*sesv2.ListEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListEmailTemplatesOutput), nil
	}
	out, err := c.SESV2API.ListEmailTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListEmailTemplatesPages(input *sesv2.ListEmailTemplatesInput, fn func(*sesv2.ListEmailTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListEmailTemplatesOutput{}
	fnCacher := func(out *sesv2.ListEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESV2API.ListEmailTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListEmailTemplatesPagesWithContext(ctx context.Context, input *sesv2.ListEmailTemplatesInput, fn func(*sesv2.ListEmailTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListEmailTemplatesOutput{}
	fnCacher := func(out *sesv2.ListEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESV2API.ListEmailTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListEmailTemplatesWithContext(ctx context.Context, input *sesv2.ListEmailTemplatesInput, opts ...request.Option) (*sesv2.ListEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListEmailTemplatesOutput), nil
	}
	out, err := c.SESV2API.ListEmailTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListImportJobs(input *sesv2.ListImportJobsInput) (*sesv2.ListImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListImportJobsOutput), nil
	}
	out, err := c.SESV2API.ListImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListImportJobsPages(input *sesv2.ListImportJobsInput, fn func(*sesv2.ListImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListImportJobsOutput{}
	fnCacher := func(out *sesv2.ListImportJobsOutput, more bool) bool {
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
	if err := c.SESV2API.ListImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListImportJobsPagesWithContext(ctx context.Context, input *sesv2.ListImportJobsInput, fn func(*sesv2.ListImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListImportJobsOutput{}
	fnCacher := func(out *sesv2.ListImportJobsOutput, more bool) bool {
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
	if err := c.SESV2API.ListImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListImportJobsWithContext(ctx context.Context, input *sesv2.ListImportJobsInput, opts ...request.Option) (*sesv2.ListImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListImportJobsOutput), nil
	}
	out, err := c.SESV2API.ListImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListRecommendations(input *sesv2.ListRecommendationsInput) (*sesv2.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListRecommendationsOutput), nil
	}
	out, err := c.SESV2API.ListRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListRecommendationsPages(input *sesv2.ListRecommendationsInput, fn func(*sesv2.ListRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListRecommendationsOutput{}
	fnCacher := func(out *sesv2.ListRecommendationsOutput, more bool) bool {
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
	if err := c.SESV2API.ListRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListRecommendationsPagesWithContext(ctx context.Context, input *sesv2.ListRecommendationsInput, fn func(*sesv2.ListRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListRecommendationsOutput{}
	fnCacher := func(out *sesv2.ListRecommendationsOutput, more bool) bool {
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
	if err := c.SESV2API.ListRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListRecommendationsWithContext(ctx context.Context, input *sesv2.ListRecommendationsInput, opts ...request.Option) (*sesv2.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListRecommendationsOutput), nil
	}
	out, err := c.SESV2API.ListRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListSuppressedDestinations(input *sesv2.ListSuppressedDestinationsInput) (*sesv2.ListSuppressedDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListSuppressedDestinationsOutput), nil
	}
	out, err := c.SESV2API.ListSuppressedDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListSuppressedDestinationsPages(input *sesv2.ListSuppressedDestinationsInput, fn func(*sesv2.ListSuppressedDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListSuppressedDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListSuppressedDestinationsOutput{}
	fnCacher := func(out *sesv2.ListSuppressedDestinationsOutput, more bool) bool {
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
	if err := c.SESV2API.ListSuppressedDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListSuppressedDestinationsPagesWithContext(ctx context.Context, input *sesv2.ListSuppressedDestinationsInput, fn func(*sesv2.ListSuppressedDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sesv2.ListSuppressedDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &sesv2.ListSuppressedDestinationsOutput{}
	fnCacher := func(out *sesv2.ListSuppressedDestinationsOutput, more bool) bool {
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
	if err := c.SESV2API.ListSuppressedDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SESV2) ListSuppressedDestinationsWithContext(ctx context.Context, input *sesv2.ListSuppressedDestinationsInput, opts ...request.Option) (*sesv2.ListSuppressedDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListSuppressedDestinationsOutput), nil
	}
	out, err := c.SESV2API.ListSuppressedDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListTagsForResource(input *sesv2.ListTagsForResourceInput) (*sesv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListTagsForResourceOutput), nil
	}
	out, err := c.SESV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SESV2) ListTagsForResourceWithContext(ctx context.Context, input *sesv2.ListTagsForResourceInput, opts ...request.Option) (*sesv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sesv2.ListTagsForResourceOutput), nil
	}
	out, err := c.SESV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
