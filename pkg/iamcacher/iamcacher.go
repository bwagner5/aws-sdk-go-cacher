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
package iamcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/iam/iamiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IAM struct {
	iamiface.IAMAPI
	cache *cache.Cache
}

func New(iamapi iamiface.IAMAPI) *IAM {
	return &IAM{
		IAMAPI: iamapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IAM) GetAccessKeyLastUsed(input *iam.GetAccessKeyLastUsedInput) (*iam.GetAccessKeyLastUsedOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccessKeyLastUsedOutput), nil
	}
	out, err := c.IAMAPI.GetAccessKeyLastUsed(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccessKeyLastUsedWithContext(ctx context.Context, input *iam.GetAccessKeyLastUsedInput, opts ...request.Option) (*iam.GetAccessKeyLastUsedOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccessKeyLastUsedOutput), nil
	}
	out, err := c.IAMAPI.GetAccessKeyLastUsedWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountAuthorizationDetails(input *iam.GetAccountAuthorizationDetailsInput) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountAuthorizationDetailsOutput), nil
	}
	out, err := c.IAMAPI.GetAccountAuthorizationDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountAuthorizationDetailsPages(input *iam.GetAccountAuthorizationDetailsInput, fn func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.GetAccountAuthorizationDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.GetAccountAuthorizationDetailsOutput{}
	fnCacher := func(out *iam.GetAccountAuthorizationDetailsOutput, more bool) bool {
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
	if err := c.IAMAPI.GetAccountAuthorizationDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) GetAccountAuthorizationDetailsPagesWithContext(ctx context.Context, input *iam.GetAccountAuthorizationDetailsInput, fn func(*iam.GetAccountAuthorizationDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.GetAccountAuthorizationDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.GetAccountAuthorizationDetailsOutput{}
	fnCacher := func(out *iam.GetAccountAuthorizationDetailsOutput, more bool) bool {
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
	if err := c.IAMAPI.GetAccountAuthorizationDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) GetAccountAuthorizationDetailsWithContext(ctx context.Context, input *iam.GetAccountAuthorizationDetailsInput, opts ...request.Option) (*iam.GetAccountAuthorizationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountAuthorizationDetailsOutput), nil
	}
	out, err := c.IAMAPI.GetAccountAuthorizationDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountPasswordPolicy(input *iam.GetAccountPasswordPolicyInput) (*iam.GetAccountPasswordPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountPasswordPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetAccountPasswordPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountPasswordPolicyWithContext(ctx context.Context, input *iam.GetAccountPasswordPolicyInput, opts ...request.Option) (*iam.GetAccountPasswordPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountPasswordPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetAccountPasswordPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountSummary(input *iam.GetAccountSummaryInput) (*iam.GetAccountSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountSummaryOutput), nil
	}
	out, err := c.IAMAPI.GetAccountSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetAccountSummaryWithContext(ctx context.Context, input *iam.GetAccountSummaryInput, opts ...request.Option) (*iam.GetAccountSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetAccountSummaryOutput), nil
	}
	out, err := c.IAMAPI.GetAccountSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetCredentialReport(input *iam.GetCredentialReportInput) (*iam.GetCredentialReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetCredentialReportOutput), nil
	}
	out, err := c.IAMAPI.GetCredentialReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetCredentialReportWithContext(ctx context.Context, input *iam.GetCredentialReportInput, opts ...request.Option) (*iam.GetCredentialReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetCredentialReportOutput), nil
	}
	out, err := c.IAMAPI.GetCredentialReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetGroup(input *iam.GetGroupInput) (*iam.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetGroupOutput), nil
	}
	out, err := c.IAMAPI.GetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetGroupPages(input *iam.GetGroupInput, fn func(*iam.GetGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.GetGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iam.GetGroupOutput{}
	fnCacher := func(out *iam.GetGroupOutput, more bool) bool {
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
	if err := c.IAMAPI.GetGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) GetGroupPagesWithContext(ctx context.Context, input *iam.GetGroupInput, fn func(*iam.GetGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.GetGroupOutput), false)
		return nil
	}
	cachable := true
	output := &iam.GetGroupOutput{}
	fnCacher := func(out *iam.GetGroupOutput, more bool) bool {
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
	if err := c.IAMAPI.GetGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) GetGroupPolicy(input *iam.GetGroupPolicyInput) (*iam.GetGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetGroupPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetGroupPolicyWithContext(ctx context.Context, input *iam.GetGroupPolicyInput, opts ...request.Option) (*iam.GetGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetGroupPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetGroupWithContext(ctx context.Context, input *iam.GetGroupInput, opts ...request.Option) (*iam.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetGroupOutput), nil
	}
	out, err := c.IAMAPI.GetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetInstanceProfile(input *iam.GetInstanceProfileInput) (*iam.GetInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetInstanceProfileOutput), nil
	}
	out, err := c.IAMAPI.GetInstanceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetInstanceProfileWithContext(ctx context.Context, input *iam.GetInstanceProfileInput, opts ...request.Option) (*iam.GetInstanceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetInstanceProfileOutput), nil
	}
	out, err := c.IAMAPI.GetInstanceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetLoginProfile(input *iam.GetLoginProfileInput) (*iam.GetLoginProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetLoginProfileOutput), nil
	}
	out, err := c.IAMAPI.GetLoginProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetLoginProfileWithContext(ctx context.Context, input *iam.GetLoginProfileInput, opts ...request.Option) (*iam.GetLoginProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetLoginProfileOutput), nil
	}
	out, err := c.IAMAPI.GetLoginProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetOpenIDConnectProvider(input *iam.GetOpenIDConnectProviderInput) (*iam.GetOpenIDConnectProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetOpenIDConnectProviderOutput), nil
	}
	out, err := c.IAMAPI.GetOpenIDConnectProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetOpenIDConnectProviderWithContext(ctx context.Context, input *iam.GetOpenIDConnectProviderInput, opts ...request.Option) (*iam.GetOpenIDConnectProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetOpenIDConnectProviderOutput), nil
	}
	out, err := c.IAMAPI.GetOpenIDConnectProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetOrganizationsAccessReport(input *iam.GetOrganizationsAccessReportInput) (*iam.GetOrganizationsAccessReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetOrganizationsAccessReportOutput), nil
	}
	out, err := c.IAMAPI.GetOrganizationsAccessReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetOrganizationsAccessReportWithContext(ctx context.Context, input *iam.GetOrganizationsAccessReportInput, opts ...request.Option) (*iam.GetOrganizationsAccessReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetOrganizationsAccessReportOutput), nil
	}
	out, err := c.IAMAPI.GetOrganizationsAccessReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetPolicy(input *iam.GetPolicyInput) (*iam.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetPolicyVersion(input *iam.GetPolicyVersionInput) (*iam.GetPolicyVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetPolicyVersionOutput), nil
	}
	out, err := c.IAMAPI.GetPolicyVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetPolicyVersionWithContext(ctx context.Context, input *iam.GetPolicyVersionInput, opts ...request.Option) (*iam.GetPolicyVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetPolicyVersionOutput), nil
	}
	out, err := c.IAMAPI.GetPolicyVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetPolicyWithContext(ctx context.Context, input *iam.GetPolicyInput, opts ...request.Option) (*iam.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetRole(input *iam.GetRoleInput) (*iam.GetRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetRoleOutput), nil
	}
	out, err := c.IAMAPI.GetRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetRolePolicy(input *iam.GetRolePolicyInput) (*iam.GetRolePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetRolePolicyOutput), nil
	}
	out, err := c.IAMAPI.GetRolePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetRolePolicyWithContext(ctx context.Context, input *iam.GetRolePolicyInput, opts ...request.Option) (*iam.GetRolePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetRolePolicyOutput), nil
	}
	out, err := c.IAMAPI.GetRolePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetRoleWithContext(ctx context.Context, input *iam.GetRoleInput, opts ...request.Option) (*iam.GetRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetRoleOutput), nil
	}
	out, err := c.IAMAPI.GetRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetSAMLProvider(input *iam.GetSAMLProviderInput) (*iam.GetSAMLProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetSAMLProviderOutput), nil
	}
	out, err := c.IAMAPI.GetSAMLProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetSAMLProviderWithContext(ctx context.Context, input *iam.GetSAMLProviderInput, opts ...request.Option) (*iam.GetSAMLProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetSAMLProviderOutput), nil
	}
	out, err := c.IAMAPI.GetSAMLProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetSSHPublicKey(input *iam.GetSSHPublicKeyInput) (*iam.GetSSHPublicKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetSSHPublicKeyOutput), nil
	}
	out, err := c.IAMAPI.GetSSHPublicKey(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetSSHPublicKeyWithContext(ctx context.Context, input *iam.GetSSHPublicKeyInput, opts ...request.Option) (*iam.GetSSHPublicKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetSSHPublicKeyOutput), nil
	}
	out, err := c.IAMAPI.GetSSHPublicKeyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServerCertificate(input *iam.GetServerCertificateInput) (*iam.GetServerCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServerCertificateOutput), nil
	}
	out, err := c.IAMAPI.GetServerCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServerCertificateWithContext(ctx context.Context, input *iam.GetServerCertificateInput, opts ...request.Option) (*iam.GetServerCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServerCertificateOutput), nil
	}
	out, err := c.IAMAPI.GetServerCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLastAccessedDetails(input *iam.GetServiceLastAccessedDetailsInput) (*iam.GetServiceLastAccessedDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLastAccessedDetailsOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLastAccessedDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLastAccessedDetailsWithContext(ctx context.Context, input *iam.GetServiceLastAccessedDetailsInput, opts ...request.Option) (*iam.GetServiceLastAccessedDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLastAccessedDetailsOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLastAccessedDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLastAccessedDetailsWithEntities(input *iam.GetServiceLastAccessedDetailsWithEntitiesInput) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLastAccessedDetailsWithEntitiesOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLastAccessedDetailsWithEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLastAccessedDetailsWithEntitiesWithContext(ctx context.Context, input *iam.GetServiceLastAccessedDetailsWithEntitiesInput, opts ...request.Option) (*iam.GetServiceLastAccessedDetailsWithEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLastAccessedDetailsWithEntitiesOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLastAccessedDetailsWithEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLinkedRoleDeletionStatus(input *iam.GetServiceLinkedRoleDeletionStatusInput) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLinkedRoleDeletionStatusOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLinkedRoleDeletionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetServiceLinkedRoleDeletionStatusWithContext(ctx context.Context, input *iam.GetServiceLinkedRoleDeletionStatusInput, opts ...request.Option) (*iam.GetServiceLinkedRoleDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetServiceLinkedRoleDeletionStatusOutput), nil
	}
	out, err := c.IAMAPI.GetServiceLinkedRoleDeletionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetUser(input *iam.GetUserInput) (*iam.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetUserOutput), nil
	}
	out, err := c.IAMAPI.GetUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetUserPolicy(input *iam.GetUserPolicyInput) (*iam.GetUserPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetUserPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetUserPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetUserPolicyWithContext(ctx context.Context, input *iam.GetUserPolicyInput, opts ...request.Option) (*iam.GetUserPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetUserPolicyOutput), nil
	}
	out, err := c.IAMAPI.GetUserPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) GetUserWithContext(ctx context.Context, input *iam.GetUserInput, opts ...request.Option) (*iam.GetUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.GetUserOutput), nil
	}
	out, err := c.IAMAPI.GetUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAccessKeys(input *iam.ListAccessKeysInput) (*iam.ListAccessKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAccessKeysOutput), nil
	}
	out, err := c.IAMAPI.ListAccessKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAccessKeysPages(input *iam.ListAccessKeysInput, fn func(*iam.ListAccessKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAccessKeysOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAccessKeysOutput{}
	fnCacher := func(out *iam.ListAccessKeysOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAccessKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAccessKeysPagesWithContext(ctx context.Context, input *iam.ListAccessKeysInput, fn func(*iam.ListAccessKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAccessKeysOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAccessKeysOutput{}
	fnCacher := func(out *iam.ListAccessKeysOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAccessKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAccessKeysWithContext(ctx context.Context, input *iam.ListAccessKeysInput, opts ...request.Option) (*iam.ListAccessKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAccessKeysOutput), nil
	}
	out, err := c.IAMAPI.ListAccessKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAccountAliases(input *iam.ListAccountAliasesInput) (*iam.ListAccountAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAccountAliasesOutput), nil
	}
	out, err := c.IAMAPI.ListAccountAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAccountAliasesPages(input *iam.ListAccountAliasesInput, fn func(*iam.ListAccountAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAccountAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAccountAliasesOutput{}
	fnCacher := func(out *iam.ListAccountAliasesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAccountAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAccountAliasesPagesWithContext(ctx context.Context, input *iam.ListAccountAliasesInput, fn func(*iam.ListAccountAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAccountAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAccountAliasesOutput{}
	fnCacher := func(out *iam.ListAccountAliasesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAccountAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAccountAliasesWithContext(ctx context.Context, input *iam.ListAccountAliasesInput, opts ...request.Option) (*iam.ListAccountAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAccountAliasesOutput), nil
	}
	out, err := c.IAMAPI.ListAccountAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedGroupPolicies(input *iam.ListAttachedGroupPoliciesInput) (*iam.ListAttachedGroupPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedGroupPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedGroupPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedGroupPoliciesPages(input *iam.ListAttachedGroupPoliciesInput, fn func(*iam.ListAttachedGroupPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedGroupPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedGroupPoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedGroupPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedGroupPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedGroupPoliciesPagesWithContext(ctx context.Context, input *iam.ListAttachedGroupPoliciesInput, fn func(*iam.ListAttachedGroupPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedGroupPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedGroupPoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedGroupPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedGroupPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedGroupPoliciesWithContext(ctx context.Context, input *iam.ListAttachedGroupPoliciesInput, opts ...request.Option) (*iam.ListAttachedGroupPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedGroupPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedGroupPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedRolePolicies(input *iam.ListAttachedRolePoliciesInput) (*iam.ListAttachedRolePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedRolePoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedRolePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedRolePoliciesPages(input *iam.ListAttachedRolePoliciesInput, fn func(*iam.ListAttachedRolePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedRolePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedRolePoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedRolePoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedRolePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedRolePoliciesPagesWithContext(ctx context.Context, input *iam.ListAttachedRolePoliciesInput, fn func(*iam.ListAttachedRolePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedRolePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedRolePoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedRolePoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedRolePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedRolePoliciesWithContext(ctx context.Context, input *iam.ListAttachedRolePoliciesInput, opts ...request.Option) (*iam.ListAttachedRolePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedRolePoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedRolePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedUserPolicies(input *iam.ListAttachedUserPoliciesInput) (*iam.ListAttachedUserPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedUserPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedUserPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListAttachedUserPoliciesPages(input *iam.ListAttachedUserPoliciesInput, fn func(*iam.ListAttachedUserPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedUserPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedUserPoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedUserPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedUserPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedUserPoliciesPagesWithContext(ctx context.Context, input *iam.ListAttachedUserPoliciesInput, fn func(*iam.ListAttachedUserPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListAttachedUserPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListAttachedUserPoliciesOutput{}
	fnCacher := func(out *iam.ListAttachedUserPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListAttachedUserPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListAttachedUserPoliciesWithContext(ctx context.Context, input *iam.ListAttachedUserPoliciesInput, opts ...request.Option) (*iam.ListAttachedUserPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListAttachedUserPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListAttachedUserPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListEntitiesForPolicy(input *iam.ListEntitiesForPolicyInput) (*iam.ListEntitiesForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListEntitiesForPolicyOutput), nil
	}
	out, err := c.IAMAPI.ListEntitiesForPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListEntitiesForPolicyPages(input *iam.ListEntitiesForPolicyInput, fn func(*iam.ListEntitiesForPolicyOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListEntitiesForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListEntitiesForPolicyOutput{}
	fnCacher := func(out *iam.ListEntitiesForPolicyOutput, more bool) bool {
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
	if err := c.IAMAPI.ListEntitiesForPolicyPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListEntitiesForPolicyPagesWithContext(ctx context.Context, input *iam.ListEntitiesForPolicyInput, fn func(*iam.ListEntitiesForPolicyOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListEntitiesForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListEntitiesForPolicyOutput{}
	fnCacher := func(out *iam.ListEntitiesForPolicyOutput, more bool) bool {
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
	if err := c.IAMAPI.ListEntitiesForPolicyPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListEntitiesForPolicyWithContext(ctx context.Context, input *iam.ListEntitiesForPolicyInput, opts ...request.Option) (*iam.ListEntitiesForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListEntitiesForPolicyOutput), nil
	}
	out, err := c.IAMAPI.ListEntitiesForPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroupPolicies(input *iam.ListGroupPoliciesInput) (*iam.ListGroupPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListGroupPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroupPoliciesPages(input *iam.ListGroupPoliciesInput, fn func(*iam.ListGroupPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupPoliciesOutput{}
	fnCacher := func(out *iam.ListGroupPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupPoliciesPagesWithContext(ctx context.Context, input *iam.ListGroupPoliciesInput, fn func(*iam.ListGroupPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupPoliciesOutput{}
	fnCacher := func(out *iam.ListGroupPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupPoliciesWithContext(ctx context.Context, input *iam.ListGroupPoliciesInput, opts ...request.Option) (*iam.ListGroupPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListGroupPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroups(input *iam.ListGroupsInput) (*iam.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupsOutput), nil
	}
	out, err := c.IAMAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroupsForUser(input *iam.ListGroupsForUserInput) (*iam.ListGroupsForUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupsForUserOutput), nil
	}
	out, err := c.IAMAPI.ListGroupsForUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroupsForUserPages(input *iam.ListGroupsForUserInput, fn func(*iam.ListGroupsForUserOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupsForUserOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupsForUserOutput{}
	fnCacher := func(out *iam.ListGroupsForUserOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupsForUserPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupsForUserPagesWithContext(ctx context.Context, input *iam.ListGroupsForUserInput, fn func(*iam.ListGroupsForUserOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupsForUserOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupsForUserOutput{}
	fnCacher := func(out *iam.ListGroupsForUserOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupsForUserPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupsForUserWithContext(ctx context.Context, input *iam.ListGroupsForUserInput, opts ...request.Option) (*iam.ListGroupsForUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupsForUserOutput), nil
	}
	out, err := c.IAMAPI.ListGroupsForUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListGroupsPages(input *iam.ListGroupsInput, fn func(*iam.ListGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupsOutput{}
	fnCacher := func(out *iam.ListGroupsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupsPagesWithContext(ctx context.Context, input *iam.ListGroupsInput, fn func(*iam.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListGroupsOutput{}
	fnCacher := func(out *iam.ListGroupsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListGroupsWithContext(ctx context.Context, input *iam.ListGroupsInput, opts ...request.Option) (*iam.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListGroupsOutput), nil
	}
	out, err := c.IAMAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfileTags(input *iam.ListInstanceProfileTagsInput) (*iam.ListInstanceProfileTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfileTagsOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfileTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfileTagsWithContext(ctx context.Context, input *iam.ListInstanceProfileTagsInput, opts ...request.Option) (*iam.ListInstanceProfileTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfileTagsOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfileTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfiles(input *iam.ListInstanceProfilesInput) (*iam.ListInstanceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfilesOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfilesForRole(input *iam.ListInstanceProfilesForRoleInput) (*iam.ListInstanceProfilesForRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfilesForRoleOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfilesForRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfilesForRolePages(input *iam.ListInstanceProfilesForRoleInput, fn func(*iam.ListInstanceProfilesForRoleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListInstanceProfilesForRoleOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListInstanceProfilesForRoleOutput{}
	fnCacher := func(out *iam.ListInstanceProfilesForRoleOutput, more bool) bool {
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
	if err := c.IAMAPI.ListInstanceProfilesForRolePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListInstanceProfilesForRolePagesWithContext(ctx context.Context, input *iam.ListInstanceProfilesForRoleInput, fn func(*iam.ListInstanceProfilesForRoleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListInstanceProfilesForRoleOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListInstanceProfilesForRoleOutput{}
	fnCacher := func(out *iam.ListInstanceProfilesForRoleOutput, more bool) bool {
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
	if err := c.IAMAPI.ListInstanceProfilesForRolePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListInstanceProfilesForRoleWithContext(ctx context.Context, input *iam.ListInstanceProfilesForRoleInput, opts ...request.Option) (*iam.ListInstanceProfilesForRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfilesForRoleOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfilesForRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListInstanceProfilesPages(input *iam.ListInstanceProfilesInput, fn func(*iam.ListInstanceProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListInstanceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListInstanceProfilesOutput{}
	fnCacher := func(out *iam.ListInstanceProfilesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListInstanceProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListInstanceProfilesPagesWithContext(ctx context.Context, input *iam.ListInstanceProfilesInput, fn func(*iam.ListInstanceProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListInstanceProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListInstanceProfilesOutput{}
	fnCacher := func(out *iam.ListInstanceProfilesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListInstanceProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListInstanceProfilesWithContext(ctx context.Context, input *iam.ListInstanceProfilesInput, opts ...request.Option) (*iam.ListInstanceProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListInstanceProfilesOutput), nil
	}
	out, err := c.IAMAPI.ListInstanceProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListMFADeviceTags(input *iam.ListMFADeviceTagsInput) (*iam.ListMFADeviceTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListMFADeviceTagsOutput), nil
	}
	out, err := c.IAMAPI.ListMFADeviceTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListMFADeviceTagsWithContext(ctx context.Context, input *iam.ListMFADeviceTagsInput, opts ...request.Option) (*iam.ListMFADeviceTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListMFADeviceTagsOutput), nil
	}
	out, err := c.IAMAPI.ListMFADeviceTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListMFADevices(input *iam.ListMFADevicesInput) (*iam.ListMFADevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListMFADevicesOutput), nil
	}
	out, err := c.IAMAPI.ListMFADevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListMFADevicesPages(input *iam.ListMFADevicesInput, fn func(*iam.ListMFADevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListMFADevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListMFADevicesOutput{}
	fnCacher := func(out *iam.ListMFADevicesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListMFADevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListMFADevicesPagesWithContext(ctx context.Context, input *iam.ListMFADevicesInput, fn func(*iam.ListMFADevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListMFADevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListMFADevicesOutput{}
	fnCacher := func(out *iam.ListMFADevicesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListMFADevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListMFADevicesWithContext(ctx context.Context, input *iam.ListMFADevicesInput, opts ...request.Option) (*iam.ListMFADevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListMFADevicesOutput), nil
	}
	out, err := c.IAMAPI.ListMFADevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListOpenIDConnectProviderTags(input *iam.ListOpenIDConnectProviderTagsInput) (*iam.ListOpenIDConnectProviderTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListOpenIDConnectProviderTagsOutput), nil
	}
	out, err := c.IAMAPI.ListOpenIDConnectProviderTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListOpenIDConnectProviderTagsWithContext(ctx context.Context, input *iam.ListOpenIDConnectProviderTagsInput, opts ...request.Option) (*iam.ListOpenIDConnectProviderTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListOpenIDConnectProviderTagsOutput), nil
	}
	out, err := c.IAMAPI.ListOpenIDConnectProviderTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListOpenIDConnectProviders(input *iam.ListOpenIDConnectProvidersInput) (*iam.ListOpenIDConnectProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListOpenIDConnectProvidersOutput), nil
	}
	out, err := c.IAMAPI.ListOpenIDConnectProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListOpenIDConnectProvidersWithContext(ctx context.Context, input *iam.ListOpenIDConnectProvidersInput, opts ...request.Option) (*iam.ListOpenIDConnectProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListOpenIDConnectProvidersOutput), nil
	}
	out, err := c.IAMAPI.ListOpenIDConnectProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPolicies(input *iam.ListPoliciesInput) (*iam.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPoliciesGrantingServiceAccess(input *iam.ListPoliciesGrantingServiceAccessInput) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPoliciesGrantingServiceAccessOutput), nil
	}
	out, err := c.IAMAPI.ListPoliciesGrantingServiceAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPoliciesGrantingServiceAccessWithContext(ctx context.Context, input *iam.ListPoliciesGrantingServiceAccessInput, opts ...request.Option) (*iam.ListPoliciesGrantingServiceAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPoliciesGrantingServiceAccessOutput), nil
	}
	out, err := c.IAMAPI.ListPoliciesGrantingServiceAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPoliciesPages(input *iam.ListPoliciesInput, fn func(*iam.ListPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListPoliciesOutput{}
	fnCacher := func(out *iam.ListPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListPoliciesPagesWithContext(ctx context.Context, input *iam.ListPoliciesInput, fn func(*iam.ListPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListPoliciesOutput{}
	fnCacher := func(out *iam.ListPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListPoliciesWithContext(ctx context.Context, input *iam.ListPoliciesInput, opts ...request.Option) (*iam.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPolicyTags(input *iam.ListPolicyTagsInput) (*iam.ListPolicyTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPolicyTagsOutput), nil
	}
	out, err := c.IAMAPI.ListPolicyTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPolicyTagsWithContext(ctx context.Context, input *iam.ListPolicyTagsInput, opts ...request.Option) (*iam.ListPolicyTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPolicyTagsOutput), nil
	}
	out, err := c.IAMAPI.ListPolicyTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPolicyVersions(input *iam.ListPolicyVersionsInput) (*iam.ListPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPolicyVersionsOutput), nil
	}
	out, err := c.IAMAPI.ListPolicyVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListPolicyVersionsPages(input *iam.ListPolicyVersionsInput, fn func(*iam.ListPolicyVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListPolicyVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListPolicyVersionsOutput{}
	fnCacher := func(out *iam.ListPolicyVersionsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListPolicyVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListPolicyVersionsPagesWithContext(ctx context.Context, input *iam.ListPolicyVersionsInput, fn func(*iam.ListPolicyVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListPolicyVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListPolicyVersionsOutput{}
	fnCacher := func(out *iam.ListPolicyVersionsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListPolicyVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListPolicyVersionsWithContext(ctx context.Context, input *iam.ListPolicyVersionsInput, opts ...request.Option) (*iam.ListPolicyVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListPolicyVersionsOutput), nil
	}
	out, err := c.IAMAPI.ListPolicyVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRolePolicies(input *iam.ListRolePoliciesInput) (*iam.ListRolePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRolePoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListRolePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRolePoliciesPages(input *iam.ListRolePoliciesInput, fn func(*iam.ListRolePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListRolePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListRolePoliciesOutput{}
	fnCacher := func(out *iam.ListRolePoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListRolePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListRolePoliciesPagesWithContext(ctx context.Context, input *iam.ListRolePoliciesInput, fn func(*iam.ListRolePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListRolePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListRolePoliciesOutput{}
	fnCacher := func(out *iam.ListRolePoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListRolePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListRolePoliciesWithContext(ctx context.Context, input *iam.ListRolePoliciesInput, opts ...request.Option) (*iam.ListRolePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRolePoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListRolePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRoleTags(input *iam.ListRoleTagsInput) (*iam.ListRoleTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRoleTagsOutput), nil
	}
	out, err := c.IAMAPI.ListRoleTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRoleTagsWithContext(ctx context.Context, input *iam.ListRoleTagsInput, opts ...request.Option) (*iam.ListRoleTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRoleTagsOutput), nil
	}
	out, err := c.IAMAPI.ListRoleTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRoles(input *iam.ListRolesInput) (*iam.ListRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRolesOutput), nil
	}
	out, err := c.IAMAPI.ListRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListRolesPages(input *iam.ListRolesInput, fn func(*iam.ListRolesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListRolesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListRolesOutput{}
	fnCacher := func(out *iam.ListRolesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListRolesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListRolesPagesWithContext(ctx context.Context, input *iam.ListRolesInput, fn func(*iam.ListRolesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListRolesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListRolesOutput{}
	fnCacher := func(out *iam.ListRolesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListRolesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListRolesWithContext(ctx context.Context, input *iam.ListRolesInput, opts ...request.Option) (*iam.ListRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListRolesOutput), nil
	}
	out, err := c.IAMAPI.ListRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSAMLProviderTags(input *iam.ListSAMLProviderTagsInput) (*iam.ListSAMLProviderTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSAMLProviderTagsOutput), nil
	}
	out, err := c.IAMAPI.ListSAMLProviderTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSAMLProviderTagsWithContext(ctx context.Context, input *iam.ListSAMLProviderTagsInput, opts ...request.Option) (*iam.ListSAMLProviderTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSAMLProviderTagsOutput), nil
	}
	out, err := c.IAMAPI.ListSAMLProviderTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSAMLProviders(input *iam.ListSAMLProvidersInput) (*iam.ListSAMLProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSAMLProvidersOutput), nil
	}
	out, err := c.IAMAPI.ListSAMLProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSAMLProvidersWithContext(ctx context.Context, input *iam.ListSAMLProvidersInput, opts ...request.Option) (*iam.ListSAMLProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSAMLProvidersOutput), nil
	}
	out, err := c.IAMAPI.ListSAMLProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSSHPublicKeys(input *iam.ListSSHPublicKeysInput) (*iam.ListSSHPublicKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSSHPublicKeysOutput), nil
	}
	out, err := c.IAMAPI.ListSSHPublicKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSSHPublicKeysPages(input *iam.ListSSHPublicKeysInput, fn func(*iam.ListSSHPublicKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListSSHPublicKeysOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListSSHPublicKeysOutput{}
	fnCacher := func(out *iam.ListSSHPublicKeysOutput, more bool) bool {
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
	if err := c.IAMAPI.ListSSHPublicKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListSSHPublicKeysPagesWithContext(ctx context.Context, input *iam.ListSSHPublicKeysInput, fn func(*iam.ListSSHPublicKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListSSHPublicKeysOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListSSHPublicKeysOutput{}
	fnCacher := func(out *iam.ListSSHPublicKeysOutput, more bool) bool {
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
	if err := c.IAMAPI.ListSSHPublicKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListSSHPublicKeysWithContext(ctx context.Context, input *iam.ListSSHPublicKeysInput, opts ...request.Option) (*iam.ListSSHPublicKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSSHPublicKeysOutput), nil
	}
	out, err := c.IAMAPI.ListSSHPublicKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServerCertificateTags(input *iam.ListServerCertificateTagsInput) (*iam.ListServerCertificateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServerCertificateTagsOutput), nil
	}
	out, err := c.IAMAPI.ListServerCertificateTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServerCertificateTagsWithContext(ctx context.Context, input *iam.ListServerCertificateTagsInput, opts ...request.Option) (*iam.ListServerCertificateTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServerCertificateTagsOutput), nil
	}
	out, err := c.IAMAPI.ListServerCertificateTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServerCertificates(input *iam.ListServerCertificatesInput) (*iam.ListServerCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServerCertificatesOutput), nil
	}
	out, err := c.IAMAPI.ListServerCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServerCertificatesPages(input *iam.ListServerCertificatesInput, fn func(*iam.ListServerCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListServerCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListServerCertificatesOutput{}
	fnCacher := func(out *iam.ListServerCertificatesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListServerCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListServerCertificatesPagesWithContext(ctx context.Context, input *iam.ListServerCertificatesInput, fn func(*iam.ListServerCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListServerCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListServerCertificatesOutput{}
	fnCacher := func(out *iam.ListServerCertificatesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListServerCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListServerCertificatesWithContext(ctx context.Context, input *iam.ListServerCertificatesInput, opts ...request.Option) (*iam.ListServerCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServerCertificatesOutput), nil
	}
	out, err := c.IAMAPI.ListServerCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServiceSpecificCredentials(input *iam.ListServiceSpecificCredentialsInput) (*iam.ListServiceSpecificCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServiceSpecificCredentialsOutput), nil
	}
	out, err := c.IAMAPI.ListServiceSpecificCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListServiceSpecificCredentialsWithContext(ctx context.Context, input *iam.ListServiceSpecificCredentialsInput, opts ...request.Option) (*iam.ListServiceSpecificCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListServiceSpecificCredentialsOutput), nil
	}
	out, err := c.IAMAPI.ListServiceSpecificCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSigningCertificates(input *iam.ListSigningCertificatesInput) (*iam.ListSigningCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSigningCertificatesOutput), nil
	}
	out, err := c.IAMAPI.ListSigningCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListSigningCertificatesPages(input *iam.ListSigningCertificatesInput, fn func(*iam.ListSigningCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListSigningCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListSigningCertificatesOutput{}
	fnCacher := func(out *iam.ListSigningCertificatesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListSigningCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListSigningCertificatesPagesWithContext(ctx context.Context, input *iam.ListSigningCertificatesInput, fn func(*iam.ListSigningCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListSigningCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListSigningCertificatesOutput{}
	fnCacher := func(out *iam.ListSigningCertificatesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListSigningCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListSigningCertificatesWithContext(ctx context.Context, input *iam.ListSigningCertificatesInput, opts ...request.Option) (*iam.ListSigningCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListSigningCertificatesOutput), nil
	}
	out, err := c.IAMAPI.ListSigningCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUserPolicies(input *iam.ListUserPoliciesInput) (*iam.ListUserPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUserPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListUserPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUserPoliciesPages(input *iam.ListUserPoliciesInput, fn func(*iam.ListUserPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUserPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUserPoliciesOutput{}
	fnCacher := func(out *iam.ListUserPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUserPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUserPoliciesPagesWithContext(ctx context.Context, input *iam.ListUserPoliciesInput, fn func(*iam.ListUserPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUserPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUserPoliciesOutput{}
	fnCacher := func(out *iam.ListUserPoliciesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUserPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUserPoliciesWithContext(ctx context.Context, input *iam.ListUserPoliciesInput, opts ...request.Option) (*iam.ListUserPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUserPoliciesOutput), nil
	}
	out, err := c.IAMAPI.ListUserPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUserTags(input *iam.ListUserTagsInput) (*iam.ListUserTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUserTagsOutput), nil
	}
	out, err := c.IAMAPI.ListUserTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUserTagsPages(input *iam.ListUserTagsInput, fn func(*iam.ListUserTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUserTagsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUserTagsOutput{}
	fnCacher := func(out *iam.ListUserTagsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUserTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUserTagsPagesWithContext(ctx context.Context, input *iam.ListUserTagsInput, fn func(*iam.ListUserTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUserTagsOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUserTagsOutput{}
	fnCacher := func(out *iam.ListUserTagsOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUserTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUserTagsWithContext(ctx context.Context, input *iam.ListUserTagsInput, opts ...request.Option) (*iam.ListUserTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUserTagsOutput), nil
	}
	out, err := c.IAMAPI.ListUserTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUsers(input *iam.ListUsersInput) (*iam.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUsersOutput), nil
	}
	out, err := c.IAMAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListUsersPages(input *iam.ListUsersInput, fn func(*iam.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUsersOutput{}
	fnCacher := func(out *iam.ListUsersOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUsersPagesWithContext(ctx context.Context, input *iam.ListUsersInput, fn func(*iam.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListUsersOutput{}
	fnCacher := func(out *iam.ListUsersOutput, more bool) bool {
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
	if err := c.IAMAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListUsersWithContext(ctx context.Context, input *iam.ListUsersInput, opts ...request.Option) (*iam.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListUsersOutput), nil
	}
	out, err := c.IAMAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListVirtualMFADevices(input *iam.ListVirtualMFADevicesInput) (*iam.ListVirtualMFADevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListVirtualMFADevicesOutput), nil
	}
	out, err := c.IAMAPI.ListVirtualMFADevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IAM) ListVirtualMFADevicesPages(input *iam.ListVirtualMFADevicesInput, fn func(*iam.ListVirtualMFADevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListVirtualMFADevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListVirtualMFADevicesOutput{}
	fnCacher := func(out *iam.ListVirtualMFADevicesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListVirtualMFADevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListVirtualMFADevicesPagesWithContext(ctx context.Context, input *iam.ListVirtualMFADevicesInput, fn func(*iam.ListVirtualMFADevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iam.ListVirtualMFADevicesOutput), false)
		return nil
	}
	cachable := true
	output := &iam.ListVirtualMFADevicesOutput{}
	fnCacher := func(out *iam.ListVirtualMFADevicesOutput, more bool) bool {
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
	if err := c.IAMAPI.ListVirtualMFADevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IAM) ListVirtualMFADevicesWithContext(ctx context.Context, input *iam.ListVirtualMFADevicesInput, opts ...request.Option) (*iam.ListVirtualMFADevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iam.ListVirtualMFADevicesOutput), nil
	}
	out, err := c.IAMAPI.ListVirtualMFADevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
