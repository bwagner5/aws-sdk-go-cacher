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
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DynamoDB struct {
	dynamodbiface.DynamoDBAPI
	cache *cache.Cache
}

func NewDynamoDB(dynamodbapi dynamodbiface.DynamoDBAPI) *DynamoDB {
	return &DynamoDB{
		DynamoDBAPI: dynamodbapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DynamoDB) BatchExecuteStatement(input *dynamodb.BatchExecuteStatementInput) (*dynamodb.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchExecuteStatementOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchExecuteStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) BatchExecuteStatementWithContext(ctx context.Context, input *dynamodb.BatchExecuteStatementInput, opts ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchExecuteStatementOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchExecuteStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) BatchGetItem(input *dynamodb.BatchGetItemInput) (*dynamodb.BatchGetItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchGetItemOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchGetItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) BatchGetItemPages(input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.BatchGetItemOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.BatchGetItemOutput{}
	fnCacher := func(out *dynamodb.BatchGetItemOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.BatchGetItemPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) BatchGetItemPagesWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, fn func(*dynamodb.BatchGetItemOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.BatchGetItemOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.BatchGetItemOutput{}
	fnCacher := func(out *dynamodb.BatchGetItemOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.BatchGetItemPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) BatchGetItemWithContext(ctx context.Context, input *dynamodb.BatchGetItemInput, opts ...request.Option) (*dynamodb.BatchGetItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchGetItemOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchGetItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchWriteItemOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchWriteItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) BatchWriteItemWithContext(ctx context.Context, input *dynamodb.BatchWriteItemInput, opts ...request.Option) (*dynamodb.BatchWriteItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.BatchWriteItemOutput), nil
	}
	out, err := c.DynamoDBAPI.BatchWriteItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeBackup(input *dynamodb.DescribeBackupInput) (*dynamodb.DescribeBackupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeBackupOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeBackup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeBackupWithContext(ctx context.Context, input *dynamodb.DescribeBackupInput, opts ...request.Option) (*dynamodb.DescribeBackupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeBackupOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeBackupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeContinuousBackups(input *dynamodb.DescribeContinuousBackupsInput) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeContinuousBackupsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeContinuousBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeContinuousBackupsWithContext(ctx context.Context, input *dynamodb.DescribeContinuousBackupsInput, opts ...request.Option) (*dynamodb.DescribeContinuousBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeContinuousBackupsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeContinuousBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeContributorInsights(input *dynamodb.DescribeContributorInsightsInput) (*dynamodb.DescribeContributorInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeContributorInsightsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeContributorInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeContributorInsightsWithContext(ctx context.Context, input *dynamodb.DescribeContributorInsightsInput, opts ...request.Option) (*dynamodb.DescribeContributorInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeContributorInsightsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeContributorInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeEndpoints(input *dynamodb.DescribeEndpointsInput) (*dynamodb.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeEndpointsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeEndpointsWithContext(ctx context.Context, input *dynamodb.DescribeEndpointsInput, opts ...request.Option) (*dynamodb.DescribeEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeEndpointsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeExport(input *dynamodb.DescribeExportInput) (*dynamodb.DescribeExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeExportOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeExportWithContext(ctx context.Context, input *dynamodb.DescribeExportInput, opts ...request.Option) (*dynamodb.DescribeExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeExportOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeGlobalTable(input *dynamodb.DescribeGlobalTableInput) (*dynamodb.DescribeGlobalTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeGlobalTableOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeGlobalTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeGlobalTableSettings(input *dynamodb.DescribeGlobalTableSettingsInput) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeGlobalTableSettingsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeGlobalTableSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeGlobalTableSettingsWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableSettingsInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeGlobalTableSettingsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeGlobalTableSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeGlobalTableWithContext(ctx context.Context, input *dynamodb.DescribeGlobalTableInput, opts ...request.Option) (*dynamodb.DescribeGlobalTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeGlobalTableOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeGlobalTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeImport(input *dynamodb.DescribeImportInput) (*dynamodb.DescribeImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeImportOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeImport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeImportWithContext(ctx context.Context, input *dynamodb.DescribeImportInput, opts ...request.Option) (*dynamodb.DescribeImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeImportOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeImportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeKinesisStreamingDestination(input *dynamodb.DescribeKinesisStreamingDestinationInput) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeKinesisStreamingDestinationOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeKinesisStreamingDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeKinesisStreamingDestinationWithContext(ctx context.Context, input *dynamodb.DescribeKinesisStreamingDestinationInput, opts ...request.Option) (*dynamodb.DescribeKinesisStreamingDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeKinesisStreamingDestinationOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeKinesisStreamingDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeLimits(input *dynamodb.DescribeLimitsInput) (*dynamodb.DescribeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeLimitsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeLimitsWithContext(ctx context.Context, input *dynamodb.DescribeLimitsInput, opts ...request.Option) (*dynamodb.DescribeLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeLimitsOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTableOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTableReplicaAutoScaling(input *dynamodb.DescribeTableReplicaAutoScalingInput) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTableReplicaAutoScalingOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTableReplicaAutoScaling(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTableReplicaAutoScalingWithContext(ctx context.Context, input *dynamodb.DescribeTableReplicaAutoScalingInput, opts ...request.Option) (*dynamodb.DescribeTableReplicaAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTableReplicaAutoScalingOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTableReplicaAutoScalingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTableWithContext(ctx context.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTableOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTimeToLive(input *dynamodb.DescribeTimeToLiveInput) (*dynamodb.DescribeTimeToLiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTimeToLiveOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTimeToLive(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) DescribeTimeToLiveWithContext(ctx context.Context, input *dynamodb.DescribeTimeToLiveInput, opts ...request.Option) (*dynamodb.DescribeTimeToLiveOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.DescribeTimeToLiveOutput), nil
	}
	out, err := c.DynamoDBAPI.DescribeTimeToLiveWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.GetItemOutput), nil
	}
	out, err := c.DynamoDBAPI.GetItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) GetItemWithContext(ctx context.Context, input *dynamodb.GetItemInput, opts ...request.Option) (*dynamodb.GetItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.GetItemOutput), nil
	}
	out, err := c.DynamoDBAPI.GetItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListBackups(input *dynamodb.ListBackupsInput) (*dynamodb.ListBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListBackupsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListBackups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListBackupsWithContext(ctx context.Context, input *dynamodb.ListBackupsInput, opts ...request.Option) (*dynamodb.ListBackupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListBackupsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListBackupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListContributorInsights(input *dynamodb.ListContributorInsightsInput) (*dynamodb.ListContributorInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListContributorInsightsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListContributorInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListContributorInsightsPages(input *dynamodb.ListContributorInsightsInput, fn func(*dynamodb.ListContributorInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListContributorInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListContributorInsightsOutput{}
	fnCacher := func(out *dynamodb.ListContributorInsightsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListContributorInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListContributorInsightsPagesWithContext(ctx context.Context, input *dynamodb.ListContributorInsightsInput, fn func(*dynamodb.ListContributorInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListContributorInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListContributorInsightsOutput{}
	fnCacher := func(out *dynamodb.ListContributorInsightsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListContributorInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListContributorInsightsWithContext(ctx context.Context, input *dynamodb.ListContributorInsightsInput, opts ...request.Option) (*dynamodb.ListContributorInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListContributorInsightsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListContributorInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListExports(input *dynamodb.ListExportsInput) (*dynamodb.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListExportsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListExportsPages(input *dynamodb.ListExportsInput, fn func(*dynamodb.ListExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListExportsOutput{}
	fnCacher := func(out *dynamodb.ListExportsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListExportsPagesWithContext(ctx context.Context, input *dynamodb.ListExportsInput, fn func(*dynamodb.ListExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListExportsOutput{}
	fnCacher := func(out *dynamodb.ListExportsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListExportsWithContext(ctx context.Context, input *dynamodb.ListExportsInput, opts ...request.Option) (*dynamodb.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListExportsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListGlobalTables(input *dynamodb.ListGlobalTablesInput) (*dynamodb.ListGlobalTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListGlobalTablesOutput), nil
	}
	out, err := c.DynamoDBAPI.ListGlobalTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListGlobalTablesWithContext(ctx context.Context, input *dynamodb.ListGlobalTablesInput, opts ...request.Option) (*dynamodb.ListGlobalTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListGlobalTablesOutput), nil
	}
	out, err := c.DynamoDBAPI.ListGlobalTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListImports(input *dynamodb.ListImportsInput) (*dynamodb.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListImportsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListImports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListImportsPages(input *dynamodb.ListImportsInput, fn func(*dynamodb.ListImportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListImportsOutput{}
	fnCacher := func(out *dynamodb.ListImportsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListImportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListImportsPagesWithContext(ctx context.Context, input *dynamodb.ListImportsInput, fn func(*dynamodb.ListImportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListImportsOutput{}
	fnCacher := func(out *dynamodb.ListImportsOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListImportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListImportsWithContext(ctx context.Context, input *dynamodb.ListImportsInput, opts ...request.Option) (*dynamodb.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListImportsOutput), nil
	}
	out, err := c.DynamoDBAPI.ListImportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListTables(input *dynamodb.ListTablesInput) (*dynamodb.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListTablesOutput), nil
	}
	out, err := c.DynamoDBAPI.ListTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListTablesPages(input *dynamodb.ListTablesInput, fn func(*dynamodb.ListTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListTablesOutput{}
	fnCacher := func(out *dynamodb.ListTablesOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListTablesPagesWithContext(ctx context.Context, input *dynamodb.ListTablesInput, fn func(*dynamodb.ListTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ListTablesOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ListTablesOutput{}
	fnCacher := func(out *dynamodb.ListTablesOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ListTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ListTablesWithContext(ctx context.Context, input *dynamodb.ListTablesInput, opts ...request.Option) (*dynamodb.ListTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListTablesOutput), nil
	}
	out, err := c.DynamoDBAPI.ListTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListTagsOfResource(input *dynamodb.ListTagsOfResourceInput) (*dynamodb.ListTagsOfResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListTagsOfResourceOutput), nil
	}
	out, err := c.DynamoDBAPI.ListTagsOfResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ListTagsOfResourceWithContext(ctx context.Context, input *dynamodb.ListTagsOfResourceInput, opts ...request.Option) (*dynamodb.ListTagsOfResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ListTagsOfResourceOutput), nil
	}
	out, err := c.DynamoDBAPI.ListTagsOfResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) Query(input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.QueryOutput), nil
	}
	out, err := c.DynamoDBAPI.Query(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) QueryPages(input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.QueryOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.QueryOutput{}
	fnCacher := func(out *dynamodb.QueryOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.QueryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) QueryPagesWithContext(ctx context.Context, input *dynamodb.QueryInput, fn func(*dynamodb.QueryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.QueryOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.QueryOutput{}
	fnCacher := func(out *dynamodb.QueryOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.QueryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) QueryWithContext(ctx context.Context, input *dynamodb.QueryInput, opts ...request.Option) (*dynamodb.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.QueryOutput), nil
	}
	out, err := c.DynamoDBAPI.QueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ScanOutput), nil
	}
	out, err := c.DynamoDBAPI.Scan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DynamoDB) ScanPages(input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ScanOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ScanOutput{}
	fnCacher := func(out *dynamodb.ScanOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ScanPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ScanPagesWithContext(ctx context.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*dynamodb.ScanOutput), false)
		return nil
	}
	cachable := true
	output := &dynamodb.ScanOutput{}
	fnCacher := func(out *dynamodb.ScanOutput, more bool) bool {
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
	if err := c.DynamoDBAPI.ScanPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DynamoDB) ScanWithContext(ctx context.Context, input *dynamodb.ScanInput, opts ...request.Option) (*dynamodb.ScanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*dynamodb.ScanOutput), nil
	}
	out, err := c.DynamoDBAPI.ScanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
