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
package kendrarankingcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kendraranking"
	"github.com/aws/aws-sdk-go/service/kendraranking/kendrarankingiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type KendraRanking struct {
	kendrarankingiface.KendraRankingAPI
	cache *cache.Cache
}

func New(kendrarankingapi kendrarankingiface.KendraRankingAPI) *KendraRanking {
	return &KendraRanking{
		KendraRankingAPI: kendrarankingapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *KendraRanking) DescribeRescoreExecutionPlan(input *kendraranking.DescribeRescoreExecutionPlanInput) (*kendraranking.DescribeRescoreExecutionPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.DescribeRescoreExecutionPlanOutput), nil
	}
	out, err := c.KendraRankingAPI.DescribeRescoreExecutionPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KendraRanking) DescribeRescoreExecutionPlanWithContext(ctx context.Context, input *kendraranking.DescribeRescoreExecutionPlanInput, opts ...request.Option) (*kendraranking.DescribeRescoreExecutionPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.DescribeRescoreExecutionPlanOutput), nil
	}
	out, err := c.KendraRankingAPI.DescribeRescoreExecutionPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KendraRanking) ListRescoreExecutionPlans(input *kendraranking.ListRescoreExecutionPlansInput) (*kendraranking.ListRescoreExecutionPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.ListRescoreExecutionPlansOutput), nil
	}
	out, err := c.KendraRankingAPI.ListRescoreExecutionPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KendraRanking) ListRescoreExecutionPlansPages(input *kendraranking.ListRescoreExecutionPlansInput, fn func(*kendraranking.ListRescoreExecutionPlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendraranking.ListRescoreExecutionPlansOutput), false)
		return nil
	}
	cachable := true
	output := &kendraranking.ListRescoreExecutionPlansOutput{}
	fnCacher := func(out *kendraranking.ListRescoreExecutionPlansOutput, more bool) bool {
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
	if err := c.KendraRankingAPI.ListRescoreExecutionPlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KendraRanking) ListRescoreExecutionPlansPagesWithContext(ctx context.Context, input *kendraranking.ListRescoreExecutionPlansInput, fn func(*kendraranking.ListRescoreExecutionPlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendraranking.ListRescoreExecutionPlansOutput), false)
		return nil
	}
	cachable := true
	output := &kendraranking.ListRescoreExecutionPlansOutput{}
	fnCacher := func(out *kendraranking.ListRescoreExecutionPlansOutput, more bool) bool {
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
	if err := c.KendraRankingAPI.ListRescoreExecutionPlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *KendraRanking) ListRescoreExecutionPlansWithContext(ctx context.Context, input *kendraranking.ListRescoreExecutionPlansInput, opts ...request.Option) (*kendraranking.ListRescoreExecutionPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.ListRescoreExecutionPlansOutput), nil
	}
	out, err := c.KendraRankingAPI.ListRescoreExecutionPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KendraRanking) ListTagsForResource(input *kendraranking.ListTagsForResourceInput) (*kendraranking.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.ListTagsForResourceOutput), nil
	}
	out, err := c.KendraRankingAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *KendraRanking) ListTagsForResourceWithContext(ctx context.Context, input *kendraranking.ListTagsForResourceInput, opts ...request.Option) (*kendraranking.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendraranking.ListTagsForResourceOutput), nil
	}
	out, err := c.KendraRankingAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
