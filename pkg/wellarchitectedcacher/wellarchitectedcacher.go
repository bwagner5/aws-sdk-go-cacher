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
package wellarchitectedcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/wellarchitected"
	"github.com/aws/aws-sdk-go/service/wellarchitected/wellarchitectediface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type WellArchitected struct {
	wellarchitectediface.WellArchitectedAPI
	cache *cache.Cache
}

func New(wellarchitectedapi wellarchitectediface.WellArchitectedAPI) *WellArchitected {
	return &WellArchitected{
		WellArchitectedAPI: wellarchitectedapi,
		cache:              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *WellArchitected) GetAnswer(input *wellarchitected.GetAnswerInput) (*wellarchitected.GetAnswerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetAnswerOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetAnswer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetAnswerWithContext(ctx context.Context, input *wellarchitected.GetAnswerInput, opts ...request.Option) (*wellarchitected.GetAnswerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetAnswerOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetAnswerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLens(input *wellarchitected.GetLensInput) (*wellarchitected.GetLensOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLens(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensReview(input *wellarchitected.GetLensReviewInput) (*wellarchitected.GetLensReviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensReviewOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensReview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensReviewReport(input *wellarchitected.GetLensReviewReportInput) (*wellarchitected.GetLensReviewReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensReviewReportOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensReviewReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensReviewReportWithContext(ctx context.Context, input *wellarchitected.GetLensReviewReportInput, opts ...request.Option) (*wellarchitected.GetLensReviewReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensReviewReportOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensReviewReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensReviewWithContext(ctx context.Context, input *wellarchitected.GetLensReviewInput, opts ...request.Option) (*wellarchitected.GetLensReviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensReviewOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensReviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensVersionDifference(input *wellarchitected.GetLensVersionDifferenceInput) (*wellarchitected.GetLensVersionDifferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensVersionDifferenceOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensVersionDifference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensVersionDifferenceWithContext(ctx context.Context, input *wellarchitected.GetLensVersionDifferenceInput, opts ...request.Option) (*wellarchitected.GetLensVersionDifferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensVersionDifferenceOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensVersionDifferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetLensWithContext(ctx context.Context, input *wellarchitected.GetLensInput, opts ...request.Option) (*wellarchitected.GetLensOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetLensOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetLensWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetMilestone(input *wellarchitected.GetMilestoneInput) (*wellarchitected.GetMilestoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetMilestoneOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetMilestone(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetMilestoneWithContext(ctx context.Context, input *wellarchitected.GetMilestoneInput, opts ...request.Option) (*wellarchitected.GetMilestoneOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetMilestoneOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetMilestoneWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetWorkload(input *wellarchitected.GetWorkloadInput) (*wellarchitected.GetWorkloadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetWorkloadOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetWorkload(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) GetWorkloadWithContext(ctx context.Context, input *wellarchitected.GetWorkloadInput, opts ...request.Option) (*wellarchitected.GetWorkloadOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.GetWorkloadOutput), nil
	}
	out, err := c.WellArchitectedAPI.GetWorkloadWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListAnswers(input *wellarchitected.ListAnswersInput) (*wellarchitected.ListAnswersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListAnswersOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListAnswers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListAnswersPages(input *wellarchitected.ListAnswersInput, fn func(*wellarchitected.ListAnswersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListAnswersOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListAnswersOutput{}
	fnCacher := func(out *wellarchitected.ListAnswersOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListAnswersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListAnswersPagesWithContext(ctx context.Context, input *wellarchitected.ListAnswersInput, fn func(*wellarchitected.ListAnswersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListAnswersOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListAnswersOutput{}
	fnCacher := func(out *wellarchitected.ListAnswersOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListAnswersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListAnswersWithContext(ctx context.Context, input *wellarchitected.ListAnswersInput, opts ...request.Option) (*wellarchitected.ListAnswersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListAnswersOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListAnswersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListCheckDetails(input *wellarchitected.ListCheckDetailsInput) (*wellarchitected.ListCheckDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListCheckDetailsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListCheckDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListCheckDetailsPages(input *wellarchitected.ListCheckDetailsInput, fn func(*wellarchitected.ListCheckDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListCheckDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListCheckDetailsOutput{}
	fnCacher := func(out *wellarchitected.ListCheckDetailsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListCheckDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListCheckDetailsPagesWithContext(ctx context.Context, input *wellarchitected.ListCheckDetailsInput, fn func(*wellarchitected.ListCheckDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListCheckDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListCheckDetailsOutput{}
	fnCacher := func(out *wellarchitected.ListCheckDetailsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListCheckDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListCheckDetailsWithContext(ctx context.Context, input *wellarchitected.ListCheckDetailsInput, opts ...request.Option) (*wellarchitected.ListCheckDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListCheckDetailsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListCheckDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListCheckSummaries(input *wellarchitected.ListCheckSummariesInput) (*wellarchitected.ListCheckSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListCheckSummariesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListCheckSummaries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListCheckSummariesPages(input *wellarchitected.ListCheckSummariesInput, fn func(*wellarchitected.ListCheckSummariesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListCheckSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListCheckSummariesOutput{}
	fnCacher := func(out *wellarchitected.ListCheckSummariesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListCheckSummariesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListCheckSummariesPagesWithContext(ctx context.Context, input *wellarchitected.ListCheckSummariesInput, fn func(*wellarchitected.ListCheckSummariesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListCheckSummariesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListCheckSummariesOutput{}
	fnCacher := func(out *wellarchitected.ListCheckSummariesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListCheckSummariesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListCheckSummariesWithContext(ctx context.Context, input *wellarchitected.ListCheckSummariesInput, opts ...request.Option) (*wellarchitected.ListCheckSummariesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListCheckSummariesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListCheckSummariesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensReviewImprovements(input *wellarchitected.ListLensReviewImprovementsInput) (*wellarchitected.ListLensReviewImprovementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensReviewImprovementsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensReviewImprovements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensReviewImprovementsPages(input *wellarchitected.ListLensReviewImprovementsInput, fn func(*wellarchitected.ListLensReviewImprovementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensReviewImprovementsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensReviewImprovementsOutput{}
	fnCacher := func(out *wellarchitected.ListLensReviewImprovementsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensReviewImprovementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensReviewImprovementsPagesWithContext(ctx context.Context, input *wellarchitected.ListLensReviewImprovementsInput, fn func(*wellarchitected.ListLensReviewImprovementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensReviewImprovementsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensReviewImprovementsOutput{}
	fnCacher := func(out *wellarchitected.ListLensReviewImprovementsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensReviewImprovementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensReviewImprovementsWithContext(ctx context.Context, input *wellarchitected.ListLensReviewImprovementsInput, opts ...request.Option) (*wellarchitected.ListLensReviewImprovementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensReviewImprovementsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensReviewImprovementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensReviews(input *wellarchitected.ListLensReviewsInput) (*wellarchitected.ListLensReviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensReviewsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensReviews(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensReviewsPages(input *wellarchitected.ListLensReviewsInput, fn func(*wellarchitected.ListLensReviewsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensReviewsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensReviewsOutput{}
	fnCacher := func(out *wellarchitected.ListLensReviewsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensReviewsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensReviewsPagesWithContext(ctx context.Context, input *wellarchitected.ListLensReviewsInput, fn func(*wellarchitected.ListLensReviewsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensReviewsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensReviewsOutput{}
	fnCacher := func(out *wellarchitected.ListLensReviewsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensReviewsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensReviewsWithContext(ctx context.Context, input *wellarchitected.ListLensReviewsInput, opts ...request.Option) (*wellarchitected.ListLensReviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensReviewsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensReviewsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensShares(input *wellarchitected.ListLensSharesInput) (*wellarchitected.ListLensSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensSharesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensSharesPages(input *wellarchitected.ListLensSharesInput, fn func(*wellarchitected.ListLensSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensSharesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensSharesOutput{}
	fnCacher := func(out *wellarchitected.ListLensSharesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensSharesPagesWithContext(ctx context.Context, input *wellarchitected.ListLensSharesInput, fn func(*wellarchitected.ListLensSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensSharesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensSharesOutput{}
	fnCacher := func(out *wellarchitected.ListLensSharesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensSharesWithContext(ctx context.Context, input *wellarchitected.ListLensSharesInput, opts ...request.Option) (*wellarchitected.ListLensSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensSharesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLenses(input *wellarchitected.ListLensesInput) (*wellarchitected.ListLensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLenses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListLensesPages(input *wellarchitected.ListLensesInput, fn func(*wellarchitected.ListLensesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensesOutput{}
	fnCacher := func(out *wellarchitected.ListLensesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensesPagesWithContext(ctx context.Context, input *wellarchitected.ListLensesInput, fn func(*wellarchitected.ListLensesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListLensesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListLensesOutput{}
	fnCacher := func(out *wellarchitected.ListLensesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListLensesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListLensesWithContext(ctx context.Context, input *wellarchitected.ListLensesInput, opts ...request.Option) (*wellarchitected.ListLensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListLensesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListLensesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListMilestones(input *wellarchitected.ListMilestonesInput) (*wellarchitected.ListMilestonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListMilestonesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListMilestones(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListMilestonesPages(input *wellarchitected.ListMilestonesInput, fn func(*wellarchitected.ListMilestonesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListMilestonesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListMilestonesOutput{}
	fnCacher := func(out *wellarchitected.ListMilestonesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListMilestonesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListMilestonesPagesWithContext(ctx context.Context, input *wellarchitected.ListMilestonesInput, fn func(*wellarchitected.ListMilestonesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListMilestonesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListMilestonesOutput{}
	fnCacher := func(out *wellarchitected.ListMilestonesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListMilestonesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListMilestonesWithContext(ctx context.Context, input *wellarchitected.ListMilestonesInput, opts ...request.Option) (*wellarchitected.ListMilestonesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListMilestonesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListMilestonesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListNotifications(input *wellarchitected.ListNotificationsInput) (*wellarchitected.ListNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListNotificationsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListNotifications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListNotificationsPages(input *wellarchitected.ListNotificationsInput, fn func(*wellarchitected.ListNotificationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListNotificationsOutput{}
	fnCacher := func(out *wellarchitected.ListNotificationsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListNotificationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListNotificationsPagesWithContext(ctx context.Context, input *wellarchitected.ListNotificationsInput, fn func(*wellarchitected.ListNotificationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListNotificationsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListNotificationsOutput{}
	fnCacher := func(out *wellarchitected.ListNotificationsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListNotificationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListNotificationsWithContext(ctx context.Context, input *wellarchitected.ListNotificationsInput, opts ...request.Option) (*wellarchitected.ListNotificationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListNotificationsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListNotificationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListShareInvitations(input *wellarchitected.ListShareInvitationsInput) (*wellarchitected.ListShareInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListShareInvitationsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListShareInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListShareInvitationsPages(input *wellarchitected.ListShareInvitationsInput, fn func(*wellarchitected.ListShareInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListShareInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListShareInvitationsOutput{}
	fnCacher := func(out *wellarchitected.ListShareInvitationsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListShareInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListShareInvitationsPagesWithContext(ctx context.Context, input *wellarchitected.ListShareInvitationsInput, fn func(*wellarchitected.ListShareInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListShareInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListShareInvitationsOutput{}
	fnCacher := func(out *wellarchitected.ListShareInvitationsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListShareInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListShareInvitationsWithContext(ctx context.Context, input *wellarchitected.ListShareInvitationsInput, opts ...request.Option) (*wellarchitected.ListShareInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListShareInvitationsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListShareInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListTagsForResource(input *wellarchitected.ListTagsForResourceInput) (*wellarchitected.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListTagsForResourceOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListTagsForResourceWithContext(ctx context.Context, input *wellarchitected.ListTagsForResourceInput, opts ...request.Option) (*wellarchitected.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListTagsForResourceOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListWorkloadShares(input *wellarchitected.ListWorkloadSharesInput) (*wellarchitected.ListWorkloadSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListWorkloadSharesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListWorkloadShares(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListWorkloadSharesPages(input *wellarchitected.ListWorkloadSharesInput, fn func(*wellarchitected.ListWorkloadSharesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListWorkloadSharesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListWorkloadSharesOutput{}
	fnCacher := func(out *wellarchitected.ListWorkloadSharesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListWorkloadSharesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListWorkloadSharesPagesWithContext(ctx context.Context, input *wellarchitected.ListWorkloadSharesInput, fn func(*wellarchitected.ListWorkloadSharesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListWorkloadSharesOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListWorkloadSharesOutput{}
	fnCacher := func(out *wellarchitected.ListWorkloadSharesOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListWorkloadSharesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListWorkloadSharesWithContext(ctx context.Context, input *wellarchitected.ListWorkloadSharesInput, opts ...request.Option) (*wellarchitected.ListWorkloadSharesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListWorkloadSharesOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListWorkloadSharesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListWorkloads(input *wellarchitected.ListWorkloadsInput) (*wellarchitected.ListWorkloadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListWorkloadsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListWorkloads(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *WellArchitected) ListWorkloadsPages(input *wellarchitected.ListWorkloadsInput, fn func(*wellarchitected.ListWorkloadsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListWorkloadsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListWorkloadsOutput{}
	fnCacher := func(out *wellarchitected.ListWorkloadsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListWorkloadsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListWorkloadsPagesWithContext(ctx context.Context, input *wellarchitected.ListWorkloadsInput, fn func(*wellarchitected.ListWorkloadsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*wellarchitected.ListWorkloadsOutput), false)
		return nil
	}
	cachable := true
	output := &wellarchitected.ListWorkloadsOutput{}
	fnCacher := func(out *wellarchitected.ListWorkloadsOutput, more bool) bool {
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
	if err := c.WellArchitectedAPI.ListWorkloadsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *WellArchitected) ListWorkloadsWithContext(ctx context.Context, input *wellarchitected.ListWorkloadsInput, opts ...request.Option) (*wellarchitected.ListWorkloadsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*wellarchitected.ListWorkloadsOutput), nil
	}
	out, err := c.WellArchitectedAPI.ListWorkloadsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
