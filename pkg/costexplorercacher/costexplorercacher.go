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
package costexplorercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CostExplorer struct {
	costexploreriface.CostExplorerAPI
	cache *cache.Cache
}

func New(costexplorerapi costexploreriface.CostExplorerAPI) *CostExplorer {
	return &CostExplorer{
		CostExplorerAPI: costexplorerapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CostExplorer) DescribeCostCategoryDefinition(input *costexplorer.DescribeCostCategoryDefinitionInput) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.DescribeCostCategoryDefinitionOutput), nil
	}
	out, err := c.CostExplorerAPI.DescribeCostCategoryDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) DescribeCostCategoryDefinitionWithContext(ctx context.Context, input *costexplorer.DescribeCostCategoryDefinitionInput, opts ...request.Option) (*costexplorer.DescribeCostCategoryDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.DescribeCostCategoryDefinitionOutput), nil
	}
	out, err := c.CostExplorerAPI.DescribeCostCategoryDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomalies(input *costexplorer.GetAnomaliesInput) (*costexplorer.GetAnomaliesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomaliesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomalies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomaliesWithContext(ctx context.Context, input *costexplorer.GetAnomaliesInput, opts ...request.Option) (*costexplorer.GetAnomaliesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomaliesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomaliesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomalyMonitors(input *costexplorer.GetAnomalyMonitorsInput) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomalyMonitorsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomalyMonitors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomalyMonitorsWithContext(ctx context.Context, input *costexplorer.GetAnomalyMonitorsInput, opts ...request.Option) (*costexplorer.GetAnomalyMonitorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomalyMonitorsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomalyMonitorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomalySubscriptions(input *costexplorer.GetAnomalySubscriptionsInput) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomalySubscriptionsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomalySubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetAnomalySubscriptionsWithContext(ctx context.Context, input *costexplorer.GetAnomalySubscriptionsInput, opts ...request.Option) (*costexplorer.GetAnomalySubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetAnomalySubscriptionsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetAnomalySubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostAndUsageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostAndUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostAndUsageWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageInput, opts ...request.Option) (*costexplorer.GetCostAndUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostAndUsageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostAndUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostAndUsageWithResources(input *costexplorer.GetCostAndUsageWithResourcesInput) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostAndUsageWithResourcesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostAndUsageWithResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostAndUsageWithResourcesWithContext(ctx context.Context, input *costexplorer.GetCostAndUsageWithResourcesInput, opts ...request.Option) (*costexplorer.GetCostAndUsageWithResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostAndUsageWithResourcesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostAndUsageWithResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostCategories(input *costexplorer.GetCostCategoriesInput) (*costexplorer.GetCostCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostCategoriesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostCategoriesWithContext(ctx context.Context, input *costexplorer.GetCostCategoriesInput, opts ...request.Option) (*costexplorer.GetCostCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostCategoriesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostForecast(input *costexplorer.GetCostForecastInput) (*costexplorer.GetCostForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostForecastOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetCostForecastWithContext(ctx context.Context, input *costexplorer.GetCostForecastInput, opts ...request.Option) (*costexplorer.GetCostForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetCostForecastOutput), nil
	}
	out, err := c.CostExplorerAPI.GetCostForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetDimensionValues(input *costexplorer.GetDimensionValuesInput) (*costexplorer.GetDimensionValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetDimensionValuesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetDimensionValues(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetDimensionValuesWithContext(ctx context.Context, input *costexplorer.GetDimensionValuesInput, opts ...request.Option) (*costexplorer.GetDimensionValuesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetDimensionValuesOutput), nil
	}
	out, err := c.CostExplorerAPI.GetDimensionValuesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationCoverage(input *costexplorer.GetReservationCoverageInput) (*costexplorer.GetReservationCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationCoverageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationCoverage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationCoverageWithContext(ctx context.Context, input *costexplorer.GetReservationCoverageInput, opts ...request.Option) (*costexplorer.GetReservationCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationCoverageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationCoverageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationPurchaseRecommendation(input *costexplorer.GetReservationPurchaseRecommendationInput) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationPurchaseRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationPurchaseRecommendation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetReservationPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetReservationPurchaseRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationPurchaseRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationPurchaseRecommendationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationUtilization(input *costexplorer.GetReservationUtilizationInput) (*costexplorer.GetReservationUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationUtilizationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationUtilization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetReservationUtilizationWithContext(ctx context.Context, input *costexplorer.GetReservationUtilizationInput, opts ...request.Option) (*costexplorer.GetReservationUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetReservationUtilizationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetReservationUtilizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetRightsizingRecommendation(input *costexplorer.GetRightsizingRecommendationInput) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetRightsizingRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetRightsizingRecommendation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetRightsizingRecommendationWithContext(ctx context.Context, input *costexplorer.GetRightsizingRecommendationInput, opts ...request.Option) (*costexplorer.GetRightsizingRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetRightsizingRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetRightsizingRecommendationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansCoverage(input *costexplorer.GetSavingsPlansCoverageInput) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansCoverageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansCoverage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansCoveragePages(input *costexplorer.GetSavingsPlansCoverageInput, fn func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.GetSavingsPlansCoverageOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.GetSavingsPlansCoverageOutput{}
	fnCacher := func(out *costexplorer.GetSavingsPlansCoverageOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.GetSavingsPlansCoveragePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) GetSavingsPlansCoveragePagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, fn func(*costexplorer.GetSavingsPlansCoverageOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.GetSavingsPlansCoverageOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.GetSavingsPlansCoverageOutput{}
	fnCacher := func(out *costexplorer.GetSavingsPlansCoverageOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.GetSavingsPlansCoveragePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) GetSavingsPlansCoverageWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansCoverageInput, opts ...request.Option) (*costexplorer.GetSavingsPlansCoverageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansCoverageOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansCoverageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansPurchaseRecommendation(input *costexplorer.GetSavingsPlansPurchaseRecommendationInput) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansPurchaseRecommendation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansPurchaseRecommendationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansPurchaseRecommendationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansPurchaseRecommendationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansPurchaseRecommendationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansPurchaseRecommendationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansUtilization(input *costexplorer.GetSavingsPlansUtilizationInput) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansUtilizationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansUtilization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansUtilizationDetails(input *costexplorer.GetSavingsPlansUtilizationDetailsInput) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansUtilizationDetailsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansUtilizationDetails(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansUtilizationDetailsPages(input *costexplorer.GetSavingsPlansUtilizationDetailsInput, fn func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.GetSavingsPlansUtilizationDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.GetSavingsPlansUtilizationDetailsOutput{}
	fnCacher := func(out *costexplorer.GetSavingsPlansUtilizationDetailsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.GetSavingsPlansUtilizationDetailsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) GetSavingsPlansUtilizationDetailsPagesWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, fn func(*costexplorer.GetSavingsPlansUtilizationDetailsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.GetSavingsPlansUtilizationDetailsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.GetSavingsPlansUtilizationDetailsOutput{}
	fnCacher := func(out *costexplorer.GetSavingsPlansUtilizationDetailsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.GetSavingsPlansUtilizationDetailsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) GetSavingsPlansUtilizationDetailsWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationDetailsInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationDetailsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansUtilizationDetailsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansUtilizationDetailsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetSavingsPlansUtilizationWithContext(ctx context.Context, input *costexplorer.GetSavingsPlansUtilizationInput, opts ...request.Option) (*costexplorer.GetSavingsPlansUtilizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetSavingsPlansUtilizationOutput), nil
	}
	out, err := c.CostExplorerAPI.GetSavingsPlansUtilizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetTags(input *costexplorer.GetTagsInput) (*costexplorer.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetTagsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetTagsWithContext(ctx context.Context, input *costexplorer.GetTagsInput, opts ...request.Option) (*costexplorer.GetTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetTagsOutput), nil
	}
	out, err := c.CostExplorerAPI.GetTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetUsageForecast(input *costexplorer.GetUsageForecastInput) (*costexplorer.GetUsageForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetUsageForecastOutput), nil
	}
	out, err := c.CostExplorerAPI.GetUsageForecast(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) GetUsageForecastWithContext(ctx context.Context, input *costexplorer.GetUsageForecastInput, opts ...request.Option) (*costexplorer.GetUsageForecastOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.GetUsageForecastOutput), nil
	}
	out, err := c.CostExplorerAPI.GetUsageForecastWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListCostAllocationTags(input *costexplorer.ListCostAllocationTagsInput) (*costexplorer.ListCostAllocationTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListCostAllocationTagsOutput), nil
	}
	out, err := c.CostExplorerAPI.ListCostAllocationTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListCostAllocationTagsPages(input *costexplorer.ListCostAllocationTagsInput, fn func(*costexplorer.ListCostAllocationTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.ListCostAllocationTagsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.ListCostAllocationTagsOutput{}
	fnCacher := func(out *costexplorer.ListCostAllocationTagsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.ListCostAllocationTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) ListCostAllocationTagsPagesWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, fn func(*costexplorer.ListCostAllocationTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.ListCostAllocationTagsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.ListCostAllocationTagsOutput{}
	fnCacher := func(out *costexplorer.ListCostAllocationTagsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.ListCostAllocationTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) ListCostAllocationTagsWithContext(ctx context.Context, input *costexplorer.ListCostAllocationTagsInput, opts ...request.Option) (*costexplorer.ListCostAllocationTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListCostAllocationTagsOutput), nil
	}
	out, err := c.CostExplorerAPI.ListCostAllocationTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListCostCategoryDefinitions(input *costexplorer.ListCostCategoryDefinitionsInput) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListCostCategoryDefinitionsOutput), nil
	}
	out, err := c.CostExplorerAPI.ListCostCategoryDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListCostCategoryDefinitionsPages(input *costexplorer.ListCostCategoryDefinitionsInput, fn func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.ListCostCategoryDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.ListCostCategoryDefinitionsOutput{}
	fnCacher := func(out *costexplorer.ListCostCategoryDefinitionsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.ListCostCategoryDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) ListCostCategoryDefinitionsPagesWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, fn func(*costexplorer.ListCostCategoryDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*costexplorer.ListCostCategoryDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &costexplorer.ListCostCategoryDefinitionsOutput{}
	fnCacher := func(out *costexplorer.ListCostCategoryDefinitionsOutput, more bool) bool {
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
	if err := c.CostExplorerAPI.ListCostCategoryDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CostExplorer) ListCostCategoryDefinitionsWithContext(ctx context.Context, input *costexplorer.ListCostCategoryDefinitionsInput, opts ...request.Option) (*costexplorer.ListCostCategoryDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListCostCategoryDefinitionsOutput), nil
	}
	out, err := c.CostExplorerAPI.ListCostCategoryDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListSavingsPlansPurchaseRecommendationGeneration(input *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput), nil
	}
	out, err := c.CostExplorerAPI.ListSavingsPlansPurchaseRecommendationGeneration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListSavingsPlansPurchaseRecommendationGenerationWithContext(ctx context.Context, input *costexplorer.ListSavingsPlansPurchaseRecommendationGenerationInput, opts ...request.Option) (*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListSavingsPlansPurchaseRecommendationGenerationOutput), nil
	}
	out, err := c.CostExplorerAPI.ListSavingsPlansPurchaseRecommendationGenerationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListTagsForResource(input *costexplorer.ListTagsForResourceInput) (*costexplorer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListTagsForResourceOutput), nil
	}
	out, err := c.CostExplorerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CostExplorer) ListTagsForResourceWithContext(ctx context.Context, input *costexplorer.ListTagsForResourceInput, opts ...request.Option) (*costexplorer.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*costexplorer.ListTagsForResourceOutput), nil
	}
	out, err := c.CostExplorerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
