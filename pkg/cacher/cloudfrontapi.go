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
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"github.com/aws/aws-sdk-go/service/cloudfront/cloudfrontiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudFront struct {
	cloudfrontiface.CloudFrontAPI
	cache *cache.Cache
}

func NewCloudFront(cloudfrontapi cloudfrontiface.CloudFrontAPI) *CloudFront {
	return &CloudFront{
		CloudFrontAPI: cloudfrontapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudFront) DescribeFunction(input *cloudfront.DescribeFunctionInput) (*cloudfront.DescribeFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.DescribeFunctionOutput), nil
	}
	out, err := c.CloudFrontAPI.DescribeFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) DescribeFunctionWithContext(ctx context.Context, input *cloudfront.DescribeFunctionInput, opts ...request.Option) (*cloudfront.DescribeFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.DescribeFunctionOutput), nil
	}
	out, err := c.CloudFrontAPI.DescribeFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCachePolicy(input *cloudfront.GetCachePolicyInput) (*cloudfront.GetCachePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCachePolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCachePolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCachePolicyConfig(input *cloudfront.GetCachePolicyConfigInput) (*cloudfront.GetCachePolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCachePolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCachePolicyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCachePolicyConfigWithContext(ctx context.Context, input *cloudfront.GetCachePolicyConfigInput, opts ...request.Option) (*cloudfront.GetCachePolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCachePolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCachePolicyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCachePolicyWithContext(ctx context.Context, input *cloudfront.GetCachePolicyInput, opts ...request.Option) (*cloudfront.GetCachePolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCachePolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCachePolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCloudFrontOriginAccessIdentity(input *cloudfront.GetCloudFrontOriginAccessIdentityInput) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCloudFrontOriginAccessIdentityOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCloudFrontOriginAccessIdentity(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCloudFrontOriginAccessIdentityConfig(input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCloudFrontOriginAccessIdentityConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCloudFrontOriginAccessIdentityConfigWithContext(ctx context.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityConfigInput, opts ...request.Option) (*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCloudFrontOriginAccessIdentityConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCloudFrontOriginAccessIdentityConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetCloudFrontOriginAccessIdentityWithContext(ctx context.Context, input *cloudfront.GetCloudFrontOriginAccessIdentityInput, opts ...request.Option) (*cloudfront.GetCloudFrontOriginAccessIdentityOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetCloudFrontOriginAccessIdentityOutput), nil
	}
	out, err := c.CloudFrontAPI.GetCloudFrontOriginAccessIdentityWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetContinuousDeploymentPolicy(input *cloudfront.GetContinuousDeploymentPolicyInput) (*cloudfront.GetContinuousDeploymentPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetContinuousDeploymentPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetContinuousDeploymentPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetContinuousDeploymentPolicyConfig(input *cloudfront.GetContinuousDeploymentPolicyConfigInput) (*cloudfront.GetContinuousDeploymentPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetContinuousDeploymentPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetContinuousDeploymentPolicyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetContinuousDeploymentPolicyConfigWithContext(ctx context.Context, input *cloudfront.GetContinuousDeploymentPolicyConfigInput, opts ...request.Option) (*cloudfront.GetContinuousDeploymentPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetContinuousDeploymentPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetContinuousDeploymentPolicyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetContinuousDeploymentPolicyWithContext(ctx context.Context, input *cloudfront.GetContinuousDeploymentPolicyInput, opts ...request.Option) (*cloudfront.GetContinuousDeploymentPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetContinuousDeploymentPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetContinuousDeploymentPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetDistribution(input *cloudfront.GetDistributionInput) (*cloudfront.GetDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetDistributionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetDistribution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetDistributionConfig(input *cloudfront.GetDistributionConfigInput) (*cloudfront.GetDistributionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetDistributionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetDistributionConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetDistributionConfigWithContext(ctx context.Context, input *cloudfront.GetDistributionConfigInput, opts ...request.Option) (*cloudfront.GetDistributionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetDistributionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetDistributionConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetDistributionWithContext(ctx context.Context, input *cloudfront.GetDistributionInput, opts ...request.Option) (*cloudfront.GetDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetDistributionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetDistributionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryption(input *cloudfront.GetFieldLevelEncryptionInput) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryption(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionConfig(input *cloudfront.GetFieldLevelEncryptionConfigInput) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionConfigWithContext(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionConfigInput, opts ...request.Option) (*cloudfront.GetFieldLevelEncryptionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionProfile(input *cloudfront.GetFieldLevelEncryptionProfileInput) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionProfileOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionProfileConfig(input *cloudfront.GetFieldLevelEncryptionProfileConfigInput) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionProfileConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionProfileConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionProfileConfigWithContext(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionProfileConfigInput, opts ...request.Option) (*cloudfront.GetFieldLevelEncryptionProfileConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionProfileConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionProfileConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionProfileWithContext(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionProfileInput, opts ...request.Option) (*cloudfront.GetFieldLevelEncryptionProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionProfileOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFieldLevelEncryptionWithContext(ctx context.Context, input *cloudfront.GetFieldLevelEncryptionInput, opts ...request.Option) (*cloudfront.GetFieldLevelEncryptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFieldLevelEncryptionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFieldLevelEncryptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFunction(input *cloudfront.GetFunctionInput) (*cloudfront.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFunctionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFunction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetFunctionWithContext(ctx context.Context, input *cloudfront.GetFunctionInput, opts ...request.Option) (*cloudfront.GetFunctionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetFunctionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetFunctionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetInvalidation(input *cloudfront.GetInvalidationInput) (*cloudfront.GetInvalidationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetInvalidationOutput), nil
	}
	out, err := c.CloudFrontAPI.GetInvalidation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetInvalidationWithContext(ctx context.Context, input *cloudfront.GetInvalidationInput, opts ...request.Option) (*cloudfront.GetInvalidationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetInvalidationOutput), nil
	}
	out, err := c.CloudFrontAPI.GetInvalidationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetKeyGroup(input *cloudfront.GetKeyGroupInput) (*cloudfront.GetKeyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetKeyGroupOutput), nil
	}
	out, err := c.CloudFrontAPI.GetKeyGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetKeyGroupConfig(input *cloudfront.GetKeyGroupConfigInput) (*cloudfront.GetKeyGroupConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetKeyGroupConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetKeyGroupConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetKeyGroupConfigWithContext(ctx context.Context, input *cloudfront.GetKeyGroupConfigInput, opts ...request.Option) (*cloudfront.GetKeyGroupConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetKeyGroupConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetKeyGroupConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetKeyGroupWithContext(ctx context.Context, input *cloudfront.GetKeyGroupInput, opts ...request.Option) (*cloudfront.GetKeyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetKeyGroupOutput), nil
	}
	out, err := c.CloudFrontAPI.GetKeyGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetMonitoringSubscription(input *cloudfront.GetMonitoringSubscriptionInput) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetMonitoringSubscriptionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetMonitoringSubscription(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetMonitoringSubscriptionWithContext(ctx context.Context, input *cloudfront.GetMonitoringSubscriptionInput, opts ...request.Option) (*cloudfront.GetMonitoringSubscriptionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetMonitoringSubscriptionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetMonitoringSubscriptionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginAccessControl(input *cloudfront.GetOriginAccessControlInput) (*cloudfront.GetOriginAccessControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginAccessControlOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginAccessControl(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginAccessControlConfig(input *cloudfront.GetOriginAccessControlConfigInput) (*cloudfront.GetOriginAccessControlConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginAccessControlConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginAccessControlConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginAccessControlConfigWithContext(ctx context.Context, input *cloudfront.GetOriginAccessControlConfigInput, opts ...request.Option) (*cloudfront.GetOriginAccessControlConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginAccessControlConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginAccessControlConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginAccessControlWithContext(ctx context.Context, input *cloudfront.GetOriginAccessControlInput, opts ...request.Option) (*cloudfront.GetOriginAccessControlOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginAccessControlOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginAccessControlWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginRequestPolicy(input *cloudfront.GetOriginRequestPolicyInput) (*cloudfront.GetOriginRequestPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginRequestPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginRequestPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginRequestPolicyConfig(input *cloudfront.GetOriginRequestPolicyConfigInput) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginRequestPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginRequestPolicyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginRequestPolicyConfigWithContext(ctx context.Context, input *cloudfront.GetOriginRequestPolicyConfigInput, opts ...request.Option) (*cloudfront.GetOriginRequestPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginRequestPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginRequestPolicyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetOriginRequestPolicyWithContext(ctx context.Context, input *cloudfront.GetOriginRequestPolicyInput, opts ...request.Option) (*cloudfront.GetOriginRequestPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetOriginRequestPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetOriginRequestPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetPublicKey(input *cloudfront.GetPublicKeyInput) (*cloudfront.GetPublicKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetPublicKeyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetPublicKey(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetPublicKeyConfig(input *cloudfront.GetPublicKeyConfigInput) (*cloudfront.GetPublicKeyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetPublicKeyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetPublicKeyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetPublicKeyConfigWithContext(ctx context.Context, input *cloudfront.GetPublicKeyConfigInput, opts ...request.Option) (*cloudfront.GetPublicKeyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetPublicKeyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetPublicKeyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetPublicKeyWithContext(ctx context.Context, input *cloudfront.GetPublicKeyInput, opts ...request.Option) (*cloudfront.GetPublicKeyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetPublicKeyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetPublicKeyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetRealtimeLogConfig(input *cloudfront.GetRealtimeLogConfigInput) (*cloudfront.GetRealtimeLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetRealtimeLogConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetRealtimeLogConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetRealtimeLogConfigWithContext(ctx context.Context, input *cloudfront.GetRealtimeLogConfigInput, opts ...request.Option) (*cloudfront.GetRealtimeLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetRealtimeLogConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetRealtimeLogConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetResponseHeadersPolicy(input *cloudfront.GetResponseHeadersPolicyInput) (*cloudfront.GetResponseHeadersPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetResponseHeadersPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetResponseHeadersPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetResponseHeadersPolicyConfig(input *cloudfront.GetResponseHeadersPolicyConfigInput) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetResponseHeadersPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetResponseHeadersPolicyConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetResponseHeadersPolicyConfigWithContext(ctx context.Context, input *cloudfront.GetResponseHeadersPolicyConfigInput, opts ...request.Option) (*cloudfront.GetResponseHeadersPolicyConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetResponseHeadersPolicyConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetResponseHeadersPolicyConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetResponseHeadersPolicyWithContext(ctx context.Context, input *cloudfront.GetResponseHeadersPolicyInput, opts ...request.Option) (*cloudfront.GetResponseHeadersPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetResponseHeadersPolicyOutput), nil
	}
	out, err := c.CloudFrontAPI.GetResponseHeadersPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetStreamingDistribution(input *cloudfront.GetStreamingDistributionInput) (*cloudfront.GetStreamingDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetStreamingDistributionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetStreamingDistribution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetStreamingDistributionConfig(input *cloudfront.GetStreamingDistributionConfigInput) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetStreamingDistributionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetStreamingDistributionConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetStreamingDistributionConfigWithContext(ctx context.Context, input *cloudfront.GetStreamingDistributionConfigInput, opts ...request.Option) (*cloudfront.GetStreamingDistributionConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetStreamingDistributionConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.GetStreamingDistributionConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) GetStreamingDistributionWithContext(ctx context.Context, input *cloudfront.GetStreamingDistributionInput, opts ...request.Option) (*cloudfront.GetStreamingDistributionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.GetStreamingDistributionOutput), nil
	}
	out, err := c.CloudFrontAPI.GetStreamingDistributionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListCachePolicies(input *cloudfront.ListCachePoliciesInput) (*cloudfront.ListCachePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListCachePoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListCachePolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListCachePoliciesWithContext(ctx context.Context, input *cloudfront.ListCachePoliciesInput, opts ...request.Option) (*cloudfront.ListCachePoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListCachePoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListCachePoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListCloudFrontOriginAccessIdentities(input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListCloudFrontOriginAccessIdentities(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListCloudFrontOriginAccessIdentitiesPages(input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, fn func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListCloudFrontOriginAccessIdentitiesOutput{}
	fnCacher := func(out *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListCloudFrontOriginAccessIdentitiesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListCloudFrontOriginAccessIdentitiesPagesWithContext(ctx context.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, fn func(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListCloudFrontOriginAccessIdentitiesOutput{}
	fnCacher := func(out *cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListCloudFrontOriginAccessIdentitiesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListCloudFrontOriginAccessIdentitiesWithContext(ctx context.Context, input *cloudfront.ListCloudFrontOriginAccessIdentitiesInput, opts ...request.Option) (*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListCloudFrontOriginAccessIdentitiesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListCloudFrontOriginAccessIdentitiesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListConflictingAliases(input *cloudfront.ListConflictingAliasesInput) (*cloudfront.ListConflictingAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListConflictingAliasesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListConflictingAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListConflictingAliasesWithContext(ctx context.Context, input *cloudfront.ListConflictingAliasesInput, opts ...request.Option) (*cloudfront.ListConflictingAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListConflictingAliasesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListConflictingAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListContinuousDeploymentPolicies(input *cloudfront.ListContinuousDeploymentPoliciesInput) (*cloudfront.ListContinuousDeploymentPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListContinuousDeploymentPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListContinuousDeploymentPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListContinuousDeploymentPoliciesWithContext(ctx context.Context, input *cloudfront.ListContinuousDeploymentPoliciesInput, opts ...request.Option) (*cloudfront.ListContinuousDeploymentPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListContinuousDeploymentPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListContinuousDeploymentPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributions(input *cloudfront.ListDistributionsInput) (*cloudfront.ListDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByCachePolicyId(input *cloudfront.ListDistributionsByCachePolicyIdInput) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByCachePolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByCachePolicyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByCachePolicyIdWithContext(ctx context.Context, input *cloudfront.ListDistributionsByCachePolicyIdInput, opts ...request.Option) (*cloudfront.ListDistributionsByCachePolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByCachePolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByCachePolicyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByKeyGroup(input *cloudfront.ListDistributionsByKeyGroupInput) (*cloudfront.ListDistributionsByKeyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByKeyGroupOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByKeyGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByKeyGroupWithContext(ctx context.Context, input *cloudfront.ListDistributionsByKeyGroupInput, opts ...request.Option) (*cloudfront.ListDistributionsByKeyGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByKeyGroupOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByKeyGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByOriginRequestPolicyId(input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByOriginRequestPolicyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByOriginRequestPolicyIdWithContext(ctx context.Context, input *cloudfront.ListDistributionsByOriginRequestPolicyIdInput, opts ...request.Option) (*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByOriginRequestPolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByOriginRequestPolicyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByRealtimeLogConfig(input *cloudfront.ListDistributionsByRealtimeLogConfigInput) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByRealtimeLogConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByRealtimeLogConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByRealtimeLogConfigWithContext(ctx context.Context, input *cloudfront.ListDistributionsByRealtimeLogConfigInput, opts ...request.Option) (*cloudfront.ListDistributionsByRealtimeLogConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByRealtimeLogConfigOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByRealtimeLogConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByResponseHeadersPolicyId(input *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByResponseHeadersPolicyId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByResponseHeadersPolicyIdWithContext(ctx context.Context, input *cloudfront.ListDistributionsByResponseHeadersPolicyIdInput, opts ...request.Option) (*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByResponseHeadersPolicyIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByResponseHeadersPolicyIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByWebACLId(input *cloudfront.ListDistributionsByWebACLIdInput) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByWebACLIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByWebACLId(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsByWebACLIdWithContext(ctx context.Context, input *cloudfront.ListDistributionsByWebACLIdInput, opts ...request.Option) (*cloudfront.ListDistributionsByWebACLIdOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsByWebACLIdOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsByWebACLIdWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListDistributionsPages(input *cloudfront.ListDistributionsInput, fn func(*cloudfront.ListDistributionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListDistributionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListDistributionsOutput{}
	fnCacher := func(out *cloudfront.ListDistributionsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListDistributionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListDistributionsPagesWithContext(ctx context.Context, input *cloudfront.ListDistributionsInput, fn func(*cloudfront.ListDistributionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListDistributionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListDistributionsOutput{}
	fnCacher := func(out *cloudfront.ListDistributionsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListDistributionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListDistributionsWithContext(ctx context.Context, input *cloudfront.ListDistributionsInput, opts ...request.Option) (*cloudfront.ListDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListDistributionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListDistributionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFieldLevelEncryptionConfigs(input *cloudfront.ListFieldLevelEncryptionConfigsInput) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFieldLevelEncryptionConfigsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFieldLevelEncryptionConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFieldLevelEncryptionConfigsWithContext(ctx context.Context, input *cloudfront.ListFieldLevelEncryptionConfigsInput, opts ...request.Option) (*cloudfront.ListFieldLevelEncryptionConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFieldLevelEncryptionConfigsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFieldLevelEncryptionConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFieldLevelEncryptionProfiles(input *cloudfront.ListFieldLevelEncryptionProfilesInput) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFieldLevelEncryptionProfilesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFieldLevelEncryptionProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFieldLevelEncryptionProfilesWithContext(ctx context.Context, input *cloudfront.ListFieldLevelEncryptionProfilesInput, opts ...request.Option) (*cloudfront.ListFieldLevelEncryptionProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFieldLevelEncryptionProfilesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFieldLevelEncryptionProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFunctions(input *cloudfront.ListFunctionsInput) (*cloudfront.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFunctionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFunctions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListFunctionsWithContext(ctx context.Context, input *cloudfront.ListFunctionsInput, opts ...request.Option) (*cloudfront.ListFunctionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListFunctionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListFunctionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListInvalidations(input *cloudfront.ListInvalidationsInput) (*cloudfront.ListInvalidationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListInvalidationsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListInvalidations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListInvalidationsPages(input *cloudfront.ListInvalidationsInput, fn func(*cloudfront.ListInvalidationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListInvalidationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListInvalidationsOutput{}
	fnCacher := func(out *cloudfront.ListInvalidationsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListInvalidationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListInvalidationsPagesWithContext(ctx context.Context, input *cloudfront.ListInvalidationsInput, fn func(*cloudfront.ListInvalidationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListInvalidationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListInvalidationsOutput{}
	fnCacher := func(out *cloudfront.ListInvalidationsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListInvalidationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListInvalidationsWithContext(ctx context.Context, input *cloudfront.ListInvalidationsInput, opts ...request.Option) (*cloudfront.ListInvalidationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListInvalidationsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListInvalidationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListKeyGroups(input *cloudfront.ListKeyGroupsInput) (*cloudfront.ListKeyGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListKeyGroupsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListKeyGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListKeyGroupsWithContext(ctx context.Context, input *cloudfront.ListKeyGroupsInput, opts ...request.Option) (*cloudfront.ListKeyGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListKeyGroupsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListKeyGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListOriginAccessControls(input *cloudfront.ListOriginAccessControlsInput) (*cloudfront.ListOriginAccessControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListOriginAccessControlsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListOriginAccessControls(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListOriginAccessControlsWithContext(ctx context.Context, input *cloudfront.ListOriginAccessControlsInput, opts ...request.Option) (*cloudfront.ListOriginAccessControlsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListOriginAccessControlsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListOriginAccessControlsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListOriginRequestPolicies(input *cloudfront.ListOriginRequestPoliciesInput) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListOriginRequestPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListOriginRequestPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListOriginRequestPoliciesWithContext(ctx context.Context, input *cloudfront.ListOriginRequestPoliciesInput, opts ...request.Option) (*cloudfront.ListOriginRequestPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListOriginRequestPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListOriginRequestPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListPublicKeys(input *cloudfront.ListPublicKeysInput) (*cloudfront.ListPublicKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListPublicKeysOutput), nil
	}
	out, err := c.CloudFrontAPI.ListPublicKeys(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListPublicKeysWithContext(ctx context.Context, input *cloudfront.ListPublicKeysInput, opts ...request.Option) (*cloudfront.ListPublicKeysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListPublicKeysOutput), nil
	}
	out, err := c.CloudFrontAPI.ListPublicKeysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListRealtimeLogConfigs(input *cloudfront.ListRealtimeLogConfigsInput) (*cloudfront.ListRealtimeLogConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListRealtimeLogConfigsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListRealtimeLogConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListRealtimeLogConfigsWithContext(ctx context.Context, input *cloudfront.ListRealtimeLogConfigsInput, opts ...request.Option) (*cloudfront.ListRealtimeLogConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListRealtimeLogConfigsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListRealtimeLogConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListResponseHeadersPolicies(input *cloudfront.ListResponseHeadersPoliciesInput) (*cloudfront.ListResponseHeadersPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListResponseHeadersPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListResponseHeadersPolicies(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListResponseHeadersPoliciesWithContext(ctx context.Context, input *cloudfront.ListResponseHeadersPoliciesInput, opts ...request.Option) (*cloudfront.ListResponseHeadersPoliciesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListResponseHeadersPoliciesOutput), nil
	}
	out, err := c.CloudFrontAPI.ListResponseHeadersPoliciesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListStreamingDistributions(input *cloudfront.ListStreamingDistributionsInput) (*cloudfront.ListStreamingDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListStreamingDistributionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListStreamingDistributions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListStreamingDistributionsPages(input *cloudfront.ListStreamingDistributionsInput, fn func(*cloudfront.ListStreamingDistributionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListStreamingDistributionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListStreamingDistributionsOutput{}
	fnCacher := func(out *cloudfront.ListStreamingDistributionsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListStreamingDistributionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListStreamingDistributionsPagesWithContext(ctx context.Context, input *cloudfront.ListStreamingDistributionsInput, fn func(*cloudfront.ListStreamingDistributionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudfront.ListStreamingDistributionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudfront.ListStreamingDistributionsOutput{}
	fnCacher := func(out *cloudfront.ListStreamingDistributionsOutput, more bool) bool {
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
	if err := c.CloudFrontAPI.ListStreamingDistributionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFront) ListStreamingDistributionsWithContext(ctx context.Context, input *cloudfront.ListStreamingDistributionsInput, opts ...request.Option) (*cloudfront.ListStreamingDistributionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListStreamingDistributionsOutput), nil
	}
	out, err := c.CloudFrontAPI.ListStreamingDistributionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListTagsForResource(input *cloudfront.ListTagsForResourceInput) (*cloudfront.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudFrontAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFront) ListTagsForResourceWithContext(ctx context.Context, input *cloudfront.ListTagsForResourceInput, opts ...request.Option) (*cloudfront.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudfront.ListTagsForResourceOutput), nil
	}
	out, err := c.CloudFrontAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
