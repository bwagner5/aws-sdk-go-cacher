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
package ecrpubliccacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/ecrpublic"
	"github.com/aws/aws-sdk-go/service/ecrpublic/ecrpubliciface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ECRPublic struct {
	ecrpubliciface.ECRPublicAPI
	cache *cache.Cache
}

func New(ecrpublicapi ecrpubliciface.ECRPublicAPI) *ECRPublic {
	return &ECRPublic{
		ECRPublicAPI: ecrpublicapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ECRPublic) BatchCheckLayerAvailability(input *ecrpublic.BatchCheckLayerAvailabilityInput) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.BatchCheckLayerAvailabilityOutput), nil
	}
	out, err := c.ECRPublicAPI.BatchCheckLayerAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) BatchCheckLayerAvailabilityWithContext(ctx context.Context, input *ecrpublic.BatchCheckLayerAvailabilityInput, opts ...request.Option) (*ecrpublic.BatchCheckLayerAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.BatchCheckLayerAvailabilityOutput), nil
	}
	out, err := c.ECRPublicAPI.BatchCheckLayerAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) BatchDeleteImage(input *ecrpublic.BatchDeleteImageInput) (*ecrpublic.BatchDeleteImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.BatchDeleteImageOutput), nil
	}
	out, err := c.ECRPublicAPI.BatchDeleteImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) BatchDeleteImageWithContext(ctx context.Context, input *ecrpublic.BatchDeleteImageInput, opts ...request.Option) (*ecrpublic.BatchDeleteImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.BatchDeleteImageOutput), nil
	}
	out, err := c.ECRPublicAPI.BatchDeleteImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeImageTags(input *ecrpublic.DescribeImageTagsInput) (*ecrpublic.DescribeImageTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeImageTagsOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeImageTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeImageTagsPages(input *ecrpublic.DescribeImageTagsInput, fn func(*ecrpublic.DescribeImageTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeImageTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeImageTagsOutput{}
	fnCacher := func(out *ecrpublic.DescribeImageTagsOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeImageTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeImageTagsPagesWithContext(ctx context.Context, input *ecrpublic.DescribeImageTagsInput, fn func(*ecrpublic.DescribeImageTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeImageTagsOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeImageTagsOutput{}
	fnCacher := func(out *ecrpublic.DescribeImageTagsOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeImageTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeImageTagsWithContext(ctx context.Context, input *ecrpublic.DescribeImageTagsInput, opts ...request.Option) (*ecrpublic.DescribeImageTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeImageTagsOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeImageTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeImages(input *ecrpublic.DescribeImagesInput) (*ecrpublic.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeImagesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeImagesPages(input *ecrpublic.DescribeImagesInput, fn func(*ecrpublic.DescribeImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeImagesOutput{}
	fnCacher := func(out *ecrpublic.DescribeImagesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeImagesPagesWithContext(ctx context.Context, input *ecrpublic.DescribeImagesInput, fn func(*ecrpublic.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeImagesOutput{}
	fnCacher := func(out *ecrpublic.DescribeImagesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeImagesWithContext(ctx context.Context, input *ecrpublic.DescribeImagesInput, opts ...request.Option) (*ecrpublic.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeImagesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeRegistries(input *ecrpublic.DescribeRegistriesInput) (*ecrpublic.DescribeRegistriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeRegistriesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeRegistries(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeRegistriesPages(input *ecrpublic.DescribeRegistriesInput, fn func(*ecrpublic.DescribeRegistriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeRegistriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeRegistriesOutput{}
	fnCacher := func(out *ecrpublic.DescribeRegistriesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeRegistriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeRegistriesPagesWithContext(ctx context.Context, input *ecrpublic.DescribeRegistriesInput, fn func(*ecrpublic.DescribeRegistriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeRegistriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeRegistriesOutput{}
	fnCacher := func(out *ecrpublic.DescribeRegistriesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeRegistriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeRegistriesWithContext(ctx context.Context, input *ecrpublic.DescribeRegistriesInput, opts ...request.Option) (*ecrpublic.DescribeRegistriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeRegistriesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeRegistriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeRepositories(input *ecrpublic.DescribeRepositoriesInput) (*ecrpublic.DescribeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeRepositoriesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) DescribeRepositoriesPages(input *ecrpublic.DescribeRepositoriesInput, fn func(*ecrpublic.DescribeRepositoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeRepositoriesOutput{}
	fnCacher := func(out *ecrpublic.DescribeRepositoriesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeRepositoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeRepositoriesPagesWithContext(ctx context.Context, input *ecrpublic.DescribeRepositoriesInput, fn func(*ecrpublic.DescribeRepositoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecrpublic.DescribeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecrpublic.DescribeRepositoriesOutput{}
	fnCacher := func(out *ecrpublic.DescribeRepositoriesOutput, more bool) bool {
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
	if err := c.ECRPublicAPI.DescribeRepositoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECRPublic) DescribeRepositoriesWithContext(ctx context.Context, input *ecrpublic.DescribeRepositoriesInput, opts ...request.Option) (*ecrpublic.DescribeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.DescribeRepositoriesOutput), nil
	}
	out, err := c.ECRPublicAPI.DescribeRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetAuthorizationToken(input *ecrpublic.GetAuthorizationTokenInput) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetAuthorizationTokenOutput), nil
	}
	out, err := c.ECRPublicAPI.GetAuthorizationToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetAuthorizationTokenWithContext(ctx context.Context, input *ecrpublic.GetAuthorizationTokenInput, opts ...request.Option) (*ecrpublic.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetAuthorizationTokenOutput), nil
	}
	out, err := c.ECRPublicAPI.GetAuthorizationTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRegistryCatalogData(input *ecrpublic.GetRegistryCatalogDataInput) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRegistryCatalogDataOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRegistryCatalogData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRegistryCatalogDataWithContext(ctx context.Context, input *ecrpublic.GetRegistryCatalogDataInput, opts ...request.Option) (*ecrpublic.GetRegistryCatalogDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRegistryCatalogDataOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRegistryCatalogDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRepositoryCatalogData(input *ecrpublic.GetRepositoryCatalogDataInput) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRepositoryCatalogDataOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRepositoryCatalogData(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRepositoryCatalogDataWithContext(ctx context.Context, input *ecrpublic.GetRepositoryCatalogDataInput, opts ...request.Option) (*ecrpublic.GetRepositoryCatalogDataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRepositoryCatalogDataOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRepositoryCatalogDataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRepositoryPolicy(input *ecrpublic.GetRepositoryPolicyInput) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRepositoryPolicyOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRepositoryPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) GetRepositoryPolicyWithContext(ctx context.Context, input *ecrpublic.GetRepositoryPolicyInput, opts ...request.Option) (*ecrpublic.GetRepositoryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.GetRepositoryPolicyOutput), nil
	}
	out, err := c.ECRPublicAPI.GetRepositoryPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) ListTagsForResource(input *ecrpublic.ListTagsForResourceInput) (*ecrpublic.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.ListTagsForResourceOutput), nil
	}
	out, err := c.ECRPublicAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECRPublic) ListTagsForResourceWithContext(ctx context.Context, input *ecrpublic.ListTagsForResourceInput, opts ...request.Option) (*ecrpublic.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecrpublic.ListTagsForResourceOutput), nil
	}
	out, err := c.ECRPublicAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
