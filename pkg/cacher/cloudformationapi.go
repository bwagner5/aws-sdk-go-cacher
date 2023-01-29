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
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"github.com/imdario/mergo"
	"github.com/mitchellh/hashstructure/v2"
	"github.com/patrickmn/go-cache"
)

type CloudFormation struct {
	cloudformationiface.CloudFormationAPI
	cache *cache.Cache
}

func NewCloudFormation(cloudformationapi cloudformationiface.CloudFormationAPI) *CloudFormation {
	return &CloudFormation{
		CloudFormationAPI: cloudformationapi,
		cache:             cache.New(1*time.Minute, 2*time.Minute),
	}
}
func (c *CloudFormation) BatchDescribeTypeConfigurations(input *cloudformation.BatchDescribeTypeConfigurationsInput) (*cloudformation.BatchDescribeTypeConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.BatchDescribeTypeConfigurationsOutput), nil
	}
	out, err := c.CloudFormationAPI.BatchDescribeTypeConfigurations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) BatchDescribeTypeConfigurationsWithContext(ctx context.Context, input *cloudformation.BatchDescribeTypeConfigurationsInput, opts ...request.Option) (*cloudformation.BatchDescribeTypeConfigurationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.BatchDescribeTypeConfigurationsOutput), nil
	}
	out, err := c.CloudFormationAPI.BatchDescribeTypeConfigurationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeAccountLimits(input *cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeAccountLimitsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeAccountLimits(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeAccountLimitsPages(input *cloudformation.DescribeAccountLimitsInput, fn func(*cloudformation.DescribeAccountLimitsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeAccountLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeAccountLimitsOutput{}
	fnCacher := func(out *cloudformation.DescribeAccountLimitsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeAccountLimitsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeAccountLimitsPagesWithContext(ctx context.Context, input *cloudformation.DescribeAccountLimitsInput, fn func(*cloudformation.DescribeAccountLimitsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeAccountLimitsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeAccountLimitsOutput{}
	fnCacher := func(out *cloudformation.DescribeAccountLimitsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeAccountLimitsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeAccountLimitsWithContext(ctx context.Context, input *cloudformation.DescribeAccountLimitsInput, opts ...request.Option) (*cloudformation.DescribeAccountLimitsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeAccountLimitsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeAccountLimitsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeChangeSet(input *cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeChangeSetOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeChangeSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeChangeSetHooks(input *cloudformation.DescribeChangeSetHooksInput) (*cloudformation.DescribeChangeSetHooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeChangeSetHooksOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeChangeSetHooks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeChangeSetHooksWithContext(ctx context.Context, input *cloudformation.DescribeChangeSetHooksInput, opts ...request.Option) (*cloudformation.DescribeChangeSetHooksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeChangeSetHooksOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeChangeSetHooksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeChangeSetWithContext(ctx context.Context, input *cloudformation.DescribeChangeSetInput, opts ...request.Option) (*cloudformation.DescribeChangeSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeChangeSetOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeChangeSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribePublisher(input *cloudformation.DescribePublisherInput) (*cloudformation.DescribePublisherOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribePublisherOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribePublisher(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribePublisherWithContext(ctx context.Context, input *cloudformation.DescribePublisherInput, opts ...request.Option) (*cloudformation.DescribePublisherOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribePublisherOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribePublisherWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackDriftDetectionStatus(input *cloudformation.DescribeStackDriftDetectionStatusInput) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackDriftDetectionStatusOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackDriftDetectionStatus(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackDriftDetectionStatusWithContext(ctx context.Context, input *cloudformation.DescribeStackDriftDetectionStatusInput, opts ...request.Option) (*cloudformation.DescribeStackDriftDetectionStatusOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackDriftDetectionStatusOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackDriftDetectionStatusWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackEvents(input *cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackEventsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackEvents(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackEventsPages(input *cloudformation.DescribeStackEventsInput, fn func(*cloudformation.DescribeStackEventsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStackEventsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStackEventsOutput{}
	fnCacher := func(out *cloudformation.DescribeStackEventsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStackEventsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStackEventsPagesWithContext(ctx context.Context, input *cloudformation.DescribeStackEventsInput, fn func(*cloudformation.DescribeStackEventsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStackEventsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStackEventsOutput{}
	fnCacher := func(out *cloudformation.DescribeStackEventsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStackEventsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStackEventsWithContext(ctx context.Context, input *cloudformation.DescribeStackEventsInput, opts ...request.Option) (*cloudformation.DescribeStackEventsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackEventsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackEventsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackInstance(input *cloudformation.DescribeStackInstanceInput) (*cloudformation.DescribeStackInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackInstanceOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackInstance(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackInstanceWithContext(ctx context.Context, input *cloudformation.DescribeStackInstanceInput, opts ...request.Option) (*cloudformation.DescribeStackInstanceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackInstanceOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackInstanceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResource(input *cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourceOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResource(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResourceDrifts(input *cloudformation.DescribeStackResourceDriftsInput) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourceDriftsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResourceDrifts(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResourceDriftsPages(input *cloudformation.DescribeStackResourceDriftsInput, fn func(*cloudformation.DescribeStackResourceDriftsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStackResourceDriftsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStackResourceDriftsOutput{}
	fnCacher := func(out *cloudformation.DescribeStackResourceDriftsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStackResourceDriftsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStackResourceDriftsPagesWithContext(ctx context.Context, input *cloudformation.DescribeStackResourceDriftsInput, fn func(*cloudformation.DescribeStackResourceDriftsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStackResourceDriftsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStackResourceDriftsOutput{}
	fnCacher := func(out *cloudformation.DescribeStackResourceDriftsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStackResourceDriftsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStackResourceDriftsWithContext(ctx context.Context, input *cloudformation.DescribeStackResourceDriftsInput, opts ...request.Option) (*cloudformation.DescribeStackResourceDriftsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourceDriftsOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResourceDriftsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResourceWithContext(ctx context.Context, input *cloudformation.DescribeStackResourceInput, opts ...request.Option) (*cloudformation.DescribeStackResourceOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourceOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResourceWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResources(input *cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourcesOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackResourcesWithContext(ctx context.Context, input *cloudformation.DescribeStackResourcesInput, opts ...request.Option) (*cloudformation.DescribeStackResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackResourcesOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackSet(input *cloudformation.DescribeStackSetInput) (*cloudformation.DescribeStackSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackSetOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackSet(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackSetOperation(input *cloudformation.DescribeStackSetOperationInput) (*cloudformation.DescribeStackSetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackSetOperationOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackSetOperation(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackSetOperationWithContext(ctx context.Context, input *cloudformation.DescribeStackSetOperationInput, opts ...request.Option) (*cloudformation.DescribeStackSetOperationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackSetOperationOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackSetOperationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStackSetWithContext(ctx context.Context, input *cloudformation.DescribeStackSetInput, opts ...request.Option) (*cloudformation.DescribeStackSetOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStackSetOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStackSetWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStacksOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeStacksPages(input *cloudformation.DescribeStacksInput, fn func(*cloudformation.DescribeStacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStacksOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStacksOutput{}
	fnCacher := func(out *cloudformation.DescribeStacksOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStacksPagesWithContext(ctx context.Context, input *cloudformation.DescribeStacksInput, fn func(*cloudformation.DescribeStacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.DescribeStacksOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.DescribeStacksOutput{}
	fnCacher := func(out *cloudformation.DescribeStacksOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.DescribeStacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) DescribeStacksWithContext(ctx context.Context, input *cloudformation.DescribeStacksInput, opts ...request.Option) (*cloudformation.DescribeStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeStacksOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeType(input *cloudformation.DescribeTypeInput) (*cloudformation.DescribeTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeTypeOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeType(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeTypeRegistration(input *cloudformation.DescribeTypeRegistrationInput) (*cloudformation.DescribeTypeRegistrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeTypeRegistrationOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeTypeRegistration(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeTypeRegistrationWithContext(ctx context.Context, input *cloudformation.DescribeTypeRegistrationInput, opts ...request.Option) (*cloudformation.DescribeTypeRegistrationOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeTypeRegistrationOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeTypeRegistrationWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) DescribeTypeWithContext(ctx context.Context, input *cloudformation.DescribeTypeInput, opts ...request.Option) (*cloudformation.DescribeTypeOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.DescribeTypeOutput), nil
	}
	out, err := c.CloudFormationAPI.DescribeTypeWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetStackPolicy(input *cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetStackPolicyOutput), nil
	}
	out, err := c.CloudFormationAPI.GetStackPolicy(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetStackPolicyWithContext(ctx context.Context, input *cloudformation.GetStackPolicyInput, opts ...request.Option) (*cloudformation.GetStackPolicyOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetStackPolicyOutput), nil
	}
	out, err := c.CloudFormationAPI.GetStackPolicyWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetTemplate(input *cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetTemplateOutput), nil
	}
	out, err := c.CloudFormationAPI.GetTemplate(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetTemplateSummary(input *cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetTemplateSummaryOutput), nil
	}
	out, err := c.CloudFormationAPI.GetTemplateSummary(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetTemplateSummaryWithContext(ctx context.Context, input *cloudformation.GetTemplateSummaryInput, opts ...request.Option) (*cloudformation.GetTemplateSummaryOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetTemplateSummaryOutput), nil
	}
	out, err := c.CloudFormationAPI.GetTemplateSummaryWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) GetTemplateWithContext(ctx context.Context, input *cloudformation.GetTemplateInput, opts ...request.Option) (*cloudformation.GetTemplateOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.GetTemplateOutput), nil
	}
	out, err := c.CloudFormationAPI.GetTemplateWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListChangeSets(input *cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListChangeSetsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListChangeSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListChangeSetsPages(input *cloudformation.ListChangeSetsInput, fn func(*cloudformation.ListChangeSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListChangeSetsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListChangeSetsOutput{}
	fnCacher := func(out *cloudformation.ListChangeSetsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListChangeSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListChangeSetsPagesWithContext(ctx context.Context, input *cloudformation.ListChangeSetsInput, fn func(*cloudformation.ListChangeSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListChangeSetsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListChangeSetsOutput{}
	fnCacher := func(out *cloudformation.ListChangeSetsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListChangeSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListChangeSetsWithContext(ctx context.Context, input *cloudformation.ListChangeSetsInput, opts ...request.Option) (*cloudformation.ListChangeSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListChangeSetsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListChangeSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListExports(input *cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListExportsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListExports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListExportsPages(input *cloudformation.ListExportsInput, fn func(*cloudformation.ListExportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListExportsOutput{}
	fnCacher := func(out *cloudformation.ListExportsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListExportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListExportsPagesWithContext(ctx context.Context, input *cloudformation.ListExportsInput, fn func(*cloudformation.ListExportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListExportsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListExportsOutput{}
	fnCacher := func(out *cloudformation.ListExportsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListExportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListExportsWithContext(ctx context.Context, input *cloudformation.ListExportsInput, opts ...request.Option) (*cloudformation.ListExportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListExportsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListExportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListImports(input *cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListImportsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListImports(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListImportsPages(input *cloudformation.ListImportsInput, fn func(*cloudformation.ListImportsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListImportsOutput{}
	fnCacher := func(out *cloudformation.ListImportsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListImportsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListImportsPagesWithContext(ctx context.Context, input *cloudformation.ListImportsInput, fn func(*cloudformation.ListImportsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListImportsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListImportsOutput{}
	fnCacher := func(out *cloudformation.ListImportsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListImportsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListImportsWithContext(ctx context.Context, input *cloudformation.ListImportsInput, opts ...request.Option) (*cloudformation.ListImportsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListImportsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListImportsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackInstances(input *cloudformation.ListStackInstancesInput) (*cloudformation.ListStackInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackInstancesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackInstances(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackInstancesPages(input *cloudformation.ListStackInstancesInput, fn func(*cloudformation.ListStackInstancesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackInstancesOutput{}
	fnCacher := func(out *cloudformation.ListStackInstancesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackInstancesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackInstancesPagesWithContext(ctx context.Context, input *cloudformation.ListStackInstancesInput, fn func(*cloudformation.ListStackInstancesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackInstancesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackInstancesOutput{}
	fnCacher := func(out *cloudformation.ListStackInstancesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackInstancesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackInstancesWithContext(ctx context.Context, input *cloudformation.ListStackInstancesInput, opts ...request.Option) (*cloudformation.ListStackInstancesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackInstancesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackInstancesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackResources(input *cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackResourcesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackResources(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackResourcesPages(input *cloudformation.ListStackResourcesInput, fn func(*cloudformation.ListStackResourcesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackResourcesOutput{}
	fnCacher := func(out *cloudformation.ListStackResourcesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackResourcesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackResourcesPagesWithContext(ctx context.Context, input *cloudformation.ListStackResourcesInput, fn func(*cloudformation.ListStackResourcesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackResourcesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackResourcesOutput{}
	fnCacher := func(out *cloudformation.ListStackResourcesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackResourcesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackResourcesWithContext(ctx context.Context, input *cloudformation.ListStackResourcesInput, opts ...request.Option) (*cloudformation.ListStackResourcesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackResourcesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackResourcesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSetOperationResults(input *cloudformation.ListStackSetOperationResultsInput) (*cloudformation.ListStackSetOperationResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetOperationResultsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSetOperationResults(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSetOperationResultsPages(input *cloudformation.ListStackSetOperationResultsInput, fn func(*cloudformation.ListStackSetOperationResultsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetOperationResultsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetOperationResultsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetOperationResultsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetOperationResultsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetOperationResultsPagesWithContext(ctx context.Context, input *cloudformation.ListStackSetOperationResultsInput, fn func(*cloudformation.ListStackSetOperationResultsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetOperationResultsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetOperationResultsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetOperationResultsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetOperationResultsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetOperationResultsWithContext(ctx context.Context, input *cloudformation.ListStackSetOperationResultsInput, opts ...request.Option) (*cloudformation.ListStackSetOperationResultsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetOperationResultsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSetOperationResultsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSetOperations(input *cloudformation.ListStackSetOperationsInput) (*cloudformation.ListStackSetOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetOperationsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSetOperations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSetOperationsPages(input *cloudformation.ListStackSetOperationsInput, fn func(*cloudformation.ListStackSetOperationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetOperationsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetOperationsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetOperationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetOperationsPagesWithContext(ctx context.Context, input *cloudformation.ListStackSetOperationsInput, fn func(*cloudformation.ListStackSetOperationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetOperationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetOperationsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetOperationsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetOperationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetOperationsWithContext(ctx context.Context, input *cloudformation.ListStackSetOperationsInput, opts ...request.Option) (*cloudformation.ListStackSetOperationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetOperationsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSetOperationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSets(input *cloudformation.ListStackSetsInput) (*cloudformation.ListStackSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSets(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStackSetsPages(input *cloudformation.ListStackSetsInput, fn func(*cloudformation.ListStackSetsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetsPagesWithContext(ctx context.Context, input *cloudformation.ListStackSetsInput, fn func(*cloudformation.ListStackSetsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStackSetsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStackSetsOutput{}
	fnCacher := func(out *cloudformation.ListStackSetsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStackSetsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStackSetsWithContext(ctx context.Context, input *cloudformation.ListStackSetsInput, opts ...request.Option) (*cloudformation.ListStackSetsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStackSetsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStackSetsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStacks(input *cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStacksOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStacks(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListStacksPages(input *cloudformation.ListStacksInput, fn func(*cloudformation.ListStacksOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStacksOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStacksOutput{}
	fnCacher := func(out *cloudformation.ListStacksOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStacksPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStacksPagesWithContext(ctx context.Context, input *cloudformation.ListStacksInput, fn func(*cloudformation.ListStacksOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListStacksOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListStacksOutput{}
	fnCacher := func(out *cloudformation.ListStacksOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListStacksPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListStacksWithContext(ctx context.Context, input *cloudformation.ListStacksInput, opts ...request.Option) (*cloudformation.ListStacksOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListStacksOutput), nil
	}
	out, err := c.CloudFormationAPI.ListStacksWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypeRegistrations(input *cloudformation.ListTypeRegistrationsInput) (*cloudformation.ListTypeRegistrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypeRegistrationsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypeRegistrations(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypeRegistrationsPages(input *cloudformation.ListTypeRegistrationsInput, fn func(*cloudformation.ListTypeRegistrationsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypeRegistrationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypeRegistrationsOutput{}
	fnCacher := func(out *cloudformation.ListTypeRegistrationsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypeRegistrationsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypeRegistrationsPagesWithContext(ctx context.Context, input *cloudformation.ListTypeRegistrationsInput, fn func(*cloudformation.ListTypeRegistrationsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypeRegistrationsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypeRegistrationsOutput{}
	fnCacher := func(out *cloudformation.ListTypeRegistrationsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypeRegistrationsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypeRegistrationsWithContext(ctx context.Context, input *cloudformation.ListTypeRegistrationsInput, opts ...request.Option) (*cloudformation.ListTypeRegistrationsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypeRegistrationsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypeRegistrationsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypeVersions(input *cloudformation.ListTypeVersionsInput) (*cloudformation.ListTypeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypeVersionsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypeVersions(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypeVersionsPages(input *cloudformation.ListTypeVersionsInput, fn func(*cloudformation.ListTypeVersionsOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypeVersionsOutput{}
	fnCacher := func(out *cloudformation.ListTypeVersionsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypeVersionsPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypeVersionsPagesWithContext(ctx context.Context, input *cloudformation.ListTypeVersionsInput, fn func(*cloudformation.ListTypeVersionsOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypeVersionsOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypeVersionsOutput{}
	fnCacher := func(out *cloudformation.ListTypeVersionsOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypeVersionsPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypeVersionsWithContext(ctx context.Context, input *cloudformation.ListTypeVersionsInput, opts ...request.Option) (*cloudformation.ListTypeVersionsOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypeVersionsOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypeVersionsWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypes(input *cloudformation.ListTypesInput) (*cloudformation.ListTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypes(input)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
func (c *CloudFormation) ListTypesPages(input *cloudformation.ListTypesInput, fn func(*cloudformation.ListTypesOutput, bool) bool) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypesOutput{}
	fnCacher := func(out *cloudformation.ListTypesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypesPages(input, fnCacher); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypesPagesWithContext(ctx context.Context, input *cloudformation.ListTypesInput, fn func(*cloudformation.ListTypesOutput, bool) bool, opts ...request.Option) error {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		fn(cachedOutput.(*cloudformation.ListTypesOutput), false)
		return nil
	}
	cachable := true
	output := &cloudformation.ListTypesOutput{}
	fnCacher := func(out *cloudformation.ListTypesOutput, more bool) bool {
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
	if err := c.CloudFormationAPI.ListTypesPagesWithContext(ctx, input, fnCacher, opts...); err != nil {
		return err
	}
	if cachable {
		c.cache.SetDefault(strconv.FormatUint(hash, 16), output)
	}
	return nil
}
func (c *CloudFormation) ListTypesWithContext(ctx context.Context, input *cloudformation.ListTypesInput, opts ...request.Option) (*cloudformation.ListTypesOutput, error) {
	hash, _ := hashstructure.Hash(input, hashstructure.FormatV2, &hashstructure.HashOptions{SlicesAsSets: true})
	if cachedOutput, ok := c.cache.Get(strconv.FormatUint(hash, 16)); ok {
		return cachedOutput.(*cloudformation.ListTypesOutput), nil
	}
	out, err := c.CloudFormationAPI.ListTypesWithContext(ctx, input, opts...)
	if err != nil {
		return nil, err
	}
	c.cache.SetDefault(strconv.FormatUint(hash, 16), out)
	return out, err
}
