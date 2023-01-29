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
package codestarnotificationscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"github.com/aws/aws-sdk-go/service/codestarnotifications/codestarnotificationsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeStarNotifications struct {
	codestarnotificationsiface.CodeStarNotificationsAPI
	cache *cache.Cache
}

func New(codestarnotificationsapi codestarnotificationsiface.CodeStarNotificationsAPI) *CodeStarNotifications {
	return &CodeStarNotifications{
		CodeStarNotificationsAPI: codestarnotificationsapi,
		cache:                    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeStarNotifications) DescribeNotificationRule(input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.DescribeNotificationRuleOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.DescribeNotificationRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) DescribeNotificationRuleWithContext(ctx context.Context, input *codestarnotifications.DescribeNotificationRuleInput, opts ...request.Option) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.DescribeNotificationRuleOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.DescribeNotificationRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListEventTypes(input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListEventTypesOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListEventTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListEventTypesPages(input *codestarnotifications.ListEventTypesInput, fn func(*codestarnotifications.ListEventTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListEventTypesOutput{}
	fnCacher := func(out *codestarnotifications.ListEventTypesOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListEventTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListEventTypesPagesWithContext(ctx context.Context, input *codestarnotifications.ListEventTypesInput, fn func(*codestarnotifications.ListEventTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListEventTypesOutput{}
	fnCacher := func(out *codestarnotifications.ListEventTypesOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListEventTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListEventTypesWithContext(ctx context.Context, input *codestarnotifications.ListEventTypesInput, opts ...request.Option) (*codestarnotifications.ListEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListEventTypesOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListEventTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListNotificationRules(input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListNotificationRulesOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListNotificationRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListNotificationRulesPages(input *codestarnotifications.ListNotificationRulesInput, fn func(*codestarnotifications.ListNotificationRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListNotificationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListNotificationRulesOutput{}
	fnCacher := func(out *codestarnotifications.ListNotificationRulesOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListNotificationRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListNotificationRulesPagesWithContext(ctx context.Context, input *codestarnotifications.ListNotificationRulesInput, fn func(*codestarnotifications.ListNotificationRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListNotificationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListNotificationRulesOutput{}
	fnCacher := func(out *codestarnotifications.ListNotificationRulesOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListNotificationRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListNotificationRulesWithContext(ctx context.Context, input *codestarnotifications.ListNotificationRulesInput, opts ...request.Option) (*codestarnotifications.ListNotificationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListNotificationRulesOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListNotificationRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListTagsForResource(input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListTagsForResourceWithContext(ctx context.Context, input *codestarnotifications.ListTagsForResourceInput, opts ...request.Option) (*codestarnotifications.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListTargets(input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListTargetsOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeStarNotifications) ListTargetsPages(input *codestarnotifications.ListTargetsInput, fn func(*codestarnotifications.ListTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListTargetsOutput{}
	fnCacher := func(out *codestarnotifications.ListTargetsOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListTargetsPagesWithContext(ctx context.Context, input *codestarnotifications.ListTargetsInput, fn func(*codestarnotifications.ListTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codestarnotifications.ListTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &codestarnotifications.ListTargetsOutput{}
	fnCacher := func(out *codestarnotifications.ListTargetsOutput, more bool) bool {
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
	if err := c.CodeStarNotificationsAPI.ListTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeStarNotifications) ListTargetsWithContext(ctx context.Context, input *codestarnotifications.ListTargetsInput, opts ...request.Option) (*codestarnotifications.ListTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codestarnotifications.ListTargetsOutput), nil
	}
	out, err := c.CodeStarNotificationsAPI.ListTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
