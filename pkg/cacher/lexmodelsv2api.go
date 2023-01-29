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
	"github.com/aws/aws-sdk-go/service/lexmodelsv2"
	"github.com/aws/aws-sdk-go/service/lexmodelsv2/lexmodelsv2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LexModelsV2 struct {
	lexmodelsv2iface.LexModelsV2API
	cache *cache.Cache
}

func NewLexModelsV2(lexmodelsv2api lexmodelsv2iface.LexModelsV2API) *LexModelsV2 {
	return &LexModelsV2{
		LexModelsV2API: lexmodelsv2api,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LexModelsV2) BatchCreateCustomVocabularyItem(input *lexmodelsv2.BatchCreateCustomVocabularyItemInput) (*lexmodelsv2.BatchCreateCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchCreateCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchCreateCustomVocabularyItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) BatchCreateCustomVocabularyItemWithContext(ctx context.Context, input *lexmodelsv2.BatchCreateCustomVocabularyItemInput, opts ...request.Option) (*lexmodelsv2.BatchCreateCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchCreateCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchCreateCustomVocabularyItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) BatchDeleteCustomVocabularyItem(input *lexmodelsv2.BatchDeleteCustomVocabularyItemInput) (*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchDeleteCustomVocabularyItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) BatchDeleteCustomVocabularyItemWithContext(ctx context.Context, input *lexmodelsv2.BatchDeleteCustomVocabularyItemInput, opts ...request.Option) (*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchDeleteCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchDeleteCustomVocabularyItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) BatchUpdateCustomVocabularyItem(input *lexmodelsv2.BatchUpdateCustomVocabularyItemInput) (*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchUpdateCustomVocabularyItem(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) BatchUpdateCustomVocabularyItemWithContext(ctx context.Context, input *lexmodelsv2.BatchUpdateCustomVocabularyItemInput, opts ...request.Option) (*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.BatchUpdateCustomVocabularyItemOutput), nil
	}
	out, err := c.LexModelsV2API.BatchUpdateCustomVocabularyItemWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBot(input *lexmodelsv2.DescribeBotInput) (*lexmodelsv2.DescribeBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotAlias(input *lexmodelsv2.DescribeBotAliasInput) (*lexmodelsv2.DescribeBotAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotAliasOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotAliasWithContext(ctx context.Context, input *lexmodelsv2.DescribeBotAliasInput, opts ...request.Option) (*lexmodelsv2.DescribeBotAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotAliasOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotLocale(input *lexmodelsv2.DescribeBotLocaleInput) (*lexmodelsv2.DescribeBotLocaleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotLocaleOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotLocale(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotLocaleWithContext(ctx context.Context, input *lexmodelsv2.DescribeBotLocaleInput, opts ...request.Option) (*lexmodelsv2.DescribeBotLocaleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotLocaleOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotLocaleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotRecommendation(input *lexmodelsv2.DescribeBotRecommendationInput) (*lexmodelsv2.DescribeBotRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotRecommendationOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotRecommendation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotRecommendationWithContext(ctx context.Context, input *lexmodelsv2.DescribeBotRecommendationInput, opts ...request.Option) (*lexmodelsv2.DescribeBotRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotRecommendationOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotRecommendationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotVersion(input *lexmodelsv2.DescribeBotVersionInput) (*lexmodelsv2.DescribeBotVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotVersionOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotVersionWithContext(ctx context.Context, input *lexmodelsv2.DescribeBotVersionInput, opts ...request.Option) (*lexmodelsv2.DescribeBotVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotVersionOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeBotWithContext(ctx context.Context, input *lexmodelsv2.DescribeBotInput, opts ...request.Option) (*lexmodelsv2.DescribeBotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeBotOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeBotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeCustomVocabularyMetadata(input *lexmodelsv2.DescribeCustomVocabularyMetadataInput) (*lexmodelsv2.DescribeCustomVocabularyMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeCustomVocabularyMetadataOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeCustomVocabularyMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeCustomVocabularyMetadataWithContext(ctx context.Context, input *lexmodelsv2.DescribeCustomVocabularyMetadataInput, opts ...request.Option) (*lexmodelsv2.DescribeCustomVocabularyMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeCustomVocabularyMetadataOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeCustomVocabularyMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeExport(input *lexmodelsv2.DescribeExportInput) (*lexmodelsv2.DescribeExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeExportOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeExport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeExportWithContext(ctx context.Context, input *lexmodelsv2.DescribeExportInput, opts ...request.Option) (*lexmodelsv2.DescribeExportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeExportOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeExportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeImport(input *lexmodelsv2.DescribeImportInput) (*lexmodelsv2.DescribeImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeImportOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeImport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeImportWithContext(ctx context.Context, input *lexmodelsv2.DescribeImportInput, opts ...request.Option) (*lexmodelsv2.DescribeImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeImportOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeImportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeIntent(input *lexmodelsv2.DescribeIntentInput) (*lexmodelsv2.DescribeIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeIntentOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeIntent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeIntentWithContext(ctx context.Context, input *lexmodelsv2.DescribeIntentInput, opts ...request.Option) (*lexmodelsv2.DescribeIntentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeIntentOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeIntentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeResourcePolicy(input *lexmodelsv2.DescribeResourcePolicyInput) (*lexmodelsv2.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeResourcePolicyOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeResourcePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeResourcePolicyWithContext(ctx context.Context, input *lexmodelsv2.DescribeResourcePolicyInput, opts ...request.Option) (*lexmodelsv2.DescribeResourcePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeResourcePolicyOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeResourcePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeSlot(input *lexmodelsv2.DescribeSlotInput) (*lexmodelsv2.DescribeSlotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeSlotOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeSlot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeSlotType(input *lexmodelsv2.DescribeSlotTypeInput) (*lexmodelsv2.DescribeSlotTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeSlotTypeOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeSlotType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeSlotTypeWithContext(ctx context.Context, input *lexmodelsv2.DescribeSlotTypeInput, opts ...request.Option) (*lexmodelsv2.DescribeSlotTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeSlotTypeOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeSlotTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) DescribeSlotWithContext(ctx context.Context, input *lexmodelsv2.DescribeSlotInput, opts ...request.Option) (*lexmodelsv2.DescribeSlotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.DescribeSlotOutput), nil
	}
	out, err := c.LexModelsV2API.DescribeSlotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListAggregatedUtterances(input *lexmodelsv2.ListAggregatedUtterancesInput) (*lexmodelsv2.ListAggregatedUtterancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListAggregatedUtterancesOutput), nil
	}
	out, err := c.LexModelsV2API.ListAggregatedUtterances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListAggregatedUtterancesPages(input *lexmodelsv2.ListAggregatedUtterancesInput, fn func(*lexmodelsv2.ListAggregatedUtterancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListAggregatedUtterancesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListAggregatedUtterancesOutput{}
	fnCacher := func(out *lexmodelsv2.ListAggregatedUtterancesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListAggregatedUtterancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListAggregatedUtterancesPagesWithContext(ctx context.Context, input *lexmodelsv2.ListAggregatedUtterancesInput, fn func(*lexmodelsv2.ListAggregatedUtterancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListAggregatedUtterancesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListAggregatedUtterancesOutput{}
	fnCacher := func(out *lexmodelsv2.ListAggregatedUtterancesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListAggregatedUtterancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListAggregatedUtterancesWithContext(ctx context.Context, input *lexmodelsv2.ListAggregatedUtterancesInput, opts ...request.Option) (*lexmodelsv2.ListAggregatedUtterancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListAggregatedUtterancesOutput), nil
	}
	out, err := c.LexModelsV2API.ListAggregatedUtterancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotAliases(input *lexmodelsv2.ListBotAliasesInput) (*lexmodelsv2.ListBotAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotAliasesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotAliasesPages(input *lexmodelsv2.ListBotAliasesInput, fn func(*lexmodelsv2.ListBotAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotAliasesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotAliasesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotAliasesPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBotAliasesInput, fn func(*lexmodelsv2.ListBotAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotAliasesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotAliasesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotAliasesWithContext(ctx context.Context, input *lexmodelsv2.ListBotAliasesInput, opts ...request.Option) (*lexmodelsv2.ListBotAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotAliasesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotLocales(input *lexmodelsv2.ListBotLocalesInput) (*lexmodelsv2.ListBotLocalesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotLocalesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotLocales(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotLocalesPages(input *lexmodelsv2.ListBotLocalesInput, fn func(*lexmodelsv2.ListBotLocalesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotLocalesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotLocalesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotLocalesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotLocalesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotLocalesPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBotLocalesInput, fn func(*lexmodelsv2.ListBotLocalesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotLocalesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotLocalesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotLocalesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotLocalesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotLocalesWithContext(ctx context.Context, input *lexmodelsv2.ListBotLocalesInput, opts ...request.Option) (*lexmodelsv2.ListBotLocalesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotLocalesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotLocalesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotRecommendations(input *lexmodelsv2.ListBotRecommendationsInput) (*lexmodelsv2.ListBotRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotRecommendationsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotRecommendations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotRecommendationsPages(input *lexmodelsv2.ListBotRecommendationsInput, fn func(*lexmodelsv2.ListBotRecommendationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotRecommendationsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotRecommendationsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotRecommendationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotRecommendationsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBotRecommendationsInput, fn func(*lexmodelsv2.ListBotRecommendationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotRecommendationsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotRecommendationsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotRecommendationsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotRecommendationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotRecommendationsWithContext(ctx context.Context, input *lexmodelsv2.ListBotRecommendationsInput, opts ...request.Option) (*lexmodelsv2.ListBotRecommendationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotRecommendationsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotRecommendationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotVersions(input *lexmodelsv2.ListBotVersionsInput) (*lexmodelsv2.ListBotVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotVersionsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotVersionsPages(input *lexmodelsv2.ListBotVersionsInput, fn func(*lexmodelsv2.ListBotVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotVersionsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotVersionsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotVersionsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBotVersionsInput, fn func(*lexmodelsv2.ListBotVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotVersionsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotVersionsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotVersionsWithContext(ctx context.Context, input *lexmodelsv2.ListBotVersionsInput, opts ...request.Option) (*lexmodelsv2.ListBotVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotVersionsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBots(input *lexmodelsv2.ListBotsInput) (*lexmodelsv2.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBotsPages(input *lexmodelsv2.ListBotsInput, fn func(*lexmodelsv2.ListBotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBotsInput, fn func(*lexmodelsv2.ListBotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBotsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBotsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBotsWithContext(ctx context.Context, input *lexmodelsv2.ListBotsInput, opts ...request.Option) (*lexmodelsv2.ListBotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBotsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBuiltInIntents(input *lexmodelsv2.ListBuiltInIntentsInput) (*lexmodelsv2.ListBuiltInIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBuiltInIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBuiltInIntents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBuiltInIntentsPages(input *lexmodelsv2.ListBuiltInIntentsInput, fn func(*lexmodelsv2.ListBuiltInIntentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBuiltInIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBuiltInIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBuiltInIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBuiltInIntentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBuiltInIntentsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBuiltInIntentsInput, fn func(*lexmodelsv2.ListBuiltInIntentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBuiltInIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBuiltInIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListBuiltInIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBuiltInIntentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBuiltInIntentsWithContext(ctx context.Context, input *lexmodelsv2.ListBuiltInIntentsInput, opts ...request.Option) (*lexmodelsv2.ListBuiltInIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBuiltInIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListBuiltInIntentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBuiltInSlotTypes(input *lexmodelsv2.ListBuiltInSlotTypesInput) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBuiltInSlotTypesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBuiltInSlotTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListBuiltInSlotTypesPages(input *lexmodelsv2.ListBuiltInSlotTypesInput, fn func(*lexmodelsv2.ListBuiltInSlotTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBuiltInSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBuiltInSlotTypesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBuiltInSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBuiltInSlotTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBuiltInSlotTypesPagesWithContext(ctx context.Context, input *lexmodelsv2.ListBuiltInSlotTypesInput, fn func(*lexmodelsv2.ListBuiltInSlotTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListBuiltInSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListBuiltInSlotTypesOutput{}
	fnCacher := func(out *lexmodelsv2.ListBuiltInSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListBuiltInSlotTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListBuiltInSlotTypesWithContext(ctx context.Context, input *lexmodelsv2.ListBuiltInSlotTypesInput, opts ...request.Option) (*lexmodelsv2.ListBuiltInSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListBuiltInSlotTypesOutput), nil
	}
	out, err := c.LexModelsV2API.ListBuiltInSlotTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListCustomVocabularyItems(input *lexmodelsv2.ListCustomVocabularyItemsInput) (*lexmodelsv2.ListCustomVocabularyItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListCustomVocabularyItemsOutput), nil
	}
	out, err := c.LexModelsV2API.ListCustomVocabularyItems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListCustomVocabularyItemsPages(input *lexmodelsv2.ListCustomVocabularyItemsInput, fn func(*lexmodelsv2.ListCustomVocabularyItemsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListCustomVocabularyItemsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListCustomVocabularyItemsOutput{}
	fnCacher := func(out *lexmodelsv2.ListCustomVocabularyItemsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListCustomVocabularyItemsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListCustomVocabularyItemsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListCustomVocabularyItemsInput, fn func(*lexmodelsv2.ListCustomVocabularyItemsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListCustomVocabularyItemsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListCustomVocabularyItemsOutput{}
	fnCacher := func(out *lexmodelsv2.ListCustomVocabularyItemsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListCustomVocabularyItemsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListCustomVocabularyItemsWithContext(ctx context.Context, input *lexmodelsv2.ListCustomVocabularyItemsInput, opts ...request.Option) (*lexmodelsv2.ListCustomVocabularyItemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListCustomVocabularyItemsOutput), nil
	}
	out, err := c.LexModelsV2API.ListCustomVocabularyItemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListExports(input *lexmodelsv2.ListExportsInput) (*lexmodelsv2.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListExportsOutput), nil
	}
	out, err := c.LexModelsV2API.ListExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListExportsPages(input *lexmodelsv2.ListExportsInput, fn func(*lexmodelsv2.ListExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListExportsOutput{}
	fnCacher := func(out *lexmodelsv2.ListExportsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListExportsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListExportsInput, fn func(*lexmodelsv2.ListExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListExportsOutput{}
	fnCacher := func(out *lexmodelsv2.ListExportsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListExportsWithContext(ctx context.Context, input *lexmodelsv2.ListExportsInput, opts ...request.Option) (*lexmodelsv2.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListExportsOutput), nil
	}
	out, err := c.LexModelsV2API.ListExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListImports(input *lexmodelsv2.ListImportsInput) (*lexmodelsv2.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListImportsOutput), nil
	}
	out, err := c.LexModelsV2API.ListImports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListImportsPages(input *lexmodelsv2.ListImportsInput, fn func(*lexmodelsv2.ListImportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListImportsOutput{}
	fnCacher := func(out *lexmodelsv2.ListImportsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListImportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListImportsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListImportsInput, fn func(*lexmodelsv2.ListImportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListImportsOutput{}
	fnCacher := func(out *lexmodelsv2.ListImportsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListImportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListImportsWithContext(ctx context.Context, input *lexmodelsv2.ListImportsInput, opts ...request.Option) (*lexmodelsv2.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListImportsOutput), nil
	}
	out, err := c.LexModelsV2API.ListImportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListIntents(input *lexmodelsv2.ListIntentsInput) (*lexmodelsv2.ListIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListIntents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListIntentsPages(input *lexmodelsv2.ListIntentsInput, fn func(*lexmodelsv2.ListIntentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListIntentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListIntentsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListIntentsInput, fn func(*lexmodelsv2.ListIntentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListIntentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListIntentsWithContext(ctx context.Context, input *lexmodelsv2.ListIntentsInput, opts ...request.Option) (*lexmodelsv2.ListIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListIntentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListRecommendedIntents(input *lexmodelsv2.ListRecommendedIntentsInput) (*lexmodelsv2.ListRecommendedIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListRecommendedIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListRecommendedIntents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListRecommendedIntentsPages(input *lexmodelsv2.ListRecommendedIntentsInput, fn func(*lexmodelsv2.ListRecommendedIntentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListRecommendedIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListRecommendedIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListRecommendedIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListRecommendedIntentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListRecommendedIntentsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListRecommendedIntentsInput, fn func(*lexmodelsv2.ListRecommendedIntentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListRecommendedIntentsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListRecommendedIntentsOutput{}
	fnCacher := func(out *lexmodelsv2.ListRecommendedIntentsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListRecommendedIntentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListRecommendedIntentsWithContext(ctx context.Context, input *lexmodelsv2.ListRecommendedIntentsInput, opts ...request.Option) (*lexmodelsv2.ListRecommendedIntentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListRecommendedIntentsOutput), nil
	}
	out, err := c.LexModelsV2API.ListRecommendedIntentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListSlotTypes(input *lexmodelsv2.ListSlotTypesInput) (*lexmodelsv2.ListSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListSlotTypesOutput), nil
	}
	out, err := c.LexModelsV2API.ListSlotTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListSlotTypesPages(input *lexmodelsv2.ListSlotTypesInput, fn func(*lexmodelsv2.ListSlotTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListSlotTypesOutput{}
	fnCacher := func(out *lexmodelsv2.ListSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListSlotTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListSlotTypesPagesWithContext(ctx context.Context, input *lexmodelsv2.ListSlotTypesInput, fn func(*lexmodelsv2.ListSlotTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListSlotTypesOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListSlotTypesOutput{}
	fnCacher := func(out *lexmodelsv2.ListSlotTypesOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListSlotTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListSlotTypesWithContext(ctx context.Context, input *lexmodelsv2.ListSlotTypesInput, opts ...request.Option) (*lexmodelsv2.ListSlotTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListSlotTypesOutput), nil
	}
	out, err := c.LexModelsV2API.ListSlotTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListSlots(input *lexmodelsv2.ListSlotsInput) (*lexmodelsv2.ListSlotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListSlotsOutput), nil
	}
	out, err := c.LexModelsV2API.ListSlots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListSlotsPages(input *lexmodelsv2.ListSlotsInput, fn func(*lexmodelsv2.ListSlotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListSlotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListSlotsOutput{}
	fnCacher := func(out *lexmodelsv2.ListSlotsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListSlotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListSlotsPagesWithContext(ctx context.Context, input *lexmodelsv2.ListSlotsInput, fn func(*lexmodelsv2.ListSlotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lexmodelsv2.ListSlotsOutput), false)
		return nil
	}
	cachable := true
	output := &lexmodelsv2.ListSlotsOutput{}
	fnCacher := func(out *lexmodelsv2.ListSlotsOutput, more bool) bool {
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
	if err := c.LexModelsV2API.ListSlotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LexModelsV2) ListSlotsWithContext(ctx context.Context, input *lexmodelsv2.ListSlotsInput, opts ...request.Option) (*lexmodelsv2.ListSlotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListSlotsOutput), nil
	}
	out, err := c.LexModelsV2API.ListSlotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListTagsForResource(input *lexmodelsv2.ListTagsForResourceInput) (*lexmodelsv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListTagsForResourceOutput), nil
	}
	out, err := c.LexModelsV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) ListTagsForResourceWithContext(ctx context.Context, input *lexmodelsv2.ListTagsForResourceInput, opts ...request.Option) (*lexmodelsv2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.ListTagsForResourceOutput), nil
	}
	out, err := c.LexModelsV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) SearchAssociatedTranscripts(input *lexmodelsv2.SearchAssociatedTranscriptsInput) (*lexmodelsv2.SearchAssociatedTranscriptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.SearchAssociatedTranscriptsOutput), nil
	}
	out, err := c.LexModelsV2API.SearchAssociatedTranscripts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LexModelsV2) SearchAssociatedTranscriptsWithContext(ctx context.Context, input *lexmodelsv2.SearchAssociatedTranscriptsInput, opts ...request.Option) (*lexmodelsv2.SearchAssociatedTranscriptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lexmodelsv2.SearchAssociatedTranscriptsOutput), nil
	}
	out, err := c.LexModelsV2API.SearchAssociatedTranscriptsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
