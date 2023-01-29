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
package worklinkcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/worklink"
	"github.com/aws/aws-sdk-go/service/worklink/worklinkiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkLink struct {
	worklinkiface.WorkLinkAPI
	cache *cache.Cache
}

func New(worklinkapi worklinkiface.WorkLinkAPI) *WorkLink {
	return &WorkLink{
		WorkLinkAPI: worklinkapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkLink) DescribeAuditStreamConfiguration(input *worklink.DescribeAuditStreamConfigurationInput) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeAuditStreamConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeAuditStreamConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeAuditStreamConfigurationWithContext(ctx context.Context, input *worklink.DescribeAuditStreamConfigurationInput, opts ...request.Option) (*worklink.DescribeAuditStreamConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeAuditStreamConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeAuditStreamConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeCompanyNetworkConfiguration(input *worklink.DescribeCompanyNetworkConfigurationInput) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeCompanyNetworkConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeCompanyNetworkConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeCompanyNetworkConfigurationWithContext(ctx context.Context, input *worklink.DescribeCompanyNetworkConfigurationInput, opts ...request.Option) (*worklink.DescribeCompanyNetworkConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeCompanyNetworkConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeCompanyNetworkConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDevice(input *worklink.DescribeDeviceInput) (*worklink.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDeviceOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDevicePolicyConfiguration(input *worklink.DescribeDevicePolicyConfigurationInput) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDevicePolicyConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDevicePolicyConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDevicePolicyConfigurationWithContext(ctx context.Context, input *worklink.DescribeDevicePolicyConfigurationInput, opts ...request.Option) (*worklink.DescribeDevicePolicyConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDevicePolicyConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDevicePolicyConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDeviceWithContext(ctx context.Context, input *worklink.DescribeDeviceInput, opts ...request.Option) (*worklink.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDeviceOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDomain(input *worklink.DescribeDomainInput) (*worklink.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDomainOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeDomainWithContext(ctx context.Context, input *worklink.DescribeDomainInput, opts ...request.Option) (*worklink.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeDomainOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeFleetMetadata(input *worklink.DescribeFleetMetadataInput) (*worklink.DescribeFleetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeFleetMetadataOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeFleetMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeFleetMetadataWithContext(ctx context.Context, input *worklink.DescribeFleetMetadataInput, opts ...request.Option) (*worklink.DescribeFleetMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeFleetMetadataOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeFleetMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeIdentityProviderConfiguration(input *worklink.DescribeIdentityProviderConfigurationInput) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeIdentityProviderConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeIdentityProviderConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeIdentityProviderConfigurationWithContext(ctx context.Context, input *worklink.DescribeIdentityProviderConfigurationInput, opts ...request.Option) (*worklink.DescribeIdentityProviderConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeIdentityProviderConfigurationOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeIdentityProviderConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeWebsiteCertificateAuthority(input *worklink.DescribeWebsiteCertificateAuthorityInput) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeWebsiteCertificateAuthorityOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeWebsiteCertificateAuthority(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) DescribeWebsiteCertificateAuthorityWithContext(ctx context.Context, input *worklink.DescribeWebsiteCertificateAuthorityInput, opts ...request.Option) (*worklink.DescribeWebsiteCertificateAuthorityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.DescribeWebsiteCertificateAuthorityOutput), nil
	}
	out, err := c.WorkLinkAPI.DescribeWebsiteCertificateAuthorityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListDevices(input *worklink.ListDevicesInput) (*worklink.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListDevicesOutput), nil
	}
	out, err := c.WorkLinkAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListDevicesPages(input *worklink.ListDevicesInput, fn func(*worklink.ListDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListDevicesOutput{}
	fnCacher := func(out *worklink.ListDevicesOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListDevicesPagesWithContext(ctx context.Context, input *worklink.ListDevicesInput, fn func(*worklink.ListDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListDevicesOutput{}
	fnCacher := func(out *worklink.ListDevicesOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListDevicesWithContext(ctx context.Context, input *worklink.ListDevicesInput, opts ...request.Option) (*worklink.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListDevicesOutput), nil
	}
	out, err := c.WorkLinkAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListDomains(input *worklink.ListDomainsInput) (*worklink.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListDomainsOutput), nil
	}
	out, err := c.WorkLinkAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListDomainsPages(input *worklink.ListDomainsInput, fn func(*worklink.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListDomainsOutput{}
	fnCacher := func(out *worklink.ListDomainsOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListDomainsPagesWithContext(ctx context.Context, input *worklink.ListDomainsInput, fn func(*worklink.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListDomainsOutput{}
	fnCacher := func(out *worklink.ListDomainsOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListDomainsWithContext(ctx context.Context, input *worklink.ListDomainsInput, opts ...request.Option) (*worklink.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListDomainsOutput), nil
	}
	out, err := c.WorkLinkAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListFleets(input *worklink.ListFleetsInput) (*worklink.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListFleetsOutput), nil
	}
	out, err := c.WorkLinkAPI.ListFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListFleetsPages(input *worklink.ListFleetsInput, fn func(*worklink.ListFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListFleetsOutput{}
	fnCacher := func(out *worklink.ListFleetsOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListFleetsPagesWithContext(ctx context.Context, input *worklink.ListFleetsInput, fn func(*worklink.ListFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListFleetsOutput{}
	fnCacher := func(out *worklink.ListFleetsOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListFleetsWithContext(ctx context.Context, input *worklink.ListFleetsInput, opts ...request.Option) (*worklink.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListFleetsOutput), nil
	}
	out, err := c.WorkLinkAPI.ListFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListTagsForResource(input *worklink.ListTagsForResourceInput) (*worklink.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkLinkAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListTagsForResourceWithContext(ctx context.Context, input *worklink.ListTagsForResourceInput, opts ...request.Option) (*worklink.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkLinkAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListWebsiteAuthorizationProviders(input *worklink.ListWebsiteAuthorizationProvidersInput) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListWebsiteAuthorizationProvidersOutput), nil
	}
	out, err := c.WorkLinkAPI.ListWebsiteAuthorizationProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListWebsiteAuthorizationProvidersPages(input *worklink.ListWebsiteAuthorizationProvidersInput, fn func(*worklink.ListWebsiteAuthorizationProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListWebsiteAuthorizationProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListWebsiteAuthorizationProvidersOutput{}
	fnCacher := func(out *worklink.ListWebsiteAuthorizationProvidersOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListWebsiteAuthorizationProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListWebsiteAuthorizationProvidersPagesWithContext(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput, fn func(*worklink.ListWebsiteAuthorizationProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListWebsiteAuthorizationProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListWebsiteAuthorizationProvidersOutput{}
	fnCacher := func(out *worklink.ListWebsiteAuthorizationProvidersOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListWebsiteAuthorizationProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListWebsiteAuthorizationProvidersWithContext(ctx context.Context, input *worklink.ListWebsiteAuthorizationProvidersInput, opts ...request.Option) (*worklink.ListWebsiteAuthorizationProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListWebsiteAuthorizationProvidersOutput), nil
	}
	out, err := c.WorkLinkAPI.ListWebsiteAuthorizationProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListWebsiteCertificateAuthorities(input *worklink.ListWebsiteCertificateAuthoritiesInput) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListWebsiteCertificateAuthoritiesOutput), nil
	}
	out, err := c.WorkLinkAPI.ListWebsiteCertificateAuthorities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkLink) ListWebsiteCertificateAuthoritiesPages(input *worklink.ListWebsiteCertificateAuthoritiesInput, fn func(*worklink.ListWebsiteCertificateAuthoritiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListWebsiteCertificateAuthoritiesOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListWebsiteCertificateAuthoritiesOutput{}
	fnCacher := func(out *worklink.ListWebsiteCertificateAuthoritiesOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListWebsiteCertificateAuthoritiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListWebsiteCertificateAuthoritiesPagesWithContext(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput, fn func(*worklink.ListWebsiteCertificateAuthoritiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*worklink.ListWebsiteCertificateAuthoritiesOutput), false)
		return nil
	}
	cachable := true
	output := &worklink.ListWebsiteCertificateAuthoritiesOutput{}
	fnCacher := func(out *worklink.ListWebsiteCertificateAuthoritiesOutput, more bool) bool {
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
	if err := c.WorkLinkAPI.ListWebsiteCertificateAuthoritiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkLink) ListWebsiteCertificateAuthoritiesWithContext(ctx context.Context, input *worklink.ListWebsiteCertificateAuthoritiesInput, opts ...request.Option) (*worklink.ListWebsiteCertificateAuthoritiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*worklink.ListWebsiteCertificateAuthoritiesOutput), nil
	}
	out, err := c.WorkLinkAPI.ListWebsiteCertificateAuthoritiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
