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
	"github.com/aws/aws-sdk-go/service/neptune"
	"github.com/aws/aws-sdk-go/service/neptune/neptuneiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Neptune struct {
	neptuneiface.NeptuneAPI
	cache *cache.Cache
}

func NewNeptune(neptuneapi neptuneiface.NeptuneAPI) *Neptune {
	return &Neptune{
		NeptuneAPI: neptuneapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Neptune) DescribeDBClusterEndpoints(input *neptune.DescribeDBClusterEndpointsInput) (*neptune.DescribeDBClusterEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterEndpointsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterEndpointsPages(input *neptune.DescribeDBClusterEndpointsInput, fn func(*neptune.DescribeDBClusterEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterEndpointsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterEndpointsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterEndpointsPagesWithContext(ctx context.Context, input *neptune.DescribeDBClusterEndpointsInput, fn func(*neptune.DescribeDBClusterEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterEndpointsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterEndpointsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterEndpointsWithContext(ctx context.Context, input *neptune.DescribeDBClusterEndpointsInput, opts ...request.Option) (*neptune.DescribeDBClusterEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterEndpointsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterParameterGroups(input *neptune.DescribeDBClusterParameterGroupsInput) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterParameterGroupsPages(input *neptune.DescribeDBClusterParameterGroupsInput, fn func(*neptune.DescribeDBClusterParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterParameterGroupsPagesWithContext(ctx context.Context, input *neptune.DescribeDBClusterParameterGroupsInput, fn func(*neptune.DescribeDBClusterParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterParameterGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterParameterGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterParameterGroupsWithContext(ctx context.Context, input *neptune.DescribeDBClusterParameterGroupsInput, opts ...request.Option) (*neptune.DescribeDBClusterParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterParameterGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterParameters(input *neptune.DescribeDBClusterParametersInput) (*neptune.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterParametersPages(input *neptune.DescribeDBClusterParametersInput, fn func(*neptune.DescribeDBClusterParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterParametersPagesWithContext(ctx context.Context, input *neptune.DescribeDBClusterParametersInput, fn func(*neptune.DescribeDBClusterParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterParametersOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterParametersWithContext(ctx context.Context, input *neptune.DescribeDBClusterParametersInput, opts ...request.Option) (*neptune.DescribeDBClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterSnapshotAttributes(input *neptune.DescribeDBClusterSnapshotAttributesInput) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterSnapshotAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterSnapshotAttributesWithContext(ctx context.Context, input *neptune.DescribeDBClusterSnapshotAttributesInput, opts ...request.Option) (*neptune.DescribeDBClusterSnapshotAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterSnapshotAttributesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterSnapshotAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterSnapshots(input *neptune.DescribeDBClusterSnapshotsInput) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusterSnapshotsPages(input *neptune.DescribeDBClusterSnapshotsInput, fn func(*neptune.DescribeDBClusterSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterSnapshotsPagesWithContext(ctx context.Context, input *neptune.DescribeDBClusterSnapshotsInput, fn func(*neptune.DescribeDBClusterSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClusterSnapshotsOutput{}
	fnCacher := func(out *neptune.DescribeDBClusterSnapshotsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClusterSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClusterSnapshotsWithContext(ctx context.Context, input *neptune.DescribeDBClusterSnapshotsInput, opts ...request.Option) (*neptune.DescribeDBClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClusterSnapshotsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClusters(input *neptune.DescribeDBClustersInput) (*neptune.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClustersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBClustersPages(input *neptune.DescribeDBClustersInput, fn func(*neptune.DescribeDBClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClustersOutput{}
	fnCacher := func(out *neptune.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClustersPagesWithContext(ctx context.Context, input *neptune.DescribeDBClustersInput, fn func(*neptune.DescribeDBClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBClustersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBClustersOutput{}
	fnCacher := func(out *neptune.DescribeDBClustersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBClustersWithContext(ctx context.Context, input *neptune.DescribeDBClustersInput, opts ...request.Option) (*neptune.DescribeDBClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBClustersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBEngineVersions(input *neptune.DescribeDBEngineVersionsInput) (*neptune.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBEngineVersionsPages(input *neptune.DescribeDBEngineVersionsInput, fn func(*neptune.DescribeDBEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *neptune.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBEngineVersionsPagesWithContext(ctx context.Context, input *neptune.DescribeDBEngineVersionsInput, fn func(*neptune.DescribeDBEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBEngineVersionsOutput{}
	fnCacher := func(out *neptune.DescribeDBEngineVersionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBEngineVersionsWithContext(ctx context.Context, input *neptune.DescribeDBEngineVersionsInput, opts ...request.Option) (*neptune.DescribeDBEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBEngineVersionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBInstances(input *neptune.DescribeDBInstancesInput) (*neptune.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBInstancesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBInstancesPages(input *neptune.DescribeDBInstancesInput, fn func(*neptune.DescribeDBInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBInstancesOutput{}
	fnCacher := func(out *neptune.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBInstancesPagesWithContext(ctx context.Context, input *neptune.DescribeDBInstancesInput, fn func(*neptune.DescribeDBInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBInstancesOutput{}
	fnCacher := func(out *neptune.DescribeDBInstancesOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBInstancesWithContext(ctx context.Context, input *neptune.DescribeDBInstancesInput, opts ...request.Option) (*neptune.DescribeDBInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBInstancesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBParameterGroups(input *neptune.DescribeDBParameterGroupsInput) (*neptune.DescribeDBParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBParameterGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBParameterGroupsPages(input *neptune.DescribeDBParameterGroupsInput, fn func(*neptune.DescribeDBParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBParameterGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBParameterGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBParameterGroupsPagesWithContext(ctx context.Context, input *neptune.DescribeDBParameterGroupsInput, fn func(*neptune.DescribeDBParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBParameterGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBParameterGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBParameterGroupsWithContext(ctx context.Context, input *neptune.DescribeDBParameterGroupsInput, opts ...request.Option) (*neptune.DescribeDBParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBParameterGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBParameters(input *neptune.DescribeDBParametersInput) (*neptune.DescribeDBParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBParametersPages(input *neptune.DescribeDBParametersInput, fn func(*neptune.DescribeDBParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBParametersOutput{}
	fnCacher := func(out *neptune.DescribeDBParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBParametersPagesWithContext(ctx context.Context, input *neptune.DescribeDBParametersInput, fn func(*neptune.DescribeDBParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBParametersOutput{}
	fnCacher := func(out *neptune.DescribeDBParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBParametersWithContext(ctx context.Context, input *neptune.DescribeDBParametersInput, opts ...request.Option) (*neptune.DescribeDBParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBSubnetGroups(input *neptune.DescribeDBSubnetGroupsInput) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeDBSubnetGroupsPages(input *neptune.DescribeDBSubnetGroupsInput, fn func(*neptune.DescribeDBSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBSubnetGroupsPagesWithContext(ctx context.Context, input *neptune.DescribeDBSubnetGroupsInput, fn func(*neptune.DescribeDBSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeDBSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeDBSubnetGroupsOutput{}
	fnCacher := func(out *neptune.DescribeDBSubnetGroupsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeDBSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeDBSubnetGroupsWithContext(ctx context.Context, input *neptune.DescribeDBSubnetGroupsInput, opts ...request.Option) (*neptune.DescribeDBSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeDBSubnetGroupsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeDBSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEngineDefaultClusterParameters(input *neptune.DescribeEngineDefaultClusterParametersInput) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEngineDefaultClusterParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEngineDefaultClusterParametersWithContext(ctx context.Context, input *neptune.DescribeEngineDefaultClusterParametersInput, opts ...request.Option) (*neptune.DescribeEngineDefaultClusterParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEngineDefaultClusterParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEngineDefaultClusterParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEngineDefaultParameters(input *neptune.DescribeEngineDefaultParametersInput) (*neptune.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEngineDefaultParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEngineDefaultParametersPages(input *neptune.DescribeEngineDefaultParametersInput, fn func(*neptune.DescribeEngineDefaultParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *neptune.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEngineDefaultParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEngineDefaultParametersPagesWithContext(ctx context.Context, input *neptune.DescribeEngineDefaultParametersInput, fn func(*neptune.DescribeEngineDefaultParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *neptune.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEngineDefaultParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEngineDefaultParametersWithContext(ctx context.Context, input *neptune.DescribeEngineDefaultParametersInput, opts ...request.Option) (*neptune.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEngineDefaultParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEventCategories(input *neptune.DescribeEventCategoriesInput) (*neptune.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventCategoriesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEventCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEventCategoriesWithContext(ctx context.Context, input *neptune.DescribeEventCategoriesInput, opts ...request.Option) (*neptune.DescribeEventCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventCategoriesOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEventCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEventSubscriptions(input *neptune.DescribeEventSubscriptionsInput) (*neptune.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEventSubscriptionsPages(input *neptune.DescribeEventSubscriptionsInput, fn func(*neptune.DescribeEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *neptune.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEventSubscriptionsPagesWithContext(ctx context.Context, input *neptune.DescribeEventSubscriptionsInput, fn func(*neptune.DescribeEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEventSubscriptionsOutput{}
	fnCacher := func(out *neptune.DescribeEventSubscriptionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEventSubscriptionsWithContext(ctx context.Context, input *neptune.DescribeEventSubscriptionsInput, opts ...request.Option) (*neptune.DescribeEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventSubscriptionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEvents(input *neptune.DescribeEventsInput) (*neptune.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeEventsPages(input *neptune.DescribeEventsInput, fn func(*neptune.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEventsOutput{}
	fnCacher := func(out *neptune.DescribeEventsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEventsPagesWithContext(ctx context.Context, input *neptune.DescribeEventsInput, fn func(*neptune.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeEventsOutput{}
	fnCacher := func(out *neptune.DescribeEventsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeEventsWithContext(ctx context.Context, input *neptune.DescribeEventsInput, opts ...request.Option) (*neptune.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeEventsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeGlobalClusters(input *neptune.DescribeGlobalClustersInput) (*neptune.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeGlobalClustersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeGlobalClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeGlobalClustersPages(input *neptune.DescribeGlobalClustersInput, fn func(*neptune.DescribeGlobalClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeGlobalClustersOutput{}
	fnCacher := func(out *neptune.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeGlobalClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeGlobalClustersPagesWithContext(ctx context.Context, input *neptune.DescribeGlobalClustersInput, fn func(*neptune.DescribeGlobalClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeGlobalClustersOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeGlobalClustersOutput{}
	fnCacher := func(out *neptune.DescribeGlobalClustersOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeGlobalClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeGlobalClustersWithContext(ctx context.Context, input *neptune.DescribeGlobalClustersInput, opts ...request.Option) (*neptune.DescribeGlobalClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeGlobalClustersOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeGlobalClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeOrderableDBInstanceOptions(input *neptune.DescribeOrderableDBInstanceOptionsInput) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeOrderableDBInstanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeOrderableDBInstanceOptionsPages(input *neptune.DescribeOrderableDBInstanceOptionsInput, fn func(*neptune.DescribeOrderableDBInstanceOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *neptune.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeOrderableDBInstanceOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeOrderableDBInstanceOptionsPagesWithContext(ctx context.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput, fn func(*neptune.DescribeOrderableDBInstanceOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribeOrderableDBInstanceOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribeOrderableDBInstanceOptionsOutput{}
	fnCacher := func(out *neptune.DescribeOrderableDBInstanceOptionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribeOrderableDBInstanceOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribeOrderableDBInstanceOptionsWithContext(ctx context.Context, input *neptune.DescribeOrderableDBInstanceOptionsInput, opts ...request.Option) (*neptune.DescribeOrderableDBInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeOrderableDBInstanceOptionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeOrderableDBInstanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribePendingMaintenanceActions(input *neptune.DescribePendingMaintenanceActionsInput) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribePendingMaintenanceActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribePendingMaintenanceActionsPages(input *neptune.DescribePendingMaintenanceActionsInput, fn func(*neptune.DescribePendingMaintenanceActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *neptune.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribePendingMaintenanceActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribePendingMaintenanceActionsPagesWithContext(ctx context.Context, input *neptune.DescribePendingMaintenanceActionsInput, fn func(*neptune.DescribePendingMaintenanceActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*neptune.DescribePendingMaintenanceActionsOutput), false)
		return nil
	}
	cachable := true
	output := &neptune.DescribePendingMaintenanceActionsOutput{}
	fnCacher := func(out *neptune.DescribePendingMaintenanceActionsOutput, more bool) bool {
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
	if err := c.NeptuneAPI.DescribePendingMaintenanceActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Neptune) DescribePendingMaintenanceActionsWithContext(ctx context.Context, input *neptune.DescribePendingMaintenanceActionsInput, opts ...request.Option) (*neptune.DescribePendingMaintenanceActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribePendingMaintenanceActionsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribePendingMaintenanceActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeValidDBInstanceModifications(input *neptune.DescribeValidDBInstanceModificationsInput) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeValidDBInstanceModificationsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeValidDBInstanceModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) DescribeValidDBInstanceModificationsWithContext(ctx context.Context, input *neptune.DescribeValidDBInstanceModificationsInput, opts ...request.Option) (*neptune.DescribeValidDBInstanceModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.DescribeValidDBInstanceModificationsOutput), nil
	}
	out, err := c.NeptuneAPI.DescribeValidDBInstanceModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) ListTagsForResource(input *neptune.ListTagsForResourceInput) (*neptune.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.ListTagsForResourceOutput), nil
	}
	out, err := c.NeptuneAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Neptune) ListTagsForResourceWithContext(ctx context.Context, input *neptune.ListTagsForResourceInput, opts ...request.Option) (*neptune.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*neptune.ListTagsForResourceOutput), nil
	}
	out, err := c.NeptuneAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
