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
package inspector2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/inspector2"
	"github.com/aws/aws-sdk-go/service/inspector2/inspector2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Inspector2 struct {
	inspector2iface.Inspector2API
	cache *cache.Cache
}

func New(inspector2api inspector2iface.Inspector2API) *Inspector2 {
	return &Inspector2{
		Inspector2API: inspector2api,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Inspector2) BatchGetAccountStatus(input *inspector2.BatchGetAccountStatusInput) (*inspector2.BatchGetAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.BatchGetAccountStatusOutput), nil
	}
	out, err := c.Inspector2API.BatchGetAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) BatchGetAccountStatusWithContext(ctx context.Context, input *inspector2.BatchGetAccountStatusInput, opts ...request.Option) (*inspector2.BatchGetAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.BatchGetAccountStatusOutput), nil
	}
	out, err := c.Inspector2API.BatchGetAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) BatchGetFreeTrialInfo(input *inspector2.BatchGetFreeTrialInfoInput) (*inspector2.BatchGetFreeTrialInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.BatchGetFreeTrialInfoOutput), nil
	}
	out, err := c.Inspector2API.BatchGetFreeTrialInfo(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) BatchGetFreeTrialInfoWithContext(ctx context.Context, input *inspector2.BatchGetFreeTrialInfoInput, opts ...request.Option) (*inspector2.BatchGetFreeTrialInfoOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.BatchGetFreeTrialInfoOutput), nil
	}
	out, err := c.Inspector2API.BatchGetFreeTrialInfoWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) DescribeOrganizationConfiguration(input *inspector2.DescribeOrganizationConfigurationInput) (*inspector2.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.Inspector2API.DescribeOrganizationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) DescribeOrganizationConfigurationWithContext(ctx context.Context, input *inspector2.DescribeOrganizationConfigurationInput, opts ...request.Option) (*inspector2.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.Inspector2API.DescribeOrganizationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetConfiguration(input *inspector2.GetConfigurationInput) (*inspector2.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetConfigurationOutput), nil
	}
	out, err := c.Inspector2API.GetConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetConfigurationWithContext(ctx context.Context, input *inspector2.GetConfigurationInput, opts ...request.Option) (*inspector2.GetConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetConfigurationOutput), nil
	}
	out, err := c.Inspector2API.GetConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetDelegatedAdminAccount(input *inspector2.GetDelegatedAdminAccountInput) (*inspector2.GetDelegatedAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetDelegatedAdminAccountOutput), nil
	}
	out, err := c.Inspector2API.GetDelegatedAdminAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetDelegatedAdminAccountWithContext(ctx context.Context, input *inspector2.GetDelegatedAdminAccountInput, opts ...request.Option) (*inspector2.GetDelegatedAdminAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetDelegatedAdminAccountOutput), nil
	}
	out, err := c.Inspector2API.GetDelegatedAdminAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetFindingsReportStatus(input *inspector2.GetFindingsReportStatusInput) (*inspector2.GetFindingsReportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetFindingsReportStatusOutput), nil
	}
	out, err := c.Inspector2API.GetFindingsReportStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetFindingsReportStatusWithContext(ctx context.Context, input *inspector2.GetFindingsReportStatusInput, opts ...request.Option) (*inspector2.GetFindingsReportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetFindingsReportStatusOutput), nil
	}
	out, err := c.Inspector2API.GetFindingsReportStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetMember(input *inspector2.GetMemberInput) (*inspector2.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetMemberOutput), nil
	}
	out, err := c.Inspector2API.GetMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) GetMemberWithContext(ctx context.Context, input *inspector2.GetMemberInput, opts ...request.Option) (*inspector2.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.GetMemberOutput), nil
	}
	out, err := c.Inspector2API.GetMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListAccountPermissions(input *inspector2.ListAccountPermissionsInput) (*inspector2.ListAccountPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListAccountPermissionsOutput), nil
	}
	out, err := c.Inspector2API.ListAccountPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListAccountPermissionsPages(input *inspector2.ListAccountPermissionsInput, fn func(*inspector2.ListAccountPermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListAccountPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListAccountPermissionsOutput{}
	fnCacher := func(out *inspector2.ListAccountPermissionsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListAccountPermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListAccountPermissionsPagesWithContext(ctx context.Context, input *inspector2.ListAccountPermissionsInput, fn func(*inspector2.ListAccountPermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListAccountPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListAccountPermissionsOutput{}
	fnCacher := func(out *inspector2.ListAccountPermissionsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListAccountPermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListAccountPermissionsWithContext(ctx context.Context, input *inspector2.ListAccountPermissionsInput, opts ...request.Option) (*inspector2.ListAccountPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListAccountPermissionsOutput), nil
	}
	out, err := c.Inspector2API.ListAccountPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListCoverage(input *inspector2.ListCoverageInput) (*inspector2.ListCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListCoverageOutput), nil
	}
	out, err := c.Inspector2API.ListCoverage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListCoveragePages(input *inspector2.ListCoverageInput, fn func(*inspector2.ListCoverageOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListCoverageOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListCoverageOutput{}
	fnCacher := func(out *inspector2.ListCoverageOutput, more bool) bool {
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
	if err := c.Inspector2API.ListCoveragePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListCoveragePagesWithContext(ctx context.Context, input *inspector2.ListCoverageInput, fn func(*inspector2.ListCoverageOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListCoverageOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListCoverageOutput{}
	fnCacher := func(out *inspector2.ListCoverageOutput, more bool) bool {
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
	if err := c.Inspector2API.ListCoveragePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListCoverageStatistics(input *inspector2.ListCoverageStatisticsInput) (*inspector2.ListCoverageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListCoverageStatisticsOutput), nil
	}
	out, err := c.Inspector2API.ListCoverageStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListCoverageStatisticsPages(input *inspector2.ListCoverageStatisticsInput, fn func(*inspector2.ListCoverageStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListCoverageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListCoverageStatisticsOutput{}
	fnCacher := func(out *inspector2.ListCoverageStatisticsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListCoverageStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListCoverageStatisticsPagesWithContext(ctx context.Context, input *inspector2.ListCoverageStatisticsInput, fn func(*inspector2.ListCoverageStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListCoverageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListCoverageStatisticsOutput{}
	fnCacher := func(out *inspector2.ListCoverageStatisticsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListCoverageStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListCoverageStatisticsWithContext(ctx context.Context, input *inspector2.ListCoverageStatisticsInput, opts ...request.Option) (*inspector2.ListCoverageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListCoverageStatisticsOutput), nil
	}
	out, err := c.Inspector2API.ListCoverageStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListCoverageWithContext(ctx context.Context, input *inspector2.ListCoverageInput, opts ...request.Option) (*inspector2.ListCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListCoverageOutput), nil
	}
	out, err := c.Inspector2API.ListCoverageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListDelegatedAdminAccounts(input *inspector2.ListDelegatedAdminAccountsInput) (*inspector2.ListDelegatedAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListDelegatedAdminAccountsOutput), nil
	}
	out, err := c.Inspector2API.ListDelegatedAdminAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListDelegatedAdminAccountsPages(input *inspector2.ListDelegatedAdminAccountsInput, fn func(*inspector2.ListDelegatedAdminAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListDelegatedAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListDelegatedAdminAccountsOutput{}
	fnCacher := func(out *inspector2.ListDelegatedAdminAccountsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListDelegatedAdminAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListDelegatedAdminAccountsPagesWithContext(ctx context.Context, input *inspector2.ListDelegatedAdminAccountsInput, fn func(*inspector2.ListDelegatedAdminAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListDelegatedAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListDelegatedAdminAccountsOutput{}
	fnCacher := func(out *inspector2.ListDelegatedAdminAccountsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListDelegatedAdminAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListDelegatedAdminAccountsWithContext(ctx context.Context, input *inspector2.ListDelegatedAdminAccountsInput, opts ...request.Option) (*inspector2.ListDelegatedAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListDelegatedAdminAccountsOutput), nil
	}
	out, err := c.Inspector2API.ListDelegatedAdminAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFilters(input *inspector2.ListFiltersInput) (*inspector2.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFiltersOutput), nil
	}
	out, err := c.Inspector2API.ListFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFiltersPages(input *inspector2.ListFiltersInput, fn func(*inspector2.ListFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFiltersOutput{}
	fnCacher := func(out *inspector2.ListFiltersOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFiltersPagesWithContext(ctx context.Context, input *inspector2.ListFiltersInput, fn func(*inspector2.ListFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFiltersOutput{}
	fnCacher := func(out *inspector2.ListFiltersOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFiltersWithContext(ctx context.Context, input *inspector2.ListFiltersInput, opts ...request.Option) (*inspector2.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFiltersOutput), nil
	}
	out, err := c.Inspector2API.ListFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFindingAggregations(input *inspector2.ListFindingAggregationsInput) (*inspector2.ListFindingAggregationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFindingAggregationsOutput), nil
	}
	out, err := c.Inspector2API.ListFindingAggregations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFindingAggregationsPages(input *inspector2.ListFindingAggregationsInput, fn func(*inspector2.ListFindingAggregationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFindingAggregationsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFindingAggregationsOutput{}
	fnCacher := func(out *inspector2.ListFindingAggregationsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFindingAggregationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFindingAggregationsPagesWithContext(ctx context.Context, input *inspector2.ListFindingAggregationsInput, fn func(*inspector2.ListFindingAggregationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFindingAggregationsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFindingAggregationsOutput{}
	fnCacher := func(out *inspector2.ListFindingAggregationsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFindingAggregationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFindingAggregationsWithContext(ctx context.Context, input *inspector2.ListFindingAggregationsInput, opts ...request.Option) (*inspector2.ListFindingAggregationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFindingAggregationsOutput), nil
	}
	out, err := c.Inspector2API.ListFindingAggregationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFindings(input *inspector2.ListFindingsInput) (*inspector2.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFindingsOutput), nil
	}
	out, err := c.Inspector2API.ListFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListFindingsPages(input *inspector2.ListFindingsInput, fn func(*inspector2.ListFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFindingsOutput{}
	fnCacher := func(out *inspector2.ListFindingsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFindingsPagesWithContext(ctx context.Context, input *inspector2.ListFindingsInput, fn func(*inspector2.ListFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListFindingsOutput{}
	fnCacher := func(out *inspector2.ListFindingsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListFindingsWithContext(ctx context.Context, input *inspector2.ListFindingsInput, opts ...request.Option) (*inspector2.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListFindingsOutput), nil
	}
	out, err := c.Inspector2API.ListFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListMembers(input *inspector2.ListMembersInput) (*inspector2.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListMembersOutput), nil
	}
	out, err := c.Inspector2API.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListMembersPages(input *inspector2.ListMembersInput, fn func(*inspector2.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListMembersOutput{}
	fnCacher := func(out *inspector2.ListMembersOutput, more bool) bool {
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
	if err := c.Inspector2API.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListMembersPagesWithContext(ctx context.Context, input *inspector2.ListMembersInput, fn func(*inspector2.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListMembersOutput{}
	fnCacher := func(out *inspector2.ListMembersOutput, more bool) bool {
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
	if err := c.Inspector2API.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListMembersWithContext(ctx context.Context, input *inspector2.ListMembersInput, opts ...request.Option) (*inspector2.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListMembersOutput), nil
	}
	out, err := c.Inspector2API.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListTagsForResource(input *inspector2.ListTagsForResourceInput) (*inspector2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListTagsForResourceOutput), nil
	}
	out, err := c.Inspector2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListTagsForResourceWithContext(ctx context.Context, input *inspector2.ListTagsForResourceInput, opts ...request.Option) (*inspector2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListTagsForResourceOutput), nil
	}
	out, err := c.Inspector2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListUsageTotals(input *inspector2.ListUsageTotalsInput) (*inspector2.ListUsageTotalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListUsageTotalsOutput), nil
	}
	out, err := c.Inspector2API.ListUsageTotals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector2) ListUsageTotalsPages(input *inspector2.ListUsageTotalsInput, fn func(*inspector2.ListUsageTotalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListUsageTotalsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListUsageTotalsOutput{}
	fnCacher := func(out *inspector2.ListUsageTotalsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListUsageTotalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListUsageTotalsPagesWithContext(ctx context.Context, input *inspector2.ListUsageTotalsInput, fn func(*inspector2.ListUsageTotalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector2.ListUsageTotalsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector2.ListUsageTotalsOutput{}
	fnCacher := func(out *inspector2.ListUsageTotalsOutput, more bool) bool {
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
	if err := c.Inspector2API.ListUsageTotalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector2) ListUsageTotalsWithContext(ctx context.Context, input *inspector2.ListUsageTotalsInput, opts ...request.Option) (*inspector2.ListUsageTotalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector2.ListUsageTotalsOutput), nil
	}
	out, err := c.Inspector2API.ListUsageTotalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
