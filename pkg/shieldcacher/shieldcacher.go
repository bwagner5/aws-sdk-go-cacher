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
package shieldcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/shield"
	"github.com/aws/aws-sdk-go/service/shield/shieldiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Shield struct {
	shieldiface.ShieldAPI
	cache *cache.Cache
}

func New(shieldapi shieldiface.ShieldAPI) *Shield {
	return &Shield{
		ShieldAPI: shieldapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Shield) DescribeAttack(input *shield.DescribeAttackInput) (*shield.DescribeAttackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeAttackOutput), nil
	}
	out, err := c.ShieldAPI.DescribeAttack(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeAttackStatistics(input *shield.DescribeAttackStatisticsInput) (*shield.DescribeAttackStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeAttackStatisticsOutput), nil
	}
	out, err := c.ShieldAPI.DescribeAttackStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeAttackStatisticsWithContext(ctx context.Context, input *shield.DescribeAttackStatisticsInput, opts ...request.Option) (*shield.DescribeAttackStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeAttackStatisticsOutput), nil
	}
	out, err := c.ShieldAPI.DescribeAttackStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeAttackWithContext(ctx context.Context, input *shield.DescribeAttackInput, opts ...request.Option) (*shield.DescribeAttackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeAttackOutput), nil
	}
	out, err := c.ShieldAPI.DescribeAttackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeDRTAccess(input *shield.DescribeDRTAccessInput) (*shield.DescribeDRTAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeDRTAccessOutput), nil
	}
	out, err := c.ShieldAPI.DescribeDRTAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeDRTAccessWithContext(ctx context.Context, input *shield.DescribeDRTAccessInput, opts ...request.Option) (*shield.DescribeDRTAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeDRTAccessOutput), nil
	}
	out, err := c.ShieldAPI.DescribeDRTAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeEmergencyContactSettings(input *shield.DescribeEmergencyContactSettingsInput) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeEmergencyContactSettingsOutput), nil
	}
	out, err := c.ShieldAPI.DescribeEmergencyContactSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeEmergencyContactSettingsWithContext(ctx context.Context, input *shield.DescribeEmergencyContactSettingsInput, opts ...request.Option) (*shield.DescribeEmergencyContactSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeEmergencyContactSettingsOutput), nil
	}
	out, err := c.ShieldAPI.DescribeEmergencyContactSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeProtection(input *shield.DescribeProtectionInput) (*shield.DescribeProtectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeProtectionOutput), nil
	}
	out, err := c.ShieldAPI.DescribeProtection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeProtectionGroup(input *shield.DescribeProtectionGroupInput) (*shield.DescribeProtectionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeProtectionGroupOutput), nil
	}
	out, err := c.ShieldAPI.DescribeProtectionGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeProtectionGroupWithContext(ctx context.Context, input *shield.DescribeProtectionGroupInput, opts ...request.Option) (*shield.DescribeProtectionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeProtectionGroupOutput), nil
	}
	out, err := c.ShieldAPI.DescribeProtectionGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeProtectionWithContext(ctx context.Context, input *shield.DescribeProtectionInput, opts ...request.Option) (*shield.DescribeProtectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeProtectionOutput), nil
	}
	out, err := c.ShieldAPI.DescribeProtectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeSubscription(input *shield.DescribeSubscriptionInput) (*shield.DescribeSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeSubscriptionOutput), nil
	}
	out, err := c.ShieldAPI.DescribeSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) DescribeSubscriptionWithContext(ctx context.Context, input *shield.DescribeSubscriptionInput, opts ...request.Option) (*shield.DescribeSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.DescribeSubscriptionOutput), nil
	}
	out, err := c.ShieldAPI.DescribeSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) GetSubscriptionState(input *shield.GetSubscriptionStateInput) (*shield.GetSubscriptionStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.GetSubscriptionStateOutput), nil
	}
	out, err := c.ShieldAPI.GetSubscriptionState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) GetSubscriptionStateWithContext(ctx context.Context, input *shield.GetSubscriptionStateInput, opts ...request.Option) (*shield.GetSubscriptionStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.GetSubscriptionStateOutput), nil
	}
	out, err := c.ShieldAPI.GetSubscriptionStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListAttacks(input *shield.ListAttacksInput) (*shield.ListAttacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListAttacksOutput), nil
	}
	out, err := c.ShieldAPI.ListAttacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListAttacksPages(input *shield.ListAttacksInput, fn func(*shield.ListAttacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListAttacksOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListAttacksOutput{}
	fnCacher := func(out *shield.ListAttacksOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListAttacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListAttacksPagesWithContext(ctx context.Context, input *shield.ListAttacksInput, fn func(*shield.ListAttacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListAttacksOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListAttacksOutput{}
	fnCacher := func(out *shield.ListAttacksOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListAttacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListAttacksWithContext(ctx context.Context, input *shield.ListAttacksInput, opts ...request.Option) (*shield.ListAttacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListAttacksOutput), nil
	}
	out, err := c.ShieldAPI.ListAttacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListProtectionGroups(input *shield.ListProtectionGroupsInput) (*shield.ListProtectionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListProtectionGroupsOutput), nil
	}
	out, err := c.ShieldAPI.ListProtectionGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListProtectionGroupsPages(input *shield.ListProtectionGroupsInput, fn func(*shield.ListProtectionGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListProtectionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListProtectionGroupsOutput{}
	fnCacher := func(out *shield.ListProtectionGroupsOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListProtectionGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListProtectionGroupsPagesWithContext(ctx context.Context, input *shield.ListProtectionGroupsInput, fn func(*shield.ListProtectionGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListProtectionGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListProtectionGroupsOutput{}
	fnCacher := func(out *shield.ListProtectionGroupsOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListProtectionGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListProtectionGroupsWithContext(ctx context.Context, input *shield.ListProtectionGroupsInput, opts ...request.Option) (*shield.ListProtectionGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListProtectionGroupsOutput), nil
	}
	out, err := c.ShieldAPI.ListProtectionGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListProtections(input *shield.ListProtectionsInput) (*shield.ListProtectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListProtectionsOutput), nil
	}
	out, err := c.ShieldAPI.ListProtections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListProtectionsPages(input *shield.ListProtectionsInput, fn func(*shield.ListProtectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListProtectionsOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListProtectionsOutput{}
	fnCacher := func(out *shield.ListProtectionsOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListProtectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListProtectionsPagesWithContext(ctx context.Context, input *shield.ListProtectionsInput, fn func(*shield.ListProtectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListProtectionsOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListProtectionsOutput{}
	fnCacher := func(out *shield.ListProtectionsOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListProtectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListProtectionsWithContext(ctx context.Context, input *shield.ListProtectionsInput, opts ...request.Option) (*shield.ListProtectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListProtectionsOutput), nil
	}
	out, err := c.ShieldAPI.ListProtectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListResourcesInProtectionGroup(input *shield.ListResourcesInProtectionGroupInput) (*shield.ListResourcesInProtectionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListResourcesInProtectionGroupOutput), nil
	}
	out, err := c.ShieldAPI.ListResourcesInProtectionGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListResourcesInProtectionGroupPages(input *shield.ListResourcesInProtectionGroupInput, fn func(*shield.ListResourcesInProtectionGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListResourcesInProtectionGroupOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListResourcesInProtectionGroupOutput{}
	fnCacher := func(out *shield.ListResourcesInProtectionGroupOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListResourcesInProtectionGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListResourcesInProtectionGroupPagesWithContext(ctx context.Context, input *shield.ListResourcesInProtectionGroupInput, fn func(*shield.ListResourcesInProtectionGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*shield.ListResourcesInProtectionGroupOutput), false)
		return nil
	}
	cachable := true
	output := &shield.ListResourcesInProtectionGroupOutput{}
	fnCacher := func(out *shield.ListResourcesInProtectionGroupOutput, more bool) bool {
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
	if err := c.ShieldAPI.ListResourcesInProtectionGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Shield) ListResourcesInProtectionGroupWithContext(ctx context.Context, input *shield.ListResourcesInProtectionGroupInput, opts ...request.Option) (*shield.ListResourcesInProtectionGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListResourcesInProtectionGroupOutput), nil
	}
	out, err := c.ShieldAPI.ListResourcesInProtectionGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListTagsForResource(input *shield.ListTagsForResourceInput) (*shield.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListTagsForResourceOutput), nil
	}
	out, err := c.ShieldAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Shield) ListTagsForResourceWithContext(ctx context.Context, input *shield.ListTagsForResourceInput, opts ...request.Option) (*shield.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*shield.ListTagsForResourceOutput), nil
	}
	out, err := c.ShieldAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
