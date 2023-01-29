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
package ssmincidentscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ssmincidents"
	"github.com/aws/aws-sdk-go/service/ssmincidents/ssmincidentsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SSMIncidents struct {
	ssmincidentsiface.SSMIncidentsAPI
	cache *cache.Cache
}

func New(ssmincidentsapi ssmincidentsiface.SSMIncidentsAPI) *SSMIncidents {
	return &SSMIncidents{
		SSMIncidentsAPI: ssmincidentsapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SSMIncidents) GetIncidentRecord(input *ssmincidents.GetIncidentRecordInput) (*ssmincidents.GetIncidentRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetIncidentRecordOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetIncidentRecord(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetIncidentRecordWithContext(ctx context.Context, input *ssmincidents.GetIncidentRecordInput, opts ...request.Option) (*ssmincidents.GetIncidentRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetIncidentRecordOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetIncidentRecordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetReplicationSet(input *ssmincidents.GetReplicationSetInput) (*ssmincidents.GetReplicationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetReplicationSetOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetReplicationSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetReplicationSetWithContext(ctx context.Context, input *ssmincidents.GetReplicationSetInput, opts ...request.Option) (*ssmincidents.GetReplicationSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetReplicationSetOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetReplicationSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetResourcePolicies(input *ssmincidents.GetResourcePoliciesInput) (*ssmincidents.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetResourcePoliciesOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetResourcePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetResourcePoliciesPages(input *ssmincidents.GetResourcePoliciesInput, fn func(*ssmincidents.GetResourcePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.GetResourcePoliciesOutput{}
	fnCacher := func(out *ssmincidents.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.GetResourcePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) GetResourcePoliciesPagesWithContext(ctx context.Context, input *ssmincidents.GetResourcePoliciesInput, fn func(*ssmincidents.GetResourcePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.GetResourcePoliciesOutput{}
	fnCacher := func(out *ssmincidents.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.GetResourcePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) GetResourcePoliciesWithContext(ctx context.Context, input *ssmincidents.GetResourcePoliciesInput, opts ...request.Option) (*ssmincidents.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetResourcePoliciesOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetResourcePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetResponsePlan(input *ssmincidents.GetResponsePlanInput) (*ssmincidents.GetResponsePlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetResponsePlanOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetResponsePlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetResponsePlanWithContext(ctx context.Context, input *ssmincidents.GetResponsePlanInput, opts ...request.Option) (*ssmincidents.GetResponsePlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetResponsePlanOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetResponsePlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetTimelineEvent(input *ssmincidents.GetTimelineEventInput) (*ssmincidents.GetTimelineEventOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetTimelineEventOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetTimelineEvent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) GetTimelineEventWithContext(ctx context.Context, input *ssmincidents.GetTimelineEventInput, opts ...request.Option) (*ssmincidents.GetTimelineEventOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.GetTimelineEventOutput), nil
	}
	out, err := c.SSMIncidentsAPI.GetTimelineEventWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListIncidentRecords(input *ssmincidents.ListIncidentRecordsInput) (*ssmincidents.ListIncidentRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListIncidentRecordsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListIncidentRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListIncidentRecordsPages(input *ssmincidents.ListIncidentRecordsInput, fn func(*ssmincidents.ListIncidentRecordsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListIncidentRecordsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListIncidentRecordsOutput{}
	fnCacher := func(out *ssmincidents.ListIncidentRecordsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListIncidentRecordsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListIncidentRecordsPagesWithContext(ctx context.Context, input *ssmincidents.ListIncidentRecordsInput, fn func(*ssmincidents.ListIncidentRecordsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListIncidentRecordsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListIncidentRecordsOutput{}
	fnCacher := func(out *ssmincidents.ListIncidentRecordsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListIncidentRecordsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListIncidentRecordsWithContext(ctx context.Context, input *ssmincidents.ListIncidentRecordsInput, opts ...request.Option) (*ssmincidents.ListIncidentRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListIncidentRecordsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListIncidentRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListRelatedItems(input *ssmincidents.ListRelatedItemsInput) (*ssmincidents.ListRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListRelatedItemsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListRelatedItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListRelatedItemsPages(input *ssmincidents.ListRelatedItemsInput, fn func(*ssmincidents.ListRelatedItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListRelatedItemsOutput{}
	fnCacher := func(out *ssmincidents.ListRelatedItemsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListRelatedItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListRelatedItemsPagesWithContext(ctx context.Context, input *ssmincidents.ListRelatedItemsInput, fn func(*ssmincidents.ListRelatedItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListRelatedItemsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListRelatedItemsOutput{}
	fnCacher := func(out *ssmincidents.ListRelatedItemsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListRelatedItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListRelatedItemsWithContext(ctx context.Context, input *ssmincidents.ListRelatedItemsInput, opts ...request.Option) (*ssmincidents.ListRelatedItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListRelatedItemsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListRelatedItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListReplicationSets(input *ssmincidents.ListReplicationSetsInput) (*ssmincidents.ListReplicationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListReplicationSetsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListReplicationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListReplicationSetsPages(input *ssmincidents.ListReplicationSetsInput, fn func(*ssmincidents.ListReplicationSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListReplicationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListReplicationSetsOutput{}
	fnCacher := func(out *ssmincidents.ListReplicationSetsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListReplicationSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListReplicationSetsPagesWithContext(ctx context.Context, input *ssmincidents.ListReplicationSetsInput, fn func(*ssmincidents.ListReplicationSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListReplicationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListReplicationSetsOutput{}
	fnCacher := func(out *ssmincidents.ListReplicationSetsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListReplicationSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListReplicationSetsWithContext(ctx context.Context, input *ssmincidents.ListReplicationSetsInput, opts ...request.Option) (*ssmincidents.ListReplicationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListReplicationSetsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListReplicationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListResponsePlans(input *ssmincidents.ListResponsePlansInput) (*ssmincidents.ListResponsePlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListResponsePlansOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListResponsePlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListResponsePlansPages(input *ssmincidents.ListResponsePlansInput, fn func(*ssmincidents.ListResponsePlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListResponsePlansOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListResponsePlansOutput{}
	fnCacher := func(out *ssmincidents.ListResponsePlansOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListResponsePlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListResponsePlansPagesWithContext(ctx context.Context, input *ssmincidents.ListResponsePlansInput, fn func(*ssmincidents.ListResponsePlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListResponsePlansOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListResponsePlansOutput{}
	fnCacher := func(out *ssmincidents.ListResponsePlansOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListResponsePlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListResponsePlansWithContext(ctx context.Context, input *ssmincidents.ListResponsePlansInput, opts ...request.Option) (*ssmincidents.ListResponsePlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListResponsePlansOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListResponsePlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListTagsForResource(input *ssmincidents.ListTagsForResourceInput) (*ssmincidents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListTagsForResourceWithContext(ctx context.Context, input *ssmincidents.ListTagsForResourceInput, opts ...request.Option) (*ssmincidents.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListTimelineEvents(input *ssmincidents.ListTimelineEventsInput) (*ssmincidents.ListTimelineEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListTimelineEventsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListTimelineEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMIncidents) ListTimelineEventsPages(input *ssmincidents.ListTimelineEventsInput, fn func(*ssmincidents.ListTimelineEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListTimelineEventsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListTimelineEventsOutput{}
	fnCacher := func(out *ssmincidents.ListTimelineEventsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListTimelineEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListTimelineEventsPagesWithContext(ctx context.Context, input *ssmincidents.ListTimelineEventsInput, fn func(*ssmincidents.ListTimelineEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmincidents.ListTimelineEventsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmincidents.ListTimelineEventsOutput{}
	fnCacher := func(out *ssmincidents.ListTimelineEventsOutput, more bool) bool {
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
	if err := c.SSMIncidentsAPI.ListTimelineEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMIncidents) ListTimelineEventsWithContext(ctx context.Context, input *ssmincidents.ListTimelineEventsInput, opts ...request.Option) (*ssmincidents.ListTimelineEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmincidents.ListTimelineEventsOutput), nil
	}
	out, err := c.SSMIncidentsAPI.ListTimelineEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
