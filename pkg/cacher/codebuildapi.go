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
	"github.com/aws/aws-sdk-go/service/codebuild"
	"github.com/aws/aws-sdk-go/service/codebuild/codebuildiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeBuild struct {
	codebuildiface.CodeBuildAPI
	cache *cache.Cache
}

func NewCodeBuild(codebuildapi codebuildiface.CodeBuildAPI) *CodeBuild {
	return &CodeBuild{
		CodeBuildAPI: codebuildapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeBuild) BatchDeleteBuilds(input *codebuild.BatchDeleteBuildsInput) (*codebuild.BatchDeleteBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchDeleteBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchDeleteBuilds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchDeleteBuildsWithContext(ctx context.Context, input *codebuild.BatchDeleteBuildsInput, opts ...request.Option) (*codebuild.BatchDeleteBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchDeleteBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchDeleteBuildsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetBuildBatches(input *codebuild.BatchGetBuildBatchesInput) (*codebuild.BatchGetBuildBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetBuildBatchesOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetBuildBatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetBuildBatchesWithContext(ctx context.Context, input *codebuild.BatchGetBuildBatchesInput, opts ...request.Option) (*codebuild.BatchGetBuildBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetBuildBatchesOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetBuildBatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetBuilds(input *codebuild.BatchGetBuildsInput) (*codebuild.BatchGetBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetBuilds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetBuildsWithContext(ctx context.Context, input *codebuild.BatchGetBuildsInput, opts ...request.Option) (*codebuild.BatchGetBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetBuildsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetProjects(input *codebuild.BatchGetProjectsInput) (*codebuild.BatchGetProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetProjectsWithContext(ctx context.Context, input *codebuild.BatchGetProjectsInput, opts ...request.Option) (*codebuild.BatchGetProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetReportGroups(input *codebuild.BatchGetReportGroupsInput) (*codebuild.BatchGetReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetReportGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetReportGroupsWithContext(ctx context.Context, input *codebuild.BatchGetReportGroupsInput, opts ...request.Option) (*codebuild.BatchGetReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetReportGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetReports(input *codebuild.BatchGetReportsInput) (*codebuild.BatchGetReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetReportsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) BatchGetReportsWithContext(ctx context.Context, input *codebuild.BatchGetReportsInput, opts ...request.Option) (*codebuild.BatchGetReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.BatchGetReportsOutput), nil
	}
	out, err := c.CodeBuildAPI.BatchGetReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) DescribeCodeCoverages(input *codebuild.DescribeCodeCoveragesInput) (*codebuild.DescribeCodeCoveragesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.DescribeCodeCoveragesOutput), nil
	}
	out, err := c.CodeBuildAPI.DescribeCodeCoverages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) DescribeCodeCoveragesPages(input *codebuild.DescribeCodeCoveragesInput, fn func(*codebuild.DescribeCodeCoveragesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.DescribeCodeCoveragesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.DescribeCodeCoveragesOutput{}
	fnCacher := func(out *codebuild.DescribeCodeCoveragesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.DescribeCodeCoveragesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) DescribeCodeCoveragesPagesWithContext(ctx context.Context, input *codebuild.DescribeCodeCoveragesInput, fn func(*codebuild.DescribeCodeCoveragesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.DescribeCodeCoveragesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.DescribeCodeCoveragesOutput{}
	fnCacher := func(out *codebuild.DescribeCodeCoveragesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.DescribeCodeCoveragesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) DescribeCodeCoveragesWithContext(ctx context.Context, input *codebuild.DescribeCodeCoveragesInput, opts ...request.Option) (*codebuild.DescribeCodeCoveragesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.DescribeCodeCoveragesOutput), nil
	}
	out, err := c.CodeBuildAPI.DescribeCodeCoveragesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) DescribeTestCases(input *codebuild.DescribeTestCasesInput) (*codebuild.DescribeTestCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.DescribeTestCasesOutput), nil
	}
	out, err := c.CodeBuildAPI.DescribeTestCases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) DescribeTestCasesPages(input *codebuild.DescribeTestCasesInput, fn func(*codebuild.DescribeTestCasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.DescribeTestCasesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.DescribeTestCasesOutput{}
	fnCacher := func(out *codebuild.DescribeTestCasesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.DescribeTestCasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) DescribeTestCasesPagesWithContext(ctx context.Context, input *codebuild.DescribeTestCasesInput, fn func(*codebuild.DescribeTestCasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.DescribeTestCasesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.DescribeTestCasesOutput{}
	fnCacher := func(out *codebuild.DescribeTestCasesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.DescribeTestCasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) DescribeTestCasesWithContext(ctx context.Context, input *codebuild.DescribeTestCasesInput, opts ...request.Option) (*codebuild.DescribeTestCasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.DescribeTestCasesOutput), nil
	}
	out, err := c.CodeBuildAPI.DescribeTestCasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) GetReportGroupTrend(input *codebuild.GetReportGroupTrendInput) (*codebuild.GetReportGroupTrendOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.GetReportGroupTrendOutput), nil
	}
	out, err := c.CodeBuildAPI.GetReportGroupTrend(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) GetReportGroupTrendWithContext(ctx context.Context, input *codebuild.GetReportGroupTrendInput, opts ...request.Option) (*codebuild.GetReportGroupTrendOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.GetReportGroupTrendOutput), nil
	}
	out, err := c.CodeBuildAPI.GetReportGroupTrendWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) GetResourcePolicy(input *codebuild.GetResourcePolicyInput) (*codebuild.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.GetResourcePolicyOutput), nil
	}
	out, err := c.CodeBuildAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) GetResourcePolicyWithContext(ctx context.Context, input *codebuild.GetResourcePolicyInput, opts ...request.Option) (*codebuild.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.GetResourcePolicyOutput), nil
	}
	out, err := c.CodeBuildAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildBatches(input *codebuild.ListBuildBatchesInput) (*codebuild.ListBuildBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildBatchesOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildBatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildBatchesForProject(input *codebuild.ListBuildBatchesForProjectInput) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildBatchesForProjectOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildBatchesForProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildBatchesForProjectPages(input *codebuild.ListBuildBatchesForProjectInput, fn func(*codebuild.ListBuildBatchesForProjectOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildBatchesForProjectOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildBatchesForProjectOutput{}
	fnCacher := func(out *codebuild.ListBuildBatchesForProjectOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildBatchesForProjectPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildBatchesForProjectPagesWithContext(ctx context.Context, input *codebuild.ListBuildBatchesForProjectInput, fn func(*codebuild.ListBuildBatchesForProjectOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildBatchesForProjectOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildBatchesForProjectOutput{}
	fnCacher := func(out *codebuild.ListBuildBatchesForProjectOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildBatchesForProjectPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildBatchesForProjectWithContext(ctx context.Context, input *codebuild.ListBuildBatchesForProjectInput, opts ...request.Option) (*codebuild.ListBuildBatchesForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildBatchesForProjectOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildBatchesForProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildBatchesPages(input *codebuild.ListBuildBatchesInput, fn func(*codebuild.ListBuildBatchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildBatchesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildBatchesOutput{}
	fnCacher := func(out *codebuild.ListBuildBatchesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildBatchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildBatchesPagesWithContext(ctx context.Context, input *codebuild.ListBuildBatchesInput, fn func(*codebuild.ListBuildBatchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildBatchesOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildBatchesOutput{}
	fnCacher := func(out *codebuild.ListBuildBatchesOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildBatchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildBatchesWithContext(ctx context.Context, input *codebuild.ListBuildBatchesInput, opts ...request.Option) (*codebuild.ListBuildBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildBatchesOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildBatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuilds(input *codebuild.ListBuildsInput) (*codebuild.ListBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuilds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildsForProject(input *codebuild.ListBuildsForProjectInput) (*codebuild.ListBuildsForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildsForProjectOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildsForProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildsForProjectPages(input *codebuild.ListBuildsForProjectInput, fn func(*codebuild.ListBuildsForProjectOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildsForProjectOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildsForProjectOutput{}
	fnCacher := func(out *codebuild.ListBuildsForProjectOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildsForProjectPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildsForProjectPagesWithContext(ctx context.Context, input *codebuild.ListBuildsForProjectInput, fn func(*codebuild.ListBuildsForProjectOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildsForProjectOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildsForProjectOutput{}
	fnCacher := func(out *codebuild.ListBuildsForProjectOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildsForProjectPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildsForProjectWithContext(ctx context.Context, input *codebuild.ListBuildsForProjectInput, opts ...request.Option) (*codebuild.ListBuildsForProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildsForProjectOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildsForProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListBuildsPages(input *codebuild.ListBuildsInput, fn func(*codebuild.ListBuildsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildsOutput{}
	fnCacher := func(out *codebuild.ListBuildsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildsPagesWithContext(ctx context.Context, input *codebuild.ListBuildsInput, fn func(*codebuild.ListBuildsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListBuildsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListBuildsOutput{}
	fnCacher := func(out *codebuild.ListBuildsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListBuildsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListBuildsWithContext(ctx context.Context, input *codebuild.ListBuildsInput, opts ...request.Option) (*codebuild.ListBuildsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListBuildsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListBuildsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListCuratedEnvironmentImages(input *codebuild.ListCuratedEnvironmentImagesInput) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListCuratedEnvironmentImagesOutput), nil
	}
	out, err := c.CodeBuildAPI.ListCuratedEnvironmentImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListCuratedEnvironmentImagesWithContext(ctx context.Context, input *codebuild.ListCuratedEnvironmentImagesInput, opts ...request.Option) (*codebuild.ListCuratedEnvironmentImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListCuratedEnvironmentImagesOutput), nil
	}
	out, err := c.CodeBuildAPI.ListCuratedEnvironmentImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListProjects(input *codebuild.ListProjectsInput) (*codebuild.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListProjectsPages(input *codebuild.ListProjectsInput, fn func(*codebuild.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListProjectsOutput{}
	fnCacher := func(out *codebuild.ListProjectsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListProjectsPagesWithContext(ctx context.Context, input *codebuild.ListProjectsInput, fn func(*codebuild.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListProjectsOutput{}
	fnCacher := func(out *codebuild.ListProjectsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListProjectsWithContext(ctx context.Context, input *codebuild.ListProjectsInput, opts ...request.Option) (*codebuild.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReportGroups(input *codebuild.ListReportGroupsInput) (*codebuild.ListReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReportGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReportGroupsPages(input *codebuild.ListReportGroupsInput, fn func(*codebuild.ListReportGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportGroupsOutput{}
	fnCacher := func(out *codebuild.ListReportGroupsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportGroupsPagesWithContext(ctx context.Context, input *codebuild.ListReportGroupsInput, fn func(*codebuild.ListReportGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportGroupsOutput{}
	fnCacher := func(out *codebuild.ListReportGroupsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportGroupsWithContext(ctx context.Context, input *codebuild.ListReportGroupsInput, opts ...request.Option) (*codebuild.ListReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReportGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReports(input *codebuild.ListReportsInput) (*codebuild.ListReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReportsForReportGroup(input *codebuild.ListReportsForReportGroupInput) (*codebuild.ListReportsForReportGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportsForReportGroupOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReportsForReportGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReportsForReportGroupPages(input *codebuild.ListReportsForReportGroupInput, fn func(*codebuild.ListReportsForReportGroupOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportsForReportGroupOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportsForReportGroupOutput{}
	fnCacher := func(out *codebuild.ListReportsForReportGroupOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportsForReportGroupPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportsForReportGroupPagesWithContext(ctx context.Context, input *codebuild.ListReportsForReportGroupInput, fn func(*codebuild.ListReportsForReportGroupOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportsForReportGroupOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportsForReportGroupOutput{}
	fnCacher := func(out *codebuild.ListReportsForReportGroupOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportsForReportGroupPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportsForReportGroupWithContext(ctx context.Context, input *codebuild.ListReportsForReportGroupInput, opts ...request.Option) (*codebuild.ListReportsForReportGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportsForReportGroupOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReportsForReportGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListReportsPages(input *codebuild.ListReportsInput, fn func(*codebuild.ListReportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportsOutput{}
	fnCacher := func(out *codebuild.ListReportsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportsPagesWithContext(ctx context.Context, input *codebuild.ListReportsInput, fn func(*codebuild.ListReportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListReportsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListReportsOutput{}
	fnCacher := func(out *codebuild.ListReportsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListReportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListReportsWithContext(ctx context.Context, input *codebuild.ListReportsInput, opts ...request.Option) (*codebuild.ListReportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListReportsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListReportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSharedProjects(input *codebuild.ListSharedProjectsInput) (*codebuild.ListSharedProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSharedProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSharedProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSharedProjectsPages(input *codebuild.ListSharedProjectsInput, fn func(*codebuild.ListSharedProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListSharedProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListSharedProjectsOutput{}
	fnCacher := func(out *codebuild.ListSharedProjectsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListSharedProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListSharedProjectsPagesWithContext(ctx context.Context, input *codebuild.ListSharedProjectsInput, fn func(*codebuild.ListSharedProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListSharedProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListSharedProjectsOutput{}
	fnCacher := func(out *codebuild.ListSharedProjectsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListSharedProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListSharedProjectsWithContext(ctx context.Context, input *codebuild.ListSharedProjectsInput, opts ...request.Option) (*codebuild.ListSharedProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSharedProjectsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSharedProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSharedReportGroups(input *codebuild.ListSharedReportGroupsInput) (*codebuild.ListSharedReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSharedReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSharedReportGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSharedReportGroupsPages(input *codebuild.ListSharedReportGroupsInput, fn func(*codebuild.ListSharedReportGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListSharedReportGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListSharedReportGroupsOutput{}
	fnCacher := func(out *codebuild.ListSharedReportGroupsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListSharedReportGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListSharedReportGroupsPagesWithContext(ctx context.Context, input *codebuild.ListSharedReportGroupsInput, fn func(*codebuild.ListSharedReportGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codebuild.ListSharedReportGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codebuild.ListSharedReportGroupsOutput{}
	fnCacher := func(out *codebuild.ListSharedReportGroupsOutput, more bool) bool {
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
	if err := c.CodeBuildAPI.ListSharedReportGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeBuild) ListSharedReportGroupsWithContext(ctx context.Context, input *codebuild.ListSharedReportGroupsInput, opts ...request.Option) (*codebuild.ListSharedReportGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSharedReportGroupsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSharedReportGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSourceCredentials(input *codebuild.ListSourceCredentialsInput) (*codebuild.ListSourceCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSourceCredentialsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSourceCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeBuild) ListSourceCredentialsWithContext(ctx context.Context, input *codebuild.ListSourceCredentialsInput, opts ...request.Option) (*codebuild.ListSourceCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codebuild.ListSourceCredentialsOutput), nil
	}
	out, err := c.CodeBuildAPI.ListSourceCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
