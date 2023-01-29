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
	"github.com/aws/aws-sdk-go/service/codedeploy"
	"github.com/aws/aws-sdk-go/service/codedeploy/codedeployiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CodeDeploy struct {
	codedeployiface.CodeDeployAPI
	cache *cache.Cache
}

func NewCodeDeploy(codedeployapi codedeployiface.CodeDeployAPI) *CodeDeploy {
	return &CodeDeploy{
		CodeDeployAPI: codedeployapi,
		cache:         cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CodeDeploy) BatchGetApplicationRevisions(input *codedeploy.BatchGetApplicationRevisionsInput) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetApplicationRevisionsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetApplicationRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetApplicationRevisionsWithContext(ctx context.Context, input *codedeploy.BatchGetApplicationRevisionsInput, opts ...request.Option) (*codedeploy.BatchGetApplicationRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetApplicationRevisionsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetApplicationRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetApplications(input *codedeploy.BatchGetApplicationsInput) (*codedeploy.BatchGetApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetApplicationsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetApplicationsWithContext(ctx context.Context, input *codedeploy.BatchGetApplicationsInput, opts ...request.Option) (*codedeploy.BatchGetApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetApplicationsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentGroups(input *codedeploy.BatchGetDeploymentGroupsInput) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentGroupsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentGroupsWithContext(ctx context.Context, input *codedeploy.BatchGetDeploymentGroupsInput, opts ...request.Option) (*codedeploy.BatchGetDeploymentGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentGroupsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentInstances(input *codedeploy.BatchGetDeploymentInstancesInput) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentInstancesWithContext(ctx context.Context, input *codedeploy.BatchGetDeploymentInstancesInput, opts ...request.Option) (*codedeploy.BatchGetDeploymentInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentTargets(input *codedeploy.BatchGetDeploymentTargetsInput) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentTargetsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentTargetsWithContext(ctx context.Context, input *codedeploy.BatchGetDeploymentTargetsInput, opts ...request.Option) (*codedeploy.BatchGetDeploymentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentTargetsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeployments(input *codedeploy.BatchGetDeploymentsInput) (*codedeploy.BatchGetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetDeploymentsWithContext(ctx context.Context, input *codedeploy.BatchGetDeploymentsInput, opts ...request.Option) (*codedeploy.BatchGetDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetDeploymentsOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetOnPremisesInstances(input *codedeploy.BatchGetOnPremisesInstancesInput) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetOnPremisesInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetOnPremisesInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) BatchGetOnPremisesInstancesWithContext(ctx context.Context, input *codedeploy.BatchGetOnPremisesInstancesInput, opts ...request.Option) (*codedeploy.BatchGetOnPremisesInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.BatchGetOnPremisesInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.BatchGetOnPremisesInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetApplication(input *codedeploy.GetApplicationInput) (*codedeploy.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetApplicationOutput), nil
	}
	out, err := c.CodeDeployAPI.GetApplication(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetApplicationRevision(input *codedeploy.GetApplicationRevisionInput) (*codedeploy.GetApplicationRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetApplicationRevisionOutput), nil
	}
	out, err := c.CodeDeployAPI.GetApplicationRevision(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetApplicationRevisionWithContext(ctx context.Context, input *codedeploy.GetApplicationRevisionInput, opts ...request.Option) (*codedeploy.GetApplicationRevisionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetApplicationRevisionOutput), nil
	}
	out, err := c.CodeDeployAPI.GetApplicationRevisionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetApplicationWithContext(ctx context.Context, input *codedeploy.GetApplicationInput, opts ...request.Option) (*codedeploy.GetApplicationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetApplicationOutput), nil
	}
	out, err := c.CodeDeployAPI.GetApplicationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeployment(input *codedeploy.GetDeploymentInput) (*codedeploy.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeployment(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentConfig(input *codedeploy.GetDeploymentConfigInput) (*codedeploy.GetDeploymentConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentConfigOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentConfig(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentConfigWithContext(ctx context.Context, input *codedeploy.GetDeploymentConfigInput, opts ...request.Option) (*codedeploy.GetDeploymentConfigOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentConfigOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentConfigWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentGroup(input *codedeploy.GetDeploymentGroupInput) (*codedeploy.GetDeploymentGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentGroupOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentGroup(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentGroupWithContext(ctx context.Context, input *codedeploy.GetDeploymentGroupInput, opts ...request.Option) (*codedeploy.GetDeploymentGroupOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentGroupOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentGroupWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentInstance(input *codedeploy.GetDeploymentInstanceInput) (*codedeploy.GetDeploymentInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentInstanceOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentInstanceWithContext(ctx context.Context, input *codedeploy.GetDeploymentInstanceInput, opts ...request.Option) (*codedeploy.GetDeploymentInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentInstanceOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentTarget(input *codedeploy.GetDeploymentTargetInput) (*codedeploy.GetDeploymentTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentTargetOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentTarget(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentTargetWithContext(ctx context.Context, input *codedeploy.GetDeploymentTargetInput, opts ...request.Option) (*codedeploy.GetDeploymentTargetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentTargetOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentTargetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetDeploymentWithContext(ctx context.Context, input *codedeploy.GetDeploymentInput, opts ...request.Option) (*codedeploy.GetDeploymentOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetDeploymentOutput), nil
	}
	out, err := c.CodeDeployAPI.GetDeploymentWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetOnPremisesInstance(input *codedeploy.GetOnPremisesInstanceInput) (*codedeploy.GetOnPremisesInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetOnPremisesInstanceOutput), nil
	}
	out, err := c.CodeDeployAPI.GetOnPremisesInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) GetOnPremisesInstanceWithContext(ctx context.Context, input *codedeploy.GetOnPremisesInstanceInput, opts ...request.Option) (*codedeploy.GetOnPremisesInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.GetOnPremisesInstanceOutput), nil
	}
	out, err := c.CodeDeployAPI.GetOnPremisesInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListApplicationRevisions(input *codedeploy.ListApplicationRevisionsInput) (*codedeploy.ListApplicationRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListApplicationRevisionsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListApplicationRevisions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListApplicationRevisionsPages(input *codedeploy.ListApplicationRevisionsInput, fn func(*codedeploy.ListApplicationRevisionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListApplicationRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListApplicationRevisionsOutput{}
	fnCacher := func(out *codedeploy.ListApplicationRevisionsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListApplicationRevisionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListApplicationRevisionsPagesWithContext(ctx context.Context, input *codedeploy.ListApplicationRevisionsInput, fn func(*codedeploy.ListApplicationRevisionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListApplicationRevisionsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListApplicationRevisionsOutput{}
	fnCacher := func(out *codedeploy.ListApplicationRevisionsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListApplicationRevisionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListApplicationRevisionsWithContext(ctx context.Context, input *codedeploy.ListApplicationRevisionsInput, opts ...request.Option) (*codedeploy.ListApplicationRevisionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListApplicationRevisionsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListApplicationRevisionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListApplications(input *codedeploy.ListApplicationsInput) (*codedeploy.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListApplicationsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListApplications(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListApplicationsPages(input *codedeploy.ListApplicationsInput, fn func(*codedeploy.ListApplicationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListApplicationsOutput{}
	fnCacher := func(out *codedeploy.ListApplicationsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListApplicationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListApplicationsPagesWithContext(ctx context.Context, input *codedeploy.ListApplicationsInput, fn func(*codedeploy.ListApplicationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListApplicationsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListApplicationsOutput{}
	fnCacher := func(out *codedeploy.ListApplicationsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListApplicationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListApplicationsWithContext(ctx context.Context, input *codedeploy.ListApplicationsInput, opts ...request.Option) (*codedeploy.ListApplicationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListApplicationsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListApplicationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentConfigs(input *codedeploy.ListDeploymentConfigsInput) (*codedeploy.ListDeploymentConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentConfigsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentConfigs(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentConfigsPages(input *codedeploy.ListDeploymentConfigsInput, fn func(*codedeploy.ListDeploymentConfigsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentConfigsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentConfigsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentConfigsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentConfigsPagesWithContext(ctx context.Context, input *codedeploy.ListDeploymentConfigsInput, fn func(*codedeploy.ListDeploymentConfigsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentConfigsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentConfigsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentConfigsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentConfigsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentConfigsWithContext(ctx context.Context, input *codedeploy.ListDeploymentConfigsInput, opts ...request.Option) (*codedeploy.ListDeploymentConfigsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentConfigsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentConfigsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentGroups(input *codedeploy.ListDeploymentGroupsInput) (*codedeploy.ListDeploymentGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentGroupsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentGroupsPages(input *codedeploy.ListDeploymentGroupsInput, fn func(*codedeploy.ListDeploymentGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentGroupsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentGroupsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentGroupsPagesWithContext(ctx context.Context, input *codedeploy.ListDeploymentGroupsInput, fn func(*codedeploy.ListDeploymentGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentGroupsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentGroupsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentGroupsWithContext(ctx context.Context, input *codedeploy.ListDeploymentGroupsInput, opts ...request.Option) (*codedeploy.ListDeploymentGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentGroupsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentInstances(input *codedeploy.ListDeploymentInstancesInput) (*codedeploy.ListDeploymentInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentInstancesPages(input *codedeploy.ListDeploymentInstancesInput, fn func(*codedeploy.ListDeploymentInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentInstancesOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentInstancesOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentInstancesPagesWithContext(ctx context.Context, input *codedeploy.ListDeploymentInstancesInput, fn func(*codedeploy.ListDeploymentInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentInstancesOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentInstancesOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentInstancesWithContext(ctx context.Context, input *codedeploy.ListDeploymentInstancesInput, opts ...request.Option) (*codedeploy.ListDeploymentInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentTargets(input *codedeploy.ListDeploymentTargetsInput) (*codedeploy.ListDeploymentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentTargetsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentTargets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentTargetsWithContext(ctx context.Context, input *codedeploy.ListDeploymentTargetsInput, opts ...request.Option) (*codedeploy.ListDeploymentTargetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentTargetsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentTargetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeployments(input *codedeploy.ListDeploymentsInput) (*codedeploy.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListDeploymentsPages(input *codedeploy.ListDeploymentsInput, fn func(*codedeploy.ListDeploymentsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentsPagesWithContext(ctx context.Context, input *codedeploy.ListDeploymentsInput, fn func(*codedeploy.ListDeploymentsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*codedeploy.ListDeploymentsOutput), false)
		return nil
	}
	cachable := true
	output := &codedeploy.ListDeploymentsOutput{}
	fnCacher := func(out *codedeploy.ListDeploymentsOutput, more bool) bool {
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
	if err := c.CodeDeployAPI.ListDeploymentsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CodeDeploy) ListDeploymentsWithContext(ctx context.Context, input *codedeploy.ListDeploymentsInput, opts ...request.Option) (*codedeploy.ListDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListDeploymentsOutput), nil
	}
	out, err := c.CodeDeployAPI.ListDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListGitHubAccountTokenNames(input *codedeploy.ListGitHubAccountTokenNamesInput) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListGitHubAccountTokenNamesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListGitHubAccountTokenNames(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListGitHubAccountTokenNamesWithContext(ctx context.Context, input *codedeploy.ListGitHubAccountTokenNamesInput, opts ...request.Option) (*codedeploy.ListGitHubAccountTokenNamesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListGitHubAccountTokenNamesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListGitHubAccountTokenNamesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListOnPremisesInstances(input *codedeploy.ListOnPremisesInstancesInput) (*codedeploy.ListOnPremisesInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListOnPremisesInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListOnPremisesInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListOnPremisesInstancesWithContext(ctx context.Context, input *codedeploy.ListOnPremisesInstancesInput, opts ...request.Option) (*codedeploy.ListOnPremisesInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListOnPremisesInstancesOutput), nil
	}
	out, err := c.CodeDeployAPI.ListOnPremisesInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListTagsForResource(input *codedeploy.ListTagsForResourceInput) (*codedeploy.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeDeployAPI.ListTagsForResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CodeDeploy) ListTagsForResourceWithContext(ctx context.Context, input *codedeploy.ListTagsForResourceInput, opts ...request.Option) (*codedeploy.ListTagsForResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*codedeploy.ListTagsForResourceOutput), nil
	}
	out, err := c.CodeDeployAPI.ListTagsForResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
