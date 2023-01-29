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
package s3controlcacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type S3Control struct {
	s3controliface.S3ControlAPI
	cache *cache.Cache
}

func New(s3controlapi s3controliface.S3ControlAPI) *S3Control {
	return &S3Control{
		S3ControlAPI: s3controlapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *S3Control) DescribeJob(input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.DescribeJobOutput), nil
	}
	out, err := c.S3ControlAPI.DescribeJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) DescribeJobWithContext(ctx context.Context, input *s3control.DescribeJobInput, opts ...request.Option) (*s3control.DescribeJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.DescribeJobOutput), nil
	}
	out, err := c.S3ControlAPI.DescribeJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) DescribeMultiRegionAccessPointOperation(input *s3control.DescribeMultiRegionAccessPointOperationInput) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.DescribeMultiRegionAccessPointOperationOutput), nil
	}
	out, err := c.S3ControlAPI.DescribeMultiRegionAccessPointOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) DescribeMultiRegionAccessPointOperationWithContext(ctx context.Context, input *s3control.DescribeMultiRegionAccessPointOperationInput, opts ...request.Option) (*s3control.DescribeMultiRegionAccessPointOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.DescribeMultiRegionAccessPointOperationOutput), nil
	}
	out, err := c.S3ControlAPI.DescribeMultiRegionAccessPointOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPoint(input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointConfigurationForObjectLambda(input *s3control.GetAccessPointConfigurationForObjectLambdaInput) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointConfigurationForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointConfigurationForObjectLambda(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointConfigurationForObjectLambdaWithContext(ctx context.Context, input *s3control.GetAccessPointConfigurationForObjectLambdaInput, opts ...request.Option) (*s3control.GetAccessPointConfigurationForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointConfigurationForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointConfigurationForObjectLambdaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointForObjectLambda(input *s3control.GetAccessPointForObjectLambdaInput) (*s3control.GetAccessPointForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointForObjectLambda(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointForObjectLambdaWithContext(ctx context.Context, input *s3control.GetAccessPointForObjectLambdaInput, opts ...request.Option) (*s3control.GetAccessPointForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointForObjectLambdaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicy(input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyForObjectLambda(input *s3control.GetAccessPointPolicyForObjectLambdaInput) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyForObjectLambda(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyForObjectLambdaWithContext(ctx context.Context, input *s3control.GetAccessPointPolicyForObjectLambdaInput, opts ...request.Option) (*s3control.GetAccessPointPolicyForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyForObjectLambdaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyStatus(input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyStatusOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyStatusForObjectLambda(input *s3control.GetAccessPointPolicyStatusForObjectLambdaInput) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyStatusForObjectLambda(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyStatusForObjectLambdaWithContext(ctx context.Context, input *s3control.GetAccessPointPolicyStatusForObjectLambdaInput, opts ...request.Option) (*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyStatusForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyStatusForObjectLambdaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyStatusWithContext(ctx context.Context, input *s3control.GetAccessPointPolicyStatusInput, opts ...request.Option) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyStatusOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointPolicyWithContext(ctx context.Context, input *s3control.GetAccessPointPolicyInput, opts ...request.Option) (*s3control.GetAccessPointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetAccessPointWithContext(ctx context.Context, input *s3control.GetAccessPointInput, opts ...request.Option) (*s3control.GetAccessPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetAccessPointOutput), nil
	}
	out, err := c.S3ControlAPI.GetAccessPointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucket(input *s3control.GetBucketInput) (*s3control.GetBucketOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucket(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketLifecycleConfiguration(input *s3control.GetBucketLifecycleConfigurationInput) (*s3control.GetBucketLifecycleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketLifecycleConfigurationOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketLifecycleConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketLifecycleConfigurationWithContext(ctx context.Context, input *s3control.GetBucketLifecycleConfigurationInput, opts ...request.Option) (*s3control.GetBucketLifecycleConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketLifecycleConfigurationOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketLifecycleConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketPolicy(input *s3control.GetBucketPolicyInput) (*s3control.GetBucketPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketPolicyWithContext(ctx context.Context, input *s3control.GetBucketPolicyInput, opts ...request.Option) (*s3control.GetBucketPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketTagging(input *s3control.GetBucketTaggingInput) (*s3control.GetBucketTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketTagging(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketTaggingWithContext(ctx context.Context, input *s3control.GetBucketTaggingInput, opts ...request.Option) (*s3control.GetBucketTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketTaggingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketVersioning(input *s3control.GetBucketVersioningInput) (*s3control.GetBucketVersioningOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketVersioningOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketVersioning(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketVersioningWithContext(ctx context.Context, input *s3control.GetBucketVersioningInput, opts ...request.Option) (*s3control.GetBucketVersioningOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketVersioningOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketVersioningWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetBucketWithContext(ctx context.Context, input *s3control.GetBucketInput, opts ...request.Option) (*s3control.GetBucketOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetBucketOutput), nil
	}
	out, err := c.S3ControlAPI.GetBucketWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetJobTagging(input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetJobTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetJobTagging(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetJobTaggingWithContext(ctx context.Context, input *s3control.GetJobTaggingInput, opts ...request.Option) (*s3control.GetJobTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetJobTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetJobTaggingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPoint(input *s3control.GetMultiRegionAccessPointInput) (*s3control.GetMultiRegionAccessPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointPolicy(input *s3control.GetMultiRegionAccessPointPolicyInput) (*s3control.GetMultiRegionAccessPointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointPolicyStatus(input *s3control.GetMultiRegionAccessPointPolicyStatusInput) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointPolicyStatusOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointPolicyStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointPolicyStatusWithContext(ctx context.Context, input *s3control.GetMultiRegionAccessPointPolicyStatusInput, opts ...request.Option) (*s3control.GetMultiRegionAccessPointPolicyStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointPolicyStatusOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointPolicyStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointPolicyWithContext(ctx context.Context, input *s3control.GetMultiRegionAccessPointPolicyInput, opts ...request.Option) (*s3control.GetMultiRegionAccessPointPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointPolicyOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointRoutes(input *s3control.GetMultiRegionAccessPointRoutesInput) (*s3control.GetMultiRegionAccessPointRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointRoutesOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointRoutes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointRoutesWithContext(ctx context.Context, input *s3control.GetMultiRegionAccessPointRoutesInput, opts ...request.Option) (*s3control.GetMultiRegionAccessPointRoutesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointRoutesOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointRoutesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetMultiRegionAccessPointWithContext(ctx context.Context, input *s3control.GetMultiRegionAccessPointInput, opts ...request.Option) (*s3control.GetMultiRegionAccessPointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetMultiRegionAccessPointOutput), nil
	}
	out, err := c.S3ControlAPI.GetMultiRegionAccessPointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetPublicAccessBlock(input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetPublicAccessBlockOutput), nil
	}
	out, err := c.S3ControlAPI.GetPublicAccessBlock(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetPublicAccessBlockWithContext(ctx context.Context, input *s3control.GetPublicAccessBlockInput, opts ...request.Option) (*s3control.GetPublicAccessBlockOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetPublicAccessBlockOutput), nil
	}
	out, err := c.S3ControlAPI.GetPublicAccessBlockWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetStorageLensConfiguration(input *s3control.GetStorageLensConfigurationInput) (*s3control.GetStorageLensConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetStorageLensConfigurationOutput), nil
	}
	out, err := c.S3ControlAPI.GetStorageLensConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetStorageLensConfigurationTagging(input *s3control.GetStorageLensConfigurationTaggingInput) (*s3control.GetStorageLensConfigurationTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetStorageLensConfigurationTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetStorageLensConfigurationTagging(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetStorageLensConfigurationTaggingWithContext(ctx context.Context, input *s3control.GetStorageLensConfigurationTaggingInput, opts ...request.Option) (*s3control.GetStorageLensConfigurationTaggingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetStorageLensConfigurationTaggingOutput), nil
	}
	out, err := c.S3ControlAPI.GetStorageLensConfigurationTaggingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) GetStorageLensConfigurationWithContext(ctx context.Context, input *s3control.GetStorageLensConfigurationInput, opts ...request.Option) (*s3control.GetStorageLensConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.GetStorageLensConfigurationOutput), nil
	}
	out, err := c.S3ControlAPI.GetStorageLensConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListAccessPoints(input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListAccessPointsOutput), nil
	}
	out, err := c.S3ControlAPI.ListAccessPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListAccessPointsForObjectLambda(input *s3control.ListAccessPointsForObjectLambdaInput) (*s3control.ListAccessPointsForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListAccessPointsForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.ListAccessPointsForObjectLambda(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListAccessPointsForObjectLambdaPages(input *s3control.ListAccessPointsForObjectLambdaInput, fn func(*s3control.ListAccessPointsForObjectLambdaOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListAccessPointsForObjectLambdaOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListAccessPointsForObjectLambdaOutput{}
	fnCacher := func(out *s3control.ListAccessPointsForObjectLambdaOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListAccessPointsForObjectLambdaPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListAccessPointsForObjectLambdaPagesWithContext(ctx context.Context, input *s3control.ListAccessPointsForObjectLambdaInput, fn func(*s3control.ListAccessPointsForObjectLambdaOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListAccessPointsForObjectLambdaOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListAccessPointsForObjectLambdaOutput{}
	fnCacher := func(out *s3control.ListAccessPointsForObjectLambdaOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListAccessPointsForObjectLambdaPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListAccessPointsForObjectLambdaWithContext(ctx context.Context, input *s3control.ListAccessPointsForObjectLambdaInput, opts ...request.Option) (*s3control.ListAccessPointsForObjectLambdaOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListAccessPointsForObjectLambdaOutput), nil
	}
	out, err := c.S3ControlAPI.ListAccessPointsForObjectLambdaWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListAccessPointsPages(input *s3control.ListAccessPointsInput, fn func(*s3control.ListAccessPointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListAccessPointsOutput{}
	fnCacher := func(out *s3control.ListAccessPointsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListAccessPointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListAccessPointsPagesWithContext(ctx context.Context, input *s3control.ListAccessPointsInput, fn func(*s3control.ListAccessPointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListAccessPointsOutput{}
	fnCacher := func(out *s3control.ListAccessPointsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListAccessPointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListAccessPointsWithContext(ctx context.Context, input *s3control.ListAccessPointsInput, opts ...request.Option) (*s3control.ListAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListAccessPointsOutput), nil
	}
	out, err := c.S3ControlAPI.ListAccessPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListJobs(input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListJobsOutput), nil
	}
	out, err := c.S3ControlAPI.ListJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListJobsPages(input *s3control.ListJobsInput, fn func(*s3control.ListJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListJobsOutput{}
	fnCacher := func(out *s3control.ListJobsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListJobsPagesWithContext(ctx context.Context, input *s3control.ListJobsInput, fn func(*s3control.ListJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListJobsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListJobsOutput{}
	fnCacher := func(out *s3control.ListJobsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListJobsWithContext(ctx context.Context, input *s3control.ListJobsInput, opts ...request.Option) (*s3control.ListJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListJobsOutput), nil
	}
	out, err := c.S3ControlAPI.ListJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListMultiRegionAccessPoints(input *s3control.ListMultiRegionAccessPointsInput) (*s3control.ListMultiRegionAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListMultiRegionAccessPointsOutput), nil
	}
	out, err := c.S3ControlAPI.ListMultiRegionAccessPoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListMultiRegionAccessPointsPages(input *s3control.ListMultiRegionAccessPointsInput, fn func(*s3control.ListMultiRegionAccessPointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListMultiRegionAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListMultiRegionAccessPointsOutput{}
	fnCacher := func(out *s3control.ListMultiRegionAccessPointsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListMultiRegionAccessPointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListMultiRegionAccessPointsPagesWithContext(ctx context.Context, input *s3control.ListMultiRegionAccessPointsInput, fn func(*s3control.ListMultiRegionAccessPointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListMultiRegionAccessPointsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListMultiRegionAccessPointsOutput{}
	fnCacher := func(out *s3control.ListMultiRegionAccessPointsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListMultiRegionAccessPointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListMultiRegionAccessPointsWithContext(ctx context.Context, input *s3control.ListMultiRegionAccessPointsInput, opts ...request.Option) (*s3control.ListMultiRegionAccessPointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListMultiRegionAccessPointsOutput), nil
	}
	out, err := c.S3ControlAPI.ListMultiRegionAccessPointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListRegionalBuckets(input *s3control.ListRegionalBucketsInput) (*s3control.ListRegionalBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListRegionalBucketsOutput), nil
	}
	out, err := c.S3ControlAPI.ListRegionalBuckets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListRegionalBucketsPages(input *s3control.ListRegionalBucketsInput, fn func(*s3control.ListRegionalBucketsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListRegionalBucketsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListRegionalBucketsOutput{}
	fnCacher := func(out *s3control.ListRegionalBucketsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListRegionalBucketsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListRegionalBucketsPagesWithContext(ctx context.Context, input *s3control.ListRegionalBucketsInput, fn func(*s3control.ListRegionalBucketsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListRegionalBucketsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListRegionalBucketsOutput{}
	fnCacher := func(out *s3control.ListRegionalBucketsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListRegionalBucketsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListRegionalBucketsWithContext(ctx context.Context, input *s3control.ListRegionalBucketsInput, opts ...request.Option) (*s3control.ListRegionalBucketsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListRegionalBucketsOutput), nil
	}
	out, err := c.S3ControlAPI.ListRegionalBucketsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListStorageLensConfigurations(input *s3control.ListStorageLensConfigurationsInput) (*s3control.ListStorageLensConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListStorageLensConfigurationsOutput), nil
	}
	out, err := c.S3ControlAPI.ListStorageLensConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *S3Control) ListStorageLensConfigurationsPages(input *s3control.ListStorageLensConfigurationsInput, fn func(*s3control.ListStorageLensConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListStorageLensConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListStorageLensConfigurationsOutput{}
	fnCacher := func(out *s3control.ListStorageLensConfigurationsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListStorageLensConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListStorageLensConfigurationsPagesWithContext(ctx context.Context, input *s3control.ListStorageLensConfigurationsInput, fn func(*s3control.ListStorageLensConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*s3control.ListStorageLensConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &s3control.ListStorageLensConfigurationsOutput{}
	fnCacher := func(out *s3control.ListStorageLensConfigurationsOutput, more bool) bool {
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
	if err := c.S3ControlAPI.ListStorageLensConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *S3Control) ListStorageLensConfigurationsWithContext(ctx context.Context, input *s3control.ListStorageLensConfigurationsInput, opts ...request.Option) (*s3control.ListStorageLensConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*s3control.ListStorageLensConfigurationsOutput), nil
	}
	out, err := c.S3ControlAPI.ListStorageLensConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
