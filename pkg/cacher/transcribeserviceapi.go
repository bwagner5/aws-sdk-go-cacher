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
	"github.com/aws/aws-sdk-go/service/transcribeservice"
	"github.com/aws/aws-sdk-go/service/transcribeservice/transcribeserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type TranscribeService struct {
	transcribeserviceiface.TranscribeServiceAPI
	cache *cache.Cache
}

func NewTranscribeService(transcribeserviceapi transcribeserviceiface.TranscribeServiceAPI) *TranscribeService {
	return &TranscribeService{
		TranscribeServiceAPI: transcribeserviceapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *TranscribeService) DescribeLanguageModel(input *transcribeservice.DescribeLanguageModelInput) (*transcribeservice.DescribeLanguageModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.DescribeLanguageModelOutput), nil
	}
	out, err := c.TranscribeServiceAPI.DescribeLanguageModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) DescribeLanguageModelWithContext(ctx context.Context, input *transcribeservice.DescribeLanguageModelInput, opts ...request.Option) (*transcribeservice.DescribeLanguageModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.DescribeLanguageModelOutput), nil
	}
	out, err := c.TranscribeServiceAPI.DescribeLanguageModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetCallAnalyticsCategory(input *transcribeservice.GetCallAnalyticsCategoryInput) (*transcribeservice.GetCallAnalyticsCategoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetCallAnalyticsCategoryOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetCallAnalyticsCategory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetCallAnalyticsCategoryWithContext(ctx context.Context, input *transcribeservice.GetCallAnalyticsCategoryInput, opts ...request.Option) (*transcribeservice.GetCallAnalyticsCategoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetCallAnalyticsCategoryOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetCallAnalyticsCategoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetCallAnalyticsJob(input *transcribeservice.GetCallAnalyticsJobInput) (*transcribeservice.GetCallAnalyticsJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetCallAnalyticsJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetCallAnalyticsJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetCallAnalyticsJobWithContext(ctx context.Context, input *transcribeservice.GetCallAnalyticsJobInput, opts ...request.Option) (*transcribeservice.GetCallAnalyticsJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetCallAnalyticsJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetCallAnalyticsJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetMedicalTranscriptionJob(input *transcribeservice.GetMedicalTranscriptionJobInput) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetMedicalTranscriptionJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetMedicalTranscriptionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetMedicalTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.GetMedicalTranscriptionJobInput, opts ...request.Option) (*transcribeservice.GetMedicalTranscriptionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetMedicalTranscriptionJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetMedicalTranscriptionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetMedicalVocabulary(input *transcribeservice.GetMedicalVocabularyInput) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetMedicalVocabularyOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetMedicalVocabulary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetMedicalVocabularyWithContext(ctx context.Context, input *transcribeservice.GetMedicalVocabularyInput, opts ...request.Option) (*transcribeservice.GetMedicalVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetMedicalVocabularyOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetMedicalVocabularyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetTranscriptionJob(input *transcribeservice.GetTranscriptionJobInput) (*transcribeservice.GetTranscriptionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetTranscriptionJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetTranscriptionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetTranscriptionJobWithContext(ctx context.Context, input *transcribeservice.GetTranscriptionJobInput, opts ...request.Option) (*transcribeservice.GetTranscriptionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetTranscriptionJobOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetTranscriptionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetVocabulary(input *transcribeservice.GetVocabularyInput) (*transcribeservice.GetVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetVocabularyOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetVocabulary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetVocabularyFilter(input *transcribeservice.GetVocabularyFilterInput) (*transcribeservice.GetVocabularyFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetVocabularyFilterOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetVocabularyFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetVocabularyFilterWithContext(ctx context.Context, input *transcribeservice.GetVocabularyFilterInput, opts ...request.Option) (*transcribeservice.GetVocabularyFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetVocabularyFilterOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetVocabularyFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) GetVocabularyWithContext(ctx context.Context, input *transcribeservice.GetVocabularyInput, opts ...request.Option) (*transcribeservice.GetVocabularyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.GetVocabularyOutput), nil
	}
	out, err := c.TranscribeServiceAPI.GetVocabularyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListCallAnalyticsCategories(input *transcribeservice.ListCallAnalyticsCategoriesInput) (*transcribeservice.ListCallAnalyticsCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListCallAnalyticsCategoriesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListCallAnalyticsCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListCallAnalyticsCategoriesPages(input *transcribeservice.ListCallAnalyticsCategoriesInput, fn func(*transcribeservice.ListCallAnalyticsCategoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListCallAnalyticsCategoriesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListCallAnalyticsCategoriesOutput{}
	fnCacher := func(out *transcribeservice.ListCallAnalyticsCategoriesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListCallAnalyticsCategoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListCallAnalyticsCategoriesPagesWithContext(ctx context.Context, input *transcribeservice.ListCallAnalyticsCategoriesInput, fn func(*transcribeservice.ListCallAnalyticsCategoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListCallAnalyticsCategoriesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListCallAnalyticsCategoriesOutput{}
	fnCacher := func(out *transcribeservice.ListCallAnalyticsCategoriesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListCallAnalyticsCategoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListCallAnalyticsCategoriesWithContext(ctx context.Context, input *transcribeservice.ListCallAnalyticsCategoriesInput, opts ...request.Option) (*transcribeservice.ListCallAnalyticsCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListCallAnalyticsCategoriesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListCallAnalyticsCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListCallAnalyticsJobs(input *transcribeservice.ListCallAnalyticsJobsInput) (*transcribeservice.ListCallAnalyticsJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListCallAnalyticsJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListCallAnalyticsJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListCallAnalyticsJobsPages(input *transcribeservice.ListCallAnalyticsJobsInput, fn func(*transcribeservice.ListCallAnalyticsJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListCallAnalyticsJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListCallAnalyticsJobsOutput{}
	fnCacher := func(out *transcribeservice.ListCallAnalyticsJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListCallAnalyticsJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListCallAnalyticsJobsPagesWithContext(ctx context.Context, input *transcribeservice.ListCallAnalyticsJobsInput, fn func(*transcribeservice.ListCallAnalyticsJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListCallAnalyticsJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListCallAnalyticsJobsOutput{}
	fnCacher := func(out *transcribeservice.ListCallAnalyticsJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListCallAnalyticsJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListCallAnalyticsJobsWithContext(ctx context.Context, input *transcribeservice.ListCallAnalyticsJobsInput, opts ...request.Option) (*transcribeservice.ListCallAnalyticsJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListCallAnalyticsJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListCallAnalyticsJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListLanguageModels(input *transcribeservice.ListLanguageModelsInput) (*transcribeservice.ListLanguageModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListLanguageModelsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListLanguageModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListLanguageModelsPages(input *transcribeservice.ListLanguageModelsInput, fn func(*transcribeservice.ListLanguageModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListLanguageModelsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListLanguageModelsOutput{}
	fnCacher := func(out *transcribeservice.ListLanguageModelsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListLanguageModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListLanguageModelsPagesWithContext(ctx context.Context, input *transcribeservice.ListLanguageModelsInput, fn func(*transcribeservice.ListLanguageModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListLanguageModelsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListLanguageModelsOutput{}
	fnCacher := func(out *transcribeservice.ListLanguageModelsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListLanguageModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListLanguageModelsWithContext(ctx context.Context, input *transcribeservice.ListLanguageModelsInput, opts ...request.Option) (*transcribeservice.ListLanguageModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListLanguageModelsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListLanguageModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListMedicalTranscriptionJobs(input *transcribeservice.ListMedicalTranscriptionJobsInput) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListMedicalTranscriptionJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListMedicalTranscriptionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListMedicalTranscriptionJobsPages(input *transcribeservice.ListMedicalTranscriptionJobsInput, fn func(*transcribeservice.ListMedicalTranscriptionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListMedicalTranscriptionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListMedicalTranscriptionJobsOutput{}
	fnCacher := func(out *transcribeservice.ListMedicalTranscriptionJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListMedicalTranscriptionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListMedicalTranscriptionJobsPagesWithContext(ctx context.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput, fn func(*transcribeservice.ListMedicalTranscriptionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListMedicalTranscriptionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListMedicalTranscriptionJobsOutput{}
	fnCacher := func(out *transcribeservice.ListMedicalTranscriptionJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListMedicalTranscriptionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListMedicalTranscriptionJobsWithContext(ctx context.Context, input *transcribeservice.ListMedicalTranscriptionJobsInput, opts ...request.Option) (*transcribeservice.ListMedicalTranscriptionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListMedicalTranscriptionJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListMedicalTranscriptionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListMedicalVocabularies(input *transcribeservice.ListMedicalVocabulariesInput) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListMedicalVocabulariesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListMedicalVocabularies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListMedicalVocabulariesPages(input *transcribeservice.ListMedicalVocabulariesInput, fn func(*transcribeservice.ListMedicalVocabulariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListMedicalVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListMedicalVocabulariesOutput{}
	fnCacher := func(out *transcribeservice.ListMedicalVocabulariesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListMedicalVocabulariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListMedicalVocabulariesPagesWithContext(ctx context.Context, input *transcribeservice.ListMedicalVocabulariesInput, fn func(*transcribeservice.ListMedicalVocabulariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListMedicalVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListMedicalVocabulariesOutput{}
	fnCacher := func(out *transcribeservice.ListMedicalVocabulariesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListMedicalVocabulariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListMedicalVocabulariesWithContext(ctx context.Context, input *transcribeservice.ListMedicalVocabulariesInput, opts ...request.Option) (*transcribeservice.ListMedicalVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListMedicalVocabulariesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListMedicalVocabulariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListTagsForResource(input *transcribeservice.ListTagsForResourceInput) (*transcribeservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListTagsForResourceOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListTagsForResourceWithContext(ctx context.Context, input *transcribeservice.ListTagsForResourceInput, opts ...request.Option) (*transcribeservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListTagsForResourceOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListTranscriptionJobs(input *transcribeservice.ListTranscriptionJobsInput) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListTranscriptionJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListTranscriptionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListTranscriptionJobsPages(input *transcribeservice.ListTranscriptionJobsInput, fn func(*transcribeservice.ListTranscriptionJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListTranscriptionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListTranscriptionJobsOutput{}
	fnCacher := func(out *transcribeservice.ListTranscriptionJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListTranscriptionJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListTranscriptionJobsPagesWithContext(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput, fn func(*transcribeservice.ListTranscriptionJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListTranscriptionJobsOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListTranscriptionJobsOutput{}
	fnCacher := func(out *transcribeservice.ListTranscriptionJobsOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListTranscriptionJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListTranscriptionJobsWithContext(ctx context.Context, input *transcribeservice.ListTranscriptionJobsInput, opts ...request.Option) (*transcribeservice.ListTranscriptionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListTranscriptionJobsOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListTranscriptionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListVocabularies(input *transcribeservice.ListVocabulariesInput) (*transcribeservice.ListVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListVocabulariesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListVocabularies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListVocabulariesPages(input *transcribeservice.ListVocabulariesInput, fn func(*transcribeservice.ListVocabulariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListVocabulariesOutput{}
	fnCacher := func(out *transcribeservice.ListVocabulariesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListVocabulariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListVocabulariesPagesWithContext(ctx context.Context, input *transcribeservice.ListVocabulariesInput, fn func(*transcribeservice.ListVocabulariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListVocabulariesOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListVocabulariesOutput{}
	fnCacher := func(out *transcribeservice.ListVocabulariesOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListVocabulariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListVocabulariesWithContext(ctx context.Context, input *transcribeservice.ListVocabulariesInput, opts ...request.Option) (*transcribeservice.ListVocabulariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListVocabulariesOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListVocabulariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListVocabularyFilters(input *transcribeservice.ListVocabularyFiltersInput) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListVocabularyFiltersOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListVocabularyFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *TranscribeService) ListVocabularyFiltersPages(input *transcribeservice.ListVocabularyFiltersInput, fn func(*transcribeservice.ListVocabularyFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListVocabularyFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListVocabularyFiltersOutput{}
	fnCacher := func(out *transcribeservice.ListVocabularyFiltersOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListVocabularyFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListVocabularyFiltersPagesWithContext(ctx context.Context, input *transcribeservice.ListVocabularyFiltersInput, fn func(*transcribeservice.ListVocabularyFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*transcribeservice.ListVocabularyFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &transcribeservice.ListVocabularyFiltersOutput{}
	fnCacher := func(out *transcribeservice.ListVocabularyFiltersOutput, more bool) bool {
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
	if err := c.TranscribeServiceAPI.ListVocabularyFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *TranscribeService) ListVocabularyFiltersWithContext(ctx context.Context, input *transcribeservice.ListVocabularyFiltersInput, opts ...request.Option) (*transcribeservice.ListVocabularyFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*transcribeservice.ListVocabularyFiltersOutput), nil
	}
	out, err := c.TranscribeServiceAPI.ListVocabularyFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
