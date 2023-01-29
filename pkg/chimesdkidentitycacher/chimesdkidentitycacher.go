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
package chimesdkidentitycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/chimesdkidentity"
	"github.com/aws/aws-sdk-go/service/chimesdkidentity/chimesdkidentityiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ChimeSDKIdentity struct {
	chimesdkidentityiface.ChimeSDKIdentityAPI
	cache *cache.Cache
}

func New(chimesdkidentityapi chimesdkidentityiface.ChimeSDKIdentityAPI) *ChimeSDKIdentity {
	return &ChimeSDKIdentity{
		ChimeSDKIdentityAPI: chimesdkidentityapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ChimeSDKIdentity) DescribeAppInstance(input *chimesdkidentity.DescribeAppInstanceInput) (*chimesdkidentity.DescribeAppInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceAdmin(input *chimesdkidentity.DescribeAppInstanceAdminInput) (*chimesdkidentity.DescribeAppInstanceAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceAdminOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceAdmin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceAdminWithContext(ctx context.Context, input *chimesdkidentity.DescribeAppInstanceAdminInput, opts ...request.Option) (*chimesdkidentity.DescribeAppInstanceAdminOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceAdminOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceAdminWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceUser(input *chimesdkidentity.DescribeAppInstanceUserInput) (*chimesdkidentity.DescribeAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceUserEndpoint(input *chimesdkidentity.DescribeAppInstanceUserEndpointInput) (*chimesdkidentity.DescribeAppInstanceUserEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceUserEndpointOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceUserEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceUserEndpointWithContext(ctx context.Context, input *chimesdkidentity.DescribeAppInstanceUserEndpointInput, opts ...request.Option) (*chimesdkidentity.DescribeAppInstanceUserEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceUserEndpointOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceUserEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceUserWithContext(ctx context.Context, input *chimesdkidentity.DescribeAppInstanceUserInput, opts ...request.Option) (*chimesdkidentity.DescribeAppInstanceUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceUserOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) DescribeAppInstanceWithContext(ctx context.Context, input *chimesdkidentity.DescribeAppInstanceInput, opts ...request.Option) (*chimesdkidentity.DescribeAppInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.DescribeAppInstanceOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.DescribeAppInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) GetAppInstanceRetentionSettings(input *chimesdkidentity.GetAppInstanceRetentionSettingsInput) (*chimesdkidentity.GetAppInstanceRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.GetAppInstanceRetentionSettingsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.GetAppInstanceRetentionSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) GetAppInstanceRetentionSettingsWithContext(ctx context.Context, input *chimesdkidentity.GetAppInstanceRetentionSettingsInput, opts ...request.Option) (*chimesdkidentity.GetAppInstanceRetentionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.GetAppInstanceRetentionSettingsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.GetAppInstanceRetentionSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceAdmins(input *chimesdkidentity.ListAppInstanceAdminsInput) (*chimesdkidentity.ListAppInstanceAdminsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceAdminsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceAdmins(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceAdminsPages(input *chimesdkidentity.ListAppInstanceAdminsInput, fn func(*chimesdkidentity.ListAppInstanceAdminsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceAdminsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceAdminsOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceAdminsOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceAdminsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceAdminsPagesWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceAdminsInput, fn func(*chimesdkidentity.ListAppInstanceAdminsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceAdminsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceAdminsOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceAdminsOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceAdminsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceAdminsWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceAdminsInput, opts ...request.Option) (*chimesdkidentity.ListAppInstanceAdminsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceAdminsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceAdminsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceUserEndpoints(input *chimesdkidentity.ListAppInstanceUserEndpointsInput) (*chimesdkidentity.ListAppInstanceUserEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceUserEndpointsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceUserEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceUserEndpointsPages(input *chimesdkidentity.ListAppInstanceUserEndpointsInput, fn func(*chimesdkidentity.ListAppInstanceUserEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceUserEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceUserEndpointsOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceUserEndpointsOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceUserEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceUserEndpointsPagesWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceUserEndpointsInput, fn func(*chimesdkidentity.ListAppInstanceUserEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceUserEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceUserEndpointsOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceUserEndpointsOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceUserEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceUserEndpointsWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceUserEndpointsInput, opts ...request.Option) (*chimesdkidentity.ListAppInstanceUserEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceUserEndpointsOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceUserEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceUsers(input *chimesdkidentity.ListAppInstanceUsersInput) (*chimesdkidentity.ListAppInstanceUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceUsersOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstanceUsersPages(input *chimesdkidentity.ListAppInstanceUsersInput, fn func(*chimesdkidentity.ListAppInstanceUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceUsersOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceUsersOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceUsersPagesWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceUsersInput, fn func(*chimesdkidentity.ListAppInstanceUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstanceUsersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstanceUsersOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstanceUsersOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstanceUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstanceUsersWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstanceUsersInput, opts ...request.Option) (*chimesdkidentity.ListAppInstanceUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstanceUsersOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstanceUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstances(input *chimesdkidentity.ListAppInstancesInput) (*chimesdkidentity.ListAppInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstancesOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListAppInstancesPages(input *chimesdkidentity.ListAppInstancesInput, fn func(*chimesdkidentity.ListAppInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstancesOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstancesOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstancesPagesWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstancesInput, fn func(*chimesdkidentity.ListAppInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkidentity.ListAppInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkidentity.ListAppInstancesOutput{}
	fnCacher := func(out *chimesdkidentity.ListAppInstancesOutput, more bool) bool {
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
	if err := c.ChimeSDKIdentityAPI.ListAppInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKIdentity) ListAppInstancesWithContext(ctx context.Context, input *chimesdkidentity.ListAppInstancesInput, opts ...request.Option) (*chimesdkidentity.ListAppInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListAppInstancesOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListAppInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListTagsForResource(input *chimesdkidentity.ListTagsForResourceInput) (*chimesdkidentity.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKIdentity) ListTagsForResourceWithContext(ctx context.Context, input *chimesdkidentity.ListTagsForResourceInput, opts ...request.Option) (*chimesdkidentity.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkidentity.ListTagsForResourceOutput), nil
	}
	out, err := c.ChimeSDKIdentityAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
