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
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice/pinpointsmsvoiceiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PinpointSMSVoice struct {
	pinpointsmsvoiceiface.PinpointSMSVoiceAPI
	cache *cache.Cache
}

func NewPinpointSMSVoice(pinpointsmsvoiceapi pinpointsmsvoiceiface.PinpointSMSVoiceAPI) *PinpointSMSVoice {
	return &PinpointSMSVoice{
		PinpointSMSVoiceAPI: pinpointsmsvoiceapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PinpointSMSVoice) GetConfigurationSetEventDestinations(input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.PinpointSMSVoiceAPI.GetConfigurationSetEventDestinations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoice) GetConfigurationSetEventDestinationsWithContext(ctx context.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput, opts ...request.Option) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput), nil
	}
	out, err := c.PinpointSMSVoiceAPI.GetConfigurationSetEventDestinationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoice) ListConfigurationSets(input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoice.ListConfigurationSetsOutput), nil
	}
	out, err := c.PinpointSMSVoiceAPI.ListConfigurationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoice) ListConfigurationSetsWithContext(ctx context.Context, input *pinpointsmsvoice.ListConfigurationSetsInput, opts ...request.Option) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoice.ListConfigurationSetsOutput), nil
	}
	out, err := c.PinpointSMSVoiceAPI.ListConfigurationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
