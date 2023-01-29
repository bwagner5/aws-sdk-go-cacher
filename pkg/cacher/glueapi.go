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
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/glue/glueiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Glue struct {
	glueiface.GlueAPI
	cache *cache.Cache
}

func NewGlue(glueapi glueiface.GlueAPI) *Glue {
	return &Glue{
		GlueAPI: glueapi,
		cache:   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Glue) BatchCreatePartition(input *glue.BatchCreatePartitionInput) (*glue.BatchCreatePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchCreatePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchCreatePartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchCreatePartitionWithContext(ctx context.Context, input *glue.BatchCreatePartitionInput, opts ...request.Option) (*glue.BatchCreatePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchCreatePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchCreatePartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteConnection(input *glue.BatchDeleteConnectionInput) (*glue.BatchDeleteConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteConnectionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteConnectionWithContext(ctx context.Context, input *glue.BatchDeleteConnectionInput, opts ...request.Option) (*glue.BatchDeleteConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteConnectionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeletePartition(input *glue.BatchDeletePartitionInput) (*glue.BatchDeletePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeletePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeletePartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeletePartitionWithContext(ctx context.Context, input *glue.BatchDeletePartitionInput, opts ...request.Option) (*glue.BatchDeletePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeletePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeletePartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteTable(input *glue.BatchDeleteTableInput) (*glue.BatchDeleteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteTableOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteTableVersion(input *glue.BatchDeleteTableVersionInput) (*glue.BatchDeleteTableVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteTableVersionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteTableVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteTableVersionWithContext(ctx context.Context, input *glue.BatchDeleteTableVersionInput, opts ...request.Option) (*glue.BatchDeleteTableVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteTableVersionOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteTableVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchDeleteTableWithContext(ctx context.Context, input *glue.BatchDeleteTableInput, opts ...request.Option) (*glue.BatchDeleteTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchDeleteTableOutput), nil
	}
	out, err := c.GlueAPI.BatchDeleteTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetBlueprints(input *glue.BatchGetBlueprintsInput) (*glue.BatchGetBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetBlueprintsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetBlueprints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetBlueprintsWithContext(ctx context.Context, input *glue.BatchGetBlueprintsInput, opts ...request.Option) (*glue.BatchGetBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetBlueprintsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetBlueprintsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetCrawlers(input *glue.BatchGetCrawlersInput) (*glue.BatchGetCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetCrawlersOutput), nil
	}
	out, err := c.GlueAPI.BatchGetCrawlers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetCrawlersWithContext(ctx context.Context, input *glue.BatchGetCrawlersInput, opts ...request.Option) (*glue.BatchGetCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetCrawlersOutput), nil
	}
	out, err := c.GlueAPI.BatchGetCrawlersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetCustomEntityTypes(input *glue.BatchGetCustomEntityTypesInput) (*glue.BatchGetCustomEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetCustomEntityTypesOutput), nil
	}
	out, err := c.GlueAPI.BatchGetCustomEntityTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetCustomEntityTypesWithContext(ctx context.Context, input *glue.BatchGetCustomEntityTypesInput, opts ...request.Option) (*glue.BatchGetCustomEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetCustomEntityTypesOutput), nil
	}
	out, err := c.GlueAPI.BatchGetCustomEntityTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetDataQualityResult(input *glue.BatchGetDataQualityResultInput) (*glue.BatchGetDataQualityResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetDataQualityResultOutput), nil
	}
	out, err := c.GlueAPI.BatchGetDataQualityResult(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetDataQualityResultWithContext(ctx context.Context, input *glue.BatchGetDataQualityResultInput, opts ...request.Option) (*glue.BatchGetDataQualityResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetDataQualityResultOutput), nil
	}
	out, err := c.GlueAPI.BatchGetDataQualityResultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetDevEndpoints(input *glue.BatchGetDevEndpointsInput) (*glue.BatchGetDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetDevEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetDevEndpointsWithContext(ctx context.Context, input *glue.BatchGetDevEndpointsInput, opts ...request.Option) (*glue.BatchGetDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetDevEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetJobs(input *glue.BatchGetJobsInput) (*glue.BatchGetJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetJobsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetJobsWithContext(ctx context.Context, input *glue.BatchGetJobsInput, opts ...request.Option) (*glue.BatchGetJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetJobsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetPartition(input *glue.BatchGetPartitionInput) (*glue.BatchGetPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetPartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchGetPartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetPartitionWithContext(ctx context.Context, input *glue.BatchGetPartitionInput, opts ...request.Option) (*glue.BatchGetPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetPartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchGetPartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetTriggers(input *glue.BatchGetTriggersInput) (*glue.BatchGetTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetTriggersOutput), nil
	}
	out, err := c.GlueAPI.BatchGetTriggers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetTriggersWithContext(ctx context.Context, input *glue.BatchGetTriggersInput, opts ...request.Option) (*glue.BatchGetTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetTriggersOutput), nil
	}
	out, err := c.GlueAPI.BatchGetTriggersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetWorkflows(input *glue.BatchGetWorkflowsInput) (*glue.BatchGetWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetWorkflowsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetWorkflows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchGetWorkflowsWithContext(ctx context.Context, input *glue.BatchGetWorkflowsInput, opts ...request.Option) (*glue.BatchGetWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchGetWorkflowsOutput), nil
	}
	out, err := c.GlueAPI.BatchGetWorkflowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchStopJobRun(input *glue.BatchStopJobRunInput) (*glue.BatchStopJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchStopJobRunOutput), nil
	}
	out, err := c.GlueAPI.BatchStopJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchStopJobRunWithContext(ctx context.Context, input *glue.BatchStopJobRunInput, opts ...request.Option) (*glue.BatchStopJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchStopJobRunOutput), nil
	}
	out, err := c.GlueAPI.BatchStopJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchUpdatePartition(input *glue.BatchUpdatePartitionInput) (*glue.BatchUpdatePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchUpdatePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchUpdatePartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) BatchUpdatePartitionWithContext(ctx context.Context, input *glue.BatchUpdatePartitionInput, opts ...request.Option) (*glue.BatchUpdatePartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.BatchUpdatePartitionOutput), nil
	}
	out, err := c.GlueAPI.BatchUpdatePartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprint(input *glue.GetBlueprintInput) (*glue.GetBlueprintOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprintRun(input *glue.GetBlueprintRunInput) (*glue.GetBlueprintRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintRunOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprintRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprintRunWithContext(ctx context.Context, input *glue.GetBlueprintRunInput, opts ...request.Option) (*glue.GetBlueprintRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintRunOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprintRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprintRuns(input *glue.GetBlueprintRunsInput) (*glue.GetBlueprintRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintRunsOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprintRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprintRunsPages(input *glue.GetBlueprintRunsInput, fn func(*glue.GetBlueprintRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetBlueprintRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetBlueprintRunsOutput{}
	fnCacher := func(out *glue.GetBlueprintRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetBlueprintRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetBlueprintRunsPagesWithContext(ctx context.Context, input *glue.GetBlueprintRunsInput, fn func(*glue.GetBlueprintRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetBlueprintRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetBlueprintRunsOutput{}
	fnCacher := func(out *glue.GetBlueprintRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetBlueprintRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetBlueprintRunsWithContext(ctx context.Context, input *glue.GetBlueprintRunsInput, opts ...request.Option) (*glue.GetBlueprintRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintRunsOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprintRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetBlueprintWithContext(ctx context.Context, input *glue.GetBlueprintInput, opts ...request.Option) (*glue.GetBlueprintOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetBlueprintOutput), nil
	}
	out, err := c.GlueAPI.GetBlueprintWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCatalogImportStatus(input *glue.GetCatalogImportStatusInput) (*glue.GetCatalogImportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCatalogImportStatusOutput), nil
	}
	out, err := c.GlueAPI.GetCatalogImportStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCatalogImportStatusWithContext(ctx context.Context, input *glue.GetCatalogImportStatusInput, opts ...request.Option) (*glue.GetCatalogImportStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCatalogImportStatusOutput), nil
	}
	out, err := c.GlueAPI.GetCatalogImportStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetClassifier(input *glue.GetClassifierInput) (*glue.GetClassifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetClassifierOutput), nil
	}
	out, err := c.GlueAPI.GetClassifier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetClassifierWithContext(ctx context.Context, input *glue.GetClassifierInput, opts ...request.Option) (*glue.GetClassifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetClassifierOutput), nil
	}
	out, err := c.GlueAPI.GetClassifierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetClassifiers(input *glue.GetClassifiersInput) (*glue.GetClassifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetClassifiersOutput), nil
	}
	out, err := c.GlueAPI.GetClassifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetClassifiersPages(input *glue.GetClassifiersInput, fn func(*glue.GetClassifiersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetClassifiersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetClassifiersOutput{}
	fnCacher := func(out *glue.GetClassifiersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetClassifiersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetClassifiersPagesWithContext(ctx context.Context, input *glue.GetClassifiersInput, fn func(*glue.GetClassifiersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetClassifiersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetClassifiersOutput{}
	fnCacher := func(out *glue.GetClassifiersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetClassifiersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetClassifiersWithContext(ctx context.Context, input *glue.GetClassifiersInput, opts ...request.Option) (*glue.GetClassifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetClassifiersOutput), nil
	}
	out, err := c.GlueAPI.GetClassifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetColumnStatisticsForPartition(input *glue.GetColumnStatisticsForPartitionInput) (*glue.GetColumnStatisticsForPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetColumnStatisticsForPartitionOutput), nil
	}
	out, err := c.GlueAPI.GetColumnStatisticsForPartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetColumnStatisticsForPartitionWithContext(ctx context.Context, input *glue.GetColumnStatisticsForPartitionInput, opts ...request.Option) (*glue.GetColumnStatisticsForPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetColumnStatisticsForPartitionOutput), nil
	}
	out, err := c.GlueAPI.GetColumnStatisticsForPartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetColumnStatisticsForTable(input *glue.GetColumnStatisticsForTableInput) (*glue.GetColumnStatisticsForTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetColumnStatisticsForTableOutput), nil
	}
	out, err := c.GlueAPI.GetColumnStatisticsForTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetColumnStatisticsForTableWithContext(ctx context.Context, input *glue.GetColumnStatisticsForTableInput, opts ...request.Option) (*glue.GetColumnStatisticsForTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetColumnStatisticsForTableOutput), nil
	}
	out, err := c.GlueAPI.GetColumnStatisticsForTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetConnection(input *glue.GetConnectionInput) (*glue.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetConnectionOutput), nil
	}
	out, err := c.GlueAPI.GetConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetConnectionWithContext(ctx context.Context, input *glue.GetConnectionInput, opts ...request.Option) (*glue.GetConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetConnectionOutput), nil
	}
	out, err := c.GlueAPI.GetConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetConnections(input *glue.GetConnectionsInput) (*glue.GetConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetConnectionsOutput), nil
	}
	out, err := c.GlueAPI.GetConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetConnectionsPages(input *glue.GetConnectionsInput, fn func(*glue.GetConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetConnectionsOutput{}
	fnCacher := func(out *glue.GetConnectionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetConnectionsPagesWithContext(ctx context.Context, input *glue.GetConnectionsInput, fn func(*glue.GetConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetConnectionsOutput{}
	fnCacher := func(out *glue.GetConnectionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetConnectionsWithContext(ctx context.Context, input *glue.GetConnectionsInput, opts ...request.Option) (*glue.GetConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetConnectionsOutput), nil
	}
	out, err := c.GlueAPI.GetConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawler(input *glue.GetCrawlerInput) (*glue.GetCrawlerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlerOutput), nil
	}
	out, err := c.GlueAPI.GetCrawler(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawlerMetrics(input *glue.GetCrawlerMetricsInput) (*glue.GetCrawlerMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlerMetricsOutput), nil
	}
	out, err := c.GlueAPI.GetCrawlerMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawlerMetricsPages(input *glue.GetCrawlerMetricsInput, fn func(*glue.GetCrawlerMetricsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetCrawlerMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetCrawlerMetricsOutput{}
	fnCacher := func(out *glue.GetCrawlerMetricsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetCrawlerMetricsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetCrawlerMetricsPagesWithContext(ctx context.Context, input *glue.GetCrawlerMetricsInput, fn func(*glue.GetCrawlerMetricsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetCrawlerMetricsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetCrawlerMetricsOutput{}
	fnCacher := func(out *glue.GetCrawlerMetricsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetCrawlerMetricsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetCrawlerMetricsWithContext(ctx context.Context, input *glue.GetCrawlerMetricsInput, opts ...request.Option) (*glue.GetCrawlerMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlerMetricsOutput), nil
	}
	out, err := c.GlueAPI.GetCrawlerMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawlerWithContext(ctx context.Context, input *glue.GetCrawlerInput, opts ...request.Option) (*glue.GetCrawlerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlerOutput), nil
	}
	out, err := c.GlueAPI.GetCrawlerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawlers(input *glue.GetCrawlersInput) (*glue.GetCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlersOutput), nil
	}
	out, err := c.GlueAPI.GetCrawlers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCrawlersPages(input *glue.GetCrawlersInput, fn func(*glue.GetCrawlersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetCrawlersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetCrawlersOutput{}
	fnCacher := func(out *glue.GetCrawlersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetCrawlersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetCrawlersPagesWithContext(ctx context.Context, input *glue.GetCrawlersInput, fn func(*glue.GetCrawlersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetCrawlersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetCrawlersOutput{}
	fnCacher := func(out *glue.GetCrawlersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetCrawlersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetCrawlersWithContext(ctx context.Context, input *glue.GetCrawlersInput, opts ...request.Option) (*glue.GetCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCrawlersOutput), nil
	}
	out, err := c.GlueAPI.GetCrawlersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCustomEntityType(input *glue.GetCustomEntityTypeInput) (*glue.GetCustomEntityTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCustomEntityTypeOutput), nil
	}
	out, err := c.GlueAPI.GetCustomEntityType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetCustomEntityTypeWithContext(ctx context.Context, input *glue.GetCustomEntityTypeInput, opts ...request.Option) (*glue.GetCustomEntityTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetCustomEntityTypeOutput), nil
	}
	out, err := c.GlueAPI.GetCustomEntityTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataCatalogEncryptionSettings(input *glue.GetDataCatalogEncryptionSettingsInput) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataCatalogEncryptionSettingsOutput), nil
	}
	out, err := c.GlueAPI.GetDataCatalogEncryptionSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataCatalogEncryptionSettingsWithContext(ctx context.Context, input *glue.GetDataCatalogEncryptionSettingsInput, opts ...request.Option) (*glue.GetDataCatalogEncryptionSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataCatalogEncryptionSettingsOutput), nil
	}
	out, err := c.GlueAPI.GetDataCatalogEncryptionSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityResult(input *glue.GetDataQualityResultInput) (*glue.GetDataQualityResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityResultOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityResult(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityResultWithContext(ctx context.Context, input *glue.GetDataQualityResultInput, opts ...request.Option) (*glue.GetDataQualityResultOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityResultOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityResultWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRuleRecommendationRun(input *glue.GetDataQualityRuleRecommendationRunInput) (*glue.GetDataQualityRuleRecommendationRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRuleRecommendationRunOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRuleRecommendationRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRuleRecommendationRunWithContext(ctx context.Context, input *glue.GetDataQualityRuleRecommendationRunInput, opts ...request.Option) (*glue.GetDataQualityRuleRecommendationRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRuleRecommendationRunOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRuleRecommendationRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRuleset(input *glue.GetDataQualityRulesetInput) (*glue.GetDataQualityRulesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRulesetOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRuleset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRulesetEvaluationRun(input *glue.GetDataQualityRulesetEvaluationRunInput) (*glue.GetDataQualityRulesetEvaluationRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRulesetEvaluationRunOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRulesetEvaluationRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRulesetEvaluationRunWithContext(ctx context.Context, input *glue.GetDataQualityRulesetEvaluationRunInput, opts ...request.Option) (*glue.GetDataQualityRulesetEvaluationRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRulesetEvaluationRunOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRulesetEvaluationRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataQualityRulesetWithContext(ctx context.Context, input *glue.GetDataQualityRulesetInput, opts ...request.Option) (*glue.GetDataQualityRulesetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataQualityRulesetOutput), nil
	}
	out, err := c.GlueAPI.GetDataQualityRulesetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDatabase(input *glue.GetDatabaseInput) (*glue.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDatabaseOutput), nil
	}
	out, err := c.GlueAPI.GetDatabase(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDatabaseWithContext(ctx context.Context, input *glue.GetDatabaseInput, opts ...request.Option) (*glue.GetDatabaseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDatabaseOutput), nil
	}
	out, err := c.GlueAPI.GetDatabaseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDatabases(input *glue.GetDatabasesInput) (*glue.GetDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDatabasesOutput), nil
	}
	out, err := c.GlueAPI.GetDatabases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDatabasesPages(input *glue.GetDatabasesInput, fn func(*glue.GetDatabasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetDatabasesOutput{}
	fnCacher := func(out *glue.GetDatabasesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetDatabasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetDatabasesPagesWithContext(ctx context.Context, input *glue.GetDatabasesInput, fn func(*glue.GetDatabasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetDatabasesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetDatabasesOutput{}
	fnCacher := func(out *glue.GetDatabasesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetDatabasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetDatabasesWithContext(ctx context.Context, input *glue.GetDatabasesInput, opts ...request.Option) (*glue.GetDatabasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDatabasesOutput), nil
	}
	out, err := c.GlueAPI.GetDatabasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataflowGraph(input *glue.GetDataflowGraphInput) (*glue.GetDataflowGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataflowGraphOutput), nil
	}
	out, err := c.GlueAPI.GetDataflowGraph(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDataflowGraphWithContext(ctx context.Context, input *glue.GetDataflowGraphInput, opts ...request.Option) (*glue.GetDataflowGraphOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDataflowGraphOutput), nil
	}
	out, err := c.GlueAPI.GetDataflowGraphWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDevEndpoint(input *glue.GetDevEndpointInput) (*glue.GetDevEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDevEndpointOutput), nil
	}
	out, err := c.GlueAPI.GetDevEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDevEndpointWithContext(ctx context.Context, input *glue.GetDevEndpointInput, opts ...request.Option) (*glue.GetDevEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDevEndpointOutput), nil
	}
	out, err := c.GlueAPI.GetDevEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDevEndpoints(input *glue.GetDevEndpointsInput) (*glue.GetDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.GetDevEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetDevEndpointsPages(input *glue.GetDevEndpointsInput, fn func(*glue.GetDevEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetDevEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetDevEndpointsOutput{}
	fnCacher := func(out *glue.GetDevEndpointsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetDevEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetDevEndpointsPagesWithContext(ctx context.Context, input *glue.GetDevEndpointsInput, fn func(*glue.GetDevEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetDevEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetDevEndpointsOutput{}
	fnCacher := func(out *glue.GetDevEndpointsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetDevEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetDevEndpointsWithContext(ctx context.Context, input *glue.GetDevEndpointsInput, opts ...request.Option) (*glue.GetDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.GetDevEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJob(input *glue.GetJobInput) (*glue.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobOutput), nil
	}
	out, err := c.GlueAPI.GetJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobBookmark(input *glue.GetJobBookmarkInput) (*glue.GetJobBookmarkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobBookmarkOutput), nil
	}
	out, err := c.GlueAPI.GetJobBookmark(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobBookmarkWithContext(ctx context.Context, input *glue.GetJobBookmarkInput, opts ...request.Option) (*glue.GetJobBookmarkOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobBookmarkOutput), nil
	}
	out, err := c.GlueAPI.GetJobBookmarkWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobRun(input *glue.GetJobRunInput) (*glue.GetJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobRunOutput), nil
	}
	out, err := c.GlueAPI.GetJobRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobRunWithContext(ctx context.Context, input *glue.GetJobRunInput, opts ...request.Option) (*glue.GetJobRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobRunOutput), nil
	}
	out, err := c.GlueAPI.GetJobRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobRuns(input *glue.GetJobRunsInput) (*glue.GetJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobRunsOutput), nil
	}
	out, err := c.GlueAPI.GetJobRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobRunsPages(input *glue.GetJobRunsInput, fn func(*glue.GetJobRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetJobRunsOutput{}
	fnCacher := func(out *glue.GetJobRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetJobRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetJobRunsPagesWithContext(ctx context.Context, input *glue.GetJobRunsInput, fn func(*glue.GetJobRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetJobRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetJobRunsOutput{}
	fnCacher := func(out *glue.GetJobRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetJobRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetJobRunsWithContext(ctx context.Context, input *glue.GetJobRunsInput, opts ...request.Option) (*glue.GetJobRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobRunsOutput), nil
	}
	out, err := c.GlueAPI.GetJobRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobWithContext(ctx context.Context, input *glue.GetJobInput, opts ...request.Option) (*glue.GetJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobOutput), nil
	}
	out, err := c.GlueAPI.GetJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobs(input *glue.GetJobsInput) (*glue.GetJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobsOutput), nil
	}
	out, err := c.GlueAPI.GetJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetJobsPages(input *glue.GetJobsInput, fn func(*glue.GetJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetJobsOutput{}
	fnCacher := func(out *glue.GetJobsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetJobsPagesWithContext(ctx context.Context, input *glue.GetJobsInput, fn func(*glue.GetJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetJobsOutput{}
	fnCacher := func(out *glue.GetJobsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetJobsWithContext(ctx context.Context, input *glue.GetJobsInput, opts ...request.Option) (*glue.GetJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetJobsOutput), nil
	}
	out, err := c.GlueAPI.GetJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTaskRun(input *glue.GetMLTaskRunInput) (*glue.GetMLTaskRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTaskRunOutput), nil
	}
	out, err := c.GlueAPI.GetMLTaskRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTaskRunWithContext(ctx context.Context, input *glue.GetMLTaskRunInput, opts ...request.Option) (*glue.GetMLTaskRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTaskRunOutput), nil
	}
	out, err := c.GlueAPI.GetMLTaskRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTaskRuns(input *glue.GetMLTaskRunsInput) (*glue.GetMLTaskRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTaskRunsOutput), nil
	}
	out, err := c.GlueAPI.GetMLTaskRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTaskRunsPages(input *glue.GetMLTaskRunsInput, fn func(*glue.GetMLTaskRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetMLTaskRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetMLTaskRunsOutput{}
	fnCacher := func(out *glue.GetMLTaskRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetMLTaskRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetMLTaskRunsPagesWithContext(ctx context.Context, input *glue.GetMLTaskRunsInput, fn func(*glue.GetMLTaskRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetMLTaskRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetMLTaskRunsOutput{}
	fnCacher := func(out *glue.GetMLTaskRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetMLTaskRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetMLTaskRunsWithContext(ctx context.Context, input *glue.GetMLTaskRunsInput, opts ...request.Option) (*glue.GetMLTaskRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTaskRunsOutput), nil
	}
	out, err := c.GlueAPI.GetMLTaskRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTransform(input *glue.GetMLTransformInput) (*glue.GetMLTransformOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTransformOutput), nil
	}
	out, err := c.GlueAPI.GetMLTransform(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTransformWithContext(ctx context.Context, input *glue.GetMLTransformInput, opts ...request.Option) (*glue.GetMLTransformOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTransformOutput), nil
	}
	out, err := c.GlueAPI.GetMLTransformWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTransforms(input *glue.GetMLTransformsInput) (*glue.GetMLTransformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTransformsOutput), nil
	}
	out, err := c.GlueAPI.GetMLTransforms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMLTransformsPages(input *glue.GetMLTransformsInput, fn func(*glue.GetMLTransformsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetMLTransformsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetMLTransformsOutput{}
	fnCacher := func(out *glue.GetMLTransformsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetMLTransformsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetMLTransformsPagesWithContext(ctx context.Context, input *glue.GetMLTransformsInput, fn func(*glue.GetMLTransformsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetMLTransformsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetMLTransformsOutput{}
	fnCacher := func(out *glue.GetMLTransformsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetMLTransformsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetMLTransformsWithContext(ctx context.Context, input *glue.GetMLTransformsInput, opts ...request.Option) (*glue.GetMLTransformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMLTransformsOutput), nil
	}
	out, err := c.GlueAPI.GetMLTransformsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMapping(input *glue.GetMappingInput) (*glue.GetMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMappingOutput), nil
	}
	out, err := c.GlueAPI.GetMapping(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetMappingWithContext(ctx context.Context, input *glue.GetMappingInput, opts ...request.Option) (*glue.GetMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetMappingOutput), nil
	}
	out, err := c.GlueAPI.GetMappingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartition(input *glue.GetPartitionInput) (*glue.GetPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionOutput), nil
	}
	out, err := c.GlueAPI.GetPartition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartitionIndexes(input *glue.GetPartitionIndexesInput) (*glue.GetPartitionIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionIndexesOutput), nil
	}
	out, err := c.GlueAPI.GetPartitionIndexes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartitionIndexesPages(input *glue.GetPartitionIndexesInput, fn func(*glue.GetPartitionIndexesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetPartitionIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetPartitionIndexesOutput{}
	fnCacher := func(out *glue.GetPartitionIndexesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetPartitionIndexesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetPartitionIndexesPagesWithContext(ctx context.Context, input *glue.GetPartitionIndexesInput, fn func(*glue.GetPartitionIndexesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetPartitionIndexesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetPartitionIndexesOutput{}
	fnCacher := func(out *glue.GetPartitionIndexesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetPartitionIndexesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetPartitionIndexesWithContext(ctx context.Context, input *glue.GetPartitionIndexesInput, opts ...request.Option) (*glue.GetPartitionIndexesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionIndexesOutput), nil
	}
	out, err := c.GlueAPI.GetPartitionIndexesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartitionWithContext(ctx context.Context, input *glue.GetPartitionInput, opts ...request.Option) (*glue.GetPartitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionOutput), nil
	}
	out, err := c.GlueAPI.GetPartitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartitions(input *glue.GetPartitionsInput) (*glue.GetPartitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionsOutput), nil
	}
	out, err := c.GlueAPI.GetPartitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPartitionsPages(input *glue.GetPartitionsInput, fn func(*glue.GetPartitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetPartitionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetPartitionsOutput{}
	fnCacher := func(out *glue.GetPartitionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetPartitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetPartitionsPagesWithContext(ctx context.Context, input *glue.GetPartitionsInput, fn func(*glue.GetPartitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetPartitionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetPartitionsOutput{}
	fnCacher := func(out *glue.GetPartitionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetPartitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetPartitionsWithContext(ctx context.Context, input *glue.GetPartitionsInput, opts ...request.Option) (*glue.GetPartitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPartitionsOutput), nil
	}
	out, err := c.GlueAPI.GetPartitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPlan(input *glue.GetPlanInput) (*glue.GetPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPlanOutput), nil
	}
	out, err := c.GlueAPI.GetPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetPlanWithContext(ctx context.Context, input *glue.GetPlanInput, opts ...request.Option) (*glue.GetPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetPlanOutput), nil
	}
	out, err := c.GlueAPI.GetPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetRegistry(input *glue.GetRegistryInput) (*glue.GetRegistryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetRegistryOutput), nil
	}
	out, err := c.GlueAPI.GetRegistry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetRegistryWithContext(ctx context.Context, input *glue.GetRegistryInput, opts ...request.Option) (*glue.GetRegistryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetRegistryOutput), nil
	}
	out, err := c.GlueAPI.GetRegistryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetResourcePolicies(input *glue.GetResourcePoliciesInput) (*glue.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetResourcePoliciesOutput), nil
	}
	out, err := c.GlueAPI.GetResourcePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetResourcePoliciesPages(input *glue.GetResourcePoliciesInput, fn func(*glue.GetResourcePoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetResourcePoliciesOutput{}
	fnCacher := func(out *glue.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetResourcePoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetResourcePoliciesPagesWithContext(ctx context.Context, input *glue.GetResourcePoliciesInput, fn func(*glue.GetResourcePoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetResourcePoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetResourcePoliciesOutput{}
	fnCacher := func(out *glue.GetResourcePoliciesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetResourcePoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetResourcePoliciesWithContext(ctx context.Context, input *glue.GetResourcePoliciesInput, opts ...request.Option) (*glue.GetResourcePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetResourcePoliciesOutput), nil
	}
	out, err := c.GlueAPI.GetResourcePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetResourcePolicy(input *glue.GetResourcePolicyInput) (*glue.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetResourcePolicyOutput), nil
	}
	out, err := c.GlueAPI.GetResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetResourcePolicyWithContext(ctx context.Context, input *glue.GetResourcePolicyInput, opts ...request.Option) (*glue.GetResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetResourcePolicyOutput), nil
	}
	out, err := c.GlueAPI.GetResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchema(input *glue.GetSchemaInput) (*glue.GetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaOutput), nil
	}
	out, err := c.GlueAPI.GetSchema(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaByDefinition(input *glue.GetSchemaByDefinitionInput) (*glue.GetSchemaByDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaByDefinitionOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaByDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaByDefinitionWithContext(ctx context.Context, input *glue.GetSchemaByDefinitionInput, opts ...request.Option) (*glue.GetSchemaByDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaByDefinitionOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaByDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaVersion(input *glue.GetSchemaVersionInput) (*glue.GetSchemaVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaVersionOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaVersionWithContext(ctx context.Context, input *glue.GetSchemaVersionInput, opts ...request.Option) (*glue.GetSchemaVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaVersionOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaVersionsDiff(input *glue.GetSchemaVersionsDiffInput) (*glue.GetSchemaVersionsDiffOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaVersionsDiffOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaVersionsDiff(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaVersionsDiffWithContext(ctx context.Context, input *glue.GetSchemaVersionsDiffInput, opts ...request.Option) (*glue.GetSchemaVersionsDiffOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaVersionsDiffOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaVersionsDiffWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSchemaWithContext(ctx context.Context, input *glue.GetSchemaInput, opts ...request.Option) (*glue.GetSchemaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSchemaOutput), nil
	}
	out, err := c.GlueAPI.GetSchemaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSecurityConfiguration(input *glue.GetSecurityConfigurationInput) (*glue.GetSecurityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSecurityConfigurationOutput), nil
	}
	out, err := c.GlueAPI.GetSecurityConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSecurityConfigurationWithContext(ctx context.Context, input *glue.GetSecurityConfigurationInput, opts ...request.Option) (*glue.GetSecurityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSecurityConfigurationOutput), nil
	}
	out, err := c.GlueAPI.GetSecurityConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSecurityConfigurations(input *glue.GetSecurityConfigurationsInput) (*glue.GetSecurityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSecurityConfigurationsOutput), nil
	}
	out, err := c.GlueAPI.GetSecurityConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSecurityConfigurationsPages(input *glue.GetSecurityConfigurationsInput, fn func(*glue.GetSecurityConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetSecurityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetSecurityConfigurationsOutput{}
	fnCacher := func(out *glue.GetSecurityConfigurationsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetSecurityConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetSecurityConfigurationsPagesWithContext(ctx context.Context, input *glue.GetSecurityConfigurationsInput, fn func(*glue.GetSecurityConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetSecurityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetSecurityConfigurationsOutput{}
	fnCacher := func(out *glue.GetSecurityConfigurationsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetSecurityConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetSecurityConfigurationsWithContext(ctx context.Context, input *glue.GetSecurityConfigurationsInput, opts ...request.Option) (*glue.GetSecurityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSecurityConfigurationsOutput), nil
	}
	out, err := c.GlueAPI.GetSecurityConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSession(input *glue.GetSessionInput) (*glue.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSessionOutput), nil
	}
	out, err := c.GlueAPI.GetSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetSessionWithContext(ctx context.Context, input *glue.GetSessionInput, opts ...request.Option) (*glue.GetSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetSessionOutput), nil
	}
	out, err := c.GlueAPI.GetSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetStatement(input *glue.GetStatementInput) (*glue.GetStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetStatementOutput), nil
	}
	out, err := c.GlueAPI.GetStatement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetStatementWithContext(ctx context.Context, input *glue.GetStatementInput, opts ...request.Option) (*glue.GetStatementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetStatementOutput), nil
	}
	out, err := c.GlueAPI.GetStatementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTable(input *glue.GetTableInput) (*glue.GetTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableOutput), nil
	}
	out, err := c.GlueAPI.GetTable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTableVersion(input *glue.GetTableVersionInput) (*glue.GetTableVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableVersionOutput), nil
	}
	out, err := c.GlueAPI.GetTableVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTableVersionWithContext(ctx context.Context, input *glue.GetTableVersionInput, opts ...request.Option) (*glue.GetTableVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableVersionOutput), nil
	}
	out, err := c.GlueAPI.GetTableVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTableVersions(input *glue.GetTableVersionsInput) (*glue.GetTableVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableVersionsOutput), nil
	}
	out, err := c.GlueAPI.GetTableVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTableVersionsPages(input *glue.GetTableVersionsInput, fn func(*glue.GetTableVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTableVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTableVersionsOutput{}
	fnCacher := func(out *glue.GetTableVersionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTableVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTableVersionsPagesWithContext(ctx context.Context, input *glue.GetTableVersionsInput, fn func(*glue.GetTableVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTableVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTableVersionsOutput{}
	fnCacher := func(out *glue.GetTableVersionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTableVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTableVersionsWithContext(ctx context.Context, input *glue.GetTableVersionsInput, opts ...request.Option) (*glue.GetTableVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableVersionsOutput), nil
	}
	out, err := c.GlueAPI.GetTableVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTableWithContext(ctx context.Context, input *glue.GetTableInput, opts ...request.Option) (*glue.GetTableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTableOutput), nil
	}
	out, err := c.GlueAPI.GetTableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTables(input *glue.GetTablesInput) (*glue.GetTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTablesOutput), nil
	}
	out, err := c.GlueAPI.GetTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTablesPages(input *glue.GetTablesInput, fn func(*glue.GetTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTablesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTablesOutput{}
	fnCacher := func(out *glue.GetTablesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTablesPagesWithContext(ctx context.Context, input *glue.GetTablesInput, fn func(*glue.GetTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTablesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTablesOutput{}
	fnCacher := func(out *glue.GetTablesOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTablesWithContext(ctx context.Context, input *glue.GetTablesInput, opts ...request.Option) (*glue.GetTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTablesOutput), nil
	}
	out, err := c.GlueAPI.GetTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTags(input *glue.GetTagsInput) (*glue.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTagsOutput), nil
	}
	out, err := c.GlueAPI.GetTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTagsWithContext(ctx context.Context, input *glue.GetTagsInput, opts ...request.Option) (*glue.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTagsOutput), nil
	}
	out, err := c.GlueAPI.GetTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTrigger(input *glue.GetTriggerInput) (*glue.GetTriggerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTriggerOutput), nil
	}
	out, err := c.GlueAPI.GetTrigger(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTriggerWithContext(ctx context.Context, input *glue.GetTriggerInput, opts ...request.Option) (*glue.GetTriggerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTriggerOutput), nil
	}
	out, err := c.GlueAPI.GetTriggerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTriggers(input *glue.GetTriggersInput) (*glue.GetTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTriggersOutput), nil
	}
	out, err := c.GlueAPI.GetTriggers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetTriggersPages(input *glue.GetTriggersInput, fn func(*glue.GetTriggersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTriggersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTriggersOutput{}
	fnCacher := func(out *glue.GetTriggersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTriggersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTriggersPagesWithContext(ctx context.Context, input *glue.GetTriggersInput, fn func(*glue.GetTriggersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetTriggersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetTriggersOutput{}
	fnCacher := func(out *glue.GetTriggersOutput, more bool) bool {
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
	if err := c.GlueAPI.GetTriggersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetTriggersWithContext(ctx context.Context, input *glue.GetTriggersInput, opts ...request.Option) (*glue.GetTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetTriggersOutput), nil
	}
	out, err := c.GlueAPI.GetTriggersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredPartitionMetadata(input *glue.GetUnfilteredPartitionMetadataInput) (*glue.GetUnfilteredPartitionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredPartitionMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredPartitionMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredPartitionMetadataWithContext(ctx context.Context, input *glue.GetUnfilteredPartitionMetadataInput, opts ...request.Option) (*glue.GetUnfilteredPartitionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredPartitionMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredPartitionMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredPartitionsMetadata(input *glue.GetUnfilteredPartitionsMetadataInput) (*glue.GetUnfilteredPartitionsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredPartitionsMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredPartitionsMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredPartitionsMetadataPages(input *glue.GetUnfilteredPartitionsMetadataInput, fn func(*glue.GetUnfilteredPartitionsMetadataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetUnfilteredPartitionsMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetUnfilteredPartitionsMetadataOutput{}
	fnCacher := func(out *glue.GetUnfilteredPartitionsMetadataOutput, more bool) bool {
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
	if err := c.GlueAPI.GetUnfilteredPartitionsMetadataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetUnfilteredPartitionsMetadataPagesWithContext(ctx context.Context, input *glue.GetUnfilteredPartitionsMetadataInput, fn func(*glue.GetUnfilteredPartitionsMetadataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetUnfilteredPartitionsMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetUnfilteredPartitionsMetadataOutput{}
	fnCacher := func(out *glue.GetUnfilteredPartitionsMetadataOutput, more bool) bool {
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
	if err := c.GlueAPI.GetUnfilteredPartitionsMetadataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetUnfilteredPartitionsMetadataWithContext(ctx context.Context, input *glue.GetUnfilteredPartitionsMetadataInput, opts ...request.Option) (*glue.GetUnfilteredPartitionsMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredPartitionsMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredPartitionsMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredTableMetadata(input *glue.GetUnfilteredTableMetadataInput) (*glue.GetUnfilteredTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredTableMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredTableMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUnfilteredTableMetadataWithContext(ctx context.Context, input *glue.GetUnfilteredTableMetadataInput, opts ...request.Option) (*glue.GetUnfilteredTableMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUnfilteredTableMetadataOutput), nil
	}
	out, err := c.GlueAPI.GetUnfilteredTableMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUserDefinedFunction(input *glue.GetUserDefinedFunctionInput) (*glue.GetUserDefinedFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUserDefinedFunctionOutput), nil
	}
	out, err := c.GlueAPI.GetUserDefinedFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUserDefinedFunctionWithContext(ctx context.Context, input *glue.GetUserDefinedFunctionInput, opts ...request.Option) (*glue.GetUserDefinedFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUserDefinedFunctionOutput), nil
	}
	out, err := c.GlueAPI.GetUserDefinedFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUserDefinedFunctions(input *glue.GetUserDefinedFunctionsInput) (*glue.GetUserDefinedFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUserDefinedFunctionsOutput), nil
	}
	out, err := c.GlueAPI.GetUserDefinedFunctions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetUserDefinedFunctionsPages(input *glue.GetUserDefinedFunctionsInput, fn func(*glue.GetUserDefinedFunctionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetUserDefinedFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetUserDefinedFunctionsOutput{}
	fnCacher := func(out *glue.GetUserDefinedFunctionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetUserDefinedFunctionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetUserDefinedFunctionsPagesWithContext(ctx context.Context, input *glue.GetUserDefinedFunctionsInput, fn func(*glue.GetUserDefinedFunctionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetUserDefinedFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetUserDefinedFunctionsOutput{}
	fnCacher := func(out *glue.GetUserDefinedFunctionsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetUserDefinedFunctionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetUserDefinedFunctionsWithContext(ctx context.Context, input *glue.GetUserDefinedFunctionsInput, opts ...request.Option) (*glue.GetUserDefinedFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetUserDefinedFunctionsOutput), nil
	}
	out, err := c.GlueAPI.GetUserDefinedFunctionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflow(input *glue.GetWorkflowInput) (*glue.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRun(input *glue.GetWorkflowRunInput) (*glue.GetWorkflowRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRun(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRunProperties(input *glue.GetWorkflowRunPropertiesInput) (*glue.GetWorkflowRunPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunPropertiesOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRunProperties(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRunPropertiesWithContext(ctx context.Context, input *glue.GetWorkflowRunPropertiesInput, opts ...request.Option) (*glue.GetWorkflowRunPropertiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunPropertiesOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRunPropertiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRunWithContext(ctx context.Context, input *glue.GetWorkflowRunInput, opts ...request.Option) (*glue.GetWorkflowRunOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRunWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRuns(input *glue.GetWorkflowRunsInput) (*glue.GetWorkflowRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunsOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowRunsPages(input *glue.GetWorkflowRunsInput, fn func(*glue.GetWorkflowRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetWorkflowRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetWorkflowRunsOutput{}
	fnCacher := func(out *glue.GetWorkflowRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetWorkflowRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetWorkflowRunsPagesWithContext(ctx context.Context, input *glue.GetWorkflowRunsInput, fn func(*glue.GetWorkflowRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.GetWorkflowRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.GetWorkflowRunsOutput{}
	fnCacher := func(out *glue.GetWorkflowRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.GetWorkflowRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) GetWorkflowRunsWithContext(ctx context.Context, input *glue.GetWorkflowRunsInput, opts ...request.Option) (*glue.GetWorkflowRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowRunsOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) GetWorkflowWithContext(ctx context.Context, input *glue.GetWorkflowInput, opts ...request.Option) (*glue.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.GetWorkflowOutput), nil
	}
	out, err := c.GlueAPI.GetWorkflowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListBlueprints(input *glue.ListBlueprintsInput) (*glue.ListBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListBlueprintsOutput), nil
	}
	out, err := c.GlueAPI.ListBlueprints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListBlueprintsPages(input *glue.ListBlueprintsInput, fn func(*glue.ListBlueprintsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListBlueprintsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListBlueprintsOutput{}
	fnCacher := func(out *glue.ListBlueprintsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListBlueprintsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListBlueprintsPagesWithContext(ctx context.Context, input *glue.ListBlueprintsInput, fn func(*glue.ListBlueprintsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListBlueprintsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListBlueprintsOutput{}
	fnCacher := func(out *glue.ListBlueprintsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListBlueprintsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListBlueprintsWithContext(ctx context.Context, input *glue.ListBlueprintsInput, opts ...request.Option) (*glue.ListBlueprintsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListBlueprintsOutput), nil
	}
	out, err := c.GlueAPI.ListBlueprintsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCrawlers(input *glue.ListCrawlersInput) (*glue.ListCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCrawlersOutput), nil
	}
	out, err := c.GlueAPI.ListCrawlers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCrawlersPages(input *glue.ListCrawlersInput, fn func(*glue.ListCrawlersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListCrawlersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListCrawlersOutput{}
	fnCacher := func(out *glue.ListCrawlersOutput, more bool) bool {
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
	if err := c.GlueAPI.ListCrawlersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListCrawlersPagesWithContext(ctx context.Context, input *glue.ListCrawlersInput, fn func(*glue.ListCrawlersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListCrawlersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListCrawlersOutput{}
	fnCacher := func(out *glue.ListCrawlersOutput, more bool) bool {
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
	if err := c.GlueAPI.ListCrawlersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListCrawlersWithContext(ctx context.Context, input *glue.ListCrawlersInput, opts ...request.Option) (*glue.ListCrawlersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCrawlersOutput), nil
	}
	out, err := c.GlueAPI.ListCrawlersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCrawls(input *glue.ListCrawlsInput) (*glue.ListCrawlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCrawlsOutput), nil
	}
	out, err := c.GlueAPI.ListCrawls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCrawlsWithContext(ctx context.Context, input *glue.ListCrawlsInput, opts ...request.Option) (*glue.ListCrawlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCrawlsOutput), nil
	}
	out, err := c.GlueAPI.ListCrawlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCustomEntityTypes(input *glue.ListCustomEntityTypesInput) (*glue.ListCustomEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCustomEntityTypesOutput), nil
	}
	out, err := c.GlueAPI.ListCustomEntityTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListCustomEntityTypesPages(input *glue.ListCustomEntityTypesInput, fn func(*glue.ListCustomEntityTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListCustomEntityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListCustomEntityTypesOutput{}
	fnCacher := func(out *glue.ListCustomEntityTypesOutput, more bool) bool {
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
	if err := c.GlueAPI.ListCustomEntityTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListCustomEntityTypesPagesWithContext(ctx context.Context, input *glue.ListCustomEntityTypesInput, fn func(*glue.ListCustomEntityTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListCustomEntityTypesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListCustomEntityTypesOutput{}
	fnCacher := func(out *glue.ListCustomEntityTypesOutput, more bool) bool {
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
	if err := c.GlueAPI.ListCustomEntityTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListCustomEntityTypesWithContext(ctx context.Context, input *glue.ListCustomEntityTypesInput, opts ...request.Option) (*glue.ListCustomEntityTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListCustomEntityTypesOutput), nil
	}
	out, err := c.GlueAPI.ListCustomEntityTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityResults(input *glue.ListDataQualityResultsInput) (*glue.ListDataQualityResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityResultsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityResultsPages(input *glue.ListDataQualityResultsInput, fn func(*glue.ListDataQualityResultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityResultsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityResultsOutput{}
	fnCacher := func(out *glue.ListDataQualityResultsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityResultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityResultsPagesWithContext(ctx context.Context, input *glue.ListDataQualityResultsInput, fn func(*glue.ListDataQualityResultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityResultsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityResultsOutput{}
	fnCacher := func(out *glue.ListDataQualityResultsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityResultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityResultsWithContext(ctx context.Context, input *glue.ListDataQualityResultsInput, opts ...request.Option) (*glue.ListDataQualityResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityResultsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRuleRecommendationRuns(input *glue.ListDataQualityRuleRecommendationRunsInput) (*glue.ListDataQualityRuleRecommendationRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRuleRecommendationRunsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRuleRecommendationRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRuleRecommendationRunsPages(input *glue.ListDataQualityRuleRecommendationRunsInput, fn func(*glue.ListDataQualityRuleRecommendationRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRuleRecommendationRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRuleRecommendationRunsOutput{}
	fnCacher := func(out *glue.ListDataQualityRuleRecommendationRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRuleRecommendationRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRuleRecommendationRunsPagesWithContext(ctx context.Context, input *glue.ListDataQualityRuleRecommendationRunsInput, fn func(*glue.ListDataQualityRuleRecommendationRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRuleRecommendationRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRuleRecommendationRunsOutput{}
	fnCacher := func(out *glue.ListDataQualityRuleRecommendationRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRuleRecommendationRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRuleRecommendationRunsWithContext(ctx context.Context, input *glue.ListDataQualityRuleRecommendationRunsInput, opts ...request.Option) (*glue.ListDataQualityRuleRecommendationRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRuleRecommendationRunsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRuleRecommendationRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRulesetEvaluationRuns(input *glue.ListDataQualityRulesetEvaluationRunsInput) (*glue.ListDataQualityRulesetEvaluationRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRulesetEvaluationRunsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRulesetEvaluationRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRulesetEvaluationRunsPages(input *glue.ListDataQualityRulesetEvaluationRunsInput, fn func(*glue.ListDataQualityRulesetEvaluationRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRulesetEvaluationRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRulesetEvaluationRunsOutput{}
	fnCacher := func(out *glue.ListDataQualityRulesetEvaluationRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRulesetEvaluationRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRulesetEvaluationRunsPagesWithContext(ctx context.Context, input *glue.ListDataQualityRulesetEvaluationRunsInput, fn func(*glue.ListDataQualityRulesetEvaluationRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRulesetEvaluationRunsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRulesetEvaluationRunsOutput{}
	fnCacher := func(out *glue.ListDataQualityRulesetEvaluationRunsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRulesetEvaluationRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRulesetEvaluationRunsWithContext(ctx context.Context, input *glue.ListDataQualityRulesetEvaluationRunsInput, opts ...request.Option) (*glue.ListDataQualityRulesetEvaluationRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRulesetEvaluationRunsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRulesetEvaluationRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRulesets(input *glue.ListDataQualityRulesetsInput) (*glue.ListDataQualityRulesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRulesetsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRulesets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDataQualityRulesetsPages(input *glue.ListDataQualityRulesetsInput, fn func(*glue.ListDataQualityRulesetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRulesetsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRulesetsOutput{}
	fnCacher := func(out *glue.ListDataQualityRulesetsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRulesetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRulesetsPagesWithContext(ctx context.Context, input *glue.ListDataQualityRulesetsInput, fn func(*glue.ListDataQualityRulesetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDataQualityRulesetsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDataQualityRulesetsOutput{}
	fnCacher := func(out *glue.ListDataQualityRulesetsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDataQualityRulesetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDataQualityRulesetsWithContext(ctx context.Context, input *glue.ListDataQualityRulesetsInput, opts ...request.Option) (*glue.ListDataQualityRulesetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDataQualityRulesetsOutput), nil
	}
	out, err := c.GlueAPI.ListDataQualityRulesetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDevEndpoints(input *glue.ListDevEndpointsInput) (*glue.ListDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.ListDevEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListDevEndpointsPages(input *glue.ListDevEndpointsInput, fn func(*glue.ListDevEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDevEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDevEndpointsOutput{}
	fnCacher := func(out *glue.ListDevEndpointsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDevEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDevEndpointsPagesWithContext(ctx context.Context, input *glue.ListDevEndpointsInput, fn func(*glue.ListDevEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListDevEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListDevEndpointsOutput{}
	fnCacher := func(out *glue.ListDevEndpointsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListDevEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListDevEndpointsWithContext(ctx context.Context, input *glue.ListDevEndpointsInput, opts ...request.Option) (*glue.ListDevEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListDevEndpointsOutput), nil
	}
	out, err := c.GlueAPI.ListDevEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListJobs(input *glue.ListJobsInput) (*glue.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListJobsOutput), nil
	}
	out, err := c.GlueAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListJobsPages(input *glue.ListJobsInput, fn func(*glue.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListJobsOutput{}
	fnCacher := func(out *glue.ListJobsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListJobsPagesWithContext(ctx context.Context, input *glue.ListJobsInput, fn func(*glue.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListJobsOutput{}
	fnCacher := func(out *glue.ListJobsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListJobsWithContext(ctx context.Context, input *glue.ListJobsInput, opts ...request.Option) (*glue.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListJobsOutput), nil
	}
	out, err := c.GlueAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListMLTransforms(input *glue.ListMLTransformsInput) (*glue.ListMLTransformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListMLTransformsOutput), nil
	}
	out, err := c.GlueAPI.ListMLTransforms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListMLTransformsPages(input *glue.ListMLTransformsInput, fn func(*glue.ListMLTransformsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListMLTransformsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListMLTransformsOutput{}
	fnCacher := func(out *glue.ListMLTransformsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListMLTransformsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListMLTransformsPagesWithContext(ctx context.Context, input *glue.ListMLTransformsInput, fn func(*glue.ListMLTransformsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListMLTransformsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListMLTransformsOutput{}
	fnCacher := func(out *glue.ListMLTransformsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListMLTransformsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListMLTransformsWithContext(ctx context.Context, input *glue.ListMLTransformsInput, opts ...request.Option) (*glue.ListMLTransformsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListMLTransformsOutput), nil
	}
	out, err := c.GlueAPI.ListMLTransformsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListRegistries(input *glue.ListRegistriesInput) (*glue.ListRegistriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListRegistriesOutput), nil
	}
	out, err := c.GlueAPI.ListRegistries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListRegistriesPages(input *glue.ListRegistriesInput, fn func(*glue.ListRegistriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListRegistriesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListRegistriesOutput{}
	fnCacher := func(out *glue.ListRegistriesOutput, more bool) bool {
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
	if err := c.GlueAPI.ListRegistriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListRegistriesPagesWithContext(ctx context.Context, input *glue.ListRegistriesInput, fn func(*glue.ListRegistriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListRegistriesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListRegistriesOutput{}
	fnCacher := func(out *glue.ListRegistriesOutput, more bool) bool {
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
	if err := c.GlueAPI.ListRegistriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListRegistriesWithContext(ctx context.Context, input *glue.ListRegistriesInput, opts ...request.Option) (*glue.ListRegistriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListRegistriesOutput), nil
	}
	out, err := c.GlueAPI.ListRegistriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSchemaVersions(input *glue.ListSchemaVersionsInput) (*glue.ListSchemaVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSchemaVersionsOutput), nil
	}
	out, err := c.GlueAPI.ListSchemaVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSchemaVersionsPages(input *glue.ListSchemaVersionsInput, fn func(*glue.ListSchemaVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSchemaVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSchemaVersionsOutput{}
	fnCacher := func(out *glue.ListSchemaVersionsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSchemaVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSchemaVersionsPagesWithContext(ctx context.Context, input *glue.ListSchemaVersionsInput, fn func(*glue.ListSchemaVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSchemaVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSchemaVersionsOutput{}
	fnCacher := func(out *glue.ListSchemaVersionsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSchemaVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSchemaVersionsWithContext(ctx context.Context, input *glue.ListSchemaVersionsInput, opts ...request.Option) (*glue.ListSchemaVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSchemaVersionsOutput), nil
	}
	out, err := c.GlueAPI.ListSchemaVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSchemas(input *glue.ListSchemasInput) (*glue.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSchemasOutput), nil
	}
	out, err := c.GlueAPI.ListSchemas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSchemasPages(input *glue.ListSchemasInput, fn func(*glue.ListSchemasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSchemasOutput{}
	fnCacher := func(out *glue.ListSchemasOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSchemasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSchemasPagesWithContext(ctx context.Context, input *glue.ListSchemasInput, fn func(*glue.ListSchemasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSchemasOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSchemasOutput{}
	fnCacher := func(out *glue.ListSchemasOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSchemasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSchemasWithContext(ctx context.Context, input *glue.ListSchemasInput, opts ...request.Option) (*glue.ListSchemasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSchemasOutput), nil
	}
	out, err := c.GlueAPI.ListSchemasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSessions(input *glue.ListSessionsInput) (*glue.ListSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSessionsOutput), nil
	}
	out, err := c.GlueAPI.ListSessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListSessionsPages(input *glue.ListSessionsInput, fn func(*glue.ListSessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSessionsOutput{}
	fnCacher := func(out *glue.ListSessionsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSessionsPagesWithContext(ctx context.Context, input *glue.ListSessionsInput, fn func(*glue.ListSessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListSessionsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListSessionsOutput{}
	fnCacher := func(out *glue.ListSessionsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListSessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListSessionsWithContext(ctx context.Context, input *glue.ListSessionsInput, opts ...request.Option) (*glue.ListSessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListSessionsOutput), nil
	}
	out, err := c.GlueAPI.ListSessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListStatements(input *glue.ListStatementsInput) (*glue.ListStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListStatementsOutput), nil
	}
	out, err := c.GlueAPI.ListStatements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListStatementsWithContext(ctx context.Context, input *glue.ListStatementsInput, opts ...request.Option) (*glue.ListStatementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListStatementsOutput), nil
	}
	out, err := c.GlueAPI.ListStatementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListTriggers(input *glue.ListTriggersInput) (*glue.ListTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListTriggersOutput), nil
	}
	out, err := c.GlueAPI.ListTriggers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListTriggersPages(input *glue.ListTriggersInput, fn func(*glue.ListTriggersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListTriggersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListTriggersOutput{}
	fnCacher := func(out *glue.ListTriggersOutput, more bool) bool {
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
	if err := c.GlueAPI.ListTriggersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListTriggersPagesWithContext(ctx context.Context, input *glue.ListTriggersInput, fn func(*glue.ListTriggersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListTriggersOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListTriggersOutput{}
	fnCacher := func(out *glue.ListTriggersOutput, more bool) bool {
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
	if err := c.GlueAPI.ListTriggersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListTriggersWithContext(ctx context.Context, input *glue.ListTriggersInput, opts ...request.Option) (*glue.ListTriggersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListTriggersOutput), nil
	}
	out, err := c.GlueAPI.ListTriggersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListWorkflows(input *glue.ListWorkflowsInput) (*glue.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListWorkflowsOutput), nil
	}
	out, err := c.GlueAPI.ListWorkflows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) ListWorkflowsPages(input *glue.ListWorkflowsInput, fn func(*glue.ListWorkflowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListWorkflowsOutput{}
	fnCacher := func(out *glue.ListWorkflowsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListWorkflowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListWorkflowsPagesWithContext(ctx context.Context, input *glue.ListWorkflowsInput, fn func(*glue.ListWorkflowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &glue.ListWorkflowsOutput{}
	fnCacher := func(out *glue.ListWorkflowsOutput, more bool) bool {
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
	if err := c.GlueAPI.ListWorkflowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) ListWorkflowsWithContext(ctx context.Context, input *glue.ListWorkflowsInput, opts ...request.Option) (*glue.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.ListWorkflowsOutput), nil
	}
	out, err := c.GlueAPI.ListWorkflowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) QuerySchemaVersionMetadata(input *glue.QuerySchemaVersionMetadataInput) (*glue.QuerySchemaVersionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.QuerySchemaVersionMetadataOutput), nil
	}
	out, err := c.GlueAPI.QuerySchemaVersionMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) QuerySchemaVersionMetadataWithContext(ctx context.Context, input *glue.QuerySchemaVersionMetadataInput, opts ...request.Option) (*glue.QuerySchemaVersionMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.QuerySchemaVersionMetadataOutput), nil
	}
	out, err := c.GlueAPI.QuerySchemaVersionMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) SearchTables(input *glue.SearchTablesInput) (*glue.SearchTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.SearchTablesOutput), nil
	}
	out, err := c.GlueAPI.SearchTables(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Glue) SearchTablesPages(input *glue.SearchTablesInput, fn func(*glue.SearchTablesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.SearchTablesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.SearchTablesOutput{}
	fnCacher := func(out *glue.SearchTablesOutput, more bool) bool {
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
	if err := c.GlueAPI.SearchTablesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) SearchTablesPagesWithContext(ctx context.Context, input *glue.SearchTablesInput, fn func(*glue.SearchTablesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*glue.SearchTablesOutput), false)
		return nil
	}
	cachable := true
	output := &glue.SearchTablesOutput{}
	fnCacher := func(out *glue.SearchTablesOutput, more bool) bool {
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
	if err := c.GlueAPI.SearchTablesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Glue) SearchTablesWithContext(ctx context.Context, input *glue.SearchTablesInput, opts ...request.Option) (*glue.SearchTablesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*glue.SearchTablesOutput), nil
	}
	out, err := c.GlueAPI.SearchTablesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
