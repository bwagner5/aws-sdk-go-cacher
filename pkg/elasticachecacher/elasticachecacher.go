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
package elasticachecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/elasticache"
	"github.com/aws/aws-sdk-go/service/elasticache/elasticacheiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ElastiCache struct {
	elasticacheiface.ElastiCacheAPI
	cache *cache.Cache
}

func New(elasticacheapi elasticacheiface.ElastiCacheAPI) *ElastiCache {
	return &ElastiCache{
		ElastiCacheAPI: elasticacheapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ElastiCache) BatchApplyUpdateAction(input *elasticache.BatchApplyUpdateActionInput) (*elasticache.BatchApplyUpdateActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.BatchApplyUpdateActionOutput), nil
	}
	out, err := c.ElastiCacheAPI.BatchApplyUpdateAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) BatchApplyUpdateActionWithContext(ctx context.Context, input *elasticache.BatchApplyUpdateActionInput, opts ...request.Option) (*elasticache.BatchApplyUpdateActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.BatchApplyUpdateActionOutput), nil
	}
	out, err := c.ElastiCacheAPI.BatchApplyUpdateActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) BatchStopUpdateAction(input *elasticache.BatchStopUpdateActionInput) (*elasticache.BatchStopUpdateActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.BatchStopUpdateActionOutput), nil
	}
	out, err := c.ElastiCacheAPI.BatchStopUpdateAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) BatchStopUpdateActionWithContext(ctx context.Context, input *elasticache.BatchStopUpdateActionInput, opts ...request.Option) (*elasticache.BatchStopUpdateActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.BatchStopUpdateActionOutput), nil
	}
	out, err := c.ElastiCacheAPI.BatchStopUpdateActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheClusters(input *elasticache.DescribeCacheClustersInput) (*elasticache.DescribeCacheClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheClustersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheClustersPages(input *elasticache.DescribeCacheClustersInput, fn func(*elasticache.DescribeCacheClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheClustersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheClustersOutput{}
	fnCacher := func(out *elasticache.DescribeCacheClustersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheClustersPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheClustersInput, fn func(*elasticache.DescribeCacheClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheClustersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheClustersOutput{}
	fnCacher := func(out *elasticache.DescribeCacheClustersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheClustersWithContext(ctx context.Context, input *elasticache.DescribeCacheClustersInput, opts ...request.Option) (*elasticache.DescribeCacheClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheClustersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheEngineVersions(input *elasticache.DescribeCacheEngineVersionsInput) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheEngineVersionsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheEngineVersionsPages(input *elasticache.DescribeCacheEngineVersionsInput, fn func(*elasticache.DescribeCacheEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheEngineVersionsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheEngineVersionsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheEngineVersionsPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheEngineVersionsInput, fn func(*elasticache.DescribeCacheEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheEngineVersionsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheEngineVersionsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheEngineVersionsWithContext(ctx context.Context, input *elasticache.DescribeCacheEngineVersionsInput, opts ...request.Option) (*elasticache.DescribeCacheEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheEngineVersionsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheParameterGroups(input *elasticache.DescribeCacheParameterGroupsInput) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheParameterGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheParameterGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheParameterGroupsPages(input *elasticache.DescribeCacheParameterGroupsInput, fn func(*elasticache.DescribeCacheParameterGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheParameterGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheParameterGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheParameterGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheParameterGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheParameterGroupsInput, fn func(*elasticache.DescribeCacheParameterGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheParameterGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheParameterGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheParameterGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheParameterGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheParameterGroupsWithContext(ctx context.Context, input *elasticache.DescribeCacheParameterGroupsInput, opts ...request.Option) (*elasticache.DescribeCacheParameterGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheParameterGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheParameterGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheParameters(input *elasticache.DescribeCacheParametersInput) (*elasticache.DescribeCacheParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheParametersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheParametersPages(input *elasticache.DescribeCacheParametersInput, fn func(*elasticache.DescribeCacheParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheParametersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheParametersOutput{}
	fnCacher := func(out *elasticache.DescribeCacheParametersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheParametersPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheParametersInput, fn func(*elasticache.DescribeCacheParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheParametersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheParametersOutput{}
	fnCacher := func(out *elasticache.DescribeCacheParametersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheParametersWithContext(ctx context.Context, input *elasticache.DescribeCacheParametersInput, opts ...request.Option) (*elasticache.DescribeCacheParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheParametersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheSecurityGroups(input *elasticache.DescribeCacheSecurityGroupsInput) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheSecurityGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheSecurityGroupsPages(input *elasticache.DescribeCacheSecurityGroupsInput, fn func(*elasticache.DescribeCacheSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheSecurityGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheSecurityGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheSecurityGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheSecurityGroupsInput, fn func(*elasticache.DescribeCacheSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheSecurityGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheSecurityGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheSecurityGroupsWithContext(ctx context.Context, input *elasticache.DescribeCacheSecurityGroupsInput, opts ...request.Option) (*elasticache.DescribeCacheSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheSecurityGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheSubnetGroups(input *elasticache.DescribeCacheSubnetGroupsInput) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheSubnetGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheSubnetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeCacheSubnetGroupsPages(input *elasticache.DescribeCacheSubnetGroupsInput, fn func(*elasticache.DescribeCacheSubnetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheSubnetGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheSubnetGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheSubnetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheSubnetGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeCacheSubnetGroupsInput, fn func(*elasticache.DescribeCacheSubnetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeCacheSubnetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeCacheSubnetGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeCacheSubnetGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeCacheSubnetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeCacheSubnetGroupsWithContext(ctx context.Context, input *elasticache.DescribeCacheSubnetGroupsInput, opts ...request.Option) (*elasticache.DescribeCacheSubnetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeCacheSubnetGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeCacheSubnetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeEngineDefaultParameters(input *elasticache.DescribeEngineDefaultParametersInput) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeEngineDefaultParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeEngineDefaultParametersPages(input *elasticache.DescribeEngineDefaultParametersInput, fn func(*elasticache.DescribeEngineDefaultParametersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *elasticache.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeEngineDefaultParametersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeEngineDefaultParametersPagesWithContext(ctx context.Context, input *elasticache.DescribeEngineDefaultParametersInput, fn func(*elasticache.DescribeEngineDefaultParametersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeEngineDefaultParametersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeEngineDefaultParametersOutput{}
	fnCacher := func(out *elasticache.DescribeEngineDefaultParametersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeEngineDefaultParametersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeEngineDefaultParametersWithContext(ctx context.Context, input *elasticache.DescribeEngineDefaultParametersInput, opts ...request.Option) (*elasticache.DescribeEngineDefaultParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeEngineDefaultParametersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeEngineDefaultParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeEvents(input *elasticache.DescribeEventsInput) (*elasticache.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeEventsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeEventsPages(input *elasticache.DescribeEventsInput, fn func(*elasticache.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeEventsOutput{}
	fnCacher := func(out *elasticache.DescribeEventsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeEventsPagesWithContext(ctx context.Context, input *elasticache.DescribeEventsInput, fn func(*elasticache.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeEventsOutput{}
	fnCacher := func(out *elasticache.DescribeEventsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeEventsWithContext(ctx context.Context, input *elasticache.DescribeEventsInput, opts ...request.Option) (*elasticache.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeEventsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeGlobalReplicationGroups(input *elasticache.DescribeGlobalReplicationGroupsInput) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeGlobalReplicationGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeGlobalReplicationGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeGlobalReplicationGroupsPages(input *elasticache.DescribeGlobalReplicationGroupsInput, fn func(*elasticache.DescribeGlobalReplicationGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeGlobalReplicationGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeGlobalReplicationGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeGlobalReplicationGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeGlobalReplicationGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeGlobalReplicationGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeGlobalReplicationGroupsInput, fn func(*elasticache.DescribeGlobalReplicationGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeGlobalReplicationGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeGlobalReplicationGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeGlobalReplicationGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeGlobalReplicationGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeGlobalReplicationGroupsWithContext(ctx context.Context, input *elasticache.DescribeGlobalReplicationGroupsInput, opts ...request.Option) (*elasticache.DescribeGlobalReplicationGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeGlobalReplicationGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeGlobalReplicationGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReplicationGroups(input *elasticache.DescribeReplicationGroupsInput) (*elasticache.DescribeReplicationGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReplicationGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReplicationGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReplicationGroupsPages(input *elasticache.DescribeReplicationGroupsInput, fn func(*elasticache.DescribeReplicationGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReplicationGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReplicationGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeReplicationGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReplicationGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReplicationGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput, fn func(*elasticache.DescribeReplicationGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReplicationGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReplicationGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeReplicationGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReplicationGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReplicationGroupsWithContext(ctx context.Context, input *elasticache.DescribeReplicationGroupsInput, opts ...request.Option) (*elasticache.DescribeReplicationGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReplicationGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReplicationGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReservedCacheNodes(input *elasticache.DescribeReservedCacheNodesInput) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReservedCacheNodesOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReservedCacheNodes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReservedCacheNodesOfferings(input *elasticache.DescribeReservedCacheNodesOfferingsInput) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReservedCacheNodesOfferingsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReservedCacheNodesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReservedCacheNodesOfferingsPages(input *elasticache.DescribeReservedCacheNodesOfferingsInput, fn func(*elasticache.DescribeReservedCacheNodesOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReservedCacheNodesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReservedCacheNodesOfferingsOutput{}
	fnCacher := func(out *elasticache.DescribeReservedCacheNodesOfferingsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReservedCacheNodesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReservedCacheNodesOfferingsPagesWithContext(ctx context.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput, fn func(*elasticache.DescribeReservedCacheNodesOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReservedCacheNodesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReservedCacheNodesOfferingsOutput{}
	fnCacher := func(out *elasticache.DescribeReservedCacheNodesOfferingsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReservedCacheNodesOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReservedCacheNodesOfferingsWithContext(ctx context.Context, input *elasticache.DescribeReservedCacheNodesOfferingsInput, opts ...request.Option) (*elasticache.DescribeReservedCacheNodesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReservedCacheNodesOfferingsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReservedCacheNodesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeReservedCacheNodesPages(input *elasticache.DescribeReservedCacheNodesInput, fn func(*elasticache.DescribeReservedCacheNodesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReservedCacheNodesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReservedCacheNodesOutput{}
	fnCacher := func(out *elasticache.DescribeReservedCacheNodesOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReservedCacheNodesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReservedCacheNodesPagesWithContext(ctx context.Context, input *elasticache.DescribeReservedCacheNodesInput, fn func(*elasticache.DescribeReservedCacheNodesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeReservedCacheNodesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeReservedCacheNodesOutput{}
	fnCacher := func(out *elasticache.DescribeReservedCacheNodesOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeReservedCacheNodesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeReservedCacheNodesWithContext(ctx context.Context, input *elasticache.DescribeReservedCacheNodesInput, opts ...request.Option) (*elasticache.DescribeReservedCacheNodesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeReservedCacheNodesOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeReservedCacheNodesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeServiceUpdates(input *elasticache.DescribeServiceUpdatesInput) (*elasticache.DescribeServiceUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeServiceUpdatesOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeServiceUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeServiceUpdatesPages(input *elasticache.DescribeServiceUpdatesInput, fn func(*elasticache.DescribeServiceUpdatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeServiceUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeServiceUpdatesOutput{}
	fnCacher := func(out *elasticache.DescribeServiceUpdatesOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeServiceUpdatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeServiceUpdatesPagesWithContext(ctx context.Context, input *elasticache.DescribeServiceUpdatesInput, fn func(*elasticache.DescribeServiceUpdatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeServiceUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeServiceUpdatesOutput{}
	fnCacher := func(out *elasticache.DescribeServiceUpdatesOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeServiceUpdatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeServiceUpdatesWithContext(ctx context.Context, input *elasticache.DescribeServiceUpdatesInput, opts ...request.Option) (*elasticache.DescribeServiceUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeServiceUpdatesOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeServiceUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeSnapshots(input *elasticache.DescribeSnapshotsInput) (*elasticache.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeSnapshotsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeSnapshotsPages(input *elasticache.DescribeSnapshotsInput, fn func(*elasticache.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeSnapshotsOutput{}
	fnCacher := func(out *elasticache.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeSnapshotsPagesWithContext(ctx context.Context, input *elasticache.DescribeSnapshotsInput, fn func(*elasticache.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeSnapshotsOutput{}
	fnCacher := func(out *elasticache.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeSnapshotsWithContext(ctx context.Context, input *elasticache.DescribeSnapshotsInput, opts ...request.Option) (*elasticache.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeSnapshotsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUpdateActions(input *elasticache.DescribeUpdateActionsInput) (*elasticache.DescribeUpdateActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUpdateActionsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUpdateActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUpdateActionsPages(input *elasticache.DescribeUpdateActionsInput, fn func(*elasticache.DescribeUpdateActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUpdateActionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUpdateActionsOutput{}
	fnCacher := func(out *elasticache.DescribeUpdateActionsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUpdateActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUpdateActionsPagesWithContext(ctx context.Context, input *elasticache.DescribeUpdateActionsInput, fn func(*elasticache.DescribeUpdateActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUpdateActionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUpdateActionsOutput{}
	fnCacher := func(out *elasticache.DescribeUpdateActionsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUpdateActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUpdateActionsWithContext(ctx context.Context, input *elasticache.DescribeUpdateActionsInput, opts ...request.Option) (*elasticache.DescribeUpdateActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUpdateActionsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUpdateActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUserGroups(input *elasticache.DescribeUserGroupsInput) (*elasticache.DescribeUserGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUserGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUserGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUserGroupsPages(input *elasticache.DescribeUserGroupsInput, fn func(*elasticache.DescribeUserGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUserGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUserGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeUserGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUserGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUserGroupsPagesWithContext(ctx context.Context, input *elasticache.DescribeUserGroupsInput, fn func(*elasticache.DescribeUserGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUserGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUserGroupsOutput{}
	fnCacher := func(out *elasticache.DescribeUserGroupsOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUserGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUserGroupsWithContext(ctx context.Context, input *elasticache.DescribeUserGroupsInput, opts ...request.Option) (*elasticache.DescribeUserGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUserGroupsOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUserGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUsers(input *elasticache.DescribeUsersInput) (*elasticache.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUsersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) DescribeUsersPages(input *elasticache.DescribeUsersInput, fn func(*elasticache.DescribeUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUsersOutput{}
	fnCacher := func(out *elasticache.DescribeUsersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUsersPagesWithContext(ctx context.Context, input *elasticache.DescribeUsersInput, fn func(*elasticache.DescribeUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticache.DescribeUsersOutput), false)
		return nil
	}
	cachable := true
	output := &elasticache.DescribeUsersOutput{}
	fnCacher := func(out *elasticache.DescribeUsersOutput, more bool) bool {
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
	if err := c.ElastiCacheAPI.DescribeUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElastiCache) DescribeUsersWithContext(ctx context.Context, input *elasticache.DescribeUsersInput, opts ...request.Option) (*elasticache.DescribeUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.DescribeUsersOutput), nil
	}
	out, err := c.ElastiCacheAPI.DescribeUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) ListAllowedNodeTypeModifications(input *elasticache.ListAllowedNodeTypeModificationsInput) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.ListAllowedNodeTypeModificationsOutput), nil
	}
	out, err := c.ElastiCacheAPI.ListAllowedNodeTypeModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElastiCache) ListAllowedNodeTypeModificationsWithContext(ctx context.Context, input *elasticache.ListAllowedNodeTypeModificationsInput, opts ...request.Option) (*elasticache.ListAllowedNodeTypeModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticache.ListAllowedNodeTypeModificationsOutput), nil
	}
	out, err := c.ElastiCacheAPI.ListAllowedNodeTypeModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
