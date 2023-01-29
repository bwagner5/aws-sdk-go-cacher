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
package iotjobsdataplanecacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane"
	"github.com/aws/aws-sdk-go/service/iotjobsdataplane/iotjobsdataplaneiface"
	
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type IoTJobsDataPlane struct {
	iotjobsdataplaneiface.IoTJobsDataPlaneAPI
	cache *cache.Cache
}

func New(iotjobsdataplaneapi iotjobsdataplaneiface.IoTJobsDataPlaneAPI) *IoTJobsDataPlane {
	return &IoTJobsDataPlane{
		IoTJobsDataPlaneAPI: iotjobsdataplaneapi,
		cache:               cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *IoTJobsDataPlane) DescribeJobExecution(input *iotjobsdataplane.DescribeJobExecutionInput) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotjobsdataplane.DescribeJobExecutionOutput), nil
	}
	out, err := c.IoTJobsDataPlaneAPI.DescribeJobExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTJobsDataPlane) DescribeJobExecutionWithContext(ctx context.Context, input *iotjobsdataplane.DescribeJobExecutionInput, opts ...request.Option) (*iotjobsdataplane.DescribeJobExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotjobsdataplane.DescribeJobExecutionOutput), nil
	}
	out, err := c.IoTJobsDataPlaneAPI.DescribeJobExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTJobsDataPlane) GetPendingJobExecutions(input *iotjobsdataplane.GetPendingJobExecutionsInput) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotjobsdataplane.GetPendingJobExecutionsOutput), nil
	}
	out, err := c.IoTJobsDataPlaneAPI.GetPendingJobExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *IoTJobsDataPlane) GetPendingJobExecutionsWithContext(ctx context.Context, input *iotjobsdataplane.GetPendingJobExecutionsInput, opts ...request.Option) (*iotjobsdataplane.GetPendingJobExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*iotjobsdataplane.GetPendingJobExecutionsOutput), nil
	}
	out, err := c.IoTJobsDataPlaneAPI.GetPendingJobExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
