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
package dataexchangecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dataexchange"
	"github.com/aws/aws-sdk-go/service/dataexchange/dataexchangeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DataExchange struct {
	dataexchangeiface.DataExchangeAPI
	cache *cache.Cache
}

func New(dataexchangeapi dataexchangeiface.DataExchangeAPI) *DataExchange {
	return &DataExchange{
		DataExchangeAPI: dataexchangeapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DataExchange) GetAsset(input *dataexchange.GetAssetInput) (*dataexchange.GetAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetAssetOutput), nil
	}
	out, err := c.DataExchangeAPI.GetAsset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetAssetWithContext(ctx context.Context, input *dataexchange.GetAssetInput, opts ...request.Option) (*dataexchange.GetAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetAssetOutput), nil
	}
	out, err := c.DataExchangeAPI.GetAssetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetDataSet(input *dataexchange.GetDataSetInput) (*dataexchange.GetDataSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetDataSetOutput), nil
	}
	out, err := c.DataExchangeAPI.GetDataSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetDataSetWithContext(ctx context.Context, input *dataexchange.GetDataSetInput, opts ...request.Option) (*dataexchange.GetDataSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetDataSetOutput), nil
	}
	out, err := c.DataExchangeAPI.GetDataSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetEventAction(input *dataexchange.GetEventActionInput) (*dataexchange.GetEventActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetEventActionOutput), nil
	}
	out, err := c.DataExchangeAPI.GetEventAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetEventActionWithContext(ctx context.Context, input *dataexchange.GetEventActionInput, opts ...request.Option) (*dataexchange.GetEventActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetEventActionOutput), nil
	}
	out, err := c.DataExchangeAPI.GetEventActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetJob(input *dataexchange.GetJobInput) (*dataexchange.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetJobOutput), nil
	}
	out, err := c.DataExchangeAPI.GetJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetJobWithContext(ctx context.Context, input *dataexchange.GetJobInput, opts ...request.Option) (*dataexchange.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetJobOutput), nil
	}
	out, err := c.DataExchangeAPI.GetJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetRevision(input *dataexchange.GetRevisionInput) (*dataexchange.GetRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetRevisionOutput), nil
	}
	out, err := c.DataExchangeAPI.GetRevision(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) GetRevisionWithContext(ctx context.Context, input *dataexchange.GetRevisionInput, opts ...request.Option) (*dataexchange.GetRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.GetRevisionOutput), nil
	}
	out, err := c.DataExchangeAPI.GetRevisionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListDataSetRevisions(input *dataexchange.ListDataSetRevisionsInput) (*dataexchange.ListDataSetRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListDataSetRevisionsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListDataSetRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListDataSetRevisionsPages(input *dataexchange.ListDataSetRevisionsInput, fn func(*dataexchange.ListDataSetRevisionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListDataSetRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListDataSetRevisionsOutput{}
	fnCacher := func(out *dataexchange.ListDataSetRevisionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListDataSetRevisionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListDataSetRevisionsPagesWithContext(ctx context.Context, input *dataexchange.ListDataSetRevisionsInput, fn func(*dataexchange.ListDataSetRevisionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListDataSetRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListDataSetRevisionsOutput{}
	fnCacher := func(out *dataexchange.ListDataSetRevisionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListDataSetRevisionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListDataSetRevisionsWithContext(ctx context.Context, input *dataexchange.ListDataSetRevisionsInput, opts ...request.Option) (*dataexchange.ListDataSetRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListDataSetRevisionsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListDataSetRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListDataSets(input *dataexchange.ListDataSetsInput) (*dataexchange.ListDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListDataSetsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListDataSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListDataSetsPages(input *dataexchange.ListDataSetsInput, fn func(*dataexchange.ListDataSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListDataSetsOutput{}
	fnCacher := func(out *dataexchange.ListDataSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListDataSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListDataSetsPagesWithContext(ctx context.Context, input *dataexchange.ListDataSetsInput, fn func(*dataexchange.ListDataSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListDataSetsOutput{}
	fnCacher := func(out *dataexchange.ListDataSetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListDataSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListDataSetsWithContext(ctx context.Context, input *dataexchange.ListDataSetsInput, opts ...request.Option) (*dataexchange.ListDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListDataSetsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListDataSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListEventActions(input *dataexchange.ListEventActionsInput) (*dataexchange.ListEventActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListEventActionsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListEventActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListEventActionsPages(input *dataexchange.ListEventActionsInput, fn func(*dataexchange.ListEventActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListEventActionsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListEventActionsOutput{}
	fnCacher := func(out *dataexchange.ListEventActionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListEventActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListEventActionsPagesWithContext(ctx context.Context, input *dataexchange.ListEventActionsInput, fn func(*dataexchange.ListEventActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListEventActionsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListEventActionsOutput{}
	fnCacher := func(out *dataexchange.ListEventActionsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListEventActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListEventActionsWithContext(ctx context.Context, input *dataexchange.ListEventActionsInput, opts ...request.Option) (*dataexchange.ListEventActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListEventActionsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListEventActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListJobs(input *dataexchange.ListJobsInput) (*dataexchange.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListJobsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListJobsPages(input *dataexchange.ListJobsInput, fn func(*dataexchange.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListJobsOutput{}
	fnCacher := func(out *dataexchange.ListJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListJobsPagesWithContext(ctx context.Context, input *dataexchange.ListJobsInput, fn func(*dataexchange.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListJobsOutput{}
	fnCacher := func(out *dataexchange.ListJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListJobsWithContext(ctx context.Context, input *dataexchange.ListJobsInput, opts ...request.Option) (*dataexchange.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListJobsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListRevisionAssets(input *dataexchange.ListRevisionAssetsInput) (*dataexchange.ListRevisionAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListRevisionAssetsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListRevisionAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListRevisionAssetsPages(input *dataexchange.ListRevisionAssetsInput, fn func(*dataexchange.ListRevisionAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListRevisionAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListRevisionAssetsOutput{}
	fnCacher := func(out *dataexchange.ListRevisionAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListRevisionAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListRevisionAssetsPagesWithContext(ctx context.Context, input *dataexchange.ListRevisionAssetsInput, fn func(*dataexchange.ListRevisionAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dataexchange.ListRevisionAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &dataexchange.ListRevisionAssetsOutput{}
	fnCacher := func(out *dataexchange.ListRevisionAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.DataExchangeAPI.ListRevisionAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataExchange) ListRevisionAssetsWithContext(ctx context.Context, input *dataexchange.ListRevisionAssetsInput, opts ...request.Option) (*dataexchange.ListRevisionAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListRevisionAssetsOutput), nil
	}
	out, err := c.DataExchangeAPI.ListRevisionAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListTagsForResource(input *dataexchange.ListTagsForResourceInput) (*dataexchange.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListTagsForResourceOutput), nil
	}
	out, err := c.DataExchangeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataExchange) ListTagsForResourceWithContext(ctx context.Context, input *dataexchange.ListTagsForResourceInput, opts ...request.Option) (*dataexchange.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dataexchange.ListTagsForResourceOutput), nil
	}
	out, err := c.DataExchangeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
