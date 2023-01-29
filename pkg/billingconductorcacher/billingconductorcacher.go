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
package billingconductorcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/billingconductor"
	"github.com/aws/aws-sdk-go/service/billingconductor/billingconductoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type BillingConductor struct {
	billingconductoriface.BillingConductorAPI
	cache *cache.Cache
}

func New(billingconductorapi billingconductoriface.BillingConductorAPI) *BillingConductor {
	return &BillingConductor{
		BillingConductorAPI: billingconductorapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *BillingConductor) BatchAssociateResourcesToCustomLineItem(input *billingconductor.BatchAssociateResourcesToCustomLineItemInput) (*billingconductor.BatchAssociateResourcesToCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.BatchAssociateResourcesToCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.BatchAssociateResourcesToCustomLineItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) BatchAssociateResourcesToCustomLineItemWithContext(ctx context.Context, input *billingconductor.BatchAssociateResourcesToCustomLineItemInput, opts ...request.Option) (*billingconductor.BatchAssociateResourcesToCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.BatchAssociateResourcesToCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.BatchAssociateResourcesToCustomLineItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) BatchDisassociateResourcesFromCustomLineItem(input *billingconductor.BatchDisassociateResourcesFromCustomLineItemInput) (*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.BatchDisassociateResourcesFromCustomLineItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) BatchDisassociateResourcesFromCustomLineItemWithContext(ctx context.Context, input *billingconductor.BatchDisassociateResourcesFromCustomLineItemInput, opts ...request.Option) (*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.BatchDisassociateResourcesFromCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.BatchDisassociateResourcesFromCustomLineItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListAccountAssociations(input *billingconductor.ListAccountAssociationsInput) (*billingconductor.ListAccountAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListAccountAssociationsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListAccountAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListAccountAssociationsPages(input *billingconductor.ListAccountAssociationsInput, fn func(*billingconductor.ListAccountAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListAccountAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListAccountAssociationsOutput{}
	fnCacher := func(out *billingconductor.ListAccountAssociationsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListAccountAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListAccountAssociationsPagesWithContext(ctx context.Context, input *billingconductor.ListAccountAssociationsInput, fn func(*billingconductor.ListAccountAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListAccountAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListAccountAssociationsOutput{}
	fnCacher := func(out *billingconductor.ListAccountAssociationsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListAccountAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListAccountAssociationsWithContext(ctx context.Context, input *billingconductor.ListAccountAssociationsInput, opts ...request.Option) (*billingconductor.ListAccountAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListAccountAssociationsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListAccountAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListBillingGroupCostReports(input *billingconductor.ListBillingGroupCostReportsInput) (*billingconductor.ListBillingGroupCostReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListBillingGroupCostReportsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListBillingGroupCostReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListBillingGroupCostReportsPages(input *billingconductor.ListBillingGroupCostReportsInput, fn func(*billingconductor.ListBillingGroupCostReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListBillingGroupCostReportsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListBillingGroupCostReportsOutput{}
	fnCacher := func(out *billingconductor.ListBillingGroupCostReportsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListBillingGroupCostReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListBillingGroupCostReportsPagesWithContext(ctx context.Context, input *billingconductor.ListBillingGroupCostReportsInput, fn func(*billingconductor.ListBillingGroupCostReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListBillingGroupCostReportsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListBillingGroupCostReportsOutput{}
	fnCacher := func(out *billingconductor.ListBillingGroupCostReportsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListBillingGroupCostReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListBillingGroupCostReportsWithContext(ctx context.Context, input *billingconductor.ListBillingGroupCostReportsInput, opts ...request.Option) (*billingconductor.ListBillingGroupCostReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListBillingGroupCostReportsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListBillingGroupCostReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListBillingGroups(input *billingconductor.ListBillingGroupsInput) (*billingconductor.ListBillingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListBillingGroupsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListBillingGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListBillingGroupsPages(input *billingconductor.ListBillingGroupsInput, fn func(*billingconductor.ListBillingGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListBillingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListBillingGroupsOutput{}
	fnCacher := func(out *billingconductor.ListBillingGroupsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListBillingGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListBillingGroupsPagesWithContext(ctx context.Context, input *billingconductor.ListBillingGroupsInput, fn func(*billingconductor.ListBillingGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListBillingGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListBillingGroupsOutput{}
	fnCacher := func(out *billingconductor.ListBillingGroupsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListBillingGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListBillingGroupsWithContext(ctx context.Context, input *billingconductor.ListBillingGroupsInput, opts ...request.Option) (*billingconductor.ListBillingGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListBillingGroupsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListBillingGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListCustomLineItemVersions(input *billingconductor.ListCustomLineItemVersionsInput) (*billingconductor.ListCustomLineItemVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListCustomLineItemVersionsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListCustomLineItemVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListCustomLineItemVersionsPages(input *billingconductor.ListCustomLineItemVersionsInput, fn func(*billingconductor.ListCustomLineItemVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListCustomLineItemVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListCustomLineItemVersionsOutput{}
	fnCacher := func(out *billingconductor.ListCustomLineItemVersionsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListCustomLineItemVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListCustomLineItemVersionsPagesWithContext(ctx context.Context, input *billingconductor.ListCustomLineItemVersionsInput, fn func(*billingconductor.ListCustomLineItemVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListCustomLineItemVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListCustomLineItemVersionsOutput{}
	fnCacher := func(out *billingconductor.ListCustomLineItemVersionsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListCustomLineItemVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListCustomLineItemVersionsWithContext(ctx context.Context, input *billingconductor.ListCustomLineItemVersionsInput, opts ...request.Option) (*billingconductor.ListCustomLineItemVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListCustomLineItemVersionsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListCustomLineItemVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListCustomLineItems(input *billingconductor.ListCustomLineItemsInput) (*billingconductor.ListCustomLineItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListCustomLineItemsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListCustomLineItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListCustomLineItemsPages(input *billingconductor.ListCustomLineItemsInput, fn func(*billingconductor.ListCustomLineItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListCustomLineItemsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListCustomLineItemsOutput{}
	fnCacher := func(out *billingconductor.ListCustomLineItemsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListCustomLineItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListCustomLineItemsPagesWithContext(ctx context.Context, input *billingconductor.ListCustomLineItemsInput, fn func(*billingconductor.ListCustomLineItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListCustomLineItemsOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListCustomLineItemsOutput{}
	fnCacher := func(out *billingconductor.ListCustomLineItemsOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListCustomLineItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListCustomLineItemsWithContext(ctx context.Context, input *billingconductor.ListCustomLineItemsInput, opts ...request.Option) (*billingconductor.ListCustomLineItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListCustomLineItemsOutput), nil
	}
	out, err := c.BillingConductorAPI.ListCustomLineItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingPlans(input *billingconductor.ListPricingPlansInput) (*billingconductor.ListPricingPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingPlansOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingPlansAssociatedWithPricingRule(input *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput) (*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingPlansAssociatedWithPricingRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingPlansAssociatedWithPricingRulePages(input *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, fn func(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput{}
	fnCacher := func(out *billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingPlansAssociatedWithPricingRulePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingPlansAssociatedWithPricingRulePagesWithContext(ctx context.Context, input *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, fn func(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput{}
	fnCacher := func(out *billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingPlansAssociatedWithPricingRulePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingPlansAssociatedWithPricingRuleWithContext(ctx context.Context, input *billingconductor.ListPricingPlansAssociatedWithPricingRuleInput, opts ...request.Option) (*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingPlansAssociatedWithPricingRuleOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingPlansAssociatedWithPricingRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingPlansPages(input *billingconductor.ListPricingPlansInput, fn func(*billingconductor.ListPricingPlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingPlansOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingPlansOutput{}
	fnCacher := func(out *billingconductor.ListPricingPlansOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingPlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingPlansPagesWithContext(ctx context.Context, input *billingconductor.ListPricingPlansInput, fn func(*billingconductor.ListPricingPlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingPlansOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingPlansOutput{}
	fnCacher := func(out *billingconductor.ListPricingPlansOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingPlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingPlansWithContext(ctx context.Context, input *billingconductor.ListPricingPlansInput, opts ...request.Option) (*billingconductor.ListPricingPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingPlansOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingRules(input *billingconductor.ListPricingRulesInput) (*billingconductor.ListPricingRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingRulesOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingRulesAssociatedToPricingPlan(input *billingconductor.ListPricingRulesAssociatedToPricingPlanInput) (*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingRulesAssociatedToPricingPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingRulesAssociatedToPricingPlanPages(input *billingconductor.ListPricingRulesAssociatedToPricingPlanInput, fn func(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingRulesAssociatedToPricingPlanOutput{}
	fnCacher := func(out *billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingRulesAssociatedToPricingPlanPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingRulesAssociatedToPricingPlanPagesWithContext(ctx context.Context, input *billingconductor.ListPricingRulesAssociatedToPricingPlanInput, fn func(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingRulesAssociatedToPricingPlanOutput{}
	fnCacher := func(out *billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingRulesAssociatedToPricingPlanPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingRulesAssociatedToPricingPlanWithContext(ctx context.Context, input *billingconductor.ListPricingRulesAssociatedToPricingPlanInput, opts ...request.Option) (*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingRulesAssociatedToPricingPlanOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingRulesAssociatedToPricingPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListPricingRulesPages(input *billingconductor.ListPricingRulesInput, fn func(*billingconductor.ListPricingRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingRulesOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingRulesOutput{}
	fnCacher := func(out *billingconductor.ListPricingRulesOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingRulesPagesWithContext(ctx context.Context, input *billingconductor.ListPricingRulesInput, fn func(*billingconductor.ListPricingRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListPricingRulesOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListPricingRulesOutput{}
	fnCacher := func(out *billingconductor.ListPricingRulesOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListPricingRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListPricingRulesWithContext(ctx context.Context, input *billingconductor.ListPricingRulesInput, opts ...request.Option) (*billingconductor.ListPricingRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListPricingRulesOutput), nil
	}
	out, err := c.BillingConductorAPI.ListPricingRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListResourcesAssociatedToCustomLineItem(input *billingconductor.ListResourcesAssociatedToCustomLineItemInput) (*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.ListResourcesAssociatedToCustomLineItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListResourcesAssociatedToCustomLineItemPages(input *billingconductor.ListResourcesAssociatedToCustomLineItemInput, fn func(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListResourcesAssociatedToCustomLineItemOutput{}
	fnCacher := func(out *billingconductor.ListResourcesAssociatedToCustomLineItemOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListResourcesAssociatedToCustomLineItemPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListResourcesAssociatedToCustomLineItemPagesWithContext(ctx context.Context, input *billingconductor.ListResourcesAssociatedToCustomLineItemInput, fn func(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput), false)
		return nil
	}
	cachable := true
	output := &billingconductor.ListResourcesAssociatedToCustomLineItemOutput{}
	fnCacher := func(out *billingconductor.ListResourcesAssociatedToCustomLineItemOutput, more bool) bool {
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
	if err := c.BillingConductorAPI.ListResourcesAssociatedToCustomLineItemPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *BillingConductor) ListResourcesAssociatedToCustomLineItemWithContext(ctx context.Context, input *billingconductor.ListResourcesAssociatedToCustomLineItemInput, opts ...request.Option) (*billingconductor.ListResourcesAssociatedToCustomLineItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListResourcesAssociatedToCustomLineItemOutput), nil
	}
	out, err := c.BillingConductorAPI.ListResourcesAssociatedToCustomLineItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListTagsForResource(input *billingconductor.ListTagsForResourceInput) (*billingconductor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListTagsForResourceOutput), nil
	}
	out, err := c.BillingConductorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *BillingConductor) ListTagsForResourceWithContext(ctx context.Context, input *billingconductor.ListTagsForResourceInput, opts ...request.Option) (*billingconductor.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*billingconductor.ListTagsForResourceOutput), nil
	}
	out, err := c.BillingConductorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
