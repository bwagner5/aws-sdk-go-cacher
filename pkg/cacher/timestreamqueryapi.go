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
	"github.com/aws/aws-sdk-go/service/timestreamquery"
	"github.com/aws/aws-sdk-go/service/timestreamquery/timestreamqueryiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type TimestreamQuery struct {
	timestreamqueryiface.TimestreamQueryAPI
	cache *cache.Cache
}

func NewTimestreamQuery(timestreamqueryapi timestreamqueryiface.TimestreamQueryAPI) *TimestreamQuery {
	return &TimestreamQuery{
		TimestreamQueryAPI: timestreamqueryapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *TimestreamQuery) DescribeEndpoints(input *timestreamquery.DescribeEndpointsInput) (*timestreamquery.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.DescribeEndpointsOutput), nil
	}
	out, err := c.TimestreamQueryAPI.DescribeEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) DescribeEndpointsWithContext(ctx context.Context, input *timestreamquery.DescribeEndpointsInput, opts ...request.Option) (*timestreamquery.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.DescribeEndpointsOutput), nil
	}
	out, err := c.TimestreamQueryAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) DescribeScheduledQuery(input *timestreamquery.DescribeScheduledQueryInput) (*timestreamquery.DescribeScheduledQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.DescribeScheduledQueryOutput), nil
	}
	out, err := c.TimestreamQueryAPI.DescribeScheduledQuery(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) DescribeScheduledQueryWithContext(ctx context.Context, input *timestreamquery.DescribeScheduledQueryInput, opts ...request.Option) (*timestreamquery.DescribeScheduledQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.DescribeScheduledQueryOutput), nil
	}
	out, err := c.TimestreamQueryAPI.DescribeScheduledQueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) ListScheduledQueries(input *timestreamquery.ListScheduledQueriesInput) (*timestreamquery.ListScheduledQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.ListScheduledQueriesOutput), nil
	}
	out, err := c.TimestreamQueryAPI.ListScheduledQueries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) ListScheduledQueriesPages(input *timestreamquery.ListScheduledQueriesInput, fn func(*timestreamquery.ListScheduledQueriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.ListScheduledQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.ListScheduledQueriesOutput{}
	fnCacher := func(out *timestreamquery.ListScheduledQueriesOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.ListScheduledQueriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) ListScheduledQueriesPagesWithContext(ctx context.Context, input *timestreamquery.ListScheduledQueriesInput, fn func(*timestreamquery.ListScheduledQueriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.ListScheduledQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.ListScheduledQueriesOutput{}
	fnCacher := func(out *timestreamquery.ListScheduledQueriesOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.ListScheduledQueriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) ListScheduledQueriesWithContext(ctx context.Context, input *timestreamquery.ListScheduledQueriesInput, opts ...request.Option) (*timestreamquery.ListScheduledQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.ListScheduledQueriesOutput), nil
	}
	out, err := c.TimestreamQueryAPI.ListScheduledQueriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) ListTagsForResource(input *timestreamquery.ListTagsForResourceInput) (*timestreamquery.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.ListTagsForResourceOutput), nil
	}
	out, err := c.TimestreamQueryAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) ListTagsForResourcePages(input *timestreamquery.ListTagsForResourceInput, fn func(*timestreamquery.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.ListTagsForResourceOutput{}
	fnCacher := func(out *timestreamquery.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) ListTagsForResourcePagesWithContext(ctx context.Context, input *timestreamquery.ListTagsForResourceInput, fn func(*timestreamquery.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.ListTagsForResourceOutput{}
	fnCacher := func(out *timestreamquery.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) ListTagsForResourceWithContext(ctx context.Context, input *timestreamquery.ListTagsForResourceInput, opts ...request.Option) (*timestreamquery.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.ListTagsForResourceOutput), nil
	}
	out, err := c.TimestreamQueryAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) Query(input *timestreamquery.QueryInput) (*timestreamquery.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.QueryOutput), nil
	}
	out, err := c.TimestreamQueryAPI.Query(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamQuery) QueryPages(input *timestreamquery.QueryInput, fn func(*timestreamquery.QueryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.QueryOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.QueryOutput{}
	fnCacher := func(out *timestreamquery.QueryOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.QueryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) QueryPagesWithContext(ctx context.Context, input *timestreamquery.QueryInput, fn func(*timestreamquery.QueryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamquery.QueryOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamquery.QueryOutput{}
	fnCacher := func(out *timestreamquery.QueryOutput, more bool) bool {
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
	if err := c.TimestreamQueryAPI.QueryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamQuery) QueryWithContext(ctx context.Context, input *timestreamquery.QueryInput, opts ...request.Option) (*timestreamquery.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamquery.QueryOutput), nil
	}
	out, err := c.TimestreamQueryAPI.QueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
