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
package organizationscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/organizations"
	"github.com/aws/aws-sdk-go/service/organizations/organizationsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Organizations struct {
	organizationsiface.OrganizationsAPI
	cache *cache.Cache
}

func New(organizationsapi organizationsiface.OrganizationsAPI) *Organizations {
	return &Organizations{
		OrganizationsAPI: organizationsapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Organizations) DescribeAccount(input *organizations.DescribeAccountInput) (*organizations.DescribeAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeAccountWithContext(ctx context.Context, input *organizations.DescribeAccountInput, opts ...request.Option) (*organizations.DescribeAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeCreateAccountStatus(input *organizations.DescribeCreateAccountStatusInput) (*organizations.DescribeCreateAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeCreateAccountStatusOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeCreateAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeCreateAccountStatusWithContext(ctx context.Context, input *organizations.DescribeCreateAccountStatusInput, opts ...request.Option) (*organizations.DescribeCreateAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeCreateAccountStatusOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeCreateAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeEffectivePolicy(input *organizations.DescribeEffectivePolicyInput) (*organizations.DescribeEffectivePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeEffectivePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeEffectivePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeEffectivePolicyWithContext(ctx context.Context, input *organizations.DescribeEffectivePolicyInput, opts ...request.Option) (*organizations.DescribeEffectivePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeEffectivePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeEffectivePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeHandshake(input *organizations.DescribeHandshakeInput) (*organizations.DescribeHandshakeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeHandshakeOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeHandshake(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeHandshakeWithContext(ctx context.Context, input *organizations.DescribeHandshakeInput, opts ...request.Option) (*organizations.DescribeHandshakeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeHandshakeOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeHandshakeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeOrganization(input *organizations.DescribeOrganizationInput) (*organizations.DescribeOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeOrganizationWithContext(ctx context.Context, input *organizations.DescribeOrganizationInput, opts ...request.Option) (*organizations.DescribeOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeOrganizationalUnit(input *organizations.DescribeOrganizationalUnitInput) (*organizations.DescribeOrganizationalUnitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeOrganizationalUnitOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeOrganizationalUnit(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeOrganizationalUnitWithContext(ctx context.Context, input *organizations.DescribeOrganizationalUnitInput, opts ...request.Option) (*organizations.DescribeOrganizationalUnitOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeOrganizationalUnitOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeOrganizationalUnitWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribePolicy(input *organizations.DescribePolicyInput) (*organizations.DescribePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribePolicyWithContext(ctx context.Context, input *organizations.DescribePolicyInput, opts ...request.Option) (*organizations.DescribePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeResourcePolicy(input *organizations.DescribeResourcePolicyInput) (*organizations.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeResourcePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) DescribeResourcePolicyWithContext(ctx context.Context, input *organizations.DescribeResourcePolicyInput, opts ...request.Option) (*organizations.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.DescribeResourcePolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.DescribeResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAWSServiceAccessForOrganization(input *organizations.ListAWSServiceAccessForOrganizationInput) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAWSServiceAccessForOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAWSServiceAccessForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAWSServiceAccessForOrganizationPages(input *organizations.ListAWSServiceAccessForOrganizationInput, fn func(*organizations.ListAWSServiceAccessForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAWSServiceAccessForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAWSServiceAccessForOrganizationOutput{}
	fnCacher := func(out *organizations.ListAWSServiceAccessForOrganizationOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAWSServiceAccessForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAWSServiceAccessForOrganizationPagesWithContext(ctx context.Context, input *organizations.ListAWSServiceAccessForOrganizationInput, fn func(*organizations.ListAWSServiceAccessForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAWSServiceAccessForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAWSServiceAccessForOrganizationOutput{}
	fnCacher := func(out *organizations.ListAWSServiceAccessForOrganizationOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAWSServiceAccessForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAWSServiceAccessForOrganizationWithContext(ctx context.Context, input *organizations.ListAWSServiceAccessForOrganizationInput, opts ...request.Option) (*organizations.ListAWSServiceAccessForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAWSServiceAccessForOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAWSServiceAccessForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAccounts(input *organizations.ListAccountsInput) (*organizations.ListAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAccountsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAccountsForParent(input *organizations.ListAccountsForParentInput) (*organizations.ListAccountsForParentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAccountsForParentOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAccountsForParent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAccountsForParentPages(input *organizations.ListAccountsForParentInput, fn func(*organizations.ListAccountsForParentOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAccountsForParentOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAccountsForParentOutput{}
	fnCacher := func(out *organizations.ListAccountsForParentOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAccountsForParentPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAccountsForParentPagesWithContext(ctx context.Context, input *organizations.ListAccountsForParentInput, fn func(*organizations.ListAccountsForParentOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAccountsForParentOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAccountsForParentOutput{}
	fnCacher := func(out *organizations.ListAccountsForParentOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAccountsForParentPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAccountsForParentWithContext(ctx context.Context, input *organizations.ListAccountsForParentInput, opts ...request.Option) (*organizations.ListAccountsForParentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAccountsForParentOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAccountsForParentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListAccountsPages(input *organizations.ListAccountsInput, fn func(*organizations.ListAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAccountsOutput{}
	fnCacher := func(out *organizations.ListAccountsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAccountsPagesWithContext(ctx context.Context, input *organizations.ListAccountsInput, fn func(*organizations.ListAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListAccountsOutput{}
	fnCacher := func(out *organizations.ListAccountsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListAccountsWithContext(ctx context.Context, input *organizations.ListAccountsInput, opts ...request.Option) (*organizations.ListAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListAccountsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListChildren(input *organizations.ListChildrenInput) (*organizations.ListChildrenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListChildrenOutput), nil
	}
	out, err := c.OrganizationsAPI.ListChildren(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListChildrenPages(input *organizations.ListChildrenInput, fn func(*organizations.ListChildrenOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListChildrenOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListChildrenOutput{}
	fnCacher := func(out *organizations.ListChildrenOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListChildrenPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListChildrenPagesWithContext(ctx context.Context, input *organizations.ListChildrenInput, fn func(*organizations.ListChildrenOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListChildrenOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListChildrenOutput{}
	fnCacher := func(out *organizations.ListChildrenOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListChildrenPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListChildrenWithContext(ctx context.Context, input *organizations.ListChildrenInput, opts ...request.Option) (*organizations.ListChildrenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListChildrenOutput), nil
	}
	out, err := c.OrganizationsAPI.ListChildrenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListCreateAccountStatus(input *organizations.ListCreateAccountStatusInput) (*organizations.ListCreateAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListCreateAccountStatusOutput), nil
	}
	out, err := c.OrganizationsAPI.ListCreateAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListCreateAccountStatusPages(input *organizations.ListCreateAccountStatusInput, fn func(*organizations.ListCreateAccountStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListCreateAccountStatusOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListCreateAccountStatusOutput{}
	fnCacher := func(out *organizations.ListCreateAccountStatusOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListCreateAccountStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListCreateAccountStatusPagesWithContext(ctx context.Context, input *organizations.ListCreateAccountStatusInput, fn func(*organizations.ListCreateAccountStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListCreateAccountStatusOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListCreateAccountStatusOutput{}
	fnCacher := func(out *organizations.ListCreateAccountStatusOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListCreateAccountStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListCreateAccountStatusWithContext(ctx context.Context, input *organizations.ListCreateAccountStatusInput, opts ...request.Option) (*organizations.ListCreateAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListCreateAccountStatusOutput), nil
	}
	out, err := c.OrganizationsAPI.ListCreateAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListDelegatedAdministrators(input *organizations.ListDelegatedAdministratorsInput) (*organizations.ListDelegatedAdministratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListDelegatedAdministratorsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListDelegatedAdministrators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListDelegatedAdministratorsPages(input *organizations.ListDelegatedAdministratorsInput, fn func(*organizations.ListDelegatedAdministratorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListDelegatedAdministratorsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListDelegatedAdministratorsOutput{}
	fnCacher := func(out *organizations.ListDelegatedAdministratorsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListDelegatedAdministratorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListDelegatedAdministratorsPagesWithContext(ctx context.Context, input *organizations.ListDelegatedAdministratorsInput, fn func(*organizations.ListDelegatedAdministratorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListDelegatedAdministratorsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListDelegatedAdministratorsOutput{}
	fnCacher := func(out *organizations.ListDelegatedAdministratorsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListDelegatedAdministratorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListDelegatedAdministratorsWithContext(ctx context.Context, input *organizations.ListDelegatedAdministratorsInput, opts ...request.Option) (*organizations.ListDelegatedAdministratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListDelegatedAdministratorsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListDelegatedAdministratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListDelegatedServicesForAccount(input *organizations.ListDelegatedServicesForAccountInput) (*organizations.ListDelegatedServicesForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListDelegatedServicesForAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.ListDelegatedServicesForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListDelegatedServicesForAccountPages(input *organizations.ListDelegatedServicesForAccountInput, fn func(*organizations.ListDelegatedServicesForAccountOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListDelegatedServicesForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListDelegatedServicesForAccountOutput{}
	fnCacher := func(out *organizations.ListDelegatedServicesForAccountOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListDelegatedServicesForAccountPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListDelegatedServicesForAccountPagesWithContext(ctx context.Context, input *organizations.ListDelegatedServicesForAccountInput, fn func(*organizations.ListDelegatedServicesForAccountOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListDelegatedServicesForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListDelegatedServicesForAccountOutput{}
	fnCacher := func(out *organizations.ListDelegatedServicesForAccountOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListDelegatedServicesForAccountPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListDelegatedServicesForAccountWithContext(ctx context.Context, input *organizations.ListDelegatedServicesForAccountInput, opts ...request.Option) (*organizations.ListDelegatedServicesForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListDelegatedServicesForAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.ListDelegatedServicesForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListHandshakesForAccount(input *organizations.ListHandshakesForAccountInput) (*organizations.ListHandshakesForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListHandshakesForAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.ListHandshakesForAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListHandshakesForAccountPages(input *organizations.ListHandshakesForAccountInput, fn func(*organizations.ListHandshakesForAccountOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListHandshakesForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListHandshakesForAccountOutput{}
	fnCacher := func(out *organizations.ListHandshakesForAccountOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListHandshakesForAccountPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListHandshakesForAccountPagesWithContext(ctx context.Context, input *organizations.ListHandshakesForAccountInput, fn func(*organizations.ListHandshakesForAccountOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListHandshakesForAccountOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListHandshakesForAccountOutput{}
	fnCacher := func(out *organizations.ListHandshakesForAccountOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListHandshakesForAccountPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListHandshakesForAccountWithContext(ctx context.Context, input *organizations.ListHandshakesForAccountInput, opts ...request.Option) (*organizations.ListHandshakesForAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListHandshakesForAccountOutput), nil
	}
	out, err := c.OrganizationsAPI.ListHandshakesForAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListHandshakesForOrganization(input *organizations.ListHandshakesForOrganizationInput) (*organizations.ListHandshakesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListHandshakesForOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.ListHandshakesForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListHandshakesForOrganizationPages(input *organizations.ListHandshakesForOrganizationInput, fn func(*organizations.ListHandshakesForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListHandshakesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListHandshakesForOrganizationOutput{}
	fnCacher := func(out *organizations.ListHandshakesForOrganizationOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListHandshakesForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListHandshakesForOrganizationPagesWithContext(ctx context.Context, input *organizations.ListHandshakesForOrganizationInput, fn func(*organizations.ListHandshakesForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListHandshakesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListHandshakesForOrganizationOutput{}
	fnCacher := func(out *organizations.ListHandshakesForOrganizationOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListHandshakesForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListHandshakesForOrganizationWithContext(ctx context.Context, input *organizations.ListHandshakesForOrganizationInput, opts ...request.Option) (*organizations.ListHandshakesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListHandshakesForOrganizationOutput), nil
	}
	out, err := c.OrganizationsAPI.ListHandshakesForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListOrganizationalUnitsForParent(input *organizations.ListOrganizationalUnitsForParentInput) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListOrganizationalUnitsForParentOutput), nil
	}
	out, err := c.OrganizationsAPI.ListOrganizationalUnitsForParent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListOrganizationalUnitsForParentPages(input *organizations.ListOrganizationalUnitsForParentInput, fn func(*organizations.ListOrganizationalUnitsForParentOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListOrganizationalUnitsForParentOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListOrganizationalUnitsForParentOutput{}
	fnCacher := func(out *organizations.ListOrganizationalUnitsForParentOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListOrganizationalUnitsForParentPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListOrganizationalUnitsForParentPagesWithContext(ctx context.Context, input *organizations.ListOrganizationalUnitsForParentInput, fn func(*organizations.ListOrganizationalUnitsForParentOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListOrganizationalUnitsForParentOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListOrganizationalUnitsForParentOutput{}
	fnCacher := func(out *organizations.ListOrganizationalUnitsForParentOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListOrganizationalUnitsForParentPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListOrganizationalUnitsForParentWithContext(ctx context.Context, input *organizations.ListOrganizationalUnitsForParentInput, opts ...request.Option) (*organizations.ListOrganizationalUnitsForParentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListOrganizationalUnitsForParentOutput), nil
	}
	out, err := c.OrganizationsAPI.ListOrganizationalUnitsForParentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListParents(input *organizations.ListParentsInput) (*organizations.ListParentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListParentsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListParents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListParentsPages(input *organizations.ListParentsInput, fn func(*organizations.ListParentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListParentsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListParentsOutput{}
	fnCacher := func(out *organizations.ListParentsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListParentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListParentsPagesWithContext(ctx context.Context, input *organizations.ListParentsInput, fn func(*organizations.ListParentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListParentsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListParentsOutput{}
	fnCacher := func(out *organizations.ListParentsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListParentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListParentsWithContext(ctx context.Context, input *organizations.ListParentsInput, opts ...request.Option) (*organizations.ListParentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListParentsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListParentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListPolicies(input *organizations.ListPoliciesInput) (*organizations.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListPoliciesOutput), nil
	}
	out, err := c.OrganizationsAPI.ListPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListPoliciesForTarget(input *organizations.ListPoliciesForTargetInput) (*organizations.ListPoliciesForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListPoliciesForTargetOutput), nil
	}
	out, err := c.OrganizationsAPI.ListPoliciesForTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListPoliciesForTargetPages(input *organizations.ListPoliciesForTargetInput, fn func(*organizations.ListPoliciesForTargetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListPoliciesForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListPoliciesForTargetOutput{}
	fnCacher := func(out *organizations.ListPoliciesForTargetOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListPoliciesForTargetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListPoliciesForTargetPagesWithContext(ctx context.Context, input *organizations.ListPoliciesForTargetInput, fn func(*organizations.ListPoliciesForTargetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListPoliciesForTargetOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListPoliciesForTargetOutput{}
	fnCacher := func(out *organizations.ListPoliciesForTargetOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListPoliciesForTargetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListPoliciesForTargetWithContext(ctx context.Context, input *organizations.ListPoliciesForTargetInput, opts ...request.Option) (*organizations.ListPoliciesForTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListPoliciesForTargetOutput), nil
	}
	out, err := c.OrganizationsAPI.ListPoliciesForTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListPoliciesPages(input *organizations.ListPoliciesInput, fn func(*organizations.ListPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListPoliciesOutput{}
	fnCacher := func(out *organizations.ListPoliciesOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListPoliciesPagesWithContext(ctx context.Context, input *organizations.ListPoliciesInput, fn func(*organizations.ListPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListPoliciesOutput{}
	fnCacher := func(out *organizations.ListPoliciesOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListPoliciesWithContext(ctx context.Context, input *organizations.ListPoliciesInput, opts ...request.Option) (*organizations.ListPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListPoliciesOutput), nil
	}
	out, err := c.OrganizationsAPI.ListPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListRoots(input *organizations.ListRootsInput) (*organizations.ListRootsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListRootsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListRoots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListRootsPages(input *organizations.ListRootsInput, fn func(*organizations.ListRootsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListRootsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListRootsOutput{}
	fnCacher := func(out *organizations.ListRootsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListRootsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListRootsPagesWithContext(ctx context.Context, input *organizations.ListRootsInput, fn func(*organizations.ListRootsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListRootsOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListRootsOutput{}
	fnCacher := func(out *organizations.ListRootsOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListRootsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListRootsWithContext(ctx context.Context, input *organizations.ListRootsInput, opts ...request.Option) (*organizations.ListRootsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListRootsOutput), nil
	}
	out, err := c.OrganizationsAPI.ListRootsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListTagsForResource(input *organizations.ListTagsForResourceInput) (*organizations.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListTagsForResourceOutput), nil
	}
	out, err := c.OrganizationsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListTagsForResourcePages(input *organizations.ListTagsForResourceInput, fn func(*organizations.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListTagsForResourceOutput{}
	fnCacher := func(out *organizations.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListTagsForResourcePagesWithContext(ctx context.Context, input *organizations.ListTagsForResourceInput, fn func(*organizations.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListTagsForResourceOutput{}
	fnCacher := func(out *organizations.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListTagsForResourceWithContext(ctx context.Context, input *organizations.ListTagsForResourceInput, opts ...request.Option) (*organizations.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListTagsForResourceOutput), nil
	}
	out, err := c.OrganizationsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListTargetsForPolicy(input *organizations.ListTargetsForPolicyInput) (*organizations.ListTargetsForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListTargetsForPolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.ListTargetsForPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Organizations) ListTargetsForPolicyPages(input *organizations.ListTargetsForPolicyInput, fn func(*organizations.ListTargetsForPolicyOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListTargetsForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListTargetsForPolicyOutput{}
	fnCacher := func(out *organizations.ListTargetsForPolicyOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListTargetsForPolicyPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListTargetsForPolicyPagesWithContext(ctx context.Context, input *organizations.ListTargetsForPolicyInput, fn func(*organizations.ListTargetsForPolicyOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*organizations.ListTargetsForPolicyOutput), false)
		return nil
	}
	cachable := true
	output := &organizations.ListTargetsForPolicyOutput{}
	fnCacher := func(out *organizations.ListTargetsForPolicyOutput, more bool) bool {
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
	if err := c.OrganizationsAPI.ListTargetsForPolicyPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Organizations) ListTargetsForPolicyWithContext(ctx context.Context, input *organizations.ListTargetsForPolicyInput, opts ...request.Option) (*organizations.ListTargetsForPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*organizations.ListTargetsForPolicyOutput), nil
	}
	out, err := c.OrganizationsAPI.ListTargetsForPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
