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
	"github.com/aws/aws-sdk-go/service/iotsitewise"
	"github.com/aws/aws-sdk-go/service/iotsitewise/iotsitewiseiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTSiteWise struct {
	iotsitewiseiface.IoTSiteWiseAPI
	cache *cache.Cache
}

func NewIoTSiteWise(iotsitewiseapi iotsitewiseiface.IoTSiteWiseAPI) *IoTSiteWise {
	return &IoTSiteWise{
		IoTSiteWiseAPI: iotsitewiseapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTSiteWise) BatchAssociateProjectAssets(input *iotsitewise.BatchAssociateProjectAssetsInput) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchAssociateProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchAssociateProjectAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchAssociateProjectAssetsWithContext(ctx context.Context, input *iotsitewise.BatchAssociateProjectAssetsInput, opts ...request.Option) (*iotsitewise.BatchAssociateProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchAssociateProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchAssociateProjectAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchDisassociateProjectAssets(input *iotsitewise.BatchDisassociateProjectAssetsInput) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchDisassociateProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchDisassociateProjectAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchDisassociateProjectAssetsWithContext(ctx context.Context, input *iotsitewise.BatchDisassociateProjectAssetsInput, opts ...request.Option) (*iotsitewise.BatchDisassociateProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchDisassociateProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchDisassociateProjectAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyAggregates(input *iotsitewise.BatchGetAssetPropertyAggregatesInput) (*iotsitewise.BatchGetAssetPropertyAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyAggregatesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyAggregates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyAggregatesPages(input *iotsitewise.BatchGetAssetPropertyAggregatesInput, fn func(*iotsitewise.BatchGetAssetPropertyAggregatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyAggregatesOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyAggregatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyAggregatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyAggregatesPagesWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyAggregatesInput, fn func(*iotsitewise.BatchGetAssetPropertyAggregatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyAggregatesOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyAggregatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyAggregatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyAggregatesWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyAggregatesInput, opts ...request.Option) (*iotsitewise.BatchGetAssetPropertyAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyAggregatesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyAggregatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyValue(input *iotsitewise.BatchGetAssetPropertyValueInput) (*iotsitewise.BatchGetAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyValueHistory(input *iotsitewise.BatchGetAssetPropertyValueHistoryInput) (*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValueHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyValueHistoryPages(input *iotsitewise.BatchGetAssetPropertyValueHistoryInput, fn func(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyValueHistoryOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValueHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyValueHistoryPagesWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyValueHistoryInput, fn func(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyValueHistoryOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValueHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyValueHistoryWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyValueHistoryInput, opts ...request.Option) (*iotsitewise.BatchGetAssetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValueHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchGetAssetPropertyValuePages(input *iotsitewise.BatchGetAssetPropertyValueInput, fn func(*iotsitewise.BatchGetAssetPropertyValueOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyValueOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyValueOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValuePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyValuePagesWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyValueInput, fn func(*iotsitewise.BatchGetAssetPropertyValueOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.BatchGetAssetPropertyValueOutput{}
	fnCacher := func(out *iotsitewise.BatchGetAssetPropertyValueOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValuePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) BatchGetAssetPropertyValueWithContext(ctx context.Context, input *iotsitewise.BatchGetAssetPropertyValueInput, opts ...request.Option) (*iotsitewise.BatchGetAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchGetAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchGetAssetPropertyValueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchPutAssetPropertyValue(input *iotsitewise.BatchPutAssetPropertyValueInput) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchPutAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchPutAssetPropertyValue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) BatchPutAssetPropertyValueWithContext(ctx context.Context, input *iotsitewise.BatchPutAssetPropertyValueInput, opts ...request.Option) (*iotsitewise.BatchPutAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.BatchPutAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.BatchPutAssetPropertyValueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAccessPolicy(input *iotsitewise.DescribeAccessPolicyInput) (*iotsitewise.DescribeAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAccessPolicyOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAccessPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAccessPolicyWithContext(ctx context.Context, input *iotsitewise.DescribeAccessPolicyInput, opts ...request.Option) (*iotsitewise.DescribeAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAccessPolicyOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAccessPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAsset(input *iotsitewise.DescribeAssetInput) (*iotsitewise.DescribeAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAsset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAssetModel(input *iotsitewise.DescribeAssetModelInput) (*iotsitewise.DescribeAssetModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetModelOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAssetModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAssetModelWithContext(ctx context.Context, input *iotsitewise.DescribeAssetModelInput, opts ...request.Option) (*iotsitewise.DescribeAssetModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetModelOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAssetModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAssetProperty(input *iotsitewise.DescribeAssetPropertyInput) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetPropertyOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAssetProperty(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAssetPropertyWithContext(ctx context.Context, input *iotsitewise.DescribeAssetPropertyInput, opts ...request.Option) (*iotsitewise.DescribeAssetPropertyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetPropertyOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAssetPropertyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeAssetWithContext(ctx context.Context, input *iotsitewise.DescribeAssetInput, opts ...request.Option) (*iotsitewise.DescribeAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeAssetOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeAssetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeBulkImportJob(input *iotsitewise.DescribeBulkImportJobInput) (*iotsitewise.DescribeBulkImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeBulkImportJobOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeBulkImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeBulkImportJobWithContext(ctx context.Context, input *iotsitewise.DescribeBulkImportJobInput, opts ...request.Option) (*iotsitewise.DescribeBulkImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeBulkImportJobOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeBulkImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeDashboard(input *iotsitewise.DescribeDashboardInput) (*iotsitewise.DescribeDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeDashboardOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeDashboard(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeDashboardWithContext(ctx context.Context, input *iotsitewise.DescribeDashboardInput, opts ...request.Option) (*iotsitewise.DescribeDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeDashboardOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeDashboardWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeDefaultEncryptionConfiguration(input *iotsitewise.DescribeDefaultEncryptionConfigurationInput) (*iotsitewise.DescribeDefaultEncryptionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeDefaultEncryptionConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeDefaultEncryptionConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeDefaultEncryptionConfigurationWithContext(ctx context.Context, input *iotsitewise.DescribeDefaultEncryptionConfigurationInput, opts ...request.Option) (*iotsitewise.DescribeDefaultEncryptionConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeDefaultEncryptionConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeDefaultEncryptionConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeGateway(input *iotsitewise.DescribeGatewayInput) (*iotsitewise.DescribeGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeGatewayOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeGatewayCapabilityConfiguration(input *iotsitewise.DescribeGatewayCapabilityConfigurationInput) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeGatewayCapabilityConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeGatewayCapabilityConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeGatewayCapabilityConfigurationWithContext(ctx context.Context, input *iotsitewise.DescribeGatewayCapabilityConfigurationInput, opts ...request.Option) (*iotsitewise.DescribeGatewayCapabilityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeGatewayCapabilityConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeGatewayCapabilityConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeGatewayWithContext(ctx context.Context, input *iotsitewise.DescribeGatewayInput, opts ...request.Option) (*iotsitewise.DescribeGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeGatewayOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeLoggingOptions(input *iotsitewise.DescribeLoggingOptionsInput) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeLoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeLoggingOptionsWithContext(ctx context.Context, input *iotsitewise.DescribeLoggingOptionsInput, opts ...request.Option) (*iotsitewise.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeLoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribePortal(input *iotsitewise.DescribePortalInput) (*iotsitewise.DescribePortalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribePortalOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribePortal(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribePortalWithContext(ctx context.Context, input *iotsitewise.DescribePortalInput, opts ...request.Option) (*iotsitewise.DescribePortalOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribePortalOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribePortalWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeProject(input *iotsitewise.DescribeProjectInput) (*iotsitewise.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeProjectOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeProjectWithContext(ctx context.Context, input *iotsitewise.DescribeProjectInput, opts ...request.Option) (*iotsitewise.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeProjectOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeStorageConfiguration(input *iotsitewise.DescribeStorageConfigurationInput) (*iotsitewise.DescribeStorageConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeStorageConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeStorageConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeStorageConfigurationWithContext(ctx context.Context, input *iotsitewise.DescribeStorageConfigurationInput, opts ...request.Option) (*iotsitewise.DescribeStorageConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeStorageConfigurationOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeStorageConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeTimeSeries(input *iotsitewise.DescribeTimeSeriesInput) (*iotsitewise.DescribeTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeTimeSeriesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeTimeSeries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) DescribeTimeSeriesWithContext(ctx context.Context, input *iotsitewise.DescribeTimeSeriesInput, opts ...request.Option) (*iotsitewise.DescribeTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.DescribeTimeSeriesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.DescribeTimeSeriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyAggregates(input *iotsitewise.GetAssetPropertyAggregatesInput) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyAggregatesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyAggregates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyAggregatesPages(input *iotsitewise.GetAssetPropertyAggregatesInput, fn func(*iotsitewise.GetAssetPropertyAggregatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetAssetPropertyAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetAssetPropertyAggregatesOutput{}
	fnCacher := func(out *iotsitewise.GetAssetPropertyAggregatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetAssetPropertyAggregatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetAssetPropertyAggregatesPagesWithContext(ctx context.Context, input *iotsitewise.GetAssetPropertyAggregatesInput, fn func(*iotsitewise.GetAssetPropertyAggregatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetAssetPropertyAggregatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetAssetPropertyAggregatesOutput{}
	fnCacher := func(out *iotsitewise.GetAssetPropertyAggregatesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetAssetPropertyAggregatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetAssetPropertyAggregatesWithContext(ctx context.Context, input *iotsitewise.GetAssetPropertyAggregatesInput, opts ...request.Option) (*iotsitewise.GetAssetPropertyAggregatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyAggregatesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyAggregatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyValue(input *iotsitewise.GetAssetPropertyValueInput) (*iotsitewise.GetAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyValue(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyValueHistory(input *iotsitewise.GetAssetPropertyValueHistoryInput) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyValueHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyValueHistoryPages(input *iotsitewise.GetAssetPropertyValueHistoryInput, fn func(*iotsitewise.GetAssetPropertyValueHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetAssetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetAssetPropertyValueHistoryOutput{}
	fnCacher := func(out *iotsitewise.GetAssetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetAssetPropertyValueHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetAssetPropertyValueHistoryPagesWithContext(ctx context.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput, fn func(*iotsitewise.GetAssetPropertyValueHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetAssetPropertyValueHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetAssetPropertyValueHistoryOutput{}
	fnCacher := func(out *iotsitewise.GetAssetPropertyValueHistoryOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetAssetPropertyValueHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetAssetPropertyValueHistoryWithContext(ctx context.Context, input *iotsitewise.GetAssetPropertyValueHistoryInput, opts ...request.Option) (*iotsitewise.GetAssetPropertyValueHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyValueHistoryOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyValueHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetAssetPropertyValueWithContext(ctx context.Context, input *iotsitewise.GetAssetPropertyValueInput, opts ...request.Option) (*iotsitewise.GetAssetPropertyValueOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetAssetPropertyValueOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetAssetPropertyValueWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetInterpolatedAssetPropertyValues(input *iotsitewise.GetInterpolatedAssetPropertyValuesInput) (*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetInterpolatedAssetPropertyValues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) GetInterpolatedAssetPropertyValuesPages(input *iotsitewise.GetInterpolatedAssetPropertyValuesInput, fn func(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetInterpolatedAssetPropertyValuesOutput{}
	fnCacher := func(out *iotsitewise.GetInterpolatedAssetPropertyValuesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetInterpolatedAssetPropertyValuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetInterpolatedAssetPropertyValuesPagesWithContext(ctx context.Context, input *iotsitewise.GetInterpolatedAssetPropertyValuesInput, fn func(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.GetInterpolatedAssetPropertyValuesOutput{}
	fnCacher := func(out *iotsitewise.GetInterpolatedAssetPropertyValuesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.GetInterpolatedAssetPropertyValuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) GetInterpolatedAssetPropertyValuesWithContext(ctx context.Context, input *iotsitewise.GetInterpolatedAssetPropertyValuesInput, opts ...request.Option) (*iotsitewise.GetInterpolatedAssetPropertyValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.GetInterpolatedAssetPropertyValuesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.GetInterpolatedAssetPropertyValuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAccessPolicies(input *iotsitewise.ListAccessPoliciesInput) (*iotsitewise.ListAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAccessPoliciesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAccessPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAccessPoliciesPages(input *iotsitewise.ListAccessPoliciesInput, fn func(*iotsitewise.ListAccessPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAccessPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAccessPoliciesOutput{}
	fnCacher := func(out *iotsitewise.ListAccessPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAccessPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAccessPoliciesPagesWithContext(ctx context.Context, input *iotsitewise.ListAccessPoliciesInput, fn func(*iotsitewise.ListAccessPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAccessPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAccessPoliciesOutput{}
	fnCacher := func(out *iotsitewise.ListAccessPoliciesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAccessPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAccessPoliciesWithContext(ctx context.Context, input *iotsitewise.ListAccessPoliciesInput, opts ...request.Option) (*iotsitewise.ListAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAccessPoliciesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAccessPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetModelProperties(input *iotsitewise.ListAssetModelPropertiesInput) (*iotsitewise.ListAssetModelPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetModelPropertiesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetModelProperties(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetModelPropertiesPages(input *iotsitewise.ListAssetModelPropertiesInput, fn func(*iotsitewise.ListAssetModelPropertiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetModelPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetModelPropertiesOutput{}
	fnCacher := func(out *iotsitewise.ListAssetModelPropertiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetModelPropertiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetModelPropertiesPagesWithContext(ctx context.Context, input *iotsitewise.ListAssetModelPropertiesInput, fn func(*iotsitewise.ListAssetModelPropertiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetModelPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetModelPropertiesOutput{}
	fnCacher := func(out *iotsitewise.ListAssetModelPropertiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetModelPropertiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetModelPropertiesWithContext(ctx context.Context, input *iotsitewise.ListAssetModelPropertiesInput, opts ...request.Option) (*iotsitewise.ListAssetModelPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetModelPropertiesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetModelPropertiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetModels(input *iotsitewise.ListAssetModelsInput) (*iotsitewise.ListAssetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetModelsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetModelsPages(input *iotsitewise.ListAssetModelsInput, fn func(*iotsitewise.ListAssetModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetModelsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetModelsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetModelsPagesWithContext(ctx context.Context, input *iotsitewise.ListAssetModelsInput, fn func(*iotsitewise.ListAssetModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetModelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetModelsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetModelsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetModelsWithContext(ctx context.Context, input *iotsitewise.ListAssetModelsInput, opts ...request.Option) (*iotsitewise.ListAssetModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetModelsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetProperties(input *iotsitewise.ListAssetPropertiesInput) (*iotsitewise.ListAssetPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetPropertiesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetProperties(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetPropertiesPages(input *iotsitewise.ListAssetPropertiesInput, fn func(*iotsitewise.ListAssetPropertiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetPropertiesOutput{}
	fnCacher := func(out *iotsitewise.ListAssetPropertiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetPropertiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetPropertiesPagesWithContext(ctx context.Context, input *iotsitewise.ListAssetPropertiesInput, fn func(*iotsitewise.ListAssetPropertiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetPropertiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetPropertiesOutput{}
	fnCacher := func(out *iotsitewise.ListAssetPropertiesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetPropertiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetPropertiesWithContext(ctx context.Context, input *iotsitewise.ListAssetPropertiesInput, opts ...request.Option) (*iotsitewise.ListAssetPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetPropertiesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetPropertiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetRelationships(input *iotsitewise.ListAssetRelationshipsInput) (*iotsitewise.ListAssetRelationshipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetRelationshipsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetRelationships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetRelationshipsPages(input *iotsitewise.ListAssetRelationshipsInput, fn func(*iotsitewise.ListAssetRelationshipsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetRelationshipsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetRelationshipsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetRelationshipsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetRelationshipsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetRelationshipsPagesWithContext(ctx context.Context, input *iotsitewise.ListAssetRelationshipsInput, fn func(*iotsitewise.ListAssetRelationshipsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetRelationshipsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetRelationshipsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetRelationshipsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetRelationshipsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetRelationshipsWithContext(ctx context.Context, input *iotsitewise.ListAssetRelationshipsInput, opts ...request.Option) (*iotsitewise.ListAssetRelationshipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetRelationshipsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetRelationshipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssets(input *iotsitewise.ListAssetsInput) (*iotsitewise.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssetsPages(input *iotsitewise.ListAssetsInput, fn func(*iotsitewise.ListAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetsPagesWithContext(ctx context.Context, input *iotsitewise.ListAssetsInput, fn func(*iotsitewise.ListAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssetsWithContext(ctx context.Context, input *iotsitewise.ListAssetsInput, opts ...request.Option) (*iotsitewise.ListAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssociatedAssets(input *iotsitewise.ListAssociatedAssetsInput) (*iotsitewise.ListAssociatedAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssociatedAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssociatedAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListAssociatedAssetsPages(input *iotsitewise.ListAssociatedAssetsInput, fn func(*iotsitewise.ListAssociatedAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssociatedAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssociatedAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListAssociatedAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssociatedAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssociatedAssetsPagesWithContext(ctx context.Context, input *iotsitewise.ListAssociatedAssetsInput, fn func(*iotsitewise.ListAssociatedAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListAssociatedAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListAssociatedAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListAssociatedAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListAssociatedAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListAssociatedAssetsWithContext(ctx context.Context, input *iotsitewise.ListAssociatedAssetsInput, opts ...request.Option) (*iotsitewise.ListAssociatedAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListAssociatedAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListAssociatedAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListBulkImportJobs(input *iotsitewise.ListBulkImportJobsInput) (*iotsitewise.ListBulkImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListBulkImportJobsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListBulkImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListBulkImportJobsPages(input *iotsitewise.ListBulkImportJobsInput, fn func(*iotsitewise.ListBulkImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListBulkImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListBulkImportJobsOutput{}
	fnCacher := func(out *iotsitewise.ListBulkImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListBulkImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListBulkImportJobsPagesWithContext(ctx context.Context, input *iotsitewise.ListBulkImportJobsInput, fn func(*iotsitewise.ListBulkImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListBulkImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListBulkImportJobsOutput{}
	fnCacher := func(out *iotsitewise.ListBulkImportJobsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListBulkImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListBulkImportJobsWithContext(ctx context.Context, input *iotsitewise.ListBulkImportJobsInput, opts ...request.Option) (*iotsitewise.ListBulkImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListBulkImportJobsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListBulkImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListDashboards(input *iotsitewise.ListDashboardsInput) (*iotsitewise.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListDashboardsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListDashboards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListDashboardsPages(input *iotsitewise.ListDashboardsInput, fn func(*iotsitewise.ListDashboardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListDashboardsOutput{}
	fnCacher := func(out *iotsitewise.ListDashboardsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListDashboardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListDashboardsPagesWithContext(ctx context.Context, input *iotsitewise.ListDashboardsInput, fn func(*iotsitewise.ListDashboardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListDashboardsOutput{}
	fnCacher := func(out *iotsitewise.ListDashboardsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListDashboardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListDashboardsWithContext(ctx context.Context, input *iotsitewise.ListDashboardsInput, opts ...request.Option) (*iotsitewise.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListDashboardsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListDashboardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListGateways(input *iotsitewise.ListGatewaysInput) (*iotsitewise.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListGatewaysOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListGatewaysPages(input *iotsitewise.ListGatewaysInput, fn func(*iotsitewise.ListGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListGatewaysOutput{}
	fnCacher := func(out *iotsitewise.ListGatewaysOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListGatewaysPagesWithContext(ctx context.Context, input *iotsitewise.ListGatewaysInput, fn func(*iotsitewise.ListGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListGatewaysOutput{}
	fnCacher := func(out *iotsitewise.ListGatewaysOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListGatewaysWithContext(ctx context.Context, input *iotsitewise.ListGatewaysInput, opts ...request.Option) (*iotsitewise.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListGatewaysOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListPortals(input *iotsitewise.ListPortalsInput) (*iotsitewise.ListPortalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListPortalsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListPortals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListPortalsPages(input *iotsitewise.ListPortalsInput, fn func(*iotsitewise.ListPortalsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListPortalsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListPortalsOutput{}
	fnCacher := func(out *iotsitewise.ListPortalsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListPortalsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListPortalsPagesWithContext(ctx context.Context, input *iotsitewise.ListPortalsInput, fn func(*iotsitewise.ListPortalsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListPortalsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListPortalsOutput{}
	fnCacher := func(out *iotsitewise.ListPortalsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListPortalsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListPortalsWithContext(ctx context.Context, input *iotsitewise.ListPortalsInput, opts ...request.Option) (*iotsitewise.ListPortalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListPortalsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListPortalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListProjectAssets(input *iotsitewise.ListProjectAssetsInput) (*iotsitewise.ListProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListProjectAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListProjectAssetsPages(input *iotsitewise.ListProjectAssetsInput, fn func(*iotsitewise.ListProjectAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListProjectAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListProjectAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListProjectAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListProjectAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListProjectAssetsPagesWithContext(ctx context.Context, input *iotsitewise.ListProjectAssetsInput, fn func(*iotsitewise.ListProjectAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListProjectAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListProjectAssetsOutput{}
	fnCacher := func(out *iotsitewise.ListProjectAssetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListProjectAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListProjectAssetsWithContext(ctx context.Context, input *iotsitewise.ListProjectAssetsInput, opts ...request.Option) (*iotsitewise.ListProjectAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListProjectAssetsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListProjectAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListProjects(input *iotsitewise.ListProjectsInput) (*iotsitewise.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListProjectsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListProjectsPages(input *iotsitewise.ListProjectsInput, fn func(*iotsitewise.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListProjectsOutput{}
	fnCacher := func(out *iotsitewise.ListProjectsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListProjectsPagesWithContext(ctx context.Context, input *iotsitewise.ListProjectsInput, fn func(*iotsitewise.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListProjectsOutput{}
	fnCacher := func(out *iotsitewise.ListProjectsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListProjectsWithContext(ctx context.Context, input *iotsitewise.ListProjectsInput, opts ...request.Option) (*iotsitewise.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListProjectsOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListTagsForResource(input *iotsitewise.ListTagsForResourceInput) (*iotsitewise.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListTagsForResourceWithContext(ctx context.Context, input *iotsitewise.ListTagsForResourceInput, opts ...request.Option) (*iotsitewise.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListTimeSeries(input *iotsitewise.ListTimeSeriesInput) (*iotsitewise.ListTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListTimeSeriesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListTimeSeries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTSiteWise) ListTimeSeriesPages(input *iotsitewise.ListTimeSeriesInput, fn func(*iotsitewise.ListTimeSeriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListTimeSeriesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListTimeSeriesOutput{}
	fnCacher := func(out *iotsitewise.ListTimeSeriesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListTimeSeriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListTimeSeriesPagesWithContext(ctx context.Context, input *iotsitewise.ListTimeSeriesInput, fn func(*iotsitewise.ListTimeSeriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotsitewise.ListTimeSeriesOutput), false)
		return nil
	}
	cachable := true
	output := &iotsitewise.ListTimeSeriesOutput{}
	fnCacher := func(out *iotsitewise.ListTimeSeriesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTSiteWiseAPI.ListTimeSeriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTSiteWise) ListTimeSeriesWithContext(ctx context.Context, input *iotsitewise.ListTimeSeriesInput, opts ...request.Option) (*iotsitewise.ListTimeSeriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotsitewise.ListTimeSeriesOutput), nil
	}
	out, err := c.IoTSiteWiseAPI.ListTimeSeriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
