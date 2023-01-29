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
package codeartifactcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"github.com/aws/aws-sdk-go/service/codeartifact/codeartifactiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeArtifact struct {
	codeartifactiface.CodeArtifactAPI
	cache *cache.Cache
}

func New(codeartifactapi codeartifactiface.CodeArtifactAPI) *CodeArtifact {
	return &CodeArtifact{
		CodeArtifactAPI: codeartifactapi,
		cache:           cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeArtifact) DescribeDomain(input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribeDomainOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribeDomainWithContext(ctx context.Context, input *codeartifact.DescribeDomainInput, opts ...request.Option) (*codeartifact.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribeDomainOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribePackage(input *codeartifact.DescribePackageInput) (*codeartifact.DescribePackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribePackageOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribePackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribePackageVersion(input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribePackageVersionOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribePackageVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribePackageVersionWithContext(ctx context.Context, input *codeartifact.DescribePackageVersionInput, opts ...request.Option) (*codeartifact.DescribePackageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribePackageVersionOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribePackageVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribePackageWithContext(ctx context.Context, input *codeartifact.DescribePackageInput, opts ...request.Option) (*codeartifact.DescribePackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribePackageOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribePackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribeRepository(input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribeRepositoryOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribeRepository(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) DescribeRepositoryWithContext(ctx context.Context, input *codeartifact.DescribeRepositoryInput, opts ...request.Option) (*codeartifact.DescribeRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.DescribeRepositoryOutput), nil
	}
	out, err := c.CodeArtifactAPI.DescribeRepositoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetAuthorizationToken(input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetAuthorizationTokenOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetAuthorizationToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetAuthorizationTokenWithContext(ctx context.Context, input *codeartifact.GetAuthorizationTokenInput, opts ...request.Option) (*codeartifact.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetAuthorizationTokenOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetAuthorizationTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetDomainPermissionsPolicy(input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetDomainPermissionsPolicyOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetDomainPermissionsPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetDomainPermissionsPolicyWithContext(ctx context.Context, input *codeartifact.GetDomainPermissionsPolicyInput, opts ...request.Option) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetDomainPermissionsPolicyOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetDomainPermissionsPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetPackageVersionAsset(input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetPackageVersionAssetOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetPackageVersionAsset(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetPackageVersionAssetWithContext(ctx context.Context, input *codeartifact.GetPackageVersionAssetInput, opts ...request.Option) (*codeartifact.GetPackageVersionAssetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetPackageVersionAssetOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetPackageVersionAssetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetPackageVersionReadme(input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetPackageVersionReadmeOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetPackageVersionReadme(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetPackageVersionReadmeWithContext(ctx context.Context, input *codeartifact.GetPackageVersionReadmeInput, opts ...request.Option) (*codeartifact.GetPackageVersionReadmeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetPackageVersionReadmeOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetPackageVersionReadmeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetRepositoryEndpoint(input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetRepositoryEndpointOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetRepositoryEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetRepositoryEndpointWithContext(ctx context.Context, input *codeartifact.GetRepositoryEndpointInput, opts ...request.Option) (*codeartifact.GetRepositoryEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetRepositoryEndpointOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetRepositoryEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetRepositoryPermissionsPolicy(input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetRepositoryPermissionsPolicyOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetRepositoryPermissionsPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) GetRepositoryPermissionsPolicyWithContext(ctx context.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput, opts ...request.Option) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.GetRepositoryPermissionsPolicyOutput), nil
	}
	out, err := c.CodeArtifactAPI.GetRepositoryPermissionsPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListDomains(input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListDomainsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListDomainsPages(input *codeartifact.ListDomainsInput, fn func(*codeartifact.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListDomainsOutput{}
	fnCacher := func(out *codeartifact.ListDomainsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListDomainsPagesWithContext(ctx context.Context, input *codeartifact.ListDomainsInput, fn func(*codeartifact.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListDomainsOutput{}
	fnCacher := func(out *codeartifact.ListDomainsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListDomainsWithContext(ctx context.Context, input *codeartifact.ListDomainsInput, opts ...request.Option) (*codeartifact.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListDomainsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersionAssets(input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionAssetsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersionAssets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersionAssetsPages(input *codeartifact.ListPackageVersionAssetsInput, fn func(*codeartifact.ListPackageVersionAssetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackageVersionAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackageVersionAssetsOutput{}
	fnCacher := func(out *codeartifact.ListPackageVersionAssetsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackageVersionAssetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackageVersionAssetsPagesWithContext(ctx context.Context, input *codeartifact.ListPackageVersionAssetsInput, fn func(*codeartifact.ListPackageVersionAssetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackageVersionAssetsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackageVersionAssetsOutput{}
	fnCacher := func(out *codeartifact.ListPackageVersionAssetsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackageVersionAssetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackageVersionAssetsWithContext(ctx context.Context, input *codeartifact.ListPackageVersionAssetsInput, opts ...request.Option) (*codeartifact.ListPackageVersionAssetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionAssetsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersionAssetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersionDependencies(input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionDependenciesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersionDependencies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersionDependenciesWithContext(ctx context.Context, input *codeartifact.ListPackageVersionDependenciesInput, opts ...request.Option) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionDependenciesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersionDependenciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersions(input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackageVersionsPages(input *codeartifact.ListPackageVersionsInput, fn func(*codeartifact.ListPackageVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackageVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackageVersionsOutput{}
	fnCacher := func(out *codeartifact.ListPackageVersionsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackageVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackageVersionsPagesWithContext(ctx context.Context, input *codeartifact.ListPackageVersionsInput, fn func(*codeartifact.ListPackageVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackageVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackageVersionsOutput{}
	fnCacher := func(out *codeartifact.ListPackageVersionsOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackageVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackageVersionsWithContext(ctx context.Context, input *codeartifact.ListPackageVersionsInput, opts ...request.Option) (*codeartifact.ListPackageVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackageVersionsOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackageVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackages(input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackagesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListPackagesPages(input *codeartifact.ListPackagesInput, fn func(*codeartifact.ListPackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackagesOutput{}
	fnCacher := func(out *codeartifact.ListPackagesOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackagesPagesWithContext(ctx context.Context, input *codeartifact.ListPackagesInput, fn func(*codeartifact.ListPackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListPackagesOutput{}
	fnCacher := func(out *codeartifact.ListPackagesOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListPackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListPackagesWithContext(ctx context.Context, input *codeartifact.ListPackagesInput, opts ...request.Option) (*codeartifact.ListPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListPackagesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListPackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListRepositories(input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListRepositoriesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListRepositoriesInDomain(input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListRepositoriesInDomainOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListRepositoriesInDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListRepositoriesInDomainPages(input *codeartifact.ListRepositoriesInDomainInput, fn func(*codeartifact.ListRepositoriesInDomainOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListRepositoriesInDomainOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListRepositoriesInDomainOutput{}
	fnCacher := func(out *codeartifact.ListRepositoriesInDomainOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListRepositoriesInDomainPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListRepositoriesInDomainPagesWithContext(ctx context.Context, input *codeartifact.ListRepositoriesInDomainInput, fn func(*codeartifact.ListRepositoriesInDomainOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListRepositoriesInDomainOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListRepositoriesInDomainOutput{}
	fnCacher := func(out *codeartifact.ListRepositoriesInDomainOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListRepositoriesInDomainPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListRepositoriesInDomainWithContext(ctx context.Context, input *codeartifact.ListRepositoriesInDomainInput, opts ...request.Option) (*codeartifact.ListRepositoriesInDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListRepositoriesInDomainOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListRepositoriesInDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListRepositoriesPages(input *codeartifact.ListRepositoriesInput, fn func(*codeartifact.ListRepositoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListRepositoriesOutput{}
	fnCacher := func(out *codeartifact.ListRepositoriesOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListRepositoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListRepositoriesPagesWithContext(ctx context.Context, input *codeartifact.ListRepositoriesInput, fn func(*codeartifact.ListRepositoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codeartifact.ListRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &codeartifact.ListRepositoriesOutput{}
	fnCacher := func(out *codeartifact.ListRepositoriesOutput, more bool) bool {
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
	if err := c.CodeArtifactAPI.ListRepositoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeArtifact) ListRepositoriesWithContext(ctx context.Context, input *codeartifact.ListRepositoriesInput, opts ...request.Option) (*codeartifact.ListRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListRepositoriesOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListTagsForResource(input *codeartifact.ListTagsForResourceInput) (*codeartifact.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeArtifact) ListTagsForResourceWithContext(ctx context.Context, input *codeartifact.ListTagsForResourceInput, opts ...request.Option) (*codeartifact.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codeartifact.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeArtifactAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
