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
package licensemanagerusersubscriptionscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanagerusersubscriptions"
	"github.com/aws/aws-sdk-go/service/licensemanagerusersubscriptions/licensemanagerusersubscriptionsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LicenseManagerUserSubscriptions struct {
	licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI
	cache *cache.Cache
}

func New(licensemanagerusersubscriptionsapi licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI) *LicenseManagerUserSubscriptions {
	return &LicenseManagerUserSubscriptions{
		LicenseManagerUserSubscriptionsAPI: licensemanagerusersubscriptionsapi,
		cache:                              cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LicenseManagerUserSubscriptions) ListIdentityProviders(input *licensemanagerusersubscriptions.ListIdentityProvidersInput) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListIdentityProvidersOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListIdentityProviders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListIdentityProvidersPages(input *licensemanagerusersubscriptions.ListIdentityProvidersInput, fn func(*licensemanagerusersubscriptions.ListIdentityProvidersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListIdentityProvidersOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListIdentityProvidersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListIdentityProvidersPagesWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListIdentityProvidersInput, fn func(*licensemanagerusersubscriptions.ListIdentityProvidersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListIdentityProvidersOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListIdentityProvidersOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListIdentityProvidersOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListIdentityProvidersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListIdentityProvidersWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListIdentityProvidersInput, opts ...request.Option) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListIdentityProvidersOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListIdentityProvidersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListInstances(input *licensemanagerusersubscriptions.ListInstancesInput) (*licensemanagerusersubscriptions.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListInstancesOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListInstancesPages(input *licensemanagerusersubscriptions.ListInstancesInput, fn func(*licensemanagerusersubscriptions.ListInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListInstancesOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListInstancesOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListInstancesPagesWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListInstancesInput, fn func(*licensemanagerusersubscriptions.ListInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListInstancesOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListInstancesOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListInstancesWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListInstancesInput, opts ...request.Option) (*licensemanagerusersubscriptions.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListInstancesOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListProductSubscriptions(input *licensemanagerusersubscriptions.ListProductSubscriptionsInput) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListProductSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListProductSubscriptionsPages(input *licensemanagerusersubscriptions.ListProductSubscriptionsInput, fn func(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListProductSubscriptionsOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListProductSubscriptionsOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListProductSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListProductSubscriptionsPagesWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListProductSubscriptionsInput, fn func(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListProductSubscriptionsOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListProductSubscriptionsOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListProductSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListProductSubscriptionsWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListProductSubscriptionsInput, opts ...request.Option) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListProductSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListUserAssociations(input *licensemanagerusersubscriptions.ListUserAssociationsInput) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListUserAssociationsOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListUserAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerUserSubscriptions) ListUserAssociationsPages(input *licensemanagerusersubscriptions.ListUserAssociationsInput, fn func(*licensemanagerusersubscriptions.ListUserAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListUserAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListUserAssociationsOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListUserAssociationsOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListUserAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListUserAssociationsPagesWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListUserAssociationsInput, fn func(*licensemanagerusersubscriptions.ListUserAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerusersubscriptions.ListUserAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerusersubscriptions.ListUserAssociationsOutput{}
	fnCacher := func(out *licensemanagerusersubscriptions.ListUserAssociationsOutput, more bool) bool {
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
	if err := c.LicenseManagerUserSubscriptionsAPI.ListUserAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerUserSubscriptions) ListUserAssociationsWithContext(ctx context.Context, input *licensemanagerusersubscriptions.ListUserAssociationsInput, opts ...request.Option) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerusersubscriptions.ListUserAssociationsOutput), nil
	}
	out, err := c.LicenseManagerUserSubscriptionsAPI.ListUserAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
