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
package datapipelinecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DataPipeline struct {
	datapipelineiface.DataPipelineAPI
	cache *cache.Cache
}

func New(datapipelineapi datapipelineiface.DataPipelineAPI) *DataPipeline {
	return &DataPipeline{
		DataPipelineAPI: datapipelineapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DataPipeline) DescribeObjects(input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.DescribeObjectsOutput), nil
	}
	out, err := c.DataPipelineAPI.DescribeObjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) DescribeObjectsPages(input *datapipeline.DescribeObjectsInput, fn func(*datapipeline.DescribeObjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.DescribeObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.DescribeObjectsOutput{}
	fnCacher := func(out *datapipeline.DescribeObjectsOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.DescribeObjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) DescribeObjectsPagesWithContext(ctx context.Context, input *datapipeline.DescribeObjectsInput, fn func(*datapipeline.DescribeObjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.DescribeObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.DescribeObjectsOutput{}
	fnCacher := func(out *datapipeline.DescribeObjectsOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.DescribeObjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) DescribeObjectsWithContext(ctx context.Context, input *datapipeline.DescribeObjectsInput, opts ...request.Option) (*datapipeline.DescribeObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.DescribeObjectsOutput), nil
	}
	out, err := c.DataPipelineAPI.DescribeObjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) DescribePipelines(input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.DescribePipelinesOutput), nil
	}
	out, err := c.DataPipelineAPI.DescribePipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) DescribePipelinesWithContext(ctx context.Context, input *datapipeline.DescribePipelinesInput, opts ...request.Option) (*datapipeline.DescribePipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.DescribePipelinesOutput), nil
	}
	out, err := c.DataPipelineAPI.DescribePipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) GetPipelineDefinition(input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.GetPipelineDefinitionOutput), nil
	}
	out, err := c.DataPipelineAPI.GetPipelineDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) GetPipelineDefinitionWithContext(ctx context.Context, input *datapipeline.GetPipelineDefinitionInput, opts ...request.Option) (*datapipeline.GetPipelineDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.GetPipelineDefinitionOutput), nil
	}
	out, err := c.DataPipelineAPI.GetPipelineDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) ListPipelines(input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.ListPipelinesOutput), nil
	}
	out, err := c.DataPipelineAPI.ListPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) ListPipelinesPages(input *datapipeline.ListPipelinesInput, fn func(*datapipeline.ListPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.ListPipelinesOutput{}
	fnCacher := func(out *datapipeline.ListPipelinesOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.ListPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) ListPipelinesPagesWithContext(ctx context.Context, input *datapipeline.ListPipelinesInput, fn func(*datapipeline.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.ListPipelinesOutput{}
	fnCacher := func(out *datapipeline.ListPipelinesOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.ListPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) ListPipelinesWithContext(ctx context.Context, input *datapipeline.ListPipelinesInput, opts ...request.Option) (*datapipeline.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.ListPipelinesOutput), nil
	}
	out, err := c.DataPipelineAPI.ListPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) QueryObjects(input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.QueryObjectsOutput), nil
	}
	out, err := c.DataPipelineAPI.QueryObjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DataPipeline) QueryObjectsPages(input *datapipeline.QueryObjectsInput, fn func(*datapipeline.QueryObjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.QueryObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.QueryObjectsOutput{}
	fnCacher := func(out *datapipeline.QueryObjectsOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.QueryObjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) QueryObjectsPagesWithContext(ctx context.Context, input *datapipeline.QueryObjectsInput, fn func(*datapipeline.QueryObjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*datapipeline.QueryObjectsOutput), false)
		return nil
	}
	cachable := true
	output := &datapipeline.QueryObjectsOutput{}
	fnCacher := func(out *datapipeline.QueryObjectsOutput, more bool) bool {
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
	if err := c.DataPipelineAPI.QueryObjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DataPipeline) QueryObjectsWithContext(ctx context.Context, input *datapipeline.QueryObjectsInput, opts ...request.Option) (*datapipeline.QueryObjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*datapipeline.QueryObjectsOutput), nil
	}
	out, err := c.DataPipelineAPI.QueryObjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
