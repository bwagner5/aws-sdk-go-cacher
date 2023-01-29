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
	"github.com/aws/aws-sdk-go/service/computeoptimizer"
	"github.com/aws/aws-sdk-go/service/computeoptimizer/computeoptimizeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ComputeOptimizer struct {
	computeoptimizeriface.ComputeOptimizerAPI
	cache *cache.Cache
}

func NewComputeOptimizer(computeoptimizerapi computeoptimizeriface.ComputeOptimizerAPI) *ComputeOptimizer {
	return &ComputeOptimizer{
		ComputeOptimizerAPI: computeoptimizerapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ComputeOptimizer) DescribeRecommendationExportJobs(input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.DescribeRecommendationExportJobsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.DescribeRecommendationExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) DescribeRecommendationExportJobsPages(input *computeoptimizer.DescribeRecommendationExportJobsInput, fn func(*computeoptimizer.DescribeRecommendationExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.DescribeRecommendationExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.DescribeRecommendationExportJobsOutput{}
	fnCacher := func(out *computeoptimizer.DescribeRecommendationExportJobsOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.DescribeRecommendationExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) DescribeRecommendationExportJobsPagesWithContext(ctx context.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput, fn func(*computeoptimizer.DescribeRecommendationExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.DescribeRecommendationExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.DescribeRecommendationExportJobsOutput{}
	fnCacher := func(out *computeoptimizer.DescribeRecommendationExportJobsOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.DescribeRecommendationExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) DescribeRecommendationExportJobsWithContext(ctx context.Context, input *computeoptimizer.DescribeRecommendationExportJobsInput, opts ...request.Option) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.DescribeRecommendationExportJobsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.DescribeRecommendationExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetAutoScalingGroupRecommendations(input *computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetAutoScalingGroupRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetAutoScalingGroupRecommendationsWithContext(ctx context.Context, input *computeoptimizer.GetAutoScalingGroupRecommendationsInput, opts ...request.Option) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetAutoScalingGroupRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetAutoScalingGroupRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEBSVolumeRecommendations(input *computeoptimizer.GetEBSVolumeRecommendationsInput) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEBSVolumeRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEBSVolumeRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEBSVolumeRecommendationsWithContext(ctx context.Context, input *computeoptimizer.GetEBSVolumeRecommendationsInput, opts ...request.Option) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEBSVolumeRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEBSVolumeRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEC2InstanceRecommendations(input *computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEC2InstanceRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEC2InstanceRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEC2InstanceRecommendationsWithContext(ctx context.Context, input *computeoptimizer.GetEC2InstanceRecommendationsInput, opts ...request.Option) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEC2InstanceRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEC2InstanceRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEC2RecommendationProjectedMetrics(input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEC2RecommendationProjectedMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEC2RecommendationProjectedMetricsWithContext(ctx context.Context, input *computeoptimizer.GetEC2RecommendationProjectedMetricsInput, opts ...request.Option) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEC2RecommendationProjectedMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetECSServiceRecommendationProjectedMetrics(input *computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput) (*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetECSServiceRecommendationProjectedMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetECSServiceRecommendationProjectedMetricsWithContext(ctx context.Context, input *computeoptimizer.GetECSServiceRecommendationProjectedMetricsInput, opts ...request.Option) (*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetECSServiceRecommendationProjectedMetricsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetECSServiceRecommendationProjectedMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetECSServiceRecommendations(input *computeoptimizer.GetECSServiceRecommendationsInput) (*computeoptimizer.GetECSServiceRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetECSServiceRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetECSServiceRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetECSServiceRecommendationsWithContext(ctx context.Context, input *computeoptimizer.GetECSServiceRecommendationsInput, opts ...request.Option) (*computeoptimizer.GetECSServiceRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetECSServiceRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetECSServiceRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEffectiveRecommendationPreferences(input *computeoptimizer.GetEffectiveRecommendationPreferencesInput) (*computeoptimizer.GetEffectiveRecommendationPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEffectiveRecommendationPreferencesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEffectiveRecommendationPreferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEffectiveRecommendationPreferencesWithContext(ctx context.Context, input *computeoptimizer.GetEffectiveRecommendationPreferencesInput, opts ...request.Option) (*computeoptimizer.GetEffectiveRecommendationPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEffectiveRecommendationPreferencesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEffectiveRecommendationPreferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEnrollmentStatus(input *computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEnrollmentStatusOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEnrollmentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEnrollmentStatusWithContext(ctx context.Context, input *computeoptimizer.GetEnrollmentStatusInput, opts ...request.Option) (*computeoptimizer.GetEnrollmentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEnrollmentStatusOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEnrollmentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEnrollmentStatusesForOrganization(input *computeoptimizer.GetEnrollmentStatusesForOrganizationInput) (*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEnrollmentStatusesForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetEnrollmentStatusesForOrganizationPages(input *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, fn func(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetEnrollmentStatusesForOrganizationOutput{}
	fnCacher := func(out *computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetEnrollmentStatusesForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetEnrollmentStatusesForOrganizationPagesWithContext(ctx context.Context, input *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, fn func(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetEnrollmentStatusesForOrganizationOutput{}
	fnCacher := func(out *computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetEnrollmentStatusesForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetEnrollmentStatusesForOrganizationWithContext(ctx context.Context, input *computeoptimizer.GetEnrollmentStatusesForOrganizationInput, opts ...request.Option) (*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetEnrollmentStatusesForOrganizationOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetEnrollmentStatusesForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetLambdaFunctionRecommendations(input *computeoptimizer.GetLambdaFunctionRecommendationsInput) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetLambdaFunctionRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetLambdaFunctionRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetLambdaFunctionRecommendationsPages(input *computeoptimizer.GetLambdaFunctionRecommendationsInput, fn func(*computeoptimizer.GetLambdaFunctionRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetLambdaFunctionRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetLambdaFunctionRecommendationsOutput{}
	fnCacher := func(out *computeoptimizer.GetLambdaFunctionRecommendationsOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetLambdaFunctionRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetLambdaFunctionRecommendationsPagesWithContext(ctx context.Context, input *computeoptimizer.GetLambdaFunctionRecommendationsInput, fn func(*computeoptimizer.GetLambdaFunctionRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetLambdaFunctionRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetLambdaFunctionRecommendationsOutput{}
	fnCacher := func(out *computeoptimizer.GetLambdaFunctionRecommendationsOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetLambdaFunctionRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetLambdaFunctionRecommendationsWithContext(ctx context.Context, input *computeoptimizer.GetLambdaFunctionRecommendationsInput, opts ...request.Option) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetLambdaFunctionRecommendationsOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetLambdaFunctionRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetRecommendationPreferences(input *computeoptimizer.GetRecommendationPreferencesInput) (*computeoptimizer.GetRecommendationPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetRecommendationPreferencesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetRecommendationPreferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetRecommendationPreferencesPages(input *computeoptimizer.GetRecommendationPreferencesInput, fn func(*computeoptimizer.GetRecommendationPreferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetRecommendationPreferencesOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetRecommendationPreferencesOutput{}
	fnCacher := func(out *computeoptimizer.GetRecommendationPreferencesOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetRecommendationPreferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetRecommendationPreferencesPagesWithContext(ctx context.Context, input *computeoptimizer.GetRecommendationPreferencesInput, fn func(*computeoptimizer.GetRecommendationPreferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetRecommendationPreferencesOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetRecommendationPreferencesOutput{}
	fnCacher := func(out *computeoptimizer.GetRecommendationPreferencesOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetRecommendationPreferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetRecommendationPreferencesWithContext(ctx context.Context, input *computeoptimizer.GetRecommendationPreferencesInput, opts ...request.Option) (*computeoptimizer.GetRecommendationPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetRecommendationPreferencesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetRecommendationPreferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetRecommendationSummaries(input *computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetRecommendationSummariesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetRecommendationSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComputeOptimizer) GetRecommendationSummariesPages(input *computeoptimizer.GetRecommendationSummariesInput, fn func(*computeoptimizer.GetRecommendationSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetRecommendationSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetRecommendationSummariesOutput{}
	fnCacher := func(out *computeoptimizer.GetRecommendationSummariesOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetRecommendationSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetRecommendationSummariesPagesWithContext(ctx context.Context, input *computeoptimizer.GetRecommendationSummariesInput, fn func(*computeoptimizer.GetRecommendationSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*computeoptimizer.GetRecommendationSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &computeoptimizer.GetRecommendationSummariesOutput{}
	fnCacher := func(out *computeoptimizer.GetRecommendationSummariesOutput, more bool) bool {
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
	if err := c.ComputeOptimizerAPI.GetRecommendationSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ComputeOptimizer) GetRecommendationSummariesWithContext(ctx context.Context, input *computeoptimizer.GetRecommendationSummariesInput, opts ...request.Option) (*computeoptimizer.GetRecommendationSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*computeoptimizer.GetRecommendationSummariesOutput), nil
	}
	out, err := c.ComputeOptimizerAPI.GetRecommendationSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
