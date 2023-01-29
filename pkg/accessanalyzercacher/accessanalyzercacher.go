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
package accessanalyzercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/accessanalyzer"
	"github.com/aws/aws-sdk-go/service/accessanalyzer/accessanalyzeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AccessAnalyzer struct {
	accessanalyzeriface.AccessAnalyzerAPI
	cache *cache.Cache
}

func New(accessanalyzerapi accessanalyzeriface.AccessAnalyzerAPI) *AccessAnalyzer {
	return &AccessAnalyzer{
		AccessAnalyzerAPI: accessanalyzerapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AccessAnalyzer) GetAccessPreview(input *accessanalyzer.GetAccessPreviewInput) (*accessanalyzer.GetAccessPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAccessPreviewOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAccessPreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetAccessPreviewWithContext(ctx context.Context, input *accessanalyzer.GetAccessPreviewInput, opts ...request.Option) (*accessanalyzer.GetAccessPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAccessPreviewOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAccessPreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetAnalyzedResource(input *accessanalyzer.GetAnalyzedResourceInput) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAnalyzedResourceOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAnalyzedResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetAnalyzedResourceWithContext(ctx context.Context, input *accessanalyzer.GetAnalyzedResourceInput, opts ...request.Option) (*accessanalyzer.GetAnalyzedResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAnalyzedResourceOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAnalyzedResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetAnalyzer(input *accessanalyzer.GetAnalyzerInput) (*accessanalyzer.GetAnalyzerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAnalyzerOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAnalyzer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetAnalyzerWithContext(ctx context.Context, input *accessanalyzer.GetAnalyzerInput, opts ...request.Option) (*accessanalyzer.GetAnalyzerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetAnalyzerOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetAnalyzerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetArchiveRule(input *accessanalyzer.GetArchiveRuleInput) (*accessanalyzer.GetArchiveRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetArchiveRuleOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetArchiveRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetArchiveRuleWithContext(ctx context.Context, input *accessanalyzer.GetArchiveRuleInput, opts ...request.Option) (*accessanalyzer.GetArchiveRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetArchiveRuleOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetArchiveRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetFinding(input *accessanalyzer.GetFindingInput) (*accessanalyzer.GetFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetFindingOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetFinding(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetFindingWithContext(ctx context.Context, input *accessanalyzer.GetFindingInput, opts ...request.Option) (*accessanalyzer.GetFindingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetFindingOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetFindingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetGeneratedPolicy(input *accessanalyzer.GetGeneratedPolicyInput) (*accessanalyzer.GetGeneratedPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetGeneratedPolicyOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetGeneratedPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) GetGeneratedPolicyWithContext(ctx context.Context, input *accessanalyzer.GetGeneratedPolicyInput, opts ...request.Option) (*accessanalyzer.GetGeneratedPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.GetGeneratedPolicyOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.GetGeneratedPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAccessPreviewFindings(input *accessanalyzer.ListAccessPreviewFindingsInput) (*accessanalyzer.ListAccessPreviewFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAccessPreviewFindingsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAccessPreviewFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAccessPreviewFindingsPages(input *accessanalyzer.ListAccessPreviewFindingsInput, fn func(*accessanalyzer.ListAccessPreviewFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAccessPreviewFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAccessPreviewFindingsOutput{}
	fnCacher := func(out *accessanalyzer.ListAccessPreviewFindingsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAccessPreviewFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAccessPreviewFindingsPagesWithContext(ctx context.Context, input *accessanalyzer.ListAccessPreviewFindingsInput, fn func(*accessanalyzer.ListAccessPreviewFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAccessPreviewFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAccessPreviewFindingsOutput{}
	fnCacher := func(out *accessanalyzer.ListAccessPreviewFindingsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAccessPreviewFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAccessPreviewFindingsWithContext(ctx context.Context, input *accessanalyzer.ListAccessPreviewFindingsInput, opts ...request.Option) (*accessanalyzer.ListAccessPreviewFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAccessPreviewFindingsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAccessPreviewFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAccessPreviews(input *accessanalyzer.ListAccessPreviewsInput) (*accessanalyzer.ListAccessPreviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAccessPreviewsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAccessPreviews(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAccessPreviewsPages(input *accessanalyzer.ListAccessPreviewsInput, fn func(*accessanalyzer.ListAccessPreviewsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAccessPreviewsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAccessPreviewsOutput{}
	fnCacher := func(out *accessanalyzer.ListAccessPreviewsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAccessPreviewsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAccessPreviewsPagesWithContext(ctx context.Context, input *accessanalyzer.ListAccessPreviewsInput, fn func(*accessanalyzer.ListAccessPreviewsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAccessPreviewsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAccessPreviewsOutput{}
	fnCacher := func(out *accessanalyzer.ListAccessPreviewsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAccessPreviewsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAccessPreviewsWithContext(ctx context.Context, input *accessanalyzer.ListAccessPreviewsInput, opts ...request.Option) (*accessanalyzer.ListAccessPreviewsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAccessPreviewsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAccessPreviewsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAnalyzedResources(input *accessanalyzer.ListAnalyzedResourcesInput) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAnalyzedResourcesOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAnalyzedResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAnalyzedResourcesPages(input *accessanalyzer.ListAnalyzedResourcesInput, fn func(*accessanalyzer.ListAnalyzedResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAnalyzedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAnalyzedResourcesOutput{}
	fnCacher := func(out *accessanalyzer.ListAnalyzedResourcesOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAnalyzedResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAnalyzedResourcesPagesWithContext(ctx context.Context, input *accessanalyzer.ListAnalyzedResourcesInput, fn func(*accessanalyzer.ListAnalyzedResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAnalyzedResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAnalyzedResourcesOutput{}
	fnCacher := func(out *accessanalyzer.ListAnalyzedResourcesOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAnalyzedResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAnalyzedResourcesWithContext(ctx context.Context, input *accessanalyzer.ListAnalyzedResourcesInput, opts ...request.Option) (*accessanalyzer.ListAnalyzedResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAnalyzedResourcesOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAnalyzedResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAnalyzers(input *accessanalyzer.ListAnalyzersInput) (*accessanalyzer.ListAnalyzersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAnalyzersOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAnalyzers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListAnalyzersPages(input *accessanalyzer.ListAnalyzersInput, fn func(*accessanalyzer.ListAnalyzersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAnalyzersOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAnalyzersOutput{}
	fnCacher := func(out *accessanalyzer.ListAnalyzersOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAnalyzersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAnalyzersPagesWithContext(ctx context.Context, input *accessanalyzer.ListAnalyzersInput, fn func(*accessanalyzer.ListAnalyzersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListAnalyzersOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListAnalyzersOutput{}
	fnCacher := func(out *accessanalyzer.ListAnalyzersOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListAnalyzersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListAnalyzersWithContext(ctx context.Context, input *accessanalyzer.ListAnalyzersInput, opts ...request.Option) (*accessanalyzer.ListAnalyzersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListAnalyzersOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListAnalyzersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListArchiveRules(input *accessanalyzer.ListArchiveRulesInput) (*accessanalyzer.ListArchiveRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListArchiveRulesOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListArchiveRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListArchiveRulesPages(input *accessanalyzer.ListArchiveRulesInput, fn func(*accessanalyzer.ListArchiveRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListArchiveRulesOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListArchiveRulesOutput{}
	fnCacher := func(out *accessanalyzer.ListArchiveRulesOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListArchiveRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListArchiveRulesPagesWithContext(ctx context.Context, input *accessanalyzer.ListArchiveRulesInput, fn func(*accessanalyzer.ListArchiveRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListArchiveRulesOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListArchiveRulesOutput{}
	fnCacher := func(out *accessanalyzer.ListArchiveRulesOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListArchiveRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListArchiveRulesWithContext(ctx context.Context, input *accessanalyzer.ListArchiveRulesInput, opts ...request.Option) (*accessanalyzer.ListArchiveRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListArchiveRulesOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListArchiveRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListFindings(input *accessanalyzer.ListFindingsInput) (*accessanalyzer.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListFindingsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListFindingsPages(input *accessanalyzer.ListFindingsInput, fn func(*accessanalyzer.ListFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListFindingsOutput{}
	fnCacher := func(out *accessanalyzer.ListFindingsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListFindingsPagesWithContext(ctx context.Context, input *accessanalyzer.ListFindingsInput, fn func(*accessanalyzer.ListFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListFindingsOutput{}
	fnCacher := func(out *accessanalyzer.ListFindingsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListFindingsWithContext(ctx context.Context, input *accessanalyzer.ListFindingsInput, opts ...request.Option) (*accessanalyzer.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListFindingsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListPolicyGenerations(input *accessanalyzer.ListPolicyGenerationsInput) (*accessanalyzer.ListPolicyGenerationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListPolicyGenerationsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListPolicyGenerations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListPolicyGenerationsPages(input *accessanalyzer.ListPolicyGenerationsInput, fn func(*accessanalyzer.ListPolicyGenerationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListPolicyGenerationsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListPolicyGenerationsOutput{}
	fnCacher := func(out *accessanalyzer.ListPolicyGenerationsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListPolicyGenerationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListPolicyGenerationsPagesWithContext(ctx context.Context, input *accessanalyzer.ListPolicyGenerationsInput, fn func(*accessanalyzer.ListPolicyGenerationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*accessanalyzer.ListPolicyGenerationsOutput), false)
		return nil
	}
	cachable := true
	output := &accessanalyzer.ListPolicyGenerationsOutput{}
	fnCacher := func(out *accessanalyzer.ListPolicyGenerationsOutput, more bool) bool {
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
	if err := c.AccessAnalyzerAPI.ListPolicyGenerationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AccessAnalyzer) ListPolicyGenerationsWithContext(ctx context.Context, input *accessanalyzer.ListPolicyGenerationsInput, opts ...request.Option) (*accessanalyzer.ListPolicyGenerationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListPolicyGenerationsOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListPolicyGenerationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListTagsForResource(input *accessanalyzer.ListTagsForResourceInput) (*accessanalyzer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListTagsForResourceOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AccessAnalyzer) ListTagsForResourceWithContext(ctx context.Context, input *accessanalyzer.ListTagsForResourceInput, opts ...request.Option) (*accessanalyzer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*accessanalyzer.ListTagsForResourceOutput), nil
	}
	out, err := c.AccessAnalyzerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
