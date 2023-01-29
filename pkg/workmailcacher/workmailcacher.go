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
package workmailcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workmail"
	"github.com/aws/aws-sdk-go/service/workmail/workmailiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkMail struct {
	workmailiface.WorkMailAPI
	cache *cache.Cache
}

func New(workmailapi workmailiface.WorkMailAPI) *WorkMail {
	return &WorkMail{
		WorkMailAPI: workmailapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkMail) DescribeEmailMonitoringConfiguration(input *workmail.DescribeEmailMonitoringConfigurationInput) (*workmail.DescribeEmailMonitoringConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeEmailMonitoringConfigurationOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeEmailMonitoringConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeEmailMonitoringConfigurationWithContext(ctx context.Context, input *workmail.DescribeEmailMonitoringConfigurationInput, opts ...request.Option) (*workmail.DescribeEmailMonitoringConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeEmailMonitoringConfigurationOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeEmailMonitoringConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeGroup(input *workmail.DescribeGroupInput) (*workmail.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeGroupOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeGroupWithContext(ctx context.Context, input *workmail.DescribeGroupInput, opts ...request.Option) (*workmail.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeGroupOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeInboundDmarcSettings(input *workmail.DescribeInboundDmarcSettingsInput) (*workmail.DescribeInboundDmarcSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeInboundDmarcSettingsOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeInboundDmarcSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeInboundDmarcSettingsWithContext(ctx context.Context, input *workmail.DescribeInboundDmarcSettingsInput, opts ...request.Option) (*workmail.DescribeInboundDmarcSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeInboundDmarcSettingsOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeInboundDmarcSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeMailboxExportJob(input *workmail.DescribeMailboxExportJobInput) (*workmail.DescribeMailboxExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeMailboxExportJobOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeMailboxExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeMailboxExportJobWithContext(ctx context.Context, input *workmail.DescribeMailboxExportJobInput, opts ...request.Option) (*workmail.DescribeMailboxExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeMailboxExportJobOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeMailboxExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeOrganization(input *workmail.DescribeOrganizationInput) (*workmail.DescribeOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeOrganizationOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeOrganizationWithContext(ctx context.Context, input *workmail.DescribeOrganizationInput, opts ...request.Option) (*workmail.DescribeOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeOrganizationOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeResource(input *workmail.DescribeResourceInput) (*workmail.DescribeResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeResourceOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeResourceWithContext(ctx context.Context, input *workmail.DescribeResourceInput, opts ...request.Option) (*workmail.DescribeResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeResourceOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeUser(input *workmail.DescribeUserInput) (*workmail.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeUserOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) DescribeUserWithContext(ctx context.Context, input *workmail.DescribeUserInput, opts ...request.Option) (*workmail.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.DescribeUserOutput), nil
	}
	out, err := c.WorkMailAPI.DescribeUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetAccessControlEffect(input *workmail.GetAccessControlEffectInput) (*workmail.GetAccessControlEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetAccessControlEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetAccessControlEffect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetAccessControlEffectWithContext(ctx context.Context, input *workmail.GetAccessControlEffectInput, opts ...request.Option) (*workmail.GetAccessControlEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetAccessControlEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetAccessControlEffectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetDefaultRetentionPolicy(input *workmail.GetDefaultRetentionPolicyInput) (*workmail.GetDefaultRetentionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetDefaultRetentionPolicyOutput), nil
	}
	out, err := c.WorkMailAPI.GetDefaultRetentionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetDefaultRetentionPolicyWithContext(ctx context.Context, input *workmail.GetDefaultRetentionPolicyInput, opts ...request.Option) (*workmail.GetDefaultRetentionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetDefaultRetentionPolicyOutput), nil
	}
	out, err := c.WorkMailAPI.GetDefaultRetentionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetImpersonationRole(input *workmail.GetImpersonationRoleInput) (*workmail.GetImpersonationRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetImpersonationRoleOutput), nil
	}
	out, err := c.WorkMailAPI.GetImpersonationRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetImpersonationRoleEffect(input *workmail.GetImpersonationRoleEffectInput) (*workmail.GetImpersonationRoleEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetImpersonationRoleEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetImpersonationRoleEffect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetImpersonationRoleEffectWithContext(ctx context.Context, input *workmail.GetImpersonationRoleEffectInput, opts ...request.Option) (*workmail.GetImpersonationRoleEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetImpersonationRoleEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetImpersonationRoleEffectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetImpersonationRoleWithContext(ctx context.Context, input *workmail.GetImpersonationRoleInput, opts ...request.Option) (*workmail.GetImpersonationRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetImpersonationRoleOutput), nil
	}
	out, err := c.WorkMailAPI.GetImpersonationRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMailDomain(input *workmail.GetMailDomainInput) (*workmail.GetMailDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMailDomainOutput), nil
	}
	out, err := c.WorkMailAPI.GetMailDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMailDomainWithContext(ctx context.Context, input *workmail.GetMailDomainInput, opts ...request.Option) (*workmail.GetMailDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMailDomainOutput), nil
	}
	out, err := c.WorkMailAPI.GetMailDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMailboxDetails(input *workmail.GetMailboxDetailsInput) (*workmail.GetMailboxDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMailboxDetailsOutput), nil
	}
	out, err := c.WorkMailAPI.GetMailboxDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMailboxDetailsWithContext(ctx context.Context, input *workmail.GetMailboxDetailsInput, opts ...request.Option) (*workmail.GetMailboxDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMailboxDetailsOutput), nil
	}
	out, err := c.WorkMailAPI.GetMailboxDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMobileDeviceAccessEffect(input *workmail.GetMobileDeviceAccessEffectInput) (*workmail.GetMobileDeviceAccessEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMobileDeviceAccessEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetMobileDeviceAccessEffect(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMobileDeviceAccessEffectWithContext(ctx context.Context, input *workmail.GetMobileDeviceAccessEffectInput, opts ...request.Option) (*workmail.GetMobileDeviceAccessEffectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMobileDeviceAccessEffectOutput), nil
	}
	out, err := c.WorkMailAPI.GetMobileDeviceAccessEffectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMobileDeviceAccessOverride(input *workmail.GetMobileDeviceAccessOverrideInput) (*workmail.GetMobileDeviceAccessOverrideOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMobileDeviceAccessOverrideOutput), nil
	}
	out, err := c.WorkMailAPI.GetMobileDeviceAccessOverride(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) GetMobileDeviceAccessOverrideWithContext(ctx context.Context, input *workmail.GetMobileDeviceAccessOverrideInput, opts ...request.Option) (*workmail.GetMobileDeviceAccessOverrideOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.GetMobileDeviceAccessOverrideOutput), nil
	}
	out, err := c.WorkMailAPI.GetMobileDeviceAccessOverrideWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAccessControlRules(input *workmail.ListAccessControlRulesInput) (*workmail.ListAccessControlRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAccessControlRulesOutput), nil
	}
	out, err := c.WorkMailAPI.ListAccessControlRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAccessControlRulesWithContext(ctx context.Context, input *workmail.ListAccessControlRulesInput, opts ...request.Option) (*workmail.ListAccessControlRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAccessControlRulesOutput), nil
	}
	out, err := c.WorkMailAPI.ListAccessControlRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAliases(input *workmail.ListAliasesInput) (*workmail.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAliasesOutput), nil
	}
	out, err := c.WorkMailAPI.ListAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAliasesPages(input *workmail.ListAliasesInput, fn func(*workmail.ListAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListAliasesOutput{}
	fnCacher := func(out *workmail.ListAliasesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListAliasesPagesWithContext(ctx context.Context, input *workmail.ListAliasesInput, fn func(*workmail.ListAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListAliasesOutput{}
	fnCacher := func(out *workmail.ListAliasesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListAliasesWithContext(ctx context.Context, input *workmail.ListAliasesInput, opts ...request.Option) (*workmail.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAliasesOutput), nil
	}
	out, err := c.WorkMailAPI.ListAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAvailabilityConfigurations(input *workmail.ListAvailabilityConfigurationsInput) (*workmail.ListAvailabilityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAvailabilityConfigurationsOutput), nil
	}
	out, err := c.WorkMailAPI.ListAvailabilityConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListAvailabilityConfigurationsPages(input *workmail.ListAvailabilityConfigurationsInput, fn func(*workmail.ListAvailabilityConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListAvailabilityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListAvailabilityConfigurationsOutput{}
	fnCacher := func(out *workmail.ListAvailabilityConfigurationsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListAvailabilityConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListAvailabilityConfigurationsPagesWithContext(ctx context.Context, input *workmail.ListAvailabilityConfigurationsInput, fn func(*workmail.ListAvailabilityConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListAvailabilityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListAvailabilityConfigurationsOutput{}
	fnCacher := func(out *workmail.ListAvailabilityConfigurationsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListAvailabilityConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListAvailabilityConfigurationsWithContext(ctx context.Context, input *workmail.ListAvailabilityConfigurationsInput, opts ...request.Option) (*workmail.ListAvailabilityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListAvailabilityConfigurationsOutput), nil
	}
	out, err := c.WorkMailAPI.ListAvailabilityConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListGroupMembers(input *workmail.ListGroupMembersInput) (*workmail.ListGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListGroupMembersOutput), nil
	}
	out, err := c.WorkMailAPI.ListGroupMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListGroupMembersPages(input *workmail.ListGroupMembersInput, fn func(*workmail.ListGroupMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListGroupMembersOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListGroupMembersOutput{}
	fnCacher := func(out *workmail.ListGroupMembersOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListGroupMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListGroupMembersPagesWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, fn func(*workmail.ListGroupMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListGroupMembersOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListGroupMembersOutput{}
	fnCacher := func(out *workmail.ListGroupMembersOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListGroupMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListGroupMembersWithContext(ctx context.Context, input *workmail.ListGroupMembersInput, opts ...request.Option) (*workmail.ListGroupMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListGroupMembersOutput), nil
	}
	out, err := c.WorkMailAPI.ListGroupMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListGroups(input *workmail.ListGroupsInput) (*workmail.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListGroupsOutput), nil
	}
	out, err := c.WorkMailAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListGroupsPages(input *workmail.ListGroupsInput, fn func(*workmail.ListGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListGroupsOutput{}
	fnCacher := func(out *workmail.ListGroupsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListGroupsPagesWithContext(ctx context.Context, input *workmail.ListGroupsInput, fn func(*workmail.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListGroupsOutput{}
	fnCacher := func(out *workmail.ListGroupsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListGroupsWithContext(ctx context.Context, input *workmail.ListGroupsInput, opts ...request.Option) (*workmail.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListGroupsOutput), nil
	}
	out, err := c.WorkMailAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListImpersonationRoles(input *workmail.ListImpersonationRolesInput) (*workmail.ListImpersonationRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListImpersonationRolesOutput), nil
	}
	out, err := c.WorkMailAPI.ListImpersonationRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListImpersonationRolesPages(input *workmail.ListImpersonationRolesInput, fn func(*workmail.ListImpersonationRolesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListImpersonationRolesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListImpersonationRolesOutput{}
	fnCacher := func(out *workmail.ListImpersonationRolesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListImpersonationRolesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListImpersonationRolesPagesWithContext(ctx context.Context, input *workmail.ListImpersonationRolesInput, fn func(*workmail.ListImpersonationRolesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListImpersonationRolesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListImpersonationRolesOutput{}
	fnCacher := func(out *workmail.ListImpersonationRolesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListImpersonationRolesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListImpersonationRolesWithContext(ctx context.Context, input *workmail.ListImpersonationRolesInput, opts ...request.Option) (*workmail.ListImpersonationRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListImpersonationRolesOutput), nil
	}
	out, err := c.WorkMailAPI.ListImpersonationRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailDomains(input *workmail.ListMailDomainsInput) (*workmail.ListMailDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailDomainsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailDomainsPages(input *workmail.ListMailDomainsInput, fn func(*workmail.ListMailDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailDomainsOutput{}
	fnCacher := func(out *workmail.ListMailDomainsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailDomainsPagesWithContext(ctx context.Context, input *workmail.ListMailDomainsInput, fn func(*workmail.ListMailDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailDomainsOutput{}
	fnCacher := func(out *workmail.ListMailDomainsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailDomainsWithContext(ctx context.Context, input *workmail.ListMailDomainsInput, opts ...request.Option) (*workmail.ListMailDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailDomainsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailboxExportJobs(input *workmail.ListMailboxExportJobsInput) (*workmail.ListMailboxExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailboxExportJobsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailboxExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailboxExportJobsPages(input *workmail.ListMailboxExportJobsInput, fn func(*workmail.ListMailboxExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailboxExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailboxExportJobsOutput{}
	fnCacher := func(out *workmail.ListMailboxExportJobsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailboxExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailboxExportJobsPagesWithContext(ctx context.Context, input *workmail.ListMailboxExportJobsInput, fn func(*workmail.ListMailboxExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailboxExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailboxExportJobsOutput{}
	fnCacher := func(out *workmail.ListMailboxExportJobsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailboxExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailboxExportJobsWithContext(ctx context.Context, input *workmail.ListMailboxExportJobsInput, opts ...request.Option) (*workmail.ListMailboxExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailboxExportJobsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailboxExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailboxPermissions(input *workmail.ListMailboxPermissionsInput) (*workmail.ListMailboxPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailboxPermissionsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailboxPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMailboxPermissionsPages(input *workmail.ListMailboxPermissionsInput, fn func(*workmail.ListMailboxPermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailboxPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailboxPermissionsOutput{}
	fnCacher := func(out *workmail.ListMailboxPermissionsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailboxPermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailboxPermissionsPagesWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, fn func(*workmail.ListMailboxPermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMailboxPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMailboxPermissionsOutput{}
	fnCacher := func(out *workmail.ListMailboxPermissionsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMailboxPermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMailboxPermissionsWithContext(ctx context.Context, input *workmail.ListMailboxPermissionsInput, opts ...request.Option) (*workmail.ListMailboxPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMailboxPermissionsOutput), nil
	}
	out, err := c.WorkMailAPI.ListMailboxPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMobileDeviceAccessOverrides(input *workmail.ListMobileDeviceAccessOverridesInput) (*workmail.ListMobileDeviceAccessOverridesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMobileDeviceAccessOverridesOutput), nil
	}
	out, err := c.WorkMailAPI.ListMobileDeviceAccessOverrides(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMobileDeviceAccessOverridesPages(input *workmail.ListMobileDeviceAccessOverridesInput, fn func(*workmail.ListMobileDeviceAccessOverridesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMobileDeviceAccessOverridesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMobileDeviceAccessOverridesOutput{}
	fnCacher := func(out *workmail.ListMobileDeviceAccessOverridesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMobileDeviceAccessOverridesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMobileDeviceAccessOverridesPagesWithContext(ctx context.Context, input *workmail.ListMobileDeviceAccessOverridesInput, fn func(*workmail.ListMobileDeviceAccessOverridesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListMobileDeviceAccessOverridesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListMobileDeviceAccessOverridesOutput{}
	fnCacher := func(out *workmail.ListMobileDeviceAccessOverridesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListMobileDeviceAccessOverridesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListMobileDeviceAccessOverridesWithContext(ctx context.Context, input *workmail.ListMobileDeviceAccessOverridesInput, opts ...request.Option) (*workmail.ListMobileDeviceAccessOverridesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMobileDeviceAccessOverridesOutput), nil
	}
	out, err := c.WorkMailAPI.ListMobileDeviceAccessOverridesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMobileDeviceAccessRules(input *workmail.ListMobileDeviceAccessRulesInput) (*workmail.ListMobileDeviceAccessRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMobileDeviceAccessRulesOutput), nil
	}
	out, err := c.WorkMailAPI.ListMobileDeviceAccessRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListMobileDeviceAccessRulesWithContext(ctx context.Context, input *workmail.ListMobileDeviceAccessRulesInput, opts ...request.Option) (*workmail.ListMobileDeviceAccessRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListMobileDeviceAccessRulesOutput), nil
	}
	out, err := c.WorkMailAPI.ListMobileDeviceAccessRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListOrganizations(input *workmail.ListOrganizationsInput) (*workmail.ListOrganizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListOrganizationsOutput), nil
	}
	out, err := c.WorkMailAPI.ListOrganizations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListOrganizationsPages(input *workmail.ListOrganizationsInput, fn func(*workmail.ListOrganizationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListOrganizationsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListOrganizationsOutput{}
	fnCacher := func(out *workmail.ListOrganizationsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListOrganizationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListOrganizationsPagesWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, fn func(*workmail.ListOrganizationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListOrganizationsOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListOrganizationsOutput{}
	fnCacher := func(out *workmail.ListOrganizationsOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListOrganizationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListOrganizationsWithContext(ctx context.Context, input *workmail.ListOrganizationsInput, opts ...request.Option) (*workmail.ListOrganizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListOrganizationsOutput), nil
	}
	out, err := c.WorkMailAPI.ListOrganizationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListResourceDelegates(input *workmail.ListResourceDelegatesInput) (*workmail.ListResourceDelegatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListResourceDelegatesOutput), nil
	}
	out, err := c.WorkMailAPI.ListResourceDelegates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListResourceDelegatesPages(input *workmail.ListResourceDelegatesInput, fn func(*workmail.ListResourceDelegatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListResourceDelegatesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListResourceDelegatesOutput{}
	fnCacher := func(out *workmail.ListResourceDelegatesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListResourceDelegatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListResourceDelegatesPagesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, fn func(*workmail.ListResourceDelegatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListResourceDelegatesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListResourceDelegatesOutput{}
	fnCacher := func(out *workmail.ListResourceDelegatesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListResourceDelegatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListResourceDelegatesWithContext(ctx context.Context, input *workmail.ListResourceDelegatesInput, opts ...request.Option) (*workmail.ListResourceDelegatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListResourceDelegatesOutput), nil
	}
	out, err := c.WorkMailAPI.ListResourceDelegatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListResources(input *workmail.ListResourcesInput) (*workmail.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListResourcesOutput), nil
	}
	out, err := c.WorkMailAPI.ListResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListResourcesPages(input *workmail.ListResourcesInput, fn func(*workmail.ListResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListResourcesOutput{}
	fnCacher := func(out *workmail.ListResourcesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListResourcesPagesWithContext(ctx context.Context, input *workmail.ListResourcesInput, fn func(*workmail.ListResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListResourcesOutput{}
	fnCacher := func(out *workmail.ListResourcesOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListResourcesWithContext(ctx context.Context, input *workmail.ListResourcesInput, opts ...request.Option) (*workmail.ListResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListResourcesOutput), nil
	}
	out, err := c.WorkMailAPI.ListResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListTagsForResource(input *workmail.ListTagsForResourceInput) (*workmail.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkMailAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListTagsForResourceWithContext(ctx context.Context, input *workmail.ListTagsForResourceInput, opts ...request.Option) (*workmail.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkMailAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListUsers(input *workmail.ListUsersInput) (*workmail.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListUsersOutput), nil
	}
	out, err := c.WorkMailAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkMail) ListUsersPages(input *workmail.ListUsersInput, fn func(*workmail.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListUsersOutput{}
	fnCacher := func(out *workmail.ListUsersOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListUsersPagesWithContext(ctx context.Context, input *workmail.ListUsersInput, fn func(*workmail.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workmail.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &workmail.ListUsersOutput{}
	fnCacher := func(out *workmail.ListUsersOutput, more bool) bool {
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
	if err := c.WorkMailAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkMail) ListUsersWithContext(ctx context.Context, input *workmail.ListUsersInput, opts ...request.Option) (*workmail.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workmail.ListUsersOutput), nil
	}
	out, err := c.WorkMailAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
