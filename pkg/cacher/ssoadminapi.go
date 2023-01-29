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
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/aws/aws-sdk-go/service/ssoadmin/ssoadminiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SSOAdmin struct {
	ssoadminiface.SSOAdminAPI
	cache *cache.Cache
}

func NewSSOAdmin(ssoadminapi ssoadminiface.SSOAdminAPI) *SSOAdmin {
	return &SSOAdmin{
		SSOAdminAPI: ssoadminapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SSOAdmin) DescribeAccountAssignmentCreationStatus(input *ssoadmin.DescribeAccountAssignmentCreationStatusInput) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeAccountAssignmentCreationStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeAccountAssignmentCreationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribeAccountAssignmentCreationStatusWithContext(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentCreationStatusInput, opts ...request.Option) (*ssoadmin.DescribeAccountAssignmentCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeAccountAssignmentCreationStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeAccountAssignmentCreationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribeAccountAssignmentDeletionStatus(input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeAccountAssignmentDeletionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribeAccountAssignmentDeletionStatusWithContext(ctx context.Context, input *ssoadmin.DescribeAccountAssignmentDeletionStatusInput, opts ...request.Option) (*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeAccountAssignmentDeletionStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeAccountAssignmentDeletionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribeInstanceAccessControlAttributeConfiguration(input *ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput) (*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeInstanceAccessControlAttributeConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribeInstanceAccessControlAttributeConfigurationWithContext(ctx context.Context, input *ssoadmin.DescribeInstanceAccessControlAttributeConfigurationInput, opts ...request.Option) (*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribeInstanceAccessControlAttributeConfigurationOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribeInstanceAccessControlAttributeConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribePermissionSet(input *ssoadmin.DescribePermissionSetInput) (*ssoadmin.DescribePermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribePermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribePermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribePermissionSetProvisioningStatus(input *ssoadmin.DescribePermissionSetProvisioningStatusInput) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribePermissionSetProvisioningStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribePermissionSetProvisioningStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribePermissionSetProvisioningStatusWithContext(ctx context.Context, input *ssoadmin.DescribePermissionSetProvisioningStatusInput, opts ...request.Option) (*ssoadmin.DescribePermissionSetProvisioningStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribePermissionSetProvisioningStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribePermissionSetProvisioningStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) DescribePermissionSetWithContext(ctx context.Context, input *ssoadmin.DescribePermissionSetInput, opts ...request.Option) (*ssoadmin.DescribePermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.DescribePermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.DescribePermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) GetInlinePolicyForPermissionSet(input *ssoadmin.GetInlinePolicyForPermissionSetInput) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.GetInlinePolicyForPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.GetInlinePolicyForPermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) GetInlinePolicyForPermissionSetWithContext(ctx context.Context, input *ssoadmin.GetInlinePolicyForPermissionSetInput, opts ...request.Option) (*ssoadmin.GetInlinePolicyForPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.GetInlinePolicyForPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.GetInlinePolicyForPermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) GetPermissionsBoundaryForPermissionSet(input *ssoadmin.GetPermissionsBoundaryForPermissionSetInput) (*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.GetPermissionsBoundaryForPermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) GetPermissionsBoundaryForPermissionSetWithContext(ctx context.Context, input *ssoadmin.GetPermissionsBoundaryForPermissionSetInput, opts ...request.Option) (*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.GetPermissionsBoundaryForPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.GetPermissionsBoundaryForPermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignmentCreationStatus(input *ssoadmin.ListAccountAssignmentCreationStatusInput) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentCreationStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignmentCreationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignmentCreationStatusPages(input *ssoadmin.ListAccountAssignmentCreationStatusInput, fn func(*ssoadmin.ListAccountAssignmentCreationStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentCreationStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentCreationStatusOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentCreationStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentCreationStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentCreationStatusPagesWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput, fn func(*ssoadmin.ListAccountAssignmentCreationStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentCreationStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentCreationStatusOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentCreationStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentCreationStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentCreationStatusWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentCreationStatusInput, opts ...request.Option) (*ssoadmin.ListAccountAssignmentCreationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentCreationStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignmentCreationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignmentDeletionStatus(input *ssoadmin.ListAccountAssignmentDeletionStatusInput) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentDeletionStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignmentDeletionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignmentDeletionStatusPages(input *ssoadmin.ListAccountAssignmentDeletionStatusInput, fn func(*ssoadmin.ListAccountAssignmentDeletionStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentDeletionStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentDeletionStatusOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentDeletionStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentDeletionStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentDeletionStatusPagesWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput, fn func(*ssoadmin.ListAccountAssignmentDeletionStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentDeletionStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentDeletionStatusOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentDeletionStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentDeletionStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentDeletionStatusWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentDeletionStatusInput, opts ...request.Option) (*ssoadmin.ListAccountAssignmentDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentDeletionStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignmentDeletionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignments(input *ssoadmin.ListAccountAssignmentsInput) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentsOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountAssignmentsPages(input *ssoadmin.ListAccountAssignmentsInput, fn func(*ssoadmin.ListAccountAssignmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentsOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentsOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentsPagesWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentsInput, fn func(*ssoadmin.ListAccountAssignmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountAssignmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountAssignmentsOutput{}
	fnCacher := func(out *ssoadmin.ListAccountAssignmentsOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountAssignmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountAssignmentsWithContext(ctx context.Context, input *ssoadmin.ListAccountAssignmentsInput, opts ...request.Option) (*ssoadmin.ListAccountAssignmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountAssignmentsOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountAssignmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountsForProvisionedPermissionSet(input *ssoadmin.ListAccountsForProvisionedPermissionSetInput) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountsForProvisionedPermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListAccountsForProvisionedPermissionSetPages(input *ssoadmin.ListAccountsForProvisionedPermissionSetInput, fn func(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountsForProvisionedPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListAccountsForProvisionedPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountsForProvisionedPermissionSetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountsForProvisionedPermissionSetPagesWithContext(ctx context.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput, fn func(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListAccountsForProvisionedPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListAccountsForProvisionedPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListAccountsForProvisionedPermissionSetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListAccountsForProvisionedPermissionSetWithContext(ctx context.Context, input *ssoadmin.ListAccountsForProvisionedPermissionSetInput, opts ...request.Option) (*ssoadmin.ListAccountsForProvisionedPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListAccountsForProvisionedPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListAccountsForProvisionedPermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListCustomerManagedPolicyReferencesInPermissionSet(input *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput) (*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListCustomerManagedPolicyReferencesInPermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListCustomerManagedPolicyReferencesInPermissionSetPages(input *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, fn func(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListCustomerManagedPolicyReferencesInPermissionSetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListCustomerManagedPolicyReferencesInPermissionSetPagesWithContext(ctx context.Context, input *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, fn func(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListCustomerManagedPolicyReferencesInPermissionSetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListCustomerManagedPolicyReferencesInPermissionSetWithContext(ctx context.Context, input *ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetInput, opts ...request.Option) (*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListCustomerManagedPolicyReferencesInPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListCustomerManagedPolicyReferencesInPermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListInstances(input *ssoadmin.ListInstancesInput) (*ssoadmin.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListInstancesOutput), nil
	}
	out, err := c.SSOAdminAPI.ListInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListInstancesPages(input *ssoadmin.ListInstancesInput, fn func(*ssoadmin.ListInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListInstancesOutput{}
	fnCacher := func(out *ssoadmin.ListInstancesOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListInstancesPagesWithContext(ctx context.Context, input *ssoadmin.ListInstancesInput, fn func(*ssoadmin.ListInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListInstancesOutput{}
	fnCacher := func(out *ssoadmin.ListInstancesOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListInstancesWithContext(ctx context.Context, input *ssoadmin.ListInstancesInput, opts ...request.Option) (*ssoadmin.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListInstancesOutput), nil
	}
	out, err := c.SSOAdminAPI.ListInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListManagedPoliciesInPermissionSet(input *ssoadmin.ListManagedPoliciesInPermissionSetInput) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListManagedPoliciesInPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListManagedPoliciesInPermissionSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListManagedPoliciesInPermissionSetPages(input *ssoadmin.ListManagedPoliciesInPermissionSetInput, fn func(*ssoadmin.ListManagedPoliciesInPermissionSetOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListManagedPoliciesInPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListManagedPoliciesInPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListManagedPoliciesInPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListManagedPoliciesInPermissionSetPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListManagedPoliciesInPermissionSetPagesWithContext(ctx context.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput, fn func(*ssoadmin.ListManagedPoliciesInPermissionSetOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListManagedPoliciesInPermissionSetOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListManagedPoliciesInPermissionSetOutput{}
	fnCacher := func(out *ssoadmin.ListManagedPoliciesInPermissionSetOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListManagedPoliciesInPermissionSetPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListManagedPoliciesInPermissionSetWithContext(ctx context.Context, input *ssoadmin.ListManagedPoliciesInPermissionSetInput, opts ...request.Option) (*ssoadmin.ListManagedPoliciesInPermissionSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListManagedPoliciesInPermissionSetOutput), nil
	}
	out, err := c.SSOAdminAPI.ListManagedPoliciesInPermissionSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSetProvisioningStatus(input *ssoadmin.ListPermissionSetProvisioningStatusInput) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetProvisioningStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSetProvisioningStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSetProvisioningStatusPages(input *ssoadmin.ListPermissionSetProvisioningStatusInput, fn func(*ssoadmin.ListPermissionSetProvisioningStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetProvisioningStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetProvisioningStatusOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetProvisioningStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetProvisioningStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetProvisioningStatusPagesWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput, fn func(*ssoadmin.ListPermissionSetProvisioningStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetProvisioningStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetProvisioningStatusOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetProvisioningStatusOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetProvisioningStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetProvisioningStatusWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetProvisioningStatusInput, opts ...request.Option) (*ssoadmin.ListPermissionSetProvisioningStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetProvisioningStatusOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSetProvisioningStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSets(input *ssoadmin.ListPermissionSetsInput) (*ssoadmin.ListPermissionSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetsOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSetsPages(input *ssoadmin.ListPermissionSetsInput, fn func(*ssoadmin.ListPermissionSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetsOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetsOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetsPagesWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetsInput, fn func(*ssoadmin.ListPermissionSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetsOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetsOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetsProvisionedToAccount(input *ssoadmin.ListPermissionSetsProvisionedToAccountInput) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSetsProvisionedToAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSetsProvisionedToAccountPages(input *ssoadmin.ListPermissionSetsProvisionedToAccountInput, fn func(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetsProvisionedToAccountOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetsProvisionedToAccountOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetsProvisionedToAccountPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetsProvisionedToAccountPagesWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput, fn func(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListPermissionSetsProvisionedToAccountOutput{}
	fnCacher := func(out *ssoadmin.ListPermissionSetsProvisionedToAccountOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListPermissionSetsProvisionedToAccountPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListPermissionSetsProvisionedToAccountWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetsProvisionedToAccountInput, opts ...request.Option) (*ssoadmin.ListPermissionSetsProvisionedToAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetsProvisionedToAccountOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSetsProvisionedToAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListPermissionSetsWithContext(ctx context.Context, input *ssoadmin.ListPermissionSetsInput, opts ...request.Option) (*ssoadmin.ListPermissionSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListPermissionSetsOutput), nil
	}
	out, err := c.SSOAdminAPI.ListPermissionSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListTagsForResource(input *ssoadmin.ListTagsForResourceInput) (*ssoadmin.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListTagsForResourceOutput), nil
	}
	out, err := c.SSOAdminAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSOAdmin) ListTagsForResourcePages(input *ssoadmin.ListTagsForResourceInput, fn func(*ssoadmin.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListTagsForResourceOutput{}
	fnCacher := func(out *ssoadmin.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListTagsForResourcePagesWithContext(ctx context.Context, input *ssoadmin.ListTagsForResourceInput, fn func(*ssoadmin.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssoadmin.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &ssoadmin.ListTagsForResourceOutput{}
	fnCacher := func(out *ssoadmin.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.SSOAdminAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSOAdmin) ListTagsForResourceWithContext(ctx context.Context, input *ssoadmin.ListTagsForResourceInput, opts ...request.Option) (*ssoadmin.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssoadmin.ListTagsForResourceOutput), nil
	}
	out, err := c.SSOAdminAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
