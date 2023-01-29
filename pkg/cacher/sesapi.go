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
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/ses/sesiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SES struct {
	sesiface.SESAPI
	cache *cache.Cache
}

func NewSES(sesapi sesiface.SESAPI) *SES {
	return &SES{
		SESAPI: sesapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SES) DescribeActiveReceiptRuleSet(input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeActiveReceiptRuleSetOutput), nil
	}
	out, err := c.SESAPI.DescribeActiveReceiptRuleSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeActiveReceiptRuleSetWithContext(ctx context.Context, input *ses.DescribeActiveReceiptRuleSetInput, opts ...request.Option) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeActiveReceiptRuleSetOutput), nil
	}
	out, err := c.SESAPI.DescribeActiveReceiptRuleSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeConfigurationSet(input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeConfigurationSetOutput), nil
	}
	out, err := c.SESAPI.DescribeConfigurationSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeConfigurationSetWithContext(ctx context.Context, input *ses.DescribeConfigurationSetInput, opts ...request.Option) (*ses.DescribeConfigurationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeConfigurationSetOutput), nil
	}
	out, err := c.SESAPI.DescribeConfigurationSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeReceiptRule(input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeReceiptRuleOutput), nil
	}
	out, err := c.SESAPI.DescribeReceiptRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeReceiptRuleSet(input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeReceiptRuleSetOutput), nil
	}
	out, err := c.SESAPI.DescribeReceiptRuleSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeReceiptRuleSetWithContext(ctx context.Context, input *ses.DescribeReceiptRuleSetInput, opts ...request.Option) (*ses.DescribeReceiptRuleSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeReceiptRuleSetOutput), nil
	}
	out, err := c.SESAPI.DescribeReceiptRuleSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) DescribeReceiptRuleWithContext(ctx context.Context, input *ses.DescribeReceiptRuleInput, opts ...request.Option) (*ses.DescribeReceiptRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.DescribeReceiptRuleOutput), nil
	}
	out, err := c.SESAPI.DescribeReceiptRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetAccountSendingEnabled(input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetAccountSendingEnabledOutput), nil
	}
	out, err := c.SESAPI.GetAccountSendingEnabled(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetAccountSendingEnabledWithContext(ctx context.Context, input *ses.GetAccountSendingEnabledInput, opts ...request.Option) (*ses.GetAccountSendingEnabledOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetAccountSendingEnabledOutput), nil
	}
	out, err := c.SESAPI.GetAccountSendingEnabledWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetCustomVerificationEmailTemplate(input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetCustomVerificationEmailTemplateOutput), nil
	}
	out, err := c.SESAPI.GetCustomVerificationEmailTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetCustomVerificationEmailTemplateWithContext(ctx context.Context, input *ses.GetCustomVerificationEmailTemplateInput, opts ...request.Option) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetCustomVerificationEmailTemplateOutput), nil
	}
	out, err := c.SESAPI.GetCustomVerificationEmailTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityDkimAttributes(input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityDkimAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityDkimAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityDkimAttributesWithContext(ctx context.Context, input *ses.GetIdentityDkimAttributesInput, opts ...request.Option) (*ses.GetIdentityDkimAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityDkimAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityDkimAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityMailFromDomainAttributes(input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityMailFromDomainAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityMailFromDomainAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityMailFromDomainAttributesWithContext(ctx context.Context, input *ses.GetIdentityMailFromDomainAttributesInput, opts ...request.Option) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityMailFromDomainAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityMailFromDomainAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityNotificationAttributes(input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityNotificationAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityNotificationAttributesWithContext(ctx context.Context, input *ses.GetIdentityNotificationAttributesInput, opts ...request.Option) (*ses.GetIdentityNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityNotificationAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityPolicies(input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityPoliciesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityPoliciesWithContext(ctx context.Context, input *ses.GetIdentityPoliciesInput, opts ...request.Option) (*ses.GetIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityPoliciesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityVerificationAttributes(input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityVerificationAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityVerificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetIdentityVerificationAttributesWithContext(ctx context.Context, input *ses.GetIdentityVerificationAttributesInput, opts ...request.Option) (*ses.GetIdentityVerificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetIdentityVerificationAttributesOutput), nil
	}
	out, err := c.SESAPI.GetIdentityVerificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetSendQuota(input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetSendQuotaOutput), nil
	}
	out, err := c.SESAPI.GetSendQuota(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetSendQuotaWithContext(ctx context.Context, input *ses.GetSendQuotaInput, opts ...request.Option) (*ses.GetSendQuotaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetSendQuotaOutput), nil
	}
	out, err := c.SESAPI.GetSendQuotaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetSendStatistics(input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetSendStatisticsOutput), nil
	}
	out, err := c.SESAPI.GetSendStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetSendStatisticsWithContext(ctx context.Context, input *ses.GetSendStatisticsInput, opts ...request.Option) (*ses.GetSendStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetSendStatisticsOutput), nil
	}
	out, err := c.SESAPI.GetSendStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetTemplate(input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetTemplateOutput), nil
	}
	out, err := c.SESAPI.GetTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) GetTemplateWithContext(ctx context.Context, input *ses.GetTemplateInput, opts ...request.Option) (*ses.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.GetTemplateOutput), nil
	}
	out, err := c.SESAPI.GetTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListConfigurationSets(input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListConfigurationSetsOutput), nil
	}
	out, err := c.SESAPI.ListConfigurationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListConfigurationSetsWithContext(ctx context.Context, input *ses.ListConfigurationSetsInput, opts ...request.Option) (*ses.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListConfigurationSetsOutput), nil
	}
	out, err := c.SESAPI.ListConfigurationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListCustomVerificationEmailTemplates(input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListCustomVerificationEmailTemplatesOutput), nil
	}
	out, err := c.SESAPI.ListCustomVerificationEmailTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListCustomVerificationEmailTemplatesPages(input *ses.ListCustomVerificationEmailTemplatesInput, fn func(*ses.ListCustomVerificationEmailTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ses.ListCustomVerificationEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ses.ListCustomVerificationEmailTemplatesOutput{}
	fnCacher := func(out *ses.ListCustomVerificationEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESAPI.ListCustomVerificationEmailTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SES) ListCustomVerificationEmailTemplatesPagesWithContext(ctx context.Context, input *ses.ListCustomVerificationEmailTemplatesInput, fn func(*ses.ListCustomVerificationEmailTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ses.ListCustomVerificationEmailTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ses.ListCustomVerificationEmailTemplatesOutput{}
	fnCacher := func(out *ses.ListCustomVerificationEmailTemplatesOutput, more bool) bool {
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
	if err := c.SESAPI.ListCustomVerificationEmailTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SES) ListCustomVerificationEmailTemplatesWithContext(ctx context.Context, input *ses.ListCustomVerificationEmailTemplatesInput, opts ...request.Option) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListCustomVerificationEmailTemplatesOutput), nil
	}
	out, err := c.SESAPI.ListCustomVerificationEmailTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListIdentities(input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListIdentitiesOutput), nil
	}
	out, err := c.SESAPI.ListIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListIdentitiesPages(input *ses.ListIdentitiesInput, fn func(*ses.ListIdentitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ses.ListIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &ses.ListIdentitiesOutput{}
	fnCacher := func(out *ses.ListIdentitiesOutput, more bool) bool {
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
	if err := c.SESAPI.ListIdentitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SES) ListIdentitiesPagesWithContext(ctx context.Context, input *ses.ListIdentitiesInput, fn func(*ses.ListIdentitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ses.ListIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &ses.ListIdentitiesOutput{}
	fnCacher := func(out *ses.ListIdentitiesOutput, more bool) bool {
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
	if err := c.SESAPI.ListIdentitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SES) ListIdentitiesWithContext(ctx context.Context, input *ses.ListIdentitiesInput, opts ...request.Option) (*ses.ListIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListIdentitiesOutput), nil
	}
	out, err := c.SESAPI.ListIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListIdentityPolicies(input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListIdentityPoliciesOutput), nil
	}
	out, err := c.SESAPI.ListIdentityPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListIdentityPoliciesWithContext(ctx context.Context, input *ses.ListIdentityPoliciesInput, opts ...request.Option) (*ses.ListIdentityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListIdentityPoliciesOutput), nil
	}
	out, err := c.SESAPI.ListIdentityPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListReceiptFilters(input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListReceiptFiltersOutput), nil
	}
	out, err := c.SESAPI.ListReceiptFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListReceiptFiltersWithContext(ctx context.Context, input *ses.ListReceiptFiltersInput, opts ...request.Option) (*ses.ListReceiptFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListReceiptFiltersOutput), nil
	}
	out, err := c.SESAPI.ListReceiptFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListReceiptRuleSets(input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListReceiptRuleSetsOutput), nil
	}
	out, err := c.SESAPI.ListReceiptRuleSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListReceiptRuleSetsWithContext(ctx context.Context, input *ses.ListReceiptRuleSetsInput, opts ...request.Option) (*ses.ListReceiptRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListReceiptRuleSetsOutput), nil
	}
	out, err := c.SESAPI.ListReceiptRuleSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListTemplates(input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListTemplatesOutput), nil
	}
	out, err := c.SESAPI.ListTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListTemplatesWithContext(ctx context.Context, input *ses.ListTemplatesInput, opts ...request.Option) (*ses.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListTemplatesOutput), nil
	}
	out, err := c.SESAPI.ListTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListVerifiedEmailAddresses(input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListVerifiedEmailAddressesOutput), nil
	}
	out, err := c.SESAPI.ListVerifiedEmailAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SES) ListVerifiedEmailAddressesWithContext(ctx context.Context, input *ses.ListVerifiedEmailAddressesInput, opts ...request.Option) (*ses.ListVerifiedEmailAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ses.ListVerifiedEmailAddressesOutput), nil
	}
	out, err := c.SESAPI.ListVerifiedEmailAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
