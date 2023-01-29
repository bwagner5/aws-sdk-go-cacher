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
	"github.com/aws/aws-sdk-go/service/personalize"
	"github.com/aws/aws-sdk-go/service/personalize/personalizeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Personalize struct {
	personalizeiface.PersonalizeAPI
	cache *cache.Cache
}

func NewPersonalize(personalizeapi personalizeiface.PersonalizeAPI) *Personalize {
	return &Personalize{
		PersonalizeAPI: personalizeapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Personalize) DescribeAlgorithm(input *personalize.DescribeAlgorithmInput) (*personalize.DescribeAlgorithmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeAlgorithmOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeAlgorithm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeAlgorithmWithContext(ctx context.Context, input *personalize.DescribeAlgorithmInput, opts ...request.Option) (*personalize.DescribeAlgorithmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeAlgorithmOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeAlgorithmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeBatchInferenceJob(input *personalize.DescribeBatchInferenceJobInput) (*personalize.DescribeBatchInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeBatchInferenceJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeBatchInferenceJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeBatchInferenceJobWithContext(ctx context.Context, input *personalize.DescribeBatchInferenceJobInput, opts ...request.Option) (*personalize.DescribeBatchInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeBatchInferenceJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeBatchInferenceJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeBatchSegmentJob(input *personalize.DescribeBatchSegmentJobInput) (*personalize.DescribeBatchSegmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeBatchSegmentJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeBatchSegmentJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeBatchSegmentJobWithContext(ctx context.Context, input *personalize.DescribeBatchSegmentJobInput, opts ...request.Option) (*personalize.DescribeBatchSegmentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeBatchSegmentJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeBatchSegmentJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeCampaign(input *personalize.DescribeCampaignInput) (*personalize.DescribeCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeCampaignOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeCampaignWithContext(ctx context.Context, input *personalize.DescribeCampaignInput, opts ...request.Option) (*personalize.DescribeCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeCampaignOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDataset(input *personalize.DescribeDatasetInput) (*personalize.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetExportJob(input *personalize.DescribeDatasetExportJobInput) (*personalize.DescribeDatasetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetExportJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetExportJobWithContext(ctx context.Context, input *personalize.DescribeDatasetExportJobInput, opts ...request.Option) (*personalize.DescribeDatasetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetExportJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetGroup(input *personalize.DescribeDatasetGroupInput) (*personalize.DescribeDatasetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetGroupOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetGroupWithContext(ctx context.Context, input *personalize.DescribeDatasetGroupInput, opts ...request.Option) (*personalize.DescribeDatasetGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetGroupOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetImportJob(input *personalize.DescribeDatasetImportJobInput) (*personalize.DescribeDatasetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetImportJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetImportJobWithContext(ctx context.Context, input *personalize.DescribeDatasetImportJobInput, opts ...request.Option) (*personalize.DescribeDatasetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetImportJobOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeDatasetWithContext(ctx context.Context, input *personalize.DescribeDatasetInput, opts ...request.Option) (*personalize.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeDatasetOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeEventTracker(input *personalize.DescribeEventTrackerInput) (*personalize.DescribeEventTrackerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeEventTrackerOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeEventTracker(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeEventTrackerWithContext(ctx context.Context, input *personalize.DescribeEventTrackerInput, opts ...request.Option) (*personalize.DescribeEventTrackerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeEventTrackerOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeEventTrackerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeFeatureTransformation(input *personalize.DescribeFeatureTransformationInput) (*personalize.DescribeFeatureTransformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeFeatureTransformationOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeFeatureTransformation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeFeatureTransformationWithContext(ctx context.Context, input *personalize.DescribeFeatureTransformationInput, opts ...request.Option) (*personalize.DescribeFeatureTransformationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeFeatureTransformationOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeFeatureTransformationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeFilter(input *personalize.DescribeFilterInput) (*personalize.DescribeFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeFilterOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeFilterWithContext(ctx context.Context, input *personalize.DescribeFilterInput, opts ...request.Option) (*personalize.DescribeFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeFilterOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeMetricAttribution(input *personalize.DescribeMetricAttributionInput) (*personalize.DescribeMetricAttributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeMetricAttributionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeMetricAttribution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeMetricAttributionWithContext(ctx context.Context, input *personalize.DescribeMetricAttributionInput, opts ...request.Option) (*personalize.DescribeMetricAttributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeMetricAttributionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeMetricAttributionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeRecipe(input *personalize.DescribeRecipeInput) (*personalize.DescribeRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeRecipeOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeRecipe(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeRecipeWithContext(ctx context.Context, input *personalize.DescribeRecipeInput, opts ...request.Option) (*personalize.DescribeRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeRecipeOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeRecipeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeRecommender(input *personalize.DescribeRecommenderInput) (*personalize.DescribeRecommenderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeRecommenderOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeRecommender(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeRecommenderWithContext(ctx context.Context, input *personalize.DescribeRecommenderInput, opts ...request.Option) (*personalize.DescribeRecommenderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeRecommenderOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeRecommenderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSchema(input *personalize.DescribeSchemaInput) (*personalize.DescribeSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSchemaOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSchemaWithContext(ctx context.Context, input *personalize.DescribeSchemaInput, opts ...request.Option) (*personalize.DescribeSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSchemaOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSolution(input *personalize.DescribeSolutionInput) (*personalize.DescribeSolutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSolutionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSolution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSolutionVersion(input *personalize.DescribeSolutionVersionInput) (*personalize.DescribeSolutionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSolutionVersionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSolutionVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSolutionVersionWithContext(ctx context.Context, input *personalize.DescribeSolutionVersionInput, opts ...request.Option) (*personalize.DescribeSolutionVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSolutionVersionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSolutionVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) DescribeSolutionWithContext(ctx context.Context, input *personalize.DescribeSolutionInput, opts ...request.Option) (*personalize.DescribeSolutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.DescribeSolutionOutput), nil
	}
	out, err := c.PersonalizeAPI.DescribeSolutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) GetSolutionMetrics(input *personalize.GetSolutionMetricsInput) (*personalize.GetSolutionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.GetSolutionMetricsOutput), nil
	}
	out, err := c.PersonalizeAPI.GetSolutionMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) GetSolutionMetricsWithContext(ctx context.Context, input *personalize.GetSolutionMetricsInput, opts ...request.Option) (*personalize.GetSolutionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.GetSolutionMetricsOutput), nil
	}
	out, err := c.PersonalizeAPI.GetSolutionMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListBatchInferenceJobs(input *personalize.ListBatchInferenceJobsInput) (*personalize.ListBatchInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListBatchInferenceJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListBatchInferenceJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListBatchInferenceJobsPages(input *personalize.ListBatchInferenceJobsInput, fn func(*personalize.ListBatchInferenceJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListBatchInferenceJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListBatchInferenceJobsOutput{}
	fnCacher := func(out *personalize.ListBatchInferenceJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListBatchInferenceJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListBatchInferenceJobsPagesWithContext(ctx context.Context, input *personalize.ListBatchInferenceJobsInput, fn func(*personalize.ListBatchInferenceJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListBatchInferenceJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListBatchInferenceJobsOutput{}
	fnCacher := func(out *personalize.ListBatchInferenceJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListBatchInferenceJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListBatchInferenceJobsWithContext(ctx context.Context, input *personalize.ListBatchInferenceJobsInput, opts ...request.Option) (*personalize.ListBatchInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListBatchInferenceJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListBatchInferenceJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListBatchSegmentJobs(input *personalize.ListBatchSegmentJobsInput) (*personalize.ListBatchSegmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListBatchSegmentJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListBatchSegmentJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListBatchSegmentJobsPages(input *personalize.ListBatchSegmentJobsInput, fn func(*personalize.ListBatchSegmentJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListBatchSegmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListBatchSegmentJobsOutput{}
	fnCacher := func(out *personalize.ListBatchSegmentJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListBatchSegmentJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListBatchSegmentJobsPagesWithContext(ctx context.Context, input *personalize.ListBatchSegmentJobsInput, fn func(*personalize.ListBatchSegmentJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListBatchSegmentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListBatchSegmentJobsOutput{}
	fnCacher := func(out *personalize.ListBatchSegmentJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListBatchSegmentJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListBatchSegmentJobsWithContext(ctx context.Context, input *personalize.ListBatchSegmentJobsInput, opts ...request.Option) (*personalize.ListBatchSegmentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListBatchSegmentJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListBatchSegmentJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListCampaigns(input *personalize.ListCampaignsInput) (*personalize.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListCampaignsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListCampaignsPages(input *personalize.ListCampaignsInput, fn func(*personalize.ListCampaignsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListCampaignsOutput{}
	fnCacher := func(out *personalize.ListCampaignsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListCampaignsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListCampaignsPagesWithContext(ctx context.Context, input *personalize.ListCampaignsInput, fn func(*personalize.ListCampaignsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListCampaignsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListCampaignsOutput{}
	fnCacher := func(out *personalize.ListCampaignsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListCampaignsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListCampaignsWithContext(ctx context.Context, input *personalize.ListCampaignsInput, opts ...request.Option) (*personalize.ListCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListCampaignsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetExportJobs(input *personalize.ListDatasetExportJobsInput) (*personalize.ListDatasetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetExportJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetExportJobsPages(input *personalize.ListDatasetExportJobsInput, fn func(*personalize.ListDatasetExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetExportJobsOutput{}
	fnCacher := func(out *personalize.ListDatasetExportJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetExportJobsPagesWithContext(ctx context.Context, input *personalize.ListDatasetExportJobsInput, fn func(*personalize.ListDatasetExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetExportJobsOutput{}
	fnCacher := func(out *personalize.ListDatasetExportJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetExportJobsWithContext(ctx context.Context, input *personalize.ListDatasetExportJobsInput, opts ...request.Option) (*personalize.ListDatasetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetExportJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetGroups(input *personalize.ListDatasetGroupsInput) (*personalize.ListDatasetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetGroupsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetGroupsPages(input *personalize.ListDatasetGroupsInput, fn func(*personalize.ListDatasetGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetGroupsOutput{}
	fnCacher := func(out *personalize.ListDatasetGroupsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetGroupsPagesWithContext(ctx context.Context, input *personalize.ListDatasetGroupsInput, fn func(*personalize.ListDatasetGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetGroupsOutput{}
	fnCacher := func(out *personalize.ListDatasetGroupsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetGroupsWithContext(ctx context.Context, input *personalize.ListDatasetGroupsInput, opts ...request.Option) (*personalize.ListDatasetGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetGroupsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetImportJobs(input *personalize.ListDatasetImportJobsInput) (*personalize.ListDatasetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetImportJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetImportJobsPages(input *personalize.ListDatasetImportJobsInput, fn func(*personalize.ListDatasetImportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetImportJobsOutput{}
	fnCacher := func(out *personalize.ListDatasetImportJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetImportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetImportJobsPagesWithContext(ctx context.Context, input *personalize.ListDatasetImportJobsInput, fn func(*personalize.ListDatasetImportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetImportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetImportJobsOutput{}
	fnCacher := func(out *personalize.ListDatasetImportJobsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetImportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetImportJobsWithContext(ctx context.Context, input *personalize.ListDatasetImportJobsInput, opts ...request.Option) (*personalize.ListDatasetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetImportJobsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasets(input *personalize.ListDatasetsInput) (*personalize.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListDatasetsPages(input *personalize.ListDatasetsInput, fn func(*personalize.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetsOutput{}
	fnCacher := func(out *personalize.ListDatasetsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetsPagesWithContext(ctx context.Context, input *personalize.ListDatasetsInput, fn func(*personalize.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListDatasetsOutput{}
	fnCacher := func(out *personalize.ListDatasetsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListDatasetsWithContext(ctx context.Context, input *personalize.ListDatasetsInput, opts ...request.Option) (*personalize.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListDatasetsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListEventTrackers(input *personalize.ListEventTrackersInput) (*personalize.ListEventTrackersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListEventTrackersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListEventTrackers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListEventTrackersPages(input *personalize.ListEventTrackersInput, fn func(*personalize.ListEventTrackersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListEventTrackersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListEventTrackersOutput{}
	fnCacher := func(out *personalize.ListEventTrackersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListEventTrackersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListEventTrackersPagesWithContext(ctx context.Context, input *personalize.ListEventTrackersInput, fn func(*personalize.ListEventTrackersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListEventTrackersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListEventTrackersOutput{}
	fnCacher := func(out *personalize.ListEventTrackersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListEventTrackersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListEventTrackersWithContext(ctx context.Context, input *personalize.ListEventTrackersInput, opts ...request.Option) (*personalize.ListEventTrackersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListEventTrackersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListEventTrackersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListFilters(input *personalize.ListFiltersInput) (*personalize.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListFiltersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListFiltersPages(input *personalize.ListFiltersInput, fn func(*personalize.ListFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListFiltersOutput{}
	fnCacher := func(out *personalize.ListFiltersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListFiltersPagesWithContext(ctx context.Context, input *personalize.ListFiltersInput, fn func(*personalize.ListFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListFiltersOutput{}
	fnCacher := func(out *personalize.ListFiltersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListFiltersWithContext(ctx context.Context, input *personalize.ListFiltersInput, opts ...request.Option) (*personalize.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListFiltersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListMetricAttributionMetrics(input *personalize.ListMetricAttributionMetricsInput) (*personalize.ListMetricAttributionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListMetricAttributionMetricsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListMetricAttributionMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListMetricAttributionMetricsPages(input *personalize.ListMetricAttributionMetricsInput, fn func(*personalize.ListMetricAttributionMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListMetricAttributionMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListMetricAttributionMetricsOutput{}
	fnCacher := func(out *personalize.ListMetricAttributionMetricsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListMetricAttributionMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListMetricAttributionMetricsPagesWithContext(ctx context.Context, input *personalize.ListMetricAttributionMetricsInput, fn func(*personalize.ListMetricAttributionMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListMetricAttributionMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListMetricAttributionMetricsOutput{}
	fnCacher := func(out *personalize.ListMetricAttributionMetricsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListMetricAttributionMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListMetricAttributionMetricsWithContext(ctx context.Context, input *personalize.ListMetricAttributionMetricsInput, opts ...request.Option) (*personalize.ListMetricAttributionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListMetricAttributionMetricsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListMetricAttributionMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListMetricAttributions(input *personalize.ListMetricAttributionsInput) (*personalize.ListMetricAttributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListMetricAttributionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListMetricAttributions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListMetricAttributionsPages(input *personalize.ListMetricAttributionsInput, fn func(*personalize.ListMetricAttributionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListMetricAttributionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListMetricAttributionsOutput{}
	fnCacher := func(out *personalize.ListMetricAttributionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListMetricAttributionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListMetricAttributionsPagesWithContext(ctx context.Context, input *personalize.ListMetricAttributionsInput, fn func(*personalize.ListMetricAttributionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListMetricAttributionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListMetricAttributionsOutput{}
	fnCacher := func(out *personalize.ListMetricAttributionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListMetricAttributionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListMetricAttributionsWithContext(ctx context.Context, input *personalize.ListMetricAttributionsInput, opts ...request.Option) (*personalize.ListMetricAttributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListMetricAttributionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListMetricAttributionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListRecipes(input *personalize.ListRecipesInput) (*personalize.ListRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListRecipesOutput), nil
	}
	out, err := c.PersonalizeAPI.ListRecipes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListRecipesPages(input *personalize.ListRecipesInput, fn func(*personalize.ListRecipesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListRecipesOutput{}
	fnCacher := func(out *personalize.ListRecipesOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListRecipesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListRecipesPagesWithContext(ctx context.Context, input *personalize.ListRecipesInput, fn func(*personalize.ListRecipesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListRecipesOutput{}
	fnCacher := func(out *personalize.ListRecipesOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListRecipesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListRecipesWithContext(ctx context.Context, input *personalize.ListRecipesInput, opts ...request.Option) (*personalize.ListRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListRecipesOutput), nil
	}
	out, err := c.PersonalizeAPI.ListRecipesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListRecommenders(input *personalize.ListRecommendersInput) (*personalize.ListRecommendersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListRecommendersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListRecommenders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListRecommendersPages(input *personalize.ListRecommendersInput, fn func(*personalize.ListRecommendersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListRecommendersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListRecommendersOutput{}
	fnCacher := func(out *personalize.ListRecommendersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListRecommendersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListRecommendersPagesWithContext(ctx context.Context, input *personalize.ListRecommendersInput, fn func(*personalize.ListRecommendersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListRecommendersOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListRecommendersOutput{}
	fnCacher := func(out *personalize.ListRecommendersOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListRecommendersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListRecommendersWithContext(ctx context.Context, input *personalize.ListRecommendersInput, opts ...request.Option) (*personalize.ListRecommendersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListRecommendersOutput), nil
	}
	out, err := c.PersonalizeAPI.ListRecommendersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSchemas(input *personalize.ListSchemasInput) (*personalize.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSchemasOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSchemasPages(input *personalize.ListSchemasInput, fn func(*personalize.ListSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSchemasOutput{}
	fnCacher := func(out *personalize.ListSchemasOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSchemasPagesWithContext(ctx context.Context, input *personalize.ListSchemasInput, fn func(*personalize.ListSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSchemasOutput{}
	fnCacher := func(out *personalize.ListSchemasOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSchemasWithContext(ctx context.Context, input *personalize.ListSchemasInput, opts ...request.Option) (*personalize.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSchemasOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSolutionVersions(input *personalize.ListSolutionVersionsInput) (*personalize.ListSolutionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSolutionVersionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSolutionVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSolutionVersionsPages(input *personalize.ListSolutionVersionsInput, fn func(*personalize.ListSolutionVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSolutionVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSolutionVersionsOutput{}
	fnCacher := func(out *personalize.ListSolutionVersionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSolutionVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSolutionVersionsPagesWithContext(ctx context.Context, input *personalize.ListSolutionVersionsInput, fn func(*personalize.ListSolutionVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSolutionVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSolutionVersionsOutput{}
	fnCacher := func(out *personalize.ListSolutionVersionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSolutionVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSolutionVersionsWithContext(ctx context.Context, input *personalize.ListSolutionVersionsInput, opts ...request.Option) (*personalize.ListSolutionVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSolutionVersionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSolutionVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSolutions(input *personalize.ListSolutionsInput) (*personalize.ListSolutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSolutionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSolutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListSolutionsPages(input *personalize.ListSolutionsInput, fn func(*personalize.ListSolutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSolutionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSolutionsOutput{}
	fnCacher := func(out *personalize.ListSolutionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSolutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSolutionsPagesWithContext(ctx context.Context, input *personalize.ListSolutionsInput, fn func(*personalize.ListSolutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*personalize.ListSolutionsOutput), false)
		return nil
	}
	cachable := true
	output := &personalize.ListSolutionsOutput{}
	fnCacher := func(out *personalize.ListSolutionsOutput, more bool) bool {
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
	if err := c.PersonalizeAPI.ListSolutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Personalize) ListSolutionsWithContext(ctx context.Context, input *personalize.ListSolutionsInput, opts ...request.Option) (*personalize.ListSolutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListSolutionsOutput), nil
	}
	out, err := c.PersonalizeAPI.ListSolutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListTagsForResource(input *personalize.ListTagsForResourceInput) (*personalize.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListTagsForResourceOutput), nil
	}
	out, err := c.PersonalizeAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Personalize) ListTagsForResourceWithContext(ctx context.Context, input *personalize.ListTagsForResourceInput, opts ...request.Option) (*personalize.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*personalize.ListTagsForResourceOutput), nil
	}
	out, err := c.PersonalizeAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
