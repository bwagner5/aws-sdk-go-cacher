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
package mturkcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mturk"
	"github.com/aws/aws-sdk-go/service/mturk/mturkiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type MTurk struct {
	mturkiface.MTurkAPI
	cache *cache.Cache
}

func New(mturkapi mturkiface.MTurkAPI) *MTurk {
	return &MTurk{
		MTurkAPI: mturkapi,
		cache:    cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *MTurk) GetAccountBalance(input *mturk.GetAccountBalanceInput) (*mturk.GetAccountBalanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetAccountBalanceOutput), nil
	}
	out, err := c.MTurkAPI.GetAccountBalance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetAccountBalanceWithContext(ctx context.Context, input *mturk.GetAccountBalanceInput, opts ...request.Option) (*mturk.GetAccountBalanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetAccountBalanceOutput), nil
	}
	out, err := c.MTurkAPI.GetAccountBalanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetAssignment(input *mturk.GetAssignmentInput) (*mturk.GetAssignmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetAssignmentOutput), nil
	}
	out, err := c.MTurkAPI.GetAssignment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetAssignmentWithContext(ctx context.Context, input *mturk.GetAssignmentInput, opts ...request.Option) (*mturk.GetAssignmentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetAssignmentOutput), nil
	}
	out, err := c.MTurkAPI.GetAssignmentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetFileUploadURL(input *mturk.GetFileUploadURLInput) (*mturk.GetFileUploadURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetFileUploadURLOutput), nil
	}
	out, err := c.MTurkAPI.GetFileUploadURL(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetFileUploadURLWithContext(ctx context.Context, input *mturk.GetFileUploadURLInput, opts ...request.Option) (*mturk.GetFileUploadURLOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetFileUploadURLOutput), nil
	}
	out, err := c.MTurkAPI.GetFileUploadURLWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetHIT(input *mturk.GetHITInput) (*mturk.GetHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetHITOutput), nil
	}
	out, err := c.MTurkAPI.GetHIT(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetHITWithContext(ctx context.Context, input *mturk.GetHITInput, opts ...request.Option) (*mturk.GetHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetHITOutput), nil
	}
	out, err := c.MTurkAPI.GetHITWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetQualificationScore(input *mturk.GetQualificationScoreInput) (*mturk.GetQualificationScoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetQualificationScoreOutput), nil
	}
	out, err := c.MTurkAPI.GetQualificationScore(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetQualificationScoreWithContext(ctx context.Context, input *mturk.GetQualificationScoreInput, opts ...request.Option) (*mturk.GetQualificationScoreOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetQualificationScoreOutput), nil
	}
	out, err := c.MTurkAPI.GetQualificationScoreWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetQualificationType(input *mturk.GetQualificationTypeInput) (*mturk.GetQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.GetQualificationType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) GetQualificationTypeWithContext(ctx context.Context, input *mturk.GetQualificationTypeInput, opts ...request.Option) (*mturk.GetQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.GetQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.GetQualificationTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListAssignmentsForHIT(input *mturk.ListAssignmentsForHITInput) (*mturk.ListAssignmentsForHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListAssignmentsForHITOutput), nil
	}
	out, err := c.MTurkAPI.ListAssignmentsForHIT(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListAssignmentsForHITPages(input *mturk.ListAssignmentsForHITInput, fn func(*mturk.ListAssignmentsForHITOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListAssignmentsForHITOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListAssignmentsForHITOutput{}
	fnCacher := func(out *mturk.ListAssignmentsForHITOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListAssignmentsForHITPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListAssignmentsForHITPagesWithContext(ctx context.Context, input *mturk.ListAssignmentsForHITInput, fn func(*mturk.ListAssignmentsForHITOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListAssignmentsForHITOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListAssignmentsForHITOutput{}
	fnCacher := func(out *mturk.ListAssignmentsForHITOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListAssignmentsForHITPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListAssignmentsForHITWithContext(ctx context.Context, input *mturk.ListAssignmentsForHITInput, opts ...request.Option) (*mturk.ListAssignmentsForHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListAssignmentsForHITOutput), nil
	}
	out, err := c.MTurkAPI.ListAssignmentsForHITWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListBonusPayments(input *mturk.ListBonusPaymentsInput) (*mturk.ListBonusPaymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListBonusPaymentsOutput), nil
	}
	out, err := c.MTurkAPI.ListBonusPayments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListBonusPaymentsPages(input *mturk.ListBonusPaymentsInput, fn func(*mturk.ListBonusPaymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListBonusPaymentsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListBonusPaymentsOutput{}
	fnCacher := func(out *mturk.ListBonusPaymentsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListBonusPaymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListBonusPaymentsPagesWithContext(ctx context.Context, input *mturk.ListBonusPaymentsInput, fn func(*mturk.ListBonusPaymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListBonusPaymentsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListBonusPaymentsOutput{}
	fnCacher := func(out *mturk.ListBonusPaymentsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListBonusPaymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListBonusPaymentsWithContext(ctx context.Context, input *mturk.ListBonusPaymentsInput, opts ...request.Option) (*mturk.ListBonusPaymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListBonusPaymentsOutput), nil
	}
	out, err := c.MTurkAPI.ListBonusPaymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListHITs(input *mturk.ListHITsInput) (*mturk.ListHITsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListHITsOutput), nil
	}
	out, err := c.MTurkAPI.ListHITs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListHITsForQualificationType(input *mturk.ListHITsForQualificationTypeInput) (*mturk.ListHITsForQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListHITsForQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.ListHITsForQualificationType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListHITsForQualificationTypePages(input *mturk.ListHITsForQualificationTypeInput, fn func(*mturk.ListHITsForQualificationTypeOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListHITsForQualificationTypeOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListHITsForQualificationTypeOutput{}
	fnCacher := func(out *mturk.ListHITsForQualificationTypeOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListHITsForQualificationTypePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListHITsForQualificationTypePagesWithContext(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput, fn func(*mturk.ListHITsForQualificationTypeOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListHITsForQualificationTypeOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListHITsForQualificationTypeOutput{}
	fnCacher := func(out *mturk.ListHITsForQualificationTypeOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListHITsForQualificationTypePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListHITsForQualificationTypeWithContext(ctx context.Context, input *mturk.ListHITsForQualificationTypeInput, opts ...request.Option) (*mturk.ListHITsForQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListHITsForQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.ListHITsForQualificationTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListHITsPages(input *mturk.ListHITsInput, fn func(*mturk.ListHITsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListHITsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListHITsOutput{}
	fnCacher := func(out *mturk.ListHITsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListHITsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListHITsPagesWithContext(ctx context.Context, input *mturk.ListHITsInput, fn func(*mturk.ListHITsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListHITsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListHITsOutput{}
	fnCacher := func(out *mturk.ListHITsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListHITsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListHITsWithContext(ctx context.Context, input *mturk.ListHITsInput, opts ...request.Option) (*mturk.ListHITsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListHITsOutput), nil
	}
	out, err := c.MTurkAPI.ListHITsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListQualificationRequests(input *mturk.ListQualificationRequestsInput) (*mturk.ListQualificationRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListQualificationRequestsOutput), nil
	}
	out, err := c.MTurkAPI.ListQualificationRequests(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListQualificationRequestsPages(input *mturk.ListQualificationRequestsInput, fn func(*mturk.ListQualificationRequestsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListQualificationRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListQualificationRequestsOutput{}
	fnCacher := func(out *mturk.ListQualificationRequestsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListQualificationRequestsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListQualificationRequestsPagesWithContext(ctx context.Context, input *mturk.ListQualificationRequestsInput, fn func(*mturk.ListQualificationRequestsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListQualificationRequestsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListQualificationRequestsOutput{}
	fnCacher := func(out *mturk.ListQualificationRequestsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListQualificationRequestsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListQualificationRequestsWithContext(ctx context.Context, input *mturk.ListQualificationRequestsInput, opts ...request.Option) (*mturk.ListQualificationRequestsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListQualificationRequestsOutput), nil
	}
	out, err := c.MTurkAPI.ListQualificationRequestsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListQualificationTypes(input *mturk.ListQualificationTypesInput) (*mturk.ListQualificationTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListQualificationTypesOutput), nil
	}
	out, err := c.MTurkAPI.ListQualificationTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListQualificationTypesPages(input *mturk.ListQualificationTypesInput, fn func(*mturk.ListQualificationTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListQualificationTypesOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListQualificationTypesOutput{}
	fnCacher := func(out *mturk.ListQualificationTypesOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListQualificationTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListQualificationTypesPagesWithContext(ctx context.Context, input *mturk.ListQualificationTypesInput, fn func(*mturk.ListQualificationTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListQualificationTypesOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListQualificationTypesOutput{}
	fnCacher := func(out *mturk.ListQualificationTypesOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListQualificationTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListQualificationTypesWithContext(ctx context.Context, input *mturk.ListQualificationTypesInput, opts ...request.Option) (*mturk.ListQualificationTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListQualificationTypesOutput), nil
	}
	out, err := c.MTurkAPI.ListQualificationTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListReviewPolicyResultsForHIT(input *mturk.ListReviewPolicyResultsForHITInput) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListReviewPolicyResultsForHITOutput), nil
	}
	out, err := c.MTurkAPI.ListReviewPolicyResultsForHIT(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListReviewPolicyResultsForHITPages(input *mturk.ListReviewPolicyResultsForHITInput, fn func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListReviewPolicyResultsForHITOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListReviewPolicyResultsForHITOutput{}
	fnCacher := func(out *mturk.ListReviewPolicyResultsForHITOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListReviewPolicyResultsForHITPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListReviewPolicyResultsForHITPagesWithContext(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput, fn func(*mturk.ListReviewPolicyResultsForHITOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListReviewPolicyResultsForHITOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListReviewPolicyResultsForHITOutput{}
	fnCacher := func(out *mturk.ListReviewPolicyResultsForHITOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListReviewPolicyResultsForHITPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListReviewPolicyResultsForHITWithContext(ctx context.Context, input *mturk.ListReviewPolicyResultsForHITInput, opts ...request.Option) (*mturk.ListReviewPolicyResultsForHITOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListReviewPolicyResultsForHITOutput), nil
	}
	out, err := c.MTurkAPI.ListReviewPolicyResultsForHITWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListReviewableHITs(input *mturk.ListReviewableHITsInput) (*mturk.ListReviewableHITsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListReviewableHITsOutput), nil
	}
	out, err := c.MTurkAPI.ListReviewableHITs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListReviewableHITsPages(input *mturk.ListReviewableHITsInput, fn func(*mturk.ListReviewableHITsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListReviewableHITsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListReviewableHITsOutput{}
	fnCacher := func(out *mturk.ListReviewableHITsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListReviewableHITsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListReviewableHITsPagesWithContext(ctx context.Context, input *mturk.ListReviewableHITsInput, fn func(*mturk.ListReviewableHITsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListReviewableHITsOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListReviewableHITsOutput{}
	fnCacher := func(out *mturk.ListReviewableHITsOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListReviewableHITsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListReviewableHITsWithContext(ctx context.Context, input *mturk.ListReviewableHITsInput, opts ...request.Option) (*mturk.ListReviewableHITsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListReviewableHITsOutput), nil
	}
	out, err := c.MTurkAPI.ListReviewableHITsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListWorkerBlocks(input *mturk.ListWorkerBlocksInput) (*mturk.ListWorkerBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListWorkerBlocksOutput), nil
	}
	out, err := c.MTurkAPI.ListWorkerBlocks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListWorkerBlocksPages(input *mturk.ListWorkerBlocksInput, fn func(*mturk.ListWorkerBlocksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListWorkerBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListWorkerBlocksOutput{}
	fnCacher := func(out *mturk.ListWorkerBlocksOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListWorkerBlocksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListWorkerBlocksPagesWithContext(ctx context.Context, input *mturk.ListWorkerBlocksInput, fn func(*mturk.ListWorkerBlocksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListWorkerBlocksOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListWorkerBlocksOutput{}
	fnCacher := func(out *mturk.ListWorkerBlocksOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListWorkerBlocksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListWorkerBlocksWithContext(ctx context.Context, input *mturk.ListWorkerBlocksInput, opts ...request.Option) (*mturk.ListWorkerBlocksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListWorkerBlocksOutput), nil
	}
	out, err := c.MTurkAPI.ListWorkerBlocksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListWorkersWithQualificationType(input *mturk.ListWorkersWithQualificationTypeInput) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListWorkersWithQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.ListWorkersWithQualificationType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *MTurk) ListWorkersWithQualificationTypePages(input *mturk.ListWorkersWithQualificationTypeInput, fn func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListWorkersWithQualificationTypeOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListWorkersWithQualificationTypeOutput{}
	fnCacher := func(out *mturk.ListWorkersWithQualificationTypeOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListWorkersWithQualificationTypePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListWorkersWithQualificationTypePagesWithContext(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput, fn func(*mturk.ListWorkersWithQualificationTypeOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*mturk.ListWorkersWithQualificationTypeOutput), false)
		return nil
	}
	cachable := true
	output := &mturk.ListWorkersWithQualificationTypeOutput{}
	fnCacher := func(out *mturk.ListWorkersWithQualificationTypeOutput, more bool) bool {
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
	if err := c.MTurkAPI.ListWorkersWithQualificationTypePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *MTurk) ListWorkersWithQualificationTypeWithContext(ctx context.Context, input *mturk.ListWorkersWithQualificationTypeInput, opts ...request.Option) (*mturk.ListWorkersWithQualificationTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*mturk.ListWorkersWithQualificationTypeOutput), nil
	}
	out, err := c.MTurkAPI.ListWorkersWithQualificationTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
