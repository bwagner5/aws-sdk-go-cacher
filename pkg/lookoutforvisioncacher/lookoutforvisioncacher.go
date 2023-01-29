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
package lookoutforvisioncacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lookoutforvision"
	"github.com/aws/aws-sdk-go/service/lookoutforvision/lookoutforvisioniface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LookoutForVision struct {
	lookoutforvisioniface.LookoutForVisionAPI
	cache *cache.Cache
}

func New(lookoutforvisionapi lookoutforvisioniface.LookoutForVisionAPI) *LookoutForVision {
	return &LookoutForVision{
		LookoutForVisionAPI: lookoutforvisionapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LookoutForVision) DescribeDataset(input *lookoutforvision.DescribeDatasetInput) (*lookoutforvision.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeDatasetOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeDatasetWithContext(ctx context.Context, input *lookoutforvision.DescribeDatasetInput, opts ...request.Option) (*lookoutforvision.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeDatasetOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeModel(input *lookoutforvision.DescribeModelInput) (*lookoutforvision.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeModelOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeModelPackagingJob(input *lookoutforvision.DescribeModelPackagingJobInput) (*lookoutforvision.DescribeModelPackagingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeModelPackagingJobOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeModelPackagingJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeModelPackagingJobWithContext(ctx context.Context, input *lookoutforvision.DescribeModelPackagingJobInput, opts ...request.Option) (*lookoutforvision.DescribeModelPackagingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeModelPackagingJobOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeModelPackagingJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeModelWithContext(ctx context.Context, input *lookoutforvision.DescribeModelInput, opts ...request.Option) (*lookoutforvision.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeModelOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeProject(input *lookoutforvision.DescribeProjectInput) (*lookoutforvision.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeProjectOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) DescribeProjectWithContext(ctx context.Context, input *lookoutforvision.DescribeProjectInput, opts ...request.Option) (*lookoutforvision.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.DescribeProjectOutput), nil
	}
	out, err := c.LookoutForVisionAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListDatasetEntries(input *lookoutforvision.ListDatasetEntriesInput) (*lookoutforvision.ListDatasetEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListDatasetEntriesOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListDatasetEntries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListDatasetEntriesPages(input *lookoutforvision.ListDatasetEntriesInput, fn func(*lookoutforvision.ListDatasetEntriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListDatasetEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListDatasetEntriesOutput{}
	fnCacher := func(out *lookoutforvision.ListDatasetEntriesOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListDatasetEntriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListDatasetEntriesPagesWithContext(ctx context.Context, input *lookoutforvision.ListDatasetEntriesInput, fn func(*lookoutforvision.ListDatasetEntriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListDatasetEntriesOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListDatasetEntriesOutput{}
	fnCacher := func(out *lookoutforvision.ListDatasetEntriesOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListDatasetEntriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListDatasetEntriesWithContext(ctx context.Context, input *lookoutforvision.ListDatasetEntriesInput, opts ...request.Option) (*lookoutforvision.ListDatasetEntriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListDatasetEntriesOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListDatasetEntriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListModelPackagingJobs(input *lookoutforvision.ListModelPackagingJobsInput) (*lookoutforvision.ListModelPackagingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListModelPackagingJobsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListModelPackagingJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListModelPackagingJobsPages(input *lookoutforvision.ListModelPackagingJobsInput, fn func(*lookoutforvision.ListModelPackagingJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListModelPackagingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListModelPackagingJobsOutput{}
	fnCacher := func(out *lookoutforvision.ListModelPackagingJobsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListModelPackagingJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListModelPackagingJobsPagesWithContext(ctx context.Context, input *lookoutforvision.ListModelPackagingJobsInput, fn func(*lookoutforvision.ListModelPackagingJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListModelPackagingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListModelPackagingJobsOutput{}
	fnCacher := func(out *lookoutforvision.ListModelPackagingJobsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListModelPackagingJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListModelPackagingJobsWithContext(ctx context.Context, input *lookoutforvision.ListModelPackagingJobsInput, opts ...request.Option) (*lookoutforvision.ListModelPackagingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListModelPackagingJobsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListModelPackagingJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListModels(input *lookoutforvision.ListModelsInput) (*lookoutforvision.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListModelsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListModelsPages(input *lookoutforvision.ListModelsInput, fn func(*lookoutforvision.ListModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListModelsOutput{}
	fnCacher := func(out *lookoutforvision.ListModelsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListModelsPagesWithContext(ctx context.Context, input *lookoutforvision.ListModelsInput, fn func(*lookoutforvision.ListModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListModelsOutput{}
	fnCacher := func(out *lookoutforvision.ListModelsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListModelsWithContext(ctx context.Context, input *lookoutforvision.ListModelsInput, opts ...request.Option) (*lookoutforvision.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListModelsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListProjects(input *lookoutforvision.ListProjectsInput) (*lookoutforvision.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListProjectsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListProjectsPages(input *lookoutforvision.ListProjectsInput, fn func(*lookoutforvision.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListProjectsOutput{}
	fnCacher := func(out *lookoutforvision.ListProjectsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListProjectsPagesWithContext(ctx context.Context, input *lookoutforvision.ListProjectsInput, fn func(*lookoutforvision.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lookoutforvision.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &lookoutforvision.ListProjectsOutput{}
	fnCacher := func(out *lookoutforvision.ListProjectsOutput, more bool) bool {
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
	if err := c.LookoutForVisionAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LookoutForVision) ListProjectsWithContext(ctx context.Context, input *lookoutforvision.ListProjectsInput, opts ...request.Option) (*lookoutforvision.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListProjectsOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListTagsForResource(input *lookoutforvision.ListTagsForResourceInput) (*lookoutforvision.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LookoutForVision) ListTagsForResourceWithContext(ctx context.Context, input *lookoutforvision.ListTagsForResourceInput, opts ...request.Option) (*lookoutforvision.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lookoutforvision.ListTagsForResourceOutput), nil
	}
	out, err := c.LookoutForVisionAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
