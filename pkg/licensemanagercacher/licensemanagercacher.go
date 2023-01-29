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
package licensemanagercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanager"
	"github.com/aws/aws-sdk-go/service/licensemanager/licensemanageriface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LicenseManager struct {
	licensemanageriface.LicenseManagerAPI
	cache *cache.Cache
}

func New(licensemanagerapi licensemanageriface.LicenseManagerAPI) *LicenseManager {
	return &LicenseManager{
		LicenseManagerAPI: licensemanagerapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LicenseManager) GetAccessToken(input *licensemanager.GetAccessTokenInput) (*licensemanager.GetAccessTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetAccessTokenOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetAccessToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetAccessTokenWithContext(ctx context.Context, input *licensemanager.GetAccessTokenInput, opts ...request.Option) (*licensemanager.GetAccessTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetAccessTokenOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetAccessTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetGrant(input *licensemanager.GetGrantInput) (*licensemanager.GetGrantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetGrantOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetGrant(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetGrantWithContext(ctx context.Context, input *licensemanager.GetGrantInput, opts ...request.Option) (*licensemanager.GetGrantOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetGrantOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetGrantWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicense(input *licensemanager.GetLicenseInput) (*licensemanager.GetLicenseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicense(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseConfiguration(input *licensemanager.GetLicenseConfigurationInput) (*licensemanager.GetLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseConfigurationWithContext(ctx context.Context, input *licensemanager.GetLicenseConfigurationInput, opts ...request.Option) (*licensemanager.GetLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseConversionTask(input *licensemanager.GetLicenseConversionTaskInput) (*licensemanager.GetLicenseConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseConversionTaskOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseConversionTask(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseConversionTaskWithContext(ctx context.Context, input *licensemanager.GetLicenseConversionTaskInput, opts ...request.Option) (*licensemanager.GetLicenseConversionTaskOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseConversionTaskOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseConversionTaskWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseManagerReportGenerator(input *licensemanager.GetLicenseManagerReportGeneratorInput) (*licensemanager.GetLicenseManagerReportGeneratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseManagerReportGeneratorOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseManagerReportGenerator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseManagerReportGeneratorWithContext(ctx context.Context, input *licensemanager.GetLicenseManagerReportGeneratorInput, opts ...request.Option) (*licensemanager.GetLicenseManagerReportGeneratorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseManagerReportGeneratorOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseManagerReportGeneratorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseUsage(input *licensemanager.GetLicenseUsageInput) (*licensemanager.GetLicenseUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseUsageOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseUsage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseUsageWithContext(ctx context.Context, input *licensemanager.GetLicenseUsageInput, opts ...request.Option) (*licensemanager.GetLicenseUsageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseUsageOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseUsageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetLicenseWithContext(ctx context.Context, input *licensemanager.GetLicenseInput, opts ...request.Option) (*licensemanager.GetLicenseOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetLicenseOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetLicenseWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetServiceSettings(input *licensemanager.GetServiceSettingsInput) (*licensemanager.GetServiceSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetServiceSettingsOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetServiceSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) GetServiceSettingsWithContext(ctx context.Context, input *licensemanager.GetServiceSettingsInput, opts ...request.Option) (*licensemanager.GetServiceSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.GetServiceSettingsOutput), nil
	}
	out, err := c.LicenseManagerAPI.GetServiceSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListAssociationsForLicenseConfiguration(input *licensemanager.ListAssociationsForLicenseConfigurationInput) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListAssociationsForLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListAssociationsForLicenseConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListAssociationsForLicenseConfigurationWithContext(ctx context.Context, input *licensemanager.ListAssociationsForLicenseConfigurationInput, opts ...request.Option) (*licensemanager.ListAssociationsForLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListAssociationsForLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListAssociationsForLicenseConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListDistributedGrants(input *licensemanager.ListDistributedGrantsInput) (*licensemanager.ListDistributedGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListDistributedGrantsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListDistributedGrants(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListDistributedGrantsWithContext(ctx context.Context, input *licensemanager.ListDistributedGrantsInput, opts ...request.Option) (*licensemanager.ListDistributedGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListDistributedGrantsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListDistributedGrantsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListFailuresForLicenseConfigurationOperations(input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListFailuresForLicenseConfigurationOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListFailuresForLicenseConfigurationOperationsWithContext(ctx context.Context, input *licensemanager.ListFailuresForLicenseConfigurationOperationsInput, opts ...request.Option) (*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListFailuresForLicenseConfigurationOperationsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListFailuresForLicenseConfigurationOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseConfigurations(input *licensemanager.ListLicenseConfigurationsInput) (*licensemanager.ListLicenseConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseConfigurationsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseConfigurationsWithContext(ctx context.Context, input *licensemanager.ListLicenseConfigurationsInput, opts ...request.Option) (*licensemanager.ListLicenseConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseConfigurationsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseConversionTasks(input *licensemanager.ListLicenseConversionTasksInput) (*licensemanager.ListLicenseConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseConversionTasksOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseConversionTasks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseConversionTasksWithContext(ctx context.Context, input *licensemanager.ListLicenseConversionTasksInput, opts ...request.Option) (*licensemanager.ListLicenseConversionTasksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseConversionTasksOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseConversionTasksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseManagerReportGenerators(input *licensemanager.ListLicenseManagerReportGeneratorsInput) (*licensemanager.ListLicenseManagerReportGeneratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseManagerReportGeneratorsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseManagerReportGenerators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseManagerReportGeneratorsWithContext(ctx context.Context, input *licensemanager.ListLicenseManagerReportGeneratorsInput, opts ...request.Option) (*licensemanager.ListLicenseManagerReportGeneratorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseManagerReportGeneratorsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseManagerReportGeneratorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseSpecificationsForResource(input *licensemanager.ListLicenseSpecificationsForResourceInput) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseSpecificationsForResourceOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseSpecificationsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseSpecificationsForResourceWithContext(ctx context.Context, input *licensemanager.ListLicenseSpecificationsForResourceInput, opts ...request.Option) (*licensemanager.ListLicenseSpecificationsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseSpecificationsForResourceOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseSpecificationsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseVersions(input *licensemanager.ListLicenseVersionsInput) (*licensemanager.ListLicenseVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseVersionsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenseVersionsWithContext(ctx context.Context, input *licensemanager.ListLicenseVersionsInput, opts ...request.Option) (*licensemanager.ListLicenseVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicenseVersionsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenseVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicenses(input *licensemanager.ListLicensesInput) (*licensemanager.ListLicensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicensesOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicenses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListLicensesWithContext(ctx context.Context, input *licensemanager.ListLicensesInput, opts ...request.Option) (*licensemanager.ListLicensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListLicensesOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListLicensesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedGrants(input *licensemanager.ListReceivedGrantsInput) (*licensemanager.ListReceivedGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedGrantsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedGrants(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedGrantsForOrganization(input *licensemanager.ListReceivedGrantsForOrganizationInput) (*licensemanager.ListReceivedGrantsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedGrantsForOrganizationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedGrantsForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedGrantsForOrganizationWithContext(ctx context.Context, input *licensemanager.ListReceivedGrantsForOrganizationInput, opts ...request.Option) (*licensemanager.ListReceivedGrantsForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedGrantsForOrganizationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedGrantsForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedGrantsWithContext(ctx context.Context, input *licensemanager.ListReceivedGrantsInput, opts ...request.Option) (*licensemanager.ListReceivedGrantsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedGrantsOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedGrantsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedLicenses(input *licensemanager.ListReceivedLicensesInput) (*licensemanager.ListReceivedLicensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedLicensesOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedLicenses(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedLicensesForOrganization(input *licensemanager.ListReceivedLicensesForOrganizationInput) (*licensemanager.ListReceivedLicensesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedLicensesForOrganizationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedLicensesForOrganization(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedLicensesForOrganizationWithContext(ctx context.Context, input *licensemanager.ListReceivedLicensesForOrganizationInput, opts ...request.Option) (*licensemanager.ListReceivedLicensesForOrganizationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedLicensesForOrganizationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedLicensesForOrganizationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListReceivedLicensesWithContext(ctx context.Context, input *licensemanager.ListReceivedLicensesInput, opts ...request.Option) (*licensemanager.ListReceivedLicensesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListReceivedLicensesOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListReceivedLicensesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListResourceInventory(input *licensemanager.ListResourceInventoryInput) (*licensemanager.ListResourceInventoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListResourceInventoryOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListResourceInventory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListResourceInventoryWithContext(ctx context.Context, input *licensemanager.ListResourceInventoryInput, opts ...request.Option) (*licensemanager.ListResourceInventoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListResourceInventoryOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListResourceInventoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListTagsForResource(input *licensemanager.ListTagsForResourceInput) (*licensemanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListTagsForResourceOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListTagsForResourceWithContext(ctx context.Context, input *licensemanager.ListTagsForResourceInput, opts ...request.Option) (*licensemanager.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListTagsForResourceOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListTokens(input *licensemanager.ListTokensInput) (*licensemanager.ListTokensOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListTokensOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListTokens(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListTokensWithContext(ctx context.Context, input *licensemanager.ListTokensInput, opts ...request.Option) (*licensemanager.ListTokensOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListTokensOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListTokensWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListUsageForLicenseConfiguration(input *licensemanager.ListUsageForLicenseConfigurationInput) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListUsageForLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListUsageForLicenseConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManager) ListUsageForLicenseConfigurationWithContext(ctx context.Context, input *licensemanager.ListUsageForLicenseConfigurationInput, opts ...request.Option) (*licensemanager.ListUsageForLicenseConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanager.ListUsageForLicenseConfigurationOutput), nil
	}
	out, err := c.LicenseManagerAPI.ListUsageForLicenseConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
