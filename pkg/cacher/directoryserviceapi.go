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
	"github.com/aws/aws-sdk-go/service/directoryservice"
	"github.com/aws/aws-sdk-go/service/directoryservice/directoryserviceiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type DirectoryService struct {
	directoryserviceiface.DirectoryServiceAPI
	cache *cache.Cache
}

func NewDirectoryService(directoryserviceapi directoryserviceiface.DirectoryServiceAPI) *DirectoryService {
	return &DirectoryService{
		DirectoryServiceAPI: directoryserviceapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *DirectoryService) DescribeCertificate(input *directoryservice.DescribeCertificateInput) (*directoryservice.DescribeCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeCertificateOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeCertificate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeCertificateWithContext(ctx context.Context, input *directoryservice.DescribeCertificateInput, opts ...request.Option) (*directoryservice.DescribeCertificateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeCertificateOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeCertificateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeClientAuthenticationSettings(input *directoryservice.DescribeClientAuthenticationSettingsInput) (*directoryservice.DescribeClientAuthenticationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeClientAuthenticationSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeClientAuthenticationSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeClientAuthenticationSettingsPages(input *directoryservice.DescribeClientAuthenticationSettingsInput, fn func(*directoryservice.DescribeClientAuthenticationSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeClientAuthenticationSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeClientAuthenticationSettingsOutput{}
	fnCacher := func(out *directoryservice.DescribeClientAuthenticationSettingsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeClientAuthenticationSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeClientAuthenticationSettingsPagesWithContext(ctx context.Context, input *directoryservice.DescribeClientAuthenticationSettingsInput, fn func(*directoryservice.DescribeClientAuthenticationSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeClientAuthenticationSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeClientAuthenticationSettingsOutput{}
	fnCacher := func(out *directoryservice.DescribeClientAuthenticationSettingsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeClientAuthenticationSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeClientAuthenticationSettingsWithContext(ctx context.Context, input *directoryservice.DescribeClientAuthenticationSettingsInput, opts ...request.Option) (*directoryservice.DescribeClientAuthenticationSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeClientAuthenticationSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeClientAuthenticationSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeConditionalForwarders(input *directoryservice.DescribeConditionalForwardersInput) (*directoryservice.DescribeConditionalForwardersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeConditionalForwardersOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeConditionalForwarders(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeConditionalForwardersWithContext(ctx context.Context, input *directoryservice.DescribeConditionalForwardersInput, opts ...request.Option) (*directoryservice.DescribeConditionalForwardersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeConditionalForwardersOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeConditionalForwardersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeDirectories(input *directoryservice.DescribeDirectoriesInput) (*directoryservice.DescribeDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeDirectoriesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeDirectories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeDirectoriesPages(input *directoryservice.DescribeDirectoriesInput, fn func(*directoryservice.DescribeDirectoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeDirectoriesOutput{}
	fnCacher := func(out *directoryservice.DescribeDirectoriesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeDirectoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeDirectoriesPagesWithContext(ctx context.Context, input *directoryservice.DescribeDirectoriesInput, fn func(*directoryservice.DescribeDirectoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeDirectoriesOutput{}
	fnCacher := func(out *directoryservice.DescribeDirectoriesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeDirectoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeDirectoriesWithContext(ctx context.Context, input *directoryservice.DescribeDirectoriesInput, opts ...request.Option) (*directoryservice.DescribeDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeDirectoriesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeDirectoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeDomainControllers(input *directoryservice.DescribeDomainControllersInput) (*directoryservice.DescribeDomainControllersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeDomainControllersOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeDomainControllers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeDomainControllersPages(input *directoryservice.DescribeDomainControllersInput, fn func(*directoryservice.DescribeDomainControllersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeDomainControllersOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeDomainControllersOutput{}
	fnCacher := func(out *directoryservice.DescribeDomainControllersOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeDomainControllersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeDomainControllersPagesWithContext(ctx context.Context, input *directoryservice.DescribeDomainControllersInput, fn func(*directoryservice.DescribeDomainControllersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeDomainControllersOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeDomainControllersOutput{}
	fnCacher := func(out *directoryservice.DescribeDomainControllersOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeDomainControllersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeDomainControllersWithContext(ctx context.Context, input *directoryservice.DescribeDomainControllersInput, opts ...request.Option) (*directoryservice.DescribeDomainControllersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeDomainControllersOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeDomainControllersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeEventTopics(input *directoryservice.DescribeEventTopicsInput) (*directoryservice.DescribeEventTopicsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeEventTopicsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeEventTopics(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeEventTopicsWithContext(ctx context.Context, input *directoryservice.DescribeEventTopicsInput, opts ...request.Option) (*directoryservice.DescribeEventTopicsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeEventTopicsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeEventTopicsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeLDAPSSettings(input *directoryservice.DescribeLDAPSSettingsInput) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeLDAPSSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeLDAPSSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeLDAPSSettingsPages(input *directoryservice.DescribeLDAPSSettingsInput, fn func(*directoryservice.DescribeLDAPSSettingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeLDAPSSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeLDAPSSettingsOutput{}
	fnCacher := func(out *directoryservice.DescribeLDAPSSettingsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeLDAPSSettingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeLDAPSSettingsPagesWithContext(ctx context.Context, input *directoryservice.DescribeLDAPSSettingsInput, fn func(*directoryservice.DescribeLDAPSSettingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeLDAPSSettingsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeLDAPSSettingsOutput{}
	fnCacher := func(out *directoryservice.DescribeLDAPSSettingsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeLDAPSSettingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeLDAPSSettingsWithContext(ctx context.Context, input *directoryservice.DescribeLDAPSSettingsInput, opts ...request.Option) (*directoryservice.DescribeLDAPSSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeLDAPSSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeLDAPSSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeRegions(input *directoryservice.DescribeRegionsInput) (*directoryservice.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeRegionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeRegions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeRegionsPages(input *directoryservice.DescribeRegionsInput, fn func(*directoryservice.DescribeRegionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeRegionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeRegionsOutput{}
	fnCacher := func(out *directoryservice.DescribeRegionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeRegionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeRegionsPagesWithContext(ctx context.Context, input *directoryservice.DescribeRegionsInput, fn func(*directoryservice.DescribeRegionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeRegionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeRegionsOutput{}
	fnCacher := func(out *directoryservice.DescribeRegionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeRegionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeRegionsWithContext(ctx context.Context, input *directoryservice.DescribeRegionsInput, opts ...request.Option) (*directoryservice.DescribeRegionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeRegionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeRegionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSettings(input *directoryservice.DescribeSettingsInput) (*directoryservice.DescribeSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSettingsWithContext(ctx context.Context, input *directoryservice.DescribeSettingsInput, opts ...request.Option) (*directoryservice.DescribeSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSettingsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSharedDirectories(input *directoryservice.DescribeSharedDirectoriesInput) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSharedDirectoriesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSharedDirectories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSharedDirectoriesPages(input *directoryservice.DescribeSharedDirectoriesInput, fn func(*directoryservice.DescribeSharedDirectoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeSharedDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeSharedDirectoriesOutput{}
	fnCacher := func(out *directoryservice.DescribeSharedDirectoriesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeSharedDirectoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeSharedDirectoriesPagesWithContext(ctx context.Context, input *directoryservice.DescribeSharedDirectoriesInput, fn func(*directoryservice.DescribeSharedDirectoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeSharedDirectoriesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeSharedDirectoriesOutput{}
	fnCacher := func(out *directoryservice.DescribeSharedDirectoriesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeSharedDirectoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeSharedDirectoriesWithContext(ctx context.Context, input *directoryservice.DescribeSharedDirectoriesInput, opts ...request.Option) (*directoryservice.DescribeSharedDirectoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSharedDirectoriesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSharedDirectoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSnapshots(input *directoryservice.DescribeSnapshotsInput) (*directoryservice.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSnapshotsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSnapshots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeSnapshotsPages(input *directoryservice.DescribeSnapshotsInput, fn func(*directoryservice.DescribeSnapshotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeSnapshotsOutput{}
	fnCacher := func(out *directoryservice.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeSnapshotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeSnapshotsPagesWithContext(ctx context.Context, input *directoryservice.DescribeSnapshotsInput, fn func(*directoryservice.DescribeSnapshotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeSnapshotsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeSnapshotsOutput{}
	fnCacher := func(out *directoryservice.DescribeSnapshotsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeSnapshotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeSnapshotsWithContext(ctx context.Context, input *directoryservice.DescribeSnapshotsInput, opts ...request.Option) (*directoryservice.DescribeSnapshotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeSnapshotsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeSnapshotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeTrusts(input *directoryservice.DescribeTrustsInput) (*directoryservice.DescribeTrustsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeTrustsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeTrusts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeTrustsPages(input *directoryservice.DescribeTrustsInput, fn func(*directoryservice.DescribeTrustsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeTrustsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeTrustsOutput{}
	fnCacher := func(out *directoryservice.DescribeTrustsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeTrustsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeTrustsPagesWithContext(ctx context.Context, input *directoryservice.DescribeTrustsInput, fn func(*directoryservice.DescribeTrustsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeTrustsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeTrustsOutput{}
	fnCacher := func(out *directoryservice.DescribeTrustsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeTrustsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeTrustsWithContext(ctx context.Context, input *directoryservice.DescribeTrustsInput, opts ...request.Option) (*directoryservice.DescribeTrustsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeTrustsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeTrustsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeUpdateDirectory(input *directoryservice.DescribeUpdateDirectoryInput) (*directoryservice.DescribeUpdateDirectoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeUpdateDirectoryOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeUpdateDirectory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) DescribeUpdateDirectoryPages(input *directoryservice.DescribeUpdateDirectoryInput, fn func(*directoryservice.DescribeUpdateDirectoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeUpdateDirectoryOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeUpdateDirectoryOutput{}
	fnCacher := func(out *directoryservice.DescribeUpdateDirectoryOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeUpdateDirectoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeUpdateDirectoryPagesWithContext(ctx context.Context, input *directoryservice.DescribeUpdateDirectoryInput, fn func(*directoryservice.DescribeUpdateDirectoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.DescribeUpdateDirectoryOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.DescribeUpdateDirectoryOutput{}
	fnCacher := func(out *directoryservice.DescribeUpdateDirectoryOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.DescribeUpdateDirectoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) DescribeUpdateDirectoryWithContext(ctx context.Context, input *directoryservice.DescribeUpdateDirectoryInput, opts ...request.Option) (*directoryservice.DescribeUpdateDirectoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.DescribeUpdateDirectoryOutput), nil
	}
	out, err := c.DirectoryServiceAPI.DescribeUpdateDirectoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) GetDirectoryLimits(input *directoryservice.GetDirectoryLimitsInput) (*directoryservice.GetDirectoryLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.GetDirectoryLimitsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.GetDirectoryLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) GetDirectoryLimitsWithContext(ctx context.Context, input *directoryservice.GetDirectoryLimitsInput, opts ...request.Option) (*directoryservice.GetDirectoryLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.GetDirectoryLimitsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.GetDirectoryLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) GetSnapshotLimits(input *directoryservice.GetSnapshotLimitsInput) (*directoryservice.GetSnapshotLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.GetSnapshotLimitsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.GetSnapshotLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) GetSnapshotLimitsWithContext(ctx context.Context, input *directoryservice.GetSnapshotLimitsInput, opts ...request.Option) (*directoryservice.GetSnapshotLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.GetSnapshotLimitsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.GetSnapshotLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListCertificates(input *directoryservice.ListCertificatesInput) (*directoryservice.ListCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListCertificatesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListCertificates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListCertificatesPages(input *directoryservice.ListCertificatesInput, fn func(*directoryservice.ListCertificatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListCertificatesOutput{}
	fnCacher := func(out *directoryservice.ListCertificatesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListCertificatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListCertificatesPagesWithContext(ctx context.Context, input *directoryservice.ListCertificatesInput, fn func(*directoryservice.ListCertificatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListCertificatesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListCertificatesOutput{}
	fnCacher := func(out *directoryservice.ListCertificatesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListCertificatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListCertificatesWithContext(ctx context.Context, input *directoryservice.ListCertificatesInput, opts ...request.Option) (*directoryservice.ListCertificatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListCertificatesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListCertificatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListIpRoutes(input *directoryservice.ListIpRoutesInput) (*directoryservice.ListIpRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListIpRoutesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListIpRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListIpRoutesPages(input *directoryservice.ListIpRoutesInput, fn func(*directoryservice.ListIpRoutesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListIpRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListIpRoutesOutput{}
	fnCacher := func(out *directoryservice.ListIpRoutesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListIpRoutesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListIpRoutesPagesWithContext(ctx context.Context, input *directoryservice.ListIpRoutesInput, fn func(*directoryservice.ListIpRoutesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListIpRoutesOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListIpRoutesOutput{}
	fnCacher := func(out *directoryservice.ListIpRoutesOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListIpRoutesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListIpRoutesWithContext(ctx context.Context, input *directoryservice.ListIpRoutesInput, opts ...request.Option) (*directoryservice.ListIpRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListIpRoutesOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListIpRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListLogSubscriptions(input *directoryservice.ListLogSubscriptionsInput) (*directoryservice.ListLogSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListLogSubscriptionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListLogSubscriptions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListLogSubscriptionsPages(input *directoryservice.ListLogSubscriptionsInput, fn func(*directoryservice.ListLogSubscriptionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListLogSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListLogSubscriptionsOutput{}
	fnCacher := func(out *directoryservice.ListLogSubscriptionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListLogSubscriptionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListLogSubscriptionsPagesWithContext(ctx context.Context, input *directoryservice.ListLogSubscriptionsInput, fn func(*directoryservice.ListLogSubscriptionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListLogSubscriptionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListLogSubscriptionsOutput{}
	fnCacher := func(out *directoryservice.ListLogSubscriptionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListLogSubscriptionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListLogSubscriptionsWithContext(ctx context.Context, input *directoryservice.ListLogSubscriptionsInput, opts ...request.Option) (*directoryservice.ListLogSubscriptionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListLogSubscriptionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListLogSubscriptionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListSchemaExtensions(input *directoryservice.ListSchemaExtensionsInput) (*directoryservice.ListSchemaExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListSchemaExtensionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListSchemaExtensions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListSchemaExtensionsPages(input *directoryservice.ListSchemaExtensionsInput, fn func(*directoryservice.ListSchemaExtensionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListSchemaExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListSchemaExtensionsOutput{}
	fnCacher := func(out *directoryservice.ListSchemaExtensionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListSchemaExtensionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListSchemaExtensionsPagesWithContext(ctx context.Context, input *directoryservice.ListSchemaExtensionsInput, fn func(*directoryservice.ListSchemaExtensionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListSchemaExtensionsOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListSchemaExtensionsOutput{}
	fnCacher := func(out *directoryservice.ListSchemaExtensionsOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListSchemaExtensionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListSchemaExtensionsWithContext(ctx context.Context, input *directoryservice.ListSchemaExtensionsInput, opts ...request.Option) (*directoryservice.ListSchemaExtensionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListSchemaExtensionsOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListSchemaExtensionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListTagsForResource(input *directoryservice.ListTagsForResourceInput) (*directoryservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListTagsForResourceOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *DirectoryService) ListTagsForResourcePages(input *directoryservice.ListTagsForResourceInput, fn func(*directoryservice.ListTagsForResourceOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListTagsForResourceOutput{}
	fnCacher := func(out *directoryservice.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListTagsForResourcePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListTagsForResourcePagesWithContext(ctx context.Context, input *directoryservice.ListTagsForResourceInput, fn func(*directoryservice.ListTagsForResourceOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*directoryservice.ListTagsForResourceOutput), false)
		return nil
	}
	cachable := true
	output := &directoryservice.ListTagsForResourceOutput{}
	fnCacher := func(out *directoryservice.ListTagsForResourceOutput, more bool) bool {
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
	if err := c.DirectoryServiceAPI.ListTagsForResourcePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *DirectoryService) ListTagsForResourceWithContext(ctx context.Context, input *directoryservice.ListTagsForResourceInput, opts ...request.Option) (*directoryservice.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*directoryservice.ListTagsForResourceOutput), nil
	}
	out, err := c.DirectoryServiceAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
