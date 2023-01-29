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
package pinpointsmsvoicev2cacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoicev2"
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoicev2/pinpointsmsvoicev2iface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type PinpointSMSVoiceV2 struct {
	pinpointsmsvoicev2iface.PinpointSMSVoiceV2API
	cache *cache.Cache
}

func New(pinpointsmsvoicev2api pinpointsmsvoicev2iface.PinpointSMSVoiceV2API) *PinpointSMSVoiceV2 {
	return &PinpointSMSVoiceV2{
		PinpointSMSVoiceV2API: pinpointsmsvoicev2api,
		cache:                 cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *PinpointSMSVoiceV2) DescribeAccountAttributes(input *pinpointsmsvoicev2.DescribeAccountAttributesInput) (*pinpointsmsvoicev2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeAccountAttributes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeAccountAttributesPages(input *pinpointsmsvoicev2.DescribeAccountAttributesInput, fn func(*pinpointsmsvoicev2.DescribeAccountAttributesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeAccountAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeAccountAttributesOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeAccountAttributesOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeAccountAttributesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeAccountAttributesPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeAccountAttributesInput, fn func(*pinpointsmsvoicev2.DescribeAccountAttributesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeAccountAttributesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeAccountAttributesOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeAccountAttributesOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeAccountAttributesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeAccountAttributesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeAccountAttributesInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeAccountAttributesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeAccountAttributesOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeAccountAttributesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeAccountLimits(input *pinpointsmsvoicev2.DescribeAccountLimitsInput) (*pinpointsmsvoicev2.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeAccountLimitsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeAccountLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeAccountLimitsPages(input *pinpointsmsvoicev2.DescribeAccountLimitsInput, fn func(*pinpointsmsvoicev2.DescribeAccountLimitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeAccountLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeAccountLimitsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeAccountLimitsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeAccountLimitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeAccountLimitsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeAccountLimitsInput, fn func(*pinpointsmsvoicev2.DescribeAccountLimitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeAccountLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeAccountLimitsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeAccountLimitsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeAccountLimitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeAccountLimitsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeAccountLimitsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeAccountLimitsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeAccountLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeConfigurationSets(input *pinpointsmsvoicev2.DescribeConfigurationSetsInput) (*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeConfigurationSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeConfigurationSetsPages(input *pinpointsmsvoicev2.DescribeConfigurationSetsInput, fn func(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeConfigurationSetsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeConfigurationSetsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeConfigurationSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeConfigurationSetsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeConfigurationSetsInput, fn func(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeConfigurationSetsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeConfigurationSetsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeConfigurationSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeConfigurationSetsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeConfigurationSetsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeConfigurationSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeConfigurationSetsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeConfigurationSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeKeywords(input *pinpointsmsvoicev2.DescribeKeywordsInput) (*pinpointsmsvoicev2.DescribeKeywordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeKeywordsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeKeywords(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeKeywordsPages(input *pinpointsmsvoicev2.DescribeKeywordsInput, fn func(*pinpointsmsvoicev2.DescribeKeywordsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeKeywordsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeKeywordsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeKeywordsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeKeywordsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeKeywordsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeKeywordsInput, fn func(*pinpointsmsvoicev2.DescribeKeywordsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeKeywordsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeKeywordsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeKeywordsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeKeywordsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeKeywordsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeKeywordsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeKeywordsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeKeywordsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeKeywordsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeOptOutLists(input *pinpointsmsvoicev2.DescribeOptOutListsInput) (*pinpointsmsvoicev2.DescribeOptOutListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeOptOutListsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeOptOutLists(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeOptOutListsPages(input *pinpointsmsvoicev2.DescribeOptOutListsInput, fn func(*pinpointsmsvoicev2.DescribeOptOutListsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeOptOutListsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeOptOutListsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeOptOutListsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeOptOutListsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeOptOutListsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeOptOutListsInput, fn func(*pinpointsmsvoicev2.DescribeOptOutListsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeOptOutListsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeOptOutListsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeOptOutListsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeOptOutListsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeOptOutListsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeOptOutListsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeOptOutListsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeOptOutListsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeOptOutListsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeOptedOutNumbers(input *pinpointsmsvoicev2.DescribeOptedOutNumbersInput) (*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeOptedOutNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeOptedOutNumbersPages(input *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, fn func(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeOptedOutNumbersOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeOptedOutNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeOptedOutNumbersPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, fn func(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeOptedOutNumbersOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeOptedOutNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeOptedOutNumbersWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeOptedOutNumbersInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeOptedOutNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribePhoneNumbers(input *pinpointsmsvoicev2.DescribePhoneNumbersInput) (*pinpointsmsvoicev2.DescribePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribePhoneNumbersOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribePhoneNumbers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribePhoneNumbersPages(input *pinpointsmsvoicev2.DescribePhoneNumbersInput, fn func(*pinpointsmsvoicev2.DescribePhoneNumbersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribePhoneNumbersOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribePhoneNumbersOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribePhoneNumbersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribePhoneNumbersPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribePhoneNumbersInput, fn func(*pinpointsmsvoicev2.DescribePhoneNumbersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribePhoneNumbersOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribePhoneNumbersOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribePhoneNumbersOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribePhoneNumbersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribePhoneNumbersWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribePhoneNumbersInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribePhoneNumbersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribePhoneNumbersOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribePhoneNumbersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribePools(input *pinpointsmsvoicev2.DescribePoolsInput) (*pinpointsmsvoicev2.DescribePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribePoolsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribePools(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribePoolsPages(input *pinpointsmsvoicev2.DescribePoolsInput, fn func(*pinpointsmsvoicev2.DescribePoolsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribePoolsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribePoolsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribePoolsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribePoolsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribePoolsInput, fn func(*pinpointsmsvoicev2.DescribePoolsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribePoolsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribePoolsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribePoolsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribePoolsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribePoolsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribePoolsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribePoolsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribePoolsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribePoolsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeSenderIds(input *pinpointsmsvoicev2.DescribeSenderIdsInput) (*pinpointsmsvoicev2.DescribeSenderIdsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeSenderIdsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeSenderIds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeSenderIdsPages(input *pinpointsmsvoicev2.DescribeSenderIdsInput, fn func(*pinpointsmsvoicev2.DescribeSenderIdsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeSenderIdsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeSenderIdsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeSenderIdsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeSenderIdsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeSenderIdsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeSenderIdsInput, fn func(*pinpointsmsvoicev2.DescribeSenderIdsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeSenderIdsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeSenderIdsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeSenderIdsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeSenderIdsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeSenderIdsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeSenderIdsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeSenderIdsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeSenderIdsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeSenderIdsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeSpendLimits(input *pinpointsmsvoicev2.DescribeSpendLimitsInput) (*pinpointsmsvoicev2.DescribeSpendLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeSpendLimitsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeSpendLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) DescribeSpendLimitsPages(input *pinpointsmsvoicev2.DescribeSpendLimitsInput, fn func(*pinpointsmsvoicev2.DescribeSpendLimitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeSpendLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeSpendLimitsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeSpendLimitsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeSpendLimitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeSpendLimitsPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeSpendLimitsInput, fn func(*pinpointsmsvoicev2.DescribeSpendLimitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.DescribeSpendLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.DescribeSpendLimitsOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.DescribeSpendLimitsOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.DescribeSpendLimitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) DescribeSpendLimitsWithContext(ctx context.Context, input *pinpointsmsvoicev2.DescribeSpendLimitsInput, opts ...request.Option) (*pinpointsmsvoicev2.DescribeSpendLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.DescribeSpendLimitsOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.DescribeSpendLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) ListPoolOriginationIdentities(input *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput) (*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.ListPoolOriginationIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) ListPoolOriginationIdentitiesPages(input *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, fn func(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.ListPoolOriginationIdentitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) ListPoolOriginationIdentitiesPagesWithContext(ctx context.Context, input *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, fn func(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput{}
	fnCacher := func(out *pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, more bool) bool {
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
	if err := c.PinpointSMSVoiceV2API.ListPoolOriginationIdentitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *PinpointSMSVoiceV2) ListPoolOriginationIdentitiesWithContext(ctx context.Context, input *pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput, opts ...request.Option) (*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.ListPoolOriginationIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) ListTagsForResource(input *pinpointsmsvoicev2.ListTagsForResourceInput) (*pinpointsmsvoicev2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *PinpointSMSVoiceV2) ListTagsForResourceWithContext(ctx context.Context, input *pinpointsmsvoicev2.ListTagsForResourceInput, opts ...request.Option) (*pinpointsmsvoicev2.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*pinpointsmsvoicev2.ListTagsForResourceOutput), nil
	}
	out, err := c.PinpointSMSVoiceV2API.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
