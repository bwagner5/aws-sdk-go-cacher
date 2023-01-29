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
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/groundstation/groundstationiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GroundStation struct {
	groundstationiface.GroundStationAPI
	cache *cache.Cache
}

func NewGroundStation(groundstationapi groundstationiface.GroundStationAPI) *GroundStation {
	return &GroundStation{
		GroundStationAPI: groundstationapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GroundStation) DescribeContact(input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.DescribeContactOutput), nil
	}
	out, err := c.GroundStationAPI.DescribeContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) DescribeContactWithContext(ctx context.Context, input *groundstation.DescribeContactInput, opts ...request.Option) (*groundstation.DescribeContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.DescribeContactOutput), nil
	}
	out, err := c.GroundStationAPI.DescribeContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) DescribeEphemeris(input *groundstation.DescribeEphemerisInput) (*groundstation.DescribeEphemerisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.DescribeEphemerisOutput), nil
	}
	out, err := c.GroundStationAPI.DescribeEphemeris(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) DescribeEphemerisWithContext(ctx context.Context, input *groundstation.DescribeEphemerisInput, opts ...request.Option) (*groundstation.DescribeEphemerisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.DescribeEphemerisOutput), nil
	}
	out, err := c.GroundStationAPI.DescribeEphemerisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetConfig(input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetConfigOutput), nil
	}
	out, err := c.GroundStationAPI.GetConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetConfigWithContext(ctx context.Context, input *groundstation.GetConfigInput, opts ...request.Option) (*groundstation.GetConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetConfigOutput), nil
	}
	out, err := c.GroundStationAPI.GetConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetDataflowEndpointGroup(input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetDataflowEndpointGroupOutput), nil
	}
	out, err := c.GroundStationAPI.GetDataflowEndpointGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetDataflowEndpointGroupWithContext(ctx context.Context, input *groundstation.GetDataflowEndpointGroupInput, opts ...request.Option) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetDataflowEndpointGroupOutput), nil
	}
	out, err := c.GroundStationAPI.GetDataflowEndpointGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetMinuteUsage(input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetMinuteUsageOutput), nil
	}
	out, err := c.GroundStationAPI.GetMinuteUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetMinuteUsageWithContext(ctx context.Context, input *groundstation.GetMinuteUsageInput, opts ...request.Option) (*groundstation.GetMinuteUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetMinuteUsageOutput), nil
	}
	out, err := c.GroundStationAPI.GetMinuteUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetMissionProfile(input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetMissionProfileOutput), nil
	}
	out, err := c.GroundStationAPI.GetMissionProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetMissionProfileWithContext(ctx context.Context, input *groundstation.GetMissionProfileInput, opts ...request.Option) (*groundstation.GetMissionProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetMissionProfileOutput), nil
	}
	out, err := c.GroundStationAPI.GetMissionProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetSatellite(input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetSatelliteOutput), nil
	}
	out, err := c.GroundStationAPI.GetSatellite(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) GetSatelliteWithContext(ctx context.Context, input *groundstation.GetSatelliteInput, opts ...request.Option) (*groundstation.GetSatelliteOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.GetSatelliteOutput), nil
	}
	out, err := c.GroundStationAPI.GetSatelliteWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListConfigs(input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListConfigsOutput), nil
	}
	out, err := c.GroundStationAPI.ListConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListConfigsPages(input *groundstation.ListConfigsInput, fn func(*groundstation.ListConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListConfigsOutput{}
	fnCacher := func(out *groundstation.ListConfigsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListConfigsPagesWithContext(ctx context.Context, input *groundstation.ListConfigsInput, fn func(*groundstation.ListConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListConfigsOutput{}
	fnCacher := func(out *groundstation.ListConfigsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListConfigsWithContext(ctx context.Context, input *groundstation.ListConfigsInput, opts ...request.Option) (*groundstation.ListConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListConfigsOutput), nil
	}
	out, err := c.GroundStationAPI.ListConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListContacts(input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListContactsOutput), nil
	}
	out, err := c.GroundStationAPI.ListContacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListContactsPages(input *groundstation.ListContactsInput, fn func(*groundstation.ListContactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListContactsOutput{}
	fnCacher := func(out *groundstation.ListContactsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListContactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListContactsPagesWithContext(ctx context.Context, input *groundstation.ListContactsInput, fn func(*groundstation.ListContactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListContactsOutput{}
	fnCacher := func(out *groundstation.ListContactsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListContactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListContactsWithContext(ctx context.Context, input *groundstation.ListContactsInput, opts ...request.Option) (*groundstation.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListContactsOutput), nil
	}
	out, err := c.GroundStationAPI.ListContactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListDataflowEndpointGroups(input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListDataflowEndpointGroupsOutput), nil
	}
	out, err := c.GroundStationAPI.ListDataflowEndpointGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListDataflowEndpointGroupsPages(input *groundstation.ListDataflowEndpointGroupsInput, fn func(*groundstation.ListDataflowEndpointGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListDataflowEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListDataflowEndpointGroupsOutput{}
	fnCacher := func(out *groundstation.ListDataflowEndpointGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListDataflowEndpointGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListDataflowEndpointGroupsPagesWithContext(ctx context.Context, input *groundstation.ListDataflowEndpointGroupsInput, fn func(*groundstation.ListDataflowEndpointGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListDataflowEndpointGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListDataflowEndpointGroupsOutput{}
	fnCacher := func(out *groundstation.ListDataflowEndpointGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListDataflowEndpointGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListDataflowEndpointGroupsWithContext(ctx context.Context, input *groundstation.ListDataflowEndpointGroupsInput, opts ...request.Option) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListDataflowEndpointGroupsOutput), nil
	}
	out, err := c.GroundStationAPI.ListDataflowEndpointGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListEphemerides(input *groundstation.ListEphemeridesInput) (*groundstation.ListEphemeridesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListEphemeridesOutput), nil
	}
	out, err := c.GroundStationAPI.ListEphemerides(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListEphemeridesPages(input *groundstation.ListEphemeridesInput, fn func(*groundstation.ListEphemeridesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListEphemeridesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListEphemeridesOutput{}
	fnCacher := func(out *groundstation.ListEphemeridesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListEphemeridesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListEphemeridesPagesWithContext(ctx context.Context, input *groundstation.ListEphemeridesInput, fn func(*groundstation.ListEphemeridesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListEphemeridesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListEphemeridesOutput{}
	fnCacher := func(out *groundstation.ListEphemeridesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListEphemeridesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListEphemeridesWithContext(ctx context.Context, input *groundstation.ListEphemeridesInput, opts ...request.Option) (*groundstation.ListEphemeridesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListEphemeridesOutput), nil
	}
	out, err := c.GroundStationAPI.ListEphemeridesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListGroundStations(input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListGroundStationsOutput), nil
	}
	out, err := c.GroundStationAPI.ListGroundStations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListGroundStationsPages(input *groundstation.ListGroundStationsInput, fn func(*groundstation.ListGroundStationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListGroundStationsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListGroundStationsOutput{}
	fnCacher := func(out *groundstation.ListGroundStationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListGroundStationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListGroundStationsPagesWithContext(ctx context.Context, input *groundstation.ListGroundStationsInput, fn func(*groundstation.ListGroundStationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListGroundStationsOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListGroundStationsOutput{}
	fnCacher := func(out *groundstation.ListGroundStationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListGroundStationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListGroundStationsWithContext(ctx context.Context, input *groundstation.ListGroundStationsInput, opts ...request.Option) (*groundstation.ListGroundStationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListGroundStationsOutput), nil
	}
	out, err := c.GroundStationAPI.ListGroundStationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListMissionProfiles(input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListMissionProfilesOutput), nil
	}
	out, err := c.GroundStationAPI.ListMissionProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListMissionProfilesPages(input *groundstation.ListMissionProfilesInput, fn func(*groundstation.ListMissionProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListMissionProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListMissionProfilesOutput{}
	fnCacher := func(out *groundstation.ListMissionProfilesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListMissionProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListMissionProfilesPagesWithContext(ctx context.Context, input *groundstation.ListMissionProfilesInput, fn func(*groundstation.ListMissionProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListMissionProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListMissionProfilesOutput{}
	fnCacher := func(out *groundstation.ListMissionProfilesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListMissionProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListMissionProfilesWithContext(ctx context.Context, input *groundstation.ListMissionProfilesInput, opts ...request.Option) (*groundstation.ListMissionProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListMissionProfilesOutput), nil
	}
	out, err := c.GroundStationAPI.ListMissionProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListSatellites(input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListSatellitesOutput), nil
	}
	out, err := c.GroundStationAPI.ListSatellites(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListSatellitesPages(input *groundstation.ListSatellitesInput, fn func(*groundstation.ListSatellitesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListSatellitesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListSatellitesOutput{}
	fnCacher := func(out *groundstation.ListSatellitesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListSatellitesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListSatellitesPagesWithContext(ctx context.Context, input *groundstation.ListSatellitesInput, fn func(*groundstation.ListSatellitesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*groundstation.ListSatellitesOutput), false)
		return nil
	}
	cachable := true
	output := &groundstation.ListSatellitesOutput{}
	fnCacher := func(out *groundstation.ListSatellitesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GroundStationAPI.ListSatellitesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GroundStation) ListSatellitesWithContext(ctx context.Context, input *groundstation.ListSatellitesInput, opts ...request.Option) (*groundstation.ListSatellitesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListSatellitesOutput), nil
	}
	out, err := c.GroundStationAPI.ListSatellitesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListTagsForResource(input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListTagsForResourceOutput), nil
	}
	out, err := c.GroundStationAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GroundStation) ListTagsForResourceWithContext(ctx context.Context, input *groundstation.ListTagsForResourceInput, opts ...request.Option) (*groundstation.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*groundstation.ListTagsForResourceOutput), nil
	}
	out, err := c.GroundStationAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
