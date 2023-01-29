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
package lambdacacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type Lambda struct {
	lambdaiface.LambdaAPI
	cache *cache.Cache
}

func New(lambdaapi lambdaiface.LambdaAPI) *Lambda {
	return &Lambda{
		LambdaAPI: lambdaapi,
		cache:     cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *Lambda) GetAccountSettings(input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetAccountSettingsOutput), nil
	}
	out, err := c.LambdaAPI.GetAccountSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetAccountSettingsWithContext(ctx context.Context, input *lambda.GetAccountSettingsInput, opts ...request.Option) (*lambda.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetAccountSettingsOutput), nil
	}
	out, err := c.LambdaAPI.GetAccountSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetCodeSigningConfig(input *lambda.GetCodeSigningConfigInput) (*lambda.GetCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetCodeSigningConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetCodeSigningConfigWithContext(ctx context.Context, input *lambda.GetCodeSigningConfigInput, opts ...request.Option) (*lambda.GetCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetCodeSigningConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionOutput), nil
	}
	out, err := c.LambdaAPI.GetFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionCodeSigningConfig(input *lambda.GetFunctionCodeSigningConfigInput) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionCodeSigningConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionCodeSigningConfigWithContext(ctx context.Context, input *lambda.GetFunctionCodeSigningConfigInput, opts ...request.Option) (*lambda.GetFunctionCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionCodeSigningConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionConcurrency(input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionConcurrencyOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionConcurrency(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionConcurrencyWithContext(ctx context.Context, input *lambda.GetFunctionConcurrencyInput, opts ...request.Option) (*lambda.GetFunctionConcurrencyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionConcurrencyOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionConcurrencyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionEventInvokeConfig(input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionEventInvokeConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionEventInvokeConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionEventInvokeConfigWithContext(ctx context.Context, input *lambda.GetFunctionEventInvokeConfigInput, opts ...request.Option) (*lambda.GetFunctionEventInvokeConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionEventInvokeConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionEventInvokeConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionUrlConfig(input *lambda.GetFunctionUrlConfigInput) (*lambda.GetFunctionUrlConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionUrlConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionUrlConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionUrlConfigWithContext(ctx context.Context, input *lambda.GetFunctionUrlConfigInput, opts ...request.Option) (*lambda.GetFunctionUrlConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionUrlConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionUrlConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetFunctionWithContext(ctx context.Context, input *lambda.GetFunctionInput, opts ...request.Option) (*lambda.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetFunctionOutput), nil
	}
	out, err := c.LambdaAPI.GetFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersion(input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersionByArn(input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionByArnOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersionByArn(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersionByArnWithContext(ctx context.Context, input *lambda.GetLayerVersionByArnInput, opts ...request.Option) (*lambda.GetLayerVersionByArnOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionByArnOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersionByArnWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersionPolicy(input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionPolicyOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersionPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersionPolicyWithContext(ctx context.Context, input *lambda.GetLayerVersionPolicyInput, opts ...request.Option) (*lambda.GetLayerVersionPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionPolicyOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersionPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetLayerVersionWithContext(ctx context.Context, input *lambda.GetLayerVersionInput, opts ...request.Option) (*lambda.GetLayerVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetLayerVersionOutput), nil
	}
	out, err := c.LambdaAPI.GetLayerVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetPolicy(input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetPolicyOutput), nil
	}
	out, err := c.LambdaAPI.GetPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetPolicyWithContext(ctx context.Context, input *lambda.GetPolicyInput, opts ...request.Option) (*lambda.GetPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetPolicyOutput), nil
	}
	out, err := c.LambdaAPI.GetPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetProvisionedConcurrencyConfig(input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetProvisionedConcurrencyConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetProvisionedConcurrencyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) GetProvisionedConcurrencyConfigWithContext(ctx context.Context, input *lambda.GetProvisionedConcurrencyConfigInput, opts ...request.Option) (*lambda.GetProvisionedConcurrencyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.GetProvisionedConcurrencyConfigOutput), nil
	}
	out, err := c.LambdaAPI.GetProvisionedConcurrencyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListAliases(input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListAliasesOutput), nil
	}
	out, err := c.LambdaAPI.ListAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListAliasesPages(input *lambda.ListAliasesInput, fn func(*lambda.ListAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListAliasesOutput{}
	fnCacher := func(out *lambda.ListAliasesOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListAliasesPagesWithContext(ctx context.Context, input *lambda.ListAliasesInput, fn func(*lambda.ListAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListAliasesOutput{}
	fnCacher := func(out *lambda.ListAliasesOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListAliasesWithContext(ctx context.Context, input *lambda.ListAliasesInput, opts ...request.Option) (*lambda.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListAliasesOutput), nil
	}
	out, err := c.LambdaAPI.ListAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListCodeSigningConfigs(input *lambda.ListCodeSigningConfigsInput) (*lambda.ListCodeSigningConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListCodeSigningConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListCodeSigningConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListCodeSigningConfigsPages(input *lambda.ListCodeSigningConfigsInput, fn func(*lambda.ListCodeSigningConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListCodeSigningConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListCodeSigningConfigsOutput{}
	fnCacher := func(out *lambda.ListCodeSigningConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListCodeSigningConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListCodeSigningConfigsPagesWithContext(ctx context.Context, input *lambda.ListCodeSigningConfigsInput, fn func(*lambda.ListCodeSigningConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListCodeSigningConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListCodeSigningConfigsOutput{}
	fnCacher := func(out *lambda.ListCodeSigningConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListCodeSigningConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListCodeSigningConfigsWithContext(ctx context.Context, input *lambda.ListCodeSigningConfigsInput, opts ...request.Option) (*lambda.ListCodeSigningConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListCodeSigningConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListCodeSigningConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListEventSourceMappings(input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListEventSourceMappingsOutput), nil
	}
	out, err := c.LambdaAPI.ListEventSourceMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListEventSourceMappingsPages(input *lambda.ListEventSourceMappingsInput, fn func(*lambda.ListEventSourceMappingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListEventSourceMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListEventSourceMappingsOutput{}
	fnCacher := func(out *lambda.ListEventSourceMappingsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListEventSourceMappingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListEventSourceMappingsPagesWithContext(ctx context.Context, input *lambda.ListEventSourceMappingsInput, fn func(*lambda.ListEventSourceMappingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListEventSourceMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListEventSourceMappingsOutput{}
	fnCacher := func(out *lambda.ListEventSourceMappingsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListEventSourceMappingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListEventSourceMappingsWithContext(ctx context.Context, input *lambda.ListEventSourceMappingsInput, opts ...request.Option) (*lambda.ListEventSourceMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListEventSourceMappingsOutput), nil
	}
	out, err := c.LambdaAPI.ListEventSourceMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionEventInvokeConfigs(input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionEventInvokeConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionEventInvokeConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionEventInvokeConfigsPages(input *lambda.ListFunctionEventInvokeConfigsInput, fn func(*lambda.ListFunctionEventInvokeConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionEventInvokeConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionEventInvokeConfigsOutput{}
	fnCacher := func(out *lambda.ListFunctionEventInvokeConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionEventInvokeConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionEventInvokeConfigsPagesWithContext(ctx context.Context, input *lambda.ListFunctionEventInvokeConfigsInput, fn func(*lambda.ListFunctionEventInvokeConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionEventInvokeConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionEventInvokeConfigsOutput{}
	fnCacher := func(out *lambda.ListFunctionEventInvokeConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionEventInvokeConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionEventInvokeConfigsWithContext(ctx context.Context, input *lambda.ListFunctionEventInvokeConfigsInput, opts ...request.Option) (*lambda.ListFunctionEventInvokeConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionEventInvokeConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionEventInvokeConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionUrlConfigs(input *lambda.ListFunctionUrlConfigsInput) (*lambda.ListFunctionUrlConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionUrlConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionUrlConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionUrlConfigsPages(input *lambda.ListFunctionUrlConfigsInput, fn func(*lambda.ListFunctionUrlConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionUrlConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionUrlConfigsOutput{}
	fnCacher := func(out *lambda.ListFunctionUrlConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionUrlConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionUrlConfigsPagesWithContext(ctx context.Context, input *lambda.ListFunctionUrlConfigsInput, fn func(*lambda.ListFunctionUrlConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionUrlConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionUrlConfigsOutput{}
	fnCacher := func(out *lambda.ListFunctionUrlConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionUrlConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionUrlConfigsWithContext(ctx context.Context, input *lambda.ListFunctionUrlConfigsInput, opts ...request.Option) (*lambda.ListFunctionUrlConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionUrlConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionUrlConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctions(input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionsByCodeSigningConfig(input *lambda.ListFunctionsByCodeSigningConfigInput) (*lambda.ListFunctionsByCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionsByCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionsByCodeSigningConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionsByCodeSigningConfigPages(input *lambda.ListFunctionsByCodeSigningConfigInput, fn func(*lambda.ListFunctionsByCodeSigningConfigOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionsByCodeSigningConfigOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionsByCodeSigningConfigOutput{}
	fnCacher := func(out *lambda.ListFunctionsByCodeSigningConfigOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionsByCodeSigningConfigPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionsByCodeSigningConfigPagesWithContext(ctx context.Context, input *lambda.ListFunctionsByCodeSigningConfigInput, fn func(*lambda.ListFunctionsByCodeSigningConfigOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionsByCodeSigningConfigOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionsByCodeSigningConfigOutput{}
	fnCacher := func(out *lambda.ListFunctionsByCodeSigningConfigOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionsByCodeSigningConfigPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionsByCodeSigningConfigWithContext(ctx context.Context, input *lambda.ListFunctionsByCodeSigningConfigInput, opts ...request.Option) (*lambda.ListFunctionsByCodeSigningConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionsByCodeSigningConfigOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionsByCodeSigningConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListFunctionsPages(input *lambda.ListFunctionsInput, fn func(*lambda.ListFunctionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionsOutput{}
	fnCacher := func(out *lambda.ListFunctionsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionsPagesWithContext(ctx context.Context, input *lambda.ListFunctionsInput, fn func(*lambda.ListFunctionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListFunctionsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListFunctionsOutput{}
	fnCacher := func(out *lambda.ListFunctionsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListFunctionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListFunctionsWithContext(ctx context.Context, input *lambda.ListFunctionsInput, opts ...request.Option) (*lambda.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListFunctionsOutput), nil
	}
	out, err := c.LambdaAPI.ListFunctionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListLayerVersions(input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListLayerVersionsOutput), nil
	}
	out, err := c.LambdaAPI.ListLayerVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListLayerVersionsPages(input *lambda.ListLayerVersionsInput, fn func(*lambda.ListLayerVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListLayerVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListLayerVersionsOutput{}
	fnCacher := func(out *lambda.ListLayerVersionsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListLayerVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListLayerVersionsPagesWithContext(ctx context.Context, input *lambda.ListLayerVersionsInput, fn func(*lambda.ListLayerVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListLayerVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListLayerVersionsOutput{}
	fnCacher := func(out *lambda.ListLayerVersionsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListLayerVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListLayerVersionsWithContext(ctx context.Context, input *lambda.ListLayerVersionsInput, opts ...request.Option) (*lambda.ListLayerVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListLayerVersionsOutput), nil
	}
	out, err := c.LambdaAPI.ListLayerVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListLayers(input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListLayersOutput), nil
	}
	out, err := c.LambdaAPI.ListLayers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListLayersPages(input *lambda.ListLayersInput, fn func(*lambda.ListLayersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListLayersOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListLayersOutput{}
	fnCacher := func(out *lambda.ListLayersOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListLayersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListLayersPagesWithContext(ctx context.Context, input *lambda.ListLayersInput, fn func(*lambda.ListLayersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListLayersOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListLayersOutput{}
	fnCacher := func(out *lambda.ListLayersOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListLayersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListLayersWithContext(ctx context.Context, input *lambda.ListLayersInput, opts ...request.Option) (*lambda.ListLayersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListLayersOutput), nil
	}
	out, err := c.LambdaAPI.ListLayersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListProvisionedConcurrencyConfigs(input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListProvisionedConcurrencyConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListProvisionedConcurrencyConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListProvisionedConcurrencyConfigsPages(input *lambda.ListProvisionedConcurrencyConfigsInput, fn func(*lambda.ListProvisionedConcurrencyConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListProvisionedConcurrencyConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListProvisionedConcurrencyConfigsOutput{}
	fnCacher := func(out *lambda.ListProvisionedConcurrencyConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListProvisionedConcurrencyConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListProvisionedConcurrencyConfigsPagesWithContext(ctx context.Context, input *lambda.ListProvisionedConcurrencyConfigsInput, fn func(*lambda.ListProvisionedConcurrencyConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListProvisionedConcurrencyConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListProvisionedConcurrencyConfigsOutput{}
	fnCacher := func(out *lambda.ListProvisionedConcurrencyConfigsOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListProvisionedConcurrencyConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListProvisionedConcurrencyConfigsWithContext(ctx context.Context, input *lambda.ListProvisionedConcurrencyConfigsInput, opts ...request.Option) (*lambda.ListProvisionedConcurrencyConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListProvisionedConcurrencyConfigsOutput), nil
	}
	out, err := c.LambdaAPI.ListProvisionedConcurrencyConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListTags(input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListTagsOutput), nil
	}
	out, err := c.LambdaAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListTagsWithContext(ctx context.Context, input *lambda.ListTagsInput, opts ...request.Option) (*lambda.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListTagsOutput), nil
	}
	out, err := c.LambdaAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListVersionsByFunction(input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListVersionsByFunctionOutput), nil
	}
	out, err := c.LambdaAPI.ListVersionsByFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *Lambda) ListVersionsByFunctionPages(input *lambda.ListVersionsByFunctionInput, fn func(*lambda.ListVersionsByFunctionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListVersionsByFunctionOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListVersionsByFunctionOutput{}
	fnCacher := func(out *lambda.ListVersionsByFunctionOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListVersionsByFunctionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListVersionsByFunctionPagesWithContext(ctx context.Context, input *lambda.ListVersionsByFunctionInput, fn func(*lambda.ListVersionsByFunctionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*lambda.ListVersionsByFunctionOutput), false)
		return nil
	}
	cachable := true
	output := &lambda.ListVersionsByFunctionOutput{}
	fnCacher := func(out *lambda.ListVersionsByFunctionOutput, more bool) bool {
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
	if err := c.LambdaAPI.ListVersionsByFunctionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *Lambda) ListVersionsByFunctionWithContext(ctx context.Context, input *lambda.ListVersionsByFunctionInput, opts ...request.Option) (*lambda.ListVersionsByFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*lambda.ListVersionsByFunctionOutput), nil
	}
	out, err := c.LambdaAPI.ListVersionsByFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
