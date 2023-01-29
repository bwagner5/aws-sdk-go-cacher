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
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice/applicationdiscoveryserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApplicationDiscoveryService struct {
	applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI
	cache *cache.Cache
}

func NewApplicationDiscoveryService(applicationdiscoveryserviceapi applicationdiscoveryserviceiface.ApplicationDiscoveryServiceAPI) *ApplicationDiscoveryService {
	return &ApplicationDiscoveryService{
		ApplicationDiscoveryServiceAPI: applicationdiscoveryserviceapi,
		cache:                          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApplicationDiscoveryService) BatchDeleteImportData(input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.BatchDeleteImportDataOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.BatchDeleteImportData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) BatchDeleteImportDataWithContext(ctx context.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput, opts ...request.Option) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.BatchDeleteImportDataOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.BatchDeleteImportDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeAgents(input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeAgentsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeAgents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeAgentsWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeAgentsInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeAgentsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeAgentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeConfigurations(input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeConfigurationsWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeContinuousExports(input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeContinuousExportsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeContinuousExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeContinuousExportsPages(input *applicationdiscoveryservice.DescribeContinuousExportsInput, fn func(*applicationdiscoveryservice.DescribeContinuousExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationdiscoveryservice.DescribeContinuousExportsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationdiscoveryservice.DescribeContinuousExportsOutput{}
	fnCacher := func(out *applicationdiscoveryservice.DescribeContinuousExportsOutput, more bool) bool {
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
	if err := c.ApplicationDiscoveryServiceAPI.DescribeContinuousExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationDiscoveryService) DescribeContinuousExportsPagesWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput, fn func(*applicationdiscoveryservice.DescribeContinuousExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationdiscoveryservice.DescribeContinuousExportsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationdiscoveryservice.DescribeContinuousExportsOutput{}
	fnCacher := func(out *applicationdiscoveryservice.DescribeContinuousExportsOutput, more bool) bool {
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
	if err := c.ApplicationDiscoveryServiceAPI.DescribeContinuousExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationDiscoveryService) DescribeContinuousExportsWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeContinuousExportsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeContinuousExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeExportConfigurations(input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeExportConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeExportConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeExportConfigurationsWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeExportConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeExportConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeExportTasks(input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeExportTasksOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeExportTasksWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeExportTasksInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeExportTasksOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeImportTasks(input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeImportTasksOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeImportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeImportTasksPages(input *applicationdiscoveryservice.DescribeImportTasksInput, fn func(*applicationdiscoveryservice.DescribeImportTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationdiscoveryservice.DescribeImportTasksOutput), false)
		return nil
	}
	cachable := true
	output := &applicationdiscoveryservice.DescribeImportTasksOutput{}
	fnCacher := func(out *applicationdiscoveryservice.DescribeImportTasksOutput, more bool) bool {
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
	if err := c.ApplicationDiscoveryServiceAPI.DescribeImportTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationDiscoveryService) DescribeImportTasksPagesWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeImportTasksInput, fn func(*applicationdiscoveryservice.DescribeImportTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationdiscoveryservice.DescribeImportTasksOutput), false)
		return nil
	}
	cachable := true
	output := &applicationdiscoveryservice.DescribeImportTasksOutput{}
	fnCacher := func(out *applicationdiscoveryservice.DescribeImportTasksOutput, more bool) bool {
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
	if err := c.ApplicationDiscoveryServiceAPI.DescribeImportTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationDiscoveryService) DescribeImportTasksWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeImportTasksInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeImportTasksOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeImportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeTags(input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeTagsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) DescribeTagsWithContext(ctx context.Context, input *applicationdiscoveryservice.DescribeTagsInput, opts ...request.Option) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.DescribeTagsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) GetDiscoverySummary(input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.GetDiscoverySummaryOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.GetDiscoverySummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) GetDiscoverySummaryWithContext(ctx context.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput, opts ...request.Option) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.GetDiscoverySummaryOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.GetDiscoverySummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) ListConfigurations(input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.ListConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.ListConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) ListConfigurationsWithContext(ctx context.Context, input *applicationdiscoveryservice.ListConfigurationsInput, opts ...request.Option) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.ListConfigurationsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.ListConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) ListServerNeighbors(input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.ListServerNeighborsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.ListServerNeighbors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationDiscoveryService) ListServerNeighborsWithContext(ctx context.Context, input *applicationdiscoveryservice.ListServerNeighborsInput, opts ...request.Option) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationdiscoveryservice.ListServerNeighborsOutput), nil
	}
	out, err := c.ApplicationDiscoveryServiceAPI.ListServerNeighborsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
