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
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ElasticBeanstalk struct {
	elasticbeanstalkiface.ElasticBeanstalkAPI
	cache *cache.Cache
}

func NewElasticBeanstalk(elasticbeanstalkapi elasticbeanstalkiface.ElasticBeanstalkAPI) *ElasticBeanstalk {
	return &ElasticBeanstalk{
		ElasticBeanstalkAPI: elasticbeanstalkapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ElasticBeanstalk) DescribeAccountAttributes(input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeAccountAttributesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeAccountAttributesWithContext(ctx context.Context, input *elasticbeanstalk.DescribeAccountAttributesInput, opts ...request.Option) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeAccountAttributesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeApplicationVersions(input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeApplicationVersionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeApplicationVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeApplicationVersionsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput, opts ...request.Option) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeApplicationVersionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeApplicationVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeApplications(input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeApplicationsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeApplicationsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeApplicationsInput, opts ...request.Option) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeApplicationsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeConfigurationOptions(input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeConfigurationOptionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeConfigurationOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeConfigurationOptionsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput, opts ...request.Option) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeConfigurationOptionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeConfigurationOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeConfigurationSettings(input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeConfigurationSettingsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeConfigurationSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeConfigurationSettingsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput, opts ...request.Option) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeConfigurationSettingsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeConfigurationSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentHealth(input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentHealthOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentHealthWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput, opts ...request.Option) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentHealthOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActionHistory(input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActionHistoryPages(input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, fn func(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput{}
	fnCacher := func(out *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActionHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActionHistoryPagesWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, fn func(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput{}
	fnCacher := func(out *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActionHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActionHistoryWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput, opts ...request.Option) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActions(input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentManagedActionsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput, opts ...request.Option) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentManagedActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentResources(input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentResourcesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEnvironmentResourcesWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput, opts ...request.Option) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEnvironmentResourcesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEnvironmentResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEvents(input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEventsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeEventsPages(input *elasticbeanstalk.DescribeEventsInput, fn func(*elasticbeanstalk.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.DescribeEventsOutput{}
	fnCacher := func(out *elasticbeanstalk.DescribeEventsOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) DescribeEventsPagesWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEventsInput, fn func(*elasticbeanstalk.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.DescribeEventsOutput{}
	fnCacher := func(out *elasticbeanstalk.DescribeEventsOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) DescribeEventsWithContext(ctx context.Context, input *elasticbeanstalk.DescribeEventsInput, opts ...request.Option) (*elasticbeanstalk.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeEventsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeInstancesHealth(input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeInstancesHealthOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeInstancesHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribeInstancesHealthWithContext(ctx context.Context, input *elasticbeanstalk.DescribeInstancesHealthInput, opts ...request.Option) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribeInstancesHealthOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribeInstancesHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribePlatformVersion(input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribePlatformVersionOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribePlatformVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) DescribePlatformVersionWithContext(ctx context.Context, input *elasticbeanstalk.DescribePlatformVersionInput, opts ...request.Option) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.DescribePlatformVersionOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.DescribePlatformVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListAvailableSolutionStacks(input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListAvailableSolutionStacksOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListAvailableSolutionStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListAvailableSolutionStacksWithContext(ctx context.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput, opts ...request.Option) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListAvailableSolutionStacksOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListAvailableSolutionStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListPlatformBranches(input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListPlatformBranchesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListPlatformBranches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListPlatformBranchesPages(input *elasticbeanstalk.ListPlatformBranchesInput, fn func(*elasticbeanstalk.ListPlatformBranchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.ListPlatformBranchesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.ListPlatformBranchesOutput{}
	fnCacher := func(out *elasticbeanstalk.ListPlatformBranchesOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.ListPlatformBranchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) ListPlatformBranchesPagesWithContext(ctx context.Context, input *elasticbeanstalk.ListPlatformBranchesInput, fn func(*elasticbeanstalk.ListPlatformBranchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.ListPlatformBranchesOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.ListPlatformBranchesOutput{}
	fnCacher := func(out *elasticbeanstalk.ListPlatformBranchesOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.ListPlatformBranchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) ListPlatformBranchesWithContext(ctx context.Context, input *elasticbeanstalk.ListPlatformBranchesInput, opts ...request.Option) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListPlatformBranchesOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListPlatformBranchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListPlatformVersions(input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListPlatformVersionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListPlatformVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListPlatformVersionsPages(input *elasticbeanstalk.ListPlatformVersionsInput, fn func(*elasticbeanstalk.ListPlatformVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.ListPlatformVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.ListPlatformVersionsOutput{}
	fnCacher := func(out *elasticbeanstalk.ListPlatformVersionsOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.ListPlatformVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) ListPlatformVersionsPagesWithContext(ctx context.Context, input *elasticbeanstalk.ListPlatformVersionsInput, fn func(*elasticbeanstalk.ListPlatformVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*elasticbeanstalk.ListPlatformVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &elasticbeanstalk.ListPlatformVersionsOutput{}
	fnCacher := func(out *elasticbeanstalk.ListPlatformVersionsOutput, more bool) bool {
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
	if err := c.ElasticBeanstalkAPI.ListPlatformVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ElasticBeanstalk) ListPlatformVersionsWithContext(ctx context.Context, input *elasticbeanstalk.ListPlatformVersionsInput, opts ...request.Option) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListPlatformVersionsOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListPlatformVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListTagsForResource(input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListTagsForResourceOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ElasticBeanstalk) ListTagsForResourceWithContext(ctx context.Context, input *elasticbeanstalk.ListTagsForResourceInput, opts ...request.Option) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*elasticbeanstalk.ListTagsForResourceOutput), nil
	}
	out, err := c.ElasticBeanstalkAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
