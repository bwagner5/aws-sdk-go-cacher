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
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/aws/aws-sdk-go/service/xray/xrayiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type XRay struct {
	xrayiface.XRayAPI
	cache *cache.Cache
}

func NewXRay(xrayapi xrayiface.XRayAPI) *XRay {
	return &XRay{
		XRayAPI: xrayapi,
		cache:   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *XRay) BatchGetTraces(input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.BatchGetTracesOutput), nil
	}
	out, err := c.XRayAPI.BatchGetTraces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) BatchGetTracesPages(input *xray.BatchGetTracesInput, fn func(*xray.BatchGetTracesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.BatchGetTracesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.BatchGetTracesOutput{}
	fnCacher := func(out *xray.BatchGetTracesOutput, more bool) bool {
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
	if err := c.XRayAPI.BatchGetTracesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) BatchGetTracesPagesWithContext(ctx context.Context, input *xray.BatchGetTracesInput, fn func(*xray.BatchGetTracesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.BatchGetTracesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.BatchGetTracesOutput{}
	fnCacher := func(out *xray.BatchGetTracesOutput, more bool) bool {
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
	if err := c.XRayAPI.BatchGetTracesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) BatchGetTracesWithContext(ctx context.Context, input *xray.BatchGetTracesInput, opts ...request.Option) (*xray.BatchGetTracesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.BatchGetTracesOutput), nil
	}
	out, err := c.XRayAPI.BatchGetTracesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetEncryptionConfig(input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetEncryptionConfigOutput), nil
	}
	out, err := c.XRayAPI.GetEncryptionConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetEncryptionConfigWithContext(ctx context.Context, input *xray.GetEncryptionConfigInput, opts ...request.Option) (*xray.GetEncryptionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetEncryptionConfigOutput), nil
	}
	out, err := c.XRayAPI.GetEncryptionConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetGroup(input *xray.GetGroupInput) (*xray.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetGroupOutput), nil
	}
	out, err := c.XRayAPI.GetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetGroupWithContext(ctx context.Context, input *xray.GetGroupInput, opts ...request.Option) (*xray.GetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetGroupOutput), nil
	}
	out, err := c.XRayAPI.GetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetGroups(input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetGroupsOutput), nil
	}
	out, err := c.XRayAPI.GetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetGroupsPages(input *xray.GetGroupsInput, fn func(*xray.GetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetGroupsOutput{}
	fnCacher := func(out *xray.GetGroupsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetGroupsPagesWithContext(ctx context.Context, input *xray.GetGroupsInput, fn func(*xray.GetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetGroupsOutput{}
	fnCacher := func(out *xray.GetGroupsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetGroupsWithContext(ctx context.Context, input *xray.GetGroupsInput, opts ...request.Option) (*xray.GetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetGroupsOutput), nil
	}
	out, err := c.XRayAPI.GetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsight(input *xray.GetInsightInput) (*xray.GetInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightOutput), nil
	}
	out, err := c.XRayAPI.GetInsight(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightEvents(input *xray.GetInsightEventsInput) (*xray.GetInsightEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightEventsOutput), nil
	}
	out, err := c.XRayAPI.GetInsightEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightEventsPages(input *xray.GetInsightEventsInput, fn func(*xray.GetInsightEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetInsightEventsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetInsightEventsOutput{}
	fnCacher := func(out *xray.GetInsightEventsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetInsightEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetInsightEventsPagesWithContext(ctx context.Context, input *xray.GetInsightEventsInput, fn func(*xray.GetInsightEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetInsightEventsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetInsightEventsOutput{}
	fnCacher := func(out *xray.GetInsightEventsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetInsightEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetInsightEventsWithContext(ctx context.Context, input *xray.GetInsightEventsInput, opts ...request.Option) (*xray.GetInsightEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightEventsOutput), nil
	}
	out, err := c.XRayAPI.GetInsightEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightImpactGraph(input *xray.GetInsightImpactGraphInput) (*xray.GetInsightImpactGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightImpactGraphOutput), nil
	}
	out, err := c.XRayAPI.GetInsightImpactGraph(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightImpactGraphWithContext(ctx context.Context, input *xray.GetInsightImpactGraphInput, opts ...request.Option) (*xray.GetInsightImpactGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightImpactGraphOutput), nil
	}
	out, err := c.XRayAPI.GetInsightImpactGraphWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightSummaries(input *xray.GetInsightSummariesInput) (*xray.GetInsightSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetInsightSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightSummariesPages(input *xray.GetInsightSummariesInput, fn func(*xray.GetInsightSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetInsightSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetInsightSummariesOutput{}
	fnCacher := func(out *xray.GetInsightSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetInsightSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetInsightSummariesPagesWithContext(ctx context.Context, input *xray.GetInsightSummariesInput, fn func(*xray.GetInsightSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetInsightSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetInsightSummariesOutput{}
	fnCacher := func(out *xray.GetInsightSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetInsightSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetInsightSummariesWithContext(ctx context.Context, input *xray.GetInsightSummariesInput, opts ...request.Option) (*xray.GetInsightSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetInsightSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetInsightWithContext(ctx context.Context, input *xray.GetInsightInput, opts ...request.Option) (*xray.GetInsightOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetInsightOutput), nil
	}
	out, err := c.XRayAPI.GetInsightWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingRules(input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingRulesOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingRulesPages(input *xray.GetSamplingRulesInput, fn func(*xray.GetSamplingRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetSamplingRulesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetSamplingRulesOutput{}
	fnCacher := func(out *xray.GetSamplingRulesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetSamplingRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetSamplingRulesPagesWithContext(ctx context.Context, input *xray.GetSamplingRulesInput, fn func(*xray.GetSamplingRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetSamplingRulesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetSamplingRulesOutput{}
	fnCacher := func(out *xray.GetSamplingRulesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetSamplingRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetSamplingRulesWithContext(ctx context.Context, input *xray.GetSamplingRulesInput, opts ...request.Option) (*xray.GetSamplingRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingRulesOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingStatisticSummaries(input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingStatisticSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingStatisticSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingStatisticSummariesPages(input *xray.GetSamplingStatisticSummariesInput, fn func(*xray.GetSamplingStatisticSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetSamplingStatisticSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetSamplingStatisticSummariesOutput{}
	fnCacher := func(out *xray.GetSamplingStatisticSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetSamplingStatisticSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetSamplingStatisticSummariesPagesWithContext(ctx context.Context, input *xray.GetSamplingStatisticSummariesInput, fn func(*xray.GetSamplingStatisticSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetSamplingStatisticSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetSamplingStatisticSummariesOutput{}
	fnCacher := func(out *xray.GetSamplingStatisticSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetSamplingStatisticSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetSamplingStatisticSummariesWithContext(ctx context.Context, input *xray.GetSamplingStatisticSummariesInput, opts ...request.Option) (*xray.GetSamplingStatisticSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingStatisticSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingStatisticSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingTargets(input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingTargetsOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetSamplingTargetsWithContext(ctx context.Context, input *xray.GetSamplingTargetsInput, opts ...request.Option) (*xray.GetSamplingTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetSamplingTargetsOutput), nil
	}
	out, err := c.XRayAPI.GetSamplingTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetServiceGraph(input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetServiceGraphOutput), nil
	}
	out, err := c.XRayAPI.GetServiceGraph(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetServiceGraphPages(input *xray.GetServiceGraphInput, fn func(*xray.GetServiceGraphOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetServiceGraphOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetServiceGraphOutput{}
	fnCacher := func(out *xray.GetServiceGraphOutput, more bool) bool {
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
	if err := c.XRayAPI.GetServiceGraphPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetServiceGraphPagesWithContext(ctx context.Context, input *xray.GetServiceGraphInput, fn func(*xray.GetServiceGraphOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetServiceGraphOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetServiceGraphOutput{}
	fnCacher := func(out *xray.GetServiceGraphOutput, more bool) bool {
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
	if err := c.XRayAPI.GetServiceGraphPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetServiceGraphWithContext(ctx context.Context, input *xray.GetServiceGraphInput, opts ...request.Option) (*xray.GetServiceGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetServiceGraphOutput), nil
	}
	out, err := c.XRayAPI.GetServiceGraphWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTimeSeriesServiceStatistics(input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTimeSeriesServiceStatisticsOutput), nil
	}
	out, err := c.XRayAPI.GetTimeSeriesServiceStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTimeSeriesServiceStatisticsPages(input *xray.GetTimeSeriesServiceStatisticsInput, fn func(*xray.GetTimeSeriesServiceStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTimeSeriesServiceStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTimeSeriesServiceStatisticsOutput{}
	fnCacher := func(out *xray.GetTimeSeriesServiceStatisticsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTimeSeriesServiceStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTimeSeriesServiceStatisticsPagesWithContext(ctx context.Context, input *xray.GetTimeSeriesServiceStatisticsInput, fn func(*xray.GetTimeSeriesServiceStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTimeSeriesServiceStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTimeSeriesServiceStatisticsOutput{}
	fnCacher := func(out *xray.GetTimeSeriesServiceStatisticsOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTimeSeriesServiceStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTimeSeriesServiceStatisticsWithContext(ctx context.Context, input *xray.GetTimeSeriesServiceStatisticsInput, opts ...request.Option) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTimeSeriesServiceStatisticsOutput), nil
	}
	out, err := c.XRayAPI.GetTimeSeriesServiceStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTraceGraph(input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTraceGraphOutput), nil
	}
	out, err := c.XRayAPI.GetTraceGraph(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTraceGraphPages(input *xray.GetTraceGraphInput, fn func(*xray.GetTraceGraphOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTraceGraphOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTraceGraphOutput{}
	fnCacher := func(out *xray.GetTraceGraphOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTraceGraphPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTraceGraphPagesWithContext(ctx context.Context, input *xray.GetTraceGraphInput, fn func(*xray.GetTraceGraphOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTraceGraphOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTraceGraphOutput{}
	fnCacher := func(out *xray.GetTraceGraphOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTraceGraphPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTraceGraphWithContext(ctx context.Context, input *xray.GetTraceGraphInput, opts ...request.Option) (*xray.GetTraceGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTraceGraphOutput), nil
	}
	out, err := c.XRayAPI.GetTraceGraphWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTraceSummaries(input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTraceSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetTraceSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) GetTraceSummariesPages(input *xray.GetTraceSummariesInput, fn func(*xray.GetTraceSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTraceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTraceSummariesOutput{}
	fnCacher := func(out *xray.GetTraceSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTraceSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTraceSummariesPagesWithContext(ctx context.Context, input *xray.GetTraceSummariesInput, fn func(*xray.GetTraceSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.GetTraceSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.GetTraceSummariesOutput{}
	fnCacher := func(out *xray.GetTraceSummariesOutput, more bool) bool {
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
	if err := c.XRayAPI.GetTraceSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) GetTraceSummariesWithContext(ctx context.Context, input *xray.GetTraceSummariesInput, opts ...request.Option) (*xray.GetTraceSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.GetTraceSummariesOutput), nil
	}
	out, err := c.XRayAPI.GetTraceSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) ListResourcePolicies(input *xray.ListResourcePoliciesInput) (*xray.ListResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.ListResourcePoliciesOutput), nil
	}
	out, err := c.XRayAPI.ListResourcePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) ListResourcePoliciesPages(input *xray.ListResourcePoliciesInput, fn func(*xray.ListResourcePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.ListResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.ListResourcePoliciesOutput{}
	fnCacher := func(out *xray.ListResourcePoliciesOutput, more bool) bool {
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
	if err := c.XRayAPI.ListResourcePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) ListResourcePoliciesPagesWithContext(ctx context.Context, input *xray.ListResourcePoliciesInput, fn func(*xray.ListResourcePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.ListResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &xray.ListResourcePoliciesOutput{}
	fnCacher := func(out *xray.ListResourcePoliciesOutput, more bool) bool {
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
	if err := c.XRayAPI.ListResourcePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) ListResourcePoliciesWithContext(ctx context.Context, input *xray.ListResourcePoliciesInput, opts ...request.Option) (*xray.ListResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.ListResourcePoliciesOutput), nil
	}
	out, err := c.XRayAPI.ListResourcePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) ListTagsForResource(input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.ListTagsForResourceOutput), nil
	}
	out, err := c.XRayAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *XRay) ListTagsForResourcePages(input *xray.ListTagsForResourceInput, fn func(*xray.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &xray.ListTagsForResourceOutput{}
	fnCacher := func(out *xray.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.XRayAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) ListTagsForResourcePagesWithContext(ctx context.Context, input *xray.ListTagsForResourceInput, fn func(*xray.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*xray.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &xray.ListTagsForResourceOutput{}
	fnCacher := func(out *xray.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.XRayAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *XRay) ListTagsForResourceWithContext(ctx context.Context, input *xray.ListTagsForResourceInput, opts ...request.Option) (*xray.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*xray.ListTagsForResourceOutput), nil
	}
	out, err := c.XRayAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
