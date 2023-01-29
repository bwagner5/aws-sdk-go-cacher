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
package opensearchserverlesscacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/opensearchserverless"
	"github.com/aws/aws-sdk-go/service/opensearchserverless/opensearchserverlessiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type OpenSearchServerless struct {
	opensearchserverlessiface.OpenSearchServerlessAPI
	cache *cache.Cache
}

func New(opensearchserverlessapi opensearchserverlessiface.OpenSearchServerlessAPI) *OpenSearchServerless {
	return &OpenSearchServerless{
		OpenSearchServerlessAPI: opensearchserverlessapi,
		cache:                   cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *OpenSearchServerless) BatchGetCollection(input *opensearchserverless.BatchGetCollectionInput) (*opensearchserverless.BatchGetCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.BatchGetCollectionOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.BatchGetCollection(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) BatchGetCollectionWithContext(ctx context.Context, input *opensearchserverless.BatchGetCollectionInput, opts ...request.Option) (*opensearchserverless.BatchGetCollectionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.BatchGetCollectionOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.BatchGetCollectionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) BatchGetVpcEndpoint(input *opensearchserverless.BatchGetVpcEndpointInput) (*opensearchserverless.BatchGetVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.BatchGetVpcEndpointOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.BatchGetVpcEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) BatchGetVpcEndpointWithContext(ctx context.Context, input *opensearchserverless.BatchGetVpcEndpointInput, opts ...request.Option) (*opensearchserverless.BatchGetVpcEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.BatchGetVpcEndpointOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.BatchGetVpcEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetAccessPolicy(input *opensearchserverless.GetAccessPolicyInput) (*opensearchserverless.GetAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetAccessPolicyOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetAccessPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetAccessPolicyWithContext(ctx context.Context, input *opensearchserverless.GetAccessPolicyInput, opts ...request.Option) (*opensearchserverless.GetAccessPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetAccessPolicyOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetAccessPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetAccountSettings(input *opensearchserverless.GetAccountSettingsInput) (*opensearchserverless.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetAccountSettingsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetAccountSettings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetAccountSettingsWithContext(ctx context.Context, input *opensearchserverless.GetAccountSettingsInput, opts ...request.Option) (*opensearchserverless.GetAccountSettingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetAccountSettingsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetAccountSettingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetPoliciesStats(input *opensearchserverless.GetPoliciesStatsInput) (*opensearchserverless.GetPoliciesStatsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetPoliciesStatsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetPoliciesStats(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetPoliciesStatsWithContext(ctx context.Context, input *opensearchserverless.GetPoliciesStatsInput, opts ...request.Option) (*opensearchserverless.GetPoliciesStatsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetPoliciesStatsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetPoliciesStatsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetSecurityConfig(input *opensearchserverless.GetSecurityConfigInput) (*opensearchserverless.GetSecurityConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetSecurityConfigOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetSecurityConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetSecurityConfigWithContext(ctx context.Context, input *opensearchserverless.GetSecurityConfigInput, opts ...request.Option) (*opensearchserverless.GetSecurityConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetSecurityConfigOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetSecurityConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetSecurityPolicy(input *opensearchserverless.GetSecurityPolicyInput) (*opensearchserverless.GetSecurityPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetSecurityPolicyOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetSecurityPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) GetSecurityPolicyWithContext(ctx context.Context, input *opensearchserverless.GetSecurityPolicyInput, opts ...request.Option) (*opensearchserverless.GetSecurityPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.GetSecurityPolicyOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.GetSecurityPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListAccessPolicies(input *opensearchserverless.ListAccessPoliciesInput) (*opensearchserverless.ListAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListAccessPoliciesOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListAccessPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListAccessPoliciesPages(input *opensearchserverless.ListAccessPoliciesInput, fn func(*opensearchserverless.ListAccessPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListAccessPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListAccessPoliciesOutput{}
	fnCacher := func(out *opensearchserverless.ListAccessPoliciesOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListAccessPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListAccessPoliciesPagesWithContext(ctx context.Context, input *opensearchserverless.ListAccessPoliciesInput, fn func(*opensearchserverless.ListAccessPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListAccessPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListAccessPoliciesOutput{}
	fnCacher := func(out *opensearchserverless.ListAccessPoliciesOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListAccessPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListAccessPoliciesWithContext(ctx context.Context, input *opensearchserverless.ListAccessPoliciesInput, opts ...request.Option) (*opensearchserverless.ListAccessPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListAccessPoliciesOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListAccessPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListCollections(input *opensearchserverless.ListCollectionsInput) (*opensearchserverless.ListCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListCollectionsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListCollections(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListCollectionsPages(input *opensearchserverless.ListCollectionsInput, fn func(*opensearchserverless.ListCollectionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListCollectionsOutput{}
	fnCacher := func(out *opensearchserverless.ListCollectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListCollectionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListCollectionsPagesWithContext(ctx context.Context, input *opensearchserverless.ListCollectionsInput, fn func(*opensearchserverless.ListCollectionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListCollectionsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListCollectionsOutput{}
	fnCacher := func(out *opensearchserverless.ListCollectionsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListCollectionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListCollectionsWithContext(ctx context.Context, input *opensearchserverless.ListCollectionsInput, opts ...request.Option) (*opensearchserverless.ListCollectionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListCollectionsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListCollectionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListSecurityConfigs(input *opensearchserverless.ListSecurityConfigsInput) (*opensearchserverless.ListSecurityConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListSecurityConfigsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListSecurityConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListSecurityConfigsPages(input *opensearchserverless.ListSecurityConfigsInput, fn func(*opensearchserverless.ListSecurityConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListSecurityConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListSecurityConfigsOutput{}
	fnCacher := func(out *opensearchserverless.ListSecurityConfigsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListSecurityConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListSecurityConfigsPagesWithContext(ctx context.Context, input *opensearchserverless.ListSecurityConfigsInput, fn func(*opensearchserverless.ListSecurityConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListSecurityConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListSecurityConfigsOutput{}
	fnCacher := func(out *opensearchserverless.ListSecurityConfigsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListSecurityConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListSecurityConfigsWithContext(ctx context.Context, input *opensearchserverless.ListSecurityConfigsInput, opts ...request.Option) (*opensearchserverless.ListSecurityConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListSecurityConfigsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListSecurityConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListSecurityPolicies(input *opensearchserverless.ListSecurityPoliciesInput) (*opensearchserverless.ListSecurityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListSecurityPoliciesOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListSecurityPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListSecurityPoliciesPages(input *opensearchserverless.ListSecurityPoliciesInput, fn func(*opensearchserverless.ListSecurityPoliciesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListSecurityPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListSecurityPoliciesOutput{}
	fnCacher := func(out *opensearchserverless.ListSecurityPoliciesOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListSecurityPoliciesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListSecurityPoliciesPagesWithContext(ctx context.Context, input *opensearchserverless.ListSecurityPoliciesInput, fn func(*opensearchserverless.ListSecurityPoliciesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListSecurityPoliciesOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListSecurityPoliciesOutput{}
	fnCacher := func(out *opensearchserverless.ListSecurityPoliciesOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListSecurityPoliciesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListSecurityPoliciesWithContext(ctx context.Context, input *opensearchserverless.ListSecurityPoliciesInput, opts ...request.Option) (*opensearchserverless.ListSecurityPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListSecurityPoliciesOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListSecurityPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListTagsForResource(input *opensearchserverless.ListTagsForResourceInput) (*opensearchserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListTagsForResourceWithContext(ctx context.Context, input *opensearchserverless.ListTagsForResourceInput, opts ...request.Option) (*opensearchserverless.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListTagsForResourceOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListVpcEndpoints(input *opensearchserverless.ListVpcEndpointsInput) (*opensearchserverless.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListVpcEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpenSearchServerless) ListVpcEndpointsPages(input *opensearchserverless.ListVpcEndpointsInput, fn func(*opensearchserverless.ListVpcEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListVpcEndpointsOutput{}
	fnCacher := func(out *opensearchserverless.ListVpcEndpointsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListVpcEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListVpcEndpointsPagesWithContext(ctx context.Context, input *opensearchserverless.ListVpcEndpointsInput, fn func(*opensearchserverless.ListVpcEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opensearchserverless.ListVpcEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &opensearchserverless.ListVpcEndpointsOutput{}
	fnCacher := func(out *opensearchserverless.ListVpcEndpointsOutput, more bool) bool {
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
	if err := c.OpenSearchServerlessAPI.ListVpcEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpenSearchServerless) ListVpcEndpointsWithContext(ctx context.Context, input *opensearchserverless.ListVpcEndpointsInput, opts ...request.Option) (*opensearchserverless.ListVpcEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opensearchserverless.ListVpcEndpointsOutput), nil
	}
	out, err := c.OpenSearchServerlessAPI.ListVpcEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
