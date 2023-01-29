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
	"github.com/aws/aws-sdk-go/service/iotanalytics"
	"github.com/aws/aws-sdk-go/service/iotanalytics/iotanalyticsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTAnalytics struct {
	iotanalyticsiface.IoTAnalyticsAPI
	cache *cache.Cache
}

func NewIoTAnalytics(iotanalyticsapi iotanalyticsiface.IoTAnalyticsAPI) *IoTAnalytics {
	return &IoTAnalytics{
		IoTAnalyticsAPI: iotanalyticsapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTAnalytics) BatchPutMessage(input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.BatchPutMessageOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.BatchPutMessage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) BatchPutMessageWithContext(ctx context.Context, input *iotanalytics.BatchPutMessageInput, opts ...request.Option) (*iotanalytics.BatchPutMessageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.BatchPutMessageOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.BatchPutMessageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeChannel(input *iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeChannelOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeChannelWithContext(ctx context.Context, input *iotanalytics.DescribeChannelInput, opts ...request.Option) (*iotanalytics.DescribeChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeChannelOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeDataset(input *iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeDatasetOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeDatasetWithContext(ctx context.Context, input *iotanalytics.DescribeDatasetInput, opts ...request.Option) (*iotanalytics.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeDatasetOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeDatastore(input *iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeDatastoreOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeDatastore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeDatastoreWithContext(ctx context.Context, input *iotanalytics.DescribeDatastoreInput, opts ...request.Option) (*iotanalytics.DescribeDatastoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeDatastoreOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeDatastoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeLoggingOptions(input *iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeLoggingOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribeLoggingOptionsWithContext(ctx context.Context, input *iotanalytics.DescribeLoggingOptionsInput, opts ...request.Option) (*iotanalytics.DescribeLoggingOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribeLoggingOptionsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribeLoggingOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribePipeline(input *iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribePipelineOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribePipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) DescribePipelineWithContext(ctx context.Context, input *iotanalytics.DescribePipelineInput, opts ...request.Option) (*iotanalytics.DescribePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.DescribePipelineOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.DescribePipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) GetDatasetContent(input *iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.GetDatasetContentOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.GetDatasetContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) GetDatasetContentWithContext(ctx context.Context, input *iotanalytics.GetDatasetContentInput, opts ...request.Option) (*iotanalytics.GetDatasetContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.GetDatasetContentOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.GetDatasetContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListChannels(input *iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListChannelsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListChannelsPages(input *iotanalytics.ListChannelsInput, fn func(*iotanalytics.ListChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListChannelsOutput{}
	fnCacher := func(out *iotanalytics.ListChannelsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListChannelsPagesWithContext(ctx context.Context, input *iotanalytics.ListChannelsInput, fn func(*iotanalytics.ListChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListChannelsOutput{}
	fnCacher := func(out *iotanalytics.ListChannelsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListChannelsWithContext(ctx context.Context, input *iotanalytics.ListChannelsInput, opts ...request.Option) (*iotanalytics.ListChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListChannelsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatasetContents(input *iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatasetContentsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatasetContents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatasetContentsPages(input *iotanalytics.ListDatasetContentsInput, fn func(*iotanalytics.ListDatasetContentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatasetContentsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatasetContentsOutput{}
	fnCacher := func(out *iotanalytics.ListDatasetContentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatasetContentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatasetContentsPagesWithContext(ctx context.Context, input *iotanalytics.ListDatasetContentsInput, fn func(*iotanalytics.ListDatasetContentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatasetContentsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatasetContentsOutput{}
	fnCacher := func(out *iotanalytics.ListDatasetContentsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatasetContentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatasetContentsWithContext(ctx context.Context, input *iotanalytics.ListDatasetContentsInput, opts ...request.Option) (*iotanalytics.ListDatasetContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatasetContentsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatasetContentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatasets(input *iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatasetsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatasetsPages(input *iotanalytics.ListDatasetsInput, fn func(*iotanalytics.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatasetsOutput{}
	fnCacher := func(out *iotanalytics.ListDatasetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatasetsPagesWithContext(ctx context.Context, input *iotanalytics.ListDatasetsInput, fn func(*iotanalytics.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatasetsOutput{}
	fnCacher := func(out *iotanalytics.ListDatasetsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatasetsWithContext(ctx context.Context, input *iotanalytics.ListDatasetsInput, opts ...request.Option) (*iotanalytics.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatasetsOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatastores(input *iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatastoresOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatastores(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListDatastoresPages(input *iotanalytics.ListDatastoresInput, fn func(*iotanalytics.ListDatastoresOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatastoresOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatastoresOutput{}
	fnCacher := func(out *iotanalytics.ListDatastoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatastoresPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatastoresPagesWithContext(ctx context.Context, input *iotanalytics.ListDatastoresInput, fn func(*iotanalytics.ListDatastoresOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListDatastoresOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListDatastoresOutput{}
	fnCacher := func(out *iotanalytics.ListDatastoresOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListDatastoresPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListDatastoresWithContext(ctx context.Context, input *iotanalytics.ListDatastoresInput, opts ...request.Option) (*iotanalytics.ListDatastoresOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListDatastoresOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListDatastoresWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListPipelines(input *iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListPipelinesOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListPipelinesPages(input *iotanalytics.ListPipelinesInput, fn func(*iotanalytics.ListPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListPipelinesOutput{}
	fnCacher := func(out *iotanalytics.ListPipelinesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListPipelinesPagesWithContext(ctx context.Context, input *iotanalytics.ListPipelinesInput, fn func(*iotanalytics.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotanalytics.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &iotanalytics.ListPipelinesOutput{}
	fnCacher := func(out *iotanalytics.ListPipelinesOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.IoTAnalyticsAPI.ListPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTAnalytics) ListPipelinesWithContext(ctx context.Context, input *iotanalytics.ListPipelinesInput, opts ...request.Option) (*iotanalytics.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListPipelinesOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListTagsForResource(input *iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTAnalytics) ListTagsForResourceWithContext(ctx context.Context, input *iotanalytics.ListTagsForResourceInput, opts ...request.Option) (*iotanalytics.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotanalytics.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTAnalyticsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
