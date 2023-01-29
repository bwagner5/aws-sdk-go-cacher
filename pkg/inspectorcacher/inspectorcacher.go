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
package inspectorcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/inspector"
	"github.com/aws/aws-sdk-go/service/inspector/inspectoriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Inspector struct {
	inspectoriface.InspectorAPI
	cache *cache.Cache
}

func New(inspectorapi inspectoriface.InspectorAPI) *Inspector {
	return &Inspector{
		InspectorAPI: inspectorapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Inspector) DescribeAssessmentRuns(input *inspector.DescribeAssessmentRunsInput) (*inspector.DescribeAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentRunsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeAssessmentRunsWithContext(ctx context.Context, input *inspector.DescribeAssessmentRunsInput, opts ...request.Option) (*inspector.DescribeAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentRunsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeAssessmentTargets(input *inspector.DescribeAssessmentTargetsInput) (*inspector.DescribeAssessmentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentTargetsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeAssessmentTargetsWithContext(ctx context.Context, input *inspector.DescribeAssessmentTargetsInput, opts ...request.Option) (*inspector.DescribeAssessmentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentTargetsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeAssessmentTemplates(input *inspector.DescribeAssessmentTemplatesInput) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentTemplatesOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeAssessmentTemplatesWithContext(ctx context.Context, input *inspector.DescribeAssessmentTemplatesInput, opts ...request.Option) (*inspector.DescribeAssessmentTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeAssessmentTemplatesOutput), nil
	}
	out, err := c.InspectorAPI.DescribeAssessmentTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeCrossAccountAccessRole(input *inspector.DescribeCrossAccountAccessRoleInput) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeCrossAccountAccessRoleOutput), nil
	}
	out, err := c.InspectorAPI.DescribeCrossAccountAccessRole(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeCrossAccountAccessRoleWithContext(ctx context.Context, input *inspector.DescribeCrossAccountAccessRoleInput, opts ...request.Option) (*inspector.DescribeCrossAccountAccessRoleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeCrossAccountAccessRoleOutput), nil
	}
	out, err := c.InspectorAPI.DescribeCrossAccountAccessRoleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeExclusions(input *inspector.DescribeExclusionsInput) (*inspector.DescribeExclusionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeExclusionsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeExclusions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeExclusionsWithContext(ctx context.Context, input *inspector.DescribeExclusionsInput, opts ...request.Option) (*inspector.DescribeExclusionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeExclusionsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeExclusionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeFindings(input *inspector.DescribeFindingsInput) (*inspector.DescribeFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeFindingsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeFindingsWithContext(ctx context.Context, input *inspector.DescribeFindingsInput, opts ...request.Option) (*inspector.DescribeFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeFindingsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeResourceGroups(input *inspector.DescribeResourceGroupsInput) (*inspector.DescribeResourceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeResourceGroupsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeResourceGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeResourceGroupsWithContext(ctx context.Context, input *inspector.DescribeResourceGroupsInput, opts ...request.Option) (*inspector.DescribeResourceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeResourceGroupsOutput), nil
	}
	out, err := c.InspectorAPI.DescribeResourceGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeRulesPackages(input *inspector.DescribeRulesPackagesInput) (*inspector.DescribeRulesPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeRulesPackagesOutput), nil
	}
	out, err := c.InspectorAPI.DescribeRulesPackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) DescribeRulesPackagesWithContext(ctx context.Context, input *inspector.DescribeRulesPackagesInput, opts ...request.Option) (*inspector.DescribeRulesPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.DescribeRulesPackagesOutput), nil
	}
	out, err := c.InspectorAPI.DescribeRulesPackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetAssessmentReport(input *inspector.GetAssessmentReportInput) (*inspector.GetAssessmentReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetAssessmentReportOutput), nil
	}
	out, err := c.InspectorAPI.GetAssessmentReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetAssessmentReportWithContext(ctx context.Context, input *inspector.GetAssessmentReportInput, opts ...request.Option) (*inspector.GetAssessmentReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetAssessmentReportOutput), nil
	}
	out, err := c.InspectorAPI.GetAssessmentReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetExclusionsPreview(input *inspector.GetExclusionsPreviewInput) (*inspector.GetExclusionsPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetExclusionsPreviewOutput), nil
	}
	out, err := c.InspectorAPI.GetExclusionsPreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetExclusionsPreviewPages(input *inspector.GetExclusionsPreviewInput, fn func(*inspector.GetExclusionsPreviewOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.GetExclusionsPreviewOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.GetExclusionsPreviewOutput{}
	fnCacher := func(out *inspector.GetExclusionsPreviewOutput, more bool) bool {
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
	if err := c.InspectorAPI.GetExclusionsPreviewPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) GetExclusionsPreviewPagesWithContext(ctx context.Context, input *inspector.GetExclusionsPreviewInput, fn func(*inspector.GetExclusionsPreviewOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.GetExclusionsPreviewOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.GetExclusionsPreviewOutput{}
	fnCacher := func(out *inspector.GetExclusionsPreviewOutput, more bool) bool {
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
	if err := c.InspectorAPI.GetExclusionsPreviewPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) GetExclusionsPreviewWithContext(ctx context.Context, input *inspector.GetExclusionsPreviewInput, opts ...request.Option) (*inspector.GetExclusionsPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetExclusionsPreviewOutput), nil
	}
	out, err := c.InspectorAPI.GetExclusionsPreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetTelemetryMetadata(input *inspector.GetTelemetryMetadataInput) (*inspector.GetTelemetryMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetTelemetryMetadataOutput), nil
	}
	out, err := c.InspectorAPI.GetTelemetryMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) GetTelemetryMetadataWithContext(ctx context.Context, input *inspector.GetTelemetryMetadataInput, opts ...request.Option) (*inspector.GetTelemetryMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.GetTelemetryMetadataOutput), nil
	}
	out, err := c.InspectorAPI.GetTelemetryMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentRunAgents(input *inspector.ListAssessmentRunAgentsInput) (*inspector.ListAssessmentRunAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentRunAgentsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentRunAgents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentRunAgentsPages(input *inspector.ListAssessmentRunAgentsInput, fn func(*inspector.ListAssessmentRunAgentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentRunAgentsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentRunAgentsOutput{}
	fnCacher := func(out *inspector.ListAssessmentRunAgentsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentRunAgentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentRunAgentsPagesWithContext(ctx context.Context, input *inspector.ListAssessmentRunAgentsInput, fn func(*inspector.ListAssessmentRunAgentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentRunAgentsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentRunAgentsOutput{}
	fnCacher := func(out *inspector.ListAssessmentRunAgentsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentRunAgentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentRunAgentsWithContext(ctx context.Context, input *inspector.ListAssessmentRunAgentsInput, opts ...request.Option) (*inspector.ListAssessmentRunAgentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentRunAgentsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentRunAgentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentRuns(input *inspector.ListAssessmentRunsInput) (*inspector.ListAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentRunsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentRuns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentRunsPages(input *inspector.ListAssessmentRunsInput, fn func(*inspector.ListAssessmentRunsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentRunsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentRunsOutput{}
	fnCacher := func(out *inspector.ListAssessmentRunsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentRunsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentRunsPagesWithContext(ctx context.Context, input *inspector.ListAssessmentRunsInput, fn func(*inspector.ListAssessmentRunsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentRunsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentRunsOutput{}
	fnCacher := func(out *inspector.ListAssessmentRunsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentRunsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentRunsWithContext(ctx context.Context, input *inspector.ListAssessmentRunsInput, opts ...request.Option) (*inspector.ListAssessmentRunsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentRunsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentRunsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentTargets(input *inspector.ListAssessmentTargetsInput) (*inspector.ListAssessmentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentTargetsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentTargetsPages(input *inspector.ListAssessmentTargetsInput, fn func(*inspector.ListAssessmentTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentTargetsOutput{}
	fnCacher := func(out *inspector.ListAssessmentTargetsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentTargetsPagesWithContext(ctx context.Context, input *inspector.ListAssessmentTargetsInput, fn func(*inspector.ListAssessmentTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentTargetsOutput{}
	fnCacher := func(out *inspector.ListAssessmentTargetsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentTargetsWithContext(ctx context.Context, input *inspector.ListAssessmentTargetsInput, opts ...request.Option) (*inspector.ListAssessmentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentTargetsOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentTemplates(input *inspector.ListAssessmentTemplatesInput) (*inspector.ListAssessmentTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentTemplatesOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListAssessmentTemplatesPages(input *inspector.ListAssessmentTemplatesInput, fn func(*inspector.ListAssessmentTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentTemplatesOutput{}
	fnCacher := func(out *inspector.ListAssessmentTemplatesOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentTemplatesPagesWithContext(ctx context.Context, input *inspector.ListAssessmentTemplatesInput, fn func(*inspector.ListAssessmentTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListAssessmentTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListAssessmentTemplatesOutput{}
	fnCacher := func(out *inspector.ListAssessmentTemplatesOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListAssessmentTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListAssessmentTemplatesWithContext(ctx context.Context, input *inspector.ListAssessmentTemplatesInput, opts ...request.Option) (*inspector.ListAssessmentTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListAssessmentTemplatesOutput), nil
	}
	out, err := c.InspectorAPI.ListAssessmentTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListEventSubscriptions(input *inspector.ListEventSubscriptionsInput) (*inspector.ListEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListEventSubscriptionsOutput), nil
	}
	out, err := c.InspectorAPI.ListEventSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListEventSubscriptionsPages(input *inspector.ListEventSubscriptionsInput, fn func(*inspector.ListEventSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListEventSubscriptionsOutput{}
	fnCacher := func(out *inspector.ListEventSubscriptionsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListEventSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListEventSubscriptionsPagesWithContext(ctx context.Context, input *inspector.ListEventSubscriptionsInput, fn func(*inspector.ListEventSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListEventSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListEventSubscriptionsOutput{}
	fnCacher := func(out *inspector.ListEventSubscriptionsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListEventSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListEventSubscriptionsWithContext(ctx context.Context, input *inspector.ListEventSubscriptionsInput, opts ...request.Option) (*inspector.ListEventSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListEventSubscriptionsOutput), nil
	}
	out, err := c.InspectorAPI.ListEventSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListExclusions(input *inspector.ListExclusionsInput) (*inspector.ListExclusionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListExclusionsOutput), nil
	}
	out, err := c.InspectorAPI.ListExclusions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListExclusionsPages(input *inspector.ListExclusionsInput, fn func(*inspector.ListExclusionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListExclusionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListExclusionsOutput{}
	fnCacher := func(out *inspector.ListExclusionsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListExclusionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListExclusionsPagesWithContext(ctx context.Context, input *inspector.ListExclusionsInput, fn func(*inspector.ListExclusionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListExclusionsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListExclusionsOutput{}
	fnCacher := func(out *inspector.ListExclusionsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListExclusionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListExclusionsWithContext(ctx context.Context, input *inspector.ListExclusionsInput, opts ...request.Option) (*inspector.ListExclusionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListExclusionsOutput), nil
	}
	out, err := c.InspectorAPI.ListExclusionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListFindings(input *inspector.ListFindingsInput) (*inspector.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListFindingsOutput), nil
	}
	out, err := c.InspectorAPI.ListFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListFindingsPages(input *inspector.ListFindingsInput, fn func(*inspector.ListFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListFindingsOutput{}
	fnCacher := func(out *inspector.ListFindingsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListFindingsPagesWithContext(ctx context.Context, input *inspector.ListFindingsInput, fn func(*inspector.ListFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListFindingsOutput{}
	fnCacher := func(out *inspector.ListFindingsOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListFindingsWithContext(ctx context.Context, input *inspector.ListFindingsInput, opts ...request.Option) (*inspector.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListFindingsOutput), nil
	}
	out, err := c.InspectorAPI.ListFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListRulesPackages(input *inspector.ListRulesPackagesInput) (*inspector.ListRulesPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListRulesPackagesOutput), nil
	}
	out, err := c.InspectorAPI.ListRulesPackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListRulesPackagesPages(input *inspector.ListRulesPackagesInput, fn func(*inspector.ListRulesPackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListRulesPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListRulesPackagesOutput{}
	fnCacher := func(out *inspector.ListRulesPackagesOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListRulesPackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListRulesPackagesPagesWithContext(ctx context.Context, input *inspector.ListRulesPackagesInput, fn func(*inspector.ListRulesPackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*inspector.ListRulesPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &inspector.ListRulesPackagesOutput{}
	fnCacher := func(out *inspector.ListRulesPackagesOutput, more bool) bool {
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
	if err := c.InspectorAPI.ListRulesPackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Inspector) ListRulesPackagesWithContext(ctx context.Context, input *inspector.ListRulesPackagesInput, opts ...request.Option) (*inspector.ListRulesPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListRulesPackagesOutput), nil
	}
	out, err := c.InspectorAPI.ListRulesPackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListTagsForResource(input *inspector.ListTagsForResourceInput) (*inspector.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListTagsForResourceOutput), nil
	}
	out, err := c.InspectorAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Inspector) ListTagsForResourceWithContext(ctx context.Context, input *inspector.ListTagsForResourceInput, opts ...request.Option) (*inspector.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*inspector.ListTagsForResourceOutput), nil
	}
	out, err := c.InspectorAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
