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
package sqscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SQS struct {
	sqsiface.SQSAPI
	cache *cache.Cache
}

func New(sqsapi sqsiface.SQSAPI) *SQS {
	return &SQS{
		SQSAPI: sqsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SQS) GetQueueAttributes(input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.GetQueueAttributesOutput), nil
	}
	out, err := c.SQSAPI.GetQueueAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) GetQueueAttributesWithContext(ctx context.Context, input *sqs.GetQueueAttributesInput, opts ...request.Option) (*sqs.GetQueueAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.GetQueueAttributesOutput), nil
	}
	out, err := c.SQSAPI.GetQueueAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) GetQueueUrl(input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.GetQueueUrlOutput), nil
	}
	out, err := c.SQSAPI.GetQueueUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) GetQueueUrlWithContext(ctx context.Context, input *sqs.GetQueueUrlInput, opts ...request.Option) (*sqs.GetQueueUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.GetQueueUrlOutput), nil
	}
	out, err := c.SQSAPI.GetQueueUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListDeadLetterSourceQueues(input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListDeadLetterSourceQueuesOutput), nil
	}
	out, err := c.SQSAPI.ListDeadLetterSourceQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListDeadLetterSourceQueuesPages(input *sqs.ListDeadLetterSourceQueuesInput, fn func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sqs.ListDeadLetterSourceQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &sqs.ListDeadLetterSourceQueuesOutput{}
	fnCacher := func(out *sqs.ListDeadLetterSourceQueuesOutput, more bool) bool {
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
	if err := c.SQSAPI.ListDeadLetterSourceQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SQS) ListDeadLetterSourceQueuesPagesWithContext(ctx context.Context, input *sqs.ListDeadLetterSourceQueuesInput, fn func(*sqs.ListDeadLetterSourceQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sqs.ListDeadLetterSourceQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &sqs.ListDeadLetterSourceQueuesOutput{}
	fnCacher := func(out *sqs.ListDeadLetterSourceQueuesOutput, more bool) bool {
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
	if err := c.SQSAPI.ListDeadLetterSourceQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SQS) ListDeadLetterSourceQueuesWithContext(ctx context.Context, input *sqs.ListDeadLetterSourceQueuesInput, opts ...request.Option) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListDeadLetterSourceQueuesOutput), nil
	}
	out, err := c.SQSAPI.ListDeadLetterSourceQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListQueueTags(input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListQueueTagsOutput), nil
	}
	out, err := c.SQSAPI.ListQueueTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListQueueTagsWithContext(ctx context.Context, input *sqs.ListQueueTagsInput, opts ...request.Option) (*sqs.ListQueueTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListQueueTagsOutput), nil
	}
	out, err := c.SQSAPI.ListQueueTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListQueues(input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListQueuesOutput), nil
	}
	out, err := c.SQSAPI.ListQueues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SQS) ListQueuesPages(input *sqs.ListQueuesInput, fn func(*sqs.ListQueuesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sqs.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &sqs.ListQueuesOutput{}
	fnCacher := func(out *sqs.ListQueuesOutput, more bool) bool {
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
	if err := c.SQSAPI.ListQueuesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SQS) ListQueuesPagesWithContext(ctx context.Context, input *sqs.ListQueuesInput, fn func(*sqs.ListQueuesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sqs.ListQueuesOutput), false)
		return nil
	}
	cachable := true
	output := &sqs.ListQueuesOutput{}
	fnCacher := func(out *sqs.ListQueuesOutput, more bool) bool {
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
	if err := c.SQSAPI.ListQueuesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SQS) ListQueuesWithContext(ctx context.Context, input *sqs.ListQueuesInput, opts ...request.Option) (*sqs.ListQueuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sqs.ListQueuesOutput), nil
	}
	out, err := c.SQSAPI.ListQueuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
