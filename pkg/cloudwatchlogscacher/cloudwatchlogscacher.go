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
package cloudwatchlogscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs/cloudwatchlogsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudWatchLogs struct {
	cloudwatchlogsiface.CloudWatchLogsAPI
	cache *cache.Cache
}

func New(cloudwatchlogsapi cloudwatchlogsiface.CloudWatchLogsAPI) *CloudWatchLogs {
	return &CloudWatchLogs{
		CloudWatchLogsAPI: cloudwatchlogsapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudWatchLogs) DescribeDestinations(input *cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeDestinationsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeDestinationsPages(input *cloudwatchlogs.DescribeDestinationsInput, fn func(*cloudwatchlogs.DescribeDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeDestinationsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeDestinationsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeDestinationsPagesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeDestinationsInput, fn func(*cloudwatchlogs.DescribeDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeDestinationsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeDestinationsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeDestinationsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeDestinationsInput, opts ...request.Option) (*cloudwatchlogs.DescribeDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeDestinationsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeExportTasks(input *cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeExportTasksOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeExportTasksWithContext(ctx context.Context, input *cloudwatchlogs.DescribeExportTasksInput, opts ...request.Option) (*cloudwatchlogs.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeExportTasksOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeLogGroups(input *cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeLogGroupsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeLogGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeLogGroupsPages(input *cloudwatchlogs.DescribeLogGroupsInput, fn func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeLogGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeLogGroupsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeLogGroupsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeLogGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeLogGroupsPagesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, fn func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeLogGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeLogGroupsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeLogGroupsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeLogGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeLogGroupsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogGroupsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeLogGroupsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeLogGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeLogStreams(input *cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeLogStreamsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeLogStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeLogStreamsPages(input *cloudwatchlogs.DescribeLogStreamsInput, fn func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeLogStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeLogStreamsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeLogStreamsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeLogStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeLogStreamsPagesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogStreamsInput, fn func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeLogStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeLogStreamsOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeLogStreamsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeLogStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeLogStreamsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeLogStreamsInput, opts ...request.Option) (*cloudwatchlogs.DescribeLogStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeLogStreamsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeLogStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeMetricFilters(input *cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeMetricFiltersOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeMetricFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeMetricFiltersPages(input *cloudwatchlogs.DescribeMetricFiltersInput, fn func(*cloudwatchlogs.DescribeMetricFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeMetricFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeMetricFiltersOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeMetricFiltersOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeMetricFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeMetricFiltersPagesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeMetricFiltersInput, fn func(*cloudwatchlogs.DescribeMetricFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeMetricFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeMetricFiltersOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeMetricFiltersOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeMetricFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeMetricFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeMetricFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeMetricFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeMetricFiltersOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeMetricFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeQueries(input *cloudwatchlogs.DescribeQueriesInput) (*cloudwatchlogs.DescribeQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeQueriesOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeQueries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeQueriesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeQueriesInput, opts ...request.Option) (*cloudwatchlogs.DescribeQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeQueriesOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeQueriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeQueryDefinitions(input *cloudwatchlogs.DescribeQueryDefinitionsInput) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeQueryDefinitionsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeQueryDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeQueryDefinitionsWithContext(ctx context.Context, input *cloudwatchlogs.DescribeQueryDefinitionsInput, opts ...request.Option) (*cloudwatchlogs.DescribeQueryDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeQueryDefinitionsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeQueryDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeResourcePolicies(input *cloudwatchlogs.DescribeResourcePoliciesInput) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeResourcePoliciesOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeResourcePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeResourcePoliciesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeResourcePoliciesInput, opts ...request.Option) (*cloudwatchlogs.DescribeResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeResourcePoliciesOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeResourcePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeSubscriptionFilters(input *cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeSubscriptionFiltersOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeSubscriptionFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) DescribeSubscriptionFiltersPages(input *cloudwatchlogs.DescribeSubscriptionFiltersInput, fn func(*cloudwatchlogs.DescribeSubscriptionFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeSubscriptionFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeSubscriptionFiltersOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeSubscriptionFiltersOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeSubscriptionFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeSubscriptionFiltersPagesWithContext(ctx context.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput, fn func(*cloudwatchlogs.DescribeSubscriptionFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.DescribeSubscriptionFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.DescribeSubscriptionFiltersOutput{}
	fnCacher := func(out *cloudwatchlogs.DescribeSubscriptionFiltersOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.DescribeSubscriptionFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) DescribeSubscriptionFiltersWithContext(ctx context.Context, input *cloudwatchlogs.DescribeSubscriptionFiltersInput, opts ...request.Option) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.DescribeSubscriptionFiltersOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.DescribeSubscriptionFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetDataProtectionPolicy(input *cloudwatchlogs.GetDataProtectionPolicyInput) (*cloudwatchlogs.GetDataProtectionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetDataProtectionPolicyOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetDataProtectionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetDataProtectionPolicyWithContext(ctx context.Context, input *cloudwatchlogs.GetDataProtectionPolicyInput, opts ...request.Option) (*cloudwatchlogs.GetDataProtectionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetDataProtectionPolicyOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetDataProtectionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogEvents(input *cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogEventsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogEventsPages(input *cloudwatchlogs.GetLogEventsInput, fn func(*cloudwatchlogs.GetLogEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.GetLogEventsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.GetLogEventsOutput{}
	fnCacher := func(out *cloudwatchlogs.GetLogEventsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.GetLogEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) GetLogEventsPagesWithContext(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput, fn func(*cloudwatchlogs.GetLogEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudwatchlogs.GetLogEventsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudwatchlogs.GetLogEventsOutput{}
	fnCacher := func(out *cloudwatchlogs.GetLogEventsOutput, more bool) bool {
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
	if err := c.CloudWatchLogsAPI.GetLogEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudWatchLogs) GetLogEventsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogEventsInput, opts ...request.Option) (*cloudwatchlogs.GetLogEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogEventsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogGroupFields(input *cloudwatchlogs.GetLogGroupFieldsInput) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogGroupFieldsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogGroupFields(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogGroupFieldsWithContext(ctx context.Context, input *cloudwatchlogs.GetLogGroupFieldsInput, opts ...request.Option) (*cloudwatchlogs.GetLogGroupFieldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogGroupFieldsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogGroupFieldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogRecord(input *cloudwatchlogs.GetLogRecordInput) (*cloudwatchlogs.GetLogRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogRecordOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogRecord(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetLogRecordWithContext(ctx context.Context, input *cloudwatchlogs.GetLogRecordInput, opts ...request.Option) (*cloudwatchlogs.GetLogRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetLogRecordOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetLogRecordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetQueryResults(input *cloudwatchlogs.GetQueryResultsInput) (*cloudwatchlogs.GetQueryResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetQueryResultsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetQueryResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) GetQueryResultsWithContext(ctx context.Context, input *cloudwatchlogs.GetQueryResultsInput, opts ...request.Option) (*cloudwatchlogs.GetQueryResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.GetQueryResultsOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.GetQueryResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) ListTagsForResource(input *cloudwatchlogs.ListTagsForResourceInput) (*cloudwatchlogs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) ListTagsForResourceWithContext(ctx context.Context, input *cloudwatchlogs.ListTagsForResourceInput, opts ...request.Option) (*cloudwatchlogs.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) ListTagsLogGroup(input *cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.ListTagsLogGroupOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.ListTagsLogGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudWatchLogs) ListTagsLogGroupWithContext(ctx context.Context, input *cloudwatchlogs.ListTagsLogGroupInput, opts ...request.Option) (*cloudwatchlogs.ListTagsLogGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudwatchlogs.ListTagsLogGroupOutput), nil
	}
	out, err := c.CloudWatchLogsAPI.ListTagsLogGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
