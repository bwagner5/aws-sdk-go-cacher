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
	"github.com/aws/aws-sdk-go/service/datasync"
	"github.com/aws/aws-sdk-go/service/datasync/datasynciface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DataSync struct {
	datasynciface.DataSyncAPI
	cache *cache.Cache
}

func NewDataSync(datasyncapi datasynciface.DataSyncAPI) *DataSync {
	return &DataSync{
		DataSyncAPI: datasyncapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DataSync) DescribeAgent(input *datasync.DescribeAgentInput) (*datasync.DescribeAgentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeAgentOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeAgent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeAgentWithContext(ctx context.Context, input *datasync.DescribeAgentInput, opts ...request.Option) (*datasync.DescribeAgentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeAgentOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeAgentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationEfs(input *datasync.DescribeLocationEfsInput) (*datasync.DescribeLocationEfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationEfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationEfs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationEfsWithContext(ctx context.Context, input *datasync.DescribeLocationEfsInput, opts ...request.Option) (*datasync.DescribeLocationEfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationEfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationEfsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxLustre(input *datasync.DescribeLocationFsxLustreInput) (*datasync.DescribeLocationFsxLustreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxLustreOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxLustre(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxLustreWithContext(ctx context.Context, input *datasync.DescribeLocationFsxLustreInput, opts ...request.Option) (*datasync.DescribeLocationFsxLustreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxLustreOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxLustreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxOntap(input *datasync.DescribeLocationFsxOntapInput) (*datasync.DescribeLocationFsxOntapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxOntapOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxOntap(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxOntapWithContext(ctx context.Context, input *datasync.DescribeLocationFsxOntapInput, opts ...request.Option) (*datasync.DescribeLocationFsxOntapOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxOntapOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxOntapWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxOpenZfs(input *datasync.DescribeLocationFsxOpenZfsInput) (*datasync.DescribeLocationFsxOpenZfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxOpenZfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxOpenZfs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxOpenZfsWithContext(ctx context.Context, input *datasync.DescribeLocationFsxOpenZfsInput, opts ...request.Option) (*datasync.DescribeLocationFsxOpenZfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxOpenZfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxOpenZfsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxWindows(input *datasync.DescribeLocationFsxWindowsInput) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxWindowsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxWindows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationFsxWindowsWithContext(ctx context.Context, input *datasync.DescribeLocationFsxWindowsInput, opts ...request.Option) (*datasync.DescribeLocationFsxWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationFsxWindowsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationFsxWindowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationHdfs(input *datasync.DescribeLocationHdfsInput) (*datasync.DescribeLocationHdfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationHdfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationHdfs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationHdfsWithContext(ctx context.Context, input *datasync.DescribeLocationHdfsInput, opts ...request.Option) (*datasync.DescribeLocationHdfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationHdfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationHdfsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationNfs(input *datasync.DescribeLocationNfsInput) (*datasync.DescribeLocationNfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationNfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationNfs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationNfsWithContext(ctx context.Context, input *datasync.DescribeLocationNfsInput, opts ...request.Option) (*datasync.DescribeLocationNfsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationNfsOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationNfsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationObjectStorage(input *datasync.DescribeLocationObjectStorageInput) (*datasync.DescribeLocationObjectStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationObjectStorageOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationObjectStorage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationObjectStorageWithContext(ctx context.Context, input *datasync.DescribeLocationObjectStorageInput, opts ...request.Option) (*datasync.DescribeLocationObjectStorageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationObjectStorageOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationObjectStorageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationS3(input *datasync.DescribeLocationS3Input) (*datasync.DescribeLocationS3Output, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationS3Output), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationS3(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationS3WithContext(ctx context.Context, input *datasync.DescribeLocationS3Input, opts ...request.Option) (*datasync.DescribeLocationS3Output, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationS3Output), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationS3WithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationSmb(input *datasync.DescribeLocationSmbInput) (*datasync.DescribeLocationSmbOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationSmbOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationSmb(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeLocationSmbWithContext(ctx context.Context, input *datasync.DescribeLocationSmbInput, opts ...request.Option) (*datasync.DescribeLocationSmbOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeLocationSmbOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeLocationSmbWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeTask(input *datasync.DescribeTaskInput) (*datasync.DescribeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeTaskOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeTaskExecution(input *datasync.DescribeTaskExecutionInput) (*datasync.DescribeTaskExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeTaskExecutionOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeTaskExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeTaskExecutionWithContext(ctx context.Context, input *datasync.DescribeTaskExecutionInput, opts ...request.Option) (*datasync.DescribeTaskExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeTaskExecutionOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeTaskExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) DescribeTaskWithContext(ctx context.Context, input *datasync.DescribeTaskInput, opts ...request.Option) (*datasync.DescribeTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.DescribeTaskOutput), nil
	}
	out, err := c.DataSyncAPI.DescribeTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListAgents(input *datasync.ListAgentsInput) (*datasync.ListAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListAgentsOutput), nil
	}
	out, err := c.DataSyncAPI.ListAgents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListAgentsPages(input *datasync.ListAgentsInput, fn func(*datasync.ListAgentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListAgentsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListAgentsOutput{}
	fnCacher := func(out *datasync.ListAgentsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListAgentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListAgentsPagesWithContext(ctx context.Context, input *datasync.ListAgentsInput, fn func(*datasync.ListAgentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListAgentsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListAgentsOutput{}
	fnCacher := func(out *datasync.ListAgentsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListAgentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListAgentsWithContext(ctx context.Context, input *datasync.ListAgentsInput, opts ...request.Option) (*datasync.ListAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListAgentsOutput), nil
	}
	out, err := c.DataSyncAPI.ListAgentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListLocations(input *datasync.ListLocationsInput) (*datasync.ListLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListLocationsOutput), nil
	}
	out, err := c.DataSyncAPI.ListLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListLocationsPages(input *datasync.ListLocationsInput, fn func(*datasync.ListLocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListLocationsOutput{}
	fnCacher := func(out *datasync.ListLocationsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListLocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListLocationsPagesWithContext(ctx context.Context, input *datasync.ListLocationsInput, fn func(*datasync.ListLocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListLocationsOutput{}
	fnCacher := func(out *datasync.ListLocationsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListLocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListLocationsWithContext(ctx context.Context, input *datasync.ListLocationsInput, opts ...request.Option) (*datasync.ListLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListLocationsOutput), nil
	}
	out, err := c.DataSyncAPI.ListLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTagsForResource(input *datasync.ListTagsForResourceInput) (*datasync.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTagsForResourceOutput), nil
	}
	out, err := c.DataSyncAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTagsForResourcePages(input *datasync.ListTagsForResourceInput, fn func(*datasync.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTagsForResourceOutput{}
	fnCacher := func(out *datasync.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTagsForResourcePagesWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, fn func(*datasync.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTagsForResourceOutput{}
	fnCacher := func(out *datasync.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTagsForResourceWithContext(ctx context.Context, input *datasync.ListTagsForResourceInput, opts ...request.Option) (*datasync.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTagsForResourceOutput), nil
	}
	out, err := c.DataSyncAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTaskExecutions(input *datasync.ListTaskExecutionsInput) (*datasync.ListTaskExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTaskExecutionsOutput), nil
	}
	out, err := c.DataSyncAPI.ListTaskExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTaskExecutionsPages(input *datasync.ListTaskExecutionsInput, fn func(*datasync.ListTaskExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTaskExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTaskExecutionsOutput{}
	fnCacher := func(out *datasync.ListTaskExecutionsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTaskExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTaskExecutionsPagesWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, fn func(*datasync.ListTaskExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTaskExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTaskExecutionsOutput{}
	fnCacher := func(out *datasync.ListTaskExecutionsOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTaskExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTaskExecutionsWithContext(ctx context.Context, input *datasync.ListTaskExecutionsInput, opts ...request.Option) (*datasync.ListTaskExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTaskExecutionsOutput), nil
	}
	out, err := c.DataSyncAPI.ListTaskExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTasks(input *datasync.ListTasksInput) (*datasync.ListTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTasksOutput), nil
	}
	out, err := c.DataSyncAPI.ListTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataSync) ListTasksPages(input *datasync.ListTasksInput, fn func(*datasync.ListTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTasksOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTasksOutput{}
	fnCacher := func(out *datasync.ListTasksOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTasksPagesWithContext(ctx context.Context, input *datasync.ListTasksInput, fn func(*datasync.ListTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datasync.ListTasksOutput), false)
		return nil
	}
	cachable := true
	output := &datasync.ListTasksOutput{}
	fnCacher := func(out *datasync.ListTasksOutput, more bool) bool {
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
	if err := c.DataSyncAPI.ListTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataSync) ListTasksWithContext(ctx context.Context, input *datasync.ListTasksInput, opts ...request.Option) (*datasync.ListTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datasync.ListTasksOutput), nil
	}
	out, err := c.DataSyncAPI.ListTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
