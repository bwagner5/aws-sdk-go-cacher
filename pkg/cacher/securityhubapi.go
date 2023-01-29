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
	"github.com/aws/aws-sdk-go/service/securityhub"
	"github.com/aws/aws-sdk-go/service/securityhub/securityhubiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SecurityHub struct {
	securityhubiface.SecurityHubAPI
	cache *cache.Cache
}

func NewSecurityHub(securityhubapi securityhubiface.SecurityHubAPI) *SecurityHub {
	return &SecurityHub{
		SecurityHubAPI: securityhubapi,
		cache:          cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SecurityHub) BatchDisableStandards(input *securityhub.BatchDisableStandardsInput) (*securityhub.BatchDisableStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchDisableStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchDisableStandards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchDisableStandardsWithContext(ctx context.Context, input *securityhub.BatchDisableStandardsInput, opts ...request.Option) (*securityhub.BatchDisableStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchDisableStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchDisableStandardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchEnableStandards(input *securityhub.BatchEnableStandardsInput) (*securityhub.BatchEnableStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchEnableStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchEnableStandards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchEnableStandardsWithContext(ctx context.Context, input *securityhub.BatchEnableStandardsInput, opts ...request.Option) (*securityhub.BatchEnableStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchEnableStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchEnableStandardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchImportFindings(input *securityhub.BatchImportFindingsInput) (*securityhub.BatchImportFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchImportFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchImportFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchImportFindingsWithContext(ctx context.Context, input *securityhub.BatchImportFindingsInput, opts ...request.Option) (*securityhub.BatchImportFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchImportFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchImportFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchUpdateFindings(input *securityhub.BatchUpdateFindingsInput) (*securityhub.BatchUpdateFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchUpdateFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchUpdateFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) BatchUpdateFindingsWithContext(ctx context.Context, input *securityhub.BatchUpdateFindingsInput, opts ...request.Option) (*securityhub.BatchUpdateFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.BatchUpdateFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.BatchUpdateFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeActionTargets(input *securityhub.DescribeActionTargetsInput) (*securityhub.DescribeActionTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeActionTargetsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeActionTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeActionTargetsPages(input *securityhub.DescribeActionTargetsInput, fn func(*securityhub.DescribeActionTargetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeActionTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeActionTargetsOutput{}
	fnCacher := func(out *securityhub.DescribeActionTargetsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeActionTargetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeActionTargetsPagesWithContext(ctx context.Context, input *securityhub.DescribeActionTargetsInput, fn func(*securityhub.DescribeActionTargetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeActionTargetsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeActionTargetsOutput{}
	fnCacher := func(out *securityhub.DescribeActionTargetsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeActionTargetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeActionTargetsWithContext(ctx context.Context, input *securityhub.DescribeActionTargetsInput, opts ...request.Option) (*securityhub.DescribeActionTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeActionTargetsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeActionTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeHub(input *securityhub.DescribeHubInput) (*securityhub.DescribeHubOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeHubOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeHub(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeHubWithContext(ctx context.Context, input *securityhub.DescribeHubInput, opts ...request.Option) (*securityhub.DescribeHubOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeHubOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeHubWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeOrganizationConfiguration(input *securityhub.DescribeOrganizationConfigurationInput) (*securityhub.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeOrganizationConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeOrganizationConfigurationWithContext(ctx context.Context, input *securityhub.DescribeOrganizationConfigurationInput, opts ...request.Option) (*securityhub.DescribeOrganizationConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeOrganizationConfigurationOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeOrganizationConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeProducts(input *securityhub.DescribeProductsInput) (*securityhub.DescribeProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeProductsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeProducts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeProductsPages(input *securityhub.DescribeProductsInput, fn func(*securityhub.DescribeProductsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeProductsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeProductsOutput{}
	fnCacher := func(out *securityhub.DescribeProductsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeProductsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeProductsPagesWithContext(ctx context.Context, input *securityhub.DescribeProductsInput, fn func(*securityhub.DescribeProductsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeProductsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeProductsOutput{}
	fnCacher := func(out *securityhub.DescribeProductsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeProductsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeProductsWithContext(ctx context.Context, input *securityhub.DescribeProductsInput, opts ...request.Option) (*securityhub.DescribeProductsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeProductsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeProductsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeStandards(input *securityhub.DescribeStandardsInput) (*securityhub.DescribeStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeStandards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeStandardsControls(input *securityhub.DescribeStandardsControlsInput) (*securityhub.DescribeStandardsControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeStandardsControlsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeStandardsControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeStandardsControlsPages(input *securityhub.DescribeStandardsControlsInput, fn func(*securityhub.DescribeStandardsControlsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeStandardsControlsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeStandardsControlsOutput{}
	fnCacher := func(out *securityhub.DescribeStandardsControlsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeStandardsControlsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeStandardsControlsPagesWithContext(ctx context.Context, input *securityhub.DescribeStandardsControlsInput, fn func(*securityhub.DescribeStandardsControlsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeStandardsControlsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeStandardsControlsOutput{}
	fnCacher := func(out *securityhub.DescribeStandardsControlsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeStandardsControlsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeStandardsControlsWithContext(ctx context.Context, input *securityhub.DescribeStandardsControlsInput, opts ...request.Option) (*securityhub.DescribeStandardsControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeStandardsControlsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeStandardsControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) DescribeStandardsPages(input *securityhub.DescribeStandardsInput, fn func(*securityhub.DescribeStandardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeStandardsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeStandardsOutput{}
	fnCacher := func(out *securityhub.DescribeStandardsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeStandardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeStandardsPagesWithContext(ctx context.Context, input *securityhub.DescribeStandardsInput, fn func(*securityhub.DescribeStandardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.DescribeStandardsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.DescribeStandardsOutput{}
	fnCacher := func(out *securityhub.DescribeStandardsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.DescribeStandardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) DescribeStandardsWithContext(ctx context.Context, input *securityhub.DescribeStandardsInput, opts ...request.Option) (*securityhub.DescribeStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.DescribeStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.DescribeStandardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetAdministratorAccount(input *securityhub.GetAdministratorAccountInput) (*securityhub.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetAdministratorAccountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetAdministratorAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetAdministratorAccountWithContext(ctx context.Context, input *securityhub.GetAdministratorAccountInput, opts ...request.Option) (*securityhub.GetAdministratorAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetAdministratorAccountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetAdministratorAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetEnabledStandards(input *securityhub.GetEnabledStandardsInput) (*securityhub.GetEnabledStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetEnabledStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetEnabledStandards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetEnabledStandardsPages(input *securityhub.GetEnabledStandardsInput, fn func(*securityhub.GetEnabledStandardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetEnabledStandardsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetEnabledStandardsOutput{}
	fnCacher := func(out *securityhub.GetEnabledStandardsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetEnabledStandardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetEnabledStandardsPagesWithContext(ctx context.Context, input *securityhub.GetEnabledStandardsInput, fn func(*securityhub.GetEnabledStandardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetEnabledStandardsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetEnabledStandardsOutput{}
	fnCacher := func(out *securityhub.GetEnabledStandardsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetEnabledStandardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetEnabledStandardsWithContext(ctx context.Context, input *securityhub.GetEnabledStandardsInput, opts ...request.Option) (*securityhub.GetEnabledStandardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetEnabledStandardsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetEnabledStandardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetFindingAggregator(input *securityhub.GetFindingAggregatorInput) (*securityhub.GetFindingAggregatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetFindingAggregatorOutput), nil
	}
	out, err := c.SecurityHubAPI.GetFindingAggregator(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetFindingAggregatorWithContext(ctx context.Context, input *securityhub.GetFindingAggregatorInput, opts ...request.Option) (*securityhub.GetFindingAggregatorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetFindingAggregatorOutput), nil
	}
	out, err := c.SecurityHubAPI.GetFindingAggregatorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetFindings(input *securityhub.GetFindingsInput) (*securityhub.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetFindingsPages(input *securityhub.GetFindingsInput, fn func(*securityhub.GetFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetFindingsOutput{}
	fnCacher := func(out *securityhub.GetFindingsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetFindingsPagesWithContext(ctx context.Context, input *securityhub.GetFindingsInput, fn func(*securityhub.GetFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetFindingsOutput{}
	fnCacher := func(out *securityhub.GetFindingsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetFindingsWithContext(ctx context.Context, input *securityhub.GetFindingsInput, opts ...request.Option) (*securityhub.GetFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetFindingsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInsightResults(input *securityhub.GetInsightResultsInput) (*securityhub.GetInsightResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInsightResultsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInsightResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInsightResultsWithContext(ctx context.Context, input *securityhub.GetInsightResultsInput, opts ...request.Option) (*securityhub.GetInsightResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInsightResultsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInsightResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInsights(input *securityhub.GetInsightsInput) (*securityhub.GetInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInsightsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInsights(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInsightsPages(input *securityhub.GetInsightsInput, fn func(*securityhub.GetInsightsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetInsightsOutput{}
	fnCacher := func(out *securityhub.GetInsightsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetInsightsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetInsightsPagesWithContext(ctx context.Context, input *securityhub.GetInsightsInput, fn func(*securityhub.GetInsightsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.GetInsightsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.GetInsightsOutput{}
	fnCacher := func(out *securityhub.GetInsightsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.GetInsightsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) GetInsightsWithContext(ctx context.Context, input *securityhub.GetInsightsInput, opts ...request.Option) (*securityhub.GetInsightsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInsightsOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInsightsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInvitationsCount(input *securityhub.GetInvitationsCountInput) (*securityhub.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInvitationsCountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInvitationsCount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetInvitationsCountWithContext(ctx context.Context, input *securityhub.GetInvitationsCountInput, opts ...request.Option) (*securityhub.GetInvitationsCountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetInvitationsCountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetInvitationsCountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetMasterAccount(input *securityhub.GetMasterAccountInput) (*securityhub.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetMasterAccountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetMasterAccount(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetMasterAccountWithContext(ctx context.Context, input *securityhub.GetMasterAccountInput, opts ...request.Option) (*securityhub.GetMasterAccountOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetMasterAccountOutput), nil
	}
	out, err := c.SecurityHubAPI.GetMasterAccountWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetMembers(input *securityhub.GetMembersInput) (*securityhub.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetMembersOutput), nil
	}
	out, err := c.SecurityHubAPI.GetMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) GetMembersWithContext(ctx context.Context, input *securityhub.GetMembersInput, opts ...request.Option) (*securityhub.GetMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.GetMembersOutput), nil
	}
	out, err := c.SecurityHubAPI.GetMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListEnabledProductsForImport(input *securityhub.ListEnabledProductsForImportInput) (*securityhub.ListEnabledProductsForImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListEnabledProductsForImportOutput), nil
	}
	out, err := c.SecurityHubAPI.ListEnabledProductsForImport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListEnabledProductsForImportPages(input *securityhub.ListEnabledProductsForImportInput, fn func(*securityhub.ListEnabledProductsForImportOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListEnabledProductsForImportOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListEnabledProductsForImportOutput{}
	fnCacher := func(out *securityhub.ListEnabledProductsForImportOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListEnabledProductsForImportPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListEnabledProductsForImportPagesWithContext(ctx context.Context, input *securityhub.ListEnabledProductsForImportInput, fn func(*securityhub.ListEnabledProductsForImportOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListEnabledProductsForImportOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListEnabledProductsForImportOutput{}
	fnCacher := func(out *securityhub.ListEnabledProductsForImportOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListEnabledProductsForImportPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListEnabledProductsForImportWithContext(ctx context.Context, input *securityhub.ListEnabledProductsForImportInput, opts ...request.Option) (*securityhub.ListEnabledProductsForImportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListEnabledProductsForImportOutput), nil
	}
	out, err := c.SecurityHubAPI.ListEnabledProductsForImportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListFindingAggregators(input *securityhub.ListFindingAggregatorsInput) (*securityhub.ListFindingAggregatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListFindingAggregatorsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListFindingAggregators(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListFindingAggregatorsPages(input *securityhub.ListFindingAggregatorsInput, fn func(*securityhub.ListFindingAggregatorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListFindingAggregatorsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListFindingAggregatorsOutput{}
	fnCacher := func(out *securityhub.ListFindingAggregatorsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListFindingAggregatorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListFindingAggregatorsPagesWithContext(ctx context.Context, input *securityhub.ListFindingAggregatorsInput, fn func(*securityhub.ListFindingAggregatorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListFindingAggregatorsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListFindingAggregatorsOutput{}
	fnCacher := func(out *securityhub.ListFindingAggregatorsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListFindingAggregatorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListFindingAggregatorsWithContext(ctx context.Context, input *securityhub.ListFindingAggregatorsInput, opts ...request.Option) (*securityhub.ListFindingAggregatorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListFindingAggregatorsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListFindingAggregatorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListInvitations(input *securityhub.ListInvitationsInput) (*securityhub.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListInvitationsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListInvitations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListInvitationsPages(input *securityhub.ListInvitationsInput, fn func(*securityhub.ListInvitationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListInvitationsOutput{}
	fnCacher := func(out *securityhub.ListInvitationsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListInvitationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListInvitationsPagesWithContext(ctx context.Context, input *securityhub.ListInvitationsInput, fn func(*securityhub.ListInvitationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListInvitationsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListInvitationsOutput{}
	fnCacher := func(out *securityhub.ListInvitationsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListInvitationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListInvitationsWithContext(ctx context.Context, input *securityhub.ListInvitationsInput, opts ...request.Option) (*securityhub.ListInvitationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListInvitationsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListInvitationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListMembers(input *securityhub.ListMembersInput) (*securityhub.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListMembersOutput), nil
	}
	out, err := c.SecurityHubAPI.ListMembers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListMembersPages(input *securityhub.ListMembersInput, fn func(*securityhub.ListMembersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListMembersOutput{}
	fnCacher := func(out *securityhub.ListMembersOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListMembersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListMembersPagesWithContext(ctx context.Context, input *securityhub.ListMembersInput, fn func(*securityhub.ListMembersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListMembersOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListMembersOutput{}
	fnCacher := func(out *securityhub.ListMembersOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListMembersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListMembersWithContext(ctx context.Context, input *securityhub.ListMembersInput, opts ...request.Option) (*securityhub.ListMembersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListMembersOutput), nil
	}
	out, err := c.SecurityHubAPI.ListMembersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListOrganizationAdminAccounts(input *securityhub.ListOrganizationAdminAccountsInput) (*securityhub.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListOrganizationAdminAccounts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListOrganizationAdminAccountsPages(input *securityhub.ListOrganizationAdminAccountsInput, fn func(*securityhub.ListOrganizationAdminAccountsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *securityhub.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListOrganizationAdminAccountsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListOrganizationAdminAccountsPagesWithContext(ctx context.Context, input *securityhub.ListOrganizationAdminAccountsInput, fn func(*securityhub.ListOrganizationAdminAccountsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*securityhub.ListOrganizationAdminAccountsOutput), false)
		return nil
	}
	cachable := true
	output := &securityhub.ListOrganizationAdminAccountsOutput{}
	fnCacher := func(out *securityhub.ListOrganizationAdminAccountsOutput, more bool) bool {
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
	if err := c.SecurityHubAPI.ListOrganizationAdminAccountsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SecurityHub) ListOrganizationAdminAccountsWithContext(ctx context.Context, input *securityhub.ListOrganizationAdminAccountsInput, opts ...request.Option) (*securityhub.ListOrganizationAdminAccountsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListOrganizationAdminAccountsOutput), nil
	}
	out, err := c.SecurityHubAPI.ListOrganizationAdminAccountsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListTagsForResource(input *securityhub.ListTagsForResourceInput) (*securityhub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListTagsForResourceOutput), nil
	}
	out, err := c.SecurityHubAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SecurityHub) ListTagsForResourceWithContext(ctx context.Context, input *securityhub.ListTagsForResourceInput, opts ...request.Option) (*securityhub.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*securityhub.ListTagsForResourceOutput), nil
	}
	out, err := c.SecurityHubAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
