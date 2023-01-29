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
	"github.com/aws/aws-sdk-go/service/honeycode"
	"github.com/aws/aws-sdk-go/service/honeycode/honeycodeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Honeycode struct {
	honeycodeiface.HoneycodeAPI
	cache *cache.Cache
}

func NewHoneycode(honeycodeapi honeycodeiface.HoneycodeAPI) *Honeycode {
	return &Honeycode{
		HoneycodeAPI: honeycodeapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Honeycode) BatchCreateTableRows(input *honeycode.BatchCreateTableRowsInput) (*honeycode.BatchCreateTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchCreateTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchCreateTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchCreateTableRowsWithContext(ctx context.Context, input *honeycode.BatchCreateTableRowsInput, opts ...request.Option) (*honeycode.BatchCreateTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchCreateTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchCreateTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchDeleteTableRows(input *honeycode.BatchDeleteTableRowsInput) (*honeycode.BatchDeleteTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchDeleteTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchDeleteTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchDeleteTableRowsWithContext(ctx context.Context, input *honeycode.BatchDeleteTableRowsInput, opts ...request.Option) (*honeycode.BatchDeleteTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchDeleteTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchDeleteTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchUpdateTableRows(input *honeycode.BatchUpdateTableRowsInput) (*honeycode.BatchUpdateTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchUpdateTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchUpdateTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchUpdateTableRowsWithContext(ctx context.Context, input *honeycode.BatchUpdateTableRowsInput, opts ...request.Option) (*honeycode.BatchUpdateTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchUpdateTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchUpdateTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchUpsertTableRows(input *honeycode.BatchUpsertTableRowsInput) (*honeycode.BatchUpsertTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchUpsertTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchUpsertTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) BatchUpsertTableRowsWithContext(ctx context.Context, input *honeycode.BatchUpsertTableRowsInput, opts ...request.Option) (*honeycode.BatchUpsertTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.BatchUpsertTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.BatchUpsertTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) DescribeTableDataImportJob(input *honeycode.DescribeTableDataImportJobInput) (*honeycode.DescribeTableDataImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.DescribeTableDataImportJobOutput), nil
	}
	out, err := c.HoneycodeAPI.DescribeTableDataImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) DescribeTableDataImportJobWithContext(ctx context.Context, input *honeycode.DescribeTableDataImportJobInput, opts ...request.Option) (*honeycode.DescribeTableDataImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.DescribeTableDataImportJobOutput), nil
	}
	out, err := c.HoneycodeAPI.DescribeTableDataImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) GetScreenData(input *honeycode.GetScreenDataInput) (*honeycode.GetScreenDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.GetScreenDataOutput), nil
	}
	out, err := c.HoneycodeAPI.GetScreenData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) GetScreenDataWithContext(ctx context.Context, input *honeycode.GetScreenDataInput, opts ...request.Option) (*honeycode.GetScreenDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.GetScreenDataOutput), nil
	}
	out, err := c.HoneycodeAPI.GetScreenDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTableColumns(input *honeycode.ListTableColumnsInput) (*honeycode.ListTableColumnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTableColumnsOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTableColumns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTableColumnsPages(input *honeycode.ListTableColumnsInput, fn func(*honeycode.ListTableColumnsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTableColumnsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTableColumnsOutput{}
	fnCacher := func(out *honeycode.ListTableColumnsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTableColumnsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTableColumnsPagesWithContext(ctx context.Context, input *honeycode.ListTableColumnsInput, fn func(*honeycode.ListTableColumnsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTableColumnsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTableColumnsOutput{}
	fnCacher := func(out *honeycode.ListTableColumnsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTableColumnsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTableColumnsWithContext(ctx context.Context, input *honeycode.ListTableColumnsInput, opts ...request.Option) (*honeycode.ListTableColumnsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTableColumnsOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTableColumnsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTableRows(input *honeycode.ListTableRowsInput) (*honeycode.ListTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTableRowsPages(input *honeycode.ListTableRowsInput, fn func(*honeycode.ListTableRowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTableRowsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTableRowsOutput{}
	fnCacher := func(out *honeycode.ListTableRowsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTableRowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTableRowsPagesWithContext(ctx context.Context, input *honeycode.ListTableRowsInput, fn func(*honeycode.ListTableRowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTableRowsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTableRowsOutput{}
	fnCacher := func(out *honeycode.ListTableRowsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTableRowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTableRowsWithContext(ctx context.Context, input *honeycode.ListTableRowsInput, opts ...request.Option) (*honeycode.ListTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTables(input *honeycode.ListTablesInput) (*honeycode.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTablesOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTablesPages(input *honeycode.ListTablesInput, fn func(*honeycode.ListTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTablesOutput{}
	fnCacher := func(out *honeycode.ListTablesOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTablesPagesWithContext(ctx context.Context, input *honeycode.ListTablesInput, fn func(*honeycode.ListTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.ListTablesOutput{}
	fnCacher := func(out *honeycode.ListTablesOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.ListTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) ListTablesWithContext(ctx context.Context, input *honeycode.ListTablesInput, opts ...request.Option) (*honeycode.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTablesOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTagsForResource(input *honeycode.ListTagsForResourceInput) (*honeycode.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTagsForResourceOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) ListTagsForResourceWithContext(ctx context.Context, input *honeycode.ListTagsForResourceInput, opts ...request.Option) (*honeycode.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.ListTagsForResourceOutput), nil
	}
	out, err := c.HoneycodeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) QueryTableRows(input *honeycode.QueryTableRowsInput) (*honeycode.QueryTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.QueryTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.QueryTableRows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Honeycode) QueryTableRowsPages(input *honeycode.QueryTableRowsInput, fn func(*honeycode.QueryTableRowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.QueryTableRowsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.QueryTableRowsOutput{}
	fnCacher := func(out *honeycode.QueryTableRowsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.QueryTableRowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) QueryTableRowsPagesWithContext(ctx context.Context, input *honeycode.QueryTableRowsInput, fn func(*honeycode.QueryTableRowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*honeycode.QueryTableRowsOutput), false)
		return nil
	}
	cachable := true
	output := &honeycode.QueryTableRowsOutput{}
	fnCacher := func(out *honeycode.QueryTableRowsOutput, more bool) bool {
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
	if err := c.HoneycodeAPI.QueryTableRowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Honeycode) QueryTableRowsWithContext(ctx context.Context, input *honeycode.QueryTableRowsInput, opts ...request.Option) (*honeycode.QueryTableRowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*honeycode.QueryTableRowsOutput), nil
	}
	out, err := c.HoneycodeAPI.QueryTableRowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
