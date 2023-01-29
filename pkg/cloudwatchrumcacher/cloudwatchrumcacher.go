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
package cloudwatchrumcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchrum"
	"github.com/aws/aws-sdk-go/service/cloudwatchrum/cloudwatchrumiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudWatchRUM struct {
	cloudwatchrumiface.CloudWatchRUMAPI
	cache *cache.Cache
}

func New(cloudwatchrumapi cloudwatchrumiface.CloudWatchRUMAPI) *CloudWatchRUM {
	return &CloudWatchRUM{
		CloudWatchRUMAPI: cloudwatchrumapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudWatchRUM) BatchCreateRumMetricDefinitions(input *cloudwatchrum.BatchCreateRumMetricDefinitionsInput) (*cloudwatchrum.BatchCreateRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchCreateRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchCreateRumMetricDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) BatchCreateRumMetricDefinitionsWithContext(ctx context.Context, input *cloudwatchrum.BatchCreateRumMetricDefinitionsInput, opts ...request.Option) (*cloudwatchrum.BatchCreateRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchCreateRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchCreateRumMetricDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) BatchDeleteRumMetricDefinitions(input *cloudwatchrum.BatchDeleteRumMetricDefinitionsInput) (*cloudwatchrum.BatchDeleteRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchDeleteRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchDeleteRumMetricDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) BatchDeleteRumMetricDefinitionsWithContext(ctx context.Context, input *cloudwatchrum.BatchDeleteRumMetricDefinitionsInput, opts ...request.Option) (*cloudwatchrum.BatchDeleteRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchDeleteRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchDeleteRumMetricDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) BatchGetRumMetricDefinitions(input *cloudwatchrum.BatchGetRumMetricDefinitionsInput) (*cloudwatchrum.BatchGetRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchGetRumMetricDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) BatchGetRumMetricDefinitionsPages(input *cloudwatchrum.BatchGetRumMetricDefinitionsInput, fn func(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.BatchGetRumMetricDefinitionsOutput{}
	fnCacher := func(out *cloudwatchrum.BatchGetRumMetricDefinitionsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.BatchGetRumMetricDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) BatchGetRumMetricDefinitionsPagesWithContext(ctx context.Context, input *cloudwatchrum.BatchGetRumMetricDefinitionsInput, fn func(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.BatchGetRumMetricDefinitionsOutput{}
	fnCacher := func(out *cloudwatchrum.BatchGetRumMetricDefinitionsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.BatchGetRumMetricDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) BatchGetRumMetricDefinitionsWithContext(ctx context.Context, input *cloudwatchrum.BatchGetRumMetricDefinitionsInput, opts ...request.Option) (*cloudwatchrum.BatchGetRumMetricDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.BatchGetRumMetricDefinitionsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.BatchGetRumMetricDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) GetAppMonitor(input *cloudwatchrum.GetAppMonitorInput) (*cloudwatchrum.GetAppMonitorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.GetAppMonitorOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.GetAppMonitor(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) GetAppMonitorData(input *cloudwatchrum.GetAppMonitorDataInput) (*cloudwatchrum.GetAppMonitorDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.GetAppMonitorDataOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.GetAppMonitorData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) GetAppMonitorDataPages(input *cloudwatchrum.GetAppMonitorDataInput, fn func(*cloudwatchrum.GetAppMonitorDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.GetAppMonitorDataOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.GetAppMonitorDataOutput{}
	fnCacher := func(out *cloudwatchrum.GetAppMonitorDataOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.GetAppMonitorDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) GetAppMonitorDataPagesWithContext(ctx context.Context, input *cloudwatchrum.GetAppMonitorDataInput, fn func(*cloudwatchrum.GetAppMonitorDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.GetAppMonitorDataOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.GetAppMonitorDataOutput{}
	fnCacher := func(out *cloudwatchrum.GetAppMonitorDataOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.GetAppMonitorDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) GetAppMonitorDataWithContext(ctx context.Context, input *cloudwatchrum.GetAppMonitorDataInput, opts ...request.Option) (*cloudwatchrum.GetAppMonitorDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.GetAppMonitorDataOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.GetAppMonitorDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) GetAppMonitorWithContext(ctx context.Context, input *cloudwatchrum.GetAppMonitorInput, opts ...request.Option) (*cloudwatchrum.GetAppMonitorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.GetAppMonitorOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.GetAppMonitorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListAppMonitors(input *cloudwatchrum.ListAppMonitorsInput) (*cloudwatchrum.ListAppMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListAppMonitorsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListAppMonitors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListAppMonitorsPages(input *cloudwatchrum.ListAppMonitorsInput, fn func(*cloudwatchrum.ListAppMonitorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.ListAppMonitorsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.ListAppMonitorsOutput{}
	fnCacher := func(out *cloudwatchrum.ListAppMonitorsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.ListAppMonitorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) ListAppMonitorsPagesWithContext(ctx context.Context, input *cloudwatchrum.ListAppMonitorsInput, fn func(*cloudwatchrum.ListAppMonitorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.ListAppMonitorsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.ListAppMonitorsOutput{}
	fnCacher := func(out *cloudwatchrum.ListAppMonitorsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.ListAppMonitorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) ListAppMonitorsWithContext(ctx context.Context, input *cloudwatchrum.ListAppMonitorsInput, opts ...request.Option) (*cloudwatchrum.ListAppMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListAppMonitorsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListAppMonitorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListRumMetricsDestinations(input *cloudwatchrum.ListRumMetricsDestinationsInput) (*cloudwatchrum.ListRumMetricsDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListRumMetricsDestinationsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListRumMetricsDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListRumMetricsDestinationsPages(input *cloudwatchrum.ListRumMetricsDestinationsInput, fn func(*cloudwatchrum.ListRumMetricsDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.ListRumMetricsDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.ListRumMetricsDestinationsOutput{}
	fnCacher := func(out *cloudwatchrum.ListRumMetricsDestinationsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.ListRumMetricsDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) ListRumMetricsDestinationsPagesWithContext(ctx context.Context, input *cloudwatchrum.ListRumMetricsDestinationsInput, fn func(*cloudwatchrum.ListRumMetricsDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchrum.ListRumMetricsDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchrum.ListRumMetricsDestinationsOutput{}
	fnCacher := func(out *cloudwatchrum.ListRumMetricsDestinationsOutput, more bool) bool {
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
	if err := c.CloudWatchRUMAPI.ListRumMetricsDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchRUM) ListRumMetricsDestinationsWithContext(ctx context.Context, input *cloudwatchrum.ListRumMetricsDestinationsInput, opts ...request.Option) (*cloudwatchrum.ListRumMetricsDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListRumMetricsDestinationsOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListRumMetricsDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListTagsForResource(input *cloudwatchrum.ListTagsForResourceInput) (*cloudwatchrum.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchRUM) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatchrum.ListTagsForResourceInput, opts ...request.Option) (*cloudwatchrum.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchrum.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchRUMAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
