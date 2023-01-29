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
	"github.com/aws/aws-sdk-go/service/redshiftdataservice"
	"github.com/aws/aws-sdk-go/service/redshiftdataservice/redshiftdataserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RedshiftDataAPIService struct {
	redshiftdataserviceiface.RedshiftDataAPIServiceAPI
	cache *cache.Cache
}

func NewRedshiftDataAPIService(redshiftdataserviceapi redshiftdataserviceiface.RedshiftDataAPIServiceAPI) *RedshiftDataAPIService {
	return &RedshiftDataAPIService{
		RedshiftDataAPIServiceAPI: redshiftdataserviceapi,
		cache:                     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RedshiftDataAPIService) BatchExecuteStatement(input *redshiftdataapiservice.BatchExecuteStatementInput) (*redshiftdataapiservice.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.BatchExecuteStatementOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.BatchExecuteStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) BatchExecuteStatementWithContext(ctx context.Context, input *redshiftdataapiservice.BatchExecuteStatementInput, opts ...request.Option) (*redshiftdataapiservice.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.BatchExecuteStatementOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.BatchExecuteStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) DescribeStatement(input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.DescribeStatementOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.DescribeStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) DescribeStatementWithContext(ctx context.Context, input *redshiftdataapiservice.DescribeStatementInput, opts ...request.Option) (*redshiftdataapiservice.DescribeStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.DescribeStatementOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.DescribeStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) DescribeTable(input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.DescribeTableOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.DescribeTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) DescribeTablePages(input *redshiftdataapiservice.DescribeTableInput, fn func(*redshiftdataapiservice.DescribeTableOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.DescribeTableOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.DescribeTableOutput{}
	fnCacher := func(out *redshiftdataapiservice.DescribeTableOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.DescribeTablePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) DescribeTablePagesWithContext(ctx context.Context, input *redshiftdataapiservice.DescribeTableInput, fn func(*redshiftdataapiservice.DescribeTableOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.DescribeTableOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.DescribeTableOutput{}
	fnCacher := func(out *redshiftdataapiservice.DescribeTableOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.DescribeTablePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) DescribeTableWithContext(ctx context.Context, input *redshiftdataapiservice.DescribeTableInput, opts ...request.Option) (*redshiftdataapiservice.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.DescribeTableOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.DescribeTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) GetStatementResult(input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.GetStatementResultOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.GetStatementResult(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) GetStatementResultPages(input *redshiftdataapiservice.GetStatementResultInput, fn func(*redshiftdataapiservice.GetStatementResultOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.GetStatementResultOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.GetStatementResultOutput{}
	fnCacher := func(out *redshiftdataapiservice.GetStatementResultOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.GetStatementResultPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) GetStatementResultPagesWithContext(ctx context.Context, input *redshiftdataapiservice.GetStatementResultInput, fn func(*redshiftdataapiservice.GetStatementResultOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.GetStatementResultOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.GetStatementResultOutput{}
	fnCacher := func(out *redshiftdataapiservice.GetStatementResultOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.GetStatementResultPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) GetStatementResultWithContext(ctx context.Context, input *redshiftdataapiservice.GetStatementResultInput, opts ...request.Option) (*redshiftdataapiservice.GetStatementResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.GetStatementResultOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.GetStatementResultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListDatabases(input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListDatabasesOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListDatabasesPages(input *redshiftdataapiservice.ListDatabasesInput, fn func(*redshiftdataapiservice.ListDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListDatabasesOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListDatabasesOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListDatabasesPagesWithContext(ctx context.Context, input *redshiftdataapiservice.ListDatabasesInput, fn func(*redshiftdataapiservice.ListDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListDatabasesOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListDatabasesOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListDatabasesWithContext(ctx context.Context, input *redshiftdataapiservice.ListDatabasesInput, opts ...request.Option) (*redshiftdataapiservice.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListDatabasesOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListSchemas(input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListSchemasOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListSchemasPages(input *redshiftdataapiservice.ListSchemasInput, fn func(*redshiftdataapiservice.ListSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListSchemasOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListSchemasOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListSchemasPagesWithContext(ctx context.Context, input *redshiftdataapiservice.ListSchemasInput, fn func(*redshiftdataapiservice.ListSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListSchemasOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListSchemasOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListSchemasWithContext(ctx context.Context, input *redshiftdataapiservice.ListSchemasInput, opts ...request.Option) (*redshiftdataapiservice.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListSchemasOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListStatements(input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListStatementsOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListStatements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListStatementsPages(input *redshiftdataapiservice.ListStatementsInput, fn func(*redshiftdataapiservice.ListStatementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListStatementsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListStatementsOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListStatementsOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListStatementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListStatementsPagesWithContext(ctx context.Context, input *redshiftdataapiservice.ListStatementsInput, fn func(*redshiftdataapiservice.ListStatementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListStatementsOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListStatementsOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListStatementsOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListStatementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListStatementsWithContext(ctx context.Context, input *redshiftdataapiservice.ListStatementsInput, opts ...request.Option) (*redshiftdataapiservice.ListStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListStatementsOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListStatementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListTables(input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListTablesOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RedshiftDataAPIService) ListTablesPages(input *redshiftdataapiservice.ListTablesInput, fn func(*redshiftdataapiservice.ListTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListTablesOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListTablesOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListTablesPagesWithContext(ctx context.Context, input *redshiftdataapiservice.ListTablesInput, fn func(*redshiftdataapiservice.ListTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*redshiftdataapiservice.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &redshiftdataapiservice.ListTablesOutput{}
	fnCacher := func(out *redshiftdataapiservice.ListTablesOutput, more bool) bool {
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
	if err := c.RedshiftDataAPIServiceAPI.ListTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RedshiftDataAPIService) ListTablesWithContext(ctx context.Context, input *redshiftdataapiservice.ListTablesInput, opts ...request.Option) (*redshiftdataapiservice.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*redshiftdataapiservice.ListTablesOutput), nil
	}
	out, err := c.RedshiftDataAPIServiceAPI.ListTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
