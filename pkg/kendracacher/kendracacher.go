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
package kendracacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kendra"
	"github.com/aws/aws-sdk-go/service/kendra/kendraiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Kendra struct {
	kendraiface.KendraAPI
	cache *cache.Cache
}

func New(kendraapi kendraiface.KendraAPI) *Kendra {
	return &Kendra{
		KendraAPI: kendraapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Kendra) BatchDeleteDocument(input *kendra.BatchDeleteDocumentInput) (*kendra.BatchDeleteDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchDeleteDocumentOutput), nil
	}
	out, err := c.KendraAPI.BatchDeleteDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) BatchDeleteDocumentWithContext(ctx context.Context, input *kendra.BatchDeleteDocumentInput, opts ...request.Option) (*kendra.BatchDeleteDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchDeleteDocumentOutput), nil
	}
	out, err := c.KendraAPI.BatchDeleteDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) BatchGetDocumentStatus(input *kendra.BatchGetDocumentStatusInput) (*kendra.BatchGetDocumentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchGetDocumentStatusOutput), nil
	}
	out, err := c.KendraAPI.BatchGetDocumentStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) BatchGetDocumentStatusWithContext(ctx context.Context, input *kendra.BatchGetDocumentStatusInput, opts ...request.Option) (*kendra.BatchGetDocumentStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchGetDocumentStatusOutput), nil
	}
	out, err := c.KendraAPI.BatchGetDocumentStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) BatchPutDocument(input *kendra.BatchPutDocumentInput) (*kendra.BatchPutDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchPutDocumentOutput), nil
	}
	out, err := c.KendraAPI.BatchPutDocument(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) BatchPutDocumentWithContext(ctx context.Context, input *kendra.BatchPutDocumentInput, opts ...request.Option) (*kendra.BatchPutDocumentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.BatchPutDocumentOutput), nil
	}
	out, err := c.KendraAPI.BatchPutDocumentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeAccessControlConfiguration(input *kendra.DescribeAccessControlConfigurationInput) (*kendra.DescribeAccessControlConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeAccessControlConfigurationOutput), nil
	}
	out, err := c.KendraAPI.DescribeAccessControlConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeAccessControlConfigurationWithContext(ctx context.Context, input *kendra.DescribeAccessControlConfigurationInput, opts ...request.Option) (*kendra.DescribeAccessControlConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeAccessControlConfigurationOutput), nil
	}
	out, err := c.KendraAPI.DescribeAccessControlConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeDataSource(input *kendra.DescribeDataSourceInput) (*kendra.DescribeDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeDataSourceOutput), nil
	}
	out, err := c.KendraAPI.DescribeDataSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeDataSourceWithContext(ctx context.Context, input *kendra.DescribeDataSourceInput, opts ...request.Option) (*kendra.DescribeDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeDataSourceOutput), nil
	}
	out, err := c.KendraAPI.DescribeDataSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeExperience(input *kendra.DescribeExperienceInput) (*kendra.DescribeExperienceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeExperienceOutput), nil
	}
	out, err := c.KendraAPI.DescribeExperience(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeExperienceWithContext(ctx context.Context, input *kendra.DescribeExperienceInput, opts ...request.Option) (*kendra.DescribeExperienceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeExperienceOutput), nil
	}
	out, err := c.KendraAPI.DescribeExperienceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeFaq(input *kendra.DescribeFaqInput) (*kendra.DescribeFaqOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeFaqOutput), nil
	}
	out, err := c.KendraAPI.DescribeFaq(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeFaqWithContext(ctx context.Context, input *kendra.DescribeFaqInput, opts ...request.Option) (*kendra.DescribeFaqOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeFaqOutput), nil
	}
	out, err := c.KendraAPI.DescribeFaqWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeIndex(input *kendra.DescribeIndexInput) (*kendra.DescribeIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeIndexOutput), nil
	}
	out, err := c.KendraAPI.DescribeIndex(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeIndexWithContext(ctx context.Context, input *kendra.DescribeIndexInput, opts ...request.Option) (*kendra.DescribeIndexOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeIndexOutput), nil
	}
	out, err := c.KendraAPI.DescribeIndexWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribePrincipalMapping(input *kendra.DescribePrincipalMappingInput) (*kendra.DescribePrincipalMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribePrincipalMappingOutput), nil
	}
	out, err := c.KendraAPI.DescribePrincipalMapping(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribePrincipalMappingWithContext(ctx context.Context, input *kendra.DescribePrincipalMappingInput, opts ...request.Option) (*kendra.DescribePrincipalMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribePrincipalMappingOutput), nil
	}
	out, err := c.KendraAPI.DescribePrincipalMappingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeQuerySuggestionsBlockList(input *kendra.DescribeQuerySuggestionsBlockListInput) (*kendra.DescribeQuerySuggestionsBlockListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeQuerySuggestionsBlockListOutput), nil
	}
	out, err := c.KendraAPI.DescribeQuerySuggestionsBlockList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeQuerySuggestionsBlockListWithContext(ctx context.Context, input *kendra.DescribeQuerySuggestionsBlockListInput, opts ...request.Option) (*kendra.DescribeQuerySuggestionsBlockListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeQuerySuggestionsBlockListOutput), nil
	}
	out, err := c.KendraAPI.DescribeQuerySuggestionsBlockListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeQuerySuggestionsConfig(input *kendra.DescribeQuerySuggestionsConfigInput) (*kendra.DescribeQuerySuggestionsConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeQuerySuggestionsConfigOutput), nil
	}
	out, err := c.KendraAPI.DescribeQuerySuggestionsConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeQuerySuggestionsConfigWithContext(ctx context.Context, input *kendra.DescribeQuerySuggestionsConfigInput, opts ...request.Option) (*kendra.DescribeQuerySuggestionsConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeQuerySuggestionsConfigOutput), nil
	}
	out, err := c.KendraAPI.DescribeQuerySuggestionsConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeThesaurus(input *kendra.DescribeThesaurusInput) (*kendra.DescribeThesaurusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeThesaurusOutput), nil
	}
	out, err := c.KendraAPI.DescribeThesaurus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) DescribeThesaurusWithContext(ctx context.Context, input *kendra.DescribeThesaurusInput, opts ...request.Option) (*kendra.DescribeThesaurusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.DescribeThesaurusOutput), nil
	}
	out, err := c.KendraAPI.DescribeThesaurusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) GetQuerySuggestions(input *kendra.GetQuerySuggestionsInput) (*kendra.GetQuerySuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.GetQuerySuggestionsOutput), nil
	}
	out, err := c.KendraAPI.GetQuerySuggestions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) GetQuerySuggestionsWithContext(ctx context.Context, input *kendra.GetQuerySuggestionsInput, opts ...request.Option) (*kendra.GetQuerySuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.GetQuerySuggestionsOutput), nil
	}
	out, err := c.KendraAPI.GetQuerySuggestionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) GetSnapshots(input *kendra.GetSnapshotsInput) (*kendra.GetSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.GetSnapshotsOutput), nil
	}
	out, err := c.KendraAPI.GetSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) GetSnapshotsPages(input *kendra.GetSnapshotsInput, fn func(*kendra.GetSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.GetSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.GetSnapshotsOutput{}
	fnCacher := func(out *kendra.GetSnapshotsOutput, more bool) bool {
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
	if err := c.KendraAPI.GetSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) GetSnapshotsPagesWithContext(ctx context.Context, input *kendra.GetSnapshotsInput, fn func(*kendra.GetSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.GetSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.GetSnapshotsOutput{}
	fnCacher := func(out *kendra.GetSnapshotsOutput, more bool) bool {
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
	if err := c.KendraAPI.GetSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) GetSnapshotsWithContext(ctx context.Context, input *kendra.GetSnapshotsInput, opts ...request.Option) (*kendra.GetSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.GetSnapshotsOutput), nil
	}
	out, err := c.KendraAPI.GetSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListAccessControlConfigurations(input *kendra.ListAccessControlConfigurationsInput) (*kendra.ListAccessControlConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListAccessControlConfigurationsOutput), nil
	}
	out, err := c.KendraAPI.ListAccessControlConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListAccessControlConfigurationsPages(input *kendra.ListAccessControlConfigurationsInput, fn func(*kendra.ListAccessControlConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListAccessControlConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListAccessControlConfigurationsOutput{}
	fnCacher := func(out *kendra.ListAccessControlConfigurationsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListAccessControlConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListAccessControlConfigurationsPagesWithContext(ctx context.Context, input *kendra.ListAccessControlConfigurationsInput, fn func(*kendra.ListAccessControlConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListAccessControlConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListAccessControlConfigurationsOutput{}
	fnCacher := func(out *kendra.ListAccessControlConfigurationsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListAccessControlConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListAccessControlConfigurationsWithContext(ctx context.Context, input *kendra.ListAccessControlConfigurationsInput, opts ...request.Option) (*kendra.ListAccessControlConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListAccessControlConfigurationsOutput), nil
	}
	out, err := c.KendraAPI.ListAccessControlConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListDataSourceSyncJobs(input *kendra.ListDataSourceSyncJobsInput) (*kendra.ListDataSourceSyncJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListDataSourceSyncJobsOutput), nil
	}
	out, err := c.KendraAPI.ListDataSourceSyncJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListDataSourceSyncJobsPages(input *kendra.ListDataSourceSyncJobsInput, fn func(*kendra.ListDataSourceSyncJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListDataSourceSyncJobsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListDataSourceSyncJobsOutput{}
	fnCacher := func(out *kendra.ListDataSourceSyncJobsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListDataSourceSyncJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListDataSourceSyncJobsPagesWithContext(ctx context.Context, input *kendra.ListDataSourceSyncJobsInput, fn func(*kendra.ListDataSourceSyncJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListDataSourceSyncJobsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListDataSourceSyncJobsOutput{}
	fnCacher := func(out *kendra.ListDataSourceSyncJobsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListDataSourceSyncJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListDataSourceSyncJobsWithContext(ctx context.Context, input *kendra.ListDataSourceSyncJobsInput, opts ...request.Option) (*kendra.ListDataSourceSyncJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListDataSourceSyncJobsOutput), nil
	}
	out, err := c.KendraAPI.ListDataSourceSyncJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListDataSources(input *kendra.ListDataSourcesInput) (*kendra.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListDataSourcesOutput), nil
	}
	out, err := c.KendraAPI.ListDataSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListDataSourcesPages(input *kendra.ListDataSourcesInput, fn func(*kendra.ListDataSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListDataSourcesOutput{}
	fnCacher := func(out *kendra.ListDataSourcesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListDataSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListDataSourcesPagesWithContext(ctx context.Context, input *kendra.ListDataSourcesInput, fn func(*kendra.ListDataSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListDataSourcesOutput{}
	fnCacher := func(out *kendra.ListDataSourcesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListDataSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListDataSourcesWithContext(ctx context.Context, input *kendra.ListDataSourcesInput, opts ...request.Option) (*kendra.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListDataSourcesOutput), nil
	}
	out, err := c.KendraAPI.ListDataSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListEntityPersonas(input *kendra.ListEntityPersonasInput) (*kendra.ListEntityPersonasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListEntityPersonasOutput), nil
	}
	out, err := c.KendraAPI.ListEntityPersonas(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListEntityPersonasPages(input *kendra.ListEntityPersonasInput, fn func(*kendra.ListEntityPersonasOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListEntityPersonasOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListEntityPersonasOutput{}
	fnCacher := func(out *kendra.ListEntityPersonasOutput, more bool) bool {
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
	if err := c.KendraAPI.ListEntityPersonasPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListEntityPersonasPagesWithContext(ctx context.Context, input *kendra.ListEntityPersonasInput, fn func(*kendra.ListEntityPersonasOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListEntityPersonasOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListEntityPersonasOutput{}
	fnCacher := func(out *kendra.ListEntityPersonasOutput, more bool) bool {
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
	if err := c.KendraAPI.ListEntityPersonasPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListEntityPersonasWithContext(ctx context.Context, input *kendra.ListEntityPersonasInput, opts ...request.Option) (*kendra.ListEntityPersonasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListEntityPersonasOutput), nil
	}
	out, err := c.KendraAPI.ListEntityPersonasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListExperienceEntities(input *kendra.ListExperienceEntitiesInput) (*kendra.ListExperienceEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListExperienceEntitiesOutput), nil
	}
	out, err := c.KendraAPI.ListExperienceEntities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListExperienceEntitiesPages(input *kendra.ListExperienceEntitiesInput, fn func(*kendra.ListExperienceEntitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListExperienceEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListExperienceEntitiesOutput{}
	fnCacher := func(out *kendra.ListExperienceEntitiesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListExperienceEntitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListExperienceEntitiesPagesWithContext(ctx context.Context, input *kendra.ListExperienceEntitiesInput, fn func(*kendra.ListExperienceEntitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListExperienceEntitiesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListExperienceEntitiesOutput{}
	fnCacher := func(out *kendra.ListExperienceEntitiesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListExperienceEntitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListExperienceEntitiesWithContext(ctx context.Context, input *kendra.ListExperienceEntitiesInput, opts ...request.Option) (*kendra.ListExperienceEntitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListExperienceEntitiesOutput), nil
	}
	out, err := c.KendraAPI.ListExperienceEntitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListExperiences(input *kendra.ListExperiencesInput) (*kendra.ListExperiencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListExperiencesOutput), nil
	}
	out, err := c.KendraAPI.ListExperiences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListExperiencesPages(input *kendra.ListExperiencesInput, fn func(*kendra.ListExperiencesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListExperiencesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListExperiencesOutput{}
	fnCacher := func(out *kendra.ListExperiencesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListExperiencesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListExperiencesPagesWithContext(ctx context.Context, input *kendra.ListExperiencesInput, fn func(*kendra.ListExperiencesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListExperiencesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListExperiencesOutput{}
	fnCacher := func(out *kendra.ListExperiencesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListExperiencesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListExperiencesWithContext(ctx context.Context, input *kendra.ListExperiencesInput, opts ...request.Option) (*kendra.ListExperiencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListExperiencesOutput), nil
	}
	out, err := c.KendraAPI.ListExperiencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListFaqs(input *kendra.ListFaqsInput) (*kendra.ListFaqsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListFaqsOutput), nil
	}
	out, err := c.KendraAPI.ListFaqs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListFaqsPages(input *kendra.ListFaqsInput, fn func(*kendra.ListFaqsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListFaqsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListFaqsOutput{}
	fnCacher := func(out *kendra.ListFaqsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListFaqsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListFaqsPagesWithContext(ctx context.Context, input *kendra.ListFaqsInput, fn func(*kendra.ListFaqsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListFaqsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListFaqsOutput{}
	fnCacher := func(out *kendra.ListFaqsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListFaqsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListFaqsWithContext(ctx context.Context, input *kendra.ListFaqsInput, opts ...request.Option) (*kendra.ListFaqsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListFaqsOutput), nil
	}
	out, err := c.KendraAPI.ListFaqsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListGroupsOlderThanOrderingId(input *kendra.ListGroupsOlderThanOrderingIdInput) (*kendra.ListGroupsOlderThanOrderingIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListGroupsOlderThanOrderingIdOutput), nil
	}
	out, err := c.KendraAPI.ListGroupsOlderThanOrderingId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListGroupsOlderThanOrderingIdPages(input *kendra.ListGroupsOlderThanOrderingIdInput, fn func(*kendra.ListGroupsOlderThanOrderingIdOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListGroupsOlderThanOrderingIdOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListGroupsOlderThanOrderingIdOutput{}
	fnCacher := func(out *kendra.ListGroupsOlderThanOrderingIdOutput, more bool) bool {
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
	if err := c.KendraAPI.ListGroupsOlderThanOrderingIdPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListGroupsOlderThanOrderingIdPagesWithContext(ctx context.Context, input *kendra.ListGroupsOlderThanOrderingIdInput, fn func(*kendra.ListGroupsOlderThanOrderingIdOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListGroupsOlderThanOrderingIdOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListGroupsOlderThanOrderingIdOutput{}
	fnCacher := func(out *kendra.ListGroupsOlderThanOrderingIdOutput, more bool) bool {
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
	if err := c.KendraAPI.ListGroupsOlderThanOrderingIdPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListGroupsOlderThanOrderingIdWithContext(ctx context.Context, input *kendra.ListGroupsOlderThanOrderingIdInput, opts ...request.Option) (*kendra.ListGroupsOlderThanOrderingIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListGroupsOlderThanOrderingIdOutput), nil
	}
	out, err := c.KendraAPI.ListGroupsOlderThanOrderingIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListIndices(input *kendra.ListIndicesInput) (*kendra.ListIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListIndicesOutput), nil
	}
	out, err := c.KendraAPI.ListIndices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListIndicesPages(input *kendra.ListIndicesInput, fn func(*kendra.ListIndicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListIndicesOutput{}
	fnCacher := func(out *kendra.ListIndicesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListIndicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListIndicesPagesWithContext(ctx context.Context, input *kendra.ListIndicesInput, fn func(*kendra.ListIndicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListIndicesOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListIndicesOutput{}
	fnCacher := func(out *kendra.ListIndicesOutput, more bool) bool {
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
	if err := c.KendraAPI.ListIndicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListIndicesWithContext(ctx context.Context, input *kendra.ListIndicesInput, opts ...request.Option) (*kendra.ListIndicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListIndicesOutput), nil
	}
	out, err := c.KendraAPI.ListIndicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListQuerySuggestionsBlockLists(input *kendra.ListQuerySuggestionsBlockListsInput) (*kendra.ListQuerySuggestionsBlockListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListQuerySuggestionsBlockListsOutput), nil
	}
	out, err := c.KendraAPI.ListQuerySuggestionsBlockLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListQuerySuggestionsBlockListsPages(input *kendra.ListQuerySuggestionsBlockListsInput, fn func(*kendra.ListQuerySuggestionsBlockListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListQuerySuggestionsBlockListsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListQuerySuggestionsBlockListsOutput{}
	fnCacher := func(out *kendra.ListQuerySuggestionsBlockListsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListQuerySuggestionsBlockListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListQuerySuggestionsBlockListsPagesWithContext(ctx context.Context, input *kendra.ListQuerySuggestionsBlockListsInput, fn func(*kendra.ListQuerySuggestionsBlockListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListQuerySuggestionsBlockListsOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListQuerySuggestionsBlockListsOutput{}
	fnCacher := func(out *kendra.ListQuerySuggestionsBlockListsOutput, more bool) bool {
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
	if err := c.KendraAPI.ListQuerySuggestionsBlockListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListQuerySuggestionsBlockListsWithContext(ctx context.Context, input *kendra.ListQuerySuggestionsBlockListsInput, opts ...request.Option) (*kendra.ListQuerySuggestionsBlockListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListQuerySuggestionsBlockListsOutput), nil
	}
	out, err := c.KendraAPI.ListQuerySuggestionsBlockListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListTagsForResource(input *kendra.ListTagsForResourceInput) (*kendra.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListTagsForResourceOutput), nil
	}
	out, err := c.KendraAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListTagsForResourceWithContext(ctx context.Context, input *kendra.ListTagsForResourceInput, opts ...request.Option) (*kendra.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListTagsForResourceOutput), nil
	}
	out, err := c.KendraAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListThesauri(input *kendra.ListThesauriInput) (*kendra.ListThesauriOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListThesauriOutput), nil
	}
	out, err := c.KendraAPI.ListThesauri(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) ListThesauriPages(input *kendra.ListThesauriInput, fn func(*kendra.ListThesauriOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListThesauriOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListThesauriOutput{}
	fnCacher := func(out *kendra.ListThesauriOutput, more bool) bool {
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
	if err := c.KendraAPI.ListThesauriPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListThesauriPagesWithContext(ctx context.Context, input *kendra.ListThesauriInput, fn func(*kendra.ListThesauriOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*kendra.ListThesauriOutput), false)
		return nil
	}
	cachable := true
	output := &kendra.ListThesauriOutput{}
	fnCacher := func(out *kendra.ListThesauriOutput, more bool) bool {
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
	if err := c.KendraAPI.ListThesauriPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Kendra) ListThesauriWithContext(ctx context.Context, input *kendra.ListThesauriInput, opts ...request.Option) (*kendra.ListThesauriOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.ListThesauriOutput), nil
	}
	out, err := c.KendraAPI.ListThesauriWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) Query(input *kendra.QueryInput) (*kendra.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.QueryOutput), nil
	}
	out, err := c.KendraAPI.Query(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Kendra) QueryWithContext(ctx context.Context, input *kendra.QueryInput, opts ...request.Option) (*kendra.QueryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*kendra.QueryOutput), nil
	}
	out, err := c.KendraAPI.QueryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
