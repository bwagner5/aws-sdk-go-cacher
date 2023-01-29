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
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"github.com/aws/aws-sdk-go/service/comprehendmedical/comprehendmedicaliface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type ComprehendMedical struct {
	comprehendmedicaliface.ComprehendMedicalAPI
	cache *cache.Cache
}

func NewComprehendMedical(comprehendmedicalapi comprehendmedicaliface.ComprehendMedicalAPI) *ComprehendMedical {
	return &ComprehendMedical{
		ComprehendMedicalAPI: comprehendmedicalapi,
		cache:                cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *ComprehendMedical) DescribeEntitiesDetectionV2Job(input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeEntitiesDetectionV2Job(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeEntitiesDetectionV2JobWithContext(ctx context.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput, opts ...request.Option) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeEntitiesDetectionV2JobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeEntitiesDetectionV2JobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeICD10CMInferenceJob(input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeICD10CMInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeICD10CMInferenceJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeICD10CMInferenceJobWithContext(ctx context.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput, opts ...request.Option) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeICD10CMInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeICD10CMInferenceJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribePHIDetectionJob(input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribePHIDetectionJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribePHIDetectionJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribePHIDetectionJobWithContext(ctx context.Context, input *comprehendmedical.DescribePHIDetectionJobInput, opts ...request.Option) (*comprehendmedical.DescribePHIDetectionJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribePHIDetectionJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribePHIDetectionJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeRxNormInferenceJob(input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeRxNormInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeRxNormInferenceJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeRxNormInferenceJobWithContext(ctx context.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput, opts ...request.Option) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeRxNormInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeRxNormInferenceJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeSNOMEDCTInferenceJob(input *comprehendmedical.DescribeSNOMEDCTInferenceJobInput) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeSNOMEDCTInferenceJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) DescribeSNOMEDCTInferenceJobWithContext(ctx context.Context, input *comprehendmedical.DescribeSNOMEDCTInferenceJobInput, opts ...request.Option) (*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.DescribeSNOMEDCTInferenceJobOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.DescribeSNOMEDCTInferenceJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListEntitiesDetectionV2Jobs(input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListEntitiesDetectionV2JobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListEntitiesDetectionV2Jobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListEntitiesDetectionV2JobsWithContext(ctx context.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput, opts ...request.Option) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListEntitiesDetectionV2JobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListEntitiesDetectionV2JobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListICD10CMInferenceJobs(input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListICD10CMInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListICD10CMInferenceJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListICD10CMInferenceJobsWithContext(ctx context.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput, opts ...request.Option) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListICD10CMInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListICD10CMInferenceJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListPHIDetectionJobs(input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListPHIDetectionJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListPHIDetectionJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListPHIDetectionJobsWithContext(ctx context.Context, input *comprehendmedical.ListPHIDetectionJobsInput, opts ...request.Option) (*comprehendmedical.ListPHIDetectionJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListPHIDetectionJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListPHIDetectionJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListRxNormInferenceJobs(input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListRxNormInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListRxNormInferenceJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListRxNormInferenceJobsWithContext(ctx context.Context, input *comprehendmedical.ListRxNormInferenceJobsInput, opts ...request.Option) (*comprehendmedical.ListRxNormInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListRxNormInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListRxNormInferenceJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListSNOMEDCTInferenceJobs(input *comprehendmedical.ListSNOMEDCTInferenceJobsInput) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListSNOMEDCTInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListSNOMEDCTInferenceJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *ComprehendMedical) ListSNOMEDCTInferenceJobsWithContext(ctx context.Context, input *comprehendmedical.ListSNOMEDCTInferenceJobsInput, opts ...request.Option) (*comprehendmedical.ListSNOMEDCTInferenceJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*comprehendmedical.ListSNOMEDCTInferenceJobsOutput), nil
	}
	out, err := c.ComprehendMedicalAPI.ListSNOMEDCTInferenceJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
