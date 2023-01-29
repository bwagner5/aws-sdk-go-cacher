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
package codeguruprofilercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler/codeguruprofileriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeGuruProfiler struct {
	codeguruprofileriface.CodeGuruProfilerAPI
	cache *cache.Cache
}

func New(codeguruprofilerapi codeguruprofileriface.CodeGuruProfilerAPI) *CodeGuruProfiler {
	return &CodeGuruProfiler{
		CodeGuruProfilerAPI: codeguruprofilerapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeGuruProfiler) BatchGetFrameMetricData(input *codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.BatchGetFrameMetricDataOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.BatchGetFrameMetricData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) BatchGetFrameMetricDataWithContext(ctx context.Context, input *codeguruprofiler.BatchGetFrameMetricDataInput, opts ...request.Option) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.BatchGetFrameMetricDataOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.BatchGetFrameMetricDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) DescribeProfilingGroup(input *codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.DescribeProfilingGroupOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.DescribeProfilingGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) DescribeProfilingGroupWithContext(ctx context.Context, input *codeguruprofiler.DescribeProfilingGroupInput, opts ...request.Option) (*codeguruprofiler.DescribeProfilingGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.DescribeProfilingGroupOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.DescribeProfilingGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetFindingsReportAccountSummary(input *codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetFindingsReportAccountSummaryOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetFindingsReportAccountSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetFindingsReportAccountSummaryPages(input *codeguruprofiler.GetFindingsReportAccountSummaryInput, fn func(*codeguruprofiler.GetFindingsReportAccountSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.GetFindingsReportAccountSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.GetFindingsReportAccountSummaryOutput{}
	fnCacher := func(out *codeguruprofiler.GetFindingsReportAccountSummaryOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.GetFindingsReportAccountSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) GetFindingsReportAccountSummaryPagesWithContext(ctx context.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput, fn func(*codeguruprofiler.GetFindingsReportAccountSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.GetFindingsReportAccountSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.GetFindingsReportAccountSummaryOutput{}
	fnCacher := func(out *codeguruprofiler.GetFindingsReportAccountSummaryOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.GetFindingsReportAccountSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) GetFindingsReportAccountSummaryWithContext(ctx context.Context, input *codeguruprofiler.GetFindingsReportAccountSummaryInput, opts ...request.Option) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetFindingsReportAccountSummaryOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetFindingsReportAccountSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetNotificationConfiguration(input *codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetNotificationConfigurationOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetNotificationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetNotificationConfigurationWithContext(ctx context.Context, input *codeguruprofiler.GetNotificationConfigurationInput, opts ...request.Option) (*codeguruprofiler.GetNotificationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetNotificationConfigurationOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetNotificationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetPolicy(input *codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetPolicyOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetPolicyWithContext(ctx context.Context, input *codeguruprofiler.GetPolicyInput, opts ...request.Option) (*codeguruprofiler.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetPolicyOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetProfile(input *codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetProfileOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetProfileWithContext(ctx context.Context, input *codeguruprofiler.GetProfileInput, opts ...request.Option) (*codeguruprofiler.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetProfileOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetRecommendations(input *codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetRecommendationsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) GetRecommendationsWithContext(ctx context.Context, input *codeguruprofiler.GetRecommendationsInput, opts ...request.Option) (*codeguruprofiler.GetRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.GetRecommendationsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.GetRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListFindingsReports(input *codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListFindingsReportsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListFindingsReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListFindingsReportsPages(input *codeguruprofiler.ListFindingsReportsInput, fn func(*codeguruprofiler.ListFindingsReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListFindingsReportsOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListFindingsReportsOutput{}
	fnCacher := func(out *codeguruprofiler.ListFindingsReportsOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListFindingsReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListFindingsReportsPagesWithContext(ctx context.Context, input *codeguruprofiler.ListFindingsReportsInput, fn func(*codeguruprofiler.ListFindingsReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListFindingsReportsOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListFindingsReportsOutput{}
	fnCacher := func(out *codeguruprofiler.ListFindingsReportsOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListFindingsReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListFindingsReportsWithContext(ctx context.Context, input *codeguruprofiler.ListFindingsReportsInput, opts ...request.Option) (*codeguruprofiler.ListFindingsReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListFindingsReportsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListFindingsReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListProfileTimes(input *codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListProfileTimesOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListProfileTimes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListProfileTimesPages(input *codeguruprofiler.ListProfileTimesInput, fn func(*codeguruprofiler.ListProfileTimesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListProfileTimesOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListProfileTimesOutput{}
	fnCacher := func(out *codeguruprofiler.ListProfileTimesOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListProfileTimesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListProfileTimesPagesWithContext(ctx context.Context, input *codeguruprofiler.ListProfileTimesInput, fn func(*codeguruprofiler.ListProfileTimesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListProfileTimesOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListProfileTimesOutput{}
	fnCacher := func(out *codeguruprofiler.ListProfileTimesOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListProfileTimesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListProfileTimesWithContext(ctx context.Context, input *codeguruprofiler.ListProfileTimesInput, opts ...request.Option) (*codeguruprofiler.ListProfileTimesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListProfileTimesOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListProfileTimesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListProfilingGroups(input *codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListProfilingGroupsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListProfilingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListProfilingGroupsPages(input *codeguruprofiler.ListProfilingGroupsInput, fn func(*codeguruprofiler.ListProfilingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListProfilingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListProfilingGroupsOutput{}
	fnCacher := func(out *codeguruprofiler.ListProfilingGroupsOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListProfilingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListProfilingGroupsPagesWithContext(ctx context.Context, input *codeguruprofiler.ListProfilingGroupsInput, fn func(*codeguruprofiler.ListProfilingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeguruprofiler.ListProfilingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codeguruprofiler.ListProfilingGroupsOutput{}
	fnCacher := func(out *codeguruprofiler.ListProfilingGroupsOutput, more bool) bool {
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
	if err := c.CodeGuruProfilerAPI.ListProfilingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruProfiler) ListProfilingGroupsWithContext(ctx context.Context, input *codeguruprofiler.ListProfilingGroupsInput, opts ...request.Option) (*codeguruprofiler.ListProfilingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListProfilingGroupsOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListProfilingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListTagsForResource(input *codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruProfiler) ListTagsForResourceWithContext(ctx context.Context, input *codeguruprofiler.ListTagsForResourceInput, opts ...request.Option) (*codeguruprofiler.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeguruprofiler.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeGuruProfilerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
