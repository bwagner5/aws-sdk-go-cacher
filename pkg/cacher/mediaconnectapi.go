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
	"github.com/aws/aws-sdk-go/service/mediaconnect"
	"github.com/aws/aws-sdk-go/service/mediaconnect/mediaconnectiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MediaConnect struct {
	mediaconnectiface.MediaConnectAPI
	cache *cache.Cache
}

func NewMediaConnect(mediaconnectapi mediaconnectiface.MediaConnectAPI) *MediaConnect {
	return &MediaConnect{
		MediaConnectAPI: mediaconnectapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MediaConnect) DescribeFlow(input *mediaconnect.DescribeFlowInput) (*mediaconnect.DescribeFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeFlowOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeFlow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) DescribeFlowWithContext(ctx context.Context, input *mediaconnect.DescribeFlowInput, opts ...request.Option) (*mediaconnect.DescribeFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeFlowOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeFlowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) DescribeOffering(input *mediaconnect.DescribeOfferingInput) (*mediaconnect.DescribeOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeOfferingOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeOffering(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) DescribeOfferingWithContext(ctx context.Context, input *mediaconnect.DescribeOfferingInput, opts ...request.Option) (*mediaconnect.DescribeOfferingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeOfferingOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeOfferingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) DescribeReservation(input *mediaconnect.DescribeReservationInput) (*mediaconnect.DescribeReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeReservationOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) DescribeReservationWithContext(ctx context.Context, input *mediaconnect.DescribeReservationInput, opts ...request.Option) (*mediaconnect.DescribeReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.DescribeReservationOutput), nil
	}
	out, err := c.MediaConnectAPI.DescribeReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListEntitlements(input *mediaconnect.ListEntitlementsInput) (*mediaconnect.ListEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListEntitlementsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListEntitlements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListEntitlementsPages(input *mediaconnect.ListEntitlementsInput, fn func(*mediaconnect.ListEntitlementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListEntitlementsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListEntitlementsOutput{}
	fnCacher := func(out *mediaconnect.ListEntitlementsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListEntitlementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListEntitlementsPagesWithContext(ctx context.Context, input *mediaconnect.ListEntitlementsInput, fn func(*mediaconnect.ListEntitlementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListEntitlementsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListEntitlementsOutput{}
	fnCacher := func(out *mediaconnect.ListEntitlementsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListEntitlementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListEntitlementsWithContext(ctx context.Context, input *mediaconnect.ListEntitlementsInput, opts ...request.Option) (*mediaconnect.ListEntitlementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListEntitlementsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListEntitlementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListFlows(input *mediaconnect.ListFlowsInput) (*mediaconnect.ListFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListFlowsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListFlows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListFlowsPages(input *mediaconnect.ListFlowsInput, fn func(*mediaconnect.ListFlowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListFlowsOutput{}
	fnCacher := func(out *mediaconnect.ListFlowsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListFlowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListFlowsPagesWithContext(ctx context.Context, input *mediaconnect.ListFlowsInput, fn func(*mediaconnect.ListFlowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListFlowsOutput{}
	fnCacher := func(out *mediaconnect.ListFlowsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListFlowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListFlowsWithContext(ctx context.Context, input *mediaconnect.ListFlowsInput, opts ...request.Option) (*mediaconnect.ListFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListFlowsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListFlowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListOfferings(input *mediaconnect.ListOfferingsInput) (*mediaconnect.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListOfferingsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListOfferingsPages(input *mediaconnect.ListOfferingsInput, fn func(*mediaconnect.ListOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListOfferingsOutput{}
	fnCacher := func(out *mediaconnect.ListOfferingsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListOfferingsPagesWithContext(ctx context.Context, input *mediaconnect.ListOfferingsInput, fn func(*mediaconnect.ListOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListOfferingsOutput{}
	fnCacher := func(out *mediaconnect.ListOfferingsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListOfferingsWithContext(ctx context.Context, input *mediaconnect.ListOfferingsInput, opts ...request.Option) (*mediaconnect.ListOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListOfferingsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListReservations(input *mediaconnect.ListReservationsInput) (*mediaconnect.ListReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListReservationsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListReservationsPages(input *mediaconnect.ListReservationsInput, fn func(*mediaconnect.ListReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListReservationsOutput{}
	fnCacher := func(out *mediaconnect.ListReservationsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListReservationsPagesWithContext(ctx context.Context, input *mediaconnect.ListReservationsInput, fn func(*mediaconnect.ListReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mediaconnect.ListReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &mediaconnect.ListReservationsOutput{}
	fnCacher := func(out *mediaconnect.ListReservationsOutput, more bool) bool {
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
	if err := c.MediaConnectAPI.ListReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MediaConnect) ListReservationsWithContext(ctx context.Context, input *mediaconnect.ListReservationsInput, opts ...request.Option) (*mediaconnect.ListReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListReservationsOutput), nil
	}
	out, err := c.MediaConnectAPI.ListReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListTagsForResource(input *mediaconnect.ListTagsForResourceInput) (*mediaconnect.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaConnectAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MediaConnect) ListTagsForResourceWithContext(ctx context.Context, input *mediaconnect.ListTagsForResourceInput, opts ...request.Option) (*mediaconnect.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mediaconnect.ListTagsForResourceOutput), nil
	}
	out, err := c.MediaConnectAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
