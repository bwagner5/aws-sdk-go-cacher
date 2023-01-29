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
package efscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EFS struct {
	efsiface.EFSAPI
	cache *cache.Cache
}

func New(efsapi efsiface.EFSAPI) *EFS {
	return &EFS{
		EFSAPI: efsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EFS) DescribeAccessPoints(input *efs.DescribeAccessPointsInput) (*efs.DescribeAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeAccessPointsOutput), nil
	}
	out, err := c.EFSAPI.DescribeAccessPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeAccessPointsPages(input *efs.DescribeAccessPointsInput, fn func(*efs.DescribeAccessPointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeAccessPointsOutput{}
	fnCacher := func(out *efs.DescribeAccessPointsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeAccessPointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeAccessPointsPagesWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, fn func(*efs.DescribeAccessPointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeAccessPointsOutput{}
	fnCacher := func(out *efs.DescribeAccessPointsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeAccessPointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeAccessPointsWithContext(ctx context.Context, input *efs.DescribeAccessPointsInput, opts ...request.Option) (*efs.DescribeAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeAccessPointsOutput), nil
	}
	out, err := c.EFSAPI.DescribeAccessPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeAccountPreferences(input *efs.DescribeAccountPreferencesInput) (*efs.DescribeAccountPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeAccountPreferencesOutput), nil
	}
	out, err := c.EFSAPI.DescribeAccountPreferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeAccountPreferencesWithContext(ctx context.Context, input *efs.DescribeAccountPreferencesInput, opts ...request.Option) (*efs.DescribeAccountPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeAccountPreferencesOutput), nil
	}
	out, err := c.EFSAPI.DescribeAccountPreferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeBackupPolicy(input *efs.DescribeBackupPolicyInput) (*efs.DescribeBackupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeBackupPolicyOutput), nil
	}
	out, err := c.EFSAPI.DescribeBackupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeBackupPolicyWithContext(ctx context.Context, input *efs.DescribeBackupPolicyInput, opts ...request.Option) (*efs.DescribeBackupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeBackupPolicyOutput), nil
	}
	out, err := c.EFSAPI.DescribeBackupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeFileSystemPolicy(input *efs.DescribeFileSystemPolicyInput) (*efs.DescribeFileSystemPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeFileSystemPolicyOutput), nil
	}
	out, err := c.EFSAPI.DescribeFileSystemPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeFileSystemPolicyWithContext(ctx context.Context, input *efs.DescribeFileSystemPolicyInput, opts ...request.Option) (*efs.DescribeFileSystemPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeFileSystemPolicyOutput), nil
	}
	out, err := c.EFSAPI.DescribeFileSystemPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeFileSystems(input *efs.DescribeFileSystemsInput) (*efs.DescribeFileSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeFileSystemsOutput), nil
	}
	out, err := c.EFSAPI.DescribeFileSystems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeFileSystemsPages(input *efs.DescribeFileSystemsInput, fn func(*efs.DescribeFileSystemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeFileSystemsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeFileSystemsOutput{}
	fnCacher := func(out *efs.DescribeFileSystemsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeFileSystemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeFileSystemsPagesWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, fn func(*efs.DescribeFileSystemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeFileSystemsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeFileSystemsOutput{}
	fnCacher := func(out *efs.DescribeFileSystemsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeFileSystemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeFileSystemsWithContext(ctx context.Context, input *efs.DescribeFileSystemsInput, opts ...request.Option) (*efs.DescribeFileSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeFileSystemsOutput), nil
	}
	out, err := c.EFSAPI.DescribeFileSystemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeLifecycleConfiguration(input *efs.DescribeLifecycleConfigurationInput) (*efs.DescribeLifecycleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeLifecycleConfigurationOutput), nil
	}
	out, err := c.EFSAPI.DescribeLifecycleConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeLifecycleConfigurationWithContext(ctx context.Context, input *efs.DescribeLifecycleConfigurationInput, opts ...request.Option) (*efs.DescribeLifecycleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeLifecycleConfigurationOutput), nil
	}
	out, err := c.EFSAPI.DescribeLifecycleConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeMountTargetSecurityGroups(input *efs.DescribeMountTargetSecurityGroupsInput) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeMountTargetSecurityGroupsOutput), nil
	}
	out, err := c.EFSAPI.DescribeMountTargetSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeMountTargetSecurityGroupsWithContext(ctx context.Context, input *efs.DescribeMountTargetSecurityGroupsInput, opts ...request.Option) (*efs.DescribeMountTargetSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeMountTargetSecurityGroupsOutput), nil
	}
	out, err := c.EFSAPI.DescribeMountTargetSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeMountTargets(input *efs.DescribeMountTargetsInput) (*efs.DescribeMountTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeMountTargetsOutput), nil
	}
	out, err := c.EFSAPI.DescribeMountTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeMountTargetsWithContext(ctx context.Context, input *efs.DescribeMountTargetsInput, opts ...request.Option) (*efs.DescribeMountTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeMountTargetsOutput), nil
	}
	out, err := c.EFSAPI.DescribeMountTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeReplicationConfigurations(input *efs.DescribeReplicationConfigurationsInput) (*efs.DescribeReplicationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeReplicationConfigurationsOutput), nil
	}
	out, err := c.EFSAPI.DescribeReplicationConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeReplicationConfigurationsWithContext(ctx context.Context, input *efs.DescribeReplicationConfigurationsInput, opts ...request.Option) (*efs.DescribeReplicationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeReplicationConfigurationsOutput), nil
	}
	out, err := c.EFSAPI.DescribeReplicationConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeTags(input *efs.DescribeTagsInput) (*efs.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeTagsOutput), nil
	}
	out, err := c.EFSAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) DescribeTagsPages(input *efs.DescribeTagsInput, fn func(*efs.DescribeTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeTagsOutput{}
	fnCacher := func(out *efs.DescribeTagsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeTagsPagesWithContext(ctx context.Context, input *efs.DescribeTagsInput, fn func(*efs.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &efs.DescribeTagsOutput{}
	fnCacher := func(out *efs.DescribeTagsOutput, more bool) bool {
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
	if err := c.EFSAPI.DescribeTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) DescribeTagsWithContext(ctx context.Context, input *efs.DescribeTagsInput, opts ...request.Option) (*efs.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.DescribeTagsOutput), nil
	}
	out, err := c.EFSAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) ListTagsForResource(input *efs.ListTagsForResourceInput) (*efs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.ListTagsForResourceOutput), nil
	}
	out, err := c.EFSAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EFS) ListTagsForResourcePages(input *efs.ListTagsForResourceInput, fn func(*efs.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &efs.ListTagsForResourceOutput{}
	fnCacher := func(out *efs.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.EFSAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) ListTagsForResourcePagesWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, fn func(*efs.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*efs.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &efs.ListTagsForResourceOutput{}
	fnCacher := func(out *efs.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.EFSAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EFS) ListTagsForResourceWithContext(ctx context.Context, input *efs.ListTagsForResourceInput, opts ...request.Option) (*efs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*efs.ListTagsForResourceOutput), nil
	}
	out, err := c.EFSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
