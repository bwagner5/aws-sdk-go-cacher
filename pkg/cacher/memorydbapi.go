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
	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/aws/aws-sdk-go/service/memorydb/memorydbiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MemoryDB struct {
	memorydbiface.MemoryDBAPI
	cache *cache.Cache
}

func NewMemoryDB(memorydbapi memorydbiface.MemoryDBAPI) *MemoryDB {
	return &MemoryDB{
		MemoryDBAPI: memorydbapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MemoryDB) BatchUpdateCluster(input *memorydb.BatchUpdateClusterInput) (*memorydb.BatchUpdateClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.BatchUpdateClusterOutput), nil
	}
	out, err := c.MemoryDBAPI.BatchUpdateCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) BatchUpdateClusterWithContext(ctx context.Context, input *memorydb.BatchUpdateClusterInput, opts ...request.Option) (*memorydb.BatchUpdateClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.BatchUpdateClusterOutput), nil
	}
	out, err := c.MemoryDBAPI.BatchUpdateClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeACLs(input *memorydb.DescribeACLsInput) (*memorydb.DescribeACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeACLsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeACLs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeACLsPages(input *memorydb.DescribeACLsInput, fn func(*memorydb.DescribeACLsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeACLsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeACLsOutput{}
	fnCacher := func(out *memorydb.DescribeACLsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeACLsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeACLsPagesWithContext(ctx context.Context, input *memorydb.DescribeACLsInput, fn func(*memorydb.DescribeACLsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeACLsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeACLsOutput{}
	fnCacher := func(out *memorydb.DescribeACLsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeACLsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeACLsWithContext(ctx context.Context, input *memorydb.DescribeACLsInput, opts ...request.Option) (*memorydb.DescribeACLsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeACLsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeACLsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeClusters(input *memorydb.DescribeClustersInput) (*memorydb.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeClustersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeClustersPages(input *memorydb.DescribeClustersInput, fn func(*memorydb.DescribeClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeClustersOutput{}
	fnCacher := func(out *memorydb.DescribeClustersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeClustersPagesWithContext(ctx context.Context, input *memorydb.DescribeClustersInput, fn func(*memorydb.DescribeClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeClustersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeClustersOutput{}
	fnCacher := func(out *memorydb.DescribeClustersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeClustersWithContext(ctx context.Context, input *memorydb.DescribeClustersInput, opts ...request.Option) (*memorydb.DescribeClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeClustersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeEngineVersions(input *memorydb.DescribeEngineVersionsInput) (*memorydb.DescribeEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeEngineVersionsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeEngineVersionsPages(input *memorydb.DescribeEngineVersionsInput, fn func(*memorydb.DescribeEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeEngineVersionsOutput{}
	fnCacher := func(out *memorydb.DescribeEngineVersionsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeEngineVersionsPagesWithContext(ctx context.Context, input *memorydb.DescribeEngineVersionsInput, fn func(*memorydb.DescribeEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeEngineVersionsOutput{}
	fnCacher := func(out *memorydb.DescribeEngineVersionsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeEngineVersionsWithContext(ctx context.Context, input *memorydb.DescribeEngineVersionsInput, opts ...request.Option) (*memorydb.DescribeEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeEngineVersionsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeEvents(input *memorydb.DescribeEventsInput) (*memorydb.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeEventsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeEventsPages(input *memorydb.DescribeEventsInput, fn func(*memorydb.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeEventsOutput{}
	fnCacher := func(out *memorydb.DescribeEventsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeEventsPagesWithContext(ctx context.Context, input *memorydb.DescribeEventsInput, fn func(*memorydb.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeEventsOutput{}
	fnCacher := func(out *memorydb.DescribeEventsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeEventsWithContext(ctx context.Context, input *memorydb.DescribeEventsInput, opts ...request.Option) (*memorydb.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeEventsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeParameterGroups(input *memorydb.DescribeParameterGroupsInput) (*memorydb.DescribeParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeParameterGroupsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeParameterGroupsPages(input *memorydb.DescribeParameterGroupsInput, fn func(*memorydb.DescribeParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeParameterGroupsOutput{}
	fnCacher := func(out *memorydb.DescribeParameterGroupsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeParameterGroupsPagesWithContext(ctx context.Context, input *memorydb.DescribeParameterGroupsInput, fn func(*memorydb.DescribeParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeParameterGroupsOutput{}
	fnCacher := func(out *memorydb.DescribeParameterGroupsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeParameterGroupsWithContext(ctx context.Context, input *memorydb.DescribeParameterGroupsInput, opts ...request.Option) (*memorydb.DescribeParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeParameterGroupsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeParameters(input *memorydb.DescribeParametersInput) (*memorydb.DescribeParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeParametersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeParametersPages(input *memorydb.DescribeParametersInput, fn func(*memorydb.DescribeParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeParametersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeParametersOutput{}
	fnCacher := func(out *memorydb.DescribeParametersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeParametersPagesWithContext(ctx context.Context, input *memorydb.DescribeParametersInput, fn func(*memorydb.DescribeParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeParametersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeParametersOutput{}
	fnCacher := func(out *memorydb.DescribeParametersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeParametersWithContext(ctx context.Context, input *memorydb.DescribeParametersInput, opts ...request.Option) (*memorydb.DescribeParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeParametersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeReservedNodes(input *memorydb.DescribeReservedNodesInput) (*memorydb.DescribeReservedNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeReservedNodesOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeReservedNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeReservedNodesOfferings(input *memorydb.DescribeReservedNodesOfferingsInput) (*memorydb.DescribeReservedNodesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeReservedNodesOfferingsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeReservedNodesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeReservedNodesOfferingsPages(input *memorydb.DescribeReservedNodesOfferingsInput, fn func(*memorydb.DescribeReservedNodesOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeReservedNodesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeReservedNodesOfferingsOutput{}
	fnCacher := func(out *memorydb.DescribeReservedNodesOfferingsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeReservedNodesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeReservedNodesOfferingsPagesWithContext(ctx context.Context, input *memorydb.DescribeReservedNodesOfferingsInput, fn func(*memorydb.DescribeReservedNodesOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeReservedNodesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeReservedNodesOfferingsOutput{}
	fnCacher := func(out *memorydb.DescribeReservedNodesOfferingsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeReservedNodesOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeReservedNodesOfferingsWithContext(ctx context.Context, input *memorydb.DescribeReservedNodesOfferingsInput, opts ...request.Option) (*memorydb.DescribeReservedNodesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeReservedNodesOfferingsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeReservedNodesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeReservedNodesPages(input *memorydb.DescribeReservedNodesInput, fn func(*memorydb.DescribeReservedNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeReservedNodesOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeReservedNodesOutput{}
	fnCacher := func(out *memorydb.DescribeReservedNodesOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeReservedNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeReservedNodesPagesWithContext(ctx context.Context, input *memorydb.DescribeReservedNodesInput, fn func(*memorydb.DescribeReservedNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeReservedNodesOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeReservedNodesOutput{}
	fnCacher := func(out *memorydb.DescribeReservedNodesOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeReservedNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeReservedNodesWithContext(ctx context.Context, input *memorydb.DescribeReservedNodesInput, opts ...request.Option) (*memorydb.DescribeReservedNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeReservedNodesOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeReservedNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeServiceUpdates(input *memorydb.DescribeServiceUpdatesInput) (*memorydb.DescribeServiceUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeServiceUpdatesOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeServiceUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeServiceUpdatesPages(input *memorydb.DescribeServiceUpdatesInput, fn func(*memorydb.DescribeServiceUpdatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeServiceUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeServiceUpdatesOutput{}
	fnCacher := func(out *memorydb.DescribeServiceUpdatesOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeServiceUpdatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeServiceUpdatesPagesWithContext(ctx context.Context, input *memorydb.DescribeServiceUpdatesInput, fn func(*memorydb.DescribeServiceUpdatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeServiceUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeServiceUpdatesOutput{}
	fnCacher := func(out *memorydb.DescribeServiceUpdatesOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeServiceUpdatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeServiceUpdatesWithContext(ctx context.Context, input *memorydb.DescribeServiceUpdatesInput, opts ...request.Option) (*memorydb.DescribeServiceUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeServiceUpdatesOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeServiceUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeSnapshots(input *memorydb.DescribeSnapshotsInput) (*memorydb.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeSnapshotsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeSnapshotsPages(input *memorydb.DescribeSnapshotsInput, fn func(*memorydb.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeSnapshotsOutput{}
	fnCacher := func(out *memorydb.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeSnapshotsPagesWithContext(ctx context.Context, input *memorydb.DescribeSnapshotsInput, fn func(*memorydb.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeSnapshotsOutput{}
	fnCacher := func(out *memorydb.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeSnapshotsWithContext(ctx context.Context, input *memorydb.DescribeSnapshotsInput, opts ...request.Option) (*memorydb.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeSnapshotsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeSubnetGroups(input *memorydb.DescribeSubnetGroupsInput) (*memorydb.DescribeSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeSubnetGroupsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeSubnetGroupsPages(input *memorydb.DescribeSubnetGroupsInput, fn func(*memorydb.DescribeSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeSubnetGroupsOutput{}
	fnCacher := func(out *memorydb.DescribeSubnetGroupsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeSubnetGroupsPagesWithContext(ctx context.Context, input *memorydb.DescribeSubnetGroupsInput, fn func(*memorydb.DescribeSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeSubnetGroupsOutput{}
	fnCacher := func(out *memorydb.DescribeSubnetGroupsOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeSubnetGroupsWithContext(ctx context.Context, input *memorydb.DescribeSubnetGroupsInput, opts ...request.Option) (*memorydb.DescribeSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeSubnetGroupsOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeUsers(input *memorydb.DescribeUsersInput) (*memorydb.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeUsersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) DescribeUsersPages(input *memorydb.DescribeUsersInput, fn func(*memorydb.DescribeUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeUsersOutput{}
	fnCacher := func(out *memorydb.DescribeUsersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeUsersPagesWithContext(ctx context.Context, input *memorydb.DescribeUsersInput, fn func(*memorydb.DescribeUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*memorydb.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &memorydb.DescribeUsersOutput{}
	fnCacher := func(out *memorydb.DescribeUsersOutput, more bool) bool {
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
	if err := c.MemoryDBAPI.DescribeUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MemoryDB) DescribeUsersWithContext(ctx context.Context, input *memorydb.DescribeUsersInput, opts ...request.Option) (*memorydb.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.DescribeUsersOutput), nil
	}
	out, err := c.MemoryDBAPI.DescribeUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) ListAllowedNodeTypeUpdates(input *memorydb.ListAllowedNodeTypeUpdatesInput) (*memorydb.ListAllowedNodeTypeUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.ListAllowedNodeTypeUpdatesOutput), nil
	}
	out, err := c.MemoryDBAPI.ListAllowedNodeTypeUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) ListAllowedNodeTypeUpdatesWithContext(ctx context.Context, input *memorydb.ListAllowedNodeTypeUpdatesInput, opts ...request.Option) (*memorydb.ListAllowedNodeTypeUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.ListAllowedNodeTypeUpdatesOutput), nil
	}
	out, err := c.MemoryDBAPI.ListAllowedNodeTypeUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) ListTags(input *memorydb.ListTagsInput) (*memorydb.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.ListTagsOutput), nil
	}
	out, err := c.MemoryDBAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MemoryDB) ListTagsWithContext(ctx context.Context, input *memorydb.ListTagsInput, opts ...request.Option) (*memorydb.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*memorydb.ListTagsOutput), nil
	}
	out, err := c.MemoryDBAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
