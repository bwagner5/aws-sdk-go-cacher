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
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/configservice/configserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ConfigService struct {
	configserviceiface.ConfigServiceAPI
	cache *cache.Cache
}

func NewConfigService(configserviceapi configserviceiface.ConfigServiceAPI) *ConfigService {
	return &ConfigService{
		ConfigServiceAPI: configserviceapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ConfigService) BatchGetAggregateResourceConfig(input *configservice.BatchGetAggregateResourceConfigInput) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.BatchGetAggregateResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.BatchGetAggregateResourceConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) BatchGetAggregateResourceConfigWithContext(ctx context.Context, input *configservice.BatchGetAggregateResourceConfigInput, opts ...request.Option) (*configservice.BatchGetAggregateResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.BatchGetAggregateResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.BatchGetAggregateResourceConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) BatchGetResourceConfig(input *configservice.BatchGetResourceConfigInput) (*configservice.BatchGetResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.BatchGetResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.BatchGetResourceConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) BatchGetResourceConfigWithContext(ctx context.Context, input *configservice.BatchGetResourceConfigInput, opts ...request.Option) (*configservice.BatchGetResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.BatchGetResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.BatchGetResourceConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregateComplianceByConfigRules(input *configservice.DescribeAggregateComplianceByConfigRulesInput) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregateComplianceByConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregateComplianceByConfigRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregateComplianceByConfigRulesPages(input *configservice.DescribeAggregateComplianceByConfigRulesInput, fn func(*configservice.DescribeAggregateComplianceByConfigRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregateComplianceByConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregateComplianceByConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeAggregateComplianceByConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregateComplianceByConfigRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregateComplianceByConfigRulesPagesWithContext(ctx context.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput, fn func(*configservice.DescribeAggregateComplianceByConfigRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregateComplianceByConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregateComplianceByConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeAggregateComplianceByConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregateComplianceByConfigRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregateComplianceByConfigRulesWithContext(ctx context.Context, input *configservice.DescribeAggregateComplianceByConfigRulesInput, opts ...request.Option) (*configservice.DescribeAggregateComplianceByConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregateComplianceByConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregateComplianceByConfigRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregateComplianceByConformancePacks(input *configservice.DescribeAggregateComplianceByConformancePacksInput) (*configservice.DescribeAggregateComplianceByConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregateComplianceByConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregateComplianceByConformancePacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregateComplianceByConformancePacksPages(input *configservice.DescribeAggregateComplianceByConformancePacksInput, fn func(*configservice.DescribeAggregateComplianceByConformancePacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregateComplianceByConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregateComplianceByConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeAggregateComplianceByConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregateComplianceByConformancePacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregateComplianceByConformancePacksPagesWithContext(ctx context.Context, input *configservice.DescribeAggregateComplianceByConformancePacksInput, fn func(*configservice.DescribeAggregateComplianceByConformancePacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregateComplianceByConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregateComplianceByConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeAggregateComplianceByConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregateComplianceByConformancePacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregateComplianceByConformancePacksWithContext(ctx context.Context, input *configservice.DescribeAggregateComplianceByConformancePacksInput, opts ...request.Option) (*configservice.DescribeAggregateComplianceByConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregateComplianceByConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregateComplianceByConformancePacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregationAuthorizations(input *configservice.DescribeAggregationAuthorizationsInput) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregationAuthorizationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregationAuthorizations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeAggregationAuthorizationsPages(input *configservice.DescribeAggregationAuthorizationsInput, fn func(*configservice.DescribeAggregationAuthorizationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregationAuthorizationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregationAuthorizationsOutput{}
	fnCacher := func(out *configservice.DescribeAggregationAuthorizationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregationAuthorizationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregationAuthorizationsPagesWithContext(ctx context.Context, input *configservice.DescribeAggregationAuthorizationsInput, fn func(*configservice.DescribeAggregationAuthorizationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeAggregationAuthorizationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeAggregationAuthorizationsOutput{}
	fnCacher := func(out *configservice.DescribeAggregationAuthorizationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeAggregationAuthorizationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeAggregationAuthorizationsWithContext(ctx context.Context, input *configservice.DescribeAggregationAuthorizationsInput, opts ...request.Option) (*configservice.DescribeAggregationAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeAggregationAuthorizationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeAggregationAuthorizationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeComplianceByConfigRule(input *configservice.DescribeComplianceByConfigRuleInput) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeComplianceByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeComplianceByConfigRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeComplianceByConfigRulePages(input *configservice.DescribeComplianceByConfigRuleInput, fn func(*configservice.DescribeComplianceByConfigRuleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeComplianceByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeComplianceByConfigRuleOutput{}
	fnCacher := func(out *configservice.DescribeComplianceByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeComplianceByConfigRulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeComplianceByConfigRulePagesWithContext(ctx context.Context, input *configservice.DescribeComplianceByConfigRuleInput, fn func(*configservice.DescribeComplianceByConfigRuleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeComplianceByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeComplianceByConfigRuleOutput{}
	fnCacher := func(out *configservice.DescribeComplianceByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeComplianceByConfigRulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeComplianceByConfigRuleWithContext(ctx context.Context, input *configservice.DescribeComplianceByConfigRuleInput, opts ...request.Option) (*configservice.DescribeComplianceByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeComplianceByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeComplianceByConfigRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeComplianceByResource(input *configservice.DescribeComplianceByResourceInput) (*configservice.DescribeComplianceByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeComplianceByResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeComplianceByResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeComplianceByResourcePages(input *configservice.DescribeComplianceByResourceInput, fn func(*configservice.DescribeComplianceByResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeComplianceByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeComplianceByResourceOutput{}
	fnCacher := func(out *configservice.DescribeComplianceByResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeComplianceByResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeComplianceByResourcePagesWithContext(ctx context.Context, input *configservice.DescribeComplianceByResourceInput, fn func(*configservice.DescribeComplianceByResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeComplianceByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeComplianceByResourceOutput{}
	fnCacher := func(out *configservice.DescribeComplianceByResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeComplianceByResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeComplianceByResourceWithContext(ctx context.Context, input *configservice.DescribeComplianceByResourceInput, opts ...request.Option) (*configservice.DescribeComplianceByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeComplianceByResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeComplianceByResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigRuleEvaluationStatus(input *configservice.DescribeConfigRuleEvaluationStatusInput) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigRuleEvaluationStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigRuleEvaluationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigRuleEvaluationStatusPages(input *configservice.DescribeConfigRuleEvaluationStatusInput, fn func(*configservice.DescribeConfigRuleEvaluationStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigRuleEvaluationStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigRuleEvaluationStatusOutput{}
	fnCacher := func(out *configservice.DescribeConfigRuleEvaluationStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigRuleEvaluationStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigRuleEvaluationStatusPagesWithContext(ctx context.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput, fn func(*configservice.DescribeConfigRuleEvaluationStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigRuleEvaluationStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigRuleEvaluationStatusOutput{}
	fnCacher := func(out *configservice.DescribeConfigRuleEvaluationStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigRuleEvaluationStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigRuleEvaluationStatusWithContext(ctx context.Context, input *configservice.DescribeConfigRuleEvaluationStatusInput, opts ...request.Option) (*configservice.DescribeConfigRuleEvaluationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigRuleEvaluationStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigRuleEvaluationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigRules(input *configservice.DescribeConfigRulesInput) (*configservice.DescribeConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigRulesPages(input *configservice.DescribeConfigRulesInput, fn func(*configservice.DescribeConfigRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigRulesPagesWithContext(ctx context.Context, input *configservice.DescribeConfigRulesInput, fn func(*configservice.DescribeConfigRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigRulesWithContext(ctx context.Context, input *configservice.DescribeConfigRulesInput, opts ...request.Option) (*configservice.DescribeConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationAggregatorSourcesStatus(input *configservice.DescribeConfigurationAggregatorSourcesStatusInput) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationAggregatorSourcesStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationAggregatorSourcesStatusPages(input *configservice.DescribeConfigurationAggregatorSourcesStatusInput, fn func(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigurationAggregatorSourcesStatusOutput{}
	fnCacher := func(out *configservice.DescribeConfigurationAggregatorSourcesStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigurationAggregatorSourcesStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigurationAggregatorSourcesStatusPagesWithContext(ctx context.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput, fn func(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigurationAggregatorSourcesStatusOutput{}
	fnCacher := func(out *configservice.DescribeConfigurationAggregatorSourcesStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigurationAggregatorSourcesStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigurationAggregatorSourcesStatusWithContext(ctx context.Context, input *configservice.DescribeConfigurationAggregatorSourcesStatusInput, opts ...request.Option) (*configservice.DescribeConfigurationAggregatorSourcesStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationAggregatorSourcesStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationAggregatorSourcesStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationAggregators(input *configservice.DescribeConfigurationAggregatorsInput) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationAggregatorsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationAggregators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationAggregatorsPages(input *configservice.DescribeConfigurationAggregatorsInput, fn func(*configservice.DescribeConfigurationAggregatorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigurationAggregatorsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigurationAggregatorsOutput{}
	fnCacher := func(out *configservice.DescribeConfigurationAggregatorsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigurationAggregatorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigurationAggregatorsPagesWithContext(ctx context.Context, input *configservice.DescribeConfigurationAggregatorsInput, fn func(*configservice.DescribeConfigurationAggregatorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConfigurationAggregatorsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConfigurationAggregatorsOutput{}
	fnCacher := func(out *configservice.DescribeConfigurationAggregatorsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConfigurationAggregatorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConfigurationAggregatorsWithContext(ctx context.Context, input *configservice.DescribeConfigurationAggregatorsInput, opts ...request.Option) (*configservice.DescribeConfigurationAggregatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationAggregatorsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationAggregatorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationRecorderStatus(input *configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationRecorderStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationRecorderStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationRecorderStatusWithContext(ctx context.Context, input *configservice.DescribeConfigurationRecorderStatusInput, opts ...request.Option) (*configservice.DescribeConfigurationRecorderStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationRecorderStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationRecorderStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationRecorders(input *configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationRecordersOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationRecorders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConfigurationRecordersWithContext(ctx context.Context, input *configservice.DescribeConfigurationRecordersInput, opts ...request.Option) (*configservice.DescribeConfigurationRecordersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConfigurationRecordersOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConfigurationRecordersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePackCompliance(input *configservice.DescribeConformancePackComplianceInput) (*configservice.DescribeConformancePackComplianceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePackComplianceOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePackCompliance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePackCompliancePages(input *configservice.DescribeConformancePackComplianceInput, fn func(*configservice.DescribeConformancePackComplianceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePackComplianceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePackComplianceOutput{}
	fnCacher := func(out *configservice.DescribeConformancePackComplianceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePackCompliancePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePackCompliancePagesWithContext(ctx context.Context, input *configservice.DescribeConformancePackComplianceInput, fn func(*configservice.DescribeConformancePackComplianceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePackComplianceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePackComplianceOutput{}
	fnCacher := func(out *configservice.DescribeConformancePackComplianceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePackCompliancePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePackComplianceWithContext(ctx context.Context, input *configservice.DescribeConformancePackComplianceInput, opts ...request.Option) (*configservice.DescribeConformancePackComplianceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePackComplianceOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePackComplianceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePackStatus(input *configservice.DescribeConformancePackStatusInput) (*configservice.DescribeConformancePackStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePackStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePackStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePackStatusPages(input *configservice.DescribeConformancePackStatusInput, fn func(*configservice.DescribeConformancePackStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePackStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePackStatusOutput{}
	fnCacher := func(out *configservice.DescribeConformancePackStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePackStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePackStatusPagesWithContext(ctx context.Context, input *configservice.DescribeConformancePackStatusInput, fn func(*configservice.DescribeConformancePackStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePackStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePackStatusOutput{}
	fnCacher := func(out *configservice.DescribeConformancePackStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePackStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePackStatusWithContext(ctx context.Context, input *configservice.DescribeConformancePackStatusInput, opts ...request.Option) (*configservice.DescribeConformancePackStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePackStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePackStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePacks(input *configservice.DescribeConformancePacksInput) (*configservice.DescribeConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeConformancePacksPages(input *configservice.DescribeConformancePacksInput, fn func(*configservice.DescribeConformancePacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePacksPagesWithContext(ctx context.Context, input *configservice.DescribeConformancePacksInput, fn func(*configservice.DescribeConformancePacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeConformancePacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeConformancePacksWithContext(ctx context.Context, input *configservice.DescribeConformancePacksInput, opts ...request.Option) (*configservice.DescribeConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeConformancePacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeDeliveryChannelStatus(input *configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeDeliveryChannelStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeDeliveryChannelStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeDeliveryChannelStatusWithContext(ctx context.Context, input *configservice.DescribeDeliveryChannelStatusInput, opts ...request.Option) (*configservice.DescribeDeliveryChannelStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeDeliveryChannelStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeDeliveryChannelStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeDeliveryChannels(input *configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeDeliveryChannelsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeDeliveryChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeDeliveryChannelsWithContext(ctx context.Context, input *configservice.DescribeDeliveryChannelsInput, opts ...request.Option) (*configservice.DescribeDeliveryChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeDeliveryChannelsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeDeliveryChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConfigRuleStatuses(input *configservice.DescribeOrganizationConfigRuleStatusesInput) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConfigRuleStatusesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConfigRuleStatuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConfigRuleStatusesPages(input *configservice.DescribeOrganizationConfigRuleStatusesInput, fn func(*configservice.DescribeOrganizationConfigRuleStatusesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConfigRuleStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConfigRuleStatusesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConfigRuleStatusesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConfigRuleStatusesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConfigRuleStatusesPagesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput, fn func(*configservice.DescribeOrganizationConfigRuleStatusesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConfigRuleStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConfigRuleStatusesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConfigRuleStatusesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConfigRuleStatusesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConfigRuleStatusesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConfigRuleStatusesInput, opts ...request.Option) (*configservice.DescribeOrganizationConfigRuleStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConfigRuleStatusesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConfigRuleStatusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConfigRules(input *configservice.DescribeOrganizationConfigRulesInput) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConfigRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConfigRulesPages(input *configservice.DescribeOrganizationConfigRulesInput, fn func(*configservice.DescribeOrganizationConfigRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConfigRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConfigRulesPagesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConfigRulesInput, fn func(*configservice.DescribeOrganizationConfigRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConfigRulesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConfigRulesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConfigRulesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConfigRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConfigRulesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConfigRulesInput, opts ...request.Option) (*configservice.DescribeOrganizationConfigRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConfigRulesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConfigRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConformancePackStatuses(input *configservice.DescribeOrganizationConformancePackStatusesInput) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConformancePackStatusesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConformancePackStatuses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConformancePackStatusesPages(input *configservice.DescribeOrganizationConformancePackStatusesInput, fn func(*configservice.DescribeOrganizationConformancePackStatusesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConformancePackStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConformancePackStatusesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConformancePackStatusesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConformancePackStatusesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConformancePackStatusesPagesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput, fn func(*configservice.DescribeOrganizationConformancePackStatusesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConformancePackStatusesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConformancePackStatusesOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConformancePackStatusesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConformancePackStatusesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConformancePackStatusesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConformancePackStatusesInput, opts ...request.Option) (*configservice.DescribeOrganizationConformancePackStatusesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConformancePackStatusesOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConformancePackStatusesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConformancePacks(input *configservice.DescribeOrganizationConformancePacksInput) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConformancePacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeOrganizationConformancePacksPages(input *configservice.DescribeOrganizationConformancePacksInput, fn func(*configservice.DescribeOrganizationConformancePacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConformancePacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConformancePacksPagesWithContext(ctx context.Context, input *configservice.DescribeOrganizationConformancePacksInput, fn func(*configservice.DescribeOrganizationConformancePacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeOrganizationConformancePacksOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeOrganizationConformancePacksOutput{}
	fnCacher := func(out *configservice.DescribeOrganizationConformancePacksOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeOrganizationConformancePacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeOrganizationConformancePacksWithContext(ctx context.Context, input *configservice.DescribeOrganizationConformancePacksInput, opts ...request.Option) (*configservice.DescribeOrganizationConformancePacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeOrganizationConformancePacksOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeOrganizationConformancePacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribePendingAggregationRequests(input *configservice.DescribePendingAggregationRequestsInput) (*configservice.DescribePendingAggregationRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribePendingAggregationRequestsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribePendingAggregationRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribePendingAggregationRequestsPages(input *configservice.DescribePendingAggregationRequestsInput, fn func(*configservice.DescribePendingAggregationRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribePendingAggregationRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribePendingAggregationRequestsOutput{}
	fnCacher := func(out *configservice.DescribePendingAggregationRequestsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribePendingAggregationRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribePendingAggregationRequestsPagesWithContext(ctx context.Context, input *configservice.DescribePendingAggregationRequestsInput, fn func(*configservice.DescribePendingAggregationRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribePendingAggregationRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribePendingAggregationRequestsOutput{}
	fnCacher := func(out *configservice.DescribePendingAggregationRequestsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribePendingAggregationRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribePendingAggregationRequestsWithContext(ctx context.Context, input *configservice.DescribePendingAggregationRequestsInput, opts ...request.Option) (*configservice.DescribePendingAggregationRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribePendingAggregationRequestsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribePendingAggregationRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationConfigurations(input *configservice.DescribeRemediationConfigurationsInput) (*configservice.DescribeRemediationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationConfigurationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationConfigurationsWithContext(ctx context.Context, input *configservice.DescribeRemediationConfigurationsInput, opts ...request.Option) (*configservice.DescribeRemediationConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationConfigurationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationExceptions(input *configservice.DescribeRemediationExceptionsInput) (*configservice.DescribeRemediationExceptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationExceptionsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationExceptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationExceptionsPages(input *configservice.DescribeRemediationExceptionsInput, fn func(*configservice.DescribeRemediationExceptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRemediationExceptionsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRemediationExceptionsOutput{}
	fnCacher := func(out *configservice.DescribeRemediationExceptionsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRemediationExceptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRemediationExceptionsPagesWithContext(ctx context.Context, input *configservice.DescribeRemediationExceptionsInput, fn func(*configservice.DescribeRemediationExceptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRemediationExceptionsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRemediationExceptionsOutput{}
	fnCacher := func(out *configservice.DescribeRemediationExceptionsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRemediationExceptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRemediationExceptionsWithContext(ctx context.Context, input *configservice.DescribeRemediationExceptionsInput, opts ...request.Option) (*configservice.DescribeRemediationExceptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationExceptionsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationExceptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationExecutionStatus(input *configservice.DescribeRemediationExecutionStatusInput) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationExecutionStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationExecutionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRemediationExecutionStatusPages(input *configservice.DescribeRemediationExecutionStatusInput, fn func(*configservice.DescribeRemediationExecutionStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRemediationExecutionStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRemediationExecutionStatusOutput{}
	fnCacher := func(out *configservice.DescribeRemediationExecutionStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRemediationExecutionStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRemediationExecutionStatusPagesWithContext(ctx context.Context, input *configservice.DescribeRemediationExecutionStatusInput, fn func(*configservice.DescribeRemediationExecutionStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRemediationExecutionStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRemediationExecutionStatusOutput{}
	fnCacher := func(out *configservice.DescribeRemediationExecutionStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRemediationExecutionStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRemediationExecutionStatusWithContext(ctx context.Context, input *configservice.DescribeRemediationExecutionStatusInput, opts ...request.Option) (*configservice.DescribeRemediationExecutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRemediationExecutionStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRemediationExecutionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRetentionConfigurations(input *configservice.DescribeRetentionConfigurationsInput) (*configservice.DescribeRetentionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRetentionConfigurationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRetentionConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) DescribeRetentionConfigurationsPages(input *configservice.DescribeRetentionConfigurationsInput, fn func(*configservice.DescribeRetentionConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRetentionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRetentionConfigurationsOutput{}
	fnCacher := func(out *configservice.DescribeRetentionConfigurationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRetentionConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRetentionConfigurationsPagesWithContext(ctx context.Context, input *configservice.DescribeRetentionConfigurationsInput, fn func(*configservice.DescribeRetentionConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.DescribeRetentionConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.DescribeRetentionConfigurationsOutput{}
	fnCacher := func(out *configservice.DescribeRetentionConfigurationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.DescribeRetentionConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) DescribeRetentionConfigurationsWithContext(ctx context.Context, input *configservice.DescribeRetentionConfigurationsInput, opts ...request.Option) (*configservice.DescribeRetentionConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.DescribeRetentionConfigurationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.DescribeRetentionConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateComplianceDetailsByConfigRule(input *configservice.GetAggregateComplianceDetailsByConfigRuleInput) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateComplianceDetailsByConfigRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateComplianceDetailsByConfigRulePages(input *configservice.GetAggregateComplianceDetailsByConfigRuleInput, fn func(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateComplianceDetailsByConfigRuleOutput{}
	fnCacher := func(out *configservice.GetAggregateComplianceDetailsByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateComplianceDetailsByConfigRulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateComplianceDetailsByConfigRulePagesWithContext(ctx context.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput, fn func(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateComplianceDetailsByConfigRuleOutput{}
	fnCacher := func(out *configservice.GetAggregateComplianceDetailsByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateComplianceDetailsByConfigRulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateComplianceDetailsByConfigRuleWithContext(ctx context.Context, input *configservice.GetAggregateComplianceDetailsByConfigRuleInput, opts ...request.Option) (*configservice.GetAggregateComplianceDetailsByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateComplianceDetailsByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateComplianceDetailsByConfigRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateConfigRuleComplianceSummary(input *configservice.GetAggregateConfigRuleComplianceSummaryInput) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateConfigRuleComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateConfigRuleComplianceSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateConfigRuleComplianceSummaryPages(input *configservice.GetAggregateConfigRuleComplianceSummaryInput, fn func(*configservice.GetAggregateConfigRuleComplianceSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateConfigRuleComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateConfigRuleComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetAggregateConfigRuleComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateConfigRuleComplianceSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateConfigRuleComplianceSummaryPagesWithContext(ctx context.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput, fn func(*configservice.GetAggregateConfigRuleComplianceSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateConfigRuleComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateConfigRuleComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetAggregateConfigRuleComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateConfigRuleComplianceSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateConfigRuleComplianceSummaryWithContext(ctx context.Context, input *configservice.GetAggregateConfigRuleComplianceSummaryInput, opts ...request.Option) (*configservice.GetAggregateConfigRuleComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateConfigRuleComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateConfigRuleComplianceSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateConformancePackComplianceSummary(input *configservice.GetAggregateConformancePackComplianceSummaryInput) (*configservice.GetAggregateConformancePackComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateConformancePackComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateConformancePackComplianceSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateConformancePackComplianceSummaryPages(input *configservice.GetAggregateConformancePackComplianceSummaryInput, fn func(*configservice.GetAggregateConformancePackComplianceSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateConformancePackComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateConformancePackComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetAggregateConformancePackComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateConformancePackComplianceSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateConformancePackComplianceSummaryPagesWithContext(ctx context.Context, input *configservice.GetAggregateConformancePackComplianceSummaryInput, fn func(*configservice.GetAggregateConformancePackComplianceSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateConformancePackComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateConformancePackComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetAggregateConformancePackComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateConformancePackComplianceSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateConformancePackComplianceSummaryWithContext(ctx context.Context, input *configservice.GetAggregateConformancePackComplianceSummaryInput, opts ...request.Option) (*configservice.GetAggregateConformancePackComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateConformancePackComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateConformancePackComplianceSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateDiscoveredResourceCounts(input *configservice.GetAggregateDiscoveredResourceCountsInput) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateDiscoveredResourceCountsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateDiscoveredResourceCounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateDiscoveredResourceCountsPages(input *configservice.GetAggregateDiscoveredResourceCountsInput, fn func(*configservice.GetAggregateDiscoveredResourceCountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateDiscoveredResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateDiscoveredResourceCountsOutput{}
	fnCacher := func(out *configservice.GetAggregateDiscoveredResourceCountsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateDiscoveredResourceCountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateDiscoveredResourceCountsPagesWithContext(ctx context.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput, fn func(*configservice.GetAggregateDiscoveredResourceCountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetAggregateDiscoveredResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetAggregateDiscoveredResourceCountsOutput{}
	fnCacher := func(out *configservice.GetAggregateDiscoveredResourceCountsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetAggregateDiscoveredResourceCountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetAggregateDiscoveredResourceCountsWithContext(ctx context.Context, input *configservice.GetAggregateDiscoveredResourceCountsInput, opts ...request.Option) (*configservice.GetAggregateDiscoveredResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateDiscoveredResourceCountsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateDiscoveredResourceCountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateResourceConfig(input *configservice.GetAggregateResourceConfigInput) (*configservice.GetAggregateResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateResourceConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetAggregateResourceConfigWithContext(ctx context.Context, input *configservice.GetAggregateResourceConfigInput, opts ...request.Option) (*configservice.GetAggregateResourceConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetAggregateResourceConfigOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetAggregateResourceConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceDetailsByConfigRule(input *configservice.GetComplianceDetailsByConfigRuleInput) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceDetailsByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceDetailsByConfigRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceDetailsByConfigRulePages(input *configservice.GetComplianceDetailsByConfigRuleInput, fn func(*configservice.GetComplianceDetailsByConfigRuleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetComplianceDetailsByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetComplianceDetailsByConfigRuleOutput{}
	fnCacher := func(out *configservice.GetComplianceDetailsByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetComplianceDetailsByConfigRulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetComplianceDetailsByConfigRulePagesWithContext(ctx context.Context, input *configservice.GetComplianceDetailsByConfigRuleInput, fn func(*configservice.GetComplianceDetailsByConfigRuleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetComplianceDetailsByConfigRuleOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetComplianceDetailsByConfigRuleOutput{}
	fnCacher := func(out *configservice.GetComplianceDetailsByConfigRuleOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetComplianceDetailsByConfigRulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetComplianceDetailsByConfigRuleWithContext(ctx context.Context, input *configservice.GetComplianceDetailsByConfigRuleInput, opts ...request.Option) (*configservice.GetComplianceDetailsByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceDetailsByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceDetailsByConfigRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceDetailsByResource(input *configservice.GetComplianceDetailsByResourceInput) (*configservice.GetComplianceDetailsByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceDetailsByResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceDetailsByResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceDetailsByResourcePages(input *configservice.GetComplianceDetailsByResourceInput, fn func(*configservice.GetComplianceDetailsByResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetComplianceDetailsByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetComplianceDetailsByResourceOutput{}
	fnCacher := func(out *configservice.GetComplianceDetailsByResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetComplianceDetailsByResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetComplianceDetailsByResourcePagesWithContext(ctx context.Context, input *configservice.GetComplianceDetailsByResourceInput, fn func(*configservice.GetComplianceDetailsByResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetComplianceDetailsByResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetComplianceDetailsByResourceOutput{}
	fnCacher := func(out *configservice.GetComplianceDetailsByResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetComplianceDetailsByResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetComplianceDetailsByResourceWithContext(ctx context.Context, input *configservice.GetComplianceDetailsByResourceInput, opts ...request.Option) (*configservice.GetComplianceDetailsByResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceDetailsByResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceDetailsByResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceSummaryByConfigRule(input *configservice.GetComplianceSummaryByConfigRuleInput) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceSummaryByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceSummaryByConfigRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceSummaryByConfigRuleWithContext(ctx context.Context, input *configservice.GetComplianceSummaryByConfigRuleInput, opts ...request.Option) (*configservice.GetComplianceSummaryByConfigRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceSummaryByConfigRuleOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceSummaryByConfigRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceSummaryByResourceType(input *configservice.GetComplianceSummaryByResourceTypeInput) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceSummaryByResourceTypeOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceSummaryByResourceType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetComplianceSummaryByResourceTypeWithContext(ctx context.Context, input *configservice.GetComplianceSummaryByResourceTypeInput, opts ...request.Option) (*configservice.GetComplianceSummaryByResourceTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetComplianceSummaryByResourceTypeOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetComplianceSummaryByResourceTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetConformancePackComplianceDetails(input *configservice.GetConformancePackComplianceDetailsInput) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetConformancePackComplianceDetailsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetConformancePackComplianceDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetConformancePackComplianceDetailsPages(input *configservice.GetConformancePackComplianceDetailsInput, fn func(*configservice.GetConformancePackComplianceDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetConformancePackComplianceDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetConformancePackComplianceDetailsOutput{}
	fnCacher := func(out *configservice.GetConformancePackComplianceDetailsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetConformancePackComplianceDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetConformancePackComplianceDetailsPagesWithContext(ctx context.Context, input *configservice.GetConformancePackComplianceDetailsInput, fn func(*configservice.GetConformancePackComplianceDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetConformancePackComplianceDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetConformancePackComplianceDetailsOutput{}
	fnCacher := func(out *configservice.GetConformancePackComplianceDetailsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetConformancePackComplianceDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetConformancePackComplianceDetailsWithContext(ctx context.Context, input *configservice.GetConformancePackComplianceDetailsInput, opts ...request.Option) (*configservice.GetConformancePackComplianceDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetConformancePackComplianceDetailsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetConformancePackComplianceDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetConformancePackComplianceSummary(input *configservice.GetConformancePackComplianceSummaryInput) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetConformancePackComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetConformancePackComplianceSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetConformancePackComplianceSummaryPages(input *configservice.GetConformancePackComplianceSummaryInput, fn func(*configservice.GetConformancePackComplianceSummaryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetConformancePackComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetConformancePackComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetConformancePackComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetConformancePackComplianceSummaryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetConformancePackComplianceSummaryPagesWithContext(ctx context.Context, input *configservice.GetConformancePackComplianceSummaryInput, fn func(*configservice.GetConformancePackComplianceSummaryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetConformancePackComplianceSummaryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetConformancePackComplianceSummaryOutput{}
	fnCacher := func(out *configservice.GetConformancePackComplianceSummaryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetConformancePackComplianceSummaryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetConformancePackComplianceSummaryWithContext(ctx context.Context, input *configservice.GetConformancePackComplianceSummaryInput, opts ...request.Option) (*configservice.GetConformancePackComplianceSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetConformancePackComplianceSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetConformancePackComplianceSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetCustomRulePolicy(input *configservice.GetCustomRulePolicyInput) (*configservice.GetCustomRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetCustomRulePolicyOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetCustomRulePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetCustomRulePolicyWithContext(ctx context.Context, input *configservice.GetCustomRulePolicyInput, opts ...request.Option) (*configservice.GetCustomRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetCustomRulePolicyOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetCustomRulePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetDiscoveredResourceCounts(input *configservice.GetDiscoveredResourceCountsInput) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetDiscoveredResourceCountsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetDiscoveredResourceCounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetDiscoveredResourceCountsPages(input *configservice.GetDiscoveredResourceCountsInput, fn func(*configservice.GetDiscoveredResourceCountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetDiscoveredResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetDiscoveredResourceCountsOutput{}
	fnCacher := func(out *configservice.GetDiscoveredResourceCountsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetDiscoveredResourceCountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetDiscoveredResourceCountsPagesWithContext(ctx context.Context, input *configservice.GetDiscoveredResourceCountsInput, fn func(*configservice.GetDiscoveredResourceCountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetDiscoveredResourceCountsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetDiscoveredResourceCountsOutput{}
	fnCacher := func(out *configservice.GetDiscoveredResourceCountsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetDiscoveredResourceCountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetDiscoveredResourceCountsWithContext(ctx context.Context, input *configservice.GetDiscoveredResourceCountsInput, opts ...request.Option) (*configservice.GetDiscoveredResourceCountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetDiscoveredResourceCountsOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetDiscoveredResourceCountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationConfigRuleDetailedStatus(input *configservice.GetOrganizationConfigRuleDetailedStatusInput) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationConfigRuleDetailedStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationConfigRuleDetailedStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationConfigRuleDetailedStatusPages(input *configservice.GetOrganizationConfigRuleDetailedStatusInput, fn func(*configservice.GetOrganizationConfigRuleDetailedStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetOrganizationConfigRuleDetailedStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetOrganizationConfigRuleDetailedStatusOutput{}
	fnCacher := func(out *configservice.GetOrganizationConfigRuleDetailedStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetOrganizationConfigRuleDetailedStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetOrganizationConfigRuleDetailedStatusPagesWithContext(ctx context.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput, fn func(*configservice.GetOrganizationConfigRuleDetailedStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetOrganizationConfigRuleDetailedStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetOrganizationConfigRuleDetailedStatusOutput{}
	fnCacher := func(out *configservice.GetOrganizationConfigRuleDetailedStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetOrganizationConfigRuleDetailedStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetOrganizationConfigRuleDetailedStatusWithContext(ctx context.Context, input *configservice.GetOrganizationConfigRuleDetailedStatusInput, opts ...request.Option) (*configservice.GetOrganizationConfigRuleDetailedStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationConfigRuleDetailedStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationConfigRuleDetailedStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationConformancePackDetailedStatus(input *configservice.GetOrganizationConformancePackDetailedStatusInput) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationConformancePackDetailedStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationConformancePackDetailedStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationConformancePackDetailedStatusPages(input *configservice.GetOrganizationConformancePackDetailedStatusInput, fn func(*configservice.GetOrganizationConformancePackDetailedStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetOrganizationConformancePackDetailedStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetOrganizationConformancePackDetailedStatusOutput{}
	fnCacher := func(out *configservice.GetOrganizationConformancePackDetailedStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetOrganizationConformancePackDetailedStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetOrganizationConformancePackDetailedStatusPagesWithContext(ctx context.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput, fn func(*configservice.GetOrganizationConformancePackDetailedStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetOrganizationConformancePackDetailedStatusOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetOrganizationConformancePackDetailedStatusOutput{}
	fnCacher := func(out *configservice.GetOrganizationConformancePackDetailedStatusOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetOrganizationConformancePackDetailedStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetOrganizationConformancePackDetailedStatusWithContext(ctx context.Context, input *configservice.GetOrganizationConformancePackDetailedStatusInput, opts ...request.Option) (*configservice.GetOrganizationConformancePackDetailedStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationConformancePackDetailedStatusOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationConformancePackDetailedStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationCustomRulePolicy(input *configservice.GetOrganizationCustomRulePolicyInput) (*configservice.GetOrganizationCustomRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationCustomRulePolicyOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationCustomRulePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetOrganizationCustomRulePolicyWithContext(ctx context.Context, input *configservice.GetOrganizationCustomRulePolicyInput, opts ...request.Option) (*configservice.GetOrganizationCustomRulePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetOrganizationCustomRulePolicyOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetOrganizationCustomRulePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetResourceConfigHistory(input *configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetResourceConfigHistoryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetResourceConfigHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetResourceConfigHistoryPages(input *configservice.GetResourceConfigHistoryInput, fn func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetResourceConfigHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetResourceConfigHistoryOutput{}
	fnCacher := func(out *configservice.GetResourceConfigHistoryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetResourceConfigHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetResourceConfigHistoryPagesWithContext(ctx context.Context, input *configservice.GetResourceConfigHistoryInput, fn func(*configservice.GetResourceConfigHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.GetResourceConfigHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.GetResourceConfigHistoryOutput{}
	fnCacher := func(out *configservice.GetResourceConfigHistoryOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.GetResourceConfigHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) GetResourceConfigHistoryWithContext(ctx context.Context, input *configservice.GetResourceConfigHistoryInput, opts ...request.Option) (*configservice.GetResourceConfigHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetResourceConfigHistoryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetResourceConfigHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetResourceEvaluationSummary(input *configservice.GetResourceEvaluationSummaryInput) (*configservice.GetResourceEvaluationSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetResourceEvaluationSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetResourceEvaluationSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetResourceEvaluationSummaryWithContext(ctx context.Context, input *configservice.GetResourceEvaluationSummaryInput, opts ...request.Option) (*configservice.GetResourceEvaluationSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetResourceEvaluationSummaryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetResourceEvaluationSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetStoredQuery(input *configservice.GetStoredQueryInput) (*configservice.GetStoredQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetStoredQueryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetStoredQuery(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) GetStoredQueryWithContext(ctx context.Context, input *configservice.GetStoredQueryInput, opts ...request.Option) (*configservice.GetStoredQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.GetStoredQueryOutput), nil
	}
	out, err := c.ConfigServiceAPI.GetStoredQueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListAggregateDiscoveredResources(input *configservice.ListAggregateDiscoveredResourcesInput) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListAggregateDiscoveredResourcesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListAggregateDiscoveredResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListAggregateDiscoveredResourcesPages(input *configservice.ListAggregateDiscoveredResourcesInput, fn func(*configservice.ListAggregateDiscoveredResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListAggregateDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListAggregateDiscoveredResourcesOutput{}
	fnCacher := func(out *configservice.ListAggregateDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListAggregateDiscoveredResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListAggregateDiscoveredResourcesPagesWithContext(ctx context.Context, input *configservice.ListAggregateDiscoveredResourcesInput, fn func(*configservice.ListAggregateDiscoveredResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListAggregateDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListAggregateDiscoveredResourcesOutput{}
	fnCacher := func(out *configservice.ListAggregateDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListAggregateDiscoveredResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListAggregateDiscoveredResourcesWithContext(ctx context.Context, input *configservice.ListAggregateDiscoveredResourcesInput, opts ...request.Option) (*configservice.ListAggregateDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListAggregateDiscoveredResourcesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListAggregateDiscoveredResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListConformancePackComplianceScores(input *configservice.ListConformancePackComplianceScoresInput) (*configservice.ListConformancePackComplianceScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListConformancePackComplianceScoresOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListConformancePackComplianceScores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListConformancePackComplianceScoresPages(input *configservice.ListConformancePackComplianceScoresInput, fn func(*configservice.ListConformancePackComplianceScoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListConformancePackComplianceScoresOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListConformancePackComplianceScoresOutput{}
	fnCacher := func(out *configservice.ListConformancePackComplianceScoresOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListConformancePackComplianceScoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListConformancePackComplianceScoresPagesWithContext(ctx context.Context, input *configservice.ListConformancePackComplianceScoresInput, fn func(*configservice.ListConformancePackComplianceScoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListConformancePackComplianceScoresOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListConformancePackComplianceScoresOutput{}
	fnCacher := func(out *configservice.ListConformancePackComplianceScoresOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListConformancePackComplianceScoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListConformancePackComplianceScoresWithContext(ctx context.Context, input *configservice.ListConformancePackComplianceScoresInput, opts ...request.Option) (*configservice.ListConformancePackComplianceScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListConformancePackComplianceScoresOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListConformancePackComplianceScoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListDiscoveredResources(input *configservice.ListDiscoveredResourcesInput) (*configservice.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListDiscoveredResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListDiscoveredResourcesPages(input *configservice.ListDiscoveredResourcesInput, fn func(*configservice.ListDiscoveredResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListDiscoveredResourcesOutput{}
	fnCacher := func(out *configservice.ListDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListDiscoveredResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListDiscoveredResourcesPagesWithContext(ctx context.Context, input *configservice.ListDiscoveredResourcesInput, fn func(*configservice.ListDiscoveredResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListDiscoveredResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListDiscoveredResourcesOutput{}
	fnCacher := func(out *configservice.ListDiscoveredResourcesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListDiscoveredResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListDiscoveredResourcesWithContext(ctx context.Context, input *configservice.ListDiscoveredResourcesInput, opts ...request.Option) (*configservice.ListDiscoveredResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListDiscoveredResourcesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListDiscoveredResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListResourceEvaluations(input *configservice.ListResourceEvaluationsInput) (*configservice.ListResourceEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListResourceEvaluationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListResourceEvaluations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListResourceEvaluationsPages(input *configservice.ListResourceEvaluationsInput, fn func(*configservice.ListResourceEvaluationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListResourceEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListResourceEvaluationsOutput{}
	fnCacher := func(out *configservice.ListResourceEvaluationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListResourceEvaluationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListResourceEvaluationsPagesWithContext(ctx context.Context, input *configservice.ListResourceEvaluationsInput, fn func(*configservice.ListResourceEvaluationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListResourceEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListResourceEvaluationsOutput{}
	fnCacher := func(out *configservice.ListResourceEvaluationsOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListResourceEvaluationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListResourceEvaluationsWithContext(ctx context.Context, input *configservice.ListResourceEvaluationsInput, opts ...request.Option) (*configservice.ListResourceEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListResourceEvaluationsOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListResourceEvaluationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListStoredQueries(input *configservice.ListStoredQueriesInput) (*configservice.ListStoredQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListStoredQueriesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListStoredQueries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListStoredQueriesPages(input *configservice.ListStoredQueriesInput, fn func(*configservice.ListStoredQueriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListStoredQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListStoredQueriesOutput{}
	fnCacher := func(out *configservice.ListStoredQueriesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListStoredQueriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListStoredQueriesPagesWithContext(ctx context.Context, input *configservice.ListStoredQueriesInput, fn func(*configservice.ListStoredQueriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListStoredQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListStoredQueriesOutput{}
	fnCacher := func(out *configservice.ListStoredQueriesOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListStoredQueriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListStoredQueriesWithContext(ctx context.Context, input *configservice.ListStoredQueriesInput, opts ...request.Option) (*configservice.ListStoredQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListStoredQueriesOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListStoredQueriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListTagsForResource(input *configservice.ListTagsForResourceInput) (*configservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ConfigService) ListTagsForResourcePages(input *configservice.ListTagsForResourceInput, fn func(*configservice.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListTagsForResourceOutput{}
	fnCacher := func(out *configservice.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListTagsForResourcePagesWithContext(ctx context.Context, input *configservice.ListTagsForResourceInput, fn func(*configservice.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*configservice.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &configservice.ListTagsForResourceOutput{}
	fnCacher := func(out *configservice.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.ConfigServiceAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ConfigService) ListTagsForResourceWithContext(ctx context.Context, input *configservice.ListTagsForResourceInput, opts ...request.Option) (*configservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*configservice.ListTagsForResourceOutput), nil
	}
	out, err := c.ConfigServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
