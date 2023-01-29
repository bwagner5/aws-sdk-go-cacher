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
	"github.com/aws/aws-sdk-go/service/identitystore"
	"github.com/aws/aws-sdk-go/service/identitystore/identitystoreiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IdentityStore struct {
	identitystoreiface.IdentityStoreAPI
	cache *cache.Cache
}

func NewIdentityStore(identitystoreapi identitystoreiface.IdentityStoreAPI) *IdentityStore {
	return &IdentityStore{
		IdentityStoreAPI: identitystoreapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IdentityStore) DescribeGroup(input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeGroupOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) DescribeGroupMembership(input *identitystore.DescribeGroupMembershipInput) (*identitystore.DescribeGroupMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeGroupMembershipOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeGroupMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) DescribeGroupMembershipWithContext(ctx context.Context, input *identitystore.DescribeGroupMembershipInput, opts ...request.Option) (*identitystore.DescribeGroupMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeGroupMembershipOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeGroupMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) DescribeGroupWithContext(ctx context.Context, input *identitystore.DescribeGroupInput, opts ...request.Option) (*identitystore.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeGroupOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) DescribeUser(input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeUserOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) DescribeUserWithContext(ctx context.Context, input *identitystore.DescribeUserInput, opts ...request.Option) (*identitystore.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.DescribeUserOutput), nil
	}
	out, err := c.IdentityStoreAPI.DescribeUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetGroupId(input *identitystore.GetGroupIdInput) (*identitystore.GetGroupIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetGroupIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetGroupId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetGroupIdWithContext(ctx context.Context, input *identitystore.GetGroupIdInput, opts ...request.Option) (*identitystore.GetGroupIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetGroupIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetGroupIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetGroupMembershipId(input *identitystore.GetGroupMembershipIdInput) (*identitystore.GetGroupMembershipIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetGroupMembershipIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetGroupMembershipId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetGroupMembershipIdWithContext(ctx context.Context, input *identitystore.GetGroupMembershipIdInput, opts ...request.Option) (*identitystore.GetGroupMembershipIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetGroupMembershipIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetGroupMembershipIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetUserId(input *identitystore.GetUserIdInput) (*identitystore.GetUserIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetUserIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetUserId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) GetUserIdWithContext(ctx context.Context, input *identitystore.GetUserIdInput, opts ...request.Option) (*identitystore.GetUserIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.GetUserIdOutput), nil
	}
	out, err := c.IdentityStoreAPI.GetUserIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroupMemberships(input *identitystore.ListGroupMembershipsInput) (*identitystore.ListGroupMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupMembershipsOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroupMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroupMembershipsForMember(input *identitystore.ListGroupMembershipsForMemberInput) (*identitystore.ListGroupMembershipsForMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupMembershipsForMemberOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroupMembershipsForMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroupMembershipsForMemberPages(input *identitystore.ListGroupMembershipsForMemberInput, fn func(*identitystore.ListGroupMembershipsForMemberOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupMembershipsForMemberOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupMembershipsForMemberOutput{}
	fnCacher := func(out *identitystore.ListGroupMembershipsForMemberOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupMembershipsForMemberPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupMembershipsForMemberPagesWithContext(ctx context.Context, input *identitystore.ListGroupMembershipsForMemberInput, fn func(*identitystore.ListGroupMembershipsForMemberOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupMembershipsForMemberOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupMembershipsForMemberOutput{}
	fnCacher := func(out *identitystore.ListGroupMembershipsForMemberOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupMembershipsForMemberPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupMembershipsForMemberWithContext(ctx context.Context, input *identitystore.ListGroupMembershipsForMemberInput, opts ...request.Option) (*identitystore.ListGroupMembershipsForMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupMembershipsForMemberOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroupMembershipsForMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroupMembershipsPages(input *identitystore.ListGroupMembershipsInput, fn func(*identitystore.ListGroupMembershipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupMembershipsOutput{}
	fnCacher := func(out *identitystore.ListGroupMembershipsOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupMembershipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupMembershipsPagesWithContext(ctx context.Context, input *identitystore.ListGroupMembershipsInput, fn func(*identitystore.ListGroupMembershipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupMembershipsOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupMembershipsOutput{}
	fnCacher := func(out *identitystore.ListGroupMembershipsOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupMembershipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupMembershipsWithContext(ctx context.Context, input *identitystore.ListGroupMembershipsInput, opts ...request.Option) (*identitystore.ListGroupMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupMembershipsOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroupMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroups(input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupsOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListGroupsPages(input *identitystore.ListGroupsInput, fn func(*identitystore.ListGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupsOutput{}
	fnCacher := func(out *identitystore.ListGroupsOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupsPagesWithContext(ctx context.Context, input *identitystore.ListGroupsInput, fn func(*identitystore.ListGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListGroupsOutput{}
	fnCacher := func(out *identitystore.ListGroupsOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListGroupsWithContext(ctx context.Context, input *identitystore.ListGroupsInput, opts ...request.Option) (*identitystore.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListGroupsOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListUsers(input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListUsersOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IdentityStore) ListUsersPages(input *identitystore.ListUsersInput, fn func(*identitystore.ListUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListUsersOutput{}
	fnCacher := func(out *identitystore.ListUsersOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListUsersPagesWithContext(ctx context.Context, input *identitystore.ListUsersInput, fn func(*identitystore.ListUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*identitystore.ListUsersOutput), false)
		return nil
	}
	cachable := true
	output := &identitystore.ListUsersOutput{}
	fnCacher := func(out *identitystore.ListUsersOutput, more bool) bool {
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
	if err := c.IdentityStoreAPI.ListUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IdentityStore) ListUsersWithContext(ctx context.Context, input *identitystore.ListUsersInput, opts ...request.Option) (*identitystore.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*identitystore.ListUsersOutput), nil
	}
	out, err := c.IdentityStoreAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
