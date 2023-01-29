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
package migrationhubstrategyrecommendationscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/migrationhubstrategyrecommendations"
	"github.com/aws/aws-sdk-go/service/migrationhubstrategyrecommendations/migrationhubstrategyrecommendationsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MigrationHubStrategyRecommendations struct {
	migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI
	cache *cache.Cache
}

func New(migrationhubstrategyrecommendationsapi migrationhubstrategyrecommendationsiface.MigrationHubStrategyRecommendationsAPI) *MigrationHubStrategyRecommendations {
	return &MigrationHubStrategyRecommendations{
		MigrationHubStrategyRecommendationsAPI: migrationhubstrategyrecommendationsapi,
		cache:                                  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MigrationHubStrategyRecommendations) GetApplicationComponentDetails(input *migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput) (*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetApplicationComponentDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetApplicationComponentDetailsWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetApplicationComponentDetailsInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetApplicationComponentDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetApplicationComponentDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetApplicationComponentStrategies(input *migrationhubstrategyrecommendations.GetApplicationComponentStrategiesInput) (*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetApplicationComponentStrategies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetApplicationComponentStrategiesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetApplicationComponentStrategiesInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetApplicationComponentStrategiesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetApplicationComponentStrategiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetAssessment(input *migrationhubstrategyrecommendations.GetAssessmentInput) (*migrationhubstrategyrecommendations.GetAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetAssessmentOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetAssessment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetAssessmentWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetAssessmentInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetAssessmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetAssessmentOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetAssessmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetImportFileTask(input *migrationhubstrategyrecommendations.GetImportFileTaskInput) (*migrationhubstrategyrecommendations.GetImportFileTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetImportFileTaskOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetImportFileTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetImportFileTaskWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetImportFileTaskInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetImportFileTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetImportFileTaskOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetImportFileTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetLatestAssessmentId(input *migrationhubstrategyrecommendations.GetLatestAssessmentIdInput) (*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetLatestAssessmentId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetLatestAssessmentIdWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetLatestAssessmentIdInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetLatestAssessmentIdOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetLatestAssessmentIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetPortfolioPreferences(input *migrationhubstrategyrecommendations.GetPortfolioPreferencesInput) (*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetPortfolioPreferences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetPortfolioPreferencesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetPortfolioPreferencesInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetPortfolioPreferencesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetPortfolioPreferencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetPortfolioSummary(input *migrationhubstrategyrecommendations.GetPortfolioSummaryInput) (*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetPortfolioSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetPortfolioSummaryWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetPortfolioSummaryInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetPortfolioSummaryOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetPortfolioSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetRecommendationReportDetails(input *migrationhubstrategyrecommendations.GetRecommendationReportDetailsInput) (*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetRecommendationReportDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetRecommendationReportDetailsWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetRecommendationReportDetailsInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetRecommendationReportDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetRecommendationReportDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetServerDetails(input *migrationhubstrategyrecommendations.GetServerDetailsInput) (*migrationhubstrategyrecommendations.GetServerDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetServerDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetServerDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetServerDetailsPages(input *migrationhubstrategyrecommendations.GetServerDetailsInput, fn func(*migrationhubstrategyrecommendations.GetServerDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.GetServerDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.GetServerDetailsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.GetServerDetailsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.GetServerDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) GetServerDetailsPagesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetServerDetailsInput, fn func(*migrationhubstrategyrecommendations.GetServerDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.GetServerDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.GetServerDetailsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.GetServerDetailsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.GetServerDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) GetServerDetailsWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetServerDetailsInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetServerDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetServerDetailsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetServerDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetServerStrategies(input *migrationhubstrategyrecommendations.GetServerStrategiesInput) (*migrationhubstrategyrecommendations.GetServerStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetServerStrategiesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetServerStrategies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) GetServerStrategiesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.GetServerStrategiesInput, opts ...request.Option) (*migrationhubstrategyrecommendations.GetServerStrategiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.GetServerStrategiesOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.GetServerStrategiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListApplicationComponents(input *migrationhubstrategyrecommendations.ListApplicationComponentsInput) (*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListApplicationComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListApplicationComponentsPages(input *migrationhubstrategyrecommendations.ListApplicationComponentsInput, fn func(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListApplicationComponentsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListApplicationComponentsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListApplicationComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListApplicationComponentsPagesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListApplicationComponentsInput, fn func(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListApplicationComponentsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListApplicationComponentsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListApplicationComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListApplicationComponentsWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListApplicationComponentsInput, opts ...request.Option) (*migrationhubstrategyrecommendations.ListApplicationComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListApplicationComponentsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListApplicationComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListCollectors(input *migrationhubstrategyrecommendations.ListCollectorsInput) (*migrationhubstrategyrecommendations.ListCollectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListCollectorsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListCollectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListCollectorsPages(input *migrationhubstrategyrecommendations.ListCollectorsInput, fn func(*migrationhubstrategyrecommendations.ListCollectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListCollectorsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListCollectorsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListCollectorsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListCollectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListCollectorsPagesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListCollectorsInput, fn func(*migrationhubstrategyrecommendations.ListCollectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListCollectorsOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListCollectorsOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListCollectorsOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListCollectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListCollectorsWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListCollectorsInput, opts ...request.Option) (*migrationhubstrategyrecommendations.ListCollectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListCollectorsOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListCollectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListImportFileTask(input *migrationhubstrategyrecommendations.ListImportFileTaskInput) (*migrationhubstrategyrecommendations.ListImportFileTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListImportFileTaskOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListImportFileTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListImportFileTaskPages(input *migrationhubstrategyrecommendations.ListImportFileTaskInput, fn func(*migrationhubstrategyrecommendations.ListImportFileTaskOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListImportFileTaskOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListImportFileTaskOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListImportFileTaskOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListImportFileTaskPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListImportFileTaskPagesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListImportFileTaskInput, fn func(*migrationhubstrategyrecommendations.ListImportFileTaskOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListImportFileTaskOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListImportFileTaskOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListImportFileTaskOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListImportFileTaskPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListImportFileTaskWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListImportFileTaskInput, opts ...request.Option) (*migrationhubstrategyrecommendations.ListImportFileTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListImportFileTaskOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListImportFileTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListServers(input *migrationhubstrategyrecommendations.ListServersInput) (*migrationhubstrategyrecommendations.ListServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListServersOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListServers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MigrationHubStrategyRecommendations) ListServersPages(input *migrationhubstrategyrecommendations.ListServersInput, fn func(*migrationhubstrategyrecommendations.ListServersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListServersOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListServersOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListServersOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListServersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListServersPagesWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListServersInput, fn func(*migrationhubstrategyrecommendations.ListServersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*migrationhubstrategyrecommendations.ListServersOutput), false)
		return nil
	}
	cachable := true
	output := &migrationhubstrategyrecommendations.ListServersOutput{}
	fnCacher := func(out *migrationhubstrategyrecommendations.ListServersOutput, more bool) bool {
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
	if err := c.MigrationHubStrategyRecommendationsAPI.ListServersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MigrationHubStrategyRecommendations) ListServersWithContext(ctx context.Context, input *migrationhubstrategyrecommendations.ListServersInput, opts ...request.Option) (*migrationhubstrategyrecommendations.ListServersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*migrationhubstrategyrecommendations.ListServersOutput), nil
	}
	out, err := c.MigrationHubStrategyRecommendationsAPI.ListServersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
