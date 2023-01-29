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
package healthcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/health"
	"github.com/aws/aws-sdk-go/service/health/healthiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Health struct {
	healthiface.HealthAPI
	cache *cache.Cache
}

func New(healthapi healthiface.HealthAPI) *Health {
	return &Health{
		HealthAPI: healthapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Health) DescribeAffectedAccountsForOrganization(input *health.DescribeAffectedAccountsForOrganizationInput) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedAccountsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedAccountsForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeAffectedAccountsForOrganizationPages(input *health.DescribeAffectedAccountsForOrganizationInput, fn func(*health.DescribeAffectedAccountsForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedAccountsForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedAccountsForOrganizationOutput{}
	fnCacher := func(out *health.DescribeAffectedAccountsForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedAccountsForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedAccountsForOrganizationPagesWithContext(ctx context.Context, input *health.DescribeAffectedAccountsForOrganizationInput, fn func(*health.DescribeAffectedAccountsForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedAccountsForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedAccountsForOrganizationOutput{}
	fnCacher := func(out *health.DescribeAffectedAccountsForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedAccountsForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedAccountsForOrganizationWithContext(ctx context.Context, input *health.DescribeAffectedAccountsForOrganizationInput, opts ...request.Option) (*health.DescribeAffectedAccountsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedAccountsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedAccountsForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeAffectedEntities(input *health.DescribeAffectedEntitiesInput) (*health.DescribeAffectedEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedEntitiesOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeAffectedEntitiesForOrganization(input *health.DescribeAffectedEntitiesForOrganizationInput) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedEntitiesForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedEntitiesForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeAffectedEntitiesForOrganizationPages(input *health.DescribeAffectedEntitiesForOrganizationInput, fn func(*health.DescribeAffectedEntitiesForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedEntitiesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedEntitiesForOrganizationOutput{}
	fnCacher := func(out *health.DescribeAffectedEntitiesForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedEntitiesForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedEntitiesForOrganizationPagesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesForOrganizationInput, fn func(*health.DescribeAffectedEntitiesForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedEntitiesForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedEntitiesForOrganizationOutput{}
	fnCacher := func(out *health.DescribeAffectedEntitiesForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedEntitiesForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedEntitiesForOrganizationWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesForOrganizationInput, opts ...request.Option) (*health.DescribeAffectedEntitiesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedEntitiesForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedEntitiesForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeAffectedEntitiesPages(input *health.DescribeAffectedEntitiesInput, fn func(*health.DescribeAffectedEntitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedEntitiesOutput{}
	fnCacher := func(out *health.DescribeAffectedEntitiesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedEntitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedEntitiesPagesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, fn func(*health.DescribeAffectedEntitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeAffectedEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeAffectedEntitiesOutput{}
	fnCacher := func(out *health.DescribeAffectedEntitiesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeAffectedEntitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeAffectedEntitiesWithContext(ctx context.Context, input *health.DescribeAffectedEntitiesInput, opts ...request.Option) (*health.DescribeAffectedEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeAffectedEntitiesOutput), nil
	}
	out, err := c.HealthAPI.DescribeAffectedEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEntityAggregates(input *health.DescribeEntityAggregatesInput) (*health.DescribeEntityAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEntityAggregatesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEntityAggregates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEntityAggregatesWithContext(ctx context.Context, input *health.DescribeEntityAggregatesInput, opts ...request.Option) (*health.DescribeEntityAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEntityAggregatesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEntityAggregatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventAggregates(input *health.DescribeEventAggregatesInput) (*health.DescribeEventAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventAggregatesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventAggregates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventAggregatesPages(input *health.DescribeEventAggregatesInput, fn func(*health.DescribeEventAggregatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventAggregatesOutput{}
	fnCacher := func(out *health.DescribeEventAggregatesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventAggregatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventAggregatesPagesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, fn func(*health.DescribeEventAggregatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventAggregatesOutput{}
	fnCacher := func(out *health.DescribeEventAggregatesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventAggregatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventAggregatesWithContext(ctx context.Context, input *health.DescribeEventAggregatesInput, opts ...request.Option) (*health.DescribeEventAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventAggregatesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventAggregatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventDetails(input *health.DescribeEventDetailsInput) (*health.DescribeEventDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventDetailsOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventDetailsForOrganization(input *health.DescribeEventDetailsForOrganizationInput) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventDetailsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventDetailsForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventDetailsForOrganizationWithContext(ctx context.Context, input *health.DescribeEventDetailsForOrganizationInput, opts ...request.Option) (*health.DescribeEventDetailsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventDetailsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventDetailsForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventDetailsWithContext(ctx context.Context, input *health.DescribeEventDetailsInput, opts ...request.Option) (*health.DescribeEventDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventDetailsOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventTypes(input *health.DescribeEventTypesInput) (*health.DescribeEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventTypesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventTypesPages(input *health.DescribeEventTypesInput, fn func(*health.DescribeEventTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventTypesOutput{}
	fnCacher := func(out *health.DescribeEventTypesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventTypesPagesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, fn func(*health.DescribeEventTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventTypesOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventTypesOutput{}
	fnCacher := func(out *health.DescribeEventTypesOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventTypesWithContext(ctx context.Context, input *health.DescribeEventTypesInput, opts ...request.Option) (*health.DescribeEventTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventTypesOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEvents(input *health.DescribeEventsInput) (*health.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventsOutput), nil
	}
	out, err := c.HealthAPI.DescribeEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventsForOrganization(input *health.DescribeEventsForOrganizationInput) (*health.DescribeEventsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventsForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventsForOrganizationPages(input *health.DescribeEventsForOrganizationInput, fn func(*health.DescribeEventsForOrganizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventsForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventsForOrganizationOutput{}
	fnCacher := func(out *health.DescribeEventsForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventsForOrganizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventsForOrganizationPagesWithContext(ctx context.Context, input *health.DescribeEventsForOrganizationInput, fn func(*health.DescribeEventsForOrganizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventsForOrganizationOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventsForOrganizationOutput{}
	fnCacher := func(out *health.DescribeEventsForOrganizationOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventsForOrganizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventsForOrganizationWithContext(ctx context.Context, input *health.DescribeEventsForOrganizationInput, opts ...request.Option) (*health.DescribeEventsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventsForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventsForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeEventsPages(input *health.DescribeEventsInput, fn func(*health.DescribeEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventsOutput{}
	fnCacher := func(out *health.DescribeEventsOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventsPagesWithContext(ctx context.Context, input *health.DescribeEventsInput, fn func(*health.DescribeEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*health.DescribeEventsOutput), false)
		return nil
	}
	cachable := true
	output := &health.DescribeEventsOutput{}
	fnCacher := func(out *health.DescribeEventsOutput, more bool) bool {
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
	if err := c.HealthAPI.DescribeEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Health) DescribeEventsWithContext(ctx context.Context, input *health.DescribeEventsInput, opts ...request.Option) (*health.DescribeEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeEventsOutput), nil
	}
	out, err := c.HealthAPI.DescribeEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeHealthServiceStatusForOrganization(input *health.DescribeHealthServiceStatusForOrganizationInput) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeHealthServiceStatusForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeHealthServiceStatusForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Health) DescribeHealthServiceStatusForOrganizationWithContext(ctx context.Context, input *health.DescribeHealthServiceStatusForOrganizationInput, opts ...request.Option) (*health.DescribeHealthServiceStatusForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*health.DescribeHealthServiceStatusForOrganizationOutput), nil
	}
	out, err := c.HealthAPI.DescribeHealthServiceStatusForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
