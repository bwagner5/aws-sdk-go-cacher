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
package robomakercacher

// DO NOT EDIT
// THIS FILE IS AUTO GENERATED
import (
	"context"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/robomaker"
	"github.com/aws/aws-sdk-go/service/robomaker/robomakeriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type RoboMaker struct {
	robomakeriface.RoboMakerAPI
	cache *cache.Cache
}

func New(robomakerapi robomakeriface.RoboMakerAPI) *RoboMaker {
	return &RoboMaker{
		RoboMakerAPI: robomakerapi,
		cache:        cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *RoboMaker) BatchDeleteWorlds(input *robomaker.BatchDeleteWorldsInput) (*robomaker.BatchDeleteWorldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.BatchDeleteWorldsOutput), nil
	}
	out, err := c.RoboMakerAPI.BatchDeleteWorlds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) BatchDeleteWorldsWithContext(ctx context.Context, input *robomaker.BatchDeleteWorldsInput, opts ...request.Option) (*robomaker.BatchDeleteWorldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.BatchDeleteWorldsOutput), nil
	}
	out, err := c.RoboMakerAPI.BatchDeleteWorldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) BatchDescribeSimulationJob(input *robomaker.BatchDescribeSimulationJobInput) (*robomaker.BatchDescribeSimulationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.BatchDescribeSimulationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.BatchDescribeSimulationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) BatchDescribeSimulationJobWithContext(ctx context.Context, input *robomaker.BatchDescribeSimulationJobInput, opts ...request.Option) (*robomaker.BatchDescribeSimulationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.BatchDescribeSimulationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.BatchDescribeSimulationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeDeploymentJob(input *robomaker.DescribeDeploymentJobInput) (*robomaker.DescribeDeploymentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeDeploymentJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeDeploymentJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeDeploymentJobWithContext(ctx context.Context, input *robomaker.DescribeDeploymentJobInput, opts ...request.Option) (*robomaker.DescribeDeploymentJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeDeploymentJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeDeploymentJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeFleet(input *robomaker.DescribeFleetInput) (*robomaker.DescribeFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeFleetOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeFleet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeFleetWithContext(ctx context.Context, input *robomaker.DescribeFleetInput, opts ...request.Option) (*robomaker.DescribeFleetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeFleetOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeFleetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeRobot(input *robomaker.DescribeRobotInput) (*robomaker.DescribeRobotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeRobotOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeRobot(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeRobotApplication(input *robomaker.DescribeRobotApplicationInput) (*robomaker.DescribeRobotApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeRobotApplicationOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeRobotApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeRobotApplicationWithContext(ctx context.Context, input *robomaker.DescribeRobotApplicationInput, opts ...request.Option) (*robomaker.DescribeRobotApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeRobotApplicationOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeRobotApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeRobotWithContext(ctx context.Context, input *robomaker.DescribeRobotInput, opts ...request.Option) (*robomaker.DescribeRobotOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeRobotOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeRobotWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationApplication(input *robomaker.DescribeSimulationApplicationInput) (*robomaker.DescribeSimulationApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationApplicationOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationApplicationWithContext(ctx context.Context, input *robomaker.DescribeSimulationApplicationInput, opts ...request.Option) (*robomaker.DescribeSimulationApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationApplicationOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationJob(input *robomaker.DescribeSimulationJobInput) (*robomaker.DescribeSimulationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationJobBatch(input *robomaker.DescribeSimulationJobBatchInput) (*robomaker.DescribeSimulationJobBatchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationJobBatchOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationJobBatch(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationJobBatchWithContext(ctx context.Context, input *robomaker.DescribeSimulationJobBatchInput, opts ...request.Option) (*robomaker.DescribeSimulationJobBatchOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationJobBatchOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationJobBatchWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeSimulationJobWithContext(ctx context.Context, input *robomaker.DescribeSimulationJobInput, opts ...request.Option) (*robomaker.DescribeSimulationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeSimulationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeSimulationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorld(input *robomaker.DescribeWorldInput) (*robomaker.DescribeWorldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorld(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldExportJob(input *robomaker.DescribeWorldExportJobInput) (*robomaker.DescribeWorldExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldExportJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldExportJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldExportJobWithContext(ctx context.Context, input *robomaker.DescribeWorldExportJobInput, opts ...request.Option) (*robomaker.DescribeWorldExportJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldExportJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldExportJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldGenerationJob(input *robomaker.DescribeWorldGenerationJobInput) (*robomaker.DescribeWorldGenerationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldGenerationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldGenerationJob(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldGenerationJobWithContext(ctx context.Context, input *robomaker.DescribeWorldGenerationJobInput, opts ...request.Option) (*robomaker.DescribeWorldGenerationJobOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldGenerationJobOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldGenerationJobWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldTemplate(input *robomaker.DescribeWorldTemplateInput) (*robomaker.DescribeWorldTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldTemplateOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldTemplateWithContext(ctx context.Context, input *robomaker.DescribeWorldTemplateInput, opts ...request.Option) (*robomaker.DescribeWorldTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldTemplateOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) DescribeWorldWithContext(ctx context.Context, input *robomaker.DescribeWorldInput, opts ...request.Option) (*robomaker.DescribeWorldOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.DescribeWorldOutput), nil
	}
	out, err := c.RoboMakerAPI.DescribeWorldWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) GetWorldTemplateBody(input *robomaker.GetWorldTemplateBodyInput) (*robomaker.GetWorldTemplateBodyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.GetWorldTemplateBodyOutput), nil
	}
	out, err := c.RoboMakerAPI.GetWorldTemplateBody(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) GetWorldTemplateBodyWithContext(ctx context.Context, input *robomaker.GetWorldTemplateBodyInput, opts ...request.Option) (*robomaker.GetWorldTemplateBodyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.GetWorldTemplateBodyOutput), nil
	}
	out, err := c.RoboMakerAPI.GetWorldTemplateBodyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListDeploymentJobs(input *robomaker.ListDeploymentJobsInput) (*robomaker.ListDeploymentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListDeploymentJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListDeploymentJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListDeploymentJobsPages(input *robomaker.ListDeploymentJobsInput, fn func(*robomaker.ListDeploymentJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListDeploymentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListDeploymentJobsOutput{}
	fnCacher := func(out *robomaker.ListDeploymentJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListDeploymentJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListDeploymentJobsPagesWithContext(ctx context.Context, input *robomaker.ListDeploymentJobsInput, fn func(*robomaker.ListDeploymentJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListDeploymentJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListDeploymentJobsOutput{}
	fnCacher := func(out *robomaker.ListDeploymentJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListDeploymentJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListDeploymentJobsWithContext(ctx context.Context, input *robomaker.ListDeploymentJobsInput, opts ...request.Option) (*robomaker.ListDeploymentJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListDeploymentJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListDeploymentJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListFleets(input *robomaker.ListFleetsInput) (*robomaker.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListFleetsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListFleetsPages(input *robomaker.ListFleetsInput, fn func(*robomaker.ListFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListFleetsOutput{}
	fnCacher := func(out *robomaker.ListFleetsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListFleetsPagesWithContext(ctx context.Context, input *robomaker.ListFleetsInput, fn func(*robomaker.ListFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListFleetsOutput{}
	fnCacher := func(out *robomaker.ListFleetsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListFleetsWithContext(ctx context.Context, input *robomaker.ListFleetsInput, opts ...request.Option) (*robomaker.ListFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListFleetsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListRobotApplications(input *robomaker.ListRobotApplicationsInput) (*robomaker.ListRobotApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListRobotApplicationsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListRobotApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListRobotApplicationsPages(input *robomaker.ListRobotApplicationsInput, fn func(*robomaker.ListRobotApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListRobotApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListRobotApplicationsOutput{}
	fnCacher := func(out *robomaker.ListRobotApplicationsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListRobotApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListRobotApplicationsPagesWithContext(ctx context.Context, input *robomaker.ListRobotApplicationsInput, fn func(*robomaker.ListRobotApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListRobotApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListRobotApplicationsOutput{}
	fnCacher := func(out *robomaker.ListRobotApplicationsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListRobotApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListRobotApplicationsWithContext(ctx context.Context, input *robomaker.ListRobotApplicationsInput, opts ...request.Option) (*robomaker.ListRobotApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListRobotApplicationsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListRobotApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListRobots(input *robomaker.ListRobotsInput) (*robomaker.ListRobotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListRobotsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListRobots(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListRobotsPages(input *robomaker.ListRobotsInput, fn func(*robomaker.ListRobotsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListRobotsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListRobotsOutput{}
	fnCacher := func(out *robomaker.ListRobotsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListRobotsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListRobotsPagesWithContext(ctx context.Context, input *robomaker.ListRobotsInput, fn func(*robomaker.ListRobotsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListRobotsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListRobotsOutput{}
	fnCacher := func(out *robomaker.ListRobotsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListRobotsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListRobotsWithContext(ctx context.Context, input *robomaker.ListRobotsInput, opts ...request.Option) (*robomaker.ListRobotsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListRobotsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListRobotsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationApplications(input *robomaker.ListSimulationApplicationsInput) (*robomaker.ListSimulationApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationApplicationsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationApplicationsPages(input *robomaker.ListSimulationApplicationsInput, fn func(*robomaker.ListSimulationApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationApplicationsOutput{}
	fnCacher := func(out *robomaker.ListSimulationApplicationsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationApplicationsPagesWithContext(ctx context.Context, input *robomaker.ListSimulationApplicationsInput, fn func(*robomaker.ListSimulationApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationApplicationsOutput{}
	fnCacher := func(out *robomaker.ListSimulationApplicationsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationApplicationsWithContext(ctx context.Context, input *robomaker.ListSimulationApplicationsInput, opts ...request.Option) (*robomaker.ListSimulationApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationApplicationsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationJobBatches(input *robomaker.ListSimulationJobBatchesInput) (*robomaker.ListSimulationJobBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationJobBatchesOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationJobBatches(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationJobBatchesPages(input *robomaker.ListSimulationJobBatchesInput, fn func(*robomaker.ListSimulationJobBatchesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationJobBatchesOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationJobBatchesOutput{}
	fnCacher := func(out *robomaker.ListSimulationJobBatchesOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationJobBatchesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationJobBatchesPagesWithContext(ctx context.Context, input *robomaker.ListSimulationJobBatchesInput, fn func(*robomaker.ListSimulationJobBatchesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationJobBatchesOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationJobBatchesOutput{}
	fnCacher := func(out *robomaker.ListSimulationJobBatchesOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationJobBatchesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationJobBatchesWithContext(ctx context.Context, input *robomaker.ListSimulationJobBatchesInput, opts ...request.Option) (*robomaker.ListSimulationJobBatchesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationJobBatchesOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationJobBatchesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationJobs(input *robomaker.ListSimulationJobsInput) (*robomaker.ListSimulationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListSimulationJobsPages(input *robomaker.ListSimulationJobsInput, fn func(*robomaker.ListSimulationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationJobsOutput{}
	fnCacher := func(out *robomaker.ListSimulationJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationJobsPagesWithContext(ctx context.Context, input *robomaker.ListSimulationJobsInput, fn func(*robomaker.ListSimulationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListSimulationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListSimulationJobsOutput{}
	fnCacher := func(out *robomaker.ListSimulationJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListSimulationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListSimulationJobsWithContext(ctx context.Context, input *robomaker.ListSimulationJobsInput, opts ...request.Option) (*robomaker.ListSimulationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListSimulationJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListSimulationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListTagsForResource(input *robomaker.ListTagsForResourceInput) (*robomaker.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListTagsForResourceOutput), nil
	}
	out, err := c.RoboMakerAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListTagsForResourceWithContext(ctx context.Context, input *robomaker.ListTagsForResourceInput, opts ...request.Option) (*robomaker.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListTagsForResourceOutput), nil
	}
	out, err := c.RoboMakerAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldExportJobs(input *robomaker.ListWorldExportJobsInput) (*robomaker.ListWorldExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldExportJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldExportJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldExportJobsPages(input *robomaker.ListWorldExportJobsInput, fn func(*robomaker.ListWorldExportJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldExportJobsOutput{}
	fnCacher := func(out *robomaker.ListWorldExportJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldExportJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldExportJobsPagesWithContext(ctx context.Context, input *robomaker.ListWorldExportJobsInput, fn func(*robomaker.ListWorldExportJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldExportJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldExportJobsOutput{}
	fnCacher := func(out *robomaker.ListWorldExportJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldExportJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldExportJobsWithContext(ctx context.Context, input *robomaker.ListWorldExportJobsInput, opts ...request.Option) (*robomaker.ListWorldExportJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldExportJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldExportJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldGenerationJobs(input *robomaker.ListWorldGenerationJobsInput) (*robomaker.ListWorldGenerationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldGenerationJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldGenerationJobs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldGenerationJobsPages(input *robomaker.ListWorldGenerationJobsInput, fn func(*robomaker.ListWorldGenerationJobsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldGenerationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldGenerationJobsOutput{}
	fnCacher := func(out *robomaker.ListWorldGenerationJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldGenerationJobsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldGenerationJobsPagesWithContext(ctx context.Context, input *robomaker.ListWorldGenerationJobsInput, fn func(*robomaker.ListWorldGenerationJobsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldGenerationJobsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldGenerationJobsOutput{}
	fnCacher := func(out *robomaker.ListWorldGenerationJobsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldGenerationJobsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldGenerationJobsWithContext(ctx context.Context, input *robomaker.ListWorldGenerationJobsInput, opts ...request.Option) (*robomaker.ListWorldGenerationJobsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldGenerationJobsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldGenerationJobsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldTemplates(input *robomaker.ListWorldTemplatesInput) (*robomaker.ListWorldTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldTemplatesOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldTemplates(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldTemplatesPages(input *robomaker.ListWorldTemplatesInput, fn func(*robomaker.ListWorldTemplatesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldTemplatesOutput{}
	fnCacher := func(out *robomaker.ListWorldTemplatesOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldTemplatesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldTemplatesPagesWithContext(ctx context.Context, input *robomaker.ListWorldTemplatesInput, fn func(*robomaker.ListWorldTemplatesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldTemplatesOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldTemplatesOutput{}
	fnCacher := func(out *robomaker.ListWorldTemplatesOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldTemplatesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldTemplatesWithContext(ctx context.Context, input *robomaker.ListWorldTemplatesInput, opts ...request.Option) (*robomaker.ListWorldTemplatesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldTemplatesOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldTemplatesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorlds(input *robomaker.ListWorldsInput) (*robomaker.ListWorldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorlds(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *RoboMaker) ListWorldsPages(input *robomaker.ListWorldsInput, fn func(*robomaker.ListWorldsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldsOutput{}
	fnCacher := func(out *robomaker.ListWorldsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldsPagesWithContext(ctx context.Context, input *robomaker.ListWorldsInput, fn func(*robomaker.ListWorldsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*robomaker.ListWorldsOutput), false)
		return nil
	}
	cachable := true
	output := &robomaker.ListWorldsOutput{}
	fnCacher := func(out *robomaker.ListWorldsOutput, more bool) bool {
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
	if err := c.RoboMakerAPI.ListWorldsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *RoboMaker) ListWorldsWithContext(ctx context.Context, input *robomaker.ListWorldsInput, opts ...request.Option) (*robomaker.ListWorldsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*robomaker.ListWorldsOutput), nil
	}
	out, err := c.RoboMakerAPI.ListWorldsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
