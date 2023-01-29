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
	"github.com/aws/aws-sdk-go/service/resourcegroupstagging"
	"github.com/aws/aws-sdk-go/service/resourcegroupstagging/resourcegroupstaggingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ResourceGroupsTaggingAPI struct {
	resourcegroupstaggingiface.ResourceGroupsTaggingAPIAPI
	cache *cache.Cache
}

func NewResourceGroupsTaggingAPI(resourcegroupstaggingapi resourcegroupstaggingiface.ResourceGroupsTaggingAPIAPI) *ResourceGroupsTaggingAPI {
	return &ResourceGroupsTaggingAPI{
		ResourceGroupsTaggingAPIAPI: resourcegroupstaggingapi,
		cache:                       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ResourceGroupsTaggingAPI) DescribeReportCreation(input *resourcegroupstaggingapi.DescribeReportCreationInput) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.DescribeReportCreationOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.DescribeReportCreation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) DescribeReportCreationWithContext(ctx context.Context, input *resourcegroupstaggingapi.DescribeReportCreationInput, opts ...request.Option) (*resourcegroupstaggingapi.DescribeReportCreationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.DescribeReportCreationOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.DescribeReportCreationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetComplianceSummary(input *resourcegroupstaggingapi.GetComplianceSummaryInput) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetComplianceSummaryOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetComplianceSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetComplianceSummaryPages(input *resourcegroupstaggingapi.GetComplianceSummaryInput, fn func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetComplianceSummaryOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetComplianceSummaryOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetComplianceSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetComplianceSummaryPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, fn func(*resourcegroupstaggingapi.GetComplianceSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetComplianceSummaryOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetComplianceSummaryOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetComplianceSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetComplianceSummaryWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetComplianceSummaryInput, opts ...request.Option) (*resourcegroupstaggingapi.GetComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetComplianceSummaryOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetComplianceSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetResources(input *resourcegroupstaggingapi.GetResourcesInput) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetResourcesOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetResourcesPages(input *resourcegroupstaggingapi.GetResourcesInput, fn func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetResourcesOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetResourcesOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetResourcesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, fn func(*resourcegroupstaggingapi.GetResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetResourcesOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetResourcesOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetResourcesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetResourcesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetResourcesOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetTagKeys(input *resourcegroupstaggingapi.GetTagKeysInput) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetTagKeysOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetTagKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetTagKeysPages(input *resourcegroupstaggingapi.GetTagKeysInput, fn func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetTagKeysOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetTagKeysOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetTagKeysOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetTagKeysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetTagKeysPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, fn func(*resourcegroupstaggingapi.GetTagKeysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetTagKeysOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetTagKeysOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetTagKeysOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetTagKeysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetTagKeysWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagKeysInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetTagKeysOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetTagKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetTagValues(input *resourcegroupstaggingapi.GetTagValuesInput) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetTagValuesOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetTagValues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ResourceGroupsTaggingAPI) GetTagValuesPages(input *resourcegroupstaggingapi.GetTagValuesInput, fn func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetTagValuesOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetTagValuesOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetTagValuesOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetTagValuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetTagValuesPagesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, fn func(*resourcegroupstaggingapi.GetTagValuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*resourcegroupstaggingapi.GetTagValuesOutput), false)
		return nil
	}
	cachable := true
	output := &resourcegroupstaggingapi.GetTagValuesOutput{}
	fnCacher := func(out *resourcegroupstaggingapi.GetTagValuesOutput, more bool) bool {
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
	if err := c.ResourceGroupsTaggingAPIAPI.GetTagValuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ResourceGroupsTaggingAPI) GetTagValuesWithContext(ctx context.Context, input *resourcegroupstaggingapi.GetTagValuesInput, opts ...request.Option) (*resourcegroupstaggingapi.GetTagValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*resourcegroupstaggingapi.GetTagValuesOutput), nil
	}
	out, err := c.ResourceGroupsTaggingAPIAPI.GetTagValuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
