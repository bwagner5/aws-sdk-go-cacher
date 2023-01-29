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
package codepipelinecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/aws/aws-sdk-go/service/codepipeline/codepipelineiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodePipeline struct {
	codepipelineiface.CodePipelineAPI
	cache *cache.Cache
}

func New(codepipelineapi codepipelineiface.CodePipelineAPI) *CodePipeline {
	return &CodePipeline{
		CodePipelineAPI: codepipelineapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodePipeline) GetActionType(input *codepipeline.GetActionTypeInput) (*codepipeline.GetActionTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetActionTypeOutput), nil
	}
	out, err := c.CodePipelineAPI.GetActionType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetActionTypeWithContext(ctx context.Context, input *codepipeline.GetActionTypeInput, opts ...request.Option) (*codepipeline.GetActionTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetActionTypeOutput), nil
	}
	out, err := c.CodePipelineAPI.GetActionTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetJobDetails(input *codepipeline.GetJobDetailsInput) (*codepipeline.GetJobDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetJobDetailsOutput), nil
	}
	out, err := c.CodePipelineAPI.GetJobDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetJobDetailsWithContext(ctx context.Context, input *codepipeline.GetJobDetailsInput, opts ...request.Option) (*codepipeline.GetJobDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetJobDetailsOutput), nil
	}
	out, err := c.CodePipelineAPI.GetJobDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipeline(input *codepipeline.GetPipelineInput) (*codepipeline.GetPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipelineExecution(input *codepipeline.GetPipelineExecutionInput) (*codepipeline.GetPipelineExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineExecutionOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipelineExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipelineExecutionWithContext(ctx context.Context, input *codepipeline.GetPipelineExecutionInput, opts ...request.Option) (*codepipeline.GetPipelineExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineExecutionOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipelineExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipelineState(input *codepipeline.GetPipelineStateInput) (*codepipeline.GetPipelineStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineStateOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipelineState(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipelineStateWithContext(ctx context.Context, input *codepipeline.GetPipelineStateInput, opts ...request.Option) (*codepipeline.GetPipelineStateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineStateOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipelineStateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetPipelineWithContext(ctx context.Context, input *codepipeline.GetPipelineInput, opts ...request.Option) (*codepipeline.GetPipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetPipelineOutput), nil
	}
	out, err := c.CodePipelineAPI.GetPipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetThirdPartyJobDetails(input *codepipeline.GetThirdPartyJobDetailsInput) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetThirdPartyJobDetailsOutput), nil
	}
	out, err := c.CodePipelineAPI.GetThirdPartyJobDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) GetThirdPartyJobDetailsWithContext(ctx context.Context, input *codepipeline.GetThirdPartyJobDetailsInput, opts ...request.Option) (*codepipeline.GetThirdPartyJobDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.GetThirdPartyJobDetailsOutput), nil
	}
	out, err := c.CodePipelineAPI.GetThirdPartyJobDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListActionExecutions(input *codepipeline.ListActionExecutionsInput) (*codepipeline.ListActionExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListActionExecutionsOutput), nil
	}
	out, err := c.CodePipelineAPI.ListActionExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListActionExecutionsPages(input *codepipeline.ListActionExecutionsInput, fn func(*codepipeline.ListActionExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListActionExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListActionExecutionsOutput{}
	fnCacher := func(out *codepipeline.ListActionExecutionsOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListActionExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListActionExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, fn func(*codepipeline.ListActionExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListActionExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListActionExecutionsOutput{}
	fnCacher := func(out *codepipeline.ListActionExecutionsOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListActionExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListActionExecutionsWithContext(ctx context.Context, input *codepipeline.ListActionExecutionsInput, opts ...request.Option) (*codepipeline.ListActionExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListActionExecutionsOutput), nil
	}
	out, err := c.CodePipelineAPI.ListActionExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListActionTypes(input *codepipeline.ListActionTypesInput) (*codepipeline.ListActionTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListActionTypesOutput), nil
	}
	out, err := c.CodePipelineAPI.ListActionTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListActionTypesPages(input *codepipeline.ListActionTypesInput, fn func(*codepipeline.ListActionTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListActionTypesOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListActionTypesOutput{}
	fnCacher := func(out *codepipeline.ListActionTypesOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListActionTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListActionTypesPagesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, fn func(*codepipeline.ListActionTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListActionTypesOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListActionTypesOutput{}
	fnCacher := func(out *codepipeline.ListActionTypesOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListActionTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListActionTypesWithContext(ctx context.Context, input *codepipeline.ListActionTypesInput, opts ...request.Option) (*codepipeline.ListActionTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListActionTypesOutput), nil
	}
	out, err := c.CodePipelineAPI.ListActionTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListPipelineExecutions(input *codepipeline.ListPipelineExecutionsInput) (*codepipeline.ListPipelineExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListPipelineExecutionsOutput), nil
	}
	out, err := c.CodePipelineAPI.ListPipelineExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListPipelineExecutionsPages(input *codepipeline.ListPipelineExecutionsInput, fn func(*codepipeline.ListPipelineExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListPipelineExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListPipelineExecutionsOutput{}
	fnCacher := func(out *codepipeline.ListPipelineExecutionsOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListPipelineExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListPipelineExecutionsPagesWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, fn func(*codepipeline.ListPipelineExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListPipelineExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListPipelineExecutionsOutput{}
	fnCacher := func(out *codepipeline.ListPipelineExecutionsOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListPipelineExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListPipelineExecutionsWithContext(ctx context.Context, input *codepipeline.ListPipelineExecutionsInput, opts ...request.Option) (*codepipeline.ListPipelineExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListPipelineExecutionsOutput), nil
	}
	out, err := c.CodePipelineAPI.ListPipelineExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListPipelines(input *codepipeline.ListPipelinesInput) (*codepipeline.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListPipelinesOutput), nil
	}
	out, err := c.CodePipelineAPI.ListPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListPipelinesPages(input *codepipeline.ListPipelinesInput, fn func(*codepipeline.ListPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListPipelinesOutput{}
	fnCacher := func(out *codepipeline.ListPipelinesOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListPipelinesPagesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, fn func(*codepipeline.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListPipelinesOutput{}
	fnCacher := func(out *codepipeline.ListPipelinesOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListPipelinesWithContext(ctx context.Context, input *codepipeline.ListPipelinesInput, opts ...request.Option) (*codepipeline.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListPipelinesOutput), nil
	}
	out, err := c.CodePipelineAPI.ListPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListTagsForResource(input *codepipeline.ListTagsForResourceInput) (*codepipeline.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListTagsForResourceOutput), nil
	}
	out, err := c.CodePipelineAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListTagsForResourcePages(input *codepipeline.ListTagsForResourceInput, fn func(*codepipeline.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListTagsForResourceOutput{}
	fnCacher := func(out *codepipeline.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListTagsForResourcePagesWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, fn func(*codepipeline.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListTagsForResourceOutput{}
	fnCacher := func(out *codepipeline.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListTagsForResourceWithContext(ctx context.Context, input *codepipeline.ListTagsForResourceInput, opts ...request.Option) (*codepipeline.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListTagsForResourceOutput), nil
	}
	out, err := c.CodePipelineAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListWebhooks(input *codepipeline.ListWebhooksInput) (*codepipeline.ListWebhooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListWebhooksOutput), nil
	}
	out, err := c.CodePipelineAPI.ListWebhooks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodePipeline) ListWebhooksPages(input *codepipeline.ListWebhooksInput, fn func(*codepipeline.ListWebhooksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListWebhooksOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListWebhooksOutput{}
	fnCacher := func(out *codepipeline.ListWebhooksOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListWebhooksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListWebhooksPagesWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, fn func(*codepipeline.ListWebhooksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codepipeline.ListWebhooksOutput), false)
		return nil
	}
	cachable := true
	output := &codepipeline.ListWebhooksOutput{}
	fnCacher := func(out *codepipeline.ListWebhooksOutput, more bool) bool {
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
	if err := c.CodePipelineAPI.ListWebhooksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodePipeline) ListWebhooksWithContext(ctx context.Context, input *codepipeline.ListWebhooksInput, opts ...request.Option) (*codepipeline.ListWebhooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codepipeline.ListWebhooksOutput), nil
	}
	out, err := c.CodePipelineAPI.ListWebhooksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
