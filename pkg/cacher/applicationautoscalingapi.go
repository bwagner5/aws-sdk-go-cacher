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
	"github.com/aws/aws-sdk-go/service/applicationautoscaling"
	"github.com/aws/aws-sdk-go/service/applicationautoscaling/applicationautoscalingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ApplicationAutoScaling struct {
	applicationautoscalingiface.ApplicationAutoScalingAPI
	cache *cache.Cache
}

func NewApplicationAutoScaling(applicationautoscalingapi applicationautoscalingiface.ApplicationAutoScalingAPI) *ApplicationAutoScaling {
	return &ApplicationAutoScaling{
		ApplicationAutoScalingAPI: applicationautoscalingapi,
		cache:                     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ApplicationAutoScaling) DescribeScalableTargets(input *applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalableTargetsOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalableTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScalableTargetsPages(input *applicationautoscaling.DescribeScalableTargetsInput, fn func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalableTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalableTargetsOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalableTargetsOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalableTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalableTargetsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, fn func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalableTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalableTargetsOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalableTargetsOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalableTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalableTargetsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalableTargetsInput, opts ...request.Option) (*applicationautoscaling.DescribeScalableTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalableTargetsOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalableTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScalingActivities(input *applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalingActivitiesOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalingActivities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScalingActivitiesPages(input *applicationautoscaling.DescribeScalingActivitiesInput, fn func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalingActivitiesOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalingActivitiesOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalingActivitiesOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalingActivitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalingActivitiesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, fn func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalingActivitiesOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalingActivitiesOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalingActivitiesOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalingActivitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalingActivitiesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingActivitiesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalingActivitiesOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalingActivitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScalingPolicies(input *applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalingPoliciesOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalingPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScalingPoliciesPages(input *applicationautoscaling.DescribeScalingPoliciesInput, fn func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalingPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalingPoliciesOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalingPoliciesOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalingPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalingPoliciesPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, fn func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScalingPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScalingPoliciesOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScalingPoliciesOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScalingPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScalingPoliciesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScalingPoliciesInput, opts ...request.Option) (*applicationautoscaling.DescribeScalingPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScalingPoliciesOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScalingPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScheduledActions(input *applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScheduledActionsOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScheduledActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ApplicationAutoScaling) DescribeScheduledActionsPages(input *applicationautoscaling.DescribeScheduledActionsInput, fn func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScheduledActionsOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScheduledActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScheduledActionsPagesWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, fn func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*applicationautoscaling.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &applicationautoscaling.DescribeScheduledActionsOutput{}
	fnCacher := func(out *applicationautoscaling.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.ApplicationAutoScalingAPI.DescribeScheduledActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ApplicationAutoScaling) DescribeScheduledActionsWithContext(ctx context.Context, input *applicationautoscaling.DescribeScheduledActionsInput, opts ...request.Option) (*applicationautoscaling.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*applicationautoscaling.DescribeScheduledActionsOutput), nil
	}
	out, err := c.ApplicationAutoScalingAPI.DescribeScheduledActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
