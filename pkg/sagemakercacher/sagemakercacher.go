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
package sagemakercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sagemaker"
	"github.com/aws/aws-sdk-go/service/sagemaker/sagemakeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type SageMaker struct {
	sagemakeriface.SageMakerAPI
	cache *cache.Cache
}

func New(sagemakerapi sagemakeriface.SageMakerAPI) *SageMaker {
	return &SageMaker{
		SageMakerAPI: sagemakerapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *SageMaker) BatchDescribeModelPackage(input *sagemaker.BatchDescribeModelPackageInput) (*sagemaker.BatchDescribeModelPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.BatchDescribeModelPackageOutput), nil
	}
	out, err := c.SageMakerAPI.BatchDescribeModelPackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) BatchDescribeModelPackageWithContext(ctx context.Context, input *sagemaker.BatchDescribeModelPackageInput, opts ...request.Option) (*sagemaker.BatchDescribeModelPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.BatchDescribeModelPackageOutput), nil
	}
	out, err := c.SageMakerAPI.BatchDescribeModelPackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAction(input *sagemaker.DescribeActionInput) (*sagemaker.DescribeActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeActionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAction(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeActionWithContext(ctx context.Context, input *sagemaker.DescribeActionInput, opts ...request.Option) (*sagemaker.DescribeActionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeActionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeActionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAlgorithm(input *sagemaker.DescribeAlgorithmInput) (*sagemaker.DescribeAlgorithmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAlgorithmOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAlgorithm(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAlgorithmWithContext(ctx context.Context, input *sagemaker.DescribeAlgorithmInput, opts ...request.Option) (*sagemaker.DescribeAlgorithmOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAlgorithmOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAlgorithmWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeApp(input *sagemaker.DescribeAppInput) (*sagemaker.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAppOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeApp(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAppImageConfig(input *sagemaker.DescribeAppImageConfigInput) (*sagemaker.DescribeAppImageConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAppImageConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAppImageConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAppImageConfigWithContext(ctx context.Context, input *sagemaker.DescribeAppImageConfigInput, opts ...request.Option) (*sagemaker.DescribeAppImageConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAppImageConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAppImageConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAppWithContext(ctx context.Context, input *sagemaker.DescribeAppInput, opts ...request.Option) (*sagemaker.DescribeAppOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAppOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAppWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeArtifact(input *sagemaker.DescribeArtifactInput) (*sagemaker.DescribeArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeArtifactOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeArtifact(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeArtifactWithContext(ctx context.Context, input *sagemaker.DescribeArtifactInput, opts ...request.Option) (*sagemaker.DescribeArtifactOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeArtifactOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeArtifactWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAutoMLJob(input *sagemaker.DescribeAutoMLJobInput) (*sagemaker.DescribeAutoMLJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAutoMLJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAutoMLJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeAutoMLJobWithContext(ctx context.Context, input *sagemaker.DescribeAutoMLJobInput, opts ...request.Option) (*sagemaker.DescribeAutoMLJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeAutoMLJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeAutoMLJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeCodeRepository(input *sagemaker.DescribeCodeRepositoryInput) (*sagemaker.DescribeCodeRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeCodeRepositoryOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeCodeRepository(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeCodeRepositoryWithContext(ctx context.Context, input *sagemaker.DescribeCodeRepositoryInput, opts ...request.Option) (*sagemaker.DescribeCodeRepositoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeCodeRepositoryOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeCodeRepositoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeCompilationJob(input *sagemaker.DescribeCompilationJobInput) (*sagemaker.DescribeCompilationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeCompilationJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeCompilationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeCompilationJobWithContext(ctx context.Context, input *sagemaker.DescribeCompilationJobInput, opts ...request.Option) (*sagemaker.DescribeCompilationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeCompilationJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeCompilationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeContext(input *sagemaker.DescribeContextInput) (*sagemaker.DescribeContextOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeContextOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeContext(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeContextWithContext(ctx context.Context, input *sagemaker.DescribeContextInput, opts ...request.Option) (*sagemaker.DescribeContextOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeContextOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeContextWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDataQualityJobDefinition(input *sagemaker.DescribeDataQualityJobDefinitionInput) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDataQualityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDataQualityJobDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDataQualityJobDefinitionWithContext(ctx context.Context, input *sagemaker.DescribeDataQualityJobDefinitionInput, opts ...request.Option) (*sagemaker.DescribeDataQualityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDataQualityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDataQualityJobDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDevice(input *sagemaker.DescribeDeviceInput) (*sagemaker.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDeviceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDevice(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDeviceFleet(input *sagemaker.DescribeDeviceFleetInput) (*sagemaker.DescribeDeviceFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDeviceFleetOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDeviceFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDeviceFleetWithContext(ctx context.Context, input *sagemaker.DescribeDeviceFleetInput, opts ...request.Option) (*sagemaker.DescribeDeviceFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDeviceFleetOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDeviceFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDeviceWithContext(ctx context.Context, input *sagemaker.DescribeDeviceInput, opts ...request.Option) (*sagemaker.DescribeDeviceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDeviceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDeviceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDomain(input *sagemaker.DescribeDomainInput) (*sagemaker.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDomainOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDomain(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeDomainWithContext(ctx context.Context, input *sagemaker.DescribeDomainInput, opts ...request.Option) (*sagemaker.DescribeDomainOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeDomainOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeDomainWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEdgeDeploymentPlan(input *sagemaker.DescribeEdgeDeploymentPlanInput) (*sagemaker.DescribeEdgeDeploymentPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEdgeDeploymentPlanOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEdgeDeploymentPlan(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEdgeDeploymentPlanWithContext(ctx context.Context, input *sagemaker.DescribeEdgeDeploymentPlanInput, opts ...request.Option) (*sagemaker.DescribeEdgeDeploymentPlanOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEdgeDeploymentPlanOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEdgeDeploymentPlanWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEdgePackagingJob(input *sagemaker.DescribeEdgePackagingJobInput) (*sagemaker.DescribeEdgePackagingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEdgePackagingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEdgePackagingJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEdgePackagingJobWithContext(ctx context.Context, input *sagemaker.DescribeEdgePackagingJobInput, opts ...request.Option) (*sagemaker.DescribeEdgePackagingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEdgePackagingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEdgePackagingJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEndpoint(input *sagemaker.DescribeEndpointInput) (*sagemaker.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEndpointOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEndpoint(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEndpointConfig(input *sagemaker.DescribeEndpointConfigInput) (*sagemaker.DescribeEndpointConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEndpointConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEndpointConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEndpointConfigWithContext(ctx context.Context, input *sagemaker.DescribeEndpointConfigInput, opts ...request.Option) (*sagemaker.DescribeEndpointConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEndpointConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEndpointConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeEndpointWithContext(ctx context.Context, input *sagemaker.DescribeEndpointInput, opts ...request.Option) (*sagemaker.DescribeEndpointOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeEndpointOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeEndpointWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeExperiment(input *sagemaker.DescribeExperimentInput) (*sagemaker.DescribeExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeExperimentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeExperiment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeExperimentWithContext(ctx context.Context, input *sagemaker.DescribeExperimentInput, opts ...request.Option) (*sagemaker.DescribeExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeExperimentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeExperimentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFeatureGroup(input *sagemaker.DescribeFeatureGroupInput) (*sagemaker.DescribeFeatureGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFeatureGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFeatureGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFeatureGroupWithContext(ctx context.Context, input *sagemaker.DescribeFeatureGroupInput, opts ...request.Option) (*sagemaker.DescribeFeatureGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFeatureGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFeatureGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFeatureMetadata(input *sagemaker.DescribeFeatureMetadataInput) (*sagemaker.DescribeFeatureMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFeatureMetadataOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFeatureMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFeatureMetadataWithContext(ctx context.Context, input *sagemaker.DescribeFeatureMetadataInput, opts ...request.Option) (*sagemaker.DescribeFeatureMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFeatureMetadataOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFeatureMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFlowDefinition(input *sagemaker.DescribeFlowDefinitionInput) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFlowDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFlowDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeFlowDefinitionWithContext(ctx context.Context, input *sagemaker.DescribeFlowDefinitionInput, opts ...request.Option) (*sagemaker.DescribeFlowDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeFlowDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeFlowDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHub(input *sagemaker.DescribeHubInput) (*sagemaker.DescribeHubOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHubOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHub(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHubContent(input *sagemaker.DescribeHubContentInput) (*sagemaker.DescribeHubContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHubContentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHubContent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHubContentWithContext(ctx context.Context, input *sagemaker.DescribeHubContentInput, opts ...request.Option) (*sagemaker.DescribeHubContentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHubContentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHubContentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHubWithContext(ctx context.Context, input *sagemaker.DescribeHubInput, opts ...request.Option) (*sagemaker.DescribeHubOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHubOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHubWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHumanTaskUi(input *sagemaker.DescribeHumanTaskUiInput) (*sagemaker.DescribeHumanTaskUiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHumanTaskUiOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHumanTaskUi(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHumanTaskUiWithContext(ctx context.Context, input *sagemaker.DescribeHumanTaskUiInput, opts ...request.Option) (*sagemaker.DescribeHumanTaskUiOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHumanTaskUiOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHumanTaskUiWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHyperParameterTuningJob(input *sagemaker.DescribeHyperParameterTuningJobInput) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHyperParameterTuningJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHyperParameterTuningJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeHyperParameterTuningJobWithContext(ctx context.Context, input *sagemaker.DescribeHyperParameterTuningJobInput, opts ...request.Option) (*sagemaker.DescribeHyperParameterTuningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeHyperParameterTuningJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeHyperParameterTuningJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeImage(input *sagemaker.DescribeImageInput) (*sagemaker.DescribeImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeImageOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeImage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeImageVersion(input *sagemaker.DescribeImageVersionInput) (*sagemaker.DescribeImageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeImageVersionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeImageVersion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeImageVersionWithContext(ctx context.Context, input *sagemaker.DescribeImageVersionInput, opts ...request.Option) (*sagemaker.DescribeImageVersionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeImageVersionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeImageVersionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeImageWithContext(ctx context.Context, input *sagemaker.DescribeImageInput, opts ...request.Option) (*sagemaker.DescribeImageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeImageOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeImageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeInferenceExperiment(input *sagemaker.DescribeInferenceExperimentInput) (*sagemaker.DescribeInferenceExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeInferenceExperimentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeInferenceExperiment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeInferenceExperimentWithContext(ctx context.Context, input *sagemaker.DescribeInferenceExperimentInput, opts ...request.Option) (*sagemaker.DescribeInferenceExperimentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeInferenceExperimentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeInferenceExperimentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeInferenceRecommendationsJob(input *sagemaker.DescribeInferenceRecommendationsJobInput) (*sagemaker.DescribeInferenceRecommendationsJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeInferenceRecommendationsJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeInferenceRecommendationsJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeInferenceRecommendationsJobWithContext(ctx context.Context, input *sagemaker.DescribeInferenceRecommendationsJobInput, opts ...request.Option) (*sagemaker.DescribeInferenceRecommendationsJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeInferenceRecommendationsJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeInferenceRecommendationsJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeLabelingJob(input *sagemaker.DescribeLabelingJobInput) (*sagemaker.DescribeLabelingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeLabelingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeLabelingJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeLabelingJobWithContext(ctx context.Context, input *sagemaker.DescribeLabelingJobInput, opts ...request.Option) (*sagemaker.DescribeLabelingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeLabelingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeLabelingJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeLineageGroup(input *sagemaker.DescribeLineageGroupInput) (*sagemaker.DescribeLineageGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeLineageGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeLineageGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeLineageGroupWithContext(ctx context.Context, input *sagemaker.DescribeLineageGroupInput, opts ...request.Option) (*sagemaker.DescribeLineageGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeLineageGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeLineageGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModel(input *sagemaker.DescribeModelInput) (*sagemaker.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelBiasJobDefinition(input *sagemaker.DescribeModelBiasJobDefinitionInput) (*sagemaker.DescribeModelBiasJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelBiasJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelBiasJobDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelBiasJobDefinitionWithContext(ctx context.Context, input *sagemaker.DescribeModelBiasJobDefinitionInput, opts ...request.Option) (*sagemaker.DescribeModelBiasJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelBiasJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelBiasJobDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelCard(input *sagemaker.DescribeModelCardInput) (*sagemaker.DescribeModelCardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelCardOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelCard(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelCardExportJob(input *sagemaker.DescribeModelCardExportJobInput) (*sagemaker.DescribeModelCardExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelCardExportJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelCardExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelCardExportJobWithContext(ctx context.Context, input *sagemaker.DescribeModelCardExportJobInput, opts ...request.Option) (*sagemaker.DescribeModelCardExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelCardExportJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelCardExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelCardWithContext(ctx context.Context, input *sagemaker.DescribeModelCardInput, opts ...request.Option) (*sagemaker.DescribeModelCardOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelCardOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelCardWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelExplainabilityJobDefinition(input *sagemaker.DescribeModelExplainabilityJobDefinitionInput) (*sagemaker.DescribeModelExplainabilityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelExplainabilityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelExplainabilityJobDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelExplainabilityJobDefinitionWithContext(ctx context.Context, input *sagemaker.DescribeModelExplainabilityJobDefinitionInput, opts ...request.Option) (*sagemaker.DescribeModelExplainabilityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelExplainabilityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelExplainabilityJobDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelPackage(input *sagemaker.DescribeModelPackageInput) (*sagemaker.DescribeModelPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelPackageOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelPackage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelPackageGroup(input *sagemaker.DescribeModelPackageGroupInput) (*sagemaker.DescribeModelPackageGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelPackageGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelPackageGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelPackageGroupWithContext(ctx context.Context, input *sagemaker.DescribeModelPackageGroupInput, opts ...request.Option) (*sagemaker.DescribeModelPackageGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelPackageGroupOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelPackageGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelPackageWithContext(ctx context.Context, input *sagemaker.DescribeModelPackageInput, opts ...request.Option) (*sagemaker.DescribeModelPackageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelPackageOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelPackageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelQualityJobDefinition(input *sagemaker.DescribeModelQualityJobDefinitionInput) (*sagemaker.DescribeModelQualityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelQualityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelQualityJobDefinition(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelQualityJobDefinitionWithContext(ctx context.Context, input *sagemaker.DescribeModelQualityJobDefinitionInput, opts ...request.Option) (*sagemaker.DescribeModelQualityJobDefinitionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelQualityJobDefinitionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelQualityJobDefinitionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeModelWithContext(ctx context.Context, input *sagemaker.DescribeModelInput, opts ...request.Option) (*sagemaker.DescribeModelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeModelOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeModelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeMonitoringSchedule(input *sagemaker.DescribeMonitoringScheduleInput) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeMonitoringScheduleOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeMonitoringSchedule(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeMonitoringScheduleWithContext(ctx context.Context, input *sagemaker.DescribeMonitoringScheduleInput, opts ...request.Option) (*sagemaker.DescribeMonitoringScheduleOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeMonitoringScheduleOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeMonitoringScheduleWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeNotebookInstance(input *sagemaker.DescribeNotebookInstanceInput) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeNotebookInstanceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeNotebookInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeNotebookInstanceLifecycleConfig(input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeNotebookInstanceLifecycleConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeNotebookInstanceLifecycleConfigWithContext(ctx context.Context, input *sagemaker.DescribeNotebookInstanceLifecycleConfigInput, opts ...request.Option) (*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeNotebookInstanceLifecycleConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeNotebookInstanceLifecycleConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeNotebookInstanceWithContext(ctx context.Context, input *sagemaker.DescribeNotebookInstanceInput, opts ...request.Option) (*sagemaker.DescribeNotebookInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeNotebookInstanceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeNotebookInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipeline(input *sagemaker.DescribePipelineInput) (*sagemaker.DescribePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipeline(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipelineDefinitionForExecution(input *sagemaker.DescribePipelineDefinitionForExecutionInput) (*sagemaker.DescribePipelineDefinitionForExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineDefinitionForExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipelineDefinitionForExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipelineDefinitionForExecutionWithContext(ctx context.Context, input *sagemaker.DescribePipelineDefinitionForExecutionInput, opts ...request.Option) (*sagemaker.DescribePipelineDefinitionForExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineDefinitionForExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipelineDefinitionForExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipelineExecution(input *sagemaker.DescribePipelineExecutionInput) (*sagemaker.DescribePipelineExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipelineExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipelineExecutionWithContext(ctx context.Context, input *sagemaker.DescribePipelineExecutionInput, opts ...request.Option) (*sagemaker.DescribePipelineExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipelineExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribePipelineWithContext(ctx context.Context, input *sagemaker.DescribePipelineInput, opts ...request.Option) (*sagemaker.DescribePipelineOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribePipelineOutput), nil
	}
	out, err := c.SageMakerAPI.DescribePipelineWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeProcessingJob(input *sagemaker.DescribeProcessingJobInput) (*sagemaker.DescribeProcessingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeProcessingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeProcessingJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeProcessingJobWithContext(ctx context.Context, input *sagemaker.DescribeProcessingJobInput, opts ...request.Option) (*sagemaker.DescribeProcessingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeProcessingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeProcessingJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeProject(input *sagemaker.DescribeProjectInput) (*sagemaker.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeProjectOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeProject(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeProjectWithContext(ctx context.Context, input *sagemaker.DescribeProjectInput, opts ...request.Option) (*sagemaker.DescribeProjectOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeProjectOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeProjectWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeSpace(input *sagemaker.DescribeSpaceInput) (*sagemaker.DescribeSpaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeSpaceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeSpace(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeSpaceWithContext(ctx context.Context, input *sagemaker.DescribeSpaceInput, opts ...request.Option) (*sagemaker.DescribeSpaceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeSpaceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeSpaceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeStudioLifecycleConfig(input *sagemaker.DescribeStudioLifecycleConfigInput) (*sagemaker.DescribeStudioLifecycleConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeStudioLifecycleConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeStudioLifecycleConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeStudioLifecycleConfigWithContext(ctx context.Context, input *sagemaker.DescribeStudioLifecycleConfigInput, opts ...request.Option) (*sagemaker.DescribeStudioLifecycleConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeStudioLifecycleConfigOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeStudioLifecycleConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeSubscribedWorkteam(input *sagemaker.DescribeSubscribedWorkteamInput) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeSubscribedWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeSubscribedWorkteam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeSubscribedWorkteamWithContext(ctx context.Context, input *sagemaker.DescribeSubscribedWorkteamInput, opts ...request.Option) (*sagemaker.DescribeSubscribedWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeSubscribedWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeSubscribedWorkteamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrainingJob(input *sagemaker.DescribeTrainingJobInput) (*sagemaker.DescribeTrainingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrainingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrainingJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrainingJobWithContext(ctx context.Context, input *sagemaker.DescribeTrainingJobInput, opts ...request.Option) (*sagemaker.DescribeTrainingJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrainingJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrainingJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTransformJob(input *sagemaker.DescribeTransformJobInput) (*sagemaker.DescribeTransformJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTransformJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTransformJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTransformJobWithContext(ctx context.Context, input *sagemaker.DescribeTransformJobInput, opts ...request.Option) (*sagemaker.DescribeTransformJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTransformJobOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTransformJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrial(input *sagemaker.DescribeTrialInput) (*sagemaker.DescribeTrialOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrialOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrial(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrialComponent(input *sagemaker.DescribeTrialComponentInput) (*sagemaker.DescribeTrialComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrialComponentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrialComponent(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrialComponentWithContext(ctx context.Context, input *sagemaker.DescribeTrialComponentInput, opts ...request.Option) (*sagemaker.DescribeTrialComponentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrialComponentOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrialComponentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeTrialWithContext(ctx context.Context, input *sagemaker.DescribeTrialInput, opts ...request.Option) (*sagemaker.DescribeTrialOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeTrialOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeTrialWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeUserProfile(input *sagemaker.DescribeUserProfileInput) (*sagemaker.DescribeUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeUserProfileOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeUserProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeUserProfileWithContext(ctx context.Context, input *sagemaker.DescribeUserProfileInput, opts ...request.Option) (*sagemaker.DescribeUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeUserProfileOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeUserProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeWorkforce(input *sagemaker.DescribeWorkforceInput) (*sagemaker.DescribeWorkforceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeWorkforceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeWorkforce(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeWorkforceWithContext(ctx context.Context, input *sagemaker.DescribeWorkforceInput, opts ...request.Option) (*sagemaker.DescribeWorkforceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeWorkforceOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeWorkforceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeWorkteam(input *sagemaker.DescribeWorkteamInput) (*sagemaker.DescribeWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeWorkteam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) DescribeWorkteamWithContext(ctx context.Context, input *sagemaker.DescribeWorkteamInput, opts ...request.Option) (*sagemaker.DescribeWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.DescribeWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.DescribeWorkteamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetDeviceFleetReport(input *sagemaker.GetDeviceFleetReportInput) (*sagemaker.GetDeviceFleetReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetDeviceFleetReportOutput), nil
	}
	out, err := c.SageMakerAPI.GetDeviceFleetReport(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetDeviceFleetReportWithContext(ctx context.Context, input *sagemaker.GetDeviceFleetReportInput, opts ...request.Option) (*sagemaker.GetDeviceFleetReportOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetDeviceFleetReportOutput), nil
	}
	out, err := c.SageMakerAPI.GetDeviceFleetReportWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetLineageGroupPolicy(input *sagemaker.GetLineageGroupPolicyInput) (*sagemaker.GetLineageGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetLineageGroupPolicyOutput), nil
	}
	out, err := c.SageMakerAPI.GetLineageGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetLineageGroupPolicyWithContext(ctx context.Context, input *sagemaker.GetLineageGroupPolicyInput, opts ...request.Option) (*sagemaker.GetLineageGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetLineageGroupPolicyOutput), nil
	}
	out, err := c.SageMakerAPI.GetLineageGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetModelPackageGroupPolicy(input *sagemaker.GetModelPackageGroupPolicyInput) (*sagemaker.GetModelPackageGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetModelPackageGroupPolicyOutput), nil
	}
	out, err := c.SageMakerAPI.GetModelPackageGroupPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetModelPackageGroupPolicyWithContext(ctx context.Context, input *sagemaker.GetModelPackageGroupPolicyInput, opts ...request.Option) (*sagemaker.GetModelPackageGroupPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetModelPackageGroupPolicyOutput), nil
	}
	out, err := c.SageMakerAPI.GetModelPackageGroupPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetSagemakerServicecatalogPortfolioStatus(input *sagemaker.GetSagemakerServicecatalogPortfolioStatusInput) (*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput), nil
	}
	out, err := c.SageMakerAPI.GetSagemakerServicecatalogPortfolioStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetSagemakerServicecatalogPortfolioStatusWithContext(ctx context.Context, input *sagemaker.GetSagemakerServicecatalogPortfolioStatusInput, opts ...request.Option) (*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetSagemakerServicecatalogPortfolioStatusOutput), nil
	}
	out, err := c.SageMakerAPI.GetSagemakerServicecatalogPortfolioStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetSearchSuggestions(input *sagemaker.GetSearchSuggestionsInput) (*sagemaker.GetSearchSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetSearchSuggestionsOutput), nil
	}
	out, err := c.SageMakerAPI.GetSearchSuggestions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) GetSearchSuggestionsWithContext(ctx context.Context, input *sagemaker.GetSearchSuggestionsInput, opts ...request.Option) (*sagemaker.GetSearchSuggestionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.GetSearchSuggestionsOutput), nil
	}
	out, err := c.SageMakerAPI.GetSearchSuggestionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListActions(input *sagemaker.ListActionsInput) (*sagemaker.ListActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListActionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListActionsPages(input *sagemaker.ListActionsInput, fn func(*sagemaker.ListActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListActionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListActionsOutput{}
	fnCacher := func(out *sagemaker.ListActionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListActionsPagesWithContext(ctx context.Context, input *sagemaker.ListActionsInput, fn func(*sagemaker.ListActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListActionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListActionsOutput{}
	fnCacher := func(out *sagemaker.ListActionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListActionsWithContext(ctx context.Context, input *sagemaker.ListActionsInput, opts ...request.Option) (*sagemaker.ListActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListActionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAlgorithms(input *sagemaker.ListAlgorithmsInput) (*sagemaker.ListAlgorithmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAlgorithmsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAlgorithms(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAlgorithmsPages(input *sagemaker.ListAlgorithmsInput, fn func(*sagemaker.ListAlgorithmsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAlgorithmsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAlgorithmsOutput{}
	fnCacher := func(out *sagemaker.ListAlgorithmsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAlgorithmsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAlgorithmsPagesWithContext(ctx context.Context, input *sagemaker.ListAlgorithmsInput, fn func(*sagemaker.ListAlgorithmsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAlgorithmsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAlgorithmsOutput{}
	fnCacher := func(out *sagemaker.ListAlgorithmsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAlgorithmsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAlgorithmsWithContext(ctx context.Context, input *sagemaker.ListAlgorithmsInput, opts ...request.Option) (*sagemaker.ListAlgorithmsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAlgorithmsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAlgorithmsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAliases(input *sagemaker.ListAliasesInput) (*sagemaker.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAliasesOutput), nil
	}
	out, err := c.SageMakerAPI.ListAliases(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAliasesPages(input *sagemaker.ListAliasesInput, fn func(*sagemaker.ListAliasesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAliasesOutput{}
	fnCacher := func(out *sagemaker.ListAliasesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAliasesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAliasesPagesWithContext(ctx context.Context, input *sagemaker.ListAliasesInput, fn func(*sagemaker.ListAliasesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAliasesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAliasesOutput{}
	fnCacher := func(out *sagemaker.ListAliasesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAliasesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAliasesWithContext(ctx context.Context, input *sagemaker.ListAliasesInput, opts ...request.Option) (*sagemaker.ListAliasesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAliasesOutput), nil
	}
	out, err := c.SageMakerAPI.ListAliasesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAppImageConfigs(input *sagemaker.ListAppImageConfigsInput) (*sagemaker.ListAppImageConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAppImageConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAppImageConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAppImageConfigsPages(input *sagemaker.ListAppImageConfigsInput, fn func(*sagemaker.ListAppImageConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAppImageConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAppImageConfigsOutput{}
	fnCacher := func(out *sagemaker.ListAppImageConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAppImageConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAppImageConfigsPagesWithContext(ctx context.Context, input *sagemaker.ListAppImageConfigsInput, fn func(*sagemaker.ListAppImageConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAppImageConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAppImageConfigsOutput{}
	fnCacher := func(out *sagemaker.ListAppImageConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAppImageConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAppImageConfigsWithContext(ctx context.Context, input *sagemaker.ListAppImageConfigsInput, opts ...request.Option) (*sagemaker.ListAppImageConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAppImageConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAppImageConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListApps(input *sagemaker.ListAppsInput) (*sagemaker.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAppsOutput), nil
	}
	out, err := c.SageMakerAPI.ListApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAppsPages(input *sagemaker.ListAppsInput, fn func(*sagemaker.ListAppsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAppsOutput{}
	fnCacher := func(out *sagemaker.ListAppsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAppsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAppsPagesWithContext(ctx context.Context, input *sagemaker.ListAppsInput, fn func(*sagemaker.ListAppsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAppsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAppsOutput{}
	fnCacher := func(out *sagemaker.ListAppsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAppsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAppsWithContext(ctx context.Context, input *sagemaker.ListAppsInput, opts ...request.Option) (*sagemaker.ListAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAppsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListArtifacts(input *sagemaker.ListArtifactsInput) (*sagemaker.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListArtifactsOutput), nil
	}
	out, err := c.SageMakerAPI.ListArtifacts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListArtifactsPages(input *sagemaker.ListArtifactsInput, fn func(*sagemaker.ListArtifactsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListArtifactsOutput{}
	fnCacher := func(out *sagemaker.ListArtifactsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListArtifactsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListArtifactsPagesWithContext(ctx context.Context, input *sagemaker.ListArtifactsInput, fn func(*sagemaker.ListArtifactsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListArtifactsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListArtifactsOutput{}
	fnCacher := func(out *sagemaker.ListArtifactsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListArtifactsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListArtifactsWithContext(ctx context.Context, input *sagemaker.ListArtifactsInput, opts ...request.Option) (*sagemaker.ListArtifactsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListArtifactsOutput), nil
	}
	out, err := c.SageMakerAPI.ListArtifactsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAssociations(input *sagemaker.ListAssociationsInput) (*sagemaker.ListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAssociationsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAssociations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAssociationsPages(input *sagemaker.ListAssociationsInput, fn func(*sagemaker.ListAssociationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAssociationsOutput{}
	fnCacher := func(out *sagemaker.ListAssociationsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAssociationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAssociationsPagesWithContext(ctx context.Context, input *sagemaker.ListAssociationsInput, fn func(*sagemaker.ListAssociationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAssociationsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAssociationsOutput{}
	fnCacher := func(out *sagemaker.ListAssociationsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAssociationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAssociationsWithContext(ctx context.Context, input *sagemaker.ListAssociationsInput, opts ...request.Option) (*sagemaker.ListAssociationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAssociationsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAssociationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAutoMLJobs(input *sagemaker.ListAutoMLJobsInput) (*sagemaker.ListAutoMLJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAutoMLJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAutoMLJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListAutoMLJobsPages(input *sagemaker.ListAutoMLJobsInput, fn func(*sagemaker.ListAutoMLJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAutoMLJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAutoMLJobsOutput{}
	fnCacher := func(out *sagemaker.ListAutoMLJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAutoMLJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAutoMLJobsPagesWithContext(ctx context.Context, input *sagemaker.ListAutoMLJobsInput, fn func(*sagemaker.ListAutoMLJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListAutoMLJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListAutoMLJobsOutput{}
	fnCacher := func(out *sagemaker.ListAutoMLJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListAutoMLJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListAutoMLJobsWithContext(ctx context.Context, input *sagemaker.ListAutoMLJobsInput, opts ...request.Option) (*sagemaker.ListAutoMLJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListAutoMLJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListAutoMLJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCandidatesForAutoMLJob(input *sagemaker.ListCandidatesForAutoMLJobInput) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCandidatesForAutoMLJobOutput), nil
	}
	out, err := c.SageMakerAPI.ListCandidatesForAutoMLJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCandidatesForAutoMLJobPages(input *sagemaker.ListCandidatesForAutoMLJobInput, fn func(*sagemaker.ListCandidatesForAutoMLJobOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCandidatesForAutoMLJobOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCandidatesForAutoMLJobOutput{}
	fnCacher := func(out *sagemaker.ListCandidatesForAutoMLJobOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCandidatesForAutoMLJobPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCandidatesForAutoMLJobPagesWithContext(ctx context.Context, input *sagemaker.ListCandidatesForAutoMLJobInput, fn func(*sagemaker.ListCandidatesForAutoMLJobOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCandidatesForAutoMLJobOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCandidatesForAutoMLJobOutput{}
	fnCacher := func(out *sagemaker.ListCandidatesForAutoMLJobOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCandidatesForAutoMLJobPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCandidatesForAutoMLJobWithContext(ctx context.Context, input *sagemaker.ListCandidatesForAutoMLJobInput, opts ...request.Option) (*sagemaker.ListCandidatesForAutoMLJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCandidatesForAutoMLJobOutput), nil
	}
	out, err := c.SageMakerAPI.ListCandidatesForAutoMLJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCodeRepositories(input *sagemaker.ListCodeRepositoriesInput) (*sagemaker.ListCodeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCodeRepositoriesOutput), nil
	}
	out, err := c.SageMakerAPI.ListCodeRepositories(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCodeRepositoriesPages(input *sagemaker.ListCodeRepositoriesInput, fn func(*sagemaker.ListCodeRepositoriesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCodeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCodeRepositoriesOutput{}
	fnCacher := func(out *sagemaker.ListCodeRepositoriesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCodeRepositoriesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCodeRepositoriesPagesWithContext(ctx context.Context, input *sagemaker.ListCodeRepositoriesInput, fn func(*sagemaker.ListCodeRepositoriesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCodeRepositoriesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCodeRepositoriesOutput{}
	fnCacher := func(out *sagemaker.ListCodeRepositoriesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCodeRepositoriesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCodeRepositoriesWithContext(ctx context.Context, input *sagemaker.ListCodeRepositoriesInput, opts ...request.Option) (*sagemaker.ListCodeRepositoriesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCodeRepositoriesOutput), nil
	}
	out, err := c.SageMakerAPI.ListCodeRepositoriesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCompilationJobs(input *sagemaker.ListCompilationJobsInput) (*sagemaker.ListCompilationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCompilationJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListCompilationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListCompilationJobsPages(input *sagemaker.ListCompilationJobsInput, fn func(*sagemaker.ListCompilationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCompilationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCompilationJobsOutput{}
	fnCacher := func(out *sagemaker.ListCompilationJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCompilationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCompilationJobsPagesWithContext(ctx context.Context, input *sagemaker.ListCompilationJobsInput, fn func(*sagemaker.ListCompilationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListCompilationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListCompilationJobsOutput{}
	fnCacher := func(out *sagemaker.ListCompilationJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListCompilationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListCompilationJobsWithContext(ctx context.Context, input *sagemaker.ListCompilationJobsInput, opts ...request.Option) (*sagemaker.ListCompilationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListCompilationJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListCompilationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListContexts(input *sagemaker.ListContextsInput) (*sagemaker.ListContextsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListContextsOutput), nil
	}
	out, err := c.SageMakerAPI.ListContexts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListContextsPages(input *sagemaker.ListContextsInput, fn func(*sagemaker.ListContextsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListContextsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListContextsOutput{}
	fnCacher := func(out *sagemaker.ListContextsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListContextsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListContextsPagesWithContext(ctx context.Context, input *sagemaker.ListContextsInput, fn func(*sagemaker.ListContextsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListContextsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListContextsOutput{}
	fnCacher := func(out *sagemaker.ListContextsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListContextsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListContextsWithContext(ctx context.Context, input *sagemaker.ListContextsInput, opts ...request.Option) (*sagemaker.ListContextsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListContextsOutput), nil
	}
	out, err := c.SageMakerAPI.ListContextsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDataQualityJobDefinitions(input *sagemaker.ListDataQualityJobDefinitionsInput) (*sagemaker.ListDataQualityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDataQualityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDataQualityJobDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDataQualityJobDefinitionsPages(input *sagemaker.ListDataQualityJobDefinitionsInput, fn func(*sagemaker.ListDataQualityJobDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDataQualityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDataQualityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListDataQualityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDataQualityJobDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDataQualityJobDefinitionsPagesWithContext(ctx context.Context, input *sagemaker.ListDataQualityJobDefinitionsInput, fn func(*sagemaker.ListDataQualityJobDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDataQualityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDataQualityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListDataQualityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDataQualityJobDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDataQualityJobDefinitionsWithContext(ctx context.Context, input *sagemaker.ListDataQualityJobDefinitionsInput, opts ...request.Option) (*sagemaker.ListDataQualityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDataQualityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDataQualityJobDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDeviceFleets(input *sagemaker.ListDeviceFleetsInput) (*sagemaker.ListDeviceFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDeviceFleetsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDeviceFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDeviceFleetsPages(input *sagemaker.ListDeviceFleetsInput, fn func(*sagemaker.ListDeviceFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDeviceFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDeviceFleetsOutput{}
	fnCacher := func(out *sagemaker.ListDeviceFleetsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDeviceFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDeviceFleetsPagesWithContext(ctx context.Context, input *sagemaker.ListDeviceFleetsInput, fn func(*sagemaker.ListDeviceFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDeviceFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDeviceFleetsOutput{}
	fnCacher := func(out *sagemaker.ListDeviceFleetsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDeviceFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDeviceFleetsWithContext(ctx context.Context, input *sagemaker.ListDeviceFleetsInput, opts ...request.Option) (*sagemaker.ListDeviceFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDeviceFleetsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDeviceFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDevices(input *sagemaker.ListDevicesInput) (*sagemaker.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDevicesOutput), nil
	}
	out, err := c.SageMakerAPI.ListDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDevicesPages(input *sagemaker.ListDevicesInput, fn func(*sagemaker.ListDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDevicesOutput{}
	fnCacher := func(out *sagemaker.ListDevicesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDevicesPagesWithContext(ctx context.Context, input *sagemaker.ListDevicesInput, fn func(*sagemaker.ListDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDevicesOutput{}
	fnCacher := func(out *sagemaker.ListDevicesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDevicesWithContext(ctx context.Context, input *sagemaker.ListDevicesInput, opts ...request.Option) (*sagemaker.ListDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDevicesOutput), nil
	}
	out, err := c.SageMakerAPI.ListDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDomains(input *sagemaker.ListDomainsInput) (*sagemaker.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDomainsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDomains(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListDomainsPages(input *sagemaker.ListDomainsInput, fn func(*sagemaker.ListDomainsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDomainsOutput{}
	fnCacher := func(out *sagemaker.ListDomainsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDomainsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDomainsPagesWithContext(ctx context.Context, input *sagemaker.ListDomainsInput, fn func(*sagemaker.ListDomainsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListDomainsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListDomainsOutput{}
	fnCacher := func(out *sagemaker.ListDomainsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListDomainsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListDomainsWithContext(ctx context.Context, input *sagemaker.ListDomainsInput, opts ...request.Option) (*sagemaker.ListDomainsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListDomainsOutput), nil
	}
	out, err := c.SageMakerAPI.ListDomainsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEdgeDeploymentPlans(input *sagemaker.ListEdgeDeploymentPlansInput) (*sagemaker.ListEdgeDeploymentPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEdgeDeploymentPlansOutput), nil
	}
	out, err := c.SageMakerAPI.ListEdgeDeploymentPlans(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEdgeDeploymentPlansPages(input *sagemaker.ListEdgeDeploymentPlansInput, fn func(*sagemaker.ListEdgeDeploymentPlansOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEdgeDeploymentPlansOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEdgeDeploymentPlansOutput{}
	fnCacher := func(out *sagemaker.ListEdgeDeploymentPlansOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEdgeDeploymentPlansPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEdgeDeploymentPlansPagesWithContext(ctx context.Context, input *sagemaker.ListEdgeDeploymentPlansInput, fn func(*sagemaker.ListEdgeDeploymentPlansOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEdgeDeploymentPlansOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEdgeDeploymentPlansOutput{}
	fnCacher := func(out *sagemaker.ListEdgeDeploymentPlansOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEdgeDeploymentPlansPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEdgeDeploymentPlansWithContext(ctx context.Context, input *sagemaker.ListEdgeDeploymentPlansInput, opts ...request.Option) (*sagemaker.ListEdgeDeploymentPlansOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEdgeDeploymentPlansOutput), nil
	}
	out, err := c.SageMakerAPI.ListEdgeDeploymentPlansWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEdgePackagingJobs(input *sagemaker.ListEdgePackagingJobsInput) (*sagemaker.ListEdgePackagingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEdgePackagingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEdgePackagingJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEdgePackagingJobsPages(input *sagemaker.ListEdgePackagingJobsInput, fn func(*sagemaker.ListEdgePackagingJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEdgePackagingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEdgePackagingJobsOutput{}
	fnCacher := func(out *sagemaker.ListEdgePackagingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEdgePackagingJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEdgePackagingJobsPagesWithContext(ctx context.Context, input *sagemaker.ListEdgePackagingJobsInput, fn func(*sagemaker.ListEdgePackagingJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEdgePackagingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEdgePackagingJobsOutput{}
	fnCacher := func(out *sagemaker.ListEdgePackagingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEdgePackagingJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEdgePackagingJobsWithContext(ctx context.Context, input *sagemaker.ListEdgePackagingJobsInput, opts ...request.Option) (*sagemaker.ListEdgePackagingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEdgePackagingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEdgePackagingJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEndpointConfigs(input *sagemaker.ListEndpointConfigsInput) (*sagemaker.ListEndpointConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEndpointConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEndpointConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEndpointConfigsPages(input *sagemaker.ListEndpointConfigsInput, fn func(*sagemaker.ListEndpointConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEndpointConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEndpointConfigsOutput{}
	fnCacher := func(out *sagemaker.ListEndpointConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEndpointConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEndpointConfigsPagesWithContext(ctx context.Context, input *sagemaker.ListEndpointConfigsInput, fn func(*sagemaker.ListEndpointConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEndpointConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEndpointConfigsOutput{}
	fnCacher := func(out *sagemaker.ListEndpointConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEndpointConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEndpointConfigsWithContext(ctx context.Context, input *sagemaker.ListEndpointConfigsInput, opts ...request.Option) (*sagemaker.ListEndpointConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEndpointConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEndpointConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEndpoints(input *sagemaker.ListEndpointsInput) (*sagemaker.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEndpointsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEndpoints(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListEndpointsPages(input *sagemaker.ListEndpointsInput, fn func(*sagemaker.ListEndpointsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEndpointsOutput{}
	fnCacher := func(out *sagemaker.ListEndpointsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEndpointsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEndpointsPagesWithContext(ctx context.Context, input *sagemaker.ListEndpointsInput, fn func(*sagemaker.ListEndpointsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListEndpointsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListEndpointsOutput{}
	fnCacher := func(out *sagemaker.ListEndpointsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListEndpointsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListEndpointsWithContext(ctx context.Context, input *sagemaker.ListEndpointsInput, opts ...request.Option) (*sagemaker.ListEndpointsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListEndpointsOutput), nil
	}
	out, err := c.SageMakerAPI.ListEndpointsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListExperiments(input *sagemaker.ListExperimentsInput) (*sagemaker.ListExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListExperimentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListExperiments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListExperimentsPages(input *sagemaker.ListExperimentsInput, fn func(*sagemaker.ListExperimentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListExperimentsOutput{}
	fnCacher := func(out *sagemaker.ListExperimentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListExperimentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListExperimentsPagesWithContext(ctx context.Context, input *sagemaker.ListExperimentsInput, fn func(*sagemaker.ListExperimentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListExperimentsOutput{}
	fnCacher := func(out *sagemaker.ListExperimentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListExperimentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListExperimentsWithContext(ctx context.Context, input *sagemaker.ListExperimentsInput, opts ...request.Option) (*sagemaker.ListExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListExperimentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListExperimentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListFeatureGroups(input *sagemaker.ListFeatureGroupsInput) (*sagemaker.ListFeatureGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListFeatureGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListFeatureGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListFeatureGroupsPages(input *sagemaker.ListFeatureGroupsInput, fn func(*sagemaker.ListFeatureGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListFeatureGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListFeatureGroupsOutput{}
	fnCacher := func(out *sagemaker.ListFeatureGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListFeatureGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListFeatureGroupsPagesWithContext(ctx context.Context, input *sagemaker.ListFeatureGroupsInput, fn func(*sagemaker.ListFeatureGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListFeatureGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListFeatureGroupsOutput{}
	fnCacher := func(out *sagemaker.ListFeatureGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListFeatureGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListFeatureGroupsWithContext(ctx context.Context, input *sagemaker.ListFeatureGroupsInput, opts ...request.Option) (*sagemaker.ListFeatureGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListFeatureGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListFeatureGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListFlowDefinitions(input *sagemaker.ListFlowDefinitionsInput) (*sagemaker.ListFlowDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListFlowDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListFlowDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListFlowDefinitionsPages(input *sagemaker.ListFlowDefinitionsInput, fn func(*sagemaker.ListFlowDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListFlowDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListFlowDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListFlowDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListFlowDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListFlowDefinitionsPagesWithContext(ctx context.Context, input *sagemaker.ListFlowDefinitionsInput, fn func(*sagemaker.ListFlowDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListFlowDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListFlowDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListFlowDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListFlowDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListFlowDefinitionsWithContext(ctx context.Context, input *sagemaker.ListFlowDefinitionsInput, opts ...request.Option) (*sagemaker.ListFlowDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListFlowDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListFlowDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubContentVersions(input *sagemaker.ListHubContentVersionsInput) (*sagemaker.ListHubContentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubContentVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubContentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubContentVersionsWithContext(ctx context.Context, input *sagemaker.ListHubContentVersionsInput, opts ...request.Option) (*sagemaker.ListHubContentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubContentVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubContentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubContents(input *sagemaker.ListHubContentsInput) (*sagemaker.ListHubContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubContentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubContents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubContentsWithContext(ctx context.Context, input *sagemaker.ListHubContentsInput, opts ...request.Option) (*sagemaker.ListHubContentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubContentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubContentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubs(input *sagemaker.ListHubsInput) (*sagemaker.ListHubsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHubsWithContext(ctx context.Context, input *sagemaker.ListHubsInput, opts ...request.Option) (*sagemaker.ListHubsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHubsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHubsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHumanTaskUis(input *sagemaker.ListHumanTaskUisInput) (*sagemaker.ListHumanTaskUisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHumanTaskUisOutput), nil
	}
	out, err := c.SageMakerAPI.ListHumanTaskUis(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHumanTaskUisPages(input *sagemaker.ListHumanTaskUisInput, fn func(*sagemaker.ListHumanTaskUisOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListHumanTaskUisOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListHumanTaskUisOutput{}
	fnCacher := func(out *sagemaker.ListHumanTaskUisOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListHumanTaskUisPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListHumanTaskUisPagesWithContext(ctx context.Context, input *sagemaker.ListHumanTaskUisInput, fn func(*sagemaker.ListHumanTaskUisOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListHumanTaskUisOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListHumanTaskUisOutput{}
	fnCacher := func(out *sagemaker.ListHumanTaskUisOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListHumanTaskUisPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListHumanTaskUisWithContext(ctx context.Context, input *sagemaker.ListHumanTaskUisInput, opts ...request.Option) (*sagemaker.ListHumanTaskUisOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHumanTaskUisOutput), nil
	}
	out, err := c.SageMakerAPI.ListHumanTaskUisWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHyperParameterTuningJobs(input *sagemaker.ListHyperParameterTuningJobsInput) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHyperParameterTuningJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHyperParameterTuningJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListHyperParameterTuningJobsPages(input *sagemaker.ListHyperParameterTuningJobsInput, fn func(*sagemaker.ListHyperParameterTuningJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListHyperParameterTuningJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListHyperParameterTuningJobsOutput{}
	fnCacher := func(out *sagemaker.ListHyperParameterTuningJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListHyperParameterTuningJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListHyperParameterTuningJobsPagesWithContext(ctx context.Context, input *sagemaker.ListHyperParameterTuningJobsInput, fn func(*sagemaker.ListHyperParameterTuningJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListHyperParameterTuningJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListHyperParameterTuningJobsOutput{}
	fnCacher := func(out *sagemaker.ListHyperParameterTuningJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListHyperParameterTuningJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListHyperParameterTuningJobsWithContext(ctx context.Context, input *sagemaker.ListHyperParameterTuningJobsInput, opts ...request.Option) (*sagemaker.ListHyperParameterTuningJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListHyperParameterTuningJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListHyperParameterTuningJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListImageVersions(input *sagemaker.ListImageVersionsInput) (*sagemaker.ListImageVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListImageVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListImageVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListImageVersionsPages(input *sagemaker.ListImageVersionsInput, fn func(*sagemaker.ListImageVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListImageVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListImageVersionsOutput{}
	fnCacher := func(out *sagemaker.ListImageVersionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListImageVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListImageVersionsPagesWithContext(ctx context.Context, input *sagemaker.ListImageVersionsInput, fn func(*sagemaker.ListImageVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListImageVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListImageVersionsOutput{}
	fnCacher := func(out *sagemaker.ListImageVersionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListImageVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListImageVersionsWithContext(ctx context.Context, input *sagemaker.ListImageVersionsInput, opts ...request.Option) (*sagemaker.ListImageVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListImageVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListImageVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListImages(input *sagemaker.ListImagesInput) (*sagemaker.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListImagesOutput), nil
	}
	out, err := c.SageMakerAPI.ListImages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListImagesPages(input *sagemaker.ListImagesInput, fn func(*sagemaker.ListImagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListImagesOutput{}
	fnCacher := func(out *sagemaker.ListImagesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListImagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListImagesPagesWithContext(ctx context.Context, input *sagemaker.ListImagesInput, fn func(*sagemaker.ListImagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListImagesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListImagesOutput{}
	fnCacher := func(out *sagemaker.ListImagesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListImagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListImagesWithContext(ctx context.Context, input *sagemaker.ListImagesInput, opts ...request.Option) (*sagemaker.ListImagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListImagesOutput), nil
	}
	out, err := c.SageMakerAPI.ListImagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceExperiments(input *sagemaker.ListInferenceExperimentsInput) (*sagemaker.ListInferenceExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceExperimentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceExperiments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceExperimentsPages(input *sagemaker.ListInferenceExperimentsInput, fn func(*sagemaker.ListInferenceExperimentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceExperimentsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceExperimentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceExperimentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceExperimentsPagesWithContext(ctx context.Context, input *sagemaker.ListInferenceExperimentsInput, fn func(*sagemaker.ListInferenceExperimentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceExperimentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceExperimentsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceExperimentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceExperimentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceExperimentsWithContext(ctx context.Context, input *sagemaker.ListInferenceExperimentsInput, opts ...request.Option) (*sagemaker.ListInferenceExperimentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceExperimentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceExperimentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceRecommendationsJobSteps(input *sagemaker.ListInferenceRecommendationsJobStepsInput) (*sagemaker.ListInferenceRecommendationsJobStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceRecommendationsJobStepsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceRecommendationsJobSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceRecommendationsJobStepsPages(input *sagemaker.ListInferenceRecommendationsJobStepsInput, fn func(*sagemaker.ListInferenceRecommendationsJobStepsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceRecommendationsJobStepsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceRecommendationsJobStepsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceRecommendationsJobStepsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceRecommendationsJobStepsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceRecommendationsJobStepsPagesWithContext(ctx context.Context, input *sagemaker.ListInferenceRecommendationsJobStepsInput, fn func(*sagemaker.ListInferenceRecommendationsJobStepsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceRecommendationsJobStepsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceRecommendationsJobStepsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceRecommendationsJobStepsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceRecommendationsJobStepsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceRecommendationsJobStepsWithContext(ctx context.Context, input *sagemaker.ListInferenceRecommendationsJobStepsInput, opts ...request.Option) (*sagemaker.ListInferenceRecommendationsJobStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceRecommendationsJobStepsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceRecommendationsJobStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceRecommendationsJobs(input *sagemaker.ListInferenceRecommendationsJobsInput) (*sagemaker.ListInferenceRecommendationsJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceRecommendationsJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceRecommendationsJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListInferenceRecommendationsJobsPages(input *sagemaker.ListInferenceRecommendationsJobsInput, fn func(*sagemaker.ListInferenceRecommendationsJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceRecommendationsJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceRecommendationsJobsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceRecommendationsJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceRecommendationsJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceRecommendationsJobsPagesWithContext(ctx context.Context, input *sagemaker.ListInferenceRecommendationsJobsInput, fn func(*sagemaker.ListInferenceRecommendationsJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListInferenceRecommendationsJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListInferenceRecommendationsJobsOutput{}
	fnCacher := func(out *sagemaker.ListInferenceRecommendationsJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListInferenceRecommendationsJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListInferenceRecommendationsJobsWithContext(ctx context.Context, input *sagemaker.ListInferenceRecommendationsJobsInput, opts ...request.Option) (*sagemaker.ListInferenceRecommendationsJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListInferenceRecommendationsJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListInferenceRecommendationsJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLabelingJobs(input *sagemaker.ListLabelingJobsInput) (*sagemaker.ListLabelingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLabelingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListLabelingJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLabelingJobsForWorkteam(input *sagemaker.ListLabelingJobsForWorkteamInput) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLabelingJobsForWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.ListLabelingJobsForWorkteam(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLabelingJobsForWorkteamPages(input *sagemaker.ListLabelingJobsForWorkteamInput, fn func(*sagemaker.ListLabelingJobsForWorkteamOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLabelingJobsForWorkteamOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLabelingJobsForWorkteamOutput{}
	fnCacher := func(out *sagemaker.ListLabelingJobsForWorkteamOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLabelingJobsForWorkteamPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLabelingJobsForWorkteamPagesWithContext(ctx context.Context, input *sagemaker.ListLabelingJobsForWorkteamInput, fn func(*sagemaker.ListLabelingJobsForWorkteamOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLabelingJobsForWorkteamOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLabelingJobsForWorkteamOutput{}
	fnCacher := func(out *sagemaker.ListLabelingJobsForWorkteamOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLabelingJobsForWorkteamPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLabelingJobsForWorkteamWithContext(ctx context.Context, input *sagemaker.ListLabelingJobsForWorkteamInput, opts ...request.Option) (*sagemaker.ListLabelingJobsForWorkteamOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLabelingJobsForWorkteamOutput), nil
	}
	out, err := c.SageMakerAPI.ListLabelingJobsForWorkteamWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLabelingJobsPages(input *sagemaker.ListLabelingJobsInput, fn func(*sagemaker.ListLabelingJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLabelingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLabelingJobsOutput{}
	fnCacher := func(out *sagemaker.ListLabelingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLabelingJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLabelingJobsPagesWithContext(ctx context.Context, input *sagemaker.ListLabelingJobsInput, fn func(*sagemaker.ListLabelingJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLabelingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLabelingJobsOutput{}
	fnCacher := func(out *sagemaker.ListLabelingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLabelingJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLabelingJobsWithContext(ctx context.Context, input *sagemaker.ListLabelingJobsInput, opts ...request.Option) (*sagemaker.ListLabelingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLabelingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListLabelingJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLineageGroups(input *sagemaker.ListLineageGroupsInput) (*sagemaker.ListLineageGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLineageGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListLineageGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListLineageGroupsPages(input *sagemaker.ListLineageGroupsInput, fn func(*sagemaker.ListLineageGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLineageGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLineageGroupsOutput{}
	fnCacher := func(out *sagemaker.ListLineageGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLineageGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLineageGroupsPagesWithContext(ctx context.Context, input *sagemaker.ListLineageGroupsInput, fn func(*sagemaker.ListLineageGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListLineageGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListLineageGroupsOutput{}
	fnCacher := func(out *sagemaker.ListLineageGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListLineageGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListLineageGroupsWithContext(ctx context.Context, input *sagemaker.ListLineageGroupsInput, opts ...request.Option) (*sagemaker.ListLineageGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListLineageGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListLineageGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelBiasJobDefinitions(input *sagemaker.ListModelBiasJobDefinitionsInput) (*sagemaker.ListModelBiasJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelBiasJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelBiasJobDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelBiasJobDefinitionsPages(input *sagemaker.ListModelBiasJobDefinitionsInput, fn func(*sagemaker.ListModelBiasJobDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelBiasJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelBiasJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelBiasJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelBiasJobDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelBiasJobDefinitionsPagesWithContext(ctx context.Context, input *sagemaker.ListModelBiasJobDefinitionsInput, fn func(*sagemaker.ListModelBiasJobDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelBiasJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelBiasJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelBiasJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelBiasJobDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelBiasJobDefinitionsWithContext(ctx context.Context, input *sagemaker.ListModelBiasJobDefinitionsInput, opts ...request.Option) (*sagemaker.ListModelBiasJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelBiasJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelBiasJobDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCardExportJobs(input *sagemaker.ListModelCardExportJobsInput) (*sagemaker.ListModelCardExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardExportJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCardExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCardExportJobsPages(input *sagemaker.ListModelCardExportJobsInput, fn func(*sagemaker.ListModelCardExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardExportJobsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardExportJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardExportJobsPagesWithContext(ctx context.Context, input *sagemaker.ListModelCardExportJobsInput, fn func(*sagemaker.ListModelCardExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardExportJobsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardExportJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardExportJobsWithContext(ctx context.Context, input *sagemaker.ListModelCardExportJobsInput, opts ...request.Option) (*sagemaker.ListModelCardExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardExportJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCardExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCardVersions(input *sagemaker.ListModelCardVersionsInput) (*sagemaker.ListModelCardVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCardVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCardVersionsPages(input *sagemaker.ListModelCardVersionsInput, fn func(*sagemaker.ListModelCardVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardVersionsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardVersionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardVersionsPagesWithContext(ctx context.Context, input *sagemaker.ListModelCardVersionsInput, fn func(*sagemaker.ListModelCardVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardVersionsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardVersionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardVersionsWithContext(ctx context.Context, input *sagemaker.ListModelCardVersionsInput, opts ...request.Option) (*sagemaker.ListModelCardVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardVersionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCardVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCards(input *sagemaker.ListModelCardsInput) (*sagemaker.ListModelCardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCards(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelCardsPages(input *sagemaker.ListModelCardsInput, fn func(*sagemaker.ListModelCardsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardsPagesWithContext(ctx context.Context, input *sagemaker.ListModelCardsInput, fn func(*sagemaker.ListModelCardsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelCardsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelCardsOutput{}
	fnCacher := func(out *sagemaker.ListModelCardsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelCardsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelCardsWithContext(ctx context.Context, input *sagemaker.ListModelCardsInput, opts ...request.Option) (*sagemaker.ListModelCardsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelCardsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelCardsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelExplainabilityJobDefinitions(input *sagemaker.ListModelExplainabilityJobDefinitionsInput) (*sagemaker.ListModelExplainabilityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelExplainabilityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelExplainabilityJobDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelExplainabilityJobDefinitionsPages(input *sagemaker.ListModelExplainabilityJobDefinitionsInput, fn func(*sagemaker.ListModelExplainabilityJobDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelExplainabilityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelExplainabilityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelExplainabilityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelExplainabilityJobDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelExplainabilityJobDefinitionsPagesWithContext(ctx context.Context, input *sagemaker.ListModelExplainabilityJobDefinitionsInput, fn func(*sagemaker.ListModelExplainabilityJobDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelExplainabilityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelExplainabilityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelExplainabilityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelExplainabilityJobDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelExplainabilityJobDefinitionsWithContext(ctx context.Context, input *sagemaker.ListModelExplainabilityJobDefinitionsInput, opts ...request.Option) (*sagemaker.ListModelExplainabilityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelExplainabilityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelExplainabilityJobDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelMetadata(input *sagemaker.ListModelMetadataInput) (*sagemaker.ListModelMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelMetadataOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelMetadata(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelMetadataPages(input *sagemaker.ListModelMetadataInput, fn func(*sagemaker.ListModelMetadataOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelMetadataOutput{}
	fnCacher := func(out *sagemaker.ListModelMetadataOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelMetadataPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelMetadataPagesWithContext(ctx context.Context, input *sagemaker.ListModelMetadataInput, fn func(*sagemaker.ListModelMetadataOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelMetadataOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelMetadataOutput{}
	fnCacher := func(out *sagemaker.ListModelMetadataOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelMetadataPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelMetadataWithContext(ctx context.Context, input *sagemaker.ListModelMetadataInput, opts ...request.Option) (*sagemaker.ListModelMetadataOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelMetadataOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelMetadataWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelPackageGroups(input *sagemaker.ListModelPackageGroupsInput) (*sagemaker.ListModelPackageGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelPackageGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelPackageGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelPackageGroupsPages(input *sagemaker.ListModelPackageGroupsInput, fn func(*sagemaker.ListModelPackageGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelPackageGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelPackageGroupsOutput{}
	fnCacher := func(out *sagemaker.ListModelPackageGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelPackageGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelPackageGroupsPagesWithContext(ctx context.Context, input *sagemaker.ListModelPackageGroupsInput, fn func(*sagemaker.ListModelPackageGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelPackageGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelPackageGroupsOutput{}
	fnCacher := func(out *sagemaker.ListModelPackageGroupsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelPackageGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelPackageGroupsWithContext(ctx context.Context, input *sagemaker.ListModelPackageGroupsInput, opts ...request.Option) (*sagemaker.ListModelPackageGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelPackageGroupsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelPackageGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelPackages(input *sagemaker.ListModelPackagesInput) (*sagemaker.ListModelPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelPackagesOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelPackages(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelPackagesPages(input *sagemaker.ListModelPackagesInput, fn func(*sagemaker.ListModelPackagesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelPackagesOutput{}
	fnCacher := func(out *sagemaker.ListModelPackagesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelPackagesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelPackagesPagesWithContext(ctx context.Context, input *sagemaker.ListModelPackagesInput, fn func(*sagemaker.ListModelPackagesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelPackagesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelPackagesOutput{}
	fnCacher := func(out *sagemaker.ListModelPackagesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelPackagesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelPackagesWithContext(ctx context.Context, input *sagemaker.ListModelPackagesInput, opts ...request.Option) (*sagemaker.ListModelPackagesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelPackagesOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelPackagesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelQualityJobDefinitions(input *sagemaker.ListModelQualityJobDefinitionsInput) (*sagemaker.ListModelQualityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelQualityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelQualityJobDefinitions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelQualityJobDefinitionsPages(input *sagemaker.ListModelQualityJobDefinitionsInput, fn func(*sagemaker.ListModelQualityJobDefinitionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelQualityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelQualityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelQualityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelQualityJobDefinitionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelQualityJobDefinitionsPagesWithContext(ctx context.Context, input *sagemaker.ListModelQualityJobDefinitionsInput, fn func(*sagemaker.ListModelQualityJobDefinitionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelQualityJobDefinitionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelQualityJobDefinitionsOutput{}
	fnCacher := func(out *sagemaker.ListModelQualityJobDefinitionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelQualityJobDefinitionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelQualityJobDefinitionsWithContext(ctx context.Context, input *sagemaker.ListModelQualityJobDefinitionsInput, opts ...request.Option) (*sagemaker.ListModelQualityJobDefinitionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelQualityJobDefinitionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelQualityJobDefinitionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModels(input *sagemaker.ListModelsInput) (*sagemaker.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListModelsPages(input *sagemaker.ListModelsInput, fn func(*sagemaker.ListModelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelsOutput{}
	fnCacher := func(out *sagemaker.ListModelsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelsPagesWithContext(ctx context.Context, input *sagemaker.ListModelsInput, fn func(*sagemaker.ListModelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListModelsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListModelsOutput{}
	fnCacher := func(out *sagemaker.ListModelsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListModelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListModelsWithContext(ctx context.Context, input *sagemaker.ListModelsInput, opts ...request.Option) (*sagemaker.ListModelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListModelsOutput), nil
	}
	out, err := c.SageMakerAPI.ListModelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringAlertHistory(input *sagemaker.ListMonitoringAlertHistoryInput) (*sagemaker.ListMonitoringAlertHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringAlertHistoryOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringAlertHistory(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringAlertHistoryPages(input *sagemaker.ListMonitoringAlertHistoryInput, fn func(*sagemaker.ListMonitoringAlertHistoryOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringAlertHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringAlertHistoryOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringAlertHistoryOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringAlertHistoryPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringAlertHistoryPagesWithContext(ctx context.Context, input *sagemaker.ListMonitoringAlertHistoryInput, fn func(*sagemaker.ListMonitoringAlertHistoryOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringAlertHistoryOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringAlertHistoryOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringAlertHistoryOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringAlertHistoryPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringAlertHistoryWithContext(ctx context.Context, input *sagemaker.ListMonitoringAlertHistoryInput, opts ...request.Option) (*sagemaker.ListMonitoringAlertHistoryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringAlertHistoryOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringAlertHistoryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringAlerts(input *sagemaker.ListMonitoringAlertsInput) (*sagemaker.ListMonitoringAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringAlertsOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringAlerts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringAlertsPages(input *sagemaker.ListMonitoringAlertsInput, fn func(*sagemaker.ListMonitoringAlertsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringAlertsOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringAlertsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringAlertsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringAlertsPagesWithContext(ctx context.Context, input *sagemaker.ListMonitoringAlertsInput, fn func(*sagemaker.ListMonitoringAlertsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringAlertsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringAlertsOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringAlertsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringAlertsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringAlertsWithContext(ctx context.Context, input *sagemaker.ListMonitoringAlertsInput, opts ...request.Option) (*sagemaker.ListMonitoringAlertsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringAlertsOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringAlertsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringExecutions(input *sagemaker.ListMonitoringExecutionsInput) (*sagemaker.ListMonitoringExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringExecutionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringExecutionsPages(input *sagemaker.ListMonitoringExecutionsInput, fn func(*sagemaker.ListMonitoringExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringExecutionsOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringExecutionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringExecutionsPagesWithContext(ctx context.Context, input *sagemaker.ListMonitoringExecutionsInput, fn func(*sagemaker.ListMonitoringExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringExecutionsOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringExecutionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringExecutionsWithContext(ctx context.Context, input *sagemaker.ListMonitoringExecutionsInput, opts ...request.Option) (*sagemaker.ListMonitoringExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringExecutionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringSchedules(input *sagemaker.ListMonitoringSchedulesInput) (*sagemaker.ListMonitoringSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringSchedulesOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringSchedules(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListMonitoringSchedulesPages(input *sagemaker.ListMonitoringSchedulesInput, fn func(*sagemaker.ListMonitoringSchedulesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringSchedulesOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringSchedulesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringSchedulesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringSchedulesPagesWithContext(ctx context.Context, input *sagemaker.ListMonitoringSchedulesInput, fn func(*sagemaker.ListMonitoringSchedulesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListMonitoringSchedulesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListMonitoringSchedulesOutput{}
	fnCacher := func(out *sagemaker.ListMonitoringSchedulesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListMonitoringSchedulesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListMonitoringSchedulesWithContext(ctx context.Context, input *sagemaker.ListMonitoringSchedulesInput, opts ...request.Option) (*sagemaker.ListMonitoringSchedulesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListMonitoringSchedulesOutput), nil
	}
	out, err := c.SageMakerAPI.ListMonitoringSchedulesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListNotebookInstanceLifecycleConfigs(input *sagemaker.ListNotebookInstanceLifecycleConfigsInput) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListNotebookInstanceLifecycleConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListNotebookInstanceLifecycleConfigsPages(input *sagemaker.ListNotebookInstanceLifecycleConfigsInput, fn func(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListNotebookInstanceLifecycleConfigsOutput{}
	fnCacher := func(out *sagemaker.ListNotebookInstanceLifecycleConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListNotebookInstanceLifecycleConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListNotebookInstanceLifecycleConfigsPagesWithContext(ctx context.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput, fn func(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListNotebookInstanceLifecycleConfigsOutput{}
	fnCacher := func(out *sagemaker.ListNotebookInstanceLifecycleConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListNotebookInstanceLifecycleConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListNotebookInstanceLifecycleConfigsWithContext(ctx context.Context, input *sagemaker.ListNotebookInstanceLifecycleConfigsInput, opts ...request.Option) (*sagemaker.ListNotebookInstanceLifecycleConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListNotebookInstanceLifecycleConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListNotebookInstanceLifecycleConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListNotebookInstances(input *sagemaker.ListNotebookInstancesInput) (*sagemaker.ListNotebookInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListNotebookInstancesOutput), nil
	}
	out, err := c.SageMakerAPI.ListNotebookInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListNotebookInstancesPages(input *sagemaker.ListNotebookInstancesInput, fn func(*sagemaker.ListNotebookInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListNotebookInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListNotebookInstancesOutput{}
	fnCacher := func(out *sagemaker.ListNotebookInstancesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListNotebookInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListNotebookInstancesPagesWithContext(ctx context.Context, input *sagemaker.ListNotebookInstancesInput, fn func(*sagemaker.ListNotebookInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListNotebookInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListNotebookInstancesOutput{}
	fnCacher := func(out *sagemaker.ListNotebookInstancesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListNotebookInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListNotebookInstancesWithContext(ctx context.Context, input *sagemaker.ListNotebookInstancesInput, opts ...request.Option) (*sagemaker.ListNotebookInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListNotebookInstancesOutput), nil
	}
	out, err := c.SageMakerAPI.ListNotebookInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineExecutionSteps(input *sagemaker.ListPipelineExecutionStepsInput) (*sagemaker.ListPipelineExecutionStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineExecutionStepsOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineExecutionSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineExecutionStepsPages(input *sagemaker.ListPipelineExecutionStepsInput, fn func(*sagemaker.ListPipelineExecutionStepsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineExecutionStepsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineExecutionStepsOutput{}
	fnCacher := func(out *sagemaker.ListPipelineExecutionStepsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineExecutionStepsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineExecutionStepsPagesWithContext(ctx context.Context, input *sagemaker.ListPipelineExecutionStepsInput, fn func(*sagemaker.ListPipelineExecutionStepsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineExecutionStepsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineExecutionStepsOutput{}
	fnCacher := func(out *sagemaker.ListPipelineExecutionStepsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineExecutionStepsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineExecutionStepsWithContext(ctx context.Context, input *sagemaker.ListPipelineExecutionStepsInput, opts ...request.Option) (*sagemaker.ListPipelineExecutionStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineExecutionStepsOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineExecutionStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineExecutions(input *sagemaker.ListPipelineExecutionsInput) (*sagemaker.ListPipelineExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineExecutionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineExecutionsPages(input *sagemaker.ListPipelineExecutionsInput, fn func(*sagemaker.ListPipelineExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineExecutionsOutput{}
	fnCacher := func(out *sagemaker.ListPipelineExecutionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineExecutionsPagesWithContext(ctx context.Context, input *sagemaker.ListPipelineExecutionsInput, fn func(*sagemaker.ListPipelineExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineExecutionsOutput{}
	fnCacher := func(out *sagemaker.ListPipelineExecutionsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineExecutionsWithContext(ctx context.Context, input *sagemaker.ListPipelineExecutionsInput, opts ...request.Option) (*sagemaker.ListPipelineExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineExecutionsOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineParametersForExecution(input *sagemaker.ListPipelineParametersForExecutionInput) (*sagemaker.ListPipelineParametersForExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineParametersForExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineParametersForExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelineParametersForExecutionPages(input *sagemaker.ListPipelineParametersForExecutionInput, fn func(*sagemaker.ListPipelineParametersForExecutionOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineParametersForExecutionOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineParametersForExecutionOutput{}
	fnCacher := func(out *sagemaker.ListPipelineParametersForExecutionOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineParametersForExecutionPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineParametersForExecutionPagesWithContext(ctx context.Context, input *sagemaker.ListPipelineParametersForExecutionInput, fn func(*sagemaker.ListPipelineParametersForExecutionOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelineParametersForExecutionOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelineParametersForExecutionOutput{}
	fnCacher := func(out *sagemaker.ListPipelineParametersForExecutionOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelineParametersForExecutionPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelineParametersForExecutionWithContext(ctx context.Context, input *sagemaker.ListPipelineParametersForExecutionInput, opts ...request.Option) (*sagemaker.ListPipelineParametersForExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelineParametersForExecutionOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelineParametersForExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelines(input *sagemaker.ListPipelinesInput) (*sagemaker.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelinesOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelines(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListPipelinesPages(input *sagemaker.ListPipelinesInput, fn func(*sagemaker.ListPipelinesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelinesOutput{}
	fnCacher := func(out *sagemaker.ListPipelinesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelinesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelinesPagesWithContext(ctx context.Context, input *sagemaker.ListPipelinesInput, fn func(*sagemaker.ListPipelinesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListPipelinesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListPipelinesOutput{}
	fnCacher := func(out *sagemaker.ListPipelinesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListPipelinesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListPipelinesWithContext(ctx context.Context, input *sagemaker.ListPipelinesInput, opts ...request.Option) (*sagemaker.ListPipelinesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListPipelinesOutput), nil
	}
	out, err := c.SageMakerAPI.ListPipelinesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListProcessingJobs(input *sagemaker.ListProcessingJobsInput) (*sagemaker.ListProcessingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListProcessingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListProcessingJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListProcessingJobsPages(input *sagemaker.ListProcessingJobsInput, fn func(*sagemaker.ListProcessingJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListProcessingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListProcessingJobsOutput{}
	fnCacher := func(out *sagemaker.ListProcessingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListProcessingJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListProcessingJobsPagesWithContext(ctx context.Context, input *sagemaker.ListProcessingJobsInput, fn func(*sagemaker.ListProcessingJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListProcessingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListProcessingJobsOutput{}
	fnCacher := func(out *sagemaker.ListProcessingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListProcessingJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListProcessingJobsWithContext(ctx context.Context, input *sagemaker.ListProcessingJobsInput, opts ...request.Option) (*sagemaker.ListProcessingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListProcessingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListProcessingJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListProjects(input *sagemaker.ListProjectsInput) (*sagemaker.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListProjectsOutput), nil
	}
	out, err := c.SageMakerAPI.ListProjects(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListProjectsPages(input *sagemaker.ListProjectsInput, fn func(*sagemaker.ListProjectsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListProjectsOutput{}
	fnCacher := func(out *sagemaker.ListProjectsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListProjectsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListProjectsPagesWithContext(ctx context.Context, input *sagemaker.ListProjectsInput, fn func(*sagemaker.ListProjectsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListProjectsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListProjectsOutput{}
	fnCacher := func(out *sagemaker.ListProjectsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListProjectsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListProjectsWithContext(ctx context.Context, input *sagemaker.ListProjectsInput, opts ...request.Option) (*sagemaker.ListProjectsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListProjectsOutput), nil
	}
	out, err := c.SageMakerAPI.ListProjectsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListSpaces(input *sagemaker.ListSpacesInput) (*sagemaker.ListSpacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListSpacesOutput), nil
	}
	out, err := c.SageMakerAPI.ListSpaces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListSpacesPages(input *sagemaker.ListSpacesInput, fn func(*sagemaker.ListSpacesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListSpacesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListSpacesOutput{}
	fnCacher := func(out *sagemaker.ListSpacesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListSpacesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListSpacesPagesWithContext(ctx context.Context, input *sagemaker.ListSpacesInput, fn func(*sagemaker.ListSpacesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListSpacesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListSpacesOutput{}
	fnCacher := func(out *sagemaker.ListSpacesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListSpacesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListSpacesWithContext(ctx context.Context, input *sagemaker.ListSpacesInput, opts ...request.Option) (*sagemaker.ListSpacesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListSpacesOutput), nil
	}
	out, err := c.SageMakerAPI.ListSpacesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListStageDevices(input *sagemaker.ListStageDevicesInput) (*sagemaker.ListStageDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListStageDevicesOutput), nil
	}
	out, err := c.SageMakerAPI.ListStageDevices(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListStageDevicesPages(input *sagemaker.ListStageDevicesInput, fn func(*sagemaker.ListStageDevicesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListStageDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListStageDevicesOutput{}
	fnCacher := func(out *sagemaker.ListStageDevicesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListStageDevicesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListStageDevicesPagesWithContext(ctx context.Context, input *sagemaker.ListStageDevicesInput, fn func(*sagemaker.ListStageDevicesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListStageDevicesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListStageDevicesOutput{}
	fnCacher := func(out *sagemaker.ListStageDevicesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListStageDevicesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListStageDevicesWithContext(ctx context.Context, input *sagemaker.ListStageDevicesInput, opts ...request.Option) (*sagemaker.ListStageDevicesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListStageDevicesOutput), nil
	}
	out, err := c.SageMakerAPI.ListStageDevicesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListStudioLifecycleConfigs(input *sagemaker.ListStudioLifecycleConfigsInput) (*sagemaker.ListStudioLifecycleConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListStudioLifecycleConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListStudioLifecycleConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListStudioLifecycleConfigsPages(input *sagemaker.ListStudioLifecycleConfigsInput, fn func(*sagemaker.ListStudioLifecycleConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListStudioLifecycleConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListStudioLifecycleConfigsOutput{}
	fnCacher := func(out *sagemaker.ListStudioLifecycleConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListStudioLifecycleConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListStudioLifecycleConfigsPagesWithContext(ctx context.Context, input *sagemaker.ListStudioLifecycleConfigsInput, fn func(*sagemaker.ListStudioLifecycleConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListStudioLifecycleConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListStudioLifecycleConfigsOutput{}
	fnCacher := func(out *sagemaker.ListStudioLifecycleConfigsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListStudioLifecycleConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListStudioLifecycleConfigsWithContext(ctx context.Context, input *sagemaker.ListStudioLifecycleConfigsInput, opts ...request.Option) (*sagemaker.ListStudioLifecycleConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListStudioLifecycleConfigsOutput), nil
	}
	out, err := c.SageMakerAPI.ListStudioLifecycleConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListSubscribedWorkteams(input *sagemaker.ListSubscribedWorkteamsInput) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListSubscribedWorkteamsOutput), nil
	}
	out, err := c.SageMakerAPI.ListSubscribedWorkteams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListSubscribedWorkteamsPages(input *sagemaker.ListSubscribedWorkteamsInput, fn func(*sagemaker.ListSubscribedWorkteamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListSubscribedWorkteamsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListSubscribedWorkteamsOutput{}
	fnCacher := func(out *sagemaker.ListSubscribedWorkteamsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListSubscribedWorkteamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListSubscribedWorkteamsPagesWithContext(ctx context.Context, input *sagemaker.ListSubscribedWorkteamsInput, fn func(*sagemaker.ListSubscribedWorkteamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListSubscribedWorkteamsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListSubscribedWorkteamsOutput{}
	fnCacher := func(out *sagemaker.ListSubscribedWorkteamsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListSubscribedWorkteamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListSubscribedWorkteamsWithContext(ctx context.Context, input *sagemaker.ListSubscribedWorkteamsInput, opts ...request.Option) (*sagemaker.ListSubscribedWorkteamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListSubscribedWorkteamsOutput), nil
	}
	out, err := c.SageMakerAPI.ListSubscribedWorkteamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTags(input *sagemaker.ListTagsInput) (*sagemaker.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTagsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTagsPages(input *sagemaker.ListTagsInput, fn func(*sagemaker.ListTagsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTagsOutput{}
	fnCacher := func(out *sagemaker.ListTagsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTagsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTagsPagesWithContext(ctx context.Context, input *sagemaker.ListTagsInput, fn func(*sagemaker.ListTagsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTagsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTagsOutput{}
	fnCacher := func(out *sagemaker.ListTagsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTagsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTagsWithContext(ctx context.Context, input *sagemaker.ListTagsInput, opts ...request.Option) (*sagemaker.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTagsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrainingJobs(input *sagemaker.ListTrainingJobsInput) (*sagemaker.ListTrainingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrainingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrainingJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrainingJobsForHyperParameterTuningJob(input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrainingJobsForHyperParameterTuningJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrainingJobsForHyperParameterTuningJobPages(input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, fn func(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput{}
	fnCacher := func(out *sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrainingJobsForHyperParameterTuningJobPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrainingJobsForHyperParameterTuningJobPagesWithContext(ctx context.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, fn func(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput{}
	fnCacher := func(out *sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrainingJobsForHyperParameterTuningJobPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrainingJobsForHyperParameterTuningJobWithContext(ctx context.Context, input *sagemaker.ListTrainingJobsForHyperParameterTuningJobInput, opts ...request.Option) (*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrainingJobsForHyperParameterTuningJobOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrainingJobsForHyperParameterTuningJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrainingJobsPages(input *sagemaker.ListTrainingJobsInput, fn func(*sagemaker.ListTrainingJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrainingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrainingJobsOutput{}
	fnCacher := func(out *sagemaker.ListTrainingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrainingJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrainingJobsPagesWithContext(ctx context.Context, input *sagemaker.ListTrainingJobsInput, fn func(*sagemaker.ListTrainingJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrainingJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrainingJobsOutput{}
	fnCacher := func(out *sagemaker.ListTrainingJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrainingJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrainingJobsWithContext(ctx context.Context, input *sagemaker.ListTrainingJobsInput, opts ...request.Option) (*sagemaker.ListTrainingJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrainingJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrainingJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTransformJobs(input *sagemaker.ListTransformJobsInput) (*sagemaker.ListTransformJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTransformJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTransformJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTransformJobsPages(input *sagemaker.ListTransformJobsInput, fn func(*sagemaker.ListTransformJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTransformJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTransformJobsOutput{}
	fnCacher := func(out *sagemaker.ListTransformJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTransformJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTransformJobsPagesWithContext(ctx context.Context, input *sagemaker.ListTransformJobsInput, fn func(*sagemaker.ListTransformJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTransformJobsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTransformJobsOutput{}
	fnCacher := func(out *sagemaker.ListTransformJobsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTransformJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTransformJobsWithContext(ctx context.Context, input *sagemaker.ListTransformJobsInput, opts ...request.Option) (*sagemaker.ListTransformJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTransformJobsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTransformJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrialComponents(input *sagemaker.ListTrialComponentsInput) (*sagemaker.ListTrialComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrialComponentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrialComponents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrialComponentsPages(input *sagemaker.ListTrialComponentsInput, fn func(*sagemaker.ListTrialComponentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrialComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrialComponentsOutput{}
	fnCacher := func(out *sagemaker.ListTrialComponentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrialComponentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrialComponentsPagesWithContext(ctx context.Context, input *sagemaker.ListTrialComponentsInput, fn func(*sagemaker.ListTrialComponentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrialComponentsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrialComponentsOutput{}
	fnCacher := func(out *sagemaker.ListTrialComponentsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrialComponentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrialComponentsWithContext(ctx context.Context, input *sagemaker.ListTrialComponentsInput, opts ...request.Option) (*sagemaker.ListTrialComponentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrialComponentsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrialComponentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrials(input *sagemaker.ListTrialsInput) (*sagemaker.ListTrialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrialsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListTrialsPages(input *sagemaker.ListTrialsInput, fn func(*sagemaker.ListTrialsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrialsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrialsOutput{}
	fnCacher := func(out *sagemaker.ListTrialsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrialsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrialsPagesWithContext(ctx context.Context, input *sagemaker.ListTrialsInput, fn func(*sagemaker.ListTrialsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListTrialsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListTrialsOutput{}
	fnCacher := func(out *sagemaker.ListTrialsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListTrialsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListTrialsWithContext(ctx context.Context, input *sagemaker.ListTrialsInput, opts ...request.Option) (*sagemaker.ListTrialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListTrialsOutput), nil
	}
	out, err := c.SageMakerAPI.ListTrialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListUserProfiles(input *sagemaker.ListUserProfilesInput) (*sagemaker.ListUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListUserProfilesOutput), nil
	}
	out, err := c.SageMakerAPI.ListUserProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListUserProfilesPages(input *sagemaker.ListUserProfilesInput, fn func(*sagemaker.ListUserProfilesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListUserProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListUserProfilesOutput{}
	fnCacher := func(out *sagemaker.ListUserProfilesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListUserProfilesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListUserProfilesPagesWithContext(ctx context.Context, input *sagemaker.ListUserProfilesInput, fn func(*sagemaker.ListUserProfilesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListUserProfilesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListUserProfilesOutput{}
	fnCacher := func(out *sagemaker.ListUserProfilesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListUserProfilesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListUserProfilesWithContext(ctx context.Context, input *sagemaker.ListUserProfilesInput, opts ...request.Option) (*sagemaker.ListUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListUserProfilesOutput), nil
	}
	out, err := c.SageMakerAPI.ListUserProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListWorkforces(input *sagemaker.ListWorkforcesInput) (*sagemaker.ListWorkforcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListWorkforcesOutput), nil
	}
	out, err := c.SageMakerAPI.ListWorkforces(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListWorkforcesPages(input *sagemaker.ListWorkforcesInput, fn func(*sagemaker.ListWorkforcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListWorkforcesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListWorkforcesOutput{}
	fnCacher := func(out *sagemaker.ListWorkforcesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListWorkforcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListWorkforcesPagesWithContext(ctx context.Context, input *sagemaker.ListWorkforcesInput, fn func(*sagemaker.ListWorkforcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListWorkforcesOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListWorkforcesOutput{}
	fnCacher := func(out *sagemaker.ListWorkforcesOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListWorkforcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListWorkforcesWithContext(ctx context.Context, input *sagemaker.ListWorkforcesInput, opts ...request.Option) (*sagemaker.ListWorkforcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListWorkforcesOutput), nil
	}
	out, err := c.SageMakerAPI.ListWorkforcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListWorkteams(input *sagemaker.ListWorkteamsInput) (*sagemaker.ListWorkteamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListWorkteamsOutput), nil
	}
	out, err := c.SageMakerAPI.ListWorkteams(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) ListWorkteamsPages(input *sagemaker.ListWorkteamsInput, fn func(*sagemaker.ListWorkteamsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListWorkteamsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListWorkteamsOutput{}
	fnCacher := func(out *sagemaker.ListWorkteamsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListWorkteamsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListWorkteamsPagesWithContext(ctx context.Context, input *sagemaker.ListWorkteamsInput, fn func(*sagemaker.ListWorkteamsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.ListWorkteamsOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.ListWorkteamsOutput{}
	fnCacher := func(out *sagemaker.ListWorkteamsOutput, more bool) bool {
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
	if err := c.SageMakerAPI.ListWorkteamsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) ListWorkteamsWithContext(ctx context.Context, input *sagemaker.ListWorkteamsInput, opts ...request.Option) (*sagemaker.ListWorkteamsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.ListWorkteamsOutput), nil
	}
	out, err := c.SageMakerAPI.ListWorkteamsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) QueryLineage(input *sagemaker.QueryLineageInput) (*sagemaker.QueryLineageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.QueryLineageOutput), nil
	}
	out, err := c.SageMakerAPI.QueryLineage(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) QueryLineagePages(input *sagemaker.QueryLineageInput, fn func(*sagemaker.QueryLineageOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.QueryLineageOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.QueryLineageOutput{}
	fnCacher := func(out *sagemaker.QueryLineageOutput, more bool) bool {
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
	if err := c.SageMakerAPI.QueryLineagePages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) QueryLineagePagesWithContext(ctx context.Context, input *sagemaker.QueryLineageInput, fn func(*sagemaker.QueryLineageOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.QueryLineageOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.QueryLineageOutput{}
	fnCacher := func(out *sagemaker.QueryLineageOutput, more bool) bool {
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
	if err := c.SageMakerAPI.QueryLineagePagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) QueryLineageWithContext(ctx context.Context, input *sagemaker.QueryLineageInput, opts ...request.Option) (*sagemaker.QueryLineageOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.QueryLineageOutput), nil
	}
	out, err := c.SageMakerAPI.QueryLineageWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) Search(input *sagemaker.SearchInput) (*sagemaker.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.SearchOutput), nil
	}
	out, err := c.SageMakerAPI.Search(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *SageMaker) SearchPages(input *sagemaker.SearchInput, fn func(*sagemaker.SearchOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.SearchOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.SearchOutput{}
	fnCacher := func(out *sagemaker.SearchOutput, more bool) bool {
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
	if err := c.SageMakerAPI.SearchPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) SearchPagesWithContext(ctx context.Context, input *sagemaker.SearchInput, fn func(*sagemaker.SearchOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*sagemaker.SearchOutput), false)
		return nil
	}
	cachable := true
	output := &sagemaker.SearchOutput{}
	fnCacher := func(out *sagemaker.SearchOutput, more bool) bool {
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
	if err := c.SageMakerAPI.SearchPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *SageMaker) SearchWithContext(ctx context.Context, input *sagemaker.SearchInput, opts ...request.Option) (*sagemaker.SearchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*sagemaker.SearchOutput), nil
	}
	out, err := c.SageMakerAPI.SearchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
