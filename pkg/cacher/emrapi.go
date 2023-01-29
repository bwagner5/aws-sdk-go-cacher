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
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/aws/aws-sdk-go/service/emr/emriface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type EMR struct {
	emriface.EMRAPI
	cache *cache.Cache
}

func NewEMR(emrapi emriface.EMRAPI) *EMR {
	return &EMR{
		EMRAPI: emrapi,
		cache:  cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *EMR) DescribeCluster(input *emr.DescribeClusterInput) (*emr.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeClusterOutput), nil
	}
	out, err := c.EMRAPI.DescribeCluster(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeClusterWithContext(ctx context.Context, input *emr.DescribeClusterInput, opts ...request.Option) (*emr.DescribeClusterOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeClusterOutput), nil
	}
	out, err := c.EMRAPI.DescribeClusterWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeJobFlows(input *emr.DescribeJobFlowsInput) (*emr.DescribeJobFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeJobFlowsOutput), nil
	}
	out, err := c.EMRAPI.DescribeJobFlows(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeJobFlowsWithContext(ctx context.Context, input *emr.DescribeJobFlowsInput, opts ...request.Option) (*emr.DescribeJobFlowsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeJobFlowsOutput), nil
	}
	out, err := c.EMRAPI.DescribeJobFlowsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeNotebookExecution(input *emr.DescribeNotebookExecutionInput) (*emr.DescribeNotebookExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeNotebookExecutionOutput), nil
	}
	out, err := c.EMRAPI.DescribeNotebookExecution(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeNotebookExecutionWithContext(ctx context.Context, input *emr.DescribeNotebookExecutionInput, opts ...request.Option) (*emr.DescribeNotebookExecutionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeNotebookExecutionOutput), nil
	}
	out, err := c.EMRAPI.DescribeNotebookExecutionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeReleaseLabel(input *emr.DescribeReleaseLabelInput) (*emr.DescribeReleaseLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeReleaseLabelOutput), nil
	}
	out, err := c.EMRAPI.DescribeReleaseLabel(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeReleaseLabelWithContext(ctx context.Context, input *emr.DescribeReleaseLabelInput, opts ...request.Option) (*emr.DescribeReleaseLabelOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeReleaseLabelOutput), nil
	}
	out, err := c.EMRAPI.DescribeReleaseLabelWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeSecurityConfiguration(input *emr.DescribeSecurityConfigurationInput) (*emr.DescribeSecurityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeSecurityConfigurationOutput), nil
	}
	out, err := c.EMRAPI.DescribeSecurityConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeSecurityConfigurationWithContext(ctx context.Context, input *emr.DescribeSecurityConfigurationInput, opts ...request.Option) (*emr.DescribeSecurityConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeSecurityConfigurationOutput), nil
	}
	out, err := c.EMRAPI.DescribeSecurityConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeStep(input *emr.DescribeStepInput) (*emr.DescribeStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeStepOutput), nil
	}
	out, err := c.EMRAPI.DescribeStep(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeStepWithContext(ctx context.Context, input *emr.DescribeStepInput, opts ...request.Option) (*emr.DescribeStepOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeStepOutput), nil
	}
	out, err := c.EMRAPI.DescribeStepWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeStudio(input *emr.DescribeStudioInput) (*emr.DescribeStudioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeStudioOutput), nil
	}
	out, err := c.EMRAPI.DescribeStudio(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) DescribeStudioWithContext(ctx context.Context, input *emr.DescribeStudioInput, opts ...request.Option) (*emr.DescribeStudioOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.DescribeStudioOutput), nil
	}
	out, err := c.EMRAPI.DescribeStudioWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetAutoTerminationPolicy(input *emr.GetAutoTerminationPolicyInput) (*emr.GetAutoTerminationPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetAutoTerminationPolicyOutput), nil
	}
	out, err := c.EMRAPI.GetAutoTerminationPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetAutoTerminationPolicyWithContext(ctx context.Context, input *emr.GetAutoTerminationPolicyInput, opts ...request.Option) (*emr.GetAutoTerminationPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetAutoTerminationPolicyOutput), nil
	}
	out, err := c.EMRAPI.GetAutoTerminationPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetBlockPublicAccessConfiguration(input *emr.GetBlockPublicAccessConfigurationInput) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetBlockPublicAccessConfigurationOutput), nil
	}
	out, err := c.EMRAPI.GetBlockPublicAccessConfiguration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetBlockPublicAccessConfigurationWithContext(ctx context.Context, input *emr.GetBlockPublicAccessConfigurationInput, opts ...request.Option) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetBlockPublicAccessConfigurationOutput), nil
	}
	out, err := c.EMRAPI.GetBlockPublicAccessConfigurationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetClusterSessionCredentials(input *emr.GetClusterSessionCredentialsInput) (*emr.GetClusterSessionCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetClusterSessionCredentialsOutput), nil
	}
	out, err := c.EMRAPI.GetClusterSessionCredentials(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetClusterSessionCredentialsWithContext(ctx context.Context, input *emr.GetClusterSessionCredentialsInput, opts ...request.Option) (*emr.GetClusterSessionCredentialsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetClusterSessionCredentialsOutput), nil
	}
	out, err := c.EMRAPI.GetClusterSessionCredentialsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetManagedScalingPolicy(input *emr.GetManagedScalingPolicyInput) (*emr.GetManagedScalingPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetManagedScalingPolicyOutput), nil
	}
	out, err := c.EMRAPI.GetManagedScalingPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetManagedScalingPolicyWithContext(ctx context.Context, input *emr.GetManagedScalingPolicyInput, opts ...request.Option) (*emr.GetManagedScalingPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetManagedScalingPolicyOutput), nil
	}
	out, err := c.EMRAPI.GetManagedScalingPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetStudioSessionMapping(input *emr.GetStudioSessionMappingInput) (*emr.GetStudioSessionMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetStudioSessionMappingOutput), nil
	}
	out, err := c.EMRAPI.GetStudioSessionMapping(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) GetStudioSessionMappingWithContext(ctx context.Context, input *emr.GetStudioSessionMappingInput, opts ...request.Option) (*emr.GetStudioSessionMappingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.GetStudioSessionMappingOutput), nil
	}
	out, err := c.EMRAPI.GetStudioSessionMappingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListBootstrapActions(input *emr.ListBootstrapActionsInput) (*emr.ListBootstrapActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListBootstrapActionsOutput), nil
	}
	out, err := c.EMRAPI.ListBootstrapActions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListBootstrapActionsPages(input *emr.ListBootstrapActionsInput, fn func(*emr.ListBootstrapActionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListBootstrapActionsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListBootstrapActionsOutput{}
	fnCacher := func(out *emr.ListBootstrapActionsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListBootstrapActionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListBootstrapActionsPagesWithContext(ctx context.Context, input *emr.ListBootstrapActionsInput, fn func(*emr.ListBootstrapActionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListBootstrapActionsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListBootstrapActionsOutput{}
	fnCacher := func(out *emr.ListBootstrapActionsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListBootstrapActionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListBootstrapActionsWithContext(ctx context.Context, input *emr.ListBootstrapActionsInput, opts ...request.Option) (*emr.ListBootstrapActionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListBootstrapActionsOutput), nil
	}
	out, err := c.EMRAPI.ListBootstrapActionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListClusters(input *emr.ListClustersInput) (*emr.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListClustersOutput), nil
	}
	out, err := c.EMRAPI.ListClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListClustersPages(input *emr.ListClustersInput, fn func(*emr.ListClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListClustersOutput{}
	fnCacher := func(out *emr.ListClustersOutput, more bool) bool {
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
	if err := c.EMRAPI.ListClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListClustersPagesWithContext(ctx context.Context, input *emr.ListClustersInput, fn func(*emr.ListClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListClustersOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListClustersOutput{}
	fnCacher := func(out *emr.ListClustersOutput, more bool) bool {
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
	if err := c.EMRAPI.ListClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListClustersWithContext(ctx context.Context, input *emr.ListClustersInput, opts ...request.Option) (*emr.ListClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListClustersOutput), nil
	}
	out, err := c.EMRAPI.ListClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstanceFleets(input *emr.ListInstanceFleetsInput) (*emr.ListInstanceFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstanceFleetsOutput), nil
	}
	out, err := c.EMRAPI.ListInstanceFleets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstanceFleetsPages(input *emr.ListInstanceFleetsInput, fn func(*emr.ListInstanceFleetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstanceFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstanceFleetsOutput{}
	fnCacher := func(out *emr.ListInstanceFleetsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstanceFleetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstanceFleetsPagesWithContext(ctx context.Context, input *emr.ListInstanceFleetsInput, fn func(*emr.ListInstanceFleetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstanceFleetsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstanceFleetsOutput{}
	fnCacher := func(out *emr.ListInstanceFleetsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstanceFleetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstanceFleetsWithContext(ctx context.Context, input *emr.ListInstanceFleetsInput, opts ...request.Option) (*emr.ListInstanceFleetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstanceFleetsOutput), nil
	}
	out, err := c.EMRAPI.ListInstanceFleetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstanceGroups(input *emr.ListInstanceGroupsInput) (*emr.ListInstanceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstanceGroupsOutput), nil
	}
	out, err := c.EMRAPI.ListInstanceGroups(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstanceGroupsPages(input *emr.ListInstanceGroupsInput, fn func(*emr.ListInstanceGroupsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstanceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstanceGroupsOutput{}
	fnCacher := func(out *emr.ListInstanceGroupsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstanceGroupsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstanceGroupsPagesWithContext(ctx context.Context, input *emr.ListInstanceGroupsInput, fn func(*emr.ListInstanceGroupsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstanceGroupsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstanceGroupsOutput{}
	fnCacher := func(out *emr.ListInstanceGroupsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstanceGroupsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstanceGroupsWithContext(ctx context.Context, input *emr.ListInstanceGroupsInput, opts ...request.Option) (*emr.ListInstanceGroupsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstanceGroupsOutput), nil
	}
	out, err := c.EMRAPI.ListInstanceGroupsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstances(input *emr.ListInstancesInput) (*emr.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstancesOutput), nil
	}
	out, err := c.EMRAPI.ListInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListInstancesPages(input *emr.ListInstancesInput, fn func(*emr.ListInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstancesOutput{}
	fnCacher := func(out *emr.ListInstancesOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstancesPagesWithContext(ctx context.Context, input *emr.ListInstancesInput, fn func(*emr.ListInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListInstancesOutput{}
	fnCacher := func(out *emr.ListInstancesOutput, more bool) bool {
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
	if err := c.EMRAPI.ListInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListInstancesWithContext(ctx context.Context, input *emr.ListInstancesInput, opts ...request.Option) (*emr.ListInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListInstancesOutput), nil
	}
	out, err := c.EMRAPI.ListInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListNotebookExecutions(input *emr.ListNotebookExecutionsInput) (*emr.ListNotebookExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListNotebookExecutionsOutput), nil
	}
	out, err := c.EMRAPI.ListNotebookExecutions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListNotebookExecutionsPages(input *emr.ListNotebookExecutionsInput, fn func(*emr.ListNotebookExecutionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListNotebookExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListNotebookExecutionsOutput{}
	fnCacher := func(out *emr.ListNotebookExecutionsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListNotebookExecutionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListNotebookExecutionsPagesWithContext(ctx context.Context, input *emr.ListNotebookExecutionsInput, fn func(*emr.ListNotebookExecutionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListNotebookExecutionsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListNotebookExecutionsOutput{}
	fnCacher := func(out *emr.ListNotebookExecutionsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListNotebookExecutionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListNotebookExecutionsWithContext(ctx context.Context, input *emr.ListNotebookExecutionsInput, opts ...request.Option) (*emr.ListNotebookExecutionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListNotebookExecutionsOutput), nil
	}
	out, err := c.EMRAPI.ListNotebookExecutionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListReleaseLabels(input *emr.ListReleaseLabelsInput) (*emr.ListReleaseLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListReleaseLabelsOutput), nil
	}
	out, err := c.EMRAPI.ListReleaseLabels(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListReleaseLabelsPages(input *emr.ListReleaseLabelsInput, fn func(*emr.ListReleaseLabelsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListReleaseLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListReleaseLabelsOutput{}
	fnCacher := func(out *emr.ListReleaseLabelsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListReleaseLabelsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListReleaseLabelsPagesWithContext(ctx context.Context, input *emr.ListReleaseLabelsInput, fn func(*emr.ListReleaseLabelsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListReleaseLabelsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListReleaseLabelsOutput{}
	fnCacher := func(out *emr.ListReleaseLabelsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListReleaseLabelsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListReleaseLabelsWithContext(ctx context.Context, input *emr.ListReleaseLabelsInput, opts ...request.Option) (*emr.ListReleaseLabelsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListReleaseLabelsOutput), nil
	}
	out, err := c.EMRAPI.ListReleaseLabelsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListSecurityConfigurations(input *emr.ListSecurityConfigurationsInput) (*emr.ListSecurityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListSecurityConfigurationsOutput), nil
	}
	out, err := c.EMRAPI.ListSecurityConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListSecurityConfigurationsPages(input *emr.ListSecurityConfigurationsInput, fn func(*emr.ListSecurityConfigurationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListSecurityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListSecurityConfigurationsOutput{}
	fnCacher := func(out *emr.ListSecurityConfigurationsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListSecurityConfigurationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListSecurityConfigurationsPagesWithContext(ctx context.Context, input *emr.ListSecurityConfigurationsInput, fn func(*emr.ListSecurityConfigurationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListSecurityConfigurationsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListSecurityConfigurationsOutput{}
	fnCacher := func(out *emr.ListSecurityConfigurationsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListSecurityConfigurationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListSecurityConfigurationsWithContext(ctx context.Context, input *emr.ListSecurityConfigurationsInput, opts ...request.Option) (*emr.ListSecurityConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListSecurityConfigurationsOutput), nil
	}
	out, err := c.EMRAPI.ListSecurityConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListSteps(input *emr.ListStepsInput) (*emr.ListStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStepsOutput), nil
	}
	out, err := c.EMRAPI.ListSteps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListStepsPages(input *emr.ListStepsInput, fn func(*emr.ListStepsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStepsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStepsOutput{}
	fnCacher := func(out *emr.ListStepsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStepsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStepsPagesWithContext(ctx context.Context, input *emr.ListStepsInput, fn func(*emr.ListStepsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStepsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStepsOutput{}
	fnCacher := func(out *emr.ListStepsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStepsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStepsWithContext(ctx context.Context, input *emr.ListStepsInput, opts ...request.Option) (*emr.ListStepsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStepsOutput), nil
	}
	out, err := c.EMRAPI.ListStepsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListStudioSessionMappings(input *emr.ListStudioSessionMappingsInput) (*emr.ListStudioSessionMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStudioSessionMappingsOutput), nil
	}
	out, err := c.EMRAPI.ListStudioSessionMappings(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListStudioSessionMappingsPages(input *emr.ListStudioSessionMappingsInput, fn func(*emr.ListStudioSessionMappingsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStudioSessionMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStudioSessionMappingsOutput{}
	fnCacher := func(out *emr.ListStudioSessionMappingsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStudioSessionMappingsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStudioSessionMappingsPagesWithContext(ctx context.Context, input *emr.ListStudioSessionMappingsInput, fn func(*emr.ListStudioSessionMappingsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStudioSessionMappingsOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStudioSessionMappingsOutput{}
	fnCacher := func(out *emr.ListStudioSessionMappingsOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStudioSessionMappingsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStudioSessionMappingsWithContext(ctx context.Context, input *emr.ListStudioSessionMappingsInput, opts ...request.Option) (*emr.ListStudioSessionMappingsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStudioSessionMappingsOutput), nil
	}
	out, err := c.EMRAPI.ListStudioSessionMappingsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListStudios(input *emr.ListStudiosInput) (*emr.ListStudiosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStudiosOutput), nil
	}
	out, err := c.EMRAPI.ListStudios(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *EMR) ListStudiosPages(input *emr.ListStudiosInput, fn func(*emr.ListStudiosOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStudiosOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStudiosOutput{}
	fnCacher := func(out *emr.ListStudiosOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStudiosPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStudiosPagesWithContext(ctx context.Context, input *emr.ListStudiosInput, fn func(*emr.ListStudiosOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*emr.ListStudiosOutput), false)
		return nil
	}
	cachable := true
	output := &emr.ListStudiosOutput{}
	fnCacher := func(out *emr.ListStudiosOutput, more bool) bool {
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
	if err := c.EMRAPI.ListStudiosPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *EMR) ListStudiosWithContext(ctx context.Context, input *emr.ListStudiosInput, opts ...request.Option) (*emr.ListStudiosOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*emr.ListStudiosOutput), nil
	}
	out, err := c.EMRAPI.ListStudiosWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
