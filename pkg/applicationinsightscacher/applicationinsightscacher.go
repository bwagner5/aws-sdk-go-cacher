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
package applicationinsightscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/applicationinsights"
	"github.com/aws/aws-sdk-go/service/applicationinsights/applicationinsightsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApplicationInsights struct {
	applicationinsightsiface.ApplicationInsightsAPI
	cache *cache.Cache
}

func New(applicationinsightsapi applicationinsightsiface.ApplicationInsightsAPI) *ApplicationInsights {
	return &ApplicationInsights{
		ApplicationInsightsAPI: applicationinsightsapi,
		cache:                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApplicationInsights) DescribeApplication(input *applicationinsights.DescribeApplicationInput) (*applicationinsights.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeApplicationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeApplicationWithContext(ctx context.Context, input *applicationinsights.DescribeApplicationInput, opts ...request.Option) (*applicationinsights.DescribeApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeApplicationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponent(input *applicationinsights.DescribeComponentInput) (*applicationinsights.DescribeComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponentConfiguration(input *applicationinsights.DescribeComponentConfigurationInput) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentConfigurationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponentConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponentConfigurationRecommendation(input *applicationinsights.DescribeComponentConfigurationRecommendationInput) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentConfigurationRecommendationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponentConfigurationRecommendation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponentConfigurationRecommendationWithContext(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationRecommendationInput, opts ...request.Option) (*applicationinsights.DescribeComponentConfigurationRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentConfigurationRecommendationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponentConfigurationRecommendationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponentConfigurationWithContext(ctx context.Context, input *applicationinsights.DescribeComponentConfigurationInput, opts ...request.Option) (*applicationinsights.DescribeComponentConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentConfigurationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponentConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeComponentWithContext(ctx context.Context, input *applicationinsights.DescribeComponentInput, opts ...request.Option) (*applicationinsights.DescribeComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeComponentOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeLogPattern(input *applicationinsights.DescribeLogPatternInput) (*applicationinsights.DescribeLogPatternOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeLogPatternOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeLogPattern(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeLogPatternWithContext(ctx context.Context, input *applicationinsights.DescribeLogPatternInput, opts ...request.Option) (*applicationinsights.DescribeLogPatternOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeLogPatternOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeLogPatternWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeObservation(input *applicationinsights.DescribeObservationInput) (*applicationinsights.DescribeObservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeObservationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeObservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeObservationWithContext(ctx context.Context, input *applicationinsights.DescribeObservationInput, opts ...request.Option) (*applicationinsights.DescribeObservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeObservationOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeObservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeProblem(input *applicationinsights.DescribeProblemInput) (*applicationinsights.DescribeProblemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeProblemOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeProblem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeProblemObservations(input *applicationinsights.DescribeProblemObservationsInput) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeProblemObservationsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeProblemObservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeProblemObservationsWithContext(ctx context.Context, input *applicationinsights.DescribeProblemObservationsInput, opts ...request.Option) (*applicationinsights.DescribeProblemObservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeProblemObservationsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeProblemObservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) DescribeProblemWithContext(ctx context.Context, input *applicationinsights.DescribeProblemInput, opts ...request.Option) (*applicationinsights.DescribeProblemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.DescribeProblemOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.DescribeProblemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListApplications(input *applicationinsights.ListApplicationsInput) (*applicationinsights.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListApplicationsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListApplicationsPages(input *applicationinsights.ListApplicationsInput, fn func(*applicationinsights.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListApplicationsOutput{}
	fnCacher := func(out *applicationinsights.ListApplicationsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListApplicationsPagesWithContext(ctx context.Context, input *applicationinsights.ListApplicationsInput, fn func(*applicationinsights.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListApplicationsOutput{}
	fnCacher := func(out *applicationinsights.ListApplicationsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListApplicationsWithContext(ctx context.Context, input *applicationinsights.ListApplicationsInput, opts ...request.Option) (*applicationinsights.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListApplicationsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListComponents(input *applicationinsights.ListComponentsInput) (*applicationinsights.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListComponentsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListComponentsPages(input *applicationinsights.ListComponentsInput, fn func(*applicationinsights.ListComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListComponentsOutput{}
	fnCacher := func(out *applicationinsights.ListComponentsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListComponentsPagesWithContext(ctx context.Context, input *applicationinsights.ListComponentsInput, fn func(*applicationinsights.ListComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListComponentsOutput{}
	fnCacher := func(out *applicationinsights.ListComponentsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListComponentsWithContext(ctx context.Context, input *applicationinsights.ListComponentsInput, opts ...request.Option) (*applicationinsights.ListComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListComponentsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListConfigurationHistory(input *applicationinsights.ListConfigurationHistoryInput) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListConfigurationHistoryOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListConfigurationHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListConfigurationHistoryPages(input *applicationinsights.ListConfigurationHistoryInput, fn func(*applicationinsights.ListConfigurationHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListConfigurationHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListConfigurationHistoryOutput{}
	fnCacher := func(out *applicationinsights.ListConfigurationHistoryOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListConfigurationHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListConfigurationHistoryPagesWithContext(ctx context.Context, input *applicationinsights.ListConfigurationHistoryInput, fn func(*applicationinsights.ListConfigurationHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListConfigurationHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListConfigurationHistoryOutput{}
	fnCacher := func(out *applicationinsights.ListConfigurationHistoryOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListConfigurationHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListConfigurationHistoryWithContext(ctx context.Context, input *applicationinsights.ListConfigurationHistoryInput, opts ...request.Option) (*applicationinsights.ListConfigurationHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListConfigurationHistoryOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListConfigurationHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListLogPatternSets(input *applicationinsights.ListLogPatternSetsInput) (*applicationinsights.ListLogPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListLogPatternSetsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListLogPatternSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListLogPatternSetsPages(input *applicationinsights.ListLogPatternSetsInput, fn func(*applicationinsights.ListLogPatternSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListLogPatternSetsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListLogPatternSetsOutput{}
	fnCacher := func(out *applicationinsights.ListLogPatternSetsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListLogPatternSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListLogPatternSetsPagesWithContext(ctx context.Context, input *applicationinsights.ListLogPatternSetsInput, fn func(*applicationinsights.ListLogPatternSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListLogPatternSetsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListLogPatternSetsOutput{}
	fnCacher := func(out *applicationinsights.ListLogPatternSetsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListLogPatternSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListLogPatternSetsWithContext(ctx context.Context, input *applicationinsights.ListLogPatternSetsInput, opts ...request.Option) (*applicationinsights.ListLogPatternSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListLogPatternSetsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListLogPatternSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListLogPatterns(input *applicationinsights.ListLogPatternsInput) (*applicationinsights.ListLogPatternsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListLogPatternsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListLogPatterns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListLogPatternsPages(input *applicationinsights.ListLogPatternsInput, fn func(*applicationinsights.ListLogPatternsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListLogPatternsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListLogPatternsOutput{}
	fnCacher := func(out *applicationinsights.ListLogPatternsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListLogPatternsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListLogPatternsPagesWithContext(ctx context.Context, input *applicationinsights.ListLogPatternsInput, fn func(*applicationinsights.ListLogPatternsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListLogPatternsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListLogPatternsOutput{}
	fnCacher := func(out *applicationinsights.ListLogPatternsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListLogPatternsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListLogPatternsWithContext(ctx context.Context, input *applicationinsights.ListLogPatternsInput, opts ...request.Option) (*applicationinsights.ListLogPatternsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListLogPatternsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListLogPatternsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListProblems(input *applicationinsights.ListProblemsInput) (*applicationinsights.ListProblemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListProblemsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListProblems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListProblemsPages(input *applicationinsights.ListProblemsInput, fn func(*applicationinsights.ListProblemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListProblemsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListProblemsOutput{}
	fnCacher := func(out *applicationinsights.ListProblemsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListProblemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListProblemsPagesWithContext(ctx context.Context, input *applicationinsights.ListProblemsInput, fn func(*applicationinsights.ListProblemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationinsights.ListProblemsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationinsights.ListProblemsOutput{}
	fnCacher := func(out *applicationinsights.ListProblemsOutput, more bool) bool {
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
	if err := c.ApplicationInsightsAPI.ListProblemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationInsights) ListProblemsWithContext(ctx context.Context, input *applicationinsights.ListProblemsInput, opts ...request.Option) (*applicationinsights.ListProblemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListProblemsOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListProblemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListTagsForResource(input *applicationinsights.ListTagsForResourceInput) (*applicationinsights.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListTagsForResourceOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationInsights) ListTagsForResourceWithContext(ctx context.Context, input *applicationinsights.ListTagsForResourceInput, opts ...request.Option) (*applicationinsights.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationinsights.ListTagsForResourceOutput), nil
	}
	out, err := c.ApplicationInsightsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
