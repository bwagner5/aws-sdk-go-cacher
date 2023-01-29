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
	"github.com/aws/aws-sdk-go/service/gamelift"
	"github.com/aws/aws-sdk-go/service/gamelift/gameliftiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GameLift struct {
	gameliftiface.GameLiftAPI
	cache *cache.Cache
}

func NewGameLift(gameliftapi gameliftiface.GameLiftAPI) *GameLift {
	return &GameLift{
		GameLiftAPI: gameliftapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GameLift) DescribeAlias(input *gamelift.DescribeAliasInput) (*gamelift.DescribeAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeAliasOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeAliasWithContext(ctx context.Context, input *gamelift.DescribeAliasInput, opts ...request.Option) (*gamelift.DescribeAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeAliasOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeBuild(input *gamelift.DescribeBuildInput) (*gamelift.DescribeBuildOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeBuildOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeBuild(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeBuildWithContext(ctx context.Context, input *gamelift.DescribeBuildInput, opts ...request.Option) (*gamelift.DescribeBuildOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeBuildOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeBuildWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeCompute(input *gamelift.DescribeComputeInput) (*gamelift.DescribeComputeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeComputeOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeCompute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeComputeWithContext(ctx context.Context, input *gamelift.DescribeComputeInput, opts ...request.Option) (*gamelift.DescribeComputeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeComputeOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeComputeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeEC2InstanceLimits(input *gamelift.DescribeEC2InstanceLimitsInput) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeEC2InstanceLimitsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeEC2InstanceLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeEC2InstanceLimitsWithContext(ctx context.Context, input *gamelift.DescribeEC2InstanceLimitsInput, opts ...request.Option) (*gamelift.DescribeEC2InstanceLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeEC2InstanceLimitsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeEC2InstanceLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetAttributes(input *gamelift.DescribeFleetAttributesInput) (*gamelift.DescribeFleetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetAttributesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetAttributesPages(input *gamelift.DescribeFleetAttributesInput, fn func(*gamelift.DescribeFleetAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetAttributesOutput{}
	fnCacher := func(out *gamelift.DescribeFleetAttributesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetAttributesPagesWithContext(ctx context.Context, input *gamelift.DescribeFleetAttributesInput, fn func(*gamelift.DescribeFleetAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetAttributesOutput{}
	fnCacher := func(out *gamelift.DescribeFleetAttributesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetAttributesWithContext(ctx context.Context, input *gamelift.DescribeFleetAttributesInput, opts ...request.Option) (*gamelift.DescribeFleetAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetAttributesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetCapacity(input *gamelift.DescribeFleetCapacityInput) (*gamelift.DescribeFleetCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetCapacityOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetCapacity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetCapacityPages(input *gamelift.DescribeFleetCapacityInput, fn func(*gamelift.DescribeFleetCapacityOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetCapacityOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetCapacityOutput{}
	fnCacher := func(out *gamelift.DescribeFleetCapacityOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetCapacityPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetCapacityPagesWithContext(ctx context.Context, input *gamelift.DescribeFleetCapacityInput, fn func(*gamelift.DescribeFleetCapacityOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetCapacityOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetCapacityOutput{}
	fnCacher := func(out *gamelift.DescribeFleetCapacityOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetCapacityPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetCapacityWithContext(ctx context.Context, input *gamelift.DescribeFleetCapacityInput, opts ...request.Option) (*gamelift.DescribeFleetCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetCapacityOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetCapacityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetEvents(input *gamelift.DescribeFleetEventsInput) (*gamelift.DescribeFleetEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetEventsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetEventsPages(input *gamelift.DescribeFleetEventsInput, fn func(*gamelift.DescribeFleetEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetEventsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetEventsOutput{}
	fnCacher := func(out *gamelift.DescribeFleetEventsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetEventsPagesWithContext(ctx context.Context, input *gamelift.DescribeFleetEventsInput, fn func(*gamelift.DescribeFleetEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetEventsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetEventsOutput{}
	fnCacher := func(out *gamelift.DescribeFleetEventsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetEventsWithContext(ctx context.Context, input *gamelift.DescribeFleetEventsInput, opts ...request.Option) (*gamelift.DescribeFleetEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetEventsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationAttributes(input *gamelift.DescribeFleetLocationAttributesInput) (*gamelift.DescribeFleetLocationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationAttributesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationAttributesPages(input *gamelift.DescribeFleetLocationAttributesInput, fn func(*gamelift.DescribeFleetLocationAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetLocationAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetLocationAttributesOutput{}
	fnCacher := func(out *gamelift.DescribeFleetLocationAttributesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetLocationAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetLocationAttributesPagesWithContext(ctx context.Context, input *gamelift.DescribeFleetLocationAttributesInput, fn func(*gamelift.DescribeFleetLocationAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetLocationAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetLocationAttributesOutput{}
	fnCacher := func(out *gamelift.DescribeFleetLocationAttributesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetLocationAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetLocationAttributesWithContext(ctx context.Context, input *gamelift.DescribeFleetLocationAttributesInput, opts ...request.Option) (*gamelift.DescribeFleetLocationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationAttributesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationCapacity(input *gamelift.DescribeFleetLocationCapacityInput) (*gamelift.DescribeFleetLocationCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationCapacityOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationCapacity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationCapacityWithContext(ctx context.Context, input *gamelift.DescribeFleetLocationCapacityInput, opts ...request.Option) (*gamelift.DescribeFleetLocationCapacityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationCapacityOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationCapacityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationUtilization(input *gamelift.DescribeFleetLocationUtilizationInput) (*gamelift.DescribeFleetLocationUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationUtilizationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationUtilization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetLocationUtilizationWithContext(ctx context.Context, input *gamelift.DescribeFleetLocationUtilizationInput, opts ...request.Option) (*gamelift.DescribeFleetLocationUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetLocationUtilizationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetLocationUtilizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetPortSettings(input *gamelift.DescribeFleetPortSettingsInput) (*gamelift.DescribeFleetPortSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetPortSettingsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetPortSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetPortSettingsWithContext(ctx context.Context, input *gamelift.DescribeFleetPortSettingsInput, opts ...request.Option) (*gamelift.DescribeFleetPortSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetPortSettingsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetPortSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetUtilization(input *gamelift.DescribeFleetUtilizationInput) (*gamelift.DescribeFleetUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetUtilizationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetUtilization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeFleetUtilizationPages(input *gamelift.DescribeFleetUtilizationInput, fn func(*gamelift.DescribeFleetUtilizationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetUtilizationOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetUtilizationOutput{}
	fnCacher := func(out *gamelift.DescribeFleetUtilizationOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetUtilizationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetUtilizationPagesWithContext(ctx context.Context, input *gamelift.DescribeFleetUtilizationInput, fn func(*gamelift.DescribeFleetUtilizationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeFleetUtilizationOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeFleetUtilizationOutput{}
	fnCacher := func(out *gamelift.DescribeFleetUtilizationOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeFleetUtilizationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeFleetUtilizationWithContext(ctx context.Context, input *gamelift.DescribeFleetUtilizationInput, opts ...request.Option) (*gamelift.DescribeFleetUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeFleetUtilizationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeFleetUtilizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServer(input *gamelift.DescribeGameServerInput) (*gamelift.DescribeGameServerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServerGroup(input *gamelift.DescribeGameServerGroupInput) (*gamelift.DescribeGameServerGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerGroupOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServerGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServerGroupWithContext(ctx context.Context, input *gamelift.DescribeGameServerGroupInput, opts ...request.Option) (*gamelift.DescribeGameServerGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerGroupOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServerGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServerInstances(input *gamelift.DescribeGameServerInstancesInput) (*gamelift.DescribeGameServerInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerInstancesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServerInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServerInstancesPages(input *gamelift.DescribeGameServerInstancesInput, fn func(*gamelift.DescribeGameServerInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameServerInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameServerInstancesOutput{}
	fnCacher := func(out *gamelift.DescribeGameServerInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameServerInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameServerInstancesPagesWithContext(ctx context.Context, input *gamelift.DescribeGameServerInstancesInput, fn func(*gamelift.DescribeGameServerInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameServerInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameServerInstancesOutput{}
	fnCacher := func(out *gamelift.DescribeGameServerInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameServerInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameServerInstancesWithContext(ctx context.Context, input *gamelift.DescribeGameServerInstancesInput, opts ...request.Option) (*gamelift.DescribeGameServerInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerInstancesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServerInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameServerWithContext(ctx context.Context, input *gamelift.DescribeGameServerInput, opts ...request.Option) (*gamelift.DescribeGameServerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameServerOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameServerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionDetails(input *gamelift.DescribeGameSessionDetailsInput) (*gamelift.DescribeGameSessionDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionDetailsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionDetailsPages(input *gamelift.DescribeGameSessionDetailsInput, fn func(*gamelift.DescribeGameSessionDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionDetailsOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionDetailsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionDetailsPagesWithContext(ctx context.Context, input *gamelift.DescribeGameSessionDetailsInput, fn func(*gamelift.DescribeGameSessionDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionDetailsOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionDetailsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionDetailsWithContext(ctx context.Context, input *gamelift.DescribeGameSessionDetailsInput, opts ...request.Option) (*gamelift.DescribeGameSessionDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionDetailsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionPlacement(input *gamelift.DescribeGameSessionPlacementInput) (*gamelift.DescribeGameSessionPlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionPlacementOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionPlacement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionPlacementWithContext(ctx context.Context, input *gamelift.DescribeGameSessionPlacementInput, opts ...request.Option) (*gamelift.DescribeGameSessionPlacementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionPlacementOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionPlacementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionQueues(input *gamelift.DescribeGameSessionQueuesInput) (*gamelift.DescribeGameSessionQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionQueuesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionQueuesPages(input *gamelift.DescribeGameSessionQueuesInput, fn func(*gamelift.DescribeGameSessionQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionQueuesOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionQueuesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionQueuesPagesWithContext(ctx context.Context, input *gamelift.DescribeGameSessionQueuesInput, fn func(*gamelift.DescribeGameSessionQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionQueuesOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionQueuesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionQueuesWithContext(ctx context.Context, input *gamelift.DescribeGameSessionQueuesInput, opts ...request.Option) (*gamelift.DescribeGameSessionQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionQueuesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessions(input *gamelift.DescribeGameSessionsInput) (*gamelift.DescribeGameSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeGameSessionsPages(input *gamelift.DescribeGameSessionsInput, fn func(*gamelift.DescribeGameSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionsOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionsPagesWithContext(ctx context.Context, input *gamelift.DescribeGameSessionsInput, fn func(*gamelift.DescribeGameSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeGameSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeGameSessionsOutput{}
	fnCacher := func(out *gamelift.DescribeGameSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeGameSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeGameSessionsWithContext(ctx context.Context, input *gamelift.DescribeGameSessionsInput, opts ...request.Option) (*gamelift.DescribeGameSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeGameSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeGameSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeInstances(input *gamelift.DescribeInstancesInput) (*gamelift.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeInstancesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeInstancesPages(input *gamelift.DescribeInstancesInput, fn func(*gamelift.DescribeInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeInstancesOutput{}
	fnCacher := func(out *gamelift.DescribeInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeInstancesPagesWithContext(ctx context.Context, input *gamelift.DescribeInstancesInput, fn func(*gamelift.DescribeInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeInstancesOutput{}
	fnCacher := func(out *gamelift.DescribeInstancesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeInstancesWithContext(ctx context.Context, input *gamelift.DescribeInstancesInput, opts ...request.Option) (*gamelift.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeInstancesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmaking(input *gamelift.DescribeMatchmakingInput) (*gamelift.DescribeMatchmakingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmaking(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmakingConfigurations(input *gamelift.DescribeMatchmakingConfigurationsInput) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingConfigurationsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmakingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmakingConfigurationsPages(input *gamelift.DescribeMatchmakingConfigurationsInput, fn func(*gamelift.DescribeMatchmakingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeMatchmakingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeMatchmakingConfigurationsOutput{}
	fnCacher := func(out *gamelift.DescribeMatchmakingConfigurationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeMatchmakingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeMatchmakingConfigurationsPagesWithContext(ctx context.Context, input *gamelift.DescribeMatchmakingConfigurationsInput, fn func(*gamelift.DescribeMatchmakingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeMatchmakingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeMatchmakingConfigurationsOutput{}
	fnCacher := func(out *gamelift.DescribeMatchmakingConfigurationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeMatchmakingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeMatchmakingConfigurationsWithContext(ctx context.Context, input *gamelift.DescribeMatchmakingConfigurationsInput, opts ...request.Option) (*gamelift.DescribeMatchmakingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingConfigurationsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmakingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmakingRuleSets(input *gamelift.DescribeMatchmakingRuleSetsInput) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingRuleSetsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmakingRuleSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmakingRuleSetsPages(input *gamelift.DescribeMatchmakingRuleSetsInput, fn func(*gamelift.DescribeMatchmakingRuleSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeMatchmakingRuleSetsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeMatchmakingRuleSetsOutput{}
	fnCacher := func(out *gamelift.DescribeMatchmakingRuleSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeMatchmakingRuleSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeMatchmakingRuleSetsPagesWithContext(ctx context.Context, input *gamelift.DescribeMatchmakingRuleSetsInput, fn func(*gamelift.DescribeMatchmakingRuleSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeMatchmakingRuleSetsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeMatchmakingRuleSetsOutput{}
	fnCacher := func(out *gamelift.DescribeMatchmakingRuleSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeMatchmakingRuleSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeMatchmakingRuleSetsWithContext(ctx context.Context, input *gamelift.DescribeMatchmakingRuleSetsInput, opts ...request.Option) (*gamelift.DescribeMatchmakingRuleSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingRuleSetsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmakingRuleSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeMatchmakingWithContext(ctx context.Context, input *gamelift.DescribeMatchmakingInput, opts ...request.Option) (*gamelift.DescribeMatchmakingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeMatchmakingOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeMatchmakingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribePlayerSessions(input *gamelift.DescribePlayerSessionsInput) (*gamelift.DescribePlayerSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribePlayerSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribePlayerSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribePlayerSessionsPages(input *gamelift.DescribePlayerSessionsInput, fn func(*gamelift.DescribePlayerSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribePlayerSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribePlayerSessionsOutput{}
	fnCacher := func(out *gamelift.DescribePlayerSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribePlayerSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribePlayerSessionsPagesWithContext(ctx context.Context, input *gamelift.DescribePlayerSessionsInput, fn func(*gamelift.DescribePlayerSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribePlayerSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribePlayerSessionsOutput{}
	fnCacher := func(out *gamelift.DescribePlayerSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribePlayerSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribePlayerSessionsWithContext(ctx context.Context, input *gamelift.DescribePlayerSessionsInput, opts ...request.Option) (*gamelift.DescribePlayerSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribePlayerSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribePlayerSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeRuntimeConfiguration(input *gamelift.DescribeRuntimeConfigurationInput) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeRuntimeConfigurationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeRuntimeConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeRuntimeConfigurationWithContext(ctx context.Context, input *gamelift.DescribeRuntimeConfigurationInput, opts ...request.Option) (*gamelift.DescribeRuntimeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeRuntimeConfigurationOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeRuntimeConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeScalingPolicies(input *gamelift.DescribeScalingPoliciesInput) (*gamelift.DescribeScalingPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeScalingPoliciesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeScalingPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeScalingPoliciesPages(input *gamelift.DescribeScalingPoliciesInput, fn func(*gamelift.DescribeScalingPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeScalingPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeScalingPoliciesOutput{}
	fnCacher := func(out *gamelift.DescribeScalingPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeScalingPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeScalingPoliciesPagesWithContext(ctx context.Context, input *gamelift.DescribeScalingPoliciesInput, fn func(*gamelift.DescribeScalingPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.DescribeScalingPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.DescribeScalingPoliciesOutput{}
	fnCacher := func(out *gamelift.DescribeScalingPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.DescribeScalingPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) DescribeScalingPoliciesWithContext(ctx context.Context, input *gamelift.DescribeScalingPoliciesInput, opts ...request.Option) (*gamelift.DescribeScalingPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeScalingPoliciesOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeScalingPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeScript(input *gamelift.DescribeScriptInput) (*gamelift.DescribeScriptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeScriptOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeScript(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeScriptWithContext(ctx context.Context, input *gamelift.DescribeScriptInput, opts ...request.Option) (*gamelift.DescribeScriptOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeScriptOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeScriptWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeVpcPeeringAuthorizations(input *gamelift.DescribeVpcPeeringAuthorizationsInput) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeVpcPeeringAuthorizationsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeVpcPeeringAuthorizations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeVpcPeeringAuthorizationsWithContext(ctx context.Context, input *gamelift.DescribeVpcPeeringAuthorizationsInput, opts ...request.Option) (*gamelift.DescribeVpcPeeringAuthorizationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeVpcPeeringAuthorizationsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeVpcPeeringAuthorizationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeVpcPeeringConnections(input *gamelift.DescribeVpcPeeringConnectionsInput) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeVpcPeeringConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) DescribeVpcPeeringConnectionsWithContext(ctx context.Context, input *gamelift.DescribeVpcPeeringConnectionsInput, opts ...request.Option) (*gamelift.DescribeVpcPeeringConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.DescribeVpcPeeringConnectionsOutput), nil
	}
	out, err := c.GameLiftAPI.DescribeVpcPeeringConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetComputeAccess(input *gamelift.GetComputeAccessInput) (*gamelift.GetComputeAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetComputeAccessOutput), nil
	}
	out, err := c.GameLiftAPI.GetComputeAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetComputeAccessWithContext(ctx context.Context, input *gamelift.GetComputeAccessInput, opts ...request.Option) (*gamelift.GetComputeAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetComputeAccessOutput), nil
	}
	out, err := c.GameLiftAPI.GetComputeAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetComputeAuthToken(input *gamelift.GetComputeAuthTokenInput) (*gamelift.GetComputeAuthTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetComputeAuthTokenOutput), nil
	}
	out, err := c.GameLiftAPI.GetComputeAuthToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetComputeAuthTokenWithContext(ctx context.Context, input *gamelift.GetComputeAuthTokenInput, opts ...request.Option) (*gamelift.GetComputeAuthTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetComputeAuthTokenOutput), nil
	}
	out, err := c.GameLiftAPI.GetComputeAuthTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetGameSessionLogUrl(input *gamelift.GetGameSessionLogUrlInput) (*gamelift.GetGameSessionLogUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetGameSessionLogUrlOutput), nil
	}
	out, err := c.GameLiftAPI.GetGameSessionLogUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetGameSessionLogUrlWithContext(ctx context.Context, input *gamelift.GetGameSessionLogUrlInput, opts ...request.Option) (*gamelift.GetGameSessionLogUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetGameSessionLogUrlOutput), nil
	}
	out, err := c.GameLiftAPI.GetGameSessionLogUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetInstanceAccess(input *gamelift.GetInstanceAccessInput) (*gamelift.GetInstanceAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetInstanceAccessOutput), nil
	}
	out, err := c.GameLiftAPI.GetInstanceAccess(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) GetInstanceAccessWithContext(ctx context.Context, input *gamelift.GetInstanceAccessInput, opts ...request.Option) (*gamelift.GetInstanceAccessOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.GetInstanceAccessOutput), nil
	}
	out, err := c.GameLiftAPI.GetInstanceAccessWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListAliases(input *gamelift.ListAliasesInput) (*gamelift.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListAliasesOutput), nil
	}
	out, err := c.GameLiftAPI.ListAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListAliasesPages(input *gamelift.ListAliasesInput, fn func(*gamelift.ListAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListAliasesOutput{}
	fnCacher := func(out *gamelift.ListAliasesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListAliasesPagesWithContext(ctx context.Context, input *gamelift.ListAliasesInput, fn func(*gamelift.ListAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListAliasesOutput{}
	fnCacher := func(out *gamelift.ListAliasesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListAliasesWithContext(ctx context.Context, input *gamelift.ListAliasesInput, opts ...request.Option) (*gamelift.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListAliasesOutput), nil
	}
	out, err := c.GameLiftAPI.ListAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListBuilds(input *gamelift.ListBuildsInput) (*gamelift.ListBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListBuildsOutput), nil
	}
	out, err := c.GameLiftAPI.ListBuilds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListBuildsPages(input *gamelift.ListBuildsInput, fn func(*gamelift.ListBuildsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListBuildsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListBuildsOutput{}
	fnCacher := func(out *gamelift.ListBuildsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListBuildsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListBuildsPagesWithContext(ctx context.Context, input *gamelift.ListBuildsInput, fn func(*gamelift.ListBuildsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListBuildsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListBuildsOutput{}
	fnCacher := func(out *gamelift.ListBuildsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListBuildsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListBuildsWithContext(ctx context.Context, input *gamelift.ListBuildsInput, opts ...request.Option) (*gamelift.ListBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListBuildsOutput), nil
	}
	out, err := c.GameLiftAPI.ListBuildsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListCompute(input *gamelift.ListComputeInput) (*gamelift.ListComputeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListComputeOutput), nil
	}
	out, err := c.GameLiftAPI.ListCompute(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListComputePages(input *gamelift.ListComputeInput, fn func(*gamelift.ListComputeOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListComputeOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListComputeOutput{}
	fnCacher := func(out *gamelift.ListComputeOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListComputePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListComputePagesWithContext(ctx context.Context, input *gamelift.ListComputeInput, fn func(*gamelift.ListComputeOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListComputeOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListComputeOutput{}
	fnCacher := func(out *gamelift.ListComputeOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListComputePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListComputeWithContext(ctx context.Context, input *gamelift.ListComputeInput, opts ...request.Option) (*gamelift.ListComputeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListComputeOutput), nil
	}
	out, err := c.GameLiftAPI.ListComputeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListFleets(input *gamelift.ListFleetsInput) (*gamelift.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListFleetsOutput), nil
	}
	out, err := c.GameLiftAPI.ListFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListFleetsPages(input *gamelift.ListFleetsInput, fn func(*gamelift.ListFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListFleetsOutput{}
	fnCacher := func(out *gamelift.ListFleetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListFleetsPagesWithContext(ctx context.Context, input *gamelift.ListFleetsInput, fn func(*gamelift.ListFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListFleetsOutput{}
	fnCacher := func(out *gamelift.ListFleetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListFleetsWithContext(ctx context.Context, input *gamelift.ListFleetsInput, opts ...request.Option) (*gamelift.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListFleetsOutput), nil
	}
	out, err := c.GameLiftAPI.ListFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListGameServerGroups(input *gamelift.ListGameServerGroupsInput) (*gamelift.ListGameServerGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListGameServerGroupsOutput), nil
	}
	out, err := c.GameLiftAPI.ListGameServerGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListGameServerGroupsPages(input *gamelift.ListGameServerGroupsInput, fn func(*gamelift.ListGameServerGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListGameServerGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListGameServerGroupsOutput{}
	fnCacher := func(out *gamelift.ListGameServerGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListGameServerGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListGameServerGroupsPagesWithContext(ctx context.Context, input *gamelift.ListGameServerGroupsInput, fn func(*gamelift.ListGameServerGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListGameServerGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListGameServerGroupsOutput{}
	fnCacher := func(out *gamelift.ListGameServerGroupsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListGameServerGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListGameServerGroupsWithContext(ctx context.Context, input *gamelift.ListGameServerGroupsInput, opts ...request.Option) (*gamelift.ListGameServerGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListGameServerGroupsOutput), nil
	}
	out, err := c.GameLiftAPI.ListGameServerGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListGameServers(input *gamelift.ListGameServersInput) (*gamelift.ListGameServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListGameServersOutput), nil
	}
	out, err := c.GameLiftAPI.ListGameServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListGameServersPages(input *gamelift.ListGameServersInput, fn func(*gamelift.ListGameServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListGameServersOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListGameServersOutput{}
	fnCacher := func(out *gamelift.ListGameServersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListGameServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListGameServersPagesWithContext(ctx context.Context, input *gamelift.ListGameServersInput, fn func(*gamelift.ListGameServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListGameServersOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListGameServersOutput{}
	fnCacher := func(out *gamelift.ListGameServersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListGameServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListGameServersWithContext(ctx context.Context, input *gamelift.ListGameServersInput, opts ...request.Option) (*gamelift.ListGameServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListGameServersOutput), nil
	}
	out, err := c.GameLiftAPI.ListGameServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListLocations(input *gamelift.ListLocationsInput) (*gamelift.ListLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListLocationsOutput), nil
	}
	out, err := c.GameLiftAPI.ListLocations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListLocationsPages(input *gamelift.ListLocationsInput, fn func(*gamelift.ListLocationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListLocationsOutput{}
	fnCacher := func(out *gamelift.ListLocationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListLocationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListLocationsPagesWithContext(ctx context.Context, input *gamelift.ListLocationsInput, fn func(*gamelift.ListLocationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListLocationsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListLocationsOutput{}
	fnCacher := func(out *gamelift.ListLocationsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListLocationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListLocationsWithContext(ctx context.Context, input *gamelift.ListLocationsInput, opts ...request.Option) (*gamelift.ListLocationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListLocationsOutput), nil
	}
	out, err := c.GameLiftAPI.ListLocationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListScripts(input *gamelift.ListScriptsInput) (*gamelift.ListScriptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListScriptsOutput), nil
	}
	out, err := c.GameLiftAPI.ListScripts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListScriptsPages(input *gamelift.ListScriptsInput, fn func(*gamelift.ListScriptsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListScriptsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListScriptsOutput{}
	fnCacher := func(out *gamelift.ListScriptsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListScriptsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListScriptsPagesWithContext(ctx context.Context, input *gamelift.ListScriptsInput, fn func(*gamelift.ListScriptsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.ListScriptsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.ListScriptsOutput{}
	fnCacher := func(out *gamelift.ListScriptsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.ListScriptsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) ListScriptsWithContext(ctx context.Context, input *gamelift.ListScriptsInput, opts ...request.Option) (*gamelift.ListScriptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListScriptsOutput), nil
	}
	out, err := c.GameLiftAPI.ListScriptsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListTagsForResource(input *gamelift.ListTagsForResourceInput) (*gamelift.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListTagsForResourceOutput), nil
	}
	out, err := c.GameLiftAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) ListTagsForResourceWithContext(ctx context.Context, input *gamelift.ListTagsForResourceInput, opts ...request.Option) (*gamelift.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.ListTagsForResourceOutput), nil
	}
	out, err := c.GameLiftAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) SearchGameSessions(input *gamelift.SearchGameSessionsInput) (*gamelift.SearchGameSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.SearchGameSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.SearchGameSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GameLift) SearchGameSessionsPages(input *gamelift.SearchGameSessionsInput, fn func(*gamelift.SearchGameSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.SearchGameSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.SearchGameSessionsOutput{}
	fnCacher := func(out *gamelift.SearchGameSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.SearchGameSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) SearchGameSessionsPagesWithContext(ctx context.Context, input *gamelift.SearchGameSessionsInput, fn func(*gamelift.SearchGameSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gamelift.SearchGameSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &gamelift.SearchGameSessionsOutput{}
	fnCacher := func(out *gamelift.SearchGameSessionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.GameLiftAPI.SearchGameSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GameLift) SearchGameSessionsWithContext(ctx context.Context, input *gamelift.SearchGameSessionsInput, opts ...request.Option) (*gamelift.SearchGameSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gamelift.SearchGameSessionsOutput), nil
	}
	out, err := c.GameLiftAPI.SearchGameSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
