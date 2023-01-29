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
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Kinesis struct {
	kinesisiface.KinesisAPI
	cache *cache.Cache
}

func NewKinesis(kinesisapi kinesisiface.KinesisAPI) *Kinesis {
	return &Kinesis{
		KinesisAPI: kinesisapi,
		cache:      cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Kinesis) DescribeLimits(input *kinesis.DescribeLimitsInput) (*kinesis.DescribeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeLimitsOutput), nil
	}
	out, err := c.KinesisAPI.DescribeLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeLimitsWithContext(ctx context.Context, input *kinesis.DescribeLimitsInput, opts ...request.Option) (*kinesis.DescribeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeLimitsOutput), nil
	}
	out, err := c.KinesisAPI.DescribeLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStream(input *kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStreamConsumer(input *kinesis.DescribeStreamConsumerInput) (*kinesis.DescribeStreamConsumerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamConsumerOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStreamConsumer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStreamConsumerWithContext(ctx context.Context, input *kinesis.DescribeStreamConsumerInput, opts ...request.Option) (*kinesis.DescribeStreamConsumerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamConsumerOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStreamConsumerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStreamPages(input *kinesis.DescribeStreamInput, fn func(*kinesis.DescribeStreamOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.DescribeStreamOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.DescribeStreamOutput{}
	fnCacher := func(out *kinesis.DescribeStreamOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.DescribeStreamPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) DescribeStreamPagesWithContext(ctx context.Context, input *kinesis.DescribeStreamInput, fn func(*kinesis.DescribeStreamOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.DescribeStreamOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.DescribeStreamOutput{}
	fnCacher := func(out *kinesis.DescribeStreamOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.DescribeStreamPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) DescribeStreamSummary(input *kinesis.DescribeStreamSummaryInput) (*kinesis.DescribeStreamSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamSummaryOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStreamSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStreamSummaryWithContext(ctx context.Context, input *kinesis.DescribeStreamSummaryInput, opts ...request.Option) (*kinesis.DescribeStreamSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamSummaryOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStreamSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) DescribeStreamWithContext(ctx context.Context, input *kinesis.DescribeStreamInput, opts ...request.Option) (*kinesis.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.DescribeStreamOutput), nil
	}
	out, err := c.KinesisAPI.DescribeStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) GetRecords(input *kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.GetRecordsOutput), nil
	}
	out, err := c.KinesisAPI.GetRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) GetRecordsWithContext(ctx context.Context, input *kinesis.GetRecordsInput, opts ...request.Option) (*kinesis.GetRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.GetRecordsOutput), nil
	}
	out, err := c.KinesisAPI.GetRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) GetShardIterator(input *kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.GetShardIteratorOutput), nil
	}
	out, err := c.KinesisAPI.GetShardIterator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) GetShardIteratorWithContext(ctx context.Context, input *kinesis.GetShardIteratorInput, opts ...request.Option) (*kinesis.GetShardIteratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.GetShardIteratorOutput), nil
	}
	out, err := c.KinesisAPI.GetShardIteratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListShards(input *kinesis.ListShardsInput) (*kinesis.ListShardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListShardsOutput), nil
	}
	out, err := c.KinesisAPI.ListShards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListShardsWithContext(ctx context.Context, input *kinesis.ListShardsInput, opts ...request.Option) (*kinesis.ListShardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListShardsOutput), nil
	}
	out, err := c.KinesisAPI.ListShardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListStreamConsumers(input *kinesis.ListStreamConsumersInput) (*kinesis.ListStreamConsumersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListStreamConsumersOutput), nil
	}
	out, err := c.KinesisAPI.ListStreamConsumers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListStreamConsumersPages(input *kinesis.ListStreamConsumersInput, fn func(*kinesis.ListStreamConsumersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.ListStreamConsumersOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.ListStreamConsumersOutput{}
	fnCacher := func(out *kinesis.ListStreamConsumersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.ListStreamConsumersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) ListStreamConsumersPagesWithContext(ctx context.Context, input *kinesis.ListStreamConsumersInput, fn func(*kinesis.ListStreamConsumersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.ListStreamConsumersOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.ListStreamConsumersOutput{}
	fnCacher := func(out *kinesis.ListStreamConsumersOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.ListStreamConsumersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) ListStreamConsumersWithContext(ctx context.Context, input *kinesis.ListStreamConsumersInput, opts ...request.Option) (*kinesis.ListStreamConsumersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListStreamConsumersOutput), nil
	}
	out, err := c.KinesisAPI.ListStreamConsumersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListStreams(input *kinesis.ListStreamsInput) (*kinesis.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListStreamsOutput), nil
	}
	out, err := c.KinesisAPI.ListStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListStreamsPages(input *kinesis.ListStreamsInput, fn func(*kinesis.ListStreamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.ListStreamsOutput{}
	fnCacher := func(out *kinesis.ListStreamsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.ListStreamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) ListStreamsPagesWithContext(ctx context.Context, input *kinesis.ListStreamsInput, fn func(*kinesis.ListStreamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kinesis.ListStreamsOutput), false)
		return nil
	}
	cachable := true
	output := &kinesis.ListStreamsOutput{}
	fnCacher := func(out *kinesis.ListStreamsOutput, more bool) bool {
		ret := fn(out, more)
		if !ret && more {
			cachable = false
			return false
		}
		if err := mergo.Merge(output, out); err != nil {
			cachable = false
		}
		return true
	}
	if err := c.KinesisAPI.ListStreamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kinesis) ListStreamsWithContext(ctx context.Context, input *kinesis.ListStreamsInput, opts ...request.Option) (*kinesis.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListStreamsOutput), nil
	}
	out, err := c.KinesisAPI.ListStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListTagsForStream(input *kinesis.ListTagsForStreamInput) (*kinesis.ListTagsForStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListTagsForStreamOutput), nil
	}
	out, err := c.KinesisAPI.ListTagsForStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kinesis) ListTagsForStreamWithContext(ctx context.Context, input *kinesis.ListTagsForStreamInput, opts ...request.Option) (*kinesis.ListTagsForStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kinesis.ListTagsForStreamOutput), nil
	}
	out, err := c.KinesisAPI.ListTagsForStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
