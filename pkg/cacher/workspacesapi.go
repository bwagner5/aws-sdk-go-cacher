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
	"github.com/aws/aws-sdk-go/service/workspaces"
	"github.com/aws/aws-sdk-go/service/workspaces/workspacesiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WorkSpaces struct {
	workspacesiface.WorkSpacesAPI
	cache *cache.Cache
}

func NewWorkSpaces(workspacesapi workspacesiface.WorkSpacesAPI) *WorkSpaces {
	return &WorkSpaces{
		WorkSpacesAPI: workspacesapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WorkSpaces) DescribeAccount(input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeAccountOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeAccountModifications(input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeAccountModificationsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeAccountModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeAccountModificationsWithContext(ctx context.Context, input *workspaces.DescribeAccountModificationsInput, opts ...request.Option) (*workspaces.DescribeAccountModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeAccountModificationsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeAccountModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeAccountWithContext(ctx context.Context, input *workspaces.DescribeAccountInput, opts ...request.Option) (*workspaces.DescribeAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeAccountOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeClientBranding(input *workspaces.DescribeClientBrandingInput) (*workspaces.DescribeClientBrandingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeClientBrandingOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeClientBranding(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeClientBrandingWithContext(ctx context.Context, input *workspaces.DescribeClientBrandingInput, opts ...request.Option) (*workspaces.DescribeClientBrandingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeClientBrandingOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeClientBrandingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeClientProperties(input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeClientPropertiesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeClientProperties(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeClientPropertiesWithContext(ctx context.Context, input *workspaces.DescribeClientPropertiesInput, opts ...request.Option) (*workspaces.DescribeClientPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeClientPropertiesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeClientPropertiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectClientAddIns(input *workspaces.DescribeConnectClientAddInsInput) (*workspaces.DescribeConnectClientAddInsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectClientAddInsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectClientAddIns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectClientAddInsWithContext(ctx context.Context, input *workspaces.DescribeConnectClientAddInsInput, opts ...request.Option) (*workspaces.DescribeConnectClientAddInsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectClientAddInsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectClientAddInsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectionAliasPermissions(input *workspaces.DescribeConnectionAliasPermissionsInput) (*workspaces.DescribeConnectionAliasPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectionAliasPermissionsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectionAliasPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectionAliasPermissionsWithContext(ctx context.Context, input *workspaces.DescribeConnectionAliasPermissionsInput, opts ...request.Option) (*workspaces.DescribeConnectionAliasPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectionAliasPermissionsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectionAliasPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectionAliases(input *workspaces.DescribeConnectionAliasesInput) (*workspaces.DescribeConnectionAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectionAliasesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectionAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeConnectionAliasesWithContext(ctx context.Context, input *workspaces.DescribeConnectionAliasesInput, opts ...request.Option) (*workspaces.DescribeConnectionAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeConnectionAliasesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeConnectionAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeIpGroups(input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeIpGroupsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeIpGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeIpGroupsWithContext(ctx context.Context, input *workspaces.DescribeIpGroupsInput, opts ...request.Option) (*workspaces.DescribeIpGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeIpGroupsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeIpGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeTags(input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeTagsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeTagsWithContext(ctx context.Context, input *workspaces.DescribeTagsInput, opts ...request.Option) (*workspaces.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeTagsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceBundles(input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceBundlesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceBundles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceBundlesPages(input *workspaces.DescribeWorkspaceBundlesInput, fn func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspaceBundlesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspaceBundlesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspaceBundlesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspaceBundlesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspaceBundlesPagesWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceBundlesInput, fn func(*workspaces.DescribeWorkspaceBundlesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspaceBundlesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspaceBundlesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspaceBundlesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspaceBundlesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspaceBundlesWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceBundlesInput, opts ...request.Option) (*workspaces.DescribeWorkspaceBundlesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceBundlesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceBundlesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceDirectories(input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceDirectoriesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceDirectories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceDirectoriesPages(input *workspaces.DescribeWorkspaceDirectoriesInput, fn func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspaceDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspaceDirectoriesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspaceDirectoriesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspaceDirectoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspaceDirectoriesPagesWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceDirectoriesInput, fn func(*workspaces.DescribeWorkspaceDirectoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspaceDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspaceDirectoriesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspaceDirectoriesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspaceDirectoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspaceDirectoriesWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceDirectoriesInput, opts ...request.Option) (*workspaces.DescribeWorkspaceDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceDirectoriesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceDirectoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceImagePermissions(input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceImagePermissionsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceImagePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceImagePermissionsWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput, opts ...request.Option) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceImagePermissionsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceImagePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceImages(input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceImagesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceImagesWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceImagesInput, opts ...request.Option) (*workspaces.DescribeWorkspaceImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceImagesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceSnapshots(input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceSnapshotsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaceSnapshotsWithContext(ctx context.Context, input *workspaces.DescribeWorkspaceSnapshotsInput, opts ...request.Option) (*workspaces.DescribeWorkspaceSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspaceSnapshotsOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaceSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspaces(input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspacesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspacesConnectionStatus(input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspacesConnectionStatusOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspacesConnectionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspacesConnectionStatusWithContext(ctx context.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput, opts ...request.Option) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspacesConnectionStatusOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspacesConnectionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) DescribeWorkspacesPages(input *workspaces.DescribeWorkspacesInput, fn func(*workspaces.DescribeWorkspacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspacesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspacesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspacesPagesWithContext(ctx context.Context, input *workspaces.DescribeWorkspacesInput, fn func(*workspaces.DescribeWorkspacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*workspaces.DescribeWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &workspaces.DescribeWorkspacesOutput{}
	fnCacher := func(out *workspaces.DescribeWorkspacesOutput, more bool) bool {
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
	if err := c.WorkSpacesAPI.DescribeWorkspacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WorkSpaces) DescribeWorkspacesWithContext(ctx context.Context, input *workspaces.DescribeWorkspacesInput, opts ...request.Option) (*workspaces.DescribeWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.DescribeWorkspacesOutput), nil
	}
	out, err := c.WorkSpacesAPI.DescribeWorkspacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) ListAvailableManagementCidrRanges(input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.ListAvailableManagementCidrRangesOutput), nil
	}
	out, err := c.WorkSpacesAPI.ListAvailableManagementCidrRanges(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WorkSpaces) ListAvailableManagementCidrRangesWithContext(ctx context.Context, input *workspaces.ListAvailableManagementCidrRangesInput, opts ...request.Option) (*workspaces.ListAvailableManagementCidrRangesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*workspaces.ListAvailableManagementCidrRangesOutput), nil
	}
	out, err := c.WorkSpacesAPI.ListAvailableManagementCidrRangesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
