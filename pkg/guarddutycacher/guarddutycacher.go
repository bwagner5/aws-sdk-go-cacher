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
package guarddutycacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/guardduty"
	"github.com/aws/aws-sdk-go/service/guardduty/guarddutyiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type GuardDuty struct {
	guarddutyiface.GuardDutyAPI
	cache *cache.Cache
}

func New(guarddutyapi guarddutyiface.GuardDutyAPI) *GuardDuty {
	return &GuardDuty{
		GuardDutyAPI: guarddutyapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *GuardDuty) DescribeMalwareScans(input *guardduty.DescribeMalwareScansInput) (*guardduty.DescribeMalwareScansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribeMalwareScansOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribeMalwareScans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) DescribeMalwareScansPages(input *guardduty.DescribeMalwareScansInput, fn func(*guardduty.DescribeMalwareScansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.DescribeMalwareScansOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.DescribeMalwareScansOutput{}
	fnCacher := func(out *guardduty.DescribeMalwareScansOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.DescribeMalwareScansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) DescribeMalwareScansPagesWithContext(ctx context.Context, input *guardduty.DescribeMalwareScansInput, fn func(*guardduty.DescribeMalwareScansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.DescribeMalwareScansOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.DescribeMalwareScansOutput{}
	fnCacher := func(out *guardduty.DescribeMalwareScansOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.DescribeMalwareScansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) DescribeMalwareScansWithContext(ctx context.Context, input *guardduty.DescribeMalwareScansInput, opts ...request.Option) (*guardduty.DescribeMalwareScansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribeMalwareScansOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribeMalwareScansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) DescribeOrganizationConfiguration(input *guardduty.DescribeOrganizationConfigurationInput) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribeOrganizationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) DescribeOrganizationConfigurationWithContext(ctx context.Context, input *guardduty.DescribeOrganizationConfigurationInput, opts ...request.Option) (*guardduty.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribeOrganizationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) DescribePublishingDestination(input *guardduty.DescribePublishingDestinationInput) (*guardduty.DescribePublishingDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribePublishingDestinationOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribePublishingDestination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) DescribePublishingDestinationWithContext(ctx context.Context, input *guardduty.DescribePublishingDestinationInput, opts ...request.Option) (*guardduty.DescribePublishingDestinationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.DescribePublishingDestinationOutput), nil
	}
	out, err := c.GuardDutyAPI.DescribePublishingDestinationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetAdministratorAccount(input *guardduty.GetAdministratorAccountInput) (*guardduty.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetAdministratorAccountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetAdministratorAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetAdministratorAccountWithContext(ctx context.Context, input *guardduty.GetAdministratorAccountInput, opts ...request.Option) (*guardduty.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetAdministratorAccountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetAdministratorAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetDetector(input *guardduty.GetDetectorInput) (*guardduty.GetDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetDetectorOutput), nil
	}
	out, err := c.GuardDutyAPI.GetDetector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetDetectorWithContext(ctx context.Context, input *guardduty.GetDetectorInput, opts ...request.Option) (*guardduty.GetDetectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetDetectorOutput), nil
	}
	out, err := c.GuardDutyAPI.GetDetectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFilter(input *guardduty.GetFilterInput) (*guardduty.GetFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFilterOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFilter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFilterWithContext(ctx context.Context, input *guardduty.GetFilterInput, opts ...request.Option) (*guardduty.GetFilterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFilterOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFilterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFindings(input *guardduty.GetFindingsInput) (*guardduty.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFindingsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFindingsStatistics(input *guardduty.GetFindingsStatisticsInput) (*guardduty.GetFindingsStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFindingsStatisticsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFindingsStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFindingsStatisticsWithContext(ctx context.Context, input *guardduty.GetFindingsStatisticsInput, opts ...request.Option) (*guardduty.GetFindingsStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFindingsStatisticsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFindingsStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetFindingsWithContext(ctx context.Context, input *guardduty.GetFindingsInput, opts ...request.Option) (*guardduty.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetFindingsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetIPSet(input *guardduty.GetIPSetInput) (*guardduty.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetIPSetOutput), nil
	}
	out, err := c.GuardDutyAPI.GetIPSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetIPSetWithContext(ctx context.Context, input *guardduty.GetIPSetInput, opts ...request.Option) (*guardduty.GetIPSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetIPSetOutput), nil
	}
	out, err := c.GuardDutyAPI.GetIPSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetInvitationsCount(input *guardduty.GetInvitationsCountInput) (*guardduty.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetInvitationsCountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetInvitationsCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetInvitationsCountWithContext(ctx context.Context, input *guardduty.GetInvitationsCountInput, opts ...request.Option) (*guardduty.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetInvitationsCountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetInvitationsCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMalwareScanSettings(input *guardduty.GetMalwareScanSettingsInput) (*guardduty.GetMalwareScanSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMalwareScanSettingsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMalwareScanSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMalwareScanSettingsWithContext(ctx context.Context, input *guardduty.GetMalwareScanSettingsInput, opts ...request.Option) (*guardduty.GetMalwareScanSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMalwareScanSettingsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMalwareScanSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMasterAccount(input *guardduty.GetMasterAccountInput) (*guardduty.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMasterAccountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMasterAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMasterAccountWithContext(ctx context.Context, input *guardduty.GetMasterAccountInput, opts ...request.Option) (*guardduty.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMasterAccountOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMasterAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMemberDetectors(input *guardduty.GetMemberDetectorsInput) (*guardduty.GetMemberDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMemberDetectorsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMemberDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMemberDetectorsWithContext(ctx context.Context, input *guardduty.GetMemberDetectorsInput, opts ...request.Option) (*guardduty.GetMemberDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMemberDetectorsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMemberDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMembers(input *guardduty.GetMembersInput) (*guardduty.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMembersOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetMembersWithContext(ctx context.Context, input *guardduty.GetMembersInput, opts ...request.Option) (*guardduty.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetMembersOutput), nil
	}
	out, err := c.GuardDutyAPI.GetMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetRemainingFreeTrialDays(input *guardduty.GetRemainingFreeTrialDaysInput) (*guardduty.GetRemainingFreeTrialDaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetRemainingFreeTrialDaysOutput), nil
	}
	out, err := c.GuardDutyAPI.GetRemainingFreeTrialDays(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetRemainingFreeTrialDaysWithContext(ctx context.Context, input *guardduty.GetRemainingFreeTrialDaysInput, opts ...request.Option) (*guardduty.GetRemainingFreeTrialDaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetRemainingFreeTrialDaysOutput), nil
	}
	out, err := c.GuardDutyAPI.GetRemainingFreeTrialDaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetThreatIntelSet(input *guardduty.GetThreatIntelSetInput) (*guardduty.GetThreatIntelSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetThreatIntelSetOutput), nil
	}
	out, err := c.GuardDutyAPI.GetThreatIntelSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetThreatIntelSetWithContext(ctx context.Context, input *guardduty.GetThreatIntelSetInput, opts ...request.Option) (*guardduty.GetThreatIntelSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetThreatIntelSetOutput), nil
	}
	out, err := c.GuardDutyAPI.GetThreatIntelSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetUsageStatistics(input *guardduty.GetUsageStatisticsInput) (*guardduty.GetUsageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetUsageStatisticsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetUsageStatistics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) GetUsageStatisticsPages(input *guardduty.GetUsageStatisticsInput, fn func(*guardduty.GetUsageStatisticsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.GetUsageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.GetUsageStatisticsOutput{}
	fnCacher := func(out *guardduty.GetUsageStatisticsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.GetUsageStatisticsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) GetUsageStatisticsPagesWithContext(ctx context.Context, input *guardduty.GetUsageStatisticsInput, fn func(*guardduty.GetUsageStatisticsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.GetUsageStatisticsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.GetUsageStatisticsOutput{}
	fnCacher := func(out *guardduty.GetUsageStatisticsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.GetUsageStatisticsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) GetUsageStatisticsWithContext(ctx context.Context, input *guardduty.GetUsageStatisticsInput, opts ...request.Option) (*guardduty.GetUsageStatisticsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.GetUsageStatisticsOutput), nil
	}
	out, err := c.GuardDutyAPI.GetUsageStatisticsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListDetectors(input *guardduty.ListDetectorsInput) (*guardduty.ListDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListDetectorsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListDetectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListDetectorsPages(input *guardduty.ListDetectorsInput, fn func(*guardduty.ListDetectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListDetectorsOutput{}
	fnCacher := func(out *guardduty.ListDetectorsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListDetectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListDetectorsPagesWithContext(ctx context.Context, input *guardduty.ListDetectorsInput, fn func(*guardduty.ListDetectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListDetectorsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListDetectorsOutput{}
	fnCacher := func(out *guardduty.ListDetectorsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListDetectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListDetectorsWithContext(ctx context.Context, input *guardduty.ListDetectorsInput, opts ...request.Option) (*guardduty.ListDetectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListDetectorsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListDetectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListFilters(input *guardduty.ListFiltersInput) (*guardduty.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListFiltersOutput), nil
	}
	out, err := c.GuardDutyAPI.ListFilters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListFiltersPages(input *guardduty.ListFiltersInput, fn func(*guardduty.ListFiltersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListFiltersOutput{}
	fnCacher := func(out *guardduty.ListFiltersOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListFiltersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListFiltersPagesWithContext(ctx context.Context, input *guardduty.ListFiltersInput, fn func(*guardduty.ListFiltersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListFiltersOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListFiltersOutput{}
	fnCacher := func(out *guardduty.ListFiltersOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListFiltersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListFiltersWithContext(ctx context.Context, input *guardduty.ListFiltersInput, opts ...request.Option) (*guardduty.ListFiltersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListFiltersOutput), nil
	}
	out, err := c.GuardDutyAPI.ListFiltersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListFindings(input *guardduty.ListFindingsInput) (*guardduty.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListFindingsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListFindingsPages(input *guardduty.ListFindingsInput, fn func(*guardduty.ListFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListFindingsOutput{}
	fnCacher := func(out *guardduty.ListFindingsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListFindingsPagesWithContext(ctx context.Context, input *guardduty.ListFindingsInput, fn func(*guardduty.ListFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListFindingsOutput{}
	fnCacher := func(out *guardduty.ListFindingsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListFindingsWithContext(ctx context.Context, input *guardduty.ListFindingsInput, opts ...request.Option) (*guardduty.ListFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListFindingsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListIPSets(input *guardduty.ListIPSetsInput) (*guardduty.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListIPSetsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListIPSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListIPSetsPages(input *guardduty.ListIPSetsInput, fn func(*guardduty.ListIPSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListIPSetsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListIPSetsOutput{}
	fnCacher := func(out *guardduty.ListIPSetsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListIPSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListIPSetsPagesWithContext(ctx context.Context, input *guardduty.ListIPSetsInput, fn func(*guardduty.ListIPSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListIPSetsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListIPSetsOutput{}
	fnCacher := func(out *guardduty.ListIPSetsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListIPSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListIPSetsWithContext(ctx context.Context, input *guardduty.ListIPSetsInput, opts ...request.Option) (*guardduty.ListIPSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListIPSetsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListIPSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListInvitations(input *guardduty.ListInvitationsInput) (*guardduty.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListInvitationsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListInvitationsPages(input *guardduty.ListInvitationsInput, fn func(*guardduty.ListInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListInvitationsOutput{}
	fnCacher := func(out *guardduty.ListInvitationsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListInvitationsPagesWithContext(ctx context.Context, input *guardduty.ListInvitationsInput, fn func(*guardduty.ListInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListInvitationsOutput{}
	fnCacher := func(out *guardduty.ListInvitationsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListInvitationsWithContext(ctx context.Context, input *guardduty.ListInvitationsInput, opts ...request.Option) (*guardduty.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListInvitationsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListMembers(input *guardduty.ListMembersInput) (*guardduty.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListMembersOutput), nil
	}
	out, err := c.GuardDutyAPI.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListMembersPages(input *guardduty.ListMembersInput, fn func(*guardduty.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListMembersOutput{}
	fnCacher := func(out *guardduty.ListMembersOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListMembersPagesWithContext(ctx context.Context, input *guardduty.ListMembersInput, fn func(*guardduty.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListMembersOutput{}
	fnCacher := func(out *guardduty.ListMembersOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListMembersWithContext(ctx context.Context, input *guardduty.ListMembersInput, opts ...request.Option) (*guardduty.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListMembersOutput), nil
	}
	out, err := c.GuardDutyAPI.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListOrganizationAdminAccounts(input *guardduty.ListOrganizationAdminAccountsInput) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListOrganizationAdminAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListOrganizationAdminAccountsPages(input *guardduty.ListOrganizationAdminAccountsInput, fn func(*guardduty.ListOrganizationAdminAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *guardduty.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListOrganizationAdminAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListOrganizationAdminAccountsPagesWithContext(ctx context.Context, input *guardduty.ListOrganizationAdminAccountsInput, fn func(*guardduty.ListOrganizationAdminAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *guardduty.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListOrganizationAdminAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListOrganizationAdminAccountsWithContext(ctx context.Context, input *guardduty.ListOrganizationAdminAccountsInput, opts ...request.Option) (*guardduty.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListOrganizationAdminAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListPublishingDestinations(input *guardduty.ListPublishingDestinationsInput) (*guardduty.ListPublishingDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListPublishingDestinationsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListPublishingDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListPublishingDestinationsPages(input *guardduty.ListPublishingDestinationsInput, fn func(*guardduty.ListPublishingDestinationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListPublishingDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListPublishingDestinationsOutput{}
	fnCacher := func(out *guardduty.ListPublishingDestinationsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListPublishingDestinationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListPublishingDestinationsPagesWithContext(ctx context.Context, input *guardduty.ListPublishingDestinationsInput, fn func(*guardduty.ListPublishingDestinationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListPublishingDestinationsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListPublishingDestinationsOutput{}
	fnCacher := func(out *guardduty.ListPublishingDestinationsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListPublishingDestinationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListPublishingDestinationsWithContext(ctx context.Context, input *guardduty.ListPublishingDestinationsInput, opts ...request.Option) (*guardduty.ListPublishingDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListPublishingDestinationsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListPublishingDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListTagsForResource(input *guardduty.ListTagsForResourceInput) (*guardduty.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListTagsForResourceOutput), nil
	}
	out, err := c.GuardDutyAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListTagsForResourceWithContext(ctx context.Context, input *guardduty.ListTagsForResourceInput, opts ...request.Option) (*guardduty.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListTagsForResourceOutput), nil
	}
	out, err := c.GuardDutyAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListThreatIntelSets(input *guardduty.ListThreatIntelSetsInput) (*guardduty.ListThreatIntelSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListThreatIntelSetsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListThreatIntelSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *GuardDuty) ListThreatIntelSetsPages(input *guardduty.ListThreatIntelSetsInput, fn func(*guardduty.ListThreatIntelSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListThreatIntelSetsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListThreatIntelSetsOutput{}
	fnCacher := func(out *guardduty.ListThreatIntelSetsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListThreatIntelSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListThreatIntelSetsPagesWithContext(ctx context.Context, input *guardduty.ListThreatIntelSetsInput, fn func(*guardduty.ListThreatIntelSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*guardduty.ListThreatIntelSetsOutput), false)
		return nil
	}
	cachable := true
	output := &guardduty.ListThreatIntelSetsOutput{}
	fnCacher := func(out *guardduty.ListThreatIntelSetsOutput, more bool) bool {
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
	if err := c.GuardDutyAPI.ListThreatIntelSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *GuardDuty) ListThreatIntelSetsWithContext(ctx context.Context, input *guardduty.ListThreatIntelSetsInput, opts ...request.Option) (*guardduty.ListThreatIntelSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*guardduty.ListThreatIntelSetsOutput), nil
	}
	out, err := c.GuardDutyAPI.ListThreatIntelSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
