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
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/aws/aws-sdk-go/service/opsworks/opsworksiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type OpsWorks struct {
	opsworksiface.OpsWorksAPI
	cache *cache.Cache
}

func NewOpsWorks(opsworksapi opsworksiface.OpsWorksAPI) *OpsWorks {
	return &OpsWorks{
		OpsWorksAPI: opsworksapi,
		cache:       cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *OpsWorks) DescribeAgentVersions(input *opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeAgentVersionsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeAgentVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeAgentVersionsWithContext(ctx context.Context, input *opsworks.DescribeAgentVersionsInput, opts ...request.Option) (*opsworks.DescribeAgentVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeAgentVersionsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeAgentVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeApps(input *opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeAppsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeApps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeAppsWithContext(ctx context.Context, input *opsworks.DescribeAppsInput, opts ...request.Option) (*opsworks.DescribeAppsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeAppsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeAppsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeCommands(input *opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeCommandsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeCommands(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeCommandsWithContext(ctx context.Context, input *opsworks.DescribeCommandsInput, opts ...request.Option) (*opsworks.DescribeCommandsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeCommandsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeCommandsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeDeployments(input *opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeDeploymentsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeDeployments(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeDeploymentsWithContext(ctx context.Context, input *opsworks.DescribeDeploymentsInput, opts ...request.Option) (*opsworks.DescribeDeploymentsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeDeploymentsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeDeploymentsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeEcsClusters(input *opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeEcsClustersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeEcsClusters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeEcsClustersPages(input *opsworks.DescribeEcsClustersInput, fn func(*opsworks.DescribeEcsClustersOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworks.DescribeEcsClustersOutput), false)
		return nil
	}
	cachable := true
	output := &opsworks.DescribeEcsClustersOutput{}
	fnCacher := func(out *opsworks.DescribeEcsClustersOutput, more bool) bool {
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
	if err := c.OpsWorksAPI.DescribeEcsClustersPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorks) DescribeEcsClustersPagesWithContext(ctx context.Context, input *opsworks.DescribeEcsClustersInput, fn func(*opsworks.DescribeEcsClustersOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*opsworks.DescribeEcsClustersOutput), false)
		return nil
	}
	cachable := true
	output := &opsworks.DescribeEcsClustersOutput{}
	fnCacher := func(out *opsworks.DescribeEcsClustersOutput, more bool) bool {
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
	if err := c.OpsWorksAPI.DescribeEcsClustersPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *OpsWorks) DescribeEcsClustersWithContext(ctx context.Context, input *opsworks.DescribeEcsClustersInput, opts ...request.Option) (*opsworks.DescribeEcsClustersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeEcsClustersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeEcsClustersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeElasticIps(input *opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeElasticIpsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeElasticIps(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeElasticIpsWithContext(ctx context.Context, input *opsworks.DescribeElasticIpsInput, opts ...request.Option) (*opsworks.DescribeElasticIpsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeElasticIpsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeElasticIpsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeElasticLoadBalancers(input *opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeElasticLoadBalancersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeElasticLoadBalancers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeElasticLoadBalancersWithContext(ctx context.Context, input *opsworks.DescribeElasticLoadBalancersInput, opts ...request.Option) (*opsworks.DescribeElasticLoadBalancersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeElasticLoadBalancersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeElasticLoadBalancersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeInstances(input *opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeInstancesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeInstancesWithContext(ctx context.Context, input *opsworks.DescribeInstancesInput, opts ...request.Option) (*opsworks.DescribeInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeInstancesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeLayers(input *opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeLayersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeLayers(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeLayersWithContext(ctx context.Context, input *opsworks.DescribeLayersInput, opts ...request.Option) (*opsworks.DescribeLayersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeLayersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeLayersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeLoadBasedAutoScaling(input *opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeLoadBasedAutoScalingOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeLoadBasedAutoScaling(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeLoadBasedAutoScalingWithContext(ctx context.Context, input *opsworks.DescribeLoadBasedAutoScalingInput, opts ...request.Option) (*opsworks.DescribeLoadBasedAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeLoadBasedAutoScalingOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeLoadBasedAutoScalingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeMyUserProfile(input *opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeMyUserProfileOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeMyUserProfile(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeMyUserProfileWithContext(ctx context.Context, input *opsworks.DescribeMyUserProfileInput, opts ...request.Option) (*opsworks.DescribeMyUserProfileOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeMyUserProfileOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeMyUserProfileWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeOperatingSystems(input *opsworks.DescribeOperatingSystemsInput) (*opsworks.DescribeOperatingSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeOperatingSystemsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeOperatingSystems(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeOperatingSystemsWithContext(ctx context.Context, input *opsworks.DescribeOperatingSystemsInput, opts ...request.Option) (*opsworks.DescribeOperatingSystemsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeOperatingSystemsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeOperatingSystemsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribePermissions(input *opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribePermissionsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribePermissions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribePermissionsWithContext(ctx context.Context, input *opsworks.DescribePermissionsInput, opts ...request.Option) (*opsworks.DescribePermissionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribePermissionsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribePermissionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeRaidArrays(input *opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeRaidArraysOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeRaidArrays(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeRaidArraysWithContext(ctx context.Context, input *opsworks.DescribeRaidArraysInput, opts ...request.Option) (*opsworks.DescribeRaidArraysOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeRaidArraysOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeRaidArraysWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeRdsDbInstances(input *opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeRdsDbInstancesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeRdsDbInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeRdsDbInstancesWithContext(ctx context.Context, input *opsworks.DescribeRdsDbInstancesInput, opts ...request.Option) (*opsworks.DescribeRdsDbInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeRdsDbInstancesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeRdsDbInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeServiceErrors(input *opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeServiceErrorsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeServiceErrors(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeServiceErrorsWithContext(ctx context.Context, input *opsworks.DescribeServiceErrorsInput, opts ...request.Option) (*opsworks.DescribeServiceErrorsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeServiceErrorsOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeServiceErrorsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStackProvisioningParameters(input *opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStackProvisioningParametersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStackProvisioningParameters(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStackProvisioningParametersWithContext(ctx context.Context, input *opsworks.DescribeStackProvisioningParametersInput, opts ...request.Option) (*opsworks.DescribeStackProvisioningParametersOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStackProvisioningParametersOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStackProvisioningParametersWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStackSummary(input *opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStackSummaryOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStackSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStackSummaryWithContext(ctx context.Context, input *opsworks.DescribeStackSummaryInput, opts ...request.Option) (*opsworks.DescribeStackSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStackSummaryOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStackSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStacks(input *opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStacksOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeStacksWithContext(ctx context.Context, input *opsworks.DescribeStacksInput, opts ...request.Option) (*opsworks.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeStacksOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeTimeBasedAutoScaling(input *opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeTimeBasedAutoScalingOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeTimeBasedAutoScaling(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeTimeBasedAutoScalingWithContext(ctx context.Context, input *opsworks.DescribeTimeBasedAutoScalingInput, opts ...request.Option) (*opsworks.DescribeTimeBasedAutoScalingOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeTimeBasedAutoScalingOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeTimeBasedAutoScalingWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeUserProfiles(input *opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeUserProfilesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeUserProfiles(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeUserProfilesWithContext(ctx context.Context, input *opsworks.DescribeUserProfilesInput, opts ...request.Option) (*opsworks.DescribeUserProfilesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeUserProfilesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeUserProfilesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeVolumes(input *opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeVolumesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeVolumes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) DescribeVolumesWithContext(ctx context.Context, input *opsworks.DescribeVolumesInput, opts ...request.Option) (*opsworks.DescribeVolumesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.DescribeVolumesOutput), nil
	}
	out, err := c.OpsWorksAPI.DescribeVolumesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) GetHostnameSuggestion(input *opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.GetHostnameSuggestionOutput), nil
	}
	out, err := c.OpsWorksAPI.GetHostnameSuggestion(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) GetHostnameSuggestionWithContext(ctx context.Context, input *opsworks.GetHostnameSuggestionInput, opts ...request.Option) (*opsworks.GetHostnameSuggestionOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.GetHostnameSuggestionOutput), nil
	}
	out, err := c.OpsWorksAPI.GetHostnameSuggestionWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) ListTags(input *opsworks.ListTagsInput) (*opsworks.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.ListTagsOutput), nil
	}
	out, err := c.OpsWorksAPI.ListTags(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *OpsWorks) ListTagsWithContext(ctx context.Context, input *opsworks.ListTagsInput, opts ...request.Option) (*opsworks.ListTagsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*opsworks.ListTagsOutput), nil
	}
	out, err := c.OpsWorksAPI.ListTagsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
