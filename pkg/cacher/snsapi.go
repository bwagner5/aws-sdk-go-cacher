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
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SNS struct {
	snsiface.SNSAPI
	cache *cache.Cache
}

func NewSNS(snsapi snsiface.SNSAPI) *SNS {
	return &SNS{
		SNSAPI: snsapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SNS) GetDataProtectionPolicy(input *sns.GetDataProtectionPolicyInput) (*sns.GetDataProtectionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetDataProtectionPolicyOutput), nil
	}
	out, err := c.SNSAPI.GetDataProtectionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetDataProtectionPolicyWithContext(ctx context.Context, input *sns.GetDataProtectionPolicyInput, opts ...request.Option) (*sns.GetDataProtectionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetDataProtectionPolicyOutput), nil
	}
	out, err := c.SNSAPI.GetDataProtectionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetEndpointAttributes(input *sns.GetEndpointAttributesInput) (*sns.GetEndpointAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetEndpointAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetEndpointAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetEndpointAttributesWithContext(ctx context.Context, input *sns.GetEndpointAttributesInput, opts ...request.Option) (*sns.GetEndpointAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetEndpointAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetEndpointAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetPlatformApplicationAttributes(input *sns.GetPlatformApplicationAttributesInput) (*sns.GetPlatformApplicationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetPlatformApplicationAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetPlatformApplicationAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetPlatformApplicationAttributesWithContext(ctx context.Context, input *sns.GetPlatformApplicationAttributesInput, opts ...request.Option) (*sns.GetPlatformApplicationAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetPlatformApplicationAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetPlatformApplicationAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSMSAttributes(input *sns.GetSMSAttributesInput) (*sns.GetSMSAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSMSAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetSMSAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSMSAttributesWithContext(ctx context.Context, input *sns.GetSMSAttributesInput, opts ...request.Option) (*sns.GetSMSAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSMSAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetSMSAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSMSSandboxAccountStatus(input *sns.GetSMSSandboxAccountStatusInput) (*sns.GetSMSSandboxAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSMSSandboxAccountStatusOutput), nil
	}
	out, err := c.SNSAPI.GetSMSSandboxAccountStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSMSSandboxAccountStatusWithContext(ctx context.Context, input *sns.GetSMSSandboxAccountStatusInput, opts ...request.Option) (*sns.GetSMSSandboxAccountStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSMSSandboxAccountStatusOutput), nil
	}
	out, err := c.SNSAPI.GetSMSSandboxAccountStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSubscriptionAttributes(input *sns.GetSubscriptionAttributesInput) (*sns.GetSubscriptionAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSubscriptionAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetSubscriptionAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetSubscriptionAttributesWithContext(ctx context.Context, input *sns.GetSubscriptionAttributesInput, opts ...request.Option) (*sns.GetSubscriptionAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetSubscriptionAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetSubscriptionAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetTopicAttributes(input *sns.GetTopicAttributesInput) (*sns.GetTopicAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetTopicAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetTopicAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) GetTopicAttributesWithContext(ctx context.Context, input *sns.GetTopicAttributesInput, opts ...request.Option) (*sns.GetTopicAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.GetTopicAttributesOutput), nil
	}
	out, err := c.SNSAPI.GetTopicAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListEndpointsByPlatformApplication(input *sns.ListEndpointsByPlatformApplicationInput) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListEndpointsByPlatformApplicationOutput), nil
	}
	out, err := c.SNSAPI.ListEndpointsByPlatformApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListEndpointsByPlatformApplicationPages(input *sns.ListEndpointsByPlatformApplicationInput, fn func(*sns.ListEndpointsByPlatformApplicationOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListEndpointsByPlatformApplicationOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListEndpointsByPlatformApplicationOutput{}
	fnCacher := func(out *sns.ListEndpointsByPlatformApplicationOutput, more bool) bool {
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
	if err := c.SNSAPI.ListEndpointsByPlatformApplicationPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListEndpointsByPlatformApplicationPagesWithContext(ctx context.Context, input *sns.ListEndpointsByPlatformApplicationInput, fn func(*sns.ListEndpointsByPlatformApplicationOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListEndpointsByPlatformApplicationOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListEndpointsByPlatformApplicationOutput{}
	fnCacher := func(out *sns.ListEndpointsByPlatformApplicationOutput, more bool) bool {
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
	if err := c.SNSAPI.ListEndpointsByPlatformApplicationPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListEndpointsByPlatformApplicationWithContext(ctx context.Context, input *sns.ListEndpointsByPlatformApplicationInput, opts ...request.Option) (*sns.ListEndpointsByPlatformApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListEndpointsByPlatformApplicationOutput), nil
	}
	out, err := c.SNSAPI.ListEndpointsByPlatformApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListOriginationNumbers(input *sns.ListOriginationNumbersInput) (*sns.ListOriginationNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListOriginationNumbersOutput), nil
	}
	out, err := c.SNSAPI.ListOriginationNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListOriginationNumbersPages(input *sns.ListOriginationNumbersInput, fn func(*sns.ListOriginationNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListOriginationNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListOriginationNumbersOutput{}
	fnCacher := func(out *sns.ListOriginationNumbersOutput, more bool) bool {
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
	if err := c.SNSAPI.ListOriginationNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListOriginationNumbersPagesWithContext(ctx context.Context, input *sns.ListOriginationNumbersInput, fn func(*sns.ListOriginationNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListOriginationNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListOriginationNumbersOutput{}
	fnCacher := func(out *sns.ListOriginationNumbersOutput, more bool) bool {
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
	if err := c.SNSAPI.ListOriginationNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListOriginationNumbersWithContext(ctx context.Context, input *sns.ListOriginationNumbersInput, opts ...request.Option) (*sns.ListOriginationNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListOriginationNumbersOutput), nil
	}
	out, err := c.SNSAPI.ListOriginationNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListPhoneNumbersOptedOut(input *sns.ListPhoneNumbersOptedOutInput) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListPhoneNumbersOptedOutOutput), nil
	}
	out, err := c.SNSAPI.ListPhoneNumbersOptedOut(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListPhoneNumbersOptedOutPages(input *sns.ListPhoneNumbersOptedOutInput, fn func(*sns.ListPhoneNumbersOptedOutOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListPhoneNumbersOptedOutOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListPhoneNumbersOptedOutOutput{}
	fnCacher := func(out *sns.ListPhoneNumbersOptedOutOutput, more bool) bool {
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
	if err := c.SNSAPI.ListPhoneNumbersOptedOutPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListPhoneNumbersOptedOutPagesWithContext(ctx context.Context, input *sns.ListPhoneNumbersOptedOutInput, fn func(*sns.ListPhoneNumbersOptedOutOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListPhoneNumbersOptedOutOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListPhoneNumbersOptedOutOutput{}
	fnCacher := func(out *sns.ListPhoneNumbersOptedOutOutput, more bool) bool {
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
	if err := c.SNSAPI.ListPhoneNumbersOptedOutPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListPhoneNumbersOptedOutWithContext(ctx context.Context, input *sns.ListPhoneNumbersOptedOutInput, opts ...request.Option) (*sns.ListPhoneNumbersOptedOutOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListPhoneNumbersOptedOutOutput), nil
	}
	out, err := c.SNSAPI.ListPhoneNumbersOptedOutWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListPlatformApplications(input *sns.ListPlatformApplicationsInput) (*sns.ListPlatformApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListPlatformApplicationsOutput), nil
	}
	out, err := c.SNSAPI.ListPlatformApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListPlatformApplicationsPages(input *sns.ListPlatformApplicationsInput, fn func(*sns.ListPlatformApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListPlatformApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListPlatformApplicationsOutput{}
	fnCacher := func(out *sns.ListPlatformApplicationsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListPlatformApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListPlatformApplicationsPagesWithContext(ctx context.Context, input *sns.ListPlatformApplicationsInput, fn func(*sns.ListPlatformApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListPlatformApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListPlatformApplicationsOutput{}
	fnCacher := func(out *sns.ListPlatformApplicationsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListPlatformApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListPlatformApplicationsWithContext(ctx context.Context, input *sns.ListPlatformApplicationsInput, opts ...request.Option) (*sns.ListPlatformApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListPlatformApplicationsOutput), nil
	}
	out, err := c.SNSAPI.ListPlatformApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSMSSandboxPhoneNumbers(input *sns.ListSMSSandboxPhoneNumbersInput) (*sns.ListSMSSandboxPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSMSSandboxPhoneNumbersOutput), nil
	}
	out, err := c.SNSAPI.ListSMSSandboxPhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSMSSandboxPhoneNumbersPages(input *sns.ListSMSSandboxPhoneNumbersInput, fn func(*sns.ListSMSSandboxPhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSMSSandboxPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSMSSandboxPhoneNumbersOutput{}
	fnCacher := func(out *sns.ListSMSSandboxPhoneNumbersOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSMSSandboxPhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSMSSandboxPhoneNumbersPagesWithContext(ctx context.Context, input *sns.ListSMSSandboxPhoneNumbersInput, fn func(*sns.ListSMSSandboxPhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSMSSandboxPhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSMSSandboxPhoneNumbersOutput{}
	fnCacher := func(out *sns.ListSMSSandboxPhoneNumbersOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSMSSandboxPhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSMSSandboxPhoneNumbersWithContext(ctx context.Context, input *sns.ListSMSSandboxPhoneNumbersInput, opts ...request.Option) (*sns.ListSMSSandboxPhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSMSSandboxPhoneNumbersOutput), nil
	}
	out, err := c.SNSAPI.ListSMSSandboxPhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSubscriptions(input *sns.ListSubscriptionsInput) (*sns.ListSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSubscriptionsOutput), nil
	}
	out, err := c.SNSAPI.ListSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSubscriptionsByTopic(input *sns.ListSubscriptionsByTopicInput) (*sns.ListSubscriptionsByTopicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSubscriptionsByTopicOutput), nil
	}
	out, err := c.SNSAPI.ListSubscriptionsByTopic(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSubscriptionsByTopicPages(input *sns.ListSubscriptionsByTopicInput, fn func(*sns.ListSubscriptionsByTopicOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSubscriptionsByTopicOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSubscriptionsByTopicOutput{}
	fnCacher := func(out *sns.ListSubscriptionsByTopicOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSubscriptionsByTopicPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSubscriptionsByTopicPagesWithContext(ctx context.Context, input *sns.ListSubscriptionsByTopicInput, fn func(*sns.ListSubscriptionsByTopicOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSubscriptionsByTopicOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSubscriptionsByTopicOutput{}
	fnCacher := func(out *sns.ListSubscriptionsByTopicOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSubscriptionsByTopicPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSubscriptionsByTopicWithContext(ctx context.Context, input *sns.ListSubscriptionsByTopicInput, opts ...request.Option) (*sns.ListSubscriptionsByTopicOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSubscriptionsByTopicOutput), nil
	}
	out, err := c.SNSAPI.ListSubscriptionsByTopicWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListSubscriptionsPages(input *sns.ListSubscriptionsInput, fn func(*sns.ListSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSubscriptionsOutput{}
	fnCacher := func(out *sns.ListSubscriptionsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSubscriptionsPagesWithContext(ctx context.Context, input *sns.ListSubscriptionsInput, fn func(*sns.ListSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListSubscriptionsOutput{}
	fnCacher := func(out *sns.ListSubscriptionsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListSubscriptionsWithContext(ctx context.Context, input *sns.ListSubscriptionsInput, opts ...request.Option) (*sns.ListSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListSubscriptionsOutput), nil
	}
	out, err := c.SNSAPI.ListSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListTagsForResource(input *sns.ListTagsForResourceInput) (*sns.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListTagsForResourceOutput), nil
	}
	out, err := c.SNSAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListTagsForResourceWithContext(ctx context.Context, input *sns.ListTagsForResourceInput, opts ...request.Option) (*sns.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListTagsForResourceOutput), nil
	}
	out, err := c.SNSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListTopics(input *sns.ListTopicsInput) (*sns.ListTopicsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListTopicsOutput), nil
	}
	out, err := c.SNSAPI.ListTopics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SNS) ListTopicsPages(input *sns.ListTopicsInput, fn func(*sns.ListTopicsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListTopicsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListTopicsOutput{}
	fnCacher := func(out *sns.ListTopicsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListTopicsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListTopicsPagesWithContext(ctx context.Context, input *sns.ListTopicsInput, fn func(*sns.ListTopicsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sns.ListTopicsOutput), false)
		return nil
	}
	cachable := true
	output := &sns.ListTopicsOutput{}
	fnCacher := func(out *sns.ListTopicsOutput, more bool) bool {
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
	if err := c.SNSAPI.ListTopicsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SNS) ListTopicsWithContext(ctx context.Context, input *sns.ListTopicsInput, opts ...request.Option) (*sns.ListTopicsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sns.ListTopicsOutput), nil
	}
	out, err := c.SNSAPI.ListTopicsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
