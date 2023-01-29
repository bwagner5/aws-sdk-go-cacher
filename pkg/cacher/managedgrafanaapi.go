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
	"github.com/aws/aws-sdk-go/service/managedgrafana"
	"github.com/aws/aws-sdk-go/service/managedgrafana/managedgrafanaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ManagedGrafana struct {
	managedgrafanaiface.ManagedGrafanaAPI
	cache *cache.Cache
}

func NewManagedGrafana(managedgrafanaapi managedgrafanaiface.ManagedGrafanaAPI) *ManagedGrafana {
	return &ManagedGrafana{
		ManagedGrafanaAPI: managedgrafanaapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ManagedGrafana) DescribeWorkspace(input *managedgrafana.DescribeWorkspaceInput) (*managedgrafana.DescribeWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) DescribeWorkspaceAuthentication(input *managedgrafana.DescribeWorkspaceAuthenticationInput) (*managedgrafana.DescribeWorkspaceAuthenticationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceAuthenticationOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspaceAuthentication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) DescribeWorkspaceAuthenticationWithContext(ctx context.Context, input *managedgrafana.DescribeWorkspaceAuthenticationInput, opts ...request.Option) (*managedgrafana.DescribeWorkspaceAuthenticationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceAuthenticationOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspaceAuthenticationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) DescribeWorkspaceConfiguration(input *managedgrafana.DescribeWorkspaceConfigurationInput) (*managedgrafana.DescribeWorkspaceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceConfigurationOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspaceConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) DescribeWorkspaceConfigurationWithContext(ctx context.Context, input *managedgrafana.DescribeWorkspaceConfigurationInput, opts ...request.Option) (*managedgrafana.DescribeWorkspaceConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceConfigurationOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspaceConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) DescribeWorkspaceWithContext(ctx context.Context, input *managedgrafana.DescribeWorkspaceInput, opts ...request.Option) (*managedgrafana.DescribeWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.DescribeWorkspaceOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.DescribeWorkspaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListPermissions(input *managedgrafana.ListPermissionsInput) (*managedgrafana.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListPermissionsOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListPermissionsPages(input *managedgrafana.ListPermissionsInput, fn func(*managedgrafana.ListPermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedgrafana.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &managedgrafana.ListPermissionsOutput{}
	fnCacher := func(out *managedgrafana.ListPermissionsOutput, more bool) bool {
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
	if err := c.ManagedGrafanaAPI.ListPermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedGrafana) ListPermissionsPagesWithContext(ctx context.Context, input *managedgrafana.ListPermissionsInput, fn func(*managedgrafana.ListPermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedgrafana.ListPermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &managedgrafana.ListPermissionsOutput{}
	fnCacher := func(out *managedgrafana.ListPermissionsOutput, more bool) bool {
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
	if err := c.ManagedGrafanaAPI.ListPermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedGrafana) ListPermissionsWithContext(ctx context.Context, input *managedgrafana.ListPermissionsInput, opts ...request.Option) (*managedgrafana.ListPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListPermissionsOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListTagsForResource(input *managedgrafana.ListTagsForResourceInput) (*managedgrafana.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListTagsForResourceOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListTagsForResourceWithContext(ctx context.Context, input *managedgrafana.ListTagsForResourceInput, opts ...request.Option) (*managedgrafana.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListTagsForResourceOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListWorkspaces(input *managedgrafana.ListWorkspacesInput) (*managedgrafana.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListWorkspacesOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListWorkspaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ManagedGrafana) ListWorkspacesPages(input *managedgrafana.ListWorkspacesInput, fn func(*managedgrafana.ListWorkspacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedgrafana.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &managedgrafana.ListWorkspacesOutput{}
	fnCacher := func(out *managedgrafana.ListWorkspacesOutput, more bool) bool {
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
	if err := c.ManagedGrafanaAPI.ListWorkspacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedGrafana) ListWorkspacesPagesWithContext(ctx context.Context, input *managedgrafana.ListWorkspacesInput, fn func(*managedgrafana.ListWorkspacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*managedgrafana.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &managedgrafana.ListWorkspacesOutput{}
	fnCacher := func(out *managedgrafana.ListWorkspacesOutput, more bool) bool {
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
	if err := c.ManagedGrafanaAPI.ListWorkspacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ManagedGrafana) ListWorkspacesWithContext(ctx context.Context, input *managedgrafana.ListWorkspacesInput, opts ...request.Option) (*managedgrafana.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*managedgrafana.ListWorkspacesOutput), nil
	}
	out, err := c.ManagedGrafanaAPI.ListWorkspacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
