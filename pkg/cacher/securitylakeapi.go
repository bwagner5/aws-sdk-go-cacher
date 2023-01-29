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
	"github.com/aws/aws-sdk-go/service/securitylake"
	"github.com/aws/aws-sdk-go/service/securitylake/securitylakeiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SecurityLake struct {
	securitylakeiface.SecurityLakeAPI
	cache *cache.Cache
}

func NewSecurityLake(securitylakeapi securitylakeiface.SecurityLakeAPI) *SecurityLake {
	return &SecurityLake{
		SecurityLakeAPI: securitylakeapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SecurityLake) GetDatalake(input *securitylake.GetDatalakeInput) (*securitylake.GetDatalakeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalake(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeAutoEnable(input *securitylake.GetDatalakeAutoEnableInput) (*securitylake.GetDatalakeAutoEnableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeAutoEnableOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeAutoEnable(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeAutoEnableWithContext(ctx context.Context, input *securitylake.GetDatalakeAutoEnableInput, opts ...request.Option) (*securitylake.GetDatalakeAutoEnableOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeAutoEnableOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeAutoEnableWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeExceptionsExpiry(input *securitylake.GetDatalakeExceptionsExpiryInput) (*securitylake.GetDatalakeExceptionsExpiryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeExceptionsExpiryOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeExceptionsExpiry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeExceptionsExpiryWithContext(ctx context.Context, input *securitylake.GetDatalakeExceptionsExpiryInput, opts ...request.Option) (*securitylake.GetDatalakeExceptionsExpiryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeExceptionsExpiryOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeExceptionsExpiryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeExceptionsSubscription(input *securitylake.GetDatalakeExceptionsSubscriptionInput) (*securitylake.GetDatalakeExceptionsSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeExceptionsSubscriptionOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeExceptionsSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeExceptionsSubscriptionWithContext(ctx context.Context, input *securitylake.GetDatalakeExceptionsSubscriptionInput, opts ...request.Option) (*securitylake.GetDatalakeExceptionsSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeExceptionsSubscriptionOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeExceptionsSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeStatus(input *securitylake.GetDatalakeStatusInput) (*securitylake.GetDatalakeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeStatusOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeStatusPages(input *securitylake.GetDatalakeStatusInput, fn func(*securitylake.GetDatalakeStatusOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.GetDatalakeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.GetDatalakeStatusOutput{}
	fnCacher := func(out *securitylake.GetDatalakeStatusOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.GetDatalakeStatusPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) GetDatalakeStatusPagesWithContext(ctx context.Context, input *securitylake.GetDatalakeStatusInput, fn func(*securitylake.GetDatalakeStatusOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.GetDatalakeStatusOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.GetDatalakeStatusOutput{}
	fnCacher := func(out *securitylake.GetDatalakeStatusOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.GetDatalakeStatusPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) GetDatalakeStatusWithContext(ctx context.Context, input *securitylake.GetDatalakeStatusInput, opts ...request.Option) (*securitylake.GetDatalakeStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeStatusOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetDatalakeWithContext(ctx context.Context, input *securitylake.GetDatalakeInput, opts ...request.Option) (*securitylake.GetDatalakeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetDatalakeOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetDatalakeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetSubscriber(input *securitylake.GetSubscriberInput) (*securitylake.GetSubscriberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetSubscriberOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetSubscriber(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) GetSubscriberWithContext(ctx context.Context, input *securitylake.GetSubscriberInput, opts ...request.Option) (*securitylake.GetSubscriberOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.GetSubscriberOutput), nil
	}
	out, err := c.SecurityLakeAPI.GetSubscriberWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListDatalakeExceptions(input *securitylake.ListDatalakeExceptionsInput) (*securitylake.ListDatalakeExceptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListDatalakeExceptionsOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListDatalakeExceptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListDatalakeExceptionsPages(input *securitylake.ListDatalakeExceptionsInput, fn func(*securitylake.ListDatalakeExceptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListDatalakeExceptionsOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListDatalakeExceptionsOutput{}
	fnCacher := func(out *securitylake.ListDatalakeExceptionsOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListDatalakeExceptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListDatalakeExceptionsPagesWithContext(ctx context.Context, input *securitylake.ListDatalakeExceptionsInput, fn func(*securitylake.ListDatalakeExceptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListDatalakeExceptionsOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListDatalakeExceptionsOutput{}
	fnCacher := func(out *securitylake.ListDatalakeExceptionsOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListDatalakeExceptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListDatalakeExceptionsWithContext(ctx context.Context, input *securitylake.ListDatalakeExceptionsInput, opts ...request.Option) (*securitylake.ListDatalakeExceptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListDatalakeExceptionsOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListDatalakeExceptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListLogSources(input *securitylake.ListLogSourcesInput) (*securitylake.ListLogSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListLogSourcesOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListLogSources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListLogSourcesPages(input *securitylake.ListLogSourcesInput, fn func(*securitylake.ListLogSourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListLogSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListLogSourcesOutput{}
	fnCacher := func(out *securitylake.ListLogSourcesOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListLogSourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListLogSourcesPagesWithContext(ctx context.Context, input *securitylake.ListLogSourcesInput, fn func(*securitylake.ListLogSourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListLogSourcesOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListLogSourcesOutput{}
	fnCacher := func(out *securitylake.ListLogSourcesOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListLogSourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListLogSourcesWithContext(ctx context.Context, input *securitylake.ListLogSourcesInput, opts ...request.Option) (*securitylake.ListLogSourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListLogSourcesOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListLogSourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListSubscribers(input *securitylake.ListSubscribersInput) (*securitylake.ListSubscribersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListSubscribersOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListSubscribers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityLake) ListSubscribersPages(input *securitylake.ListSubscribersInput, fn func(*securitylake.ListSubscribersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListSubscribersOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListSubscribersOutput{}
	fnCacher := func(out *securitylake.ListSubscribersOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListSubscribersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListSubscribersPagesWithContext(ctx context.Context, input *securitylake.ListSubscribersInput, fn func(*securitylake.ListSubscribersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securitylake.ListSubscribersOutput), false)
		return nil
	}
	cachable := true
	output := &securitylake.ListSubscribersOutput{}
	fnCacher := func(out *securitylake.ListSubscribersOutput, more bool) bool {
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
	if err := c.SecurityLakeAPI.ListSubscribersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityLake) ListSubscribersWithContext(ctx context.Context, input *securitylake.ListSubscribersInput, opts ...request.Option) (*securitylake.ListSubscribersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securitylake.ListSubscribersOutput), nil
	}
	out, err := c.SecurityLakeAPI.ListSubscribersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
