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
package detectivecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/detective"
	"github.com/aws/aws-sdk-go/service/detective/detectiveiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Detective struct {
	detectiveiface.DetectiveAPI
	cache *cache.Cache
}

func New(detectiveapi detectiveiface.DetectiveAPI) *Detective {
	return &Detective{
		DetectiveAPI: detectiveapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Detective) BatchGetGraphMemberDatasources(input *detective.BatchGetGraphMemberDatasourcesInput) (*detective.BatchGetGraphMemberDatasourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.BatchGetGraphMemberDatasourcesOutput), nil
	}
	out, err := c.DetectiveAPI.BatchGetGraphMemberDatasources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) BatchGetGraphMemberDatasourcesWithContext(ctx context.Context, input *detective.BatchGetGraphMemberDatasourcesInput, opts ...request.Option) (*detective.BatchGetGraphMemberDatasourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.BatchGetGraphMemberDatasourcesOutput), nil
	}
	out, err := c.DetectiveAPI.BatchGetGraphMemberDatasourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) BatchGetMembershipDatasources(input *detective.BatchGetMembershipDatasourcesInput) (*detective.BatchGetMembershipDatasourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.BatchGetMembershipDatasourcesOutput), nil
	}
	out, err := c.DetectiveAPI.BatchGetMembershipDatasources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) BatchGetMembershipDatasourcesWithContext(ctx context.Context, input *detective.BatchGetMembershipDatasourcesInput, opts ...request.Option) (*detective.BatchGetMembershipDatasourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.BatchGetMembershipDatasourcesOutput), nil
	}
	out, err := c.DetectiveAPI.BatchGetMembershipDatasourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) DescribeOrganizationConfiguration(input *detective.DescribeOrganizationConfigurationInput) (*detective.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.DetectiveAPI.DescribeOrganizationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) DescribeOrganizationConfigurationWithContext(ctx context.Context, input *detective.DescribeOrganizationConfigurationInput, opts ...request.Option) (*detective.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.DetectiveAPI.DescribeOrganizationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) GetMembers(input *detective.GetMembersInput) (*detective.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.GetMembersOutput), nil
	}
	out, err := c.DetectiveAPI.GetMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) GetMembersWithContext(ctx context.Context, input *detective.GetMembersInput, opts ...request.Option) (*detective.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.GetMembersOutput), nil
	}
	out, err := c.DetectiveAPI.GetMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListDatasourcePackages(input *detective.ListDatasourcePackagesInput) (*detective.ListDatasourcePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListDatasourcePackagesOutput), nil
	}
	out, err := c.DetectiveAPI.ListDatasourcePackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListDatasourcePackagesPages(input *detective.ListDatasourcePackagesInput, fn func(*detective.ListDatasourcePackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListDatasourcePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListDatasourcePackagesOutput{}
	fnCacher := func(out *detective.ListDatasourcePackagesOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListDatasourcePackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListDatasourcePackagesPagesWithContext(ctx context.Context, input *detective.ListDatasourcePackagesInput, fn func(*detective.ListDatasourcePackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListDatasourcePackagesOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListDatasourcePackagesOutput{}
	fnCacher := func(out *detective.ListDatasourcePackagesOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListDatasourcePackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListDatasourcePackagesWithContext(ctx context.Context, input *detective.ListDatasourcePackagesInput, opts ...request.Option) (*detective.ListDatasourcePackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListDatasourcePackagesOutput), nil
	}
	out, err := c.DetectiveAPI.ListDatasourcePackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListGraphs(input *detective.ListGraphsInput) (*detective.ListGraphsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListGraphsOutput), nil
	}
	out, err := c.DetectiveAPI.ListGraphs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListGraphsPages(input *detective.ListGraphsInput, fn func(*detective.ListGraphsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListGraphsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListGraphsOutput{}
	fnCacher := func(out *detective.ListGraphsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListGraphsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListGraphsPagesWithContext(ctx context.Context, input *detective.ListGraphsInput, fn func(*detective.ListGraphsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListGraphsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListGraphsOutput{}
	fnCacher := func(out *detective.ListGraphsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListGraphsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListGraphsWithContext(ctx context.Context, input *detective.ListGraphsInput, opts ...request.Option) (*detective.ListGraphsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListGraphsOutput), nil
	}
	out, err := c.DetectiveAPI.ListGraphsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListInvitations(input *detective.ListInvitationsInput) (*detective.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListInvitationsOutput), nil
	}
	out, err := c.DetectiveAPI.ListInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListInvitationsPages(input *detective.ListInvitationsInput, fn func(*detective.ListInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListInvitationsOutput{}
	fnCacher := func(out *detective.ListInvitationsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListInvitationsPagesWithContext(ctx context.Context, input *detective.ListInvitationsInput, fn func(*detective.ListInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListInvitationsOutput{}
	fnCacher := func(out *detective.ListInvitationsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListInvitationsWithContext(ctx context.Context, input *detective.ListInvitationsInput, opts ...request.Option) (*detective.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListInvitationsOutput), nil
	}
	out, err := c.DetectiveAPI.ListInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListMembers(input *detective.ListMembersInput) (*detective.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListMembersOutput), nil
	}
	out, err := c.DetectiveAPI.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListMembersPages(input *detective.ListMembersInput, fn func(*detective.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListMembersOutput{}
	fnCacher := func(out *detective.ListMembersOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListMembersPagesWithContext(ctx context.Context, input *detective.ListMembersInput, fn func(*detective.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListMembersOutput{}
	fnCacher := func(out *detective.ListMembersOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListMembersWithContext(ctx context.Context, input *detective.ListMembersInput, opts ...request.Option) (*detective.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListMembersOutput), nil
	}
	out, err := c.DetectiveAPI.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListOrganizationAdminAccounts(input *detective.ListOrganizationAdminAccountsInput) (*detective.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.DetectiveAPI.ListOrganizationAdminAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListOrganizationAdminAccountsPages(input *detective.ListOrganizationAdminAccountsInput, fn func(*detective.ListOrganizationAdminAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *detective.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListOrganizationAdminAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListOrganizationAdminAccountsPagesWithContext(ctx context.Context, input *detective.ListOrganizationAdminAccountsInput, fn func(*detective.ListOrganizationAdminAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*detective.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &detective.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *detective.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.DetectiveAPI.ListOrganizationAdminAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Detective) ListOrganizationAdminAccountsWithContext(ctx context.Context, input *detective.ListOrganizationAdminAccountsInput, opts ...request.Option) (*detective.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.DetectiveAPI.ListOrganizationAdminAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListTagsForResource(input *detective.ListTagsForResourceInput) (*detective.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListTagsForResourceOutput), nil
	}
	out, err := c.DetectiveAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Detective) ListTagsForResourceWithContext(ctx context.Context, input *detective.ListTagsForResourceInput, opts ...request.Option) (*detective.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*detective.ListTagsForResourceOutput), nil
	}
	out, err := c.DetectiveAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
