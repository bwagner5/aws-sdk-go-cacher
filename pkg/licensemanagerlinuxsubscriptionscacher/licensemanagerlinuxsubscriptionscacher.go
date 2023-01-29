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
package licensemanagerlinuxsubscriptionscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanagerlinuxsubscriptions"
	"github.com/aws/aws-sdk-go/service/licensemanagerlinuxsubscriptions/licensemanagerlinuxsubscriptionsiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type LicenseManagerLinuxSubscriptions struct {
	licensemanagerlinuxsubscriptionsiface.LicenseManagerLinuxSubscriptionsAPI
	cache *cache.Cache
}

func New(licensemanagerlinuxsubscriptionsapi licensemanagerlinuxsubscriptionsiface.LicenseManagerLinuxSubscriptionsAPI) *LicenseManagerLinuxSubscriptions {
	return &LicenseManagerLinuxSubscriptions{
		LicenseManagerLinuxSubscriptionsAPI: licensemanagerlinuxsubscriptionsapi,
		cache:                               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *LicenseManagerLinuxSubscriptions) GetServiceSettings(input *licensemanagerlinuxsubscriptions.GetServiceSettingsInput) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.GetServiceSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerLinuxSubscriptions) GetServiceSettingsWithContext(ctx context.Context, input *licensemanagerlinuxsubscriptions.GetServiceSettingsInput, opts ...request.Option) (*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.GetServiceSettingsOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.GetServiceSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionInstances(input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionInstancesPages(input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, fn func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput{}
	fnCacher := func(out *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, more bool) bool {
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
	if err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionInstancesPagesWithContext(ctx context.Context, input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, fn func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput{}
	fnCacher := func(out *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, more bool) bool {
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
	if err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionInstancesWithContext(ctx context.Context, input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesInput, opts ...request.Option) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionInstancesOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptions(input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionsPages(input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, fn func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput{}
	fnCacher := func(out *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, more bool) bool {
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
	if err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionsPagesWithContext(ctx context.Context, input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, fn func(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput{}
	fnCacher := func(out *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, more bool) bool {
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
	if err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *LicenseManagerLinuxSubscriptions) ListLinuxSubscriptionsWithContext(ctx context.Context, input *licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsInput, opts ...request.Option) (*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*licensemanagerlinuxsubscriptions.ListLinuxSubscriptionsOutput), nil
	}
	out, err := c.LicenseManagerLinuxSubscriptionsAPI.ListLinuxSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
