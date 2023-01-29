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
package quicksightcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/aws/aws-sdk-go/service/quicksight/quicksightiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type QuickSight struct {
	quicksightiface.QuickSightAPI
	cache *cache.Cache
}

func New(quicksightapi quicksightiface.QuickSightAPI) *QuickSight {
	return &QuickSight{
		QuickSightAPI: quicksightapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *QuickSight) DescribeAccountCustomization(input *quicksight.DescribeAccountCustomizationInput) (*quicksight.DescribeAccountCustomizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountCustomizationOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountCustomization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAccountCustomizationWithContext(ctx context.Context, input *quicksight.DescribeAccountCustomizationInput, opts ...request.Option) (*quicksight.DescribeAccountCustomizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountCustomizationOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountCustomizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAccountSettings(input *quicksight.DescribeAccountSettingsInput) (*quicksight.DescribeAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountSettingsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAccountSettingsWithContext(ctx context.Context, input *quicksight.DescribeAccountSettingsInput, opts ...request.Option) (*quicksight.DescribeAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountSettingsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAccountSubscription(input *quicksight.DescribeAccountSubscriptionInput) (*quicksight.DescribeAccountSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountSubscriptionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAccountSubscriptionWithContext(ctx context.Context, input *quicksight.DescribeAccountSubscriptionInput, opts ...request.Option) (*quicksight.DescribeAccountSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAccountSubscriptionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAccountSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysis(input *quicksight.DescribeAnalysisInput) (*quicksight.DescribeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysisDefinition(input *quicksight.DescribeAnalysisDefinitionInput) (*quicksight.DescribeAnalysisDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysisDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysisDefinitionWithContext(ctx context.Context, input *quicksight.DescribeAnalysisDefinitionInput, opts ...request.Option) (*quicksight.DescribeAnalysisDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysisDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysisPermissions(input *quicksight.DescribeAnalysisPermissionsInput) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysisPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysisPermissionsWithContext(ctx context.Context, input *quicksight.DescribeAnalysisPermissionsInput, opts ...request.Option) (*quicksight.DescribeAnalysisPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysisPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeAnalysisWithContext(ctx context.Context, input *quicksight.DescribeAnalysisInput, opts ...request.Option) (*quicksight.DescribeAnalysisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeAnalysisOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeAnalysisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboard(input *quicksight.DescribeDashboardInput) (*quicksight.DescribeDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboard(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboardDefinition(input *quicksight.DescribeDashboardDefinitionInput) (*quicksight.DescribeDashboardDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboardDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboardDefinitionWithContext(ctx context.Context, input *quicksight.DescribeDashboardDefinitionInput, opts ...request.Option) (*quicksight.DescribeDashboardDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboardDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboardPermissions(input *quicksight.DescribeDashboardPermissionsInput) (*quicksight.DescribeDashboardPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboardPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboardPermissionsWithContext(ctx context.Context, input *quicksight.DescribeDashboardPermissionsInput, opts ...request.Option) (*quicksight.DescribeDashboardPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboardPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDashboardWithContext(ctx context.Context, input *quicksight.DescribeDashboardInput, opts ...request.Option) (*quicksight.DescribeDashboardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDashboardOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDashboardWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSet(input *quicksight.DescribeDataSetInput) (*quicksight.DescribeDataSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSetOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSetPermissions(input *quicksight.DescribeDataSetPermissionsInput) (*quicksight.DescribeDataSetPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSetPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSetPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSetPermissionsWithContext(ctx context.Context, input *quicksight.DescribeDataSetPermissionsInput, opts ...request.Option) (*quicksight.DescribeDataSetPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSetPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSetPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSetWithContext(ctx context.Context, input *quicksight.DescribeDataSetInput, opts ...request.Option) (*quicksight.DescribeDataSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSetOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSource(input *quicksight.DescribeDataSourceInput) (*quicksight.DescribeDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSourceOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSourcePermissions(input *quicksight.DescribeDataSourcePermissionsInput) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSourcePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSourcePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSourcePermissionsWithContext(ctx context.Context, input *quicksight.DescribeDataSourcePermissionsInput, opts ...request.Option) (*quicksight.DescribeDataSourcePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSourcePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSourcePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeDataSourceWithContext(ctx context.Context, input *quicksight.DescribeDataSourceInput, opts ...request.Option) (*quicksight.DescribeDataSourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeDataSourceOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeDataSourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolder(input *quicksight.DescribeFolderInput) (*quicksight.DescribeFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolderPermissions(input *quicksight.DescribeFolderPermissionsInput) (*quicksight.DescribeFolderPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolderPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolderPermissionsWithContext(ctx context.Context, input *quicksight.DescribeFolderPermissionsInput, opts ...request.Option) (*quicksight.DescribeFolderPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolderPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolderResolvedPermissions(input *quicksight.DescribeFolderResolvedPermissionsInput) (*quicksight.DescribeFolderResolvedPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderResolvedPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolderResolvedPermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolderResolvedPermissionsWithContext(ctx context.Context, input *quicksight.DescribeFolderResolvedPermissionsInput, opts ...request.Option) (*quicksight.DescribeFolderResolvedPermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderResolvedPermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolderResolvedPermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeFolderWithContext(ctx context.Context, input *quicksight.DescribeFolderInput, opts ...request.Option) (*quicksight.DescribeFolderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeFolderOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeFolderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeGroup(input *quicksight.DescribeGroupInput) (*quicksight.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeGroupOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeGroupMembership(input *quicksight.DescribeGroupMembershipInput) (*quicksight.DescribeGroupMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeGroupMembershipOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeGroupMembership(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeGroupMembershipWithContext(ctx context.Context, input *quicksight.DescribeGroupMembershipInput, opts ...request.Option) (*quicksight.DescribeGroupMembershipOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeGroupMembershipOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeGroupMembershipWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeGroupWithContext(ctx context.Context, input *quicksight.DescribeGroupInput, opts ...request.Option) (*quicksight.DescribeGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeGroupOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIAMPolicyAssignment(input *quicksight.DescribeIAMPolicyAssignmentInput) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIAMPolicyAssignmentOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIAMPolicyAssignment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIAMPolicyAssignmentWithContext(ctx context.Context, input *quicksight.DescribeIAMPolicyAssignmentInput, opts ...request.Option) (*quicksight.DescribeIAMPolicyAssignmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIAMPolicyAssignmentOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIAMPolicyAssignmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIngestion(input *quicksight.DescribeIngestionInput) (*quicksight.DescribeIngestionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIngestionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIngestion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIngestionWithContext(ctx context.Context, input *quicksight.DescribeIngestionInput, opts ...request.Option) (*quicksight.DescribeIngestionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIngestionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIngestionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIpRestriction(input *quicksight.DescribeIpRestrictionInput) (*quicksight.DescribeIpRestrictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIpRestrictionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIpRestriction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeIpRestrictionWithContext(ctx context.Context, input *quicksight.DescribeIpRestrictionInput, opts ...request.Option) (*quicksight.DescribeIpRestrictionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeIpRestrictionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeIpRestrictionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeNamespace(input *quicksight.DescribeNamespaceInput) (*quicksight.DescribeNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeNamespaceOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeNamespace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeNamespaceWithContext(ctx context.Context, input *quicksight.DescribeNamespaceInput, opts ...request.Option) (*quicksight.DescribeNamespaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeNamespaceOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeNamespaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplate(input *quicksight.DescribeTemplateInput) (*quicksight.DescribeTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplateAlias(input *quicksight.DescribeTemplateAliasInput) (*quicksight.DescribeTemplateAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateAliasOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplateAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplateAliasWithContext(ctx context.Context, input *quicksight.DescribeTemplateAliasInput, opts ...request.Option) (*quicksight.DescribeTemplateAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateAliasOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplateAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplateDefinition(input *quicksight.DescribeTemplateDefinitionInput) (*quicksight.DescribeTemplateDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplateDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplateDefinitionWithContext(ctx context.Context, input *quicksight.DescribeTemplateDefinitionInput, opts ...request.Option) (*quicksight.DescribeTemplateDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateDefinitionOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplateDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplatePermissions(input *quicksight.DescribeTemplatePermissionsInput) (*quicksight.DescribeTemplatePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplatePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplatePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplatePermissionsWithContext(ctx context.Context, input *quicksight.DescribeTemplatePermissionsInput, opts ...request.Option) (*quicksight.DescribeTemplatePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplatePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplatePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTemplateWithContext(ctx context.Context, input *quicksight.DescribeTemplateInput, opts ...request.Option) (*quicksight.DescribeTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeTemplateOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeTheme(input *quicksight.DescribeThemeInput) (*quicksight.DescribeThemeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemeOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeTheme(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeThemeAlias(input *quicksight.DescribeThemeAliasInput) (*quicksight.DescribeThemeAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemeAliasOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeThemeAlias(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeThemeAliasWithContext(ctx context.Context, input *quicksight.DescribeThemeAliasInput, opts ...request.Option) (*quicksight.DescribeThemeAliasOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemeAliasOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeThemeAliasWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeThemePermissions(input *quicksight.DescribeThemePermissionsInput) (*quicksight.DescribeThemePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeThemePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeThemePermissionsWithContext(ctx context.Context, input *quicksight.DescribeThemePermissionsInput, opts ...request.Option) (*quicksight.DescribeThemePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemePermissionsOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeThemePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeThemeWithContext(ctx context.Context, input *quicksight.DescribeThemeInput, opts ...request.Option) (*quicksight.DescribeThemeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeThemeOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeThemeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeUser(input *quicksight.DescribeUserInput) (*quicksight.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeUserOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) DescribeUserWithContext(ctx context.Context, input *quicksight.DescribeUserInput, opts ...request.Option) (*quicksight.DescribeUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.DescribeUserOutput), nil
	}
	out, err := c.QuickSightAPI.DescribeUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) GetDashboardEmbedUrl(input *quicksight.GetDashboardEmbedUrlInput) (*quicksight.GetDashboardEmbedUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.GetDashboardEmbedUrlOutput), nil
	}
	out, err := c.QuickSightAPI.GetDashboardEmbedUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) GetDashboardEmbedUrlWithContext(ctx context.Context, input *quicksight.GetDashboardEmbedUrlInput, opts ...request.Option) (*quicksight.GetDashboardEmbedUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.GetDashboardEmbedUrlOutput), nil
	}
	out, err := c.QuickSightAPI.GetDashboardEmbedUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) GetSessionEmbedUrl(input *quicksight.GetSessionEmbedUrlInput) (*quicksight.GetSessionEmbedUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.GetSessionEmbedUrlOutput), nil
	}
	out, err := c.QuickSightAPI.GetSessionEmbedUrl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) GetSessionEmbedUrlWithContext(ctx context.Context, input *quicksight.GetSessionEmbedUrlInput, opts ...request.Option) (*quicksight.GetSessionEmbedUrlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.GetSessionEmbedUrlOutput), nil
	}
	out, err := c.QuickSightAPI.GetSessionEmbedUrlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListAnalyses(input *quicksight.ListAnalysesInput) (*quicksight.ListAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListAnalysesOutput), nil
	}
	out, err := c.QuickSightAPI.ListAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListAnalysesPages(input *quicksight.ListAnalysesInput, fn func(*quicksight.ListAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListAnalysesOutput{}
	fnCacher := func(out *quicksight.ListAnalysesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListAnalysesPagesWithContext(ctx context.Context, input *quicksight.ListAnalysesInput, fn func(*quicksight.ListAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListAnalysesOutput{}
	fnCacher := func(out *quicksight.ListAnalysesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListAnalysesWithContext(ctx context.Context, input *quicksight.ListAnalysesInput, opts ...request.Option) (*quicksight.ListAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListAnalysesOutput), nil
	}
	out, err := c.QuickSightAPI.ListAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDashboardVersions(input *quicksight.ListDashboardVersionsInput) (*quicksight.ListDashboardVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDashboardVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDashboardVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDashboardVersionsPages(input *quicksight.ListDashboardVersionsInput, fn func(*quicksight.ListDashboardVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDashboardVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDashboardVersionsOutput{}
	fnCacher := func(out *quicksight.ListDashboardVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDashboardVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDashboardVersionsPagesWithContext(ctx context.Context, input *quicksight.ListDashboardVersionsInput, fn func(*quicksight.ListDashboardVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDashboardVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDashboardVersionsOutput{}
	fnCacher := func(out *quicksight.ListDashboardVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDashboardVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDashboardVersionsWithContext(ctx context.Context, input *quicksight.ListDashboardVersionsInput, opts ...request.Option) (*quicksight.ListDashboardVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDashboardVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDashboardVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDashboards(input *quicksight.ListDashboardsInput) (*quicksight.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDashboardsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDashboards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDashboardsPages(input *quicksight.ListDashboardsInput, fn func(*quicksight.ListDashboardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDashboardsOutput{}
	fnCacher := func(out *quicksight.ListDashboardsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDashboardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDashboardsPagesWithContext(ctx context.Context, input *quicksight.ListDashboardsInput, fn func(*quicksight.ListDashboardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDashboardsOutput{}
	fnCacher := func(out *quicksight.ListDashboardsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDashboardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDashboardsWithContext(ctx context.Context, input *quicksight.ListDashboardsInput, opts ...request.Option) (*quicksight.ListDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDashboardsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDashboardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDataSets(input *quicksight.ListDataSetsInput) (*quicksight.ListDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDataSetsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDataSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDataSetsPages(input *quicksight.ListDataSetsInput, fn func(*quicksight.ListDataSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDataSetsOutput{}
	fnCacher := func(out *quicksight.ListDataSetsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDataSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDataSetsPagesWithContext(ctx context.Context, input *quicksight.ListDataSetsInput, fn func(*quicksight.ListDataSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDataSetsOutput{}
	fnCacher := func(out *quicksight.ListDataSetsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDataSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDataSetsWithContext(ctx context.Context, input *quicksight.ListDataSetsInput, opts ...request.Option) (*quicksight.ListDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDataSetsOutput), nil
	}
	out, err := c.QuickSightAPI.ListDataSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDataSources(input *quicksight.ListDataSourcesInput) (*quicksight.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDataSourcesOutput), nil
	}
	out, err := c.QuickSightAPI.ListDataSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListDataSourcesPages(input *quicksight.ListDataSourcesInput, fn func(*quicksight.ListDataSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDataSourcesOutput{}
	fnCacher := func(out *quicksight.ListDataSourcesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDataSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDataSourcesPagesWithContext(ctx context.Context, input *quicksight.ListDataSourcesInput, fn func(*quicksight.ListDataSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListDataSourcesOutput{}
	fnCacher := func(out *quicksight.ListDataSourcesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListDataSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListDataSourcesWithContext(ctx context.Context, input *quicksight.ListDataSourcesInput, opts ...request.Option) (*quicksight.ListDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListDataSourcesOutput), nil
	}
	out, err := c.QuickSightAPI.ListDataSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListFolderMembers(input *quicksight.ListFolderMembersInput) (*quicksight.ListFolderMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListFolderMembersOutput), nil
	}
	out, err := c.QuickSightAPI.ListFolderMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListFolderMembersWithContext(ctx context.Context, input *quicksight.ListFolderMembersInput, opts ...request.Option) (*quicksight.ListFolderMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListFolderMembersOutput), nil
	}
	out, err := c.QuickSightAPI.ListFolderMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListFolders(input *quicksight.ListFoldersInput) (*quicksight.ListFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListFoldersOutput), nil
	}
	out, err := c.QuickSightAPI.ListFolders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListFoldersWithContext(ctx context.Context, input *quicksight.ListFoldersInput, opts ...request.Option) (*quicksight.ListFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListFoldersOutput), nil
	}
	out, err := c.QuickSightAPI.ListFoldersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListGroupMemberships(input *quicksight.ListGroupMembershipsInput) (*quicksight.ListGroupMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListGroupMembershipsOutput), nil
	}
	out, err := c.QuickSightAPI.ListGroupMemberships(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListGroupMembershipsWithContext(ctx context.Context, input *quicksight.ListGroupMembershipsInput, opts ...request.Option) (*quicksight.ListGroupMembershipsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListGroupMembershipsOutput), nil
	}
	out, err := c.QuickSightAPI.ListGroupMembershipsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListGroups(input *quicksight.ListGroupsInput) (*quicksight.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.ListGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListGroupsWithContext(ctx context.Context, input *quicksight.ListGroupsInput, opts ...request.Option) (*quicksight.ListGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.ListGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIAMPolicyAssignments(input *quicksight.ListIAMPolicyAssignmentsInput) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIAMPolicyAssignmentsOutput), nil
	}
	out, err := c.QuickSightAPI.ListIAMPolicyAssignments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIAMPolicyAssignmentsForUser(input *quicksight.ListIAMPolicyAssignmentsForUserInput) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIAMPolicyAssignmentsForUserOutput), nil
	}
	out, err := c.QuickSightAPI.ListIAMPolicyAssignmentsForUser(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIAMPolicyAssignmentsForUserWithContext(ctx context.Context, input *quicksight.ListIAMPolicyAssignmentsForUserInput, opts ...request.Option) (*quicksight.ListIAMPolicyAssignmentsForUserOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIAMPolicyAssignmentsForUserOutput), nil
	}
	out, err := c.QuickSightAPI.ListIAMPolicyAssignmentsForUserWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIAMPolicyAssignmentsWithContext(ctx context.Context, input *quicksight.ListIAMPolicyAssignmentsInput, opts ...request.Option) (*quicksight.ListIAMPolicyAssignmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIAMPolicyAssignmentsOutput), nil
	}
	out, err := c.QuickSightAPI.ListIAMPolicyAssignmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIngestions(input *quicksight.ListIngestionsInput) (*quicksight.ListIngestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIngestionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListIngestions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListIngestionsPages(input *quicksight.ListIngestionsInput, fn func(*quicksight.ListIngestionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListIngestionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListIngestionsOutput{}
	fnCacher := func(out *quicksight.ListIngestionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListIngestionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListIngestionsPagesWithContext(ctx context.Context, input *quicksight.ListIngestionsInput, fn func(*quicksight.ListIngestionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListIngestionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListIngestionsOutput{}
	fnCacher := func(out *quicksight.ListIngestionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListIngestionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListIngestionsWithContext(ctx context.Context, input *quicksight.ListIngestionsInput, opts ...request.Option) (*quicksight.ListIngestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListIngestionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListIngestionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListNamespaces(input *quicksight.ListNamespacesInput) (*quicksight.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListNamespacesOutput), nil
	}
	out, err := c.QuickSightAPI.ListNamespaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListNamespacesPages(input *quicksight.ListNamespacesInput, fn func(*quicksight.ListNamespacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListNamespacesOutput{}
	fnCacher := func(out *quicksight.ListNamespacesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListNamespacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListNamespacesPagesWithContext(ctx context.Context, input *quicksight.ListNamespacesInput, fn func(*quicksight.ListNamespacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListNamespacesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListNamespacesOutput{}
	fnCacher := func(out *quicksight.ListNamespacesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListNamespacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListNamespacesWithContext(ctx context.Context, input *quicksight.ListNamespacesInput, opts ...request.Option) (*quicksight.ListNamespacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListNamespacesOutput), nil
	}
	out, err := c.QuickSightAPI.ListNamespacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTagsForResource(input *quicksight.ListTagsForResourceInput) (*quicksight.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTagsForResourceOutput), nil
	}
	out, err := c.QuickSightAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTagsForResourceWithContext(ctx context.Context, input *quicksight.ListTagsForResourceInput, opts ...request.Option) (*quicksight.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTagsForResourceOutput), nil
	}
	out, err := c.QuickSightAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplateAliases(input *quicksight.ListTemplateAliasesInput) (*quicksight.ListTemplateAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplateAliasesOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplateAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplateAliasesPages(input *quicksight.ListTemplateAliasesInput, fn func(*quicksight.ListTemplateAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplateAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplateAliasesOutput{}
	fnCacher := func(out *quicksight.ListTemplateAliasesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplateAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplateAliasesPagesWithContext(ctx context.Context, input *quicksight.ListTemplateAliasesInput, fn func(*quicksight.ListTemplateAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplateAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplateAliasesOutput{}
	fnCacher := func(out *quicksight.ListTemplateAliasesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplateAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplateAliasesWithContext(ctx context.Context, input *quicksight.ListTemplateAliasesInput, opts ...request.Option) (*quicksight.ListTemplateAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplateAliasesOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplateAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplateVersions(input *quicksight.ListTemplateVersionsInput) (*quicksight.ListTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplateVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplateVersionsPages(input *quicksight.ListTemplateVersionsInput, fn func(*quicksight.ListTemplateVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplateVersionsOutput{}
	fnCacher := func(out *quicksight.ListTemplateVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplateVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplateVersionsPagesWithContext(ctx context.Context, input *quicksight.ListTemplateVersionsInput, fn func(*quicksight.ListTemplateVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplateVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplateVersionsOutput{}
	fnCacher := func(out *quicksight.ListTemplateVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplateVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplateVersionsWithContext(ctx context.Context, input *quicksight.ListTemplateVersionsInput, opts ...request.Option) (*quicksight.ListTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplateVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplates(input *quicksight.ListTemplatesInput) (*quicksight.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplatesOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListTemplatesPages(input *quicksight.ListTemplatesInput, fn func(*quicksight.ListTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplatesOutput{}
	fnCacher := func(out *quicksight.ListTemplatesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplatesPagesWithContext(ctx context.Context, input *quicksight.ListTemplatesInput, fn func(*quicksight.ListTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListTemplatesOutput{}
	fnCacher := func(out *quicksight.ListTemplatesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListTemplatesWithContext(ctx context.Context, input *quicksight.ListTemplatesInput, opts ...request.Option) (*quicksight.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListTemplatesOutput), nil
	}
	out, err := c.QuickSightAPI.ListTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemeAliases(input *quicksight.ListThemeAliasesInput) (*quicksight.ListThemeAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemeAliasesOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemeAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemeAliasesWithContext(ctx context.Context, input *quicksight.ListThemeAliasesInput, opts ...request.Option) (*quicksight.ListThemeAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemeAliasesOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemeAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemeVersions(input *quicksight.ListThemeVersionsInput) (*quicksight.ListThemeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemeVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemeVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemeVersionsPages(input *quicksight.ListThemeVersionsInput, fn func(*quicksight.ListThemeVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListThemeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListThemeVersionsOutput{}
	fnCacher := func(out *quicksight.ListThemeVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListThemeVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListThemeVersionsPagesWithContext(ctx context.Context, input *quicksight.ListThemeVersionsInput, fn func(*quicksight.ListThemeVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListThemeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListThemeVersionsOutput{}
	fnCacher := func(out *quicksight.ListThemeVersionsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListThemeVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListThemeVersionsWithContext(ctx context.Context, input *quicksight.ListThemeVersionsInput, opts ...request.Option) (*quicksight.ListThemeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemeVersionsOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemeVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemes(input *quicksight.ListThemesInput) (*quicksight.ListThemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemesOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListThemesPages(input *quicksight.ListThemesInput, fn func(*quicksight.ListThemesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListThemesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListThemesOutput{}
	fnCacher := func(out *quicksight.ListThemesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListThemesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListThemesPagesWithContext(ctx context.Context, input *quicksight.ListThemesInput, fn func(*quicksight.ListThemesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.ListThemesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.ListThemesOutput{}
	fnCacher := func(out *quicksight.ListThemesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.ListThemesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) ListThemesWithContext(ctx context.Context, input *quicksight.ListThemesInput, opts ...request.Option) (*quicksight.ListThemesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListThemesOutput), nil
	}
	out, err := c.QuickSightAPI.ListThemesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListUserGroups(input *quicksight.ListUserGroupsInput) (*quicksight.ListUserGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListUserGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.ListUserGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListUserGroupsWithContext(ctx context.Context, input *quicksight.ListUserGroupsInput, opts ...request.Option) (*quicksight.ListUserGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListUserGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.ListUserGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListUsers(input *quicksight.ListUsersInput) (*quicksight.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListUsersOutput), nil
	}
	out, err := c.QuickSightAPI.ListUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) ListUsersWithContext(ctx context.Context, input *quicksight.ListUsersInput, opts ...request.Option) (*quicksight.ListUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.ListUsersOutput), nil
	}
	out, err := c.QuickSightAPI.ListUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchAnalyses(input *quicksight.SearchAnalysesInput) (*quicksight.SearchAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchAnalysesOutput), nil
	}
	out, err := c.QuickSightAPI.SearchAnalyses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchAnalysesPages(input *quicksight.SearchAnalysesInput, fn func(*quicksight.SearchAnalysesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchAnalysesOutput{}
	fnCacher := func(out *quicksight.SearchAnalysesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchAnalysesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchAnalysesPagesWithContext(ctx context.Context, input *quicksight.SearchAnalysesInput, fn func(*quicksight.SearchAnalysesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchAnalysesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchAnalysesOutput{}
	fnCacher := func(out *quicksight.SearchAnalysesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchAnalysesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchAnalysesWithContext(ctx context.Context, input *quicksight.SearchAnalysesInput, opts ...request.Option) (*quicksight.SearchAnalysesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchAnalysesOutput), nil
	}
	out, err := c.QuickSightAPI.SearchAnalysesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDashboards(input *quicksight.SearchDashboardsInput) (*quicksight.SearchDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDashboardsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDashboards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDashboardsPages(input *quicksight.SearchDashboardsInput, fn func(*quicksight.SearchDashboardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDashboardsOutput{}
	fnCacher := func(out *quicksight.SearchDashboardsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDashboardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDashboardsPagesWithContext(ctx context.Context, input *quicksight.SearchDashboardsInput, fn func(*quicksight.SearchDashboardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDashboardsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDashboardsOutput{}
	fnCacher := func(out *quicksight.SearchDashboardsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDashboardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDashboardsWithContext(ctx context.Context, input *quicksight.SearchDashboardsInput, opts ...request.Option) (*quicksight.SearchDashboardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDashboardsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDashboardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDataSets(input *quicksight.SearchDataSetsInput) (*quicksight.SearchDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDataSetsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDataSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDataSetsPages(input *quicksight.SearchDataSetsInput, fn func(*quicksight.SearchDataSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDataSetsOutput{}
	fnCacher := func(out *quicksight.SearchDataSetsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDataSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDataSetsPagesWithContext(ctx context.Context, input *quicksight.SearchDataSetsInput, fn func(*quicksight.SearchDataSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDataSetsOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDataSetsOutput{}
	fnCacher := func(out *quicksight.SearchDataSetsOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDataSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDataSetsWithContext(ctx context.Context, input *quicksight.SearchDataSetsInput, opts ...request.Option) (*quicksight.SearchDataSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDataSetsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDataSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDataSources(input *quicksight.SearchDataSourcesInput) (*quicksight.SearchDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDataSourcesOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDataSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchDataSourcesPages(input *quicksight.SearchDataSourcesInput, fn func(*quicksight.SearchDataSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDataSourcesOutput{}
	fnCacher := func(out *quicksight.SearchDataSourcesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDataSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDataSourcesPagesWithContext(ctx context.Context, input *quicksight.SearchDataSourcesInput, fn func(*quicksight.SearchDataSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*quicksight.SearchDataSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &quicksight.SearchDataSourcesOutput{}
	fnCacher := func(out *quicksight.SearchDataSourcesOutput, more bool) bool {
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
	if err := c.QuickSightAPI.SearchDataSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *QuickSight) SearchDataSourcesWithContext(ctx context.Context, input *quicksight.SearchDataSourcesInput, opts ...request.Option) (*quicksight.SearchDataSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchDataSourcesOutput), nil
	}
	out, err := c.QuickSightAPI.SearchDataSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchFolders(input *quicksight.SearchFoldersInput) (*quicksight.SearchFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchFoldersOutput), nil
	}
	out, err := c.QuickSightAPI.SearchFolders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchFoldersWithContext(ctx context.Context, input *quicksight.SearchFoldersInput, opts ...request.Option) (*quicksight.SearchFoldersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchFoldersOutput), nil
	}
	out, err := c.QuickSightAPI.SearchFoldersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchGroups(input *quicksight.SearchGroupsInput) (*quicksight.SearchGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *QuickSight) SearchGroupsWithContext(ctx context.Context, input *quicksight.SearchGroupsInput, opts ...request.Option) (*quicksight.SearchGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*quicksight.SearchGroupsOutput), nil
	}
	out, err := c.QuickSightAPI.SearchGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
