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
package timestreamwritecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/timestreamwrite"
	"github.com/aws/aws-sdk-go/service/timestreamwrite/timestreamwriteiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type TimestreamWrite struct {
	timestreamwriteiface.TimestreamWriteAPI
	cache *cache.Cache
}

func New(timestreamwriteapi timestreamwriteiface.TimestreamWriteAPI) *TimestreamWrite {
	return &TimestreamWrite{
		TimestreamWriteAPI: timestreamwriteapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *TimestreamWrite) DescribeDatabase(input *timestreamwrite.DescribeDatabaseInput) (*timestreamwrite.DescribeDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeDatabaseOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeDatabase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) DescribeDatabaseWithContext(ctx context.Context, input *timestreamwrite.DescribeDatabaseInput, opts ...request.Option) (*timestreamwrite.DescribeDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeDatabaseOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeDatabaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) DescribeEndpoints(input *timestreamwrite.DescribeEndpointsInput) (*timestreamwrite.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeEndpointsOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) DescribeEndpointsWithContext(ctx context.Context, input *timestreamwrite.DescribeEndpointsInput, opts ...request.Option) (*timestreamwrite.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeEndpointsOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) DescribeTable(input *timestreamwrite.DescribeTableInput) (*timestreamwrite.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeTableOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) DescribeTableWithContext(ctx context.Context, input *timestreamwrite.DescribeTableInput, opts ...request.Option) (*timestreamwrite.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.DescribeTableOutput), nil
	}
	out, err := c.TimestreamWriteAPI.DescribeTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListDatabases(input *timestreamwrite.ListDatabasesInput) (*timestreamwrite.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListDatabasesOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListDatabasesPages(input *timestreamwrite.ListDatabasesInput, fn func(*timestreamwrite.ListDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamwrite.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamwrite.ListDatabasesOutput{}
	fnCacher := func(out *timestreamwrite.ListDatabasesOutput, more bool) bool {
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
	if err := c.TimestreamWriteAPI.ListDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamWrite) ListDatabasesPagesWithContext(ctx context.Context, input *timestreamwrite.ListDatabasesInput, fn func(*timestreamwrite.ListDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamwrite.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamwrite.ListDatabasesOutput{}
	fnCacher := func(out *timestreamwrite.ListDatabasesOutput, more bool) bool {
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
	if err := c.TimestreamWriteAPI.ListDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamWrite) ListDatabasesWithContext(ctx context.Context, input *timestreamwrite.ListDatabasesInput, opts ...request.Option) (*timestreamwrite.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListDatabasesOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListTables(input *timestreamwrite.ListTablesInput) (*timestreamwrite.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListTablesOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListTablesPages(input *timestreamwrite.ListTablesInput, fn func(*timestreamwrite.ListTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamwrite.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamwrite.ListTablesOutput{}
	fnCacher := func(out *timestreamwrite.ListTablesOutput, more bool) bool {
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
	if err := c.TimestreamWriteAPI.ListTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamWrite) ListTablesPagesWithContext(ctx context.Context, input *timestreamwrite.ListTablesInput, fn func(*timestreamwrite.ListTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*timestreamwrite.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &timestreamwrite.ListTablesOutput{}
	fnCacher := func(out *timestreamwrite.ListTablesOutput, more bool) bool {
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
	if err := c.TimestreamWriteAPI.ListTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TimestreamWrite) ListTablesWithContext(ctx context.Context, input *timestreamwrite.ListTablesInput, opts ...request.Option) (*timestreamwrite.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListTablesOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListTagsForResource(input *timestreamwrite.ListTagsForResourceInput) (*timestreamwrite.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListTagsForResourceOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TimestreamWrite) ListTagsForResourceWithContext(ctx context.Context, input *timestreamwrite.ListTagsForResourceInput, opts ...request.Option) (*timestreamwrite.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*timestreamwrite.ListTagsForResourceOutput), nil
	}
	out, err := c.TimestreamWriteAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
