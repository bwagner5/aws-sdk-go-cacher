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
package autoscalingcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AutoScaling struct {
	autoscalingiface.AutoScalingAPI
	cache *cache.Cache
}

func New(autoscalingapi autoscalingiface.AutoScalingAPI) *AutoScaling {
	return &AutoScaling{
		AutoScalingAPI: autoscalingapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AutoScaling) BatchDeleteScheduledAction(input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.BatchDeleteScheduledActionOutput), nil
	}
	out, err := c.AutoScalingAPI.BatchDeleteScheduledAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) BatchDeleteScheduledActionWithContext(ctx context.Context, input *autoscaling.BatchDeleteScheduledActionInput, opts ...request.Option) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.BatchDeleteScheduledActionOutput), nil
	}
	out, err := c.AutoScalingAPI.BatchDeleteScheduledActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) BatchPutScheduledUpdateGroupAction(input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.BatchPutScheduledUpdateGroupActionOutput), nil
	}
	out, err := c.AutoScalingAPI.BatchPutScheduledUpdateGroupAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) BatchPutScheduledUpdateGroupActionWithContext(ctx context.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput, opts ...request.Option) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.BatchPutScheduledUpdateGroupActionOutput), nil
	}
	out, err := c.AutoScalingAPI.BatchPutScheduledUpdateGroupActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAccountLimits(input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAccountLimitsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAccountLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAccountLimitsWithContext(ctx context.Context, input *autoscaling.DescribeAccountLimitsInput, opts ...request.Option) (*autoscaling.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAccountLimitsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAccountLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAdjustmentTypes(input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAdjustmentTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAdjustmentTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAdjustmentTypesWithContext(ctx context.Context, input *autoscaling.DescribeAdjustmentTypesInput, opts ...request.Option) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAdjustmentTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAdjustmentTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingGroups(input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingGroupsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingGroupsPages(input *autoscaling.DescribeAutoScalingGroupsInput, fn func(*autoscaling.DescribeAutoScalingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeAutoScalingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeAutoScalingGroupsOutput{}
	fnCacher := func(out *autoscaling.DescribeAutoScalingGroupsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeAutoScalingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeAutoScalingGroupsPagesWithContext(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput, fn func(*autoscaling.DescribeAutoScalingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeAutoScalingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeAutoScalingGroupsOutput{}
	fnCacher := func(out *autoscaling.DescribeAutoScalingGroupsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeAutoScalingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeAutoScalingGroupsWithContext(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput, opts ...request.Option) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingGroupsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingInstances(input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingInstancesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingInstancesPages(input *autoscaling.DescribeAutoScalingInstancesInput, fn func(*autoscaling.DescribeAutoScalingInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeAutoScalingInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeAutoScalingInstancesOutput{}
	fnCacher := func(out *autoscaling.DescribeAutoScalingInstancesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeAutoScalingInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeAutoScalingInstancesPagesWithContext(ctx context.Context, input *autoscaling.DescribeAutoScalingInstancesInput, fn func(*autoscaling.DescribeAutoScalingInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeAutoScalingInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeAutoScalingInstancesOutput{}
	fnCacher := func(out *autoscaling.DescribeAutoScalingInstancesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeAutoScalingInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeAutoScalingInstancesWithContext(ctx context.Context, input *autoscaling.DescribeAutoScalingInstancesInput, opts ...request.Option) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingInstancesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingNotificationTypes(input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingNotificationTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingNotificationTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeAutoScalingNotificationTypesWithContext(ctx context.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput, opts ...request.Option) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeAutoScalingNotificationTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeAutoScalingNotificationTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeInstanceRefreshes(input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeInstanceRefreshesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeInstanceRefreshes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeInstanceRefreshesWithContext(ctx context.Context, input *autoscaling.DescribeInstanceRefreshesInput, opts ...request.Option) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeInstanceRefreshesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeInstanceRefreshesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLaunchConfigurations(input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLaunchConfigurationsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLaunchConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLaunchConfigurationsPages(input *autoscaling.DescribeLaunchConfigurationsInput, fn func(*autoscaling.DescribeLaunchConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeLaunchConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeLaunchConfigurationsOutput{}
	fnCacher := func(out *autoscaling.DescribeLaunchConfigurationsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeLaunchConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeLaunchConfigurationsPagesWithContext(ctx context.Context, input *autoscaling.DescribeLaunchConfigurationsInput, fn func(*autoscaling.DescribeLaunchConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeLaunchConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeLaunchConfigurationsOutput{}
	fnCacher := func(out *autoscaling.DescribeLaunchConfigurationsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeLaunchConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeLaunchConfigurationsWithContext(ctx context.Context, input *autoscaling.DescribeLaunchConfigurationsInput, opts ...request.Option) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLaunchConfigurationsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLaunchConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLifecycleHookTypes(input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLifecycleHookTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLifecycleHookTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLifecycleHookTypesWithContext(ctx context.Context, input *autoscaling.DescribeLifecycleHookTypesInput, opts ...request.Option) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLifecycleHookTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLifecycleHookTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLifecycleHooks(input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLifecycleHooksOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLifecycleHooks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLifecycleHooksWithContext(ctx context.Context, input *autoscaling.DescribeLifecycleHooksInput, opts ...request.Option) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLifecycleHooksOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLifecycleHooksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLoadBalancerTargetGroups(input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLoadBalancerTargetGroupsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLoadBalancerTargetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLoadBalancerTargetGroupsWithContext(ctx context.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput, opts ...request.Option) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLoadBalancerTargetGroupsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLoadBalancerTargetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLoadBalancers(input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLoadBalancersOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeLoadBalancersWithContext(ctx context.Context, input *autoscaling.DescribeLoadBalancersInput, opts ...request.Option) (*autoscaling.DescribeLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeLoadBalancersOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeLoadBalancersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeMetricCollectionTypes(input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeMetricCollectionTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeMetricCollectionTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeMetricCollectionTypesWithContext(ctx context.Context, input *autoscaling.DescribeMetricCollectionTypesInput, opts ...request.Option) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeMetricCollectionTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeMetricCollectionTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeNotificationConfigurations(input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeNotificationConfigurationsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeNotificationConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeNotificationConfigurationsPages(input *autoscaling.DescribeNotificationConfigurationsInput, fn func(*autoscaling.DescribeNotificationConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeNotificationConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeNotificationConfigurationsOutput{}
	fnCacher := func(out *autoscaling.DescribeNotificationConfigurationsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeNotificationConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeNotificationConfigurationsPagesWithContext(ctx context.Context, input *autoscaling.DescribeNotificationConfigurationsInput, fn func(*autoscaling.DescribeNotificationConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeNotificationConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeNotificationConfigurationsOutput{}
	fnCacher := func(out *autoscaling.DescribeNotificationConfigurationsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeNotificationConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeNotificationConfigurationsWithContext(ctx context.Context, input *autoscaling.DescribeNotificationConfigurationsInput, opts ...request.Option) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeNotificationConfigurationsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeNotificationConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribePolicies(input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribePoliciesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribePoliciesPages(input *autoscaling.DescribePoliciesInput, fn func(*autoscaling.DescribePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribePoliciesOutput{}
	fnCacher := func(out *autoscaling.DescribePoliciesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribePoliciesPagesWithContext(ctx context.Context, input *autoscaling.DescribePoliciesInput, fn func(*autoscaling.DescribePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribePoliciesOutput{}
	fnCacher := func(out *autoscaling.DescribePoliciesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribePoliciesWithContext(ctx context.Context, input *autoscaling.DescribePoliciesInput, opts ...request.Option) (*autoscaling.DescribePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribePoliciesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScalingActivities(input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScalingActivitiesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScalingActivities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScalingActivitiesPages(input *autoscaling.DescribeScalingActivitiesInput, fn func(*autoscaling.DescribeScalingActivitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeScalingActivitiesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeScalingActivitiesOutput{}
	fnCacher := func(out *autoscaling.DescribeScalingActivitiesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeScalingActivitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeScalingActivitiesPagesWithContext(ctx context.Context, input *autoscaling.DescribeScalingActivitiesInput, fn func(*autoscaling.DescribeScalingActivitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeScalingActivitiesOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeScalingActivitiesOutput{}
	fnCacher := func(out *autoscaling.DescribeScalingActivitiesOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeScalingActivitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeScalingActivitiesWithContext(ctx context.Context, input *autoscaling.DescribeScalingActivitiesInput, opts ...request.Option) (*autoscaling.DescribeScalingActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScalingActivitiesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScalingActivitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScalingProcessTypes(input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScalingProcessTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScalingProcessTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScalingProcessTypesWithContext(ctx context.Context, input *autoscaling.DescribeScalingProcessTypesInput, opts ...request.Option) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScalingProcessTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScalingProcessTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScheduledActions(input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScheduledActionsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScheduledActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeScheduledActionsPages(input *autoscaling.DescribeScheduledActionsInput, fn func(*autoscaling.DescribeScheduledActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeScheduledActionsOutput{}
	fnCacher := func(out *autoscaling.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeScheduledActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeScheduledActionsPagesWithContext(ctx context.Context, input *autoscaling.DescribeScheduledActionsInput, fn func(*autoscaling.DescribeScheduledActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeScheduledActionsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeScheduledActionsOutput{}
	fnCacher := func(out *autoscaling.DescribeScheduledActionsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeScheduledActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeScheduledActionsWithContext(ctx context.Context, input *autoscaling.DescribeScheduledActionsInput, opts ...request.Option) (*autoscaling.DescribeScheduledActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeScheduledActionsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeScheduledActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTags(input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTagsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTagsPages(input *autoscaling.DescribeTagsInput, fn func(*autoscaling.DescribeTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeTagsOutput{}
	fnCacher := func(out *autoscaling.DescribeTagsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeTagsPagesWithContext(ctx context.Context, input *autoscaling.DescribeTagsInput, fn func(*autoscaling.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*autoscaling.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &autoscaling.DescribeTagsOutput{}
	fnCacher := func(out *autoscaling.DescribeTagsOutput, more bool) bool {
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
	if err := c.AutoScalingAPI.DescribeTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AutoScaling) DescribeTagsWithContext(ctx context.Context, input *autoscaling.DescribeTagsInput, opts ...request.Option) (*autoscaling.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTagsOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTerminationPolicyTypes(input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTerminationPolicyTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTerminationPolicyTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTerminationPolicyTypesWithContext(ctx context.Context, input *autoscaling.DescribeTerminationPolicyTypesInput, opts ...request.Option) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTerminationPolicyTypesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTerminationPolicyTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTrafficSources(input *autoscaling.DescribeTrafficSourcesInput) (*autoscaling.DescribeTrafficSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTrafficSourcesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTrafficSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeTrafficSourcesWithContext(ctx context.Context, input *autoscaling.DescribeTrafficSourcesInput, opts ...request.Option) (*autoscaling.DescribeTrafficSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeTrafficSourcesOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeTrafficSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeWarmPool(input *autoscaling.DescribeWarmPoolInput) (*autoscaling.DescribeWarmPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeWarmPoolOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeWarmPool(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) DescribeWarmPoolWithContext(ctx context.Context, input *autoscaling.DescribeWarmPoolInput, opts ...request.Option) (*autoscaling.DescribeWarmPoolOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.DescribeWarmPoolOutput), nil
	}
	out, err := c.AutoScalingAPI.DescribeWarmPoolWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) GetPredictiveScalingForecast(input *autoscaling.GetPredictiveScalingForecastInput) (*autoscaling.GetPredictiveScalingForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.GetPredictiveScalingForecastOutput), nil
	}
	out, err := c.AutoScalingAPI.GetPredictiveScalingForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AutoScaling) GetPredictiveScalingForecastWithContext(ctx context.Context, input *autoscaling.GetPredictiveScalingForecastInput, opts ...request.Option) (*autoscaling.GetPredictiveScalingForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*autoscaling.GetPredictiveScalingForecastOutput), nil
	}
	out, err := c.AutoScalingAPI.GetPredictiveScalingForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
