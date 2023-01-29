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
	"github.com/aws/aws-sdk-go/service/resiliencehub"
	"github.com/aws/aws-sdk-go/service/resiliencehub/resiliencehubiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ResilienceHub struct {
	resiliencehubiface.ResilienceHubAPI
	cache *cache.Cache
}

func NewResilienceHub(resiliencehubapi resiliencehubiface.ResilienceHubAPI) *ResilienceHub {
	return &ResilienceHub{
		ResilienceHubAPI: resiliencehubapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ResilienceHub) DescribeApp(input *resiliencehub.DescribeAppInput) (*resiliencehub.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeApp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppAssessment(input *resiliencehub.DescribeAppAssessmentInput) (*resiliencehub.DescribeAppAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppAssessmentOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppAssessmentWithContext(ctx context.Context, input *resiliencehub.DescribeAppAssessmentInput, opts ...request.Option) (*resiliencehub.DescribeAppAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppAssessmentOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppVersionResourcesResolutionStatus(input *resiliencehub.DescribeAppVersionResourcesResolutionStatusInput) (*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppVersionResourcesResolutionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppVersionResourcesResolutionStatusWithContext(ctx context.Context, input *resiliencehub.DescribeAppVersionResourcesResolutionStatusInput, opts ...request.Option) (*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppVersionResourcesResolutionStatusOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppVersionResourcesResolutionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppVersionTemplate(input *resiliencehub.DescribeAppVersionTemplateInput) (*resiliencehub.DescribeAppVersionTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppVersionTemplateOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppVersionTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppVersionTemplateWithContext(ctx context.Context, input *resiliencehub.DescribeAppVersionTemplateInput, opts ...request.Option) (*resiliencehub.DescribeAppVersionTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppVersionTemplateOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppVersionTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeAppWithContext(ctx context.Context, input *resiliencehub.DescribeAppInput, opts ...request.Option) (*resiliencehub.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeAppOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeAppWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeDraftAppVersionResourcesImportStatus(input *resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput) (*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeDraftAppVersionResourcesImportStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeDraftAppVersionResourcesImportStatusWithContext(ctx context.Context, input *resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput, opts ...request.Option) (*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeDraftAppVersionResourcesImportStatusOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeDraftAppVersionResourcesImportStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeResiliencyPolicy(input *resiliencehub.DescribeResiliencyPolicyInput) (*resiliencehub.DescribeResiliencyPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeResiliencyPolicyOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeResiliencyPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) DescribeResiliencyPolicyWithContext(ctx context.Context, input *resiliencehub.DescribeResiliencyPolicyInput, opts ...request.Option) (*resiliencehub.DescribeResiliencyPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.DescribeResiliencyPolicyOutput), nil
	}
	out, err := c.ResilienceHubAPI.DescribeResiliencyPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAlarmRecommendations(input *resiliencehub.ListAlarmRecommendationsInput) (*resiliencehub.ListAlarmRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAlarmRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAlarmRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAlarmRecommendationsPages(input *resiliencehub.ListAlarmRecommendationsInput, fn func(*resiliencehub.ListAlarmRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAlarmRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAlarmRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListAlarmRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAlarmRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAlarmRecommendationsPagesWithContext(ctx context.Context, input *resiliencehub.ListAlarmRecommendationsInput, fn func(*resiliencehub.ListAlarmRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAlarmRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAlarmRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListAlarmRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAlarmRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAlarmRecommendationsWithContext(ctx context.Context, input *resiliencehub.ListAlarmRecommendationsInput, opts ...request.Option) (*resiliencehub.ListAlarmRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAlarmRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAlarmRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppAssessments(input *resiliencehub.ListAppAssessmentsInput) (*resiliencehub.ListAppAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppAssessmentsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppAssessments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppAssessmentsPages(input *resiliencehub.ListAppAssessmentsInput, fn func(*resiliencehub.ListAppAssessmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppAssessmentsOutput{}
	fnCacher := func(out *resiliencehub.ListAppAssessmentsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppAssessmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppAssessmentsPagesWithContext(ctx context.Context, input *resiliencehub.ListAppAssessmentsInput, fn func(*resiliencehub.ListAppAssessmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppAssessmentsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppAssessmentsOutput{}
	fnCacher := func(out *resiliencehub.ListAppAssessmentsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppAssessmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppAssessmentsWithContext(ctx context.Context, input *resiliencehub.ListAppAssessmentsInput, opts ...request.Option) (*resiliencehub.ListAppAssessmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppAssessmentsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppAssessmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppComponentCompliances(input *resiliencehub.ListAppComponentCompliancesInput) (*resiliencehub.ListAppComponentCompliancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppComponentCompliancesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppComponentCompliances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppComponentCompliancesPages(input *resiliencehub.ListAppComponentCompliancesInput, fn func(*resiliencehub.ListAppComponentCompliancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppComponentCompliancesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppComponentCompliancesOutput{}
	fnCacher := func(out *resiliencehub.ListAppComponentCompliancesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppComponentCompliancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppComponentCompliancesPagesWithContext(ctx context.Context, input *resiliencehub.ListAppComponentCompliancesInput, fn func(*resiliencehub.ListAppComponentCompliancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppComponentCompliancesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppComponentCompliancesOutput{}
	fnCacher := func(out *resiliencehub.ListAppComponentCompliancesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppComponentCompliancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppComponentCompliancesWithContext(ctx context.Context, input *resiliencehub.ListAppComponentCompliancesInput, opts ...request.Option) (*resiliencehub.ListAppComponentCompliancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppComponentCompliancesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppComponentCompliancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppComponentRecommendations(input *resiliencehub.ListAppComponentRecommendationsInput) (*resiliencehub.ListAppComponentRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppComponentRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppComponentRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppComponentRecommendationsPages(input *resiliencehub.ListAppComponentRecommendationsInput, fn func(*resiliencehub.ListAppComponentRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppComponentRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppComponentRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListAppComponentRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppComponentRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppComponentRecommendationsPagesWithContext(ctx context.Context, input *resiliencehub.ListAppComponentRecommendationsInput, fn func(*resiliencehub.ListAppComponentRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppComponentRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppComponentRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListAppComponentRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppComponentRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppComponentRecommendationsWithContext(ctx context.Context, input *resiliencehub.ListAppComponentRecommendationsInput, opts ...request.Option) (*resiliencehub.ListAppComponentRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppComponentRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppComponentRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersionResourceMappings(input *resiliencehub.ListAppVersionResourceMappingsInput) (*resiliencehub.ListAppVersionResourceMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionResourceMappingsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersionResourceMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersionResourceMappingsPages(input *resiliencehub.ListAppVersionResourceMappingsInput, fn func(*resiliencehub.ListAppVersionResourceMappingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionResourceMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionResourceMappingsOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionResourceMappingsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionResourceMappingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionResourceMappingsPagesWithContext(ctx context.Context, input *resiliencehub.ListAppVersionResourceMappingsInput, fn func(*resiliencehub.ListAppVersionResourceMappingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionResourceMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionResourceMappingsOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionResourceMappingsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionResourceMappingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionResourceMappingsWithContext(ctx context.Context, input *resiliencehub.ListAppVersionResourceMappingsInput, opts ...request.Option) (*resiliencehub.ListAppVersionResourceMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionResourceMappingsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersionResourceMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersionResources(input *resiliencehub.ListAppVersionResourcesInput) (*resiliencehub.ListAppVersionResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionResourcesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersionResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersionResourcesPages(input *resiliencehub.ListAppVersionResourcesInput, fn func(*resiliencehub.ListAppVersionResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionResourcesOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionResourcesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionResourcesPagesWithContext(ctx context.Context, input *resiliencehub.ListAppVersionResourcesInput, fn func(*resiliencehub.ListAppVersionResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionResourcesOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionResourcesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionResourcesWithContext(ctx context.Context, input *resiliencehub.ListAppVersionResourcesInput, opts ...request.Option) (*resiliencehub.ListAppVersionResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionResourcesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersionResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersions(input *resiliencehub.ListAppVersionsInput) (*resiliencehub.ListAppVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppVersionsPages(input *resiliencehub.ListAppVersionsInput, fn func(*resiliencehub.ListAppVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionsOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionsPagesWithContext(ctx context.Context, input *resiliencehub.ListAppVersionsInput, fn func(*resiliencehub.ListAppVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppVersionsOutput{}
	fnCacher := func(out *resiliencehub.ListAppVersionsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppVersionsWithContext(ctx context.Context, input *resiliencehub.ListAppVersionsInput, opts ...request.Option) (*resiliencehub.ListAppVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppVersionsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListApps(input *resiliencehub.ListAppsInput) (*resiliencehub.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListAppsPages(input *resiliencehub.ListAppsInput, fn func(*resiliencehub.ListAppsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppsOutput{}
	fnCacher := func(out *resiliencehub.ListAppsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppsPagesWithContext(ctx context.Context, input *resiliencehub.ListAppsInput, fn func(*resiliencehub.ListAppsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListAppsOutput{}
	fnCacher := func(out *resiliencehub.ListAppsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListAppsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListAppsWithContext(ctx context.Context, input *resiliencehub.ListAppsInput, opts ...request.Option) (*resiliencehub.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListAppsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListRecommendationTemplates(input *resiliencehub.ListRecommendationTemplatesInput) (*resiliencehub.ListRecommendationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListRecommendationTemplatesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListRecommendationTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListRecommendationTemplatesPages(input *resiliencehub.ListRecommendationTemplatesInput, fn func(*resiliencehub.ListRecommendationTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListRecommendationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListRecommendationTemplatesOutput{}
	fnCacher := func(out *resiliencehub.ListRecommendationTemplatesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListRecommendationTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListRecommendationTemplatesPagesWithContext(ctx context.Context, input *resiliencehub.ListRecommendationTemplatesInput, fn func(*resiliencehub.ListRecommendationTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListRecommendationTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListRecommendationTemplatesOutput{}
	fnCacher := func(out *resiliencehub.ListRecommendationTemplatesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListRecommendationTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListRecommendationTemplatesWithContext(ctx context.Context, input *resiliencehub.ListRecommendationTemplatesInput, opts ...request.Option) (*resiliencehub.ListRecommendationTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListRecommendationTemplatesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListRecommendationTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListResiliencyPolicies(input *resiliencehub.ListResiliencyPoliciesInput) (*resiliencehub.ListResiliencyPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListResiliencyPoliciesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListResiliencyPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListResiliencyPoliciesPages(input *resiliencehub.ListResiliencyPoliciesInput, fn func(*resiliencehub.ListResiliencyPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListResiliencyPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListResiliencyPoliciesOutput{}
	fnCacher := func(out *resiliencehub.ListResiliencyPoliciesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListResiliencyPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListResiliencyPoliciesPagesWithContext(ctx context.Context, input *resiliencehub.ListResiliencyPoliciesInput, fn func(*resiliencehub.ListResiliencyPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListResiliencyPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListResiliencyPoliciesOutput{}
	fnCacher := func(out *resiliencehub.ListResiliencyPoliciesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListResiliencyPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListResiliencyPoliciesWithContext(ctx context.Context, input *resiliencehub.ListResiliencyPoliciesInput, opts ...request.Option) (*resiliencehub.ListResiliencyPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListResiliencyPoliciesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListResiliencyPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListSopRecommendations(input *resiliencehub.ListSopRecommendationsInput) (*resiliencehub.ListSopRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListSopRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListSopRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListSopRecommendationsPages(input *resiliencehub.ListSopRecommendationsInput, fn func(*resiliencehub.ListSopRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListSopRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListSopRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListSopRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListSopRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListSopRecommendationsPagesWithContext(ctx context.Context, input *resiliencehub.ListSopRecommendationsInput, fn func(*resiliencehub.ListSopRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListSopRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListSopRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListSopRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListSopRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListSopRecommendationsWithContext(ctx context.Context, input *resiliencehub.ListSopRecommendationsInput, opts ...request.Option) (*resiliencehub.ListSopRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListSopRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListSopRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListSuggestedResiliencyPolicies(input *resiliencehub.ListSuggestedResiliencyPoliciesInput) (*resiliencehub.ListSuggestedResiliencyPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListSuggestedResiliencyPoliciesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListSuggestedResiliencyPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListSuggestedResiliencyPoliciesPages(input *resiliencehub.ListSuggestedResiliencyPoliciesInput, fn func(*resiliencehub.ListSuggestedResiliencyPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListSuggestedResiliencyPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListSuggestedResiliencyPoliciesOutput{}
	fnCacher := func(out *resiliencehub.ListSuggestedResiliencyPoliciesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListSuggestedResiliencyPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListSuggestedResiliencyPoliciesPagesWithContext(ctx context.Context, input *resiliencehub.ListSuggestedResiliencyPoliciesInput, fn func(*resiliencehub.ListSuggestedResiliencyPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListSuggestedResiliencyPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListSuggestedResiliencyPoliciesOutput{}
	fnCacher := func(out *resiliencehub.ListSuggestedResiliencyPoliciesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListSuggestedResiliencyPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListSuggestedResiliencyPoliciesWithContext(ctx context.Context, input *resiliencehub.ListSuggestedResiliencyPoliciesInput, opts ...request.Option) (*resiliencehub.ListSuggestedResiliencyPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListSuggestedResiliencyPoliciesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListSuggestedResiliencyPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListTagsForResource(input *resiliencehub.ListTagsForResourceInput) (*resiliencehub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListTagsForResourceOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListTagsForResourceWithContext(ctx context.Context, input *resiliencehub.ListTagsForResourceInput, opts ...request.Option) (*resiliencehub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListTagsForResourceOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListTestRecommendations(input *resiliencehub.ListTestRecommendationsInput) (*resiliencehub.ListTestRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListTestRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListTestRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListTestRecommendationsPages(input *resiliencehub.ListTestRecommendationsInput, fn func(*resiliencehub.ListTestRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListTestRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListTestRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListTestRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListTestRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListTestRecommendationsPagesWithContext(ctx context.Context, input *resiliencehub.ListTestRecommendationsInput, fn func(*resiliencehub.ListTestRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListTestRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListTestRecommendationsOutput{}
	fnCacher := func(out *resiliencehub.ListTestRecommendationsOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListTestRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListTestRecommendationsWithContext(ctx context.Context, input *resiliencehub.ListTestRecommendationsInput, opts ...request.Option) (*resiliencehub.ListTestRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListTestRecommendationsOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListTestRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListUnsupportedAppVersionResources(input *resiliencehub.ListUnsupportedAppVersionResourcesInput) (*resiliencehub.ListUnsupportedAppVersionResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListUnsupportedAppVersionResourcesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListUnsupportedAppVersionResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResilienceHub) ListUnsupportedAppVersionResourcesPages(input *resiliencehub.ListUnsupportedAppVersionResourcesInput, fn func(*resiliencehub.ListUnsupportedAppVersionResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListUnsupportedAppVersionResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListUnsupportedAppVersionResourcesOutput{}
	fnCacher := func(out *resiliencehub.ListUnsupportedAppVersionResourcesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListUnsupportedAppVersionResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListUnsupportedAppVersionResourcesPagesWithContext(ctx context.Context, input *resiliencehub.ListUnsupportedAppVersionResourcesInput, fn func(*resiliencehub.ListUnsupportedAppVersionResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resiliencehub.ListUnsupportedAppVersionResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resiliencehub.ListUnsupportedAppVersionResourcesOutput{}
	fnCacher := func(out *resiliencehub.ListUnsupportedAppVersionResourcesOutput, more bool) bool {
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
	if err := c.ResilienceHubAPI.ListUnsupportedAppVersionResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResilienceHub) ListUnsupportedAppVersionResourcesWithContext(ctx context.Context, input *resiliencehub.ListUnsupportedAppVersionResourcesInput, opts ...request.Option) (*resiliencehub.ListUnsupportedAppVersionResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resiliencehub.ListUnsupportedAppVersionResourcesOutput), nil
	}
	out, err := c.ResilienceHubAPI.ListUnsupportedAppVersionResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
