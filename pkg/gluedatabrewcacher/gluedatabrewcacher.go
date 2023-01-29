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
package gluedatabrewcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/gluedatabrew"
	"github.com/aws/aws-sdk-go/service/gluedatabrew/gluedatabrewiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GlueDataBrew struct {
	gluedatabrewiface.GlueDataBrewAPI
	cache *cache.Cache
}

func New(gluedatabrewapi gluedatabrewiface.GlueDataBrewAPI) *GlueDataBrew {
	return &GlueDataBrew{
		GlueDataBrewAPI: gluedatabrewapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GlueDataBrew) BatchDeleteRecipeVersion(input *gluedatabrew.BatchDeleteRecipeVersionInput) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.BatchDeleteRecipeVersionOutput), nil
	}
	out, err := c.GlueDataBrewAPI.BatchDeleteRecipeVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) BatchDeleteRecipeVersionWithContext(ctx context.Context, input *gluedatabrew.BatchDeleteRecipeVersionInput, opts ...request.Option) (*gluedatabrew.BatchDeleteRecipeVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.BatchDeleteRecipeVersionOutput), nil
	}
	out, err := c.GlueDataBrewAPI.BatchDeleteRecipeVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeDataset(input *gluedatabrew.DescribeDatasetInput) (*gluedatabrew.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeDatasetOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeDataset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeDatasetWithContext(ctx context.Context, input *gluedatabrew.DescribeDatasetInput, opts ...request.Option) (*gluedatabrew.DescribeDatasetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeDatasetOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeDatasetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeJob(input *gluedatabrew.DescribeJobInput) (*gluedatabrew.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeJobOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeJobRun(input *gluedatabrew.DescribeJobRunInput) (*gluedatabrew.DescribeJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeJobRunOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeJobRunWithContext(ctx context.Context, input *gluedatabrew.DescribeJobRunInput, opts ...request.Option) (*gluedatabrew.DescribeJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeJobRunOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeJobWithContext(ctx context.Context, input *gluedatabrew.DescribeJobInput, opts ...request.Option) (*gluedatabrew.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeJobOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeProject(input *gluedatabrew.DescribeProjectInput) (*gluedatabrew.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeProjectOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeProjectWithContext(ctx context.Context, input *gluedatabrew.DescribeProjectInput, opts ...request.Option) (*gluedatabrew.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeProjectOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeRecipe(input *gluedatabrew.DescribeRecipeInput) (*gluedatabrew.DescribeRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeRecipeOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeRecipe(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeRecipeWithContext(ctx context.Context, input *gluedatabrew.DescribeRecipeInput, opts ...request.Option) (*gluedatabrew.DescribeRecipeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeRecipeOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeRecipeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeRuleset(input *gluedatabrew.DescribeRulesetInput) (*gluedatabrew.DescribeRulesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeRulesetOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeRuleset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeRulesetWithContext(ctx context.Context, input *gluedatabrew.DescribeRulesetInput, opts ...request.Option) (*gluedatabrew.DescribeRulesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeRulesetOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeRulesetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeSchedule(input *gluedatabrew.DescribeScheduleInput) (*gluedatabrew.DescribeScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeScheduleOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) DescribeScheduleWithContext(ctx context.Context, input *gluedatabrew.DescribeScheduleInput, opts ...request.Option) (*gluedatabrew.DescribeScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.DescribeScheduleOutput), nil
	}
	out, err := c.GlueDataBrewAPI.DescribeScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListDatasets(input *gluedatabrew.ListDatasetsInput) (*gluedatabrew.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListDatasetsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListDatasets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListDatasetsPages(input *gluedatabrew.ListDatasetsInput, fn func(*gluedatabrew.ListDatasetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListDatasetsOutput{}
	fnCacher := func(out *gluedatabrew.ListDatasetsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListDatasetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListDatasetsPagesWithContext(ctx context.Context, input *gluedatabrew.ListDatasetsInput, fn func(*gluedatabrew.ListDatasetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListDatasetsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListDatasetsOutput{}
	fnCacher := func(out *gluedatabrew.ListDatasetsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListDatasetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListDatasetsWithContext(ctx context.Context, input *gluedatabrew.ListDatasetsInput, opts ...request.Option) (*gluedatabrew.ListDatasetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListDatasetsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListDatasetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListJobRuns(input *gluedatabrew.ListJobRunsInput) (*gluedatabrew.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListJobRunsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListJobRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListJobRunsPages(input *gluedatabrew.ListJobRunsInput, fn func(*gluedatabrew.ListJobRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListJobRunsOutput{}
	fnCacher := func(out *gluedatabrew.ListJobRunsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListJobRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListJobRunsPagesWithContext(ctx context.Context, input *gluedatabrew.ListJobRunsInput, fn func(*gluedatabrew.ListJobRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListJobRunsOutput{}
	fnCacher := func(out *gluedatabrew.ListJobRunsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListJobRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListJobRunsWithContext(ctx context.Context, input *gluedatabrew.ListJobRunsInput, opts ...request.Option) (*gluedatabrew.ListJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListJobRunsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListJobRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListJobs(input *gluedatabrew.ListJobsInput) (*gluedatabrew.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListJobsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListJobsPages(input *gluedatabrew.ListJobsInput, fn func(*gluedatabrew.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListJobsOutput{}
	fnCacher := func(out *gluedatabrew.ListJobsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListJobsPagesWithContext(ctx context.Context, input *gluedatabrew.ListJobsInput, fn func(*gluedatabrew.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListJobsOutput{}
	fnCacher := func(out *gluedatabrew.ListJobsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListJobsWithContext(ctx context.Context, input *gluedatabrew.ListJobsInput, opts ...request.Option) (*gluedatabrew.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListJobsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListProjects(input *gluedatabrew.ListProjectsInput) (*gluedatabrew.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListProjectsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListProjectsPages(input *gluedatabrew.ListProjectsInput, fn func(*gluedatabrew.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListProjectsOutput{}
	fnCacher := func(out *gluedatabrew.ListProjectsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListProjectsPagesWithContext(ctx context.Context, input *gluedatabrew.ListProjectsInput, fn func(*gluedatabrew.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListProjectsOutput{}
	fnCacher := func(out *gluedatabrew.ListProjectsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListProjectsWithContext(ctx context.Context, input *gluedatabrew.ListProjectsInput, opts ...request.Option) (*gluedatabrew.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListProjectsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRecipeVersions(input *gluedatabrew.ListRecipeVersionsInput) (*gluedatabrew.ListRecipeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRecipeVersionsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRecipeVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRecipeVersionsPages(input *gluedatabrew.ListRecipeVersionsInput, fn func(*gluedatabrew.ListRecipeVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRecipeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRecipeVersionsOutput{}
	fnCacher := func(out *gluedatabrew.ListRecipeVersionsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRecipeVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRecipeVersionsPagesWithContext(ctx context.Context, input *gluedatabrew.ListRecipeVersionsInput, fn func(*gluedatabrew.ListRecipeVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRecipeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRecipeVersionsOutput{}
	fnCacher := func(out *gluedatabrew.ListRecipeVersionsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRecipeVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRecipeVersionsWithContext(ctx context.Context, input *gluedatabrew.ListRecipeVersionsInput, opts ...request.Option) (*gluedatabrew.ListRecipeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRecipeVersionsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRecipeVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRecipes(input *gluedatabrew.ListRecipesInput) (*gluedatabrew.ListRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRecipesOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRecipes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRecipesPages(input *gluedatabrew.ListRecipesInput, fn func(*gluedatabrew.ListRecipesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRecipesOutput{}
	fnCacher := func(out *gluedatabrew.ListRecipesOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRecipesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRecipesPagesWithContext(ctx context.Context, input *gluedatabrew.ListRecipesInput, fn func(*gluedatabrew.ListRecipesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRecipesOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRecipesOutput{}
	fnCacher := func(out *gluedatabrew.ListRecipesOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRecipesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRecipesWithContext(ctx context.Context, input *gluedatabrew.ListRecipesInput, opts ...request.Option) (*gluedatabrew.ListRecipesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRecipesOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRecipesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRulesets(input *gluedatabrew.ListRulesetsInput) (*gluedatabrew.ListRulesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRulesetsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRulesets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListRulesetsPages(input *gluedatabrew.ListRulesetsInput, fn func(*gluedatabrew.ListRulesetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRulesetsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRulesetsOutput{}
	fnCacher := func(out *gluedatabrew.ListRulesetsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRulesetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRulesetsPagesWithContext(ctx context.Context, input *gluedatabrew.ListRulesetsInput, fn func(*gluedatabrew.ListRulesetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListRulesetsOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListRulesetsOutput{}
	fnCacher := func(out *gluedatabrew.ListRulesetsOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListRulesetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListRulesetsWithContext(ctx context.Context, input *gluedatabrew.ListRulesetsInput, opts ...request.Option) (*gluedatabrew.ListRulesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListRulesetsOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListRulesetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListSchedules(input *gluedatabrew.ListSchedulesInput) (*gluedatabrew.ListSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListSchedulesOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListSchedules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListSchedulesPages(input *gluedatabrew.ListSchedulesInput, fn func(*gluedatabrew.ListSchedulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListSchedulesOutput{}
	fnCacher := func(out *gluedatabrew.ListSchedulesOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListSchedulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListSchedulesPagesWithContext(ctx context.Context, input *gluedatabrew.ListSchedulesInput, fn func(*gluedatabrew.ListSchedulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*gluedatabrew.ListSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &gluedatabrew.ListSchedulesOutput{}
	fnCacher := func(out *gluedatabrew.ListSchedulesOutput, more bool) bool {
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
	if err := c.GlueDataBrewAPI.ListSchedulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GlueDataBrew) ListSchedulesWithContext(ctx context.Context, input *gluedatabrew.ListSchedulesInput, opts ...request.Option) (*gluedatabrew.ListSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListSchedulesOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListSchedulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListTagsForResource(input *gluedatabrew.ListTagsForResourceInput) (*gluedatabrew.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListTagsForResourceOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GlueDataBrew) ListTagsForResourceWithContext(ctx context.Context, input *gluedatabrew.ListTagsForResourceInput, opts ...request.Option) (*gluedatabrew.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*gluedatabrew.ListTagsForResourceOutput), nil
	}
	out, err := c.GlueDataBrewAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
