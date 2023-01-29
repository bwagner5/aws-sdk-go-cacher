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
package alexaforbusinesscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness/alexaforbusinessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AlexaForBusiness struct {
	alexaforbusinessiface.AlexaForBusinessAPI
	cache *cache.Cache
}

func New(alexaforbusinessapi alexaforbusinessiface.AlexaForBusinessAPI) *AlexaForBusiness {
	return &AlexaForBusiness{
		AlexaForBusinessAPI: alexaforbusinessapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AlexaForBusiness) GetAddressBook(input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetAddressBookOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetAddressBook(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetAddressBookWithContext(ctx context.Context, input *alexaforbusiness.GetAddressBookInput, opts ...request.Option) (*alexaforbusiness.GetAddressBookOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetAddressBookOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetAddressBookWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetConferencePreference(input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetConferencePreferenceOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetConferencePreference(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetConferencePreferenceWithContext(ctx context.Context, input *alexaforbusiness.GetConferencePreferenceInput, opts ...request.Option) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetConferencePreferenceOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetConferencePreferenceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetConferenceProvider(input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetConferenceProviderOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetConferenceProvider(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetConferenceProviderWithContext(ctx context.Context, input *alexaforbusiness.GetConferenceProviderInput, opts ...request.Option) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetConferenceProviderOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetConferenceProviderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetContact(input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetContactOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetContactWithContext(ctx context.Context, input *alexaforbusiness.GetContactInput, opts ...request.Option) (*alexaforbusiness.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetContactOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetDevice(input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetDeviceOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetDeviceWithContext(ctx context.Context, input *alexaforbusiness.GetDeviceInput, opts ...request.Option) (*alexaforbusiness.GetDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetDeviceOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetGateway(input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetGatewayOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetGateway(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetGatewayGroup(input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetGatewayGroupOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetGatewayGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetGatewayGroupWithContext(ctx context.Context, input *alexaforbusiness.GetGatewayGroupInput, opts ...request.Option) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetGatewayGroupOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetGatewayGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetGatewayWithContext(ctx context.Context, input *alexaforbusiness.GetGatewayInput, opts ...request.Option) (*alexaforbusiness.GetGatewayOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetGatewayOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetGatewayWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetInvitationConfiguration(input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetInvitationConfigurationOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetInvitationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetInvitationConfigurationWithContext(ctx context.Context, input *alexaforbusiness.GetInvitationConfigurationInput, opts ...request.Option) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetInvitationConfigurationOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetInvitationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetNetworkProfile(input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetNetworkProfileOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetNetworkProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetNetworkProfileWithContext(ctx context.Context, input *alexaforbusiness.GetNetworkProfileInput, opts ...request.Option) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetNetworkProfileOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetNetworkProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetProfile(input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetProfileOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetProfileWithContext(ctx context.Context, input *alexaforbusiness.GetProfileInput, opts ...request.Option) (*alexaforbusiness.GetProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetProfileOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetRoom(input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetRoomOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetRoom(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetRoomSkillParameter(input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetRoomSkillParameterOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetRoomSkillParameter(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetRoomSkillParameterWithContext(ctx context.Context, input *alexaforbusiness.GetRoomSkillParameterInput, opts ...request.Option) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetRoomSkillParameterOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetRoomSkillParameterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetRoomWithContext(ctx context.Context, input *alexaforbusiness.GetRoomInput, opts ...request.Option) (*alexaforbusiness.GetRoomOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetRoomOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetRoomWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetSkillGroup(input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetSkillGroupOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetSkillGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) GetSkillGroupWithContext(ctx context.Context, input *alexaforbusiness.GetSkillGroupInput, opts ...request.Option) (*alexaforbusiness.GetSkillGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.GetSkillGroupOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.GetSkillGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListBusinessReportSchedules(input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListBusinessReportSchedulesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListBusinessReportSchedules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListBusinessReportSchedulesPages(input *alexaforbusiness.ListBusinessReportSchedulesInput, fn func(*alexaforbusiness.ListBusinessReportSchedulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListBusinessReportSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListBusinessReportSchedulesOutput{}
	fnCacher := func(out *alexaforbusiness.ListBusinessReportSchedulesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListBusinessReportSchedulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListBusinessReportSchedulesPagesWithContext(ctx context.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput, fn func(*alexaforbusiness.ListBusinessReportSchedulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListBusinessReportSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListBusinessReportSchedulesOutput{}
	fnCacher := func(out *alexaforbusiness.ListBusinessReportSchedulesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListBusinessReportSchedulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListBusinessReportSchedulesWithContext(ctx context.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput, opts ...request.Option) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListBusinessReportSchedulesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListBusinessReportSchedulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListConferenceProviders(input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListConferenceProvidersOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListConferenceProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListConferenceProvidersPages(input *alexaforbusiness.ListConferenceProvidersInput, fn func(*alexaforbusiness.ListConferenceProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListConferenceProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListConferenceProvidersOutput{}
	fnCacher := func(out *alexaforbusiness.ListConferenceProvidersOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListConferenceProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListConferenceProvidersPagesWithContext(ctx context.Context, input *alexaforbusiness.ListConferenceProvidersInput, fn func(*alexaforbusiness.ListConferenceProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListConferenceProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListConferenceProvidersOutput{}
	fnCacher := func(out *alexaforbusiness.ListConferenceProvidersOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListConferenceProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListConferenceProvidersWithContext(ctx context.Context, input *alexaforbusiness.ListConferenceProvidersInput, opts ...request.Option) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListConferenceProvidersOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListConferenceProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListDeviceEvents(input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListDeviceEventsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListDeviceEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListDeviceEventsPages(input *alexaforbusiness.ListDeviceEventsInput, fn func(*alexaforbusiness.ListDeviceEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListDeviceEventsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListDeviceEventsOutput{}
	fnCacher := func(out *alexaforbusiness.ListDeviceEventsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListDeviceEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListDeviceEventsPagesWithContext(ctx context.Context, input *alexaforbusiness.ListDeviceEventsInput, fn func(*alexaforbusiness.ListDeviceEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListDeviceEventsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListDeviceEventsOutput{}
	fnCacher := func(out *alexaforbusiness.ListDeviceEventsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListDeviceEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListDeviceEventsWithContext(ctx context.Context, input *alexaforbusiness.ListDeviceEventsInput, opts ...request.Option) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListDeviceEventsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListDeviceEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListGatewayGroups(input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListGatewayGroupsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListGatewayGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListGatewayGroupsPages(input *alexaforbusiness.ListGatewayGroupsInput, fn func(*alexaforbusiness.ListGatewayGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListGatewayGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListGatewayGroupsOutput{}
	fnCacher := func(out *alexaforbusiness.ListGatewayGroupsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListGatewayGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListGatewayGroupsPagesWithContext(ctx context.Context, input *alexaforbusiness.ListGatewayGroupsInput, fn func(*alexaforbusiness.ListGatewayGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListGatewayGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListGatewayGroupsOutput{}
	fnCacher := func(out *alexaforbusiness.ListGatewayGroupsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListGatewayGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListGatewayGroupsWithContext(ctx context.Context, input *alexaforbusiness.ListGatewayGroupsInput, opts ...request.Option) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListGatewayGroupsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListGatewayGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListGateways(input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListGatewaysOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListGateways(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListGatewaysPages(input *alexaforbusiness.ListGatewaysInput, fn func(*alexaforbusiness.ListGatewaysOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListGatewaysOutput{}
	fnCacher := func(out *alexaforbusiness.ListGatewaysOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListGatewaysPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListGatewaysPagesWithContext(ctx context.Context, input *alexaforbusiness.ListGatewaysInput, fn func(*alexaforbusiness.ListGatewaysOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListGatewaysOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListGatewaysOutput{}
	fnCacher := func(out *alexaforbusiness.ListGatewaysOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListGatewaysPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListGatewaysWithContext(ctx context.Context, input *alexaforbusiness.ListGatewaysInput, opts ...request.Option) (*alexaforbusiness.ListGatewaysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListGatewaysOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListGatewaysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkills(input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkills(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkillsPages(input *alexaforbusiness.ListSkillsInput, fn func(*alexaforbusiness.ListSkillsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsPagesWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsInput, fn func(*alexaforbusiness.ListSkillsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsStoreCategories(input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsStoreCategoriesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkillsStoreCategories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkillsStoreCategoriesPages(input *alexaforbusiness.ListSkillsStoreCategoriesInput, fn func(*alexaforbusiness.ListSkillsStoreCategoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsStoreCategoriesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsStoreCategoriesOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsStoreCategoriesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsStoreCategoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsStoreCategoriesPagesWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput, fn func(*alexaforbusiness.ListSkillsStoreCategoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsStoreCategoriesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsStoreCategoriesOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsStoreCategoriesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsStoreCategoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsStoreCategoriesWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput, opts ...request.Option) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsStoreCategoriesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkillsStoreCategoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkillsStoreSkillsByCategory(input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkillsStoreSkillsByCategory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkillsStoreSkillsByCategoryPages(input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, fn func(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsStoreSkillsByCategoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsStoreSkillsByCategoryPagesWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, fn func(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput{}
	fnCacher := func(out *alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSkillsStoreSkillsByCategoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSkillsStoreSkillsByCategoryWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput, opts ...request.Option) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkillsStoreSkillsByCategoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSkillsWithContext(ctx context.Context, input *alexaforbusiness.ListSkillsInput, opts ...request.Option) (*alexaforbusiness.ListSkillsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSkillsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSkillsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSmartHomeAppliances(input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSmartHomeAppliancesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSmartHomeAppliances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListSmartHomeAppliancesPages(input *alexaforbusiness.ListSmartHomeAppliancesInput, fn func(*alexaforbusiness.ListSmartHomeAppliancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSmartHomeAppliancesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSmartHomeAppliancesOutput{}
	fnCacher := func(out *alexaforbusiness.ListSmartHomeAppliancesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSmartHomeAppliancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSmartHomeAppliancesPagesWithContext(ctx context.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput, fn func(*alexaforbusiness.ListSmartHomeAppliancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListSmartHomeAppliancesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListSmartHomeAppliancesOutput{}
	fnCacher := func(out *alexaforbusiness.ListSmartHomeAppliancesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListSmartHomeAppliancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListSmartHomeAppliancesWithContext(ctx context.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput, opts ...request.Option) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListSmartHomeAppliancesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListSmartHomeAppliancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListTags(input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListTagsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) ListTagsPages(input *alexaforbusiness.ListTagsInput, fn func(*alexaforbusiness.ListTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListTagsOutput{}
	fnCacher := func(out *alexaforbusiness.ListTagsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListTagsPagesWithContext(ctx context.Context, input *alexaforbusiness.ListTagsInput, fn func(*alexaforbusiness.ListTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.ListTagsOutput{}
	fnCacher := func(out *alexaforbusiness.ListTagsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.ListTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) ListTagsWithContext(ctx context.Context, input *alexaforbusiness.ListTagsInput, opts ...request.Option) (*alexaforbusiness.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.ListTagsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchAddressBooks(input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchAddressBooksOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchAddressBooks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchAddressBooksPages(input *alexaforbusiness.SearchAddressBooksInput, fn func(*alexaforbusiness.SearchAddressBooksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchAddressBooksOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchAddressBooksOutput{}
	fnCacher := func(out *alexaforbusiness.SearchAddressBooksOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchAddressBooksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchAddressBooksPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchAddressBooksInput, fn func(*alexaforbusiness.SearchAddressBooksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchAddressBooksOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchAddressBooksOutput{}
	fnCacher := func(out *alexaforbusiness.SearchAddressBooksOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchAddressBooksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchAddressBooksWithContext(ctx context.Context, input *alexaforbusiness.SearchAddressBooksInput, opts ...request.Option) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchAddressBooksOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchAddressBooksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchContacts(input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchContactsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchContacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchContactsPages(input *alexaforbusiness.SearchContactsInput, fn func(*alexaforbusiness.SearchContactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchContactsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchContactsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchContactsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchContactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchContactsPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchContactsInput, fn func(*alexaforbusiness.SearchContactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchContactsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchContactsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchContactsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchContactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchContactsWithContext(ctx context.Context, input *alexaforbusiness.SearchContactsInput, opts ...request.Option) (*alexaforbusiness.SearchContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchContactsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchContactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchDevices(input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchDevicesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchDevicesPages(input *alexaforbusiness.SearchDevicesInput, fn func(*alexaforbusiness.SearchDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchDevicesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchDevicesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchDevicesPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchDevicesInput, fn func(*alexaforbusiness.SearchDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchDevicesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchDevicesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchDevicesWithContext(ctx context.Context, input *alexaforbusiness.SearchDevicesInput, opts ...request.Option) (*alexaforbusiness.SearchDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchDevicesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchNetworkProfiles(input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchNetworkProfilesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchNetworkProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchNetworkProfilesPages(input *alexaforbusiness.SearchNetworkProfilesInput, fn func(*alexaforbusiness.SearchNetworkProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchNetworkProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchNetworkProfilesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchNetworkProfilesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchNetworkProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchNetworkProfilesPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchNetworkProfilesInput, fn func(*alexaforbusiness.SearchNetworkProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchNetworkProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchNetworkProfilesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchNetworkProfilesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchNetworkProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchNetworkProfilesWithContext(ctx context.Context, input *alexaforbusiness.SearchNetworkProfilesInput, opts ...request.Option) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchNetworkProfilesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchNetworkProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchProfiles(input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchProfilesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchProfilesPages(input *alexaforbusiness.SearchProfilesInput, fn func(*alexaforbusiness.SearchProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchProfilesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchProfilesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchProfilesPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchProfilesInput, fn func(*alexaforbusiness.SearchProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchProfilesOutput{}
	fnCacher := func(out *alexaforbusiness.SearchProfilesOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchProfilesWithContext(ctx context.Context, input *alexaforbusiness.SearchProfilesInput, opts ...request.Option) (*alexaforbusiness.SearchProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchProfilesOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchRooms(input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchRoomsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchRooms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchRoomsPages(input *alexaforbusiness.SearchRoomsInput, fn func(*alexaforbusiness.SearchRoomsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchRoomsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchRoomsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchRoomsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchRoomsPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchRoomsInput, fn func(*alexaforbusiness.SearchRoomsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchRoomsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchRoomsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchRoomsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchRoomsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchRoomsWithContext(ctx context.Context, input *alexaforbusiness.SearchRoomsInput, opts ...request.Option) (*alexaforbusiness.SearchRoomsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchRoomsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchRoomsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchSkillGroups(input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchSkillGroupsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchSkillGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchSkillGroupsPages(input *alexaforbusiness.SearchSkillGroupsInput, fn func(*alexaforbusiness.SearchSkillGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchSkillGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchSkillGroupsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchSkillGroupsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchSkillGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchSkillGroupsPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchSkillGroupsInput, fn func(*alexaforbusiness.SearchSkillGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchSkillGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchSkillGroupsOutput{}
	fnCacher := func(out *alexaforbusiness.SearchSkillGroupsOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchSkillGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchSkillGroupsWithContext(ctx context.Context, input *alexaforbusiness.SearchSkillGroupsInput, opts ...request.Option) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchSkillGroupsOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchSkillGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchUsers(input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchUsersOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchUsers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AlexaForBusiness) SearchUsersPages(input *alexaforbusiness.SearchUsersInput, fn func(*alexaforbusiness.SearchUsersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchUsersOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchUsersOutput{}
	fnCacher := func(out *alexaforbusiness.SearchUsersOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchUsersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchUsersPagesWithContext(ctx context.Context, input *alexaforbusiness.SearchUsersInput, fn func(*alexaforbusiness.SearchUsersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*alexaforbusiness.SearchUsersOutput), false)
		return nil
	}
	cachable := true
	output := &alexaforbusiness.SearchUsersOutput{}
	fnCacher := func(out *alexaforbusiness.SearchUsersOutput, more bool) bool {
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
	if err := c.AlexaForBusinessAPI.SearchUsersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AlexaForBusiness) SearchUsersWithContext(ctx context.Context, input *alexaforbusiness.SearchUsersInput, opts ...request.Option) (*alexaforbusiness.SearchUsersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*alexaforbusiness.SearchUsersOutput), nil
	}
	out, err := c.AlexaForBusinessAPI.SearchUsersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
