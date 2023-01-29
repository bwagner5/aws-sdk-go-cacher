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
package macie2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/macie2"
	"github.com/aws/aws-sdk-go/service/macie2/macie2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Macie2 struct {
	macie2iface.Macie2API
	cache *cache.Cache
}

func New(macie2api macie2iface.Macie2API) *Macie2 {
	return &Macie2{
		Macie2API: macie2api,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Macie2) BatchGetCustomDataIdentifiers(input *macie2.BatchGetCustomDataIdentifiersInput) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.BatchGetCustomDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.BatchGetCustomDataIdentifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) BatchGetCustomDataIdentifiersWithContext(ctx context.Context, input *macie2.BatchGetCustomDataIdentifiersInput, opts ...request.Option) (*macie2.BatchGetCustomDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.BatchGetCustomDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.BatchGetCustomDataIdentifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeBuckets(input *macie2.DescribeBucketsInput) (*macie2.DescribeBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeBucketsOutput), nil
	}
	out, err := c.Macie2API.DescribeBuckets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeBucketsPages(input *macie2.DescribeBucketsInput, fn func(*macie2.DescribeBucketsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.DescribeBucketsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.DescribeBucketsOutput{}
	fnCacher := func(out *macie2.DescribeBucketsOutput, more bool) bool {
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
	if err := c.Macie2API.DescribeBucketsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) DescribeBucketsPagesWithContext(ctx context.Context, input *macie2.DescribeBucketsInput, fn func(*macie2.DescribeBucketsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.DescribeBucketsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.DescribeBucketsOutput{}
	fnCacher := func(out *macie2.DescribeBucketsOutput, more bool) bool {
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
	if err := c.Macie2API.DescribeBucketsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) DescribeBucketsWithContext(ctx context.Context, input *macie2.DescribeBucketsInput, opts ...request.Option) (*macie2.DescribeBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeBucketsOutput), nil
	}
	out, err := c.Macie2API.DescribeBucketsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeClassificationJob(input *macie2.DescribeClassificationJobInput) (*macie2.DescribeClassificationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeClassificationJobOutput), nil
	}
	out, err := c.Macie2API.DescribeClassificationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeClassificationJobWithContext(ctx context.Context, input *macie2.DescribeClassificationJobInput, opts ...request.Option) (*macie2.DescribeClassificationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeClassificationJobOutput), nil
	}
	out, err := c.Macie2API.DescribeClassificationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeOrganizationConfiguration(input *macie2.DescribeOrganizationConfigurationInput) (*macie2.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.Macie2API.DescribeOrganizationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) DescribeOrganizationConfigurationWithContext(ctx context.Context, input *macie2.DescribeOrganizationConfigurationInput, opts ...request.Option) (*macie2.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.Macie2API.DescribeOrganizationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAdministratorAccount(input *macie2.GetAdministratorAccountInput) (*macie2.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAdministratorAccountOutput), nil
	}
	out, err := c.Macie2API.GetAdministratorAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAdministratorAccountWithContext(ctx context.Context, input *macie2.GetAdministratorAccountInput, opts ...request.Option) (*macie2.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAdministratorAccountOutput), nil
	}
	out, err := c.Macie2API.GetAdministratorAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAllowList(input *macie2.GetAllowListInput) (*macie2.GetAllowListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAllowListOutput), nil
	}
	out, err := c.Macie2API.GetAllowList(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAllowListWithContext(ctx context.Context, input *macie2.GetAllowListInput, opts ...request.Option) (*macie2.GetAllowListOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAllowListOutput), nil
	}
	out, err := c.Macie2API.GetAllowListWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAutomatedDiscoveryConfiguration(input *macie2.GetAutomatedDiscoveryConfigurationInput) (*macie2.GetAutomatedDiscoveryConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAutomatedDiscoveryConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetAutomatedDiscoveryConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetAutomatedDiscoveryConfigurationWithContext(ctx context.Context, input *macie2.GetAutomatedDiscoveryConfigurationInput, opts ...request.Option) (*macie2.GetAutomatedDiscoveryConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetAutomatedDiscoveryConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetAutomatedDiscoveryConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetBucketStatistics(input *macie2.GetBucketStatisticsInput) (*macie2.GetBucketStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetBucketStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetBucketStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetBucketStatisticsWithContext(ctx context.Context, input *macie2.GetBucketStatisticsInput, opts ...request.Option) (*macie2.GetBucketStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetBucketStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetBucketStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetClassificationExportConfiguration(input *macie2.GetClassificationExportConfigurationInput) (*macie2.GetClassificationExportConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetClassificationExportConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetClassificationExportConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetClassificationExportConfigurationWithContext(ctx context.Context, input *macie2.GetClassificationExportConfigurationInput, opts ...request.Option) (*macie2.GetClassificationExportConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetClassificationExportConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetClassificationExportConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetClassificationScope(input *macie2.GetClassificationScopeInput) (*macie2.GetClassificationScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetClassificationScopeOutput), nil
	}
	out, err := c.Macie2API.GetClassificationScope(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetClassificationScopeWithContext(ctx context.Context, input *macie2.GetClassificationScopeInput, opts ...request.Option) (*macie2.GetClassificationScopeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetClassificationScopeOutput), nil
	}
	out, err := c.Macie2API.GetClassificationScopeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetCustomDataIdentifier(input *macie2.GetCustomDataIdentifierInput) (*macie2.GetCustomDataIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetCustomDataIdentifierOutput), nil
	}
	out, err := c.Macie2API.GetCustomDataIdentifier(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetCustomDataIdentifierWithContext(ctx context.Context, input *macie2.GetCustomDataIdentifierInput, opts ...request.Option) (*macie2.GetCustomDataIdentifierOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetCustomDataIdentifierOutput), nil
	}
	out, err := c.Macie2API.GetCustomDataIdentifierWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingStatistics(input *macie2.GetFindingStatisticsInput) (*macie2.GetFindingStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetFindingStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingStatisticsWithContext(ctx context.Context, input *macie2.GetFindingStatisticsInput, opts ...request.Option) (*macie2.GetFindingStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetFindingStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindings(input *macie2.GetFindingsInput) (*macie2.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsOutput), nil
	}
	out, err := c.Macie2API.GetFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingsFilter(input *macie2.GetFindingsFilterInput) (*macie2.GetFindingsFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsFilterOutput), nil
	}
	out, err := c.Macie2API.GetFindingsFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingsFilterWithContext(ctx context.Context, input *macie2.GetFindingsFilterInput, opts ...request.Option) (*macie2.GetFindingsFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsFilterOutput), nil
	}
	out, err := c.Macie2API.GetFindingsFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingsPublicationConfiguration(input *macie2.GetFindingsPublicationConfigurationInput) (*macie2.GetFindingsPublicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsPublicationConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetFindingsPublicationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingsPublicationConfigurationWithContext(ctx context.Context, input *macie2.GetFindingsPublicationConfigurationInput, opts ...request.Option) (*macie2.GetFindingsPublicationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsPublicationConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetFindingsPublicationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetFindingsWithContext(ctx context.Context, input *macie2.GetFindingsInput, opts ...request.Option) (*macie2.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetFindingsOutput), nil
	}
	out, err := c.Macie2API.GetFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetInvitationsCount(input *macie2.GetInvitationsCountInput) (*macie2.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetInvitationsCountOutput), nil
	}
	out, err := c.Macie2API.GetInvitationsCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetInvitationsCountWithContext(ctx context.Context, input *macie2.GetInvitationsCountInput, opts ...request.Option) (*macie2.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetInvitationsCountOutput), nil
	}
	out, err := c.Macie2API.GetInvitationsCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMacieSession(input *macie2.GetMacieSessionInput) (*macie2.GetMacieSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMacieSessionOutput), nil
	}
	out, err := c.Macie2API.GetMacieSession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMacieSessionWithContext(ctx context.Context, input *macie2.GetMacieSessionInput, opts ...request.Option) (*macie2.GetMacieSessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMacieSessionOutput), nil
	}
	out, err := c.Macie2API.GetMacieSessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMasterAccount(input *macie2.GetMasterAccountInput) (*macie2.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMasterAccountOutput), nil
	}
	out, err := c.Macie2API.GetMasterAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMasterAccountWithContext(ctx context.Context, input *macie2.GetMasterAccountInput, opts ...request.Option) (*macie2.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMasterAccountOutput), nil
	}
	out, err := c.Macie2API.GetMasterAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMember(input *macie2.GetMemberInput) (*macie2.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMemberOutput), nil
	}
	out, err := c.Macie2API.GetMember(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetMemberWithContext(ctx context.Context, input *macie2.GetMemberInput, opts ...request.Option) (*macie2.GetMemberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetMemberOutput), nil
	}
	out, err := c.Macie2API.GetMemberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetResourceProfile(input *macie2.GetResourceProfileInput) (*macie2.GetResourceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetResourceProfileOutput), nil
	}
	out, err := c.Macie2API.GetResourceProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetResourceProfileWithContext(ctx context.Context, input *macie2.GetResourceProfileInput, opts ...request.Option) (*macie2.GetResourceProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetResourceProfileOutput), nil
	}
	out, err := c.Macie2API.GetResourceProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetRevealConfiguration(input *macie2.GetRevealConfigurationInput) (*macie2.GetRevealConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetRevealConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetRevealConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetRevealConfigurationWithContext(ctx context.Context, input *macie2.GetRevealConfigurationInput, opts ...request.Option) (*macie2.GetRevealConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetRevealConfigurationOutput), nil
	}
	out, err := c.Macie2API.GetRevealConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitiveDataOccurrences(input *macie2.GetSensitiveDataOccurrencesInput) (*macie2.GetSensitiveDataOccurrencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitiveDataOccurrencesOutput), nil
	}
	out, err := c.Macie2API.GetSensitiveDataOccurrences(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitiveDataOccurrencesAvailability(input *macie2.GetSensitiveDataOccurrencesAvailabilityInput) (*macie2.GetSensitiveDataOccurrencesAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitiveDataOccurrencesAvailabilityOutput), nil
	}
	out, err := c.Macie2API.GetSensitiveDataOccurrencesAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitiveDataOccurrencesAvailabilityWithContext(ctx context.Context, input *macie2.GetSensitiveDataOccurrencesAvailabilityInput, opts ...request.Option) (*macie2.GetSensitiveDataOccurrencesAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitiveDataOccurrencesAvailabilityOutput), nil
	}
	out, err := c.Macie2API.GetSensitiveDataOccurrencesAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitiveDataOccurrencesWithContext(ctx context.Context, input *macie2.GetSensitiveDataOccurrencesInput, opts ...request.Option) (*macie2.GetSensitiveDataOccurrencesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitiveDataOccurrencesOutput), nil
	}
	out, err := c.Macie2API.GetSensitiveDataOccurrencesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitivityInspectionTemplate(input *macie2.GetSensitivityInspectionTemplateInput) (*macie2.GetSensitivityInspectionTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitivityInspectionTemplateOutput), nil
	}
	out, err := c.Macie2API.GetSensitivityInspectionTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetSensitivityInspectionTemplateWithContext(ctx context.Context, input *macie2.GetSensitivityInspectionTemplateInput, opts ...request.Option) (*macie2.GetSensitivityInspectionTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetSensitivityInspectionTemplateOutput), nil
	}
	out, err := c.Macie2API.GetSensitivityInspectionTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetUsageStatistics(input *macie2.GetUsageStatisticsInput) (*macie2.GetUsageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetUsageStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetUsageStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetUsageStatisticsPages(input *macie2.GetUsageStatisticsInput, fn func(*macie2.GetUsageStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.GetUsageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.GetUsageStatisticsOutput{}
	fnCacher := func(out *macie2.GetUsageStatisticsOutput, more bool) bool {
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
	if err := c.Macie2API.GetUsageStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) GetUsageStatisticsPagesWithContext(ctx context.Context, input *macie2.GetUsageStatisticsInput, fn func(*macie2.GetUsageStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.GetUsageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.GetUsageStatisticsOutput{}
	fnCacher := func(out *macie2.GetUsageStatisticsOutput, more bool) bool {
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
	if err := c.Macie2API.GetUsageStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) GetUsageStatisticsWithContext(ctx context.Context, input *macie2.GetUsageStatisticsInput, opts ...request.Option) (*macie2.GetUsageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetUsageStatisticsOutput), nil
	}
	out, err := c.Macie2API.GetUsageStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetUsageTotals(input *macie2.GetUsageTotalsInput) (*macie2.GetUsageTotalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetUsageTotalsOutput), nil
	}
	out, err := c.Macie2API.GetUsageTotals(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) GetUsageTotalsWithContext(ctx context.Context, input *macie2.GetUsageTotalsInput, opts ...request.Option) (*macie2.GetUsageTotalsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.GetUsageTotalsOutput), nil
	}
	out, err := c.Macie2API.GetUsageTotalsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListAllowLists(input *macie2.ListAllowListsInput) (*macie2.ListAllowListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListAllowListsOutput), nil
	}
	out, err := c.Macie2API.ListAllowLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListAllowListsPages(input *macie2.ListAllowListsInput, fn func(*macie2.ListAllowListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListAllowListsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListAllowListsOutput{}
	fnCacher := func(out *macie2.ListAllowListsOutput, more bool) bool {
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
	if err := c.Macie2API.ListAllowListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListAllowListsPagesWithContext(ctx context.Context, input *macie2.ListAllowListsInput, fn func(*macie2.ListAllowListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListAllowListsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListAllowListsOutput{}
	fnCacher := func(out *macie2.ListAllowListsOutput, more bool) bool {
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
	if err := c.Macie2API.ListAllowListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListAllowListsWithContext(ctx context.Context, input *macie2.ListAllowListsInput, opts ...request.Option) (*macie2.ListAllowListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListAllowListsOutput), nil
	}
	out, err := c.Macie2API.ListAllowListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListClassificationJobs(input *macie2.ListClassificationJobsInput) (*macie2.ListClassificationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListClassificationJobsOutput), nil
	}
	out, err := c.Macie2API.ListClassificationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListClassificationJobsPages(input *macie2.ListClassificationJobsInput, fn func(*macie2.ListClassificationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListClassificationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListClassificationJobsOutput{}
	fnCacher := func(out *macie2.ListClassificationJobsOutput, more bool) bool {
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
	if err := c.Macie2API.ListClassificationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListClassificationJobsPagesWithContext(ctx context.Context, input *macie2.ListClassificationJobsInput, fn func(*macie2.ListClassificationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListClassificationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListClassificationJobsOutput{}
	fnCacher := func(out *macie2.ListClassificationJobsOutput, more bool) bool {
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
	if err := c.Macie2API.ListClassificationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListClassificationJobsWithContext(ctx context.Context, input *macie2.ListClassificationJobsInput, opts ...request.Option) (*macie2.ListClassificationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListClassificationJobsOutput), nil
	}
	out, err := c.Macie2API.ListClassificationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListClassificationScopes(input *macie2.ListClassificationScopesInput) (*macie2.ListClassificationScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListClassificationScopesOutput), nil
	}
	out, err := c.Macie2API.ListClassificationScopes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListClassificationScopesPages(input *macie2.ListClassificationScopesInput, fn func(*macie2.ListClassificationScopesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListClassificationScopesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListClassificationScopesOutput{}
	fnCacher := func(out *macie2.ListClassificationScopesOutput, more bool) bool {
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
	if err := c.Macie2API.ListClassificationScopesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListClassificationScopesPagesWithContext(ctx context.Context, input *macie2.ListClassificationScopesInput, fn func(*macie2.ListClassificationScopesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListClassificationScopesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListClassificationScopesOutput{}
	fnCacher := func(out *macie2.ListClassificationScopesOutput, more bool) bool {
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
	if err := c.Macie2API.ListClassificationScopesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListClassificationScopesWithContext(ctx context.Context, input *macie2.ListClassificationScopesInput, opts ...request.Option) (*macie2.ListClassificationScopesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListClassificationScopesOutput), nil
	}
	out, err := c.Macie2API.ListClassificationScopesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListCustomDataIdentifiers(input *macie2.ListCustomDataIdentifiersInput) (*macie2.ListCustomDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListCustomDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.ListCustomDataIdentifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListCustomDataIdentifiersPages(input *macie2.ListCustomDataIdentifiersInput, fn func(*macie2.ListCustomDataIdentifiersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListCustomDataIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListCustomDataIdentifiersOutput{}
	fnCacher := func(out *macie2.ListCustomDataIdentifiersOutput, more bool) bool {
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
	if err := c.Macie2API.ListCustomDataIdentifiersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListCustomDataIdentifiersPagesWithContext(ctx context.Context, input *macie2.ListCustomDataIdentifiersInput, fn func(*macie2.ListCustomDataIdentifiersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListCustomDataIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListCustomDataIdentifiersOutput{}
	fnCacher := func(out *macie2.ListCustomDataIdentifiersOutput, more bool) bool {
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
	if err := c.Macie2API.ListCustomDataIdentifiersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListCustomDataIdentifiersWithContext(ctx context.Context, input *macie2.ListCustomDataIdentifiersInput, opts ...request.Option) (*macie2.ListCustomDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListCustomDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.ListCustomDataIdentifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListFindings(input *macie2.ListFindingsInput) (*macie2.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListFindingsOutput), nil
	}
	out, err := c.Macie2API.ListFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListFindingsFilters(input *macie2.ListFindingsFiltersInput) (*macie2.ListFindingsFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListFindingsFiltersOutput), nil
	}
	out, err := c.Macie2API.ListFindingsFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListFindingsFiltersPages(input *macie2.ListFindingsFiltersInput, fn func(*macie2.ListFindingsFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListFindingsFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListFindingsFiltersOutput{}
	fnCacher := func(out *macie2.ListFindingsFiltersOutput, more bool) bool {
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
	if err := c.Macie2API.ListFindingsFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListFindingsFiltersPagesWithContext(ctx context.Context, input *macie2.ListFindingsFiltersInput, fn func(*macie2.ListFindingsFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListFindingsFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListFindingsFiltersOutput{}
	fnCacher := func(out *macie2.ListFindingsFiltersOutput, more bool) bool {
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
	if err := c.Macie2API.ListFindingsFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListFindingsFiltersWithContext(ctx context.Context, input *macie2.ListFindingsFiltersInput, opts ...request.Option) (*macie2.ListFindingsFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListFindingsFiltersOutput), nil
	}
	out, err := c.Macie2API.ListFindingsFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListFindingsPages(input *macie2.ListFindingsInput, fn func(*macie2.ListFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListFindingsOutput{}
	fnCacher := func(out *macie2.ListFindingsOutput, more bool) bool {
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
	if err := c.Macie2API.ListFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListFindingsPagesWithContext(ctx context.Context, input *macie2.ListFindingsInput, fn func(*macie2.ListFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListFindingsOutput{}
	fnCacher := func(out *macie2.ListFindingsOutput, more bool) bool {
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
	if err := c.Macie2API.ListFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListFindingsWithContext(ctx context.Context, input *macie2.ListFindingsInput, opts ...request.Option) (*macie2.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListFindingsOutput), nil
	}
	out, err := c.Macie2API.ListFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListInvitations(input *macie2.ListInvitationsInput) (*macie2.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListInvitationsOutput), nil
	}
	out, err := c.Macie2API.ListInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListInvitationsPages(input *macie2.ListInvitationsInput, fn func(*macie2.ListInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListInvitationsOutput{}
	fnCacher := func(out *macie2.ListInvitationsOutput, more bool) bool {
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
	if err := c.Macie2API.ListInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListInvitationsPagesWithContext(ctx context.Context, input *macie2.ListInvitationsInput, fn func(*macie2.ListInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListInvitationsOutput{}
	fnCacher := func(out *macie2.ListInvitationsOutput, more bool) bool {
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
	if err := c.Macie2API.ListInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListInvitationsWithContext(ctx context.Context, input *macie2.ListInvitationsInput, opts ...request.Option) (*macie2.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListInvitationsOutput), nil
	}
	out, err := c.Macie2API.ListInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListManagedDataIdentifiers(input *macie2.ListManagedDataIdentifiersInput) (*macie2.ListManagedDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListManagedDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.ListManagedDataIdentifiers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListManagedDataIdentifiersPages(input *macie2.ListManagedDataIdentifiersInput, fn func(*macie2.ListManagedDataIdentifiersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListManagedDataIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListManagedDataIdentifiersOutput{}
	fnCacher := func(out *macie2.ListManagedDataIdentifiersOutput, more bool) bool {
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
	if err := c.Macie2API.ListManagedDataIdentifiersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListManagedDataIdentifiersPagesWithContext(ctx context.Context, input *macie2.ListManagedDataIdentifiersInput, fn func(*macie2.ListManagedDataIdentifiersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListManagedDataIdentifiersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListManagedDataIdentifiersOutput{}
	fnCacher := func(out *macie2.ListManagedDataIdentifiersOutput, more bool) bool {
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
	if err := c.Macie2API.ListManagedDataIdentifiersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListManagedDataIdentifiersWithContext(ctx context.Context, input *macie2.ListManagedDataIdentifiersInput, opts ...request.Option) (*macie2.ListManagedDataIdentifiersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListManagedDataIdentifiersOutput), nil
	}
	out, err := c.Macie2API.ListManagedDataIdentifiersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListMembers(input *macie2.ListMembersInput) (*macie2.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListMembersOutput), nil
	}
	out, err := c.Macie2API.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListMembersPages(input *macie2.ListMembersInput, fn func(*macie2.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListMembersOutput{}
	fnCacher := func(out *macie2.ListMembersOutput, more bool) bool {
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
	if err := c.Macie2API.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListMembersPagesWithContext(ctx context.Context, input *macie2.ListMembersInput, fn func(*macie2.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListMembersOutput{}
	fnCacher := func(out *macie2.ListMembersOutput, more bool) bool {
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
	if err := c.Macie2API.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListMembersWithContext(ctx context.Context, input *macie2.ListMembersInput, opts ...request.Option) (*macie2.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListMembersOutput), nil
	}
	out, err := c.Macie2API.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListOrganizationAdminAccounts(input *macie2.ListOrganizationAdminAccountsInput) (*macie2.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.Macie2API.ListOrganizationAdminAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListOrganizationAdminAccountsPages(input *macie2.ListOrganizationAdminAccountsInput, fn func(*macie2.ListOrganizationAdminAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *macie2.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.Macie2API.ListOrganizationAdminAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListOrganizationAdminAccountsPagesWithContext(ctx context.Context, input *macie2.ListOrganizationAdminAccountsInput, fn func(*macie2.ListOrganizationAdminAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *macie2.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.Macie2API.ListOrganizationAdminAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListOrganizationAdminAccountsWithContext(ctx context.Context, input *macie2.ListOrganizationAdminAccountsInput, opts ...request.Option) (*macie2.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.Macie2API.ListOrganizationAdminAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListResourceProfileArtifacts(input *macie2.ListResourceProfileArtifactsInput) (*macie2.ListResourceProfileArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListResourceProfileArtifactsOutput), nil
	}
	out, err := c.Macie2API.ListResourceProfileArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListResourceProfileArtifactsPages(input *macie2.ListResourceProfileArtifactsInput, fn func(*macie2.ListResourceProfileArtifactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListResourceProfileArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListResourceProfileArtifactsOutput{}
	fnCacher := func(out *macie2.ListResourceProfileArtifactsOutput, more bool) bool {
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
	if err := c.Macie2API.ListResourceProfileArtifactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListResourceProfileArtifactsPagesWithContext(ctx context.Context, input *macie2.ListResourceProfileArtifactsInput, fn func(*macie2.ListResourceProfileArtifactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListResourceProfileArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListResourceProfileArtifactsOutput{}
	fnCacher := func(out *macie2.ListResourceProfileArtifactsOutput, more bool) bool {
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
	if err := c.Macie2API.ListResourceProfileArtifactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListResourceProfileArtifactsWithContext(ctx context.Context, input *macie2.ListResourceProfileArtifactsInput, opts ...request.Option) (*macie2.ListResourceProfileArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListResourceProfileArtifactsOutput), nil
	}
	out, err := c.Macie2API.ListResourceProfileArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListResourceProfileDetections(input *macie2.ListResourceProfileDetectionsInput) (*macie2.ListResourceProfileDetectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListResourceProfileDetectionsOutput), nil
	}
	out, err := c.Macie2API.ListResourceProfileDetections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListResourceProfileDetectionsPages(input *macie2.ListResourceProfileDetectionsInput, fn func(*macie2.ListResourceProfileDetectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListResourceProfileDetectionsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListResourceProfileDetectionsOutput{}
	fnCacher := func(out *macie2.ListResourceProfileDetectionsOutput, more bool) bool {
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
	if err := c.Macie2API.ListResourceProfileDetectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListResourceProfileDetectionsPagesWithContext(ctx context.Context, input *macie2.ListResourceProfileDetectionsInput, fn func(*macie2.ListResourceProfileDetectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListResourceProfileDetectionsOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListResourceProfileDetectionsOutput{}
	fnCacher := func(out *macie2.ListResourceProfileDetectionsOutput, more bool) bool {
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
	if err := c.Macie2API.ListResourceProfileDetectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListResourceProfileDetectionsWithContext(ctx context.Context, input *macie2.ListResourceProfileDetectionsInput, opts ...request.Option) (*macie2.ListResourceProfileDetectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListResourceProfileDetectionsOutput), nil
	}
	out, err := c.Macie2API.ListResourceProfileDetectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListSensitivityInspectionTemplates(input *macie2.ListSensitivityInspectionTemplatesInput) (*macie2.ListSensitivityInspectionTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListSensitivityInspectionTemplatesOutput), nil
	}
	out, err := c.Macie2API.ListSensitivityInspectionTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListSensitivityInspectionTemplatesPages(input *macie2.ListSensitivityInspectionTemplatesInput, fn func(*macie2.ListSensitivityInspectionTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListSensitivityInspectionTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListSensitivityInspectionTemplatesOutput{}
	fnCacher := func(out *macie2.ListSensitivityInspectionTemplatesOutput, more bool) bool {
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
	if err := c.Macie2API.ListSensitivityInspectionTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListSensitivityInspectionTemplatesPagesWithContext(ctx context.Context, input *macie2.ListSensitivityInspectionTemplatesInput, fn func(*macie2.ListSensitivityInspectionTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.ListSensitivityInspectionTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.ListSensitivityInspectionTemplatesOutput{}
	fnCacher := func(out *macie2.ListSensitivityInspectionTemplatesOutput, more bool) bool {
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
	if err := c.Macie2API.ListSensitivityInspectionTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) ListSensitivityInspectionTemplatesWithContext(ctx context.Context, input *macie2.ListSensitivityInspectionTemplatesInput, opts ...request.Option) (*macie2.ListSensitivityInspectionTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListSensitivityInspectionTemplatesOutput), nil
	}
	out, err := c.Macie2API.ListSensitivityInspectionTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListTagsForResource(input *macie2.ListTagsForResourceInput) (*macie2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListTagsForResourceOutput), nil
	}
	out, err := c.Macie2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) ListTagsForResourceWithContext(ctx context.Context, input *macie2.ListTagsForResourceInput, opts ...request.Option) (*macie2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.ListTagsForResourceOutput), nil
	}
	out, err := c.Macie2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) SearchResources(input *macie2.SearchResourcesInput) (*macie2.SearchResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.SearchResourcesOutput), nil
	}
	out, err := c.Macie2API.SearchResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Macie2) SearchResourcesPages(input *macie2.SearchResourcesInput, fn func(*macie2.SearchResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.SearchResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.SearchResourcesOutput{}
	fnCacher := func(out *macie2.SearchResourcesOutput, more bool) bool {
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
	if err := c.Macie2API.SearchResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) SearchResourcesPagesWithContext(ctx context.Context, input *macie2.SearchResourcesInput, fn func(*macie2.SearchResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*macie2.SearchResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &macie2.SearchResourcesOutput{}
	fnCacher := func(out *macie2.SearchResourcesOutput, more bool) bool {
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
	if err := c.Macie2API.SearchResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Macie2) SearchResourcesWithContext(ctx context.Context, input *macie2.SearchResourcesInput, opts ...request.Option) (*macie2.SearchResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*macie2.SearchResourcesOutput), nil
	}
	out, err := c.Macie2API.SearchResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
