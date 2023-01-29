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
	"github.com/aws/aws-sdk-go/service/emrcontainers"
	"github.com/aws/aws-sdk-go/service/emrcontainers/emrcontainersiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EMRContainers struct {
	emrcontainersiface.EMRContainersAPI
	cache *cache.Cache
}

func NewEMRContainers(emrcontainersapi emrcontainersiface.EMRContainersAPI) *EMRContainers {
	return &EMRContainers{
		EMRContainersAPI: emrcontainersapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EMRContainers) DescribeJobRun(input *emrcontainers.DescribeJobRunInput) (*emrcontainers.DescribeJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeJobRunOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeJobRunWithContext(ctx context.Context, input *emrcontainers.DescribeJobRunInput, opts ...request.Option) (*emrcontainers.DescribeJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeJobRunOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeJobTemplate(input *emrcontainers.DescribeJobTemplateInput) (*emrcontainers.DescribeJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeJobTemplateOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeJobTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeJobTemplateWithContext(ctx context.Context, input *emrcontainers.DescribeJobTemplateInput, opts ...request.Option) (*emrcontainers.DescribeJobTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeJobTemplateOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeJobTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeManagedEndpoint(input *emrcontainers.DescribeManagedEndpointInput) (*emrcontainers.DescribeManagedEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeManagedEndpointOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeManagedEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeManagedEndpointWithContext(ctx context.Context, input *emrcontainers.DescribeManagedEndpointInput, opts ...request.Option) (*emrcontainers.DescribeManagedEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeManagedEndpointOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeManagedEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeVirtualCluster(input *emrcontainers.DescribeVirtualClusterInput) (*emrcontainers.DescribeVirtualClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeVirtualClusterOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeVirtualCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) DescribeVirtualClusterWithContext(ctx context.Context, input *emrcontainers.DescribeVirtualClusterInput, opts ...request.Option) (*emrcontainers.DescribeVirtualClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.DescribeVirtualClusterOutput), nil
	}
	out, err := c.EMRContainersAPI.DescribeVirtualClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListJobRuns(input *emrcontainers.ListJobRunsInput) (*emrcontainers.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListJobRunsOutput), nil
	}
	out, err := c.EMRContainersAPI.ListJobRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListJobRunsPages(input *emrcontainers.ListJobRunsInput, fn func(*emrcontainers.ListJobRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListJobRunsOutput{}
	fnCacher := func(out *emrcontainers.ListJobRunsOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListJobRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListJobRunsPagesWithContext(ctx context.Context, input *emrcontainers.ListJobRunsInput, fn func(*emrcontainers.ListJobRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListJobRunsOutput{}
	fnCacher := func(out *emrcontainers.ListJobRunsOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListJobRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListJobRunsWithContext(ctx context.Context, input *emrcontainers.ListJobRunsInput, opts ...request.Option) (*emrcontainers.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListJobRunsOutput), nil
	}
	out, err := c.EMRContainersAPI.ListJobRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListJobTemplates(input *emrcontainers.ListJobTemplatesInput) (*emrcontainers.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListJobTemplatesOutput), nil
	}
	out, err := c.EMRContainersAPI.ListJobTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListJobTemplatesPages(input *emrcontainers.ListJobTemplatesInput, fn func(*emrcontainers.ListJobTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListJobTemplatesOutput{}
	fnCacher := func(out *emrcontainers.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListJobTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListJobTemplatesPagesWithContext(ctx context.Context, input *emrcontainers.ListJobTemplatesInput, fn func(*emrcontainers.ListJobTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListJobTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListJobTemplatesOutput{}
	fnCacher := func(out *emrcontainers.ListJobTemplatesOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListJobTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListJobTemplatesWithContext(ctx context.Context, input *emrcontainers.ListJobTemplatesInput, opts ...request.Option) (*emrcontainers.ListJobTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListJobTemplatesOutput), nil
	}
	out, err := c.EMRContainersAPI.ListJobTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListManagedEndpoints(input *emrcontainers.ListManagedEndpointsInput) (*emrcontainers.ListManagedEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListManagedEndpointsOutput), nil
	}
	out, err := c.EMRContainersAPI.ListManagedEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListManagedEndpointsPages(input *emrcontainers.ListManagedEndpointsInput, fn func(*emrcontainers.ListManagedEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListManagedEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListManagedEndpointsOutput{}
	fnCacher := func(out *emrcontainers.ListManagedEndpointsOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListManagedEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListManagedEndpointsPagesWithContext(ctx context.Context, input *emrcontainers.ListManagedEndpointsInput, fn func(*emrcontainers.ListManagedEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListManagedEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListManagedEndpointsOutput{}
	fnCacher := func(out *emrcontainers.ListManagedEndpointsOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListManagedEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListManagedEndpointsWithContext(ctx context.Context, input *emrcontainers.ListManagedEndpointsInput, opts ...request.Option) (*emrcontainers.ListManagedEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListManagedEndpointsOutput), nil
	}
	out, err := c.EMRContainersAPI.ListManagedEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListTagsForResource(input *emrcontainers.ListTagsForResourceInput) (*emrcontainers.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListTagsForResourceOutput), nil
	}
	out, err := c.EMRContainersAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListTagsForResourceWithContext(ctx context.Context, input *emrcontainers.ListTagsForResourceInput, opts ...request.Option) (*emrcontainers.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListTagsForResourceOutput), nil
	}
	out, err := c.EMRContainersAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListVirtualClusters(input *emrcontainers.ListVirtualClustersInput) (*emrcontainers.ListVirtualClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListVirtualClustersOutput), nil
	}
	out, err := c.EMRContainersAPI.ListVirtualClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMRContainers) ListVirtualClustersPages(input *emrcontainers.ListVirtualClustersInput, fn func(*emrcontainers.ListVirtualClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListVirtualClustersOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListVirtualClustersOutput{}
	fnCacher := func(out *emrcontainers.ListVirtualClustersOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListVirtualClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListVirtualClustersPagesWithContext(ctx context.Context, input *emrcontainers.ListVirtualClustersInput, fn func(*emrcontainers.ListVirtualClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emrcontainers.ListVirtualClustersOutput), false)
		return nil
	}
	cachable := true
	output := &emrcontainers.ListVirtualClustersOutput{}
	fnCacher := func(out *emrcontainers.ListVirtualClustersOutput, more bool) bool {
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
	if err := c.EMRContainersAPI.ListVirtualClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMRContainers) ListVirtualClustersWithContext(ctx context.Context, input *emrcontainers.ListVirtualClustersInput, opts ...request.Option) (*emrcontainers.ListVirtualClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emrcontainers.ListVirtualClustersOutput), nil
	}
	out, err := c.EMRContainersAPI.ListVirtualClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
