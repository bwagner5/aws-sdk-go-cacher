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
package appflowcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/appflow"
	"github.com/aws/aws-sdk-go/service/appflow/appflowiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Appflow struct {
	appflowiface.AppflowAPI
	cache *cache.Cache
}

func New(appflowapi appflowiface.AppflowAPI) *Appflow {
	return &Appflow{
		AppflowAPI: appflowapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Appflow) DescribeConnector(input *appflow.DescribeConnectorInput) (*appflow.DescribeConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorEntity(input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorEntityOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorEntity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorEntityWithContext(ctx context.Context, input *appflow.DescribeConnectorEntityInput, opts ...request.Option) (*appflow.DescribeConnectorEntityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorEntityOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorEntityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorProfiles(input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorProfilesOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorProfilesPages(input *appflow.DescribeConnectorProfilesInput, fn func(*appflow.DescribeConnectorProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeConnectorProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeConnectorProfilesOutput{}
	fnCacher := func(out *appflow.DescribeConnectorProfilesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeConnectorProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeConnectorProfilesPagesWithContext(ctx context.Context, input *appflow.DescribeConnectorProfilesInput, fn func(*appflow.DescribeConnectorProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeConnectorProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeConnectorProfilesOutput{}
	fnCacher := func(out *appflow.DescribeConnectorProfilesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeConnectorProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeConnectorProfilesWithContext(ctx context.Context, input *appflow.DescribeConnectorProfilesInput, opts ...request.Option) (*appflow.DescribeConnectorProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorProfilesOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorWithContext(ctx context.Context, input *appflow.DescribeConnectorInput, opts ...request.Option) (*appflow.DescribeConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectors(input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorsOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeConnectorsPages(input *appflow.DescribeConnectorsInput, fn func(*appflow.DescribeConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeConnectorsOutput{}
	fnCacher := func(out *appflow.DescribeConnectorsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeConnectorsPagesWithContext(ctx context.Context, input *appflow.DescribeConnectorsInput, fn func(*appflow.DescribeConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeConnectorsOutput{}
	fnCacher := func(out *appflow.DescribeConnectorsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeConnectorsWithContext(ctx context.Context, input *appflow.DescribeConnectorsInput, opts ...request.Option) (*appflow.DescribeConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeConnectorsOutput), nil
	}
	out, err := c.AppflowAPI.DescribeConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeFlow(input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeFlowOutput), nil
	}
	out, err := c.AppflowAPI.DescribeFlow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeFlowExecutionRecords(input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeFlowExecutionRecordsOutput), nil
	}
	out, err := c.AppflowAPI.DescribeFlowExecutionRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeFlowExecutionRecordsPages(input *appflow.DescribeFlowExecutionRecordsInput, fn func(*appflow.DescribeFlowExecutionRecordsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeFlowExecutionRecordsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeFlowExecutionRecordsOutput{}
	fnCacher := func(out *appflow.DescribeFlowExecutionRecordsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeFlowExecutionRecordsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeFlowExecutionRecordsPagesWithContext(ctx context.Context, input *appflow.DescribeFlowExecutionRecordsInput, fn func(*appflow.DescribeFlowExecutionRecordsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.DescribeFlowExecutionRecordsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.DescribeFlowExecutionRecordsOutput{}
	fnCacher := func(out *appflow.DescribeFlowExecutionRecordsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.DescribeFlowExecutionRecordsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) DescribeFlowExecutionRecordsWithContext(ctx context.Context, input *appflow.DescribeFlowExecutionRecordsInput, opts ...request.Option) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeFlowExecutionRecordsOutput), nil
	}
	out, err := c.AppflowAPI.DescribeFlowExecutionRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) DescribeFlowWithContext(ctx context.Context, input *appflow.DescribeFlowInput, opts ...request.Option) (*appflow.DescribeFlowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.DescribeFlowOutput), nil
	}
	out, err := c.AppflowAPI.DescribeFlowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListConnectorEntities(input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListConnectorEntitiesOutput), nil
	}
	out, err := c.AppflowAPI.ListConnectorEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListConnectorEntitiesWithContext(ctx context.Context, input *appflow.ListConnectorEntitiesInput, opts ...request.Option) (*appflow.ListConnectorEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListConnectorEntitiesOutput), nil
	}
	out, err := c.AppflowAPI.ListConnectorEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListConnectors(input *appflow.ListConnectorsInput) (*appflow.ListConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListConnectorsOutput), nil
	}
	out, err := c.AppflowAPI.ListConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListConnectorsPages(input *appflow.ListConnectorsInput, fn func(*appflow.ListConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.ListConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.ListConnectorsOutput{}
	fnCacher := func(out *appflow.ListConnectorsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.ListConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) ListConnectorsPagesWithContext(ctx context.Context, input *appflow.ListConnectorsInput, fn func(*appflow.ListConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.ListConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.ListConnectorsOutput{}
	fnCacher := func(out *appflow.ListConnectorsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.ListConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) ListConnectorsWithContext(ctx context.Context, input *appflow.ListConnectorsInput, opts ...request.Option) (*appflow.ListConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListConnectorsOutput), nil
	}
	out, err := c.AppflowAPI.ListConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListFlows(input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListFlowsOutput), nil
	}
	out, err := c.AppflowAPI.ListFlows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListFlowsPages(input *appflow.ListFlowsInput, fn func(*appflow.ListFlowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.ListFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.ListFlowsOutput{}
	fnCacher := func(out *appflow.ListFlowsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.ListFlowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) ListFlowsPagesWithContext(ctx context.Context, input *appflow.ListFlowsInput, fn func(*appflow.ListFlowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*appflow.ListFlowsOutput), false)
		return nil
	}
	cachable := true
	output := &appflow.ListFlowsOutput{}
	fnCacher := func(out *appflow.ListFlowsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.AppflowAPI.ListFlowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Appflow) ListFlowsWithContext(ctx context.Context, input *appflow.ListFlowsInput, opts ...request.Option) (*appflow.ListFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListFlowsOutput), nil
	}
	out, err := c.AppflowAPI.ListFlowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListTagsForResource(input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListTagsForResourceOutput), nil
	}
	out, err := c.AppflowAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Appflow) ListTagsForResourceWithContext(ctx context.Context, input *appflow.ListTagsForResourceInput, opts ...request.Option) (*appflow.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*appflow.ListTagsForResourceOutput), nil
	}
	out, err := c.AppflowAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
