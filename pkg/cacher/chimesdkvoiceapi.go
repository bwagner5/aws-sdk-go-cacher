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
	"github.com/aws/aws-sdk-go/service/chimesdkvoice"
	"github.com/aws/aws-sdk-go/service/chimesdkvoice/chimesdkvoiceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ChimeSDKVoice struct {
	chimesdkvoiceiface.ChimeSDKVoiceAPI
	cache *cache.Cache
}

func NewChimeSDKVoice(chimesdkvoiceapi chimesdkvoiceiface.ChimeSDKVoiceAPI) *ChimeSDKVoice {
	return &ChimeSDKVoice{
		ChimeSDKVoiceAPI: chimesdkvoiceapi,
		cache:            cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ChimeSDKVoice) BatchDeletePhoneNumber(input *chimesdkvoice.BatchDeletePhoneNumberInput) (*chimesdkvoice.BatchDeletePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.BatchDeletePhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.BatchDeletePhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) BatchDeletePhoneNumberWithContext(ctx context.Context, input *chimesdkvoice.BatchDeletePhoneNumberInput, opts ...request.Option) (*chimesdkvoice.BatchDeletePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.BatchDeletePhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.BatchDeletePhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) BatchUpdatePhoneNumber(input *chimesdkvoice.BatchUpdatePhoneNumberInput) (*chimesdkvoice.BatchUpdatePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.BatchUpdatePhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.BatchUpdatePhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) BatchUpdatePhoneNumberWithContext(ctx context.Context, input *chimesdkvoice.BatchUpdatePhoneNumberInput, opts ...request.Option) (*chimesdkvoice.BatchUpdatePhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.BatchUpdatePhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.BatchUpdatePhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetGlobalSettings(input *chimesdkvoice.GetGlobalSettingsInput) (*chimesdkvoice.GetGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetGlobalSettingsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetGlobalSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetGlobalSettingsWithContext(ctx context.Context, input *chimesdkvoice.GetGlobalSettingsInput, opts ...request.Option) (*chimesdkvoice.GetGlobalSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetGlobalSettingsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetGlobalSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumber(input *chimesdkvoice.GetPhoneNumberInput) (*chimesdkvoice.GetPhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumberOrder(input *chimesdkvoice.GetPhoneNumberOrderInput) (*chimesdkvoice.GetPhoneNumberOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberOrderOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumberOrder(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumberOrderWithContext(ctx context.Context, input *chimesdkvoice.GetPhoneNumberOrderInput, opts ...request.Option) (*chimesdkvoice.GetPhoneNumberOrderOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberOrderOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumberOrderWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumberSettings(input *chimesdkvoice.GetPhoneNumberSettingsInput) (*chimesdkvoice.GetPhoneNumberSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberSettingsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumberSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumberSettingsWithContext(ctx context.Context, input *chimesdkvoice.GetPhoneNumberSettingsInput, opts ...request.Option) (*chimesdkvoice.GetPhoneNumberSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberSettingsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumberSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetPhoneNumberWithContext(ctx context.Context, input *chimesdkvoice.GetPhoneNumberInput, opts ...request.Option) (*chimesdkvoice.GetPhoneNumberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetPhoneNumberOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetPhoneNumberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetProxySession(input *chimesdkvoice.GetProxySessionInput) (*chimesdkvoice.GetProxySessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetProxySessionOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetProxySession(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetProxySessionWithContext(ctx context.Context, input *chimesdkvoice.GetProxySessionInput, opts ...request.Option) (*chimesdkvoice.GetProxySessionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetProxySessionOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetProxySessionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplication(input *chimesdkvoice.GetSipMediaApplicationInput) (*chimesdkvoice.GetSipMediaApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplicationAlexaSkillConfiguration(input *chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationInput) (*chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplicationAlexaSkillConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplicationAlexaSkillConfigurationWithContext(ctx context.Context, input *chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationInput, opts ...request.Option) (*chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationAlexaSkillConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplicationAlexaSkillConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplicationLoggingConfiguration(input *chimesdkvoice.GetSipMediaApplicationLoggingConfigurationInput) (*chimesdkvoice.GetSipMediaApplicationLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplicationLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplicationLoggingConfigurationWithContext(ctx context.Context, input *chimesdkvoice.GetSipMediaApplicationLoggingConfigurationInput, opts ...request.Option) (*chimesdkvoice.GetSipMediaApplicationLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplicationLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipMediaApplicationWithContext(ctx context.Context, input *chimesdkvoice.GetSipMediaApplicationInput, opts ...request.Option) (*chimesdkvoice.GetSipMediaApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipMediaApplicationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipMediaApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipRule(input *chimesdkvoice.GetSipRuleInput) (*chimesdkvoice.GetSipRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipRuleOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipRule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetSipRuleWithContext(ctx context.Context, input *chimesdkvoice.GetSipRuleInput, opts ...request.Option) (*chimesdkvoice.GetSipRuleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetSipRuleOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetSipRuleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnector(input *chimesdkvoice.GetVoiceConnectorInput) (*chimesdkvoice.GetVoiceConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorEmergencyCallingConfiguration(input *chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationInput) (*chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorEmergencyCallingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorEmergencyCallingConfigurationWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorEmergencyCallingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorEmergencyCallingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorGroup(input *chimesdkvoice.GetVoiceConnectorGroupInput) (*chimesdkvoice.GetVoiceConnectorGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorGroupOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorGroupWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorGroupInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorGroupOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorLoggingConfiguration(input *chimesdkvoice.GetVoiceConnectorLoggingConfigurationInput) (*chimesdkvoice.GetVoiceConnectorLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorLoggingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorLoggingConfigurationWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorLoggingConfigurationInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorLoggingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorLoggingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorLoggingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorOrigination(input *chimesdkvoice.GetVoiceConnectorOriginationInput) (*chimesdkvoice.GetVoiceConnectorOriginationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorOriginationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorOrigination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorOriginationWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorOriginationInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorOriginationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorOriginationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorOriginationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorProxy(input *chimesdkvoice.GetVoiceConnectorProxyInput) (*chimesdkvoice.GetVoiceConnectorProxyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorProxyOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorProxy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorProxyWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorProxyInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorProxyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorProxyOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorProxyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorStreamingConfiguration(input *chimesdkvoice.GetVoiceConnectorStreamingConfigurationInput) (*chimesdkvoice.GetVoiceConnectorStreamingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorStreamingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorStreamingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorStreamingConfigurationWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorStreamingConfigurationInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorStreamingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorStreamingConfigurationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorStreamingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorTermination(input *chimesdkvoice.GetVoiceConnectorTerminationInput) (*chimesdkvoice.GetVoiceConnectorTerminationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorTerminationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorTermination(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorTerminationHealth(input *chimesdkvoice.GetVoiceConnectorTerminationHealthInput) (*chimesdkvoice.GetVoiceConnectorTerminationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorTerminationHealthOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorTerminationHealth(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorTerminationHealthWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorTerminationHealthInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorTerminationHealthOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorTerminationHealthOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorTerminationHealthWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorTerminationWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorTerminationInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorTerminationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorTerminationOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorTerminationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) GetVoiceConnectorWithContext(ctx context.Context, input *chimesdkvoice.GetVoiceConnectorInput, opts ...request.Option) (*chimesdkvoice.GetVoiceConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.GetVoiceConnectorOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.GetVoiceConnectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListAvailableVoiceConnectorRegions(input *chimesdkvoice.ListAvailableVoiceConnectorRegionsInput) (*chimesdkvoice.ListAvailableVoiceConnectorRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListAvailableVoiceConnectorRegionsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListAvailableVoiceConnectorRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListAvailableVoiceConnectorRegionsWithContext(ctx context.Context, input *chimesdkvoice.ListAvailableVoiceConnectorRegionsInput, opts ...request.Option) (*chimesdkvoice.ListAvailableVoiceConnectorRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListAvailableVoiceConnectorRegionsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListAvailableVoiceConnectorRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListPhoneNumberOrders(input *chimesdkvoice.ListPhoneNumberOrdersInput) (*chimesdkvoice.ListPhoneNumberOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListPhoneNumberOrdersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListPhoneNumberOrders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListPhoneNumberOrdersPages(input *chimesdkvoice.ListPhoneNumberOrdersInput, fn func(*chimesdkvoice.ListPhoneNumberOrdersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListPhoneNumberOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListPhoneNumberOrdersOutput{}
	fnCacher := func(out *chimesdkvoice.ListPhoneNumberOrdersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListPhoneNumberOrdersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListPhoneNumberOrdersPagesWithContext(ctx context.Context, input *chimesdkvoice.ListPhoneNumberOrdersInput, fn func(*chimesdkvoice.ListPhoneNumberOrdersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListPhoneNumberOrdersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListPhoneNumberOrdersOutput{}
	fnCacher := func(out *chimesdkvoice.ListPhoneNumberOrdersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListPhoneNumberOrdersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListPhoneNumberOrdersWithContext(ctx context.Context, input *chimesdkvoice.ListPhoneNumberOrdersInput, opts ...request.Option) (*chimesdkvoice.ListPhoneNumberOrdersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListPhoneNumberOrdersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListPhoneNumberOrdersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListPhoneNumbers(input *chimesdkvoice.ListPhoneNumbersInput) (*chimesdkvoice.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListPhoneNumbersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListPhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListPhoneNumbersPages(input *chimesdkvoice.ListPhoneNumbersInput, fn func(*chimesdkvoice.ListPhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListPhoneNumbersOutput{}
	fnCacher := func(out *chimesdkvoice.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListPhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListPhoneNumbersPagesWithContext(ctx context.Context, input *chimesdkvoice.ListPhoneNumbersInput, fn func(*chimesdkvoice.ListPhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListPhoneNumbersOutput{}
	fnCacher := func(out *chimesdkvoice.ListPhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListPhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListPhoneNumbersWithContext(ctx context.Context, input *chimesdkvoice.ListPhoneNumbersInput, opts ...request.Option) (*chimesdkvoice.ListPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListPhoneNumbersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListPhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListProxySessions(input *chimesdkvoice.ListProxySessionsInput) (*chimesdkvoice.ListProxySessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListProxySessionsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListProxySessions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListProxySessionsPages(input *chimesdkvoice.ListProxySessionsInput, fn func(*chimesdkvoice.ListProxySessionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListProxySessionsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListProxySessionsOutput{}
	fnCacher := func(out *chimesdkvoice.ListProxySessionsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListProxySessionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListProxySessionsPagesWithContext(ctx context.Context, input *chimesdkvoice.ListProxySessionsInput, fn func(*chimesdkvoice.ListProxySessionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListProxySessionsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListProxySessionsOutput{}
	fnCacher := func(out *chimesdkvoice.ListProxySessionsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListProxySessionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListProxySessionsWithContext(ctx context.Context, input *chimesdkvoice.ListProxySessionsInput, opts ...request.Option) (*chimesdkvoice.ListProxySessionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListProxySessionsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListProxySessionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSipMediaApplications(input *chimesdkvoice.ListSipMediaApplicationsInput) (*chimesdkvoice.ListSipMediaApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSipMediaApplicationsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSipMediaApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSipMediaApplicationsPages(input *chimesdkvoice.ListSipMediaApplicationsInput, fn func(*chimesdkvoice.ListSipMediaApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListSipMediaApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListSipMediaApplicationsOutput{}
	fnCacher := func(out *chimesdkvoice.ListSipMediaApplicationsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListSipMediaApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListSipMediaApplicationsPagesWithContext(ctx context.Context, input *chimesdkvoice.ListSipMediaApplicationsInput, fn func(*chimesdkvoice.ListSipMediaApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListSipMediaApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListSipMediaApplicationsOutput{}
	fnCacher := func(out *chimesdkvoice.ListSipMediaApplicationsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListSipMediaApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListSipMediaApplicationsWithContext(ctx context.Context, input *chimesdkvoice.ListSipMediaApplicationsInput, opts ...request.Option) (*chimesdkvoice.ListSipMediaApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSipMediaApplicationsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSipMediaApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSipRules(input *chimesdkvoice.ListSipRulesInput) (*chimesdkvoice.ListSipRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSipRulesOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSipRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSipRulesPages(input *chimesdkvoice.ListSipRulesInput, fn func(*chimesdkvoice.ListSipRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListSipRulesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListSipRulesOutput{}
	fnCacher := func(out *chimesdkvoice.ListSipRulesOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListSipRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListSipRulesPagesWithContext(ctx context.Context, input *chimesdkvoice.ListSipRulesInput, fn func(*chimesdkvoice.ListSipRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListSipRulesOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListSipRulesOutput{}
	fnCacher := func(out *chimesdkvoice.ListSipRulesOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListSipRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListSipRulesWithContext(ctx context.Context, input *chimesdkvoice.ListSipRulesInput, opts ...request.Option) (*chimesdkvoice.ListSipRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSipRulesOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSipRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSupportedPhoneNumberCountries(input *chimesdkvoice.ListSupportedPhoneNumberCountriesInput) (*chimesdkvoice.ListSupportedPhoneNumberCountriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSupportedPhoneNumberCountriesOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSupportedPhoneNumberCountries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListSupportedPhoneNumberCountriesWithContext(ctx context.Context, input *chimesdkvoice.ListSupportedPhoneNumberCountriesInput, opts ...request.Option) (*chimesdkvoice.ListSupportedPhoneNumberCountriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListSupportedPhoneNumberCountriesOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListSupportedPhoneNumberCountriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectorGroups(input *chimesdkvoice.ListVoiceConnectorGroupsInput) (*chimesdkvoice.ListVoiceConnectorGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorGroupsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectorGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectorGroupsPages(input *chimesdkvoice.ListVoiceConnectorGroupsInput, fn func(*chimesdkvoice.ListVoiceConnectorGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListVoiceConnectorGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListVoiceConnectorGroupsOutput{}
	fnCacher := func(out *chimesdkvoice.ListVoiceConnectorGroupsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListVoiceConnectorGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListVoiceConnectorGroupsPagesWithContext(ctx context.Context, input *chimesdkvoice.ListVoiceConnectorGroupsInput, fn func(*chimesdkvoice.ListVoiceConnectorGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListVoiceConnectorGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListVoiceConnectorGroupsOutput{}
	fnCacher := func(out *chimesdkvoice.ListVoiceConnectorGroupsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListVoiceConnectorGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListVoiceConnectorGroupsWithContext(ctx context.Context, input *chimesdkvoice.ListVoiceConnectorGroupsInput, opts ...request.Option) (*chimesdkvoice.ListVoiceConnectorGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorGroupsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectorGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectorTerminationCredentials(input *chimesdkvoice.ListVoiceConnectorTerminationCredentialsInput) (*chimesdkvoice.ListVoiceConnectorTerminationCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorTerminationCredentialsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectorTerminationCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectorTerminationCredentialsWithContext(ctx context.Context, input *chimesdkvoice.ListVoiceConnectorTerminationCredentialsInput, opts ...request.Option) (*chimesdkvoice.ListVoiceConnectorTerminationCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorTerminationCredentialsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectorTerminationCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectors(input *chimesdkvoice.ListVoiceConnectorsInput) (*chimesdkvoice.ListVoiceConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) ListVoiceConnectorsPages(input *chimesdkvoice.ListVoiceConnectorsInput, fn func(*chimesdkvoice.ListVoiceConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListVoiceConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListVoiceConnectorsOutput{}
	fnCacher := func(out *chimesdkvoice.ListVoiceConnectorsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListVoiceConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListVoiceConnectorsPagesWithContext(ctx context.Context, input *chimesdkvoice.ListVoiceConnectorsInput, fn func(*chimesdkvoice.ListVoiceConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.ListVoiceConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.ListVoiceConnectorsOutput{}
	fnCacher := func(out *chimesdkvoice.ListVoiceConnectorsOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.ListVoiceConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) ListVoiceConnectorsWithContext(ctx context.Context, input *chimesdkvoice.ListVoiceConnectorsInput, opts ...request.Option) (*chimesdkvoice.ListVoiceConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.ListVoiceConnectorsOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.ListVoiceConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) SearchAvailablePhoneNumbers(input *chimesdkvoice.SearchAvailablePhoneNumbersInput) (*chimesdkvoice.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.SearchAvailablePhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ChimeSDKVoice) SearchAvailablePhoneNumbersPages(input *chimesdkvoice.SearchAvailablePhoneNumbersInput, fn func(*chimesdkvoice.SearchAvailablePhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *chimesdkvoice.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.SearchAvailablePhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) SearchAvailablePhoneNumbersPagesWithContext(ctx context.Context, input *chimesdkvoice.SearchAvailablePhoneNumbersInput, fn func(*chimesdkvoice.SearchAvailablePhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*chimesdkvoice.SearchAvailablePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &chimesdkvoice.SearchAvailablePhoneNumbersOutput{}
	fnCacher := func(out *chimesdkvoice.SearchAvailablePhoneNumbersOutput, more bool) bool {
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
	if err := c.ChimeSDKVoiceAPI.SearchAvailablePhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ChimeSDKVoice) SearchAvailablePhoneNumbersWithContext(ctx context.Context, input *chimesdkvoice.SearchAvailablePhoneNumbersInput, opts ...request.Option) (*chimesdkvoice.SearchAvailablePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*chimesdkvoice.SearchAvailablePhoneNumbersOutput), nil
	}
	out, err := c.ChimeSDKVoiceAPI.SearchAvailablePhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
