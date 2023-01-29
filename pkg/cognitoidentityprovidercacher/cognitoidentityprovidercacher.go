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
package cognitoidentityprovidercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider/cognitoidentityprovideriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CognitoIdentityProvider struct {
	cognitoidentityprovideriface.CognitoIdentityProviderAPI
	cache *cache.Cache
}

func New(cognitoidentityproviderapi cognitoidentityprovideriface.CognitoIdentityProviderAPI) *CognitoIdentityProvider {
	return &CognitoIdentityProvider{
		CognitoIdentityProviderAPI: cognitoidentityproviderapi,
		cache:                      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CognitoIdentityProvider) DescribeIdentityProvider(input *cognitoidentityprovider.DescribeIdentityProviderInput) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeIdentityProviderOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeIdentityProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeIdentityProviderWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeIdentityProviderInput, opts ...request.Option) (*cognitoidentityprovider.DescribeIdentityProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeIdentityProviderOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeIdentityProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeResourceServer(input *cognitoidentityprovider.DescribeResourceServerInput) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeResourceServerOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeResourceServer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeResourceServerWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeResourceServerInput, opts ...request.Option) (*cognitoidentityprovider.DescribeResourceServerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeResourceServerOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeResourceServerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeRiskConfiguration(input *cognitoidentityprovider.DescribeRiskConfigurationInput) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeRiskConfigurationOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeRiskConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeRiskConfigurationWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeRiskConfigurationInput, opts ...request.Option) (*cognitoidentityprovider.DescribeRiskConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeRiskConfigurationOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeRiskConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserImportJob(input *cognitoidentityprovider.DescribeUserImportJobInput) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserImportJobOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserImportJobWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeUserImportJobInput, opts ...request.Option) (*cognitoidentityprovider.DescribeUserImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserImportJobOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPool(input *cognitoidentityprovider.DescribeUserPoolInput) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPoolClient(input *cognitoidentityprovider.DescribeUserPoolClientInput) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolClientOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPoolClient(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPoolClientWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolClientInput, opts ...request.Option) (*cognitoidentityprovider.DescribeUserPoolClientOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolClientOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPoolClientWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPoolDomain(input *cognitoidentityprovider.DescribeUserPoolDomainInput) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolDomainOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPoolDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPoolDomainWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolDomainInput, opts ...request.Option) (*cognitoidentityprovider.DescribeUserPoolDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolDomainOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPoolDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) DescribeUserPoolWithContext(ctx context.Context, input *cognitoidentityprovider.DescribeUserPoolInput, opts ...request.Option) (*cognitoidentityprovider.DescribeUserPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.DescribeUserPoolOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.DescribeUserPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetCSVHeader(input *cognitoidentityprovider.GetCSVHeaderInput) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetCSVHeaderOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetCSVHeader(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetCSVHeaderWithContext(ctx context.Context, input *cognitoidentityprovider.GetCSVHeaderInput, opts ...request.Option) (*cognitoidentityprovider.GetCSVHeaderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetCSVHeaderOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetCSVHeaderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetDevice(input *cognitoidentityprovider.GetDeviceInput) (*cognitoidentityprovider.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetDeviceOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetDeviceWithContext(ctx context.Context, input *cognitoidentityprovider.GetDeviceInput, opts ...request.Option) (*cognitoidentityprovider.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetDeviceOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetGroup(input *cognitoidentityprovider.GetGroupInput) (*cognitoidentityprovider.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetGroupOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetGroupWithContext(ctx context.Context, input *cognitoidentityprovider.GetGroupInput, opts ...request.Option) (*cognitoidentityprovider.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetGroupOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetIdentityProviderByIdentifier(input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetIdentityProviderByIdentifier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetIdentityProviderByIdentifierWithContext(ctx context.Context, input *cognitoidentityprovider.GetIdentityProviderByIdentifierInput, opts ...request.Option) (*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetIdentityProviderByIdentifierOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetIdentityProviderByIdentifierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetSigningCertificate(input *cognitoidentityprovider.GetSigningCertificateInput) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetSigningCertificateOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetSigningCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetSigningCertificateWithContext(ctx context.Context, input *cognitoidentityprovider.GetSigningCertificateInput, opts ...request.Option) (*cognitoidentityprovider.GetSigningCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetSigningCertificateOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetSigningCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUICustomization(input *cognitoidentityprovider.GetUICustomizationInput) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUICustomizationOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUICustomization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUICustomizationWithContext(ctx context.Context, input *cognitoidentityprovider.GetUICustomizationInput, opts ...request.Option) (*cognitoidentityprovider.GetUICustomizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUICustomizationOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUICustomizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUser(input *cognitoidentityprovider.GetUserInput) (*cognitoidentityprovider.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUserAttributeVerificationCode(input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUserAttributeVerificationCode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUserAttributeVerificationCodeWithContext(ctx context.Context, input *cognitoidentityprovider.GetUserAttributeVerificationCodeInput, opts ...request.Option) (*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserAttributeVerificationCodeOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUserAttributeVerificationCodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUserPoolMfaConfig(input *cognitoidentityprovider.GetUserPoolMfaConfigInput) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserPoolMfaConfigOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUserPoolMfaConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUserPoolMfaConfigWithContext(ctx context.Context, input *cognitoidentityprovider.GetUserPoolMfaConfigInput, opts ...request.Option) (*cognitoidentityprovider.GetUserPoolMfaConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserPoolMfaConfigOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUserPoolMfaConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) GetUserWithContext(ctx context.Context, input *cognitoidentityprovider.GetUserInput, opts ...request.Option) (*cognitoidentityprovider.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.GetUserOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.GetUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListDevices(input *cognitoidentityprovider.ListDevicesInput) (*cognitoidentityprovider.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListDevicesOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListDevicesWithContext(ctx context.Context, input *cognitoidentityprovider.ListDevicesInput, opts ...request.Option) (*cognitoidentityprovider.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListDevicesOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListGroups(input *cognitoidentityprovider.ListGroupsInput) (*cognitoidentityprovider.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListGroupsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListGroupsPages(input *cognitoidentityprovider.ListGroupsInput, fn func(*cognitoidentityprovider.ListGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListGroupsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListGroupsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListGroupsPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListGroupsInput, fn func(*cognitoidentityprovider.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListGroupsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListGroupsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListGroupsWithContext(ctx context.Context, input *cognitoidentityprovider.ListGroupsInput, opts ...request.Option) (*cognitoidentityprovider.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListGroupsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListIdentityProviders(input *cognitoidentityprovider.ListIdentityProvidersInput) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListIdentityProvidersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListIdentityProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListIdentityProvidersPages(input *cognitoidentityprovider.ListIdentityProvidersInput, fn func(*cognitoidentityprovider.ListIdentityProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListIdentityProvidersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListIdentityProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListIdentityProvidersPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListIdentityProvidersInput, fn func(*cognitoidentityprovider.ListIdentityProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListIdentityProvidersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListIdentityProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListIdentityProvidersWithContext(ctx context.Context, input *cognitoidentityprovider.ListIdentityProvidersInput, opts ...request.Option) (*cognitoidentityprovider.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListIdentityProvidersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListIdentityProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListResourceServers(input *cognitoidentityprovider.ListResourceServersInput) (*cognitoidentityprovider.ListResourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListResourceServersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListResourceServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListResourceServersPages(input *cognitoidentityprovider.ListResourceServersInput, fn func(*cognitoidentityprovider.ListResourceServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListResourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListResourceServersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListResourceServersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListResourceServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListResourceServersPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListResourceServersInput, fn func(*cognitoidentityprovider.ListResourceServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListResourceServersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListResourceServersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListResourceServersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListResourceServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListResourceServersWithContext(ctx context.Context, input *cognitoidentityprovider.ListResourceServersInput, opts ...request.Option) (*cognitoidentityprovider.ListResourceServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListResourceServersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListResourceServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListTagsForResource(input *cognitoidentityprovider.ListTagsForResourceInput) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListTagsForResourceOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListTagsForResourceWithContext(ctx context.Context, input *cognitoidentityprovider.ListTagsForResourceInput, opts ...request.Option) (*cognitoidentityprovider.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListTagsForResourceOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserImportJobs(input *cognitoidentityprovider.ListUserImportJobsInput) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserImportJobsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserImportJobsWithContext(ctx context.Context, input *cognitoidentityprovider.ListUserImportJobsInput, opts ...request.Option) (*cognitoidentityprovider.ListUserImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserImportJobsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserPoolClients(input *cognitoidentityprovider.ListUserPoolClientsInput) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserPoolClientsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserPoolClients(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserPoolClientsPages(input *cognitoidentityprovider.ListUserPoolClientsInput, fn func(*cognitoidentityprovider.ListUserPoolClientsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUserPoolClientsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUserPoolClientsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUserPoolClientsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUserPoolClientsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUserPoolClientsPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListUserPoolClientsInput, fn func(*cognitoidentityprovider.ListUserPoolClientsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUserPoolClientsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUserPoolClientsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUserPoolClientsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUserPoolClientsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUserPoolClientsWithContext(ctx context.Context, input *cognitoidentityprovider.ListUserPoolClientsInput, opts ...request.Option) (*cognitoidentityprovider.ListUserPoolClientsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserPoolClientsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserPoolClientsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserPools(input *cognitoidentityprovider.ListUserPoolsInput) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserPoolsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUserPoolsPages(input *cognitoidentityprovider.ListUserPoolsInput, fn func(*cognitoidentityprovider.ListUserPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUserPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUserPoolsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUserPoolsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUserPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUserPoolsPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListUserPoolsInput, fn func(*cognitoidentityprovider.ListUserPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUserPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUserPoolsOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUserPoolsOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUserPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUserPoolsWithContext(ctx context.Context, input *cognitoidentityprovider.ListUserPoolsInput, opts ...request.Option) (*cognitoidentityprovider.ListUserPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUserPoolsOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUserPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUsers(input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUsersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUsersInGroup(input *cognitoidentityprovider.ListUsersInGroupInput) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUsersInGroupOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUsersInGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUsersInGroupPages(input *cognitoidentityprovider.ListUsersInGroupInput, fn func(*cognitoidentityprovider.ListUsersInGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUsersInGroupOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUsersInGroupOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUsersInGroupOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUsersInGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUsersInGroupPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListUsersInGroupInput, fn func(*cognitoidentityprovider.ListUsersInGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUsersInGroupOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUsersInGroupOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUsersInGroupOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUsersInGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUsersInGroupWithContext(ctx context.Context, input *cognitoidentityprovider.ListUsersInGroupInput, opts ...request.Option) (*cognitoidentityprovider.ListUsersInGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUsersInGroupOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUsersInGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoIdentityProvider) ListUsersPages(input *cognitoidentityprovider.ListUsersInput, fn func(*cognitoidentityprovider.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUsersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUsersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUsersPagesWithContext(ctx context.Context, input *cognitoidentityprovider.ListUsersInput, fn func(*cognitoidentityprovider.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cognitoidentityprovider.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &cognitoidentityprovider.ListUsersOutput{}
	fnCacher := func(out *cognitoidentityprovider.ListUsersOutput, more bool) bool {
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
	if err := c.CognitoIdentityProviderAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CognitoIdentityProvider) ListUsersWithContext(ctx context.Context, input *cognitoidentityprovider.ListUsersInput, opts ...request.Option) (*cognitoidentityprovider.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitoidentityprovider.ListUsersOutput), nil
	}
	out, err := c.CognitoIdentityProviderAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
