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
	"github.com/aws/aws-sdk-go/service/ecr"
	"github.com/aws/aws-sdk-go/service/ecr/ecriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ECR struct {
	ecriface.ECRAPI
	cache *cache.Cache
}

func NewECR(ecrapi ecriface.ECRAPI) *ECR {
	return &ECR{
		ECRAPI: ecrapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ECR) BatchCheckLayerAvailability(input *ecr.BatchCheckLayerAvailabilityInput) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchCheckLayerAvailabilityOutput), nil
	}
	out, err := c.ECRAPI.BatchCheckLayerAvailability(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchCheckLayerAvailabilityWithContext(ctx context.Context, input *ecr.BatchCheckLayerAvailabilityInput, opts ...request.Option) (*ecr.BatchCheckLayerAvailabilityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchCheckLayerAvailabilityOutput), nil
	}
	out, err := c.ECRAPI.BatchCheckLayerAvailabilityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchDeleteImage(input *ecr.BatchDeleteImageInput) (*ecr.BatchDeleteImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchDeleteImageOutput), nil
	}
	out, err := c.ECRAPI.BatchDeleteImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchDeleteImageWithContext(ctx context.Context, input *ecr.BatchDeleteImageInput, opts ...request.Option) (*ecr.BatchDeleteImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchDeleteImageOutput), nil
	}
	out, err := c.ECRAPI.BatchDeleteImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchGetImage(input *ecr.BatchGetImageInput) (*ecr.BatchGetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchGetImageOutput), nil
	}
	out, err := c.ECRAPI.BatchGetImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchGetImageWithContext(ctx context.Context, input *ecr.BatchGetImageInput, opts ...request.Option) (*ecr.BatchGetImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchGetImageOutput), nil
	}
	out, err := c.ECRAPI.BatchGetImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchGetRepositoryScanningConfiguration(input *ecr.BatchGetRepositoryScanningConfigurationInput) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchGetRepositoryScanningConfigurationOutput), nil
	}
	out, err := c.ECRAPI.BatchGetRepositoryScanningConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) BatchGetRepositoryScanningConfigurationWithContext(ctx context.Context, input *ecr.BatchGetRepositoryScanningConfigurationInput, opts ...request.Option) (*ecr.BatchGetRepositoryScanningConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.BatchGetRepositoryScanningConfigurationOutput), nil
	}
	out, err := c.ECRAPI.BatchGetRepositoryScanningConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImageReplicationStatus(input *ecr.DescribeImageReplicationStatusInput) (*ecr.DescribeImageReplicationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImageReplicationStatusOutput), nil
	}
	out, err := c.ECRAPI.DescribeImageReplicationStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImageReplicationStatusWithContext(ctx context.Context, input *ecr.DescribeImageReplicationStatusInput, opts ...request.Option) (*ecr.DescribeImageReplicationStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImageReplicationStatusOutput), nil
	}
	out, err := c.ECRAPI.DescribeImageReplicationStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImageScanFindings(input *ecr.DescribeImageScanFindingsInput) (*ecr.DescribeImageScanFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImageScanFindingsOutput), nil
	}
	out, err := c.ECRAPI.DescribeImageScanFindings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImageScanFindingsPages(input *ecr.DescribeImageScanFindingsInput, fn func(*ecr.DescribeImageScanFindingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeImageScanFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeImageScanFindingsOutput{}
	fnCacher := func(out *ecr.DescribeImageScanFindingsOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeImageScanFindingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeImageScanFindingsPagesWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, fn func(*ecr.DescribeImageScanFindingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeImageScanFindingsOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeImageScanFindingsOutput{}
	fnCacher := func(out *ecr.DescribeImageScanFindingsOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeImageScanFindingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeImageScanFindingsWithContext(ctx context.Context, input *ecr.DescribeImageScanFindingsInput, opts ...request.Option) (*ecr.DescribeImageScanFindingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImageScanFindingsOutput), nil
	}
	out, err := c.ECRAPI.DescribeImageScanFindingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImages(input *ecr.DescribeImagesInput) (*ecr.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImagesOutput), nil
	}
	out, err := c.ECRAPI.DescribeImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeImagesPages(input *ecr.DescribeImagesInput, fn func(*ecr.DescribeImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeImagesOutput{}
	fnCacher := func(out *ecr.DescribeImagesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeImagesPagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, fn func(*ecr.DescribeImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeImagesOutput{}
	fnCacher := func(out *ecr.DescribeImagesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeImagesWithContext(ctx context.Context, input *ecr.DescribeImagesInput, opts ...request.Option) (*ecr.DescribeImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeImagesOutput), nil
	}
	out, err := c.ECRAPI.DescribeImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribePullThroughCacheRules(input *ecr.DescribePullThroughCacheRulesInput) (*ecr.DescribePullThroughCacheRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribePullThroughCacheRulesOutput), nil
	}
	out, err := c.ECRAPI.DescribePullThroughCacheRules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribePullThroughCacheRulesPages(input *ecr.DescribePullThroughCacheRulesInput, fn func(*ecr.DescribePullThroughCacheRulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribePullThroughCacheRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribePullThroughCacheRulesOutput{}
	fnCacher := func(out *ecr.DescribePullThroughCacheRulesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribePullThroughCacheRulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribePullThroughCacheRulesPagesWithContext(ctx context.Context, input *ecr.DescribePullThroughCacheRulesInput, fn func(*ecr.DescribePullThroughCacheRulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribePullThroughCacheRulesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribePullThroughCacheRulesOutput{}
	fnCacher := func(out *ecr.DescribePullThroughCacheRulesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribePullThroughCacheRulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribePullThroughCacheRulesWithContext(ctx context.Context, input *ecr.DescribePullThroughCacheRulesInput, opts ...request.Option) (*ecr.DescribePullThroughCacheRulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribePullThroughCacheRulesOutput), nil
	}
	out, err := c.ECRAPI.DescribePullThroughCacheRulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeRegistry(input *ecr.DescribeRegistryInput) (*ecr.DescribeRegistryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeRegistryOutput), nil
	}
	out, err := c.ECRAPI.DescribeRegistry(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeRegistryWithContext(ctx context.Context, input *ecr.DescribeRegistryInput, opts ...request.Option) (*ecr.DescribeRegistryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeRegistryOutput), nil
	}
	out, err := c.ECRAPI.DescribeRegistryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeRepositories(input *ecr.DescribeRepositoriesInput) (*ecr.DescribeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeRepositoriesOutput), nil
	}
	out, err := c.ECRAPI.DescribeRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) DescribeRepositoriesPages(input *ecr.DescribeRepositoriesInput, fn func(*ecr.DescribeRepositoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeRepositoriesOutput{}
	fnCacher := func(out *ecr.DescribeRepositoriesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeRepositoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeRepositoriesPagesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, fn func(*ecr.DescribeRepositoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.DescribeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.DescribeRepositoriesOutput{}
	fnCacher := func(out *ecr.DescribeRepositoriesOutput, more bool) bool {
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
	if err := c.ECRAPI.DescribeRepositoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) DescribeRepositoriesWithContext(ctx context.Context, input *ecr.DescribeRepositoriesInput, opts ...request.Option) (*ecr.DescribeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.DescribeRepositoriesOutput), nil
	}
	out, err := c.ECRAPI.DescribeRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetAuthorizationToken(input *ecr.GetAuthorizationTokenInput) (*ecr.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetAuthorizationTokenOutput), nil
	}
	out, err := c.ECRAPI.GetAuthorizationToken(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetAuthorizationTokenWithContext(ctx context.Context, input *ecr.GetAuthorizationTokenInput, opts ...request.Option) (*ecr.GetAuthorizationTokenOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetAuthorizationTokenOutput), nil
	}
	out, err := c.ECRAPI.GetAuthorizationTokenWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetDownloadUrlForLayer(input *ecr.GetDownloadUrlForLayerInput) (*ecr.GetDownloadUrlForLayerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetDownloadUrlForLayerOutput), nil
	}
	out, err := c.ECRAPI.GetDownloadUrlForLayer(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetDownloadUrlForLayerWithContext(ctx context.Context, input *ecr.GetDownloadUrlForLayerInput, opts ...request.Option) (*ecr.GetDownloadUrlForLayerOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetDownloadUrlForLayerOutput), nil
	}
	out, err := c.ECRAPI.GetDownloadUrlForLayerWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetLifecyclePolicy(input *ecr.GetLifecyclePolicyInput) (*ecr.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetLifecyclePolicyOutput), nil
	}
	out, err := c.ECRAPI.GetLifecyclePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetLifecyclePolicyPreview(input *ecr.GetLifecyclePolicyPreviewInput) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetLifecyclePolicyPreviewOutput), nil
	}
	out, err := c.ECRAPI.GetLifecyclePolicyPreview(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetLifecyclePolicyPreviewPages(input *ecr.GetLifecyclePolicyPreviewInput, fn func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.GetLifecyclePolicyPreviewOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.GetLifecyclePolicyPreviewOutput{}
	fnCacher := func(out *ecr.GetLifecyclePolicyPreviewOutput, more bool) bool {
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
	if err := c.ECRAPI.GetLifecyclePolicyPreviewPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) GetLifecyclePolicyPreviewPagesWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, fn func(*ecr.GetLifecyclePolicyPreviewOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.GetLifecyclePolicyPreviewOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.GetLifecyclePolicyPreviewOutput{}
	fnCacher := func(out *ecr.GetLifecyclePolicyPreviewOutput, more bool) bool {
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
	if err := c.ECRAPI.GetLifecyclePolicyPreviewPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) GetLifecyclePolicyPreviewWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyPreviewInput, opts ...request.Option) (*ecr.GetLifecyclePolicyPreviewOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetLifecyclePolicyPreviewOutput), nil
	}
	out, err := c.ECRAPI.GetLifecyclePolicyPreviewWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetLifecyclePolicyWithContext(ctx context.Context, input *ecr.GetLifecyclePolicyInput, opts ...request.Option) (*ecr.GetLifecyclePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetLifecyclePolicyOutput), nil
	}
	out, err := c.ECRAPI.GetLifecyclePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRegistryPolicy(input *ecr.GetRegistryPolicyInput) (*ecr.GetRegistryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRegistryPolicyOutput), nil
	}
	out, err := c.ECRAPI.GetRegistryPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRegistryPolicyWithContext(ctx context.Context, input *ecr.GetRegistryPolicyInput, opts ...request.Option) (*ecr.GetRegistryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRegistryPolicyOutput), nil
	}
	out, err := c.ECRAPI.GetRegistryPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRegistryScanningConfiguration(input *ecr.GetRegistryScanningConfigurationInput) (*ecr.GetRegistryScanningConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRegistryScanningConfigurationOutput), nil
	}
	out, err := c.ECRAPI.GetRegistryScanningConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRegistryScanningConfigurationWithContext(ctx context.Context, input *ecr.GetRegistryScanningConfigurationInput, opts ...request.Option) (*ecr.GetRegistryScanningConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRegistryScanningConfigurationOutput), nil
	}
	out, err := c.ECRAPI.GetRegistryScanningConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRepositoryPolicy(input *ecr.GetRepositoryPolicyInput) (*ecr.GetRepositoryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRepositoryPolicyOutput), nil
	}
	out, err := c.ECRAPI.GetRepositoryPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) GetRepositoryPolicyWithContext(ctx context.Context, input *ecr.GetRepositoryPolicyInput, opts ...request.Option) (*ecr.GetRepositoryPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.GetRepositoryPolicyOutput), nil
	}
	out, err := c.ECRAPI.GetRepositoryPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) ListImages(input *ecr.ListImagesInput) (*ecr.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.ListImagesOutput), nil
	}
	out, err := c.ECRAPI.ListImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) ListImagesPages(input *ecr.ListImagesInput, fn func(*ecr.ListImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.ListImagesOutput{}
	fnCacher := func(out *ecr.ListImagesOutput, more bool) bool {
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
	if err := c.ECRAPI.ListImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) ListImagesPagesWithContext(ctx context.Context, input *ecr.ListImagesInput, fn func(*ecr.ListImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*ecr.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &ecr.ListImagesOutput{}
	fnCacher := func(out *ecr.ListImagesOutput, more bool) bool {
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
	if err := c.ECRAPI.ListImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *ECR) ListImagesWithContext(ctx context.Context, input *ecr.ListImagesInput, opts ...request.Option) (*ecr.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.ListImagesOutput), nil
	}
	out, err := c.ECRAPI.ListImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) ListTagsForResource(input *ecr.ListTagsForResourceInput) (*ecr.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.ListTagsForResourceOutput), nil
	}
	out, err := c.ECRAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ECR) ListTagsForResourceWithContext(ctx context.Context, input *ecr.ListTagsForResourceInput, opts ...request.Option) (*ecr.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*ecr.ListTagsForResourceOutput), nil
	}
	out, err := c.ECRAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
