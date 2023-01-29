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
package workspaceswebcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/workspacesweb"
	"github.com/aws/aws-sdk-go/service/workspacesweb/workspaceswebiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkSpacesWeb struct {
	workspaceswebiface.WorkSpacesWebAPI
	cache *cache.Cache
}

func New(workspaceswebapi workspaceswebiface.WorkSpacesWebAPI) *WorkSpacesWeb {
	return &WorkSpacesWeb{
		WorkSpacesWebAPI: workspaceswebapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkSpacesWeb) GetBrowserSettings(input *workspacesweb.GetBrowserSettingsInput) (*workspacesweb.GetBrowserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetBrowserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetBrowserSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetBrowserSettingsWithContext(ctx context.Context, input *workspacesweb.GetBrowserSettingsInput, opts ...request.Option) (*workspacesweb.GetBrowserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetBrowserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetBrowserSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetIdentityProvider(input *workspacesweb.GetIdentityProviderInput) (*workspacesweb.GetIdentityProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetIdentityProviderOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetIdentityProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetIdentityProviderWithContext(ctx context.Context, input *workspacesweb.GetIdentityProviderInput, opts ...request.Option) (*workspacesweb.GetIdentityProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetIdentityProviderOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetIdentityProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetNetworkSettings(input *workspacesweb.GetNetworkSettingsInput) (*workspacesweb.GetNetworkSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetNetworkSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetNetworkSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetNetworkSettingsWithContext(ctx context.Context, input *workspacesweb.GetNetworkSettingsInput, opts ...request.Option) (*workspacesweb.GetNetworkSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetNetworkSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetNetworkSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetPortal(input *workspacesweb.GetPortalInput) (*workspacesweb.GetPortalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetPortalOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetPortal(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetPortalServiceProviderMetadata(input *workspacesweb.GetPortalServiceProviderMetadataInput) (*workspacesweb.GetPortalServiceProviderMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetPortalServiceProviderMetadataOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetPortalServiceProviderMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetPortalServiceProviderMetadataWithContext(ctx context.Context, input *workspacesweb.GetPortalServiceProviderMetadataInput, opts ...request.Option) (*workspacesweb.GetPortalServiceProviderMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetPortalServiceProviderMetadataOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetPortalServiceProviderMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetPortalWithContext(ctx context.Context, input *workspacesweb.GetPortalInput, opts ...request.Option) (*workspacesweb.GetPortalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetPortalOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetPortalWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetTrustStore(input *workspacesweb.GetTrustStoreInput) (*workspacesweb.GetTrustStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetTrustStoreOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetTrustStore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetTrustStoreCertificate(input *workspacesweb.GetTrustStoreCertificateInput) (*workspacesweb.GetTrustStoreCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetTrustStoreCertificateOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetTrustStoreCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetTrustStoreCertificateWithContext(ctx context.Context, input *workspacesweb.GetTrustStoreCertificateInput, opts ...request.Option) (*workspacesweb.GetTrustStoreCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetTrustStoreCertificateOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetTrustStoreCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetTrustStoreWithContext(ctx context.Context, input *workspacesweb.GetTrustStoreInput, opts ...request.Option) (*workspacesweb.GetTrustStoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetTrustStoreOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetTrustStoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetUserAccessLoggingSettings(input *workspacesweb.GetUserAccessLoggingSettingsInput) (*workspacesweb.GetUserAccessLoggingSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetUserAccessLoggingSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetUserAccessLoggingSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetUserAccessLoggingSettingsWithContext(ctx context.Context, input *workspacesweb.GetUserAccessLoggingSettingsInput, opts ...request.Option) (*workspacesweb.GetUserAccessLoggingSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetUserAccessLoggingSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetUserAccessLoggingSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetUserSettings(input *workspacesweb.GetUserSettingsInput) (*workspacesweb.GetUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetUserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetUserSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) GetUserSettingsWithContext(ctx context.Context, input *workspacesweb.GetUserSettingsInput, opts ...request.Option) (*workspacesweb.GetUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.GetUserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.GetUserSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListBrowserSettings(input *workspacesweb.ListBrowserSettingsInput) (*workspacesweb.ListBrowserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListBrowserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListBrowserSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListBrowserSettingsPages(input *workspacesweb.ListBrowserSettingsInput, fn func(*workspacesweb.ListBrowserSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListBrowserSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListBrowserSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListBrowserSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListBrowserSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListBrowserSettingsPagesWithContext(ctx context.Context, input *workspacesweb.ListBrowserSettingsInput, fn func(*workspacesweb.ListBrowserSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListBrowserSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListBrowserSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListBrowserSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListBrowserSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListBrowserSettingsWithContext(ctx context.Context, input *workspacesweb.ListBrowserSettingsInput, opts ...request.Option) (*workspacesweb.ListBrowserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListBrowserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListBrowserSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListIdentityProviders(input *workspacesweb.ListIdentityProvidersInput) (*workspacesweb.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListIdentityProvidersOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListIdentityProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListIdentityProvidersPages(input *workspacesweb.ListIdentityProvidersInput, fn func(*workspacesweb.ListIdentityProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListIdentityProvidersOutput{}
	fnCacher := func(out *workspacesweb.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListIdentityProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListIdentityProvidersPagesWithContext(ctx context.Context, input *workspacesweb.ListIdentityProvidersInput, fn func(*workspacesweb.ListIdentityProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListIdentityProvidersOutput{}
	fnCacher := func(out *workspacesweb.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListIdentityProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListIdentityProvidersWithContext(ctx context.Context, input *workspacesweb.ListIdentityProvidersInput, opts ...request.Option) (*workspacesweb.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListIdentityProvidersOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListIdentityProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListNetworkSettings(input *workspacesweb.ListNetworkSettingsInput) (*workspacesweb.ListNetworkSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListNetworkSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListNetworkSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListNetworkSettingsPages(input *workspacesweb.ListNetworkSettingsInput, fn func(*workspacesweb.ListNetworkSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListNetworkSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListNetworkSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListNetworkSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListNetworkSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListNetworkSettingsPagesWithContext(ctx context.Context, input *workspacesweb.ListNetworkSettingsInput, fn func(*workspacesweb.ListNetworkSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListNetworkSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListNetworkSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListNetworkSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListNetworkSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListNetworkSettingsWithContext(ctx context.Context, input *workspacesweb.ListNetworkSettingsInput, opts ...request.Option) (*workspacesweb.ListNetworkSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListNetworkSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListNetworkSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListPortals(input *workspacesweb.ListPortalsInput) (*workspacesweb.ListPortalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListPortalsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListPortals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListPortalsPages(input *workspacesweb.ListPortalsInput, fn func(*workspacesweb.ListPortalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListPortalsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListPortalsOutput{}
	fnCacher := func(out *workspacesweb.ListPortalsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListPortalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListPortalsPagesWithContext(ctx context.Context, input *workspacesweb.ListPortalsInput, fn func(*workspacesweb.ListPortalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListPortalsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListPortalsOutput{}
	fnCacher := func(out *workspacesweb.ListPortalsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListPortalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListPortalsWithContext(ctx context.Context, input *workspacesweb.ListPortalsInput, opts ...request.Option) (*workspacesweb.ListPortalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListPortalsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListPortalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTagsForResource(input *workspacesweb.ListTagsForResourceInput) (*workspacesweb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTagsForResourceWithContext(ctx context.Context, input *workspacesweb.ListTagsForResourceInput, opts ...request.Option) (*workspacesweb.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTagsForResourceOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTrustStoreCertificates(input *workspacesweb.ListTrustStoreCertificatesInput) (*workspacesweb.ListTrustStoreCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTrustStoreCertificatesOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTrustStoreCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTrustStoreCertificatesPages(input *workspacesweb.ListTrustStoreCertificatesInput, fn func(*workspacesweb.ListTrustStoreCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListTrustStoreCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListTrustStoreCertificatesOutput{}
	fnCacher := func(out *workspacesweb.ListTrustStoreCertificatesOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListTrustStoreCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListTrustStoreCertificatesPagesWithContext(ctx context.Context, input *workspacesweb.ListTrustStoreCertificatesInput, fn func(*workspacesweb.ListTrustStoreCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListTrustStoreCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListTrustStoreCertificatesOutput{}
	fnCacher := func(out *workspacesweb.ListTrustStoreCertificatesOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListTrustStoreCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListTrustStoreCertificatesWithContext(ctx context.Context, input *workspacesweb.ListTrustStoreCertificatesInput, opts ...request.Option) (*workspacesweb.ListTrustStoreCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTrustStoreCertificatesOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTrustStoreCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTrustStores(input *workspacesweb.ListTrustStoresInput) (*workspacesweb.ListTrustStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTrustStoresOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTrustStores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListTrustStoresPages(input *workspacesweb.ListTrustStoresInput, fn func(*workspacesweb.ListTrustStoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListTrustStoresOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListTrustStoresOutput{}
	fnCacher := func(out *workspacesweb.ListTrustStoresOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListTrustStoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListTrustStoresPagesWithContext(ctx context.Context, input *workspacesweb.ListTrustStoresInput, fn func(*workspacesweb.ListTrustStoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListTrustStoresOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListTrustStoresOutput{}
	fnCacher := func(out *workspacesweb.ListTrustStoresOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListTrustStoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListTrustStoresWithContext(ctx context.Context, input *workspacesweb.ListTrustStoresInput, opts ...request.Option) (*workspacesweb.ListTrustStoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListTrustStoresOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListTrustStoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListUserAccessLoggingSettings(input *workspacesweb.ListUserAccessLoggingSettingsInput) (*workspacesweb.ListUserAccessLoggingSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListUserAccessLoggingSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListUserAccessLoggingSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListUserAccessLoggingSettingsPages(input *workspacesweb.ListUserAccessLoggingSettingsInput, fn func(*workspacesweb.ListUserAccessLoggingSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListUserAccessLoggingSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListUserAccessLoggingSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListUserAccessLoggingSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListUserAccessLoggingSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListUserAccessLoggingSettingsPagesWithContext(ctx context.Context, input *workspacesweb.ListUserAccessLoggingSettingsInput, fn func(*workspacesweb.ListUserAccessLoggingSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListUserAccessLoggingSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListUserAccessLoggingSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListUserAccessLoggingSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListUserAccessLoggingSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListUserAccessLoggingSettingsWithContext(ctx context.Context, input *workspacesweb.ListUserAccessLoggingSettingsInput, opts ...request.Option) (*workspacesweb.ListUserAccessLoggingSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListUserAccessLoggingSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListUserAccessLoggingSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListUserSettings(input *workspacesweb.ListUserSettingsInput) (*workspacesweb.ListUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListUserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListUserSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpacesWeb) ListUserSettingsPages(input *workspacesweb.ListUserSettingsInput, fn func(*workspacesweb.ListUserSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListUserSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListUserSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListUserSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListUserSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListUserSettingsPagesWithContext(ctx context.Context, input *workspacesweb.ListUserSettingsInput, fn func(*workspacesweb.ListUserSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspacesweb.ListUserSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &workspacesweb.ListUserSettingsOutput{}
	fnCacher := func(out *workspacesweb.ListUserSettingsOutput, more bool) bool {
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
	if err := c.WorkSpacesWebAPI.ListUserSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpacesWeb) ListUserSettingsWithContext(ctx context.Context, input *workspacesweb.ListUserSettingsInput, opts ...request.Option) (*workspacesweb.ListUserSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspacesweb.ListUserSettingsOutput), nil
	}
	out, err := c.WorkSpacesWebAPI.ListUserSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
