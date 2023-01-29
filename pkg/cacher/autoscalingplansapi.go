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
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"github.com/aws/aws-sdk-go/service/autoscalingplans/autoscalingplansiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AutoScalingPlans struct {
	autoscalingplansiface.AutoScalingPlansAPI
	cache *cache.Cache
}

func NewAutoScalingPlans(autoscalingplansapi autoscalingplansiface.AutoScalingPlansAPI) *AutoScalingPlans {
	return &AutoScalingPlans{
		AutoScalingPlansAPI: autoscalingplansapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AutoScalingPlans) DescribeScalingPlanResources(input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.DescribeScalingPlanResourcesOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.DescribeScalingPlanResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScalingPlans) DescribeScalingPlanResourcesWithContext(ctx context.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput, opts ...request.Option) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.DescribeScalingPlanResourcesOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.DescribeScalingPlanResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScalingPlans) DescribeScalingPlans(input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.DescribeScalingPlansOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.DescribeScalingPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScalingPlans) DescribeScalingPlansWithContext(ctx context.Context, input *autoscalingplans.DescribeScalingPlansInput, opts ...request.Option) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.DescribeScalingPlansOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.DescribeScalingPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScalingPlans) GetScalingPlanResourceForecastData(input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.GetScalingPlanResourceForecastDataOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.GetScalingPlanResourceForecastData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScalingPlans) GetScalingPlanResourceForecastDataWithContext(ctx context.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput, opts ...request.Option) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscalingplans.GetScalingPlanResourceForecastDataOutput), nil
	}
	out, err := c.AutoScalingPlansAPI.GetScalingPlanResourceForecastDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
