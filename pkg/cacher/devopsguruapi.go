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
	"github.com/aws/aws-sdk-go/service/devopsguru"
	"github.com/aws/aws-sdk-go/service/devopsguru/devopsguruiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DevOpsGuru struct {
	devopsguruiface.DevOpsGuruAPI
	cache *cache.Cache
}

func NewDevOpsGuru(devopsguruapi devopsguruiface.DevOpsGuruAPI) *DevOpsGuru {
	return &DevOpsGuru{
		DevOpsGuruAPI: devopsguruapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DevOpsGuru) DescribeAccountHealth(input *devopsguru.DescribeAccountHealthInput) (*devopsguru.DescribeAccountHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAccountHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAccountHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeAccountHealthWithContext(ctx context.Context, input *devopsguru.DescribeAccountHealthInput, opts ...request.Option) (*devopsguru.DescribeAccountHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAccountHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAccountHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeAccountOverview(input *devopsguru.DescribeAccountOverviewInput) (*devopsguru.DescribeAccountOverviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAccountOverviewOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAccountOverview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeAccountOverviewWithContext(ctx context.Context, input *devopsguru.DescribeAccountOverviewInput, opts ...request.Option) (*devopsguru.DescribeAccountOverviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAccountOverviewOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAccountOverviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeAnomaly(input *devopsguru.DescribeAnomalyInput) (*devopsguru.DescribeAnomalyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAnomalyOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAnomaly(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeAnomalyWithContext(ctx context.Context, input *devopsguru.DescribeAnomalyInput, opts ...request.Option) (*devopsguru.DescribeAnomalyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeAnomalyOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeAnomalyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeEventSourcesConfig(input *devopsguru.DescribeEventSourcesConfigInput) (*devopsguru.DescribeEventSourcesConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeEventSourcesConfigOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeEventSourcesConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeEventSourcesConfigWithContext(ctx context.Context, input *devopsguru.DescribeEventSourcesConfigInput, opts ...request.Option) (*devopsguru.DescribeEventSourcesConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeEventSourcesConfigOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeEventSourcesConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeFeedback(input *devopsguru.DescribeFeedbackInput) (*devopsguru.DescribeFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeFeedbackOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeFeedback(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeFeedbackWithContext(ctx context.Context, input *devopsguru.DescribeFeedbackInput, opts ...request.Option) (*devopsguru.DescribeFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeFeedbackOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeFeedbackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeInsight(input *devopsguru.DescribeInsightInput) (*devopsguru.DescribeInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeInsightOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeInsight(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeInsightWithContext(ctx context.Context, input *devopsguru.DescribeInsightInput, opts ...request.Option) (*devopsguru.DescribeInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeInsightOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeInsightWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationHealth(input *devopsguru.DescribeOrganizationHealthInput) (*devopsguru.DescribeOrganizationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationHealthWithContext(ctx context.Context, input *devopsguru.DescribeOrganizationHealthInput, opts ...request.Option) (*devopsguru.DescribeOrganizationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationOverview(input *devopsguru.DescribeOrganizationOverviewInput) (*devopsguru.DescribeOrganizationOverviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationOverviewOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationOverview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationOverviewWithContext(ctx context.Context, input *devopsguru.DescribeOrganizationOverviewInput, opts ...request.Option) (*devopsguru.DescribeOrganizationOverviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationOverviewOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationOverviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationResourceCollectionHealth(input *devopsguru.DescribeOrganizationResourceCollectionHealthInput) (*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationResourceCollectionHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeOrganizationResourceCollectionHealthPages(input *devopsguru.DescribeOrganizationResourceCollectionHealthInput, fn func(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.DescribeOrganizationResourceCollectionHealthOutput{}
	fnCacher := func(out *devopsguru.DescribeOrganizationResourceCollectionHealthOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.DescribeOrganizationResourceCollectionHealthPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) DescribeOrganizationResourceCollectionHealthPagesWithContext(ctx context.Context, input *devopsguru.DescribeOrganizationResourceCollectionHealthInput, fn func(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.DescribeOrganizationResourceCollectionHealthOutput{}
	fnCacher := func(out *devopsguru.DescribeOrganizationResourceCollectionHealthOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.DescribeOrganizationResourceCollectionHealthPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) DescribeOrganizationResourceCollectionHealthWithContext(ctx context.Context, input *devopsguru.DescribeOrganizationResourceCollectionHealthInput, opts ...request.Option) (*devopsguru.DescribeOrganizationResourceCollectionHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeOrganizationResourceCollectionHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeOrganizationResourceCollectionHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeResourceCollectionHealth(input *devopsguru.DescribeResourceCollectionHealthInput) (*devopsguru.DescribeResourceCollectionHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeResourceCollectionHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeResourceCollectionHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeResourceCollectionHealthPages(input *devopsguru.DescribeResourceCollectionHealthInput, fn func(*devopsguru.DescribeResourceCollectionHealthOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.DescribeResourceCollectionHealthOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.DescribeResourceCollectionHealthOutput{}
	fnCacher := func(out *devopsguru.DescribeResourceCollectionHealthOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.DescribeResourceCollectionHealthPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) DescribeResourceCollectionHealthPagesWithContext(ctx context.Context, input *devopsguru.DescribeResourceCollectionHealthInput, fn func(*devopsguru.DescribeResourceCollectionHealthOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.DescribeResourceCollectionHealthOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.DescribeResourceCollectionHealthOutput{}
	fnCacher := func(out *devopsguru.DescribeResourceCollectionHealthOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.DescribeResourceCollectionHealthPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) DescribeResourceCollectionHealthWithContext(ctx context.Context, input *devopsguru.DescribeResourceCollectionHealthInput, opts ...request.Option) (*devopsguru.DescribeResourceCollectionHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeResourceCollectionHealthOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeResourceCollectionHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeServiceIntegration(input *devopsguru.DescribeServiceIntegrationInput) (*devopsguru.DescribeServiceIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeServiceIntegrationOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeServiceIntegration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) DescribeServiceIntegrationWithContext(ctx context.Context, input *devopsguru.DescribeServiceIntegrationInput, opts ...request.Option) (*devopsguru.DescribeServiceIntegrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.DescribeServiceIntegrationOutput), nil
	}
	out, err := c.DevOpsGuruAPI.DescribeServiceIntegrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) GetCostEstimation(input *devopsguru.GetCostEstimationInput) (*devopsguru.GetCostEstimationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.GetCostEstimationOutput), nil
	}
	out, err := c.DevOpsGuruAPI.GetCostEstimation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) GetCostEstimationPages(input *devopsguru.GetCostEstimationInput, fn func(*devopsguru.GetCostEstimationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.GetCostEstimationOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.GetCostEstimationOutput{}
	fnCacher := func(out *devopsguru.GetCostEstimationOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.GetCostEstimationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) GetCostEstimationPagesWithContext(ctx context.Context, input *devopsguru.GetCostEstimationInput, fn func(*devopsguru.GetCostEstimationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.GetCostEstimationOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.GetCostEstimationOutput{}
	fnCacher := func(out *devopsguru.GetCostEstimationOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.GetCostEstimationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) GetCostEstimationWithContext(ctx context.Context, input *devopsguru.GetCostEstimationInput, opts ...request.Option) (*devopsguru.GetCostEstimationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.GetCostEstimationOutput), nil
	}
	out, err := c.DevOpsGuruAPI.GetCostEstimationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) GetResourceCollection(input *devopsguru.GetResourceCollectionInput) (*devopsguru.GetResourceCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.GetResourceCollectionOutput), nil
	}
	out, err := c.DevOpsGuruAPI.GetResourceCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) GetResourceCollectionPages(input *devopsguru.GetResourceCollectionInput, fn func(*devopsguru.GetResourceCollectionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.GetResourceCollectionOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.GetResourceCollectionOutput{}
	fnCacher := func(out *devopsguru.GetResourceCollectionOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.GetResourceCollectionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) GetResourceCollectionPagesWithContext(ctx context.Context, input *devopsguru.GetResourceCollectionInput, fn func(*devopsguru.GetResourceCollectionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.GetResourceCollectionOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.GetResourceCollectionOutput{}
	fnCacher := func(out *devopsguru.GetResourceCollectionOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.GetResourceCollectionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) GetResourceCollectionWithContext(ctx context.Context, input *devopsguru.GetResourceCollectionInput, opts ...request.Option) (*devopsguru.GetResourceCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.GetResourceCollectionOutput), nil
	}
	out, err := c.DevOpsGuruAPI.GetResourceCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListAnomaliesForInsight(input *devopsguru.ListAnomaliesForInsightInput) (*devopsguru.ListAnomaliesForInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListAnomaliesForInsightOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListAnomaliesForInsight(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListAnomaliesForInsightPages(input *devopsguru.ListAnomaliesForInsightInput, fn func(*devopsguru.ListAnomaliesForInsightOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListAnomaliesForInsightOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListAnomaliesForInsightOutput{}
	fnCacher := func(out *devopsguru.ListAnomaliesForInsightOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListAnomaliesForInsightPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListAnomaliesForInsightPagesWithContext(ctx context.Context, input *devopsguru.ListAnomaliesForInsightInput, fn func(*devopsguru.ListAnomaliesForInsightOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListAnomaliesForInsightOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListAnomaliesForInsightOutput{}
	fnCacher := func(out *devopsguru.ListAnomaliesForInsightOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListAnomaliesForInsightPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListAnomaliesForInsightWithContext(ctx context.Context, input *devopsguru.ListAnomaliesForInsightInput, opts ...request.Option) (*devopsguru.ListAnomaliesForInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListAnomaliesForInsightOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListAnomaliesForInsightWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListAnomalousLogGroups(input *devopsguru.ListAnomalousLogGroupsInput) (*devopsguru.ListAnomalousLogGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListAnomalousLogGroupsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListAnomalousLogGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListAnomalousLogGroupsPages(input *devopsguru.ListAnomalousLogGroupsInput, fn func(*devopsguru.ListAnomalousLogGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListAnomalousLogGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListAnomalousLogGroupsOutput{}
	fnCacher := func(out *devopsguru.ListAnomalousLogGroupsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListAnomalousLogGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListAnomalousLogGroupsPagesWithContext(ctx context.Context, input *devopsguru.ListAnomalousLogGroupsInput, fn func(*devopsguru.ListAnomalousLogGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListAnomalousLogGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListAnomalousLogGroupsOutput{}
	fnCacher := func(out *devopsguru.ListAnomalousLogGroupsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListAnomalousLogGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListAnomalousLogGroupsWithContext(ctx context.Context, input *devopsguru.ListAnomalousLogGroupsInput, opts ...request.Option) (*devopsguru.ListAnomalousLogGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListAnomalousLogGroupsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListAnomalousLogGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListEvents(input *devopsguru.ListEventsInput) (*devopsguru.ListEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListEventsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListEventsPages(input *devopsguru.ListEventsInput, fn func(*devopsguru.ListEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListEventsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListEventsOutput{}
	fnCacher := func(out *devopsguru.ListEventsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListEventsPagesWithContext(ctx context.Context, input *devopsguru.ListEventsInput, fn func(*devopsguru.ListEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListEventsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListEventsOutput{}
	fnCacher := func(out *devopsguru.ListEventsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListEventsWithContext(ctx context.Context, input *devopsguru.ListEventsInput, opts ...request.Option) (*devopsguru.ListEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListEventsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListInsights(input *devopsguru.ListInsightsInput) (*devopsguru.ListInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListInsightsPages(input *devopsguru.ListInsightsInput, fn func(*devopsguru.ListInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListInsightsOutput{}
	fnCacher := func(out *devopsguru.ListInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListInsightsPagesWithContext(ctx context.Context, input *devopsguru.ListInsightsInput, fn func(*devopsguru.ListInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListInsightsOutput{}
	fnCacher := func(out *devopsguru.ListInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListInsightsWithContext(ctx context.Context, input *devopsguru.ListInsightsInput, opts ...request.Option) (*devopsguru.ListInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListMonitoredResources(input *devopsguru.ListMonitoredResourcesInput) (*devopsguru.ListMonitoredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListMonitoredResourcesOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListMonitoredResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListMonitoredResourcesPages(input *devopsguru.ListMonitoredResourcesInput, fn func(*devopsguru.ListMonitoredResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListMonitoredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListMonitoredResourcesOutput{}
	fnCacher := func(out *devopsguru.ListMonitoredResourcesOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListMonitoredResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListMonitoredResourcesPagesWithContext(ctx context.Context, input *devopsguru.ListMonitoredResourcesInput, fn func(*devopsguru.ListMonitoredResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListMonitoredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListMonitoredResourcesOutput{}
	fnCacher := func(out *devopsguru.ListMonitoredResourcesOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListMonitoredResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListMonitoredResourcesWithContext(ctx context.Context, input *devopsguru.ListMonitoredResourcesInput, opts ...request.Option) (*devopsguru.ListMonitoredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListMonitoredResourcesOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListMonitoredResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListNotificationChannels(input *devopsguru.ListNotificationChannelsInput) (*devopsguru.ListNotificationChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListNotificationChannelsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListNotificationChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListNotificationChannelsPages(input *devopsguru.ListNotificationChannelsInput, fn func(*devopsguru.ListNotificationChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListNotificationChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListNotificationChannelsOutput{}
	fnCacher := func(out *devopsguru.ListNotificationChannelsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListNotificationChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListNotificationChannelsPagesWithContext(ctx context.Context, input *devopsguru.ListNotificationChannelsInput, fn func(*devopsguru.ListNotificationChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListNotificationChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListNotificationChannelsOutput{}
	fnCacher := func(out *devopsguru.ListNotificationChannelsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListNotificationChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListNotificationChannelsWithContext(ctx context.Context, input *devopsguru.ListNotificationChannelsInput, opts ...request.Option) (*devopsguru.ListNotificationChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListNotificationChannelsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListNotificationChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListOrganizationInsights(input *devopsguru.ListOrganizationInsightsInput) (*devopsguru.ListOrganizationInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListOrganizationInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListOrganizationInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListOrganizationInsightsPages(input *devopsguru.ListOrganizationInsightsInput, fn func(*devopsguru.ListOrganizationInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListOrganizationInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListOrganizationInsightsOutput{}
	fnCacher := func(out *devopsguru.ListOrganizationInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListOrganizationInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListOrganizationInsightsPagesWithContext(ctx context.Context, input *devopsguru.ListOrganizationInsightsInput, fn func(*devopsguru.ListOrganizationInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListOrganizationInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListOrganizationInsightsOutput{}
	fnCacher := func(out *devopsguru.ListOrganizationInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListOrganizationInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListOrganizationInsightsWithContext(ctx context.Context, input *devopsguru.ListOrganizationInsightsInput, opts ...request.Option) (*devopsguru.ListOrganizationInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListOrganizationInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListOrganizationInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListRecommendations(input *devopsguru.ListRecommendationsInput) (*devopsguru.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListRecommendationsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) ListRecommendationsPages(input *devopsguru.ListRecommendationsInput, fn func(*devopsguru.ListRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListRecommendationsOutput{}
	fnCacher := func(out *devopsguru.ListRecommendationsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListRecommendationsPagesWithContext(ctx context.Context, input *devopsguru.ListRecommendationsInput, fn func(*devopsguru.ListRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.ListRecommendationsOutput{}
	fnCacher := func(out *devopsguru.ListRecommendationsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.ListRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) ListRecommendationsWithContext(ctx context.Context, input *devopsguru.ListRecommendationsInput, opts ...request.Option) (*devopsguru.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.ListRecommendationsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.ListRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) SearchInsights(input *devopsguru.SearchInsightsInput) (*devopsguru.SearchInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.SearchInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.SearchInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) SearchInsightsPages(input *devopsguru.SearchInsightsInput, fn func(*devopsguru.SearchInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.SearchInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.SearchInsightsOutput{}
	fnCacher := func(out *devopsguru.SearchInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.SearchInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) SearchInsightsPagesWithContext(ctx context.Context, input *devopsguru.SearchInsightsInput, fn func(*devopsguru.SearchInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.SearchInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.SearchInsightsOutput{}
	fnCacher := func(out *devopsguru.SearchInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.SearchInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) SearchInsightsWithContext(ctx context.Context, input *devopsguru.SearchInsightsInput, opts ...request.Option) (*devopsguru.SearchInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.SearchInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.SearchInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) SearchOrganizationInsights(input *devopsguru.SearchOrganizationInsightsInput) (*devopsguru.SearchOrganizationInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.SearchOrganizationInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.SearchOrganizationInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DevOpsGuru) SearchOrganizationInsightsPages(input *devopsguru.SearchOrganizationInsightsInput, fn func(*devopsguru.SearchOrganizationInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.SearchOrganizationInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.SearchOrganizationInsightsOutput{}
	fnCacher := func(out *devopsguru.SearchOrganizationInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.SearchOrganizationInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) SearchOrganizationInsightsPagesWithContext(ctx context.Context, input *devopsguru.SearchOrganizationInsightsInput, fn func(*devopsguru.SearchOrganizationInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*devopsguru.SearchOrganizationInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &devopsguru.SearchOrganizationInsightsOutput{}
	fnCacher := func(out *devopsguru.SearchOrganizationInsightsOutput, more bool) bool {
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
	if err := c.DevOpsGuruAPI.SearchOrganizationInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DevOpsGuru) SearchOrganizationInsightsWithContext(ctx context.Context, input *devopsguru.SearchOrganizationInsightsInput, opts ...request.Option) (*devopsguru.SearchOrganizationInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*devopsguru.SearchOrganizationInsightsOutput), nil
	}
	out, err := c.DevOpsGuruAPI.SearchOrganizationInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
