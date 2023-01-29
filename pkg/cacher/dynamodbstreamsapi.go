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
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DynamoDBStreams struct {
	dynamodbstreamsiface.DynamoDBStreamsAPI
	cache *cache.Cache
}

func NewDynamoDBStreams(dynamodbstreamsapi dynamodbstreamsiface.DynamoDBStreamsAPI) *DynamoDBStreams {
	return &DynamoDBStreams{
		DynamoDBStreamsAPI: dynamodbstreamsapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DynamoDBStreams) DescribeStream(input *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.DescribeStreamOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.DescribeStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) DescribeStreamWithContext(ctx context.Context, input *dynamodbstreams.DescribeStreamInput, opts ...request.Option) (*dynamodbstreams.DescribeStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.DescribeStreamOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.DescribeStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) GetRecords(input *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.GetRecordsOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.GetRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) GetRecordsWithContext(ctx context.Context, input *dynamodbstreams.GetRecordsInput, opts ...request.Option) (*dynamodbstreams.GetRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.GetRecordsOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.GetRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) GetShardIterator(input *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.GetShardIteratorOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.GetShardIterator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) GetShardIteratorWithContext(ctx context.Context, input *dynamodbstreams.GetShardIteratorInput, opts ...request.Option) (*dynamodbstreams.GetShardIteratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.GetShardIteratorOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.GetShardIteratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) ListStreams(input *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.ListStreamsOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.ListStreams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDBStreams) ListStreamsWithContext(ctx context.Context, input *dynamodbstreams.ListStreamsInput, opts ...request.Option) (*dynamodbstreams.ListStreamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodbstreams.ListStreamsOutput), nil
	}
	out, err := c.DynamoDBStreamsAPI.ListStreamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
