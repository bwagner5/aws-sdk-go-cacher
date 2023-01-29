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
package mqcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mq"
	"github.com/aws/aws-sdk-go/service/mq/mqiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MQ struct {
	mqiface.MQAPI
	cache *cache.Cache
}

func New(mqapi mqiface.MQAPI) *MQ {
	return &MQ{
		MQAPI: mqapi,
		cache: cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MQ) DescribeBrokerEngineTypes(input *mq.DescribeBrokerEngineTypesInput) (*mq.DescribeBrokerEngineTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeBrokerEngineTypesOutput), nil
	}
	out, err := c.MQAPI.DescribeBrokerEngineTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) DescribeBrokerEngineTypesWithContext(ctx context.Context, input *mq.DescribeBrokerEngineTypesInput, opts ...request.Option) (*mq.DescribeBrokerEngineTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeBrokerEngineTypesOutput), nil
	}
	out, err := c.MQAPI.DescribeBrokerEngineTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) DescribeBrokerInstanceOptions(input *mq.DescribeBrokerInstanceOptionsInput) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeBrokerInstanceOptionsOutput), nil
	}
	out, err := c.MQAPI.DescribeBrokerInstanceOptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) DescribeBrokerInstanceOptionsWithContext(ctx context.Context, input *mq.DescribeBrokerInstanceOptionsInput, opts ...request.Option) (*mq.DescribeBrokerInstanceOptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeBrokerInstanceOptionsOutput), nil
	}
	out, err := c.MQAPI.DescribeBrokerInstanceOptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) DescribeConfiguration(input *mq.DescribeConfigurationInput) (*mq.DescribeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeConfigurationOutput), nil
	}
	out, err := c.MQAPI.DescribeConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) DescribeConfigurationWithContext(ctx context.Context, input *mq.DescribeConfigurationInput, opts ...request.Option) (*mq.DescribeConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.DescribeConfigurationOutput), nil
	}
	out, err := c.MQAPI.DescribeConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) ListBrokersPages(input *mq.ListBrokersInput, fn func(*mq.ListBrokersResponse, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mq.ListBrokersResponse), false)
		return nil
	}
	cachable := true
	output := &mq.ListBrokersResponse{}
	fnCacher := func(out *mq.ListBrokersResponse, more bool) bool {
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
	if err := c.MQAPI.ListBrokersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MQ) ListBrokersPagesWithContext(ctx context.Context, input *mq.ListBrokersInput, fn func(*mq.ListBrokersResponse, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mq.ListBrokersResponse), false)
		return nil
	}
	cachable := true
	output := &mq.ListBrokersResponse{}
	fnCacher := func(out *mq.ListBrokersResponse, more bool) bool {
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
	if err := c.MQAPI.ListBrokersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MQ) ListTags(input *mq.ListTagsInput) (*mq.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.ListTagsOutput), nil
	}
	out, err := c.MQAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MQ) ListTagsWithContext(ctx context.Context, input *mq.ListTagsInput, opts ...request.Option) (*mq.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mq.ListTagsOutput), nil
	}
	out, err := c.MQAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
