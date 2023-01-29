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
package apprunnercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/apprunner"
	"github.com/aws/aws-sdk-go/service/apprunner/apprunneriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type AppRunner struct {
	apprunneriface.AppRunnerAPI
	cache *cache.Cache
}

func New(apprunnerapi apprunneriface.AppRunnerAPI) *AppRunner {
	return &AppRunner{
		AppRunnerAPI: apprunnerapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *AppRunner) DescribeAutoScalingConfiguration(input *apprunner.DescribeAutoScalingConfigurationInput) (*apprunner.DescribeAutoScalingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeAutoScalingConfigurationOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeAutoScalingConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeAutoScalingConfigurationWithContext(ctx context.Context, input *apprunner.DescribeAutoScalingConfigurationInput, opts ...request.Option) (*apprunner.DescribeAutoScalingConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeAutoScalingConfigurationOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeAutoScalingConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeCustomDomains(input *apprunner.DescribeCustomDomainsInput) (*apprunner.DescribeCustomDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeCustomDomainsOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeCustomDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeCustomDomainsPages(input *apprunner.DescribeCustomDomainsInput, fn func(*apprunner.DescribeCustomDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.DescribeCustomDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.DescribeCustomDomainsOutput{}
	fnCacher := func(out *apprunner.DescribeCustomDomainsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.DescribeCustomDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) DescribeCustomDomainsPagesWithContext(ctx context.Context, input *apprunner.DescribeCustomDomainsInput, fn func(*apprunner.DescribeCustomDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.DescribeCustomDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.DescribeCustomDomainsOutput{}
	fnCacher := func(out *apprunner.DescribeCustomDomainsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.DescribeCustomDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) DescribeCustomDomainsWithContext(ctx context.Context, input *apprunner.DescribeCustomDomainsInput, opts ...request.Option) (*apprunner.DescribeCustomDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeCustomDomainsOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeCustomDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeObservabilityConfiguration(input *apprunner.DescribeObservabilityConfigurationInput) (*apprunner.DescribeObservabilityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeObservabilityConfigurationOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeObservabilityConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeObservabilityConfigurationWithContext(ctx context.Context, input *apprunner.DescribeObservabilityConfigurationInput, opts ...request.Option) (*apprunner.DescribeObservabilityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeObservabilityConfigurationOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeObservabilityConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeService(input *apprunner.DescribeServiceInput) (*apprunner.DescribeServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeServiceOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeService(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeServiceWithContext(ctx context.Context, input *apprunner.DescribeServiceInput, opts ...request.Option) (*apprunner.DescribeServiceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeServiceOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeServiceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeVpcConnector(input *apprunner.DescribeVpcConnectorInput) (*apprunner.DescribeVpcConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeVpcConnectorOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeVpcConnector(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeVpcConnectorWithContext(ctx context.Context, input *apprunner.DescribeVpcConnectorInput, opts ...request.Option) (*apprunner.DescribeVpcConnectorOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeVpcConnectorOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeVpcConnectorWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeVpcIngressConnection(input *apprunner.DescribeVpcIngressConnectionInput) (*apprunner.DescribeVpcIngressConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeVpcIngressConnectionOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeVpcIngressConnection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) DescribeVpcIngressConnectionWithContext(ctx context.Context, input *apprunner.DescribeVpcIngressConnectionInput, opts ...request.Option) (*apprunner.DescribeVpcIngressConnectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.DescribeVpcIngressConnectionOutput), nil
	}
	out, err := c.AppRunnerAPI.DescribeVpcIngressConnectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListAutoScalingConfigurations(input *apprunner.ListAutoScalingConfigurationsInput) (*apprunner.ListAutoScalingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListAutoScalingConfigurationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListAutoScalingConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListAutoScalingConfigurationsPages(input *apprunner.ListAutoScalingConfigurationsInput, fn func(*apprunner.ListAutoScalingConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListAutoScalingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListAutoScalingConfigurationsOutput{}
	fnCacher := func(out *apprunner.ListAutoScalingConfigurationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListAutoScalingConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListAutoScalingConfigurationsPagesWithContext(ctx context.Context, input *apprunner.ListAutoScalingConfigurationsInput, fn func(*apprunner.ListAutoScalingConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListAutoScalingConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListAutoScalingConfigurationsOutput{}
	fnCacher := func(out *apprunner.ListAutoScalingConfigurationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListAutoScalingConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListAutoScalingConfigurationsWithContext(ctx context.Context, input *apprunner.ListAutoScalingConfigurationsInput, opts ...request.Option) (*apprunner.ListAutoScalingConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListAutoScalingConfigurationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListAutoScalingConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListConnections(input *apprunner.ListConnectionsInput) (*apprunner.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListConnectionsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListConnectionsPages(input *apprunner.ListConnectionsInput, fn func(*apprunner.ListConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListConnectionsOutput{}
	fnCacher := func(out *apprunner.ListConnectionsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListConnectionsPagesWithContext(ctx context.Context, input *apprunner.ListConnectionsInput, fn func(*apprunner.ListConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListConnectionsOutput{}
	fnCacher := func(out *apprunner.ListConnectionsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListConnectionsWithContext(ctx context.Context, input *apprunner.ListConnectionsInput, opts ...request.Option) (*apprunner.ListConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListConnectionsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListObservabilityConfigurations(input *apprunner.ListObservabilityConfigurationsInput) (*apprunner.ListObservabilityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListObservabilityConfigurationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListObservabilityConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListObservabilityConfigurationsPages(input *apprunner.ListObservabilityConfigurationsInput, fn func(*apprunner.ListObservabilityConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListObservabilityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListObservabilityConfigurationsOutput{}
	fnCacher := func(out *apprunner.ListObservabilityConfigurationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListObservabilityConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListObservabilityConfigurationsPagesWithContext(ctx context.Context, input *apprunner.ListObservabilityConfigurationsInput, fn func(*apprunner.ListObservabilityConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListObservabilityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListObservabilityConfigurationsOutput{}
	fnCacher := func(out *apprunner.ListObservabilityConfigurationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListObservabilityConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListObservabilityConfigurationsWithContext(ctx context.Context, input *apprunner.ListObservabilityConfigurationsInput, opts ...request.Option) (*apprunner.ListObservabilityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListObservabilityConfigurationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListObservabilityConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListOperations(input *apprunner.ListOperationsInput) (*apprunner.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListOperationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListOperationsPages(input *apprunner.ListOperationsInput, fn func(*apprunner.ListOperationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListOperationsOutput{}
	fnCacher := func(out *apprunner.ListOperationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListOperationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListOperationsPagesWithContext(ctx context.Context, input *apprunner.ListOperationsInput, fn func(*apprunner.ListOperationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListOperationsOutput{}
	fnCacher := func(out *apprunner.ListOperationsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListOperationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListOperationsWithContext(ctx context.Context, input *apprunner.ListOperationsInput, opts ...request.Option) (*apprunner.ListOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListOperationsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListServices(input *apprunner.ListServicesInput) (*apprunner.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListServicesOutput), nil
	}
	out, err := c.AppRunnerAPI.ListServices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListServicesPages(input *apprunner.ListServicesInput, fn func(*apprunner.ListServicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListServicesOutput{}
	fnCacher := func(out *apprunner.ListServicesOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListServicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListServicesPagesWithContext(ctx context.Context, input *apprunner.ListServicesInput, fn func(*apprunner.ListServicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListServicesOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListServicesOutput{}
	fnCacher := func(out *apprunner.ListServicesOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListServicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListServicesWithContext(ctx context.Context, input *apprunner.ListServicesInput, opts ...request.Option) (*apprunner.ListServicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListServicesOutput), nil
	}
	out, err := c.AppRunnerAPI.ListServicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListTagsForResource(input *apprunner.ListTagsForResourceInput) (*apprunner.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListTagsForResourceOutput), nil
	}
	out, err := c.AppRunnerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListTagsForResourceWithContext(ctx context.Context, input *apprunner.ListTagsForResourceInput, opts ...request.Option) (*apprunner.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListTagsForResourceOutput), nil
	}
	out, err := c.AppRunnerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListVpcConnectors(input *apprunner.ListVpcConnectorsInput) (*apprunner.ListVpcConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListVpcConnectorsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListVpcConnectors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListVpcConnectorsPages(input *apprunner.ListVpcConnectorsInput, fn func(*apprunner.ListVpcConnectorsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListVpcConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListVpcConnectorsOutput{}
	fnCacher := func(out *apprunner.ListVpcConnectorsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListVpcConnectorsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListVpcConnectorsPagesWithContext(ctx context.Context, input *apprunner.ListVpcConnectorsInput, fn func(*apprunner.ListVpcConnectorsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListVpcConnectorsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListVpcConnectorsOutput{}
	fnCacher := func(out *apprunner.ListVpcConnectorsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListVpcConnectorsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListVpcConnectorsWithContext(ctx context.Context, input *apprunner.ListVpcConnectorsInput, opts ...request.Option) (*apprunner.ListVpcConnectorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListVpcConnectorsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListVpcConnectorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListVpcIngressConnections(input *apprunner.ListVpcIngressConnectionsInput) (*apprunner.ListVpcIngressConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListVpcIngressConnectionsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListVpcIngressConnections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *AppRunner) ListVpcIngressConnectionsPages(input *apprunner.ListVpcIngressConnectionsInput, fn func(*apprunner.ListVpcIngressConnectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListVpcIngressConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListVpcIngressConnectionsOutput{}
	fnCacher := func(out *apprunner.ListVpcIngressConnectionsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListVpcIngressConnectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListVpcIngressConnectionsPagesWithContext(ctx context.Context, input *apprunner.ListVpcIngressConnectionsInput, fn func(*apprunner.ListVpcIngressConnectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*apprunner.ListVpcIngressConnectionsOutput), false)
		return nil
	}
	cachable := true
	output := &apprunner.ListVpcIngressConnectionsOutput{}
	fnCacher := func(out *apprunner.ListVpcIngressConnectionsOutput, more bool) bool {
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
	if err := c.AppRunnerAPI.ListVpcIngressConnectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *AppRunner) ListVpcIngressConnectionsWithContext(ctx context.Context, input *apprunner.ListVpcIngressConnectionsInput, opts ...request.Option) (*apprunner.ListVpcIngressConnectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*apprunner.ListVpcIngressConnectionsOutput), nil
	}
	out, err := c.AppRunnerAPI.ListVpcIngressConnectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
