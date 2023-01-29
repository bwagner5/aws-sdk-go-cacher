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
	"github.com/aws/aws-sdk-go/service/fms"
	"github.com/aws/aws-sdk-go/service/fms/fmsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type FMS struct {
	fmsiface.FMSAPI
	cache *cache.Cache
}

func NewFMS(fmsapi fmsiface.FMSAPI) *FMS {
	return &FMS{
		FMSAPI: fmsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *FMS) BatchAssociateResource(input *fms.BatchAssociateResourceInput) (*fms.BatchAssociateResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.BatchAssociateResourceOutput), nil
	}
	out, err := c.FMSAPI.BatchAssociateResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) BatchAssociateResourceWithContext(ctx context.Context, input *fms.BatchAssociateResourceInput, opts ...request.Option) (*fms.BatchAssociateResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.BatchAssociateResourceOutput), nil
	}
	out, err := c.FMSAPI.BatchAssociateResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) BatchDisassociateResource(input *fms.BatchDisassociateResourceInput) (*fms.BatchDisassociateResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.BatchDisassociateResourceOutput), nil
	}
	out, err := c.FMSAPI.BatchDisassociateResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) BatchDisassociateResourceWithContext(ctx context.Context, input *fms.BatchDisassociateResourceInput, opts ...request.Option) (*fms.BatchDisassociateResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.BatchDisassociateResourceOutput), nil
	}
	out, err := c.FMSAPI.BatchDisassociateResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetAdminAccount(input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetAdminAccountOutput), nil
	}
	out, err := c.FMSAPI.GetAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetAdminAccountWithContext(ctx context.Context, input *fms.GetAdminAccountInput, opts ...request.Option) (*fms.GetAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetAdminAccountOutput), nil
	}
	out, err := c.FMSAPI.GetAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetAppsList(input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetAppsListOutput), nil
	}
	out, err := c.FMSAPI.GetAppsList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetAppsListWithContext(ctx context.Context, input *fms.GetAppsListInput, opts ...request.Option) (*fms.GetAppsListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetAppsListOutput), nil
	}
	out, err := c.FMSAPI.GetAppsListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetComplianceDetail(input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetComplianceDetailOutput), nil
	}
	out, err := c.FMSAPI.GetComplianceDetail(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetComplianceDetailWithContext(ctx context.Context, input *fms.GetComplianceDetailInput, opts ...request.Option) (*fms.GetComplianceDetailOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetComplianceDetailOutput), nil
	}
	out, err := c.FMSAPI.GetComplianceDetailWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetNotificationChannel(input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetNotificationChannelOutput), nil
	}
	out, err := c.FMSAPI.GetNotificationChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetNotificationChannelWithContext(ctx context.Context, input *fms.GetNotificationChannelInput, opts ...request.Option) (*fms.GetNotificationChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetNotificationChannelOutput), nil
	}
	out, err := c.FMSAPI.GetNotificationChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetPolicy(input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetPolicyOutput), nil
	}
	out, err := c.FMSAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetPolicyWithContext(ctx context.Context, input *fms.GetPolicyInput, opts ...request.Option) (*fms.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetPolicyOutput), nil
	}
	out, err := c.FMSAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetProtectionStatus(input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetProtectionStatusOutput), nil
	}
	out, err := c.FMSAPI.GetProtectionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetProtectionStatusWithContext(ctx context.Context, input *fms.GetProtectionStatusInput, opts ...request.Option) (*fms.GetProtectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetProtectionStatusOutput), nil
	}
	out, err := c.FMSAPI.GetProtectionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetProtocolsList(input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetProtocolsListOutput), nil
	}
	out, err := c.FMSAPI.GetProtocolsList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetProtocolsListWithContext(ctx context.Context, input *fms.GetProtocolsListInput, opts ...request.Option) (*fms.GetProtocolsListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetProtocolsListOutput), nil
	}
	out, err := c.FMSAPI.GetProtocolsListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetResourceSet(input *fms.GetResourceSetInput) (*fms.GetResourceSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetResourceSetOutput), nil
	}
	out, err := c.FMSAPI.GetResourceSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetResourceSetWithContext(ctx context.Context, input *fms.GetResourceSetInput, opts ...request.Option) (*fms.GetResourceSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetResourceSetOutput), nil
	}
	out, err := c.FMSAPI.GetResourceSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetThirdPartyFirewallAssociationStatus(input *fms.GetThirdPartyFirewallAssociationStatusInput) (*fms.GetThirdPartyFirewallAssociationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetThirdPartyFirewallAssociationStatusOutput), nil
	}
	out, err := c.FMSAPI.GetThirdPartyFirewallAssociationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetThirdPartyFirewallAssociationStatusWithContext(ctx context.Context, input *fms.GetThirdPartyFirewallAssociationStatusInput, opts ...request.Option) (*fms.GetThirdPartyFirewallAssociationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetThirdPartyFirewallAssociationStatusOutput), nil
	}
	out, err := c.FMSAPI.GetThirdPartyFirewallAssociationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetViolationDetails(input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetViolationDetailsOutput), nil
	}
	out, err := c.FMSAPI.GetViolationDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) GetViolationDetailsWithContext(ctx context.Context, input *fms.GetViolationDetailsInput, opts ...request.Option) (*fms.GetViolationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.GetViolationDetailsOutput), nil
	}
	out, err := c.FMSAPI.GetViolationDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListAppsLists(input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListAppsListsOutput), nil
	}
	out, err := c.FMSAPI.ListAppsLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListAppsListsPages(input *fms.ListAppsListsInput, fn func(*fms.ListAppsListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListAppsListsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListAppsListsOutput{}
	fnCacher := func(out *fms.ListAppsListsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListAppsListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListAppsListsPagesWithContext(ctx context.Context, input *fms.ListAppsListsInput, fn func(*fms.ListAppsListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListAppsListsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListAppsListsOutput{}
	fnCacher := func(out *fms.ListAppsListsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListAppsListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListAppsListsWithContext(ctx context.Context, input *fms.ListAppsListsInput, opts ...request.Option) (*fms.ListAppsListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListAppsListsOutput), nil
	}
	out, err := c.FMSAPI.ListAppsListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListComplianceStatus(input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListComplianceStatusOutput), nil
	}
	out, err := c.FMSAPI.ListComplianceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListComplianceStatusPages(input *fms.ListComplianceStatusInput, fn func(*fms.ListComplianceStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListComplianceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListComplianceStatusOutput{}
	fnCacher := func(out *fms.ListComplianceStatusOutput, more bool) bool {
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
	if err := c.FMSAPI.ListComplianceStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListComplianceStatusPagesWithContext(ctx context.Context, input *fms.ListComplianceStatusInput, fn func(*fms.ListComplianceStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListComplianceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListComplianceStatusOutput{}
	fnCacher := func(out *fms.ListComplianceStatusOutput, more bool) bool {
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
	if err := c.FMSAPI.ListComplianceStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListComplianceStatusWithContext(ctx context.Context, input *fms.ListComplianceStatusInput, opts ...request.Option) (*fms.ListComplianceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListComplianceStatusOutput), nil
	}
	out, err := c.FMSAPI.ListComplianceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListDiscoveredResources(input *fms.ListDiscoveredResourcesInput) (*fms.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.FMSAPI.ListDiscoveredResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListDiscoveredResourcesWithContext(ctx context.Context, input *fms.ListDiscoveredResourcesInput, opts ...request.Option) (*fms.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.FMSAPI.ListDiscoveredResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListMemberAccounts(input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListMemberAccountsOutput), nil
	}
	out, err := c.FMSAPI.ListMemberAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListMemberAccountsPages(input *fms.ListMemberAccountsInput, fn func(*fms.ListMemberAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListMemberAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListMemberAccountsOutput{}
	fnCacher := func(out *fms.ListMemberAccountsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListMemberAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListMemberAccountsPagesWithContext(ctx context.Context, input *fms.ListMemberAccountsInput, fn func(*fms.ListMemberAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListMemberAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListMemberAccountsOutput{}
	fnCacher := func(out *fms.ListMemberAccountsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListMemberAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListMemberAccountsWithContext(ctx context.Context, input *fms.ListMemberAccountsInput, opts ...request.Option) (*fms.ListMemberAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListMemberAccountsOutput), nil
	}
	out, err := c.FMSAPI.ListMemberAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListPolicies(input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListPoliciesOutput), nil
	}
	out, err := c.FMSAPI.ListPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListPoliciesPages(input *fms.ListPoliciesInput, fn func(*fms.ListPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListPoliciesOutput{}
	fnCacher := func(out *fms.ListPoliciesOutput, more bool) bool {
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
	if err := c.FMSAPI.ListPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListPoliciesPagesWithContext(ctx context.Context, input *fms.ListPoliciesInput, fn func(*fms.ListPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListPoliciesOutput{}
	fnCacher := func(out *fms.ListPoliciesOutput, more bool) bool {
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
	if err := c.FMSAPI.ListPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListPoliciesWithContext(ctx context.Context, input *fms.ListPoliciesInput, opts ...request.Option) (*fms.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListPoliciesOutput), nil
	}
	out, err := c.FMSAPI.ListPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListProtocolsLists(input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListProtocolsListsOutput), nil
	}
	out, err := c.FMSAPI.ListProtocolsLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListProtocolsListsPages(input *fms.ListProtocolsListsInput, fn func(*fms.ListProtocolsListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListProtocolsListsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListProtocolsListsOutput{}
	fnCacher := func(out *fms.ListProtocolsListsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListProtocolsListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListProtocolsListsPagesWithContext(ctx context.Context, input *fms.ListProtocolsListsInput, fn func(*fms.ListProtocolsListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListProtocolsListsOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListProtocolsListsOutput{}
	fnCacher := func(out *fms.ListProtocolsListsOutput, more bool) bool {
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
	if err := c.FMSAPI.ListProtocolsListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListProtocolsListsWithContext(ctx context.Context, input *fms.ListProtocolsListsInput, opts ...request.Option) (*fms.ListProtocolsListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListProtocolsListsOutput), nil
	}
	out, err := c.FMSAPI.ListProtocolsListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListResourceSetResources(input *fms.ListResourceSetResourcesInput) (*fms.ListResourceSetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListResourceSetResourcesOutput), nil
	}
	out, err := c.FMSAPI.ListResourceSetResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListResourceSetResourcesWithContext(ctx context.Context, input *fms.ListResourceSetResourcesInput, opts ...request.Option) (*fms.ListResourceSetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListResourceSetResourcesOutput), nil
	}
	out, err := c.FMSAPI.ListResourceSetResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListResourceSets(input *fms.ListResourceSetsInput) (*fms.ListResourceSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListResourceSetsOutput), nil
	}
	out, err := c.FMSAPI.ListResourceSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListResourceSetsWithContext(ctx context.Context, input *fms.ListResourceSetsInput, opts ...request.Option) (*fms.ListResourceSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListResourceSetsOutput), nil
	}
	out, err := c.FMSAPI.ListResourceSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListTagsForResource(input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListTagsForResourceOutput), nil
	}
	out, err := c.FMSAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListTagsForResourceWithContext(ctx context.Context, input *fms.ListTagsForResourceInput, opts ...request.Option) (*fms.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListTagsForResourceOutput), nil
	}
	out, err := c.FMSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListThirdPartyFirewallFirewallPolicies(input *fms.ListThirdPartyFirewallFirewallPoliciesInput) (*fms.ListThirdPartyFirewallFirewallPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListThirdPartyFirewallFirewallPoliciesOutput), nil
	}
	out, err := c.FMSAPI.ListThirdPartyFirewallFirewallPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FMS) ListThirdPartyFirewallFirewallPoliciesPages(input *fms.ListThirdPartyFirewallFirewallPoliciesInput, fn func(*fms.ListThirdPartyFirewallFirewallPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListThirdPartyFirewallFirewallPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListThirdPartyFirewallFirewallPoliciesOutput{}
	fnCacher := func(out *fms.ListThirdPartyFirewallFirewallPoliciesOutput, more bool) bool {
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
	if err := c.FMSAPI.ListThirdPartyFirewallFirewallPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListThirdPartyFirewallFirewallPoliciesPagesWithContext(ctx context.Context, input *fms.ListThirdPartyFirewallFirewallPoliciesInput, fn func(*fms.ListThirdPartyFirewallFirewallPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fms.ListThirdPartyFirewallFirewallPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &fms.ListThirdPartyFirewallFirewallPoliciesOutput{}
	fnCacher := func(out *fms.ListThirdPartyFirewallFirewallPoliciesOutput, more bool) bool {
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
	if err := c.FMSAPI.ListThirdPartyFirewallFirewallPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FMS) ListThirdPartyFirewallFirewallPoliciesWithContext(ctx context.Context, input *fms.ListThirdPartyFirewallFirewallPoliciesInput, opts ...request.Option) (*fms.ListThirdPartyFirewallFirewallPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fms.ListThirdPartyFirewallFirewallPoliciesOutput), nil
	}
	out, err := c.FMSAPI.ListThirdPartyFirewallFirewallPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
