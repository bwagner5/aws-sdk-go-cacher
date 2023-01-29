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
package migrationhuborchestratorcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhuborchestrator"
	"github.com/aws/aws-sdk-go/service/migrationhuborchestrator/migrationhuborchestratoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MigrationHubOrchestrator struct {
	migrationhuborchestratoriface.MigrationHubOrchestratorAPI
	cache *cache.Cache
}

func New(migrationhuborchestratorapi migrationhuborchestratoriface.MigrationHubOrchestratorAPI) *MigrationHubOrchestrator {
	return &MigrationHubOrchestrator{
		MigrationHubOrchestratorAPI: migrationhuborchestratorapi,
		cache:                       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MigrationHubOrchestrator) GetTemplate(input *migrationhuborchestrator.GetTemplateInput) (*migrationhuborchestrator.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetTemplateStep(input *migrationhuborchestrator.GetTemplateStepInput) (*migrationhuborchestrator.GetTemplateStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateStepOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplateStep(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetTemplateStepGroup(input *migrationhuborchestrator.GetTemplateStepGroupInput) (*migrationhuborchestrator.GetTemplateStepGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateStepGroupOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplateStepGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetTemplateStepGroupWithContext(ctx context.Context, input *migrationhuborchestrator.GetTemplateStepGroupInput, opts ...request.Option) (*migrationhuborchestrator.GetTemplateStepGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateStepGroupOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplateStepGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetTemplateStepWithContext(ctx context.Context, input *migrationhuborchestrator.GetTemplateStepInput, opts ...request.Option) (*migrationhuborchestrator.GetTemplateStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateStepOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplateStepWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetTemplateWithContext(ctx context.Context, input *migrationhuborchestrator.GetTemplateInput, opts ...request.Option) (*migrationhuborchestrator.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetTemplateOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflow(input *migrationhuborchestrator.GetWorkflowInput) (*migrationhuborchestrator.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflow(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflowStep(input *migrationhuborchestrator.GetWorkflowStepInput) (*migrationhuborchestrator.GetWorkflowStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowStepOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflowStep(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflowStepGroup(input *migrationhuborchestrator.GetWorkflowStepGroupInput) (*migrationhuborchestrator.GetWorkflowStepGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowStepGroupOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflowStepGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflowStepGroupWithContext(ctx context.Context, input *migrationhuborchestrator.GetWorkflowStepGroupInput, opts ...request.Option) (*migrationhuborchestrator.GetWorkflowStepGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowStepGroupOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflowStepGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflowStepWithContext(ctx context.Context, input *migrationhuborchestrator.GetWorkflowStepInput, opts ...request.Option) (*migrationhuborchestrator.GetWorkflowStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowStepOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflowStepWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) GetWorkflowWithContext(ctx context.Context, input *migrationhuborchestrator.GetWorkflowInput, opts ...request.Option) (*migrationhuborchestrator.GetWorkflowOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.GetWorkflowOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.GetWorkflowWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListPlugins(input *migrationhuborchestrator.ListPluginsInput) (*migrationhuborchestrator.ListPluginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListPluginsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListPlugins(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListPluginsPages(input *migrationhuborchestrator.ListPluginsInput, fn func(*migrationhuborchestrator.ListPluginsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListPluginsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListPluginsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListPluginsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListPluginsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListPluginsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListPluginsInput, fn func(*migrationhuborchestrator.ListPluginsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListPluginsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListPluginsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListPluginsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListPluginsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListPluginsWithContext(ctx context.Context, input *migrationhuborchestrator.ListPluginsInput, opts ...request.Option) (*migrationhuborchestrator.ListPluginsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListPluginsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListPluginsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTagsForResource(input *migrationhuborchestrator.ListTagsForResourceInput) (*migrationhuborchestrator.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTagsForResourceOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTagsForResourceWithContext(ctx context.Context, input *migrationhuborchestrator.ListTagsForResourceInput, opts ...request.Option) (*migrationhuborchestrator.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTagsForResourceOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplateStepGroups(input *migrationhuborchestrator.ListTemplateStepGroupsInput) (*migrationhuborchestrator.ListTemplateStepGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplateStepGroupsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplateStepGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplateStepGroupsPages(input *migrationhuborchestrator.ListTemplateStepGroupsInput, fn func(*migrationhuborchestrator.ListTemplateStepGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplateStepGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplateStepGroupsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplateStepGroupsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplateStepGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplateStepGroupsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplateStepGroupsInput, fn func(*migrationhuborchestrator.ListTemplateStepGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplateStepGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplateStepGroupsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplateStepGroupsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplateStepGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplateStepGroupsWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplateStepGroupsInput, opts ...request.Option) (*migrationhuborchestrator.ListTemplateStepGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplateStepGroupsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplateStepGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplateSteps(input *migrationhuborchestrator.ListTemplateStepsInput) (*migrationhuborchestrator.ListTemplateStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplateStepsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplateSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplateStepsPages(input *migrationhuborchestrator.ListTemplateStepsInput, fn func(*migrationhuborchestrator.ListTemplateStepsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplateStepsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplateStepsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplateStepsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplateStepsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplateStepsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplateStepsInput, fn func(*migrationhuborchestrator.ListTemplateStepsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplateStepsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplateStepsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplateStepsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplateStepsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplateStepsWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplateStepsInput, opts ...request.Option) (*migrationhuborchestrator.ListTemplateStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplateStepsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplateStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplates(input *migrationhuborchestrator.ListTemplatesInput) (*migrationhuborchestrator.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplatesOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListTemplatesPages(input *migrationhuborchestrator.ListTemplatesInput, fn func(*migrationhuborchestrator.ListTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplatesOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplatesOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplatesPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplatesInput, fn func(*migrationhuborchestrator.ListTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListTemplatesOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListTemplatesOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListTemplatesWithContext(ctx context.Context, input *migrationhuborchestrator.ListTemplatesInput, opts ...request.Option) (*migrationhuborchestrator.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListTemplatesOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflowStepGroups(input *migrationhuborchestrator.ListWorkflowStepGroupsInput) (*migrationhuborchestrator.ListWorkflowStepGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowStepGroupsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflowStepGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflowStepGroupsPages(input *migrationhuborchestrator.ListWorkflowStepGroupsInput, fn func(*migrationhuborchestrator.ListWorkflowStepGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowStepGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowStepGroupsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowStepGroupsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowStepGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowStepGroupsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowStepGroupsInput, fn func(*migrationhuborchestrator.ListWorkflowStepGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowStepGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowStepGroupsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowStepGroupsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowStepGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowStepGroupsWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowStepGroupsInput, opts ...request.Option) (*migrationhuborchestrator.ListWorkflowStepGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowStepGroupsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflowStepGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflowSteps(input *migrationhuborchestrator.ListWorkflowStepsInput) (*migrationhuborchestrator.ListWorkflowStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowStepsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflowSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflowStepsPages(input *migrationhuborchestrator.ListWorkflowStepsInput, fn func(*migrationhuborchestrator.ListWorkflowStepsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowStepsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowStepsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowStepsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowStepsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowStepsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowStepsInput, fn func(*migrationhuborchestrator.ListWorkflowStepsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowStepsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowStepsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowStepsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowStepsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowStepsWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowStepsInput, opts ...request.Option) (*migrationhuborchestrator.ListWorkflowStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowStepsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflowStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflows(input *migrationhuborchestrator.ListWorkflowsInput) (*migrationhuborchestrator.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubOrchestrator) ListWorkflowsPages(input *migrationhuborchestrator.ListWorkflowsInput, fn func(*migrationhuborchestrator.ListWorkflowsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowsPagesWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowsInput, fn func(*migrationhuborchestrator.ListWorkflowsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhuborchestrator.ListWorkflowsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhuborchestrator.ListWorkflowsOutput{}
	fnCacher := func(out *migrationhuborchestrator.ListWorkflowsOutput, more bool) bool {
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
	if err := c.MigrationHubOrchestratorAPI.ListWorkflowsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubOrchestrator) ListWorkflowsWithContext(ctx context.Context, input *migrationhuborchestrator.ListWorkflowsInput, opts ...request.Option) (*migrationhuborchestrator.ListWorkflowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhuborchestrator.ListWorkflowsOutput), nil
	}
	out, err := c.MigrationHubOrchestratorAPI.ListWorkflowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
