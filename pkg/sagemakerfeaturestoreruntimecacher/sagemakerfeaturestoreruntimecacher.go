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
package sagemakerfeaturestoreruntimecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime"
	"github.com/aws/aws-sdk-go/service/sagemakerfeaturestoreruntime/sagemakerfeaturestoreruntimeiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SageMakerFeatureStoreRuntime struct {
	sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI
	cache *cache.Cache
}

func New(sagemakerfeaturestoreruntimeapi sagemakerfeaturestoreruntimeiface.SageMakerFeatureStoreRuntimeAPI) *SageMakerFeatureStoreRuntime {
	return &SageMakerFeatureStoreRuntime{
		SageMakerFeatureStoreRuntimeAPI: sagemakerfeaturestoreruntimeapi,
		cache:                           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SageMakerFeatureStoreRuntime) BatchGetRecord(input *sagemakerfeaturestoreruntime.BatchGetRecordInput) (*sagemakerfeaturestoreruntime.BatchGetRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakerfeaturestoreruntime.BatchGetRecordOutput), nil
	}
	out, err := c.SageMakerFeatureStoreRuntimeAPI.BatchGetRecord(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerFeatureStoreRuntime) BatchGetRecordWithContext(ctx context.Context, input *sagemakerfeaturestoreruntime.BatchGetRecordInput, opts ...request.Option) (*sagemakerfeaturestoreruntime.BatchGetRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakerfeaturestoreruntime.BatchGetRecordOutput), nil
	}
	out, err := c.SageMakerFeatureStoreRuntimeAPI.BatchGetRecordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerFeatureStoreRuntime) GetRecord(input *sagemakerfeaturestoreruntime.GetRecordInput) (*sagemakerfeaturestoreruntime.GetRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakerfeaturestoreruntime.GetRecordOutput), nil
	}
	out, err := c.SageMakerFeatureStoreRuntimeAPI.GetRecord(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMakerFeatureStoreRuntime) GetRecordWithContext(ctx context.Context, input *sagemakerfeaturestoreruntime.GetRecordInput, opts ...request.Option) (*sagemakerfeaturestoreruntime.GetRecordOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemakerfeaturestoreruntime.GetRecordOutput), nil
	}
	out, err := c.SageMakerFeatureStoreRuntimeAPI.GetRecordWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
