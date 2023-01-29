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
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/athena/athenaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Athena struct {
	athenaiface.AthenaAPI
	cache *cache.Cache
}

func NewAthena(athenaapi athenaiface.AthenaAPI) *Athena {
	return &Athena{
		AthenaAPI: athenaapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Athena) BatchGetNamedQuery(input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetNamedQueryOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetNamedQuery(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) BatchGetNamedQueryWithContext(ctx context.Context, input *athena.BatchGetNamedQueryInput, opts ...request.Option) (*athena.BatchGetNamedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetNamedQueryOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetNamedQueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) BatchGetPreparedStatement(input *athena.BatchGetPreparedStatementInput) (*athena.BatchGetPreparedStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetPreparedStatementOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetPreparedStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) BatchGetPreparedStatementWithContext(ctx context.Context, input *athena.BatchGetPreparedStatementInput, opts ...request.Option) (*athena.BatchGetPreparedStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetPreparedStatementOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetPreparedStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) BatchGetQueryExecution(input *athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetQueryExecutionOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetQueryExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) BatchGetQueryExecutionWithContext(ctx context.Context, input *athena.BatchGetQueryExecutionInput, opts ...request.Option) (*athena.BatchGetQueryExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.BatchGetQueryExecutionOutput), nil
	}
	out, err := c.AthenaAPI.BatchGetQueryExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecution(input *athena.GetCalculationExecutionInput) (*athena.GetCalculationExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecutionCode(input *athena.GetCalculationExecutionCodeInput) (*athena.GetCalculationExecutionCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionCodeOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecutionCode(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecutionCodeWithContext(ctx context.Context, input *athena.GetCalculationExecutionCodeInput, opts ...request.Option) (*athena.GetCalculationExecutionCodeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionCodeOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecutionCodeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecutionStatus(input *athena.GetCalculationExecutionStatusInput) (*athena.GetCalculationExecutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionStatusOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecutionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecutionStatusWithContext(ctx context.Context, input *athena.GetCalculationExecutionStatusInput, opts ...request.Option) (*athena.GetCalculationExecutionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionStatusOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecutionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetCalculationExecutionWithContext(ctx context.Context, input *athena.GetCalculationExecutionInput, opts ...request.Option) (*athena.GetCalculationExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetCalculationExecutionOutput), nil
	}
	out, err := c.AthenaAPI.GetCalculationExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetDataCatalog(input *athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetDataCatalogOutput), nil
	}
	out, err := c.AthenaAPI.GetDataCatalog(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetDataCatalogWithContext(ctx context.Context, input *athena.GetDataCatalogInput, opts ...request.Option) (*athena.GetDataCatalogOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetDataCatalogOutput), nil
	}
	out, err := c.AthenaAPI.GetDataCatalogWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetDatabase(input *athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetDatabaseOutput), nil
	}
	out, err := c.AthenaAPI.GetDatabase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetDatabaseWithContext(ctx context.Context, input *athena.GetDatabaseInput, opts ...request.Option) (*athena.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetDatabaseOutput), nil
	}
	out, err := c.AthenaAPI.GetDatabaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetNamedQuery(input *athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetNamedQueryOutput), nil
	}
	out, err := c.AthenaAPI.GetNamedQuery(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetNamedQueryWithContext(ctx context.Context, input *athena.GetNamedQueryInput, opts ...request.Option) (*athena.GetNamedQueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetNamedQueryOutput), nil
	}
	out, err := c.AthenaAPI.GetNamedQueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetNotebookMetadata(input *athena.GetNotebookMetadataInput) (*athena.GetNotebookMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetNotebookMetadataOutput), nil
	}
	out, err := c.AthenaAPI.GetNotebookMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetNotebookMetadataWithContext(ctx context.Context, input *athena.GetNotebookMetadataInput, opts ...request.Option) (*athena.GetNotebookMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetNotebookMetadataOutput), nil
	}
	out, err := c.AthenaAPI.GetNotebookMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetPreparedStatement(input *athena.GetPreparedStatementInput) (*athena.GetPreparedStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetPreparedStatementOutput), nil
	}
	out, err := c.AthenaAPI.GetPreparedStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetPreparedStatementWithContext(ctx context.Context, input *athena.GetPreparedStatementInput, opts ...request.Option) (*athena.GetPreparedStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetPreparedStatementOutput), nil
	}
	out, err := c.AthenaAPI.GetPreparedStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryExecution(input *athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryExecutionOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryExecutionWithContext(ctx context.Context, input *athena.GetQueryExecutionInput, opts ...request.Option) (*athena.GetQueryExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryExecutionOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryResults(input *athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryResultsOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryResultsPages(input *athena.GetQueryResultsInput, fn func(*athena.GetQueryResultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.GetQueryResultsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.GetQueryResultsOutput{}
	fnCacher := func(out *athena.GetQueryResultsOutput, more bool) bool {
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
	if err := c.AthenaAPI.GetQueryResultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) GetQueryResultsPagesWithContext(ctx context.Context, input *athena.GetQueryResultsInput, fn func(*athena.GetQueryResultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.GetQueryResultsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.GetQueryResultsOutput{}
	fnCacher := func(out *athena.GetQueryResultsOutput, more bool) bool {
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
	if err := c.AthenaAPI.GetQueryResultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) GetQueryResultsWithContext(ctx context.Context, input *athena.GetQueryResultsInput, opts ...request.Option) (*athena.GetQueryResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryResultsOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryRuntimeStatistics(input *athena.GetQueryRuntimeStatisticsInput) (*athena.GetQueryRuntimeStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryRuntimeStatisticsOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryRuntimeStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetQueryRuntimeStatisticsWithContext(ctx context.Context, input *athena.GetQueryRuntimeStatisticsInput, opts ...request.Option) (*athena.GetQueryRuntimeStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetQueryRuntimeStatisticsOutput), nil
	}
	out, err := c.AthenaAPI.GetQueryRuntimeStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetSession(input *athena.GetSessionInput) (*athena.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetSessionOutput), nil
	}
	out, err := c.AthenaAPI.GetSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetSessionStatus(input *athena.GetSessionStatusInput) (*athena.GetSessionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetSessionStatusOutput), nil
	}
	out, err := c.AthenaAPI.GetSessionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetSessionStatusWithContext(ctx context.Context, input *athena.GetSessionStatusInput, opts ...request.Option) (*athena.GetSessionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetSessionStatusOutput), nil
	}
	out, err := c.AthenaAPI.GetSessionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetSessionWithContext(ctx context.Context, input *athena.GetSessionInput, opts ...request.Option) (*athena.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetSessionOutput), nil
	}
	out, err := c.AthenaAPI.GetSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetTableMetadata(input *athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetTableMetadataOutput), nil
	}
	out, err := c.AthenaAPI.GetTableMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetTableMetadataWithContext(ctx context.Context, input *athena.GetTableMetadataInput, opts ...request.Option) (*athena.GetTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetTableMetadataOutput), nil
	}
	out, err := c.AthenaAPI.GetTableMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetWorkGroup(input *athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetWorkGroupOutput), nil
	}
	out, err := c.AthenaAPI.GetWorkGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) GetWorkGroupWithContext(ctx context.Context, input *athena.GetWorkGroupInput, opts ...request.Option) (*athena.GetWorkGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.GetWorkGroupOutput), nil
	}
	out, err := c.AthenaAPI.GetWorkGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListApplicationDPUSizes(input *athena.ListApplicationDPUSizesInput) (*athena.ListApplicationDPUSizesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListApplicationDPUSizesOutput), nil
	}
	out, err := c.AthenaAPI.ListApplicationDPUSizes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListApplicationDPUSizesPages(input *athena.ListApplicationDPUSizesInput, fn func(*athena.ListApplicationDPUSizesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListApplicationDPUSizesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListApplicationDPUSizesOutput{}
	fnCacher := func(out *athena.ListApplicationDPUSizesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListApplicationDPUSizesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListApplicationDPUSizesPagesWithContext(ctx context.Context, input *athena.ListApplicationDPUSizesInput, fn func(*athena.ListApplicationDPUSizesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListApplicationDPUSizesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListApplicationDPUSizesOutput{}
	fnCacher := func(out *athena.ListApplicationDPUSizesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListApplicationDPUSizesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListApplicationDPUSizesWithContext(ctx context.Context, input *athena.ListApplicationDPUSizesInput, opts ...request.Option) (*athena.ListApplicationDPUSizesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListApplicationDPUSizesOutput), nil
	}
	out, err := c.AthenaAPI.ListApplicationDPUSizesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListCalculationExecutions(input *athena.ListCalculationExecutionsInput) (*athena.ListCalculationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListCalculationExecutionsOutput), nil
	}
	out, err := c.AthenaAPI.ListCalculationExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListCalculationExecutionsPages(input *athena.ListCalculationExecutionsInput, fn func(*athena.ListCalculationExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListCalculationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListCalculationExecutionsOutput{}
	fnCacher := func(out *athena.ListCalculationExecutionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListCalculationExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListCalculationExecutionsPagesWithContext(ctx context.Context, input *athena.ListCalculationExecutionsInput, fn func(*athena.ListCalculationExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListCalculationExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListCalculationExecutionsOutput{}
	fnCacher := func(out *athena.ListCalculationExecutionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListCalculationExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListCalculationExecutionsWithContext(ctx context.Context, input *athena.ListCalculationExecutionsInput, opts ...request.Option) (*athena.ListCalculationExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListCalculationExecutionsOutput), nil
	}
	out, err := c.AthenaAPI.ListCalculationExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListDataCatalogs(input *athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListDataCatalogsOutput), nil
	}
	out, err := c.AthenaAPI.ListDataCatalogs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListDataCatalogsPages(input *athena.ListDataCatalogsInput, fn func(*athena.ListDataCatalogsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListDataCatalogsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListDataCatalogsOutput{}
	fnCacher := func(out *athena.ListDataCatalogsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListDataCatalogsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListDataCatalogsPagesWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, fn func(*athena.ListDataCatalogsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListDataCatalogsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListDataCatalogsOutput{}
	fnCacher := func(out *athena.ListDataCatalogsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListDataCatalogsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListDataCatalogsWithContext(ctx context.Context, input *athena.ListDataCatalogsInput, opts ...request.Option) (*athena.ListDataCatalogsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListDataCatalogsOutput), nil
	}
	out, err := c.AthenaAPI.ListDataCatalogsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListDatabases(input *athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListDatabasesOutput), nil
	}
	out, err := c.AthenaAPI.ListDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListDatabasesPages(input *athena.ListDatabasesInput, fn func(*athena.ListDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListDatabasesOutput{}
	fnCacher := func(out *athena.ListDatabasesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListDatabasesPagesWithContext(ctx context.Context, input *athena.ListDatabasesInput, fn func(*athena.ListDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListDatabasesOutput{}
	fnCacher := func(out *athena.ListDatabasesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListDatabasesWithContext(ctx context.Context, input *athena.ListDatabasesInput, opts ...request.Option) (*athena.ListDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListDatabasesOutput), nil
	}
	out, err := c.AthenaAPI.ListDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListEngineVersions(input *athena.ListEngineVersionsInput) (*athena.ListEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListEngineVersionsOutput), nil
	}
	out, err := c.AthenaAPI.ListEngineVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListEngineVersionsPages(input *athena.ListEngineVersionsInput, fn func(*athena.ListEngineVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListEngineVersionsOutput{}
	fnCacher := func(out *athena.ListEngineVersionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListEngineVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListEngineVersionsPagesWithContext(ctx context.Context, input *athena.ListEngineVersionsInput, fn func(*athena.ListEngineVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListEngineVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListEngineVersionsOutput{}
	fnCacher := func(out *athena.ListEngineVersionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListEngineVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListEngineVersionsWithContext(ctx context.Context, input *athena.ListEngineVersionsInput, opts ...request.Option) (*athena.ListEngineVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListEngineVersionsOutput), nil
	}
	out, err := c.AthenaAPI.ListEngineVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListExecutors(input *athena.ListExecutorsInput) (*athena.ListExecutorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListExecutorsOutput), nil
	}
	out, err := c.AthenaAPI.ListExecutors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListExecutorsPages(input *athena.ListExecutorsInput, fn func(*athena.ListExecutorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListExecutorsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListExecutorsOutput{}
	fnCacher := func(out *athena.ListExecutorsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListExecutorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListExecutorsPagesWithContext(ctx context.Context, input *athena.ListExecutorsInput, fn func(*athena.ListExecutorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListExecutorsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListExecutorsOutput{}
	fnCacher := func(out *athena.ListExecutorsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListExecutorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListExecutorsWithContext(ctx context.Context, input *athena.ListExecutorsInput, opts ...request.Option) (*athena.ListExecutorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListExecutorsOutput), nil
	}
	out, err := c.AthenaAPI.ListExecutorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNamedQueries(input *athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNamedQueriesOutput), nil
	}
	out, err := c.AthenaAPI.ListNamedQueries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNamedQueriesPages(input *athena.ListNamedQueriesInput, fn func(*athena.ListNamedQueriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListNamedQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListNamedQueriesOutput{}
	fnCacher := func(out *athena.ListNamedQueriesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListNamedQueriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListNamedQueriesPagesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, fn func(*athena.ListNamedQueriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListNamedQueriesOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListNamedQueriesOutput{}
	fnCacher := func(out *athena.ListNamedQueriesOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListNamedQueriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListNamedQueriesWithContext(ctx context.Context, input *athena.ListNamedQueriesInput, opts ...request.Option) (*athena.ListNamedQueriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNamedQueriesOutput), nil
	}
	out, err := c.AthenaAPI.ListNamedQueriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNotebookMetadata(input *athena.ListNotebookMetadataInput) (*athena.ListNotebookMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNotebookMetadataOutput), nil
	}
	out, err := c.AthenaAPI.ListNotebookMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNotebookMetadataWithContext(ctx context.Context, input *athena.ListNotebookMetadataInput, opts ...request.Option) (*athena.ListNotebookMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNotebookMetadataOutput), nil
	}
	out, err := c.AthenaAPI.ListNotebookMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNotebookSessions(input *athena.ListNotebookSessionsInput) (*athena.ListNotebookSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNotebookSessionsOutput), nil
	}
	out, err := c.AthenaAPI.ListNotebookSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListNotebookSessionsWithContext(ctx context.Context, input *athena.ListNotebookSessionsInput, opts ...request.Option) (*athena.ListNotebookSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListNotebookSessionsOutput), nil
	}
	out, err := c.AthenaAPI.ListNotebookSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListPreparedStatements(input *athena.ListPreparedStatementsInput) (*athena.ListPreparedStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListPreparedStatementsOutput), nil
	}
	out, err := c.AthenaAPI.ListPreparedStatements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListPreparedStatementsPages(input *athena.ListPreparedStatementsInput, fn func(*athena.ListPreparedStatementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListPreparedStatementsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListPreparedStatementsOutput{}
	fnCacher := func(out *athena.ListPreparedStatementsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListPreparedStatementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListPreparedStatementsPagesWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, fn func(*athena.ListPreparedStatementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListPreparedStatementsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListPreparedStatementsOutput{}
	fnCacher := func(out *athena.ListPreparedStatementsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListPreparedStatementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListPreparedStatementsWithContext(ctx context.Context, input *athena.ListPreparedStatementsInput, opts ...request.Option) (*athena.ListPreparedStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListPreparedStatementsOutput), nil
	}
	out, err := c.AthenaAPI.ListPreparedStatementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListQueryExecutions(input *athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListQueryExecutionsOutput), nil
	}
	out, err := c.AthenaAPI.ListQueryExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListQueryExecutionsPages(input *athena.ListQueryExecutionsInput, fn func(*athena.ListQueryExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListQueryExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListQueryExecutionsOutput{}
	fnCacher := func(out *athena.ListQueryExecutionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListQueryExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListQueryExecutionsPagesWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, fn func(*athena.ListQueryExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListQueryExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListQueryExecutionsOutput{}
	fnCacher := func(out *athena.ListQueryExecutionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListQueryExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListQueryExecutionsWithContext(ctx context.Context, input *athena.ListQueryExecutionsInput, opts ...request.Option) (*athena.ListQueryExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListQueryExecutionsOutput), nil
	}
	out, err := c.AthenaAPI.ListQueryExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListSessions(input *athena.ListSessionsInput) (*athena.ListSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListSessionsOutput), nil
	}
	out, err := c.AthenaAPI.ListSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListSessionsPages(input *athena.ListSessionsInput, fn func(*athena.ListSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListSessionsOutput{}
	fnCacher := func(out *athena.ListSessionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListSessionsPagesWithContext(ctx context.Context, input *athena.ListSessionsInput, fn func(*athena.ListSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListSessionsOutput{}
	fnCacher := func(out *athena.ListSessionsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListSessionsWithContext(ctx context.Context, input *athena.ListSessionsInput, opts ...request.Option) (*athena.ListSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListSessionsOutput), nil
	}
	out, err := c.AthenaAPI.ListSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListTableMetadata(input *athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListTableMetadataOutput), nil
	}
	out, err := c.AthenaAPI.ListTableMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListTableMetadataPages(input *athena.ListTableMetadataInput, fn func(*athena.ListTableMetadataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListTableMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListTableMetadataOutput{}
	fnCacher := func(out *athena.ListTableMetadataOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListTableMetadataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListTableMetadataPagesWithContext(ctx context.Context, input *athena.ListTableMetadataInput, fn func(*athena.ListTableMetadataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListTableMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListTableMetadataOutput{}
	fnCacher := func(out *athena.ListTableMetadataOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListTableMetadataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListTableMetadataWithContext(ctx context.Context, input *athena.ListTableMetadataInput, opts ...request.Option) (*athena.ListTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListTableMetadataOutput), nil
	}
	out, err := c.AthenaAPI.ListTableMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListTagsForResource(input *athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListTagsForResourceOutput), nil
	}
	out, err := c.AthenaAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListTagsForResourcePages(input *athena.ListTagsForResourceInput, fn func(*athena.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListTagsForResourceOutput{}
	fnCacher := func(out *athena.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListTagsForResourcePagesWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, fn func(*athena.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListTagsForResourceOutput{}
	fnCacher := func(out *athena.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListTagsForResourceWithContext(ctx context.Context, input *athena.ListTagsForResourceInput, opts ...request.Option) (*athena.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListTagsForResourceOutput), nil
	}
	out, err := c.AthenaAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListWorkGroups(input *athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListWorkGroupsOutput), nil
	}
	out, err := c.AthenaAPI.ListWorkGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Athena) ListWorkGroupsPages(input *athena.ListWorkGroupsInput, fn func(*athena.ListWorkGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListWorkGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListWorkGroupsOutput{}
	fnCacher := func(out *athena.ListWorkGroupsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListWorkGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListWorkGroupsPagesWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, fn func(*athena.ListWorkGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*athena.ListWorkGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &athena.ListWorkGroupsOutput{}
	fnCacher := func(out *athena.ListWorkGroupsOutput, more bool) bool {
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
	if err := c.AthenaAPI.ListWorkGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Athena) ListWorkGroupsWithContext(ctx context.Context, input *athena.ListWorkGroupsInput, opts ...request.Option) (*athena.ListWorkGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*athena.ListWorkGroupsOutput), nil
	}
	out, err := c.AthenaAPI.ListWorkGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
