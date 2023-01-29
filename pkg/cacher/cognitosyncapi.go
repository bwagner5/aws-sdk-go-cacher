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
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"github.com/aws/aws-sdk-go/service/cognitosync/cognitosynciface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CognitoSync struct {
	cognitosynciface.CognitoSyncAPI
	cache *cache.Cache
}

func NewCognitoSync(cognitosyncapi cognitosynciface.CognitoSyncAPI) *CognitoSync {
	return &CognitoSync{
		CognitoSyncAPI: cognitosyncapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CognitoSync) DescribeDataset(input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeDatasetOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) DescribeDatasetWithContext(ctx context.Context, input *cognitosync.DescribeDatasetInput, opts ...request.Option) (*cognitosync.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeDatasetOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) DescribeIdentityPoolUsage(input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeIdentityPoolUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeIdentityPoolUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) DescribeIdentityPoolUsageWithContext(ctx context.Context, input *cognitosync.DescribeIdentityPoolUsageInput, opts ...request.Option) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeIdentityPoolUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeIdentityPoolUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) DescribeIdentityUsage(input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeIdentityUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeIdentityUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) DescribeIdentityUsageWithContext(ctx context.Context, input *cognitosync.DescribeIdentityUsageInput, opts ...request.Option) (*cognitosync.DescribeIdentityUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.DescribeIdentityUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.DescribeIdentityUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetBulkPublishDetails(input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetBulkPublishDetailsOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetBulkPublishDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetBulkPublishDetailsWithContext(ctx context.Context, input *cognitosync.GetBulkPublishDetailsInput, opts ...request.Option) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetBulkPublishDetailsOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetBulkPublishDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetCognitoEvents(input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetCognitoEventsOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetCognitoEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetCognitoEventsWithContext(ctx context.Context, input *cognitosync.GetCognitoEventsInput, opts ...request.Option) (*cognitosync.GetCognitoEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetCognitoEventsOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetCognitoEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetIdentityPoolConfiguration(input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetIdentityPoolConfigurationOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetIdentityPoolConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) GetIdentityPoolConfigurationWithContext(ctx context.Context, input *cognitosync.GetIdentityPoolConfigurationInput, opts ...request.Option) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.GetIdentityPoolConfigurationOutput), nil
	}
	out, err := c.CognitoSyncAPI.GetIdentityPoolConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListDatasets(input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListDatasetsOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListDatasetsWithContext(ctx context.Context, input *cognitosync.ListDatasetsInput, opts ...request.Option) (*cognitosync.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListDatasetsOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListIdentityPoolUsage(input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListIdentityPoolUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListIdentityPoolUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListIdentityPoolUsageWithContext(ctx context.Context, input *cognitosync.ListIdentityPoolUsageInput, opts ...request.Option) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListIdentityPoolUsageOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListIdentityPoolUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListRecords(input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListRecordsOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListRecords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CognitoSync) ListRecordsWithContext(ctx context.Context, input *cognitosync.ListRecordsInput, opts ...request.Option) (*cognitosync.ListRecordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cognitosync.ListRecordsOutput), nil
	}
	out, err := c.CognitoSyncAPI.ListRecordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
