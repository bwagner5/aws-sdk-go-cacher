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
	"github.com/aws/aws-sdk-go/service/eks"
	"github.com/aws/aws-sdk-go/service/eks/eksiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EKS struct {
	eksiface.EKSAPI
	cache *cache.Cache
}

func NewEKS(eksapi eksiface.EKSAPI) *EKS {
	return &EKS{
		EKSAPI: eksapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EKS) DescribeAddon(input *eks.DescribeAddonInput) (*eks.DescribeAddonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddon(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeAddonConfiguration(input *eks.DescribeAddonConfigurationInput) (*eks.DescribeAddonConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonConfigurationOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddonConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeAddonConfigurationWithContext(ctx context.Context, input *eks.DescribeAddonConfigurationInput, opts ...request.Option) (*eks.DescribeAddonConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonConfigurationOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddonConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeAddonVersions(input *eks.DescribeAddonVersionsInput) (*eks.DescribeAddonVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonVersionsOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddonVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeAddonVersionsPages(input *eks.DescribeAddonVersionsInput, fn func(*eks.DescribeAddonVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.DescribeAddonVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.DescribeAddonVersionsOutput{}
	fnCacher := func(out *eks.DescribeAddonVersionsOutput, more bool) bool {
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
	if err := c.EKSAPI.DescribeAddonVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) DescribeAddonVersionsPagesWithContext(ctx context.Context, input *eks.DescribeAddonVersionsInput, fn func(*eks.DescribeAddonVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.DescribeAddonVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.DescribeAddonVersionsOutput{}
	fnCacher := func(out *eks.DescribeAddonVersionsOutput, more bool) bool {
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
	if err := c.EKSAPI.DescribeAddonVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) DescribeAddonVersionsWithContext(ctx context.Context, input *eks.DescribeAddonVersionsInput, opts ...request.Option) (*eks.DescribeAddonVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonVersionsOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddonVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeAddonWithContext(ctx context.Context, input *eks.DescribeAddonInput, opts ...request.Option) (*eks.DescribeAddonOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeAddonOutput), nil
	}
	out, err := c.EKSAPI.DescribeAddonWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeCluster(input *eks.DescribeClusterInput) (*eks.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeClusterOutput), nil
	}
	out, err := c.EKSAPI.DescribeCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeClusterWithContext(ctx context.Context, input *eks.DescribeClusterInput, opts ...request.Option) (*eks.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeClusterOutput), nil
	}
	out, err := c.EKSAPI.DescribeClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeFargateProfile(input *eks.DescribeFargateProfileInput) (*eks.DescribeFargateProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeFargateProfileOutput), nil
	}
	out, err := c.EKSAPI.DescribeFargateProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeFargateProfileWithContext(ctx context.Context, input *eks.DescribeFargateProfileInput, opts ...request.Option) (*eks.DescribeFargateProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeFargateProfileOutput), nil
	}
	out, err := c.EKSAPI.DescribeFargateProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeIdentityProviderConfig(input *eks.DescribeIdentityProviderConfigInput) (*eks.DescribeIdentityProviderConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeIdentityProviderConfigOutput), nil
	}
	out, err := c.EKSAPI.DescribeIdentityProviderConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeIdentityProviderConfigWithContext(ctx context.Context, input *eks.DescribeIdentityProviderConfigInput, opts ...request.Option) (*eks.DescribeIdentityProviderConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeIdentityProviderConfigOutput), nil
	}
	out, err := c.EKSAPI.DescribeIdentityProviderConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeNodegroup(input *eks.DescribeNodegroupInput) (*eks.DescribeNodegroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeNodegroupOutput), nil
	}
	out, err := c.EKSAPI.DescribeNodegroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeNodegroupWithContext(ctx context.Context, input *eks.DescribeNodegroupInput, opts ...request.Option) (*eks.DescribeNodegroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeNodegroupOutput), nil
	}
	out, err := c.EKSAPI.DescribeNodegroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeUpdate(input *eks.DescribeUpdateInput) (*eks.DescribeUpdateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeUpdateOutput), nil
	}
	out, err := c.EKSAPI.DescribeUpdate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) DescribeUpdateWithContext(ctx context.Context, input *eks.DescribeUpdateInput, opts ...request.Option) (*eks.DescribeUpdateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.DescribeUpdateOutput), nil
	}
	out, err := c.EKSAPI.DescribeUpdateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListAddons(input *eks.ListAddonsInput) (*eks.ListAddonsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListAddonsOutput), nil
	}
	out, err := c.EKSAPI.ListAddons(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListAddonsPages(input *eks.ListAddonsInput, fn func(*eks.ListAddonsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListAddonsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListAddonsOutput{}
	fnCacher := func(out *eks.ListAddonsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListAddonsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListAddonsPagesWithContext(ctx context.Context, input *eks.ListAddonsInput, fn func(*eks.ListAddonsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListAddonsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListAddonsOutput{}
	fnCacher := func(out *eks.ListAddonsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListAddonsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListAddonsWithContext(ctx context.Context, input *eks.ListAddonsInput, opts ...request.Option) (*eks.ListAddonsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListAddonsOutput), nil
	}
	out, err := c.EKSAPI.ListAddonsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListClusters(input *eks.ListClustersInput) (*eks.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListClustersOutput), nil
	}
	out, err := c.EKSAPI.ListClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListClustersPages(input *eks.ListClustersInput, fn func(*eks.ListClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListClustersOutput{}
	fnCacher := func(out *eks.ListClustersOutput, more bool) bool {
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
	if err := c.EKSAPI.ListClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListClustersPagesWithContext(ctx context.Context, input *eks.ListClustersInput, fn func(*eks.ListClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListClustersOutput{}
	fnCacher := func(out *eks.ListClustersOutput, more bool) bool {
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
	if err := c.EKSAPI.ListClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListClustersWithContext(ctx context.Context, input *eks.ListClustersInput, opts ...request.Option) (*eks.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListClustersOutput), nil
	}
	out, err := c.EKSAPI.ListClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListFargateProfiles(input *eks.ListFargateProfilesInput) (*eks.ListFargateProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListFargateProfilesOutput), nil
	}
	out, err := c.EKSAPI.ListFargateProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListFargateProfilesPages(input *eks.ListFargateProfilesInput, fn func(*eks.ListFargateProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListFargateProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListFargateProfilesOutput{}
	fnCacher := func(out *eks.ListFargateProfilesOutput, more bool) bool {
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
	if err := c.EKSAPI.ListFargateProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListFargateProfilesPagesWithContext(ctx context.Context, input *eks.ListFargateProfilesInput, fn func(*eks.ListFargateProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListFargateProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListFargateProfilesOutput{}
	fnCacher := func(out *eks.ListFargateProfilesOutput, more bool) bool {
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
	if err := c.EKSAPI.ListFargateProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListFargateProfilesWithContext(ctx context.Context, input *eks.ListFargateProfilesInput, opts ...request.Option) (*eks.ListFargateProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListFargateProfilesOutput), nil
	}
	out, err := c.EKSAPI.ListFargateProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListIdentityProviderConfigs(input *eks.ListIdentityProviderConfigsInput) (*eks.ListIdentityProviderConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListIdentityProviderConfigsOutput), nil
	}
	out, err := c.EKSAPI.ListIdentityProviderConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListIdentityProviderConfigsPages(input *eks.ListIdentityProviderConfigsInput, fn func(*eks.ListIdentityProviderConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListIdentityProviderConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListIdentityProviderConfigsOutput{}
	fnCacher := func(out *eks.ListIdentityProviderConfigsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListIdentityProviderConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListIdentityProviderConfigsPagesWithContext(ctx context.Context, input *eks.ListIdentityProviderConfigsInput, fn func(*eks.ListIdentityProviderConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListIdentityProviderConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListIdentityProviderConfigsOutput{}
	fnCacher := func(out *eks.ListIdentityProviderConfigsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListIdentityProviderConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListIdentityProviderConfigsWithContext(ctx context.Context, input *eks.ListIdentityProviderConfigsInput, opts ...request.Option) (*eks.ListIdentityProviderConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListIdentityProviderConfigsOutput), nil
	}
	out, err := c.EKSAPI.ListIdentityProviderConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListNodegroups(input *eks.ListNodegroupsInput) (*eks.ListNodegroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListNodegroupsOutput), nil
	}
	out, err := c.EKSAPI.ListNodegroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListNodegroupsPages(input *eks.ListNodegroupsInput, fn func(*eks.ListNodegroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListNodegroupsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListNodegroupsOutput{}
	fnCacher := func(out *eks.ListNodegroupsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListNodegroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListNodegroupsPagesWithContext(ctx context.Context, input *eks.ListNodegroupsInput, fn func(*eks.ListNodegroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListNodegroupsOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListNodegroupsOutput{}
	fnCacher := func(out *eks.ListNodegroupsOutput, more bool) bool {
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
	if err := c.EKSAPI.ListNodegroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListNodegroupsWithContext(ctx context.Context, input *eks.ListNodegroupsInput, opts ...request.Option) (*eks.ListNodegroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListNodegroupsOutput), nil
	}
	out, err := c.EKSAPI.ListNodegroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListTagsForResource(input *eks.ListTagsForResourceInput) (*eks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListTagsForResourceOutput), nil
	}
	out, err := c.EKSAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListTagsForResourceWithContext(ctx context.Context, input *eks.ListTagsForResourceInput, opts ...request.Option) (*eks.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListTagsForResourceOutput), nil
	}
	out, err := c.EKSAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListUpdates(input *eks.ListUpdatesInput) (*eks.ListUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListUpdatesOutput), nil
	}
	out, err := c.EKSAPI.ListUpdates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EKS) ListUpdatesPages(input *eks.ListUpdatesInput, fn func(*eks.ListUpdatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListUpdatesOutput{}
	fnCacher := func(out *eks.ListUpdatesOutput, more bool) bool {
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
	if err := c.EKSAPI.ListUpdatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListUpdatesPagesWithContext(ctx context.Context, input *eks.ListUpdatesInput, fn func(*eks.ListUpdatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*eks.ListUpdatesOutput), false)
		return nil
	}
	cachable := true
	output := &eks.ListUpdatesOutput{}
	fnCacher := func(out *eks.ListUpdatesOutput, more bool) bool {
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
	if err := c.EKSAPI.ListUpdatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EKS) ListUpdatesWithContext(ctx context.Context, input *eks.ListUpdatesInput, opts ...request.Option) (*eks.ListUpdatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*eks.ListUpdatesOutput), nil
	}
	out, err := c.EKSAPI.ListUpdatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
