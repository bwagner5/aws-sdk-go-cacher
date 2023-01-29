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
package ec2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EC2 struct {
	ec2iface.EC2API
	cache *cache.Cache
}

func New(ec2api ec2iface.EC2API) *EC2 {
	return &EC2{
		EC2API: ec2api,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EC2) DescribeAccountAttributes(input *ec2.DescribeAccountAttributesInput) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAccountAttributesWithContext(ctx context.Context, input *ec2.DescribeAccountAttributesInput, opts ...request.Option) (*ec2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressTransfers(input *ec2.DescribeAddressTransfersInput) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.EC2API.DescribeAddressTransfers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressTransfersPages(input *ec2.DescribeAddressTransfersInput, fn func(*ec2.DescribeAddressTransfersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressTransfersOutput{}
	fnCacher := func(out *ec2.DescribeAddressTransfersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressTransfersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressTransfersPagesWithContext(ctx context.Context, input *ec2.DescribeAddressTransfersInput, fn func(*ec2.DescribeAddressTransfersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressTransfersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressTransfersOutput{}
	fnCacher := func(out *ec2.DescribeAddressTransfersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressTransfersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressTransfersWithContext(ctx context.Context, input *ec2.DescribeAddressTransfersInput, opts ...request.Option) (*ec2.DescribeAddressTransfersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressTransfersOutput), nil
	}
	out, err := c.EC2API.DescribeAddressTransfersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddresses(input *ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesAttribute(input *ec2.DescribeAddressesAttributeInput) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesAttributePages(input *ec2.DescribeAddressesAttributeInput, fn func(*ec2.DescribeAddressesAttributeOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressesAttributeOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressesAttributeOutput{}
	fnCacher := func(out *ec2.DescribeAddressesAttributeOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressesAttributePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressesAttributePagesWithContext(ctx context.Context, input *ec2.DescribeAddressesAttributeInput, fn func(*ec2.DescribeAddressesAttributeOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAddressesAttributeOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAddressesAttributeOutput{}
	fnCacher := func(out *ec2.DescribeAddressesAttributeOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAddressesAttributePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAddressesAttributeWithContext(ctx context.Context, input *ec2.DescribeAddressesAttributeInput, opts ...request.Option) (*ec2.DescribeAddressesAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAddressesWithContext(ctx context.Context, input *ec2.DescribeAddressesInput, opts ...request.Option) (*ec2.DescribeAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAggregateIdFormat(input *ec2.DescribeAggregateIdFormatInput) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeAggregateIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAggregateIdFormatWithContext(ctx context.Context, input *ec2.DescribeAggregateIdFormatInput, opts ...request.Option) (*ec2.DescribeAggregateIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAggregateIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeAggregateIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAvailabilityZones(input *ec2.DescribeAvailabilityZonesInput) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.EC2API.DescribeAvailabilityZones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAvailabilityZonesWithContext(ctx context.Context, input *ec2.DescribeAvailabilityZonesInput, opts ...request.Option) (*ec2.DescribeAvailabilityZonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAvailabilityZonesOutput), nil
	}
	out, err := c.EC2API.DescribeAvailabilityZonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptions(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, fn func(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput{}
	fnCacher := func(out *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsPagesWithContext(ctx context.Context, input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, fn func(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput{}
	fnCacher := func(out *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx context.Context, input *ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsInput, opts ...request.Option) (*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeAwsNetworkPerformanceMetricSubscriptionsOutput), nil
	}
	out, err := c.EC2API.DescribeAwsNetworkPerformanceMetricSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeBundleTasks(input *ec2.DescribeBundleTasksInput) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.EC2API.DescribeBundleTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeBundleTasksWithContext(ctx context.Context, input *ec2.DescribeBundleTasksInput, opts ...request.Option) (*ec2.DescribeBundleTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeBundleTasksOutput), nil
	}
	out, err := c.EC2API.DescribeBundleTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeByoipCidrs(input *ec2.DescribeByoipCidrsInput) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.EC2API.DescribeByoipCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeByoipCidrsPages(input *ec2.DescribeByoipCidrsInput, fn func(*ec2.DescribeByoipCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeByoipCidrsOutput{}
	fnCacher := func(out *ec2.DescribeByoipCidrsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeByoipCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeByoipCidrsPagesWithContext(ctx context.Context, input *ec2.DescribeByoipCidrsInput, fn func(*ec2.DescribeByoipCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeByoipCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeByoipCidrsOutput{}
	fnCacher := func(out *ec2.DescribeByoipCidrsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeByoipCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeByoipCidrsWithContext(ctx context.Context, input *ec2.DescribeByoipCidrsInput, opts ...request.Option) (*ec2.DescribeByoipCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeByoipCidrsOutput), nil
	}
	out, err := c.EC2API.DescribeByoipCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationFleets(input *ec2.DescribeCapacityReservationFleetsInput) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationFleetsPages(input *ec2.DescribeCapacityReservationFleetsInput, fn func(*ec2.DescribeCapacityReservationFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationFleetsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationFleetsPagesWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationFleetsInput, fn func(*ec2.DescribeCapacityReservationFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationFleetsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationFleetsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationFleetsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservations(input *ec2.DescribeCapacityReservationsInput) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCapacityReservationsPages(input *ec2.DescribeCapacityReservationsInput, fn func(*ec2.DescribeCapacityReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationsPagesWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationsInput, fn func(*ec2.DescribeCapacityReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCapacityReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCapacityReservationsOutput{}
	fnCacher := func(out *ec2.DescribeCapacityReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCapacityReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCapacityReservationsWithContext(ctx context.Context, input *ec2.DescribeCapacityReservationsInput, opts ...request.Option) (*ec2.DescribeCapacityReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCapacityReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeCapacityReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCarrierGateways(input *ec2.DescribeCarrierGatewaysInput) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCarrierGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCarrierGatewaysPages(input *ec2.DescribeCarrierGatewaysInput, fn func(*ec2.DescribeCarrierGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCarrierGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeCarrierGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCarrierGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCarrierGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput, fn func(*ec2.DescribeCarrierGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCarrierGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeCarrierGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCarrierGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCarrierGatewaysWithContext(ctx context.Context, input *ec2.DescribeCarrierGatewaysInput, opts ...request.Option) (*ec2.DescribeCarrierGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCarrierGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCarrierGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClassicLinkInstances(input *ec2.DescribeClassicLinkInstancesInput) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeClassicLinkInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClassicLinkInstancesPages(input *ec2.DescribeClassicLinkInstancesInput, fn func(*ec2.DescribeClassicLinkInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClassicLinkInstancesOutput{}
	fnCacher := func(out *ec2.DescribeClassicLinkInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClassicLinkInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClassicLinkInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput, fn func(*ec2.DescribeClassicLinkInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClassicLinkInstancesOutput{}
	fnCacher := func(out *ec2.DescribeClassicLinkInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClassicLinkInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClassicLinkInstancesWithContext(ctx context.Context, input *ec2.DescribeClassicLinkInstancesInput, opts ...request.Option) (*ec2.DescribeClassicLinkInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClassicLinkInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeClassicLinkInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnAuthorizationRules(input *ec2.DescribeClientVpnAuthorizationRulesInput) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnAuthorizationRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnAuthorizationRulesPages(input *ec2.DescribeClientVpnAuthorizationRulesInput, fn func(*ec2.DescribeClientVpnAuthorizationRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnAuthorizationRulesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnAuthorizationRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnAuthorizationRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnAuthorizationRulesPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput, fn func(*ec2.DescribeClientVpnAuthorizationRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnAuthorizationRulesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnAuthorizationRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnAuthorizationRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnAuthorizationRulesWithContext(ctx context.Context, input *ec2.DescribeClientVpnAuthorizationRulesInput, opts ...request.Option) (*ec2.DescribeClientVpnAuthorizationRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnAuthorizationRulesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnAuthorizationRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnConnections(input *ec2.DescribeClientVpnConnectionsInput) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnConnectionsPages(input *ec2.DescribeClientVpnConnectionsInput, fn func(*ec2.DescribeClientVpnConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput, fn func(*ec2.DescribeClientVpnConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeClientVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeClientVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnEndpoints(input *ec2.DescribeClientVpnEndpointsInput) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnEndpointsPages(input *ec2.DescribeClientVpnEndpointsInput, fn func(*ec2.DescribeClientVpnEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput, fn func(*ec2.DescribeClientVpnEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnEndpointsWithContext(ctx context.Context, input *ec2.DescribeClientVpnEndpointsInput, opts ...request.Option) (*ec2.DescribeClientVpnEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnRoutes(input *ec2.DescribeClientVpnRoutesInput) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnRoutesPages(input *ec2.DescribeClientVpnRoutesInput, fn func(*ec2.DescribeClientVpnRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnRoutesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnRoutesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnRoutesPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput, fn func(*ec2.DescribeClientVpnRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnRoutesOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnRoutesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnRoutesWithContext(ctx context.Context, input *ec2.DescribeClientVpnRoutesInput, opts ...request.Option) (*ec2.DescribeClientVpnRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnRoutesOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnTargetNetworks(input *ec2.DescribeClientVpnTargetNetworksInput) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnTargetNetworks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeClientVpnTargetNetworksPages(input *ec2.DescribeClientVpnTargetNetworksInput, fn func(*ec2.DescribeClientVpnTargetNetworksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnTargetNetworksOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnTargetNetworksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnTargetNetworksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnTargetNetworksPagesWithContext(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput, fn func(*ec2.DescribeClientVpnTargetNetworksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeClientVpnTargetNetworksOutput{}
	fnCacher := func(out *ec2.DescribeClientVpnTargetNetworksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeClientVpnTargetNetworksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeClientVpnTargetNetworksWithContext(ctx context.Context, input *ec2.DescribeClientVpnTargetNetworksInput, opts ...request.Option) (*ec2.DescribeClientVpnTargetNetworksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeClientVpnTargetNetworksOutput), nil
	}
	out, err := c.EC2API.DescribeClientVpnTargetNetworksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCoipPools(input *ec2.DescribeCoipPoolsInput) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeCoipPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCoipPoolsPages(input *ec2.DescribeCoipPoolsInput, fn func(*ec2.DescribeCoipPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCoipPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCoipPoolsOutput{}
	fnCacher := func(out *ec2.DescribeCoipPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCoipPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCoipPoolsPagesWithContext(ctx context.Context, input *ec2.DescribeCoipPoolsInput, fn func(*ec2.DescribeCoipPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeCoipPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeCoipPoolsOutput{}
	fnCacher := func(out *ec2.DescribeCoipPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeCoipPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeCoipPoolsWithContext(ctx context.Context, input *ec2.DescribeCoipPoolsInput, opts ...request.Option) (*ec2.DescribeCoipPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCoipPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeCoipPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeConversionTasks(input *ec2.DescribeConversionTasksInput) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.EC2API.DescribeConversionTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeConversionTasksWithContext(ctx context.Context, input *ec2.DescribeConversionTasksInput, opts ...request.Option) (*ec2.DescribeConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeConversionTasksOutput), nil
	}
	out, err := c.EC2API.DescribeConversionTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCustomerGateways(input *ec2.DescribeCustomerGatewaysInput) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCustomerGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeCustomerGatewaysWithContext(ctx context.Context, input *ec2.DescribeCustomerGatewaysInput, opts ...request.Option) (*ec2.DescribeCustomerGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeCustomerGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeCustomerGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeDhcpOptions(input *ec2.DescribeDhcpOptionsInput) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DescribeDhcpOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeDhcpOptionsPages(input *ec2.DescribeDhcpOptionsInput, fn func(*ec2.DescribeDhcpOptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeDhcpOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeDhcpOptionsOutput{}
	fnCacher := func(out *ec2.DescribeDhcpOptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeDhcpOptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeDhcpOptionsPagesWithContext(ctx context.Context, input *ec2.DescribeDhcpOptionsInput, fn func(*ec2.DescribeDhcpOptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeDhcpOptionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeDhcpOptionsOutput{}
	fnCacher := func(out *ec2.DescribeDhcpOptionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeDhcpOptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeDhcpOptionsWithContext(ctx context.Context, input *ec2.DescribeDhcpOptionsInput, opts ...request.Option) (*ec2.DescribeDhcpOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeDhcpOptionsOutput), nil
	}
	out, err := c.EC2API.DescribeDhcpOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeEgressOnlyInternetGateways(input *ec2.DescribeEgressOnlyInternetGatewaysInput) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeEgressOnlyInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysPages(input *ec2.DescribeEgressOnlyInternetGatewaysInput, fn func(*ec2.DescribeEgressOnlyInternetGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeEgressOnlyInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeEgressOnlyInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput, fn func(*ec2.DescribeEgressOnlyInternetGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeEgressOnlyInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeEgressOnlyInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeEgressOnlyInternetGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeEgressOnlyInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeEgressOnlyInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeEgressOnlyInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeEgressOnlyInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeEgressOnlyInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeElasticGpus(input *ec2.DescribeElasticGpusInput) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.EC2API.DescribeElasticGpus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeElasticGpusWithContext(ctx context.Context, input *ec2.DescribeElasticGpusInput, opts ...request.Option) (*ec2.DescribeElasticGpusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeElasticGpusOutput), nil
	}
	out, err := c.EC2API.DescribeElasticGpusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportImageTasks(input *ec2.DescribeExportImageTasksInput) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportImageTasksPages(input *ec2.DescribeExportImageTasksInput, fn func(*ec2.DescribeExportImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeExportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeExportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeExportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeExportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeExportImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeExportImageTasksInput, fn func(*ec2.DescribeExportImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeExportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeExportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeExportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeExportImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeExportImageTasksWithContext(ctx context.Context, input *ec2.DescribeExportImageTasksInput, opts ...request.Option) (*ec2.DescribeExportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportTasks(input *ec2.DescribeExportTasksInput) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeExportTasksWithContext(ctx context.Context, input *ec2.DescribeExportTasksInput, opts ...request.Option) (*ec2.DescribeExportTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeExportTasksOutput), nil
	}
	out, err := c.EC2API.DescribeExportTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastLaunchImages(input *ec2.DescribeFastLaunchImagesInput) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFastLaunchImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastLaunchImagesPages(input *ec2.DescribeFastLaunchImagesInput, fn func(*ec2.DescribeFastLaunchImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastLaunchImagesOutput{}
	fnCacher := func(out *ec2.DescribeFastLaunchImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastLaunchImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastLaunchImagesPagesWithContext(ctx context.Context, input *ec2.DescribeFastLaunchImagesInput, fn func(*ec2.DescribeFastLaunchImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastLaunchImagesOutput{}
	fnCacher := func(out *ec2.DescribeFastLaunchImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastLaunchImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastLaunchImagesWithContext(ctx context.Context, input *ec2.DescribeFastLaunchImagesInput, opts ...request.Option) (*ec2.DescribeFastLaunchImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastLaunchImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFastLaunchImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastSnapshotRestores(input *ec2.DescribeFastSnapshotRestoresInput) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DescribeFastSnapshotRestores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFastSnapshotRestoresPages(input *ec2.DescribeFastSnapshotRestoresInput, fn func(*ec2.DescribeFastSnapshotRestoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastSnapshotRestoresOutput{}
	fnCacher := func(out *ec2.DescribeFastSnapshotRestoresOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastSnapshotRestoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastSnapshotRestoresPagesWithContext(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput, fn func(*ec2.DescribeFastSnapshotRestoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFastSnapshotRestoresOutput{}
	fnCacher := func(out *ec2.DescribeFastSnapshotRestoresOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFastSnapshotRestoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFastSnapshotRestoresWithContext(ctx context.Context, input *ec2.DescribeFastSnapshotRestoresInput, opts ...request.Option) (*ec2.DescribeFastSnapshotRestoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFastSnapshotRestoresOutput), nil
	}
	out, err := c.EC2API.DescribeFastSnapshotRestoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetHistory(input *ec2.DescribeFleetHistoryInput) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeFleetHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetHistoryWithContext(ctx context.Context, input *ec2.DescribeFleetHistoryInput, opts ...request.Option) (*ec2.DescribeFleetHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeFleetHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetInstances(input *ec2.DescribeFleetInstancesInput) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeFleetInstancesInput, opts ...request.Option) (*ec2.DescribeFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleets(input *ec2.DescribeFleetsInput) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFleetsPages(input *ec2.DescribeFleetsInput, fn func(*ec2.DescribeFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFleetsOutput{}
	fnCacher := func(out *ec2.DescribeFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFleetsPagesWithContext(ctx context.Context, input *ec2.DescribeFleetsInput, fn func(*ec2.DescribeFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFleetsOutput{}
	fnCacher := func(out *ec2.DescribeFleetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFleetsWithContext(ctx context.Context, input *ec2.DescribeFleetsInput, opts ...request.Option) (*ec2.DescribeFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFleetsOutput), nil
	}
	out, err := c.EC2API.DescribeFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFlowLogs(input *ec2.DescribeFlowLogsInput) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.EC2API.DescribeFlowLogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFlowLogsPages(input *ec2.DescribeFlowLogsInput, fn func(*ec2.DescribeFlowLogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFlowLogsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFlowLogsOutput{}
	fnCacher := func(out *ec2.DescribeFlowLogsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFlowLogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFlowLogsPagesWithContext(ctx context.Context, input *ec2.DescribeFlowLogsInput, fn func(*ec2.DescribeFlowLogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFlowLogsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFlowLogsOutput{}
	fnCacher := func(out *ec2.DescribeFlowLogsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFlowLogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFlowLogsWithContext(ctx context.Context, input *ec2.DescribeFlowLogsInput, opts ...request.Option) (*ec2.DescribeFlowLogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFlowLogsOutput), nil
	}
	out, err := c.EC2API.DescribeFlowLogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImageAttribute(input *ec2.DescribeFpgaImageAttributeInput) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImageAttributeWithContext(ctx context.Context, input *ec2.DescribeFpgaImageAttributeInput, opts ...request.Option) (*ec2.DescribeFpgaImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImages(input *ec2.DescribeFpgaImagesInput) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeFpgaImagesPages(input *ec2.DescribeFpgaImagesInput, fn func(*ec2.DescribeFpgaImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFpgaImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFpgaImagesOutput{}
	fnCacher := func(out *ec2.DescribeFpgaImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFpgaImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFpgaImagesPagesWithContext(ctx context.Context, input *ec2.DescribeFpgaImagesInput, fn func(*ec2.DescribeFpgaImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeFpgaImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeFpgaImagesOutput{}
	fnCacher := func(out *ec2.DescribeFpgaImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeFpgaImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeFpgaImagesWithContext(ctx context.Context, input *ec2.DescribeFpgaImagesInput, opts ...request.Option) (*ec2.DescribeFpgaImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeFpgaImagesOutput), nil
	}
	out, err := c.EC2API.DescribeFpgaImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationOfferings(input *ec2.DescribeHostReservationOfferingsInput) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationOfferingsPages(input *ec2.DescribeHostReservationOfferingsInput, fn func(*ec2.DescribeHostReservationOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput, fn func(*ec2.DescribeHostReservationOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationOfferingsWithContext(ctx context.Context, input *ec2.DescribeHostReservationOfferingsInput, opts ...request.Option) (*ec2.DescribeHostReservationOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservations(input *ec2.DescribeHostReservationsInput) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostReservationsPages(input *ec2.DescribeHostReservationsInput, fn func(*ec2.DescribeHostReservationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationsPagesWithContext(ctx context.Context, input *ec2.DescribeHostReservationsInput, fn func(*ec2.DescribeHostReservationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostReservationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostReservationsOutput{}
	fnCacher := func(out *ec2.DescribeHostReservationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostReservationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostReservationsWithContext(ctx context.Context, input *ec2.DescribeHostReservationsInput, opts ...request.Option) (*ec2.DescribeHostReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostReservationsOutput), nil
	}
	out, err := c.EC2API.DescribeHostReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHosts(input *ec2.DescribeHostsInput) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.EC2API.DescribeHosts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeHostsPages(input *ec2.DescribeHostsInput, fn func(*ec2.DescribeHostsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostsOutput{}
	fnCacher := func(out *ec2.DescribeHostsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostsPagesWithContext(ctx context.Context, input *ec2.DescribeHostsInput, fn func(*ec2.DescribeHostsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeHostsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeHostsOutput{}
	fnCacher := func(out *ec2.DescribeHostsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeHostsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeHostsWithContext(ctx context.Context, input *ec2.DescribeHostsInput, opts ...request.Option) (*ec2.DescribeHostsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeHostsOutput), nil
	}
	out, err := c.EC2API.DescribeHostsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIamInstanceProfileAssociations(input *ec2.DescribeIamInstanceProfileAssociationsInput) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeIamInstanceProfileAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIamInstanceProfileAssociationsPages(input *ec2.DescribeIamInstanceProfileAssociationsInput, fn func(*ec2.DescribeIamInstanceProfileAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIamInstanceProfileAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeIamInstanceProfileAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIamInstanceProfileAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIamInstanceProfileAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput, fn func(*ec2.DescribeIamInstanceProfileAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIamInstanceProfileAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeIamInstanceProfileAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIamInstanceProfileAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIamInstanceProfileAssociationsWithContext(ctx context.Context, input *ec2.DescribeIamInstanceProfileAssociationsInput, opts ...request.Option) (*ec2.DescribeIamInstanceProfileAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIamInstanceProfileAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeIamInstanceProfileAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdFormat(input *ec2.DescribeIdFormatInput) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdFormatInput, opts ...request.Option) (*ec2.DescribeIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdentityIdFormat(input *ec2.DescribeIdentityIdFormatInput) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdentityIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIdentityIdFormatWithContext(ctx context.Context, input *ec2.DescribeIdentityIdFormatInput, opts ...request.Option) (*ec2.DescribeIdentityIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIdentityIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribeIdentityIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImageAttribute(input *ec2.DescribeImageAttributeInput) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeImageAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImageAttributeWithContext(ctx context.Context, input *ec2.DescribeImageAttributeInput, opts ...request.Option) (*ec2.DescribeImageAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImageAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeImageAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImages(input *ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.EC2API.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImagesPages(input *ec2.DescribeImagesInput, fn func(*ec2.DescribeImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImagesOutput{}
	fnCacher := func(out *ec2.DescribeImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImagesPagesWithContext(ctx context.Context, input *ec2.DescribeImagesInput, fn func(*ec2.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImagesOutput{}
	fnCacher := func(out *ec2.DescribeImagesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImagesWithContext(ctx context.Context, input *ec2.DescribeImagesInput, opts ...request.Option) (*ec2.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImagesOutput), nil
	}
	out, err := c.EC2API.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportImageTasks(input *ec2.DescribeImportImageTasksInput) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportImageTasksPages(input *ec2.DescribeImportImageTasksInput, fn func(*ec2.DescribeImportImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeImportImageTasksInput, fn func(*ec2.DescribeImportImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportImageTasksWithContext(ctx context.Context, input *ec2.DescribeImportImageTasksInput, opts ...request.Option) (*ec2.DescribeImportImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportSnapshotTasks(input *ec2.DescribeImportSnapshotTasksInput) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportSnapshotTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeImportSnapshotTasksPages(input *ec2.DescribeImportSnapshotTasksInput, fn func(*ec2.DescribeImportSnapshotTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportSnapshotTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportSnapshotTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportSnapshotTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportSnapshotTasksPagesWithContext(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput, fn func(*ec2.DescribeImportSnapshotTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeImportSnapshotTasksOutput{}
	fnCacher := func(out *ec2.DescribeImportSnapshotTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeImportSnapshotTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeImportSnapshotTasksWithContext(ctx context.Context, input *ec2.DescribeImportSnapshotTasksInput, opts ...request.Option) (*ec2.DescribeImportSnapshotTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeImportSnapshotTasksOutput), nil
	}
	out, err := c.EC2API.DescribeImportSnapshotTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceAttribute(input *ec2.DescribeInstanceAttributeInput) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceAttributeWithContext(ctx context.Context, input *ec2.DescribeInstanceAttributeInput, opts ...request.Option) (*ec2.DescribeInstanceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceCreditSpecifications(input *ec2.DescribeInstanceCreditSpecificationsInput) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceCreditSpecifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceCreditSpecificationsPages(input *ec2.DescribeInstanceCreditSpecificationsInput, fn func(*ec2.DescribeInstanceCreditSpecificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceCreditSpecificationsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceCreditSpecificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceCreditSpecificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceCreditSpecificationsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput, fn func(*ec2.DescribeInstanceCreditSpecificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceCreditSpecificationsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceCreditSpecificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceCreditSpecificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceCreditSpecificationsWithContext(ctx context.Context, input *ec2.DescribeInstanceCreditSpecificationsInput, opts ...request.Option) (*ec2.DescribeInstanceCreditSpecificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceCreditSpecificationsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceCreditSpecificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventNotificationAttributes(input *ec2.DescribeInstanceEventNotificationAttributesInput) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventNotificationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventNotificationAttributesWithContext(ctx context.Context, input *ec2.DescribeInstanceEventNotificationAttributesInput, opts ...request.Option) (*ec2.DescribeInstanceEventNotificationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventNotificationAttributesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventNotificationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventWindows(input *ec2.DescribeInstanceEventWindowsInput) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventWindows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceEventWindowsPages(input *ec2.DescribeInstanceEventWindowsInput, fn func(*ec2.DescribeInstanceEventWindowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceEventWindowsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceEventWindowsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceEventWindowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceEventWindowsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceEventWindowsInput, fn func(*ec2.DescribeInstanceEventWindowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceEventWindowsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceEventWindowsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceEventWindowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceEventWindowsWithContext(ctx context.Context, input *ec2.DescribeInstanceEventWindowsInput, opts ...request.Option) (*ec2.DescribeInstanceEventWindowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceEventWindowsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceEventWindowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceStatus(input *ec2.DescribeInstanceStatusInput) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceStatusPages(input *ec2.DescribeInstanceStatusInput, fn func(*ec2.DescribeInstanceStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceStatusOutput{}
	fnCacher := func(out *ec2.DescribeInstanceStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceStatusPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceStatusInput, fn func(*ec2.DescribeInstanceStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceStatusOutput{}
	fnCacher := func(out *ec2.DescribeInstanceStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceStatusWithContext(ctx context.Context, input *ec2.DescribeInstanceStatusInput, opts ...request.Option) (*ec2.DescribeInstanceStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceStatusOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypeOfferings(input *ec2.DescribeInstanceTypeOfferingsInput) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypeOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypeOfferingsPages(input *ec2.DescribeInstanceTypeOfferingsInput, fn func(*ec2.DescribeInstanceTypeOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypeOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypeOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypeOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypeOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput, fn func(*ec2.DescribeInstanceTypeOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypeOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypeOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypeOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypeOfferingsWithContext(ctx context.Context, input *ec2.DescribeInstanceTypeOfferingsInput, opts ...request.Option) (*ec2.DescribeInstanceTypeOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypeOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypeOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypes(input *ec2.DescribeInstanceTypesInput) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstanceTypesPages(input *ec2.DescribeInstanceTypesInput, fn func(*ec2.DescribeInstanceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypesOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypesPagesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypesInput, fn func(*ec2.DescribeInstanceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstanceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstanceTypesOutput{}
	fnCacher := func(out *ec2.DescribeInstanceTypesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstanceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstanceTypesWithContext(ctx context.Context, input *ec2.DescribeInstanceTypesInput, opts ...request.Option) (*ec2.DescribeInstanceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstanceTypesOutput), nil
	}
	out, err := c.EC2API.DescribeInstanceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstances(input *ec2.DescribeInstancesInput) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInstancesPages(input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstancesOutput{}
	fnCacher := func(out *ec2.DescribeInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, fn func(*ec2.DescribeInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInstancesOutput{}
	fnCacher := func(out *ec2.DescribeInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInstancesWithContext(ctx context.Context, input *ec2.DescribeInstancesInput, opts ...request.Option) (*ec2.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInternetGateways(input *ec2.DescribeInternetGatewaysInput) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeInternetGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeInternetGatewaysPages(input *ec2.DescribeInternetGatewaysInput, fn func(*ec2.DescribeInternetGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInternetGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInternetGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeInternetGatewaysInput, fn func(*ec2.DescribeInternetGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeInternetGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeInternetGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeInternetGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeInternetGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeInternetGatewaysWithContext(ctx context.Context, input *ec2.DescribeInternetGatewaysInput, opts ...request.Option) (*ec2.DescribeInternetGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeInternetGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeInternetGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamPools(input *ec2.DescribeIpamPoolsInput) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamPools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamPoolsPages(input *ec2.DescribeIpamPoolsInput, fn func(*ec2.DescribeIpamPoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamPoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpamPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamPoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamPoolsPagesWithContext(ctx context.Context, input *ec2.DescribeIpamPoolsInput, fn func(*ec2.DescribeIpamPoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamPoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamPoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpamPoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamPoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamPoolsWithContext(ctx context.Context, input *ec2.DescribeIpamPoolsInput, opts ...request.Option) (*ec2.DescribeIpamPoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamPoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamPoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamScopes(input *ec2.DescribeIpamScopesInput) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.EC2API.DescribeIpamScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamScopesPages(input *ec2.DescribeIpamScopesInput, fn func(*ec2.DescribeIpamScopesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamScopesOutput{}
	fnCacher := func(out *ec2.DescribeIpamScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamScopesPagesWithContext(ctx context.Context, input *ec2.DescribeIpamScopesInput, fn func(*ec2.DescribeIpamScopesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamScopesOutput{}
	fnCacher := func(out *ec2.DescribeIpamScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamScopesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamScopesWithContext(ctx context.Context, input *ec2.DescribeIpamScopesInput, opts ...request.Option) (*ec2.DescribeIpamScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamScopesOutput), nil
	}
	out, err := c.EC2API.DescribeIpamScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpams(input *ec2.DescribeIpamsInput) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.EC2API.DescribeIpams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpamsPages(input *ec2.DescribeIpamsInput, fn func(*ec2.DescribeIpamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamsOutput{}
	fnCacher := func(out *ec2.DescribeIpamsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamsPagesWithContext(ctx context.Context, input *ec2.DescribeIpamsInput, fn func(*ec2.DescribeIpamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpamsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpamsOutput{}
	fnCacher := func(out *ec2.DescribeIpamsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpamsWithContext(ctx context.Context, input *ec2.DescribeIpamsInput, opts ...request.Option) (*ec2.DescribeIpamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpamsOutput), nil
	}
	out, err := c.EC2API.DescribeIpamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpv6Pools(input *ec2.DescribeIpv6PoolsInput) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpv6Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeIpv6PoolsPages(input *ec2.DescribeIpv6PoolsInput, fn func(*ec2.DescribeIpv6PoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpv6PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpv6PoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpv6PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpv6PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpv6PoolsPagesWithContext(ctx context.Context, input *ec2.DescribeIpv6PoolsInput, fn func(*ec2.DescribeIpv6PoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeIpv6PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeIpv6PoolsOutput{}
	fnCacher := func(out *ec2.DescribeIpv6PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeIpv6PoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeIpv6PoolsWithContext(ctx context.Context, input *ec2.DescribeIpv6PoolsInput, opts ...request.Option) (*ec2.DescribeIpv6PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeIpv6PoolsOutput), nil
	}
	out, err := c.EC2API.DescribeIpv6PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeKeyPairs(input *ec2.DescribeKeyPairsInput) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.EC2API.DescribeKeyPairs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeKeyPairsWithContext(ctx context.Context, input *ec2.DescribeKeyPairsInput, opts ...request.Option) (*ec2.DescribeKeyPairsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeKeyPairsOutput), nil
	}
	out, err := c.EC2API.DescribeKeyPairsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplateVersions(input *ec2.DescribeLaunchTemplateVersionsInput) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplateVersionsPages(input *ec2.DescribeLaunchTemplateVersionsInput, fn func(*ec2.DescribeLaunchTemplateVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplateVersionsOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplateVersionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplateVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplateVersionsPagesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput, fn func(*ec2.DescribeLaunchTemplateVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplateVersionsOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplateVersionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplateVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplateVersionsWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplateVersionsInput, opts ...request.Option) (*ec2.DescribeLaunchTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplateVersionsOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplates(input *ec2.DescribeLaunchTemplatesInput) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLaunchTemplatesPages(input *ec2.DescribeLaunchTemplatesInput, fn func(*ec2.DescribeLaunchTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplatesOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplatesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplatesPagesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput, fn func(*ec2.DescribeLaunchTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLaunchTemplatesOutput{}
	fnCacher := func(out *ec2.DescribeLaunchTemplatesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLaunchTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLaunchTemplatesWithContext(ctx context.Context, input *ec2.DescribeLaunchTemplatesInput, opts ...request.Option) (*ec2.DescribeLaunchTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLaunchTemplatesOutput), nil
	}
	out, err := c.EC2API.DescribeLaunchTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVirtualInterfaceGroupAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociations(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsPages(input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, fn func(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTableVpcAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTables(input *ec2.DescribeLocalGatewayRouteTablesInput) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayRouteTablesPages(input *ec2.DescribeLocalGatewayRouteTablesInput, fn func(*ec2.DescribeLocalGatewayRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput, fn func(*ec2.DescribeLocalGatewayRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroups(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsPages(input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfaceGroupsInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfaceGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaceGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfaces(input *ec2.DescribeLocalGatewayVirtualInterfacesInput) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesPages(input *ec2.DescribeLocalGatewayVirtualInterfacesInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput, fn func(*ec2.DescribeLocalGatewayVirtualInterfacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewayVirtualInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewayVirtualInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewayVirtualInterfacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewayVirtualInterfacesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewayVirtualInterfacesInput, opts ...request.Option) (*ec2.DescribeLocalGatewayVirtualInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewayVirtualInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewayVirtualInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGateways(input *ec2.DescribeLocalGatewaysInput) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeLocalGatewaysPages(input *ec2.DescribeLocalGatewaysInput, fn func(*ec2.DescribeLocalGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeLocalGatewaysInput, fn func(*ec2.DescribeLocalGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeLocalGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeLocalGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeLocalGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeLocalGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeLocalGatewaysWithContext(ctx context.Context, input *ec2.DescribeLocalGatewaysInput, opts ...request.Option) (*ec2.DescribeLocalGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeLocalGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeLocalGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeManagedPrefixLists(input *ec2.DescribeManagedPrefixListsInput) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribeManagedPrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeManagedPrefixListsPages(input *ec2.DescribeManagedPrefixListsInput, fn func(*ec2.DescribeManagedPrefixListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeManagedPrefixListsOutput{}
	fnCacher := func(out *ec2.DescribeManagedPrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeManagedPrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeManagedPrefixListsPagesWithContext(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput, fn func(*ec2.DescribeManagedPrefixListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeManagedPrefixListsOutput{}
	fnCacher := func(out *ec2.DescribeManagedPrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeManagedPrefixListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeManagedPrefixListsWithContext(ctx context.Context, input *ec2.DescribeManagedPrefixListsInput, opts ...request.Option) (*ec2.DescribeManagedPrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeManagedPrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribeManagedPrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeMovingAddresses(input *ec2.DescribeMovingAddressesInput) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeMovingAddresses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeMovingAddressesPages(input *ec2.DescribeMovingAddressesInput, fn func(*ec2.DescribeMovingAddressesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeMovingAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeMovingAddressesOutput{}
	fnCacher := func(out *ec2.DescribeMovingAddressesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeMovingAddressesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeMovingAddressesPagesWithContext(ctx context.Context, input *ec2.DescribeMovingAddressesInput, fn func(*ec2.DescribeMovingAddressesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeMovingAddressesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeMovingAddressesOutput{}
	fnCacher := func(out *ec2.DescribeMovingAddressesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeMovingAddressesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeMovingAddressesWithContext(ctx context.Context, input *ec2.DescribeMovingAddressesInput, opts ...request.Option) (*ec2.DescribeMovingAddressesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeMovingAddressesOutput), nil
	}
	out, err := c.EC2API.DescribeMovingAddressesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNatGateways(input *ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeNatGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNatGatewaysPages(input *ec2.DescribeNatGatewaysInput, fn func(*ec2.DescribeNatGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNatGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNatGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeNatGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNatGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNatGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeNatGatewaysInput, fn func(*ec2.DescribeNatGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNatGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNatGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeNatGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNatGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNatGatewaysWithContext(ctx context.Context, input *ec2.DescribeNatGatewaysInput, opts ...request.Option) (*ec2.DescribeNatGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNatGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeNatGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkAcls(input *ec2.DescribeNetworkAclsInput) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkAcls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkAclsPages(input *ec2.DescribeNetworkAclsInput, fn func(*ec2.DescribeNetworkAclsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkAclsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkAclsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkAclsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkAclsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkAclsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkAclsInput, fn func(*ec2.DescribeNetworkAclsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkAclsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkAclsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkAclsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkAclsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkAclsWithContext(ctx context.Context, input *ec2.DescribeNetworkAclsInput, opts ...request.Option) (*ec2.DescribeNetworkAclsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkAclsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkAclsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalyses(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesPages(input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopeAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopeAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopeAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopes(input *ec2.DescribeNetworkInsightsAccessScopesInput) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAccessScopesPages(input *ec2.DescribeNetworkInsightsAccessScopesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopesInput, fn func(*ec2.DescribeNetworkInsightsAccessScopesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAccessScopesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAccessScopesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAccessScopesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAccessScopesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAccessScopesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAccessScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAccessScopesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAccessScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAnalyses(input *ec2.DescribeNetworkInsightsAnalysesInput) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsAnalysesPages(input *ec2.DescribeNetworkInsightsAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAnalysesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAnalysesInput, fn func(*ec2.DescribeNetworkInsightsAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsAnalysesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsAnalysesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsAnalysesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsAnalysesInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsAnalysesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsPaths(input *ec2.DescribeNetworkInsightsPathsInput) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsPaths(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInsightsPathsPages(input *ec2.DescribeNetworkInsightsPathsInput, fn func(*ec2.DescribeNetworkInsightsPathsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsPathsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsPathsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsPathsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsPathsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsPathsInput, fn func(*ec2.DescribeNetworkInsightsPathsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInsightsPathsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInsightsPathsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInsightsPathsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInsightsPathsWithContext(ctx context.Context, input *ec2.DescribeNetworkInsightsPathsInput, opts ...request.Option) (*ec2.DescribeNetworkInsightsPathsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInsightsPathsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInsightsPathsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaceAttribute(input *ec2.DescribeNetworkInterfaceAttributeInput) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaceAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaceAttributeWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfaceAttributeInput, opts ...request.Option) (*ec2.DescribeNetworkInterfaceAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfaceAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaceAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacePermissions(input *ec2.DescribeNetworkInterfacePermissionsInput) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacePermissionsPages(input *ec2.DescribeNetworkInterfacePermissionsInput, fn func(*ec2.DescribeNetworkInterfacePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacePermissionsPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput, fn func(*ec2.DescribeNetworkInterfacePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacePermissionsWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacePermissionsInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfaces(input *ec2.DescribeNetworkInterfacesInput) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeNetworkInterfacesPages(input *ec2.DescribeNetworkInterfacesInput, fn func(*ec2.DescribeNetworkInterfacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacesPagesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput, fn func(*ec2.DescribeNetworkInterfacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeNetworkInterfacesOutput{}
	fnCacher := func(out *ec2.DescribeNetworkInterfacesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeNetworkInterfacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeNetworkInterfacesWithContext(ctx context.Context, input *ec2.DescribeNetworkInterfacesInput, opts ...request.Option) (*ec2.DescribeNetworkInterfacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeNetworkInterfacesOutput), nil
	}
	out, err := c.EC2API.DescribeNetworkInterfacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePlacementGroups(input *ec2.DescribePlacementGroupsInput) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.EC2API.DescribePlacementGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePlacementGroupsWithContext(ctx context.Context, input *ec2.DescribePlacementGroupsInput, opts ...request.Option) (*ec2.DescribePlacementGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePlacementGroupsOutput), nil
	}
	out, err := c.EC2API.DescribePlacementGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrefixLists(input *ec2.DescribePrefixListsInput) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribePrefixLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrefixListsPages(input *ec2.DescribePrefixListsInput, fn func(*ec2.DescribePrefixListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrefixListsOutput{}
	fnCacher := func(out *ec2.DescribePrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrefixListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrefixListsPagesWithContext(ctx context.Context, input *ec2.DescribePrefixListsInput, fn func(*ec2.DescribePrefixListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrefixListsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrefixListsOutput{}
	fnCacher := func(out *ec2.DescribePrefixListsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrefixListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrefixListsWithContext(ctx context.Context, input *ec2.DescribePrefixListsInput, opts ...request.Option) (*ec2.DescribePrefixListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrefixListsOutput), nil
	}
	out, err := c.EC2API.DescribePrefixListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrincipalIdFormat(input *ec2.DescribePrincipalIdFormatInput) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribePrincipalIdFormat(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePrincipalIdFormatPages(input *ec2.DescribePrincipalIdFormatInput, fn func(*ec2.DescribePrincipalIdFormatOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrincipalIdFormatOutput{}
	fnCacher := func(out *ec2.DescribePrincipalIdFormatOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrincipalIdFormatPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrincipalIdFormatPagesWithContext(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput, fn func(*ec2.DescribePrincipalIdFormatOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePrincipalIdFormatOutput{}
	fnCacher := func(out *ec2.DescribePrincipalIdFormatOutput, more bool) bool {
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
	if err := c.EC2API.DescribePrincipalIdFormatPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePrincipalIdFormatWithContext(ctx context.Context, input *ec2.DescribePrincipalIdFormatInput, opts ...request.Option) (*ec2.DescribePrincipalIdFormatOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePrincipalIdFormatOutput), nil
	}
	out, err := c.EC2API.DescribePrincipalIdFormatWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePublicIpv4Pools(input *ec2.DescribePublicIpv4PoolsInput) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.EC2API.DescribePublicIpv4Pools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribePublicIpv4PoolsPages(input *ec2.DescribePublicIpv4PoolsInput, fn func(*ec2.DescribePublicIpv4PoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePublicIpv4PoolsOutput{}
	fnCacher := func(out *ec2.DescribePublicIpv4PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePublicIpv4PoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePublicIpv4PoolsPagesWithContext(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput, fn func(*ec2.DescribePublicIpv4PoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribePublicIpv4PoolsOutput{}
	fnCacher := func(out *ec2.DescribePublicIpv4PoolsOutput, more bool) bool {
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
	if err := c.EC2API.DescribePublicIpv4PoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribePublicIpv4PoolsWithContext(ctx context.Context, input *ec2.DescribePublicIpv4PoolsInput, opts ...request.Option) (*ec2.DescribePublicIpv4PoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribePublicIpv4PoolsOutput), nil
	}
	out, err := c.EC2API.DescribePublicIpv4PoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRegions(input *ec2.DescribeRegionsInput) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.EC2API.DescribeRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRegionsWithContext(ctx context.Context, input *ec2.DescribeRegionsInput, opts ...request.Option) (*ec2.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRegionsOutput), nil
	}
	out, err := c.EC2API.DescribeRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReplaceRootVolumeTasks(input *ec2.DescribeReplaceRootVolumeTasksInput) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.EC2API.DescribeReplaceRootVolumeTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReplaceRootVolumeTasksPages(input *ec2.DescribeReplaceRootVolumeTasksInput, fn func(*ec2.DescribeReplaceRootVolumeTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReplaceRootVolumeTasksOutput{}
	fnCacher := func(out *ec2.DescribeReplaceRootVolumeTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReplaceRootVolumeTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReplaceRootVolumeTasksPagesWithContext(ctx context.Context, input *ec2.DescribeReplaceRootVolumeTasksInput, fn func(*ec2.DescribeReplaceRootVolumeTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReplaceRootVolumeTasksOutput{}
	fnCacher := func(out *ec2.DescribeReplaceRootVolumeTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReplaceRootVolumeTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReplaceRootVolumeTasksWithContext(ctx context.Context, input *ec2.DescribeReplaceRootVolumeTasksInput, opts ...request.Option) (*ec2.DescribeReplaceRootVolumeTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReplaceRootVolumeTasksOutput), nil
	}
	out, err := c.EC2API.DescribeReplaceRootVolumeTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstances(input *ec2.DescribeReservedInstancesInput) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesListings(input *ec2.DescribeReservedInstancesListingsInput) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesListings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesListingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesListingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesListingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesListingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesListingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesModifications(input *ec2.DescribeReservedInstancesModificationsInput) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesModificationsPages(input *ec2.DescribeReservedInstancesModificationsInput, fn func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesModificationsPagesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput, fn func(*ec2.DescribeReservedInstancesModificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesModificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesModificationsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesModificationsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesOfferings(input *ec2.DescribeReservedInstancesOfferingsInput) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesOfferings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesOfferingsPages(input *ec2.DescribeReservedInstancesOfferingsInput, fn func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesOfferingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesOfferingsPagesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput, fn func(*ec2.DescribeReservedInstancesOfferingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeReservedInstancesOfferingsOutput{}
	fnCacher := func(out *ec2.DescribeReservedInstancesOfferingsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeReservedInstancesOfferingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeReservedInstancesOfferingsWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesOfferingsInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOfferingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOfferingsOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesOfferingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeReservedInstancesWithContext(ctx context.Context, input *ec2.DescribeReservedInstancesInput, opts ...request.Option) (*ec2.DescribeReservedInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeReservedInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeReservedInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRouteTables(input *ec2.DescribeRouteTablesInput) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeRouteTablesPages(input *ec2.DescribeRouteTablesInput, fn func(*ec2.DescribeRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeRouteTablesInput, fn func(*ec2.DescribeRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeRouteTablesWithContext(ctx context.Context, input *ec2.DescribeRouteTablesInput, opts ...request.Option) (*ec2.DescribeRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstanceAvailability(input *ec2.DescribeScheduledInstanceAvailabilityInput) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstanceAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstanceAvailabilityPages(input *ec2.DescribeScheduledInstanceAvailabilityInput, fn func(*ec2.DescribeScheduledInstanceAvailabilityOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstanceAvailabilityOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstanceAvailabilityOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstanceAvailabilityPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstanceAvailabilityPagesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput, fn func(*ec2.DescribeScheduledInstanceAvailabilityOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstanceAvailabilityOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstanceAvailabilityOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstanceAvailabilityPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstanceAvailabilityWithContext(ctx context.Context, input *ec2.DescribeScheduledInstanceAvailabilityInput, opts ...request.Option) (*ec2.DescribeScheduledInstanceAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstanceAvailabilityOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstanceAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstances(input *ec2.DescribeScheduledInstancesInput) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeScheduledInstancesPages(input *ec2.DescribeScheduledInstancesInput, fn func(*ec2.DescribeScheduledInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstancesOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstancesInput, fn func(*ec2.DescribeScheduledInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeScheduledInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeScheduledInstancesOutput{}
	fnCacher := func(out *ec2.DescribeScheduledInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeScheduledInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeScheduledInstancesWithContext(ctx context.Context, input *ec2.DescribeScheduledInstancesInput, opts ...request.Option) (*ec2.DescribeScheduledInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeScheduledInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeScheduledInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupReferences(input *ec2.DescribeSecurityGroupReferencesInput) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupReferencesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupReferencesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupReferencesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupRules(input *ec2.DescribeSecurityGroupRulesInput) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupRulesPages(input *ec2.DescribeSecurityGroupRulesInput, fn func(*ec2.DescribeSecurityGroupRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupRulesOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupRulesPagesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupRulesInput, fn func(*ec2.DescribeSecurityGroupRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupRulesOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupRulesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupRulesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupRulesInput, opts ...request.Option) (*ec2.DescribeSecurityGroupRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupRulesOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroups(input *ec2.DescribeSecurityGroupsInput) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSecurityGroupsPages(input *ec2.DescribeSecurityGroupsInput, fn func(*ec2.DescribeSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupsInput, fn func(*ec2.DescribeSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotAttribute(input *ec2.DescribeSnapshotAttributeInput) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotAttributeWithContext(ctx context.Context, input *ec2.DescribeSnapshotAttributeInput, opts ...request.Option) (*ec2.DescribeSnapshotAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotTierStatus(input *ec2.DescribeSnapshotTierStatusInput) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotTierStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotTierStatusPages(input *ec2.DescribeSnapshotTierStatusInput, fn func(*ec2.DescribeSnapshotTierStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotTierStatusOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotTierStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotTierStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotTierStatusPagesWithContext(ctx context.Context, input *ec2.DescribeSnapshotTierStatusInput, fn func(*ec2.DescribeSnapshotTierStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotTierStatusOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotTierStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotTierStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotTierStatusWithContext(ctx context.Context, input *ec2.DescribeSnapshotTierStatusInput, opts ...request.Option) (*ec2.DescribeSnapshotTierStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotTierStatusOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotTierStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshots(input *ec2.DescribeSnapshotsInput) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSnapshotsPages(input *ec2.DescribeSnapshotsInput, fn func(*ec2.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotsOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotsPagesWithContext(ctx context.Context, input *ec2.DescribeSnapshotsInput, fn func(*ec2.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSnapshotsOutput{}
	fnCacher := func(out *ec2.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSnapshotsWithContext(ctx context.Context, input *ec2.DescribeSnapshotsInput, opts ...request.Option) (*ec2.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSnapshotsOutput), nil
	}
	out, err := c.EC2API.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotDatafeedSubscription(input *ec2.DescribeSpotDatafeedSubscriptionInput) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DescribeSpotDatafeedSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotDatafeedSubscriptionWithContext(ctx context.Context, input *ec2.DescribeSpotDatafeedSubscriptionInput, opts ...request.Option) (*ec2.DescribeSpotDatafeedSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotDatafeedSubscriptionOutput), nil
	}
	out, err := c.EC2API.DescribeSpotDatafeedSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetInstances(input *ec2.DescribeSpotFleetInstancesInput) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetInstancesWithContext(ctx context.Context, input *ec2.DescribeSpotFleetInstancesInput, opts ...request.Option) (*ec2.DescribeSpotFleetInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestHistory(input *ec2.DescribeSpotFleetRequestHistoryInput) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestHistoryInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequests(input *ec2.DescribeSpotFleetRequestsInput) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotFleetRequestsPages(input *ec2.DescribeSpotFleetRequestsInput, fn func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotFleetRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotFleetRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotFleetRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotFleetRequestsPagesWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput, fn func(*ec2.DescribeSpotFleetRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotFleetRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotFleetRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotFleetRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotFleetRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotFleetRequestsInput, opts ...request.Option) (*ec2.DescribeSpotFleetRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotFleetRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotFleetRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotInstanceRequests(input *ec2.DescribeSpotInstanceRequestsInput) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotInstanceRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotInstanceRequestsPages(input *ec2.DescribeSpotInstanceRequestsInput, fn func(*ec2.DescribeSpotInstanceRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotInstanceRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotInstanceRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotInstanceRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotInstanceRequestsPagesWithContext(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput, fn func(*ec2.DescribeSpotInstanceRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotInstanceRequestsOutput{}
	fnCacher := func(out *ec2.DescribeSpotInstanceRequestsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotInstanceRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotInstanceRequestsWithContext(ctx context.Context, input *ec2.DescribeSpotInstanceRequestsInput, opts ...request.Option) (*ec2.DescribeSpotInstanceRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotInstanceRequestsOutput), nil
	}
	out, err := c.EC2API.DescribeSpotInstanceRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotPriceHistory(input *ec2.DescribeSpotPriceHistoryInput) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotPriceHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSpotPriceHistoryPages(input *ec2.DescribeSpotPriceHistoryInput, fn func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotPriceHistoryOutput{}
	fnCacher := func(out *ec2.DescribeSpotPriceHistoryOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotPriceHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotPriceHistoryPagesWithContext(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput, fn func(*ec2.DescribeSpotPriceHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSpotPriceHistoryOutput{}
	fnCacher := func(out *ec2.DescribeSpotPriceHistoryOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSpotPriceHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSpotPriceHistoryWithContext(ctx context.Context, input *ec2.DescribeSpotPriceHistoryInput, opts ...request.Option) (*ec2.DescribeSpotPriceHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSpotPriceHistoryOutput), nil
	}
	out, err := c.EC2API.DescribeSpotPriceHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStaleSecurityGroups(input *ec2.DescribeStaleSecurityGroupsInput) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeStaleSecurityGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStaleSecurityGroupsPages(input *ec2.DescribeStaleSecurityGroupsInput, fn func(*ec2.DescribeStaleSecurityGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStaleSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeStaleSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStaleSecurityGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStaleSecurityGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput, fn func(*ec2.DescribeStaleSecurityGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStaleSecurityGroupsOutput{}
	fnCacher := func(out *ec2.DescribeStaleSecurityGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStaleSecurityGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStaleSecurityGroupsWithContext(ctx context.Context, input *ec2.DescribeStaleSecurityGroupsInput, opts ...request.Option) (*ec2.DescribeStaleSecurityGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStaleSecurityGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeStaleSecurityGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStoreImageTasks(input *ec2.DescribeStoreImageTasksInput) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeStoreImageTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeStoreImageTasksPages(input *ec2.DescribeStoreImageTasksInput, fn func(*ec2.DescribeStoreImageTasksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStoreImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStoreImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeStoreImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStoreImageTasksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStoreImageTasksPagesWithContext(ctx context.Context, input *ec2.DescribeStoreImageTasksInput, fn func(*ec2.DescribeStoreImageTasksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeStoreImageTasksOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeStoreImageTasksOutput{}
	fnCacher := func(out *ec2.DescribeStoreImageTasksOutput, more bool) bool {
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
	if err := c.EC2API.DescribeStoreImageTasksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeStoreImageTasksWithContext(ctx context.Context, input *ec2.DescribeStoreImageTasksInput, opts ...request.Option) (*ec2.DescribeStoreImageTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeStoreImageTasksOutput), nil
	}
	out, err := c.EC2API.DescribeStoreImageTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.EC2API.DescribeSubnets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeSubnetsPages(input *ec2.DescribeSubnetsInput, fn func(*ec2.DescribeSubnetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSubnetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSubnetsOutput{}
	fnCacher := func(out *ec2.DescribeSubnetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSubnetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSubnetsPagesWithContext(ctx context.Context, input *ec2.DescribeSubnetsInput, fn func(*ec2.DescribeSubnetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeSubnetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeSubnetsOutput{}
	fnCacher := func(out *ec2.DescribeSubnetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeSubnetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeSubnetsWithContext(ctx context.Context, input *ec2.DescribeSubnetsInput, opts ...request.Option) (*ec2.DescribeSubnetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeSubnetsOutput), nil
	}
	out, err := c.EC2API.DescribeSubnetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTags(input *ec2.DescribeTagsInput) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.EC2API.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTagsPages(input *ec2.DescribeTagsInput, fn func(*ec2.DescribeTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTagsOutput{}
	fnCacher := func(out *ec2.DescribeTagsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTagsPagesWithContext(ctx context.Context, input *ec2.DescribeTagsInput, fn func(*ec2.DescribeTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTagsOutput{}
	fnCacher := func(out *ec2.DescribeTagsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTagsWithContext(ctx context.Context, input *ec2.DescribeTagsInput, opts ...request.Option) (*ec2.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTagsOutput), nil
	}
	out, err := c.EC2API.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorFilters(input *ec2.DescribeTrafficMirrorFiltersInput) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorFiltersPages(input *ec2.DescribeTrafficMirrorFiltersInput, fn func(*ec2.DescribeTrafficMirrorFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorFiltersOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorFiltersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorFiltersPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput, fn func(*ec2.DescribeTrafficMirrorFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorFiltersOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorFiltersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorFiltersWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorFiltersInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorFiltersOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorSessions(input *ec2.DescribeTrafficMirrorSessionsInput) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorSessionsPages(input *ec2.DescribeTrafficMirrorSessionsInput, fn func(*ec2.DescribeTrafficMirrorSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorSessionsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorSessionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorSessionsPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput, fn func(*ec2.DescribeTrafficMirrorSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorSessionsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorSessionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorSessionsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorSessionsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorSessionsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorTargets(input *ec2.DescribeTrafficMirrorTargetsInput) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrafficMirrorTargetsPages(input *ec2.DescribeTrafficMirrorTargetsInput, fn func(*ec2.DescribeTrafficMirrorTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorTargetsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorTargetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorTargetsPagesWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput, fn func(*ec2.DescribeTrafficMirrorTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrafficMirrorTargetsOutput{}
	fnCacher := func(out *ec2.DescribeTrafficMirrorTargetsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrafficMirrorTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrafficMirrorTargetsWithContext(ctx context.Context, input *ec2.DescribeTrafficMirrorTargetsInput, opts ...request.Option) (*ec2.DescribeTrafficMirrorTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrafficMirrorTargetsOutput), nil
	}
	out, err := c.EC2API.DescribeTrafficMirrorTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayAttachments(input *ec2.DescribeTransitGatewayAttachmentsInput) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayAttachmentsPages(input *ec2.DescribeTransitGatewayAttachmentsInput, fn func(*ec2.DescribeTransitGatewayAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput, fn func(*ec2.DescribeTransitGatewayAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectPeers(input *ec2.DescribeTransitGatewayConnectPeersInput) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectPeers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectPeersPages(input *ec2.DescribeTransitGatewayConnectPeersInput, fn func(*ec2.DescribeTransitGatewayConnectPeersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectPeersOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectPeersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectPeersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectPeersPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectPeersInput, fn func(*ec2.DescribeTransitGatewayConnectPeersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectPeersOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectPeersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectPeersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectPeersWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectPeersInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectPeersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectPeersOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectPeersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnects(input *ec2.DescribeTransitGatewayConnectsInput) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayConnectsPages(input *ec2.DescribeTransitGatewayConnectsInput, fn func(*ec2.DescribeTransitGatewayConnectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectsInput, fn func(*ec2.DescribeTransitGatewayConnectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayConnectsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayConnectsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayConnectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayConnectsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayConnectsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayConnectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayConnectsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayConnectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayMulticastDomains(input *ec2.DescribeTransitGatewayMulticastDomainsInput) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayMulticastDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsPages(input *ec2.DescribeTransitGatewayMulticastDomainsInput, fn func(*ec2.DescribeTransitGatewayMulticastDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayMulticastDomainsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayMulticastDomainsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayMulticastDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput, fn func(*ec2.DescribeTransitGatewayMulticastDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayMulticastDomainsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayMulticastDomainsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayMulticastDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayMulticastDomainsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayMulticastDomainsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayMulticastDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayMulticastDomainsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayMulticastDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPeeringAttachments(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPeeringAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsPages(input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, fn func(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPeeringAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPeeringAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, fn func(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPeeringAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPeeringAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPeeringAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPeeringAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPeeringAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPeeringAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPeeringAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPolicyTables(input *ec2.DescribeTransitGatewayPolicyTablesInput) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPolicyTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayPolicyTablesPages(input *ec2.DescribeTransitGatewayPolicyTablesInput, fn func(*ec2.DescribeTransitGatewayPolicyTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPolicyTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPolicyTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPolicyTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPolicyTablesPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPolicyTablesInput, fn func(*ec2.DescribeTransitGatewayPolicyTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayPolicyTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayPolicyTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayPolicyTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayPolicyTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayPolicyTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayPolicyTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayPolicyTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayPolicyTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncements(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsPages(input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, fn func(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, fn func(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTableAnnouncementsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTableAnnouncementsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTableAnnouncementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTables(input *ec2.DescribeTransitGatewayRouteTablesInput) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayRouteTablesPages(input *ec2.DescribeTransitGatewayRouteTablesInput, fn func(*ec2.DescribeTransitGatewayRouteTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTablesPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput, fn func(*ec2.DescribeTransitGatewayRouteTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayRouteTablesOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayRouteTablesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayRouteTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayRouteTablesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayRouteTablesInput, opts ...request.Option) (*ec2.DescribeTransitGatewayRouteTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayRouteTablesOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayRouteTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayVpcAttachments(input *ec2.DescribeTransitGatewayVpcAttachmentsInput) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayVpcAttachments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsPages(input *ec2.DescribeTransitGatewayVpcAttachmentsInput, fn func(*ec2.DescribeTransitGatewayVpcAttachmentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayVpcAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayVpcAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayVpcAttachmentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput, fn func(*ec2.DescribeTransitGatewayVpcAttachmentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewayVpcAttachmentsOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewayVpcAttachmentsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewayVpcAttachmentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewayVpcAttachmentsWithContext(ctx context.Context, input *ec2.DescribeTransitGatewayVpcAttachmentsInput, opts ...request.Option) (*ec2.DescribeTransitGatewayVpcAttachmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewayVpcAttachmentsOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewayVpcAttachmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGateways(input *ec2.DescribeTransitGatewaysInput) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTransitGatewaysPages(input *ec2.DescribeTransitGatewaysInput, fn func(*ec2.DescribeTransitGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewaysPagesWithContext(ctx context.Context, input *ec2.DescribeTransitGatewaysInput, fn func(*ec2.DescribeTransitGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTransitGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTransitGatewaysOutput{}
	fnCacher := func(out *ec2.DescribeTransitGatewaysOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTransitGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTransitGatewaysWithContext(ctx context.Context, input *ec2.DescribeTransitGatewaysInput, opts ...request.Option) (*ec2.DescribeTransitGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTransitGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeTransitGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrunkInterfaceAssociations(input *ec2.DescribeTrunkInterfaceAssociationsInput) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeTrunkInterfaceAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeTrunkInterfaceAssociationsPages(input *ec2.DescribeTrunkInterfaceAssociationsInput, fn func(*ec2.DescribeTrunkInterfaceAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrunkInterfaceAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeTrunkInterfaceAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrunkInterfaceAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrunkInterfaceAssociationsPagesWithContext(ctx context.Context, input *ec2.DescribeTrunkInterfaceAssociationsInput, fn func(*ec2.DescribeTrunkInterfaceAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeTrunkInterfaceAssociationsOutput{}
	fnCacher := func(out *ec2.DescribeTrunkInterfaceAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeTrunkInterfaceAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeTrunkInterfaceAssociationsWithContext(ctx context.Context, input *ec2.DescribeTrunkInterfaceAssociationsInput, opts ...request.Option) (*ec2.DescribeTrunkInterfaceAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeTrunkInterfaceAssociationsOutput), nil
	}
	out, err := c.EC2API.DescribeTrunkInterfaceAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessEndpoints(input *ec2.DescribeVerifiedAccessEndpointsInput) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessEndpointsPages(input *ec2.DescribeVerifiedAccessEndpointsInput, fn func(*ec2.DescribeVerifiedAccessEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessEndpointsInput, fn func(*ec2.DescribeVerifiedAccessEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessEndpointsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessEndpointsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessGroups(input *ec2.DescribeVerifiedAccessGroupsInput) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessGroupsPages(input *ec2.DescribeVerifiedAccessGroupsInput, fn func(*ec2.DescribeVerifiedAccessGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessGroupsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessGroupsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessGroupsInput, fn func(*ec2.DescribeVerifiedAccessGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessGroupsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessGroupsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessGroupsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessGroupsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessGroupsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurations(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, fn func(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, fn func(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstanceLoggingConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstanceLoggingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstances(input *ec2.DescribeVerifiedAccessInstancesInput) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessInstancesPages(input *ec2.DescribeVerifiedAccessInstancesInput, fn func(*ec2.DescribeVerifiedAccessInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstancesOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstancesPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstancesInput, fn func(*ec2.DescribeVerifiedAccessInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessInstancesOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessInstancesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessInstancesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessInstancesInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessInstancesOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessTrustProviders(input *ec2.DescribeVerifiedAccessTrustProvidersInput) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessTrustProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersPages(input *ec2.DescribeVerifiedAccessTrustProvidersInput, fn func(*ec2.DescribeVerifiedAccessTrustProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessTrustProvidersOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessTrustProvidersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessTrustProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersPagesWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessTrustProvidersInput, fn func(*ec2.DescribeVerifiedAccessTrustProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVerifiedAccessTrustProvidersOutput{}
	fnCacher := func(out *ec2.DescribeVerifiedAccessTrustProvidersOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVerifiedAccessTrustProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVerifiedAccessTrustProvidersWithContext(ctx context.Context, input *ec2.DescribeVerifiedAccessTrustProvidersInput, opts ...request.Option) (*ec2.DescribeVerifiedAccessTrustProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVerifiedAccessTrustProvidersOutput), nil
	}
	out, err := c.EC2API.DescribeVerifiedAccessTrustProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeAttribute(input *ec2.DescribeVolumeAttributeInput) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeAttributeWithContext(ctx context.Context, input *ec2.DescribeVolumeAttributeInput, opts ...request.Option) (*ec2.DescribeVolumeAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeStatus(input *ec2.DescribeVolumeStatusInput) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumeStatusPages(input *ec2.DescribeVolumeStatusInput, fn func(*ec2.DescribeVolumeStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumeStatusOutput{}
	fnCacher := func(out *ec2.DescribeVolumeStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumeStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumeStatusPagesWithContext(ctx context.Context, input *ec2.DescribeVolumeStatusInput, fn func(*ec2.DescribeVolumeStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumeStatusOutput{}
	fnCacher := func(out *ec2.DescribeVolumeStatusOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumeStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumeStatusWithContext(ctx context.Context, input *ec2.DescribeVolumeStatusInput, opts ...request.Option) (*ec2.DescribeVolumeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumeStatusOutput), nil
	}
	out, err := c.EC2API.DescribeVolumeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.EC2API.DescribeVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesModifications(input *ec2.DescribeVolumesModificationsInput) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesModifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesModificationsPages(input *ec2.DescribeVolumesModificationsInput, fn func(*ec2.DescribeVolumesModificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeVolumesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesModificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesModificationsPagesWithContext(ctx context.Context, input *ec2.DescribeVolumesModificationsInput, fn func(*ec2.DescribeVolumesModificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesModificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesModificationsOutput{}
	fnCacher := func(out *ec2.DescribeVolumesModificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesModificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesModificationsWithContext(ctx context.Context, input *ec2.DescribeVolumesModificationsInput, opts ...request.Option) (*ec2.DescribeVolumesModificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesModificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesModificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVolumesPages(input *ec2.DescribeVolumesInput, fn func(*ec2.DescribeVolumesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesOutput{}
	fnCacher := func(out *ec2.DescribeVolumesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesPagesWithContext(ctx context.Context, input *ec2.DescribeVolumesInput, fn func(*ec2.DescribeVolumesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVolumesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVolumesOutput{}
	fnCacher := func(out *ec2.DescribeVolumesOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVolumesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVolumesWithContext(ctx context.Context, input *ec2.DescribeVolumesInput, opts ...request.Option) (*ec2.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVolumesOutput), nil
	}
	out, err := c.EC2API.DescribeVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcAttribute(input *ec2.DescribeVpcAttributeInput) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVpcAttribute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcAttributeWithContext(ctx context.Context, input *ec2.DescribeVpcAttributeInput, opts ...request.Option) (*ec2.DescribeVpcAttributeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcAttributeOutput), nil
	}
	out, err := c.EC2API.DescribeVpcAttributeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLink(input *ec2.DescribeVpcClassicLinkInput) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLink(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkDnsSupport(input *ec2.DescribeVpcClassicLinkDnsSupportInput) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkDnsSupport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportPages(input *ec2.DescribeVpcClassicLinkDnsSupportInput, fn func(*ec2.DescribeVpcClassicLinkDnsSupportOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcClassicLinkDnsSupportOutput{}
	fnCacher := func(out *ec2.DescribeVpcClassicLinkDnsSupportOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcClassicLinkDnsSupportPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportPagesWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput, fn func(*ec2.DescribeVpcClassicLinkDnsSupportOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcClassicLinkDnsSupportOutput{}
	fnCacher := func(out *ec2.DescribeVpcClassicLinkDnsSupportOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcClassicLinkDnsSupportPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcClassicLinkDnsSupportWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkDnsSupportInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkDnsSupportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkDnsSupportOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkDnsSupportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcClassicLinkWithContext(ctx context.Context, input *ec2.DescribeVpcClassicLinkInput, opts ...request.Option) (*ec2.DescribeVpcClassicLinkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcClassicLinkOutput), nil
	}
	out, err := c.EC2API.DescribeVpcClassicLinkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionNotifications(input *ec2.DescribeVpcEndpointConnectionNotificationsInput) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsPages(input *ec2.DescribeVpcEndpointConnectionNotificationsInput, fn func(*ec2.DescribeVpcEndpointConnectionNotificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionNotificationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionNotificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionNotificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput, fn func(*ec2.DescribeVpcEndpointConnectionNotificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionNotificationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionNotificationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionNotificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionNotificationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionNotificationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionNotificationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnections(input *ec2.DescribeVpcEndpointConnectionsInput) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointConnectionsPages(input *ec2.DescribeVpcEndpointConnectionsInput, fn func(*ec2.DescribeVpcEndpointConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput, fn func(*ec2.DescribeVpcEndpointConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServiceConfigurations(input *ec2.DescribeVpcEndpointServiceConfigurationsInput) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServiceConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsPages(input *ec2.DescribeVpcEndpointServiceConfigurationsInput, fn func(*ec2.DescribeVpcEndpointServiceConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServiceConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServiceConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServiceConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, fn func(*ec2.DescribeVpcEndpointServiceConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServiceConfigurationsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServiceConfigurationsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServiceConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServiceConfigurationsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServiceConfigurationsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServiceConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServiceConfigurationsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServiceConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicePermissions(input *ec2.DescribeVpcEndpointServicePermissionsInput) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicePermissionsPages(input *ec2.DescribeVpcEndpointServicePermissionsInput, fn func(*ec2.DescribeVpcEndpointServicePermissionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServicePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServicePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServicePermissionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServicePermissionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, fn func(*ec2.DescribeVpcEndpointServicePermissionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointServicePermissionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointServicePermissionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointServicePermissionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointServicePermissionsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicePermissionsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicePermissionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServices(input *ec2.DescribeVpcEndpointServicesInput) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointServicesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointServicesInput, opts ...request.Option) (*ec2.DescribeVpcEndpointServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointServicesOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpoints(input *ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcEndpointsPages(input *ec2.DescribeVpcEndpointsInput, fn func(*ec2.DescribeVpcEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointsInput, fn func(*ec2.DescribeVpcEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcEndpointsOutput{}
	fnCacher := func(out *ec2.DescribeVpcEndpointsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcEndpointsWithContext(ctx context.Context, input *ec2.DescribeVpcEndpointsInput, opts ...request.Option) (*ec2.DescribeVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcEndpointsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcPeeringConnections(input *ec2.DescribeVpcPeeringConnectionsInput) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcPeeringConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcPeeringConnectionsPages(input *ec2.DescribeVpcPeeringConnectionsInput, fn func(*ec2.DescribeVpcPeeringConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcPeeringConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcPeeringConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcPeeringConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcPeeringConnectionsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput, fn func(*ec2.DescribeVpcPeeringConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcPeeringConnectionsOutput{}
	fnCacher := func(out *ec2.DescribeVpcPeeringConnectionsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcPeeringConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcPeeringConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpcPeeringConnectionsInput, opts ...request.Option) (*ec2.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcPeeringConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcs(input *ec2.DescribeVpcsInput) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpcsPages(input *ec2.DescribeVpcsInput, fn func(*ec2.DescribeVpcsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcsOutput{}
	fnCacher := func(out *ec2.DescribeVpcsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcsPagesWithContext(ctx context.Context, input *ec2.DescribeVpcsInput, fn func(*ec2.DescribeVpcsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.DescribeVpcsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.DescribeVpcsOutput{}
	fnCacher := func(out *ec2.DescribeVpcsOutput, more bool) bool {
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
	if err := c.EC2API.DescribeVpcsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) DescribeVpcsWithContext(ctx context.Context, input *ec2.DescribeVpcsInput, opts ...request.Option) (*ec2.DescribeVpcsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpcsOutput), nil
	}
	out, err := c.EC2API.DescribeVpcsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnConnections(input *ec2.DescribeVpnConnectionsInput) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpnConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnConnectionsWithContext(ctx context.Context, input *ec2.DescribeVpnConnectionsInput, opts ...request.Option) (*ec2.DescribeVpnConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnConnectionsOutput), nil
	}
	out, err := c.EC2API.DescribeVpnConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnGateways(input *ec2.DescribeVpnGatewaysInput) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeVpnGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) DescribeVpnGatewaysWithContext(ctx context.Context, input *ec2.DescribeVpnGatewaysInput, opts ...request.Option) (*ec2.DescribeVpnGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.DescribeVpnGatewaysOutput), nil
	}
	out, err := c.EC2API.DescribeVpnGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedEnclaveCertificateIamRoles(input *ec2.GetAssociatedEnclaveCertificateIamRolesInput) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.EC2API.GetAssociatedEnclaveCertificateIamRoles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedEnclaveCertificateIamRolesWithContext(ctx context.Context, input *ec2.GetAssociatedEnclaveCertificateIamRolesInput, opts ...request.Option) (*ec2.GetAssociatedEnclaveCertificateIamRolesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedEnclaveCertificateIamRolesOutput), nil
	}
	out, err := c.EC2API.GetAssociatedEnclaveCertificateIamRolesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedIpv6PoolCidrs(input *ec2.GetAssociatedIpv6PoolCidrsInput) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetAssociatedIpv6PoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAssociatedIpv6PoolCidrsPages(input *ec2.GetAssociatedIpv6PoolCidrsInput, fn func(*ec2.GetAssociatedIpv6PoolCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAssociatedIpv6PoolCidrsOutput{}
	fnCacher := func(out *ec2.GetAssociatedIpv6PoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetAssociatedIpv6PoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAssociatedIpv6PoolCidrsPagesWithContext(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput, fn func(*ec2.GetAssociatedIpv6PoolCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAssociatedIpv6PoolCidrsOutput{}
	fnCacher := func(out *ec2.GetAssociatedIpv6PoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetAssociatedIpv6PoolCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAssociatedIpv6PoolCidrsWithContext(ctx context.Context, input *ec2.GetAssociatedIpv6PoolCidrsInput, opts ...request.Option) (*ec2.GetAssociatedIpv6PoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAssociatedIpv6PoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetAssociatedIpv6PoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAwsNetworkPerformanceData(input *ec2.GetAwsNetworkPerformanceDataInput) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.EC2API.GetAwsNetworkPerformanceData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetAwsNetworkPerformanceDataPages(input *ec2.GetAwsNetworkPerformanceDataInput, fn func(*ec2.GetAwsNetworkPerformanceDataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAwsNetworkPerformanceDataOutput{}
	fnCacher := func(out *ec2.GetAwsNetworkPerformanceDataOutput, more bool) bool {
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
	if err := c.EC2API.GetAwsNetworkPerformanceDataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAwsNetworkPerformanceDataPagesWithContext(ctx context.Context, input *ec2.GetAwsNetworkPerformanceDataInput, fn func(*ec2.GetAwsNetworkPerformanceDataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetAwsNetworkPerformanceDataOutput{}
	fnCacher := func(out *ec2.GetAwsNetworkPerformanceDataOutput, more bool) bool {
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
	if err := c.EC2API.GetAwsNetworkPerformanceDataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetAwsNetworkPerformanceDataWithContext(ctx context.Context, input *ec2.GetAwsNetworkPerformanceDataInput, opts ...request.Option) (*ec2.GetAwsNetworkPerformanceDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetAwsNetworkPerformanceDataOutput), nil
	}
	out, err := c.EC2API.GetAwsNetworkPerformanceDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCapacityReservationUsage(input *ec2.GetCapacityReservationUsageInput) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.EC2API.GetCapacityReservationUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCapacityReservationUsageWithContext(ctx context.Context, input *ec2.GetCapacityReservationUsageInput, opts ...request.Option) (*ec2.GetCapacityReservationUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCapacityReservationUsageOutput), nil
	}
	out, err := c.EC2API.GetCapacityReservationUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCoipPoolUsage(input *ec2.GetCoipPoolUsageInput) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.EC2API.GetCoipPoolUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetCoipPoolUsageWithContext(ctx context.Context, input *ec2.GetCoipPoolUsageInput, opts ...request.Option) (*ec2.GetCoipPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetCoipPoolUsageOutput), nil
	}
	out, err := c.EC2API.GetCoipPoolUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleOutput(input *ec2.GetConsoleOutputInput) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.EC2API.GetConsoleOutput(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleOutputWithContext(ctx context.Context, input *ec2.GetConsoleOutputInput, opts ...request.Option) (*ec2.GetConsoleOutputOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleOutputOutput), nil
	}
	out, err := c.EC2API.GetConsoleOutputWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleScreenshot(input *ec2.GetConsoleScreenshotInput) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.EC2API.GetConsoleScreenshot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetConsoleScreenshotWithContext(ctx context.Context, input *ec2.GetConsoleScreenshotInput, opts ...request.Option) (*ec2.GetConsoleScreenshotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetConsoleScreenshotOutput), nil
	}
	out, err := c.EC2API.GetConsoleScreenshotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetDefaultCreditSpecification(input *ec2.GetDefaultCreditSpecificationInput) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.GetDefaultCreditSpecification(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetDefaultCreditSpecificationWithContext(ctx context.Context, input *ec2.GetDefaultCreditSpecificationInput, opts ...request.Option) (*ec2.GetDefaultCreditSpecificationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetDefaultCreditSpecificationOutput), nil
	}
	out, err := c.EC2API.GetDefaultCreditSpecificationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsDefaultKmsKeyId(input *ec2.GetEbsDefaultKmsKeyIdInput) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.GetEbsDefaultKmsKeyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsDefaultKmsKeyIdWithContext(ctx context.Context, input *ec2.GetEbsDefaultKmsKeyIdInput, opts ...request.Option) (*ec2.GetEbsDefaultKmsKeyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsDefaultKmsKeyIdOutput), nil
	}
	out, err := c.EC2API.GetEbsDefaultKmsKeyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsEncryptionByDefault(input *ec2.GetEbsEncryptionByDefaultInput) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.GetEbsEncryptionByDefault(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetEbsEncryptionByDefaultWithContext(ctx context.Context, input *ec2.GetEbsEncryptionByDefaultInput, opts ...request.Option) (*ec2.GetEbsEncryptionByDefaultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetEbsEncryptionByDefaultOutput), nil
	}
	out, err := c.EC2API.GetEbsEncryptionByDefaultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetFlowLogsIntegrationTemplate(input *ec2.GetFlowLogsIntegrationTemplateInput) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.EC2API.GetFlowLogsIntegrationTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetFlowLogsIntegrationTemplateWithContext(ctx context.Context, input *ec2.GetFlowLogsIntegrationTemplateInput, opts ...request.Option) (*ec2.GetFlowLogsIntegrationTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetFlowLogsIntegrationTemplateOutput), nil
	}
	out, err := c.EC2API.GetFlowLogsIntegrationTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetGroupsForCapacityReservation(input *ec2.GetGroupsForCapacityReservationInput) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.EC2API.GetGroupsForCapacityReservation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetGroupsForCapacityReservationPages(input *ec2.GetGroupsForCapacityReservationInput, fn func(*ec2.GetGroupsForCapacityReservationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetGroupsForCapacityReservationOutput{}
	fnCacher := func(out *ec2.GetGroupsForCapacityReservationOutput, more bool) bool {
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
	if err := c.EC2API.GetGroupsForCapacityReservationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetGroupsForCapacityReservationPagesWithContext(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput, fn func(*ec2.GetGroupsForCapacityReservationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetGroupsForCapacityReservationOutput{}
	fnCacher := func(out *ec2.GetGroupsForCapacityReservationOutput, more bool) bool {
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
	if err := c.EC2API.GetGroupsForCapacityReservationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetGroupsForCapacityReservationWithContext(ctx context.Context, input *ec2.GetGroupsForCapacityReservationInput, opts ...request.Option) (*ec2.GetGroupsForCapacityReservationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetGroupsForCapacityReservationOutput), nil
	}
	out, err := c.EC2API.GetGroupsForCapacityReservationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetHostReservationPurchasePreview(input *ec2.GetHostReservationPurchasePreviewInput) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.EC2API.GetHostReservationPurchasePreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetHostReservationPurchasePreviewWithContext(ctx context.Context, input *ec2.GetHostReservationPurchasePreviewInput, opts ...request.Option) (*ec2.GetHostReservationPurchasePreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetHostReservationPurchasePreviewOutput), nil
	}
	out, err := c.EC2API.GetHostReservationPurchasePreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceTypesFromInstanceRequirements(input *ec2.GetInstanceTypesFromInstanceRequirementsInput) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.EC2API.GetInstanceTypesFromInstanceRequirements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsPages(input *ec2.GetInstanceTypesFromInstanceRequirementsInput, fn func(*ec2.GetInstanceTypesFromInstanceRequirementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetInstanceTypesFromInstanceRequirementsOutput{}
	fnCacher := func(out *ec2.GetInstanceTypesFromInstanceRequirementsOutput, more bool) bool {
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
	if err := c.EC2API.GetInstanceTypesFromInstanceRequirementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsPagesWithContext(ctx context.Context, input *ec2.GetInstanceTypesFromInstanceRequirementsInput, fn func(*ec2.GetInstanceTypesFromInstanceRequirementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetInstanceTypesFromInstanceRequirementsOutput{}
	fnCacher := func(out *ec2.GetInstanceTypesFromInstanceRequirementsOutput, more bool) bool {
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
	if err := c.EC2API.GetInstanceTypesFromInstanceRequirementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetInstanceTypesFromInstanceRequirementsWithContext(ctx context.Context, input *ec2.GetInstanceTypesFromInstanceRequirementsInput, opts ...request.Option) (*ec2.GetInstanceTypesFromInstanceRequirementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceTypesFromInstanceRequirementsOutput), nil
	}
	out, err := c.EC2API.GetInstanceTypesFromInstanceRequirementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceUefiData(input *ec2.GetInstanceUefiDataInput) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.EC2API.GetInstanceUefiData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetInstanceUefiDataWithContext(ctx context.Context, input *ec2.GetInstanceUefiDataInput, opts ...request.Option) (*ec2.GetInstanceUefiDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetInstanceUefiDataOutput), nil
	}
	out, err := c.EC2API.GetInstanceUefiDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamAddressHistory(input *ec2.GetIpamAddressHistoryInput) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.EC2API.GetIpamAddressHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamAddressHistoryPages(input *ec2.GetIpamAddressHistoryInput, fn func(*ec2.GetIpamAddressHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamAddressHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamAddressHistoryOutput{}
	fnCacher := func(out *ec2.GetIpamAddressHistoryOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamAddressHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamAddressHistoryPagesWithContext(ctx context.Context, input *ec2.GetIpamAddressHistoryInput, fn func(*ec2.GetIpamAddressHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamAddressHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamAddressHistoryOutput{}
	fnCacher := func(out *ec2.GetIpamAddressHistoryOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamAddressHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamAddressHistoryWithContext(ctx context.Context, input *ec2.GetIpamAddressHistoryInput, opts ...request.Option) (*ec2.GetIpamAddressHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamAddressHistoryOutput), nil
	}
	out, err := c.EC2API.GetIpamAddressHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolAllocations(input *ec2.GetIpamPoolAllocationsInput) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolAllocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolAllocationsPages(input *ec2.GetIpamPoolAllocationsInput, fn func(*ec2.GetIpamPoolAllocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolAllocationsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolAllocationsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolAllocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolAllocationsPagesWithContext(ctx context.Context, input *ec2.GetIpamPoolAllocationsInput, fn func(*ec2.GetIpamPoolAllocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolAllocationsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolAllocationsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolAllocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolAllocationsWithContext(ctx context.Context, input *ec2.GetIpamPoolAllocationsInput, opts ...request.Option) (*ec2.GetIpamPoolAllocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolAllocationsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolAllocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolCidrs(input *ec2.GetIpamPoolCidrsInput) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamPoolCidrsPages(input *ec2.GetIpamPoolCidrsInput, fn func(*ec2.GetIpamPoolCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolCidrsPagesWithContext(ctx context.Context, input *ec2.GetIpamPoolCidrsInput, fn func(*ec2.GetIpamPoolCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamPoolCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamPoolCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamPoolCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamPoolCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamPoolCidrsWithContext(ctx context.Context, input *ec2.GetIpamPoolCidrsInput, opts ...request.Option) (*ec2.GetIpamPoolCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamPoolCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamPoolCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamResourceCidrs(input *ec2.GetIpamResourceCidrsInput) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamResourceCidrs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetIpamResourceCidrsPages(input *ec2.GetIpamResourceCidrsInput, fn func(*ec2.GetIpamResourceCidrsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamResourceCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamResourceCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamResourceCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamResourceCidrsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamResourceCidrsPagesWithContext(ctx context.Context, input *ec2.GetIpamResourceCidrsInput, fn func(*ec2.GetIpamResourceCidrsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetIpamResourceCidrsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetIpamResourceCidrsOutput{}
	fnCacher := func(out *ec2.GetIpamResourceCidrsOutput, more bool) bool {
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
	if err := c.EC2API.GetIpamResourceCidrsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetIpamResourceCidrsWithContext(ctx context.Context, input *ec2.GetIpamResourceCidrsInput, opts ...request.Option) (*ec2.GetIpamResourceCidrsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetIpamResourceCidrsOutput), nil
	}
	out, err := c.EC2API.GetIpamResourceCidrsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetLaunchTemplateData(input *ec2.GetLaunchTemplateDataInput) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.EC2API.GetLaunchTemplateData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetLaunchTemplateDataWithContext(ctx context.Context, input *ec2.GetLaunchTemplateDataInput, opts ...request.Option) (*ec2.GetLaunchTemplateDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetLaunchTemplateDataOutput), nil
	}
	out, err := c.EC2API.GetLaunchTemplateDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListAssociations(input *ec2.GetManagedPrefixListAssociationsInput) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListAssociationsPages(input *ec2.GetManagedPrefixListAssociationsInput, fn func(*ec2.GetManagedPrefixListAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListAssociationsOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListAssociationsPagesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput, fn func(*ec2.GetManagedPrefixListAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListAssociationsOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListAssociationsWithContext(ctx context.Context, input *ec2.GetManagedPrefixListAssociationsInput, opts ...request.Option) (*ec2.GetManagedPrefixListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListAssociationsOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListEntries(input *ec2.GetManagedPrefixListEntriesInput) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetManagedPrefixListEntriesPages(input *ec2.GetManagedPrefixListEntriesInput, fn func(*ec2.GetManagedPrefixListEntriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListEntriesOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListEntriesOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListEntriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListEntriesPagesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput, fn func(*ec2.GetManagedPrefixListEntriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetManagedPrefixListEntriesOutput{}
	fnCacher := func(out *ec2.GetManagedPrefixListEntriesOutput, more bool) bool {
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
	if err := c.EC2API.GetManagedPrefixListEntriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetManagedPrefixListEntriesWithContext(ctx context.Context, input *ec2.GetManagedPrefixListEntriesInput, opts ...request.Option) (*ec2.GetManagedPrefixListEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetManagedPrefixListEntriesOutput), nil
	}
	out, err := c.EC2API.GetManagedPrefixListEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeAnalysisFindings(input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeAnalysisFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeAnalysisFindingsInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeAnalysisFindingsOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeAnalysisFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeContent(input *ec2.GetNetworkInsightsAccessScopeContentInput) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetNetworkInsightsAccessScopeContentWithContext(ctx context.Context, input *ec2.GetNetworkInsightsAccessScopeContentInput, opts ...request.Option) (*ec2.GetNetworkInsightsAccessScopeContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetNetworkInsightsAccessScopeContentOutput), nil
	}
	out, err := c.EC2API.GetNetworkInsightsAccessScopeContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetPasswordData(input *ec2.GetPasswordDataInput) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.EC2API.GetPasswordData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetPasswordDataWithContext(ctx context.Context, input *ec2.GetPasswordDataInput, opts ...request.Option) (*ec2.GetPasswordDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetPasswordDataOutput), nil
	}
	out, err := c.EC2API.GetPasswordDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetReservedInstancesExchangeQuote(input *ec2.GetReservedInstancesExchangeQuoteInput) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.GetReservedInstancesExchangeQuote(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetReservedInstancesExchangeQuoteWithContext(ctx context.Context, input *ec2.GetReservedInstancesExchangeQuoteInput, opts ...request.Option) (*ec2.GetReservedInstancesExchangeQuoteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetReservedInstancesExchangeQuoteOutput), nil
	}
	out, err := c.EC2API.GetReservedInstancesExchangeQuoteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSerialConsoleAccessStatus(input *ec2.GetSerialConsoleAccessStatusInput) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.EC2API.GetSerialConsoleAccessStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSerialConsoleAccessStatusWithContext(ctx context.Context, input *ec2.GetSerialConsoleAccessStatusInput, opts ...request.Option) (*ec2.GetSerialConsoleAccessStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSerialConsoleAccessStatusOutput), nil
	}
	out, err := c.EC2API.GetSerialConsoleAccessStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSpotPlacementScores(input *ec2.GetSpotPlacementScoresInput) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.EC2API.GetSpotPlacementScores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSpotPlacementScoresPages(input *ec2.GetSpotPlacementScoresInput, fn func(*ec2.GetSpotPlacementScoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetSpotPlacementScoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetSpotPlacementScoresOutput{}
	fnCacher := func(out *ec2.GetSpotPlacementScoresOutput, more bool) bool {
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
	if err := c.EC2API.GetSpotPlacementScoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetSpotPlacementScoresPagesWithContext(ctx context.Context, input *ec2.GetSpotPlacementScoresInput, fn func(*ec2.GetSpotPlacementScoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetSpotPlacementScoresOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetSpotPlacementScoresOutput{}
	fnCacher := func(out *ec2.GetSpotPlacementScoresOutput, more bool) bool {
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
	if err := c.EC2API.GetSpotPlacementScoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetSpotPlacementScoresWithContext(ctx context.Context, input *ec2.GetSpotPlacementScoresInput, opts ...request.Option) (*ec2.GetSpotPlacementScoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSpotPlacementScoresOutput), nil
	}
	out, err := c.EC2API.GetSpotPlacementScoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSubnetCidrReservations(input *ec2.GetSubnetCidrReservationsInput) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.EC2API.GetSubnetCidrReservations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetSubnetCidrReservationsWithContext(ctx context.Context, input *ec2.GetSubnetCidrReservationsInput, opts ...request.Option) (*ec2.GetSubnetCidrReservationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetSubnetCidrReservationsOutput), nil
	}
	out, err := c.EC2API.GetSubnetCidrReservationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayAttachmentPropagations(input *ec2.GetTransitGatewayAttachmentPropagationsInput) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayAttachmentPropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsPages(input *ec2.GetTransitGatewayAttachmentPropagationsInput, fn func(*ec2.GetTransitGatewayAttachmentPropagationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayAttachmentPropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayAttachmentPropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayAttachmentPropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput, fn func(*ec2.GetTransitGatewayAttachmentPropagationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayAttachmentPropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayAttachmentPropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayAttachmentPropagationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayAttachmentPropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayAttachmentPropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayAttachmentPropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayAttachmentPropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayAttachmentPropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociations(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayMulticastDomainAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsPages(input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, fn func(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayMulticastDomainAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayMulticastDomainAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, fn func(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayMulticastDomainAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayMulticastDomainAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayMulticastDomainAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayMulticastDomainAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayMulticastDomainAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayMulticastDomainAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayMulticastDomainAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableAssociations(input *ec2.GetTransitGatewayPolicyTableAssociationsInput) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsPages(input *ec2.GetTransitGatewayPolicyTableAssociationsInput, fn func(*ec2.GetTransitGatewayPolicyTableAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPolicyTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPolicyTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPolicyTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableAssociationsInput, fn func(*ec2.GetTransitGatewayPolicyTableAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPolicyTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPolicyTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPolicyTableAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPolicyTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableEntries(input *ec2.GetTransitGatewayPolicyTableEntriesInput) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPolicyTableEntriesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPolicyTableEntriesInput, opts ...request.Option) (*ec2.GetTransitGatewayPolicyTableEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPolicyTableEntriesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPolicyTableEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPrefixListReferences(input *ec2.GetTransitGatewayPrefixListReferencesInput) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPrefixListReferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayPrefixListReferencesPages(input *ec2.GetTransitGatewayPrefixListReferencesInput, fn func(*ec2.GetTransitGatewayPrefixListReferencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPrefixListReferencesOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPrefixListReferencesOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPrefixListReferencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPrefixListReferencesPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput, fn func(*ec2.GetTransitGatewayPrefixListReferencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayPrefixListReferencesOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayPrefixListReferencesOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayPrefixListReferencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayPrefixListReferencesWithContext(ctx context.Context, input *ec2.GetTransitGatewayPrefixListReferencesInput, opts ...request.Option) (*ec2.GetTransitGatewayPrefixListReferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayPrefixListReferencesOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayPrefixListReferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTableAssociations(input *ec2.GetTransitGatewayRouteTableAssociationsInput) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTableAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsPages(input *ec2.GetTransitGatewayRouteTableAssociationsInput, fn func(*ec2.GetTransitGatewayRouteTableAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTableAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput, fn func(*ec2.GetTransitGatewayRouteTableAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTableAssociationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTableAssociationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTableAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTableAssociationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTableAssociationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTableAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTableAssociationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTableAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTablePropagations(input *ec2.GetTransitGatewayRouteTablePropagationsInput) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTablePropagations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsPages(input *ec2.GetTransitGatewayRouteTablePropagationsInput, fn func(*ec2.GetTransitGatewayRouteTablePropagationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTablePropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTablePropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTablePropagationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsPagesWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput, fn func(*ec2.GetTransitGatewayRouteTablePropagationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetTransitGatewayRouteTablePropagationsOutput{}
	fnCacher := func(out *ec2.GetTransitGatewayRouteTablePropagationsOutput, more bool) bool {
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
	if err := c.EC2API.GetTransitGatewayRouteTablePropagationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetTransitGatewayRouteTablePropagationsWithContext(ctx context.Context, input *ec2.GetTransitGatewayRouteTablePropagationsInput, opts ...request.Option) (*ec2.GetTransitGatewayRouteTablePropagationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetTransitGatewayRouteTablePropagationsOutput), nil
	}
	out, err := c.EC2API.GetTransitGatewayRouteTablePropagationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessEndpointPolicy(input *ec2.GetVerifiedAccessEndpointPolicyInput) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessEndpointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessEndpointPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessEndpointPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessEndpointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessEndpointPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessEndpointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessGroupPolicy(input *ec2.GetVerifiedAccessGroupPolicyInput) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVerifiedAccessGroupPolicyWithContext(ctx context.Context, input *ec2.GetVerifiedAccessGroupPolicyInput, opts ...request.Option) (*ec2.GetVerifiedAccessGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVerifiedAccessGroupPolicyOutput), nil
	}
	out, err := c.EC2API.GetVerifiedAccessGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceSampleConfiguration(input *ec2.GetVpnConnectionDeviceSampleConfigurationInput) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceSampleConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceSampleConfigurationWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceSampleConfigurationInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceSampleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceSampleConfigurationOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceSampleConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceTypes(input *ec2.GetVpnConnectionDeviceTypesInput) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) GetVpnConnectionDeviceTypesPages(input *ec2.GetVpnConnectionDeviceTypesInput, fn func(*ec2.GetVpnConnectionDeviceTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetVpnConnectionDeviceTypesOutput{}
	fnCacher := func(out *ec2.GetVpnConnectionDeviceTypesOutput, more bool) bool {
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
	if err := c.EC2API.GetVpnConnectionDeviceTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetVpnConnectionDeviceTypesPagesWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceTypesInput, fn func(*ec2.GetVpnConnectionDeviceTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.GetVpnConnectionDeviceTypesOutput{}
	fnCacher := func(out *ec2.GetVpnConnectionDeviceTypesOutput, more bool) bool {
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
	if err := c.EC2API.GetVpnConnectionDeviceTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) GetVpnConnectionDeviceTypesWithContext(ctx context.Context, input *ec2.GetVpnConnectionDeviceTypesInput, opts ...request.Option) (*ec2.GetVpnConnectionDeviceTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.GetVpnConnectionDeviceTypesOutput), nil
	}
	out, err := c.EC2API.GetVpnConnectionDeviceTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListImagesInRecycleBin(input *ec2.ListImagesInRecycleBinInput) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListImagesInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListImagesInRecycleBinPages(input *ec2.ListImagesInRecycleBinInput, fn func(*ec2.ListImagesInRecycleBinOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListImagesInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListImagesInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListImagesInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListImagesInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListImagesInRecycleBinPagesWithContext(ctx context.Context, input *ec2.ListImagesInRecycleBinInput, fn func(*ec2.ListImagesInRecycleBinOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListImagesInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListImagesInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListImagesInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListImagesInRecycleBinPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListImagesInRecycleBinWithContext(ctx context.Context, input *ec2.ListImagesInRecycleBinInput, opts ...request.Option) (*ec2.ListImagesInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListImagesInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListImagesInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListSnapshotsInRecycleBin(input *ec2.ListSnapshotsInRecycleBinInput) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListSnapshotsInRecycleBin(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) ListSnapshotsInRecycleBinPages(input *ec2.ListSnapshotsInRecycleBinInput, fn func(*ec2.ListSnapshotsInRecycleBinOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListSnapshotsInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListSnapshotsInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListSnapshotsInRecycleBinPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListSnapshotsInRecycleBinPagesWithContext(ctx context.Context, input *ec2.ListSnapshotsInRecycleBinInput, fn func(*ec2.ListSnapshotsInRecycleBinOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.ListSnapshotsInRecycleBinOutput{}
	fnCacher := func(out *ec2.ListSnapshotsInRecycleBinOutput, more bool) bool {
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
	if err := c.EC2API.ListSnapshotsInRecycleBinPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) ListSnapshotsInRecycleBinWithContext(ctx context.Context, input *ec2.ListSnapshotsInRecycleBinInput, opts ...request.Option) (*ec2.ListSnapshotsInRecycleBinOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.ListSnapshotsInRecycleBinOutput), nil
	}
	out, err := c.EC2API.ListSnapshotsInRecycleBinWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchLocalGatewayRoutes(input *ec2.SearchLocalGatewayRoutesInput) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchLocalGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchLocalGatewayRoutesPages(input *ec2.SearchLocalGatewayRoutesInput, fn func(*ec2.SearchLocalGatewayRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchLocalGatewayRoutesOutput{}
	fnCacher := func(out *ec2.SearchLocalGatewayRoutesOutput, more bool) bool {
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
	if err := c.EC2API.SearchLocalGatewayRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchLocalGatewayRoutesPagesWithContext(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput, fn func(*ec2.SearchLocalGatewayRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchLocalGatewayRoutesOutput{}
	fnCacher := func(out *ec2.SearchLocalGatewayRoutesOutput, more bool) bool {
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
	if err := c.EC2API.SearchLocalGatewayRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchLocalGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchLocalGatewayRoutesInput, opts ...request.Option) (*ec2.SearchLocalGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchLocalGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchLocalGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayMulticastGroups(input *ec2.SearchTransitGatewayMulticastGroupsInput) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayMulticastGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayMulticastGroupsPages(input *ec2.SearchTransitGatewayMulticastGroupsInput, fn func(*ec2.SearchTransitGatewayMulticastGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchTransitGatewayMulticastGroupsOutput{}
	fnCacher := func(out *ec2.SearchTransitGatewayMulticastGroupsOutput, more bool) bool {
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
	if err := c.EC2API.SearchTransitGatewayMulticastGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchTransitGatewayMulticastGroupsPagesWithContext(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput, fn func(*ec2.SearchTransitGatewayMulticastGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &ec2.SearchTransitGatewayMulticastGroupsOutput{}
	fnCacher := func(out *ec2.SearchTransitGatewayMulticastGroupsOutput, more bool) bool {
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
	if err := c.EC2API.SearchTransitGatewayMulticastGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EC2) SearchTransitGatewayMulticastGroupsWithContext(ctx context.Context, input *ec2.SearchTransitGatewayMulticastGroupsInput, opts ...request.Option) (*ec2.SearchTransitGatewayMulticastGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayMulticastGroupsOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayMulticastGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayRoutes(input *ec2.SearchTransitGatewayRoutesInput) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EC2) SearchTransitGatewayRoutesWithContext(ctx context.Context, input *ec2.SearchTransitGatewayRoutesInput, opts ...request.Option) (*ec2.SearchTransitGatewayRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ec2.SearchTransitGatewayRoutesOutput), nil
	}
	out, err := c.EC2API.SearchTransitGatewayRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
