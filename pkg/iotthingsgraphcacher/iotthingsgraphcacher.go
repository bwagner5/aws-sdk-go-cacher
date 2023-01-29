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
package iotthingsgraphcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTThingsGraph struct {
	iotthingsgraphiface.IoTThingsGraphAPI
	cache *cache.Cache
}

func New(iotthingsgraphapi iotthingsgraphiface.IoTThingsGraphAPI) *IoTThingsGraph {
	return &IoTThingsGraph{
		IoTThingsGraphAPI: iotthingsgraphapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTThingsGraph) DescribeNamespace(input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.DescribeNamespaceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.DescribeNamespace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) DescribeNamespaceWithContext(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput, opts ...request.Option) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.DescribeNamespaceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.DescribeNamespaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetEntities(input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetEntitiesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetEntitiesWithContext(ctx context.Context, input *iotthingsgraph.GetEntitiesInput, opts ...request.Option) (*iotthingsgraph.GetEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetEntitiesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetFlowTemplate(input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetFlowTemplateOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetFlowTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetFlowTemplateRevisions(input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetFlowTemplateRevisionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetFlowTemplateRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetFlowTemplateRevisionsPages(input *iotthingsgraph.GetFlowTemplateRevisionsInput, fn func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.GetFlowTemplateRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.GetFlowTemplateRevisionsOutput{}
	fnCacher := func(out *iotthingsgraph.GetFlowTemplateRevisionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.GetFlowTemplateRevisionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) GetFlowTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, fn func(*iotthingsgraph.GetFlowTemplateRevisionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.GetFlowTemplateRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.GetFlowTemplateRevisionsOutput{}
	fnCacher := func(out *iotthingsgraph.GetFlowTemplateRevisionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.GetFlowTemplateRevisionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) GetFlowTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetFlowTemplateRevisionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetFlowTemplateRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetFlowTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput, opts ...request.Option) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetFlowTemplateOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetFlowTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetNamespaceDeletionStatus(input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetNamespaceDeletionStatusOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetNamespaceDeletionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetNamespaceDeletionStatusWithContext(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput, opts ...request.Option) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetNamespaceDeletionStatusOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetNamespaceDeletionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemInstance(input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemInstanceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemInstanceWithContext(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput, opts ...request.Option) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemInstanceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemTemplate(input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemTemplateOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemTemplateRevisions(input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemTemplateRevisionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemTemplateRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemTemplateRevisionsPages(input *iotthingsgraph.GetSystemTemplateRevisionsInput, fn func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.GetSystemTemplateRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.GetSystemTemplateRevisionsOutput{}
	fnCacher := func(out *iotthingsgraph.GetSystemTemplateRevisionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.GetSystemTemplateRevisionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) GetSystemTemplateRevisionsPagesWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, fn func(*iotthingsgraph.GetSystemTemplateRevisionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.GetSystemTemplateRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.GetSystemTemplateRevisionsOutput{}
	fnCacher := func(out *iotthingsgraph.GetSystemTemplateRevisionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.GetSystemTemplateRevisionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) GetSystemTemplateRevisionsWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemTemplateRevisionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemTemplateRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetSystemTemplateWithContext(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput, opts ...request.Option) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetSystemTemplateOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetSystemTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetUploadStatus(input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetUploadStatusOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetUploadStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) GetUploadStatusWithContext(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput, opts ...request.Option) (*iotthingsgraph.GetUploadStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.GetUploadStatusOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.GetUploadStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) ListFlowExecutionMessages(input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.ListFlowExecutionMessagesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.ListFlowExecutionMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) ListFlowExecutionMessagesPages(input *iotthingsgraph.ListFlowExecutionMessagesInput, fn func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.ListFlowExecutionMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.ListFlowExecutionMessagesOutput{}
	fnCacher := func(out *iotthingsgraph.ListFlowExecutionMessagesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.ListFlowExecutionMessagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) ListFlowExecutionMessagesPagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, fn func(*iotthingsgraph.ListFlowExecutionMessagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.ListFlowExecutionMessagesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.ListFlowExecutionMessagesOutput{}
	fnCacher := func(out *iotthingsgraph.ListFlowExecutionMessagesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.ListFlowExecutionMessagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) ListFlowExecutionMessagesWithContext(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput, opts ...request.Option) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.ListFlowExecutionMessagesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.ListFlowExecutionMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) ListTagsForResource(input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) ListTagsForResourcePages(input *iotthingsgraph.ListTagsForResourceInput, fn func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.ListTagsForResourceOutput{}
	fnCacher := func(out *iotthingsgraph.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) ListTagsForResourcePagesWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, fn func(*iotthingsgraph.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.ListTagsForResourceOutput{}
	fnCacher := func(out *iotthingsgraph.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) ListTagsForResourceWithContext(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput, opts ...request.Option) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.ListTagsForResourceOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchEntities(input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchEntitiesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchEntitiesPages(input *iotthingsgraph.SearchEntitiesInput, fn func(*iotthingsgraph.SearchEntitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchEntitiesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchEntitiesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchEntitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchEntitiesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, fn func(*iotthingsgraph.SearchEntitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchEntitiesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchEntitiesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchEntitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchEntitiesWithContext(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput, opts ...request.Option) (*iotthingsgraph.SearchEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchEntitiesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchFlowExecutions(input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchFlowExecutionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchFlowExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchFlowExecutionsPages(input *iotthingsgraph.SearchFlowExecutionsInput, fn func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchFlowExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchFlowExecutionsOutput{}
	fnCacher := func(out *iotthingsgraph.SearchFlowExecutionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchFlowExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchFlowExecutionsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, fn func(*iotthingsgraph.SearchFlowExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchFlowExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchFlowExecutionsOutput{}
	fnCacher := func(out *iotthingsgraph.SearchFlowExecutionsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchFlowExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchFlowExecutionsWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput, opts ...request.Option) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchFlowExecutionsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchFlowExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchFlowTemplates(input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchFlowTemplatesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchFlowTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchFlowTemplatesPages(input *iotthingsgraph.SearchFlowTemplatesInput, fn func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchFlowTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchFlowTemplatesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchFlowTemplatesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchFlowTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchFlowTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, fn func(*iotthingsgraph.SearchFlowTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchFlowTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchFlowTemplatesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchFlowTemplatesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchFlowTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchFlowTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchFlowTemplatesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchFlowTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchSystemInstances(input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchSystemInstancesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchSystemInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchSystemInstancesPages(input *iotthingsgraph.SearchSystemInstancesInput, fn func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchSystemInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchSystemInstancesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchSystemInstancesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchSystemInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchSystemInstancesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, fn func(*iotthingsgraph.SearchSystemInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchSystemInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchSystemInstancesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchSystemInstancesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchSystemInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchSystemInstancesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchSystemInstancesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchSystemInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchSystemTemplates(input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchSystemTemplatesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchSystemTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchSystemTemplatesPages(input *iotthingsgraph.SearchSystemTemplatesInput, fn func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchSystemTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchSystemTemplatesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchSystemTemplatesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchSystemTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchSystemTemplatesPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, fn func(*iotthingsgraph.SearchSystemTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchSystemTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchSystemTemplatesOutput{}
	fnCacher := func(out *iotthingsgraph.SearchSystemTemplatesOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchSystemTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchSystemTemplatesWithContext(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput, opts ...request.Option) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchSystemTemplatesOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchSystemTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchThings(input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchThingsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchThings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTThingsGraph) SearchThingsPages(input *iotthingsgraph.SearchThingsInput, fn func(*iotthingsgraph.SearchThingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchThingsOutput{}
	fnCacher := func(out *iotthingsgraph.SearchThingsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchThingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchThingsPagesWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, fn func(*iotthingsgraph.SearchThingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*iotthingsgraph.SearchThingsOutput), false)
		return nil
	}
	cachable := true
	output := &iotthingsgraph.SearchThingsOutput{}
	fnCacher := func(out *iotthingsgraph.SearchThingsOutput, more bool) bool {
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
	if err := c.IoTThingsGraphAPI.SearchThingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *IoTThingsGraph) SearchThingsWithContext(ctx context.Context, input *iotthingsgraph.SearchThingsInput, opts ...request.Option) (*iotthingsgraph.SearchThingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotthingsgraph.SearchThingsOutput), nil
	}
	out, err := c.IoTThingsGraphAPI.SearchThingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
