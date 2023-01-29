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
package codegurureviewercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"github.com/aws/aws-sdk-go/service/codegurureviewer/codegururevieweriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeGuruReviewer struct {
	codegururevieweriface.CodeGuruReviewerAPI
	cache *cache.Cache
}

func New(codegurureviewerapi codegururevieweriface.CodeGuruReviewerAPI) *CodeGuruReviewer {
	return &CodeGuruReviewer{
		CodeGuruReviewerAPI: codegurureviewerapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeGuruReviewer) DescribeCodeReview(input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeCodeReviewOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeCodeReview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) DescribeCodeReviewWithContext(ctx context.Context, input *codegurureviewer.DescribeCodeReviewInput, opts ...request.Option) (*codegurureviewer.DescribeCodeReviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeCodeReviewOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeCodeReviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) DescribeRecommendationFeedback(input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeRecommendationFeedbackOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeRecommendationFeedback(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) DescribeRecommendationFeedbackWithContext(ctx context.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput, opts ...request.Option) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeRecommendationFeedbackOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeRecommendationFeedbackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) DescribeRepositoryAssociation(input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeRepositoryAssociationOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeRepositoryAssociation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) DescribeRepositoryAssociationWithContext(ctx context.Context, input *codegurureviewer.DescribeRepositoryAssociationInput, opts ...request.Option) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.DescribeRepositoryAssociationOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.DescribeRepositoryAssociationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListCodeReviews(input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListCodeReviewsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListCodeReviews(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListCodeReviewsPages(input *codegurureviewer.ListCodeReviewsInput, fn func(*codegurureviewer.ListCodeReviewsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListCodeReviewsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListCodeReviewsOutput{}
	fnCacher := func(out *codegurureviewer.ListCodeReviewsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListCodeReviewsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListCodeReviewsPagesWithContext(ctx context.Context, input *codegurureviewer.ListCodeReviewsInput, fn func(*codegurureviewer.ListCodeReviewsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListCodeReviewsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListCodeReviewsOutput{}
	fnCacher := func(out *codegurureviewer.ListCodeReviewsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListCodeReviewsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListCodeReviewsWithContext(ctx context.Context, input *codegurureviewer.ListCodeReviewsInput, opts ...request.Option) (*codegurureviewer.ListCodeReviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListCodeReviewsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListCodeReviewsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRecommendationFeedback(input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRecommendationFeedbackOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRecommendationFeedback(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRecommendationFeedbackPages(input *codegurureviewer.ListRecommendationFeedbackInput, fn func(*codegurureviewer.ListRecommendationFeedbackOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRecommendationFeedbackOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRecommendationFeedbackOutput{}
	fnCacher := func(out *codegurureviewer.ListRecommendationFeedbackOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRecommendationFeedbackPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRecommendationFeedbackPagesWithContext(ctx context.Context, input *codegurureviewer.ListRecommendationFeedbackInput, fn func(*codegurureviewer.ListRecommendationFeedbackOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRecommendationFeedbackOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRecommendationFeedbackOutput{}
	fnCacher := func(out *codegurureviewer.ListRecommendationFeedbackOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRecommendationFeedbackPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRecommendationFeedbackWithContext(ctx context.Context, input *codegurureviewer.ListRecommendationFeedbackInput, opts ...request.Option) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRecommendationFeedbackOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRecommendationFeedbackWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRecommendations(input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRecommendationsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRecommendationsPages(input *codegurureviewer.ListRecommendationsInput, fn func(*codegurureviewer.ListRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRecommendationsOutput{}
	fnCacher := func(out *codegurureviewer.ListRecommendationsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRecommendationsPagesWithContext(ctx context.Context, input *codegurureviewer.ListRecommendationsInput, fn func(*codegurureviewer.ListRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRecommendationsOutput{}
	fnCacher := func(out *codegurureviewer.ListRecommendationsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRecommendationsWithContext(ctx context.Context, input *codegurureviewer.ListRecommendationsInput, opts ...request.Option) (*codegurureviewer.ListRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRecommendationsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRepositoryAssociations(input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRepositoryAssociationsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRepositoryAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListRepositoryAssociationsPages(input *codegurureviewer.ListRepositoryAssociationsInput, fn func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRepositoryAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRepositoryAssociationsOutput{}
	fnCacher := func(out *codegurureviewer.ListRepositoryAssociationsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRepositoryAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRepositoryAssociationsPagesWithContext(ctx context.Context, input *codegurureviewer.ListRepositoryAssociationsInput, fn func(*codegurureviewer.ListRepositoryAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codegurureviewer.ListRepositoryAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &codegurureviewer.ListRepositoryAssociationsOutput{}
	fnCacher := func(out *codegurureviewer.ListRepositoryAssociationsOutput, more bool) bool {
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
	if err := c.CodeGuruReviewerAPI.ListRepositoryAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeGuruReviewer) ListRepositoryAssociationsWithContext(ctx context.Context, input *codegurureviewer.ListRepositoryAssociationsInput, opts ...request.Option) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListRepositoryAssociationsOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListRepositoryAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListTagsForResource(input *codegurureviewer.ListTagsForResourceInput) (*codegurureviewer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeGuruReviewer) ListTagsForResourceWithContext(ctx context.Context, input *codegurureviewer.ListTagsForResourceInput, opts ...request.Option) (*codegurureviewer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codegurureviewer.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeGuruReviewerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
