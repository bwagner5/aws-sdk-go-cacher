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
	"github.com/aws/aws-sdk-go/service/iottwinmaker"
	"github.com/aws/aws-sdk-go/service/iottwinmaker/iottwinmakeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTTwinMaker struct {
	iottwinmakeriface.IoTTwinMakerAPI
	cache *cache.Cache
}

func NewIoTTwinMaker(iottwinmakerapi iottwinmakeriface.IoTTwinMakerAPI) *IoTTwinMaker {
	return &IoTTwinMaker{
		IoTTwinMakerAPI: iottwinmakerapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTTwinMaker) BatchPutPropertyValues(input *iottwinmaker.BatchPutPropertyValuesInput) (*iottwinmaker.BatchPutPropertyValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.BatchPutPropertyValuesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.BatchPutPropertyValues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) BatchPutPropertyValuesWithContext(ctx context.Context, input *iottwinmaker.BatchPutPropertyValuesInput, opts ...request.Option) (*iottwinmaker.BatchPutPropertyValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.BatchPutPropertyValuesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.BatchPutPropertyValuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetComponentType(input *iottwinmaker.GetComponentTypeInput) (*iottwinmaker.GetComponentTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetComponentTypeOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetComponentType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetComponentTypeWithContext(ctx context.Context, input *iottwinmaker.GetComponentTypeInput, opts ...request.Option) (*iottwinmaker.GetComponentTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetComponentTypeOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetComponentTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetEntity(input *iottwinmaker.GetEntityInput) (*iottwinmaker.GetEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetEntityOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetEntity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetEntityWithContext(ctx context.Context, input *iottwinmaker.GetEntityInput, opts ...request.Option) (*iottwinmaker.GetEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetEntityOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetEntityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPricingPlan(input *iottwinmaker.GetPricingPlanInput) (*iottwinmaker.GetPricingPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPricingPlanOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPricingPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPricingPlanWithContext(ctx context.Context, input *iottwinmaker.GetPricingPlanInput, opts ...request.Option) (*iottwinmaker.GetPricingPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPricingPlanOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPricingPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPropertyValue(input *iottwinmaker.GetPropertyValueInput) (*iottwinmaker.GetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPropertyValueOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPropertyValue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPropertyValueHistory(input *iottwinmaker.GetPropertyValueHistoryInput) (*iottwinmaker.GetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPropertyValueHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPropertyValueHistoryPages(input *iottwinmaker.GetPropertyValueHistoryInput, fn func(*iottwinmaker.GetPropertyValueHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.GetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.GetPropertyValueHistoryOutput{}
	fnCacher := func(out *iottwinmaker.GetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.GetPropertyValueHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) GetPropertyValueHistoryPagesWithContext(ctx context.Context, input *iottwinmaker.GetPropertyValueHistoryInput, fn func(*iottwinmaker.GetPropertyValueHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.GetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.GetPropertyValueHistoryOutput{}
	fnCacher := func(out *iottwinmaker.GetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.GetPropertyValueHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) GetPropertyValueHistoryWithContext(ctx context.Context, input *iottwinmaker.GetPropertyValueHistoryInput, opts ...request.Option) (*iottwinmaker.GetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPropertyValueHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetPropertyValuePages(input *iottwinmaker.GetPropertyValueInput, fn func(*iottwinmaker.GetPropertyValueOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.GetPropertyValueOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.GetPropertyValueOutput{}
	fnCacher := func(out *iottwinmaker.GetPropertyValueOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.GetPropertyValuePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) GetPropertyValuePagesWithContext(ctx context.Context, input *iottwinmaker.GetPropertyValueInput, fn func(*iottwinmaker.GetPropertyValueOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.GetPropertyValueOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.GetPropertyValueOutput{}
	fnCacher := func(out *iottwinmaker.GetPropertyValueOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.GetPropertyValuePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) GetPropertyValueWithContext(ctx context.Context, input *iottwinmaker.GetPropertyValueInput, opts ...request.Option) (*iottwinmaker.GetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetPropertyValueOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetPropertyValueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetScene(input *iottwinmaker.GetSceneInput) (*iottwinmaker.GetSceneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetSceneOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetScene(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetSceneWithContext(ctx context.Context, input *iottwinmaker.GetSceneInput, opts ...request.Option) (*iottwinmaker.GetSceneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetSceneOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetSceneWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetSyncJob(input *iottwinmaker.GetSyncJobInput) (*iottwinmaker.GetSyncJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetSyncJobOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetSyncJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetSyncJobWithContext(ctx context.Context, input *iottwinmaker.GetSyncJobInput, opts ...request.Option) (*iottwinmaker.GetSyncJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetSyncJobOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetSyncJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetWorkspace(input *iottwinmaker.GetWorkspaceInput) (*iottwinmaker.GetWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetWorkspaceOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetWorkspace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) GetWorkspaceWithContext(ctx context.Context, input *iottwinmaker.GetWorkspaceInput, opts ...request.Option) (*iottwinmaker.GetWorkspaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.GetWorkspaceOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.GetWorkspaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListComponentTypes(input *iottwinmaker.ListComponentTypesInput) (*iottwinmaker.ListComponentTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListComponentTypesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListComponentTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListComponentTypesPages(input *iottwinmaker.ListComponentTypesInput, fn func(*iottwinmaker.ListComponentTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListComponentTypesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListComponentTypesOutput{}
	fnCacher := func(out *iottwinmaker.ListComponentTypesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListComponentTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListComponentTypesPagesWithContext(ctx context.Context, input *iottwinmaker.ListComponentTypesInput, fn func(*iottwinmaker.ListComponentTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListComponentTypesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListComponentTypesOutput{}
	fnCacher := func(out *iottwinmaker.ListComponentTypesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListComponentTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListComponentTypesWithContext(ctx context.Context, input *iottwinmaker.ListComponentTypesInput, opts ...request.Option) (*iottwinmaker.ListComponentTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListComponentTypesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListComponentTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListEntities(input *iottwinmaker.ListEntitiesInput) (*iottwinmaker.ListEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListEntitiesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListEntitiesPages(input *iottwinmaker.ListEntitiesInput, fn func(*iottwinmaker.ListEntitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListEntitiesOutput{}
	fnCacher := func(out *iottwinmaker.ListEntitiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListEntitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListEntitiesPagesWithContext(ctx context.Context, input *iottwinmaker.ListEntitiesInput, fn func(*iottwinmaker.ListEntitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListEntitiesOutput{}
	fnCacher := func(out *iottwinmaker.ListEntitiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListEntitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListEntitiesWithContext(ctx context.Context, input *iottwinmaker.ListEntitiesInput, opts ...request.Option) (*iottwinmaker.ListEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListEntitiesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListScenes(input *iottwinmaker.ListScenesInput) (*iottwinmaker.ListScenesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListScenesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListScenes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListScenesPages(input *iottwinmaker.ListScenesInput, fn func(*iottwinmaker.ListScenesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListScenesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListScenesOutput{}
	fnCacher := func(out *iottwinmaker.ListScenesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListScenesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListScenesPagesWithContext(ctx context.Context, input *iottwinmaker.ListScenesInput, fn func(*iottwinmaker.ListScenesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListScenesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListScenesOutput{}
	fnCacher := func(out *iottwinmaker.ListScenesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListScenesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListScenesWithContext(ctx context.Context, input *iottwinmaker.ListScenesInput, opts ...request.Option) (*iottwinmaker.ListScenesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListScenesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListScenesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListSyncJobs(input *iottwinmaker.ListSyncJobsInput) (*iottwinmaker.ListSyncJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListSyncJobsOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListSyncJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListSyncJobsPages(input *iottwinmaker.ListSyncJobsInput, fn func(*iottwinmaker.ListSyncJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListSyncJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListSyncJobsOutput{}
	fnCacher := func(out *iottwinmaker.ListSyncJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListSyncJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListSyncJobsPagesWithContext(ctx context.Context, input *iottwinmaker.ListSyncJobsInput, fn func(*iottwinmaker.ListSyncJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListSyncJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListSyncJobsOutput{}
	fnCacher := func(out *iottwinmaker.ListSyncJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListSyncJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListSyncJobsWithContext(ctx context.Context, input *iottwinmaker.ListSyncJobsInput, opts ...request.Option) (*iottwinmaker.ListSyncJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListSyncJobsOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListSyncJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListSyncResources(input *iottwinmaker.ListSyncResourcesInput) (*iottwinmaker.ListSyncResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListSyncResourcesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListSyncResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListSyncResourcesPages(input *iottwinmaker.ListSyncResourcesInput, fn func(*iottwinmaker.ListSyncResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListSyncResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListSyncResourcesOutput{}
	fnCacher := func(out *iottwinmaker.ListSyncResourcesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListSyncResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListSyncResourcesPagesWithContext(ctx context.Context, input *iottwinmaker.ListSyncResourcesInput, fn func(*iottwinmaker.ListSyncResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListSyncResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListSyncResourcesOutput{}
	fnCacher := func(out *iottwinmaker.ListSyncResourcesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListSyncResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListSyncResourcesWithContext(ctx context.Context, input *iottwinmaker.ListSyncResourcesInput, opts ...request.Option) (*iottwinmaker.ListSyncResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListSyncResourcesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListSyncResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListTagsForResource(input *iottwinmaker.ListTagsForResourceInput) (*iottwinmaker.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListTagsForResourceWithContext(ctx context.Context, input *iottwinmaker.ListTagsForResourceInput, opts ...request.Option) (*iottwinmaker.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListWorkspaces(input *iottwinmaker.ListWorkspacesInput) (*iottwinmaker.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListWorkspacesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListWorkspaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTTwinMaker) ListWorkspacesPages(input *iottwinmaker.ListWorkspacesInput, fn func(*iottwinmaker.ListWorkspacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListWorkspacesOutput{}
	fnCacher := func(out *iottwinmaker.ListWorkspacesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListWorkspacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListWorkspacesPagesWithContext(ctx context.Context, input *iottwinmaker.ListWorkspacesInput, fn func(*iottwinmaker.ListWorkspacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iottwinmaker.ListWorkspacesOutput), false)
		return nil
	}
	cachable := true
	output := &iottwinmaker.ListWorkspacesOutput{}
	fnCacher := func(out *iottwinmaker.ListWorkspacesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTTwinMakerAPI.ListWorkspacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTTwinMaker) ListWorkspacesWithContext(ctx context.Context, input *iottwinmaker.ListWorkspacesInput, opts ...request.Option) (*iottwinmaker.ListWorkspacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iottwinmaker.ListWorkspacesOutput), nil
	}
	out, err := c.IoTTwinMakerAPI.ListWorkspacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
