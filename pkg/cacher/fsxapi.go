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
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type FSx struct {
	fsxiface.FSxAPI
	cache *cache.Cache
}

func NewFSx(fsxapi fsxiface.FSxAPI) *FSx {
	return &FSx{
		FSxAPI: fsxapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *FSx) DescribeBackups(input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeBackupsOutput), nil
	}
	out, err := c.FSxAPI.DescribeBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeBackupsPages(input *fsx.DescribeBackupsInput, fn func(*fsx.DescribeBackupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeBackupsOutput{}
	fnCacher := func(out *fsx.DescribeBackupsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeBackupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeBackupsPagesWithContext(ctx context.Context, input *fsx.DescribeBackupsInput, fn func(*fsx.DescribeBackupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeBackupsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeBackupsOutput{}
	fnCacher := func(out *fsx.DescribeBackupsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeBackupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeBackupsWithContext(ctx context.Context, input *fsx.DescribeBackupsInput, opts ...request.Option) (*fsx.DescribeBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeBackupsOutput), nil
	}
	out, err := c.FSxAPI.DescribeBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeDataRepositoryAssociations(input *fsx.DescribeDataRepositoryAssociationsInput) (*fsx.DescribeDataRepositoryAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeDataRepositoryAssociationsOutput), nil
	}
	out, err := c.FSxAPI.DescribeDataRepositoryAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeDataRepositoryAssociationsPages(input *fsx.DescribeDataRepositoryAssociationsInput, fn func(*fsx.DescribeDataRepositoryAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeDataRepositoryAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeDataRepositoryAssociationsOutput{}
	fnCacher := func(out *fsx.DescribeDataRepositoryAssociationsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeDataRepositoryAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeDataRepositoryAssociationsPagesWithContext(ctx context.Context, input *fsx.DescribeDataRepositoryAssociationsInput, fn func(*fsx.DescribeDataRepositoryAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeDataRepositoryAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeDataRepositoryAssociationsOutput{}
	fnCacher := func(out *fsx.DescribeDataRepositoryAssociationsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeDataRepositoryAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeDataRepositoryAssociationsWithContext(ctx context.Context, input *fsx.DescribeDataRepositoryAssociationsInput, opts ...request.Option) (*fsx.DescribeDataRepositoryAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeDataRepositoryAssociationsOutput), nil
	}
	out, err := c.FSxAPI.DescribeDataRepositoryAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeDataRepositoryTasks(input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeDataRepositoryTasksOutput), nil
	}
	out, err := c.FSxAPI.DescribeDataRepositoryTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeDataRepositoryTasksPages(input *fsx.DescribeDataRepositoryTasksInput, fn func(*fsx.DescribeDataRepositoryTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeDataRepositoryTasksOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeDataRepositoryTasksOutput{}
	fnCacher := func(out *fsx.DescribeDataRepositoryTasksOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeDataRepositoryTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeDataRepositoryTasksPagesWithContext(ctx context.Context, input *fsx.DescribeDataRepositoryTasksInput, fn func(*fsx.DescribeDataRepositoryTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeDataRepositoryTasksOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeDataRepositoryTasksOutput{}
	fnCacher := func(out *fsx.DescribeDataRepositoryTasksOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeDataRepositoryTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeDataRepositoryTasksWithContext(ctx context.Context, input *fsx.DescribeDataRepositoryTasksInput, opts ...request.Option) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeDataRepositoryTasksOutput), nil
	}
	out, err := c.FSxAPI.DescribeDataRepositoryTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileCaches(input *fsx.DescribeFileCachesInput) (*fsx.DescribeFileCachesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileCachesOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileCaches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileCachesPages(input *fsx.DescribeFileCachesInput, fn func(*fsx.DescribeFileCachesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileCachesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileCachesOutput{}
	fnCacher := func(out *fsx.DescribeFileCachesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileCachesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileCachesPagesWithContext(ctx context.Context, input *fsx.DescribeFileCachesInput, fn func(*fsx.DescribeFileCachesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileCachesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileCachesOutput{}
	fnCacher := func(out *fsx.DescribeFileCachesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileCachesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileCachesWithContext(ctx context.Context, input *fsx.DescribeFileCachesInput, opts ...request.Option) (*fsx.DescribeFileCachesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileCachesOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileCachesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileSystemAliases(input *fsx.DescribeFileSystemAliasesInput) (*fsx.DescribeFileSystemAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileSystemAliasesOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileSystemAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileSystemAliasesPages(input *fsx.DescribeFileSystemAliasesInput, fn func(*fsx.DescribeFileSystemAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileSystemAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileSystemAliasesOutput{}
	fnCacher := func(out *fsx.DescribeFileSystemAliasesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileSystemAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileSystemAliasesPagesWithContext(ctx context.Context, input *fsx.DescribeFileSystemAliasesInput, fn func(*fsx.DescribeFileSystemAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileSystemAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileSystemAliasesOutput{}
	fnCacher := func(out *fsx.DescribeFileSystemAliasesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileSystemAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileSystemAliasesWithContext(ctx context.Context, input *fsx.DescribeFileSystemAliasesInput, opts ...request.Option) (*fsx.DescribeFileSystemAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileSystemAliasesOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileSystemAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileSystems(input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileSystemsOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileSystems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeFileSystemsPages(input *fsx.DescribeFileSystemsInput, fn func(*fsx.DescribeFileSystemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileSystemsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileSystemsOutput{}
	fnCacher := func(out *fsx.DescribeFileSystemsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileSystemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileSystemsPagesWithContext(ctx context.Context, input *fsx.DescribeFileSystemsInput, fn func(*fsx.DescribeFileSystemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeFileSystemsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeFileSystemsOutput{}
	fnCacher := func(out *fsx.DescribeFileSystemsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeFileSystemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeFileSystemsWithContext(ctx context.Context, input *fsx.DescribeFileSystemsInput, opts ...request.Option) (*fsx.DescribeFileSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeFileSystemsOutput), nil
	}
	out, err := c.FSxAPI.DescribeFileSystemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeSnapshots(input *fsx.DescribeSnapshotsInput) (*fsx.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeSnapshotsOutput), nil
	}
	out, err := c.FSxAPI.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeSnapshotsPages(input *fsx.DescribeSnapshotsInput, fn func(*fsx.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeSnapshotsOutput{}
	fnCacher := func(out *fsx.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeSnapshotsPagesWithContext(ctx context.Context, input *fsx.DescribeSnapshotsInput, fn func(*fsx.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeSnapshotsOutput{}
	fnCacher := func(out *fsx.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeSnapshotsWithContext(ctx context.Context, input *fsx.DescribeSnapshotsInput, opts ...request.Option) (*fsx.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeSnapshotsOutput), nil
	}
	out, err := c.FSxAPI.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeStorageVirtualMachines(input *fsx.DescribeStorageVirtualMachinesInput) (*fsx.DescribeStorageVirtualMachinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeStorageVirtualMachinesOutput), nil
	}
	out, err := c.FSxAPI.DescribeStorageVirtualMachines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeStorageVirtualMachinesPages(input *fsx.DescribeStorageVirtualMachinesInput, fn func(*fsx.DescribeStorageVirtualMachinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeStorageVirtualMachinesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeStorageVirtualMachinesOutput{}
	fnCacher := func(out *fsx.DescribeStorageVirtualMachinesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeStorageVirtualMachinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeStorageVirtualMachinesPagesWithContext(ctx context.Context, input *fsx.DescribeStorageVirtualMachinesInput, fn func(*fsx.DescribeStorageVirtualMachinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeStorageVirtualMachinesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeStorageVirtualMachinesOutput{}
	fnCacher := func(out *fsx.DescribeStorageVirtualMachinesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeStorageVirtualMachinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeStorageVirtualMachinesWithContext(ctx context.Context, input *fsx.DescribeStorageVirtualMachinesInput, opts ...request.Option) (*fsx.DescribeStorageVirtualMachinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeStorageVirtualMachinesOutput), nil
	}
	out, err := c.FSxAPI.DescribeStorageVirtualMachinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeVolumes(input *fsx.DescribeVolumesInput) (*fsx.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeVolumesOutput), nil
	}
	out, err := c.FSxAPI.DescribeVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) DescribeVolumesPages(input *fsx.DescribeVolumesInput, fn func(*fsx.DescribeVolumesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeVolumesOutput{}
	fnCacher := func(out *fsx.DescribeVolumesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeVolumesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeVolumesPagesWithContext(ctx context.Context, input *fsx.DescribeVolumesInput, fn func(*fsx.DescribeVolumesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.DescribeVolumesOutput{}
	fnCacher := func(out *fsx.DescribeVolumesOutput, more bool) bool {
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
	if err := c.FSxAPI.DescribeVolumesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) DescribeVolumesWithContext(ctx context.Context, input *fsx.DescribeVolumesInput, opts ...request.Option) (*fsx.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.DescribeVolumesOutput), nil
	}
	out, err := c.FSxAPI.DescribeVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) ListTagsForResource(input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.ListTagsForResourceOutput), nil
	}
	out, err := c.FSxAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *FSx) ListTagsForResourcePages(input *fsx.ListTagsForResourceInput, fn func(*fsx.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.ListTagsForResourceOutput{}
	fnCacher := func(out *fsx.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.FSxAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) ListTagsForResourcePagesWithContext(ctx context.Context, input *fsx.ListTagsForResourceInput, fn func(*fsx.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*fsx.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &fsx.ListTagsForResourceOutput{}
	fnCacher := func(out *fsx.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.FSxAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *FSx) ListTagsForResourceWithContext(ctx context.Context, input *fsx.ListTagsForResourceInput, opts ...request.Option) (*fsx.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*fsx.ListTagsForResourceOutput), nil
	}
	out, err := c.FSxAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
