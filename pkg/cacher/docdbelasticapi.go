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
	"github.com/aws/aws-sdk-go/service/docdbelastic"
	"github.com/aws/aws-sdk-go/service/docdbelastic/docdbelasticiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DocDBElastic struct {
	docdbelasticiface.DocDBElasticAPI
	cache *cache.Cache
}

func NewDocDBElastic(docdbelasticapi docdbelasticiface.DocDBElasticAPI) *DocDBElastic {
	return &DocDBElastic{
		DocDBElasticAPI: docdbelasticapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DocDBElastic) GetCluster(input *docdbelastic.GetClusterInput) (*docdbelastic.GetClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.GetClusterOutput), nil
	}
	out, err := c.DocDBElasticAPI.GetCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) GetClusterSnapshot(input *docdbelastic.GetClusterSnapshotInput) (*docdbelastic.GetClusterSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.GetClusterSnapshotOutput), nil
	}
	out, err := c.DocDBElasticAPI.GetClusterSnapshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) GetClusterSnapshotWithContext(ctx context.Context, input *docdbelastic.GetClusterSnapshotInput, opts ...request.Option) (*docdbelastic.GetClusterSnapshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.GetClusterSnapshotOutput), nil
	}
	out, err := c.DocDBElasticAPI.GetClusterSnapshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) GetClusterWithContext(ctx context.Context, input *docdbelastic.GetClusterInput, opts ...request.Option) (*docdbelastic.GetClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.GetClusterOutput), nil
	}
	out, err := c.DocDBElasticAPI.GetClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListClusterSnapshots(input *docdbelastic.ListClusterSnapshotsInput) (*docdbelastic.ListClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListClusterSnapshotsOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListClusterSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListClusterSnapshotsPages(input *docdbelastic.ListClusterSnapshotsInput, fn func(*docdbelastic.ListClusterSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdbelastic.ListClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &docdbelastic.ListClusterSnapshotsOutput{}
	fnCacher := func(out *docdbelastic.ListClusterSnapshotsOutput, more bool) bool {
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
	if err := c.DocDBElasticAPI.ListClusterSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDBElastic) ListClusterSnapshotsPagesWithContext(ctx context.Context, input *docdbelastic.ListClusterSnapshotsInput, fn func(*docdbelastic.ListClusterSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdbelastic.ListClusterSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &docdbelastic.ListClusterSnapshotsOutput{}
	fnCacher := func(out *docdbelastic.ListClusterSnapshotsOutput, more bool) bool {
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
	if err := c.DocDBElasticAPI.ListClusterSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDBElastic) ListClusterSnapshotsWithContext(ctx context.Context, input *docdbelastic.ListClusterSnapshotsInput, opts ...request.Option) (*docdbelastic.ListClusterSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListClusterSnapshotsOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListClusterSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListClusters(input *docdbelastic.ListClustersInput) (*docdbelastic.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListClustersOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListClustersPages(input *docdbelastic.ListClustersInput, fn func(*docdbelastic.ListClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdbelastic.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdbelastic.ListClustersOutput{}
	fnCacher := func(out *docdbelastic.ListClustersOutput, more bool) bool {
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
	if err := c.DocDBElasticAPI.ListClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDBElastic) ListClustersPagesWithContext(ctx context.Context, input *docdbelastic.ListClustersInput, fn func(*docdbelastic.ListClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*docdbelastic.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &docdbelastic.ListClustersOutput{}
	fnCacher := func(out *docdbelastic.ListClustersOutput, more bool) bool {
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
	if err := c.DocDBElasticAPI.ListClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DocDBElastic) ListClustersWithContext(ctx context.Context, input *docdbelastic.ListClustersInput, opts ...request.Option) (*docdbelastic.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListClustersOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListTagsForResource(input *docdbelastic.ListTagsForResourceInput) (*docdbelastic.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListTagsForResourceOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DocDBElastic) ListTagsForResourceWithContext(ctx context.Context, input *docdbelastic.ListTagsForResourceInput, opts ...request.Option) (*docdbelastic.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*docdbelastic.ListTagsForResourceOutput), nil
	}
	out, err := c.DocDBElasticAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
