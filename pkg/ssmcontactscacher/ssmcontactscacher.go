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
package ssmcontactscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ssmcontacts"
	"github.com/aws/aws-sdk-go/service/ssmcontacts/ssmcontactsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SSMContacts struct {
	ssmcontactsiface.SSMContactsAPI
	cache *cache.Cache
}

func New(ssmcontactsapi ssmcontactsiface.SSMContactsAPI) *SSMContacts {
	return &SSMContacts{
		SSMContactsAPI: ssmcontactsapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SSMContacts) DescribeEngagement(input *ssmcontacts.DescribeEngagementInput) (*ssmcontacts.DescribeEngagementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.DescribeEngagementOutput), nil
	}
	out, err := c.SSMContactsAPI.DescribeEngagement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) DescribeEngagementWithContext(ctx context.Context, input *ssmcontacts.DescribeEngagementInput, opts ...request.Option) (*ssmcontacts.DescribeEngagementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.DescribeEngagementOutput), nil
	}
	out, err := c.SSMContactsAPI.DescribeEngagementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) DescribePage(input *ssmcontacts.DescribePageInput) (*ssmcontacts.DescribePageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.DescribePageOutput), nil
	}
	out, err := c.SSMContactsAPI.DescribePage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) DescribePageWithContext(ctx context.Context, input *ssmcontacts.DescribePageInput, opts ...request.Option) (*ssmcontacts.DescribePageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.DescribePageOutput), nil
	}
	out, err := c.SSMContactsAPI.DescribePageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContact(input *ssmcontacts.GetContactInput) (*ssmcontacts.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContactChannel(input *ssmcontacts.GetContactChannelInput) (*ssmcontacts.GetContactChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactChannelOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContactChannel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContactChannelWithContext(ctx context.Context, input *ssmcontacts.GetContactChannelInput, opts ...request.Option) (*ssmcontacts.GetContactChannelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactChannelOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContactChannelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContactPolicy(input *ssmcontacts.GetContactPolicyInput) (*ssmcontacts.GetContactPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactPolicyOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContactPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContactPolicyWithContext(ctx context.Context, input *ssmcontacts.GetContactPolicyInput, opts ...request.Option) (*ssmcontacts.GetContactPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactPolicyOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContactPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) GetContactWithContext(ctx context.Context, input *ssmcontacts.GetContactInput, opts ...request.Option) (*ssmcontacts.GetContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.GetContactOutput), nil
	}
	out, err := c.SSMContactsAPI.GetContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListContactChannels(input *ssmcontacts.ListContactChannelsInput) (*ssmcontacts.ListContactChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListContactChannelsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListContactChannels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListContactChannelsPages(input *ssmcontacts.ListContactChannelsInput, fn func(*ssmcontacts.ListContactChannelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListContactChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListContactChannelsOutput{}
	fnCacher := func(out *ssmcontacts.ListContactChannelsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListContactChannelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListContactChannelsPagesWithContext(ctx context.Context, input *ssmcontacts.ListContactChannelsInput, fn func(*ssmcontacts.ListContactChannelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListContactChannelsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListContactChannelsOutput{}
	fnCacher := func(out *ssmcontacts.ListContactChannelsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListContactChannelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListContactChannelsWithContext(ctx context.Context, input *ssmcontacts.ListContactChannelsInput, opts ...request.Option) (*ssmcontacts.ListContactChannelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListContactChannelsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListContactChannelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListContacts(input *ssmcontacts.ListContactsInput) (*ssmcontacts.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListContactsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListContacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListContactsPages(input *ssmcontacts.ListContactsInput, fn func(*ssmcontacts.ListContactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListContactsOutput{}
	fnCacher := func(out *ssmcontacts.ListContactsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListContactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListContactsPagesWithContext(ctx context.Context, input *ssmcontacts.ListContactsInput, fn func(*ssmcontacts.ListContactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListContactsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListContactsOutput{}
	fnCacher := func(out *ssmcontacts.ListContactsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListContactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListContactsWithContext(ctx context.Context, input *ssmcontacts.ListContactsInput, opts ...request.Option) (*ssmcontacts.ListContactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListContactsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListContactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListEngagements(input *ssmcontacts.ListEngagementsInput) (*ssmcontacts.ListEngagementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListEngagementsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListEngagements(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListEngagementsPages(input *ssmcontacts.ListEngagementsInput, fn func(*ssmcontacts.ListEngagementsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListEngagementsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListEngagementsOutput{}
	fnCacher := func(out *ssmcontacts.ListEngagementsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListEngagementsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListEngagementsPagesWithContext(ctx context.Context, input *ssmcontacts.ListEngagementsInput, fn func(*ssmcontacts.ListEngagementsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListEngagementsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListEngagementsOutput{}
	fnCacher := func(out *ssmcontacts.ListEngagementsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListEngagementsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListEngagementsWithContext(ctx context.Context, input *ssmcontacts.ListEngagementsInput, opts ...request.Option) (*ssmcontacts.ListEngagementsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListEngagementsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListEngagementsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPageReceipts(input *ssmcontacts.ListPageReceiptsInput) (*ssmcontacts.ListPageReceiptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPageReceiptsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPageReceipts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPageReceiptsPages(input *ssmcontacts.ListPageReceiptsInput, fn func(*ssmcontacts.ListPageReceiptsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPageReceiptsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPageReceiptsOutput{}
	fnCacher := func(out *ssmcontacts.ListPageReceiptsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPageReceiptsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPageReceiptsPagesWithContext(ctx context.Context, input *ssmcontacts.ListPageReceiptsInput, fn func(*ssmcontacts.ListPageReceiptsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPageReceiptsOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPageReceiptsOutput{}
	fnCacher := func(out *ssmcontacts.ListPageReceiptsOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPageReceiptsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPageReceiptsWithContext(ctx context.Context, input *ssmcontacts.ListPageReceiptsInput, opts ...request.Option) (*ssmcontacts.ListPageReceiptsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPageReceiptsOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPageReceiptsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPagesByContact(input *ssmcontacts.ListPagesByContactInput) (*ssmcontacts.ListPagesByContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPagesByContactOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPagesByContact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPagesByContactPages(input *ssmcontacts.ListPagesByContactInput, fn func(*ssmcontacts.ListPagesByContactOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPagesByContactOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPagesByContactOutput{}
	fnCacher := func(out *ssmcontacts.ListPagesByContactOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPagesByContactPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPagesByContactPagesWithContext(ctx context.Context, input *ssmcontacts.ListPagesByContactInput, fn func(*ssmcontacts.ListPagesByContactOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPagesByContactOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPagesByContactOutput{}
	fnCacher := func(out *ssmcontacts.ListPagesByContactOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPagesByContactPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPagesByContactWithContext(ctx context.Context, input *ssmcontacts.ListPagesByContactInput, opts ...request.Option) (*ssmcontacts.ListPagesByContactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPagesByContactOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPagesByContactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPagesByEngagement(input *ssmcontacts.ListPagesByEngagementInput) (*ssmcontacts.ListPagesByEngagementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPagesByEngagementOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPagesByEngagement(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListPagesByEngagementPages(input *ssmcontacts.ListPagesByEngagementInput, fn func(*ssmcontacts.ListPagesByEngagementOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPagesByEngagementOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPagesByEngagementOutput{}
	fnCacher := func(out *ssmcontacts.ListPagesByEngagementOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPagesByEngagementPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPagesByEngagementPagesWithContext(ctx context.Context, input *ssmcontacts.ListPagesByEngagementInput, fn func(*ssmcontacts.ListPagesByEngagementOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ssmcontacts.ListPagesByEngagementOutput), false)
		return nil
	}
	cachable := true
	output := &ssmcontacts.ListPagesByEngagementOutput{}
	fnCacher := func(out *ssmcontacts.ListPagesByEngagementOutput, more bool) bool {
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
	if err := c.SSMContactsAPI.ListPagesByEngagementPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SSMContacts) ListPagesByEngagementWithContext(ctx context.Context, input *ssmcontacts.ListPagesByEngagementInput, opts ...request.Option) (*ssmcontacts.ListPagesByEngagementOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListPagesByEngagementOutput), nil
	}
	out, err := c.SSMContactsAPI.ListPagesByEngagementWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListTagsForResource(input *ssmcontacts.ListTagsForResourceInput) (*ssmcontacts.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMContactsAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SSMContacts) ListTagsForResourceWithContext(ctx context.Context, input *ssmcontacts.ListTagsForResourceInput, opts ...request.Option) (*ssmcontacts.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ssmcontacts.ListTagsForResourceOutput), nil
	}
	out, err := c.SSMContactsAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
