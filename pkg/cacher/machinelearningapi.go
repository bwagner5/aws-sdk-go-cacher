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
	"github.com/aws/aws-sdk-go/service/machinelearning"
	"github.com/aws/aws-sdk-go/service/machinelearning/machinelearningiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MachineLearning struct {
	machinelearningiface.MachineLearningAPI
	cache *cache.Cache
}

func NewMachineLearning(machinelearningapi machinelearningiface.MachineLearningAPI) *MachineLearning {
	return &MachineLearning{
		MachineLearningAPI: machinelearningapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MachineLearning) DescribeBatchPredictions(input *machinelearning.DescribeBatchPredictionsInput) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeBatchPredictionsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeBatchPredictions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeBatchPredictionsPages(input *machinelearning.DescribeBatchPredictionsInput, fn func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeBatchPredictionsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeBatchPredictionsOutput{}
	fnCacher := func(out *machinelearning.DescribeBatchPredictionsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeBatchPredictionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeBatchPredictionsPagesWithContext(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput, fn func(*machinelearning.DescribeBatchPredictionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeBatchPredictionsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeBatchPredictionsOutput{}
	fnCacher := func(out *machinelearning.DescribeBatchPredictionsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeBatchPredictionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeBatchPredictionsWithContext(ctx context.Context, input *machinelearning.DescribeBatchPredictionsInput, opts ...request.Option) (*machinelearning.DescribeBatchPredictionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeBatchPredictionsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeBatchPredictionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeDataSources(input *machinelearning.DescribeDataSourcesInput) (*machinelearning.DescribeDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeDataSourcesOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeDataSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeDataSourcesPages(input *machinelearning.DescribeDataSourcesInput, fn func(*machinelearning.DescribeDataSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeDataSourcesOutput{}
	fnCacher := func(out *machinelearning.DescribeDataSourcesOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeDataSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeDataSourcesPagesWithContext(ctx context.Context, input *machinelearning.DescribeDataSourcesInput, fn func(*machinelearning.DescribeDataSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeDataSourcesOutput{}
	fnCacher := func(out *machinelearning.DescribeDataSourcesOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeDataSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeDataSourcesWithContext(ctx context.Context, input *machinelearning.DescribeDataSourcesInput, opts ...request.Option) (*machinelearning.DescribeDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeDataSourcesOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeDataSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeEvaluations(input *machinelearning.DescribeEvaluationsInput) (*machinelearning.DescribeEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeEvaluationsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeEvaluations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeEvaluationsPages(input *machinelearning.DescribeEvaluationsInput, fn func(*machinelearning.DescribeEvaluationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeEvaluationsOutput{}
	fnCacher := func(out *machinelearning.DescribeEvaluationsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeEvaluationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeEvaluationsPagesWithContext(ctx context.Context, input *machinelearning.DescribeEvaluationsInput, fn func(*machinelearning.DescribeEvaluationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeEvaluationsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeEvaluationsOutput{}
	fnCacher := func(out *machinelearning.DescribeEvaluationsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeEvaluationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeEvaluationsWithContext(ctx context.Context, input *machinelearning.DescribeEvaluationsInput, opts ...request.Option) (*machinelearning.DescribeEvaluationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeEvaluationsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeEvaluationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeMLModels(input *machinelearning.DescribeMLModelsInput) (*machinelearning.DescribeMLModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeMLModelsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeMLModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeMLModelsPages(input *machinelearning.DescribeMLModelsInput, fn func(*machinelearning.DescribeMLModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeMLModelsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeMLModelsOutput{}
	fnCacher := func(out *machinelearning.DescribeMLModelsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeMLModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeMLModelsPagesWithContext(ctx context.Context, input *machinelearning.DescribeMLModelsInput, fn func(*machinelearning.DescribeMLModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*machinelearning.DescribeMLModelsOutput), false)
		return nil
	}
	cachable := true
	output := &machinelearning.DescribeMLModelsOutput{}
	fnCacher := func(out *machinelearning.DescribeMLModelsOutput, more bool) bool {
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
	if err := c.MachineLearningAPI.DescribeMLModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MachineLearning) DescribeMLModelsWithContext(ctx context.Context, input *machinelearning.DescribeMLModelsInput, opts ...request.Option) (*machinelearning.DescribeMLModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeMLModelsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeMLModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeTags(input *machinelearning.DescribeTagsInput) (*machinelearning.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeTagsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) DescribeTagsWithContext(ctx context.Context, input *machinelearning.DescribeTagsInput, opts ...request.Option) (*machinelearning.DescribeTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.DescribeTagsOutput), nil
	}
	out, err := c.MachineLearningAPI.DescribeTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetBatchPrediction(input *machinelearning.GetBatchPredictionInput) (*machinelearning.GetBatchPredictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetBatchPredictionOutput), nil
	}
	out, err := c.MachineLearningAPI.GetBatchPrediction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetBatchPredictionWithContext(ctx context.Context, input *machinelearning.GetBatchPredictionInput, opts ...request.Option) (*machinelearning.GetBatchPredictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetBatchPredictionOutput), nil
	}
	out, err := c.MachineLearningAPI.GetBatchPredictionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetDataSource(input *machinelearning.GetDataSourceInput) (*machinelearning.GetDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetDataSourceOutput), nil
	}
	out, err := c.MachineLearningAPI.GetDataSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetDataSourceWithContext(ctx context.Context, input *machinelearning.GetDataSourceInput, opts ...request.Option) (*machinelearning.GetDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetDataSourceOutput), nil
	}
	out, err := c.MachineLearningAPI.GetDataSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetEvaluation(input *machinelearning.GetEvaluationInput) (*machinelearning.GetEvaluationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetEvaluationOutput), nil
	}
	out, err := c.MachineLearningAPI.GetEvaluation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetEvaluationWithContext(ctx context.Context, input *machinelearning.GetEvaluationInput, opts ...request.Option) (*machinelearning.GetEvaluationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetEvaluationOutput), nil
	}
	out, err := c.MachineLearningAPI.GetEvaluationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetMLModel(input *machinelearning.GetMLModelInput) (*machinelearning.GetMLModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetMLModelOutput), nil
	}
	out, err := c.MachineLearningAPI.GetMLModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MachineLearning) GetMLModelWithContext(ctx context.Context, input *machinelearning.GetMLModelInput, opts ...request.Option) (*machinelearning.GetMLModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*machinelearning.GetMLModelOutput), nil
	}
	out, err := c.MachineLearningAPI.GetMLModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
