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
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/firehose/firehoseiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Firehose struct {
	firehoseiface.FirehoseAPI
	cache *cache.Cache
}

func NewFirehose(firehoseapi firehoseiface.FirehoseAPI) *Firehose {
	return &Firehose{
		FirehoseAPI: firehoseapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Firehose) DescribeDeliveryStream(input *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.DescribeDeliveryStreamOutput), nil
	}
	out, err := c.FirehoseAPI.DescribeDeliveryStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Firehose) DescribeDeliveryStreamWithContext(ctx context.Context, input *firehose.DescribeDeliveryStreamInput, opts ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.DescribeDeliveryStreamOutput), nil
	}
	out, err := c.FirehoseAPI.DescribeDeliveryStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Firehose) ListDeliveryStreams(input *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.ListDeliveryStreamsOutput), nil
	}
	out, err := c.FirehoseAPI.ListDeliveryStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Firehose) ListDeliveryStreamsWithContext(ctx context.Context, input *firehose.ListDeliveryStreamsInput, opts ...request.Option) (*firehose.ListDeliveryStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.ListDeliveryStreamsOutput), nil
	}
	out, err := c.FirehoseAPI.ListDeliveryStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Firehose) ListTagsForDeliveryStream(input *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.ListTagsForDeliveryStreamOutput), nil
	}
	out, err := c.FirehoseAPI.ListTagsForDeliveryStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Firehose) ListTagsForDeliveryStreamWithContext(ctx context.Context, input *firehose.ListTagsForDeliveryStreamInput, opts ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*firehose.ListTagsForDeliveryStreamOutput), nil
	}
	out, err := c.FirehoseAPI.ListTagsForDeliveryStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
