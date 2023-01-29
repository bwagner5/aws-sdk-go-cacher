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
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SWF struct {
	swfiface.SWFAPI
	cache *cache.Cache
}

func NewSWF(swfapi swfiface.SWFAPI) *SWF {
	return &SWF{
		SWFAPI: swfapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SWF) DescribeActivityType(input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeActivityTypeOutput), nil
	}
	out, err := c.SWFAPI.DescribeActivityType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeActivityTypeWithContext(ctx context.Context, input *swf.DescribeActivityTypeInput, opts ...request.Option) (*swf.DescribeActivityTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeActivityTypeOutput), nil
	}
	out, err := c.SWFAPI.DescribeActivityTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeDomain(input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeDomainOutput), nil
	}
	out, err := c.SWFAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeDomainWithContext(ctx context.Context, input *swf.DescribeDomainInput, opts ...request.Option) (*swf.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeDomainOutput), nil
	}
	out, err := c.SWFAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeWorkflowExecution(input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeWorkflowExecutionOutput), nil
	}
	out, err := c.SWFAPI.DescribeWorkflowExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeWorkflowExecutionWithContext(ctx context.Context, input *swf.DescribeWorkflowExecutionInput, opts ...request.Option) (*swf.DescribeWorkflowExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeWorkflowExecutionOutput), nil
	}
	out, err := c.SWFAPI.DescribeWorkflowExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeWorkflowType(input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeWorkflowTypeOutput), nil
	}
	out, err := c.SWFAPI.DescribeWorkflowType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) DescribeWorkflowTypeWithContext(ctx context.Context, input *swf.DescribeWorkflowTypeInput, opts ...request.Option) (*swf.DescribeWorkflowTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.DescribeWorkflowTypeOutput), nil
	}
	out, err := c.SWFAPI.DescribeWorkflowTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) GetWorkflowExecutionHistory(input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.GetWorkflowExecutionHistoryOutput), nil
	}
	out, err := c.SWFAPI.GetWorkflowExecutionHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) GetWorkflowExecutionHistoryPages(input *swf.GetWorkflowExecutionHistoryInput, fn func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.GetWorkflowExecutionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &swf.GetWorkflowExecutionHistoryOutput{}
	fnCacher := func(out *swf.GetWorkflowExecutionHistoryOutput, more bool) bool {
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
	if err := c.SWFAPI.GetWorkflowExecutionHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) GetWorkflowExecutionHistoryPagesWithContext(ctx context.Context, input *swf.GetWorkflowExecutionHistoryInput, fn func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.GetWorkflowExecutionHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &swf.GetWorkflowExecutionHistoryOutput{}
	fnCacher := func(out *swf.GetWorkflowExecutionHistoryOutput, more bool) bool {
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
	if err := c.SWFAPI.GetWorkflowExecutionHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) GetWorkflowExecutionHistoryWithContext(ctx context.Context, input *swf.GetWorkflowExecutionHistoryInput, opts ...request.Option) (*swf.GetWorkflowExecutionHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.GetWorkflowExecutionHistoryOutput), nil
	}
	out, err := c.SWFAPI.GetWorkflowExecutionHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListActivityTypes(input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListActivityTypesOutput), nil
	}
	out, err := c.SWFAPI.ListActivityTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListActivityTypesPages(input *swf.ListActivityTypesInput, fn func(*swf.ListActivityTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListActivityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListActivityTypesOutput{}
	fnCacher := func(out *swf.ListActivityTypesOutput, more bool) bool {
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
	if err := c.SWFAPI.ListActivityTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListActivityTypesPagesWithContext(ctx context.Context, input *swf.ListActivityTypesInput, fn func(*swf.ListActivityTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListActivityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListActivityTypesOutput{}
	fnCacher := func(out *swf.ListActivityTypesOutput, more bool) bool {
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
	if err := c.SWFAPI.ListActivityTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListActivityTypesWithContext(ctx context.Context, input *swf.ListActivityTypesInput, opts ...request.Option) (*swf.ListActivityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListActivityTypesOutput), nil
	}
	out, err := c.SWFAPI.ListActivityTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListClosedWorkflowExecutionsPages(input *swf.ListClosedWorkflowExecutionsInput, fn func(*swf.WorkflowExecutionInfos, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.WorkflowExecutionInfos), false)
		return nil
	}
	cachable := true
	output := &swf.WorkflowExecutionInfos{}
	fnCacher := func(out *swf.WorkflowExecutionInfos, more bool) bool {
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
	if err := c.SWFAPI.ListClosedWorkflowExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListClosedWorkflowExecutionsPagesWithContext(ctx context.Context, input *swf.ListClosedWorkflowExecutionsInput, fn func(*swf.WorkflowExecutionInfos, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.WorkflowExecutionInfos), false)
		return nil
	}
	cachable := true
	output := &swf.WorkflowExecutionInfos{}
	fnCacher := func(out *swf.WorkflowExecutionInfos, more bool) bool {
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
	if err := c.SWFAPI.ListClosedWorkflowExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListDomains(input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListDomainsOutput), nil
	}
	out, err := c.SWFAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListDomainsPages(input *swf.ListDomainsInput, fn func(*swf.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListDomainsOutput{}
	fnCacher := func(out *swf.ListDomainsOutput, more bool) bool {
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
	if err := c.SWFAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListDomainsPagesWithContext(ctx context.Context, input *swf.ListDomainsInput, fn func(*swf.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListDomainsOutput{}
	fnCacher := func(out *swf.ListDomainsOutput, more bool) bool {
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
	if err := c.SWFAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListDomainsWithContext(ctx context.Context, input *swf.ListDomainsInput, opts ...request.Option) (*swf.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListDomainsOutput), nil
	}
	out, err := c.SWFAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListOpenWorkflowExecutionsPages(input *swf.ListOpenWorkflowExecutionsInput, fn func(*swf.WorkflowExecutionInfos, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.WorkflowExecutionInfos), false)
		return nil
	}
	cachable := true
	output := &swf.WorkflowExecutionInfos{}
	fnCacher := func(out *swf.WorkflowExecutionInfos, more bool) bool {
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
	if err := c.SWFAPI.ListOpenWorkflowExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListOpenWorkflowExecutionsPagesWithContext(ctx context.Context, input *swf.ListOpenWorkflowExecutionsInput, fn func(*swf.WorkflowExecutionInfos, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.WorkflowExecutionInfos), false)
		return nil
	}
	cachable := true
	output := &swf.WorkflowExecutionInfos{}
	fnCacher := func(out *swf.WorkflowExecutionInfos, more bool) bool {
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
	if err := c.SWFAPI.ListOpenWorkflowExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListTagsForResource(input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListTagsForResourceOutput), nil
	}
	out, err := c.SWFAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListTagsForResourceWithContext(ctx context.Context, input *swf.ListTagsForResourceInput, opts ...request.Option) (*swf.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListTagsForResourceOutput), nil
	}
	out, err := c.SWFAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListWorkflowTypes(input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListWorkflowTypesOutput), nil
	}
	out, err := c.SWFAPI.ListWorkflowTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SWF) ListWorkflowTypesPages(input *swf.ListWorkflowTypesInput, fn func(*swf.ListWorkflowTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListWorkflowTypesOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListWorkflowTypesOutput{}
	fnCacher := func(out *swf.ListWorkflowTypesOutput, more bool) bool {
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
	if err := c.SWFAPI.ListWorkflowTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListWorkflowTypesPagesWithContext(ctx context.Context, input *swf.ListWorkflowTypesInput, fn func(*swf.ListWorkflowTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*swf.ListWorkflowTypesOutput), false)
		return nil
	}
	cachable := true
	output := &swf.ListWorkflowTypesOutput{}
	fnCacher := func(out *swf.ListWorkflowTypesOutput, more bool) bool {
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
	if err := c.SWFAPI.ListWorkflowTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SWF) ListWorkflowTypesWithContext(ctx context.Context, input *swf.ListWorkflowTypesInput, opts ...request.Option) (*swf.ListWorkflowTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*swf.ListWorkflowTypesOutput), nil
	}
	out, err := c.SWFAPI.ListWorkflowTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
