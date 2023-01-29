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
package pinpointcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pinpoint"
	"github.com/aws/aws-sdk-go/service/pinpoint/pinpointiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Pinpoint struct {
	pinpointiface.PinpointAPI
	cache *cache.Cache
}

func New(pinpointapi pinpointiface.PinpointAPI) *Pinpoint {
	return &Pinpoint{
		PinpointAPI: pinpointapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Pinpoint) GetAdmChannel(input *pinpoint.GetAdmChannelInput) (*pinpoint.GetAdmChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAdmChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetAdmChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetAdmChannelWithContext(ctx context.Context, input *pinpoint.GetAdmChannelInput, opts ...request.Option) (*pinpoint.GetAdmChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAdmChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetAdmChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsChannel(input *pinpoint.GetApnsChannelInput) (*pinpoint.GetApnsChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsChannelWithContext(ctx context.Context, input *pinpoint.GetApnsChannelInput, opts ...request.Option) (*pinpoint.GetApnsChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsSandboxChannel(input *pinpoint.GetApnsSandboxChannelInput) (*pinpoint.GetApnsSandboxChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsSandboxChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsSandboxChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsSandboxChannelWithContext(ctx context.Context, input *pinpoint.GetApnsSandboxChannelInput, opts ...request.Option) (*pinpoint.GetApnsSandboxChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsSandboxChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsSandboxChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsVoipChannel(input *pinpoint.GetApnsVoipChannelInput) (*pinpoint.GetApnsVoipChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsVoipChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsVoipChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsVoipChannelWithContext(ctx context.Context, input *pinpoint.GetApnsVoipChannelInput, opts ...request.Option) (*pinpoint.GetApnsVoipChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsVoipChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsVoipChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsVoipSandboxChannel(input *pinpoint.GetApnsVoipSandboxChannelInput) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsVoipSandboxChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsVoipSandboxChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApnsVoipSandboxChannelWithContext(ctx context.Context, input *pinpoint.GetApnsVoipSandboxChannelInput, opts ...request.Option) (*pinpoint.GetApnsVoipSandboxChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApnsVoipSandboxChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetApnsVoipSandboxChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApp(input *pinpoint.GetAppInput) (*pinpoint.GetAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAppOutput), nil
	}
	out, err := c.PinpointAPI.GetApp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetAppWithContext(ctx context.Context, input *pinpoint.GetAppInput, opts ...request.Option) (*pinpoint.GetAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAppOutput), nil
	}
	out, err := c.PinpointAPI.GetAppWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApplicationDateRangeKpi(input *pinpoint.GetApplicationDateRangeKpiInput) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApplicationDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetApplicationDateRangeKpi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApplicationDateRangeKpiWithContext(ctx context.Context, input *pinpoint.GetApplicationDateRangeKpiInput, opts ...request.Option) (*pinpoint.GetApplicationDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApplicationDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetApplicationDateRangeKpiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApplicationSettings(input *pinpoint.GetApplicationSettingsInput) (*pinpoint.GetApplicationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApplicationSettingsOutput), nil
	}
	out, err := c.PinpointAPI.GetApplicationSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApplicationSettingsWithContext(ctx context.Context, input *pinpoint.GetApplicationSettingsInput, opts ...request.Option) (*pinpoint.GetApplicationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetApplicationSettingsOutput), nil
	}
	out, err := c.PinpointAPI.GetApplicationSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetApps(input *pinpoint.GetAppsInput) (*pinpoint.GetAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAppsOutput), nil
	}
	out, err := c.PinpointAPI.GetApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetAppsWithContext(ctx context.Context, input *pinpoint.GetAppsInput, opts ...request.Option) (*pinpoint.GetAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetAppsOutput), nil
	}
	out, err := c.PinpointAPI.GetAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetBaiduChannel(input *pinpoint.GetBaiduChannelInput) (*pinpoint.GetBaiduChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetBaiduChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetBaiduChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetBaiduChannelWithContext(ctx context.Context, input *pinpoint.GetBaiduChannelInput, opts ...request.Option) (*pinpoint.GetBaiduChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetBaiduChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetBaiduChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaign(input *pinpoint.GetCampaignInput) (*pinpoint.GetCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaign(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignActivities(input *pinpoint.GetCampaignActivitiesInput) (*pinpoint.GetCampaignActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignActivitiesOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignActivities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignActivitiesWithContext(ctx context.Context, input *pinpoint.GetCampaignActivitiesInput, opts ...request.Option) (*pinpoint.GetCampaignActivitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignActivitiesOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignActivitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignDateRangeKpi(input *pinpoint.GetCampaignDateRangeKpiInput) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignDateRangeKpi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignDateRangeKpiWithContext(ctx context.Context, input *pinpoint.GetCampaignDateRangeKpiInput, opts ...request.Option) (*pinpoint.GetCampaignDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignDateRangeKpiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignVersion(input *pinpoint.GetCampaignVersionInput) (*pinpoint.GetCampaignVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignVersionOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignVersionWithContext(ctx context.Context, input *pinpoint.GetCampaignVersionInput, opts ...request.Option) (*pinpoint.GetCampaignVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignVersionOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignVersions(input *pinpoint.GetCampaignVersionsInput) (*pinpoint.GetCampaignVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignVersionsOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignVersionsWithContext(ctx context.Context, input *pinpoint.GetCampaignVersionsInput, opts ...request.Option) (*pinpoint.GetCampaignVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignVersionsOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignWithContext(ctx context.Context, input *pinpoint.GetCampaignInput, opts ...request.Option) (*pinpoint.GetCampaignOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaigns(input *pinpoint.GetCampaignsInput) (*pinpoint.GetCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignsOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaigns(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetCampaignsWithContext(ctx context.Context, input *pinpoint.GetCampaignsInput, opts ...request.Option) (*pinpoint.GetCampaignsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetCampaignsOutput), nil
	}
	out, err := c.PinpointAPI.GetCampaignsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetChannels(input *pinpoint.GetChannelsInput) (*pinpoint.GetChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetChannelsOutput), nil
	}
	out, err := c.PinpointAPI.GetChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetChannelsWithContext(ctx context.Context, input *pinpoint.GetChannelsInput, opts ...request.Option) (*pinpoint.GetChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetChannelsOutput), nil
	}
	out, err := c.PinpointAPI.GetChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEmailChannel(input *pinpoint.GetEmailChannelInput) (*pinpoint.GetEmailChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEmailChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetEmailChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEmailChannelWithContext(ctx context.Context, input *pinpoint.GetEmailChannelInput, opts ...request.Option) (*pinpoint.GetEmailChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEmailChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetEmailChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEmailTemplate(input *pinpoint.GetEmailTemplateInput) (*pinpoint.GetEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEmailTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetEmailTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEmailTemplateWithContext(ctx context.Context, input *pinpoint.GetEmailTemplateInput, opts ...request.Option) (*pinpoint.GetEmailTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEmailTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetEmailTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEndpoint(input *pinpoint.GetEndpointInput) (*pinpoint.GetEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEndpointOutput), nil
	}
	out, err := c.PinpointAPI.GetEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEndpointWithContext(ctx context.Context, input *pinpoint.GetEndpointInput, opts ...request.Option) (*pinpoint.GetEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEndpointOutput), nil
	}
	out, err := c.PinpointAPI.GetEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEventStream(input *pinpoint.GetEventStreamInput) (*pinpoint.GetEventStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEventStreamOutput), nil
	}
	out, err := c.PinpointAPI.GetEventStream(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetEventStreamWithContext(ctx context.Context, input *pinpoint.GetEventStreamInput, opts ...request.Option) (*pinpoint.GetEventStreamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetEventStreamOutput), nil
	}
	out, err := c.PinpointAPI.GetEventStreamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetExportJob(input *pinpoint.GetExportJobInput) (*pinpoint.GetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetExportJobOutput), nil
	}
	out, err := c.PinpointAPI.GetExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetExportJobWithContext(ctx context.Context, input *pinpoint.GetExportJobInput, opts ...request.Option) (*pinpoint.GetExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetExportJobOutput), nil
	}
	out, err := c.PinpointAPI.GetExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetExportJobs(input *pinpoint.GetExportJobsInput) (*pinpoint.GetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetExportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetExportJobsWithContext(ctx context.Context, input *pinpoint.GetExportJobsInput, opts ...request.Option) (*pinpoint.GetExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetExportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetGcmChannel(input *pinpoint.GetGcmChannelInput) (*pinpoint.GetGcmChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetGcmChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetGcmChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetGcmChannelWithContext(ctx context.Context, input *pinpoint.GetGcmChannelInput, opts ...request.Option) (*pinpoint.GetGcmChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetGcmChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetGcmChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetImportJob(input *pinpoint.GetImportJobInput) (*pinpoint.GetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetImportJobOutput), nil
	}
	out, err := c.PinpointAPI.GetImportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetImportJobWithContext(ctx context.Context, input *pinpoint.GetImportJobInput, opts ...request.Option) (*pinpoint.GetImportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetImportJobOutput), nil
	}
	out, err := c.PinpointAPI.GetImportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetImportJobs(input *pinpoint.GetImportJobsInput) (*pinpoint.GetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetImportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetImportJobsWithContext(ctx context.Context, input *pinpoint.GetImportJobsInput, opts ...request.Option) (*pinpoint.GetImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetImportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetInAppMessages(input *pinpoint.GetInAppMessagesInput) (*pinpoint.GetInAppMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetInAppMessagesOutput), nil
	}
	out, err := c.PinpointAPI.GetInAppMessages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetInAppMessagesWithContext(ctx context.Context, input *pinpoint.GetInAppMessagesInput, opts ...request.Option) (*pinpoint.GetInAppMessagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetInAppMessagesOutput), nil
	}
	out, err := c.PinpointAPI.GetInAppMessagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetInAppTemplate(input *pinpoint.GetInAppTemplateInput) (*pinpoint.GetInAppTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetInAppTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetInAppTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetInAppTemplateWithContext(ctx context.Context, input *pinpoint.GetInAppTemplateInput, opts ...request.Option) (*pinpoint.GetInAppTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetInAppTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetInAppTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourney(input *pinpoint.GetJourneyInput) (*pinpoint.GetJourneyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyOutput), nil
	}
	out, err := c.PinpointAPI.GetJourney(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyDateRangeKpi(input *pinpoint.GetJourneyDateRangeKpiInput) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyDateRangeKpi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyDateRangeKpiWithContext(ctx context.Context, input *pinpoint.GetJourneyDateRangeKpiInput, opts ...request.Option) (*pinpoint.GetJourneyDateRangeKpiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyDateRangeKpiOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyDateRangeKpiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyExecutionActivityMetrics(input *pinpoint.GetJourneyExecutionActivityMetricsInput) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyExecutionActivityMetricsOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyExecutionActivityMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyExecutionActivityMetricsWithContext(ctx context.Context, input *pinpoint.GetJourneyExecutionActivityMetricsInput, opts ...request.Option) (*pinpoint.GetJourneyExecutionActivityMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyExecutionActivityMetricsOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyExecutionActivityMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyExecutionMetrics(input *pinpoint.GetJourneyExecutionMetricsInput) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyExecutionMetricsOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyExecutionMetrics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyExecutionMetricsWithContext(ctx context.Context, input *pinpoint.GetJourneyExecutionMetricsInput, opts ...request.Option) (*pinpoint.GetJourneyExecutionMetricsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyExecutionMetricsOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyExecutionMetricsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetJourneyWithContext(ctx context.Context, input *pinpoint.GetJourneyInput, opts ...request.Option) (*pinpoint.GetJourneyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetJourneyOutput), nil
	}
	out, err := c.PinpointAPI.GetJourneyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetPushTemplate(input *pinpoint.GetPushTemplateInput) (*pinpoint.GetPushTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetPushTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetPushTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetPushTemplateWithContext(ctx context.Context, input *pinpoint.GetPushTemplateInput, opts ...request.Option) (*pinpoint.GetPushTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetPushTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetPushTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetRecommenderConfiguration(input *pinpoint.GetRecommenderConfigurationInput) (*pinpoint.GetRecommenderConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetRecommenderConfigurationOutput), nil
	}
	out, err := c.PinpointAPI.GetRecommenderConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetRecommenderConfigurationWithContext(ctx context.Context, input *pinpoint.GetRecommenderConfigurationInput, opts ...request.Option) (*pinpoint.GetRecommenderConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetRecommenderConfigurationOutput), nil
	}
	out, err := c.PinpointAPI.GetRecommenderConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetRecommenderConfigurations(input *pinpoint.GetRecommenderConfigurationsInput) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetRecommenderConfigurationsOutput), nil
	}
	out, err := c.PinpointAPI.GetRecommenderConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetRecommenderConfigurationsWithContext(ctx context.Context, input *pinpoint.GetRecommenderConfigurationsInput, opts ...request.Option) (*pinpoint.GetRecommenderConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetRecommenderConfigurationsOutput), nil
	}
	out, err := c.PinpointAPI.GetRecommenderConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegment(input *pinpoint.GetSegmentInput) (*pinpoint.GetSegmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentOutput), nil
	}
	out, err := c.PinpointAPI.GetSegment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentExportJobs(input *pinpoint.GetSegmentExportJobsInput) (*pinpoint.GetSegmentExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentExportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentExportJobsWithContext(ctx context.Context, input *pinpoint.GetSegmentExportJobsInput, opts ...request.Option) (*pinpoint.GetSegmentExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentExportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentImportJobs(input *pinpoint.GetSegmentImportJobsInput) (*pinpoint.GetSegmentImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentImportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentImportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentImportJobsWithContext(ctx context.Context, input *pinpoint.GetSegmentImportJobsInput, opts ...request.Option) (*pinpoint.GetSegmentImportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentImportJobsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentImportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentVersion(input *pinpoint.GetSegmentVersionInput) (*pinpoint.GetSegmentVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentVersionOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentVersionWithContext(ctx context.Context, input *pinpoint.GetSegmentVersionInput, opts ...request.Option) (*pinpoint.GetSegmentVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentVersionOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentVersions(input *pinpoint.GetSegmentVersionsInput) (*pinpoint.GetSegmentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentVersionsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentVersionsWithContext(ctx context.Context, input *pinpoint.GetSegmentVersionsInput, opts ...request.Option) (*pinpoint.GetSegmentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentVersionsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentWithContext(ctx context.Context, input *pinpoint.GetSegmentInput, opts ...request.Option) (*pinpoint.GetSegmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegments(input *pinpoint.GetSegmentsInput) (*pinpoint.GetSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSegmentsWithContext(ctx context.Context, input *pinpoint.GetSegmentsInput, opts ...request.Option) (*pinpoint.GetSegmentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSegmentsOutput), nil
	}
	out, err := c.PinpointAPI.GetSegmentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSmsChannel(input *pinpoint.GetSmsChannelInput) (*pinpoint.GetSmsChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSmsChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetSmsChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSmsChannelWithContext(ctx context.Context, input *pinpoint.GetSmsChannelInput, opts ...request.Option) (*pinpoint.GetSmsChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSmsChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetSmsChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSmsTemplate(input *pinpoint.GetSmsTemplateInput) (*pinpoint.GetSmsTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSmsTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetSmsTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetSmsTemplateWithContext(ctx context.Context, input *pinpoint.GetSmsTemplateInput, opts ...request.Option) (*pinpoint.GetSmsTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetSmsTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetSmsTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetUserEndpoints(input *pinpoint.GetUserEndpointsInput) (*pinpoint.GetUserEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetUserEndpointsOutput), nil
	}
	out, err := c.PinpointAPI.GetUserEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetUserEndpointsWithContext(ctx context.Context, input *pinpoint.GetUserEndpointsInput, opts ...request.Option) (*pinpoint.GetUserEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetUserEndpointsOutput), nil
	}
	out, err := c.PinpointAPI.GetUserEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetVoiceChannel(input *pinpoint.GetVoiceChannelInput) (*pinpoint.GetVoiceChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetVoiceChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetVoiceChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetVoiceChannelWithContext(ctx context.Context, input *pinpoint.GetVoiceChannelInput, opts ...request.Option) (*pinpoint.GetVoiceChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetVoiceChannelOutput), nil
	}
	out, err := c.PinpointAPI.GetVoiceChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetVoiceTemplate(input *pinpoint.GetVoiceTemplateInput) (*pinpoint.GetVoiceTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetVoiceTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetVoiceTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) GetVoiceTemplateWithContext(ctx context.Context, input *pinpoint.GetVoiceTemplateInput, opts ...request.Option) (*pinpoint.GetVoiceTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.GetVoiceTemplateOutput), nil
	}
	out, err := c.PinpointAPI.GetVoiceTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListJourneys(input *pinpoint.ListJourneysInput) (*pinpoint.ListJourneysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListJourneysOutput), nil
	}
	out, err := c.PinpointAPI.ListJourneys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListJourneysWithContext(ctx context.Context, input *pinpoint.ListJourneysInput, opts ...request.Option) (*pinpoint.ListJourneysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListJourneysOutput), nil
	}
	out, err := c.PinpointAPI.ListJourneysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTagsForResource(input *pinpoint.ListTagsForResourceInput) (*pinpoint.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTagsForResourceWithContext(ctx context.Context, input *pinpoint.ListTagsForResourceInput, opts ...request.Option) (*pinpoint.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTemplateVersions(input *pinpoint.ListTemplateVersionsInput) (*pinpoint.ListTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTemplateVersionsOutput), nil
	}
	out, err := c.PinpointAPI.ListTemplateVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTemplateVersionsWithContext(ctx context.Context, input *pinpoint.ListTemplateVersionsInput, opts ...request.Option) (*pinpoint.ListTemplateVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTemplateVersionsOutput), nil
	}
	out, err := c.PinpointAPI.ListTemplateVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTemplates(input *pinpoint.ListTemplatesInput) (*pinpoint.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTemplatesOutput), nil
	}
	out, err := c.PinpointAPI.ListTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Pinpoint) ListTemplatesWithContext(ctx context.Context, input *pinpoint.ListTemplatesInput, opts ...request.Option) (*pinpoint.ListTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpoint.ListTemplatesOutput), nil
	}
	out, err := c.PinpointAPI.ListTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
